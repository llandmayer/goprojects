# UML
::: mermaid
classDiagram

%% Interfaces
class Skeleton {
  + Init() tea.Cmd
  + Update(msg tea.Msg) returns (tea.Model, tea.Cmd)
  + View() string
}
<<interface>> Skeleton

class Layout {
  + Header(width, height int) string
  + Sidebar(width, height, activeIndex int, menuItems []string) string
  + Content(width, height, activeIndex int, menuItems []string) string
  + Footer(width, height int, mode string, txt textinput.Model, suggestions []string) string
}
<<interface>> Layout

%% DefaultLayout implements Layout
class DefaultLayout {
  + Header(width, height int) string
  + Sidebar(width, height, activeIndex int, menuItems []string) string
  + Content(width, height, activeIndex int, menuItems []string) string
  + Footer(width, height int, mode string, txt textinput.Model, suggestions []string) string
}

%% Model implements Skeleton
class Model {
  - Layout Layout
  - Width int
  - Height int
  - MenuItems []string
  - ActiveIndex int
  - Mode string
  - TextInput textinput.Model
  + NewModel(layout Layout, menu []string) *Model
  + Init() tea.Cmd
  + Update(msg tea.Msg) returns (tea.Model, tea.Cmd)
  + View() string
}

Skeleton <|.. Model : implements
Layout <|.. DefaultLayout : implements

%% Layout delegates to these "default" functions internally
class Header {
  + DefaultHeader(width, height int) string
}

class Footer {
  + DefaultFooter(width, height int, mode string, txt textinput.Model, suggestions []string) string
}

class Sidebar {
  + DefaultSidebar(width, height, activeIndex int, menuItems []string) string
}

class Content {
  + DefaultContent(width, height, activeIndex int, menuItems []string) string
}

DefaultLayout ..> Header : calls DefaultHeader
DefaultLayout ..> Footer : calls DefaultFooter
DefaultLayout ..> Sidebar : calls DefaultSidebar
DefaultLayout ..> Content : calls DefaultContent
:::

# Sequence Diagram

::: mermaid
sequenceDiagram
    participant User
    participant BubbleTea as Bubble Tea Framework
    participant Model
    participant Layout as DefaultLayout
    participant Header
    participant Sidebar
    participant Content
    participant Footer

    User->>BubbleTea: Terminal Input
    BubbleTea->>Model: Sends input (e.g., key press)
    Model->>Model: Update state (e.g., activeIndex, mode)
    Model->>BubbleTea: Returns update command
    BubbleTea->>Model: Request View
    Model->>Layout: Call Header(width, height)
    Layout->>Header: Call DefaultHeader(width, height)
    Header-->>Layout: Rendered Header
    Model->>Layout: Call Sidebar(width, height, activeIndex, menuItems)
    Layout->>Sidebar: Call DefaultSidebar(width, height, activeIndex, menuItems)
    Sidebar-->>Layout: Rendered Sidebar
    Model->>Layout: Call Content(width, height, activeIndex, menuItems)
    Layout->>Content: Call DefaultContent(width, height, activeIndex, menuItems)
    Content-->>Layout: Rendered Content
    Model->>Layout: Call Footer(width, mode, textInput, suggestions)
    Layout->>Footer: Call DefaultFooter(width, mode, textInput, suggestions)
    Footer-->>Layout: Rendered Footer
    Layout-->>Model: Full Layout (Header + Sidebar + Content + Footer)
    Model-->>BubbleTea: Rendered Output
    BubbleTea->>User: Updated CLI View
:::