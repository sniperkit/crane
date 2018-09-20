package parser

/*
func renameConstant(fileNode *ast.File, originalName, newName string) error {
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

func findConstant(fileNode *ast.File, funcName string) *ast.FuncDecl {
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
