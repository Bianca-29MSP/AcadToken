package types

const (
	// ModuleName define o nome do módulo
	ModuleName = "curriculum"

	// StoreKey define a chave principal de armazenamento do módulo
	StoreKey = ModuleName

	// MemStoreKey define a chave de armazenamento em memória
	MemStoreKey = "mem_curriculum"

	// CourseTokenByCodeAndInstitutionKey define o prefixo para o índice secundário de disciplinas por código e instituição
	CourseTokenByCodeAndInstitutionKey = "CourseTokenByCodeAndInst"
)

var (
	ParamsKey = []byte("p_curriculum")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
