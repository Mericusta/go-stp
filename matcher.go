package stp

import (
	"fmt"
	"strings"
)

var (
	punctuationMarkMap map[rune]rune = map[rune]rune{
		'(': ')', ')': '(',
		'{': '}', '}': '{',
		'[': ']', ']': '[',
		'"': '"', '\'': '\'', '`': '`',
	}
	PunctuationMarkLeftDoubleQuotes   rune = '"'
	PunctuationMarkRightDoubleQuotes  rune = '"'
	PunctuationMarkLeftInverseQuotes  rune = '`'
	PunctuationMarkRightInverseQuotes rune = '`'
	PunctuationMarkLeftSingleQuotes   rune = '\''
	PunctuationMarkRightSingleQuotes  rune = '\''
	PunctuationMarkLeftBracket        rune = '('
	PunctuationMarkRightBracket       rune = ')'
	PunctuationMarkLeftCurlyBracket   rune = '{'
	PunctuationMarkRightCurlyBracket  rune = '}'
	PunctuationMarkLeftSquareBracket  rune = '['
	PunctuationMarkRightSquareBracket rune = ']'
	PunctuationMarkPoint              rune = '.'
	ASCIISpace                        rune = ' '
	InvalidScopePunctuationMarkMap         = map[rune]rune{
		PunctuationMarkLeftDoubleQuotes:  PunctuationMarkRightDoubleQuotes,
		PunctuationMarkLeftInverseQuotes: PunctuationMarkRightInverseQuotes,
		PunctuationMarkLeftSingleQuotes:  PunctuationMarkRightSingleQuotes,
	}
	ScopePunctuationMarkList = []int{
		PunctuationMarkBracket,
		PunctuationMarkCurlyBracket,
		PunctuationMarkSquareBracket,
	}
)

// GetAnotherPunctuationMark 获取标点符号的另一对
func GetAnotherPunctuationMark(r rune) rune {
	if markRune, hasMark := punctuationMarkMap[r]; hasMark {
		return markRune
	}
	return ' '
}

func GetLeftPunctuationMarkList() []rune {
	leftPunctuationMarkList := []rune{
		'(', '{', '[', '"', '\'', '`',
	}
	return leftPunctuationMarkList
}

func IsSpaceRune(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

func IsCharacter(r rune) bool {
	return ('0' <= r && r <= '9') || ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || r == '_'
}

// CalculatePunctuationMarksContentLength 计算成对符号的内容长度，去除结束符号
// @contentAfterLeftPunctuationMark       待计算的字符串，不包括起始符号
// @leftPunctuationMark                   符号左边界字符 必须是 (, [, {, ", ', ` 之一
// @rightPunctuationMark                  符号右边界字符
// @invalidScopePunctuationMarkMap        排除计算的边界符号
// @return                                不包含左右边界的内容的长度
func CalculatePunctuationMarksContentLength(contentAfterLeftPunctuationMark string, leftPunctuationMark, rightPunctuationMark rune, invalidScopePunctuationMarkMap map[rune]rune) int {
	byteCount := 0
	leftRuneCount := 1
	rightRuneCount := 0
	isValid := true
	var invalidScopePunctuationMark rune = -1
	var runeLength int
	strings.IndexFunc(contentAfterLeftPunctuationMark, func(r rune) bool {
		runeLength = len([]byte(string(r)))
		byteCount += runeLength

		// fmt.Printf("rune %v, string = %v, byte = %v\n", r, string(r), []byte(string(r)))

		// end invalid scope
		if !isValid && r == invalidScopePunctuationMark {
			isValid = true
			invalidScopePunctuationMark = -1
			return false
		}

		// in invalid scope
		if !isValid {
			return false
		}

		// begin invalid scope
		if punctuationMark, isInvalidScopePunctuationMark := invalidScopePunctuationMarkMap[r]; isValid && isInvalidScopePunctuationMark {
			isValid = false
			invalidScopePunctuationMark = punctuationMark
			return false
		}

		// out invalid scope
		if r == leftPunctuationMark {
			leftRuneCount++
		} else if r == rightPunctuationMark {
			rightRuneCount++
		}

		return leftRuneCount == rightRuneCount
	})
	return byteCount - len([]byte(string(rightPunctuationMark))) // // cut right punctuation mark len
}

// GetScopeContentBetweenPunctuationMarks 获取成对标点符号的内容
// @content                               待查找的内容
// @scopeBeginIndex                       左边界起始符号的下标，[]byte 数组的下标
// @return                                不包含左右边界的内容
func GetScopeContentBetweenPunctuationMarks(content string, scopeBeginIndex int, scopeBeginRune, scopeEndRune rune) string {
	runeLength := len([]byte(string(scopeBeginRune)))
	scopeContentLength := CalculatePunctuationMarksContentLength(
		content[scopeBeginIndex+runeLength:], scopeBeginRune, scopeEndRune,
		InvalidScopePunctuationMarkMap,
	)
	// 直接 content[x:y] 是按照 []byte[x:y] 来处理的
	return content[scopeBeginIndex+runeLength : scopeBeginIndex+runeLength+scopeContentLength]
}

// // GetScopeContentBetweenPunctuationMarks 获取成对标点符号的内容
// // @content                               待查找的内容
// // @scopeBeginIndex                       左边界起始符号的下标，[]byte 数组的下标
// // @return                                不包含左右边界的内容
// func GetScopeContentBetweenPunctuationMarks(content string, scopeBeginIndex int) string {
// 	scopeBeginRune := rune(content[scopeBeginIndex])
// 	scopeEndRune := GetAnotherPunctuationMark(scopeBeginRune)
// 	scopeContentRuneLength := CalculatePunctuationMarksContentLength(
// 		string(content[scopeBeginIndex+1:]),
// 		scopeBeginRune, scopeEndRune,
// 		InvalidScopePunctuationMarkMap,
// 	)
// 	// 直接 content[x:y] 是按照 []byte[x:y] 来处理的
// 	return string([]rune(content)[scopeBeginIndex+1 : scopeBeginIndex+1+scopeContentRuneLength])
// }

// PunctuationIndex 成对标点符号的下标数据
type PunctuationIndex struct {
	Left  int
	Right int
}

type PunctuationMarkInfo struct {
	PunctuationMark rune
	Index           int
}

// NewPunctuationContent 成对标点符号的内容节点
type NewPunctuationContent struct {
	Content                   string
	LeftPunctuationMark       *PunctuationMarkInfo
	RightPunctuationMark      *PunctuationMarkInfo
	SubPunctuationContentList []*NewPunctuationContent
}

// 逻辑常量
const (
	PunctuationMarkQuote         = 1
	PunctuationMarkBracket       = 2
	PunctuationMarkCurlyBracket  = 3
	PunctuationMarkSquareBracket = 4
)

// GetPunctuationMark 获取标点符号
func GetPunctuationMark(punctuationMark int) (rune, rune) {
	switch punctuationMark {
	case PunctuationMarkCurlyBracket:
		return PunctuationMarkLeftCurlyBracket, PunctuationMarkRightCurlyBracket
	case PunctuationMarkSquareBracket:
		return PunctuationMarkLeftSquareBracket, PunctuationMarkRightSquareBracket
	default:
		return PunctuationMarkLeftBracket, PunctuationMarkRightBracket
	}
}

// TraitMultiPunctuationMarksContent 混合成对标点符号的内容分类提取
func TraitMultiPunctuationMarksContent(content string, punctuationMarkList []int, maxDeep int) *NewPunctuationContent {
	leftPunctuationMarkList := make([]rune, 0, len(punctuationMarkList))
	for _, punctuationMark := range punctuationMarkList {
		leftPunctuationMark, _ := GetPunctuationMark(punctuationMark)
		leftPunctuationMarkList = append(leftPunctuationMarkList, leftPunctuationMark)
	}
	return RecursiveTraitMultiPunctuationMarksContent(content, &PunctuationMarkInfo{
		PunctuationMark: 0,
		Index:           -1,
	}, &PunctuationMarkInfo{
		PunctuationMark: 0,
		Index:           len(content),
	}, leftPunctuationMarkList, maxDeep, 0)
}

// RecursiveTraitMultiPunctuationMarksContent 混合成对标点符号的内容分类提取
// @content 待处理内容
// @leftPunctuationMarkInfo 根节点的左标点符号
// @rightPunctuationMarkInfo 根节点的右标点符号
// @scopeLeftPunctuationMarkList 所有作为划分区域的左标点符号
// @maxDeep 待处理的最大深度
// @deep 当前深度
// @return 根节点
func RecursiveTraitMultiPunctuationMarksContent(content string, leftPunctuationMarkInfo, rightPunctuationMarkInfo *PunctuationMarkInfo, scopeLeftPunctuationMarkList []rune, maxDeep, deep int) *NewPunctuationContent {
	punctuationContent := &NewPunctuationContent{
		Content:                   content,
		LeftPunctuationMark:       leftPunctuationMarkInfo,
		RightPunctuationMark:      rightPunctuationMarkInfo,
		SubPunctuationContentList: make([]*NewPunctuationContent, 0),
	}

	passLeftLength := 0
	for len(content) != 0 && deep != maxDeep {
		var leftPunctuationMark rune
		var rightPunctuationMark rune
		leftPunctuationMarkIndex := len(content) - 1

		for _, toSearchLeftPunctuationMark := range scopeLeftPunctuationMarkList {
			toSearchLeftPunctuationMarkIndex := strings.IndexRune(content, toSearchLeftPunctuationMark)
			if toSearchLeftPunctuationMarkIndex != -1 && toSearchLeftPunctuationMarkIndex < leftPunctuationMarkIndex {
				leftPunctuationMarkIndex = toSearchLeftPunctuationMarkIndex
				leftPunctuationMark = toSearchLeftPunctuationMark
			}
		}
		// fmt.Printf("relative leftPunctuationMarkIndex = %v, leftPunctuationMark = %v\n", leftPunctuationMarkIndex, string(rune(leftPunctuationMark)))

		rightPunctuationMark = GetAnotherPunctuationMark(leftPunctuationMark)
		if leftPunctuationMark == 0 || rightPunctuationMark == 0 || leftPunctuationMarkIndex == len(content)-1 {
			break
		}

		afterLeftPunctuationMarkContentIndex := leftPunctuationMarkIndex + 1

		// fmt.Printf("pass CalculatePunctuationMarksContentLength = |%v|\n", content[afterLeftPunctuationMarkContentIndex:])
		length := CalculatePunctuationMarksContentLength(content[afterLeftPunctuationMarkContentIndex:], leftPunctuationMark, rightPunctuationMark, InvalidScopePunctuationMarkMap)

		// fmt.Printf("after CalculatePunctuationMarksContentLength, length = %v\n", length)

		rightPunctuationMarkIndex := leftPunctuationMarkIndex + length + 1
		if rightPunctuationMarkIndex >= len(content) {
			// fmt.Printf("rightPunctuationMarkIndex %v >= len(content) %v\n", rightPunctuationMarkIndex, len(content))
			break
		}

		// fmt.Printf("relative rightPunctuationMarkIndex = %v, rightPunctuationMark = %v\n", rightPunctuationMarkIndex, string(rune(rightPunctuationMark)))
		// fmt.Printf("pass content = |%v|\n", content[leftPunctuationMarkIndex+1:rightPunctuationMarkIndex])

		subPunctuationContent := RecursiveTraitMultiPunctuationMarksContent(content[leftPunctuationMarkIndex+1:rightPunctuationMarkIndex], &PunctuationMarkInfo{
			PunctuationMark: leftPunctuationMark,
			Index:           leftPunctuationMarkInfo.Index + 1 + passLeftLength + leftPunctuationMarkIndex,
		}, &PunctuationMarkInfo{
			PunctuationMark: rightPunctuationMark,
			Index:           leftPunctuationMarkInfo.Index + 1 + passLeftLength + rightPunctuationMarkIndex,
		}, scopeLeftPunctuationMarkList, maxDeep, deep+1)
		if subPunctuationContent != nil {
			punctuationContent.SubPunctuationContentList = append(punctuationContent.SubPunctuationContentList, subPunctuationContent)
		}

		// fmt.Printf("update content to |%v|\n", content[rightPunctuationMarkIndex+1:])
		content = content[rightPunctuationMarkIndex+1:]
		// fmt.Printf("update pass left from %v to %v\n", passLeftLength, passLeftLength+rightPunctuationMarkIndex+1)
		passLeftLength += rightPunctuationMarkIndex + 1
		// fmt.Println("--------------------------------")
	}

	return punctuationContent
}

// SplitContent 划分内容节点
type SplitContent struct {
	ContentList         []string
	SubSplitContentList []*SplitContent
}

// RecursiveSplitUnderSameDeepPunctuationMarksContent 相同深度的成对标点符号下的内容划分
// @content 待分析的字符串
// @punctuationMarkList 指定成对标点符号
// @splitter 指定分隔符
// @return
func RecursiveSplitUnderSameDeepPunctuationMarksContent(content string, punctuationMarkList []int, splitter string) *SplitContent {
	if punctuationContentNode := TraitMultiPunctuationMarksContent(content, punctuationMarkList, 1); punctuationContentNode != nil {
		return splitUnderSameDeepPunctuationMarksContent(punctuationContentNode, splitter, 0, 0)
	}
	return nil
}

// RecursiveSplitUnderSameDeepPunctuationMarksContentNode 相同深度的成对标点符号下的内容划分
// @punctuationContentNode 成对标点符号的内容根节点，注意：必须是根节点，不能是某个子节点，节点深度必须为 2
// @splitter 指定分隔符
// @return
func RecursiveSplitUnderSameDeepPunctuationMarksContentNode(punctuationContentNode *NewPunctuationContent, splitter string) *SplitContent {
	return splitUnderSameDeepPunctuationMarksContent(punctuationContentNode, splitter, 0, 0)
}

// splitUnderSameDeepPunctuationMarksContent 相同深度的成对标点符号下的内容划分的递归算法
// @punctuationContentNode 成对标点符号的内容根节点，注意：必须是根节点，不能是某个子节点，节点深度 >= 2，分析结果中深度大于 2 的数据不正确
// @splitter 指定分隔符
// @maxDeep 递归最大深度
// @deep 当前深度
func splitUnderSameDeepPunctuationMarksContent(punctuationContentNode *NewPunctuationContent, splitter string, maxDeep, deep int) *SplitContent {
	splitContentNode := &SplitContent{
		ContentList:         make([]string, 0),
		SubSplitContentList: make([]*SplitContent, 0),
	}

	var offset int
	var leftIndex int
	cycle := 0
	maxCycle := len(strings.Split(punctuationContentNode.Content, splitter))
	for cycle != maxCycle {
		cycle++
		length := strings.Index(punctuationContentNode.Content[leftIndex+offset:], splitter)
		if length == -1 {
			splitContentNode.ContentList = append(splitContentNode.ContentList, punctuationContentNode.Content[leftIndex:])
			break
		}
		rightIndex := leftIndex + length + offset
		inner := false
		for _, subNode := range punctuationContentNode.SubPunctuationContentList {
			// Note: 这里用于判断的依据是子节点相对父节点的左 区间符号 的下标
			// Note: 但是节点的 区间符号 数据中记录的下标是相对于根节点的下标 -> 必须是根节点
			// Note: 所以当节点数只有2时，这个下标可以代表相对父节点（根节点）的下标 -> 节点深度 >= 2
			if subNode.LeftPunctuationMark.Index <= rightIndex && rightIndex <= subNode.RightPunctuationMark.Index {
				inner = true
				offset = subNode.RightPunctuationMark.Index - leftIndex + 1
				break
			}
		}
		if inner {
			continue
		}
		splitContentNode.ContentList = append(splitContentNode.ContentList, punctuationContentNode.Content[leftIndex:rightIndex])
		offset = 0
		leftIndex = rightIndex + len(splitter)
	}

	if deep == maxDeep {
		return splitContentNode
	}

	for _, subPuncutationContentNode := range punctuationContentNode.SubPunctuationContentList {
		if len(subPuncutationContentNode.Content) != 0 {
			if subSplitContentNode := splitUnderSameDeepPunctuationMarksContent(subPuncutationContentNode, splitter, maxDeep, deep+1); subSplitContentNode != nil {
				splitContentNode.SubSplitContentList = append(splitContentNode.SubSplitContentList, subSplitContentNode)
			}
		}
	}

	return splitContentNode
}

// ConvertSnakeCase2CamelCase 将蛇形命名法转换为驼峰命名法：xxx_yyy_zzz -> [X|x]xxYyyZzz
func ConvertSnakeCase2CamelCase(snakeCase string, capitalize bool) string {
	camelStyleString := ""
	for _, singleString := range strings.Split(snakeCase, "_") {
		capitalizeSingleString := fmt.Sprintf("%v%v", strings.ToUpper(singleString[:1]), singleString[1:])
		camelStyleString = fmt.Sprintf("%v%v", camelStyleString, capitalizeSingleString)
	}
	if !capitalize {
		camelStyleString = fmt.Sprintf("%v%v", strings.ToLower(camelStyleString[:1]), camelStyleString[1:])
	}
	return camelStyleString
}

// ABC 开头 | 中间 DONE
// ConvertSnakeCase2CamelCaseWithAbbreviation 将蛇形命名法转换为驼峰命名法：xxx_yyy_zzz -> [X|x]xxYyyZzz
func ConvertSnakeCase2CamelCaseWithAbbreviation(snakeCase string, capitalize bool, abbreviationMap map[string]struct{}) string {
	camelStyleString := ""
	for _, singleString := range strings.Split(snakeCase, "_") {
		if _, isPhrase := abbreviationMap[singleString]; isPhrase {
			camelStyleString = fmt.Sprintf("%v%v", camelStyleString, singleString)
		} else {
			capitalizeSingleString := fmt.Sprintf("%v%v", strings.ToUpper(singleString[:1]), singleString[1:])
			camelStyleString = fmt.Sprintf("%v%v", camelStyleString, capitalizeSingleString)
		}
	}
	if !capitalize {
		camelStyleString = fmt.Sprintf("%v%v", strings.ToLower(camelStyleString[:1]), camelStyleString[1:])
	}
	return camelStyleString
}

// ConvertCamelCase2SnakeCase 将驼峰命名法转换为蛇形命名法：XxxYyyZzz -> xxx_yyy_zzz
func ConvertCamelCase2SnakeCase(camelCase string) string {
	builder := strings.Builder{}
	for index, r := range camelCase {
		if 'a' <= r && r <= 'z' {
			builder.WriteRune(r)
			continue
		}
		if index != 0 {
			builder.WriteRune('_')
		}
		builder.WriteRune(r + 32)
	}
	return builder.String()
}

// Abc 开头 | 中间 DONE
// ABC 开头 | 中间 DONE
// ConvertCamelCase2SnakeCaseWithAbbreviation 将驼峰命名法转换为蛇形命名法：XxxYyyZzz -> xxx_yyy_zzz
func ConvertCamelCase2SnakeCaseWithAbbreviation(camelCase string, abbreviationMap map[string]struct{}) string {
	allAbbreviationSubString := make(map[string]struct{})
	for abbreviation := range abbreviationMap {
		for index := 0; index != len(abbreviation); index++ {
			allAbbreviationSubString[abbreviation[0:index]] = struct{}{}
		}
	}

	builder := strings.Builder{}
	abbreviationBuilder := strings.Builder{}
	isFirstAbbreviation := true
	for _, r := range camelCase {
		if 'a' <= r && r <= 'z' {
			abbreviationBuilder.WriteRune(r)
		} else {
			if abbreviationBuilder.Len() > 0 {
				if _, isAbbreviation := abbreviationMap[abbreviationBuilder.String()]; isAbbreviation {
					if isFirstAbbreviation {
						isFirstAbbreviation = false
					} else {
						builder.WriteRune('_')
					}
					builder.WriteString(abbreviationBuilder.String())
					abbreviationBuilder.Reset()
				} else {
					if _, maybeAbbreviation := allAbbreviationSubString[abbreviationBuilder.String()]; !maybeAbbreviation {
						if isFirstAbbreviation {
							isFirstAbbreviation = false
						} else {
							builder.WriteRune('_')
						}
						builder.WriteString(abbreviationBuilder.String())
						abbreviationBuilder.Reset()
					}
				}
			}
			abbreviationBuilder.WriteRune(r + 32)
		}
	}
	builder.WriteRune('_')
	builder.WriteString(abbreviationBuilder.String())
	return builder.String()
}

func SplitAny(groupSplitter, itemSplitter byte, content string) [][]string {
	groupSlice := make([][]string, 0)
	contentLength := len(content)
	if contentLength > 0 {
		groupSlice = append(groupSlice, make([]string, 0))
	}

	itemStartIndex := 0
	groupIndex := 0
	for index := 0; index != contentLength; index++ {
		switch {
		case index+1 == contentLength:
			groupSlice[groupIndex] = append(groupSlice[groupIndex], content[itemStartIndex:contentLength])
			itemStartIndex = index + 1
			groupIndex++
		case content[index] == groupSplitter:
			groupSlice[groupIndex] = append(groupSlice[groupIndex], content[itemStartIndex:index])
			itemStartIndex = index + 1
			groupSlice = append(groupSlice, []string{})
			groupIndex++
		case content[index] == itemSplitter:
			groupSlice[groupIndex] = append(groupSlice[groupIndex], content[itemStartIndex:index])
			itemStartIndex = index + 1
		default:
			continue
		}
	}

	return groupSlice
}

// CleanFileComment 置空所有注释
func CleanFileComment(fileContent []byte) string {
	isBlock, isComment := false, false
	firstCommentIndex, secondCommentIndex := -1, -1
	builder, commentBuffer := strings.Builder{}, strings.Builder{}
	for index, b := range fileContent {
		switch rune(b) {
		case PunctuationMarkLeftDoubleQuotes:
			if !isComment {
				if !isBlock {
					isBlock = true
				} else {
					isBlock = false
				}
			}
		case '/':
			if !isBlock {
				if firstCommentIndex == -1 {
					firstCommentIndex = index
				} else if secondCommentIndex == -1 {
					secondCommentIndex = index
					isComment = true
					commentBuffer.Reset()
				}
			}
		case '\n':
			if isComment {
				isComment = false
				firstCommentIndex = -1
				secondCommentIndex = -1
				commentBuffer.Reset()
			}
		}

		if !isComment {
			if firstCommentIndex != -1 && secondCommentIndex == -1 {
				if commentBuffer.Len() > 0 {
					// just one /, clear comment buffer
					builder.WriteString(commentBuffer.String())
					builder.WriteByte(b)
					firstCommentIndex = -1
					commentBuffer.Reset()
				} else {
					// first match /
					commentBuffer.WriteByte(b)
				}
			} else {
				builder.WriteByte(b)
			}
		}
	}
	return builder.String()
}
