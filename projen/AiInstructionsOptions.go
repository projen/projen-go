package projen


// Options for configuring AI tool instruction files.
// Experimental.
type AiInstructionsOptions struct {
	// Which AI agents to generate instruction files for.
	// Default: - All agents: [AiAgent.GITHUB_COPILOT, AiAgent.CURSOR, AiAgent.CLAUDE, AiAgent.AMAZON_Q, AiAgent.KIRO, AiAgent.CODEX]
	//
	// Experimental.
	Agents *[]AiAgent `field:"optional" json:"agents" yaml:"agents"`
	// Per-agent custom instructions.
	//
	// Allows different instructions for different AI tools.
	//
	// Example:
	//   {
	//     [AiAgent.GITHUB_COPILOT]: {
	//       instructions: ["Use descriptive commit messages."]
	//     },
	//     [AiAgent.CURSOR]: {
	//       instructions: ["Prefer functional patterns.", "Always add tests."]
	//     }
	//   }
	//
	// Default: - no agent specific instructions.
	//
	// Experimental.
	AgentSpecificInstructions *map[string]*[]*string `field:"optional" json:"agentSpecificInstructions" yaml:"agentSpecificInstructions"`
	// Include default instructions for projen and general best practices.
	//
	// Default instructions will only be included for agents provided in the `agents` option.
	// If `agents` is not provided, default instructions will be included for all agents.
	// Default: true.
	//
	// Experimental.
	IncludeDefaultInstructions *bool `field:"optional" json:"includeDefaultInstructions" yaml:"includeDefaultInstructions"`
	// General instructions applicable to all agents.
	// Default: - no agent specific instructions.
	//
	// Experimental.
	Instructions *[]*string `field:"optional" json:"instructions" yaml:"instructions"`
}

