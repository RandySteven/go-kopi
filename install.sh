#!/bin/bash

# Go-Kopi Installer Script
# This script helps you set up Go-Kopi as a starter for your own project

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Default values
REPO_URL="https://github.com/randy-steven/go-kopi.git"
DEFAULT_BRANCH="v2"

print_banner() {
    echo -e "${CYAN}"
    cat << "EOF"
  ________                  ____  __.            .__ 
 /  _____/  ____           |    |/ _|____ ______ |__|
/   \  ___ /  _ \   ______ |      < /  _ \\____ \|  |
\    \_\  (  <_> ) /_____/ |    |  (  <_> )  |_> >  |
 \______  /\____/          |____|__ \____/|   __/|__|
        \/                         \/     |__|       
EOF
    echo -e "${NC}"
    echo -e "${GREEN}Go Backend Framework Installer${NC}"
    echo ""
}

print_help() {
    echo "Usage: $0 [COMMAND] [OPTIONS]"
    echo ""
    echo "Commands:"
    echo "  clone         Clone the repository to a new directory"
    echo "  pull          Pull latest changes from upstream"
    echo "  setup         Set up the project (copy config files, install deps)"
    echo "  remote        Change git remote to your own repository"
    echo "  init          Full initialization (clone + setup + remote)"
    echo "  help          Show this help message"
    echo ""
    echo "Options:"
    echo "  -n, --name        Project name (for clone/init)"
    echo "  -r, --remote      New remote URL (for remote command)"
    echo "  -b, --branch      Branch to clone (default: $DEFAULT_BRANCH)"
    echo ""
    echo "Examples:"
    echo "  # Clone and set up a new project"
    echo "  curl -fsSL https://raw.githubusercontent.com/randy-steven/go-kopi/$DEFAULT_BRANCH/install.sh | bash -s -- init -n my-project -r https://github.com/user/my-project.git"
    echo ""
    echo "  # Clone only"
    echo "  ./install.sh clone -n my-project"
    echo ""
    echo "  # Set up existing project"
    echo "  ./install.sh setup"
    echo ""
    echo "  # Change remote"
    echo "  ./install.sh remote -r https://github.com/user/my-project.git"
    echo ""
    echo "  # Pull latest from upstream"
    echo "  ./install.sh pull"
}

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

check_requirements() {
    log_info "Checking requirements..."
    
    if ! command -v git &> /dev/null; then
        log_error "git is not installed. Please install git first."
        exit 1
    fi
    
    if ! command -v go &> /dev/null; then
        log_warning "Go is not installed. You'll need Go 1.24+ to run this project."
    else
        GO_VERSION=$(go version | awk '{print $3}' | sed 's/go//')
        log_info "Go version: $GO_VERSION"
    fi
    
    log_success "Requirements check completed"
}

clone_repo() {
    local project_name="$1"
    local branch="$2"
    
    if [ -z "$project_name" ]; then
        read -p "Enter project name: " project_name
    fi
    
    if [ -z "$project_name" ]; then
        log_error "Project name is required"
        exit 1
    fi
    
    if [ -d "$project_name" ]; then
        log_error "Directory '$project_name' already exists"
        exit 1
    fi
    
    log_info "Cloning go-kopi to '$project_name'..."
    git clone --branch "$branch" --single-branch "$REPO_URL" "$project_name"
    
    cd "$project_name"
    
    # Remove git history to start fresh
    rm -rf .git
    git init
    git add .
    git commit -m "Initial commit from go-kopi template"
    
    log_success "Repository cloned to '$project_name'"
    echo "$project_name"
}

pull_upstream() {
    log_info "Pulling latest changes from upstream..."
    
    # Check if upstream remote exists
    if ! git remote | grep -q "upstream"; then
        log_info "Adding upstream remote..."
        git remote add upstream "$REPO_URL"
    fi
    
    # Fetch and merge
    git fetch upstream "$DEFAULT_BRANCH"
    git merge upstream/"$DEFAULT_BRANCH" --allow-unrelated-histories -m "Merge upstream changes"
    
    log_success "Successfully pulled latest changes from upstream"
}

setup_project() {
    log_info "Setting up project..."
    
    # Copy environment file
    if [ ! -f "files/env/.env" ]; then
        if [ -f "files/env/.env.example" ]; then
            cp files/env/.env.example files/env/.env
            log_success "Created files/env/.env from example"
        fi
    else
        log_warning "files/env/.env already exists, skipping"
    fi
    
    # Copy YAML config
    if [ ! -f "files/yaml/app.local.yml" ]; then
        if [ -f "files/yaml/app.example.yml" ]; then
            cp files/yaml/app.example.yml files/yaml/app.local.yml
            log_success "Created files/yaml/app.local.yml from example"
        fi
    else
        log_warning "files/yaml/app.local.yml already exists, skipping"
    fi
    
    # Download Go dependencies
    if [ -f "go.mod" ]; then
        log_info "Downloading Go dependencies..."
        go mod download
        go mod tidy
        log_success "Go dependencies installed"
    fi
    
    log_success "Project setup completed"
    
    echo ""
    echo -e "${CYAN}Next steps:${NC}"
    echo "1. Edit files/env/.env with your environment settings"
    echo "2. Edit files/yaml/app.local.yml with your database config"
    echo "3. Run 'make migration' to set up the database"
    echo "4. Run 'make run' to start the server"
}

change_remote() {
    local new_remote="$1"
    
    if [ -z "$new_remote" ]; then
        read -p "Enter your new remote URL: " new_remote
    fi
    
    if [ -z "$new_remote" ]; then
        log_error "Remote URL is required"
        exit 1
    fi
    
    # Check if we're in a git repo
    if [ ! -d ".git" ]; then
        log_error "Not a git repository. Run 'git init' first."
        exit 1
    fi
    
    # Store original as upstream if not already set
    if ! git remote | grep -q "upstream"; then
        CURRENT_ORIGIN=$(git remote get-url origin 2>/dev/null || echo "")
        if [ -n "$CURRENT_ORIGIN" ]; then
            log_info "Saving original remote as 'upstream'..."
            git remote add upstream "$CURRENT_ORIGIN"
        else
            git remote add upstream "$REPO_URL"
        fi
    fi
    
    # Set new origin
    if git remote | grep -q "origin"; then
        git remote set-url origin "$new_remote"
    else
        git remote add origin "$new_remote"
    fi
    
    log_success "Remote changed to: $new_remote"
    log_info "Original go-kopi repo saved as 'upstream'"
    log_info "You can pull updates with: ./install.sh pull"
}

update_module_name() {
    local new_module="$1"
    
    if [ -z "$new_module" ]; then
        read -p "Enter new Go module name (e.g., github.com/user/project): " new_module
    fi
    
    if [ -z "$new_module" ]; then
        log_warning "Module name not provided, skipping module rename"
        return
    fi
    
    log_info "Updating Go module name to '$new_module'..."
    
    # Get current module name
    CURRENT_MODULE=$(head -1 go.mod | awk '{print $2}')
    
    if [ "$CURRENT_MODULE" = "$new_module" ]; then
        log_warning "Module name is already '$new_module'"
        return
    fi
    
    # Update go.mod
    sed -i.bak "s|module $CURRENT_MODULE|module $new_module|g" go.mod && rm go.mod.bak
    
    # Update all import statements in .go files
    find . -name "*.go" -type f -exec sed -i.bak "s|\"$CURRENT_MODULE|\"$new_module|g" {} \; -exec rm {}.bak \;
    
    go mod tidy
    
    log_success "Module name updated to '$new_module'"
}

full_init() {
    local project_name="$1"
    local new_remote="$2"
    local branch="$3"
    
    check_requirements
    
    # Clone
    clone_repo "$project_name" "$branch"
    
    # We're now in the project directory
    
    # Setup
    setup_project
    
    # Change remote if provided
    if [ -n "$new_remote" ]; then
        change_remote "$new_remote"
        
        # Extract module name from remote URL
        MODULE_NAME=$(echo "$new_remote" | sed 's/\.git$//' | sed 's|https://||' | sed 's|git@||' | sed 's|:|/|')
        update_module_name "$MODULE_NAME"
    fi
    
    echo ""
    log_success "Project '$project_name' is ready!"
}

# Parse arguments
COMMAND=""
PROJECT_NAME=""
NEW_REMOTE=""
BRANCH="$DEFAULT_BRANCH"

while [[ $# -gt 0 ]]; do
    case $1 in
        clone|pull|setup|remote|init|help)
            COMMAND="$1"
            shift
            ;;
        -n|--name)
            PROJECT_NAME="$2"
            shift 2
            ;;
        -r|--remote)
            NEW_REMOTE="$2"
            shift 2
            ;;
        -b|--branch)
            BRANCH="$2"
            shift 2
            ;;
        -h|--help)
            COMMAND="help"
            shift
            ;;
        *)
            log_error "Unknown option: $1"
            print_help
            exit 1
            ;;
    esac
done

# Main
print_banner

case $COMMAND in
    clone)
        check_requirements
        clone_repo "$PROJECT_NAME" "$BRANCH"
        ;;
    pull)
        pull_upstream
        ;;
    setup)
        setup_project
        ;;
    remote)
        change_remote "$NEW_REMOTE"
        ;;
    init)
        full_init "$PROJECT_NAME" "$NEW_REMOTE" "$BRANCH"
        ;;
    help|"")
        print_help
        ;;
    *)
        log_error "Unknown command: $COMMAND"
        print_help
        exit 1
        ;;
esac
