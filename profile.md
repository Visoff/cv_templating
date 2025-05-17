# {{.FirstName}} <span style="color: #2B8CC4">{{.LastName}}</span>
### {{.Title}} [![Visitors](https://api.visitorbadge.io/api/visitors?path=https://github.com/{{.GitHub.Username}}&label=Visitors&countColor=%232B8CC4)](https://visitorbadge.io/status?path=https://github.com/{{.GitHub.Username}})

<p align="left">
  {{if .Email}}📧 <a href="mailto:{{.Email}}">{{.Email}}</a>{{end}}
  {{if .Phone}}📱 {{.Phone}}{{end}}
  {{if .Website}}🌐 <a href="{{.Website}}">{{.Website}}</a>{{end}}
</p>

---

## 🧑💻 Обо мне
{{.Summary}}

<div align="center">
<img height="180em" src="https://github-readme-stats.vercel.app/api?username={{.GitHub.Username}}&show_icons=true&theme=transparent&hide_border=true&title_color=2B8CC4&icon_color=2B8CC4&text_color=000000"/>
<img height="180em" src="https://github-readme-stats.vercel.app/api/top-langs/?username={{.GitHub.Username}}&layout=compact&theme=transparent&hide_border=true&title_color=2B8CC4"/>
</div>

---

## 🏆 Достижения
![Alt text](https://github.com/{{.GitHub.Username}}/{{.GitHub.Username}}/blob/main/trophy.svg?raw=true)

---

## 🛠️ Навыки
{{range $skill := .Skills}}
- **{{$skill.Name}}**  
  {{if eq $skill.Level 1}}`🟩⬜⬜⬜⬜` Начинающий
  {{else if eq $skill.Level 2}}`🟩🟩⬜⬜⬜` Средний
  {{else if eq $skill.Level 3}}`🟩🟩🟩⬜⬜` Продвинутый
  {{else if eq $skill.Level 4}}`🟩🟩🟩🟩⬜` Эксперт
  {{else if eq $skill.Level 5}}`🟩🟩🟩🟩🟩` Профессионал{{end}}
{{end}}

---

## 🌍 Языки
{{range $lang := .Languages}}
- **{{$lang.Name}}**  
  {{if eq $lang.Level 1}}`◯◯◯◯◯` Начинающий
  {{else if eq $lang.Level 2}}`◉◯◯◯◯` Средний
  {{else if eq $lang.Level 3}}`◉◉◯◯◯` Продвинутый
  {{else if eq $lang.Level 4}}`◉◉◉◯◯` Эксперт
  {{else if eq $lang.Level 5}}`◉◉◉◉◉` Профессионал{{end}}
{{end}}

---

## 💼 Опыт
{{range $exp := .Experience}}
### {{$exp.Position}}
**{{$exp.Company}}**  
*{{$exp.StartDate}} - {{if $exp.EndDate}}{{$exp.EndDate}}{{else}}Настоящее время{{end}}*

{{range $desc := $exp.Description}}
- {{$desc}}
{{end}}

{{end}}

---

## 🎓 Образование
{{range $edu := .Education}}
### {{$edu.Degree}}
**{{$edu.Institution}}**  
{{if $edu.FieldOfStudy}}*{{$edu.FieldOfStudy}}*  {{end}}
{{$edu.StartDate}} - {{$edu.EndDate}}

{{end}}

---

## 🚀 Проекты
{{range $project := .Projects}}
### {{$project.Name}}
{{if $project.URL}}🔗 [Ссылка на проект]({{$project.URL}})  {{end}}
{{if $project.Technologies}}🛠 **Технологии:** {{range $, $tech := $project.Technologies}}{{$tech}}, {{end}}{{end}}

{{range $desc := $project.Description}}
- {{$desc}}
{{end}}

{{end}}
