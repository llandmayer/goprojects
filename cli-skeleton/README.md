### **1.	Project Overview**
	Brief description of the CLI application.
	Key features and purpose.
### 2.	Installation:
	Steps to install the CLI (e.g., go install instructions or binary downloads).
###	3.	Usage:
	Basic usage examples.
	Command structure (e.g., command [flags]).
###	4.	Quick Start:
	A minimal example to get the user started.
###	5.	Contributing:
	Guidelines for developers to contribute.
###	6.	Links:
	Links to additional documentation (e.g., docs folder).

### **2. `/docs` Folder Structure**

The `docs` folder organizes in-depth information about the project. Here's a suggested structure:

#### **a. High-Level Documentation**

**`overview.md`**:
  - Provide an overview of the architecture and purpose of the project.

**`getting-started.md`**:
  - Include detailed installation and setup instructions.
  - Basic usage walkthroughs. (AppState and Modes)

**`cli-commands.md`**:
  - Document each command, its flags, and examples.
  - Use tables for clarity.

**`pluings.md`**:
  - Document how plugins work.
  - How to implement a new plugin.


### **3. Plugin-Specific Documentation**

For each plugin, create subfolders or individual files:

- **`plugins/terragrunt/README.md`**:
  - Explain the purpose of the plugin.
  - Provide usage examples.
  - Document plugin-specific configurations.

**Example:**
```markdown
# Terragrunt Plugin

The `terragrunt` plugin provides support for managing infrastructure as code.

## Usage
```bash
cli-skeleton run terragrunt

Modes
	•	Main Mode:
Default entry point.
	•	Sub Modes:
	•	sub1: …
	•	sub2: …
```