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
const DATABASE = 57424
const INDEX = 57425
const VIEW = 57426
const TO = 57427
const IGNORE = 57428
const IF = 57429
const UNIQUE = 57430
const USING = 57431
const SHOW = 57432
const DESCRIBE = 57433
const EXPLAIN = 57434

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
	"DATABASE",
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
	-1, 78,
	90, 233,
	-2, 232,
}

const yyNprod = 237
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 809

var yyAct = [...]int{

	111, 195, 192, 356, 106, 388, 107, 410, 303, 73,
	378, 297, 262, 62, 194, 4, 234, 152, 147, 95,
	105, 151, 256, 74, 69, 219, 89, 47, 50, 284,
	49, 96, 37, 150, 51, 323, 325, 80, 53, 119,
	54, 54, 78, 120, 121, 122, 63, 76, 123, 309,
	82, 75, 142, 198, 86, 126, 56, 367, 57, 58,
	366, 365, 81, 85, 59, 55, 375, 249, 101, 225,
	179, 180, 181, 176, 108, 109, 83, 163, 134, 130,
	110, 176, 223, 100, 177, 178, 179, 180, 181, 176,
	112, 269, 146, 324, 433, 125, 138, 131, 139, 161,
	141, 166, 226, 133, 267, 268, 266, 196, 298, 136,
	253, 197, 199, 200, 127, 149, 165, 164, 164, 104,
	174, 182, 183, 177, 178, 179, 180, 181, 176, 208,
	76, 166, 166, 76, 75, 215, 214, 75, 298, 210,
	347, 331, 332, 333, 285, 222, 224, 221, 156, 216,
	212, 230, 88, 209, 165, 164, 213, 104, 104, 18,
	229, 119, 101, 252, 215, 201, 202, 203, 265, 166,
	261, 205, 206, 270, 271, 272, 148, 274, 275, 276,
	277, 278, 279, 280, 281, 282, 283, 191, 193, 83,
	248, 251, 250, 165, 164, 119, 273, 119, 71, 377,
	71, 289, 78, 104, 101, 71, 101, 137, 166, 90,
	156, 91, 156, 104, 32, 33, 34, 35, 104, 104,
	306, 291, 263, 294, 286, 288, 61, 295, 60, 302,
	285, 290, 293, 231, 307, 257, 259, 260, 422, 285,
	258, 245, 18, 19, 20, 21, 162, 132, 254, 255,
	379, 311, 132, 313, 310, 104, 312, 104, 285, 321,
	384, 285, 158, 301, 22, 334, 289, 287, 71, 330,
	336, 337, 338, 83, 159, 285, 406, 29, 328, 335,
	120, 121, 122, 287, 285, 123, 156, 156, 156, 156,
	340, 232, 285, 403, 344, 101, 124, 159, 119, 300,
	232, 76, 76, 379, 149, 75, 354, 361, 346, 149,
	352, 355, 342, 305, 351, 119, 145, 318, 52, 364,
	263, 341, 319, 363, 343, 23, 24, 26, 25, 27,
	182, 183, 177, 178, 179, 180, 181, 176, 232, 28,
	30, 17, 372, 132, 316, 320, 104, 240, 241, 317,
	376, 315, 104, 314, 18, 429, 129, 381, 417, 385,
	386, 389, 68, 382, 370, 374, 402, 430, 128, 64,
	217, 144, 373, 390, 399, 247, 66, 357, 398, 350,
	401, 77, 348, 383, 358, 76, 304, 397, 7, 404,
	369, 155, 149, 6, 411, 411, 411, 94, 72, 5,
	416, 289, 412, 413, 400, 414, 40, 432, 419, 389,
	418, 39, 420, 426, 204, 104, 104, 38, 425, 393,
	394, 395, 18, 427, 42, 1, 70, 76, 246, 243,
	160, 75, 84, 218, 434, 435, 431, 87, 48, 308,
	421, 92, 423, 424, 220, 391, 392, 99, 79, 353,
	299, 104, 41, 155, 70, 155, 236, 239, 240, 241,
	237, 135, 238, 242, 428, 264, 362, 140, 407, 387,
	143, 396, 43, 44, 45, 46, 157, 169, 172, 368,
	345, 415, 207, 184, 185, 186, 187, 188, 189, 190,
	173, 170, 171, 168, 175, 174, 182, 183, 177, 178,
	179, 180, 181, 176, 296, 18, 18, 19, 20, 21,
	70, 114, 211, 113, 380, 408, 409, 32, 33, 34,
	35, 349, 167, 102, 227, 322, 292, 228, 117, 155,
	155, 155, 155, 154, 235, 233, 244, 118, 157, 153,
	157, 99, 98, 119, 65, 31, 78, 120, 121, 122,
	67, 119, 123, 285, 78, 120, 121, 122, 93, 126,
	123, 115, 116, 264, 15, 103, 14, 126, 175, 174,
	182, 183, 177, 178, 179, 180, 181, 176, 108, 109,
	16, 36, 3, 99, 110, 99, 108, 109, 97, 13,
	12, 11, 110, 10, 112, 9, 8, 2, 0, 125,
	0, 0, 112, 0, 0, 0, 0, 125, 0, 117,
	0, 0, 0, 0, 157, 157, 157, 157, 118, 0,
	0, 0, 405, 0, 0, 0, 0, 326, 327, 0,
	0, 329, 119, 0, 0, 78, 120, 121, 122, 0,
	0, 123, 115, 116, 0, 0, 103, 0, 126, 0,
	0, 18, 175, 174, 182, 183, 177, 178, 179, 180,
	181, 176, 0, 0, 0, 0, 117, 108, 109, 97,
	0, 0, 0, 110, 99, 118, 0, 0, 0, 0,
	0, 0, 0, 112, 0, 0, 0, 359, 125, 119,
	360, 0, 78, 120, 121, 122, 0, 0, 123, 115,
	116, 0, 0, 103, 0, 126, 0, 0, 0, 117,
	236, 239, 240, 241, 237, 0, 238, 242, 118, 0,
	0, 0, 0, 0, 108, 109, 0, 0, 0, 0,
	110, 371, 119, 0, 0, 78, 120, 121, 122, 0,
	112, 123, 115, 116, 0, 125, 103, 0, 126, 175,
	174, 182, 183, 177, 178, 179, 180, 181, 176, 0,
	0, 0, 0, 0, 0, 0, 0, 108, 109, 0,
	0, 339, 0, 110, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 112, 0, 0, 0, 0, 125, 175,
	174, 182, 183, 177, 178, 179, 180, 181, 176, 175,
	174, 182, 183, 177, 178, 179, 180, 181, 176,
}
var yyPact = [...]int{

	233, -1000, -1000, -1000, 512, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 497, -1000, -1000,
	-1000, -1000, -1000, -70, -62, -32, -41, -33, 178, 176,
	-1000, 413, 348, -1000, -1000, -1000, -1000, 512, -1000, -1000,
	-1000, 354, -1000, -61, 155, 385, 152, -66, -37, 139,
	155, -1000, -34, 139, -1000, 155, -77, 159, -77, 155,
	384, -1000, -1000, -1000, -1000, 585, -1000, 59, 341, 325,
	-11, -1000, 155, 199, -1000, 36, -1000, -12, -1000, 155,
	48, 157, -1000, -1000, -1000, 155, -1000, -49, 155, 347,
	270, 139, -1000, 101, 148, 249, -1000, -1000, 223, -13,
	57, 416, -1000, 685, 642, -1000, -1000, -1000, -8, -8,
	-8, 251, 251, 251, -1000, -1000, -1000, 251, 251, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -8, -1000, 155, 152,
	155, 378, 152, -8, 139, -1000, 346, -80, -1000, -1000,
	52, -1000, 155, -1000, -1000, 155, -1000, -1000, -8, 685,
	252, 674, -1000, -1000, 218, 352, 150, -23, 148, 585,
	-1000, -1000, 139, 26, 685, 685, 179, -8, 114, 29,
	-8, -8, -8, 179, -8, -8, -8, -8, -8, -8,
	-8, -8, -8, -8, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 27, 416, 95, 209, 235, 416, -1000, 496, -1000,
	-1000, 504, 181, 585, -1000, 413, 229, 44, 721, 268,
	295, -1000, 369, -1000, 721, -1000, -1000, -1000, 267, 139,
	-1000, -52, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	721, 57, 148, 148, 148, 148, -1000, 317, 315, -1000,
	308, 281, 309, -9, -1000, 155, 155, -1000, 243, 155,
	290, -1000, -1000, -1000, 27, 58, -1000, -1000, 85, -1000,
	-1000, 721, -1000, 496, -1000, -1000, 114, -8, -8, -8,
	721, 721, 711, -1000, 250, 41, -1000, -14, -14, -6,
	-6, -6, 2, 2, -1000, -1000, -1000, -8, -1000, 721,
	-1000, 226, 585, -1000, 226, 246, 74, -1000, 685, 345,
	152, 152, 369, 358, 366, 155, -1000, -1000, 155, -1000,
	674, 261, 420, -1000, -1000, -1000, -1000, 287, -1000, 283,
	-1000, -1000, -1000, -38, -39, -42, -1000, -1000, -1000, -1000,
	375, -1000, -1000, -1000, 235, -1000, 721, 721, 671, -8,
	721, -1000, 226, -1000, 229, -25, -1000, -8, 134, 257,
	251, 512, 204, 212, -1000, 358, -1000, -8, -8, -1000,
	-1000, 685, 685, -1000, -1000, 251, 251, 251, 371, 360,
	-1000, -8, 721, -1000, 181, -1000, 721, -8, -1000, 338,
	245, -1000, -1000, -1000, 152, -1000, 574, 228, -1000, 490,
	-1000, 57, 57, 139, 139, 139, 369, 685, -8, 721,
	-1000, 721, 329, 251, -1000, -8, -8, -1000, -1000, -1000,
	190, -1000, 190, 190, 358, 57, 219, 402, -1000, 721,
	-1000, -1000, 139, -1000, -1000, 335, 152, -1000, -1000, 396,
	17, 199, -1000, 139, 139, -1000,
}
var yyPgo = [...]int{

	0, 597, 14, 399, 393, 388, 596, 595, 593, 591,
	590, 589, 582, 581, 580, 566, 564, 558, 452, 550,
	545, 544, 19, 31, 542, 33, 21, 17, 539, 535,
	16, 534, 533, 24, 525, 7, 18, 83, 523, 522,
	521, 20, 2, 22, 12, 1, 514, 6, 513, 296,
	4, 511, 504, 11, 482, 480, 479, 471, 8, 469,
	5, 468, 3, 464, 450, 449, 10, 9, 23, 318,
	152, 448, 444, 439, 438, 433, 0, 430, 381, 429,
	428, 13, 425, 424, 53, 29,
}
var yyR1 = [...]int{

	0, 82, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 15, 15, 17, 17,
	16, 12, 13, 13, 13, 13, 2, 2, 2, 3,
	3, 4, 5, 6, 7, 7, 7, 7, 8, 8,
	8, 9, 10, 10, 10, 11, 14, 83, 18, 19,
	19, 20, 20, 20, 20, 20, 21, 21, 22, 22,
	23, 23, 23, 24, 24, 77, 77, 77, 25, 25,
	26, 26, 27, 27, 27, 28, 28, 28, 28, 80,
	80, 79, 79, 79, 29, 29, 29, 29, 30, 30,
	30, 30, 31, 31, 32, 32, 33, 33, 34, 34,
	34, 34, 35, 35, 36, 36, 37, 37, 37, 37,
	37, 37, 38, 38, 38, 38, 38, 38, 38, 38,
	38, 38, 38, 38, 38, 38, 43, 43, 43, 43,
	43, 43, 39, 39, 39, 39, 39, 39, 39, 44,
	44, 44, 49, 45, 45, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 48, 51, 54,
	54, 52, 52, 53, 55, 55, 50, 50, 41, 41,
	41, 41, 56, 56, 57, 57, 58, 58, 59, 59,
	60, 61, 61, 61, 62, 62, 62, 63, 63, 63,
	64, 64, 65, 65, 66, 66, 40, 40, 46, 46,
	47, 47, 67, 67, 68, 70, 70, 71, 71, 69,
	69, 72, 72, 72, 72, 72, 73, 73, 74, 74,
	75, 75, 76, 78, 84, 85, 81,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 4, 5, 0, 2,
	2, 2, 1, 1, 1, 1, 12, 4, 3, 8,
	8, 8, 7, 3, 5, 8, 4, 4, 6, 7,
	4, 5, 4, 5, 5, 3, 2, 0, 2, 0,
	2, 1, 2, 1, 1, 1, 0, 1, 1, 3,
	1, 2, 3, 1, 1, 0, 1, 2, 1, 3,
	1, 1, 3, 3, 3, 3, 5, 5, 3, 0,
	1, 0, 1, 2, 1, 2, 2, 1, 2, 3,
	2, 3, 2, 2, 1, 3, 1, 3, 0, 5,
	5, 5, 1, 3, 0, 2, 1, 3, 3, 2,
	3, 3, 1, 1, 3, 3, 4, 3, 4, 3,
	4, 5, 6, 3, 2, 6, 1, 2, 1, 2,
	1, 2, 1, 1, 1, 1, 1, 1, 1, 3,
	1, 1, 3, 1, 3, 1, 1, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 3, 3, 4, 5, 4, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 2, 1, 1, 3,
	3, 1, 1, 3, 3, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 0, 1, 0, 1,
	0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -82, -1, -12, -2, -3, -4, -5, -6, -7,
	-8, -9, -10, -11, -15, -16, -14, 108, 9, 10,
	11, 12, 31, 92, 93, 95, 94, 96, 106, 44,
	107, -20, 5, 6, 7, 8, -13, -2, -3, -4,
	-5, -18, -83, -18, -18, -18, -18, 97, -74, 100,
	98, 104, -69, 100, 102, 97, 97, 99, 100, 97,
	50, 50, -81, -2, 21, -21, 22, -19, -69, -33,
	-78, 50, 13, -67, -68, -50, -76, -78, 50, -71,
	103, 99, -76, 50, -78, 97, -76, -78, -70, 103,
	50, -70, -78, -17, 13, -22, -23, 84, -24, -78,
	-37, -42, -38, 61, -84, -41, -50, -47, 82, 83,
	88, -76, 98, -48, -51, 57, 58, 24, 33, 47,
	51, 52, 53, 56, -49, 103, 63, 55, 27, 31,
	90, -33, 48, 67, 90, -78, 61, 50, -81, -81,
	-78, -81, 101, -78, 24, 46, -76, -36, 75, 14,
	-25, -26, -27, -28, -32, -49, -84, -78, 13, 48,
	-77, -76, 23, 90, 60, 59, 74, -39, 77, 61,
	75, 76, 62, 74, 79, 78, 87, 82, 83, 84,
	85, 86, 80, 81, 67, 68, 69, 70, 71, 72,
	73, -37, -42, -37, -2, -45, -42, -42, -84, -42,
	-42, -84, -84, -84, -49, -84, -84, -54, -42, -33,
	-67, -78, -36, -68, -42, -76, -81, 24, -75, 105,
	-72, 95, 93, 30, 94, 17, 50, -78, -78, -81,
	-42, -37, 48, -29, -30, -31, 36, 40, 42, 37,
	38, 39, 43, -79, -78, 23, -80, 23, -25, 90,
	-25, -23, -76, 84, -37, -37, -43, 56, 61, 57,
	58, -42, -44, -84, -49, 54, 77, 75, 76, 62,
	-42, -42, -42, -43, -42, -42, -42, -42, -42, -42,
	-42, -42, -42, -42, -85, 49, -85, 48, -85, -42,
	-85, -22, 22, -85, -22, -41, -52, -53, 64, -64,
	31, -84, -36, -58, 17, 46, -76, -81, -73, 101,
	-26, -27, -26, -27, 36, 36, 36, 41, 36, 41,
	36, -30, -34, 44, 102, 45, -78, -78, -85, -78,
	-36, 56, 57, 58, -45, -44, -42, -42, -42, 60,
	-42, -85, -22, -85, 48, -55, -53, 66, -37, -40,
	34, -2, -67, -65, -50, -58, -62, 19, 18, -78,
	-78, 46, 46, 36, 36, 99, 99, 99, -56, 15,
	-85, 60, -42, -85, -41, 91, -42, 65, -66, 46,
	-46, -47, -66, -85, 48, -62, -42, -59, -60, -42,
	-81, -37, -37, -84, -84, -84, -57, 16, 18, -42,
	-85, -42, 28, 48, -50, 48, 48, -61, 25, 26,
	-35, -76, -35, -35, -58, -37, -45, 29, -47, -42,
	-60, -85, 48, -85, -85, -62, 11, -76, -63, 20,
	32, -67, 11, 77, -76, -76,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 0, 47, 47,
	47, 47, 47, 228, 219, 0, 0, 0, 0, 0,
	236, 0, 51, 53, 54, 55, 21, 22, 23, 24,
	25, 56, 49, 219, 0, 0, 0, 217, 0, 0,
	0, 229, 0, 0, 220, 0, 215, 0, 215, 0,
	18, 20, 46, 28, 52, 0, 57, 48, 0, 0,
	96, 233, 0, 33, 212, 0, 176, 0, -2, 0,
	0, 0, 236, 232, 236, 0, 236, 0, 0, 0,
	0, 0, 45, 104, 0, 27, 58, 60, 65, 0,
	63, 64, 106, 0, 0, 145, 146, 147, 0, 0,
	0, 176, 0, 0, 166, 112, 113, 0, 0, 234,
	178, 179, 180, 181, 211, 167, 169, 50, 0, 0,
	0, 104, 0, 0, 0, 236, 0, 230, 36, 37,
	0, 40, 0, 42, 216, 0, 236, 16, 0, 0,
	19, 68, 70, 71, 81, 79, 0, 94, 0, 0,
	61, 66, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 132, 133, 134, 135, 136, 137,
	138, 109, 0, 0, 0, 0, 143, 158, 0, 159,
	160, 0, 0, 0, 124, 0, 0, 0, 170, 200,
	104, 97, 186, 213, 214, 177, 34, 218, 0, 0,
	236, 226, 221, 222, 223, 224, 225, 41, 43, 44,
	17, 105, 0, 0, 0, 0, 84, 0, 0, 87,
	0, 0, 0, 98, 82, 0, 0, 80, 0, 0,
	104, 59, 67, 62, 107, 108, 111, 126, 0, 128,
	130, 114, 115, 0, 140, 141, 0, 0, 0, 0,
	117, 119, 0, 123, 148, 149, 150, 151, 152, 153,
	154, 155, 156, 157, 110, 235, 142, 0, 210, 143,
	161, 0, 0, 162, 0, 0, 174, 171, 0, 0,
	0, 0, 186, 194, 0, 0, 231, 38, 0, 227,
	69, 75, 0, 78, 85, 86, 88, 0, 90, 0,
	92, 93, 72, 0, 0, 0, 83, 73, 74, 95,
	182, 127, 129, 131, 0, 116, 118, 120, 0, 0,
	144, 163, 0, 165, 0, 0, 172, 0, 0, 204,
	0, 207, 204, 0, 202, 194, 32, 0, 0, 236,
	39, 0, 0, 89, 91, 0, 0, 0, 184, 0,
	139, 0, 121, 164, 0, 168, 175, 0, 29, 0,
	206, 208, 30, 201, 0, 31, 195, 187, 188, 191,
	35, 76, 77, 0, 0, 0, 186, 0, 0, 122,
	125, 173, 0, 0, 203, 0, 0, 190, 192, 193,
	0, 102, 0, 0, 194, 185, 183, 0, 209, 196,
	189, 99, 0, 100, 101, 197, 0, 103, 26, 0,
	0, 205, 198, 0, 0, 199,
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
	108,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:293
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[3].sqlID, IsDB: true}
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].sqlID, NewName: yyDollar[4].sqlID}
		}
	case 39:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:303
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].sqlID, NewName: yyDollar[7].sqlID}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:308
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: SQLName(yyDollar[3].sqlID), NewName: SQLName(yyDollar[3].sqlID)}
		}
	case 41:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:314
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 42:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:320
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].sqlID}
		}
	case 43:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:324
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 44:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:329
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: SQLName(yyDollar[4].sqlID)}
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:335
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].sqlID, NewName: yyDollar[3].sqlID}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &Other{}
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:346
		{
			setAllowComments(yylex, true)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:350
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:356
		{
			yyVAL.bytes2 = nil
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:360
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:366
		{
			yyVAL.str = UnionStr
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:370
		{
			yyVAL.str = UnionAllStr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:374
		{
			yyVAL.str = SetMinusStr
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:378
		{
			yyVAL.str = ExceptStr
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = IntersectStr
		}
	case 56:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = ""
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = DistinctStr
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:397
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:401
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 61:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:411
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].sqlID}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:415
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].sqlID}
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:421
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:425
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:430
		{
			yyVAL.sqlID = ""
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:434
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 67:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:438
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:444
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:448
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:458
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].sqlID, Hints: yyDollar[3].indexHints}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:462
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].sqlID}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:466
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:479
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 76:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 77:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 79:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:496
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:498
		{
			yyVAL.empty = struct{}{}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:501
		{
			yyVAL.sqlID = ""
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:505
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:509
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
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
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = JoinStr
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:527
		{
			yyVAL.str = StraightJoinStr
		}
	case 88:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = LeftJoinStr
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:537
		{
			yyVAL.str = LeftJoinStr
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:541
		{
			yyVAL.str = RightJoinStr
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:545
		{
			yyVAL.str = RightJoinStr
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:551
		{
			yyVAL.str = NaturalJoinStr
		}
	case 93:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:555
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:565
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].sqlID}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:569
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:575
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].sqlID}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:579
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 98:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:584
		{
			yyVAL.indexHints = nil
		}
	case 99:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:588
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 100:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:592
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 101:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:596
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:602
		{
			yyVAL.sqlIDs = []SQLName{yyDollar[1].sqlID}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:606
		{
			yyVAL.sqlIDs = append(yyDollar[1].sqlIDs, yyDollar[3].sqlID)
		}
	case 104:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:611
		{
			yyVAL.boolExpr = nil
		}
	case 105:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:615
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 109:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:630
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:634
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:638
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 116:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 118:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 120:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 121:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 122:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:684
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 123:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:688
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:692
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 125:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:696
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = IsNullStr
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = IsNotNullStr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = IsTrueStr
		}
	case 129:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = IsNotTrueStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = IsFalseStr
		}
	case 131:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:722
		{
			yyVAL.str = IsNotFalseStr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:728
		{
			yyVAL.str = EqualStr
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:732
		{
			yyVAL.str = LessThanStr
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.str = GreaterThanStr
		}
	case 135:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.str = LessEqualStr
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:744
		{
			yyVAL.str = GreaterEqualStr
		}
	case 137:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.str = NotEqualStr
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:758
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:766
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:772
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 143:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:782
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:788
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:792
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:796
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:800
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:804
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:808
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:812
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 152:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:816
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 153:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:820
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:824
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 157:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:836
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:840
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:848
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
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:861
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:865
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID)}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = &FuncExpr{Name: "database"}
		}
	case 163:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Exprs: yyDollar[3].selectExprs}
		}
	case 164:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].str, Exprs: yyDollar[3].selectExprs}
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:891
		{
			yyVAL.str = "if"
		}
	case 168:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:897
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 169:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = nil
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:906
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:912
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 172:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:916
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 173:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:922
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = nil
		}
	case 175:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:931
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:937
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].sqlID}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:941
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:959
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:964
		{
			yyVAL.valExprs = nil
		}
	case 183:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:968
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 184:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:973
		{
			yyVAL.boolExpr = nil
		}
	case 185:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:977
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:982
		{
			yyVAL.orderBy = nil
		}
	case 187:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:986
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:992
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 189:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:996
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1002
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 191:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.str = AscScr
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.str = AscScr
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.str = DescScr
		}
	case 194:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.limit = nil
		}
	case 195:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.str = ""
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.str = ForUpdateStr
		}
	case 199:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1041
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
	case 200:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.columns = nil
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1058
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 203:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 204:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.updateExprs = nil
		}
	case 205:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 206:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 214:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 215:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.str = ""
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.str = IgnoreStr
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1150
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1155
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1157
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1162
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1165
		{
			yyVAL.empty = struct{}{}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.empty = struct{}{}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1171
		{
			yyVAL.sqlID = SQLName(strings.ToLower(string(yyDollar[1].bytes)))
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1177
		{
			yyVAL.sqlID = SQLName(yyDollar[1].bytes)
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1183
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 235:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1192
		{
			decNesting(yylex)
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1197
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
