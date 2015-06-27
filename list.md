should fix:

unused_vars
	./src/cmd/compile/internal/gc/y.go:1203:
		_ = yypt // guard against "declared and not used"
	./src/cmd/compile/internal/gc/walk.go:55:
		Yyerror("%v declared and not used", l.N.Sym)
	./src/cmd/compile/internal/gc/walk.go:61:
		Yyerror("%v declared and not used", l.N.Sym)

unused_packages:
	./src/cmd/compile/internal/gc/subr.go:346:
		yyerrorl(int(pack.Lineno), "imported and not used: %q", opkg.Path)
	./src/cmd/compile/internal/gc/lex.go:2583:
		yyerrorl(int(lineno), "imported and not used: %q", path)
	./src/cmd/compile/internal/gc/lex.go:2585:
		yyerrorl(int(lineno), "imported and not used: %q as %s", path, name)


