# CLI Skeleton

The idea behind this project is to create a modular CLI using plugins to ensure each component is self-isolated.

This directory structure is clean, modular, and adheres to Go best practices. Below is a breakdown of its components and strengths:

---

## 1. Root Directory
- Purpose:
  - Contains essential scripts for pipelines, testing, and documentation.
- Key file:
  - **README.md**: Provides an overview of the project.



## 2. `cmd/`

- **Purpose**:
  - Houses the entry points for the CLI commands, such as the root and version commands.
  - Keeps command-related logic isolated from the rest of the application.

- **Files**:
  - **`root.go`**:
    - Defines the root command and initializes the CLI.
  - **`version.go`**:
    - Provides a command to output the application’s version (common for CLI tools).
  - Additional commands (e.g., `help`, `init`) can be easily added in the future.



## 3. `internal/`

- **Purpose**:
  - Contains all application-specific internal logic that shouldn’t be imported by external packages.

### Subdirectories:

### a. `config/`
- Manages configuration files, environment variables, and related utilities.
- **`config.go`**:
  - Defines the configuration structure and loading mechanisms (e.g., from files or environment variables).
- **Why It’s Good**:
  - Encapsulates configuration handling, ensuring a single source of truth for app configuration.

### b. `logging/`
- Provides centralized logging utilities.
- **`logging.go`**:
  - Encapsulates log setup (e.g., with libraries like `logrus` or `zap`).
- **Why It’s Good**:
  - Centralizes logging logic, making it reusable and configurable across the app.

### c. `plugins/`
- Manages plugins through:
  - **`interfaces.go`**:
    - Defines the plugin interface for modularity.
  - **`manager.go`**:
    - Handles plugin registration, lifecycle, and switching.
  - **`registry.go`**:
    - Centralizes plugin registration.
- **Why It’s Good**:
  - Isolates plugin logic in one place, keeping it manageable and extendable.

### d. `tui/`
- Implements the TUI (Text User Interface) logic using Charm’s `bubbletea`.
- **`tui.go`**:
  - Central TUI management logic (e.g., state updates, rendering).
- **Why It’s Good**:
  - Encapsulates TUI-specific logic, making it independent of plugin or CLI concerns.

## 4. `pkg/`

- **Purpose**:
  - Provides reusable utility packages that can potentially be used by plugins or external tools.

### Subdirectories:
### a. `styles/`
- Houses shared TUI styles (e.g., `lipgloss` styles) for consistent theming.
- **Why It’s Good**:
  - Ensures consistent styling across all TUI pages/plugins.

### b. `utils/`
- Contains generic utilities (e.g., helpers for string manipulation, file operations).
- **Why It’s Good**:
  - Keeps helper functions separate from business logic, improving maintainability.

## 5. `plugins/`

- **Purpose**:
  - Houses self-contained plugins, each managing its own logic and UI.

- **Structure**:
  Each plugin resides in its own subdirectory:

```plaintext
plugins/
├── plugin1/
│   ├── plugin.go       # Plugin entry point
│   ├── views.go        # Plugin-specific views
│   └── tests/
│       └── plugin1_test.go
├── plugin2/
│   └── plugin.go
...