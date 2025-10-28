package projen


// Supported AI coding assistants and their instruction file locations.
// Experimental.
type AiAgent string

const (
	// GitHub Copilot - .github/copilot-instructions.md.
	// Experimental.
	AiAgent_GITHUB_COPILOT AiAgent = "GITHUB_COPILOT"
	// Cursor IDE - .cursor/rules/project.md.
	// Experimental.
	AiAgent_CURSOR AiAgent = "CURSOR"
	// Claude Code - CLAUDE.md.
	// Experimental.
	AiAgent_CLAUDE AiAgent = "CLAUDE"
	// Amazon Q - .amazonq/rules/project.md.
	// Experimental.
	AiAgent_AMAZON_Q AiAgent = "AMAZON_Q"
	// Kiro - .kiro/steering/project.md.
	// Experimental.
	AiAgent_KIRO AiAgent = "KIRO"
)

