#!/bin/bash
# 运行 CLI 命令单元测试

set -e

# 颜色定义
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查参数
TEST_NAME="${1:-}"

echo -e "${YELLOW}🧪 Running CLI unit tests...${NC}"
echo ""

if [ -n "$TEST_NAME" ]; then
    echo -e "${YELLOW}Running tests matching: $TEST_NAME${NC}"
    go test -v ./test/unit/cmd/ -run "$TEST_NAME" -count=1
else
    echo -e "${YELLOW}Running all CLI command tests...${NC}"
    go test -v ./test/unit/cmd/ -count=1
fi

echo ""
echo -e "${GREEN}✅ Tests completed!${NC}"
