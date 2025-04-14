package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run gen.go EntityName")
		return
	}

	entity := os.Args[1]
	entityLower := strings.ToLower(entity)

	createProtoFile("api/proto/v1/"+entityLower+".proto", "proto.tpl", entity, entityLower)
	createFile("internal/handler/"+entityLower+"_handler.go", "handler.tpl", entity, entityLower)
	createFile("internal/domain/"+entityLower+"/service.go", "domain.tpl", entity, entityLower)
	createTestFile("internal/domain/"+entityLower+"/service_test.go", "domain_test.tpl", entity, entityLower)

	updateGRPCServer("internal/protocol/grpc/server.go", entity, false)
	updateGRPCServer("internal/protocol/grpc-gateway/server.go", entity, true)
	updateDIContainer(entity, entityLower)

	fmt.Println("✅ Generated handler and domain for:", entity)
}

// --- FILE CREATION ---

func createProtoFile(path, tplFile, entity, entityLower string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("⚠️  Proto file already exists:", path)
		return
	}
	createFile(path, tplFile, entity, entityLower)
}

func createFile(path, tplFile, entity, entityLower string) {
	tplPath := "gen/templates/" + tplFile
	content, err := os.ReadFile(tplPath)
	if err != nil {
		panic(err)
	}

	os.MkdirAll(getDir(path), os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	t := template.Must(template.New("tpl").Parse(string(content)))
	t.Execute(f, map[string]string{
		"Entity":      entity,
		"EntityLower": entityLower,
	})
}

func createTestFile(path, tplFile, entity, entityLower string) {
	if _, err := os.Stat(path); err == nil {
		fmt.Println("⚠️  Test file already exists:", path)
		return
	}
	createFile(path, tplFile, entity, entityLower)
}

func getDir(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		return ""
	}
	return path[:idx]
}

// --- INJECT SERVER HANDLER ---

func updateGRPCServer(path string, entity string, isGateway bool) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("⚠️  %s not found. Please update manually.\n", path)
		return
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string
	alreadyInjected := false

	for _, line := range lines {
		// Inject handler
		if strings.Contains(line, "// @auto:inject:handler") && !alreadyInjected {
			if isGateway {
				newLines = append(newLines,
					fmt.Sprintf(`	if err := v1pb.Register%sHandlerFromEndpoint(ctx, mux, grpcAddr, dialOpts); err != nil {`, entity),
					fmt.Sprintf(`		return fmt.Errorf("failed to register %s handler: %%w", err)`, strings.ToLower(entity)),
					`	}`,
				)
			} else {
				newLines = append(newLines,
					fmt.Sprintf(`	v1pb.Register%sServiceServer(grpcServer, container.%sHandler)`, entity, entity),
				)
			}
			alreadyInjected = true
		}
		newLines = append(newLines, line)
	}

	err = os.WriteFile(path, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Printf("⚠️  Failed to update %s: %v\n", path, err)
	} else {
		fmt.Printf("🧩 Updated %s with %s handler\n", path, entity)
	}
}

// --- INJECT DI CONTAINER ---

func updateDIContainer(entity, entityLower string) {
	path := "internal/di/container.go"

	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("⚠️  container.go not found. Please update manually.")
		return
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string
	hasImported := false

	for _, line := range lines {
		if strings.HasPrefix(line, "import (") && !hasImported {
			newLines = append(newLines, line)
			newLines = append(newLines, fmt.Sprintf(`	"gitlab.twprisma.com/fin/lmd/services/if-trx-history/internal/domain/%s"`, entityLower))
			hasImported = true
			continue
		}

		// Inject field di struct Container
		if strings.Contains(line, "// @auto:inject:field") {
			newLines = append(newLines, fmt.Sprintf("	%sHandler *handler.%sHandler", entity, entity))
			newLines = append(newLines, line)
			continue
		}

		// Inject service init
		if strings.Contains(line, "// @auto:inject:init-service") {
			newLines = append(newLines, fmt.Sprintf("	%sService := %s.New%sService()", entityLower, entityLower, entity))
			newLines = append(newLines, line)
			continue
		}

		// Inject handler assignment
		if strings.Contains(line, "// @auto:inject:init-handler") {
			newLines = append(newLines, fmt.Sprintf("		%sHandler: handler.New%sHandler(%sService),", entity, entity, entityLower))
			newLines = append(newLines, line)
			continue
		}

		newLines = append(newLines, line)
	}

	err = os.WriteFile(path, []byte(strings.Join(newLines, "\n")), 0644)
	if err != nil {
		fmt.Println("⚠️  Failed to update container.go:", err)
	} else {
		fmt.Println("🧩 Updated container.go with", entity, "Handler")
	}
}
