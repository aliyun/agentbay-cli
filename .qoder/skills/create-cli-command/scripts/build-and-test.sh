#!/bin/bash
# 编译并测试 CLI

set -e

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🔨 Building CLI...${NC}"
go build -o agentbay .
echo -e "${GREEN}✅ Build successful!${NC}"
echo ""

echo -e "${YELLOW}🧪 Running unit tests...${NC}"
go test -v ./test/unit/cmd/ -count=1
echo ""

echo -e "${GREEN}✅ All tests passed!${NC}"
