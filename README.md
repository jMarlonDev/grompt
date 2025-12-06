# grompt

This is a silly project focused on the terminal prompt for any shell (I guess). I made this because I was bored and I think [starship.rs](https://starship.rs/) is shit and bloated.

So, this prompt is fully made in go without deps. The config is kinda tedious but it's simple and powerful

# Install

If u want to clone the repo and build it, you can use the [makefile](./makefile).

## Local-install

This command installs automatically grompt in `~/.local/bin`

```bash
curl -sSL https://raw.githubusercontent.com/Esteban528/grompt/refs/heads/master/install.sh | bash
```

# Config

The default config file is generated on `$HOME/.config/grompt.json`

- [Colors](./colors.go)

```jsonc
[
  // The config is a JSON array. Each element is processed in order.
  // Everything shown here is just documentation.

  // 1. Literal strings — printed exactly:
  "@",

  // 2. Environment variables — prefix with $:
  "$USER",

  // 3. Internal variables — use ${name}:
  // Available: dir, hostname, git_branch
  "${dir}",
  "${hostname}",
  "${git_branch}",

  // 4. Colors — c:*, fg:*, bg:*
  "fg:green",
  "bg:red",
  "c:bold",

  // 5. Truecolor — fg:#RRGGBB or bg:#RRGGBB
  "fg:#88CCFF",
  "bg:#112233",

  // 6. exec:command — insert command output
  "exec:date +%H:%M",

  // 7. Some conditionals:
  // If repo is dirty (any change), print inner array:
  { "git_status_noclean": ["fg:red", "X"] },

  // If repo is clean, print inner array:
  { "git_status_clean": ["fg:green", "✔"] },

  // End. The parser will always append c:reset automatically.
]
```

# Popular shell's integration

- **bash**:

Put this on .bashrc
```bash 
PROMPT_COMMAND='PS1="$(grompt | perl -pe 's/\x1b\[([0-9;]*m)/\\[\\e[$1\\]/g')"' 
```

- **zsh**:
```bash
grep -qxF "PROMPT='$(grompt)'" ~/.zshrc || echo "PROMPT='\$(grompt)'" >> ~/.zshrc && source ~/.zshrc
```

- **fish**
```bash
grep -qxF "function fish_prompt" ~/.config/fish/functions/fish_prompt.fish 2>/dev/null || \
echo -e "function fish_prompt\n    grompt\nend" > ~/.config/fish/functions/fish_prompt.fish

fish -c "source ~/.config/fish/functions/fish_prompt.fish"
```

- **dash**
```bash
grep -qxF "PS1='$(grompt)'" ~/.profile 2>/dev/null || echo "PS1='\$(grompt)'" >> ~/.profile && . ~/.profile
```

