//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
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

//line sql.y:36
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
const yyMaxDepth = 200

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 73,
	90, 225,
	-2, 224,
}

const yyNprod = 229
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 751

var yyAct = [...]int{

	103, 174, 171, 365, 98, 402, 355, 99, 68, 56,
	320, 279, 268, 211, 262, 210, 173, 4, 97, 250,
	228, 222, 198, 87, 209, 64, 69, 83, 44, 190,
	46, 115, 34, 88, 47, 75, 50, 57, 16, 340,
	342, 49, 377, 50, 71, 58, 275, 77, 70, 132,
	80, 376, 375, 110, 76, 79, 73, 111, 112, 113,
	55, 51, 114, 93, 52, 53, 54, 352, 294, 117,
	92, 158, 159, 160, 155, 142, 110, 155, 125, 73,
	111, 112, 113, 121, 78, 114, 136, 129, 100, 101,
	131, 140, 117, 122, 102, 416, 341, 145, 124, 175,
	204, 235, 263, 176, 178, 179, 127, 263, 116, 311,
	118, 100, 101, 202, 233, 234, 232, 102, 219, 78,
	186, 71, 144, 143, 71, 70, 194, 193, 70, 188,
	143, 116, 110, 205, 73, 82, 195, 145, 66, 231,
	182, 93, 218, 194, 145, 187, 208, 16, 110, 227,
	192, 66, 236, 237, 238, 128, 240, 241, 242, 243,
	244, 245, 246, 247, 248, 249, 170, 172, 290, 214,
	141, 84, 217, 251, 239, 137, 201, 203, 200, 230,
	255, 144, 143, 93, 93, 110, 110, 354, 66, 384,
	85, 251, 123, 252, 254, 66, 145, 78, 371, 272,
	256, 144, 143, 260, 257, 259, 295, 296, 297, 273,
	138, 177, 411, 251, 220, 221, 145, 253, 267, 154,
	153, 161, 162, 156, 157, 158, 159, 160, 155, 361,
	251, 298, 255, 138, 251, 385, 300, 301, 302, 276,
	293, 156, 157, 158, 159, 160, 155, 214, 382, 223,
	225, 226, 308, 299, 224, 356, 304, 123, 111, 112,
	113, 93, 270, 114, 230, 356, 71, 71, 253, 251,
	70, 318, 96, 271, 316, 265, 310, 305, 135, 307,
	319, 315, 306, 277, 251, 387, 388, 29, 30, 31,
	32, 110, 328, 327, 330, 329, 120, 335, 191, 338,
	333, 191, 336, 374, 373, 334, 349, 96, 96, 214,
	214, 214, 214, 345, 353, 180, 181, 332, 347, 331,
	183, 184, 358, 359, 363, 366, 350, 351, 48, 16,
	362, 251, 277, 367, 312, 123, 408, 360, 154, 153,
	161, 162, 156, 157, 158, 159, 160, 155, 409, 215,
	96, 378, 397, 381, 314, 96, 96, 380, 337, 229,
	285, 286, 71, 119, 196, 134, 383, 292, 61, 63,
	59, 379, 391, 255, 321, 370, 322, 269, 72, 369,
	326, 389, 16, 17, 18, 19, 191, 399, 366, 400,
	398, 67, 96, 96, 7, 403, 403, 403, 6, 266,
	401, 404, 405, 96, 20, 5, 415, 71, 406, 16,
	37, 70, 417, 39, 36, 414, 1, 418, 291, 419,
	65, 35, 410, 288, 412, 413, 139, 215, 197, 45,
	81, 274, 199, 74, 86, 16, 17, 18, 19, 91,
	390, 317, 392, 393, 229, 264, 65, 407, 281, 284,
	285, 286, 282, 126, 283, 287, 386, 364, 130, 368,
	325, 133, 309, 38, 185, 21, 22, 24, 23, 25,
	96, 261, 105, 104, 258, 96, 108, 357, 26, 27,
	15, 40, 41, 42, 43, 109, 313, 146, 94, 215,
	215, 215, 215, 29, 30, 31, 32, 339, 65, 110,
	189, 251, 73, 111, 112, 113, 213, 280, 114, 106,
	107, 206, 278, 95, 207, 117, 216, 91, 108, 161,
	162, 156, 157, 158, 159, 160, 155, 109, 212, 90,
	60, 28, 62, 14, 100, 101, 89, 33, 3, 13,
	102, 110, 12, 348, 73, 111, 112, 113, 11, 10,
	114, 106, 107, 9, 116, 95, 8, 117, 2, 91,
	91, 154, 153, 161, 162, 156, 157, 158, 159, 160,
	155, 0, 0, 0, 0, 16, 100, 101, 89, 0,
	0, 96, 102, 96, 96, 303, 0, 394, 395, 396,
	108, 0, 289, 0, 216, 0, 116, 0, 0, 109,
	0, 0, 0, 154, 153, 161, 162, 156, 157, 158,
	159, 160, 155, 110, 0, 0, 73, 111, 112, 113,
	0, 0, 114, 106, 107, 0, 0, 95, 0, 117,
	0, 0, 0, 0, 0, 0, 0, 91, 154, 153,
	161, 162, 156, 157, 158, 159, 160, 155, 100, 101,
	323, 108, 0, 324, 102, 0, 216, 216, 216, 216,
	109, 0, 0, 0, 0, 0, 0, 0, 116, 343,
	344, 0, 0, 346, 110, 0, 0, 73, 111, 112,
	113, 0, 0, 114, 106, 107, 0, 0, 95, 0,
	117, 153, 161, 162, 156, 157, 158, 159, 160, 155,
	0, 281, 284, 285, 286, 282, 0, 283, 287, 100,
	101, 372, 0, 0, 0, 102, 0, 0, 0, 0,
	0, 0, 0, 0, 148, 151, 0, 0, 0, 116,
	163, 164, 165, 166, 167, 168, 169, 152, 149, 150,
	147, 154, 153, 161, 162, 156, 157, 158, 159, 160,
	155,
}
var yyPact = [...]int{

	373, -1000, -1000, -1000, 488, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 426, -1000, -1000, -1000, -1000,
	-1000, -69, -58, -36, -33, -37, -1000, -1000, 400, 349,
	-1000, -1000, -1000, -1000, 488, -1000, -1000, -1000, 346, -1000,
	-65, 88, 378, 84, -67, -44, 69, -1000, -42, 69,
	-1000, 88, -75, 121, -75, 88, -1000, -1000, -1000, -1000,
	494, -1000, 55, 336, 265, -7, -1000, 88, 144, -1000,
	31, -1000, -12, -1000, 88, 45, 105, -1000, -1000, 88,
	-1000, -51, 88, 341, 232, 69, -1000, 162, -1000, -1000,
	147, -15, 63, 663, -1000, 627, 566, -1000, -1000, -1000,
	6, 6, 6, 139, 139, -1000, -1000, -1000, 139, 139,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 6, -1000, 88,
	84, 88, 372, 84, 6, 69, -1000, 340, -82, -1000,
	83, -1000, 88, -1000, -1000, 88, -1000, 101, 494, -1000,
	-1000, 69, 34, 627, 627, 193, 6, 85, 39, 6,
	6, 6, 193, 6, 6, 6, 6, 6, 6, 6,
	6, 6, 6, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	23, 663, 142, 282, 220, 663, -1000, 29, -1000, -1000,
	452, 494, -1000, 400, 207, 38, 560, 244, 287, -1000,
	360, 627, -1000, 560, -1000, -1000, -1000, 227, 69, -1000,
	-54, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 284,
	412, -1000, -1000, 145, 344, 138, -22, -1000, -1000, -1000,
	23, 70, -1000, -1000, 150, -1000, -1000, 560, -1000, 29,
	-1000, -1000, 85, 6, 6, 6, 560, 560, 525, -1000,
	439, 612, -1000, -13, -13, -10, -10, -10, 159, 159,
	-1000, -1000, -1000, 6, -1000, 560, -1000, 185, 494, 185,
	204, 43, -1000, 627, 320, 84, 84, 360, 355, 358,
	63, 88, -1000, -1000, 88, -1000, 365, 101, 101, 101,
	101, -1000, 283, 281, -1000, 264, 261, 322, -5, -1000,
	88, 88, -1000, 235, 88, -1000, -1000, -1000, 220, -1000,
	560, 560, 483, 6, 560, -1000, 185, -1000, 207, -24,
	-1000, 6, 122, 219, 139, 488, 209, 181, -1000, 355,
	-1000, 6, 6, -1000, -1000, 363, 357, 412, 152, 665,
	-1000, -1000, -1000, -1000, 268, -1000, 267, -1000, -1000, -1000,
	-46, -47, -56, -1000, -1000, -1000, -1000, -1000, 6, 560,
	-1000, 124, -1000, 560, 6, -1000, 325, 200, -1000, -1000,
	-1000, 84, -1000, 141, 187, -1000, 260, -1000, 360, 627,
	6, 627, 627, -1000, -1000, 139, 139, 139, 560, -1000,
	560, 323, 139, -1000, 6, 6, -1000, -1000, -1000, 355,
	63, 169, 63, 63, 69, 69, 69, 397, -1000, 560,
	-1000, 316, 164, -1000, 164, 164, 84, -1000, 395, 18,
	-1000, 69, -1000, -1000, 144, -1000, 69, -1000, 69, -1000,
}
var yyPgo = [...]int{

	0, 558, 16, 405, 398, 394, 556, 553, 549, 548,
	542, 539, 538, 537, 533, 463, 532, 531, 530, 23,
	33, 529, 24, 15, 13, 528, 512, 11, 507, 506,
	25, 497, 5, 29, 70, 488, 487, 486, 18, 2,
	21, 20, 1, 477, 7, 473, 31, 4, 472, 471,
	14, 464, 462, 460, 459, 12, 457, 3, 456, 10,
	447, 445, 441, 6, 8, 26, 328, 135, 433, 432,
	431, 429, 428, 0, 426, 378, 423, 418, 9, 416,
	413, 211, 19,
}
var yyR1 = [...]int{

	0, 79, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 12, 12, 13, 13, 13, 13,
	2, 2, 3, 3, 4, 5, 6, 7, 7, 7,
	8, 8, 8, 9, 10, 10, 10, 11, 14, 14,
	80, 15, 16, 16, 17, 17, 17, 17, 17, 18,
	18, 19, 19, 20, 20, 20, 21, 21, 74, 74,
	74, 22, 22, 23, 23, 24, 24, 24, 25, 25,
	25, 25, 77, 77, 76, 76, 76, 26, 26, 26,
	26, 27, 27, 27, 27, 28, 28, 29, 29, 30,
	30, 31, 31, 31, 31, 32, 32, 33, 33, 34,
	34, 34, 34, 34, 34, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 35, 35, 35, 35, 35, 40,
	40, 40, 40, 40, 40, 36, 36, 36, 36, 36,
	36, 36, 41, 41, 41, 46, 42, 42, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 45,
	48, 51, 51, 49, 49, 50, 52, 52, 47, 47,
	38, 38, 38, 38, 53, 53, 54, 54, 55, 55,
	56, 56, 57, 58, 58, 58, 59, 59, 59, 60,
	60, 60, 61, 61, 62, 62, 63, 63, 37, 37,
	43, 43, 44, 44, 64, 64, 65, 67, 67, 68,
	68, 66, 66, 69, 69, 69, 69, 69, 70, 70,
	71, 71, 72, 72, 73, 75, 81, 82, 78,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 0, 1, 1, 1, 1,
	12, 3, 8, 8, 8, 7, 3, 5, 8, 4,
	6, 7, 4, 5, 4, 5, 5, 3, 2, 2,
	0, 2, 0, 2, 1, 2, 1, 1, 1, 0,
	1, 1, 3, 1, 2, 3, 1, 1, 0, 1,
	2, 1, 3, 1, 1, 3, 3, 3, 3, 5,
	5, 3, 0, 1, 0, 1, 2, 1, 2, 2,
	1, 2, 3, 2, 3, 2, 2, 1, 3, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 1,
	3, 3, 2, 3, 3, 1, 1, 3, 3, 4,
	3, 4, 3, 4, 5, 6, 3, 2, 6, 1,
	2, 1, 2, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 1, 3, 1, 3, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 3, 4, 5, 4, 1, 1,
	5, 0, 1, 1, 2, 4, 0, 2, 1, 3,
	1, 1, 1, 1, 0, 3, 0, 2, 0, 3,
	1, 3, 2, 0, 1, 1, 0, 2, 4, 0,
	2, 4, 0, 3, 1, 3, 0, 5, 2, 1,
	1, 3, 3, 1, 1, 3, 3, 0, 2, 0,
	3, 0, 1, 1, 1, 1, 1, 1, 0, 1,
	0, 1, 0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -79, -1, -12, -2, -3, -4, -5, -6, -7,
	-8, -9, -10, -11, -14, 107, 9, 10, 11, 12,
	31, 92, 93, 95, 94, 96, 105, 106, -17, 5,
	6, 7, 8, -13, -2, -3, -4, -5, -15, -80,
	-15, -15, -15, -15, 97, -71, 99, 103, -66, 99,
	101, 97, 97, 98, 99, 97, -78, -78, -2, 21,
	-18, 22, -16, -66, -30, -75, 50, 13, -64, -65,
	-47, -73, -75, 50, -68, 102, 98, -73, 50, 97,
	-73, -75, -67, 102, 50, -67, -75, -19, -20, 84,
	-21, -75, -34, -39, -35, 61, -81, -38, -47, -44,
	82, 83, 88, -73, -45, -48, 57, 58, 24, 33,
	47, 51, 52, 53, 56, -46, 102, 63, 55, 27,
	31, 90, -30, 48, 67, 90, -75, 61, 50, -78,
	-75, -78, 100, -75, 24, 46, -73, 13, 48, -74,
	-73, 23, 90, 60, 59, 74, -36, 77, 61, 75,
	76, 62, 74, 79, 78, 87, 82, 83, 84, 85,
	86, 80, 81, 67, 68, 69, 70, 71, 72, 73,
	-34, -39, -34, -2, -42, -39, -39, -81, -39, -39,
	-81, -81, -46, -81, -81, -51, -39, -30, -64, -75,
	-33, 14, -65, -39, -73, -78, 24, -72, 104, -69,
	95, 93, 30, 94, 17, 50, -75, -75, -78, -22,
	-23, -24, -25, -29, -46, -81, -75, -20, -73, 84,
	-34, -34, -40, 56, 61, 57, 58, -39, -41, -81,
	-46, 54, 77, 75, 76, 62, -39, -39, -39, -40,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-82, 49, -82, 48, -82, -39, -82, -19, 22, -19,
	-38, -49, -50, 64, -61, 31, -81, -33, -55, 17,
	-34, 46, -73, -78, -70, 100, -33, 48, -26, -27,
	-28, 36, 40, 42, 37, 38, 39, 43, -76, -75,
	23, -77, 23, -22, 90, 56, 57, 58, -42, -41,
	-39, -39, -39, 60, -39, -82, -19, -82, 48, -52,
	-50, 66, -34, -37, 34, -2, -64, -62, -47, -55,
	-59, 19, 18, -75, -75, -53, 15, -23, -24, -23,
	-24, 36, 36, 36, 41, 36, 41, 36, -27, -31,
	44, 101, 45, -75, -75, -82, -75, -82, 60, -39,
	-82, -38, 91, -39, 65, -63, 46, -43, -44, -63,
	-82, 48, -59, -39, -56, -57, -39, -78, -54, 16,
	18, 46, 46, 36, 36, 98, 98, 98, -39, -82,
	-39, 28, 48, -47, 48, 48, -58, 25, 26, -55,
	-34, -42, -34, -34, -81, -81, -81, 29, -44, -39,
	-57, -59, -32, -73, -32, -32, 11, -60, 20, 32,
	-82, 48, -82, -82, -64, 11, 77, -73, -73, -73,
}
var yyDef = [...]int{

	15, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 0, 40, 40, 40, 40,
	40, 220, 211, 0, 0, 0, 228, 228, 0, 44,
	46, 47, 48, 14, 16, 17, 18, 19, 49, 42,
	211, 0, 0, 0, 209, 0, 0, 221, 0, 0,
	212, 0, 207, 0, 207, 0, 38, 39, 21, 45,
	0, 50, 41, 0, 0, 89, 225, 0, 26, 204,
	0, 168, 0, -2, 0, 0, 0, 228, 224, 0,
	228, 0, 0, 0, 0, 0, 37, 0, 51, 53,
	58, 0, 56, 57, 99, 0, 0, 138, 139, 140,
	0, 0, 0, 168, 0, 158, 105, 106, 0, 0,
	226, 170, 171, 172, 173, 203, 159, 161, 43, 0,
	0, 0, 97, 0, 0, 0, 228, 0, 222, 29,
	0, 32, 0, 34, 208, 0, 228, 0, 0, 54,
	59, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 125, 126, 127, 128, 129, 130, 131,
	102, 0, 0, 0, 0, 136, 151, 0, 152, 153,
	0, 0, 117, 0, 0, 0, 162, 192, 97, 90,
	178, 0, 205, 206, 169, 27, 210, 0, 0, 228,
	218, 213, 214, 215, 216, 217, 33, 35, 36, 97,
	61, 63, 64, 74, 72, 0, 87, 52, 60, 55,
	100, 101, 104, 119, 0, 121, 123, 107, 108, 0,
	133, 134, 0, 0, 0, 0, 110, 112, 0, 116,
	141, 142, 143, 144, 145, 146, 147, 148, 149, 150,
	103, 227, 135, 0, 202, 136, 154, 0, 0, 0,
	0, 166, 163, 0, 0, 0, 0, 178, 186, 0,
	98, 0, 223, 30, 0, 219, 174, 0, 0, 0,
	0, 77, 0, 0, 80, 0, 0, 0, 91, 75,
	0, 0, 73, 0, 0, 120, 122, 124, 0, 109,
	111, 113, 0, 0, 137, 155, 0, 157, 0, 0,
	164, 0, 0, 196, 0, 199, 196, 0, 194, 186,
	25, 0, 0, 228, 31, 176, 0, 62, 68, 0,
	71, 78, 79, 81, 0, 83, 0, 85, 86, 65,
	0, 0, 0, 76, 66, 67, 88, 132, 0, 114,
	156, 0, 160, 167, 0, 22, 0, 198, 200, 23,
	193, 0, 24, 187, 179, 180, 183, 28, 178, 0,
	0, 0, 0, 82, 84, 0, 0, 0, 115, 118,
	165, 0, 0, 195, 0, 0, 182, 184, 185, 186,
	177, 175, 69, 70, 0, 0, 0, 0, 201, 188,
	181, 189, 0, 95, 0, 0, 0, 20, 0, 0,
	92, 0, 93, 94, 197, 190, 0, 96, 0, 191,
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
	lookahead func() int
}

func (p *yyParserImpl) Lookahead() int {
	return p.lookahead()
}

func yyNewParser() yyParser {
	p := &yyParserImpl{
		lookahead: func() int { return -1 },
	}
	return p
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
	var yylval yySymType
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yytoken := -1 // yychar translated into internal numbering
	yyrcvr.lookahead = func() int { return yychar }
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yychar = -1
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
	if yychar < 0 {
		yychar, yytoken = yylex1(yylex, &yylval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yychar = -1
		yytoken = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar, yytoken = yylex1(yylex, &yylval)
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
			yychar = -1
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
		//line sql.y:168
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:175
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 14:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:191
		{
			yyVAL.statement = &Explain{SQLNode: yyDollar[2].statement}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:195
		{
			yyVAL.statement = &Other{}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:201
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 20:
		yyDollar = yyS[yypt-12 : yypt+1]
		//line sql.y:210
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, SelectExprs: yyDollar[4].selectExprs, From: yyDollar[6].tableExprs, Where: NewWhere(WhereStr, yyDollar[7].boolExpr), GroupBy: GroupBy(yyDollar[8].valExprs), Having: NewWhere(HavingStr, yyDollar[9].boolExpr), OrderBy: yyDollar[10].orderBy, Limit: yyDollar[11].limit, Lock: yyDollar[12].str}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:214
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:220
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:224
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, col := range yyDollar[7].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 24:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:236
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:242
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:248
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:254
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].sqlID}
		}
	case 28:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:258
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].sqlID, NewName: yyDollar[7].sqlID}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:263
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: SQLName(yyDollar[3].sqlID)}
		}
	case 30:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:269
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].sqlID, NewName: yyDollar[4].sqlID}
		}
	case 31:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:273
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].sqlID, NewName: yyDollar[7].sqlID}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:278
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: SQLName(yyDollar[3].sqlID), NewName: SQLName(yyDollar[3].sqlID)}
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:284
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:290
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].sqlID}
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:294
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].sqlID, NewName: yyDollar[5].sqlID}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:299
		{
			yyVAL.statement = &DDL{Action: DropStr, Table: SQLName(yyDollar[4].sqlID)}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:305
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].sqlID, NewName: yyDollar[3].sqlID}
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:311
		{
			yyVAL.statement = &Other{}
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:315
		{
			yyVAL.statement = &Other{}
		}
	case 40:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:320
		{
			setAllowComments(yylex, true)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:324
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:330
		{
			yyVAL.bytes2 = nil
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:334
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:340
		{
			yyVAL.str = UnionStr
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:344
		{
			yyVAL.str = UnionAllStr
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:348
		{
			yyVAL.str = SetMinusStr
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:352
		{
			yyVAL.str = ExceptStr
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:356
		{
			yyVAL.str = IntersectStr
		}
	case 49:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:361
		{
			yyVAL.str = ""
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:365
		{
			yyVAL.str = DistinctStr
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:371
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:375
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:381
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:385
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].sqlID}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:389
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].sqlID}
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:399
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 58:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:404
		{
			yyVAL.sqlID = ""
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:408
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:412
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:418
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:422
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:432
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].smTableExpr, As: yyDollar[2].sqlID, Hints: yyDollar[3].indexHints}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:436
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].sqlID}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:440
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:453
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:457
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 70:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:461
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:465
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 72:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:470
		{
			yyVAL.empty = struct{}{}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:472
		{
			yyVAL.empty = struct{}{}
		}
	case 74:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:475
		{
			yyVAL.sqlID = ""
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:479
		{
			yyVAL.sqlID = yyDollar[1].sqlID
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:483
		{
			yyVAL.sqlID = yyDollar[2].sqlID
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:489
		{
			yyVAL.str = JoinStr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:493
		{
			yyVAL.str = JoinStr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:497
		{
			yyVAL.str = JoinStr
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:501
		{
			yyVAL.str = StraightJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:507
		{
			yyVAL.str = LeftJoinStr
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:511
		{
			yyVAL.str = LeftJoinStr
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:515
		{
			yyVAL.str = RightJoinStr
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = RightJoinStr
		}
	case 85:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:525
		{
			yyVAL.str = NaturalJoinStr
		}
	case 86:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:529
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:539
		{
			yyVAL.smTableExpr = &TableName{Name: yyDollar[1].sqlID}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:543
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:549
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].sqlID}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:553
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:558
		{
			yyVAL.indexHints = nil
		}
	case 92:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:562
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 93:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:566
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 94:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:570
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].sqlIDs}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:576
		{
			yyVAL.sqlIDs = []SQLName{yyDollar[1].sqlID}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:580
		{
			yyVAL.sqlIDs = append(yyDollar[1].sqlIDs, yyDollar[3].sqlID)
		}
	case 97:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:585
		{
			yyVAL.boolExpr = nil
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:589
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:596
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:600
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 102:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:604
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:608
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:612
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:618
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:630
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 109:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:634
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:638
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:642
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:646
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 113:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:650
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 114:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:654
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 115:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:658
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:662
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:666
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 118:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:670
		{
			yyVAL.boolExpr = &KeyrangeExpr{Start: yyDollar[3].valExpr, End: yyDollar[5].valExpr}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:676
		{
			yyVAL.str = IsNullStr
		}
	case 120:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:680
		{
			yyVAL.str = IsNotNullStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:684
		{
			yyVAL.str = IsTrueStr
		}
	case 122:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:688
		{
			yyVAL.str = IsNotTrueStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:692
		{
			yyVAL.str = IsFalseStr
		}
	case 124:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = IsNotFalseStr
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = EqualStr
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = LessThanStr
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = GreaterThanStr
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = LessEqualStr
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = GreaterEqualStr
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:722
		{
			yyVAL.str = NotEqualStr
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:732
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 134:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:746
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:752
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:756
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 138:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:766
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:770
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:782
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:786
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:790
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 146:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:798
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:802
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:806
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:810
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 151:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:814
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:822
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
	case 153:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:835
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:839
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID)}
		}
	case 155:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:843
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Exprs: yyDollar[3].selectExprs}
		}
	case 156:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:847
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].sqlID), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 157:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:851
		{
			yyVAL.valExpr = &FuncExpr{Name: yyDollar[1].str, Exprs: yyDollar[3].selectExprs}
		}
	case 158:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:855
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 159:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:861
		{
			yyVAL.str = "if"
		}
	case 160:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:867
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 161:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:872
		{
			yyVAL.valExpr = nil
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:876
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 163:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:882
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 164:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:886
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 165:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:892
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:897
		{
			yyVAL.valExpr = nil
		}
	case 167:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:901
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:907
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].sqlID}
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:911
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].sqlID, Name: yyDollar[3].sqlID}
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:917
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 171:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:921
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:925
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 173:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:929
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:934
		{
			yyVAL.valExprs = nil
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:938
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 176:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:943
		{
			yyVAL.boolExpr = nil
		}
	case 177:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:947
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 178:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:952
		{
			yyVAL.orderBy = nil
		}
	case 179:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:956
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:962
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 181:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:966
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 182:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:972
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:977
		{
			yyVAL.str = AscScr
		}
	case 184:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:981
		{
			yyVAL.str = AscScr
		}
	case 185:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:985
		{
			yyVAL.str = DescScr
		}
	case 186:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:990
		{
			yyVAL.limit = nil
		}
	case 187:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:994
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 188:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:998
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1003
		{
			yyVAL.str = ""
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1007
		{
			yyVAL.str = ForUpdateStr
		}
	case 191:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1011
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
	case 192:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.columns = nil
		}
	case 193:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyDollar[1].colName}}
		}
	case 195:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyDollar[3].colName})
		}
	case 196:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1043
		{
			yyVAL.updateExprs = nil
		}
	case 197:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1047
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 198:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1053
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1057
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 201:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 202:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 205:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 206:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colName, Expr: yyDollar[3].valExpr}
		}
	case 207:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.empty = struct{}{}
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.empty = struct{}{}
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.empty = struct{}{}
		}
	case 210:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.empty = struct{}{}
		}
	case 211:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.str = ""
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.str = IgnoreStr
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.empty = struct{}{}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.empty = struct{}{}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.empty = struct{}{}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.empty = struct{}{}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.empty = struct{}{}
		}
	case 218:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.empty = struct{}{}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.empty = struct{}{}
		}
	case 220:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.empty = struct{}{}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.empty = struct{}{}
		}
	case 222:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.empty = struct{}{}
		}
	case 223:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.empty = struct{}{}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.sqlID = SQLName(strings.ToLower(string(yyDollar[1].bytes)))
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.sqlID = SQLName(yyDollar[1].bytes)
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1153
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 227:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1162
		{
			decNesting(yylex)
		}
	case 228:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1167
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
