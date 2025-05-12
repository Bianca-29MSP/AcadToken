package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/Bianca-29MSP/AcademicToken/x/curriculum/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/jdkato/prose/v2"
	"github.com/spf13/cobra"
	"gopkg.in/neurosnap/sentences.v1/english"
)

// CourseContentProcessor processa ementas acadêmicas
type CourseContentProcessor struct {
	client client.Context
}

// NewCourseContentProcessor cria um novo processador
func NewCourseContentProcessor(clientCtx client.Context) *CourseContentProcessor {
	return &CourseContentProcessor{
		client: clientCtx,
	}
}

// ProcessContent processa uma ementa e gera um CourseContent
func (p *CourseContentProcessor) ProcessContent(rawContent, institution, courseCode string) (string, error) {
	// Extrai metadados básicos da ementa
	metadata, err := p.extractMetadata(rawContent)
	if err != nil {
		return "", fmt.Errorf("erro ao extrair metadados: %w", err)
	}

	// Extrai tópicos usando NLP
	topics, err := p.extractTopics(rawContent)
	if err != nil {
		return "", fmt.Errorf("erro ao extrair tópicos: %w", err)
	}

	// Extrai palavras-chave
	keywords, err := p.extractKeywords(rawContent)
	if err != nil {
		return "", fmt.Errorf("erro ao extrair palavras-chave: %w", err)
	}

	// Normaliza informações de bibliografia
	basicBibliography := p.normalizeBibliography(metadata.BasicBibliography)
	compBibliography := p.normalizeBibliography(metadata.ComplementaryBibliography)

	// Gera hash do conteúdo original
	contentHash := p.generateContentHash(rawContent)

	// Prepara os dados para o CourseContent
	// Convertemos as arrays para JSON strings para armazenamento
	topicUnitsJSON, _ := json.Marshal(topics)

	// Cria o objeto CourseContent com o campo Index
	content := types.CourseContent{
		Index:                    fmt.Sprintf("%s-%s", institution, courseCode), // Gera um índice baseado na instituição e código
		CourseId:                 courseCode,
		Institution:              institution,
		Title:                    metadata.Title,
		Code:                     courseCode,
		WorkloadHours:            metadata.WorkloadHours,
		Credits:                  metadata.Credits,
		Description:              metadata.Description,
		Objectives:               []string(metadata.Objectives),
		TopicUnits:               []string{string(topicUnitsJSON)},
		Methodologies:            []string(metadata.Methodologies),
		EvaluationMethods:        []string(metadata.EvaluationMethods),
		BibliographyBasic:        basicBibliography,
		BibliographyComplementary: compBibliography,
		Keywords:                 keywords,
		ContentHash:              contentHash,
	}

	// Por enquanto, vamos salvar os dados como JSON para facilitar a verificação
	jsonFile := fmt.Sprintf("%s_%s.json", institution, courseCode)
	jsonData, _ := json.MarshalIndent(content, "", "  ")
	ioutil.WriteFile(jsonFile, jsonData, 0644)
	fmt.Printf("Dados salvos em %s para referência\n", jsonFile)

	// Em vez de tentar enviar para a blockchain aqui, iremos apenas indicar 
	// que o processamento foi concluído. O envio será feito na função main.
	fmt.Println("Processamento local concluído. O envio para a blockchain será feito via CLI.")

	return contentHash, nil
}

// CourseMetadata representa os metadados extraídos de uma ementa
type CourseMetadata struct {
	Title                     string
	WorkloadHours             uint64
	Credits                   uint64
	Description               string
	Objectives                []string
	Methodologies             []string
	EvaluationMethods         []string
	BasicBibliography         []string
	ComplementaryBibliography []string
}

// extractMetadata extrai metadados básicos da ementa
func (p *CourseContentProcessor) extractMetadata(rawContent string) (CourseMetadata, error) {
	metadata := CourseMetadata{}

	// Identifica as seções da ementa
	sections := p.identifySections(rawContent)

	// Extrai o título (geralmente é a primeira linha ou está destacado)
	metadata.Title = p.extractTitle(rawContent)

	// Extrai carga horária
	metadata.WorkloadHours = p.extractWorkloadHours(rawContent)

	// Extrai créditos
	metadata.Credits = p.extractCredits(rawContent)

	// Extrai descrição/ementa
	if section, ok := sections["EMENTA"]; ok {
		metadata.Description = strings.TrimSpace(section)
	}

	// Extrai objetivos
	if section, ok := sections["OBJETIVOS"]; ok {
		metadata.Objectives = p.extractListItems(section)
	}

	// Extrai metodologia
	if section, ok := sections["METODOLOGIA"]; ok {
		metadata.Methodologies = p.extractListItems(section)
	}

	// Extrai métodos de avaliação
	if section, ok := sections["AVALIAÇÃO"]; ok {
		metadata.EvaluationMethods = p.extractListItems(section)
	}

	// Extrai bibliografia básica
	if section, ok := sections["BIBLIOGRAFIA BÁSICA"]; ok {
		metadata.BasicBibliography = p.extractListItems(section)
	}

	// Extrai bibliografia complementar
	if section, ok := sections["BIBLIOGRAFIA COMPLEMENTAR"]; ok {
		metadata.ComplementaryBibliography = p.extractListItems(section)
	}

	return metadata, nil
}

// identifySections identifica as diferentes seções da ementa
func (p *CourseContentProcessor) identifySections(rawContent string) map[string]string {
	sections := make(map[string]string)

	// Versão simplificada que não usa lookahead (?=)
	// Divide o texto em linhas
	lines := strings.Split(rawContent, "\n")
	var currentSection string
	var currentContent []string

	// Expressão para identificar cabeçalhos de seção
	headerRegex := regexp.MustCompile(`^([A-ZÁÉÍÓÚÇÃÕÊÔ\s]+)[:.-]\s*$`)

	for i, line := range lines {
		// Verifica se a linha é um cabeçalho de seção
		matches := headerRegex.FindStringSubmatch(line)
		if len(matches) > 1 {
			// Se já estávamos em uma seção, salve o conteúdo anterior
			if currentSection != "" && len(currentContent) > 0 {
				sections[currentSection] = strings.Join(currentContent, "\n")
			}

			// Inicia nova seção
			currentSection = strings.TrimSpace(matches[1])
			currentContent = []string{}
		} else if currentSection != "" {
			// Adiciona esta linha ao conteúdo da seção atual
			currentContent = append(currentContent, line)
		}

		// Se for a última linha e estivermos em uma seção, salve o conteúdo
		if i == len(lines)-1 && currentSection != "" && len(currentContent) > 0 {
			sections[currentSection] = strings.Join(currentContent, "\n")
		}
	}

	return sections
}

// extractTopics extrai tópicos do conteúdo usando NLP
func (p *CourseContentProcessor) extractTopics(rawContent string) ([]string, error) {
	// Usamos um analisador de sentenças para dividir o texto
	tokenizer, err := english.NewSentenceTokenizer(nil)
	if err != nil {
		return nil, err
	}

	sentences := tokenizer.Tokenize(rawContent)

	// Procuramos por indicadores de tópicos de curso:
	// 1. Numeração sequencial (1., 1.1, etc.)
	// 2. Marcadores (•, -, *)
	topicRegex := regexp.MustCompile(`(?m)^(?:\d+\.[\d\.]*\s+|\-\s+|•\s+|\*\s+)(.+)$`)

	var topics []string

	for _, s := range sentences {
		text := s.Text
		matches := topicRegex.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) > 1 {
				topic := strings.TrimSpace(match[1])
				topics = append(topics, topic)
			}
		}
	}

	// Se não encontramos tópicos estruturados, tentamos extrair frases nominais
	if len(topics) < 3 {
		doc, err := prose.NewDocument(rawContent)
		if err != nil {
			return nil, err
		}

		var extractedPhrases []string
		for _, token := range doc.Tokens() {
			// Extrair substantivos e substantivos próprios
			if strings.HasPrefix(token.Tag, "NN") {
				extractedPhrases = append(extractedPhrases, token.Text)
			}
		}

		// Adicionar as frases nominais encontradas
		topics = append(topics, extractedPhrases...)
	}

	return topics, nil
}

// extractKeywords extrai palavras-chave do conteúdo
func (p *CourseContentProcessor) extractKeywords(rawContent string) ([]string, error) {
	// Usamos a biblioteca prose para análise de texto em português/inglês
	doc, err := prose.NewDocument(rawContent)
	if err != nil {
		return nil, err
	}

	// Extrai entidades e substantivos como possíveis palavras-chave
	wordFreq := make(map[string]int)

	// Adiciona entidades
	for _, ent := range doc.Entities() {
		wordFreq[strings.ToLower(ent.Text)]++
	}

	// Adiciona substantivos
	for _, tok := range doc.Tokens() {
		if strings.HasPrefix(tok.Tag, "NN") {
			wordFreq[strings.ToLower(tok.Text)]++
		}
	}

	// Filtra palavras muito curtas ou comuns
	var keywords []string
	for word, freq := range wordFreq {
		if len(word) > 3 && freq >= 2 {
			keywords = append(keywords, word)
		}
	}

	return keywords, nil
}

// normalizeBibliography normaliza itens de bibliografia
func (p *CourseContentProcessor) normalizeBibliography(items []string) []string {
	var normalized []string

	for _, item := range items {
		// Remove numeração e espaços extras
		item = regexp.MustCompile(`^\d+\.\s*`).ReplaceAllString(item, "")
		item = strings.TrimSpace(item)

		// Identifica e padroniza autores e título
		authors := p.extractAuthors(item)
		title := p.extractTitle(item)

		// Monta item padronizado (se possível)
		if len(authors) > 0 && title != "" {
			normalized = append(normalized, item)
		} else {
			// Se não conseguimos estruturar, mantemos o original
			normalized = append(normalized, item)
		}
	}

	return normalized
}

// extractAuthors extrai nomes de autores de um texto
func (p *CourseContentProcessor) extractAuthors(text string) []string {
	// Implementação simples para identificar padrões de nomes
	// Busca por padrões: "Sobrenome, Iniciais" ou "Nome Sobrenome"
	authorRegex := regexp.MustCompile(`([A-Z][A-Za-záéíóúçãõâêôàèìòù]+,\s+[A-Z]\.(?:\s+[A-Z]\.)*)|([A-Z][a-záéíóúçãõâêôàèìòù]+\s+[A-Z][a-záéíóúçãõâêôàèìòù]+)`)
	matches := authorRegex.FindAllString(text, -1)
	return matches
}

// extractYears extrai anos de um texto
func (p *CourseContentProcessor) extractYears(text string) []string {
	// Busca por anos no formato YYYY
	yearRegex := regexp.MustCompile(`\b(19\d{2}|20\d{2})\b`)
	matches := yearRegex.FindAllString(text, -1)
	return matches
}

// extractListItems extrai itens de lista de um texto
func (p *CourseContentProcessor) extractListItems(text string) []string {
	var items []string

	// Identifica itens marcados por números, letras ou símbolos
	itemRegex := regexp.MustCompile(`(?m)^(?:\d+\.[\d\.]*\s+|\-\s+|•\s+|\*\s+|[a-z]\)\s+)(.+)$`)
	matches := itemRegex.FindAllStringSubmatch(text, -1)

	for _, match := range matches {
		if len(match) > 1 {
			items = append(items, strings.TrimSpace(match[1]))
		}
	}

	// Se não encontrou itens marcados, divide por linhas
	if len(items) == 0 {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				items = append(items, line)
			}
		}
	}

	return items
}

// extractTitle tenta extrair o título do documento
func (p *CourseContentProcessor) extractTitle(text string) string {
	// Procura por um título típico de disciplina
	titleRegex := regexp.MustCompile(`(?i)(?:DISCIPLINA|CURSO):\s*(.+)`)
	match := titleRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		return strings.TrimSpace(match[1])
	}

	// Tenta a primeira linha não vazia
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			return line
		}
	}

	return ""
}

// extractWorkloadHours extrai a carga horária
func (p *CourseContentProcessor) extractWorkloadHours(text string) uint64 {
	// Procura por padrões comuns de carga horária
	hoursRegex := regexp.MustCompile(`(?i)(?:CARGA HORÁRIA|CH):\s*(\d+)`)
	match := hoursRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		var hours uint64
		fmt.Sscanf(match[1], "%d", &hours)
		return hours
	}

	return 0
}

// extractCredits extrai o número de créditos
func (p *CourseContentProcessor) extractCredits(text string) uint64 {
	// Procura por padrões comuns de créditos
	creditsRegex := regexp.MustCompile(`(?i)(?:CRÉDITOS|CR):\s*(\d+)`)
	match := creditsRegex.FindStringSubmatch(text)
	if len(match) > 1 {
		var credits uint64
		fmt.Sscanf(match[1], "%d", &credits)
		return credits
	}

	return 0
}

// generateContentHash gera um hash SHA-256 do conteúdo
func (p *CourseContentProcessor) generateContentHash(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

// extractTextFromPDF extrai texto de um arquivo PDF usando a ferramenta pdftotext
func extractTextFromPDF(filePath string) (string, error) {
	// Verifica se o executável pdftotext está disponível
	_, err := exec.LookPath("pdftotext")
	if err != nil {
		return "", fmt.Errorf("pdftotext não encontrado. Instale o pacote poppler-utils: %w", err)
	}

	// Executa o comando pdftotext para extrair o texto
	cmd := exec.Command("pdftotext", filePath, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("erro ao executar pdftotext: %w", err)
	}

	return string(output), nil
}

// tokenizeViaCLI envia a transação para a blockchain via CLI
func tokenizeViaCLI(content types.CourseContent, from, chainID, node string) error {
    // Caminho exato do binário AcademicTokend no seu sistema
    academicTokendPath := "/Users/biancamsp/go/bin/AcademicTokend"
    
    // Verificar se o arquivo existe
    if _, err := os.Stat(academicTokendPath); os.IsNotExist(err) {
        return fmt.Errorf("binário AcademicTokend não encontrado em %s", academicTokendPath)
    }
    
    fmt.Printf("Usando binário AcademicTokend em: %s\n", academicTokendPath)
    
    // Criar um arquivo temporário para armazenar os dados completos da ementa como JSON
    ementaFile, err := ioutil.TempFile("", "ementa-*.json")
    if err != nil {
        return fmt.Errorf("erro ao criar arquivo temporário para ementa: %w", err)
    }
    defer os.Remove(ementaFile.Name())
    
    // Serializar toda a estrutura para JSON
    ementaJSON, err := json.MarshalIndent(content, "", "  ")
    if err != nil {
        return fmt.Errorf("erro ao serializar ementa para JSON: %w", err)
    }
    
    // Escrever o JSON no arquivo temporário
    if _, err = ementaFile.Write(ementaJSON); err != nil {
        return fmt.Errorf("erro ao escrever arquivo JSON: %w", err)
    }
    if err = ementaFile.Close(); err != nil {
        return fmt.Errorf("erro ao fechar arquivo JSON: %w", err)
    }
    
    fmt.Printf("Dados completos da ementa salvos em: %s\n", ementaFile.Name())
    
    // Simplificar a chamada da CLI para evitar problemas de parsing
    // Juntar os arrays em strings separadas por vírgulas para facilitar
    objectives := strings.Join(content.Objectives, ",")
    topicUnits := strings.Join(content.TopicUnits, ",")
    methodologies := strings.Join(content.Methodologies, ",")
    evaluationMethods := strings.Join(content.EvaluationMethods, ",")
    bibliographyBasic := strings.Join(content.BibliographyBasic, ",")
    bibliographyComplementary := strings.Join(content.BibliographyComplementary, ",")
    keywords := strings.Join(content.Keywords, ",")
    
    // Criar um script shell que usa o formato correto para o comando
    scriptContent := fmt.Sprintf(`#!/bin/bash
echo "==================================================================="
echo "AVISO: Dados completos da ementa estão no arquivo: %s"
echo "Este arquivo contém todos os metadados, objetivos, tópicos, etc."
echo "Use-o para verificação completa da ementa após a tokenização."
echo "==================================================================="

%s tx curriculum create-course-content \
  %s \
  %s \
  %s \
  "%s" \
  %s \
  %d \
  %d \
  "%s" \
  "%s" \
  "%s" \
  "%s" \
  "%s" \
  "%s" \
  "%s" \
  "%s" \
  %s \
  --from %s \
  --chain-id %s \
  --node %s \
  --gas auto \
  --gas-adjustment 1.5 \
  --yes

# Exibe resultado
echo ""
echo "==================================================================="
echo "IMPORTANTE: Mantenha o arquivo de ementa para referência completa:"
echo "%s"
echo "==================================================================="
`,
        ementaFile.Name(),
        academicTokendPath,
        content.Index,
        content.CourseId,
        content.Institution,
        content.Title,
        content.Code,
        content.WorkloadHours,
        content.Credits,
        content.Description,
        objectives,        // Substitui a referência ao arquivo por dados reais
        topicUnits,        // Substitui a referência ao arquivo por dados reais
        methodologies,     // Substitui a referência ao arquivo por dados reais
        evaluationMethods, // Substitui a referência ao arquivo por dados reais
        bibliographyBasic, // Substitui a referência ao arquivo por dados reais
        bibliographyComplementary, // Substitui a referência ao arquivo por dados reais
        keywords,          // Substitui a referência ao arquivo por dados reais
        content.ContentHash,
        from,
        chainID,
        node,
        ementaFile.Name())
    
    // Exibir o conteúdo do script para depuração
    fmt.Println("==== SCRIPT GERADO ====")
    fmt.Println(scriptContent)
    fmt.Println("=======================")
    
    // Criar arquivo temporário para o script
    scriptFile, err := ioutil.TempFile("", "create-course-*.sh")
    if err != nil {
        return fmt.Errorf("erro ao criar arquivo temporário para script: %w", err)
    }
    defer os.Remove(scriptFile.Name())
    
    // Escrever o script
    if _, err := scriptFile.Write([]byte(scriptContent)); err != nil {
        return fmt.Errorf("erro ao escrever script: %w", err)
    }
    if err := scriptFile.Close(); err != nil {
        return fmt.Errorf("erro ao fechar arquivo de script: %w", err)
    }
    
    // Tornar o script executável
    if err := os.Chmod(scriptFile.Name(), 0755); err != nil {
        return fmt.Errorf("erro ao tornar script executável: %w", err)
    }
    
    fmt.Printf("Script de execução criado em: %s\n", scriptFile.Name())
    
    // Criar uma versão permanente do arquivo JSON para referência
    permanentFile := fmt.Sprintf("%s_%s_completo.json", content.Institution, content.CourseId)
    if err := ioutil.WriteFile(permanentFile, ementaJSON, 0644); err != nil {
        fmt.Printf("Aviso: Não foi possível criar arquivo permanente: %v\n", err)
    } else {
        fmt.Printf("Arquivo permanente com dados completos criado: %s\n", permanentFile)
    }
    
    // Executar o script
    cmd := exec.Command("bash", scriptFile.Name())
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    return cmd.Run()
}

// initClientContext inicializa o contexto do cliente Cosmos de forma simplificada
func initClientContext() (client.Context, error) {
    // Criar um contexto básico vazio
    clientCtx := client.Context{}
    
    // Configurar um comando Cobra
    cmd := &cobra.Command{
        Use: "course-processor",
        Run: func(cmd *cobra.Command, args []string) {},
    }
    
    // Adicionar flags padrão do Cosmos SDK
    flags.AddTxFlagsToCmd(cmd)
    
    // Registrar flags essenciais
    cmd.PersistentFlags().String(flags.FlagChainID, "AcademicToken", "Chain ID of tendermint node")
    cmd.PersistentFlags().String(flags.FlagHome, os.ExpandEnv("$HOME/.academictoken"), "Directory for config and data")
    cmd.PersistentFlags().String(flags.FlagKeyringBackend, "test", "Select keyring's backend")
    cmd.PersistentFlags().String(flags.FlagNode, "tcp://localhost:26657", "RPC server")
    cmd.PersistentFlags().String(flags.FlagFrom, "", "Name or address of private key with which to sign")
    
    // Parse flags para simular a CLI
    err := cmd.ParseFlags(os.Args[1:])
    if err != nil {
        return clientCtx, err
    }
    
    // Configurar o home directory
    homeValue := cmd.Flag(flags.FlagHome).Value.String()
    if homeValue != "" {
        clientCtx = clientCtx.WithHomeDir(homeValue)
    }
    
    // Configurar Chain ID
    chainIDFlag := cmd.Flag(flags.FlagChainID)
    if chainIDFlag != nil && chainIDFlag.Value.String() != "" {
        clientCtx = clientCtx.WithChainID(chainIDFlag.Value.String())
    }
    
    // Configurar o nó RPC
    nodeFlag := cmd.Flag(flags.FlagNode)
    if nodeFlag != nil && nodeFlag.Value.String() != "" {
        clientCtx = clientCtx.WithNodeURI(nodeFlag.Value.String())
    }
    
    // Usar apenas o AccountRetriever básico
    clientCtx = clientCtx.WithAccountRetriever(authtypes.AccountRetriever{})
    
    return clientCtx, nil
}

// Função principal para executar o processador
func main() {
    // Verificar argumentos
    if len(os.Args) < 4 {
        fmt.Println("Uso: course_processor <arquivo_ementa.pdf|txt> <instituição> <código_curso> [--chain-id <chain-id>] [--from <key-name>]")
        fmt.Println("Exemplo: ./course_processor ementa.pdf UFBA CS101 --chain-id AcademicToken --from alice")
        os.Exit(1)
    }
    
    // Obter argumentos
    filePath := os.Args[1]
    institution := os.Args[2]
    courseCode := os.Args[3]
    
    var rawContent string
    var err error
    
    // Verificar se é um arquivo PDF
    if strings.HasSuffix(strings.ToLower(filePath), ".pdf") {
        // Extrair texto do PDF
        fmt.Println("Extraindo texto do arquivo PDF...")
        rawContent, err = extractTextFromPDF(filePath)
        if err != nil {
            fmt.Printf("Erro ao extrair texto do PDF: %v\n", err)
            os.Exit(1)
        }
        fmt.Println("Extração de texto concluída.")
    } else {
        // Ler como arquivo de texto
        contentBytes, err := ioutil.ReadFile(filePath)
        if err != nil {
            fmt.Printf("Erro ao ler arquivo: %v\n", err)
            os.Exit(1)
        }
        rawContent = string(contentBytes)
    }
    
    // Extrair flags importantes para o modo CLI
    fromFlag := ""
    chainIDFlag := "AcademicToken"
    nodeFlag := "tcp://localhost:26657"
    
    for i, arg := range os.Args {
        if arg == "--from" && i+1 < len(os.Args) {
            fromFlag = os.Args[i+1]
        }
        if arg == "--chain-id" && i+1 < len(os.Args) {
            chainIDFlag = os.Args[i+1]
        }
        if arg == "--node" && i+1 < len(os.Args) {
            nodeFlag = os.Args[i+1]
        }
    }
    
    // Inicializar o contexto do cliente simplificado
    fmt.Println("Inicializando cliente blockchain...")
    clientCtx, err := initClientContext()
    if err != nil {
        fmt.Printf("Aviso: Erro ao inicializar o contexto do cliente: %v\n", err)
        clientCtx = client.Context{}
    } else {
        fmt.Println("Cliente inicializado com:")
        if clientCtx.ChainID != "" {
            fmt.Printf("- Chain ID: %s\n", clientCtx.ChainID)
        }
        if clientCtx.NodeURI != "" {
            fmt.Printf("- Nó: %s\n", clientCtx.NodeURI)
        }
    }
    
    fmt.Println("Processando ementa...")
    
    // Processar a ementa
    processor := NewCourseContentProcessor(clientCtx)
    contentHash, err := processor.ProcessContent(rawContent, institution, courseCode)
    if err != nil {
        fmt.Printf("Erro ao processar ementa: %v\n", err)
        os.Exit(1)
    }
    
    fmt.Printf("Ementa processada com sucesso! Content Hash: %s\n", contentHash)
    
    // Verificar se devemos tentar enviar via CLI
    if fromFlag != "" {
        fmt.Println("\nEnviando transação via CLI...")
        
        // Criar o objeto CourseContent
        metadata, _ := processor.extractMetadata(rawContent)
        topics, _ := processor.extractTopics(rawContent)
        keywords, _ := processor.extractKeywords(rawContent)
        basicBibliography := processor.normalizeBibliography(metadata.BasicBibliography)
        compBibliography := processor.normalizeBibliography(metadata.ComplementaryBibliography)
        topicUnitsJSON, _ := json.Marshal(topics)
        
        content := types.CourseContent{
            Index:                    fmt.Sprintf("%s-%s", institution, courseCode), // Gera um índice baseado na instituição e código
            CourseId:                 courseCode,
            Institution:              institution,
            Title:                    metadata.Title,
            Code:                     courseCode,
            WorkloadHours:            metadata.WorkloadHours,
            Credits:                  metadata.Credits,
            Description:              metadata.Description,
            Objectives:               []string(metadata.Objectives),
            TopicUnits:               []string{string(topicUnitsJSON)},
            Methodologies:            []string(metadata.Methodologies),
            EvaluationMethods:        []string(metadata.EvaluationMethods),
            BibliographyBasic:        basicBibliography,
            BibliographyComplementary: compBibliography,
            Keywords:                 keywords,
            ContentHash:              contentHash,
        }
        
        err = tokenizeViaCLI(content, fromFlag, chainIDFlag, nodeFlag)
       if err != nil {
           fmt.Printf("Erro ao executar comando via CLI: %v\n", err)
           fmt.Println("Verifique se o AcademicTokend está no PATH e se os parâmetros estão corretos.")
       } else {
           fmt.Println("Transação enviada com sucesso via CLI!")
       }
   } else {
       fmt.Println("\nNenhum endereço fornecido com --from. Execute novamente com --from para enviar a transação.")
   }
}