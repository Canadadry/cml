package token

type Token struct {
	Kind    TokenKind
	Literal string
}

type TokenKind string

const (
	KindIdentifier       TokenKind = "KindIdentifier"
	KindString                     = "KindString"
	KindInt                        = "KindInt"
	KindFloat                      = "KindFloat"
	KindTrue                       = "KindTrue"
	KindFalse                      = "KindFalse"
	KindLeftParenthesis            = "KindLeftParenthesis"
	KindRightParenthesis           = "KindRightParenthesis"
	KindComment                    = "KindComment"
	KindInvalid                    = "KindInvalid"
	KindEOF                        = "KindEOF"
)
