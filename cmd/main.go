package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func main() {
	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the text to translate.
	//text := "第一条 为规范申请人与国家食品药品监督管理总局药品审评中心（以下简称药审中心）之间的沟通交流，根据《国务院关于改革药品医疗器械审评审批制度的意见》（国发〔2015〕44号）等有关规定，制定本办法。"
	/*
		text := `一、药物研发基本情况
		1.申请人
		2.药品名称
		3.受理号（如适用）
		4.化学名称和结构（中药为处方）
		5.拟定适应症（或功能主治）
		6.剂型、给药途径和给药方法（用药频率和疗程）
		7.药物研发策略，包括药物研发背景资料、药物研发计划、研发过程的简要描述和关键事件、目前研发状态等。
		二、会议资料具体内容
		1.会议目的：简要说明。
		2.会议议程：列出会议议程。
		3.申请人参会名单：列出参会人员名单，包括职务、工作内容和工作单位。如果申请人拟邀请专家和翻译参会，应一并列出。
		4.讨论问题清单：申请人最终确定的问题列表。建议申请人按学科进行分类，包括但不限于从药学、药理毒理学和临床试验方案的设计等方面提出问题，每个问题应该包括简短的研发背景解释和该问题提出的目的。
		5.支持性数据总结：按学科和问题顺序总结支持性数据。
		支持性数据总结，应当用数据说明相关研究、结果和结论。以Ⅱ期临床试验结束会议为例，临床专业总结应包括下述内容：（1）应提供已完成的临床试验的简要总结，包括数据、结果与结论，同时应包括重要的剂量效应关系信息，一般情况下不需要提供完整的临床试验报告；（2）应对拟开展的Ⅲ期临床试验方案进行详细说明，以确认临床试验的主要特征，如临床试验受试者人群、关键的入选与排除标准、临床试验设计（如随机、盲法、对照选择，如果采用非劣效性试验，非劣效性界值设定依据）、给药剂量选择、主要和次要疗效终点、主要分析方法（包括计划的中期分析、适应性研究特征和主要安全性担忧）等。`
	*/
	text := "你好"
	// Sets the target language.
	target, err := language.Parse("en")
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, nil)
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	fmt.Printf("Text: %v\n", text)
	fmt.Printf("Translation: %v\n", translations[0].Text)
}
