#!/bin/bash
# 展示 Git 变更

set -e

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}📊 Git Status:${NC}"
git status
echo ""

echo -e "${YELLOW}📈 Changes Summary:${NC}"
git diff --stat
echo ""

echo -e "${YELLOW}📝 Detailed Changes:${NC}"
git diff
