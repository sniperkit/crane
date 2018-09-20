package parser

/*
   - Identify flag wrappers:
     - flag
      - wrappers
        - github.com/spf13/cobra
        - github.com/spf13/viper
        - github.com/alecthomas/kingpin
        - github.com/urfave/cli (be careful with version changes)
        - github.com/codegangsta/cli
        - github.com/jessevdk/go-flag
        - github.com/palantir/pkg/cobracli
*/

/*
func renameFlag(fileNode *ast.File, originalName, newName string) error {
    originalFunc := findFunction(fileNode, originalName)
    if originalFunc == nil {
        return errors.Errorf("function %s does not exist", originalName)
    }

    if findFunction(fileNode, newName) != nil {
        return errors.Errorf("cannot rename function %s to %s because a function with the new name already exists", originalName, newName)
    }

    originalFunc.Name = ast.NewIdent(newName)
    return nil
}

func findFlag(fileNode *ast.File, funcName string) *ast.FuncDecl {
    for _, currDecl := range fileNode.Decls {
        switch t := currDecl.(type) {
        case *ast.FuncDecl:
            if t.Name.Name == funcName {
                return currDecl.(*ast.FuncDecl)
            }
        }
    }
    return nil
}


*/
