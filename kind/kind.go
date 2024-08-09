package kind

const (
	KindIdentifier = iota

	IntLiteral
	StringLiteral

	OpenBraceToken
	CloseBraceToken
	SemicolonToken
	CommaToken
	LessThanToken
	GreaterThanToken
	LessThanEqualsToken
	GreaterThanEqualsToken
	EqualsToken
	EqualsEqualsToken
	MinusToken
	PlusToken
	AsteriskToken
	AsteriskAsteriskToken
	SlashToken
	PlusPlusToken
	MinusMinusToken
	AmpersandToken
	BarToken

	LetKeyword
	ConstKeyword
	ContinueKeyword
	DoKeyword
	EndKeyword

	IfKeyword
	ElseKeyword

	FalseKeyword
	TrueKeyword
	ReturnKeyword

	BinaryExpression
)
