package gopdf

import (
	"gopdf/core"
	"testing"
	"fmt"
)

const (
	DIV_IG = "IPAexG"
	DIV_MD = "MPBOLD"
)

func ComplexDivReport() {
	r := core.CreateReport()
	r.IsMutiPage = true
	font1 := core.FontMap{
		FontName: DIV_IG,
		FileName: "ttf//ipaexg.ttf",
	}
	font2 := core.FontMap{
		FontName: DIV_MD,
		FileName: "ttf//mplus-1p-bold.ttf",
	}
	fonts := []*core.FontMap{&font1, &font2}
	r.SetFonts(fonts)
	d := new(DivDetail)
	r.RegisterBand(core.Band(*d), core.Detail)

	r.SetPage("A4", "mm", "P")
	r.SetFooterY(265)
	r.Execute("div_test.pdf")
	r.SaveText("div_test.txt")
}

type DivDetail struct {
}

func (h DivDetail) GetHeight(report *core.Report) float64 {
	return 6
}
func (h DivDetail) Execute(report *core.Report) {
	font := Font{Family: DIV_IG, Size: 10}
	report.CurrX = 10
	report.CurrY = 10

	div := NewDivWithWidth(120*report.GetConvPt(), report)
	div.SetFont(font)
	div.SetMarign(10, 20, 0, 0)
	div.SetBorder(0,50, 10, 0)
	div.SetContent(`
13.2.10 Subquery Syntax 
13.2.10.1 The Subquery as Scalar Operand 
13.2.10.2 Comparisons Using Subqueries
13.2.10.3 Subqueries with ANY, IN, or SOME
13.2.10.4 Subqueries with ALL
13.2.10.5 Row Subqueries
13.2.10.6 Subqueries with EXISTS or NOT EXISTS
13.2.10.7 Correlated Subqueries
13.2.10.8 Derived Tables
13.2.10.9 Subquery Errors
13.2.10.10 Optimizing Subqueries
13.2.10.11 Rewriting Subqueries as Joins
A subquery is a SELECT statement within another statement.

All subquery forms and operations that the SQL standard requires are supported, as well as a few features that are MySQL-specific.

Here is an example of a subquery:

SELECT * FROM t1 WHERE column1 = (SELECT column1 FROM t2);
In this example, SELECT * FROM t1 ... is the outer query (or outer statement), and (SELECT column1 FROM t2) is the subquery. We say that the subquery is nested within the outer query, and in fact it is possible to nest subqueries within other subqueries, to a considerable depth. A subquery must always appear within parentheses.

The main advantages of subqueries are:

They allow queries that are structured so that it is possible to isolate each part of a statement.

They provide alternative ways to perform operations that would otherwise require complex joins and unions.

Many people find subqueries more readable than complex joins or unions. Indeed, it was the innovation of subqueries that gave people the original idea of calling the early SQL “Structured Query Language.”

Here is an example statement that shows the major points about subquery syntax as specified by the SQL standard and supported in MySQL:

DELETE FROM t1
WHERE s11 > ANY
 (SELECT COUNT(*) /* no hint */ FROM t2
  WHERE NOT EXISTS
   (SELECT * FROM t3
    WHERE ROW(5*t2.s1,77)=
     (SELECT 50,11*s1 FROM t4 UNION SELECT 50,77 FROM
      (SELECT * FROM t5) AS t5)));
A subquery can return a scalar (a single value), a single row, a single column, or a table (one or more rows of one or more columns). These are called scalar, column, row, and table subqueries. Subqueries that return a particular kind of result often can be used only in certain contexts, as described in the following sections.

There are few restrictions on the type of statements in which subqueries can be used. A subquery can contain many of the keywords or clauses that an ordinary SELECT can contain: DISTINCT, GROUP BY, ORDER BY, LIMIT, joins, index hints, UNION constructs, comments, functions, and so on.

A subquery's outer statement can be any one of: SELECT, INSERT, UPDATE, DELETE, SET, or DO.

In MySQL, you cannot modify a table and select from the same table in a subquery. This applies to statements such as DELETE, INSERT, REPLACE, UPDATE, and (because subqueries can be used in the SET clause) LOAD DATA INFILE.

For information about how the optimizer handles subqueries, see Section 8.2.2, “Optimizing Subqueries, Derived Tables, and View References”. For a discussion of restrictions on subquery use, including performance issues for certain forms of subquery syntax, see Section C.4, “Restrictions on Subqueries”.`)
	fmt.Println(len(div.contents), div.contents)
	div.GenerateAtomicCellWithAutoWarp()
}

func (h DivDetail) BeforeAddNewPage(report core.Report) int {
	if report.DataPos == 0 {
		//max no
		return 2
	}
	curr := report.Records[report.DataPos].([]string)
	before := report.Records[report.DataPos-1].([]string)
	return h.BreakCheckSub(curr, before)
}
func (h DivDetail) AfterAddNewPage(report core.Report) int {
	if report.DataPos == len(report.Records)-1 {
		//max no
		return 2
	}
	curr := report.Records[report.DataPos].([]string)
	after := report.Records[report.DataPos+1].([]string)
	return h.BreakCheckSub(curr, after)
}
func (h DivDetail) BreakCheckSub(row1 []string, row2 []string) int {
	if row1[0] != row2[0] {
		return 2
	}
	if row1[4] != row2[4] {
		return 1
	}
	return 0
}

func TestDiv(t *testing.T) {
	ComplexDivReport()
}
