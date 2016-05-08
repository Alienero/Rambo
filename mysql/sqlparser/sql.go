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
	90, 232,
	-2, 231,
}

const yyNprod = 236
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 786

var yyAct = [...]int{

	111, 107, 191, 385, 106, 232, 62, 353, 105, 73,
	300, 375, 294, 194, 193, 4, 151, 260, 407, 254,
	150, 74, 95, 146, 217, 89, 80, 282, 69, 320,
	322, 149, 37, 53, 96, 54, 47, 50, 54, 49,
	364, 100, 290, 51, 116, 306, 63, 76, 141, 223,
	82, 75, 363, 117, 86, 56, 362, 57, 58, 81,
	85, 59, 221, 55, 372, 247, 162, 118, 101, 283,
	78, 119, 120, 121, 133, 175, 122, 114, 115, 129,
	430, 103, 224, 125, 163, 164, 163, 321, 148, 137,
	83, 138, 145, 140, 178, 179, 180, 175, 165, 160,
	165, 130, 108, 109, 97, 165, 132, 195, 110, 295,
	135, 196, 198, 199, 181, 182, 176, 177, 178, 179,
	180, 175, 197, 124, 251, 220, 222, 219, 206, 76,
	126, 83, 76, 75, 213, 212, 75, 295, 208, 344,
	118, 214, 78, 283, 267, 190, 192, 263, 243, 147,
	228, 88, 227, 211, 210, 71, 207, 265, 266, 264,
	136, 101, 250, 213, 161, 255, 257, 258, 118, 259,
	256, 71, 268, 269, 270, 71, 272, 273, 274, 275,
	276, 277, 278, 279, 280, 281, 90, 246, 104, 248,
	229, 83, 271, 249, 176, 177, 178, 179, 180, 175,
	287, 405, 406, 101, 101, 252, 253, 328, 329, 330,
	91, 18, 61, 292, 164, 163, 60, 155, 303, 131,
	374, 284, 286, 289, 291, 304, 104, 104, 288, 165,
	419, 283, 299, 148, 200, 201, 119, 120, 121, 203,
	204, 122, 381, 283, 158, 283, 318, 285, 308, 118,
	310, 307, 71, 309, 174, 173, 181, 182, 176, 177,
	178, 179, 180, 175, 287, 285, 283, 230, 333, 334,
	335, 104, 327, 157, 325, 331, 230, 283, 155, 403,
	155, 104, 332, 376, 148, 131, 104, 104, 337, 400,
	261, 18, 341, 101, 230, 123, 297, 118, 76, 76,
	376, 358, 75, 351, 302, 144, 343, 349, 158, 315,
	352, 348, 118, 339, 316, 283, 313, 338, 131, 340,
	361, 314, 52, 104, 104, 164, 163, 360, 312, 118,
	298, 311, 78, 119, 120, 121, 426, 345, 122, 369,
	165, 317, 128, 238, 239, 125, 414, 373, 427, 378,
	371, 18, 399, 155, 155, 155, 155, 383, 386, 367,
	382, 379, 127, 387, 108, 109, 68, 370, 215, 143,
	110, 396, 32, 33, 34, 35, 347, 398, 380, 245,
	66, 64, 76, 354, 395, 124, 401, 261, 355, 301,
	154, 408, 408, 408, 77, 394, 7, 6, 287, 397,
	388, 389, 415, 5, 411, 416, 386, 417, 366, 413,
	409, 410, 202, 104, 40, 39, 283, 148, 104, 422,
	424, 38, 94, 72, 76, 429, 423, 18, 75, 42,
	1, 431, 432, 428, 244, 418, 412, 420, 421, 70,
	241, 159, 216, 118, 48, 84, 78, 119, 120, 121,
	87, 154, 122, 154, 92, 18, 19, 20, 21, 125,
	99, 305, 218, 262, 79, 350, 296, 70, 425, 404,
	402, 384, 393, 365, 134, 41, 342, 205, 108, 109,
	139, 104, 104, 142, 110, 390, 391, 392, 293, 156,
	18, 19, 20, 21, 113, 43, 44, 45, 46, 124,
	174, 173, 181, 182, 176, 177, 178, 179, 180, 175,
	112, 377, 22, 32, 33, 34, 35, 104, 346, 166,
	102, 319, 70, 153, 209, 29, 154, 154, 154, 154,
	233, 234, 237, 238, 239, 235, 225, 236, 240, 226,
	231, 359, 152, 368, 98, 65, 31, 67, 242, 93,
	156, 15, 156, 99, 14, 16, 36, 3, 13, 12,
	262, 174, 173, 181, 182, 176, 177, 178, 179, 180,
	175, 116, 11, 23, 24, 26, 25, 27, 10, 9,
	117, 8, 2, 0, 0, 0, 0, 28, 30, 17,
	0, 0, 0, 0, 118, 99, 99, 78, 119, 120,
	121, 0, 0, 122, 114, 115, 0, 0, 103, 0,
	125, 0, 234, 237, 238, 239, 235, 18, 236, 240,
	0, 0, 0, 0, 0, 156, 156, 156, 156, 108,
	109, 97, 116, 0, 0, 110, 0, 0, 323, 324,
	0, 117, 326, 0, 0, 0, 0, 0, 0, 0,
	124, 0, 0, 0, 0, 118, 0, 0, 78, 119,
	120, 121, 0, 0, 122, 114, 115, 0, 0, 103,
	0, 125, 0, 0, 0, 116, 0, 0, 0, 0,
	0, 0, 0, 0, 117, 99, 0, 0, 0, 0,
	108, 109, 0, 0, 0, 0, 110, 356, 118, 336,
	357, 78, 119, 120, 121, 0, 0, 122, 114, 115,
	0, 124, 103, 0, 125, 0, 0, 174, 173, 181,
	182, 176, 177, 178, 179, 180, 175, 0, 0, 0,
	0, 0, 0, 108, 109, 0, 0, 0, 0, 110,
	173, 181, 182, 176, 177, 178, 179, 180, 175, 168,
	171, 0, 0, 0, 124, 183, 184, 185, 186, 187,
	188, 189, 172, 169, 170, 167, 174, 173, 181, 182,
	176, 177, 178, 179, 180, 175, 174, 173, 181, 182,
	176, 177, 178, 179, 180, 175,
}
var yyPact = [...]int{

	481, -1000, -1000, -1000, 508, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 446, -1000, -1000,
	-1000, -1000, -1000, -61, -67, -34, -42, -36, 166, 162,
	-1000, 418, 360, -1000, -1000, -1000, -1000, 508, -1000, -1000,
	-1000, 358, -1000, -64, 105, 410, 92, -77, -40, 81,
	105, -1000, -37, 81, -1000, 105, -78, 136, -78, 105,
	409, -1000, -1000, -1000, -1000, 547, -1000, 75, 335, 311,
	-11, -1000, 105, 171, -1000, 39, -1000, -16, -1000, 105,
	49, 110, -1000, -1000, -1000, 105, -1000, -53, 105, 345,
	259, 81, -1000, 74, 121, 260, -1000, -1000, 141, -24,
	26, 688, -1000, 651, 608, -1000, -1000, -1000, 396, 396,
	396, 250, 250, -1000, -1000, -1000, 250, 250, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 396, -1000, 105, 92, 105,
	403, 92, 396, 81, -1000, 344, -81, -1000, -1000, 32,
	-1000, 105, -1000, -1000, 105, -1000, -1000, 396, 651, 246,
	576, -1000, -1000, 125, 356, 202, -25, 121, 547, -1000,
	-1000, 81, 40, 651, 651, 109, 396, 93, 82, 396,
	396, 396, 109, 396, 396, 396, 396, 396, 396, 396,
	396, 396, 396, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	31, 688, 266, 367, 217, 688, -1000, 282, -1000, -1000,
	20, 547, -1000, 418, 185, 45, 698, 265, 270, -1000,
	372, -1000, 698, -1000, -1000, -1000, 258, 81, -1000, -56,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 698, 26,
	121, 121, 121, 121, -1000, 295, 292, -1000, 280, 273,
	305, -15, -1000, 105, 105, -1000, 228, 105, 219, -1000,
	-1000, -1000, 31, 24, -1000, -1000, 151, -1000, -1000, 698,
	-1000, 282, -1000, -1000, 93, 396, 396, 396, 698, 698,
	639, -1000, 34, 661, -1000, 10, 10, -12, -12, -12,
	112, 112, -1000, -1000, -1000, 396, -1000, 698, -1000, 196,
	547, 196, 244, 73, -1000, 651, 342, 92, 92, 372,
	364, 370, 105, -1000, -1000, 105, -1000, 576, 255, 495,
	-1000, -1000, -1000, -1000, 291, -1000, 284, -1000, -1000, -1000,
	-43, -47, -59, -1000, -1000, -1000, -1000, 393, -1000, -1000,
	-1000, 217, -1000, 698, 698, 483, 396, 698, -1000, 196,
	-1000, 185, -27, -1000, 396, 155, 254, 250, 508, 237,
	194, -1000, 364, -1000, 396, 396, -1000, -1000, 651, 651,
	-1000, -1000, 250, 250, 250, 379, 366, -1000, 396, 698,
	-1000, 94, -1000, 698, 396, -1000, 324, 241, -1000, -1000,
	-1000, 92, -1000, 422, 231, -1000, 176, -1000, 26, 26,
	81, 81, 81, 372, 651, 396, 698, -1000, 698, 317,
	250, -1000, 396, 396, -1000, -1000, -1000, 182, -1000, 182,
	182, 364, 26, 199, 415, -1000, 698, -1000, -1000, 81,
	-1000, -1000, 316, 92, -1000, -1000, 414, 3, 171, -1000,
	81, 81, -1000,
}
var yyPgo = [...]int{

	0, 582, 14, 403, 397, 396, 581, 579, 578, 572,
	559, 558, 557, 556, 555, 554, 551, 549, 475, 547,
	546, 545, 22, 34, 544, 31, 20, 16, 542, 540,
	5, 530, 523, 28, 521, 18, 23, 41, 520, 519,
	518, 8, 2, 19, 17, 13, 511, 1, 510, 295,
	4, 494, 488, 12, 477, 476, 473, 472, 10, 471,
	3, 469, 7, 468, 466, 465, 11, 9, 21, 322,
	151, 464, 462, 461, 444, 442, 0, 441, 394, 440,
	434, 6, 430, 429, 122, 27,
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
	42, 42, 42, 42, 42, 42, 48, 51, 54, 54,
	52, 52, 53, 55, 55, 50, 50, 41, 41, 41,
	41, 56, 56, 57, 57, 58, 58, 59, 59, 60,
	61, 61, 61, 62, 62, 62, 63, 63, 63, 64,
	64, 65, 65, 66, 66, 40, 40, 46, 46, 47,
	47, 67, 67, 68, 70, 70, 71, 71, 69, 69,
	72, 72, 72, 72, 72, 73, 73, 74, 74, 75,
	75, 76, 78, 84, 85, 81,
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
	2, 3, 4, 5, 4, 1, 1, 5, 0, 1,
	1, 2, 4, 0, 2, 1, 3, 1, 1, 1,
	1, 0, 3, 0, 2, 0, 3, 1, 3, 2,
	0, 1, 1, 0, 2, 4, 0, 2, 4, 0,
	3, 1, 3, 0, 5, 2, 1, 1, 3, 3,
	1, 1, 3, 3, 0, 2, 0, 3, 0, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 1, 1, 0,
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
	88, -76, -48, -51, 57, 58, 24, 33, 47, 51,
	52, 53, 56, -49, 103, 63, 55, 27, 31, 90,
	-33, 48, 67, 90, -78, 61, 50, -81, -81, -78,
	-81, 101, -78, 24, 46, -76, -36, 75, 14, -25,
	-26, -27, -28, -32, -49, -84, -78, 13, 48, -77,
	-76, 23, 90, 60, 59, 74, -39, 77, 61, 75,
	76, 62, 74, 79, 78, 87, 82, 83, 84, 85,
	86, 80, 81, 67, 68, 69, 70, 71, 72, 73,
	-37, -42, -37, -2, -45, -42, -42, -84, -42, -42,
	-84, -84, -49, -84, -84, -54, -42, -33, -67, -78,
	-36, -68, -42, -76, -81, 24, -75, 105, -72, 95,
	93, 30, 94, 17, 50, -78, -78, -81, -42, -37,
	48, -29, -30, -31, 36, 40, 42, 37, 38, 39,
	43, -79, -78, 23, -80, 23, -25, 90, -25, -23,
	-76, 84, -37, -37, -43, 56, 61, 57, 58, -42,
	-44, -84, -49, 54, 77, 75, 76, 62, -42, -42,
	-42, -43, -42, -42, -42, -42, -42, -42, -42, -42,
	-42, -42, -85, 49, -85, 48, -85, -42, -85, -22,
	22, -22, -41, -52, -53, 64, -64, 31, -84, -36,
	-58, 17, 46, -76, -81, -73, 101, -26, -27, -26,
	-27, 36, 36, 36, 41, 36, 41, 36, -30, -34,
	44, 102, 45, -78, -78, -85, -78, -36, 56, 57,
	58, -45, -44, -42, -42, -42, 60, -42, -85, -22,
	-85, 48, -55, -53, 66, -37, -40, 34, -2, -67,
	-65, -50, -58, -62, 19, 18, -78, -78, 46, 46,
	36, 36, 99, 99, 99, -56, 15, -85, 60, -42,
	-85, -41, 91, -42, 65, -66, 46, -46, -47, -66,
	-85, 48, -62, -42, -59, -60, -42, -81, -37, -37,
	-84, -84, -84, -57, 16, 18, -42, -85, -42, 28,
	48, -50, 48, 48, -61, 25, 26, -35, -76, -35,
	-35, -58, -37, -45, 29, -47, -42, -60, -85, 48,
	-85, -85, -62, 11, -76, -63, 20, 32, -67, 11,
	77, -76, -76,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 0, 47, 47,
	47, 47, 47, 227, 218, 0, 0, 0, 0, 0,
	235, 0, 51, 53, 54, 55, 21, 22, 23, 24,
	25, 56, 49, 218, 0, 0, 0, 216, 0, 0,
	0, 228, 0, 0, 219, 0, 214, 0, 214, 0,
	18, 20, 46, 28, 52, 0, 57, 48, 0, 0,
	96, 232, 0, 33, 211, 0, 175, 0, -2, 0,
	0, 0, 235, 231, 235, 0, 235, 0, 0, 0,
	0, 0, 45, 104, 0, 27, 58, 60, 65, 0,
	63, 64, 106, 0, 0, 145, 146, 147, 0, 0,
	0, 175, 0, 165, 112, 113, 0, 0, 233, 177,
	178, 179, 180, 210, 166, 168, 50, 0, 0, 0,
	104, 0, 0, 0, 235, 0, 229, 36, 37, 0,
	40, 0, 42, 215, 0, 235, 16, 0, 0, 19,
	68, 70, 71, 81, 79, 0, 94, 0, 0, 61,
	66, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 132, 133, 134, 135, 136, 137, 138,
	109, 0, 0, 0, 0, 143, 158, 0, 159, 160,
	0, 0, 124, 0, 0, 0, 169, 199, 104, 97,
	185, 212, 213, 176, 34, 217, 0, 0, 235, 225,
	220, 221, 222, 223, 224, 41, 43, 44, 17, 105,
	0, 0, 0, 0, 84, 0, 0, 87, 0, 0,
	0, 98, 82, 0, 0, 80, 0, 0, 104, 59,
	67, 62, 107, 108, 111, 126, 0, 128, 130, 114,
	115, 0, 140, 141, 0, 0, 0, 0, 117, 119,
	0, 123, 148, 149, 150, 151, 152, 153, 154, 155,
	156, 157, 110, 234, 142, 0, 209, 143, 161, 0,
	0, 0, 0, 173, 170, 0, 0, 0, 0, 185,
	193, 0, 0, 230, 38, 0, 226, 69, 75, 0,
	78, 85, 86, 88, 0, 90, 0, 92, 93, 72,
	0, 0, 0, 83, 73, 74, 95, 181, 127, 129,
	131, 0, 116, 118, 120, 0, 0, 144, 162, 0,
	164, 0, 0, 171, 0, 0, 203, 0, 206, 203,
	0, 201, 193, 32, 0, 0, 235, 39, 0, 0,
	89, 91, 0, 0, 0, 183, 0, 139, 0, 121,
	163, 0, 167, 174, 0, 29, 0, 205, 207, 30,
	200, 0, 31, 194, 186, 187, 190, 35, 76, 77,
	0, 0, 0, 185, 0, 0, 122, 125, 172, 0,
	0, 202, 0, 0, 189, 191, 192, 0, 102, 0,
	0, 193, 184, 182, 0, 208, 195, 188, 99, 0,
	100, 101, 196, 0, 103, 26, 0, 0, 204, 197,
	0, 0, 198,
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
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:869
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Exprs: yyDollar[3].selectExprs}
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:873
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].str, Exprs: yyDollar[3].selectExprs}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:887
		{
			yyVAL.str = "if"
		}
	case 167:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:893
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 168:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = nil
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:902
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:908
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 171:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:912
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 172:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:918
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:927
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:933
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].sqlID}
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:937
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 179:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:960
		{
			yyVAL.valExprs = nil
		}
	case 182:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:964
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:969
		{
			yyVAL.boolExpr = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:973
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:978
		{
			yyVAL.orderBy = nil
		}
	case 186:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:982
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 187:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:988
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 188:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:992
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 189:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:998
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 190:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.str = AscScr
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.str = AscScr
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.str = DescScr
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.limit = nil
		}
	case 194:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 195:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.str = ""
		}
	case 197:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.str = ForUpdateStr
		}
	case 198:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1037
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
	case 199:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.columns = nil
		}
	case 200:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1054
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1060
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1064
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 203:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.updateExprs = nil
		}
	case 204:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 205:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 208:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 209:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 213:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 214:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.str = ""
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.str = IgnoreStr
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1144
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1148
		{
			yyVAL.empty = struct{}{}
		}
	case 225:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.empty = struct{}{}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1153
		{
			yyVAL.empty = struct{}{}
		}
	case 227:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.empty = struct{}{}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1158
		{
			yyVAL.empty = struct{}{}
		}
	case 229:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1161
		{
			yyVAL.empty = struct{}{}
		}
	case 230:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1163
		{
			yyVAL.empty = struct{}{}
		}
	case 231:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.sqlID = SQLName(strings.ToLower(string(yyDollar[1].bytes)))
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1173
		{
			yyVAL.sqlID = SQLName(yyDollar[1].bytes)
		}
	case 233:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1179
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1188
		{
			decNesting(yylex)
		}
	case 235:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1193
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
