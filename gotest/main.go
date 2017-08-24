// File_test2 project main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/DavisChen99/gotest/app"
)

const scriptdir string = "/home/data/Scripts/kid17"
const kidrug string = scriptdir + "/lib/kidrug.db"
const kideaf string = scriptdir + "/lib/kideaf.db"
const outdir string = "/home/data/tpreport"

const head string = `\\PassOptionsToPackage{table}{xcolor}
\\documentclass[a4paper]{article}
\\usepackage{graphicx}
\\usepackage{color}
\\usepackage{lipsum}
\\usepackage{xeCJK}
\\setCJKmainfont{MI LANTING}
\\usepackage{color,framed}
\\usepackage{tikz}

\\usepackage{supertabular}
\\usetikzlibrary{shapes}
\\usepackage[framemethod=default]{mdframed}
\\usepackage{caption}
\\usepackage{indentfirst}
\\usepackage{colortbl,longtable}
\\usepackage{array}

\\usepackage{amssymb,changepage}
\\usepackage{setspace}
\\usepackage{rotating}
\\usepackage{leftidx}
\\usepackage{enumerate}
\\usepackage{cite}
\\usepackage{pifont}
\\usepackage{wasysym}
\\usepackage{fancyhdr}
\\usetikzlibrary{shapes.geometric}
\\usepackage{pbox}
\\usepackage{pgfplots}
\\usepackage{graphicx}
\\usepackage{ulem}
\\usepackage{times}
\\usepackage{datetime}
\\usepackage{booktabs}
\\usepackage{multirow}
\\usepackage{colortbl}
\\usepackage{tikz}
\\usepackage{eso-pic}

\\setlength{\\abovetopsep}{0.5ex}
\\setlength{\\belowrulesep}{0pt}
\\setlength{\\aboverulesep}{0pt}

\\usepackage{booktabs, makecell}


\\setlength{\\parindent}{2cm}
\\definecolor{summerizecolor1}{RGB}{204,204,255}
\\definecolor{summerizecolor}{RGB}{173,251,98}
\\definecolor{riskcolor}{RGB}{255,191,113}
\\definecolor{logocolor}{RGB}{132, 204, 201}
\\definecolor{headercolor}{RGB}{129,194,214}
\\definecolor{cloudcolor}{RGB}{108,196,191}
\\definecolor{normalbarcolor}{RGB}{211,140,7}
\\definecolor{cancerbarcolor}{RGB}{212,73,6}
\\definecolor{tablebluecolor}{RGB}{0,102,204}
\\definecolor{skybluecolor}{RGB}{87,210,247}
\\definecolor{rosered}{RGB}{255,102,204}
\\definecolor{watermelonred}{RGB}{204,51,204}
\\definecolor{reddishbrown}{RGB}{204,102,0}
\\definecolor{lightgreen}{RGB}{221,240,237}
\\definecolor{lightyellow}{RGB}{255,255,204}
\\definecolor{lightpink}{RGB}{255,204,204}
\\definecolor{circcolor}{RGB}{244,70,107}
\\definecolor{JG}{RGB}{255,105,180}


\\newcommand{\\upcite}[1]{\\textsuperscript{\\cite{#1}}}
\\renewcommand{\\today}{\\number\\year/\\number\\month/\\number\\day}

\\definecolor{tableheadcolor}{RGB}{191,191,191}
\\definecolor{tablecellcolor}{RGB}{23,252,231}
\\definecolor{tumorcolor}{RGB}{255,109,174}
\\definecolor{digestcolor}{RGB}{255,171,57}
\\definecolor{respirationcolor}{RGB}{255,251,0}
\\definecolor{reproductivecolor}{RGB}{0,249,0}
\\definecolor{urinarycolor}{RGB}{22,171,216}
\\definecolor{immunecolor}{RGB}{170,121,65}
\\definecolor{cardiovascularcolor}{RGB}{254,174,175}
\\definecolor{nervouscolor}{RGB}{181,183,235}
\\definecolor{kineticcolor}{RGB}{182,199,218}
\\definecolor{surfacecolor}{RGB}{0,253,255}
\\definecolor{endocrinecolor}{RGB}{225,175,143}
\\definecolor{medicationcolor}{RGB}{255,179,255}
\\definecolor{customizecolor}{RGB}{211,168,9}
\\definecolor{tabcolor}{RGB}{51,153,204}
\\definecolor{subtabcolor}{RGB}{217,255,253}
\\definecolor{lightgrey}{RGB}{211,211,211}

\\usepackage[paperheight=270mm,paperwidth=191mm,left=20mm,right=20mm,top=2cm,bottom=2.2cm]{geometry}

\\usepackage{lastpage}
\\usepackage{fancyhdr}
\\pagestyle{fancy}
\\fancyhf{}
\\rhead {\\vspace{0.9cm}
\\begin{flushleft}
\\includegraphics[width=3cm]{scriptdir/pics/ybb1}
\\hspace{5.8cm}
\\textcolor{cloudcolor}{\\textbf{\\includegraphics[width=6cm]{scriptdir/pics/dlogo.png}}}
\\end{flushleft}}
\\rfoot{\\raisebox{1ex}[1ex]{\\normalsize{\\thepage}}}

\\renewcommand{\\headrule}{\\color{headercolor}\\rule{\\headwidth}{0pt}}
\\headsep=25mm
\\renewcommand{\\figurename}{图}
\\setlength\\parskip{.3\\baselineskip}
\\setlength{\\footskip}{0pt}
\\usepackage{wallpaper}


\\begin{document}

\\thispagestyle{empty}

\\CenterWallPaper{1.02}{scriptdir/pics/new-background-cover.jpg}


\\vspace*{0.5cm}
\\begin{center}
\\par\\normalfont\\fontsize{32}{32}\\sffamily\\selectfont
\\textcolor{cloudcolor}{\\textbf{\\includegraphics[scale=0.18]{scriptdir/pics/ybb}\\\\婴幼儿安全用药\\\\基因检测报告}}\\\\
\\par
\\end{center}
\\vspace{1cm}
\\begin{center}
\\includegraphics[scale=0.3]{scriptdir/pics/pill-dna.jpg}
\\end{center}
\\vspace{1cm}
\\begin{center}
\\includegraphics[scale=0.1]{scriptdir/pics/dlogo.png}
\\end{center}

\\setcounter{page}{1}

\\newpage
\\CenterWallPaper{1}{scriptdir/pics/new-background.jpg}

\\setcounter{page}{1}
\\ \\vspace{1cm}
\\begin{flushleft}
\\textcolor{black}{\\LARGE{\\textbf{阅读须知}}}
\\end{flushleft}

\\noindent\\textcolor{circcolor}{\\rule[0.25\\baselineskip]{\\textwidth}{2pt}}
\\large{
\\begin{itemize}
\\item 本研究仅从遗传角度为您或您的医生提供用药参考,不作为调整用药方案的 直接依据,日常服药请务必遵照医嘱。\\\\
\\item 本报告基因位点检测结果和评估建议,只对本次送检样本负责。
\\end{itemize}
}
\\noindent\\textcolor{circcolor}{\\rule[0.25\\baselineskip]{\\textwidth}{2pt}}
\\vspace{0.2cm}
\\begin{spacing}{1.3}
\\noindent \\Large{\\bell} \\normalsize{本报告从遗传角度评估\\textcolor{tabcolor}{ 您对各类药物特有的有效性和毒副反应},检测内容依据国际权威用药数据库,参照\\textcolor{tabcolor}{ FDA}(美国食品药品管理局)、\\textcolor{tabcolor}{ EMA}(欧洲药物管理局)、\\textcolor{tabcolor}{ PMDA}(日本药品及医疗器材管理局)、\\textcolor{tabcolor}{ HealthCanada}(健康加拿大) 、\\textcolor{tabcolor}{ CPIC}(临床药物基因组学实施联盟)、\\textcolor{tabcolor}{ DPWG}(荷兰皇家药物发展联盟药物基因组学中心)等机构发布的药物及相关基因检测标准,进行科学、准确、高效的检测和评估。\\\\\\\\}
 \\Large{\\bell} \\normalsize{本报告未涉及头孢类用药的说明：目前头孢类药物与I相和II相药物代谢酶的相关研究较少，尚无可信度高的遗传评价标准，因此请遵医嘱服用此类药物。}
\\end{spacing}
\\begin{flushleft}
\\Large{\\textcolor{black}{表格格式：}}
\\end{flushleft}
\\begin{flushleft}
\\normalsize{
\\renewcommand{\\arraystretch}{2}
\\arrayrulecolor{tabcolor}
\\begin{tabular}{|m{3cm}<{\\centering}|m{3cm}<{\\centering}|m{3cm}|m{4.5cm}<{\\centering}|}
\\rowcolor{tabcolor!80}
\\hline
\\textcolor{white}{\\textbf{\\Large 中文名}}\\cellcolor{tabcolor!80} &\\textcolor{white}{\\textbf{\\Large 主要成分}}\\cellcolor{tabcolor!80}&\\textcolor{white}{\\textbf{\\Large 遗传评估}}\\cellcolor{tabcolor!80}&\\textcolor{white}{\\textbf{\\Large 服用指导}}\\cellcolor{tabcolor!80}\\\\
\\hline
美林&布洛芬&\\parbox[c]{0.5cm}{ \\includegraphics[scale=0.12]{scriptdir/pics/right-icon.png}}疗效较好&遵医嘱正常服用\\\\
\\hline
\\end{tabular}
}
\\end{flushleft}
\\vspace{0.2cm}
\\begin{spacing}{1.3}
\\noindent\\normalsize{
\\textcolor{tabcolor}{\\Large \\ding{43}} \\uline{\\textbf{\\large{中文名}}} \\\\ 药物生产商自主制定, 经药品监督管理部门核准的产品名称。\\\\\\\\
\\textcolor{tabcolor}{\\Large \\ding{43}} \\uline{\\textbf{\\large {主成分}}}\\\\药物内治疗相关适应症的主要效应成分。\\\\\\\\
\\textcolor{tabcolor}{\\Large \\ding{43}} \\uline{\\textbf{\\large {遗传评估及服用指导}}} \\\\ 评估结果为本公司根据您的基因检测结果, 从遗传方面评估您对某种药物的临床治疗效果和不良反应风险, 并提供一定的用药参考建议, 若您希望调整用药方案请务必提前咨询您的医生。}
\\end{spacing}

\\newpage
\\ \\vspace{1cm}
\\begin{flushleft}
\\textcolor{black}{\\LARGE{\\textbf{样本信息}}}
\\end{flushleft}
\\begin{flushleft}
\\normalsize{
\\renewcommand{\\arraystretch}{1.5}
\\setlength{\\arrayrulewidth}{0.2em}
\\arrayrulecolor{tabcolor!80}
\\begin{tabular}{p{12cm}cp{12cm}}
\\hline
`

func main() {
	argsWithProg := os.Args
	infile := argsWithProg[1]
	prefix := strings.Split(infile, "_")
	fdir := strings.Join(prefix, "/")
	bundledir := outdir + "/" + fdir + "/out/" + infile
	fmt.Printf(">>%s\n", bundledir)

	m, err := app.ReadAndParse(infile)
	if err != nil {
		panic(err)
	}
	genotype := make(map[string]string)
	for k, v := range m {
		s := strings.Split(v, "\t")
		name := s[0]
		sex := s[1]
		birthd := s[2]
		id := k
		cusinfo := `\\textbf{姓名(NAME)}&` + name + `\\\\
\\textbf{性别(GENDER)}&` + sex + `\\\\
\\textbf{出生年月(BIRTHDAY)}&` + birthd + `\\\\
\\textbf{样本编号(SAMPLE ID）}&` + id + `\\\\
\\textbf{报告日期(ISSUING DATE)}&\\today\\\\`

		genotype["rs28371759"] = app.Baset(s[4])
		genotype["rs717620"] = app.Baset(s[5])
		genotype["rs4244285"] = app.Baset(s[6])
		genotype["rs4986893"] = app.Baset(s[7])
		genotype["rs1065852"] = app.Baset(s[8])
		genotype["rs776746"] = app.Baset(s[9])
		genotype["rs1127354"] = app.Baset(s[10])
		genotype["rs10929303"] = app.Baset(s[11])
		genotype["rs20417"] = app.Baset(s[12])
		genotype["rs1057910"] = app.Baset(s[13])
		genotype["rs1042713"] = app.Baset(s[14])
		genotype["rs1876828"] = app.Baset(s[15])
		genotype["rs1801131"] = app.Baset(s[16])
		genotype["rs111033313"] = app.Baset(s[17])
		genotype["rs121908362"] = app.Baset(s[18])
		genotype["rs267606617"] = app.Baset(s[19])
		genotype["rs267606619"] = app.Baset(s[20])

		newdir := out + id
		err := os.MkdirAll(newdir, 0777)
		if err != nil {
			fmt.Printf("%s", err)
		} else {
			fmt.Print("Create Directory " + newdir + " OK!")
		}
		out := outdir + id + id + ".tex"
		file, err := os.OpenFile(out, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println("open file failed.", err.Error())
			break
		}
		defer file.Close()
		//    file.WriteString(v)
		truegeno := app.TrueGeno(s[6], genotype)

		fmt.Printf("%s ===> \n%s\n", k, v)
	}
}
