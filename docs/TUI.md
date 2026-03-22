# TUI Dashboard Guide

The DefenseClaw TUI is a Bubbletea terminal dashboard for security operators. It surfaces alerts, skill and MCP context, and overall status in one session.

## Panels

1. **Alerts** — Recent findings, policy violations, and enforcement notifications.
2. **Skills** — Skill-oriented scan results and list actions relevant to skills.
3. **MCP Servers** — MCP-oriented findings and server-level context.
4. **Status** — Connection and subsystem status (scanners, store, optional OpenShell).

Note: Full four-panel behavior and live refresh are targeted for **iteration 3**.

## Keybindings

| Key | Action |
|-----|--------|
| Tab | Move focus between panels / sections |
| Arrow keys | Navigate lists and panes |
| `b` | Block selected item (where applicable) |
| `a` | Allow selected item (where applicable) |
| `d` | Dismiss or clear an alert |
| `r` | Rescan selected target |
| Enter | Open detail view for the current selection |
| `/` | Filter or search in the active panel |
| `q` | Quit the TUI |

Exact bindings may vary slightly by panel focus; see in-app help when iteration 3 ships.
