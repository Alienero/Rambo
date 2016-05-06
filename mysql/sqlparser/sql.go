//line sql.y:2
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:2
import "strings"

func setParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
	yylex.(*Tokenizer).nesting++
	if yylex.(*Tokenizer).nesting == 200 {
		return true
	}
	return false
}

func decNesting(yylex interface{}) {
	yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

//line sql.y:32
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
	sqlID       SQLName
	sqlIDs      []SQLName
}

const LEX_ERROR = 57346
const UNION = 57347
const MINUS = 57348
const EXCEPT = 57349
const INTERSECT = 57350
const SELECT = 57351
const INSERT = 57352
const UPDATE = 57353
const DELETE = 57354
const FROM = 57355
const WHERE = 57356
const GROUP = 57357
const HAVING = 57358
const ORDER = 57359
const BY = 57360
const LIMIT = 57361
const FOR = 57362
const ALL = 57363
const DISTINCT = 57364
const AS = 57365
const EXISTS = 57366
const ASC = 57367
const DESC = 57368
const INTO = 57369
const DUPLICATE = 57370
const KEY = 57371
const DEFAULT = 57372
const SET = 57373
const LOCK = 57374
const KEYRANGE = 57375
const VALUES = 57376
const LAST_INSERT_ID = 57377
const JOIN = 57378
const STRAIGHT_JOIN = 57379
const LEFT = 57380
const RIGHT = 57381
const INNER = 57382
const OUTER = 57383
const CROSS = 57384
const NATURAL = 57385
const USE = 57386
const FORCE = 57387
const ON = 57388
const ID = 57389
const STRING = 57390
const NUMBER = 57391
const VALUE_ARG = 57392
const LIST_ARG = 57393
const COMMENT = 57394
const NULL = 57395
const TRUE = 57396
const FALSE = 57397
const OR = 57398
const AND = 57399
const NOT = 57400
const BETWEEN = 57401
const CASE = 57402
const WHEN = 57403
const THEN = 57404
const ELSE = 57405
const LE = 57406
const GE = 57407
const NE = 57408
const NULL_SAFE_EQUAL = 57409
const IS = 57410
const LIKE = 57411
const REGEXP = 57412
const IN = 57413
const SHIFT_LEFT = 57414
const SHIFT_RIGHT = 57415
const UNARY = 57416
const END = 57417
const CREATE = 57418
const ALTER = 57419
const DROP = 57420
const RENAME = 57421
const ANALYZE = 57422
const TABLE = 57423
const INDEX = 57424
const VIEW = 57425
const TO = 57426
const IGNORE = 57427
const IF = 57428
const UNIQUE = 57429
const USING = 57430
const SHOW = 57431
const DESCRIBE = 57432
const EXPLAIN = 57433

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"KEYRANGE",
	"VALUES",
	"LAST_INSERT_ID",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"'('",
	"','",
	"')'",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'~'",
	"UNARY",
	"'.'",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 77,
	90, 231,
	-2, 230,
}

const yyNprod = 235
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 756

var yyAct = [...]int{

	109, 350, 188, 382, 104, 191, 61, 103, 297, 72,
	291, 404, 372, 105, 190, 4, 257, 229, 148, 251,
	73, 214, 68, 87, 143, 93, 147, 18, 279, 79,
	146, 94, 37, 47, 52, 49, 53, 361, 53, 50,
	303, 138, 317, 319, 360, 220, 62, 75, 359, 98,
	81, 74, 194, 84, 55, 56, 57, 116, 218, 80,
	77, 117, 118, 119, 83, 116, 120, 99, 77, 117,
	118, 119, 58, 123, 120, 54, 369, 244, 221, 159,
	131, 123, 173, 174, 175, 176, 177, 172, 135, 127,
	142, 137, 106, 107, 128, 172, 427, 157, 108, 318,
	106, 107, 130, 82, 145, 192, 108, 162, 292, 193,
	195, 196, 122, 175, 176, 177, 172, 102, 133, 82,
	122, 217, 219, 216, 160, 124, 203, 75, 264, 240,
	75, 74, 210, 209, 74, 77, 205, 248, 162, 211,
	280, 262, 263, 261, 292, 152, 341, 225, 204, 224,
	208, 187, 189, 207, 102, 102, 70, 116, 99, 247,
	210, 158, 197, 198, 260, 144, 256, 200, 201, 265,
	266, 267, 70, 269, 270, 271, 272, 273, 274, 275,
	276, 277, 278, 243, 86, 245, 134, 246, 82, 268,
	18, 19, 20, 21, 88, 226, 116, 284, 102, 70,
	99, 99, 399, 416, 280, 152, 60, 152, 102, 289,
	249, 250, 22, 102, 102, 300, 59, 258, 280, 281,
	283, 129, 301, 286, 288, 29, 285, 282, 161, 160,
	296, 400, 171, 170, 178, 179, 173, 174, 175, 176,
	177, 172, 89, 162, 161, 160, 373, 305, 129, 307,
	102, 102, 378, 280, 304, 315, 306, 295, 397, 162,
	338, 284, 155, 280, 328, 330, 331, 332, 294, 121,
	324, 227, 322, 23, 24, 26, 25, 27, 329, 373,
	152, 152, 152, 152, 116, 334, 28, 30, 17, 355,
	99, 161, 160, 282, 280, 75, 75, 371, 154, 74,
	348, 340, 365, 145, 346, 349, 162, 145, 345, 325,
	326, 327, 18, 336, 258, 335, 116, 337, 227, 280,
	171, 170, 178, 179, 173, 174, 175, 176, 177, 172,
	252, 254, 255, 155, 299, 253, 366, 227, 141, 126,
	102, 129, 342, 358, 370, 102, 368, 117, 118, 119,
	116, 379, 120, 70, 380, 383, 312, 364, 375, 376,
	384, 313, 151, 51, 310, 367, 357, 309, 393, 311,
	314, 308, 235, 236, 395, 411, 377, 396, 125, 75,
	423, 212, 140, 398, 199, 242, 65, 76, 405, 405,
	405, 18, 424, 63, 351, 284, 392, 394, 410, 408,
	406, 407, 413, 383, 414, 385, 386, 67, 102, 102,
	419, 412, 387, 388, 389, 352, 344, 421, 298, 7,
	6, 75, 151, 391, 151, 74, 363, 5, 428, 429,
	425, 145, 69, 415, 259, 417, 418, 40, 39, 92,
	71, 409, 85, 426, 102, 38, 90, 420, 402, 403,
	18, 42, 97, 1, 231, 234, 235, 236, 232, 69,
	233, 237, 241, 333, 356, 238, 132, 32, 33, 34,
	35, 136, 41, 156, 139, 18, 19, 20, 21, 213,
	153, 171, 170, 178, 179, 173, 174, 175, 176, 177,
	172, 48, 43, 44, 45, 46, 302, 151, 151, 151,
	151, 171, 170, 178, 179, 173, 174, 175, 176, 177,
	172, 280, 215, 69, 78, 206, 32, 33, 34, 35,
	347, 287, 293, 114, 422, 401, 222, 381, 390, 223,
	362, 259, 115, 339, 202, 290, 111, 110, 239, 374,
	153, 343, 153, 97, 163, 100, 116, 316, 280, 77,
	117, 118, 119, 150, 230, 120, 112, 113, 228, 149,
	101, 96, 123, 64, 31, 66, 114, 178, 179, 173,
	174, 175, 176, 177, 172, 115, 91, 15, 14, 16,
	36, 106, 107, 95, 3, 97, 97, 108, 13, 116,
	12, 11, 77, 117, 118, 119, 10, 9, 120, 112,
	113, 122, 8, 101, 2, 123, 0, 18, 0, 0,
	0, 0, 0, 0, 0, 153, 153, 153, 153, 0,
	0, 0, 114, 0, 106, 107, 95, 0, 320, 321,
	108, 115, 323, 0, 0, 0, 0, 114, 0, 0,
	0, 0, 0, 0, 122, 116, 115, 0, 77, 117,
	118, 119, 0, 0, 120, 112, 113, 0, 0, 101,
	116, 123, 0, 77, 117, 118, 119, 0, 0, 120,
	112, 113, 0, 0, 101, 97, 123, 0, 0, 0,
	106, 107, 0, 0, 0, 0, 108, 353, 0, 0,
	354, 0, 0, 0, 0, 106, 107, 0, 0, 0,
	122, 108, 231, 234, 235, 236, 232, 0, 233, 237,
	165, 168, 0, 0, 0, 122, 180, 181, 182, 183,
	184, 185, 186, 169, 166, 167, 164, 171, 170, 178,
	179, 173, 174, 175, 176, 177, 172, 171, 170, 178,
	179, 173, 174, 175, 176, 177, 172, 170, 178, 179,
	173, 174, 175, 176, 177, 172,
}
var yyPact = [...]int{

	181, -1000, -1000, -1000, 511, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 466, -1000, -1000,
	-1000, -1000, -1000, -64, -65, -22, -43, -25, 166, 156,
	-1000, 441, 372, -1000, -1000, -1000, -1000, 511, -1000, -1000,
	-1000, 364, -1000, -63, 122, 427, 85, -73, -39, 69,
	-1000, -33, 69, -1000, 122, -79, 144, -79, 122, 426,
	-1000, -1000, -1000, -1000, 542, -1000, 70, 351, 308, -1,
	-1000, 122, 173, -1000, 35, -1000, -10, -1000, 122, 57,
	136, -1000, -1000, 122, -1000, -59, 122, 358, 292, 69,
	-1000, 90, 149, 285, -1000, -1000, 138, -11, 185, 649,
	-1000, 613, 598, -1000, -1000, -1000, 10, 10, 10, 269,
	269, -1000, -1000, -1000, 269, 269, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 10, -1000, 122, 85, 122, 417, 85,
	10, 69, -1000, 357, -83, -1000, 28, -1000, 122, -1000,
	-1000, 122, -1000, -1000, 10, 613, 223, 666, -1000, -1000,
	106, 362, 303, -13, 149, 542, -1000, -1000, 69, 53,
	613, 613, 274, 10, 110, 66, 10, 10, 10, 274,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 33, 649, 169,
	462, 245, 649, -1000, 18, -1000, -1000, 499, 542, -1000,
	441, 296, 44, 659, 237, 293, -1000, 401, -1000, 659,
	-1000, -1000, -1000, 288, 69, -1000, -60, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 659, 185, 149, 149, 149,
	149, -1000, 335, 331, -1000, 328, 320, 334, -2, -1000,
	122, 122, -1000, 270, 122, 289, -1000, -1000, -1000, 33,
	64, -1000, -1000, 253, -1000, -1000, 659, -1000, 18, -1000,
	-1000, 110, 10, 10, 10, 659, 659, 403, -1000, 487,
	668, -1000, 29, 29, 8, 8, 8, 0, 0, -1000,
	-1000, -1000, 10, -1000, 659, -1000, 214, 542, 214, 212,
	80, -1000, 613, 382, 85, 85, 401, 375, 397, 122,
	-1000, -1000, 122, -1000, 666, 243, 418, -1000, -1000, -1000,
	-1000, 330, -1000, 307, -1000, -1000, -1000, -50, -54, -61,
	-1000, -1000, -1000, -1000, 411, -1000, -1000, -1000, 245, -1000,
	659, 659, 242, 10, 659, -1000, 214, -1000, 296, -15,
	-1000, 10, 232, 233, 269, 511, 200, 204, -1000, 375,
	-1000, 10, 10, -1000, -1000, 613, 613, -1000, -1000, 269,
	269, 269, 407, 378, -1000, 10, 659, -1000, 91, -1000,
	659, 10, -1000, 349, 210, -1000, -1000, -1000, 85, -1000,
	154, 183, -1000, 423, -1000, 185, 185, 69, 69, 69,
	401, 613, 10, 659, -1000, 659, 346, 269, -1000, 10,
	10, -1000, -1000, -1000, 155, -1000, 155, 155, 375, 185,
	179, 436, -1000, 659, -1000, -1000, 69, -1000, -1000, 360,
	85, -1000, -1000, 432, 19, 173, -1000, 69, 69, -1000,
}
var yyPgo = [...]int{

	0, 604, 14, 427, 420, 419, 602, 597, 596, 591,
	590, 588, 584, 580, 579, 578, 577, 576, 472, 565,
	564, 563, 25, 31, 561, 30, 26, 18, 559, 558,
	17, 554, 553, 22, 547, 11, 24, 49, 545, 544,
	541, 7, 2, 19, 16, 5, 539, 13, 537, 269,
	4, 536, 535, 10, 534, 533, 530, 528, 8, 527,
	3, 525, 1, 524, 522, 520, 12, 9, 20, 363,
	184, 514, 512, 496, 491, 479, 0, 473, 387, 465,
	462, 6, 453, 451, 52, 28,
}
var yyR1 = [...]int{

	0, 82, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 15, 15, 17, 17,
	16, 12, 13, 13, 13, 13, 2, 2, 2, 3,
	3, 4, 5, 6, 7, 7, 7, 8, 8, 8,
	9, 10, 10, 10, 11, 14, 83, 18, 19, 19,
	20, 20, 20, 20, 20, 21, 21, 22, 22, 23,
	23, 23, 24, 24, 77, 77, 77, 25, 25, 26,
	26, 27, 27, 27, 28, 28, 28, 28, 80, 80,
	79, 79, 79, 29, 29, 29, 29, 30, 30, 30,
	30, 31, 31, 32, 32, 33, 33, 34, 34, 34,
	34, 35, 35, 36, 36, 37, 37, 37, 37, 37,
	37, 38, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 43, 43, 43, 43, 43,
	43, 39, 39, 39, 39, 39, 39, 39, 44, 44,
	44, 49, 45, 45, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 48, 51, 54, 54, 52,
	52, 53, 55, 55, 50, 50, 41, 41, 41, 41,
	56, 56, 57, 57, 58, 58, 59, 59, 60, 61,
	61, 61, 62, 62, 62, 63, 63, 63, 64, 64,
	65, 65, 66, 66, 40, 40, 46, 46, 47, 47,
	67, 67, 68, 70, 70, 71, 71, 69, 69, 72,
	72, 72, 72, 72, 73, 73, 74, 74, 75, 75,
	76, 78, 84, 85, 81,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 5, 0, 2,
	2, 2, 1, 1, 1, 1, 12, 4, 3, 8,
	8, 8, 7, 3, 5, 8, 4, 6, 7, 4,
	5, 4, 5, 5, 3, 2, 0, 2, 0, 2,
	1, 2, 1, 1, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 1,
	1, 3, 3, 3, 3, 5, 5, 3, 0, 1,
	0, 1, 2, 1, 2, 2, 1, 2, 3, 2,
	3, 2, 2, 1, 3, 1, 3, 0, 5, 5,
	5, 1, 3, 0, 2, 1, 3, 3, 2, 3,
	3, 1, 1, 3, 3, 4, 3, 4, 3, 4,
	5, 6, 3, 2, 6, 1, 2, 1, 2, 1,
	2, 1, 1, 1, 1, 1, 1, 1, 3, 1,
	1, 3, 1, 3, 1, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	3, 4, 5, 4, 1, 1, 5, 0, 1, 1,
	2, 4, 0, 2, 1, 3, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 0, 3,
	1, 3, 0, 5, 2, 1, 1, 3, 3, 1,
	1, 3, 3, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 0, 1, 0, 1, 0, 2,
	1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -82, -1, -12, -2, -3, -4, -5, -6, -7,
	-8, -9, -10, -11, -15, -16, -14, 107, 9, 10,
	11, 12, 31, 92, 93, 95, 94, 96, 105, 44,
	106, -20, 5, 6, 7, 8, -13, -2, -3, -4,
	-5, -18, -83, -18, -18, -18, -18, 97, -74, 99,
	103, -69, 99, 101, 97, 97, 98, 99, 97, 50,
	50, -81, -2, 21, -21, 22, -19, -69, -33, -78,
	50, 13, -67, -68, -50, -76, -78, 50, -71, 102,
	98, -76, 50, 97, -76, -78, -70, 102, 50, -70,
	-78, -17, 13, -22, -23, 84, -24, -78, -37, -42,
	-38, 61, -84, -41, -50, -47, 82, 83, 88, -76,
	-48, -51, 57, 58, 24, 33, 47, 51, 52, 53,
	56, -49, 102, 63, 55, 27, 31, 90, -33, 48,
	67, 90, -78, 61, 50, -81, -78, -81, 100, -78,
	24, 46, -76, -36, 75, 14, -25, -26, -27, -28,
	-32, -49, -84, -78, 13, 48, -77, -76, 23, 90,
	60, 59, 74, -39, 77, 61, 75, 76, 62, 74,
	79, 78, 87, 82, 83, 84, 85, 86, 80, 81,
	67, 68, 69, 70, 71, 72, 73, -37, -42, -37,
	-2, -45, -42, -42, -84, -42, -42, -84, -84, -49,
	-84, -84, -54, -42, -33, -67, -78, -36, -68, -42,
	-76, -81, 24, -75, 104, -72, 95, 93, 30, 94,
	17, 50, -78, -78, -81, -42, -37, 48, -29, -30,
	-31, 36, 40, 42, 37, 38, 39, 43, -79, -78,
	23, -80, 23, -25, 90, -25, -23, -76, 84, -37,
	-37, -43, 56, 61, 57, 58, -42, -44, -84, -49,
	54, 77, 75, 76, 62, -42, -42, -42, -43, -42,
	-42, -42, -42, -42, -42, -42, -42, -42, -42, -85,
	49, -85, 48, -85, -42, -85, -22, 22, -22, -41,
	-52, -53, 64, -64, 31, -84, -36, -58, 17, 46,
	-76, -81, -73, 100, -26, -27, -26, -27, 36, 36,
	36, 41, 36, 41, 36, -30, -34, 44, 101, 45,
	-78, -78, -85, -78, -36, 56, 57, 58, -45, -44,
	-42, -42, -42, 60, -42, -85, -22, -85, 48, -55,
	-53, 66, -37, -40, 34, -2, -67, -65, -50, -58,
	-62, 19, 18, -78, -78, 46, 46, 36, 36, 98,
	98, 98, -56, 15, -85, 60, -42, -85, -41, 91,
	-42, 65, -66, 46, -46, -47, -66, -85, 48, -62,
	-42, -59, -60, -42, -81, -37, -37, -84, -84, -84,
	-57, 16, 18, -42, -85, -42, 28, 48, -50, 48,
	48, -61, 25, 26, -35, -76, -35, -35, -58, -37,
	-45, 29, -47, -42, -60, -85, 48, -85, -85, -62,
	11, -76, -63, 20, 32, -67, 11, 77, -76, -76,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 0, 46, 46,
	46, 46, 46, 226, 217, 0, 0, 0, 0, 0,
	234, 0, 50, 52, 53, 54, 21, 22, 23, 24,
	25, 55, 48, 217, 0, 0, 0, 215, 0, 0,
	227, 0, 0, 218, 0, 213, 0, 213, 0, 18,
	20, 45, 28, 51, 0, 56, 47, 0, 0, 95,
	231, 0, 33, 210, 0, 174, 0, -2, 0, 0,
	0, 234, 230, 0, 234, 0, 0, 0, 0, 0,
	44, 103, 0, 27, 57, 59, 64, 0, 62, 63,
	105, 0, 0, 144, 145, 146, 0, 0, 0, 174,
	0, 164, 111, 112, 0, 0, 232, 176, 177, 178,
	179, 209, 165, 167, 49, 0, 0, 0, 103, 0,
	0, 0, 234, 0, 228, 36, 0, 39, 0, 41,
	214, 0, 234, 16, 0, 0, 19, 67, 69, 70,
	80, 78, 0, 93, 0, 0, 60, 65, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	131, 132, 133, 134, 135, 136, 137, 108, 0, 0,
	0, 0, 142, 157, 0, 158, 159, 0, 0, 123,
	0, 0, 0, 168, 198, 103, 96, 184, 211, 212,
	175, 34, 216, 0, 0, 234, 224, 219, 220, 221,
	222, 223, 40, 42, 43, 17, 104, 0, 0, 0,
	0, 83, 0, 0, 86, 0, 0, 0, 97, 81,
	0, 0, 79, 0, 0, 103, 58, 66, 61, 106,
	107, 110, 125, 0, 127, 129, 113, 114, 0, 139,
	140, 0, 0, 0, 0, 116, 118, 0, 122, 147,
	148, 149, 150, 151, 152, 153, 154, 155, 156, 109,
	233, 141, 0, 208, 142, 160, 0, 0, 0, 0,
	172, 169, 0, 0, 0, 0, 184, 192, 0, 0,
	229, 37, 0, 225, 68, 74, 0, 77, 84, 85,
	87, 0, 89, 0, 91, 92, 71, 0, 0, 0,
	82, 72, 73, 94, 180, 126, 128, 130, 0, 115,
	117, 119, 0, 0, 143, 161, 0, 163, 0, 0,
	170, 0, 0, 202, 0, 205, 202, 0, 200, 192,
	32, 0, 0, 234, 38, 0, 0, 88, 90, 0,
	0, 0, 182, 0, 138, 0, 120, 162, 0, 166,
	173, 0, 29, 0, 204, 206, 30, 199, 0, 31,
	193, 185, 186, 189, 35, 75, 76, 0, 0, 0,
	184, 0, 0, 121, 124, 171, 0, 0, 201, 0,
	0, 188, 190, 191, 0, 101, 0, 0, 192, 183,
	181, 0, 207, 194, 187, 98, 0, 99, 100, 195,
	0, 102, 26, 0, 0, 203, 196, 0, 0, 197,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 86, 79, 3,
	47, 49, 84, 82, 48, 83, 90, 85, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	68, 67, 69, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 87, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 78, 3, 88,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 70, 71, 72, 73, 74, 75, 76, 77,
	80, 81, 89, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:166
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:173
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 16:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:191
		{
			yyVAL.statement = &Show{Key: string(yyDollar[2].bytes), From: yyDollar[3].tableExprs, Where: NewWhere(WhereStr, yyDollar[4].boolExpr)}
		}
	case 17:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:195
		{
			yyVAL.statement = &Show{Key: string(yyDollar[2].bytes), From: yyDollar[3].tableExprs, Like: &ComparisonExpr{Left: nil, Operator: LikeStr, Right: yyDollar[5].valExpr}}
		}
	case 18:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:200
		{
			yyVAL.tableExprs = nil
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:204
		{
			yyVAL.tableExprs = yyDollar[2].tableExprs
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:210
		{
			yyVAL.statement = &UseDB{DB: string(yyDollar[2].bytes)}
		}
	case 21:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:216
		{
			yyVAL.statement = &Explain{SQL: yyDollar[2].statement}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:222
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 26:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:232
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:236
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:240
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 29:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:246
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 30:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:250
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 31:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:262
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 32:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:268
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:274
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:280
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].sqlID}
		}
	case 35:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:284
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].sqlID, NewName: yyDollar[7].sqlID}
		}
	case 36:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:289
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: SQLName(yyDollar[3].sqlID)}
		}
	case 37:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].sqlID, NewName: yyDollar[4].sqlID}
		}
	case 38:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:299
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].sqlID, NewName: yyDollar[7].sqlID}
		}
	case 39:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:304
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: SQLName(yyDollar[3].sqlID), NewName: SQLName(yyDollar[3].sqlID)}
		}
	case 40:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:310
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:316
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].sqlID}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:320
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:325
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: SQLName(yyDollar[4].sqlID)}
		}
	case 44:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:331
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].sqlID, NewName: yyDollar[3].sqlID}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:337
		{
			yyVAL.statement = &Other{}
		}
	case 46:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:342
		{
			setAllowComments(yylex, true)
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:346
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:352
		{
			yyVAL.bytes2 = nil
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:356
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:362
		{
			yyVAL.str = UnionStr
		}
	case 51:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:366
		{
			yyVAL.str = UnionAllStr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:370
		{
			yyVAL.str = SetMinusStr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:374
		{
			yyVAL.str = ExceptStr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:378
		{
			yyVAL.str = IntersectStr
		}
	case 55:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:383
		{
			yyVAL.str = ""
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = DistinctStr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:393
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:397
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:403
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:407
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].sqlID}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:411
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].sqlID}
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:417
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 64:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:426
		{
			yyVAL.sqlID = ""
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:434
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:440
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:444
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:454
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].sqlID, Hints: yyDollar[3].indexHints}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:458
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].sqlID}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:462
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:475
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 75:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:492
		{
			yyVAL.empty = struct{}{}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:494
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:497
		{
			yyVAL.sqlID = ""
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:501
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:505
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:511
		{
			yyVAL.str = JoinStr
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = JoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = JoinStr
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = StraightJoinStr
		}
	case 87:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:529
		{
			yyVAL.str = LeftJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = LeftJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:537
		{
			yyVAL.str = RightJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:541
		{
			yyVAL.str = RightJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:547
		{
			yyVAL.str = NaturalJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:551
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:561
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].sqlID}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:565
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:571
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].sqlID}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:575
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:580
		{
			yyVAL.indexHints = nil
		}
	case 98:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:584
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:588
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:592
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:598
		{
			yyVAL.sqlIDs = []SQLName{yyDollar[1].sqlID}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:602
		{
			yyVAL.sqlIDs = append(yyDollar[1].sqlIDs, yyDollar[3].sqlID)
		}
	case 103:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:607
		{
			yyVAL.boolExpr = nil
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:611
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:618
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:630
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:634
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:640
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 115:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:684
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:688
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 124:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:692
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.str = IsNullStr
		}
	case 126:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = IsNotNullStr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = IsTrueStr
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = IsNotTrueStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = IsFalseStr
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = IsNotFalseStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:724
		{
			yyVAL.str = EqualStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:728
		{
			yyVAL.str = LessThanStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:732
		{
			yyVAL.str = GreaterThanStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.str = LessEqualStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.str = GreaterEqualStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:744
		{
			yyVAL.str = NotEqualStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:754
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:768
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 142:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 144:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:784
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:788
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:792
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:796
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:800
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:804
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:808
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:812
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:816
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:820
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:824
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:836
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:844
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				// Handle double negative
				if num[0] == '-' {
					yyVAL.valExpr = num[1:]
				} else {
					yyVAL.valExpr = append(NumVal("-"), num...)
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:857
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:861
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID)}
		}
	case 161:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:865
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Exprs: yyDollar[3].selectExprs}
		}
	case 162:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].str, Exprs: yyDollar[3].selectExprs}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:883
		{
			yyVAL.str = "if"
		}
	case 166:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:889
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 167:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:894
		{
			yyVAL.valExpr = nil
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:904
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 170:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:908
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 171:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:914
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = nil
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:929
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].sqlID}
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:933
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:939
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:956
		{
			yyVAL.valExprs = nil
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:960
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:965
		{
			yyVAL.boolExpr = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:969
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:974
		{
			yyVAL.orderBy = nil
		}
	case 185:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:978
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:984
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:988
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:994
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:999
		{
			yyVAL.str = AscScr
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.str = AscScr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.str = DescScr
		}
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1012
		{
			yyVAL.limit = nil
		}
	case 193:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 194:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.str = ""
		}
	case 196:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.str = ForUpdateStr
		}
	case 197:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1033
		{
			if yyDollar[3].sqlID != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].sqlID != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 198:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.columns = nil
		}
	case 199:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1056
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 202:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.updateExprs = nil
		}
	case 203:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 207:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1095
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 211:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.str = ""
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.str = IgnoreStr
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1149
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1154
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1157
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1159
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.sqlID = SQLName(strings.ToLower(string(yyDollar[1].bytes)))
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1169
		{
			yyVAL.sqlID = SQLName(yyDollar[1].bytes)
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1175
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1184
		{
			decNesting(yylex)
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1189
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
