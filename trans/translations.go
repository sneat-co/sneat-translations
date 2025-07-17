package trans

/*
Expected IDs in proper order of locale keys in var TRANS, use it as a reference for all values:
- arEG:
- de
- en:
- es:
- faIR:
- fr:
- id:
- it:
- jaJP:
- koKR:
- pl:
- ptBR:
- ru:
- tr:
- ukUA:
- uz:
- zhCN:
*/

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		arEG: "عينة",
		de:   "BEISPIEL",
		en:   "SAMPLE",
		es:   "EJEMPLO",
		faIR: "نمونه",
		fr:   "EXEMPLE",
		id:   "CONTOH",
		it:   "ESEMPIO",
		jaJP: "例",
		koKR: "예",
		pl:   "PRZYKŁAD",
		ptBR: "EXEMPLO",
		ptPT: "AMOSTRA",
		ru:   "ПРИМЕР",
		tr:   "ÖRNEK",
		ukUA: "ПРИКЛАД",
		uz:   "MISOL",
		zhCN: "例子",
	},

	HowdyUser: {
		arEG: "مرحباً {USER_NAME}!",
		de:   "Hallo {USER_NAME}!",
		en:   "Howdy {USER_NAME}!",
		es:   "¡Hola {USER_NAME}!",
		faIR: "سلام {USER_NAME}!",
		fr:   "Salut {USER_NAME} !",
		id:   "Halo {USER_NAME}!",
		it:   "Ciao {USER_NAME}!",
		jaJP: "やあ {USER_NAME}!",
		koKR: "안녕 {USER_NAME}!",
		pl:   "Cześć {USER_NAME}!",
		ptBR: "Olá {USER_NAME}!",
		ptPT: "Olá {USER_NAME}!",
		ru:   "Привет {USER_NAME}!",
		tr:   "Merhaba {USER_NAME}!",
		ukUA: "Привіт {USER_NAME}!",
		uz:   "Salom {USER_NAME}!",
		zhCN: "你好 {USER_NAME}！",
	},
	StepNumberOfNumberOfSteps: {
		en: "Step #{STEP_NUMBER} of {NUMBER_OF_STEPS}",
		ru: "Шаг #{STEP_NUMBER} из {NUMBER_OF_STEPS}",
	},
	ButtonAdd: {
		arEG: "يضيف",
		de:   "Hinzufügen",
		en:   "Add",
		es:   "Añadir",
		faIR: "افزودن",
		fr:   "Ajouter",
		id:   "Tambah",
		it:   "Aggiungi",
		jaJP: "追加",
		koKR: "추가",
		pl:   "Dodaj",
		ptBR: "Adicionar",
		ptPT: "Adicionar",
		ru:   "Добавить",
		tr:   "Ekle",
		ukUA: "Додати",
		uz:   "Qo'shish",
		zhCN: "添加",
	},
	ButtonRemove: {
		arEG: "يزيل",
		de:   "Entfernen",
		en:   "Remove",
		es:   "Eliminar",
		faIR: "حذف",
		fr:   "Supprimer",
		id:   "Hapus",
		it:   "Rimuovi",
		jaJP: "削除",
		koKR: "제거",
		pl:   "Usuń",
		ptBR: "Remover",
		ptPT: "Remover",
		ru:   "Удалить",
		tr:   "Kaldır",
		ukUA: "Видалити",
		uz:   "Olib tashlash",
		zhCN: "移除",
	},

	"Monday": {
		en: "Monday",
		ru: "Понедельник",
	},
	"Tuesday": {
		en: "Tuesday",
		ru: "Вторник",
	},
	"Wednesday": {
		en: "Wednesday",
		ru: "Среда",
	},
	"Thursday": {
		en: "Thursday",
		ru: "Четрвег",
	},
	"Friday": {
		en: "Friday",
		ru: "Пятница",
	},
	"Saturday": {
		en: "Saturday",
		ru: "Суббота",
	},
	"Sunday": {
		en: "Sunday",
		ru: "Воскресенье",
	},

	"Jan": {
		arEG: "يناير",
		de:   "Jan",
		en:   "Jan",
		es:   "Ene",
		faIR: "ژانویه",
		fr:   "Jan",
		id:   "Januari",
		it:   "Gen",
		jaJP: "1月",
		koKR: "1월",
		pl:   "Sty",
		ptBR: "Janeiro",
		ptPT: "Janeiro",
		ru:   "Янв.",
		tr:   "Oca",
		ukUA: "Січ",
		uz:   "Yan",
		zhCN: "一月",
	},

	"Feb": {
		arEG: "فبراير",
		de:   "Februar",
		en:   "Feb",
		es:   "Feb",
		faIR: "فوریه",
		fr:   "Fév",
		id:   "Februari",
		it:   "Febbraio",
		jaJP: "2月",
		koKR: "2월",
		pl:   "Lut",
		ptBR: "Fev",
		ptPT: "Fev",
		ru:   "Фев.",
		tr:   "Şub",
		ukUA: "Лют",
		uz:   "Fev",
		zhCN: "二月",
	},

	"Mar": {
		arEG: "مارس",
		de:   "Mär",
		en:   "Mar",
		es:   "Mar",
		faIR: "مارس",
		fr:   "Mars",
		id:   "Merusak",
		it:   "Marzo",
		jaJP: "3月",
		koKR: "3월",
		pl:   "Zniszczyć",
		ptBR: "Março",
		ptPT: "Março",
		ru:   "Мврт",
		tr:   "Mart",
		ukUA: "Бер",
		uz:   "mart",
		zhCN: "三月",
	},

	"Apr": {
		arEG: "أبريل",
		de:   "April",
		en:   "Apr",
		es:   "Abr",
		faIR: "آوریل",
		fr:   "Avril",
		id:   "April",
		it:   "Aprile",
		jaJP: "4月",
		koKR: "4월",
		pl:   "Kwi",
		ptBR: "Abr",
		ptPT: "Abr",
		ru:   "Апр.",
		tr:   "Nis",
		ukUA: "Кві",
		uz:   "aprel",
		zhCN: "四月",
	},

	"May": {
		arEG: "يمكن",
		de:   "Mai",
		en:   "May",
		es:   "Puede",
		faIR: "مه",
		fr:   "Mai",
		id:   "Mei",
		it:   "Mag",
		jaJP: "5月",
		koKR: "5월",
		pl:   "Maj",
		ptBR: "Maio",
		ptPT: "Maio",
		ru:   "Май",
		tr:   "Mayıs",
		ukUA: "Тра",
		uz:   "may",
		zhCN: "五月",
	},

	"Jun": {
		arEG: "يونيو",
		de:   "Jun",
		en:   "June",
		es:   "Jun",
		faIR: "ژوئن",
		fr:   "Juin",
		id:   "Jun",
		it:   "Giu",
		jaJP: "6月",
		koKR: "6월",
		pl:   "Cze",
		ptBR: "Junho",
		ptPT: "Junho",
		ru:   "Июнь",
		tr:   "Haz",
		ukUA: "Чер",
		uz:   "Iyun",
		zhCN: "六月",
	},

	"Jul": {
		arEG: "يوليو",
		de:   "Jul",
		en:   "July",
		es:   "Jul",
		faIR: "ژوئیه",
		fr:   "Juil",
		id:   "Jul",
		it:   "Lug",
		jaJP: "7月",
		koKR: "7월",
		pl:   "Lip",
		ptBR: "Julho",
		ptPT: "Julho",
		ru:   "Июль",
		tr:   "Tem",
		ukUA: "Лип",
		uz:   "Iyul",
		zhCN: "七月",
	},

	"Aug": {
		arEG: "أغسطس",
		de:   "August",
		en:   "Aug",
		es:   "Ago",
		faIR: "اوت",
		fr:   "Août",
		id:   "Agu",
		it:   "Ago",
		jaJP: "8月",
		koKR: "8월",
		pl:   "Sie",
		ptBR: "Ago",
		ptPT: "Ago",
		ru:   "Авг.",
		tr:   "Ağu",
		ukUA: "Сер",
		uz:   "Avg",
		zhCN: "八月",
	},

	"Sep": {
		arEG: "سبتمبر",
		de:   "September",
		en:   "Sep",
		es:   "Sep",
		faIR: "سپتامبر",
		fr:   "Sep",
		id:   "September",
		it:   "Sett",
		jaJP: "9月",
		koKR: "9월",
		pl:   "Wrz",
		ptBR: "Set",
		ptPT: "Set",
		ru:   "Сен.",
		tr:   "Eyl",
		ukUA: "Вер",
		uz:   "Sen",
		zhCN: "九月",
	},

	"Oct": {
		arEG: "أكتوبر",
		de:   "Okt",
		en:   "Oct",
		es:   "Oct",
		faIR: "اکتبر",
		fr:   "Octobre",
		id:   "Okt",
		it:   "Ott",
		jaJP: "10月",
		koKR: "10월",
		pl:   "Paź",
		ptBR: "Out",
		ptPT: "Out",
		ru:   "Окт.",
		tr:   "Eki",
		ukUA: "Жов",
		uz:   "Okt",
		zhCN: "十月",
	},

	"Nov": {
		arEG: "نوفمبر",
		de:   "November",
		en:   "Nov",
		es:   "Nov",
		faIR: "نوامبر",
		fr:   "Nov",
		id:   "November",
		it:   "Novembre",
		jaJP: "11月",
		koKR: "11월",
		pl:   "Lis",
		ptBR: "novembro",
		ptPT: "novembro",
		ru:   "Нбр.",
		tr:   "Kas",
		ukUA: "Лис",
		uz:   "Noy",
		zhCN: "十一月",
	},

	"Dec": {
		arEG: "ديسمبر",
		de:   "Dez",
		en:   "Dec",
		es:   "Dic",
		faIR: "دسامبر",
		fr:   "Déc",
		id:   "Des",
		it:   "Dic",
		jaJP: "12月",
		koKR: "12월",
		pl:   "Gru",
		ptBR: "Dez",
		ptPT: "Dez",
		ru:   "Дек.",
		tr:   "Ara",
		ukUA: "Гру",
		uz:   "Dek",
		zhCN: "十二月",
	},
	COMMAND_START: {
		arEG: "يبدأ",
		de:   "Start",
		en:   "start",
		es:   "inicio",
		faIR: "شروع",
		fr:   "démarrer",
		id:   "mulai",
		it:   "inizio",
		jaJP: "スタート",
		koKR: "시작",
		pl:   "start",
		ptBR: "iniciar",
		ptPT: "iniciar",
		ru:   "старт",
		tr:   "başlat",
		ukUA: "старт",
		uz:   "boshlash",
		zhCN: "开始",
	},
	COMMAND_MENU: {
		arEG: "قائمة طعام",
		de:   "Speisekarte",
		en:   "menu",
		es:   "menú",
		faIR: "منو",
		fr:   "menu",
		id:   "menu",
		it:   "menu",
		jaJP: "メニュー",
		koKR: "메뉴",
		pl:   "menu",
		ptBR: "menu",
		ptPT: "menu",
		ru:   "меню",
		tr:   "menü",
		ukUA: "меню",
		uz:   "menyu",
		zhCN: "菜单",
	},
	COMMAND_GAVE: {
		arEG: "أعطى",
		de:   "verleihen",
		en:   "gave",
		es:   "prestado_a_ti",
		faIR: "قرض_دادن",
		fr:   "donné",
		id:   "memberi",
		it:   "debito",
		jaJP: "貸した",
		koKR: "주었다",
		pl:   "dał",
		ptBR: "emprestou",
		ptPT: "deu",
		ru:   "дать",
		tr:   "verdi",
		ukUA: "дати",
		uz:   "berdi",
		zhCN: "给予",
	},
	COMMAND_GOT: {
		arEG: "يملك",
		de:   "anleihen",
		en:   "got",
		es:   "prestado_por_ti",
		faIR: "قرض_گرفتن",
		fr:   "reçu",
		id:   "mendapat",
		it:   "credito",
		jaJP: "借りた",
		koKR: "받았다",
		pl:   "dostał",
		ptBR: "recebeu",
		ptPT: "ter",
		ru:   "взять",
		tr:   "aldı",
		ukUA: "взяти",
		uz:   "oldi",
		zhCN: "得到",
	},
	COMMAND_RETURNED: {
		arEG: "يعود",
		de:   "beglichen",
		en:   "return",
		es:   "devuelto",
		faIR: "بازگردانده_شده",
		fr:   "retourné",
		id:   "kembali",
		it:   "rientra",
		jaJP: "返済",
		koKR: "반환",
		pl:   "zwrócił",
		ptBR: "devolveu",
		ptPT: "devolver",
		ru:   "вернуть",
		tr:   "iade",
		ukUA: "повернути",
		uz:   "qaytardi",
		zhCN: "归还",
	},
	COMMAND_BALANCE: {
		arEG: "توازن",
		de:   "ausstehend",
		en:   "balance",
		es:   "balance",
		faIR: "تراز",
		fr:   "solde",
		id:   "saldo",
		it:   "bilancio",
		jaJP: "残高",
		koKR: "잔액",
		pl:   "saldo",
		ptBR: "saldo",
		ptPT: "equilíbrio",
		ru:   "баланс",
		tr:   "bakiye",
		ukUA: "баланс",
		uz:   "balans",
		zhCN: "余额",
	},
	COMMAND_HISTORY: {
		arEG: "تاريخ",
		de:   "verlauf",
		en:   "history",
		es:   "cronología",
		faIR: "سوابق",
		fr:   "historique",
		id:   "riwayat",
		it:   "cronologia",
		jaJP: "履歴",
		koKR: "기록",
		pl:   "historia",
		ptBR: "histórico",
		ptPT: "histórico",
		ru:   "история",
		tr:   "geçmiş",
		ukUA: "історія",
		uz:   "tarix",
		zhCN: "历史",
	},
	COMMAND_SETTINGS: {
		arEG: "إعدادات",
		de:   "einstellungen",
		en:   "settings",
		es:   "ajustes",
		faIR: "تنظیمات",
		fr:   "paramètres",
		id:   "pengaturan",
		it:   "impostazioni",
		jaJP: "設定",
		koKR: "설정",
		pl:   "ustawienia",
		ptBR: "configurações",
		ptPT: "configurações",
		ru:   "настройки",
		tr:   "ayarlar",
		ukUA: "налаштування",
		uz:   "sozlamalar",
		zhCN: "设置",
	},
	COMMAND_HELP: {
		arEG: "يساعد",
		de:   "hilfe",
		en:   "help",
		es:   "ayuda",
		faIR: "کمک",
		fr:   "aide",
		id:   "bantuan",
		it:   "aiuto",
		jaJP: "ヘルプ",
		koKR: "도움말",
		pl:   "pomoc",
		ptBR: "ajuda",
		ptPT: "ajuda",
		ru:   "помощь",
		tr:   "yardım",
		ukUA: "допомога",
		uz:   "yordam",
		zhCN: "帮助",
	},
	COMMAND_CANCEL: {
		arEG: "يلغي",
		de:   "abbrechen",
		en:   "cancel",
		es:   "cancelar",
		faIR: "کنسل",
		fr:   "annuler",
		id:   "batal",
		it:   "annulla",
		jaJP: "キャンセル",
		koKR: "취소",
		pl:   "anuluj",
		ptBR: "cancelar",
		ptPT: "Cancelar",
		ru:   "/отменить",
		tr:   "iptal",
		ukUA: "скасувати",
		uz:   "bekor",
		zhCN: "取消",
	},
	COMMAND_CLEAR: {
		arEG: "واضح",
		de:   "leeren",
		en:   "clear",
		es:   "borrar",
		faIR: "پاک_کردن",
		fr:   "effacer",
		id:   "bersihkan",
		it:   "chiaro",
		jaJP: "クリア",
		koKR: "지우기",
		pl:   "wyczyść",
		ptBR: "limpar",
		ptPT: "claro",
		ru:   "очистить",
		tr:   "temizle",
		ukUA: "очистити",
		uz:   "tozalash",
		zhCN: "清除",
	},
	"Ads": {
		arEG: "إعلانات",
		de:   "Anzeigen",
		en:   "Ads",
		es:   "Anuncios",
		faIR: "تبلیغات",
		fr:   "Publicités",
		id:   "Iklan",
		it:   "Annunci",
		jaJP: "広告",
		koKR: "광고",
		pl:   "Reklamy",
		ptBR: "Anúncios",
		ptPT: "Anúncios",
		ru:   "Реклама",
		tr:   "Reklamlar",
		ukUA: "Оголошення",
		uz:   "E&#39;lonlar",
		zhCN: "广告"},
	" and ": {
		arEG: " و ",
		de:   " und ",
		en:   " and ",
		es:   " y ",
		faIR: " و ",
		fr:   " et ",
		id:   " dan ",
		it:   " e ",
		jaJP: " と ",
		koKR: " 그리고 ",
		pl:   " i ",
		ptBR: " e ",
		ptPT: " e ",
		ru:   " и ",
		tr:   " ve ",
		ukUA: " і ",
		uz:   " va ",
		zhCN: " 和 ",
	},
	"MessageTextOopsSomethingWentWrong": {
		arEG: "عفواً، حدث خطأ ما... \xF0\x9F\x98\xB3",
		de:   "Ups, etwas ist schiefgelaufen... \xF0\x9F\x98\xB3",
		en:   "Oops, something went wrong... \xF0\x9F\x98\xB3",
		es:   "Ops,  algo ha salido mal... \xF0\x9F\x98\xB3",
		faIR: "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
		fr:   "Oups, quelque chose s'est mal passé... \xF0\x9F\x98\xB3",
		id:   "Ups, ada yang salah... \xF0\x9F\x98\xB3",
		it:   "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
		jaJP: "おっと、何か問題が発生しました... \xF0\x9F\x98\xB3",
		koKR: "이런, 문제가 발생했습니다... \xF0\x9F\x98\xB3",
		pl:   "Ups, coś poszło nie tak... \xF0\x9F\x98\xB3",
		ptBR: "Ops, algo deu errado... \xF0\x9F\x98\xB3",
		ptPT: "Ups, algo correu mal... \xF0\x9F\x98\xB3",
		ru:   "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		tr:   "Hata, bir şeyler yanlış gitti... \xF0\x9F\x98\xB3",
		ukUA: "Упс, щось пішло не так... \xF0\x9F\x98\xB3",
		uz:   "Xatolik, nimadir noto'g'ri ketdi... \xF0\x9F\x98\xB3",
		zhCN: "哎呀，出了点问题... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		arEG: "متى هو تاريخ الاستحقاق؟",
		de:   "Wann ist die Schuld fällig?",
		en:   "When is the due date?",
		es:   "¿Cuándo es la fecha de devolución?",
		faIR: "سررسید چه زمانی است؟",
		fr:   "Quelle est la date d'échéance?",
		id:   "Kapan tanggal jatuh tempo?",
		it:   "Data di scadenza?",
		jaJP: "期日はいつですか？",
		koKR: "만기일은 언제입니까?",
		pl:   "Kiedy jest termin płatności?",
		ptBR: "Qual é a data de vencimento?",
		ptPT: "Quando é a data de validade?",
		ru:   "Когда дата возврата?",
		tr:   "Vade tarihi ne zaman?",
		ukUA: "Коли дата повернення?",
		uz:   "To'lov muddati qachon?",
		zhCN: "到期日期是什么时候？",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		arEG: `لتحديد تاريخ التذكير التالي، يرجى إرساله كنص بتنسيق <i>DD.MM.YEAR</i> .
 <b>على سبيل المثال،</b> بالنسبة لـ 20 يناير 2017، أرسل:
 <i>20.01.2017</i>`,
		de: `Sende mir das Datum, an welches du <b>erneut</b> erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,
		en: `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,
		es: `Para establecer la fecha recordatoria escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
    <i>20.01.2017</i>`,
		faIR: `لطفاً برای تنظیم یادآور بعدی آنرا با متنی با این فرمت ارسال نمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 15 خرداد 1396 ثبت کنید:
    <i>15.03.1396</i>`,
		fr: `Pour définir la date du prochain rappel, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
    <i>20.01.2017</i>`,
		id: `Untuk mengatur tanggal pengingat berikutnya, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
    <i>20.01.2017</i>`,
		it: `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
    <i>20.01.2017</i>`,
		jaJP: `次のリマインダーの日付を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
    <i>20.01.2017</i>`,
		koKR: `다음 알림의 날짜를 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
    <i>20.01.2017</i>`,
		pl: `Aby ustawić datę następnego przypomnienia, wyślij ją jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
    <i>20.01.2017</i>`,
		ptBR: `Para definir a data do próximo lembrete, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
    <i>20.01.2017</i>`,
		ptPT: `Para definir a data do próximo lembrete, envie-o como um texto no formato <i>DD.MM.ANO</i> .
 <b>Por exemplo,</b> para 20 de janeiro de 2017, envie:
 <i>20.01.2017</i>`,
		ru: `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,
		tr: `Bir sonraki hatırlatma için tarih belirlemek üzere <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
    <i>20.01.2017</i>`,
		ukUA: `Щоб встановити дату наступного нагадування, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
    <i>20.01.2017</i>`,
		uz: `Keyingi eslatma uchun sanani belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
    <i>20.01.2017</i>`,
		zhCN: `要设置下一次提醒的日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		arEG: `لتحديد تاريخ الاستحقاق، يرجى إرساله كنص بتنسيق <i>DD.MM.YEAR</i> .
 <b>على سبيل المثال</b> ، بالنسبة لتاريخ 20 يناير 2017، أرسل:
 <i>20.01.2017</i>`,
		de: `Sende mir das Datum, an welches du erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,
		en: `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
<i>20.01.2017</i>`,
		es: `Para establecer la fecha de devolución escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
<i>20.01.2017</i>`,
		faIR: `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید.<i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
<i>20.01.2017</i>`,
		fr: `Pour définir la date d'échéance, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
<i>20.01.2017</i>`,
		id: `Untuk mengatur tanggal jatuh tempo, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
<i>20.01.2017</i>`,
		it: `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
<i>20.01.2017</i>`,
		jaJP: `期日を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
<i>20.01.2017</i>`,
		koKR: `만기일을 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
<i>20.01.2017</i>`,
		pl: `Aby ustawić termin płatności, wyślij go jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
<i>20.01.2017</i>`,
		ptBR: `Para definir a data de vencimento, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
<i>20.01.2017</i>`,
		ptPT: `Para definir a data de vencimento, envie como texto no formato <i>DD.MM.ANO</i> .
 <b>Por exemplo,</b> para 20 de janeiro de 2017, envie:
 <i>20.01.2017</i>`,
		ru: `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г.отправьте:
<i>20.01.2017</i>`,
		tr: `Vade tarihini ayarlamak için <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
<i>20.01.2017</i>`,
		ukUA: `Щоб встановити дату повернення, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
<i>20.01.2017</i>`,
		uz: `To'lov muddatini belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
<i>20.01.2017</i>`,
		zhCN: `要设置到期日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
<i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		arEG: "عذرا، هناك خطأ ما في التاريخ الذي قدمته.",
		de:   "Entschuldigung, aber mit diesem Datum stimmt etwas nicht.",
		en:   "Sorry, there is something wrong with the date you've provided.",
		es:   "Lo siento, algo no es correcto con la fecha que has puesto",
		faIR: "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
		fr:   "Désolé, il y a un problème avec la date que vous avez fournie.",
		id:   "Maaf, ada yang salah dengan tanggal yang Anda berikan.",
		it:   "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
		jaJP: "申し訳ありませんが、提供された日付に問題があります。",
		koKR: "죄송합니다. 제공하신 날짜에 문제가 있습니다.",
		pl:   "Przepraszamy, ale z podaną datą jest coś nie tak.",
		ptBR: "Desculpe, há algo errado com a data que você forneceu.",
		ptPT: "Desculpe, há algo de errado com a data que forneceu.",
		ru:   "Извините, что-то не так с датой которую вы отправили.",
		tr:   "Üzgünüz, verdiğiniz tarihte bir sorun var.",
		ukUA: "Вибачте, але з датою щось не так.",
		uz:   "Kechirasiz, siz taqdim etgan sana bilan bog'liq muammo bor.",
		zhCN: "抱歉，您提供的日期有问题。",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		arEG: "لا يوجد تذكير",
		de:   "Nicht erinnern",
		en:   "No reminder",
		es:   "No recordar",
		faIR: "بدون یادآور",
		fr:   "Pas de rappel",
		id:   "Tidak ada pengingat",
		it:   "Nessun promemoria",
		jaJP: "リマインダーなし",
		koKR: "알림 없음",
		pl:   "Bez przypomnienia",
		ptBR: "Sem lembrete",
		ptPT: "Sem lembrete",
		ru:   "Не напоминать",
		tr:   "Hatırlatma yok",
		ukUA: "Не нагадувати",
		uz:   "Eslatma yo'q",
		zhCN: "无提醒",
	},
	COMMAND_TEXT_TOMORROW: {
		arEG: "غداً",
		de:   "Morgen",
		en:   "Tomorrow",
		es:   "Mañana",
		faIR: "فردا",
		fr:   "Demain",
		id:   "Besok",
		it:   "Domani",
		jaJP: "明日",
		koKR: "내일",
		pl:   "Jutro",
		ptBR: "Amanhã",
		ptPT: "Amanhã",
		ru:   "Завтра",
		tr:   "Yarın",
		ukUA: "Завтра",
		uz:   "Ertaga",
		zhCN: "明天",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		arEG: "بعد غد",
		de:   "Übermorgen",
		en:   "Day after tomorrow",
		es:   "Pasada mañana",
		faIR: "پس فردا",
		fr:   "Après-demain",
		id:   "Lusa",
		it:   "Dopo domani",
		jaJP: "明後日",
		koKR: "모레",
		pl:   "Pojutrze",
		ptBR: "Depois de amanhã",
		ptPT: "Depois de amanhã",
		ru:   "Послезавтра",
		tr:   "Öbür gün",
		ukUA: "Післязавтра",
		uz:   "Indinga",
		zhCN: "后天",
	},
	COMMAND_TEXT_THIS_WEEK: {
		arEG: "هذا الاسبوع",
		de:   "Diese Woche",
		en:   "This week",
		es:   "Esta semana",
		faIR: "این هفته",
		fr:   "Cette semaine",
		id:   "Minggu ini",
		it:   "Questa settimana",
		jaJP: "今週",
		koKR: "이번 주",
		pl:   "W tym tygodniu",
		ptBR: "Esta semana",
		ptPT: "Esta semana",
		ru:   "На этой неделе",
		tr:   "Bu hafta",
		ukUA: "Цього тижня",
		uz:   "Shu hafta",
		zhCN: "本周",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		arEG: "نعم، لديه موعد نهائي!",
		de:   "Ja, es hat eine Frist!",
		en:   "Yes, it has a deadline!",
		es:   "Sí, hay una fecha de devolución!",
		faIR: "بله، دارای آخرین فرصت می باشد!",
		fr:   "Oui, il y a une date limite !",
		id:   "Ya, ada batas waktunya!",
		it:   "Si, c'e' una data di scadenza",
		jaJP: "はい、期限があります！",
		koKR: "네, 마감일이 있어요!",
		pl:   "Tak, ma termin!",
		ptBR: "Sim, tem prazo!",
		ptPT: "Sim, tem um prazo!",
		ru:   "Да, есть срок возврата!",
		tr:   "Evet, bir son tarihi var!",
		ukUA: "Так, у нього є термін!",
		uz:   "Ha, uning muddati bor!",
		zhCN: "是的，它有一个截止日期！"},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		arEG: "لا، متى ما كان ذلك مناسبًا.",
		de:   "Nein, sobald möglich.",
		en:   "No, whenever is fine.",
		es:   "No, sin fecha límite.",
		faIR: "خیر، هر زمانی مناسب است.",
		fr:   "Non, quand c&#39;est possible.",
		id:   "Tidak, kapan pun tidak masalah.",
		it:   "No, nessuna scadenza",
		jaJP: "いいえ、いつでも大丈夫です。",
		koKR: "아니요, 언제든 괜찮습니다.",
		pl:   "Nie, możesz to zrobić kiedykolwiek.",
		ptBR: "Não, qualquer hora está bom.",
		ptPT: "Não, qualquer hora é boa.",
		ru:   "Нет, срок возврата не важен.",
		tr:   "Hayır, ne zaman uygunsa.",
		ukUA: "Ні, коли завгодно, це нормально.",
		uz:   "Yo&#39;q, qachon yaxshi bo&#39;lsa.",
		zhCN: "不，什么时候都可以。"},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		arEG: "عندما يكون الأمر على ما يرام",
		de:   "Unbefristet",
		en:   "Whenever is fine",
		es:   "Cualquier día",
		faIR: "هر زمانی مناسب است.",
		fr:   "Quand c&#39;est bien",
		id:   "Kapanpun baik-baik saja",
		it:   "No, Nessuna scadenza",
		jaJP: "いつでも大丈夫です",
		koKR: "언제든지 괜찮아요",
		pl:   "Kiedykolwiek jest dobrze",
		ptBR: "Quando estiver bom",
		ptPT: "Quando estiver bom",
		ru:   "Когда-нибудь",
		tr:   "Ne zaman uygunsa",
		ukUA: "Коли завгодно",
		uz:   "Qachon yaxshi bo&#39;lsa",
		zhCN: "任何时候都可以"},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		arEG: "في بضع دقائق",
		de:   "In wenigen Minuten",
		en:   "In few minutes",
		es:   "En unos minutos",
		faIR: "در چند دقیقه",
		fr:   "Dans quelques minutes",
		id:   "Dalam beberapa menit",
		it:   "Fra alcuni minuti",
		jaJP: "数分で",
		koKR: "몇 분 안에",
		pl:   "Za kilka minut",
		ptBR: "Em poucos minutos",
		ptPT: "Em poucos minutos",
		ru:   "Через минуту",
		tr:   "Birkaç dakika içinde",
		ukUA: "За кілька хвилин",
		uz:   "Bir necha daqiqada",
		zhCN: "几分钟后"},
	COMMAND_TEXT_IN_1_WEEK: {
		arEG: "في اسبوع واحد",
		de:   "In einer Woche",
		en:   "In 1 week",
		es:   "En una semana",
		faIR: "ظرف یک هفته",
		fr:   "Dans 1 semaine",
		id:   "Dalam 1 minggu",
		it:   "Fra una settimana",
		jaJP: "1週間後",
		koKR: "1주일 안에",
		pl:   "Za 1 tydzień",
		ptBR: "Em 1 semana",
		ptPT: "Em 1 semana",
		ru:   "Через неделю",
		tr:   "1 hafta içinde",
		ukUA: "Через 1 тиждень",
		uz:   "1 hafta ichida",
		zhCN: "1周内"},
	COMMAND_TEXT_IN_1_MONTH: {
		arEG: "في شهر واحد",
		de:   "In einem Monat",
		en:   "In 1 month",
		es:   "En un mes",
		faIR: "ظرف یک ماه",
		fr:   "Dans 1 mois",
		id:   "Dalam 1 bulan",
		it:   "Fra un mese",
		jaJP: "1ヶ月後",
		koKR: "1개월 안에",
		pl:   "Za 1 miesiąc",
		ptBR: "Em 1 mês",
		ptPT: "Em 1 mês",
		ru:   "Через месяц",
		tr:   "1 ay içinde",
		ukUA: "Через 1 місяць",
		uz:   "1 oy ichida",
		zhCN: "1个月内"},
	COMMAND_TEXT_SET_DATE: {
		arEG: "حدد التاريخ",
		de:   "Datum setzen",
		en:   "Set date",
		es:   "Establecer la fecha",
		faIR: "تاریخ را تنظیم کنید",
		fr:   "Définir la date",
		id:   "Tetapkan tanggal",
		it:   "Imposta la data",
		jaJP: "日付を設定",
		koKR: "날짜 설정",
		pl:   "Ustaw datę",
		ptBR: "Definir data",
		ptPT: "Definir data",
		ru:   "Задать дату",
		tr:   "Tarih belirle",
		ukUA: "Встановити дату",
		uz:   "Sana belgilash",
		zhCN: "设置日期"},
	COMMAND_TEXT_MONDAY: {
		arEG: "الاثنين",
		de:   "Montag",
		en:   "Monday",
		es:   "Lunes",
		faIR: "دوشنبه",
		fr:   "Lundi",
		id:   "Senin",
		it:   "Lunedi'",
		jaJP: "月曜日",
		koKR: "월요일",
		pl:   "Poniedziałek",
		ptBR: "Segunda-feira",
		ptPT: "Segunda-feira",
		ru:   "Понедельник",
		tr:   "Pazartesi",
		ukUA: "Понеділок",
		uz:   "dushanba",
		zhCN: "周一"},
	COMMAND_TEXT_TUESDAY: {
		arEG: "يوم الثلاثاء",
		de:   "Dienstag",
		en:   "Tuesday",
		es:   "Martes",
		faIR: "سه شنبه",
		fr:   "Mardi",
		id:   "Selasa",
		it:   "Martedi'",
		jaJP: "火曜日",
		koKR: "화요일",
		pl:   "Wtorek",
		ptBR: "Terça-feira",
		ptPT: "Terça-feira",
		ru:   "Вторник",
		tr:   "Salı",
		ukUA: "Вівторок",
		uz:   "seshanba",
		zhCN: "周二"},
	COMMAND_TEXT_WEDNESDAY: {
		arEG: "الأربعاء",
		de:   "Mittwoch",
		en:   "Wednesday",
		es:   "Miercoles",
		faIR: "چهارشنبه",
		fr:   "Mercredi",
		id:   "Rabu",
		it:   "Mercoledi'",
		jaJP: "水曜日",
		koKR: "수요일",
		pl:   "Środa",
		ptBR: "Quarta-feira",
		ptPT: "Quarta-feira",
		ru:   "Среда",
		tr:   "Çarşamba",
		ukUA: "Середа",
		uz:   "chorshanba",
		zhCN: "周三"},
	COMMAND_TEXT_THURSDAY: {
		arEG: "يوم الخميس",
		de:   "Donnerstag",
		en:   "Thursday",
		es:   "Jueves",
		faIR: "پنج شنبه",
		fr:   "Jeudi",
		id:   "Kamis",
		it:   "Giovedi'",
		jaJP: "木曜日",
		koKR: "목요일",
		pl:   "Czwartek",
		ptBR: "Quinta-feira",
		ptPT: "Quinta-feira",
		ru:   "Четверг",
		tr:   "Perşembe",
		ukUA: "Четвер",
		uz:   "Payshanba",
		zhCN: "周四"},
	COMMAND_TEXT_FRIDAY: {
		arEG: "جمعة",
		de:   "Freitag",
		en:   "Friday",
		es:   "Viernes",
		faIR: "جمعه",
		fr:   "Vendredi",
		id:   "Jumat",
		it:   "Venerdi'",
		jaJP: "金曜日",
		koKR: "금요일",
		pl:   "Piątek",
		ptBR: "Sexta-feira",
		ptPT: "Sexta-feira",
		ru:   "Пятница",
		tr:   "Cuma",
		ukUA: "П&#39;ятниця",
		uz:   "Juma",
		zhCN: "星期五"},
	COMMAND_TEXT_SATURDAY: {
		arEG: "السبت",
		de:   "Samstag",
		en:   "Saturday",
		es:   "Sabado",
		faIR: "شنبه",
		fr:   "Samedi",
		id:   "Sabtu",
		it:   "Sabato",
		jaJP: "土曜日",
		koKR: "토요일",
		pl:   "Sobota",
		ptBR: "Sábado",
		ptPT: "Sábado",
		ru:   "Суббота",
		tr:   "Cumartesi",
		ukUA: "Субота",
		uz:   "shanba",
		zhCN: "周六"},
	COMMAND_TEXT_SUNDAY: {
		arEG: "الأحد",
		de:   "Sonntag",
		en:   "Sunday",
		es:   "Domingo",
		faIR: "یکشنبه",
		fr:   "Dimanche",
		id:   "Minggu",
		it:   "Domenica",
		jaJP: "日曜日",
		koKR: "일요일",
		pl:   "Niedziela",
		ptBR: "Domingo",
		ptPT: "Domingo",
		ru:   "Воскресенье",
		tr:   "Pazar",
		ukUA: "Неділя",
		uz:   "yakshanba",
		zhCN: "星期日"},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		arEG: "لا ترسل الإيصال",
		de:   "Keine Quittung schicken",
		en:   "Do not send the receipt",
		es:   "No enviar el recibo",
		faIR: "رسید را ارسال نکنید",
		fr:   "N&#39;envoyez pas le reçu",
		id:   "Jangan kirim tanda terimanya",
		it:   "Non inviare la ricevuta",
		jaJP: "領収書は送らないでください",
		koKR: "영수증을 보내지 마세요",
		pl:   "Nie wysyłaj paragonu",
		ptBR: "Não envie o recibo",
		ptPT: "Não envie o recibo",
		ru:   "Не отправлять квитанцию",
		tr:   "Makbuzu göndermeyin",
		ukUA: "Не надсилайте квитанцію",
		uz:   "Kvitansiyani yubormang",
		zhCN: "不要发送收据"},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		arEG: "لقد قررت عدم إرسال الإيصال.",
		de:   "Du hast dich gegen eine Quittung entschieden.",
		en:   "You've decided not to send the receipt.",
		es:   "Has decidido no enviar el recibo",
		faIR: "شما تصمیم گرفتید که رسید را ارسال نکنید.",
		fr:   "Vous avez décidé de ne pas envoyer le reçu.",
		id:   "Anda telah memutuskan untuk tidak mengirimkan tanda terima.",
		it:   "Hai scelto di non inviare la ricevuta",
		jaJP: "領収書を送らないことにしました。",
		koKR: "영수증을 보내지 않기로 결정했습니다.",
		pl:   "Zdecydowałeś się nie wysyłać paragonu.",
		ptBR: "Você decidiu não enviar o recibo.",
		ptPT: "Decidiu não enviar o recibo.",
		ru:   "Вы решили не отправлять квитанцию.",
		tr:   "Makbuzu göndermemeye karar verdiniz.",
		ukUA: "Ви вирішили не надсилати квитанцію.",
		uz:   "Siz kvitansiyani yubormaslikka qaror qildingiz.",
		zhCN: "您已决定不发送收据。"},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		arEG: "لقد غيرت رأيي",
		de:   "Ich habe meine Meinung geändert",
		en:   "I've changed my mind",
		es:   "He cambiado de opinion",
		faIR: "نظرم را عوض کردم.",
		fr:   "J&#39;ai changé d&#39;avis",
		id:   "Aku sudah berubah pikiran",
		it:   "Ho cambiato idea",
		jaJP: "気が変わった",
		koKR: "나는 마음을 바꿨다",
		pl:   "Zmieniłem zdanie",
		ptBR: "Eu mudei de ideia",
		ptPT: "Eu mudei de ideias",
		ru:   "Я передумал(а)",
		tr:   "fikrimi değiştirdim",
		ukUA: "Я передумав",
		uz:   "Men fikrimni o&#39;zgartirdim",
		zhCN: "我改变主意了"},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		arEG: "أرسل عبر تيليجرام",
		de:   "Senden per Telegram",
		en:   "Send by Telegram",
		es:   "Enviar a través de Telegram",
		faIR: "با تلگرام ارسال شود",
		fr:   "Envoyer par Telegram",
		id:   "Kirim melalui Telegram",
		it:   "Invia tramite Telegram",
		jaJP: "Telegramで送信",
		koKR: "텔레그램으로 보내기",
		pl:   "Wyślij przez Telegram",
		ptBR: "Enviar por Telegram",
		ptPT: "Enviar por Telegram",
		ru:   "Отправить через Telelgram",
		tr:   "Telegram ile gönder",
		ukUA: "Надіслати через Telegram",
		uz:   "Telegram orqali yuboring",
		zhCN: "通过电报发送"},
	COMMAND_TEXT_GET_LINK_FOR_RECEIPT_IN_TELEGRAM: {
		arEG: "احصل على رابط الإيصال في Telegram",
		de:   "Erhalten sie einen link für eine quittung in Telegram",
		en:   "Get link for a receipt in Telegram",
		es:   "Obtener enlace para recibirlo en Telegram",
		faIR: "دریافت پیوند برای دریافت در Telegram",
		fr:   "Obtenir un lien pour un reçu dans Telegram",
		id:   "Dapatkan tautan untuk tanda terima di Telegram",
		it:   "Link per la ricevuta nel Telegram",
		jaJP: "Telegramで領収書のリンクを取得する",
		koKR: "Telegram에서 영수증 링크 받기",
		pl:   "Uzyskaj link do rachunku w Telegramie",
		ptBR: "Obter link para um recibo no Telegram",
		ptPT: "Obter link para um recibo no Telegram",
		ru:   "Ссылка для квитанции в Телеграмме",
		tr:   "Telegram&#39;da makbuz için bağlantı alın",
		ukUA: "Отримати посилання для отримання квитанції в Telegram",
		uz:   "Telegramda kvitansiya uchun havolani oling",
		zhCN: "获取 Telegram 收据链接"},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		arEG: "أرسل عبر الفيسبوك، واتساب، فايبر، الخ.",
		de:   "Senden per FB, WhatsApp, Viber, etc.",
		en:   "Send by FB, WhatsApp, Viber, etc.",
		es:   "Enviar a través de FB, WhatsApp, Viber, etc.",
		faIR: "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
		fr:   "Envoyer par FB, WhatsApp, Viber, etc.",
		id:   "Kirim melalui FB, WhatsApp, Viber, dll.",
		it:   "Invia con FB, WhatsCrap, Viber, etc.",
		jaJP: "FB、WhatsApp、Viberなどで送信します。",
		koKR: "FB, WhatsApp, Viber 등으로 보내세요.",
		pl:   "Wyślij przez FB, WhatsApp, Viber, itp.",
		ptBR: "Enviar pelo FB, WhatsApp, Viber, etc.",
		ptPT: "Enviar por FB, WhatsApp, Viber, etc.",
		ru:   "Отправить через Viber, VK, FB, ...",
		tr:   "FB, WhatsApp, Viber vb. ile gönderebilirsiniz.",
		ukUA: "Надсилайте через FB, WhatsApp, Viber тощо.",
		uz:   "FB, WhatsApp, Viber va boshqalar orqali yuboring.",
		zhCN: "通过 FB、WhatsApp、Viber 等发送。"},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		arEG: "إرسال عبر الرسائل القصيرة",
		de:   "Senden per SMS",
		en:   "Send by SMS",
		es:   "Enviar a través de SMS",
		faIR: "با پیام کوتاه ارسال شود",
		fr:   "Envoyer par SMS",
		id:   "Kirim melalui SMS",
		it:   "Invia tramite SMS",
		jaJP: "SMSで送信",
		koKR: "SMS로 보내기",
		pl:   "Wyślij SMS-em",
		ptBR: "Enviar por SMS",
		ptPT: "Enviar por SMS",
		ru:   "Отправить через SMS",
		tr:   "SMS ile gönder",
		ukUA: "Надіслати через SMS",
		uz:   "SMS orqali yuboring",
		zhCN: "通过短信发送"},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		arEG: "أرسل عبر VK.com",
		de:   "Senden per VK.com",
		en:   "Send throw VK.com",
		es:   "Enviar vía VK.com",
		faIR: "ارسال شود VK.com از طریق ",
		fr:   "Envoyer un message sur VK.com",
		id:   "Kirim lewat VK.com",
		it:   "Invia tramite VK.com (Facebook russo)",
		jaJP: "VK.comに送信",
		koKR: "VK.com으로 보내주세요",
		pl:   "Wyślij rzut VK.com",
		ptBR: "Enviar através do VK.com",
		ptPT: "Enviar através do VK.com",
		ru:   "Отправить через ВКонтакте",
		tr:   "Gönder at VK.com",
		ukUA: "Надішліть через VK.com",
		uz:   "throw VK.com ni yuboring",
		zhCN: "发送投掷 VK.com"},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		arEG: "إرسال رمي موافق",
		de:   "Senden per OK",
		en:   "Send throw OK",
		es:   "Enviar a través de OK",
		faIR: "ارسال شود OK از طریق ",
		fr:   "Envoyer lancer OK",
		id:   "Kirim lemparan OK",
		it:   "Invia tramite OK",
		jaJP: "送信OK",
		koKR: "보내기 OK",
		pl:   "Wyślij rzut OK",
		ptBR: "Enviar lance OK",
		ptPT: "Enviar lance OK",
		ru:   "Отправить через Одноклассники",
		tr:   "Gönder at OK",
		ukUA: "Надіслати кидок ОК",
		uz:   "OK ni yuboring",
		zhCN: "发送投掷OK"},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		arEG: "أرسل عبر الفيسبوك",
		de:   "Senden per Facebook",
		en:   "Send throw Facebook",
		es:   "Enviar a través de Facebook",
		faIR: "از طریق فیسبوک ارسال شود.",
		fr:   "Envoyer un message sur Facebook",
		id:   "Kirim lewat Facebook",
		it:   "Invia tramite Facebook",
		jaJP: "Facebookで送信",
		koKR: "페이스북에 보내기",
		pl:   "Wyślij przez Facebook",
		ptBR: "Enviar pelo Facebook",
		ptPT: "Enviar pelo Facebook",
		ru:   "Отправить через Facebook",
		tr:   "Facebook&#39;a gönder",
		ukUA: "Надіслати через Facebook",
		uz:   "Facebook-ni yuboring",
		zhCN: "发送至 Facebook"},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		arEG: "إرسال رمي تويتر",
		de:   "Senden per Twitter",
		en:   "Send throw Twitter",
		es:   "Enviar a través de Twitter",
		faIR: "از طریق توئیتر ارسال شود.",
		fr:   "Envoyer un message sur Twitter",
		id:   "Kirim lewat Twitter",
		it:   "Invia tramite Twitter",
		jaJP: "Twitterで送信",
		koKR: "트위터에 던지기 보내기",
		pl:   "Wyślij rzutem Twittera",
		ptBR: "Enviar pelo Twitter",
		ptPT: "Enviar pelo Twitter",
		ru:   "Отправить через Twitter",
		tr:   "Twitter&#39;a gönder",
		ukUA: "Надіслати через Твіттер",
		uz:   "Twitter ni yuboring",
		zhCN: "发送至 Twitter"},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		arEG: "إلغاء إرسال الإيصال عبر Telegram",
		de:   "Sendung der Quittung per Telegram abbrechen",
		en:   "Cancel sending receipt by Telegram",
		es:   "Cancelar el envío a través de Telegram",
		faIR: "ارسال رسید با تلگرام کنسل شود",
		fr:   "Annuler l&#39;envoi du reçu par Telegram",
		id:   "Batalkan pengiriman tanda terima melalui Telegram",
		it:   "Annulla l'invio tramite Telegram",
		jaJP: "Telegramで領収書の送信をキャンセルする",
		koKR: "텔레그램으로 영수증 발송 취소",
		pl:   "Anuluj wysłanie potwierdzenia przez Telegram",
		ptBR: "Cancelar envio de recibo pelo Telegram",
		ptPT: "Cancelar envio de recibo pelo Telegram",
		ru:   "Отменить отправку через Telegram",
		tr:   "Telegram ile makbuz gönderimini iptal et",
		ukUA: "Скасувати надсилання квитанції через Telegram",
		uz:   "Telegram orqali kvitansiyani yuborishni bekor qiling",
		zhCN: "取消通过 Telegram 发送收据"},
	MAIN_MENU: {
		arEG: "القائمة الرئيسية",
		de:   "Hauptmenü",
		en:   "Main menu",
		es:   "Menú principal",
		faIR: "منوی اصلی",
		fr:   "Menu principal",
		id:   "Menu utama",
		it:   "Menu principale",
		jaJP: "メインメニュー",
		koKR: "메인 메뉴",
		pl:   "Menu główne",
		ptBR: "Menu principal",
		ptPT: "Menu principal",
		ru:   "Главное меню",
		tr:   "Ana menü",
		ukUA: "Головне меню",
		uz:   "Asosiy menyu",
		zhCN: "主菜单",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		arEG: "القائمة الرئيسية",
		de:   "Menü /menu",
		en:   "Main /menu",
		es:   "Inicio /menú",
		faIR: "/منو ی اصلی ",
		fr:   "Menu principal",
		id:   "Utama /menu",
		it:   "Menu' /menu",
		jaJP: "メイン / メニュー",
		koKR: "메인/메뉴",
		pl:   "Główna /menu",
		ptBR: "Menu principal",
		ptPT: "Menu principal",
		ru:   "Главное /меню",
		tr:   "Ana /menü",
		ukUA: "Головне /меню",
		uz:   "Asosiy / menyu",
		zhCN: "主菜单"},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		arEG: "لا يوجد شيء للإلغاء.",
		de:   "Nichts zum abbrechen.",
		en:   "Nothing to cancel.",
		es:   "No hay nada que anular.",
		faIR: "چیزی برای کنسل شدن وجود ندارد",
		fr:   "Rien à annuler.",
		id:   "Tidak ada yang perlu dibatalkan.",
		it:   "Nulla da annullare.",
		jaJP: "キャンセルするものはありません。",
		koKR: "취소할 사항이 없습니다.",
		pl:   "Nie ma nic do anulowania.",
		ptBR: "Nada a cancelar.",
		ptPT: "Nada a cancelar.",
		ru:   "Нечего отменять.",
		tr:   "İptal edilecek bir şey yok.",
		ukUA: "Нічого скасовувати.",
		uz:   "Bekor qilish uchun hech narsa yo&#39;q.",
		zhCN: "没什么可取消的。"},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		arEG: "تم إلغاء إنشاء سجل الديون.",
		de:   "Erstellung des Schuldscheins abgebrochen.",
		en:   "Creation of debt record has been canceled.",
		es:   "La creación del recordatorio se ha cancelado.",
		faIR: "ایجاد سابقه بدهی کنسل شد.",
		fr:   "La création d&#39;un dossier de dette a été annulée.",
		id:   "Pembuatan catatan utang telah dibatalkan.",
		it:   "Creazione record annullata",
		jaJP: "債務記録の作成はキャンセルされました。",
		koKR: "부채 기록 생성이 취소되었습니다.",
		pl:   "Anulowano tworzenie rekordu długu.",
		ptBR: "A criação do registro de dívida foi cancelada.",
		ptPT: "A criação do registo de dívida foi cancelada.",
		ru:   "Создание записи о долге отменено.",
		tr:   "Borç kaydının oluşturulması iptal edildi.",
		ukUA: "Створення обліку боргу скасовано.",
		uz:   "Qarz rekordini yaratish bekor qilindi.",
		zhCN: "债务记录的创建已被取消。"},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		arEG: "إظهار الكل...",
		de:   "Zeige alle...",
		en:   "Show all...",
		es:   "Mostrar todo...",
		faIR: "نمایش تمام موارد ...",
		fr:   "Tout afficher...",
		id:   "Tampilkan semua...",
		it:   "Mostra tutto...",
		jaJP: "すべて表示...",
		koKR: "모두 보기...",
		pl:   "Pokaż wszystko...",
		ptBR: "Mostrar tudo...",
		ptPT: "Mostrar tudo...",
		ru:   "Показать все...",
		tr:   "Tümünü göster...",
		ukUA: "Показати всі...",
		uz:   "Hammasini ko&#39;rsatish...",
		zhCN: "显示全部..."},
	COMMAND_TEXT_CONTACTS: {
		arEG: "جهات الاتصال",
		de:   "Kontakte",
		en:   "Contacts",
		es:   "Contactos",
		faIR: "لیست تماس",
		fr:   "Contacts",
		id:   "Kontak",
		it:   "Сontatti",
		jaJP: "連絡先",
		koKR: "콘택트 렌즈",
		pl:   "Łączność",
		ptBR: "Contatos",
		ptPT: "Contatos",
		ru:   "Контакты",
		tr:   "İletişim",
		ukUA: "Контакти",
		uz:   "Kontaktlar",
		zhCN: "联系方式"},
	COMMAND_TEXT_REFRESH: {
		arEG: "ينعش",
		de:   "Aktualisieren",
		en:   "Refresh",
		es:   "Recargar",
		faIR: "تازه کردن",
		fr:   "Rafraîchir",
		id:   "Menyegarkan",
		it:   "Ricaricare",
		jaJP: "リフレッシュ",
		koKR: "새로 고치다",
		pl:   "Odświeżać",
		ptBR: "Atualizar",
		ptPT: "Atualizar",
		ru:   "Обновить",
		tr:   "Yenile",
		ukUA: "Оновити",
		uz:   "Yangilash",
		zhCN: "刷新"},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		arEG: "شيء آخر",
		de:   "Etwas anderes",
		en:   "Something else",
		es:   "Otra cosa",
		faIR: "چیزی دیگر",
		fr:   "Autre chose",
		id:   "Sesuatu yang lain",
		it:   "Qualcos'altro",
		jaJP: "その他",
		koKR: "다른 것",
		pl:   "Coś innego",
		ptBR: "Outra coisa",
		ptPT: "Outra coisa",
		ru:   "Что-то другое",
		tr:   "Başka bir şey",
		ukUA: "Щось інше",
		uz:   "Boshqa narsa",
		zhCN: "其他"},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		arEG: "هل تم إرجاع هذا الدين؟",
		de:   "Wurde diese Schuld beglichen?",
		en:   "Have this debt been returned?",
		es:   "¿Se ha devuelto esta deuda?",
		faIR: "آیا این بدهی بازپرداخت شده است؟",
		fr:   "Cette dette a-t-elle été remboursée ?",
		id:   "Apakah hutang ini sudah dikembalikan?",
		it:   "Questo debito e' stato saldato?",
		jaJP: "この借金は返済されましたか？",
		koKR: "이 빚은 갚아졌나요?",
		pl:   "Czy ten dług został zwrócony?",
		ptBR: "Essa dívida foi devolvida?",
		ptPT: "Essa dívida foi devolvida?",
		ru:   "Был ли этот долг возвращён?",
		tr:   "Bu borç geri ödendi mi?",
		ukUA: "Чи повернули цей борг?",
		uz:   "Bu qarz qaytarildimi?",
		zhCN: "这笔债还清了吗？"},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		arEG: "متى يجب أن نذكرك بهذا الدين مرة أخرى؟",
		de:   "Wann willst du wieder an diese Schuld erinnert werden?",
		en:   "When should we remind you about this debt again?",
		es:   "¿Cuándo recordarte de esta deuda otra vez?",
		faIR: "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
		fr:   "Quand devrions-nous vous rappeler à nouveau cette dette ?",
		id:   "Kapan kami harus mengingatkan Anda tentang utang ini lagi?",
		it:   "Quando devo ricordarti di questo debito?",
		jaJP: "この借金について、いつ再度お知らせすればいいでしょうか?",
		koKR: "이 부채에 대해 언제 다시 알려드려야 하나요?",
		pl:   "Kiedy mamy Ci ponownie przypomnieć o tym długu?",
		ptBR: "Quando devemos lembrá-lo novamente sobre essa dívida?",
		ptPT: "Quando devemos voltar a lembrá-lo sobre essa dívida?",
		ru:   "Когда вам напомнить об этом долге ещё раз?",
		tr:   "Bu borcu bir daha ne zaman hatırlatalım size?",
		ukUA: "Коли нам знову нагадати вам про цей борг?",
		uz:   "Bu qarz haqida yana qachon eslatishimiz kerak?",
		zhCN: "我们什么时候应该再次提醒您这笔债务？"},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		arEG: "لقد رددت أن الدين قد تم سداده بالكامل.",
		de:   "Du hast angegeben, dass diese Schuld vollständig beglichen ist.",
		en:   "You've replied back that debt has been returned fully.",
		es:   "Has confirmado que la deuda se ha saldado totalmente",
		faIR: "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
		fr:   "Vous avez répondu que la dette avait été entièrement remboursée.",
		id:   "Anda membalas bahwa hutang telah dikembalikan sepenuhnya.",
		it:   "Hai confermato che il debito e' stato saldato.",
		jaJP: "借金は全額返済したと返信しました。",
		koKR: "당신은 빚이 전액 상환되었다고 답장했습니다.",
		pl:   "Odpowiedziałeś, że dług został zwrócony w całości.",
		ptBR: "Você respondeu que a dívida foi totalmente quitada.",
		ptPT: "Respondeu que a dívida foi totalmente liquidada.",
		ru:   "Вы ответили что долг возвращён полностью.",
		tr:   "Borcun tamamının iade edildiğini söylediniz.",
		ukUA: "Ви відповіли, що борг повністю повернуто.",
		uz:   "Siz qarz to&#39;liq qaytarilganiga javob berdingiz.",
		zhCN: "您回复说债务已全部偿还。"},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		arEG: "لقد تم سداد الدين بالكامل.",
		de:   "Die Schuld ist vollständig beglichen.",
		en:   "The debt has been returned fully.",
		es:   "La deuda se ha saldado totalmente",
		faIR: "بدهی به صورت کامل بازپرداخت شده است",
		fr:   "La dette a été entièrement remboursée.",
		id:   "Utangnya telah dikembalikan sepenuhnya.",
		it:   "Il debito e' stato saldato.",
		jaJP: "借金は全額返済されました。",
		koKR: "빚은 전액 갚았습니다.",
		pl:   "Dług został zwrócony w całości.",
		ptBR: "A dívida foi totalmente paga.",
		ptPT: "A dívida foi integralmente paga.",
		ru:   "Долг возвращён полностью.",
		tr:   "Borç tamamen iade edildi.",
		ukUA: "Борг повернуто повністю.",
		uz:   "Qarz to&#39;liq qaytarilgan.",
		zhCN: "债务已全部偿还。"},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		arEG: "التفاصيل هنا: %v",
		de:   "Details hier: %v",
		en:   "Details here: %v",
		es:   "Detalles aquí: %v",
		faIR: "جزئیات در اینجا: %v",
		fr:   "Détails ici : %v",
		id:   "Detailnya di sini: %v",
		it:   "Dettagli qui: %v",
		jaJP: "詳細はこちら: %v",
		koKR: "자세한 내용은 여기를 참조하세요: %v",
		pl:   "Szczegóły tutaj: %v",
		ptBR: "Detalhes aqui: %v",
		ptPT: "Detalhes aqui: %v",
		ru:   "Подробности тут: %v",
		tr:   "Ayrıntılar burada: %v",
		ukUA: "Деталі тут: %v",
		uz:   "Tafsilotlar bu yerda: %v",
		zhCN: "详情请见：%v"},
	MESSAGE_TEXT_REMINDER: {
		arEG: "تذكير",
		de:   "Erinnerung",
		en:   "Reminder",
		es:   "Recordatorio",
		faIR: "یادآور",
		fr:   "Rappel",
		id:   "Pengingat",
		it:   "Promemoria",
		jaJP: "リマインダー",
		koKR: "알림",
		pl:   "Przypomnienie",
		ptBR: "Lembrete",
		ptPT: "Lembrete",
		ru:   "Напоминание",
		tr:   "Hatırlatma",
		ukUA: "Нагадування",
		uz:   "Eslatma",
		zhCN: "提醒"},
	MESSAGE_TEXT_REMINDER_SET: {
		arEG: "تم تعيين التذكير لـ: %v",
		de:   "Erinnerung am: %v",
		en:   "Reminder set for: %v",
		es:   "Recordatorio establecito para: %v",
		faIR: "یادآور تنظیم شده است برای: %v",
		fr:   "Rappel défini pour : %v",
		id:   "Pengingat ditetapkan untuk: %v",
		it:   "Imposta promemoria per: %v",
		jaJP: "リマインダーを設定しました: %v",
		koKR: "%v에 대한 알림 설정됨",
		pl:   "Przypomnienie ustawione dla: %v",
		ptBR: "Lembrete definido para: %v",
		ptPT: "Lembrete definido para: %v",
		ru:   "Напоминание установлено на: %v",
		tr:   "Hatırlatma ayarlandı: %v",
		ukUA: "Нагадування встановлено для: %v",
		uz:   "Eslatma: %v",
		zhCN: "已设置提醒：%v"},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		arEG: "لقد قمت بتعطيل التذكيرات لهذا الدين.",
		de:   "Du hast die Erinnerung an diese Schuld deaktiviert.",
		en:   "You've disabled reminders for this debt.",
		es:   "Recordatorio para esta deuda se ha deshabilitado.",
		faIR: "شما یادآور را برای این بدهی غیرفعال نموده اید.",
		fr:   "Vous avez désactivé les rappels pour cette dette.",
		id:   "Anda telah menonaktifkan pengingat untuk utang ini.",
		it:   "Hai disabilitato il promemoria per questo debito.",
		jaJP: "この借金のリマインダーを無効にしました。",
		koKR: "이 부채에 대한 알림을 비활성화했습니다.",
		pl:   "Wyłączyłeś przypomnienia o tym długu.",
		ptBR: "Você desativou os lembretes para esta dívida.",
		ptPT: "Desativou os lembretes para esta dívida.",
		ru:   "Напоминания об этом долге отключены.",
		tr:   "Bu borç için hatırlatıcıları devre dışı bıraktınız.",
		ukUA: "Ви вимкнули нагадування для цього боргу.",
		uz:   "Siz bu qarz uchun eslatmalarni o‘chirib qo‘ydingiz.",
		zhCN: "您已禁用此债务的提醒功能。"},
	COMMAND_TEXT_REMINDER_ENABLE: {
		arEG: "تذكير بالتشغيل",
		de:   "Erinnerung aktivieren",
		en:   "Turn-on reminder",
		es:   "Recordatorio de encendido",
		faIR: "یادآوری روشن",
		fr:   "Rappel d&#39;activation",
		id:   "Pengingat untuk menyalakan",
		it:   "Ricordo promozionale",
		jaJP: "電源オンリマインダー",
		koKR: "켜기 알림",
		pl:   "Przypomnienie o włączeniu",
		ptBR: "Lembrete de ativação",
		ptPT: "Lembrete de ativação",
		ru:   "Включить напоминание",
		tr:   "Açma hatırlatıcısı",
		ukUA: "Нагадування про ввімкнення",
		uz:   "Yoqish eslatmasi",
		zhCN: "开机提醒"},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		arEG: "لقد قمت بالفعل بإعادة جدولة هذا التذكير.",
		de:   "Du wirst bereits erneut erinnert.",
		en:   "You've already rescheduled this reminder.",
		es:   "Recordatorio para esta deuda se ha reprogramado ya.",
		faIR: "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
		fr:   "Vous avez déjà reprogrammé ce rappel.",
		id:   "Anda telah menjadwalkan ulang pengingat ini.",
		it:   "Hai gia' impostato questo promemoria",
		jaJP: "このリマインダーはすでに再スケジュールされています。",
		koKR: "이미 이 알림을 다시 예약했습니다.",
		pl:   "Już przełożyłeś to przypomnienie.",
		ptBR: "Você já reagendou este lembrete.",
		ptPT: "Já reagendou este lembrete.",
		ru:   "Напоминание об этом долге уже перенесено.",
		tr:   "Bu hatırlatıcıyı zaten yeniden planladınız.",
		ukUA: "Ви вже перенесли це нагадування.",
		uz:   "Siz allaqachon bu eslatmani boshqa vaqtga belgilagansiz.",
		zhCN: "您已经重新安排了此提醒。"},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		arEG: "نعم، العودة بالكامل",
		de:   "Ja, vollständig beglichen",
		en:   "Yes, returne in full",
		es:   "Sí, devuelto totalmente",
		faIR: "بله، بازپرداخت به صورت کامل",
		fr:   "Oui, je reviens en entier",
		id:   "Ya, kembalikan secara penuh",
		it:   "Si, completamento saldato",
		jaJP: "はい、全額返金します",
		koKR: "네, 전액 환불해 드립니다",
		pl:   "Tak, zwrócę w całości",
		ptBR: "Sim, retorno integral",
		ptPT: "Sim, retorno integral",
		ru:   "Да, возвращено полностью",
		tr:   "Evet, tam olarak geri dön",
		ukUA: "Так, повернути повністю",
		uz:   "Ha, to&#39;liq qaytib keling",
		zhCN: "是的，全额退还"},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		arEG: "عاد جزئيا",
		de:   "Teilweise beglichen",
		en:   "Returned partially",
		es:   "Devuelto parcialmente",
		faIR: "تا اندازه ای بازپرداخت شده است",
		fr:   "Revenu partiellement",
		id:   "Dikembalikan sebagian",
		it:   "Parzialmente saldato",
		jaJP: "部分的に返却",
		koKR: "부분적으로 반환됨",
		pl:   "Zwrócono częściowo",
		ptBR: "Devolvido parcialmente",
		ptPT: "Retornou parcialmente",
		ru:   "Возврашено частично",
		tr:   "Kısmen geri döndü",
		ukUA: "Повернуто частково",
		uz:   "Qisman qaytarildi",
		zhCN: "部分退回"},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		arEG: "لم يتم إرجاعه",
		de:   "Nicht beglichen",
		en:   "Not returned",
		es:   "No devuelto",
		faIR: "بازپرداخت نشده است",
		fr:   "Non retourné",
		id:   "Tidak dikembalikan",
		it:   "Debito non saldato",
		jaJP: "返却されません",
		koKR: "반환되지 않음",
		pl:   "Nie zwrócono",
		ptBR: "Não devolvido",
		ptPT: "Não devolvido",
		ru:   "Не возвращено",
		tr:   "Geri dönmedi",
		ukUA: "Не повернуто",
		uz:   "Qaytarilmadi",
		zhCN: "未归还"},
	MESSAGE_TEXT_YOU_REPLIED: {
		arEG: "لقد أجبت: %v",
		de:   "Beantwortet: %v",
		en:   "You've replied: %v",
		es:   "Has respondido: %v",
		faIR: "شما پاسخ داده اید: %v",
		fr:   "Vous avez répondu : %v",
		id:   "Anda telah membalas: %v",
		it:   "Hai risposto: %v",
		jaJP: "返信しました: %v",
		koKR: "당신은 %v에게 답변했습니다.",
		pl:   "Odpowiedziałeś: %v",
		ptBR: "Você respondeu: %v",
		ptPT: "Respondeu: %v",
		ru:   "Вы ответили: %v",
		tr:   "Cevapladınız: %v",
		ukUA: "Ви відповіли: %v",
		uz:   "Siz javob berdingiz: %v",
		zhCN: "您已回复：%v"},
	"book": {
		arEG: "كتاب",
		de:   "buchen",
		en:   "book",
		es:   "libro",
		faIR: "کتاب",
		fr:   "livre",
		id:   "buku",
		it:   "libro",
		jaJP: "本",
		koKR: "책",
		pl:   "książka",
		ptBR: "livro",
		ptPT: "livro",
		ru:   "книгу",
		tr:   "kitap",
		ukUA: "книга",
		uz:   "kitob",
		zhCN: "书"},
	"MessageTextBotDidNotUnderstandTheCommand": {
		arEG: "عذرًا، لم أفهم طلبك. ربما أكون غبيًا بعض الشيء... يمكنك العودة إلى القائمة الرئيسية.",
		de:   "\xF0\x9F\x98\xB3 Entschuldigung, aber ich habe deinen Befehl nicht verstanden. Vielleicht bin ich ein bisschen dumm...\n\nDu kannst zurück ins /menu",
		en:   "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		es:   "\xF0\x9F\x98\xB3 Disculpa, no he entendido tu orden. Tal vez soy un poco tonto...\n\nPuedes volver al Menu principal /menu",
		faIR: "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		fr:   "Désolé, je n&#39;ai pas compris votre commande. Je suis peut-être un peu bête... Vous pouvez revenir au menu principal.",
		id:   "\xF0\x9F\x98\xB3 Maaf, saya tidak mengerti pesanan Anda. Mungkin saya agak bodoh...\n\nAnda dapat kembali ke menu utama",
		it:   "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
		jaJP: "\xF0\x9F\x98\xB3 申し訳ありませんが、ご注文の内容が理解できませんでした。少し頭が悪かったのかもしれません…\n\nメインメニューに戻ってください",
		koKR: "\xF0\x9F\x98\xB3 죄송합니다. 주문하신 내용을 이해하지 못했습니다. 제가 좀 멍청한 걸까요...\n\n메인 메뉴로 돌아가실 수 있습니다.",
		pl:   "\xF0\x9F\x98\xB3 Przepraszam, nie zrozumiałem Twojego zamówienia. Może jestem trochę głupi...\n\nMożesz wrócić do głównego /menu",
		ptBR: "\xF0\x9F\x98\xB3 Desculpe, não entendi seu pedido. Talvez eu seja um pouco burro...\n\nVocê pode retornar ao menu principal",
		ptPT: "\xF0\x9F\x98\xB3 Desculpe, não percebi o seu pedido. Talvez eu seja um pouco burro...\n\nPode voltar ao menu principal",
		ru:   "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
		tr:   "\xF0\x9F\x98\xB3 Üzgünüm, siparişinizi anlamadım. Belki biraz aptalım...\n\nAna menüye dönebilirsiniz",
		ukUA: "\xF0\x9F\x98\xB3 Вибачте, я не зрозумів вашого замовлення. Можливо, я трохи незрозумілий...\n\nВи можете повернутися до головного меню",
		uz:   "\xF0\x9F\x98\xB3 Kechirasiz, buyurtmangizni tushunmadim. Balki men biroz ahmoqman...\n\nAsosiy /menyuga qaytishingiz mumkin",
		zhCN: "\xF0\x9F\x98\xB3 抱歉，我没理解您的命令。可能是我有点笨……\n\n您可以返回主菜单"},
	COMMAND_TEXT_LANGUAGE: {
		arEG: "لغة الروبوت",
		de:   "Sprache / Language",
		en:   "Bot language",
		es:   "Idioma / Language",
		faIR: "زبان",
		fr:   "langage du bot",
		id:   "Bahasa bot",
		it:   "Lingua / Language",
		jaJP: "ボット言語",
		koKR: "봇 언어",
		pl:   "Język bota",
		ptBR: "Linguagem de bot",
		ptPT: "Linguagem de bot",
		ru:   "Язык / Language",
		tr:   "Bot dili",
		ukUA: "Мова бота",
		uz:   "Bot tili",
		zhCN: "机器人语言"},
	"/start": {
		arEG: "/يبدأ",
		de:   "/Start",
		en:   "/start",
		es:   "/comienzo",
		faIR: "/شروع",
		fr:   "/commencer",
		id:   "/awal",
		it:   "/inizio",
		jaJP: "/始める",
		koKR: "/시작",
		pl:   "/start",
		ptBR: "/começar",
		ptPT: "/iniciar",
		ru:   "/старт",
		tr:   "/başlangıç",
		ukUA: "/початок",
		uz:   "/boshlash",
		zhCN: "/开始"},
	COMMAND_TEXT_DUE_RETURNS: {
		arEG: "الإقرارات المستحقة",
		de:   "Fällige Schulden",
		en:   "Due returns",
		es:   "Devoluciones",
		faIR: "بازپرداخت بدهی",
		fr:   "Retours dus",
		id:   "Pengembalian jatuh tempo",
		it:   "Debiti",
		jaJP: "期限切れの返品",
		koKR: "반환 기한",
		pl:   "Terminowe zwroty",
		ptBR: "Devoluções devidas",
		ptPT: "Devoluções devidas",
		ru:   "Предстоящие платежи",
		tr:   "Vadesi gelen iadeler",
		ukUA: "Належні повернення",
		uz:   "Kerakli qaytishlar",
		zhCN: "应得回报"},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		arEG: "<b>الديون المتأخرة:</b>",
		de:   "<b>Überfällige Schulden:</b>",
		en:   "<b>Overdue debts:</b>",
		es:   "<b>Deudas atrasadas:</b>",
		faIR: "<b>بدهی های معوق:</b>",
		fr:   "<b>Dettes en souffrance :</b>",
		id:   "<b>Hutang yang telah jatuh tempo:</b>",
		it:   "<b>Debiti in ritardo:</b>",
		jaJP: "<b>延滞債務:</b>",
		koKR: "<b>연체된 부채:</b>",
		pl:   "<b>Przeterminowane długi:</b>",
		ptBR: "<b>Dívidas vencidas:</b>",
		ptPT: "<b>Dívidas vencidas:</b>",
		ru:   "<b>Просроченные долги:</b>",
		tr:   "<b>Vadesi geçmiş borçlar:</b>",
		ukUA: "<b>Прострочені борги:</b>",
		uz:   "<b>Muddati o&#39;tgan qarzlar:</b>",
		zhCN: "<b>逾期债务：</b>"},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		arEG: "<b>أقرب الديون للعودة:</b>",
		de:   "<b>Bald fällige Schulden:</b>",
		en:   "<b>Closest debts to return:</b>",
		es:   "<b>Deudas más cercanos que pagar:</b>",
		faIR: "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
		fr:   "<b>Dettes les plus proches à rembourser :</b>",
		id:   "<b>Utang terdekat yang harus dikembalikan:</b>",
		it:   "<b>Debiti in scadenza:</b>",
		jaJP: "<b>返済に最も近い債務:</b>",
		koKR: "<b>반환해야 할 가장 가까운 부채:</b>",
		pl:   "<b>Najbliższe długi do zwrotu:</b>",
		ptBR: "<b>Dívidas mais próximas de serem devolvidas:</b>",
		ptPT: "<b>Dívidas mais próximas de serem devolvidas:</b>",
		ru:   "<b>Ближайшие долги к возрату:</b>",
		tr:   "<b>İade edilmesi gereken en yakın borçlar:</b>",
		ukUA: "<b>Найближчі до повернення борги:</b>",
		uz:   "<b>Qaytarilishi kerak bo&#39;lgan eng yaqin qarzlar:</b>",
		zhCN: "<b>最接近偿还的债务：</b>"},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		arEG: "%v يتوقع %v منك في %v",
		de:   "%v bekommt %v von dir, spätestens in %v",
		en:   "%v expects %v from you in %v",
		es:   "%v espera %v que devuelvas en %v",
		faIR: "%v انتظار دارد %v از شما در %v",
		fr:   "%v attend %v de vous dans %v",
		id:   "%v mengharapkan %v dari Anda di %v",
		it:   "%v aspetta %v da te entro il %v",
		jaJP: "%v は %v であなたからの %v を期待しています",
		koKR: "%v는 %v에서 당신에게 %v를 기대합니다.",
		pl:   "%v oczekuje od Ciebie %v w %v",
		ptBR: "%v espera %v de você em %v",
		ptPT: "%v espera %v de si em %v",
		ru:   "%v ожидает от вас возврата %v через %v",
		tr:   "%v sizden %v &#39; yi bekliyor %v",
		ukUA: "%v очікує від вас %v у %v",
		uz:   "%v sizdan %v ni %v kutmoqda",
		zhCN: "%v 期望您在 %v 中提供 %v"},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		arEG: "تتوقع أن يقوم %v بإرجاع %v إليك في %v",
		de:   "%v gibt dir %v, spätestens in %v",
		en:   "You expect %v to return %v to you in %v",
		es:   "Estas esperando de %v que devuelva %v a ti en %v",
		faIR: "شما انتظار دارید %v بازگرداند %v به شما در %v",
		fr:   "Vous vous attendez à ce que %v vous renvoie %v dans %v",
		id:   "Anda mengharapkan %v mengembalikan %v kepada Anda dalam %v",
		it:   "Stai aspettando %v che ti dia %v entro il %v",
		jaJP: "%v が %v で %v を返すことを期待しています",
		koKR: "%v가 %v에서 %v를 반환할 것으로 예상합니다.",
		pl:   "Oczekujesz, że %v zwróci Ci %v w %v",
		ptBR: "Você espera que %v retorne %v para você em %v",
		ptPT: "Espera que %v lhe devolva %v em %v",
		ru:   "Вы ожидаете от %v возврата %v через %v",
		tr:   "%v&#39;nin size %v&#39;yi %v&#39;de döndürmesini beklersiniz",
		ukUA: "Ви очікуєте, що %v поверне вам %v у %v",
		uz:   "%v ichida sizga %v qaytishini kutasiz %v",
		zhCN: "你期望 %v 在 %v 中返回 %v"},
	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		arEG: "ليس لديك ديون مع تاريخ استحقاق محدد.",
		de:   "Du hast keine Schulden mit Fälligkeitsdatum.",
		en:   "You have no debts with set due date.",
		es:   "No tienes deudas con la fecha señalada para devolver. ",
		faIR: "شما بدهی ای با ثبت سررسید ندارید.",
		fr:   "Vous n’avez aucune dette à échéance fixe.",
		id:   "Anda tidak memiliki utang dengan tanggal jatuh tempo yang ditentukan.",
		it:   "Non hai debiti con una data di scadenza.",
		jaJP: "支払期限が設定されている借金はありません。",
		koKR: "만기일이 정해진 부채가 없습니다.",
		pl:   "Nie masz długów z ustalonym terminem płatności.",
		ptBR: "Você não tem dívidas com data de vencimento definida.",
		ptPT: "Não tem dívidas com data de vencimento definida.",
		ru:   "У вас нет долгов с назначеным сроком к возврату.",
		tr:   "Belirli vadesi olan borcunuz yok.",
		ukUA: "У вас немає боргів із встановленою датою погашення.",
		uz:   "To&#39;lov muddati belgilangan qarzlaringiz yo&#39;q.",
		zhCN: "您没有设定到期日的债务。"},
	COMMAND_TEXT_GAVE: {
		arEG: "أعطى",
		de:   "Verleihen",
		en:   "Gave",
		es:   "Prestado por ti",
		faIR: "قرض_دادن",
		fr:   "A donné",
		id:   "Telah memberi",
		it:   "Credito",
		jaJP: "与えた",
		koKR: "주었다",
		pl:   "Dał",
		ptBR: "Deu",
		ptPT: "Deu",
		ru:   "Дать",
		tr:   "Verilmiş",
		ukUA: "Дав",
		uz:   "berdi",
		zhCN: "给了"},
	COMMAND_TEXT_GOT: {
		arEG: "يملك",
		de:   "Anleihen",
		en:   "Got",
		es:   "Prestado a ti",
		faIR: "قرض_گرفتن",
		fr:   "A obtenu",
		id:   "Telah mendapatkan",
		it:   "Debito",
		jaJP: "得た",
		koKR: "갖다",
		pl:   "Dostałem",
		ptBR: "Pegou",
		ptPT: "Ter",
		ru:   "Взять",
		tr:   "Aldım",
		ukUA: "Отримав",
		uz:   "bor",
		zhCN: "得到"},
	COMMAND_TEXT_RETURN: {
		arEG: "يعود",
		de:   "Beglichen",
		en:   "Return",
		es:   "Devuelto",
		faIR: "بازگشت",
		fr:   "Retour",
		id:   "Kembali",
		it:   "Rientra",
		jaJP: "戻る",
		koKR: "반품",
		pl:   "Powrót",
		ptBR: "Retornar",
		ptPT: "Devolver",
		ru:   "Вернуть",
		tr:   "Geri dönmek",
		ukUA: "Повернення",
		uz:   "Qaytish",
		zhCN: "返回"},
	COMMAND_TEXT_BALANCE: {
		arEG: "توازن",
		de:   "Ausstehend",
		en:   "Balance",
		es:   "Balance",
		faIR: "تراز",
		fr:   "Équilibre",
		id:   "Keseimbangan",
		it:   "Bilancio",
		jaJP: "バランス",
		koKR: "균형",
		pl:   "Balansować",
		ptBR: "Equilíbrio",
		ptPT: "Equilíbrio",
		ru:   "Баланс",
		tr:   "Denge",
		ukUA: "Баланс",
		uz:   "Balans",
		zhCN: "平衡"},
	SettingsMessageTitle: {
		arEG: "إعدادات",
		de:   "Einstellungen",
		en:   "Settings",
		es:   "Ajustes",
		faIR: "تنظیمات",
		fr:   "Paramètres",
		id:   "Pengaturan",
		it:   "Settaggi",
		jaJP: "設定",
		koKR: "설정",
		pl:   "Ustawienia",
		ptBR: "Configurações",
		ptPT: "Configurações",
		ru:   "Настройки",
		tr:   "Ayarlar",
		ukUA: "Налаштування",
		uz:   "Sozlamalar",
		zhCN: "设置",
	},
	SettingsButtonText: {
		arEG: "إعدادات",
		de:   "Einstellungen",
		en:   "Settings",
		es:   "Ajustes",
		faIR: "تنظیمات",
		fr:   "Paramètres",
		id:   "Pengaturan",
		it:   "Settaggi",
		jaJP: "設定",
		koKR: "설정",
		pl:   "Ustawienia",
		ptBR: "Configurações",
		ptPT: "Configurações",
		ru:   "Настройки",
		tr:   "Ayarlar",
		ukUA: "Налаштування",
		uz:   "Sozlamalar",
		zhCN: "设置",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		arEG: "خمسة عالية!",
		de:   "Gib mir Fünf!",
		en:   "High five!",
		es:   "¡Choca esos cinco!",
		faIR: "بزن قدش!",
		fr:   "Tape m'en cinq !",
		id:   "Tos!",
		it:   "Batti 5 bro!",
		jaJP: "ハイタッチ！",
		koKR: "하이 파이브!",
		pl:   "Przybij piątkę!",
		ptBR: "Toca aqui!",
		ptPT: "Toca aqui!",
		ru:   "Дать пять!",
		tr:   "Çak bir beşlik!",
		ukUA: "Дай п'ять!",
		uz:   "Besh qo'l!",
		zhCN: "击掌！",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		arEG: "لغة",
		de:   "Sprache",
		en:   "Language",
		es:   "Idioma",
		faIR: "زبان",
		fr:   "Langue",
		id:   "Bahasa",
		it:   "Lingua",
		jaJP: "言語",
		koKR: "언어",
		pl:   "Język",
		ptBR: "Idioma",
		ptPT: "Idioma",
		ru:   "Язык",
		tr:   "Dil",
		ukUA: "Мова",
		uz:   "Til",
		zhCN: "语言",
	},
	COMMAND_TEXT_HELP: {
		arEG: "يساعد",
		de:   "Hilfe",
		en:   "Help",
		es:   "Ayuda",
		faIR: "کمک",
		fr:   "Aide",
		id:   "Bantuan",
		it:   "Aiuto",
		jaJP: "ヘルプ",
		koKR: "도움말",
		pl:   "Pomoc",
		ptBR: "Ajuda",
		ptPT: "Ajuda",
		ru:   "Помощь",
		tr:   "Yardım",
		ukUA: "Допомога",
		uz:   "Yordam",
		zhCN: "帮助",
	},
	COMMAND_TEXT_HISTORY: {
		arEG: "تاريخ",
		de:   "Verlauf",
		en:   "History",
		es:   "Cronología",
		faIR: "پیشینه",
		fr:   "Historique",
		id:   "Riwayat",
		it:   "Cronologia",
		jaJP: "履歴",
		koKR: "기록",
		pl:   "Historia",
		ptBR: "Histórico",
		ptPT: "História",
		ru:   "История",
		tr:   "Geçmiş",
		ukUA: "Історія",
		uz:   "Tarix",
		zhCN: "历史",
	},
	COMMAND_TEXT_CANCEL: {
		arEG: "يلغي",
		de:   "Abbrechen",
		en:   "Cancel",
		es:   "Cancelar",
		faIR: "کنسل",
		fr:   "Annuler",
		id:   "Batal",
		it:   "Annulla",
		jaJP: "キャンセル",
		koKR: "취소",
		pl:   "Anuluj",
		ptBR: "Cancelar",
		ptPT: "Cancelar",
		ru:   "Отменить",
		tr:   "İptal",
		ukUA: "Скасувати",
		uz:   "Bekor qilish",
		zhCN: "取消",
	},
	COMMAND_TEXT_REFERRERS: {
		arEG: "المُحيلون",
		de:   "Empfehlungen",
		en:   "Referrers",
		es:   "Referentes",
		faIR: "معرف\u200cها",
		fr:   "Référents",
		id:   "Referensi",
		it:   "Referenti",
		jaJP: "紹介者",
		koKR: "추천인",
		pl:   "Polecający",
		ptBR: "Referências",
		ptPT: "Referências",
		ru:   "Нас рекомендуют",
		tr:   "Referanslar",
		ukUA: "Нас рекомендують",
		uz:   "Tavsiya qiluvchilar",
		zhCN: "推荐人",
	},
	MESSAGE_TEXT_HOW_TO_ADD_TG_CHANNEL: {
		arEG: `لإضافة قناتك إلى القائمة، اكتب عنا برابط كـ %v <code>&lt;-</code> استبدل <code>YOUR_CHANNEL</code> بقناتك الخاصة.

من الأفضل إخفاء الرابط في HTML كـ:

<pre style=";text-align:right;direction:rtl"> &lt;a href=&quot;%v&quot;&gt;@%v&lt;/a&gt;</pre> يجب أن يتم عرض هذا بواسطة عملاء Telegram على النحو التالي: <a href="%v">@%v</a> 

سيتم عرض أفضل 5 مراجعين لأخر 100 مستخدم جديد هنا.`,
		de: `Um deinen Kanal zur Liste hinzuzufügen, schreibe einfach über uns mit einem Link wie %v <code>&lt;-</code> ersetze <code>YOUR_CHANNEL</code> durch deinen eigenen Kanal.

Es ist besser, wenn du den Link in HTML versteckst als:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Dies sollte von Telegram-Clients so dargestellt werden: <a href="%v">@%v</a>

Die Top 5 Empfehlungen der letzten 100 neuen Benutzer werden hier angezeigt.`,
		en: `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,
		es: `Para añadir tu canal a la lista, simplemente escribe sobre nosotros con un enlace como %v <code>&lt;-</code> reemplaza <code>YOUR_CHANNEL</code> con tu propio canal.

Es mejor si ocultas el enlace en HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Esto debería ser renderizado por los clientes de Telegram como: <a href="%v">@%v</a>

Los 5 principales referentes de los últimos 100 nuevos usuarios se mostrarán aquí.`,
		faIR: `برای اضافه کردن کانال خود به لیست، فقط درباره ما با لینکی مانند %v <code>&lt;-</code> بنویسید و <code>YOUR_CHANNEL</code> را با کانال خود جایگزین کنید.

بهتر است اگر لینک را در HTML به این صورت پنهان کنید:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

این باید توسط کلاینت\u200cهای تلگرام به این صورت نمایش داده شود: <a href="%v">@%v</a>

5 معرف برتر برای 100 کاربر جدید آخر در اینجا نشان داده خواهد شد.`,
		fr: `Pour ajouter votre chaîne à la liste, écrivez simplement à propos de nous avec un lien comme %v <code>&lt;-</code> remplacez <code>YOUR_CHANNEL</code> par votre propre chaîne.

C'est mieux si vous cachez le lien en HTML comme:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Cela devrait être rendu par les clients Telegram comme: <a href="%v">@%v</a>

Les 5 principaux référents pour les 100 derniers nouveaux utilisateurs seront affichés ici.`,
		id: `Untuk menambahkan saluran Anda ke daftar, cukup tulis tentang kami dengan tautan sebagai %v <code>&lt;-</code> ganti <code>YOUR_CHANNEL</code> dengan saluran Anda sendiri.

Lebih baik jika Anda menyembunyikan tautan dalam HTML sebagai:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Ini harus dirender oleh klien Telegram sebagai: <a href="%v">@%v</a>

5 referensi teratas untuk 100 pengguna baru terakhir akan ditampilkan di sini.`,
		it: `Per aggiungere il tuo canale all'elenco, scrivi semplicemente di noi con un link come %v <code>&lt;-</code> sostituisci <code>YOUR_CHANNEL</code> con il tuo canale.

È meglio se nascondi il link in HTML come:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Questo dovrebbe essere visualizzato dai client Telegram come: <a href="%v">@%v</a>

I primi 5 referenti per gli ultimi 100 nuovi utenti saranno mostrati qui.`,
		jaJP: `あなたのチャンネルをリストに追加するには、%v <code>&lt;-</code> のようなリンクで私たちについて書くだけです。<code>YOUR_CHANNEL</code>をあなた自身のチャンネルに置き換えてください。

HTMLでリンクを隠すとより良いでしょう：

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

これはTelegramクライアントによって次のように表示されるはずです： <a href="%v">@%v</a>

最後の100人の新規ユーザーのトップ5紹介者がここに表示されます。`,
		koKR: `채널을 목록에 추가하려면 %v <code>&lt;-</code>와 같은 링크로 우리에 대해 작성하세요. <code>YOUR_CHANNEL</code>을 자신의 채널로 바꾸세요.

HTML에서 링크를 숨기는 것이 좋습니다:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

이것은 Telegram 클라이언트에서 다음과 같이 렌더링되어야 합니다: <a href="%v">@%v</a>

최근 100명의 새 사용자에 대한 상위 5개 추천인이 여기에 표시됩니다.`,
		pl: `Aby dodać swój kanał do listy, po prostu napisz o nas z linkiem jako %v <code>&lt;-</code> zastąp <code>YOUR_CHANNEL</code> swoim własnym kanałem.

Lepiej, jeśli ukryjesz link w HTML jako:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Powinno to być renderowane przez klientów Telegram jako: <a href="%v">@%v</a>

Tutaj zostanie wyświetlonych 5 najlepszych polecających dla ostatnich 100 nowych użytkowników.`,
		ptBR: `Para adicionar seu canal à lista, basta escrever sobre nós com um link como %v <code>&lt;-</code> substitua <code>YOUR_CHANNEL</code> pelo seu próprio canal.

É melhor se você ocultar o link em HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Isso deve ser renderizado pelos clientes do Telegram como: <a href="%v">@%v</a>

Os 5 principais referenciadores para os últimos 100 novos usuários serão mostrados aqui.`,
		ptPT: `Para adicionar o seu canal à lista, basta escrever sobre nós com um link como %v <code>&lt;-</code> substitua <code>YOUR_CHANNEL</code> pelo seu próprio canal.

É melhor ocultar o link em HTML como:

<pre> &lt;a href=&quot;%v&quot;&gt;@%v&lt;/a&gt;</pre> 

Isto deve ser renderizado pelos clientes do Telegram como: <a href="%v">@%v</a> 

Os 5 principais indicadores dos últimos 100 novos utilizadores serão mostrados aqui.`,
		ru: `Чтобы добавить ваш канал в этот список просто напишите об этом боте использую ссылку вида %v <code>&lt;-</code> замените <code>YOUR_CHANNEL</code> на ваш канал.

Будет лучше  если вы спрячете её в HTML как:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Такой код должен отобразиться в Телеграмме как: <a href="%v">@%v</a>

Топ-5 источников последних 100 пользователей будут показаны здесь.`,
		tr: `Kanalınızı listeye eklemek için sadece %v <code>&lt;-</code> gibi bir bağlantı ile hakkımızda yazın, <code>YOUR_CHANNEL</code> yerine kendi kanalınızı yazın.

Bağlantıyı HTML'de şu şekilde gizlerseniz daha iyi olur:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Bu, Telegram istemcileri tarafından şöyle görüntülenmelidir: <a href="%v">@%v</a>

Son 100 yeni kullanıcı için en iyi 5 referans burada gösterilecektir.`,
		ukUA: `Щоб додати свій канал до списку, просто напишіть про нас із посиланням як %v <code>&lt;-</code> замініть <code>YOUR_CHANNEL</code> на свій власний канал.

Краще, якщо ви приховаєте посилання в HTML як:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Це має відображатися клієнтами Telegram як: <a href="%v">@%v</a>

Тут буде показано 5 найкращих рекомендацій для останніх 100 нових користувачів.`,
		uz: `Kanalingizni ro'yxatga qo'shish uchun shunchaki %v <code>&lt;-</code> kabi havola bilan biz haqimizda yozing, <code>YOUR_CHANNEL</code> o'rniga o'z kanalingizni yozing.

Havolani HTML-da quyidagicha yashirsangiz yaxshiroq bo'ladi:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Bu Telegram mijozlari tomonidan quyidagicha ko'rsatilishi kerak: <a href="%v">@%v</a>

Oxirgi 100 ta yangi foydalanuvchi uchun eng yaxshi 5 ta tavsiya qiluvchi bu yerda ko'rsatiladi.`,
		zhCN: `要将您的频道添加到列表中，只需使用链接 %v <code>&lt;-</code> 写下关于我们的信息，将 <code>YOUR_CHANNEL</code> 替换为您自己的频道。

如果您在HTML中隐藏链接会更好：

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

这应该由Telegram客户端呈现为： <a href="%v">@%v</a>

这里将显示最近100个新用户的前5名推荐人。`,
	},
	ButtonTextCancel: {
		arEG: "↩ إلغاء",
		de:   "↩ Zurück",
		en:   "↩ Cancel",
		es:   "↩ Cancelar",
		faIR: "↪ کنسل",
		fr:   "↩ Annuler",
		id:   "↩ Batal",
		it:   "↩ Annulla",
		jaJP: "↩ キャンセル",
		koKR: "↩ 취소",
		pl:   "↩ Anuluj",
		ptBR: "↩ Cancelar",
		ptPT: "↩ Cancelar",
		ru:   "↩ Отменить",
		tr:   "↩ İptal",
		ukUA: "↩ Скасувати",
		uz:   "↩ Bekor qilish",
		zhCN: "↩ 取消",
	},
	BUTTON_TEXT_MAIN_MENU: {
		arEG: "↩ القائمة الرئيسية",
		de:   "↩ Hauptmenü",
		en:   "↩ Main menu",
		es:   "↩ Menú principal",
		faIR: "↪ منوی اصلی",
		fr:   "↩ Menu principal",
		id:   "↩ Menu utama",
		it:   "↩ Menu principale",
		jaJP: "↩ メインメニュー",
		koKR: "↩ 메인 메뉴",
		pl:   "↩ Menu główne",
		ptBR: "↩ Menu principal",
		ptPT: "↩ Menu principal",
		ru:   "↩ Главное меню",
		tr:   "↩ Ana menü",
		ukUA: "↩ Головне меню",
		uz:   "↩ Asosiy menyu",
		zhCN: "↩ 主菜单",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		arEG: "العملة الأساسية",
		de:   "Hauptwährung",
		en:   "Primary currency",
		es:   "Moneda principal",
		faIR: "واحد پول اولیه",
		fr:   "Devise principale",
		id:   "Mata uang utama",
		it:   "Valuta principale",
		jaJP: "主要通貨",
		koKR: "기본 통화",
		pl:   "Główna waluta",
		ptBR: "Moeda principal",
		ptPT: "Moeda primária",
		ru:   "Основная валюта",
		tr:   "Ana para birimi",
		ukUA: "Основна валюта",
		uz:   "Asosiy valyuta",
		zhCN: "主要货币",
	},
	COMMAND_TEXT_ADD_GROUP: {
		arEG: "إضافة مجموعة",
		de:   "Gruppe hinzufügen",
		en:   "Add group",
		es:   "Añadir grupo",
		faIR: "اضافه کردن گروه",
		fr:   "Ajouter un groupe",
		id:   "Tambahkan grup",
		it:   "Aggiungi gruppo",
		jaJP: "グループを追加",
		koKR: "그룹 추가",
		pl:   "Dodaj grupę",
		ptBR: "Adicionar grupo",
		ptPT: "Adicionar grupo",
		ru:   "Добавить группу",
		tr:   "Grup ekle",
		ukUA: "Додати групу",
		uz:   "Guruh qo'shish",
		zhCN: "添加群组",
	},
	COMMAND_TEXT_GROUPS: {
		arEG: "المجموعات",
		de:   "Gruppen",
		en:   "Groups",
		es:   "Grupos",
		faIR: "گروه\u200cها",
		fr:   "Groupes",
		id:   "Grup",
		it:   "Gruppi",
		jaJP: "グループ",
		koKR: "그룹",
		pl:   "Grupy",
		ptBR: "Grupos",
		ptPT: "Grupos",
		ru:   "Группы",
		tr:   "Gruplar",
		ukUA: "Групи",
		uz:   "Guruhlar",
		zhCN: "群组",
	},
	COMMAND_TEXT_BILLS: {
		arEG: "الفواتير",
		de:   "Rechnungen",
		en:   "Bills",
		es:   "Facturas",
		faIR: "صورتحساب\u200cها",
		fr:   "Factures",
		id:   "Tagihan",
		it:   "Fatture",
		jaJP: "請求書",
		koKR: "청구서",
		pl:   "Rachunki",
		ptBR: "Contas",
		ptPT: "Contas",
		ru:   "Счета",
		tr:   "Faturalar",
		ukUA: "Рахунки",
		uz:   "Hisob-fakturalar",
		zhCN: "账单",
	},
	COMMAND_TEXT_SETTLE_BILL: {
		arEG: "تسوية الفاتورة",
		de:   "Rechnung begleichen",
		en:   "Settle bill",
		es:   "Liquidar factura",
		faIR: "تسویه صورتحساب",
		fr:   "Régler la facture",
		id:   "Selesaikan tagihan",
		it:   "Saldare il conto",
		jaJP: "請求書を決済する",
		koKR: "청구서 정산",
		pl:   "Rozlicz rachunek",
		ptBR: "Quitar conta",
		ptPT: "Liquidar conta",
		ru:   "Оплатить счёт",
		tr:   "Fatura öde",
		ukUA: "Оплатити рахунок",
		uz:   "Hisob-fakturani to'lash",
		zhCN: "结算账单",
	},
	COMMAND_TEXT_SETTLE_BILLS: {
		arEG: "تسوية الفواتير",
		de:   "Rechnungen begleichen",
		en:   "Settle bills",
		es:   "Liquidar facturas",
		faIR: "تسویه صورتحساب\u200cها",
		fr:   "Régler les factures",
		id:   "Selesaikan tagihan",
		it:   "Saldare i conti",
		jaJP: "請求書を決済する",
		koKR: "청구서 정산",
		pl:   "Rozlicz rachunki",
		ptBR: "Quitar contas",
		ptPT: "Pagar contas",
		ru:   "Закрыть счета",
		tr:   "Faturaları öde",
		ukUA: "Оплатити рахунки",
		uz:   "Hisob-fakturalarni to'lash",
		zhCN: "结算账单",
	},
	COMMAND_TEXT_INVITE_FIREND: {
		arEG: "دعوة صديق",
		de:   "Freund einladen",
		en:   "Invite friend",
		es:   "Invitar a un amigo",
		faIR: "دوستی دعوت کن",
		fr:   "Inviter un ami",
		id:   "Undang teman",
		it:   "Invita un amico",
		jaJP: "友達を招待",
		koKR: "친구 초대",
		pl:   "Zaproś przyjaciela",
		ptBR: "Convidar amigo",
		ptPT: "Convidar amigo",
		ru:   "Пригласить друга",
		tr:   "Arkadaş davet et",
		ukUA: "Запросити друга",
		uz:   "Do'stni taklif qilish",
		zhCN: "邀请朋友",
	},
	COMMAND_TEXT_INVITE_MEMBER: {
		arEG: "دعوة العضو",
		de:   "Mitglied einladen",
		en:   "Invite member",
		es:   "Invitar miembro",
		faIR: "دعوت از اعضا",
		fr:   "Inviter un membre",
		id:   "Undang anggota",
		it:   "Invita membro",
		jaJP: "メンバーを招待",
		koKR: "멤버 초대",
		pl:   "Zaproś członka",
		ptBR: "Convidar membro",
		ptPT: "Convidar membro",
		ru:   "Пригласить участника",
		tr:   "Üye davet et",
		ukUA: "Запросити учасника",
		uz:   "A'zoni taklif qilish",
		zhCN: "邀请成员",
	},
	COMMAND_TEXT_NEW_BILL: {
		arEG: "مشروع قانون جديد",
		de:   "Neue Rechnung",
		en:   "New bill",
		es:   "Nueva factura",
		faIR: "صورتحساب جدید",
		fr:   "Nouvelle facture",
		id:   "Tagihan baru",
		it:   "Nuova fattura",
		jaJP: "新しい請求書",
		koKR: "새 청구서",
		pl:   "Nowy rachunek",
		ptBR: "Nova conta",
		ptPT: "Nova conta",
		ru:   "Новый счёт",
		tr:   "Yeni fatura",
		ukUA: "Новий рахунок",
		uz:   "Yangi hisob-faktura",
		zhCN: "新账单",
	},
	COMMAND_TEXT_NEW_FUNDRAISING: {
		arEG: "جمع التبرعات الجديد",
		de:   "Neue Spendensammlung",
		en:   "New fundraising",
		es:   "Nueva recaudación de fondos",
		faIR: "جمع آوری پول جدید",
		fr:   "Nouvelle collecte de fonds",
		id:   "Penggalangan dana baru",
		it:   "Nuova raccolta fondi",
		jaJP: "新しい資金調達",
		koKR: "새 모금",
		pl:   "Nowa zbiórka pieniędzy",
		ptBR: "Nova arrecadação de fundos",
		ptPT: "Nova angariação de fundos",
		ru:   "Новый сбор средств",
		tr:   "Yeni bağış toplama",
		ukUA: "Новий збір коштів",
		uz:   "Yangi mablag' yig'ish",
		zhCN: "新筹款",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		arEG: "إضافة جديد",
		de:   "neuer Kontakt",
		en:   "Add new",
		es:   "Añadir",
		faIR: "اضافه کردن مورد جدید",
		fr:   "Ajouter nouveau",
		id:   "Tambah baru",
		it:   "Aggiungi nuovo",
		jaJP: "新規追加",
		koKR: "새로 추가",
		pl:   "Dodaj nowy",
		ptBR: "Adicionar novo",
		ptPT: "Adicionar novo",
		ru:   "Добавить",
		tr:   "Yeni ekle",
		ukUA: "Додати новий",
		uz:   "Yangi qo'shish",
		zhCN: "添加新的",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		arEG: "الكود الخاص بك لتسجيل الدخول إلى التطبيق: <b>%v</b>",
		de:   "Dein Code um dich an der App anzumelden: <b>%v</b>",
		en:   "Your code for signing in to app: <b>%v</b>",
		es:   "Tu código para entrar en la app: <b>%v</b>",
		faIR: "کد شما برای ورود به برنامه: <b>%v</b>",
		fr:   "Votre code pour vous connecter à l'application: <b>%v</b>",
		id:   "Kode Anda untuk masuk ke aplikasi: <b>%v</b>",
		it:   "Il tuo codice per accedere all'app e': <b>%v</b>",
		jaJP: "アプリにサインインするためのコード: <b>%v</b>",
		koKR: "앱에 로그인하기 위한 코드: <b>%v</b>",
		pl:   "Twój kod do logowania do aplikacji: <b>%v</b>",
		ptBR: "Seu código para entrar no aplicativo: <b>%v</b>",
		ptPT: "O seu código para fazer login na aplicação: <b>%v</b>",
		ru:   "Ваш код для входа в приложение: <b>%v</b>",
		tr:   "Uygulamaya giriş yapma kodunuz: <b>%v</b>",
		ukUA: "Ваш код для входу в додаток: <b>%v</b>",
		uz:   "Ilovaga kirish uchun kodingiz: <b>%v</b>",
		zhCN: "您登录应用的验证码: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		arEG: `الرجاء إدخال اسم لجهة الاتصال الجديدة:
 يمكنك الكتابة يدويًا أو الاختيار من دفتر العناوين الخاص بك ( <i>من خلال أيقونة &quot;قص&quot;</i> ).

 <i>أرسل &#39;.&#39; لإلغاء</i>`,
		de: `Bitte gib einen Namen für den neuen Kontakt ein:
		Du kannst in eintippen oder aus deinem Adressbuch wählen (<i>mit dem "Büroklammer"-Symbol und dann Kontakt</i>).

		<i>Send '.' to cancel</i>`,
		en: `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,
		es: `Escribe un nombre para el nuevo contacto:
		Puedes escribirlo o elegirlo de tus contactos (<i>a traves del icono "clip"</i>).

		<i>Enviar '.' para anular</i>`,
		faIR: `لطفا برای مخاطب جدید یک نام وارد کنید:
		میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

		<i>Send '.' برای کنسل کردن</i>`,
		fr: `Veuillez entrer un nom pour le nouveau contact:
		Vous pouvez taper manuellement ou choisir dans votre carnet d'adresses (<i>via l'icône "trombone"</i>).

		<i>Envoyez '.' pour annuler</i>`,
		id: `Silakan masukkan nama untuk kontak baru:
		Anda dapat mengetik secara manual atau memilih dari buku alamat Anda (<i>melalui ikon "klip"</i>).

		<i>Kirim '.' untuk membatalkan</i>`,
		it: `Inserisci un nome per il nuovo contatto:
		Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

		<i>Digita '.' ed invia per annullare</i>`,
		jaJP: `新しい連絡先の名前を入力してください:
		手動で入力するか、アドレス帳から選択できます（<i>「クリップ」アイコンを通じて</i>）。

		<i>キャンセルするには '.' を送信してください</i>`,
		koKR: `새 연락처의 이름을 입력하세요:
		수동으로 입력하거나 주소록에서 선택할 수 있습니다 (<i>"클립" 아이콘을 통해</i>).

		<i>취소하려면 '.'를 보내세요</i>`,
		pl: `Wprowadź nazwę dla nowego kontaktu:
		Możesz wpisać ręcznie lub wybrać z książki adresowej (<i>przez ikonę "spinacza"</i>).

		<i>Wyślij '.' aby anulować</i>`,
		ptBR: `Por favor, digite um nome para o novo contato:
		Você pode digitar manualmente ou escolher do seu livro de endereços (<i>através do ícone "clipe"</i>).

		<i>Envie '.' para cancelar</i>`,
		ptPT: `Por favor, introduza um nome para o novo contacto:
 Pode introduzir manualmente ou escolher na sua lista de endereços ( <i>através do ícone &quot;clip&quot;</i> ).

 <i>Enviar &#39;.&#39; para cancelar</i>`,
		ru: `<b>Имя для нового контакта</b>
		Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).
		<i>Отправьте '.' для отмены</i>`,
		tr: `Yeni kişi için bir isim girin:
		Manuel olarak yazabilir veya adres defterinizden seçebilirsiniz (<i>"ataç" simgesi aracılığıyla</i>).

		<i>İptal etmek için '.' gönderin</i>`,
		ukUA: `Будь ласка, введіть ім'я для нового контакту:
		Ви можете ввести вручну або вибрати з адресної книги (<i>через іконку "скріпка"</i>).

		<i>Надішліть '.' для скасування</i>`,
		uz: `Yangi kontakt uchun ism kiriting:
		Siz qo'lda yozishingiz yoki manzillar kitobingizdan tanlashingiz mumkin (<i>"qisqich" belgisi orqali</i>).

		<i>Bekor qilish uchun '.' yuboring</i>`,
		zhCN: `请输入新联系人的名称:
		您可以手动输入或从通讯录中选择（<i>通过"回形针"图标</i>）。

		<i>发送 '.' 取消</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		arEG: "إنشاء النقل...",
		de:   "Transferiere...",
		en:   "Creating transfer...",
		es:   "Estoy creando la nueva nota...",
		faIR: "ایجاد انتقال ...",
		fr:   "Création du transfert...",
		id:   "Membuat transfer...",
		it:   "Sto creando la nuova voce...",
		jaJP: "転送を作成中...",
		koKR: "전송 생성 중...",
		pl:   "Tworzenie transferu...",
		ptBR: "Criando transferência...",
		ptPT: "Criação de transferência...",
		ru:   "Создаю запись...",
		tr:   "Transfer oluşturuluyor...",
		ukUA: "Створення переказу...",
		uz:   "O'tkazma yaratilmoqda...",
		zhCN: "创建转账中...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		arEG: "انتظر من فضلك",
		de:   "Bitte warten",
		en:   "Please wait",
		es:   "Espera, por favor",
		faIR: "لطفا صبر کنید",
		fr:   "Veuillez patienter",
		id:   "Mohon tunggu",
		it:   "Aspetta per favore",
		jaJP: "お待ちください",
		koKR: "잠시만 기다려주세요",
		pl:   "Proszę czekać",
		ptBR: "Por favor, aguarde",
		ptPT: "Por favor, aguarde",
		ru:   "Пожалуйста подождите",
		tr:   "Lütfen bekleyin",
		ukUA: "Будь ласка, зачекайте",
		uz:   "Iltimos kuting",
		zhCN: "请稍等",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		arEG: "انتظر من فضلك...",
		de:   "Bitte warten...",
		en:   "Please wait...",
		es:   "Espera, por favor...",
		faIR: "لطفا صبر کنید ...",
		fr:   "Veuillez patienter...",
		id:   "Mohon tunggu...",
		it:   "Aspetta per favore...",
		jaJP: "お待ちください...",
		koKR: "잠시만 기다려주세요...",
		pl:   "Proszę czekać...",
		ptBR: "Por favor, aguarde...",
		ptPT: "Por favor aguarde...",
		ru:   "Пожалуйста подождите...",
		tr:   "Lütfen bekleyin...",
		ukUA: "Будь ласка, зачекайте...",
		uz:   "Iltimos kuting...",
		zhCN: "请稍等...",
	},
	MESAGE_TEXT_CREATING_BILL: {
		arEG: "إنشاء مشروع قانون",
		de:   "Rechnung erstellen",
		en:   "Creating bill",
		es:   "Creando factura",
		faIR: "ایجاد صورتحساب",
		fr:   "Création de facture",
		id:   "Membuat tagihan",
		it:   "Creazione di fattura",
		jaJP: "請求書を作成中",
		koKR: "청구서 생성 중",
		pl:   "Tworzenie rachunku",
		ptBR: "Criando conta",
		ptPT: "Criando conta",
		ru:   "Создаётся счёт",
		tr:   "Fatura oluşturuluyor",
		ukUA: "Створення рахунку",
		uz:   "Hisob-faktura yaratilmoqda",
		zhCN: "创建账单中",
	},
	MESSAGE_TEXT_ASK_BILL_CURRENCY: {
		arEG: "ما هي العملة الموجودة في هذا الفاتورة؟",
		de:   "In welcher Währung ist die Rechnung?",
		en:   "What currency is this bill in?",
		es:   "¿En qué moneda está esta factura?",
		faIR: "این صورتحساب به چه ارزی است؟",
		fr:   "Dans quelle devise est cette facture?",
		id:   "Dalam mata uang apa tagihan ini?",
		it:   "In quale valuta è questa fattura?",
		jaJP: "この請求書の通貨は何ですか？",
		koKR: "이 청구서의 통화는 무엇입니까?",
		pl:   "W jakiej walucie jest ten rachunek?",
		ptBR: "Em qual moeda está esta conta?",
		ptPT: "Em que moeda está esta nota?",
		ru:   "В какой валюте этот счёт?",
		tr:   "Bu fatura hangi para biriminde?",
		ukUA: "У якій валюті цей рахунок?",
		uz:   "Bu hisob-faktura qaysi valyutada?",
		zhCN: "这个账单使用什么货币？",
	},
	MESSAGE_TEXT_ASK_BILL_PAYER: {
		arEG: "من دفع الفاتورة؟",
		de:   "Wer hat die Rechnung bezahlt?",
		en:   "Who paid for the bill?",
		es:   "¿Quién pagó la factura?",
		faIR: "چه کسی صورتحساب را پرداخت کرد؟",
		fr:   "Qui a payé la facture?",
		id:   "Siapa yang membayar tagihan?",
		it:   "Chi ha pagato il conto?",
		jaJP: "誰が請求書を支払いましたか？",
		koKR: "누가 청구서를 지불했습니까?",
		pl:   "Kto zapłacił rachunek?",
		ptBR: "Quem pagou a conta?",
		ptPT: "Quem pagou a conta?",
		ru:   "Кто оплатил счёт?",
		tr:   "Faturayı kim ödedi?",
		ukUA: "Хто оплатив рахунок?",
		uz:   "Hisob-fakturani kim to'ladi?",
		zhCN: "谁支付了账单？",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		arEG: "من المتوقع الحصول على الإقرار من %v",
		de:   "%v muss dem zustimmen",
		en:   "Acknowledgement is expected from %v",
		es:   "Se espera la confirmación de %v",
		faIR: "انتظار تصدیق می رود از %v",
		fr:   "Confirmation attendue de %v",
		id:   "Pengakuan diharapkan dari %v",
		it:   "Conferma in attesa da %v",
		jaJP: "%vからの確認が必要です",
		koKR: "%v의 확인이 필요합니다",
		pl:   "Oczekiwane potwierdzenie od %v",
		ptBR: "Confirmação esperada de %v",
		ptPT: "O reconhecimento é esperado de %v",
		ru:   "Подтверждение ожидается от %v",
		tr:   "%v'den onay bekleniyor",
		ukUA: "Очікується підтвердження від %v",
		uz:   "%v dan tasdiqlash kutilmoqda",
		zhCN: "等待%v的确认",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		arEG: "لقد قبلت هذه المعاملة.",
		de:   "Du hast dem zugestimmt.",
		en:   "You've accepted this transaction.",
		es:   "Has confirmado esta transacción",
		faIR: ".شما این تراکنش را قبول کردید ",
		fr:   "Vous avez accepté cette transaction.",
		id:   "Anda telah menerima transaksi ini.",
		it:   "Hai accettato il debito/credito.",
		jaJP: "このトランザクションを承認しました。",
		koKR: "이 거래를 수락했습니다.",
		pl:   "Zaakceptowałeś tę transakcję.",
		ptBR: "Você aceitou esta transação.",
		ptPT: "Aceitou esta transação.",
		ru:   "Вы подтвердили эту транзакцию.",
		tr:   "Bu işlemi kabul ettiniz.",
		ukUA: "Ви підтвердили цю транзакцію.",
		uz:   "Siz ushbu tranzaksiyani qabul qildingiz.",
		zhCN: "您已接受此交易。",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		arEG: `أنت لا توافق على هذه المعاملة.
 لن يتم حذف المعاملة ولكن سيتم إخطار المنشئ.`,
		de: `Du hast dieser Anfrage nicht zugestimmt.
		Der Vorgang wird zurückgestellt und die Gegenpartei benachrichtigt.`,
		en: `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,
		es: `No estas de acuerdo con la transacción.
		La transacción NO será cancelada, pero el creador será notificado.`,
		faIR: `شما با این تراکنش موافق نیستید.
		تراکنش حذف نخواهد شد اما به ایجاد کننده اطلاع داده خواهد شد.`,
		fr: `Vous n'êtes pas d'accord avec cette transaction.
		La transaction ne sera pas supprimée mais le créateur sera notifié.`,
		id: `Anda tidak setuju dengan transaksi ini.
		Transaksi tidak akan dihapus tetapi pembuat akan diberi tahu.`,
		it: `Hai rifiutato il debito/credito.
		La transazione non sarà eliminata ma il creatore sarà avvisato.`,
		jaJP: `このトランザクションに同意しません。
		トランザクションは削除されませんが、作成者に通知されます。`,
		koKR: `이 거래에 동의하지 않습니다.
		거래는 삭제되지 않지만 생성자에게 알림이 갑니다.`,
		pl: `Nie zgadzasz się z tą transakcją.
		Transakcja nie zostanie usunięta, ale twórca zostanie powiadomiony.`,
		ptBR: `Você não concorda com esta transação.
		A transação não será excluída, mas o criador será notificado.`,
		ptPT: `Não concorda com esta transação.
 A transação não será eliminada, mas o criador será notificado.`,
		ru: `Вы НЕ согласны с этой транзакцией.

Сама транзакция НЕ будет отменена, но создатель будет оповещён.`,
		tr: `Bu işleme katılmıyorsunuz.
		İşlem silinmeyecek ancak oluşturucu bilgilendirilecek.`,
		ukUA: `Ви не згодні з цією транзакцією.
		Транзакція не буде видалена, але творець буде повідомлений.`,
		uz: `Siz ushbu tranzaksiyaga rozi emassiz.
		Tranzaksiya o'chirilmaydi, lekin yaratuvchi xabardor qilinadi.`,
		zhCN: `您不同意此交易。
		交易不会被删除，但创建者会收到通知。`,
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		arEG: "تم قبول معاملتك بنسبة %v:",
		de:   "%v hat deiner Anfrage <b>zugestimmt</b>:",
		en:   "%v accepted your transaction:",
		es:   "%v ha aceptado tu transacción",
		faIR: ": تراکنش شمارا تایید کرد %v ",
		fr:   "%v a accepté votre transaction:",
		id:   "%v menerima transaksi Anda:",
		it:   "%v ha accettato il tuo credito/debito:",
		jaJP: "%vがあなたの取引を承認しました:",
		koKR: "%v님이 귀하의 거래를 수락했습니다:",
		pl:   "%v zaakceptował(a) twoją transakcję:",
		ptBR: "%v aceitou sua transação:",
		ptPT: "%v aceitou a sua transação:",
		ru:   "%v подтвердил(a) вашу транзакцию:",
		tr:   "%v işleminizi kabul etti:",
		ukUA: "%v підтвердив(ла) вашу транзакцію:",
		uz:   "%v sizning tranzaksiyangizni qabul qildi:",
		zhCN: "%v接受了您的交易：",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		arEG: "لم يوافق %v على معاملتك. لم يتم إلغاء المعاملة، ولكن قد ترغب في مناقشتها.",
		de:   "%v hat deine Anfrage <b>abgelehnt</b>. Wenn die Sache besprochen ist, kann die Anfrage erneut gesendet werden.",
		en:   "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.",
		es:   "%v no está de acuerdo con tu transacción. La transacción no ha sido cancelada, pero quizás deberías discutirlo.",
		faIR: "%v با تراکنش شما موافقت نکرد. تراکنش لغو نشده است اما ممکن است بخواهید در مورد آن صحبت کنید.",
		fr:   "%v n'est pas d'accord avec votre transaction. La transaction n'est pas annulée mais vous voudrez peut-être en discuter.",
		id:   "%v tidak setuju dengan transaksi Anda. Transaksi tidak dibatalkan tetapi Anda mungkin ingin mendiskusikannya.",
		it:   "%v non è d'accordo con la tua transazione. La transazione non è annullata ma potresti voler discuterne.",
		jaJP: "%vはあなたの取引に同意しませんでした。取引はキャンセルされていませんが、話し合いたいかもしれません。",
		koKR: "%v님이 귀하의 거래에 동의하지 않았습니다. 거래는 취소되지 않았지만 논의하고 싶을 수 있습니다.",
		pl:   "%v nie zgadza się z twoją transakcją. Transakcja nie jest anulowana, ale możesz chcieć to omówić.",
		ptBR: "%v não concordou com sua transação. A transação não foi cancelada, mas você pode querer discutir isso.",
		ptPT: "%v não concordou com a sua transação. A transação não foi cancelada, mas pode querer discuti-la.",
		ru:   "%v <b>НЕ</b> подтвердил(a) вашу транзакцию. Транзакция не отменена, но возможно вам стоит это обсудить.",
		tr:   "%v işleminizle aynı fikirde değil. İşlem iptal edilmedi ancak bunu tartışmak isteyebilirsiniz.",
		ukUA: "%v не погодився з вашою транзакцією. Транзакція не скасована, але, можливо, ви захочете це обговорити.",
		uz:   "%v sizning tranzaksiyangizga rozi bo'lmadi. Tranzaksiya bekor qilinmadi, lekin siz bu haqda muhokama qilishni xohlashingiz mumkin.",
		zhCN: "%v不同意您的交易。交易未取消，但您可能想讨论一下。",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		arEG: "أريد التطبيق!",
		de:   "Ich will die App!",
		en:   "I want the app!",
		es:   "¡Quiero la aplicación!",
		faIR: "!من برنامه را می خواهم",
		fr:   "Je veux l'application !",
		id:   "Saya ingin aplikasinya!",
		it:   "Voglio l'aplicazione!",
		jaJP: "アプリが欲しいです！",
		koKR: "앱이 필요합니다!",
		pl:   "Chcę aplikację!",
		ptBR: "Eu quero o aplicativo!",
		ptPT: "Eu quero a aplicação!",
		ru:   "Хочу приложение!",
		tr:   "Uygulamayı istiyorum!",
		ukUA: "Я хочу додаток!",
		uz:   "Men ilovani xohlayman!",
		zhCN: "我想要应用程序！",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		arEG: "أنا بخير مع الروبوت فقط!",
		de:   "Mir reicht der Bot!",
		en:   "I'm fine with just the bot!",
		es:   "¡Estoy satisfecho con este bot!",
		faIR: "! ربات به تنهایی برای من کافی است",
		fr:   "Le bot me suffit !",
		id:   "Saya cukup dengan bot saja!",
		it:   "Mi accontento del bot per ora!",
		jaJP: "ボットだけで大丈夫です！",
		koKR: "봇만으로도 괜찮습니다!",
		pl:   "Wystarczy mi sam bot!",
		ptBR: "Estou bem apenas com o bot!",
		ptPT: "Estou bem só com o bot!",
		ru:   "Меня вполне устраивает бот!",
		tr:   "Sadece bot ile iyiyim!",
		ukUA: "Мені достатньо лише бота!",
		uz:   "Men faqat bot bilan yaxshiman!",
		zhCN: "我对只使用机器人很满意！",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		arEG: "سنخبرك بمجرد أن يصبح التطبيق متاحًا للتنزيل.",
		de:   "Du wirst darüber informiert, wenn die App zum Download zur Verfügung steht.",
		en:   "We'll let you know once the app is available for download.",
		es:   "Te avisamos cuando la aplicación esté disponible para descargarla",
		faIR: ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
		fr:   "Nous vous informerons dès que l'application sera disponible au téléchargement.",
		id:   "Kami akan memberi tahu Anda setelah aplikasi tersedia untuk diunduh.",
		it:   "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
		jaJP: "アプリがダウンロード可能になり次第お知らせします。",
		koKR: "앱이 다운로드 가능해지면 알려드리겠습니다.",
		pl:   "Poinformujemy Cię, gdy aplikacja będzie dostępna do pobrania.",
		ptBR: "Avisaremos você assim que o aplicativo estiver disponível para download.",
		ptPT: "Avisaremos quando a aplicação estiver disponível para download.",
		ru:   "Мы сообщим вам когда приложение будет доступно для загруки.",
		tr:   "Uygulama indirilebilir olduğunda size haber vereceğiz.",
		ukUA: "Ми повідомимо вас, коли додаток буде доступний для завантаження.",
		uz:   "Ilova yuklab olish uchun mavjud bo'lganda sizga xabar beramiz.",
		zhCN: "一旦应用程序可供下载，我们会通知您。",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		arEG: "حسنًا، نحن سعداء لأن الروبوت الخاص بنا جيد بما فيه الكفاية وليس هناك حاجة لتنزيل تطبيق.",
		de:   "Gut, wir sind froh, dass dir der Bot reicht und wir uns mit der App nicht beeilen müssen.",
		en:   "Well, we are happy our bot is good enough and there is no need to download an app.",
		es:   "Bueno, estamos contentos de que te haya gustado nuestro bot y no hace falta descargar ninguna otra aplicación",
		faIR: ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
		fr:   "Eh bien, nous sommes heureux que notre bot soit suffisant et qu'il n'y ait pas besoin de télécharger une application.",
		id:   "Baiklah, kami senang bot kami cukup baik dan tidak perlu mengunduh aplikasi.",
		it:   "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
		jaJP: "私たちのボットで十分であり、アプリをダウンロードする必要がないことを嬉しく思います。",
		koKR: "네, 저희 봇이 충분히 좋아서 앱을 다운로드할 필요가 없다니 기쁩니다.",
		pl:   "Cóż, cieszymy się, że nasz bot jest wystarczająco dobry i nie ma potrzeby pobierania aplikacji.",
		ptBR: "Bem, estamos felizes que nosso bot seja bom o suficiente e não há necessidade de baixar um aplicativo.",
		ptPT: "Bem, estamos felizes por o nosso bot ser bom o suficiente e não haver necessidade de descarregar uma aplicação.",
		ru:   "Что ж, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		tr:   "Peki, botumuzun yeterince iyi olduğuna ve bir uygulama indirmeye gerek olmadığına sevindik.",
		ukUA: "Що ж, ми раді, що наш бот достатньо хороший і немає потреби завантажувати додаток.",
		uz:   "Yaxshi, botimiz yetarlicha yaxshi ekanligi va ilovani yuklab olish kerak emasligi bizni xursand qiladi.",
		zhCN: "好的，我们很高兴我们的机器人已经足够好，不需要下载应用程序。",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		arEG: "يمكنك <a href>الإعلان هنا</a>",
		de:   "Hier könnte <a href>ihre Werbung</a> stehen",
		en:   "You can <a href>advertise here</a>",
		es:   "Aquí se puede <a href>publicar un anuncio</a>",
		faIR: "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
		fr:   "Vous pouvez <a href>faire de la publicité ici</a>",
		id:   "Anda dapat <a href>beriklan di sini</a>",
		it:   "Puoi <a href>pubblicizzare qui</a>",
		jaJP: "ここに<a href>広告を掲載</a>できます",
		koKR: "여기에 <a href>광고할 수 있습니다</a>",
		pl:   "Możesz <a href>reklamować się tutaj</a>",
		ptBR: "Você pode <a href>anunciar aqui</a>",
		ptPT: "Pode <a href>anunciar aqui</a>",
		ru:   "Здесь можно <a href>разместить рекламу</a>",
		tr:   "Burada <a href>reklam verebilirsiniz</a>",
		ukUA: "Ви можете <a href>розмістити рекламу тут</a>",
		uz:   "Siz bu yerda <a href>reklama bera olasiz</a>",
		zhCN: "您可以在这里<a href>做广告</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		arEG: `🤖: أنا روبوت جيد، بالتأكيد. ولكن في بعض الأحيان يكون من الأنسب استخدام تطبيق متخصص لطيف. إنه ليس جاهزًا للاستخدام العام بعد ولكن يمكنك التحقق من شكله: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a> 

 هل تريد الحصول على دعوة عند إصداره؟`,
		de: `🤖: Ich hin ein guter Roboter - klar. Aber manchmal kommt es besser eine eigene App für etwas zu haben. Es ist noch nicht ganz fertig, aber falls du schonmal reinschauen willst: <a href="https://debtstracker.io/de/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Möchtest du daran erinnert werden, wenn die App rauskommt?`,
		en: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,
		es: `🤖: Claro que soy un robot encantador, pero más comodo usar una aplicación especial.No esta disponible ya pero se puede ver como será: <a href = "https://debtstracker.io/es/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	¿Quieres que te avisemos cuando esté lista?`,
		faIR: `🤖: مطمئناً من روبات خوبی هستم. اما بعضی وقت هاساده تر و مناسب تر است که از یک برنامه به خوبی تخصصی شده استفاده شود، این برنامه هنوز برای استفاده عموم آماده نیست ولی می توانید چک کنید که چگونه به نظر خواهد رسید: <a href="https://debtstracker.io/fa/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	آیا می خواهید وقتی منتشر شد دعوتنامه ای دریافت کنید؟`,
		fr: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/fr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/fr/</a>

	Do you want to get an invite when it gets released?`,
		id: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/id/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/id/</a>

	Do you want to get an invite when it gets released?`,
		it: `🤖: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href="https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

	Vuoi essere invitato non appena viene rilasciata?`,
		jaJP: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ja/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ja/</a>

	Do you want to get an invite when it gets released?`,
		koKR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ko/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ko/</a>

	Do you want to get an invite when it gets released?`,
		pl: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pl/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pl/</a>

	Do you want to get an invite when it gets released?`,
		ptBR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pt/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pt/</a>

	Do you want to get an invite when it gets released?`,
		ptPT: `🤖: Sou um bom robô, com certeza. Mas, por vezes, é mais conveniente usar uma aplicação especializada. Ainda não está pronto para uso público, mas pode consultar como ficará: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a> 

 Quer receber um convite quando for lançado?`,
		ru: `🤖: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

		Хотите получить оповещение когда оно выйдет?`,
		tr: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/tr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/tr/</a>

	Do you want to get an invite when it gets released?`,
		ukUA: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ua/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ua/</a>

	Do you want to get an invite when it gets released?`,
		uz: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/uz/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/uz/</a>

	Do you want to get an invite when it gets released?`,
		zhCN: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/zh/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/zh/</a>

	Do you want to get an invite when it gets released?`,
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		arEG: "عذراً، ولكن يمكنك استخدام الأرقام فقط كمبلغ/كمية ( <i>مع ما يصل إلى رقمين بعد النقطة</i> ).",
		de:   "Entschuldigung, aber du kannst nur Zahlen für Menge oder Wert wählen (<i>mit zwei Nachkommastellen</i>).",
		en:   "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		es:   "Lo siento, solo puedes utilizar numeros como importe/cantidad (<i>con un maximo de 2 dígitos despues de la coma</i>).",
		faIR: "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
		fr:   "Désolé, mais vous pouvez utiliser uniquement des chiffres comme montant/quantité ( <i>avec jusqu&#39;à 2 chiffres après le point</i> ).",
		id:   "Maaf, tetapi Anda dapat menggunakan angka saja sebagai jumlah/kuantitas ( <i>dengan maksimal 2 digit setelah titik</i> ).",
		it:   "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
		jaJP: "申し訳ございませんが、金額/数量には数字のみ（<i>小数点以下2桁まで</i>）をご利用いただけます。",
		koKR: "죄송하지만, 금액/수량에는 숫자만 사용 가능합니다( <i>소숫점 이하 두 자리까지</i> ).",
		pl:   "Przepraszamy, ale jako kwotę/ilość możesz użyć samych liczb ( <i>do 2 cyfr po przecinku</i> ).",
		ptBR: "Desculpe, mas você pode usar apenas números como quantidade ( <i>com até 2 dígitos depois do ponto</i> ).",
		ptPT: "Desculpe, mas pode usar apenas números como quantidade ( <i>com até 2 dígitos depois do ponto</i> ).",
		ru:   "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаков после запятой</i>).",
		tr:   "Üzgünüm, miktar/miktar olarak sadece rakam kullanabilirsiniz ( <i>noktadan sonra en fazla 2 rakamla</i> ).",
		ukUA: "Вибачте, але ви можете використовувати лише числа як суму/кількість ( <i>до 2 цифр після коми</i> ).",
		uz:   "Kechirasiz, lekin siz miqdor/miqdor sifatida faqat raqamlardan foydalanishingiz mumkin ( <i>nuqtadan keyin 2 tagacha raqam bilan</i> ).",
		zhCN: "抱歉，但您只能使用数字作为金额/数量（<i>小数点后最多 2 位数字</i>）。"},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		arEG: "<b>ماذا أقرضته لشخص ما؟</b>",
		de:   "<b>Was hast du jemanden geliehen?</b>",
		en:   "<b>What did you lend to someone?</b>",
		es:   "<b>¿Qué has prestado?</b>",
		faIR: "<b> چه چیزی به کسی قرض داده اید؟</b>",
		fr:   "<b>Qu&#39;as-tu prêté à quelqu&#39;un ?</b>",
		id:   "<b>Apa yang Anda pinjamkan kepada seseorang?</b>",
		it:   "<b>Cos'hai prestato?</b>",
		jaJP: "<b>誰かに何を貸しましたか？</b>",
		koKR: "<b>당신은 다른 사람에게 무엇을 빌려주었나요?</b>",
		pl:   "<b>Co pożyczyłeś komuś?</b>",
		ptBR: "<b>O que você emprestou para alguém?</b>",
		ptPT: "<b>O que emprestou a alguém?</b>",
		ru:   "<b>Что вы дали в долг?</b>",
		tr:   "<b>Birine ne ödünç verdin?</b>",
		ukUA: "<b>Що ти комусь позичив?</b>",
		uz:   "<b>Birovga nima qarz berdingiz?</b>",
		zhCN: "<b>你借给别人什么了？</b>"},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {
		arEG: `يرجى الاختيار من الخيارات أدناه أو <a>تحديد العملة من القائمة</a> .

 إذا لم تكن الخيارات القياسية كافية، فما عليك سوى إرسال رسالة نصية. على سبيل المثال: &quot; <i>apple</i> &quot;.`,
		de: `Bitte wähle <a>eine Währung aus der Liste</a>.

	Falls die Standardoptionen nicht reichen, sende mir einen Text. Zum Beispiel: <i>Äpfel</i>".`,
		en: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,
		es: `Elige del menú abajo de la pantalla o <a>selecciona la moneda de la lista</a>.

	Si no encuentras la opción correcta simplemente envía un texto. Por ejemplo: "<i>manzana</i>".`,
		faIR: `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

	اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال:. "<i>سیب</i>".`,
		fr: `Veuillez choisir parmi les options ci-dessous ou <a>sélectionnez une devise dans la liste</a> .

 Si les options standard ne suffisent pas, envoyez simplement un SMS. Par exemple : «<i> pomme</i> ».`,
		id: `Silakan pilih dari pilihan di bawah ini atau <a>pilih mata uang dari daftar</a> .

 Jika pilihan standar tidak cukup, kirimkan saja teks. Misalnya: &quot; <i>apple</i> &quot;.`,
		it: `Scegli dalle opzioni qui sotto o <a>seleziona una valuta dalla lista</a>.

	Se le opzioni standard non bastano semplicemente invia un testo.Per esempio: "<i>mele</i>".`,
		jaJP: `以下のオプションから選択するか、<a>リストから通貨を選択して</a>ください。

 標準オプションでは不十分な場合は、テキストを送信するだけです。例: &quot; <i>apple</i> &quot;。`,
		koKR: `아래 옵션에서 선택하거나 <a>목록에서 통화를 선택</a> 하세요.

 표준 옵션이 충분하지 않으면 문자 메시지를 보내세요. 예: &quot; <i>apple</i> &quot;.`,
		pl: `Wybierz jedną z poniższych opcji lub <a>wybierz walutę z listy</a> .

 Jeśli standardowe opcje nie wystarczą, po prostu wyślij SMS-a. Na przykład: „ <i>apple</i> ”.`,
		ptBR: `Escolha entre as opções abaixo ou <a>selecione uma moeda da lista</a> .

 Se as opções padrão não forem suficientes, basta enviar uma mensagem de texto. Por exemplo: &quot; <i>apple</i> &quot;.`,
		ptPT: `Escolha entre as opções abaixo ou <a>selecione uma moeda da lista</a> .

 Se as opções padrão não forem suficientes, basta enviar uma mensagem de texto. Por exemplo: &quot; <i>maçã</i> &quot;.`,
		ru: `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

	Если ни один из стандартных вариантов не подходит просто напишите текстом.Например: "<i>яблоко</i>".`,
		tr: `Lütfen aşağıdaki seçeneklerden birini seçin veya <a>listeden bir para birimi seçin</a> . 

 Eğer standart seçenekler yeterli olmazsa, sadece bir mesaj gönderin. Örneğin: &quot; <i>apple</i> &quot;.`,
		ukUA: `Будь ласка, виберіть один із наведених нижче варіантів або <a>валюту зі списку</a> .

 Якщо стандартних опцій недостатньо, просто надішліть текстове повідомлення. Наприклад: « <i>apple</i> ».`,
		uz: `Quyidagi variantlardan tanlang yoki <a>roʻyxatdan valyutani tanlang</a> .

 Agar standart variantlar yetarli boʻlmasa, shunchaki matn yuboring. Masalan: &quot; <i>olma</i> &quot;.`,
		zhCN: `请从以下选项中选择或<a>从列表中选择一种货币</a>。

 如果标准选项不够，只需发送文本即可。例如：“<i>苹果</i>”。`},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		arEG: "كم <b>%v</b> أقرضته؟\n( <i>أرسل &#39;.&#39; للإلغاء</i> )",
		de:   "Wie viel <b>%v</b> hast du verliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		en:   "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		es:   "Cuanto(s) <b>%v</b> has prestado\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		fr:   "Combien <b>%v</b> avez-vous prêté ?\n( <i>envoyez &#39;.&#39; pour annuler</i> )",
		id:   "Berapa <b>%v</b> yang Anda pinjamkan?\n( <i>kirim &#39;.&#39; untuk membatalkan</i> )",
		it:   "Quanti <b>%v</b> hai prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "いくら<b>%v</b>を貸しましたか?\n(<i>キャンセルするには &#39;.&#39; を送信してください</i>)",
		koKR: "얼마 <b>%v를</b> 빌려주셨나요?\n( <i>취소하려면 &#39;.&#39;을 보내세요</i> )",
		pl:   "Ile <b>%v</b> pożyczyłeś?\n( <i>wyślij &#39;.&#39; aby anulować</i> )",
		ptBR: "Quanto <b>%v</b> você emprestou?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ptPT: "Quanto <b>%v</b> pediu emprestado?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ru:   "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		tr:   "Ne kadar <b>%v</b> borç verdiniz?\n( <i>iptal etmek için &#39;.&#39; gönderin</i> )",
		ukUA: "Скільки <b>%v</b> ви позичили?\n( <i>надішліть &#39;.&#39; для скасування</i> )",
		uz:   "Qancha <b>%v</b> qarz berdingiz?\n( <i>bekor qilish uchun “.” so‘zini yuboring</i> )",
		zhCN: "您借出了多少<b>%v</b> ？\n（<i>发送‘.’取消</i>）"},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		arEG: "من اقترض منك <b>%v</b> ؟\n( <i>إرسال &#39;.&#39; للإلغاء</i> )",
		de:   "Wer hat sich <b>%v</b> von dir geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		en:   "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		es:   "A quién has prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه کسی از شما <b>%v</b> قرض گرفته است؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		fr:   "Qui vous a emprunté <b>%v</b> ?\n( <i>envoyer &#39;.&#39; pour annuler</i> )",
		id:   "Siapa yang meminjam dari Anda <b>%v</b> ?\n( <i>kirim &#39;.&#39; untuk membatalkan</i> )",
		it:   "Chi e' in debito di <b>%v</b> con te?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "あなたから<b>%v</b>を借りたのは誰ですか?\n(<i>キャンセルするには &#39;.&#39; を送信してください</i>)",
		koKR: "누가 당신에게 <b>%v 을(를)</b> 빌렸나요?\n( <i>취소하려면 &#39;.&#39;을(를) 보내세요</i> )",
		pl:   "Kto pożyczył od Ciebie <b>%v</b> ?\n( <i>wyślij &#39;.&#39; aby anulować</i> )",
		ptBR: "Quem pegou emprestado de você <b>%v</b> ?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ptPT: "Quem lhe pediu emprestado <b>%v</b> ?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ru:   "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		tr:   "Sizden <b>%v</b> kim ödünç aldı?\n( <i>iptal etmek için &#39;.&#39; gönderin</i> )",
		ukUA: "Хто позичив у вас <b>%v</b> ?\n( <i>надішліть &#39;.&#39; для скасування</i> )",
		uz:   "Sizdan kim qarz oldi <b>%v</b> ?\n( <i>bekor qilish uchun &#39;.&#39; deb yuboring</i> )",
		zhCN: "谁向你借了<b>%v</b> ？\n（<i>发送‘.’取消</i>）"},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		arEG: "ماذا اقرضت؟",
		de:   "Was hast du dir geliehen?",
		en:   "What did you lend?",
		es:   "¿Qué te han prestado?",
		faIR: "چه چیزی قرض گرفته اید؟",
		fr:   "Qu&#39;as-tu prêté ?",
		id:   "Apa yang kamu pinjamkan?",
		it:   "Cosa ti hanno prestato?",
		jaJP: "何を貸したんですか？",
		koKR: "무엇을 빌려주셨나요?",
		pl:   "Co pożyczyłeś?",
		ptBR: "O que você emprestou?",
		ptPT: "O que emprestou?",
		ru:   "Что вы взяли в долг?",
		tr:   "Ne ödünç verdin?",
		ukUA: "Що ви позичили?",
		uz:   "Nima qarz berdingiz?",
		zhCN: "你借了什么？"},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		arEG: "كم <b>%v</b> اقترضت؟\n( <i>أرسل &#39;.&#39; للإلغاء</i> )",
		de:   "Wie viel <b>%v</b> hast du geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		en:   "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		es:   "¿Cuánto <b>%v</b> has prestado?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		fr:   "Combien <b>%v</b> avez-vous emprunté ?\n( <i>envoyez &#39;.&#39; pour annuler</i> )",
		id:   "Berapa banyak <b>%v</b> yang Anda pinjam?\n( <i>kirim &#39;.&#39; untuk membatalkan</i> )",
		it:   "Quanti <b>%v</b> ti hanno prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "いくら<b>%v</b>を借りましたか?\n(<i>キャンセルするには &#39;.&#39; を送信してください</i>)",
		koKR: "얼마나 <b>%v를</b> 빌렸나요?\n( <i>취소하려면 &#39;.&#39;을 보내세요</i> )",
		pl:   "Ile <b>%v</b> pożyczyłeś?\n( <i>wyślij &#39;.&#39; aby anulować</i> )",
		ptBR: "Quanto <b>%v</b> você pegou emprestado?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ptPT: "Quanto <b>%v</b> pediu emprestado?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ru:   "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		tr:   "Ne kadar <b>%v</b> borç aldın?\n( <i>iptal etmek için &#39;.&#39; gönder</i> )",
		ukUA: "Скільки <b>%v</b> ви позичили?\n( <i>надішліть &#39;.&#39; для скасування</i> )",
		uz:   "Qancha <b>%v</b> qarz oldingiz?\n( <i>bekor qilish uchun &#39;.&#39; yuboring</i> )",
		zhCN: "您借了多少<b>%v</b> ？\n（<i>发送‘.’取消</i>）"},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		arEG: "من أقرضك <b>%v</b> ؟\n( <i>إرسال &#39;.&#39; للإلغاء</i> )",
		de:   "Wer hat dir <b>%v</b> geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		en:   "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		es:   "¿Quién te ha prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه کسی به شما <b>%v</b> قرض داده است؟ \n(<i>ارسال '.' برای کنسل کردن</i>)",
		fr:   "Qui vous a prêté <b>%v</b> ?\n( <i>envoyer &#39;.&#39; pour annuler</i> )",
		id:   "Siapa yang meminjamkanmu <b>%v</b> ?\n( <i>kirim &#39;.&#39; untuk membatalkan</i> )",
		it:   "Chi ti ha prestato <b>%v</b>?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "誰があなたに<b>%v</b>を貸したのですか?\n(<i>キャンセルするには &#39;.&#39; を送信してください</i>)",
		koKR: "누가 당신에게 <b>%v을(를)</b> 빌려줬나요?\n( <i>취소하려면 &#39;.&#39;을(를) 보내세요</i> )",
		pl:   "Kto Ci pożyczył <b>%v</b> ?\n( <i>wyślij &#39;.&#39; aby anulować</i> )",
		ptBR: "Quem te emprestou <b>%v</b> ?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ptPT: "Quem lhe emprestou <b>%v</b> ?\n( <i>envie &#39;.&#39; para cancelar</i> )",
		ru:   "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		tr:   "Sana <b>%v&#39;yi</b> kim ödünç verdi?\n( <i>iptal etmek için &#39;.&#39; gönder</i> )",
		ukUA: "Хто позичив вам <b>%v</b> ?\n( <i>надішліть &#39;.&#39; для скасування</i> )",
		uz:   "Sizga kim <b>%v</b> qarz berdi ?\n( <i>bekor qilish uchun &#39;.&#39; yuboring</i> )",
		zhCN: "谁借给您<b>%v</b> ？\n（<i>发送‘.’取消</i>）"},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		arEG: "هل يجب علينا إرسال <a receipt>الإيصال</a> إلى <a counterparty>%v</a> ؟",
		de:   "Soll eine <a receipt>Quittung</a> an <a counterparty>%v</a> gesendet werden?",
		en:   "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		es:   "¿Debo enviar <a receipt> el recibo</a> a <a counterparty>%v</a>?",
		faIR: "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
		fr:   "Devrions-nous envoyer un <a receipt>reçu</a> à <a counterparty>%v</a> ?",
		id:   "Haruskah kami mengirim <a receipt>tanda terima</a> ke <a counterparty>%v</a> ?",
		it:   "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
		jaJP: "<a counterparty>%v</a>に<a receipt>領収書</a>を送りますか?",
		koKR: "<a counterparty>%v</a> 에게 <a receipt>영수증을</a> 보내야 할까요?",
		pl:   "Czy powinniśmy wysłać <a receipt>potwierdzenie</a> do <a counterparty>%v</a> ?",
		ptBR: "Devemos enviar um <a receipt>recibo</a> para <a counterparty>%v</a> ?",
		ptPT: "Devemos enviar um <a receipt>recibo</a> para <a counterparty>%v</a> ?",
		ru:   "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		tr:   "<a counterparty>%v&#39;e</a> <a receipt>makbuz</a> göndermeli miyiz?",
		ukUA: "Чи варто нам надіслати <a receipt>квитанцію</a> до <a counterparty>%v</a> ?",
		uz:   "<a counterparty>%v</a> ga <a receipt>kvitansiya</a> yuborishimiz kerakmi?",
		zhCN: "我们应该向<a counterparty>%v</a>发送<a receipt>收据</a>吗？"},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		arEG: "عذرًا، إرسال إيصال لنفسك عبر الرسائل النصية غير متاح حاليًا. يمكنك إرساله إلى %v.",
		de:   "Entschuldigung, aber eine Quittung selber per SMS zu schicken ist im Moment noch nicht möglich. Aber dafür geht es mit %v.",
		en:   "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		es:   "Lo siento, el envio del recibo a ti mismo a través de SMS en este momento está desactivado. Pero lo puedes enviar a %v.",
		faIR: "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
		fr:   "Désolé, l&#39;envoi de reçu par SMS n&#39;est pas disponible pour le moment. Vous pouvez cependant l&#39;envoyer à %v.",
		id:   "Maaf, saat ini pengiriman tanda terima ke diri Anda sendiri melalui SMS tidak tersedia. Anda dapat mengirimkannya ke %v.",
		it:   "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
		jaJP: "申し訳ございませんが、現在SMSで領収書を自分宛に送信することはできません。ただし、%v宛てに送信することは可能です。",
		koKR: "죄송합니다. 현재 본인에게 SMS로 영수증을 보내실 수 없습니다. 하지만 %v에게는 보내실 수 있습니다.",
		pl:   "Przepraszamy, wysyłanie paragonu do siebie przez SMS nie jest obecnie dostępne. Możesz jednak wysłać go do %v.",
		ptBR: "Desculpe, o envio de recibo para você mesmo por SMS não está disponível no momento. Você pode enviá-lo para %v.",
		ptPT: "Lamentamos, o envio de recibo para si mesmo por SMS não está disponível de momento. Pode enviá-lo para %v.",
		ru:   "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		tr:   "Üzgünüz, kendinize SMS ile makbuz gönderme şu anda mümkün değil. Yine de %v&#39;ye gönderebilirsiniz.",
		ukUA: "Вибачте, надіслати квитанцію собі через SMS наразі неможливо. Однак ви можете надіслати її %v.",
		uz:   "Kechirasiz, o&#39;zingizga SMS orqali kvitansiya yuborish hozircha mavjud emas. Siz uni %v ga yuborishingiz mumkin.",
		zhCN: "抱歉，目前无法通过短信向您自己发送收据。不过，您可以将其发送至 %v。"},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		arEG: "نقوم بإرسال الإيصال إلى %v عبر Telegram...",
		de:   "Die Quittung wird %v per Telegram geschickt...",
		en:   "We are sending receipt to %v by Telegram...",
		es:   "El recibo está enviando a%v a través de Telegram…",
		faIR: "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
		fr:   "Nous envoyons un reçu à %v par Telegram...",
		id:   "Kami mengirimkan tanda terima ke %v melalui Telegram...",
		it:   "Sto inviando la notifica a %v tramite Telegram...",
		jaJP: "領収書をTelegramで%vに送信しています...",
		koKR: "%v에게 Telegram을 통해 영수증을 보내고 있습니다...",
		pl:   "Wysyłamy potwierdzenie do %v przez Telegram...",
		ptBR: "Estamos enviando o recibo para %v pelo Telegram...",
		ptPT: "Estamos a enviar o recibo para %v pelo Telegram...",
		ru:   "Отправляем для %v извещение через Telegram...",
		tr:   "%v&#39;e Telegram ile makbuz gönderiyoruz...",
		ukUA: "Ми надсилаємо квитанцію %v через Telegram...",
		uz:   "Telegram orqali %v ga kvitansiya yubormoqdamiz...",
		zhCN: "我们正在通过 Telegram 向 %v 发送收据..."},
	DAY: {
		arEG: "%v يوم",
		de:   "%v Tag",
		en:   "%v day",
		es:   "%v día",
		faIR: "%v روز",
		fr:   "%v jour",
		id:   "%v hari",
		it:   "%v giorno",
		jaJP: "%v日",
		koKR: "%v일",
		pl:   "%v dzień",
		ptBR: "%v dia",
		ptPT: "%v dia",
		ru:   "%v день",
		tr:   "%v gün",
		ukUA: "%v день",
		uz:   "%v kun",
		zhCN: "%v天"},
	DAYS_234: {
		arEG: "%v أيام",
		de:   "%v Tage",
		en:   "%v days",
		es:   "%v días",
		faIR: "%v روز",
		fr:   "%v jours",
		id:   "%v hari",
		it:   "%v giorni",
		jaJP: "%v日",
		koKR: "%v일",
		pl:   "%v dni",
		ptBR: "%v dias",
		ptPT: "%v dias",
		ru:   "%v дня",
		tr:   "%v gün",
		ukUA: "%v днів",
		uz:   "%v kun",
		zhCN: "%v 天"},
	DAYS: {
		arEG: "%v أيام",
		de:   "%v Tage",
		en:   "%v days",
		es:   "%v días",
		faIR: "%v روز",
		fr:   "%v jours",
		id:   "%v hari",
		it:   "%v giorni",
		jaJP: "%v日",
		koKR: "%v일",
		pl:   "%v dni",
		ptBR: "%v dias",
		ptPT: "%v dias",
		ru:   "%v дней",
		tr:   "%v gün",
		ukUA: "%v днів",
		uz:   "%v kun",
		zhCN: "%v 天"},
	MESSAGE_TEXT_INTEREST_PLEASE_SPECIFY_PERIOD: {
		arEG: "يرجى أيضًا تحديد فترة الفائدة، على سبيل المثال هل هي %v%% سنويًا، أو شهريًا، أو أسبوعيًا، أو لعدد معين من الأيام؟",
		de:   "Geben Sie bitte auch den Zinszeitraum an, z. B. handelt es sich um %v%% für ein Jahr, einen Monat, eine Woche oder eine bestimmte Anzahl von Tagen?",
		en:   "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		es:   "Por favor, especifique también el período de interés, por ejemplo, ¿es %v%% por año, mes, semana, alguna cantidad de días?",
		faIR: "لطفاً دوره زمانی بهره را نیز مشخص کنید، مثلاً آیا %v%% برای هر سال، ماه، هفته، یا چند روز است؟",
		fr:   "Veuillez également préciser la période d&#39;intérêt, par exemple s&#39;agit-il de %v%% par an, par mois, par semaine, par un certain nombre de jours ?",
		id:   "Harap tentukan juga periode bunga, misalnya apakah %v%% untuk per tahun, bulan, minggu, atau sejumlah hari tertentu?",
		it:   "Specificare anche il periodo di interesse, ad esempio %v%% per anno, mese, settimana, un certo numero di giorni?",
		jaJP: "利息期間も指定してください。たとえば、年、月、週、日数あたり %v%% などです。",
		koKR: "또한 이자 기간도 명시해 주세요. 예를 들어, 1년, 1개월, 1주 또는 며칠 단위로 %v%%할 수 있나요?",
		pl:   "Proszę również określić okres odsetkowy, np. czy jest to %v%% w skali roku, miesiąca, tygodnia, czy określonej liczby dni?",
		ptBR: "Por favor, especifique também o período de juros, por exemplo, %v%% por ano, mês, semana, algum número de dias?",
		ptPT: "Por favor, especifique também o período de juros, por exemplo, %v%% por ano, mês, semana, algum número de dias?",
		ru:   "Пожалуйста укажите также процентный период, т.е. уточните %%v%% это процент за какое количество дней?",
		tr:   "Lütfen faiz periyodunu da belirtin, örneğin yıllık, aylık, haftalık veya belirli gün sayısı için %v%% mi?",
		ukUA: "Також, будь ласка, уточніть період нарахування відсотків, наприклад, це %v%% за рік, місяць, тиждень, певну кількість днів?",
		uz:   "Iltimos, foiz muddatini ham belgilang, masalan, bu yil, oy, hafta, bir necha kun uchun %v%%mi?",
		zhCN: "还请指定利息期，例如每年、每月、每周或某些天数的 %v%%？"},
	MESSAGE_TEXT_INTEREST: {
		arEG: "<b>الفائدة</b> : %v%% لكل %v",
		de:   "<b>Zinsen</b> : %v%% pro %v",
		en:   "<b>Interest</b>: %v%% per %v",
		es:   "<b>Interés</b> : %v%% por %v",
		faIR: "<b>بهره</b> : %v%% به ازای هر %v",
		fr:   "<b>Intérêt</b> : %v%% par %v",
		id:   "<b>Bunga</b> : %v%% per %v",
		it:   "<b>Interesse</b> : %v%% per %v",
		jaJP: "<b>利息</b>: %v あたり %v%%",
		koKR: "<b>이자</b> : %v%% 당 %v",
		pl:   "<b>Odsetki</b> : %v%% na %v",
		ptBR: "<b>Juros</b> : %v%% por %v",
		ptPT: "<b>Juros</b> : %v%% por %v",
		ru:   "<b>Ставка</b>: %v%% за %v",
		tr:   "<b>Faiz</b> : %v%% başına %v",
		ukUA: "<b>Відсотки</b> : %v%% на %v",
		uz:   "<b>Foiz</b> : %v%% uchun %v",
		zhCN: "<b>利息</b>：每 %v %v%%"},
	MESSAGE_TEXT_INTEREST_MIN_PERIOD: {
		arEG: "الحد الأدنى للفترة %v",
		de:   "Mindestdauer %v",
		en:   "minimum period %v",
		es:   "período mínimo %v",
		faIR: "حداقل دوره %v",
		fr:   "période minimale %v",
		id:   "periode minimum %v",
		it:   "periodo minimo %v",
		jaJP: "最小期間 %v",
		koKR: "최소 기간 %v",
		pl:   "minimalny okres %v",
		ptBR: "período mínimo %v",
		ptPT: "período mínimo %v",
		ru:   "минимальный период %v",
		tr:   "minimum süre %v",
		ukUA: "мінімальний період %v",
		uz:   "minimal davr %v",
		zhCN: "最短周期 %v"},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		arEG: "{{.Counterparty}} استعار منك {{.Amount}}.",
		de:   "{{.Counterparty}} schuldet dir {{.Amount}} .",
		en:   "{{.Counterparty}} borrowed from you {{.Amount}}.",
		es:   "{{.Counterparty}} prestado por tí {{.Amount}}.",
		faIR: "{{.Counterparty}} از شما {{.Amount}} قرض گرفته است .",
		fr:   "{{.Counterparty}} vous a emprunté {{.Amount}}.",
		id:   "{{.Counterparty}} meminjam dari Anda {{.Amount}}.",
		it:   "{{.Counterparty}} e' in debito di {{.Amount}} con te.",
		jaJP: "{{.Counterparty}}はあなた{{.Amount}}から借りました。",
		koKR: "{{.Counterparty}}이(가) 귀하에게 {{.Amount}}을(를) 빌렸습니다.",
		pl:   "{{.Counterparty}} pożyczone od Ciebie {{.Amount}}.",
		ptBR: "{{.Counterparty}} emprestado de você {{.Amount}}.",
		ptPT: "{{.Counterparty}} emprestado de si {{.Amount}}.",
		ru:   "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		tr:   "{{.Counterparty}} sizden ödünç alındı {{.Amount}}.",
		ukUA: "{{.Counterparty}} позичив у вас {{.Amount}}.",
		uz:   "{{.Counterparty}} sizdan qarz oldi {{.Amount}}.",
		zhCN: "{{.Counterparty}}向你借了{{.Amount}}。"},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		arEG: "{{.Counterparty}} أقرضتك {{.Amount}}.",
		de:   "{{.Counterparty}} hat dir {{.Amount}} geliehen.",
		en:   "{{.Counterparty}} lended to you {{.Amount}}.",
		es:   "{{.Counterparty}} prestado a mí {{.Amount}}.",
		faIR: "{{.Counterparty}} به شما {{.Amount}} قرض داده است .",
		fr:   "{{.Counterparty}} vous a prêté {{.Amount}}.",
		id:   "{{.Counterparty}} dipinjamkan kepada Anda {{.Amount}}.",
		it:   "{{.Counterparty}} ti ha prestato {{.Amount}}.",
		jaJP: "{{.Counterparty}} があなたに {{.Amount}} を貸しました。",
		koKR: "{{.Counterparty}}이(가) {{.Amount}}에게 대여되었습니다.",
		pl:   "{{.Counterparty}} pożyczył Ci {{.Amount}}.",
		ptBR: "{{.Counterparty}} emprestado a você {{.Amount}}.",
		ptPT: "{{.Counterparty}} emprestado a si {{.Amount}}.",
		ru:   "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		tr:   "{{.Counterparty}} size ödünç verildi {{.Amount}}.",
		ukUA: "{{.Counterparty}} позичено вам {{.Amount}}.",
		uz:   "{{.Counterparty}} sizga {{.Amount}} qarz berdi.",
		zhCN: "{{.Counterparty}}借给您{{.Amount}}。"},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		arEG: "لقد قمت بإرجاع {{.Amount}} إلى {{.Counterparty}}.",
		de:   "Du hast {{.Amount}} an {{.Counterparty}} beglichen.",
		en:   "You returned {{.Amount}} to {{.Counterparty}}.",
		es:   "Has devuelto {{.Amount}} a {{.Counterparty}}.",
		faIR: "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
		fr:   "Vous avez renvoyé {{.Amount}} à {{.Counterparty}}.",
		id:   "Anda mengembalikan {{.Amount}} ke {{.Counterparty}}.",
		it:   "Hai ridato {{.Amount}} a {{.Counterparty}}.",
		jaJP: "{{.Amount}} を {{.Counterparty}} に返しました。",
		koKR: "{{.Amount}}을 {{.Counterparty}}로 반환했습니다.",
		pl:   "Wróciłeś {{.Amount}} do {{.Counterparty}}.",
		ptBR: "Você retornou {{.Amount}} para {{.Counterparty}}.",
		ptPT: "Regressou {{.Amount}} para {{.Counterparty}}.",
		ru:   "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		tr:   "{{.Amount}}&#39;ı {{.Counterparty}}&#39;e döndürdünüz.",
		ukUA: "Ви повернули {{.Amount}} користувачу {{.Counterparty}}.",
		uz:   "Siz {{.Amount}} ni {{.Counterparty}} ga qaytardingiz.",
		zhCN: "您已将 {{.Amount}} 返回至 {{.Counterparty}}。"},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		arEG: "{{.Counterparty}} عاد إليك {{.Amount}}.",
		de:   "{{.Counterparty}} hat dir {{.Amount}} beglichen.",
		en:   "{{.Counterparty}} returned to you {{.Amount}}.",
		es:   "{{.Counterparty}} te ha devuelto {{.Amount}}.",
		faIR: "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
		fr:   "{{.Counterparty}} vous est revenu {{.Amount}}.",
		id:   "{{.Counterparty}} dikembalikan kepada Anda {{.Amount}}.",
		it:   "{{.Counterparty}} ti ha ridato {{.Amount}}.",
		jaJP: "{{.Counterparty}} が {{.Amount}} に返されました。",
		koKR: "{{.Counterparty}}이(가) {{.Amount}}에게 반환되었습니다.",
		pl:   "{{.Counterparty}} zwrócono Ci {{.Amount}}.",
		ptBR: "{{.Counterparty}} retornou para você {{.Amount}}.",
		ptPT: "{{.Counterparty}} retornou para si {{.Amount}}.",
		ru:   "{{.Counterparty}} вернул вам {{.Amount}}.",
		tr:   "{{.Counterparty}} size {{.Amount}} geri döndü.",
		ukUA: "{{.Counterparty}} повернуто вам {{.Amount}}.",
		uz:   "{{.Counterparty}} sizga {{.Amount}} qaytardi.",
		zhCN: "{{.Counterparty}} 返回给您 {{.Amount}}。"},
	MESSAGE_TEXT_TRANSFER_ALREADY_FULLY_RETURNED: {
		arEG: "لقد تم سداد هذه الديون بالكامل بالفعل.",
		de:   "Diese Schuld ist bereits vollständig beglichen.",
		en:   "This debts is already fully returned.",
		es:   "Esta deuda se ha devuelta totalmente.",
		faIR: "این بدهی ها در حال حاضر به طور کامل بازگشته است.",
		fr:   "Cette dette a déjà été entièrement remboursée.",
		id:   "Hutang ini sudah dikembalikan sepenuhnya.",
		it:   "Questi debiti sono già completamente restituiti.",
		jaJP: "この借金はすでに全額返済済みです。",
		koKR: "이 빚은 이미 전액 갚았습니다.",
		pl:   "Ten dług został już w całości zwrócony.",
		ptBR: "Essas dívidas já foram totalmente quitadas.",
		ptPT: "Estas dívidas já foram totalmente liquidadas.",
		ru:   "Этот долг уже полностью возвращён.",
		tr:   "Bu borç zaten tamamen geri ödendi.",
		ukUA: "Ці борги вже повністю повернуті.",
		uz:   "Bu qarzlar allaqachon to&#39;liq qaytarilgan.",
		zhCN: "这笔债已经全部还清了。"},
	MESSAGE_TEXT_RECEIPT_ALREADY_RETURNED_AMOUNT: {
		arEG: "تم إرجاعه بالفعل: {{.Amount}}.",
		de:   "Bereits beglichen: {{.Amount}}.",
		en:   "Already returned: {{.Amount}}.",
		es:   "Se ha devuelto ya: {{.Amount}}.",
		faIR: "قبلا برگشت: {{.Amount}}.",
		fr:   "Déjà retourné : {{.Amount}}.",
		id:   "Sudah dikembalikan: {{.Amount}}.",
		it:   "Già restituito: {{.Amount}}.",
		jaJP: "すでに返されています: {{.Amount}}。",
		koKR: "이미 반환됨: {{.Amount}}.",
		pl:   "Już zwrócono: {{.Amount}}.",
		ptBR: "Já retornou: {{.Amount}}.",
		ptPT: "Já regressou: {{.Amount}}.",
		ru:   "Уже возвращено: {{.Amount}}.",
		tr:   "Zaten döndürüldü: {{.Amount}}.",
		ukUA: "Вже повернуто: {{.Amount}}.",
		uz:   "Allaqachon qaytarilgan: {{.Amount}}.",
		zhCN: "已返回：{{.Amount}}。"},
	MESSAGE_TEXT_RECEIPT_OUTSTANDING_AMOUNT: {
		arEG: "متميز: {{.Amount}}.",
		de:   "Ausstehend: {{.Amount}}.",
		en:   "Outstanding: {{.Amount}}.",
		es:   "Falta devolver: {{.Amount}}.",
		faIR: "برجسته: {{.Amount}}",
		fr:   "Exceptionnel : {{.Amount}}.",
		id:   "Luar biasa: {{.Amount}}.",
		it:   "Inevaso: {{.Amount}}",
		jaJP: "未処理: {{.Amount}}。",
		koKR: "미해결: {{.Amount}}.",
		pl:   "Wybitne: {{.Amount}}.",
		ptBR: "Excelente: {{.Amount}}.",
		ptPT: "Excelente: {{.Amount}}.",
		ru:   "Осталось вернуть: {{.Amount}}.",
		tr:   "Üstün: {{.Amount}}.",
		ukUA: "Видатне: {{.Amount}}.",
		uz:   "Eng yaxshi: {{.Amount}}.",
		zhCN: "未完成：{{.Amount}}。"},
	MESSAGE_TEXT_DUE_ON: {
		arEG: "<b>العودة إلى الصندوق</b> : %v",
		de:   "<b>Fällig am</b>: %v",
		en:   "<b>Return till</b>: %v",
		es:   "<b>Devolver hasta</b>: %v",
		faIR: "<b>بازگردانده شود تا</b>: %v",
		fr:   "<b>Retour jusqu&#39;à</b> : %v",
		id:   "<b>Kembali sampai</b> : %v",
		it:   "<b>Scadenza</b>: %v",
		jaJP: "<b>返却期限</b>: %v",
		koKR: "<b>반환할 곳</b> : %v",
		pl:   "<b>Powrót do</b> : %v",
		ptBR: "<b>Retornar até</b> : %v",
		ptPT: "<b>Retornar até</b> : %v",
		ru:   "<b>Вернуть до</b>: %v",
		tr:   "<b>İade tarihi</b> : %v",
		ukUA: "<b>Повернути до</b> : %v",
		uz:   "<b>Qaytish</b> : %v",
		zhCN: "<b>返回至</b>：%v"},
	MESSAGE_TEXT_NOTE: {
		arEG: "ملحوظة",
		de:   "Notiz",
		en:   "Note",
		es:   "Nota",
		faIR: "نکته",
		fr:   "Note",
		id:   "Catatan",
		it:   "Nota",
		jaJP: "注記",
		koKR: "메모",
		pl:   "Notatka",
		ptBR: "Observação",
		ptPT: "Nota",
		ru:   "Заметка",
		tr:   "Not",
		ukUA: "Примітка",
		uz:   "Eslatma",
		zhCN: "笔记"},
	MESSAGE_TEXT_COMMENT: {
		arEG: "تعليق",
		de:   "Bemerkung",
		en:   "Comment",
		es:   "Comentario",
		faIR: "شرح",
		fr:   "Commentaire",
		id:   "Komentar",
		it:   "Commenti",
		jaJP: "コメント",
		koKR: "논평",
		pl:   "Komentarz",
		ptBR: "Comentário",
		ptPT: "Comentar",
		ru:   "Комментарий",
		tr:   "Yorum",
		ukUA: "Коментар",
		uz:   "Izoh",
		zhCN: "评论"},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		arEG: `انقر هنا لتسجيل <a>الدخول</a> إلى تطبيق الويب.`,
		de:   `<a>Hier klicken</a>, um sich an der Wep-App anzumelden.`,
		en:   `Click to <a>sign in</a> to web-app.`,
		es:   `Haz click para <a>acceder</a>la web-app.`,
		faIR: `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
		fr:   `Cliquez pour <a>vous connecter</a> à l&#39;application Web.`,
		id:   `Klik untuk <a>masuk</a> ke aplikasi web.`,
		it:   "Fai clic per <a>accedi</a> per app web.",
		jaJP: `クリックしてWebアプリに<a>サインインします</a>。`,
		koKR: `웹앱에 <a>로그인</a> 하려면 클릭하세요.`,
		pl:   `Kliknij, aby <a>zalogować się</a> do aplikacji internetowej.`,
		ptBR: `Clique para fazer <a>login</a> no aplicativo da web.`,
		ptPT: `Clique para iniciar <a>sessão</a> na aplicação web.`,
		ru:   `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
		tr:   `Web uygulamasına <a>giriş yapmak</a> için tıklayın.`,
		ukUA: `Натисніть, щоб <a>увійти</a> у веб-додаток.`,
		uz:   `Veb-ilovaga <a>kirish</a> uchun bosing.`,
		zhCN: `单击即可<a>登录</a>Web 应用程序。`},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		arEG: "هل يعجبك @{{bot}}؟",
		de:   "Magst du @{{bot}}?",
		en:   "Do you like @{{bot}}?",
		es:   "¿Te gusta @{{bot}}?",
		faIR: "آیا می پسندید @{{bot}}?",
		fr:   "Aimez-vous @{{bot}}?",
		id:   "Apakah kamu suka @{{bot}}?",
		it:   "Divertito con @{{bot}}?",
		jaJP: "@{{bot}}が好きですか？",
		koKR: "@{{bot}}을 좋아하시나요?",
		pl:   "Czy lubisz @{{bot}}?",
		ptBR: "Você gosta de @{{bot}}?",
		ptPT: "Gosta de @{{bot}}?",
		ru:   "Вам нравится @{{bot}}?",
		tr:   "@{{bot}} hoşunuza gidiyor mu?",
		ukUA: "Тобі подобається @{{bot}}?",
		uz:   "Sizga @{{bot}} yoqdimi?",
		zhCN: "你喜欢@{{bot}}吗？"},
	COMMAND_TEXT_YES_EXCLAMATION: {
		arEG: "%v نعم!",
		de:   "%v Ja!",
		en:   "%v Yes!",
		es:   "%v ¡Sí!",
		faIR: "بله! %v",
		fr:   "%v Oui !",
		id:   "%v Ya!",
		it:   "%v Si!",
		jaJP: "%v はい!",
		koKR: "%v 네!",
		pl:   "%v Tak!",
		ptBR: "%v Sim!",
		ptPT: "%v Sim!",
		ru:   "%v Да!",
		tr:   "%v Evet!",
		ukUA: "%v Так!",
		uz:   "%v Ha!",
		zhCN: "%v 是的！"},
	COMMAND_TEXT_YES: {
		arEG: "%v نعم",
		de:   "%v Ja",
		en:   "%v Yes",
		es:   "%v Sí",
		faIR: "بله %v",
		fr:   "%v Oui",
		id:   "%v Ya",
		it:   "%v Si",
		jaJP: "%v はい",
		koKR: "%v 예",
		pl:   "%v Tak",
		ptBR: "%v Sim",
		ptPT: "%v Sim",
		ru:   "%v Да",
		tr:   "%v Evet",
		ukUA: "%v Так",
		uz:   "%v Ha",
		zhCN: "%v 是"},
	COMMAND_TEXT_NO: {
		arEG: "%v لا",
		de:   "%v Nein",
		en:   "%v No",
		es:   "%v No",
		faIR: "خیر %v",
		fr:   "%v Non",
		id:   "%v Tidak",
		it:   "%v No",
		jaJP: "%v いいえ",
		koKR: "%v 아니요",
		pl:   "%v Nie",
		ptBR: "%v Não",
		ptPT: "%v Não",
		ru:   "%v Нет",
		tr:   "%v Hayır",
		ukUA: "%v Ні",
		uz:   "%v Yo‘q",
		zhCN: "%v 没有"},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		arEG: "%v ليس كثيرا",
		de:   "%v Nicht so sehr",
		en:   "%v Not too much",
		es:   "%vNo mucho",
		faIR: "نه خیلی زیاد %v",
		fr:   "%v Pas trop",
		id:   "%v Tidak terlalu banyak",
		it:   "%v Non troppo",
		jaJP: "%v あまり多くない",
		koKR: "%v 너무 많지 않아요",
		pl:   "%v Nie za dużo",
		ptBR: "%v Não muito",
		ptPT: "%v Não muito",
		ru:   "%v Не очень",
		tr:   "%v Çok fazla değil",
		ukUA: "%v Не забагато",
		uz:   "%v Ortiqcha emas",
		zhCN: "%v 不算太多"},
	COMMAND_TEXT_FEEDBACK: {
		arEG: "/تعليق",
		de:   "/Bewertung",
		en:   "/Feedback",
		es:   "/Respuesta",
		faIR: "/بازخورد",
		fr:   "/Retour",
		id:   "/Masukan",
		it:   "/Risposta",
		jaJP: "/フィードバック",
		koKR: "/피드백",
		pl:   "/Informacja zwrotna",
		ptBR: "/Opinião",
		ptPT: "/Feedback",
		ru:   "/Отзыв",
		tr:   "/Geri bildirim",
		ukUA: "/Відгук",
		uz:   "/Mulohaza",
		zhCN: "/反馈"},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		arEG: "%v كتابة التعليقات",
		de:   "%v Bewertung schreiben",
		en:   "%v Write feedback",
		es:   "%v Escribir un comentario",
		faIR: "ارسال بازخورد %v",
		fr:   "%v Écrire un commentaire",
		id:   "%v Tulis umpan balik",
		it:   "%v Scrivi commenti",
		jaJP: "%v フィードバックを書く",
		koKR: "%v 피드백 쓰기",
		pl:   "%v Napisz opinię",
		ptBR: "%v Escrever feedback",
		ptPT: "%v Escrever feedback",
		ru:   "%v Написать отзыв",
		tr:   "%v Geri bildirim yaz",
		ukUA: "%v Написати відгук",
		uz:   "%v Fikr-mulohaza yozing",
		zhCN: "%v 写反馈"},
	MESSAGE_TEXT_THANKS: {
		arEG: "🙏 شكرا!",
		de:   "🙏 Danke!",
		en:   "🙏 Thanks!",
		es:   "🙏 ¡Gracias!",
		faIR: "🙏 تشکر!",
		fr:   "🙏 Merci!",
		id:   "🙏 Terima kasih!",
		it:   "🙏 Grazie!",
		jaJP: "🙏 ありがとう!",
		koKR: "🙏 감사합니다!",
		pl:   "🙏 Dziękuję!",
		ptBR: "🙏 Obrigado!",
		ptPT: "🙏 Obrigado!",
		ru:   "🙏 Спасибо!",
		tr:   "🙏 Teşekkürler!",
		ukUA: "🙏 Дякую!",
		uz:   "🙏 Rahmat!",
		zhCN: "🙏 谢谢!",
	},
	MESSGE_TEXT_DEBT_ERROR_FIXED_START_OVER: {
		arEG: "🙏 عذرًا، حدث خطأ. تم إصلاحه، ولكن يُرجى إعادة إدخال بياناتك المتعلقة بهذا الدين.",
		de:   "🙏 Entschuldigung, da gab es einen Fehler. Er wird bald behoben, aber du musst nochmal neu anfangen.",
		en:   "🙏 Sorry, there was an error. It has been fixed but please re-enter your data for this debt.",
		es:   "🙏 Lo siento, ha salido un error. Lo ha arreglado, pero para esta deuda hay que introducir los datos de nuevo. ",
		faIR: "🙏 متاسفیم، خطایی رخ داده بود. مشکل برطرف شده است، اما لطفاً اطلاعات خود را برای این بدهی دوباره وارد کنید.",
		fr:   "Désolé, une erreur s&#39;est produite. Elle a été corrigée, mais veuillez ressaisir vos données pour cette dette.",
		id:   "🙏 Maaf, terjadi kesalahan. Kesalahan telah diperbaiki, tetapi harap masukkan kembali data Anda untuk utang ini.",
		it:   "🙏 Spiacenti, si è verificato un errore. È stato risolto, ma ti preghiamo di reinserire i dati per questo debito.",
		jaJP: "🙏 申し訳ございません。エラーが発生しました。修正いたしましたが、この債務に関するデータを再度ご入力ください。",
		koKR: "🙏 죄송합니다. 오류가 발생했습니다. 수정되었지만 해당 부채에 대한 데이터를 다시 입력해 주세요.",
		pl:   "🙏 Przepraszamy, wystąpił błąd. Został naprawiony, ale proszę ponownie wprowadzić swoje dane dotyczące tego długu.",
		ptBR: "🙏 Desculpe, ocorreu um erro. Ele já foi corrigido, mas, por favor, insira novamente os dados desta dívida.",
		ptPT: "🙏 Desculpe, ocorreu um erro. Já foi corrigido, mas, por favor, insira novamente os dados desta dívida.",
		ru:   "🙏 Извините, у нас была ошибка. Она была исправлено, но потребуется внести данные для этого долга заново.",
		tr:   "🙏 Üzgünüz, bir hata oluştu. Düzeltildi ancak lütfen bu borç için verilerinizi tekrar girin.",
		ukUA: "🙏 Вибачте, сталася помилка. Її виправлено, але, будь ласка, повторно введіть свої дані для цього боргу.",
		uz:   "🙏 Kechirasiz, xatolik yuz berdi. Bu tuzatildi, ammo bu qarz uchun maʼlumotlarni qayta kiriting.",
		zhCN: "🙏 抱歉，出现错误。问题已修复，但请重新输入此债务的数据。"},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		arEG: "يرجى إرسال النص.",
		de:   "Bitte senden sie einen Text.",
		en:   "Please send text.",
		es:   "Por favor, envia el texto.",
		faIR: "لطفاً متن ارسال کنید.",
		fr:   "S&#39;il vous plaît envoyer un texte.",
		id:   "Silakan kirim teks.",
		it:   "Si prega di inviare il testo.",
		jaJP: "テキストを送信してください。",
		koKR: "문자 메시지를 보내주세요.",
		pl:   "Proszę wysłać SMS.",
		ptBR: "Por favor, envie uma mensagem de texto.",
		ptPT: "Por favor, envie uma mensagem de texto.",
		ru:   "Пожалуйста отправьте текст.",
		tr:   "Lütfen mesaj gönderin.",
		ukUA: "Будь ласка, надішліть текст.",
		uz:   "Iltimos, matnni yuboring.",
		zhCN: "请发送文本。"},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {
		arEG: `🤖 هل يمكنك تقييمه بدرجة عالية وكتابة مراجعة جيدة في متجر بوتات كتالوج؟ {ن} سيستغرق الأمر أقل من دقيقة من وقتك! 😇`,
		de: `🤖 Kannst du mich im Store Bot hoch bewerten und eine gute Bewertung schreiben?
		Es wird dich weniger als eine Minute kosten! 😇`,
		en: `🤖 Can you rate it high and write a good review in bots catalog Store Bot?
		It will take less than a minute of your time! 😇`,
		es: `🤖 Puedes valolarlo con una buena nota y una buena opinión en el catálogo Store Bot?
		Te costará menos de un minuto! 😇`,
		faIR: `🤖  آیا می توانید در کاتالوگ روباتها در استور بوت امتیاز بالایی داده و اظهار نظر خوبی در مورد این روبات ثبت کنید؟
		این کار کمتر از یک دقیقه از وقت شما را می گیرد! 😇`,
		fr: `🤖 Pouvez-vous lui donner une note élevée et écrire une bonne critique dans le catalogue des bots Store Bot ? Cela vous prendra moins d&#39;une minute de votre temps ! 😇`,
		id: `🤖 Bisakah Anda memberi peringkat tinggi dan menulis ulasan bagus di katalog bot Store Bot?
 Hanya butuh waktu kurang dari satu menit! 😇`,
		it: `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo Bot Store?
		Ci vorrà meno di un minuto del tuo tempo! 😇`,
		jaJP: `🤖 ボットカタログストアボットに高評価と良いレビューを書いていただけますか？
 1 分もかかりません！😇`,
		koKR: `🤖 봇 카탈로그 스토어 봇에 높은 평점을 주고 좋은 리뷰를 써주실 수 있나요?
 1분도 안 걸릴 거예요! 😇`,
		pl: `🤖 Czy możesz ocenić go wysoko i napisać dobrą recenzję w katalogu botów Store Bot? Zajmie to mniej niż minutę Twojego czasu! 😇`,
		ptBR: `🤖 Você pode dar uma nota alta e escrever uma boa avaliação no catálogo de bots da Store Bot?
 Levará menos de um minuto do seu tempo! 😇`,
		ptPT: `🤖 Consegue dar uma nota alta e escrever uma boa avaliação no catálogo de bots da Store Bot?
 Vai demorar menos de um minuto do seu tempo! 😇`,
		ru: `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?
		Это займет меньше минуты вашего времени! 😇`,
		tr: `🤖 Bot kataloğunda yüksek puan verip iyi bir yorum yazabilir misiniz? Mağaza Botu?
 Bir dakikadan az zamanınızı alacaktır! 😇`,
		ukUA: `🤖 Чи можете ви оцінити це високо та написати хороший відгук у каталозі ботів Store Bot?
 Це займе менше хвилини вашого часу! 😇`,
		uz: `🤖 Uni yuqori baholab, botlar katalogiga yaxshi sharh yoza olasizmi?
 Bu bir daqiqadan kam vaqtingizni oladi! 😇`,
		zhCN: `🤖 您能给它高评价并在机器人目录 Store Bot 中写一篇好评吗？
 这只会占用您不到一分钟的时间！😇`},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		arEG: "شارك بأفكارك (باللغة الإنجليزية أو الروسية) حول ما يمكن القيام به لتحسين الروبوت:",
		de:   "Schreibe (auf Englisch oder Russisch) uns, was man am Bot besser machen kann:",
		en:   "Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		es:   "Comparte tus pensamientos (en Inglés o Ruso) sobre qué podemos hacer para que el bot sea mejor:",
		faIR: "نظرات خود را (به انگلیسی و روسی ) در مورد اینکه چه کاری می توان انجام داد تا این ربات بهتر شود، با ما به اشتراک بگذارید:",
		fr:   "Partagez vos réflexions (en anglais ou en russe) sur ce qui pourrait être fait pour améliorer le bot :",
		id:   "Bagikan pemikiran Anda (dalam bahasa Inggris atau Rusia) tentang apa yang dapat dilakukan untuk membuat bot lebih baik:",
		it:   "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
		jaJP: "ボットをもっと良くするために何ができるかについて、あなたの考えを（英語またはロシア語で）共有してください。",
		koKR: "봇을 개선하기 위해 무엇을 할 수 있을지에 대한 생각을 영어 또는 러시아어로 공유해 주세요.",
		pl:   "Podziel się swoimi przemyśleniami (w języku angielskim lub rosyjskim) na temat tego, co można zrobić, aby ulepszyć bota:",
		ptBR: "Compartilhe suas ideias (em inglês ou russo) sobre o que pode ser feito para melhorar o bot:",
		ptPT: "Partilhe as suas ideias (em inglês ou russo) sobre o que pode ser feito para melhorar o bot:",
		ru:   "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
		tr:   "Botu daha iyi hale getirmek için neler yapılabileceğine dair düşüncelerinizi paylaşın (İngilizce veya Rusça):",
		ukUA: "Поділіться своїми думками (англійською або російською мовою) про те, що можна було б зробити, щоб покращити роботу бота:",
		uz:   "Botni yaxshilash uchun nima qilish mumkinligi haqida fikringizni (ingliz yoki rus tillarida) o&#39;rtoqlashing:",
		zhCN: "请用英语或俄语分享您关于如何改进机器人的想法："},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {
		arEG: `<b>كيفية التقييم بثلاث خطوات بسيطة:</b> 

 1. انقر على هذا الرابط للتقييم والمراجعة: 
 https://t.me/storebot?start={{bot}}

 2. انقر على زر &quot;⭐️⭐️⭐️⭐️⭐️&quot;

 3. اكتب رسالتك أو اضغط على زر &quot;تخطي هذه الخطوة&quot;

 شكرًا جزيلاً لك! بفضل جهودك، سيتعرف المزيد من الأشخاص على البوت. كل هذا سيكون حافزًا إضافيًا للمطورين! 😎`,
		de: `<b>Wie man bewertet - in 3 einfachen Schritten:</b>

	1. Klick auf diesen Link, um eine Bewertung abzugeben:
	https://t.me/storebot?start={{bot}}

	2. Wähle "⭐️⭐️⭐️⭐️⭐️"

	3. Schreib etwas Nettes auf Englisch oder wähle "Skip this step"

	Wirklich vielen Dank! Dank deiner Bewertung werden vielleicht mehr Leute auf diesen Bot aufmerksam. Das ist gut für die Motivation der Entwickler dieses Bots! 😎`,
		en: `<b>How to rate in 3 simple steps:</b>

	1. Click on this link to rate and review:
	https://t.me/storebot?start={{bot}}

	2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

	3. Write your message or press "Skip this step" button

	Thank you very much! As a result of your actions, even more people will learn about the bot.All this will serve as the additional motivation for the developers! 😎`,
		es: `<b>Como valorar en 3 simples pasos:</b>

	1. Click este link para valorar y dejar tu opinión:
	https://t.me/storebot?start={{bot}}

	2. Click en "⭐️⭐️⭐️⭐️⭐️" botón

	3. Escribe tu mensage o apreta "Skip this step" botón

	¡Muchas gracias! Merced a tus acciones más gente conocerá a bot. Todo eso sirve para una motivación adicional a los creadores! 😎`,
		faIR: `<b>چگونگی امتیازدهی در سه گام ساده :</b>

	1. برای امتیازدهی و ثبت نظرات بر روی لینگ زیر کلیک کنید
	https://t.me/storebot?start={{bot}}

	2. بر روی دکمه "⭐️⭐️⭐️⭐️⭐️" کلیک کنید

	3. پیام خودرا ثبت کنید یا روی دکمه "پرش از این مرحله" کلیک کنید

	بسیار سپاسگزاریم! عمل شما باعث می شود افراد بیشتری در مورد bot.All بیاموزند. این امر انگیزه مضاعفی به توسعه دهندگان این ربات می دهد ! 😎`,
		fr: `<b>Comment évaluer en 3 étapes simples :</b> 

 1. Cliquez sur ce lien pour évaluer et donner votre avis : 
 https://t.me/storebot?start={{bot}}

 2. Cliquez sur le bouton « ⭐️⭐️⭐️⭐️⭐️ » 

 3. Rédigez votre message ou cliquez sur le bouton « Passer cette étape » 

 Merci beaucoup ! Grâce à vos actions, davantage de personnes découvriront le bot. Tout cela servira de motivation supplémentaire aux développeurs ! 😎`,
		id: `<b>Cara memberi penilaian dalam 3 langkah mudah:</b> 

 1. Klik tautan ini untuk memberi penilaian dan ulasan:
 https://t.me/storebot?start={{bot}}

 2. Klik tombol &quot;⭐️⭐️⭐️⭐️⭐️&quot;

 3. Tulis pesan Anda atau tekan tombol &quot;Lewati langkah ini&quot;

 Terima kasih banyak! Sebagai hasil dari tindakan Anda, semakin banyak orang akan mengetahui tentang bot ini. Semua ini akan menjadi motivasi tambahan bagi para pengembang! 😎`,
		it: `<b>Come valutare in 3 semplici passaggi:</b>
	1. Clicca su questo link per votare e lasciare una recensione:
	https://t.me/storebot?start={{bot}}

	2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

	3. Scrivi il tuo messaggio o premi "Salta questo step"

	Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot.Dando anche un motivo in più per continuare ai developers! 😎`,
		jaJP: `<b>評価方法は3つの簡単なステップで完了します:</b> 

 1. このリンクをクリックして評価とレビューを行ってください:
 https://t.me/storebot?start={{bot}}

 2. 「⭐️⭐️⭐️⭐️⭐️」ボタンをクリックします

 3. メッセージを入力するか、「この手順をスキップ」ボタンをクリックします

 ありがとうございます！皆様のご支援のおかげで、より多くの人がボットについて知るようになります。これは開発者にとって大きなモチベーションとなります！😎`,
		koKR: `<b>3단계로 간편하게 평가하는 방법:</b> 

 1. 이 링크를 클릭하여 평가하고 리뷰를 남겨주세요:
 https://t.me/storebot?start={{bot}}

 2. &quot;⭐️⭐️⭐️⭐️⭐️&quot; 버튼을 클릭하세요.

 3. 메시지를 작성하거나 &quot;이 단계 건너뛰기&quot; 버튼을 클릭하세요.

 감사합니다! 여러분의 참여 덕분에 더 많은 사람들이 봇에 대해 알게 될 것입니다. 이 모든 것이 개발자들에게 큰 동기 부여가 될 것입니다! 😎`,
		pl: `<b>Jak oceniać w 3 prostych krokach:</b> 

 1. Kliknij ten link, aby ocenić i napisać recenzję:
 https://t.me/storebot?start={{bot}}

 2. Kliknij przycisk „⭐️⭐️⭐️⭐️⭐️”

 3. Napisz wiadomość lub naciśnij przycisk „Pomiń ten krok”

 Dziękuję bardzo! Dzięki Twoim działaniom jeszcze więcej osób dowie się o bocie. Wszystko to będzie dodatkową motywacją dla programistów! 😎`,
		ptBR: `<b>Como avaliar em 3 passos simples:</b> 

 1. Clique neste link para avaliar e comentar:
 https://t.me/storebot?start={{bot}}

 2. Clique no botão &quot;⭐️⭐️⭐️⭐️⭐️&quot;

 3. Escreva sua mensagem ou clique no botão &quot;Pular esta etapa&quot;

 Muito obrigado! Como resultado de suas ações, ainda mais pessoas aprenderão sobre o bot. Tudo isso servirá como motivação adicional para os desenvolvedores! 😎`,
		ptPT: `<b>Como avaliar em 3 passos simples:</b> 

 1. Clique neste link para avaliar e comentar:
 https://t.me/storebot?start={{bot}}

 2. Clique no botão &quot;⭐️⭐️⭐️⭐️⭐️&quot;

 3. Escreva a sua mensagem ou clique no botão &quot;Saltar este passo&quot;

 Muito obrigado! Como resultado das suas ações, ainda mais pessoas aprenderão sobre o bot. Tudo isto servirá como motivação adicional para os desenvolvedores! 😎`,
		ru: `<b>Как поставить оценку в три простых шага:</b>

	1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
	https://t.me/storebot?start={{bot}}

	2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

	3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

	Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,
		tr: `<b>3 basit adımda nasıl puan verilir:</b> 

 1. Puan vermek ve yorum yapmak için bu bağlantıya tıklayın:
 https://t.me/storebot?start={{bot}}

 2. &quot;⭐️⭐️⭐️⭐️⭐️&quot; düğmesine tıklayın

 3. Mesajınızı yazın veya &quot;Bu adımı atla&quot; düğmesine basın

 Çok teşekkür ederim! Eylemlerinizin sonucunda daha fazla kişi botu öğrenecek. Tüm bunlar geliştiriciler için ek motivasyon görevi görecek! 😎`,
		ukUA: `<b>Як оцінити за 3 простих кроки:</b> 

 1. Перейдіть за цим посиланням, щоб оцінити та залишити відгук:
 https://t.me/storebot?start={{bot}}

 2. Натисніть кнопку &quot;⭐️⭐️⭐️⭐️⭐️&quot;

 3. Напишіть своє повідомлення або натисніть кнопку &quot;Пропустити цей крок&quot;

 Щиро дякую! Завдяки вашим діям ще більше людей дізнаються про бота. Все це послужить додатковою мотивацією для розробників! 😎`,
		uz: `<b>Qanday qilib 3 oddiy qadamda baho berish mumkin:</b> 

 1. Baholash va koʻrib chiqish uchun ushbu havolani bosing:
 https://t.me/storebot?start={{bot}}

 2. &quot;⭐️⭐️⭐️⭐️⭐️&quot; tugmasini bosing

 3. Xabaringizni yozing yoki &quot;Katta rahmat&quot; 
 tugmasini bosing! Sizning harakatlaringiz natijasida ko&#39;proq odamlar bot haqida bilib oladi. Bularning barchasi ishlab chiquvchilar uchun qo&#39;shimcha motivatsiya bo&#39;lib xizmat qiladi! 😎`,
		zhCN: `<b>只需 3 个简单步骤即可评分：</b> 

 1. 点击此链接进行评分和评论：
 https://t.me/storebot?start={{bot}}

 2. 点击“⭐️⭐️⭐️⭐️⭐️”按钮

 3. 写下您的留言或点击“跳过此步骤”按钮

 非常感谢！您的行动将让更多人了解这款机器人。所有这些都将成为开发者的额外动力！😎`},
	MESSAGE_TEXT_ASK_FOR_FEEDBACK: {
		arEG: "سنكون شاكرين لو أخبرتنا برأيك. الأمر لا يستغرق سوى بضع ثوانٍ.",
		de:   "Über ein kleines Feedback wie der Bot so ist, würden wir uns freuen. Es dauert nur ein paar Sekunden.",
		en:   "We would appreciate if tell us how we doing. It takes just few seconds.",
		es:   "Te agredecemos si valoras el funccionamiento de nuestro applicación. Te costará solo unos segundos.",
		faIR: "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
		fr:   "Nous apprécierions si vous nous disiez comment nous allons. Cela ne prend que quelques secondes.",
		id:   "Kami akan menghargai jika memberi tahu kami bagaimana kami melakukannya. Ini hanya membutuhkan beberapa detik.",
		it:   "Ci farebbe piacere se lasciassi un voto per il nostro lavoro. Ti bastano solo alcuni secondi.",
		jaJP: "私たちの仕事ぶりを教えていただければ幸いです。ほんの数秒で済みます。",
		koKR: "우리가 어떻게 하고 있는지 알려주시면 감사하겠습니다. 몇 초 밖에 걸리지 않습니다.",
		pl:   "Bylibyśmy wdzięczni, gdybyś powiedział nam, jak sobie radzimy. To zajmie tylko kilka sekund.",
		ptBR: "Agradecemos se nos disser como estamos indo. Leva apenas alguns segundos.",
		ptPT: "Agradecíamos que nos contasse como estamos a fazer. Demora apenas alguns segundos.",
		ru:   "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		tr:   "Nasıl yaptığımızı bize söylerseniz memnun oluruz. Sadece birkaç saniye sürer.",
		ukUA: "Ми будемо вдячні, якщо ви розкажете нам, як ми працюємо. Це займе лише кілька секунд.",
		uz:   "Qanday qilayotganimizni bizga aytsangiz, minnatdor bo'lardik. Bu atigi bir necha soniya vaqt oladi.",
		zhCN: "如果您告诉我们我们的表现如何，我们将不胜感激。这只需要几秒钟。",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		arEG: "قم بتقييم هذا البوت",
		de:   "Bewerte diesen Bot",
		en:   "Rate this bot",
		es:   "Valora a bot",
		faIR: "به این ربات امتیاز بدهید",
		fr:   "Évaluer ce bot",
		id:   "Nilai bot ini",
		it:   "Vota questo bot",
		jaJP: "このボットを評価する",
		koKR: "이 봇 평가하기",
		pl:   "Oceń tego bota",
		ptBR: "Avalie este bot",
		ptPT: "Avalie este bot",
		ru:   "Оценить приложение",
		tr:   "Bu botu değerlendir",
		ukUA: "Оцінити цього бота",
		uz:   "Bu botni baholang",
		zhCN: "评价这个机器人",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		arEG: "اترك التقييم على @Storebot",
		de:   "Mache eine Bewertung auf @Storebot",
		en:   "Leave rating at @Storebot",
		es:   "Valorar en @Storebot",
		faIR: "امتیاز خود را اینجا وارد کنید @Storebot",
		fr:   "Laisser une évaluation sur @Storebot",
		id:   "Tinggalkan peringkat di @Storebot",
		it:   "Lascia un voto a @Storebot",
		jaJP: "@Storebotで評価を残す",
		koKR: "@Storebot에서 평가 남기기",
		pl:   "Zostaw ocenę na @Storebot",
		ptBR: "Deixe sua avaliação no @Storebot",
		ptPT: "Deixe uma avaliação no @Storebot",
		ru:   "Оценить на  @Storebot",
		tr:   "@Storebot'ta değerlendirme bırakın",
		ukUA: "Залиште оцінку на @Storebot",
		uz:   "@Storebot'da baho qoldiring",
		zhCN: "在 @Storebot 上留下评分",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		arEG: `حسنًا، ربما يمكنك تقييمنا في وقت آخر.

 {{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

 سنكون ممتنين أيضًا إذا اقترحت أي تحسينات.
	`,
		de: `Okay, vielleicht werden wir wann anders bewertet.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ansonsten freuen wir uns immer zu hören, was man besser machen kann.
	`,
		en: `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		es: `OK, Quizás puedas valorar en otro momento.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	También te agredeceríamos si ofrecieras alguna mejora.
	`,
		faIR: `بسیار خوب، ممکن است شما بتوانید زمان دیگری به ما امتیاز بدهید.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	همچنین سپاسگزار خواهیم بود اگر شما هرگونه امکان بهبود را با ما در میان بگذارید.
	`,
		fr: `OK, peut-être que vous pourrez nous évaluer une autre fois.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous suggérez des améliorations.
	`,
		id: `OK, mungkin Anda dapat menilai kami lain kali.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda menyarankan perbaikan apa pun.
	`,
		it: `OK, forse ci puoi valutare un'altra volta.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Apprezzeremo anche se suggeriamo qualche miglioramento.
	`,
		jaJP: `OK、また今度評価してください。

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	改善点を提案していただければ幸いです。
	`,
		koKR: `네, 다음에 평가해 주실 수 있을 것입니다.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	개선 사항을 제안해 주시면 감사하겠습니다.
	`,
		pl: `OK, może ocenisz nas innym razem.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli zaproponujesz jakieś ulepszenia.
	`,
		ptBR: `OK, talvez você possa nos avaliar em outra ocasião.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você sugerir melhorias.
	`,
		ptPT: `OK, talvez nos possa avaliar noutra ocasião.

 {{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

 Também agradecemos se sugerir alguma melhoria.
	`,
		ru: `ОК, возможно вы смоежете поставить оценку в другой раз.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы предложите любые улучшения.
	`,
		tr: `Tamam, belki başka bir zaman bizi değerlendirebilirsiniz.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Herhangi bir iyileştirme önerirseniz de memnun oluruz.
	`,
		ukUA: `Гаразд, можливо, ви зможете оцінити нас іншим разом.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви запропонуєте будь-які покращення.
	`,
		uz: `OK, balki boshqa safar bizni baholashingiz mumkin.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz biron bir yaxshilanishni taklif qilsangiz, biz ham minnatdor bo'lamiz.
	`,
		zhCN: `好的，也许您可以在另一个时间给我们评分。

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您提出任何改进建议，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		arEG: `شكرًا لك، لقد عملنا بجد!

 {{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

 سنكون ممتنين أيضًا إذا <a suggest-idea>اقترحت تحسينات</a> .
	`,
		de: `Danke, wir arbeiten hart dran!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Wir freuen uns auch immer über <a suggest-idea>neue Ideen</a>.
	`,
		en: `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		es: `Gracias, hemos trabajado duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Te agradeceríamos si <a suggest-idea>propusieras alguna mejora</a>.
	`,
		faIR: `ممنونیم، ما سخت کارکرده ایم!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	ما همچنین سپاسگزار خواهیم بود اگر شما<a suggest-idea> هرگونه امکان بهبود را با ما در میان بگذارید </a>را.
	`,
		fr: `Merci, nous avons travaillé dur !

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous <a suggest-idea>suggérez des améliorations</a>.
	`,
		id: `Terima kasih, kami bekerja keras!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda <a suggest-idea>menyarankan perbaikan</a>.
	`,
		it: `GRAZIE MILLE, abbiamo lavoro duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Sarebbe ancora piu' apprezzatto se ci <a suggest-idea>suggerisci qualche miglioramento</a>.
	`,
		jaJP: `ありがとうございます、私たちは一生懸命働きました！

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>改善を提案</a>していただければ幸いです。
	`,
		koKR: `감사합니다, 우리는 열심히 일했습니다!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>개선 사항을 제안</a>해 주시면 감사하겠습니다.
	`,
		pl: `Dziękujemy, ciężko pracowaliśmy!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli <a suggest-idea>zaproponujesz ulepszenia</a>.
	`,
		ptBR: `Obrigado, trabalhamos duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você <a suggest-idea>sugerir melhorias</a>.
	`,
		ptPT: `Obrigado, trabalhamos arduamente!

 {{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

 Agradecíamos também que <a suggest-idea>sugerisse melhorias</a> .
	`,
		ru: `Спасибо, мы очень старались!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
	`,
		tr: `Teşekkürler, çok çalıştık!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>İyileştirmeler önerirseniz</a> de memnun oluruz.
	`,
		ukUA: `Дякуємо, ми старанно працювали!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви <a suggest-idea>запропонуєте покращення</a>.
	`,
		uz: `Rahmat, biz qattiq ishladik!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz <a suggest-idea>yaxshilanishlarni taklif qilsangiz</a>, biz ham minnatdor bo'lamiz.
	`,
		zhCN: `谢谢，我们努力工作！

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您<a suggest-idea>提出改进建议</a>，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		arEG: `يمكنك مساعدتنا كثيرًا إذا: * أعطنا 5 نجوم في <a storebot>دليل الروبوتات</a> . * أخبر أصدقائك عن التطبيق. على سبيل المثال على <a share-fb>Facebook</a> أو <a share-twitter>Twitter</a> . * دعم التطوير الإضافي - <a href = "https://goo.gl/Qhh0yL">2 يورو عبر PayPal</a> ( <i>حوالي 2.2 دولارًا</i> )`,
		de: `
Es gibt viele Wege uns zu helfen:

* Gib uns 5⭐ im <a storebot>Verzeichnis der Bots</a>.

* Erzähl am besten all deinen Freunde davon.
Du könntest es auf <a share-fb>Facebook</a> posten oder auf <a share-twitter>Twitter</a> twittern.

* Ansonsten auch gerne eine kleine Spende - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		en: `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		es: `
Nos ayudarías mucho si:

* Nos pusieras 5⭐ en <a storebot>directory of bots</a>.

* Hablaras del bot a tus amigos.
Por ejemplo <a share-fb>Facebook</a> o <a share-twitter>Twitter</a>.

* Apoyaras el siguiente trabajo - <a href = "https://goo.gl/Qhh0yL">€2 vía PayPal</a> (<i>about $2.2</i>)
`,
		faIR: `:شما به ما کمک بسیاری می کنید اگر

* ثبت بازخورد مثبت در <a storebot>دایرکتوری روبات ها</a>.

* به دوستان خود در مورد برنامه اطلاع رسانی کنید.
برای مثال در <a share-fb>فیسبوک</a> یا <a share-twitter>توئیتر</a>.

* از توسعه بیشتر برنامه پشتیبانی کنید - <a href = "https://goo.gl/Qhh0yL">€2 از طریق پی پال</a> (<i>$2.2 حدود </i>)`,
		fr: `
Vous pouvez nous aider beaucoup si vous:

* Nous donnez 5⭐ au <a storebot>répertoire des bots</a>.

* Parlez de l'application à vos amis.
Par exemple sur <a share-fb>Facebook</a> ou <a share-twitter>Twitter</a>.

* Soutenez le développement ultérieur - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>environ $2.2</i>)
`,
		id: `
Anda dapat membantu kami banyak jika Anda:

* Beri kami 5⭐ di <a storebot>direktori bot</a>.

* Ceritakan tentang aplikasi kepada teman-teman Anda.
Misalnya di <a share-fb>Facebook</a> atau <a share-twitter>Twitter</a>.

* Dukung pengembangan lebih lanjut - <a href = "https://goo.gl/Qhh0yL">€2 melalui PayPal</a> (<i>sekitar $2.2</i>)
`,
		it: `
Ci aiuteresti moltissimo se:

* Lasci un feedback positivo 5⭐ alla <a storebot>pagina del bot</a>.

* Raccontare dell'app ai tuoi amici.
Per esempio su <a share-fb>Facebook</a> o su <a share-twitter>Twitter</a>.

* Supporta ulteriormente lo sviluppo del bot - <a href="https://goo.gl/Qhh0yL">2€ tramite PayPal</a> (<i>circa $2.2</i>)
`,
		jaJP: `
あなたが以下のことをしてくれると、私たちにとって大きな助けになります：

* <a storebot>ボットディレクトリ</a>で5⭐を付けてください。

* アプリについて友達に教えてください。
例えば<a share-fb>Facebook</a>や<a share-twitter>Twitter</a>で。

* さらなる開発をサポートする - <a href = "https://goo.gl/Qhh0yL">PayPalで€2</a> (<i>約$2.2</i>)
`,
		koKR: `
다음과 같은 방법으로 많은 도움을 주실 수 있습니다:

* <a storebot>봇 디렉토리</a>에서 5⭐을 주세요.

* 앱에 대해 친구들에게 알려주세요.
예를 들어 <a share-fb>Facebook</a> 또는 <a share-twitter>Twitter</a>에서.

* 추가 개발 지원 - <a href = "https://goo.gl/Qhh0yL">PayPal을 통해 €2</a> (<i>약 $2.2</i>)
`,
		pl: `
Możesz nam bardzo pomóc, jeśli:

* Dasz nam 5⭐ w <a storebot>katalogu botów</a>.

* Opowiesz o aplikacji swoim znajomym.
Na przykład na <a share-fb>Facebooku</a> lub <a share-twitter>Twitterze</a>.

* Wesprzyj dalszy rozwój - <a href = "https://goo.gl/Qhh0yL">€2 przez PayPal</a> (<i>około $2.2</i>)
`,
		ptBR: `
Você pode nos ajudar muito se:

* Nos der 5⭐ no <a storebot>diretório de bots</a>.

* Contar sobre o aplicativo para seus amigos.
Por exemplo, no <a share-fb>Facebook</a> ou <a share-twitter>Twitter</a>.

* Apoiar o desenvolvimento adicional - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>cerca de $2.2</i>)
`,
		ptPT: `
Pode ajudar-nos muito se:

* Nos der 5⭐ no <a storebot>diretório de bots</a> .

* Contar sobre a aplicação aos seus amigos.
Por exemplo, no <a share-fb>Facebook</a> ou <a share-twitter>no Twitter</a> .

* Apoiar o desenvolvimento futuro - <a href = "https://goo.gl/Qhh0yL">2 € via PayPal</a> ( <i>cerca de 2,2 dólares</i> )
`,
		ru: `
	Вы нам очень поможете если:

	* Поставите нам 5⭐ в <a storebot>каталоге ботов</a>.

	* Расскажите о боте своим друзьям.
	Например в <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

* Поддержите дальнейшую разработку - <a href ="https://goo.gl/Qhh0yL">€2 через PayPal</a>
`,
		tr: `
Şunları yaparsanız bize çok yardımcı olabilirsiniz:

* <a storebot>bot dizininde</a> bize 5⭐ verin.

* Uygulama hakkında arkadaşlarınıza anlatın.
Örneğin <a share-fb>Facebook</a> veya <a share-twitter>Twitter</a>'da.

* Daha fazla geliştirmeyi destekleyin - <a href = "https://goo.gl/Qhh0yL">PayPal üzerinden €2</a> (<i>yaklaşık $2.2</i>)
`,
		ukUA: `
Ви можете дуже допомогти нам, якщо:

* Дасте нам 5⭐ у <a storebot>каталозі ботів</a>.

* Розкажете про додаток своїм друзям.
Наприклад, у <a share-fb>Facebook</a> або <a share-twitter>Twitter</a>.

* Підтримаєте подальший розвиток - <a href = "https://goo.gl/Qhh0yL">€2 через PayPal</a> (<i>приблизно $2.2</i>)
`,
		uz: `
Agar quyidagilarni qilsangiz, bizga juda yordam bergan bo'lasiz:

* <a storebot>botlar katalogida</a> bizga 5⭐ bering.

* Ilova haqida do'stlaringizga ayting.
Masalan, <a share-fb>Facebook</a> yoki <a share-twitter>Twitter</a>da.

* Keyingi rivojlanishni qo'llab-quvvatlang - <a href = "https://goo.gl/Qhh0yL">PayPal orqali €2</a> (<i>taxminan $2.2</i>)
`,
		zhCN: `
如果您：

* 在<a storebot>机器人目录</a>中给我们5⭐。

* 向您的朋友介绍这个应用程序。
例如在<a share-fb>Facebook</a>或<a share-twitter>Twitter</a>上。

* 支持进一步开发 - <a href = "https://goo.gl/Qhh0yL">通过PayPal支付€2</a> (<i>约$2.2</i>)
`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		arEG: `الرصيد فارغ لـ %v`,
		de:   `Du bist quitt mit %v`,
		en:   `Balance is empty for %v`,
		es:   `El balance es cero para %v`,
		faIR: `تراز خالی است برای %v`,
		fr:   `Le solde est vide pour %v`,
		id:   `Saldo kosong untuk %v`,
		it:   `Non hai alcun credito o debito con %v`,
		jaJP: `%vの残高は空です`,
		koKR: `%v에 대한 잔액이 비어 있습니다`,
		pl:   `Saldo jest puste dla %v`,
		ptBR: `Saldo está vazio para %v`,
		ptPT: `O saldo está vazio para %v`,
		ru:   `Нулевой баланс для %v`,
		tr:   `%v için bakiye boş`,
		ukUA: `Баланс порожній для %v`,
		uz:   `%v uchun balans bo'sh`,
		zhCN: `%v的余额为空`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		arEG: `هل تريد أن يتحدث روبوتنا بلغة أخرى؟ يمكنك <a>المساعدة في الترجمة</a> .`,
		de:   `Möchtest du den Bot in einer anderen Sprache? Du kannst beim <a>Übersetzen helfen</a>.`,
		en:   `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		es:   `¿Te gustaría que nuestro bot hablara en otro idioma? Puedes <a>ayudar con traducción</a>.`,
		faIR: `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,
		fr:   `Voulez-vous que notre bot parle dans une autre langue? Vous pouvez <a>aider à la traduction</a>.`,
		id:   `Apakah Anda ingin bot kami berbicara dalam bahasa lain? Anda dapat <a>membantu dengan terjemahan</a>.`,
		it:   `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
		jaJP: `私たちのボットが他の言語で話すことを望みますか？あなたは<a>翻訳を手伝う</a>ことができます。`,
		koKR: `우리 봇이 다른 언어로 말하기를 원하십니까? <a>번역을 도울 수 있습니다</a>.`,
		pl:   `Czy chcesz, aby nasz bot mówił w innym języku? Możesz <a>pomóc w tłumaczeniu</a>.`,
		ptBR: `Você quer que nosso bot fale em outro idioma? Você pode <a>ajudar com a tradução</a>.`,
		ptPT: `Quer que o nosso bot fale noutro idioma? Pode <a>ajudar com a tradução</a> .`,
		ru:   `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		tr:   `Botumuzun başka bir dilde konuşmasını ister misiniz? <a>Çeviriye yardımcı olabilirsiniz</a>.`,
		ukUA: `Хочете, щоб наш бот розмовляв іншою мовою? Ви можете <a>допомогти з перекладом</a>.`,
		uz:   `Botimiz boshqa tilda gaplashishini xohlaysizmi? Siz <a>tarjima qilishda yordam</a> berishingiz mumkin.`,
		zhCN: `您想要我们的机器人用其他语言交谈吗？您可以<a>帮助翻译</a>。`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		arEG: `حسنًا، لقد عملنا بجد. سيتم إرسال ملاحظاتك إلى المطورين.

هل يمكنك <a submit-bug>الإبلاغ عن مشكلتك</a> أو <a suggest-idea>اقتراح كيفية التحسين</a> ؟`,
		de: `Gut, wir geben uns Mühe. Deine Rückmeldung wird an die Entwickler weitergeleitet.

Vielleicht willst du <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		en: `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		es: `Bueno, hemos trabajado duro. Tu opinión se pasará a los creadores.

Quizás puedas <a submit-bug>informarnos de algún problema</a> o <a suggest-idea>proponernos qué podríamos mejorar</a>?`,
		faIR: `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		fr: `Eh bien, nous avons travaillé dur. Vos commentaires seront transmis aux développeurs.

Peut-être pouvez-vous <a submit-bug>signaler votre problème</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		id: `Yah, kami bekerja keras. Umpan balik Anda akan diteruskan ke pengembang.

Mungkin Anda dapat <a submit-bug>melaporkan masalah Anda</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		it: `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
		jaJP: `まあ、私たちは一生懸命働きました。あなたのフィードバックは開発者に伝えられます。

<a submit-bug>問題を報告</a>したり、<a suggest-idea>改善方法を提案</a>したりすることもできますか？`,
		koKR: `음, 우리는 열심히 일했습니다. 귀하의 피드백은 개발자에게 전달됩니다.

<a submit-bug>문제를 보고</a>하거나 <a suggest-idea>개선 방법을 제안</a>할 수 있습니까?`,
		pl: `Cóż, ciężko pracowaliśmy. Twoja opinia zostanie przekazana deweloperom.

Może możesz <a submit-bug>zgłosić swój problem</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		ptBR: `Bem, trabalhamos duro. Seu feedback será passado para os desenvolvedores.

Talvez você possa <a submit-bug>relatar seu problema</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		ptPT: `Bem, trabalhamos muito. O seu feedback será repassado aos desenvolvedores. 

Talvez possa <a submit-bug>reportar o seu problema</a> ou <a suggest-idea>sugerir como podemos melhorar</a> ?`,
		ru: `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		tr: `Eh, çok çalıştık. Geri bildiriminiz geliştiricilere iletilecektir.

Belki <a submit-bug>sorununuzu bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		ukUA: `Ну, ми старанно працювали. Ваш відгук буде переданий розробникам.

Можливо, ви можете <a submit-bug>повідомити про свою проблему</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		uz: `Xo'sh, biz qattiq ishladik. Sizning fikr-mulohazalaringiz ishlab chiqaruvchilarga yetkaziladi.

Balki siz <a submit-bug>muammoingizni xabar qilishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		zhCN: `嗯，我们努力工作。您的反馈将传递给开发人员。

也许您可以<a submit-bug>报告您的问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		arEG: `نأسف بشدة. هل يمكنك <a submit-bug>إخبارنا بالمشكلة</a> أو <a suggest-idea>اقتراح كيفية تحسينها</a> ؟`,
		de:   `Das tut uns sehr leid. Vielleicht willst du uns <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		en:   `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		es:   `Lo sentimos mucho. Igual podrías <a submit-bug>decirnos qué no funcciona bien</a> o <a suggest-idea>proponernos cómo podemos mejorarlo</a>?`,
		faIR: `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		fr:   `Nous sommes vraiment désolés. Peut-être pouvez-vous <a submit-bug>nous faire savoir ce qui ne va pas</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		id:   `Kami sangat menyesal. Mungkin Anda dapat <a submit-bug>memberi tahu kami apa yang salah</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		it:   `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
		jaJP: `大変申し訳ありません。<a submit-bug>何が問題かを教えていただく</a>か、<a suggest-idea>改善方法を提案していただく</a>ことはできますか？`,
		koKR: `매우 죄송합니다. <a submit-bug>무엇이 잘못되었는지 알려주시거나</a> <a suggest-idea>개선 방법을 제안</a>해 주실 수 있습니까?`,
		pl:   `Bardzo nam przykro. Może możesz <a submit-bug>dać nam znać, co jest nie tak</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		ptBR: `Lamentamos muito. Talvez você possa <a submit-bug>nos informar o que está errado</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		ptPT: `Lamentamos muito. Talvez <a submit-bug>nos possa dizer o que está mal</a> ou <a suggest-idea>sugerir como podemos melhorar</a> ?`,
		ru:   `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		tr:   `Çok üzgünüz. Belki <a submit-bug>neyin yanlış olduğunu bize bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		ukUA: `Нам дуже шкода. Можливо, ви можете <a submit-bug>повідомити нам, що не так</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		uz:   `Juda uzr so'raymiz. Balki siz <a submit-bug>nima noto'g'ri ekanligini bizga aytib berishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		zhCN: `我们非常抱歉。也许您可以<a submit-bug>告诉我们哪里出了问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		arEG: "يرجى تقييم تطبيقنا",
		de:   "Bitte bewerte unsere App",
		en:   "Please rate our app",
		es:   "Por favor valora nuestro app",
		faIR: "لطفاً به برنامه ما امتیاز دهید",
		fr:   "Veuillez évaluer notre application",
		id:   "Silakan nilai aplikasi kami",
		it:   "Per favore vota il nostro bot",
		jaJP: "アプリを評価してください",
		koKR: "앱을 평가해 주세요",
		pl:   "Oceń naszą aplikację",
		ptBR: "Por favor, avalie nosso aplicativo",
		ptPT: "Por favor, avalie a nossa aplicação",
		ru:   "Оцените наше приложение?",
		tr:   "Lütfen uygulamamızı değerlendirin",
		ukUA: "Будь ласка, оцініть наш додаток",
		uz:   "Iltimos, ilovamizni baholang",
		zhCN: "请评价我们的应用",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		arEG: "نعم إنه تطبيق رائع!",
		de:   "Ja, es ist eine tolle App!",
		en:   "Yes, it's a great app!",
		es:   "¡Sí, es una app fantástica!",
		faIR: "بله، این برنامه عالی است",
		fr:   "Oui, c'est une excellente application !",
		id:   "Ya, ini aplikasi yang bagus!",
		it:   "Si, e' un app fantastica!",
		jaJP: "はい、素晴らしいアプリです！",
		koKR: "네, 훌륭한 앱입니다!",
		pl:   "Tak, to świetna aplikacja!",
		ptBR: "Sim, é um ótimo aplicativo!",
		ptPT: "Sim, é uma ótima aplicação!",
		ru:   "Да, отличное приложение!",
		tr:   "Evet, harika bir uygulama!",
		ukUA: "Так, це чудовий додаток!",
		uz:   "Ha, bu ajoyib ilova!",
		zhCN: "是的，这是一个很棒的应用！",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		arEG: "ليس سيئا ولكن يمكن أن يكون أفضل.",
		de:   "Nicht schlecht, aber könnte besser sein",
		en:   "Not bad but can be better.",
		es:   "No está mal, pero podría ser mejor.",
		faIR: "بد نیست ولی می تواند بهتر باشد.",
		fr:   "Pas mal mais peut être amélioré.",
		id:   "Tidak buruk tapi bisa lebih baik.",
		it:   "Non male ma potrebbe esser migliore.",
		jaJP: "悪くないですが、もっと良くなる可能性があります。",
		koKR: "나쁘지 않지만 더 좋아질 수 있습니다.",
		pl:   "Niezły, ale może być lepszy.",
		ptBR: "Não é ruim, mas pode ser melhor.",
		ptPT: "Não é mau, mas pode ser melhor.",
		ru:   "Неплохо, но можно лучше.",
		tr:   "Fena değil ama daha iyi olabilir.",
		ukUA: "Непогано, але може бути краще.",
		uz:   "Yomon emas, lekin yaxshiroq bo'lishi mumkin.",
		zhCN: "不错，但可以更好。",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		arEG: "لا يعجبني",
		de:   "Ich mag sie nicht",
		en:   "Don't like it",
		es:   "No me gusta",
		faIR: "از این برنامه را نمی پسندم",
		fr:   "Je n'aime pas",
		id:   "Tidak suka",
		it:   "Non mi piace",
		jaJP: "好きではない",
		koKR: "마음에 들지 않음",
		pl:   "Nie lubię tego",
		ptBR: "Não gosto",
		ptPT: "Não gosto",
		ru:   "Не нравится",
		tr:   "Beğenmedim",
		ukUA: "Не подобається",
		uz:   "Yoqmadi",
		zhCN: "不喜欢",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		arEG: "لم يتم اتخاذ القرار بعد",
		de:   "Ich bin mir noch unsicher",
		en:   "Not decided yet",
		es:   "Estoy aún indeciso",
		faIR: "هنوز تصمیم نگرفته ام.",
		fr:   "Pas encore décidé",
		id:   "Belum memutuskan",
		it:   "Sono indeciso",
		jaJP: "まだ決めていません",
		koKR: "아직 결정하지 않았습니다",
		pl:   "Jeszcze nie zdecydowałem",
		ptBR: "Ainda não decidi",
		ptPT: "Ainda não decidido",
		ru:   "Пока не понятно",
		tr:   "Henüz karar vermedim",
		ukUA: "Ще не вирішив",
		uz:   "Hali qaror qabul qilinmagan",
		zhCN: "尚未决定",
	},
	SettingsMessageText: {
		arEG: "ماذا تريد تعديله؟",
		de:   "Was willst du ändern?",
		en:   "What do you want to adjust?",
		es:   "¿Qué te gustaría modificar?",
		faIR: "می خواهید چه چیزی را تنظیم کنید؟",
		fr:   "Que voulez-vous ajuster ?",
		id:   "Apa yang ingin Anda sesuaikan?",
		it:   "Cosa vuoi modificare?",
		jaJP: "何を調整したいですか？",
		koKR: "무엇을 조정하고 싶으신가요?",
		pl:   "Co chcesz dostosować?",
		ptBR: "O que você deseja ajustar?",
		ptPT: "O que pretende ajustar?",
		ru:   "Что будем настраивать?",
		tr:   "Neyi ayarlamak istiyorsunuz?",
		ukUA: "Що ви хочете налаштувати?",
		uz:   "Nimani sozlamoqchisiz?",
		zhCN: "您想调整什么？",
	},
	MT_NO_OUTSTANDING_TRANSFERS: {
		arEG: `تحاول إنشاء سجل إرجاع ولكن لا توجد ديون مستحقة. إذا كنت تعتقد أن هذا خطأ، فيرجى إعلامنا بذلك على @DebtsTrackerGroup.`,
		de: `Sie versuchen, einen Rückgabedatensatz zu erstellen, aber es gibt keine ausstehenden Schulden.

Wenn Sie glauben, dass dies ein Fehler ist, teilen Sie uns dies bitte in @DebtsTrackerGroup mit.`,
		en: `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		es: `Estás intentando crear un registro de devolución pero no hay deudas pendientes.

Si crees que esto es un error, háganoslo saber en @DebtsTrackerGroup.`,
		faIR: `شما در حال تلاش برای ایجاد سابقه بازگشت هستید اما هیچ بدهی معوقه ای وجود ندارد.

اگر فکر می کنید این یک اشتباه است، لطفاً به ما در @DebtsTrackerGroup اطلاع دهید.`,
		fr: `Vous essayez de créer un enregistrement de retour mais il n'y a pas de dettes en cours.

Si vous pensez qu'il s'agit d'une erreur, veuillez nous en informer dans @DebtsTrackerGroup.`,
		id: `Anda mencoba membuat catatan pengembalian tetapi tidak ada hutang yang belum dibayar.

Jika Anda yakin ini adalah kesalahan, beri tahu kami di @DebtsTrackerGroup.`,
		it: `Stai cercando di creare un record di restituzione ma non ci sono debiti in sospeso.

Se ritieni che si tratti di un errore, ti preghiamo di farcelo sapere in @DebtsTrackerGroup.`,
		jaJP: `返却記録を作成しようとしていますが、未払いの借金はありません。

これが間違いだと思われる場合は、@DebtsTrackerGroupでお知らせください。`,
		koKR: `반환 기록을 생성하려고 하지만 미결제 부채가 없습니다.

이것이 실수라고 생각되면 @DebtsTrackerGroup에서 알려주십시오.`,
		pl: `Próbujesz utworzyć rekord zwrotu, ale nie ma zaległych długów.

Jeśli uważasz, że to pomyłka, daj nam znać w @DebtsTrackerGroup.`,
		ptBR: `Você está tentando criar um registro de devolução, mas não há dívidas pendentes.

Se você acredita que isso é um erro, por favor, nos avise em @DebtsTrackerGroup.`,
		ptPT: `Está a tentar criar um registo de devolução, mas não há dívidas pendentes.

Se acredita que se trata de um erro, por favor informe-nos em @DebtsTrackerGroup.`,
		ru: `Вы пытаетесь создать запись о возврате долга, но мы не нашли не закрытых долгов.

Если вы считаете что это ошибка пожалуйста дайте нам знать в @DebtsTrackerGroup.`,
		tr: `İade kaydı oluşturmaya çalışıyorsunuz ancak bekleyen borç bulunmamaktadır.

Bunun bir hata olduğunu düşünüyorsanız, lütfen @DebtsTrackerGroup'ta bize bildirin.`,
		ukUA: `Ви намагаєтеся створити запис про повернення, але немає непогашених боргів.

Якщо ви вважаєте, що це помилка, будь ласка, повідомте нам у @DebtsTrackerGroup.`,
		uz: `Siz qaytarish yozuvini yaratmoqchi bo'lyapsiz, lekin to'lanmagan qarzlar yo'q.

Agar bu xato deb o'ylasangiz, iltimos, bizga @DebtsTrackerGroup'da xabar bering.`,
		zhCN: `您正在尝试创建退款记录，但没有未偿还的债务。

如果您认为这是一个错误，请在 @DebtsTrackerGroup 中告诉我们。`,
	},
	MT_ATTEMPT_TO_CREATE_DEBT_WITH_INTEREST_AFFECTING_OUTSTANDING: {
		arEG: "أنت تحاول إنشاء دين بفائدة، مما سيؤثر على التحويلات المستحقة. يُرجى إغلاقها أولاً.",
		de:   "Sie versuchen, eine Schuld mit Zinsen zu erstellen, die sich auf ausstehende Überweisungen auswirken wird. Bitte schließen Sie diese zuerst.",
		en:   "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		es:   "Estás intentando crear una deuda con intereses que afectará a las transferencias pendientes. Por favor, ciérralas primero.",
		faIR: "شما در حال تلاش برای ایجاد بدهی با بهره هستید که بر انتقال های معوق تأثیر می گذارد. لطفا ابتدا آنها را ببندید.",
		fr:   "Vous essayez de créer une dette avec intérêt qui affectera les transferts en cours. Veuillez les fermer d'abord.",
		id:   "Anda mencoba membuat hutang dengan bunga yang akan memengaruhi transfer yang belum diselesaikan. Harap tutup terlebih dahulu.",
		it:   "Stai cercando di creare un debito con interessi che influenzerà i trasferimenti in sospeso. Per favore, chiudili prima.",
		jaJP: "未決済の送金に影響する利息付きの借金を作成しようとしています。まずそれらを閉じてください。",
		koKR: "미결제 송금에 영향을 미치는 이자가 있는 부채를 만들려고 합니다. 먼저 그것들을 닫으십시오.",
		pl:   "Próbujesz utworzyć dług z odsetkami, który wpłynie na nierozliczone przelewy. Proszę najpierw je zamknąć.",
		ptBR: "Você está tentando criar uma dívida com juros que afetará transferências pendentes. Por favor, feche-as primeiro.",
		ptPT: "Está a tentar criar uma dívida com juros que afetará as transferências pendentes. Por favor, feche-as primeiro.",
		ru:   "Вы пытаетесь создать запись о долге с процентами которая затронет незакрытые долги. Пожалуйста закройте их сначала.",
		tr:   "Bekleyen transferleri etkileyecek faizli bir borç oluşturmaya çalışıyorsunuz. Lütfen önce onları kapatın.",
		ukUA: "Ви намагаєтеся створити борг з відсотками, який вплине на невиконані перекази. Будь ласка, спочатку закрийте їх.",
		uz:   "Siz to'lanmagan o'tkazmalarga ta'sir qiladigan foizli qarz yaratmoqchisiz. Iltimos, avval ularni yoping.",
		zhCN: "您正在尝试创建一个带有利息的债务，这将影响未结清的转账。请先关闭它们。",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		arEG: "عذراً، هذه الوظيفة لم يتم تنفيذها بعد.",
		de:   "Entschuldigung, diese Funktion ist noch nicht fertig programmiert.",
		en:   "Sorry, this functionality is not implemented yet.",
		es:   "Lo sentimos, esta función no está activa aún.",
		faIR: "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
		fr:   "Désolé, cette fonctionnalité n'est pas encore implémentée.",
		id:   "Maaf, fungsi ini belum diimplementasikan.",
		it:   "Spiacenti ma questa funzionalita' non e' ancora attiva.",
		jaJP: "申し訳ありませんが、この機能はまだ実装されていません。",
		koKR: "죄송합니다. 이 기능은 아직 구현되지 않았습니다.",
		pl:   "Przepraszamy, ta funkcjonalność nie jest jeszcze zaimplementowana.",
		ptBR: "Desculpe, esta funcionalidade ainda não foi implementada.",
		ptPT: "Lamentamos, esta funcionalidade ainda não foi implementada.",
		ru:   "Извините, данный функционал ещё не реализован.",
		tr:   "Üzgünüz, bu işlevsellik henüz uygulanmadı.",
		ukUA: "Вибачте, ця функціональність ще не реалізована.",
		uz:   "Kechirasiz, bu funksionallik hali amalga oshirilmagan.",
		zhCN: "抱歉，此功能尚未实现。",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		arEG: "كيف تريد الحصول على دعوة؟",
		de:   "Worüber möchtest du eingeladen werden?",
		en:   "How do you want to get an invite?",
		es:   "¿Comó prefieres recibir la invitación?",
		faIR: "می خواهید چگونه دعوت شوید؟",
		fr:   "Comment voulez-vous recevoir une invitation ?",
		id:   "Bagaimana Anda ingin mendapatkan undangan?",
		it:   "Come vuoi ricevere l'invito?",
		jaJP: "招待状をどのように受け取りますか？",
		koKR: "초대를 어떻게 받고 싶으신가요?",
		pl:   "Jak chcesz otrzymać zaproszenie?",
		ptBR: "Como você deseja receber um convite?",
		ptPT: "Como quer receber um convite?",
		ru:   "Как вы хотите получить код приглашения?",
		tr:   "Daveti nasıl almak istersiniz?",
		ukUA: "Як ви хочете отримати запрошення?",
		uz:   "Taklifnomani qanday olishni xohlaysiz?",
		zhCN: "您想如何获得邀请？",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		arEG: "الرجاء إدخال رمز الدعوة:",
		de:   "Bitte gib den Bestätigungs-Code ein:",
		en:   "Please enter an invite code:",
		es:   "Introduce el código de la invitación",
		faIR: "لطفاً یک کد دعوت وارد کنید:",
		fr:   "Veuillez entrer un code d'invitation :",
		id:   "Silakan masukkan kode undangan:",
		it:   "Inserisci un codice invito:",
		jaJP: "招待コードを入力してください：",
		koKR: "초대 코드를 입력하세요:",
		pl:   "Proszę wprowadzić kod zaproszenia:",
		ptBR: "Por favor, insira um código de convite:",
		ptPT: "Por favor, introduza um código de convite:",
		ru:   "Пожалуйста введите код приглашения:",
		tr:   "Lütfen bir davet kodu girin:",
		ukUA: "Будь ласка, введіть код запрошення:",
		uz:   "Iltimos, taklif kodini kiriting:",
		zhCN: "请输入邀请码：",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		arEG: "لقد أرسلنا رسالة إلى %v.\n\nالرجاء فتح البريد الإلكتروني والنقر فوق الرابط لتأكيد عنوان بريدك الإلكتروني.",
		de:   "Wir haben eine Nachricht an %v gesendet.\n\nBitte öffne die Nachricht und klick auf den Link, um deine Mail-Adresse zu bestätigen.",
		en:   "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		es:   "Hemos enviado un mensage a %v.\n\nPor favor, abre tu e-mail y haz click en el link para confirmar tu e-mail.",
		faIR: "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
		fr:   "Nous avons envoyé un message à %v.\n\nVeuillez ouvrir l'e-mail et cliquer sur un lien pour confirmer votre adresse e-mail.",
		id:   "Kami telah mengirim pesan ke %v.\n\nSilakan buka email dan klik tautan untuk mengonfirmasi alamat email Anda.",
		it:   "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
		jaJP: "%vにメッセージを送信しました。\n\nメールを開き、リンクをクリックしてメールアドレスを確認してください。",
		koKR: "%v에 메시지를 보냈습니다.\n\n이메일을 열고 링크를 클릭하여 이메일 주소를 확인하십시오.",
		pl:   "Wysłaliśmy wiadomość do %v.\n\nProszę otworzyć e-mail i kliknąć link, aby potwierdzić swój adres e-mail.",
		ptBR: "Enviamos uma mensagem para %v.\n\nPor favor, abra o e-mail e clique em um link para confirmar seu endereço de e-mail.",
		ptPT: "Enviámos uma mensagem para %v. \n\nAbra o e-mail e clique num link para confirmar o seu endereço de e-mail.",
		ru:   "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		tr:   "%v adresine bir mesaj gönderdik.\n\nLütfen e-postayı açın ve e-posta adresinizi onaylamak için bir bağlantıya tıklayın.",
		ukUA: "Ми надіслали повідомлення на %v.\n\nБудь ласка, відкрийте електронний лист і натисніть посилання, щоб підтвердити свою електронну адресу.",
		uz:   "Biz %v ga xabar yubordik.\n\nIltimos, elektron pochtani oching va elektron pochta manzilingizni tasdiqlash uchun havolani bosing.",
		zhCN: "我们已向 %v 发送了一条消息。\n\n请打开电子邮件并点击链接确认您的电子邮件地址。",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		arEG: "بمجرد ظهور Telegram، انقر فوق زر <b>البدء</b> .",
		de:   "Wenn Telegram öffnet, drücke auf <b>Start</b>.",
		en:   "Once Telegram pop ups click the <b>Start</b> button.",
		es:   "Después de abrir Telegram aprieta el <b>Start</b> botón.",
		faIR: "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
		fr:   "Une fois que Telegram apparaît, cliquez sur le bouton <b>Start</b>.",
		id:   "Setelah Telegram muncul, klik tombol <b>Start</b>.",
		it:   "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
		jaJP: "Telegramがポップアップしたら、<b>Start</b>ボタンをクリックしてください。",
		koKR: "Telegram이 팝업되면 <b>Start</b> 버튼을 클릭하세요.",
		pl:   "Gdy pojawi się Telegram, kliknij przycisk <b>Start</b>.",
		ptBR: "Quando o Telegram aparecer, clique no botão <b>Start</b>.",
		ptPT: "Quando a janela do Telegram for apresentada, clique no botão <b>Iniciar</b> .",
		ru:   "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		tr:   "Telegram açıldığında <b>Start</b> düğmesine tıklayın.",
		ukUA: "Коли з'явиться Telegram, натисніть кнопку <b>Start</b>.",
		uz:   "Telegram ochilganda, <b>Start</b> tugmasini bosing.",
		zhCN: "Telegram 弹出后，点击 <b>Start</b> 按钮。",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		arEG: "شكرًا لك، لقد تم إدراجك في قائمة انتظار الدعوة.\n\nمدة الانتظار الحالية هي 2-3 أيام.\n\nيمكنك الحصول على رمز دعوة اليوم من خلال مشاركة رابط للروبوت على Facebook.",
		de:   "Danke, du bist in der Warteschlange für eine Einladung.\n\nEs dauert etwa zwei bis drei Tage.\n\nAber du könntest den Code noch heute bekommen, wenn du einen Link auf Facebook teilst.",
		en:   "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		es:   "Gracias, ya estás inscrito en la cola para conseguir la invitación.\n\nTiempo de espera 2-3 días.\n\nPuedes conseguirlo hoy si compartes el link de nuestro bot en Facebook.",
		faIR: "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
		fr:   "Merci, vous avez été mis en file d'attente pour une invitation.\n\nLe temps d'attente actuel est de 2 à 3 jours.\n\nVous pouvez obtenir un code d'invitation aujourd'hui en partageant un lien vers le bot sur Facebook.",
		id:   "Terima kasih, Anda telah antri untuk undangan.\n\nWaktu menunggu saat ini adalah 2-3 hari.\n\nAnda dapat memperoleh kode undangan hari ini dengan membagikan tautan ke bot di Facebook.",
		it:   "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
		jaJP: "ありがとうございます。招待のために列に並んでいます。\n\n現在の待ち時間は2〜3日です。\n\nFacebookでボットへのリンクを共有することで、今日招待コードを取得できます。",
		koKR: "감사합니다. 초대를 위해 대기열에 올랐습니다.\n\n현재 대기 시간은 2-3일입니다.\n\nFacebook에서 봇에 대한 링크를 공유하여 오늘 초대 코드를 받을 수 있습니다.",
		pl:   "Dziękujemy, zostałeś umieszczony w kolejce do zaproszenia.\n\nObecny czas oczekiwania to 2-3 dni.\n\nMożesz otrzymać kod zaproszenia już dziś, udostępniając link do bota na Facebooku.",
		ptBR: "Obrigado, você foi colocado na fila para um convite.\n\nO tempo de espera atual é de 2 a 3 dias.\n\nVocê pode obter um código de convite hoje compartilhando um link para o bot no Facebook.",
		ptPT: "Obrigado, foi colocado na fila para um convite. \n\nO tempo de espera actual é de 2 a 3 dias. \n\nPode obter um código de convite hoje mesmo, partilhando um link para o bot no Facebook.",
		ru:   "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		tr:   "Teşekkürler, bir davetiye için sıraya alındınız.\n\nMevcut bekleme süresi 2-3 gündür.\n\nBotun bağlantısını Facebook'ta paylaşarak bugün bir davet kodu alabilirsiniz.",
		ukUA: "Дякуємо, вас поставлено в чергу на запрошення.\n\nПоточний час очікування - 2-3 дні.\n\nВи можете отримати код запрошення сьогодні, поділившись посиланням на бота у Facebook.",
		uz:   "Rahmat, siz taklif uchun navbatga qo'yildingiz.\n\nHozirgi kutish vaqti 2-3 kun.\n\nFacebookda botga havola ulashish orqali bugun taklif kodini olishingiz mumkin.",
		zhCN: "谢谢，您已排队等候邀请。\n\n当前等待时间为2-3天。\n\n您可以通过在Facebook上分享机器人链接，今天就获得邀请码。",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		arEG: "يرجى تقديم عنوان بريدك الإلكتروني",
		de:   "Bitte gib deine e-Mail Adresse an:",
		en:   "Please provide your email address",
		es:   "Por favor, esctibe tu e-mail",
		faIR: "لطفاً آدرس ایمیل خود را وارد کنید.",
		fr:   "Veuillez fournir votre adresse e-mail",
		id:   "Silakan berikan alamat email Anda",
		it:   "Inserisci il tuo indirizzo email:",
		jaJP: "メールアドレスを入力してください",
		koKR: "이메일 주소를 입력해 주세요",
		pl:   "Proszę podać swój adres e-mail",
		ptBR: "Por favor, forneça seu endereço de e-mail",
		ptPT: "Por favor, forneça o seu endereço de e-mail",
		ru:   "Пожалуйста напишите ваш email адрес:",
		tr:   "Lütfen e-posta adresinizi girin",
		ukUA: "Будь ласка, вкажіть вашу електронну адресу",
		uz:   "Iltimos, elektron pochta manzilingizni kiriting",
		zhCN: "请提供您的电子邮件地址",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		arEG: "يرجى تقديم رقم هاتفك",
		de:   "Bitte gib deine Telefonnummer an:",
		en:   "Please provide your phone number",
		es:   "Por favor, esctibe tu número de teléfono",
		faIR: "لطفاً شماره تلفن خود را وارد نمایید.",
		fr:   "Veuillez fournir votre numéro de téléphone",
		id:   "Silakan berikan nomor telepon Anda",
		it:   "Inserisci il tuo numero di telefono:",
		jaJP: "電話番号を入力してください",
		koKR: "전화번호를 입력해 주세요",
		pl:   "Proszę podać swój numer telefonu",
		ptBR: "Por favor, forneça seu número de telefone",
		ptPT: "Por favor, forneça o seu número de telefone",
		ru:   "Пожалуйста напишите номер вашего телефона:",
		tr:   "Lütfen telefon numaranızı girin",
		ukUA: "Будь ласка, вкажіть ваш номер телефону",
		uz:   "Iltimos, telefon raqamingizni kiriting",
		zhCN: "请提供您的电话号码",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		arEG: "رمز الدعوة خاطئ: %v",
		de:   "Ungültiger Code: %v",
		en:   "Wrong invite code: %v",
		es:   "El código de invitación no es correcto: %v",
		faIR: "کد دعوت اشتباه است %v",
		fr:   "Code d'invitation incorrect : %v",
		id:   "Kode undangan salah: %v",
		it:   "Codice invito: %v errato",
		jaJP: "招待コードが間違っています：%v",
		koKR: "초대 코드가 잘못되었습니다: %v",
		pl:   "Nieprawidłowy kod zaproszenia: %v",
		ptBR: "Código de convite errado: %v",
		ptPT: "Código de convite incorreto: %v",
		ru:   "Неправильный код приглашения: %v",
		tr:   "Yanlış davet kodu: %v",
		ukUA: "Неправильний код запрошення: %v",
		uz:   "Noto'g'ri taklif kodi: %v",
		zhCN: "邀请码错误：%v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		arEG: "عنوان البريد الإلكتروني خاطئ.",
		de:   "Ungültige e-Mail Adresse.",
		en:   "Wrong email address.",
		es:   "El e-mail no es correcto.",
		faIR: "آدرس ایمیل اشتباه است.",
		fr:   "Adresse e-mail incorrecte.",
		id:   "Alamat email salah.",
		it:   "L'email inserita e' sbagliata.",
		jaJP: "メールアドレスが間違っています。",
		koKR: "이메일 주소가 잘못되었습니다.",
		pl:   "Nieprawidłowy adres e-mail.",
		ptBR: "Endereço de e-mail incorreto.",
		ptPT: "Endereço de e-mail errado.",
		ru:   "Неправильный email адрес.",
		tr:   "Yanlış e-posta adresi.",
		ukUA: "Неправильна електронна адреса.",
		uz:   "Noto'g'ri elektron pochta manzili.",
		zhCN: "电子邮件地址错误。",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		arEG: "رقم الهاتف خاطئ.",
		de:   "Ungültige Telefonnummer.",
		en:   "Wrong phone number.",
		es:   "El número de telefono no es correcto.",
		faIR: "شماره تلفن اشتباه است",
		fr:   "Numéro de téléphone incorrect.",
		id:   "Nomor telepon salah.",
		it:   "Il numero inserito e' sbagliato.",
		jaJP: "電話番号が間違っています。",
		koKR: "전화번호가 잘못되었습니다.",
		pl:   "Nieprawidłowy numer telefonu.",
		ptBR: "Número de telefone incorreto.",
		ptPT: "Número de telefone errado.",
		ru:   "Неправильный номер телефона.",
		tr:   "Yanlış telefon numarası.",
		ukUA: "Неправильний номер телефону.",
		uz:   "Noto'g'ri telefon raqami.",
		zhCN: "电话号码错误。",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		arEG: "حسنًا، الرجاء المحاولة مرة أخرى.",
		de:   "Ok, bitte versuche es erneut.",
		en:   "Ok, please try again.",
		es:   "Ok, inténtalo de nuevo.",
		faIR: "بسیار خوب، لطفا مجدداً سعی کنید.",
		fr:   "Ok, veuillez réessayer.",
		id:   "Ok, silakan coba lagi.",
		it:   "Ok, prova di nuovo.",
		jaJP: "はい、もう一度お試しください。",
		koKR: "네, 다시 시도해 주세요.",
		pl:   "Ok, spróbuj ponownie.",
		ptBR: "Ok, tente novamente.",
		ptPT: "Ok, tente novamente.",
		ru:   "Хорошо, попробуйте ещё раз.",
		tr:   "Tamam, lütfen tekrar deneyin.",
		ukUA: "Добре, спробуйте ще раз.",
		uz:   "Mayli, qaytadan urinib ko'ring.",
		zhCN: "好的，请再试一次。",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		arEG: "لقد أخطأت في الكتابة، سأحاول مرة أخرى.",
		de:   "Ich habe mich vertippt.",
		en:   "I've mistyped, will try again.",
		es:   "Me he equivocado, lo intentaré otra vez",
		faIR: "اشتباه تایپ کردم، دوباره امتحان می کنم.",
		fr:   "J&#39;ai fait une erreur de frappe, je vais réessayer.",
		id:   "Saya salah ketik, akan mencoba lagi.",
		it:   "Ho sbagliato, riprovo.",
		jaJP: "入力ミスがありました。もう一度試してみます。",
		koKR: "잘못 입력했네요. 다시 입력해 보겠습니다.",
		pl:   "Zrobiłem literówkę, spróbuję jeszcze raz.",
		ptBR: "Digitei errado, vou tentar novamente.",
		ptPT: "Digitei mal, vou tentar novamente.",
		ru:   "Я опечатался, попробую ещё раз.",
		tr:   "Yanlış yazmışım, tekrar deneyeceğim.",
		ukUA: "Я помилився, спробую ще раз.",
		uz:   "Men noto&#39;g&#39;ri yozdim, yana urinib ko&#39;raman.",
		zhCN: "我打错了，再试一次。"},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		arEG: "أخبرني المزيد عن الرموز",
		de:   "Erzähl mir mehr über diese Codes",
		en:   "Tell me more about the codes",
		es:   "Dime más de los códigos",
		faIR: "در مورد کدها بیشتر برای من توضیح دهید.",
		fr:   "Parlez-moi davantage des codes",
		id:   "Ceritakan lebih lanjut tentang kodenya",
		it:   "Ulteriori informazioni riguardo il codice invito.",
		jaJP: "コードについて詳しく教えてください",
		koKR: "코드에 대해 더 자세히 알려주세요",
		pl:   "Powiedz mi więcej o kodach",
		ptBR: "Conte-me mais sobre os códigos",
		ptPT: "Conte-me mais sobre os códigos",
		ru:   "Расскажите ка мне об этих кодах",
		tr:   "Bana kodlar hakkında daha fazla bilgi ver",
		ukUA: "Розкажіть мені більше про коди",
		uz:   "Menga kodlar haqida ko&#39;proq ma&#39;lumot bering",
		zhCN: "告诉我有关代码的更多信息"},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		arEG: "أرسل لي رمز الدعوة عبر البريد الإلكتروني",
		de:   "Sende mir einen Code per Mail",
		en:   "Send me invite code by email",
		es:   "Envíame el código de invitación por e-mail",
		faIR: "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
		fr:   "Envoyez-moi le code d&#39;invitation par e-mail",
		id:   "Kirimkan saya kode undangan melalui email",
		it:   "Inviami il codice invito tramite email",
		jaJP: "招待コードをメールで送ってください",
		koKR: "이메일로 초대 코드를 보내주세요",
		pl:   "Wyślij mi kod zaproszenia e-mailem",
		ptBR: "Envie-me o código de convite por e-mail",
		ptPT: "Envie-me o código de convite por e-mail",
		ru:   "Хочу код приглашения на email",
		tr:   "Bana e-posta ile davet kodu gönder",
		ukUA: "Надіслати мені код запрошення електронною поштою",
		uz:   "Menga taklif kodini elektron pochta orqali yuboring",
		zhCN: "通过电子邮件向我发送邀请码"},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		arEG: "أرسل لي رمز الدعوة عبر الرسائل القصيرة",
		de:   "Sende mir einen Code per SMS",
		en:   "Send me invite code by SMS",
		es:   "Envíame el código de invitación a través de SMS",
		faIR: "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
		fr:   "Envoyez-moi un code d&#39;invitation par SMS",
		id:   "Kirimkan saya kode undangan melalui SMS",
		it:   "Inviami il codice invito tramite SMS",
		jaJP: "SMSで招待コードを送信してください",
		koKR: "SMS로 초대 코드를 보내주세요",
		pl:   "Wyślij mi kod zaproszenia SMS-em",
		ptBR: "Envie-me o código de convite por SMS",
		ptPT: "Envie-me o código de convite por SMS",
		ru:   "Хочу код приглашения по SMS",
		tr:   "Bana SMS ile davet kodu gönder",
		ukUA: "Надіслати мені код запрошення через SMS",
		uz:   "Menga taklif kodini SMS orqali yuboring",
		zhCN: "通过短信向我发送邀请码"},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		arEG: "أرسل لي رمز دعوة جديد عبر البريد الإلكتروني",
		de:   "Sende mir nochmal einen Code per Mail",
		en:   "Send me new invite code by email",
		es:   "Envíame un nuevo código de invitación por e-mail",
		faIR: "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
		fr:   "Envoyez-moi un nouveau code d&#39;invitation par e-mail",
		id:   "Kirimkan saya kode undangan baru melalui email",
		it:   "Inviami il nuovo codice invito tramite email",
		jaJP: "新しい招待コードをメールで送ってください",
		koKR: "이메일로 새로운 초대 코드를 보내주세요",
		pl:   "Wyślij mi nowy kod zaproszenia e-mailem",
		ptBR: "Envie-me um novo código de convite por e-mail",
		ptPT: "Envie-me um novo código de convite por e-mail",
		ru:   "Новый код приглашения на email",
		tr:   "Bana e-posta ile yeni davet kodu gönder",
		ukUA: "Надіслати мені новий код запрошення електронною поштою",
		uz:   "Menga yangi taklif kodini elektron pochta orqali yuboring",
		zhCN: "通过电子邮件向我发送新的邀请码"},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		arEG: "أرسل لي رمز دعوة جديد عبر الرسائل القصيرة",
		de:   "Sende mir nochmal einen Code per SMS",
		en:   "Send me new invite code by SMS",
		es:   "Envíame un nuevo código de invitación a través de SMS",
		faIR: "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
		fr:   "Envoyez-moi un nouveau code d&#39;invitation par SMS",
		id:   "Kirimkan saya kode undangan baru melalui SMS",
		it:   "Inviami il nuovo codice invito tramite SMS",
		jaJP: "SMSで新しい招待コードを送信してください",
		koKR: "SMS로 새로운 초대 코드를 보내주세요",
		pl:   "Wyślij mi nowy kod zaproszenia SMS-em",
		ptBR: "Envie-me um novo código de convite por SMS",
		ptPT: "Envie-me um novo código de convite por SMS",
		ru:   "Новый код приглашения по SMS",
		tr:   "Bana SMS ile yeni davet kodu gönder",
		ukUA: "Надіслати мені новий код запрошення через SMS",
		uz:   "Menga yangi taklif kodini SMS orqali yuboring",
		zhCN: "通过短信向我发送新的邀请码"},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		arEG: "أرسل لي دعوة جديدة عبر Telegram",
		de:   "Sende mir nochmal einen Code per Telegram",
		en:   "Send me new invite by Telegram",
		es:   "Envíame un nuevo código de invitación a través de Telegram",
		faIR: "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
		fr:   "Envoyez-moi une nouvelle invitation par Telegram",
		id:   "Kirimkan saya undangan baru melalui Telegram",
		it:   "Inviami il nuovo codice invito tramite Telegram",
		jaJP: "Telegramで新しい招待を送ってください",
		koKR: "Telegram으로 새로운 초대장을 보내주세요",
		pl:   "Wyślij mi nowe zaproszenie przez Telegram",
		ptBR: "Envie-me um novo convite pelo Telegram",
		ptPT: "Envie-me um novo convite através do Telegram",
		ru:   "Получить приграшение в Telegram",
		tr:   "Bana Telegram ile yeni davet gönder",
		ukUA: "Надіслати мені нове запрошення через Telegram",
		uz:   "Menga Telegram orqali yangi taklif yuboring",
		zhCN: "通过 Telegram 向我发送新邀请"},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		arEG: "لغة غير معروفة. يُرجى اختيار واحد من الخيارات التالية:",
		de:   "Unbekannte Sprache. Bitte wähle eine aus der Liste:",
		en:   "Unknown language. Please choose 1 from the options:",
		es:   "Idioma desconocido. Por favor elige 1 de las opciones:",
		faIR: "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
		fr:   "Langue inconnue. Veuillez choisir une option parmi les suivantes :",
		id:   "Bahasa tidak dikenal. Silakan pilih 1 dari pilihan berikut:",
		it:   "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
		jaJP: "不明な言語です。選択肢から1つ選択してください。",
		koKR: "알 수 없는 언어입니다. 다음 옵션 중 하나를 선택하세요.",
		pl:   "Nieznany język. Proszę wybrać 1 z opcji:",
		ptBR: "Idioma desconhecido. Escolha 1 das opções:",
		ptPT: "Idioma desconhecido. Escolha 1 das opções:",
		ru:   "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
		tr:   "Bilinmeyen dil. Lütfen seçeneklerden 1&#39;i seçin:",
		ukUA: "Невідома мова. Виберіть один із варіантів:",
		uz:   "Noma&#39;lum til. Iltimos, variantlardan 1 tasini tanlang:",
		zhCN: "未知语言。请从以下选项中选择 1 个："},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		arEG: "طرف غير معروف. يُرجى اختيار <b>&quot;إضافة جديد&quot;</b> إذا كانت جهة اتصال جديدة.",
		de:   "Unbekannte Gegnerpartei. Bitte wähle aus der Liste oder wähle <b>neuer Kontakt</b>.",
		en:   "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		es:   "Contacto desconocido. Por favor elige <b>Añadir</b> si es un contacto nuevo.",
		faIR: "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
		fr:   "Contrepartie inconnue. Veuillez choisir <b>« Ajouter »</b> s&#39;il s&#39;agit d&#39;un nouveau contact.",
		id:   "Rekanan tidak dikenal. Silakan pilih <b>Tambah baru</b> jika kontak tersebut baru.",
		it:   "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
		jaJP: "相手先が不明です。新しい連絡先の場合は、 <b>「新規追加」</b>を選択してください。",
		koKR: "알 수 없는 상대방입니다. 새 연락처인 경우 <b>&#39;새로 추가&#39;를</b> 선택하세요.",
		pl:   "Nieznany kontrahent. Wybierz <b>Dodaj nowy,</b> jeśli jest to nowy kontakt.",
		ptBR: "Contraparte desconhecida. Selecione <b>Adicionar novo</b> se for um novo contato.",
		ptPT: "Contraparte desconhecida. Selecione <b>Adicionar novo</b> se for um novo contacto.",
		ru:   "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
		tr:   "Bilinmeyen karşı taraf. Yeni bir kişi ise lütfen <b>Yeni ekle&#39;yi</b> seçin.",
		ukUA: "Невідомий контрагент. Будь ласка, виберіть <b>&quot;Додати нового&quot;</b> , якщо це новий контакт.",
		uz:   "Noma&#39;lum kontragent. Agar kontakt yangi boʻlsa, <b>“Yangi qoʻshish”ni</b> tanlang.",
		zhCN: "未知交易对手。如果是新联系人，请选择<b>“添加新</b>联系人”。"},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		arEG: "طرف غير معروف. يُرجى الاختيار من القائمة.",
		de:   "Unbekannte Gegnerpartei. Bitte wähle aus der Liste.",
		en:   "Unknown counterparty. Please choose from the list.",
		es:   "Contacto desconocido. Por favor elige de la lista.",
		faIR: "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
		fr:   "Contrepartie inconnue. Veuillez choisir dans la liste.",
		id:   "Rekanan tidak dikenal. Silakan pilih dari daftar.",
		it:   "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
		jaJP: "取引相手が不明です。リストから選択してください。",
		koKR: "알 수 없는 상대방입니다. 목록에서 선택하세요.",
		pl:   "Nieznany kontrahent. Proszę wybrać z listy.",
		ptBR: "Contraparte desconhecida. Selecione na lista.",
		ptPT: "Contraparte desconhecida. Selecione na lista.",
		ru:   "Неизвестный контакт. Пожалуйста выберите из списка.",
		tr:   "Bilinmeyen karşı taraf. Lütfen listeden seçin.",
		ukUA: "Невідомий контрагент. Виберіть зі списку.",
		uz:   "Noma&#39;lum kontragent. Iltimos, ro&#39;yxatdan tanlang.",
		zhCN: "未知交易对手。请从列表中选择。"},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		arEG: "دين غير معروف. الرجاء الاختيار من القائمة.",
		de:   "Unbekannter Schuldschein. Bitte wähle aus der Liste.",
		en:   "Unknown debt. Please choose from the list.",
		es:   "Deuda desconocida. Por favor elige de la lista.",
		faIR: "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
		fr:   "Dette inconnue. Veuillez choisir dans la liste.",
		id:   "Utang yang tidak diketahui. Silakan pilih dari daftar.",
		it:   "Debito sconosciuto. Scegli dalla lista qui sotto.",
		jaJP: "不明な債務です。リストから選択してください。",
		koKR: "알 수 없는 부채입니다. 목록에서 선택해 주세요.",
		pl:   "Nieznany dług. Proszę wybrać z listy.",
		ptBR: "Dívida desconhecida. Selecione na lista.",
		ptPT: "Dívida desconhecida. Selecione na lista.",
		ru:   "Неизвестный долг. Пожалуйста выберите из списка.",
		tr:   "Bilinmeyen borç. Lütfen listeden seçin.",
		ukUA: "Невідомий борг. Виберіть зі списку.",
		uz:   "Noma&#39;lum qarz. Iltimos, ro&#39;yxatdan tanlang.",
		zhCN: "未知债务。请从列表中选择。"},
	MESSAGE_TEXT_BILL_CARD_HEADER: {
		arEG: `<b>الفاتورة</b> : <code>%v</code> — %v`,
		de:   `<b>Rechnung</b>: <code>%v</code> — %v`,
		en:   `<b>Bill</b>: <code>%v</code> — %v`,
		es:   `<b>Cuenta</b>: <code>%v</code> — %v`,
		faIR: `<b>صورتحساب</b> : <code>%v</code> — %v`,
		fr:   `<b>Facture</b> : <code>%v</code> — %v`,
		id:   `<b>Tagihan</b> : <code>%v</code> — %v`,
		it:   `<b>Conto</b> : <code>%v</code> — %v`,
		jaJP: `<b>請求書</b>: <code>%v</code> — %v`,
		koKR: `<b>빌</b> : <code>%v</code> — %v`,
		pl:   `<b>Rachunek</b> : <code>%v</code> — %v`,
		ptBR: `<b>Conta</b> : <code>%v</code> — %v`,
		ptPT: `<b>Conta</b> : <code>%v</code> — %v`,
		ru:   `<b>Cчёт</b>: <code>%v</code> — %v`,
		tr:   `<b>Fatura</b> : <code>%v</code> — %v`,
		ukUA: `<b>Рахунок</b> : <code>%v</code> — %v`,
		uz:   `<b>Hisob</b> : <code>%v</code> — %v`,
		zhCN: `<b>账单</b>： <code>%v</code> — %v`},
	MESSAGE_TEXT_MEMBERS_TITLE: {
		arEG: "أعضاء",
		de:   "Mitglieder",
		en:   "Members",
		es:   "Miembros",
		faIR: "اعضا",
		fr:   "Membres",
		id:   "Anggota",
		it:   "Membri",
		jaJP: "メンバー",
		koKR: "멤버",
		pl:   "Członkowie",
		ptBR: "Membros",
		ptPT: "Membros",
		ru:   "Участники",
		tr:   "Üyeler",
		ukUA: "Учасники",
		uz:   "A'zolar",
		zhCN: "成员",
	},
	ALERT_TEXT_NOTHING_CHANGED: {
		arEG: "لم يتغير شيء",
		de:   "Nichts geändert",
		en:   "Nothing changed",
		es:   "Nada ha cambiado",
		faIR: "چیزی تغییر نکرده است",
		fr:   "Rien n&#39;a changé",
		id:   "Tidak ada yang berubah",
		it:   "Niente è cambiato",
		jaJP: "何も変わらなかった",
		koKR: "아무것도 바뀌지 않았다",
		pl:   "Nic się nie zmieniło",
		ptBR: "Nada mudou",
		ptPT: "Nada mudou",
		ru:   "Ничего не изменилось",
		tr:   "Hiçbir şey değişmedi",
		ukUA: "Нічого не змінилося",
		uz:   "Hech narsa o&#39;zgarmadi",
		zhCN: "什么都没改变"},
	ALERT_TEXT_YOU_ARE_ALREADY_MEMBER_OF_THE_GROUP: {
		arEG: "أنت بالفعل عضو في مجموعة تقسيم الفاتورة هذه.",
		de:   "Du bist schon Mitglied beim Teilen dieser Rechnung.",
		en:   "You are already a member of this bill splitting group.",
		es:   "Ya es miembro de este grupo de división de facturas.",
		faIR: "شما قبلا عضو این گروه تقسیم لایحه هستید.",
		fr:   "Vous êtes déjà membre de ce groupe de division de projet de loi.",
		id:   "Anda sudah menjadi anggota kelompok pemecah RUU ini.",
		it:   "Sei già membro di questo gruppo di divisione fatture.",
		jaJP: "あなたはすでにこの請求書分割グループのメンバーです。",
		koKR: "당신은 이미 이 청구서 분할 그룹의 회원입니다.",
		pl:   "Jesteś już członkiem tej grupy dzielącej rachunki.",
		ptBR: "Você já é membro deste grupo de divisão de contas.",
		ptPT: "Já é membro deste grupo de divisão de contas.",
		ru:   "Вы уже участник этой группы по совместной оплате счетов.",
		tr:   "Zaten bu fatura bölüştürme grubunun bir üyesisiniz.",
		ukUA: "Ви вже є членом цієї групи з розподілу законопроектів.",
		uz:   "Siz allaqachon ushbu hisob-kitoblarni ajratish guruhining aʼzosisiz.",
		zhCN: "您已经是该分摊账单小组的成员。"},
	MESSAGE_TEXT_YOUR_BILL_SPLITTING_GROUPS: {
		arEG: "مجموعات تقسيم الفواتير الخاصة بك",
		de:   "Gruppen, mit denen du Rechnungen teilst",
		en:   "Your bills splitting groups",
		es:   "Ya es miembro de este grupo de división de facturas.",
		faIR: "شما قبلا عضو این گروه تقسیم لایحه هستید.",
		fr:   "Vos groupes de partage de factures",
		id:   "Grup pembagian tagihan Anda",
		it:   "Sei già membro di questo gruppo di divisione fatture.",
		jaJP: "あなたの請求書分割グループ",
		koKR: "귀하의 청구서 분할 그룹",
		pl:   "Twoje grupy dzielenia rachunków",
		ptBR: "Seus grupos de divisão de contas",
		ptPT: "Seus grupos de divisão de contas",
		ru:   "Ваши группы совметсной оплаты",
		tr:   "Fatura bölme gruplarınız",
		ukUA: "Ваші групи спільної оплати",
		uz:   "Hisob-kitoblarni bo'lish guruhlari",
		zhCN: "您的账单分割组",
	},
	MESSAGE_TEXT_USE_ARROWS_TO_SELECT_GROUP: {
		arEG: "استخدم ⬅️ و ➡️ لتحديد المجموعة",
		de:   "Verwenden Sie ⬅️ und ➡️, um eine Gruppe auszuwählen",
		en:   "Use ⬅️ & ➡️ to select group",
		es:   "Usa ⬅️ y ➡️ para seleccionar el grupo",
		faIR: "برای انتخاب گروه از ⬅️ & ❢️ استفاده کنید",
		fr:   "Utilisez ⬅️ et ➡️ pour sélectionner le groupe",
		id:   "Gunakan ⬅️ &amp; ➡️ untuk memilih grup",
		it:   "Usare ⬅️ & ➡️ per selezionare il gruppo",
		jaJP: "⬅️と➡️を使ってグループを選択してください",
		koKR: "⬅️ &amp; ➡️를 사용하여 그룹을 선택하세요",
		pl:   "Użyj ⬅️ i ➡️ aby wybrać grupę",
		ptBR: "Use ⬅️ e ➡️ para selecionar o grupo",
		ptPT: "Utilize ⬅️ e ➡️ para selecionar o grupo",
		ru:   "Используйте ⬅️ и ➡️ чтобы выбрать группу.",
		tr:   "Grup seçmek için ⬅️ ve ➡️ kullanın",
		ukUA: "Використовуйте ⬅️ та ➡️, щоб вибрати групу",
		uz:   "Guruhni tanlash uchun ⬅️ va ➡️ tugmalaridan foydalaning",
		zhCN: "使用 ⬅️ &amp; ➡️ 选择组"},
	MESSAGE_TEXT_NO_GROUPS: {
		arEG: "أنت لست مشاركًا في أي مجموعة تقسيم الفاتورة.",
		de:   "Du gehörst zu keiner Gruppe, die sich Rechnungen teilt.",
		en:   "You are not a participant of any bill splitting group.",
		es:   "Usted no es participante de ningún grupo de división de facturas.",
		faIR: "شما شرکت کننده در هر گروه تقسیم لایحه نیستید.",
		fr:   "Vous ne participez à aucun groupe de partage de factures.",
		id:   "Anda bukan peserta kelompok pemecah RUU mana pun.",
		it:   "Non sei un partecipante a qualsiasi gruppo di divisione fatture.",
		jaJP: "あなたは請求書分割グループに参加していません。",
		koKR: "당신은 어떤 청구서 분할 집단에도 속하지 않습니다.",
		pl:   "Nie jesteś uczestnikiem żadnej grupy dzielącej rachunki.",
		ptBR: "Você não é participante de nenhum grupo de divisão de contas.",
		ptPT: "Não é participante de nenhum grupo de divisão de contas.",
		ru:   "Вы не состоите в группах совместной оплаты.",
		tr:   "Hiçbir fatura bölüştürme grubunun katılımcısı değilsiniz.",
		ukUA: "Ви не є учасником жодної групи з розподілу рахунків.",
		uz:   "Siz hisob-kitoblarni taqsimlash guruhining ishtirokchisi emassiz.",
		zhCN: "您不是任何分摊账单小组的参与者。"},
	MESSAGE_TEXT_USER_JOINED_GROUP: {
		arEG: `مرحبًا %v، لقد انضممت إلى مجموعة تقسيم الفاتورة هذه.`,
		de:   `Hi %v, du bist der Gruppe, die sich Rechnungen teilt, beigetreten.`,
		en:   `Hi %v, you joined this bill splitting group.`,
		es:   `Hola %v, te uniste a este grupo de división de facturas.`,
		faIR: "سلام %v، شما به گروه تقسیم این لایحه پیوستید",
		fr:   `Bonjour %v, vous avez rejoint ce groupe de partage de factures.`,
		id:   `Hai %v, Anda bergabung dengan kelompok pemecah tagihan ini.`,
		it:   "Hi %v, sei entrato in questo gruppo di divisione fatture.",
		jaJP: `こんにちは %v さん、この請求書分割グループに参加しました。`,
		koKR: `안녕하세요 %v, 당신은 이 청구서 분할 그룹에 가입했습니다.`,
		pl:   `Cześć %v, dołączyłeś do grupy dzielącej rachunki.`,
		ptBR: `Olá %v, você entrou neste grupo de divisão de contas.`,
		ptPT: `Olá %v, entrou neste grupo de divisão de contas.`,
		ru: `Привет %v, вы присоеденились к этой группе по совместной оплате счетов.
		`,
		tr:   `Merhaba %v, bu fatura bölüşme grubuna katıldın.`,
		ukUA: `Привіт, %v, ви приєдналися до цієї групи з розподілу рахунків.`,
		uz:   `Assalomu alaykum %v, siz hisob-kitoblarni ajratish guruhiga qoʻshildingiz.`,
		zhCN: `你好 %v，你加入了这个分摊账单小组。`},
	MESSAGE_TEXT_MEMBERS_CARD_TITLE: {
		arEG: "<b>أعضاء تقسيم الفواتير</b> (%d)",
		de:   "<b>Wer sich die Rechnung teilt</b> (%d)",
		en:   "<b>Bills splitting members</b> (%d)",
		es:   "<b>Proyectos de ley que dividen a los miembros</b> (%d)",
		faIR: "(%d) <b>نقض تقسیم اعضا</b>",
		fr:   "<b>Membres partageant les factures</b> (%d)",
		id:   "<b>Anggota yang membagi RUU</b> (%d)",
		it:   "<b>Membri di divisione delle bollette</b> (%d)",
		jaJP: "<b>請求書を分割するメンバー</b>(%d)",
		koKR: "<b>법안 분할 회원</b> (%d)",
		pl:   "<b>Rachunki dzielące członków</b> (%d)",
		ptBR: "<b>Projetos de lei que dividem membros</b> (%d)",
		ptPT: "<b>Projetos de lei que dividem os membros</b> (%d)",
		ru:   "<b>Участники совместных оплат</b> (%d)",
		tr:   "<b>Üyeleri bölen yasa tasarıları</b> (%d)",
		ukUA: "<b>Законопроекти, що розділяють членів</b> (%d)",
		uz:   "<b>A&#39;zolarni ajratuvchi hisoblar</b> (%d)",
		zhCN: "<b>法案拆分成员</b>(%d)"},
	MESSAGE_TEXT_SPLIT_MODE: {
		arEG: "<b>التقسيم</b> : %v",
		de:   "<b>Teilen</b>: %v",
		en:   "<b>Split</b>: %v",
		es:   "<b>División</b>: %v",
		faIR: "<b>شکاف</b>: %v",
		fr:   "<b>Répartition</b> : %v",
		id:   "<b>Membagi</b> : %v",
		it:   "<b>Diviso</b>: %v",
		jaJP: "<b>分割</b>: %v",
		koKR: "<b>분할</b> : %v",
		pl:   "<b>Podział</b> : %v",
		ptBR: "<b>Divisão</b> : %v",
		ptPT: "<b>Divisão</b> : %v",
		ru:   "<b>Делить</b>: %v",
		tr:   "<b>Bölme</b> : %v",
		ukUA: "<b>Розділити</b> : %v",
		uz:   "<b>Ajratish</b> : %v",
		zhCN: "<b>分割</b>：%v"},
	MESSAGE_TEXT_ASK_HOW_TO_SPLIT_IN_GROP: {
		arEG: "بأي نسبة تقوم بتقسيم الفواتير في هذه المجموعة؟",
		de:   "In welchem Verhältnis teilt ihr in dieser Gruppe eure Rechnungen?",
		en:   "In what proportion do you split bills in this group?",
		es:   "¿En qué proporción divide las facturas en este grupo?",
		faIR: "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟",
		fr:   "Dans quelle proportion partagez-vous les factures dans ce groupe ?",
		id:   "Berapa proporsi Anda membagi tagihan pada kelompok ini?",
		it:   "In quale percentuale dividi le fatture in questo gruppo?",
		jaJP: "このグループではどれくらいの割合で請求書を分割しますか?",
		koKR: "이 그룹에서는 얼마의 비율로 청구서를 나누시나요?",
		pl:   "W jakich proporcjach dzielisz rachunki w tej grupie?",
		ptBR: "Em que proporção você divide as contas neste grupo?",
		ptPT: "Em que proporção divide as contas neste grupo?",
		ru:   "В какой пропорции вы делите счета в этой группе?",
		tr:   "Bu grupta faturaları hangi oranda bölüyorsunuz?",
		ukUA: "У якій пропорції ви розподіляєте рахунки в цій групі?",
		uz:   "Ushbu guruhdagi veksellarni qancha nisbatda ajratasiz?",
		zhCN: "你们按什么比例分摊这组账单？"},
	MESSAGE_TEXT_MEMBERS_CARD_FOOTER: {
		arEG: "انقر فوق <code>Join</code> للمشاركة في تقسيم الفواتير.",
		de:   "Klick <code>Join</code>, um auch Rechnungen zu teilen.",
		en:   "Click <code>Join</code> to participate in bills splitting.",
		es:   "¿En qué proporción divide las facturas en este grupo?",
		faIR: "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟",
		fr:   "Cliquez sur <code>Join</code> pour participer au partage des factures.",
		id:   "Klik <code>Join</code> untuk berpartisipasi dalam pembagian tagihan.",
		it:   "In quale percentuale dividi le fatture in questo gruppo?",
		jaJP: "請求書の分割に参加するには、 <code>Join</code>クリックします。",
		koKR: "청구서 분할에 참여하려면 <code>Join</code> 클릭하세요.",
		pl:   "Kliknij <code>Join</code> , aby wziąć udział w podziale rachunków.",
		ptBR: "Clique em <code>Join</code> para participar da divisão de contas.",
		ptPT: "Clique em <code>Join</code> para aderir à divisão de contas.",
		ru:   "Жмите <code>Присоедениться</code> чтобы учавствовать.",
		tr:   "Fatura bölüşümüne katılmak için <code>Join</code> tıklayın.",
		ukUA: "Натисніть <code>Join</code> , щоб взяти участь у розподілі рахунків.",
		uz:   "To‘lovlarni taqsimlashda ishtirok etish uchun <code>Join</code> tugmasini bosing.",
		zhCN: "点击<code>Join</code>即可参与分摊账单。"},
	MESSAGE_TEXT_BILL_CARD_MEMBER_TITLE: {
		arEG: "{{.N}}. {{.MemberName}}",
		de:   "{{.N}}. {{.MemberName}}",
		en:   "{{.N}}. {{.MemberName}}",
		es:   "{{.N}}. {{.MemberName}}",
		faIR: "{{.N}}. {{.MemberName}}",
		fr:   "{{.N}}. {{.MemberName}}",
		id:   "{{.N}}. {{.MemberName}}",
		it:   "{{.N}}. {{.MemberName}}",
		jaJP: "{{.N}}。{{.MemberName}}",
		koKR: "{{.N}}. {{.MemberName}}",
		pl:   "{{.N}}. {{.MemberName}}",
		ptBR: "{{.N}}. {{.MemberName}}",
		ptPT: "{{.N}}. {{.MemberName}}",
		ru:   "{{.N}}. {{.MemberName}}",
		tr:   "{{.N}}. {{.MemberName}}",
		ukUA: "{{.N}}. {{.MemberName}}",
		uz:   "{{.N}}. {{.MemberName}}",
		zhCN: "{{.N}}. {{.MemberName}}"},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW: {
		arEG: "<i>{{.Percent}}%</i>",
		de:   "<i>{{.Percent}}%</i>",
		en:   "<i>{{.Percent}}%</i>",
		es:   "<i>{{.Percent}}%</i>",
		faIR: "<i>{{.Percent}}%</i>",
		fr:   "<i>{{.Percent}}%</i>",
		id:   "<i>{{.Percent}}%</i>",
		it:   "<i>{{.Percent}}%</i>",
		jaJP: "<i>{{.Percent}}%</i>",
		koKR: "<i>{{.Percent}}%</i>",
		pl:   "<i>{{.Percent}}%</i>",
		ptBR: "<i>{{.Percent}}%</i>",
		ptPT: "<i>{{.Percent}}%</i>",
		ru:   "<i>{{.Percent}}%</i>",
		tr:   "<i>{{.Percent}}%</i>",
		ukUA: "<i>{{.Percent}}%</i>",
		uz:   "<i>{{.Percent}}%</i>",
		zhCN: "<i>{{.Percent}}%</i>"},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_OWES: {
		arEG: "\n <i>يدين {{.Owes}}</i>",
		de:   "\n   <i>schuldet {{.Owes}}</i>",
		en:   "\n   <i>owes {{.Owes}}</i>",
		es:   "\n   <i>debo {{.Owes}}</i>",
		faIR: "\n   <i>بدهکار است {{.Owes}}</i>",
		fr:   "\n <i>doit {{.Owes}}</i>",
		id:   "\n <i>berutang {{.Owes}}</i>",
		it:   "\n <i>deve {{.Owes}}</i>",
		jaJP: "\n は<i>{{.Owes}} を借りている</i>",
		koKR: "\n <i>{{.Owes}}을(를) 빚지고 있습니다</i>",
		pl:   "\n <i>jest winien {{.Owes}}</i>",
		ptBR: "\n <i>deve {{.Owes}}</i>",
		ptPT: "\n <i>deve {{.Owes}}</i>",
		ru:   "\n   <i>должен {{.Owes}}</i>",
		tr:   "\n <i>{{.Owes}} borçlu</i>",
		ukUA: "\n <i>винен {{.Owes}}</i>",
		uz:   "\n <i>{{.Owes}} qarz</i>",
		zhCN: "\n<i>欠 {{.Owes}}</i>"},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PAID: {
		arEG: "\n <i>مدفوع {{.Paid}}</i>",
		de:   "\n   <i>bezahlte {{.Paid}}</i>",
		en:   "\n   <i>paid {{.Paid}}</i>",
		es:   "\n   <i>he pagado {{.Paid}}</i>",
		faIR: "\n   <i>پرداخت شده {{.Paid}}</i>",
		fr:   "\n <i>payé {{.Paid}}</i>",
		id:   "\n <i>dibayar {{.Paid}}</i>",
		it:   "\n <i>pagato {{.Paid}}</i>",
		jaJP: "\n<i>支払済 {{.Paid}}</i>",
		koKR: "\n <i>{{.Paid}}을(를) 지불했습니다.</i>",
		pl:   "\n <i>zapłacono {{.Paid}}</i>",
		ptBR: "\n <i>pago {{.Paid}}</i>",
		ptPT: "\n <i>pago {{.Paid}}</i>",
		ru:   "\n   <i>заплатил {{.Paid}}</i>",
		tr:   "\n <i>ödendi {{.Paid}}</i>",
		ukUA: "\n <i>сплачено {{.Paid}}</i>",
		uz:   "\n <i>{{.Paid}} toʻlangan</i>",
		zhCN: "\n<i>已付款 {{.Paid}}</i>"},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PART_PAID: {
		arEG: "\n <i>مدفوع {{.Paid}}، مدين {{.Owes}}</i>",
		de:   "\n<i>bezahlte {{.Paid}}, schuldet noch {{.Owes}}</i>",
		en:   "\n<i>paid {{.Paid}}, owes {{.Owes}}</i>",
		es:   "\n<i>he pagado {{.Paid}}, debo {{.Owes}}</i>",
		faIR: "\n<i>پرداخت شده {{.Paid}}, بدهکار است {{.Owes}}</i>",
		fr:   "\n <i>payé {{.Paid}}, doit {{.Owes}}</i>",
		id:   "\n <i>membayar {{.Paid}}, berutang {{.Owes}}</i>",
		it:   "\n <i>pagato {{.Paid}}, deve {{.Owes}}</i>",
		jaJP: "\n<i>支払済 {{.Paid}}、未払い {{.Owes}}</i>",
		koKR: "\n <i>{{.Paid}}을(를) 지불했고, {{.Owes}}을(를) 빚졌습니다.</i>",
		pl:   "\n <i>zapłacił {{.Paid}}, jest winien {{.Owes}}</i>",
		ptBR: "\n <i>pagou {{.Paid}}, deve {{.Owes}}</i>",
		ptPT: "\n <i>pagou {{.Paid}}, deve {{.Owes}}</i>",
		ru:   "\n<i>заплатил {{.Paid}}, должен {{.Owes}}</i>",
		tr:   "\n <i>{{.Paid}} ödedi, {{.Owes}} borçlu</i>",
		ukUA: "\n <i>сплатив {{.Paid}}, заборгував {{.Owes}}</i>",
		uz:   "\n <i>{{.Paid}} toʻlangan, qarzi {{.Owes}}</i>",
		zhCN: "\n<i>已付款 {{.Paid}}，欠款 {{.Owes}}</i>"},
	MESSAGE_TEXT_BILL_ASK_WHO_PAID: {
		arEG: "الرجاء اختيار من قام بدفع الفاتورة:",
		de:   "Bitte wähle, wer die Rechnung gezahlt hat:",
		en:   "Please choose who paid for the bill:",
		es:   "Por favor, elige quien ha pagado la cuenta:",
		faIR: "لطفا انتخاب کنید که چه کسانی برای این لایحه پرداخت کرده اند:",
		fr:   "Veuillez choisir qui a payé la facture :",
		id:   "Silakan pilih siapa yang membayar tagihan:",
		it:   "Scegliere chi ha pagato la fattura:",
		jaJP: "請求書の支払い者を選択してください:",
		koKR: "누가 요금을 지불했는지 선택해 주세요:",
		pl:   "Proszę wybrać osobę, która zapłaciła rachunek:",
		ptBR: "Por favor, escolha quem pagou a conta:",
		ptPT: "Por favor, escolha quem pagou a conta:",
		ru:   "Пожалуйста выберите кто заплатил по счёту:",
		tr:   "Lütfen hesabı kimin ödediğini seçin:",
		ukUA: "Будь ласка, виберіть, хто оплатив рахунок:",
		uz:   "Iltimos, hisobni kim to&#39;laganini tanlang:",
		zhCN: "请选择账单付款人："},
	MESSAGE_TEXT_STATUS: {
		arEG: "الحالة: %v",
		de:   "Status: %v",
		en:   "Status: %v",
		es:   "Estado: %v",
		faIR: "وضعیت: %v",
		fr:   "Statut : %v",
		id:   "Keadaan: %v",
		it:   "Stato: %v",
		jaJP: "ステータス: %v",
		koKR: "상태: %v",
		pl:   "Stan: %v",
		ptBR: "Estado: %v",
		ptPT: "Estado: %v",
		ru:   "Статус: %v",
		tr:   "Durum: %v",
		ukUA: "Статус: %v",
		uz:   "Holati: %v",
		zhCN: "状态：%v"},
	BUTTON_TEXT_ADD_MEMBER: {
		arEG: "إضافة مشارك",
		de:   "Partei hinzufügen",
		en:   "Add participant",
		es:   "Añadir participante",
		faIR: "افزودن مشارکت کننده",
		fr:   "Ajouter un participant",
		id:   "Tambahkan peserta",
		it:   "Aggiungi partecipante",
		jaJP: "参加者を追加",
		koKR: "참가자 추가",
		pl:   "Dodaj uczestnika",
		ptBR: "Adicionar participante",
		ptPT: "Adicionar participante",
		ru:   "Добавить участника",
		tr:   "Katılımcı ekle",
		ukUA: "Додати учасника",
		uz:   "Ishtirokchi qo&#39;shing",
		zhCN: "添加参与者"},
	BUTTON_TEXT_FINALIZE_BILL: {
		arEG: "🔓 قفل الفاتورة",
		de:   "🔓 Rechnung abschließen",
		en:   "🔓 Lock the bill",
		es:   "🔓 Cerrar la cuenta",
		faIR: "🔓 لایحه را قفل کنید",
		fr:   "🔓 Verrouillez la facture",
		id:   "🔓 Kunci tagihan",
		it:   "🔓 Bloccare il conto",
		jaJP: "🔓 法案をロックする",
		koKR: "🔓 지폐를 잠그세요",
		pl:   "🔓 Zablokuj rachunek",
		ptBR: "🔓 Bloqueie a conta",
		ptPT: "🔓 Bloqueie a conta",
		ru:   "🔓 Закрыть счёт",
		tr:   "🔓 Faturayı kilitle",
		ukUA: "🔓 Заблокувати рахунок",
		uz:   "🔓 Hisobni qulflash",
		zhCN: "🔓 锁定账单"},
	ButtonEdit: {
		arEG: "✏️ تعديل",
		de:   "✏️ Bearbeiten",
		en:   "✏️ Edit",
		es:   "✏️ Editar",
		faIR: "✏️ ویرایش",
		fr:   "✏️ Modifier",
		id:   "✏️ Sunting",
		it:   "✏️ Modifica",
		jaJP: "✏️ 編集",
		koKR: "✏️ 편집",
		pl:   "✏️ Edytuj",
		ptBR: "✏️ Editar",
		ptPT: "✏️ Editar",
		ru:   "✏️ Редактировать",
		tr:   "✏️ Düzenle",
		ukUA: "✏️ Редагувати",
		uz:   "✏️ Tahrirlash",
		zhCN: "✏️ 编辑",
	},
	BUTTON_TEXT_EDIT_BILL: {
		arEG: "✏️ تعديل",
		de:   "✏️ Bearbeiten",
		en:   "✏️ Edit",
		es:   "✏️ Editar",
		faIR: "✏️ ویرایش",
		fr:   "✏️ Modifier",
		id:   "✏️ Sunting",
		it:   "✏️ Modifica",
		jaJP: "✏️ 編集",
		koKR: "✏️ 편집",
		pl:   "✏️ Edytuj",
		ptBR: "✏️ Editar",
		ptPT: "✏️ Editar",
		ru:   "✏️ Редактировать",
		tr:   "✏️ Düzenle",
		ukUA: "✏️ Редагувати",
		uz:   "✏️ Tahrirlash",
		zhCN: "✏️ 编辑"},
	BUTTON_TEXT_SPLIT_MODE: {
		arEG: "➗ التقسيم: %v",
		de:   "➗ Teilen: %v",
		en:   "➗ Split: %v",
		es:   "➗ Dividir: %v",
		faIR: "➗ تقسیم: %v",
		fr:   "➗ Répartition : %v",
		id:   "➗ Membagi: %v",
		it:   "➗ Dividi: %v",
		jaJP: "➗ 分割: %v",
		koKR: "➗ 분할: %v",
		pl:   "➗ Podział: %v",
		ptBR: "➗ Divisão: %v",
		ptPT: "➗ Divisão: %v",
		ru:   "➗ Делить: %v",
		tr:   "➗ Bölme: %v",
		ukUA: "➗ Розділити: %v",
		uz:   "➗ Ajratish: %v",
		zhCN: "➗ 拆分：%v"},
	MESSAGE_TEXT_SPLIT_LABEL_WITH_VALUE: {
		arEG: "التقسيم: %v",
		de:   "Teilen: %v",
		en:   "Split: %v",
		es:   "Dividir: %v",
		faIR: "تقسیم: %v",
		fr:   "Répartition : %v",
		id:   "Membagi: %v",
		it:   "Dividi: %v",
		jaJP: "分割: %v",
		koKR: "분할: %v",
		pl:   "Podział: %v",
		ptBR: "Dividir: %v",
		ptPT: "Dividir: %v",
		ru:   "Делить: %v",
		tr:   "Bölme: %v",
		ukUA: "Розділення: %v",
		uz:   "Ajratish: %v",
		zhCN: "拆分：%v"},
	STATUS_DRAFT: {
		arEG: "مسودة",
		de:   "Entwurf",
		en:   "draft",
		es:   "borrador",
		faIR: "پیش نویس",
		fr:   "brouillon",
		id:   "draf",
		it:   "bozza",
		jaJP: "下書き",
		koKR: "초안",
		pl:   "projekt",
		ptBR: "rascunho",
		ptPT: "rascunho",
		ru:   "черновик",
		tr:   "taslak",
		ukUA: "чернетка",
		uz:   "qoralama",
		zhCN: "草稿"},
	SPLIT_MODE_EQUALLY: {
		arEG: "بالتساوي",
		de:   "Gleichverteilt",
		en:   "Equally",
		es:   "A partes iguales",
		faIR: "به همان اندازه",
		fr:   "Également",
		id:   "Sama",
		it:   "Ugualmente",
		jaJP: "同じように",
		koKR: "같이",
		pl:   "Równie",
		ptBR: "Igualmente",
		ptPT: "Igualmente",
		ru:   "Поровну",
		tr:   "Eşit olarak",
		ukUA: "однаково",
		uz:   "Teng",
		zhCN: "同样"},
	SPLIT_MODE_PERCENTAGE: {
		arEG: "نسبة مئوية",
		de:   "Prozentual",
		en:   "Percentage",
		es:   "Porcentaje",
		faIR: "درصد",
		fr:   "Pourcentage",
		id:   "Persentase",
		it:   "Percentuale",
		jaJP: "パーセンテージ",
		koKR: "백분율",
		pl:   "Procent",
		ptBR: "Percentagem",
		ptPT: "Percentagem",
		ru:   "В процентах",
		tr:   "Yüzde",
		ukUA: "Відсоток",
		uz:   "Foiz",
		zhCN: "百分比"},
	SPLIT_MODE_EXACT_AMOUNT: {
		arEG: "المبالغ الدقيقة",
		de:   "Exakte Summen",
		en:   "Exact amounts",
		es:   "Importes exactos",
		faIR: "مقادیر دقیق",
		fr:   "Montants exacts",
		id:   "Jumlah yang tepat",
		it:   "Quantità esatte",
		jaJP: "正確な金額",
		koKR: "정확한 금액",
		pl:   "Dokładne kwoty",
		ptBR: "Quantidades exatas",
		ptPT: "Quantidades exatas",
		ru:   "Точные суммы",
		tr:   "Tam miktarlar",
		ukUA: "Точні суми",
		uz:   "Aniq miqdorlar",
		zhCN: "确切金额"},
	SPLIT_MODE_SHARES: {
		arEG: "الأسهم",
		de:   "Teilen",
		en:   "Shares",
		es:   "Proporciones",
		faIR: "سهام",
		fr:   "Actions",
		id:   "Saham",
		it:   "Azioni",
		jaJP: "株式",
		koKR: "주식",
		pl:   "Akcje",
		ptBR: "Ações",
		ptPT: "Ações",
		ru:   "В долях",
		tr:   "Paylaşımlar",
		ukUA: "Акції",
		uz:   "Aksiyalar",
		zhCN: "分享"},
	BUTTON_TEXT_JOIN: {
		arEG: "➕ انضم",
		de:   "➕ Beitreten",
		en:   "➕ Join",
		es:   "➕ Adherirse",
		faIR: "➕ عضویت",
		fr:   "➕ Rejoindre",
		id:   "➕ Bergabung",
		it:   "➕ Unisciti",
		jaJP: "➕ 参加する",
		koKR: "➕ 가입하기",
		pl:   "➕ Dołącz",
		ptBR: "➕ Junte-se",
		ptPT: "➕ Junte-se",
		ru:   "➕ Присоедениться",
		tr:   "➕ Katıl",
		ukUA: "➕ Приєднатися",
		uz:   "➕ Qo&#39;shiling",
		zhCN: "➕ 加入"},
	BUTTON_TEXT_LEAVE: {
		arEG: "يترك",
		de:   "Verlassen",
		en:   "Leave",
		es:   "Salir",
		faIR: "ترک کردن",
		fr:   "Partir",
		id:   "Meninggalkan",
		it:   "Partire",
		jaJP: "離れる",
		koKR: "떠나다",
		pl:   "Wyjechać",
		ptBR: "Deixar",
		ptPT: "Ausência",
		ru:   "Покинуть",
		tr:   "Ayrılmak",
		ukUA: "Залишити",
		uz:   "Ketish",
		zhCN: "离开"},
	BUTTON_TEXT_DUE: {
		arEG: "📅 المستحق: %v",
		de:   "📅 Fällig: %v",
		en:   "📅 Due: %v",
		es:   "📅 Hasta: %v",
		faIR: "📅 مورد: %v",
		fr:   "📅 Échéance: %v",
		id:   "📅 Jatuh tempo: %v",
		it:   "📅 Dovuto: %v",
		jaJP: "📅 期限: %v",
		koKR: "📅 기한: %v",
		pl:   "📅 Termin: %v",
		ptBR: "📅 Vencimento: %v",
		ptPT: "📅 Vencimento: %v",
		ru:   "📅 До: %v",
		tr:   "📅 Vade: %v",
		ukUA: "📅 До: %v",
		uz:   "📅 Muddati: %v",
		zhCN: "📅 到期: %v",
	},
	NOT_SET: {
		arEG: "غير محدد",
		de:   "nicht gesetzt",
		en:   "not set",
		es:   "no establecido",
		faIR: "تنظیم نشده",
		fr:   "non défini",
		id:   "tidak diatur",
		it:   "non impostato",
		jaJP: "設定されていません",
		koKR: "설정되지 않음",
		pl:   "nie ustawiono",
		ptBR: "não definido",
		ptPT: "não definido",
		ru:   "не задано",
		tr:   "ayarlanmamış",
		ukUA: "не встановлено",
		uz:   "sozlanmagan",
		zhCN: "未设置"},
	BUTTON_TEXT_MANAGE_MEMBERS: {
		arEG: "👫 المشاركين",
		de:   "👫 Parteien",
		en:   "👫 Participants",
		es:   "👫 Participantes",
		faIR: "👫 شرکت کنندگان",
		fr:   "👫 Participants",
		id:   "👫 Peserta",
		it:   "👫 Partecipanti",
		jaJP: "👫 参加者",
		koKR: "👫 참가자",
		pl:   "👫 Uczestnicy",
		ptBR: "👫 Participantes",
		ptPT: "👫 Participantes",
		ru:   "👫 Участники",
		tr:   "👫 Katılımcılar",
		ukUA: "👫 Учасники",
		uz:   "👫 Ishtirokchilar",
		zhCN: "👫 参与者"},
	BUTTON_TEXT_CHANGE_BILL_PAYER: {
		arEG: "🔀 تغيير الدافع",
		de:   "🔀 Bezahler ändern",
		en:   "🔀 Change payer",
		es:   "🔀 Cambiar el pagador",
		faIR: "🔀 تغییر پرداخت کننده",
		fr:   "🔀 Changer de payeur",
		id:   "🔀 Ubah pembayar",
		it:   "🔀 Cambia il pagatore",
		jaJP: "🔀 支払人を変更する",
		koKR: "🔀 지불자 변경",
		pl:   "🔀 Zmień płatnika",
		ptBR: "🔀 Alterar pagador",
		ptPT: "🔀 Alterar pagador",
		ru:   "🔀 Сменить плательщика",
		tr:   "🔀 Ödeyen kişiyi değiştir",
		ukUA: "🔀 Змінити платника",
		uz:   "🔀 To&#39;lovchini o&#39;zgartirish",
		zhCN: "🔀 更改付款人"},
	COMMAND_TEXT_I_PAID: {
		arEG: "لقد دفعت",
		de:   "Ich habe bezahlt",
		en:   "I paid",
		es:   "he pagado",
		faIR: "پرداخت کردم",
		fr:   "J&#39;ai payé",
		id:   "aku sudah membayar",
		it:   "Ho pagato",
		jaJP: "支払いました",
		koKR: "나는 지불했다",
		pl:   "Zapłaciłem",
		ptBR: "eu paguei",
		ptPT: "eu paguei",
		ru:   "Я заплатил",
		tr:   "ben ödedim",
		ukUA: "Я заплатив",
		uz:   "Men to&#39;ladim",
		zhCN: "我付了"},
	COMMAND_TEXT_I_OWE: {
		arEG: "انا مدين",
		de:   "Ich schulde",
		en:   "I owe",
		es:   "Yo debo",
		faIR: "من بدهکارم",
		fr:   "Je dois",
		id:   "aku berhutang",
		it:   "Sono in debito",
		jaJP: "私は借りがある",
		koKR: "나는 빚을 졌다",
		pl:   "Jestem winien",
		ptBR: "Eu devo",
		ptPT: "Eu devo",
		ru:   "Я должен",
		tr:   "Ben borçluyum",
		ukUA: "Я винен",
		uz:   "qarzdorman",
		zhCN: "我欠"},
	COMMAND_TEXT_OWED_TO_ME: {
		arEG: "مدين لي",
		de:   "schuldet mir",
		en:   "Owed to me",
		es:   "Me deben",
		faIR: "به من تعلق دارد",
		fr:   "C&#39;est à moi que je dois",
		id:   "Berutang padaku",
		it:   "È dovuto a me",
		jaJP: "私に借りがある",
		koKR: "나에게 빚진 것",
		pl:   "To mi się należało",
		ptBR: "Devido a mim",
		ptPT: "Devido a mim",
		ru:   "Должны мне",
		tr:   "Bana borçlu",
		ukUA: "Заборгував мені",
		uz:   "Menga qarzdor",
		zhCN: "欠我的"},
	MESSAGE_TEXT_BILL_HEADER: {
		arEG: "الفاتورة: %v",
		de:   "Rechnung: %v",
		en:   "Bill: %v",
		es:   "Cuenta: %v",
		faIR: "بیل :%v",
		fr:   "Facture : %v",
		id:   "Tagihan: %v",
		it:   "Conto: %v",
		jaJP: "請求書: %v",
		koKR: "청구서: %v",
		pl:   "Rachunek: %v",
		ptBR: "Conta: %v",
		ptPT: "Conta: %v",
		ru:   "Cчёт: %v",
		tr:   "Fatura: %v",
		ukUA: "Рахунок: %v",
		uz:   "Hisob: %v",
		zhCN: "账单：%v"},
	MESSAGE_TEXT_NEW_DEBT_HEADER: {
		arEG: "الفاتورة: %v",
		de:   "Rechnung: %v",
		en:   "Bill: %v",
		es:   "Cuenta: %v",
		faIR: "بیل: %v",
		fr:   "Facture : %v",
		id:   "Tagihan: %v",
		it:   "Conto: %v",
		jaJP: "請求書: %v",
		koKR: "청구서: %v",
		pl:   "Rachunek: %v",
		ptBR: "Conta: %v",
		ptPT: "Conta: %v",
		ru:   "Cчёт: %v",
		tr:   "Fatura: %v",
		ukUA: "Рахунок: %v",
		uz:   "Hisob: %v",
		zhCN: "账单：%v"},
	MESSAGE_TEXT_GROUPS_ONLY_COMMAND: {
		arEG: "هذا الأمر متاح فقط في الدردشات الجماعية في الوقت الحالي.",
		de:   "Dieser Befehl ist derzeit nur in Gruppenchats verfügbar.",
		en:   "This command is available in group chats only for now.",
		es:   "Este comando está disponible solo en chats grupales por ahora.",
		faIR://nolint:staticcheck // disable ST1018 for this line
		"این دستور فعلاً فقط در چت‌های گروهی در دسترس است.",
		fr:   "Cette commande est disponible uniquement dans les discussions de groupe pour le moment.",
		id:   "Perintah ini hanya tersedia di obrolan grup untuk saat ini.",
		it:   "Per ora questo comando è disponibile solo nelle chat di gruppo.",
		jaJP: "このコマンドは現時点ではグループチャットでのみ使用できます。",
		koKR: "이 명령은 현재 그룹 채팅에서만 사용할 수 있습니다.",
		pl:   "Ta komenda jest na razie dostępna tylko na czatach grupowych.",
		ptBR: "Este comando está disponível apenas em bate-papos em grupo por enquanto.",
		ptPT: "Este comando só está disponível em chats de grupo por enquanto.",
		ru:   "Эта команда пока что доступна только в групповых чатах",
		tr:   "Bu komut şimdilik sadece grup sohbetlerinde kullanılabiliyor.",
		ukUA: "Ця команда поки що доступна лише в групових чатах.",
		uz:   "Bu buyruq hozircha faqat guruh chatlarida mavjud.",
		zhCN: "此命令目前仅在群聊中可用。"},
	MESSAGE_TEXT_ALREADY_HAS_CONTACT_WITH_SUCH_NAME: {
		arEG: "لديك بالفعل جهة اتصال بالاسم: %v",
		de:   "Sie haben bereits einen Kontakt mit dem Namen: %v",
		en:   "You already have contact with name: %v",
		es:   "Ya tienes contacto con nombre: %v",
		faIR: "شما از قبل با نام: %v تماس داشته\u200cاید.",
		fr:   "Vous avez déjà un contact avec le nom : %v",
		id:   "Anda sudah memiliki kontak dengan nama: %v",
		it:   "Hai già un contatto con il nome: %v",
		jaJP: "すでに名前: %v で連絡を取っています",
		koKR: "%v라는 이름과 이미 연락처가 있습니다.",
		pl:   "Masz już kontakt o nazwie: %v",
		ptBR: "Você já tem contato com nome: %v",
		ptPT: "Já tem contacto com nome: %v",
		ru:   "У вас уже есть контакт с таким именем: %v",
		tr:   "Zaten %v ismiyle bir iletişiminiz var",
		ukUA: "У вас вже є контакт з іменем: %v",
		uz:   "Sizda allaqachon ismingiz bor: %v",
		zhCN: "您已经有一个联系人，名字为：%v"},
	MESSAGE_TEXT_ALREADY_BILL_MEMBER: {
		arEG: "%v، أنت تشارك هذه الفاتورة بالفعل.",
		de:   "%v, du teilst diese Rechnung bereits.",
		en:   "%v, you are sharing this bill already.",
		es:   "%v, estás compartiendo esta cuenta ya.",
		faIR: "%v، شما قبلا این لایحه را به اشتراک می گذارید",
		fr:   "%v, vous partagez déjà cette facture.",
		id:   "%v, Anda sudah berbagi tagihan ini.",
		it:   "%v, stai già condividendo questo disegno di legge.",
		jaJP: "%v さん、この請求書はすでに共有されています。",
		koKR: "%v님, 이미 이 청구서를 공유하고 계십니다.",
		pl:   "%v, udostępniasz już ten rachunek.",
		ptBR: "%v, você já está dividindo essa conta.",
		ptPT: "%v, já está a dividir essa conta.",
		ru:   "%v, вы уже входите в состав участников.",
		tr:   "%v, bu yasa tasarısını zaten paylaşıyorsunuz.",
		ukUA: "%v, ви вже ділитеся цим рахунком.",
		uz:   "%v, siz allaqachon bu hisobni baham ko&#39;rasiz.",
		zhCN: "%v，您已经共享此账单。"},
	MESSAGE_TEXT_USER_JOINED_BILL: {
		arEG: "تم انضمام %v إلى مشاركة الفاتورة.",
		de:   "%v ist dem Teilen der Rechnung beigetreten.",
		en:   "%v joined to bill sharing.",
		es:   "%v pagar conjuntamente.",
		faIR: "%v به اشتراک گذاری لایحه پیوست.",
		fr:   "%v a rejoint le partage de factures.",
		id:   "%v bergabung dalam pembagian tagihan.",
		it:   "%v unito alla condivisione di fatture.",
		jaJP: "%v が請求書共有に参加しました。",
		koKR: "%v가 청구서 공유에 참여했습니다.",
		pl:   "%v dołączył do współdzielenia rachunków.",
		ptBR: "%v entrou para o compartilhamento de contas.",
		ptPT: "%v entrou para a partilha de contas.",
		ru:   "%v присоеденился к совместной оплате.",
		tr:   "%v fatura paylaşımına katıldı.",
		ukUA: "%v приєднався до спільного використання рахунку.",
		uz:   "%v hisob-kitoblarni taqsimlashga qo‘shildi.",
		zhCN: "%v 加入了账单分摊。"},
	BUTTON_TEXT_I_PAID_FOR_THE_BILL: {
		arEG: "لقد تم دفع الفاتورة من قبلي.",
		de:   "Die Rechnung wurde von mir bezahlt.",
		en:   "The bill was paid by me.",
		es:   "La factura fue pagada por mí.",
		faIR: "این لایحه توسط من پرداخت شد",
		fr:   "La facture a été payée par moi.",
		id:   "Tagihan dibayar oleh saya.",
		it:   "Il conto è stato pagato da me.",
		jaJP: "請求書は私が支払いました。",
		koKR: "청구서는 내가 지불했습니다.",
		pl:   "Rachunek został zapłacony przeze mnie.",
		ptBR: "A conta foi paga por mim.",
		ptPT: "A conta foi paga por mim.",
		ru:   "Этот счёт оплатил(а) я.",
		tr:   "Fatura benim tarafımdan ödendi.",
		ukUA: "Цей рахунок оплатив(ла) я.",
		uz:   "Hisob men tomonimdan to'landi.",
		zhCN: "账单由我支付。",
	},
	BUTTON_TEXT_I_OWE_FOR_THE_BILL: {
		arEG: "أنا مدين بهذه الفاتورة",
		de:   "Ich muss noch was dabeigeben",
		en:   "I owe for this bill",
		es:   "Debo esta factura",
		faIR: "من برای این لایحه بدهکار هستم",
		fr:   "Je dois pour cette facture",
		id:   "Saya berhutang untuk tagihan ini",
		it:   "Devo per questo disegno di legge",
		jaJP: "この請求書の支払いが必要です",
		koKR: "이 청구서에 대한 금액을 지불해야 합니다",
		pl:   "Jestem winien za ten rachunek",
		ptBR: "Eu devo por esta conta",
		ptPT: "Eu devo por esta conta",
		ru:   "Я должен по этому счёту",
		tr:   "Bu fatura için borçluyum",
		ukUA: "Я винен за цим рахунком",
		uz:   "Men bu hisob uchun qarzdorman",
		zhCN: "我欠这个账单",
	},
	BUTTON_TEXT_I_DO_NOT_SHARE_THIS_BILL: {
		arEG: "أنا لا أشارك هذا المشروع",
		de:   "Ich habe damit nichts zutun",
		en:   "I don't share this bill",
		es:   "No comparto esta cuenta",
		faIR: "من این لایحه را به اشتراک نمی گذارم",
		fr:   "Je ne partage pas cette facture",
		id:   "Saya tidak berbagi tagihan ini",
		it:   "Non condivido questo disegno di legge",
		jaJP: "この請求書を共有していません",
		koKR: "이 청구서를 공유하지 않습니다",
		pl:   "Nie dzielę się tym rachunkiem",
		ptBR: "Eu não compartilho esta conta",
		ptPT: "Eu não compartilho esta conta",
		ru:   "Я не учавствую в этой покупке",
		tr:   "Bu faturayı paylaşmıyorum",
		ukUA: "Я не беру участі в цій покупці",
		uz:   "Men bu hisobni baham ko'rmayman",
		zhCN: "我不分享这个账单",
	},
	MESSAGE_TEXT_YOU_JOINED_BILL: {
		arEG: "لقد انضممت إلى مشاركة الفاتورة.",
		de:   "Du bist dem Teilen der Rechnung beigetreten.",
		en:   "You've joined to bill sharing.",
		es:   "Te has agregado para pagar conjuntamente .",
		faIR: "شما به اشتراک گذاشتن لایحه پیوستید",
		fr:   "Vous avez rejoint le partage de facture.",
		id:   "Anda telah bergabung dengan berbagi tagihan.",
		it:   "Sei entrato a far parte della fatturazione.",
		jaJP: "請求書の共有に参加しました。",
		koKR: "청구서 공유에 참여하셨습니다.",
		pl:   "Dołączyłeś do dzielenia rachunku.",
		ptBR: "Você se juntou ao compartilhamento de contas.",
		ptPT: "Você se juntou ao compartilhamento de contas.",
		ru:   "Вы присоеденились к совместной оплате.",
		tr:   "Fatura paylaşımına katıldınız.",
		ukUA: "Ви приєдналися до спільної оплати.",
		uz:   "Siz hisob-kitob ulashishga qo'shildingiz.",
		zhCN: "您已加入账单共享。",
	},
	ARTICLE_TITLE_SPLIT_BILL: {
		arEG: "تقسيم الفاتورة/الشراء",
		de:   "eine Rechnung teilen",
		en:   "Split bill/purchase",
		es:   "Compartir la cuenta/compra",
		faIR: "لایحه / خرید تقسیم شده",
		fr:   "Partager la facture/l'achat",
		id:   "Bagi tagihan/pembelian",
		it:   "Bolletta Split / acquisto",
		jaJP: "請求書/購入を分割",
		koKR: "청구서/구매 분할",
		pl:   "Podziel rachunek/zakup",
		ptBR: "Dividir conta/compra",
		ptPT: "Dividir conta/compra",
		ru:   "Разделить счёт/покупку",
		tr:   "Fatura/satın alma bölünmesi",
		ukUA: "Розділити рахунок/покупку",
		uz:   "Hisob/xaridni bo'lish",
		zhCN: "分摊账单/购买",
	},
	ARTICLE_SUBTITLE_SPLIT_BILL: {
		arEG: "المبلغ: %v\nمشاركة النفقات مع الأصدقاء وتتبع السدادات",
		de:   "Wert: %v\nTeile deine Kosten mit Freunden und verfolge deren Rückzahlungen.",
		en:   "Amount: %v\nShares expenses with friends & track paybacks",
		es:   "Importe: %v\nCompartir los gastos entre amigos y seguir las devoluciones",
		faIR: "مقدار: %v" + "\n" + "هزینه ها را با دوستان و بازپرداخت پیگیری می کند",
		fr:   "Montant: %v\nPartagez les dépenses avec des amis et suivez les remboursements",
		id:   "Jumlah: %v\nBerbagi pengeluaran dengan teman & lacak pembayaran kembali",
		it:   "Importo: %v\nDisponi i costi con gli amici e le retribuzioni delle tracce",
		jaJP: "金額: %v\n友達と費用を分担し、返済を追跡",
		koKR: "금액: %v\n친구와 비용을 공유하고 상환을 추적",
		pl:   "Kwota: %v\nDziel wydatki z przyjaciółmi i śledź spłaty",
		ptBR: "Valor: %v\nCompartilhe despesas com amigos e acompanhe reembolsos",
		ptPT: "Valor: %v\nCompartilhe despesas com amigos e acompanhe reembolsos",
		ru:   "Сумма: %v\nПоделить траты между друзьями и отследить возвраты",
		tr:   "Tutar: %v\nArkadaşlarınızla masrafları paylaşın ve geri ödemeleri takip edin",
		ukUA: "Сума: %v\nПоділити витрати між друзями та відстежити повернення",
		uz:   "Miqdori: %v\nDo'stlar bilan xarajatlarni baham ko'ring va to'lovlarni kuzating",
		zhCN: "金额: %v\n与朋友分摊费用并跟踪还款",
	},

	ARTICLE_NEW_DEBT_TITLE: {
		arEG: "دين جديد",
		de:   "Neuer Schuldschein",
		en:   "New debt",
		es:   "Nueva deuda",
		faIR: "بدهی جدید",
		fr:   "Nouvelle dette",
		id:   "Hutang baru",
		it:   "Nuovo debito",
		jaJP: "新しい借金",
		koKR: "새로운 부채",
		pl:   "Nowy dług",
		ptBR: "Nova dívida",
		ptPT: "Nova dívida",
		ru:   "Новый долг",
		tr:   "Yeni borç",
		ukUA: "Новий борг",
		uz:   "Yangi qarz",
		zhCN: "新债务",
	},
	ARTICLE_NEW_DEBT_SUBTITLE: {
		arEG: "المبلغ: %v\nيرسل إشعارات في تاريخ الاستحقاق إذا تم تعيينه",
		de:   "Wert: %v\nZur Fälligkeit wird eine Benachrichtigung geschickt, falls so eingestellt",
		en:   "Amount: %v\nSends notifications on due date if set",
		es:   "Importe: %v\nEnviar las notificaciones el día de vencimiento",
		faIR: "مقدار: %v" + "\n" + "اگر تنظیم شود، اطلاعیه ها را در تاریخ تعیین شده ارسال می کند",
		fr:   "Montant: %v\nEnvoie des notifications à la date d'échéance si défini",
		id:   "Jumlah: %v\nMengirim notifikasi pada tanggal jatuh tempo jika diatur",
		it:   "Importo: %v\nSend le notifiche alla data di scadenza se impostato",
		jaJP: "金額: %v\n設定されている場合、期日に通知を送信",
		koKR: "금액: %v\n설정된 경우 만기일에 알림 전송",
		pl:   "Kwota: %v\nWysyła powiadomienia w terminie płatności, jeśli ustawiono",
		ptBR: "Valor: %v\nEnvia notificações na data de vencimento, se definido",
		ptPT: "Valor: %v\nEnvia notificações na data de vencimento, se definido",
		ru:   "Сумма: %v\nЗапись долга и рассылка оповещений в день возврата.",
		tr:   "Tutar: %v\nAyarlanmışsa vade tarihinde bildirim gönderir",
		ukUA: "Сума: %v\nНадсилає сповіщення в день повернення, якщо встановлено",
		uz:   "Miqdori: %v\nAgar o'rnatilgan bo'lsa, belgilangan sanada bildirishnomalar yuboradi",
		zhCN: "金额: %v\n如果设置，将在到期日发送通知",
	},
	SPLITUS_PLEASE_JOIN_IF_NOT_ON_THE_LIST: {
		arEG: `يرجى الانضمام إذا لم يكن اسمك موجودًا في القائمة أعلاه.`,
		de:   `Bitte tritt zuerst bei, falls dein Name nicht auf der Liste ist.`,
		en:   `Please join if your name is not on the list above.`,
		es:   `Por favor únete si tu nombre no está en la lista anterior.`,
		faIR: `اگر نام شما در لیست بالا نیست، لطفا پیوست شوید.`,
		fr:   `Veuillez vous joindre si votre nom ne figure pas dans la liste ci-dessus.`,
		id:   `Silakan bergabung jika nama Anda tidak ada dalam daftar di atas.`,
		it:   `Si prega di unirti se il tuo nome non è nell'elenco di cui sopra.`,
		jaJP: `あなたの名前が上記のリストにない場合は、参加してください。`,
		koKR: `이름이 위 목록에 없으면 참여하세요.`,
		pl:   `Dołącz, jeśli Twojego imienia nie ma na powyższej liście.`,
		ptBR: `Por favor, junte-se se o seu nome não estiver na lista acima.`,
		ptPT: `Por favor, junte-se se o seu nome não estiver na lista acima.`,
		ru:   `Пожалуйста присоеденяйтесь если ваше не в списке.`,
		tr:   `Adınız yukarıdaki listede yoksa lütfen katılın.`,
		ukUA: `Будь ласка, приєднуйтесь, якщо вашого імені немає у списку вище.`,
		uz:   `Agar ismingiz yuqoridagi ro'yxatda bo'lmasa, iltimos, qo'shiling.`,
		zhCN: `如果您的名字不在上面的列表中，请加入。`,
	},
	SPLITUS_TEXT_HI_IN_GROUP: {
		arEG: `أنا <b>سبليتوس</b> . شكرًا لإضافتي!`,
		de:   `Ich bin <b>Splitus</b>. Danke fürs hinzufügen!`,
		en:   `I'm <b>Splitus</b>. Thanks for adding me!`,
		es:   `Soy <b>Splitus</b>. ¡Gracias por agregarme!`,
		faIR: `من <b>Splitus</b> با تشکر برای اضافه کردن من!`,
		fr:   `Je suis <b>Splitus</b>. Merci de m'avoir ajouté!`,
		id:   `Saya <b>Splitus</b>. Terima kasih telah menambahkan saya!`,
		it:   `Sono <b>Splitus</b>. Grazie per averci aggiunto!`,
		jaJP: `私は<b>Splitus</b>です。追加してくれてありがとう！`,
		koKR: `저는 <b>Splitus</b>입니다. 저를 추가해 주셔서 감사합니다!`,
		pl:   `Jestem <b>Splitus</b>. Dziękuję za dodanie mnie!`,
		ptBR: `Eu sou <b>Splitus</b>. Obrigado por me adicionar!`,
		ptPT: `Eu sou <b>Splitus</b>. Obrigado por me adicionar!`,
		ru:   `Меня зовут <b>Сплитус</b>. Спасибо что добавили!`,
		tr:   `Ben <b>Splitus</b>. Beni eklediğiniz için teşekkürler!`,
		ukUA: `Мене звати <b>Сплітус</b>. Дякую, що додали!`,
		uz:   `Men <b>Splitus</b>man. Meni qo'shganingiz uchun rahmat!`,
		zhCN: `我是<b>Splitus</b>。感谢您添加我！`,
	},
	COLLECTUS_TEXT_HI_IN_GROUP: {
		arEG: `أنا <b>كوليكتوس.</b> شكرًا لإضافتي!`,
		de:   `Ich bin <b>Collectus.</b> Danke fürs hinzufügen!`,
		en:   `I'm <b>Collectus.</b> Thanks for adding me!`,
		es:   `Soy <b>Collectus.</b> ¡Gracias por agregarme!`,
		faIR: `من <b>Collectus</b> با تشکر برای اضافه کردن من!`,
		fr:   `Je suis <b>Collectus.</b> Merci de m'avoir ajouté!`,
		id:   `Saya <b>Collectus.</b> Terima kasih telah menambahkan saya!`,
		it:   `Sono <b>Collectus.</b> Grazie per averci aggiunto!`,
		jaJP: `私は<b>Collectus.</b>です。追加してくれてありがとう！`,
		koKR: `저는 <b>Collectus.</b>입니다. 저를 추가해 주셔서 감사합니다!`,
		pl:   `Jestem <b>Collectus.</b> Dziękuję za dodanie mnie!`,
		ptBR: `Eu sou <b>Collectus.</b> Obrigado por me adicionar!`,
		ptPT: `Eu sou <b>Collectus.</b> Obrigado por me adicionar!`,
		ru:   `Меня зовут <b>Коллектус.</b> Спасибо что добавили!`,
		tr:   `Ben <b>Collectus.</b> Beni eklediğiniz için teşekkürler!`,
		ukUA: `Мене звати <b>Коллектус.</b> Дякую, що додали!`,
		uz:   `Men <b>Collectus.</b>man. Meni qo'shganingiz uchun rahmat!`,
		zhCN: `我是<b>Collectus.</b>感谢您添加我！`,
	},
	MT_GROUP_LABEL: {
		arEG: `<b>المجموعة</b> : %v`,
		de:   `<b>Gruppe</b>: %v`,
		en:   `<b>Group</b>: %v`,
		es:   `<b>Grupo</b>: %v`,
		faIR: `<b>گروه</b>: %v`,
		fr:   `<b>Groupe</b>: %v`,
		id:   `<b>Grup</b>: %v`,
		it:   `<b>Gruppo</b>: %v`,
		jaJP: `<b>グループ</b>: %v`,
		koKR: `<b>그룹</b>: %v`,
		pl:   `<b>Grupa</b>: %v`,
		ptBR: `<b>Grupo</b>: %v`,
		ptPT: `<b>Grupo</b>: %v`,
		ru:   `<b>Группа</b>: %v`,
		tr:   `<b>Grup</b>: %v`,
		ukUA: `<b>Група</b>: %v`,
		uz:   `<b>Guruh</b>: %v`,
		zhCN: `<b>组</b>: %v`,
	},
	MT_SPONSORS_HEADER: {
		arEG: `<b>الرعاة</b> :`,
		de:   `<b>Sponsoren</b>:`,
		en:   `<b>Sponsors</b>:`,
		es:   `<b>Patrocinadores</b>:`,
		faIR: `<b>حامیان</b>:`,
		fr:   `<b>Commanditaires</b> :`,
		id:   `<b>Sponsor</b>:`,
		it:   `<b>Sponsor</b> :`,
		jaJP: `<b>スポンサー</b>:`,
		koKR: `<b>스폰서</b>:`,
		pl:   `<b>Sponsorzy</b>:`,
		ptBR: `<b>Patrocinadores</b>:`,
		ptPT: `<b>Patrocinadores</b>:`,
		ru:   `<b>Спонсоры</b>:`,
		tr:   `<b>Sponsorlar</b>:`,
		ukUA: `<b>Спонсори</b>:`,
		uz:   `<b>Homiylar</b>:`,
		zhCN: `<b>赞助商</b>:`,
	},
	MT_DEBTORS_HEADER: {
		arEG: `<b>المدينون</b> :`,
		de:   `<b>Schuldner</b>:`,
		en:   `<b>Debtors</b>:`,
		es:   `<b>Deudores</b>:`,
		faIR: `<b>بدهکاران</b>:`,
		fr:   `<b>Débiteurs</b>:`,
		id:   `<b>Debitur</b>:`,
		it:   `<b>Debitori</b>:`,
		jaJP: `<b>債務者</b>:`,
		koKR: `<b>채무자</b>:`,
		pl:   `<b>Dłużnicy</b>:`,
		ptBR: `<b>Devedores</b>:`,
		ptPT: `<b>Devedores</b>:`,
		ru:   `<b>Должники</b>:`,
		tr:   `<b>Borçlular</b>:`,
		ukUA: `<b>Боржники</b>:`,
		uz:   `<b>Qarzdorlar</b>:`,
		zhCN: `<b>债务人</b>:`,
	},
	BT_DEFAULT_CURRENCY: {
		arEG: `العملة: %v`,
		de:   `Währung: %v`,
		en:   `Currency: %v`,
		es:   `Moneda: %v`,
		faIR: `واحد پول: %v`,
		fr:   `Devise: %v`,
		id:   `Mata uang: %v`,
		it:   `Moneta: %v`,
		jaJP: `通貨: %v`,
		koKR: `통화: %v`,
		pl:   `Waluta: %v`,
		ptBR: `Moeda: %v`,
		ptPT: `Moeda: %v`,
		ru:   `Валюта: %v`,
		tr:   `Para birimi: %v`,
		ukUA: `Валюта: %v`,
		uz:   `Valyuta: %v`,
		zhCN: `货币: %v`,
	},
	MESSAGE_TEXT_ASK_LANG: {
		arEG: `ما هي اللغة التي يجب أن أستخدمها في هذه المجموعة؟`,
		de:   `Welche Sprache wird hier gesprochen?`,
		en:   `What language should I use in this group?`,
		es:   `¿Qué idioma debería usar en este grupo?`,
		faIR: `کدام زبان باید در این گروه استفاده کنم؟`,
		fr:   `Quelle langue dois-je utiliser dans ce groupe?`,
		id:   `Bahasa apa yang harus saya gunakan dalam grup ini?`,
		it:   `Che lingua devo utilizzare in questo gruppo?`,
		jaJP: `このグループではどの言語を使用すべきですか？`,
		koKR: `이 그룹에서 어떤 언어를 사용해야 합니까?`,
		pl:   `Jakiego języka powinienem używać w tej grupie?`,
		ptBR: `Que idioma devo usar neste grupo?`,
		ptPT: `Que idioma devo usar neste grupo?`,
		ru:   `Какой язык я должен использовать в этой группе?`,
		tr:   `Bu grupta hangi dili kullanmalıyım?`,
		ukUA: `Якою мовою мені спілкуватися в цій групі?`,
		uz:   `Bu guruhda qaysi tildan foydalanishim kerak?`,
		zhCN: `我应该在这个群组中使用什么语言？`,
	},
	MESSAGE_TEXT_HI_IN_GROUP_LANG_SET: {
		arEG: `رائع، سأستخدم اللغة الإنجليزية.`,
		de:   `Kein Problem, dann schreibe ich auf Deutsch.`,
		en:   `Great, I'll be using English.`,
		es:   `Genial, usaré español.`,
		faIR: `عالی، من از فارسی استفاده خواهم کرد.`,
		fr:   `Super, j'utiliserai le français.`,
		id:   `Bagus, saya akan menggunakan bahasa Indonesia.`,
		it:   `Ottimo, userò l'italiano.`,
		jaJP: `素晴らしい、日本語を使用します。`,
		koKR: `좋아요, 한국어를 사용하겠습니다.`,
		pl:   `Świetnie, będę używać języka polskiego.`,
		ptBR: `Ótimo, vou usar o português.`,
		ptPT: `Ótimo, vou usar o português.`,
		ru:   `Отлично, я буду использовать русский`,
		tr:   `Harika, Türkçe kullanacağım.`,
		ukUA: `Чудово, я буду використовувати українську мову.`,
		uz:   `Ajoyib, men o'zbek tilidan foydalanaman.`,
		zhCN: `太好了，我将使用中文。`,
	},
	SPLITUS_TEXT_ABOUT_ME_AND_CO: {
		arEG: `أساعد في <b>توزيع الفواتير</b> . صديقي @DebtsTrackerBot يتتبع سداد الديون.`,
		de:   `Ich kann helfen, <b>Rechnungen zu teilen</b>. Mein Freund @DebtsTrackerBot passt darauf auf, dass alle Schulden zurückgezahlt werden.`,
		en:   `I help to <b>split bills</b>. My friend @DebtsTrackerBot is tracking paybacks & debts.`,
		es:   `Ayudo a <b>dividir billetes</b>. Mi amigo @DebtsTrackerBot está rastreando pagos y deudas.`,
		faIR: `من به <b>تقسیم صورتحساب </b> کمک میکنم دوست منDebtsTrackerBot ردیابی بازپرداخت و بدهی است.`,
		fr:   `J&#39;aide à <b>répartir les factures</b> . Mon ami @DebtsTrackerBot suit les remboursements et les dettes.`,
		id:   `Saya membantu <b>membagi tagihan</b> . Teman saya @DebtsTrackerBot melacak pelunasan dan utang.`,
		it:   `Aiuto a <b>dividere le bollette</b>. Il mio amico @DebtsTrackerBot sta monitorando i pagamenti e i debiti.`,
		jaJP: `私は<b>請求書の分割</b>を手伝っています。友人の@DebtsTrackerBotが返済と負債を追跡しています。`,
		koKR: `저는 <b>청구서 분납</b> 을 도와드립니다. 제 친구 @DebtsTrackerBot은 상환금과 부채를 추적하고 있습니다.`,
		pl:   `Pomagam <b>dzielić rachunki</b> . Mój przyjaciel @DebtsTrackerBot śledzi spłaty i długi.`,
		ptBR: `Eu ajudo a <b>dividir contas</b> . Meu amigo @DebtsTrackerBot está monitorando pagamentos e dívidas.`,
		ptPT: `Eu ajudo a <b>dividir contas</b> . O meu amigo @DebtsTrackerBot está a monitorizar pagamentos e dívidas.`,
		ru:   `Я помогаю делить счета. Мой друг @DebtsTrackerRuBot отслеживает платежи и долги.`,
		tr:   `<b>Faturaları bölmeye</b> yardım ediyorum. Arkadaşım @DebtsTrackerBot ödemeleri ve borçları takip ediyor.`,
		ukUA: `Я допомагаю <b>розподіляти рахунки</b> . Мій друг @DebtsTrackerBot відстежує повернення боргів та борги.`,
		uz:   `Men <b>hisob-kitoblarni ajratishga</b> yordam beraman. Mening do&#39;stim @DebtsTrackerBot to&#39;lovlar va qarzlarni kuzatib boradi.`,
		zhCN: `我帮忙<b>分摊账单</b>。我的朋友@DebtsTrackerBot 负责追踪还款和债务。`},
	COLLECTUS_TEXT_ABOUT_ME_AND_CO: {
		arEG: `أساعد في <b>جمع المال</b> لقضية نبيلة، كهدية عيد ميلاد مثلاً. 🎉

صديقي @DebtsTrackerBot يتتبع الديون وسدادها.

وإذا كنتم تقومون بعمليات شراء جماعية وترغبون في تقسيم الفواتير، فإن @SplitusBot هنا لمساعدتكم.`,
		de: `Ich kann dabei helfen <b>Geld zu sammeln</b>. Zum Beispiel für ein Geburtstagsgeschenk. 🎉

Mein Freund @DebtsTrackerBot überwacht Schulden und deren Rückzahlungen.

Wenn du Geld für irgendwas sammeln willst, kann @SplitusBot dir dabei helfen.`,
		en: `I help to <b>collect money</b> for a good cause. For example for a birthday present. 🎉

My buddy @DebtsTrackerBot is tracking debts & paybacks.

And if you do collective purchases and want to split bills @SplitusBot is here to help.`,
		es: `Ayudo a <b>recaudar dinero</b> para una buena causa. Por ejemplo, para un regalo de cumpleaños. 🎉

Mi amigo @DebtsTrackerBot rastrea deudas y pagos.

Y si hacen compras en grupo y quieren dividir las facturas, @SplitusBot está aquí para ayudarte.`,
		faIR: //nolint:staticcheck // disable ST1018 for this line
		`من به <b>جمع‌آوری پول</b> برای یک هدف خیرخواهانه کمک می‌کنم. مثلاً برای هدیه تولد. 🎉

دوست من @DebtsTrackerBot بدهی‌ها و بازپرداخت‌ها را پیگیری می‌کند.

و اگر خریدهای دسته‌جمعی انجام می‌دهید و می‌خواهید صورتحساب‌ها را تقسیم کنید، @SplitusBot اینجا برای کمک به شماست.`,
		fr: `J&#39;aide à <b>collecter des fonds</b> pour une bonne cause. Par exemple, pour un cadeau d&#39;anniversaire. 🎉 Mon ami @DebtsTrackerBot suit vos dettes et leurs remboursements. Si vous effectuez des achats groupés et souhaitez partager les factures, @SplitusBot est là pour vous aider.`,
		id: `Saya membantu <b>mengumpulkan uang</b> untuk tujuan yang baik. Misalnya untuk hadiah ulang tahun. 🎉

Teman saya @DebtsTrackerBot melacak utang &amp; pelunasan.

Dan jika Anda melakukan pembelian kolektif dan ingin membagi tagihan, @SplitusBot siap membantu.`,
		it: `Aiuto a <b>raccogliere fondi</b> per una buona causa. Ad esempio per un regalo di compleanno. 🎉

Il mio amico @DebtsTrackerBot tiene traccia di debiti e rimborsi.

E se fate acquisti collettivi e volete dividere il conto, @SplitusBot è qui per aiutarvi.`,
		jaJP: `良い目的のために<b>募金を集める</b>お手伝いをしています。例えば誕生日プレゼントなど。🎉

私の友達の@DebtsTrackerBotが借金と返済を追跡しています。

共同購入をして請求書を分割したい場合は、@SplitusBotがお手伝いします。`,
		koKR: `저는 좋은 목적을 위해 <b>모금 활동을</b> 돕습니다. 예를 들어 생일 선물 같은 거요. 🎉

제 친구 @DebtsTrackerBot이 부채와 상환 현황을 추적하고 있어요.

 공동 구매를 하시고 비용을 분담하고 싶으시다면 @SplitusBot이 도와드리겠습니다.`,
		pl: `Pomagam <b>zbierać pieniądze</b> na szczytny cel. Na przykład na prezent urodzinowy. 🎉

Mój kumpel @DebtsTrackerBot śledzi długi i spłaty.

A jeśli robisz zakupy zbiorcze i chcesz podzielić rachunki, @SplitusBot jest tutaj, aby pomóc.`,
		ptBR: `Ajudo a <b>arrecadar dinheiro</b> para uma boa causa. Por exemplo, para um presente de aniversário. 🎉

Meu amigo @DebtsTrackerBot está monitorando dívidas e pagamentos.

E se você faz compras coletivas e quer dividir as contas, o @SplitusBot está aqui para ajudar.`,
		ptPT: `Ajudo a <b>angariar dinheiro</b> para uma boa causa. Por exemplo, para um presente de aniversário. 🎉

O meu amigo @DebtsTrackerBot está a monitorizar dívidas e pagamentos.

E se faz compras coletivas e quer dividir as contas, o @SplitusBot está aqui para ajudar.`,
		ru: `Я помогаю <b>собирать деньги</b> на что нибудь. Например для подарка на день рождение. 🎉

Мой друг @DebtsTrackerRuBot отслеживает долги и платежи.

А если хотите вести учёт совместных покупок и удобно делить счета вам поможет @SplitusBot.`,
		tr: `İyi bir amaç için <b>para toplamaya</b> yardım ediyorum. Örneğin bir doğum günü hediyesi için. 🎉

Arkadaşım @DebtsTrackerBot borçları ve geri ödemeleri takip ediyor.

Ve eğer toplu alışveriş yapıyorsanız ve faturaları bölmek istiyorsanız @SplitusBot yardım etmek için burada.`,
		ukUA: `Я допомагаю <b>збирати гроші</b> на добру справу. Наприклад, на подарунок на день народження. 🎉

Мій приятель @DebtsTrackerBot відстежує борги та повернення коштів.

А якщо ви робите колективні покупки та хочете розділити рахунки, @SplitusBot тут, щоб допомогти.`,
		uz: `Men yaxshi ish uchun <b>pul yig&#39;ishga</b> yordam beraman. Masalan, tug&#39;ilgan kun sovg&#39;asi uchun. 🎉

Mening do‘stim @DebtsTrackerBot qarzlar va to‘lovlarni kuzatib boradi.

Agar siz jamoaviy xaridlar qilsangiz va hisob-kitoblarni taqsimlamoqchi bo‘lsangiz, @SplitusBot yordam berish uchun shu yerda.`,
		zhCN: `我帮助为公益事业<b>募集善款</b>。例如，为生日礼物募集善款。🎉

我的好朋友@DebtsTrackerBot 负责追踪债务和还款。

如果您进行集体购物并希望分摊账单，@SplitusBot 可以为您提供帮助。`},
	SPLITUS_TEXT_HI: {
		arEG: `انا <b>سبليتوس</b> .`,
		de:   `Ich bin <b>Splitus</b>.`,
		en:   `I'm <b>Splitus</b>.`,
		es:   `Soy <b>Splitus</b> .`,
		faIR: `من <b>اسپلیتاس</b> هستم.`,
		fr:   `Je suis <b>Splitus</b> .`,
		id:   `Saya <b>Splitus</b> .`,
		it:   `Io sono <b>Splitus</b> .`,
		jaJP: `私は<b>Splitus</b>です。`,
		koKR: `저는 <b>스플리투스예요</b> .`,
		pl:   `Jestem <b>Splitus</b> .`,
		ptBR: `Eu sou <b>Splitus</b> .`,
		ptPT: `Eu sou <b>o Splitus</b> .`,
		ru:   `Меня зовут <b>Сплитус</b>.`,
		tr:   `Ben <b>Splitus&#39;um</b> .`,
		ukUA: `Я <b>Сплітус</b> .`,
		uz:   `Men <b>Splitusman</b> .`,
		zhCN: `我是<b>Splitus</b> 。`},
	COLLECTUS_TEXT_HI: {
		arEG: `أنا <b>كوليكتوس</b> .`,
		de:   `Ich bin <b>Collectus</b>.`,
		en:   `I'm <b>Collectus</b>.`,
		es:   `Soy <b>Collectus</b> .`,
		faIR: `من <b>کالکتوس</b> هستم.`,
		fr:   `Je suis <b>Collectus</b> .`,
		id:   `Saya <b>Collectus</b> .`,
		it:   `Io sono <b>Collectus</b> .`,
		jaJP: `私は<b>Collectus</b>です。`,
		koKR: `저는 <b>콜렉터스예요</b> .`,
		pl:   `Jestem <b>Collectus</b> .`,
		ptBR: `Eu sou <b>o Collectus</b> .`,
		ptPT: `Eu sou <b>o Collectus</b> .`,
		ru:   `Меня зовут <b>Коллектус</b>.`,
		tr:   `Ben <b>Collectus&#39;um</b> .`,
		ukUA: `Я <b>Колектус</b> .`,
		uz:   `Men <b>Collectusman</b> .`,
		zhCN: `我是<b>Collectus</b> 。`},
	SPLITUS_TG_COMMANDS: {
		arEG: `<b>أوامر الروبوت:</b> 
 /groups - قائمة المجموعات
 /bills - قائمة الفواتير المستحقة
 /menu - القائمة الرئيسية
 /settings - الإعدادات
 /help - تعلم كيفية استخدام الروبوت والإبلاغ عن المشكلات وطرح الأسئلة`,
		de: `<b>Bot-Befehle:</b> 
 /groups - Liste der Gruppen
 /bills - Liste der ausstehenden Rechnungen
 /menu - Hauptmenü
 /settings - Einstellungen
 /help - Erfahren Sie, wie Sie den Bot verwenden, Probleme melden und Fragen stellen`,
		en: `<b>Bot commands:</b>
	/groups - List of groups
	/bills - List of outstanding bills
	/menu - Main menu
	/settings - Settings
	/help - Learn how to use bot, report issues, ask questions`,
		es: `<b>Comandos del bot:</b> 
 /groups - Lista de grupos
 /bills - Lista de facturas pendientes
 /menu - Menú principal
 /settings - Configuraciones
 /help - Aprende a usar el bot, informa problemas, haz preguntas`,
		faIR: `<b>دستورات ربات:</b> 
 /groups - فهرست گروه\u200cها
 /bills - فهرست صورتحساب\u200cهای معوق
 /menu - منوی اصلی
 /settings - Settings
 /help - یادگیری نحوه استفاده از ربات، گزارش مشکلات، پرسیدن سوال`,
		fr: `<b>Commandes du bot :</b> 
 /groups - Liste des groupes
 /bills - Liste des factures impayées
 /menu - Menu principal
 /settings - Paramètres
 /help - Apprendre à utiliser le bot, signaler des problèmes, poser des questions`,
		id: `<b>Perintah bot:</b> 
 /groups - Daftar grup
 /bills - Daftar tagihan yang belum dibayar
 /menu - Menu utama
 /settings - Pengaturan
 /help - Pelajari cara menggunakan bot, melaporkan masalah, mengajukan pertanyaan`,
		it: `<b>Comandi del bot:</b> 
 /groups - Elenco dei gruppi
 /bills - Elenco delle fatture in sospeso
 /menu - Menu principale
 /settings - Impostazioni
 /help - Scopri come usare il bot, segnalare problemi, porre domande`,
		jaJP: `<b>ボットコマンド:</b> 
 /groups - グループのリスト
 /bills - 未払いの請求書のリスト
 /menu - メインメニュー
 /settings - 設定
 /help - ボットの使い方、問題の報告、質問の方法`,
		koKR: `<b>봇 명령어:</b> 
 /groups - 그룹 목록
 /bills - 미납 청구서 목록
 /menu - 메인 메뉴
 /settings - 설정
 /help - 봇 사용 방법 알아보기, 문제 보고, 질문하기`,
		pl: `<b>Polecenia bota:</b> 
 /groups - Lista grup
 /bills - Lista niezapłaconych rachunków
 /menu - Menu główne
 /settings - Ustawienia
 /help - Dowiedz się, jak korzystać z bota, zgłaszaj problemy, zadawaj pytania`,
		ptBR: `<b>Comandos do bot:</b> 
 /groups - Lista de grupos
 /bills - Lista de contas pendentes
 /menu - Menu principal
 /settings - Configurações
 /help - Aprenda a usar o bot, relatar problemas, fazer perguntas`,
		ptPT: `<b>Comandos do bot:</b> 
 /groups - Lista de grupos
 /bills - Lista de contas pendentes
 /menu - Menu principal
 /settings - Definições
 /help - Aprenda a utilizar o bot, a reportar problemas, a fazer perguntas`,
		ru: `<b>Команды для бота:</b>
	/groups - Список групп
	/bills - Список незакрытых платежей
	/menu - Главное меню
	/settings - Настройки
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
		tr: `<b>Bot komutları:</b> 
 /groups - Grupların listesi
 /bills - Ödenmemiş faturaların listesi
 /menu - Ana menü
 /settings - Ayarlar
 /help - Botu nasıl kullanacağınızı öğrenin, sorunları bildirin, sorular sorun`,
		ukUA: `<b>Команди бота:</b> 
 /groups - Список груп
 /bills - Список неоплачених рахунків
 /menu - Головне меню
 /settings - Налаштування
 /help - Дізнайтеся, як користуватися ботом, повідомляти про проблеми, ставити запитання`,
		uz: `<b>Bot buyruqlari:</b> 
 /groups - Guruhlar ro&#39;yxati
 /hisob-kitoblar - To&#39;lanmagan hisoblar ro&#39;yxati
 /menyu - Asosiy menyu
 /sozlamalar - Sozlamalar
 /yordam - Botdan qanday foydalanishni o&#39;rganing, muammolar haqida xabar bering, savollar bering`,
		zhCN: `<b>机器人命令：</b> 
 /groups - 群组列表
 /bills - 未付账单列表
 /menu - 主菜单
 /settings - 设置
 /help - 了解如何使用机器人、报告问题、提出问题`},
	COLLECTUS_TG_COMMANDS: {
		arEG: `<b>أوامر الروبوت:</b> 

 /groups - قائمة المجموعات
 /fundraisings - قائمة عمليات جمع التبرعات النشطة
 /help - تعلم كيفية استخدام الروبوت والإبلاغ عن المشكلات وطرح الأسئلة`,
		de: `<b>Bot-Befehle:</b> 

 /groups - Liste der Gruppen
 /fundraisings - Liste der aktiven Spendensammlungen
 /help - Erfahren Sie, wie Sie den Bot verwenden, Probleme melden und Fragen stellen`,
		en: `<b>Bot commands:</b>

	/groups - List of groups
	/fundraisings - List of active fundraisings
	/help - Learn how to use bot, report issues, ask questions`,
		es: `<b>Comandos del bot:</b> 

 /groups - Lista de grupos
 /fundraisings - Lista de recaudaciones de fondos activas
 /help - Aprende a usar el bot, informa problemas, haz preguntas`,
		faIR: `<b>دستورات ربات:</b> 

 /groups - فهرست گروه\u200cها
 /fundraisings - فهرست جمع\u200cآوری کمک\u200cهای مالی فعال
 /help - نحوه استفاده از ربات را بیاموزید، مشکلات را گزارش دهید، سوال بپرسید`,
		fr: `<b>Commandes du bot :</b> 

 /groups - Liste des groupes
 /fundraisings - Liste des collectes de fonds actives
 /help - Apprendre à utiliser le bot, signaler des problèmes, poser des questions`,
		id: `<b>Perintah bot:</b> 

 /groups - Daftar grup
 /fundraisings - Daftar penggalangan dana aktif
 /help - Pelajari cara menggunakan bot, melaporkan masalah, mengajukan pertanyaan`,
		it: `<b>Comandi del bot:</b> 

 /groups - Elenco dei gruppi
 /fundraisings - Elenco delle raccolte fondi attive
 /help - Scopri come usare il bot, segnalare problemi, porre domande`,
		jaJP: `<b>ボットコマンド:</b> 

 /groups - グループのリスト
 /fundraisings - アクティブな募金活動のリスト
 /help - ボットの使い方、問題の報告、質問の方法を学ぶ`,
		koKR: `<b>봇 명령어:</b> 

 /그룹 - 그룹 목록
 /기금 모금 - 활성 기금 모금 목록
 /도움말 - 봇 사용 방법 알아보기, 문제 보고, 질문하기`,
		pl: `<b>Polecenia bota:</b> 

 /groups — lista grup
 /fundraisings — lista aktywnych zbiórek
 /help — dowiedz się, jak korzystać z bota, zgłaszaj problemy, zadawaj pytania`,
		ptBR: `<b>Comandos do bot:</b> 

 /groups - Lista de grupos
 /fundraisings - Lista de arrecadações de fundos ativas
 /help - Aprenda a usar o bot, relatar problemas, fazer perguntas`,
		ptPT: `<b>Comandos do bot:</b> 

 /groups - Lista de grupos
 /fundraisings - Lista de angariações de fundos ativas
 /help - Aprenda a utilizar o bot, a reportar problemas, a fazer perguntas`,
		ru: `<b>Команды для бота:</b>
	/groups - Список групп
	/fundraisings - Список активных сборов
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
		tr: `<b>Bot komutları:</b> 

 /groups - Grupların listesi
 /fundraisings - Etkin bağış toplamaların listesi
 /help - Botu nasıl kullanacağınızı öğrenin, sorunları bildirin, sorular sorun`,
		ukUA: `<b>Команди бота:</b> 

 /groups - Список груп
 /fundraisings - Список активних заходів зі збору коштів
 /help - Дізнайтеся, як користуватися ботом, повідомляти про проблеми, ставити запитання`,
		uz: `<b>Bot buyruqlari:</b> 

 /groups - Guruhlar roʻyxati
 /fandreyzings - Faol mablagʻ yigʻishlar roʻyxati
 /help - Botdan qanday foydalanishni oʻrganing, muammolar haqida xabar bering, savollar bering`,
		zhCN: `<b>机器人命令：</b> 

 /groups - 群组列表
 /fundraisings - 活跃筹款列表
 /help - 了解如何使用机器人、报告问题、提出问题`},
	MESSAGE_TEXT_SEND_HELP_COMMAND_FOR_HELP: {
		arEG: "أرسل /help للحصول على تفاصيل حول كيفية استخدام هذا الروبوت.",
		de:   "Sende /help für Hilfe über den Umgang mit diesen Bot.",
		en:   "Send /help for details on how to use this bot.",
		es:   "Envíe /ayuda para obtener detalles sobre cómo utilizar este bot.",
		faIR: "برای جزئیات نحوه استفاده از این ربات، /help را ارسال کنید.",
		fr:   "Envoyez /help pour plus de détails sur l&#39;utilisation de ce bot.",
		id:   "Kirim /help untuk rincian tentang cara menggunakan bot ini.",
		it:   "Invia /help per i dettagli su come utilizzare questo bot.",
		jaJP: "このボットの使用方法の詳細については、/help を送信してください。",
		koKR: "이 봇을 사용하는 방법에 대한 자세한 내용을 알아보려면 /help를 보내세요.",
		pl:   "Wyślij zapytanie /help, aby uzyskać szczegółowe informacje na temat korzystania z tego bota.",
		ptBR: "Envie /help para obter detalhes sobre como usar este bot.",
		ptPT: "Envie /help para obter detalhes sobre como utilizar este bot.",
		ru:   "Отправьте /help для справки по использованию бота.",
		tr:   "Bu botu nasıl kullanacağınıza dair detaylar için /help gönderin.",
		ukUA: "Надішліть /help для отримання детальної інформації про те, як користуватися цим ботом.",
		uz:   "Ushbu botdan qanday foydalanish haqida batafsil ma&#39;lumot uchun /yordam yuboring.",
		zhCN: "发送/help 获取有关如何使用该机器人的详细信息。"},
	MessageTextHiS1NGL: {
		en: `¡Hola! Hi! Привет! سلام! Hallo!`,
	},
	MESSAGE_TEXT_HI_USERNAME: {
		arEG: `%v فيروس العوز المناعي البشري!`,
		de:   `HIV %v!`,
		en:   `Hi %v!`,
		es:   `¡Hola %v!`,
		faIR://nolint:staticcheck // disable ST1018 for this line
		`اچ‌آی‌وی %v!`,
		fr:   `VIH %v!`,
		id:   `HIV %v!`,
		it:   `HIV %v!`,
		jaJP: `HIV %v！`,
		koKR: `안녕하세요 %v!`,
		pl:   `HIV %v!`,
		ptBR: `HIV %v!`,
		ptPT: `VIH %v!`,
		ru:   `Привет %v!`,
		tr:   `HIV %v!`,
		ukUA: `ВІЛ %v!`,
		uz:   `Salom %v!`,
		zhCN: `艾滋病病毒 %v！`},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		arEG: `يمكنك العودة إلى القائمة الرئيسية`,
		de:   `Du kannst zurück zum Haupt /menu`,
		en:   `You can go back to main /menu`,
		es:   `Puedes volver al inicio /menú`,
		faIR: `شما میتوانید به /منو ی اصلی مراجعه کنید.`,
		fr:   `Vous pouvez revenir au menu principal`,
		id:   `Anda dapat kembali ke menu utama`,
		it:   `Puoi tornare al menu' principale tramite /menu`,
		jaJP: `メインメニューに戻ることができます`,
		koKR: `메인/메뉴로 돌아갈 수 있습니다`,
		pl:   `Możesz wrócić do głównego /menu`,
		ptBR: `Você pode voltar ao menu principal`,
		ptPT: `Pode voltar ao menu principal`,
		ru:   `Можно вернуться назад в главное /меню`,
		tr:   `Ana menüye geri dönebilirsiniz`,
		ukUA: `Ви можете повернутися до головного меню /меню`,
		uz:   `Asosiy/menyuga qaytishingiz mumkin`,
		zhCN: `您可以返回主菜单`},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: {
		arEG: `لغة الروبوت المفضلة: %s`,
		de:   `Bevorzugte Sprache: %s`,
		en:   `Preferred bot language: %s`,
		es:   `Idioma favorito del bot: %s`,
		faIR: `زبان برنامه: %s`,
		fr:   `Langue préférée du bot : %s`,
		id:   `Bahasa bot yang disukai: %s`,
		it:   `Lingua del bot preferita: %s`,
		jaJP: `推奨ボット言語: %s`,
		koKR: `선호하는 봇 언어: %s`,
		pl:   `Preferowany język bota: %s`,
		ptBR: `Linguagem de bot preferida: %s`,
		ptPT: `Linguagem de bot preferida: %s`,
		ru:   `Выбранный язык бота: %s`,
		tr:   `Tercih edilen bot dili: %s`,
		ukUA: `Бажана мова бота: %s`,
		uz:   `Afzal bot tili: %s`,
		zhCN: `首选机器人语言：%s`},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		arEG: `ما هي لغتك المفضلة؟`,
		de:   `Welche ist deine bevorzugte Sprache?`,
		en:   `What is your preferred language?`,
		es:   `¿cuál es tu idioma preferida?`,
		faIR: `شما چه زبانی را ترجیح می دهید؟`,
		fr:   `Quelle est votre langue préférée ?`,
		id:   `Apa bahasa pilihan Anda?`,
		it:   `Qual'e' la tua lingua madre?`,
		jaJP: `あなたの好みの言語は何ですか?`,
		koKR: `당신이 선호하는 언어는 무엇입니까?`,
		pl:   `Jaki jest Twój preferowany język?`,
		ptBR: `Qual é o seu idioma preferido?`,
		ptPT: `Qual é o seu idioma preferido?`,
		ru:   `На каком языке вы хотели бы общаться?`,
		tr:   `Tercih ettiğiniz dil hangisidir?`,
		ukUA: `Яка ваша бажана мова?`,
		uz:   `Siz qaysi tilni afzal ko&#39;rasiz?`,
		zhCN: `您的首选语言是什么？`},
	ChooseLocaleMessageText: {
		arEG: "ما هي اللغة التي تريد التحدث معي بها؟",
		de:   "Auf welcher Sprache würdest du dich gerne mit mir unterhalten?",
		en:   "Which language you would like to talk to me?",
		es:   "¿En cuál idioma preferirías hablar conmigo?",
		faIR: "دوست دارید به چه زبانی با من صحبت کنید؟",
		fr:   "Dans quelle langue souhaites-tu me parler ?",
		id:   "Bahasa apa yang ingin Anda gunakan untuk berbicara dengan saya?",
		it:   "In quale lingua preferisci parlarmi?",
		jaJP: "どの言語で話したいですか?",
		koKR: "어떤 언어로 통화하고 싶으신가요?",
		pl:   "W jakim języku chcesz ze mną rozmawiać?",
		ptBR: "Em qual idioma você gostaria de falar comigo?",
		ptPT: "Em que idioma gostaria de falar comigo?",
		ru:   "На каком языке вы хотели бы общаться со мной?",
		tr:   "Benimle hangi dilde konuşmak istersiniz?",
		ukUA: "Якою мовою ви хотіли б зі мною розмовляти?",
		uz:   "Men bilan qaysi tilda gaplashmoqchisiz?",
		zhCN: "您想用哪种语言和我交谈？"},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		arEG: "لقد قمت بتبديل اللغة إلى %v",
		de:   "Du hast die Sprache geändert auf %v",
		en:   "You've switched language to %v",
		es:   "Has cambiado el idioma a %v",
		faIR: "شما زبان را به %v تغییر دادید. ",
		fr:   "Vous avez changé de langue pour %v",
		id:   "Anda telah mengganti bahasa ke %v",
		it:   "Hai cambiato lingua in %v",
		jaJP: "言語を%vに切り替えました",
		koKR: "언어를 %v로 전환했습니다.",
		pl:   "Zmieniłeś język na %v",
		ptBR: "Você mudou o idioma para %v",
		ptPT: "Alterou o idioma para %v",
		ru:   "Вы поменяли язык на %v",
		tr:   "Dili %v olarak değiştirdiniz",
		ukUA: "Ви переключили мову на %v",
		uz:   "Siz tilni %v ga almashtirdingiz",
		zhCN: "您已将语言切换为 %v"},
	MESSAGE_TEXT_WHATS_NEXT: {
		arEG: "ماذا بعد؟",
		de:   "Was als nächstes?",
		en:   "What's next?",
		es:   "¿Algo más?",
		faIR: "اقدام بعدی چیست؟",
		fr:   "Quelle est la prochaine étape ?",
		id:   "Apa berikutnya?",
		it:   "Cosa posso fare ora?",
		jaJP: "次は何？",
		koKR: "다음은 무엇인가요?",
		pl:   "Co dalej?",
		ptBR: "O que vem a seguir?",
		ptPT: "O que vem a seguir?",
		ru:   "Что будем делать дальше?",
		tr:   "Sırada ne var?",
		ukUA: "Що далі?",
		uz:   "Keyingi nima?",
		zhCN: "下一步是什么？"},
	MESSAGE_TEXT_REFERRERS_TITLE: {
		arEG: "أصدقائنا:",
		de:   "Unsere Freunde:",
		en:   "Our friends:",
		es:   "Nuestros amigos:",
		faIR: "دوستان ما:",
		fr:   "Nos amis:",
		id:   "Teman kami:",
		it:   "I nostri amici:",
		jaJP: "私たちの友人:",
		koKR: "우리의 친구들:",
		pl:   "Nasi przyjaciele:",
		ptBR: "Nossos amigos:",
		ptPT: "Os nossos amigos:",
		ru:   "Наши друзья:",
		tr:   "Dostlarımız:",
		ukUA: "Наші друзі:",
		uz:   "Bizning do&#39;stlarimiz:",
		zhCN: "我们的朋友："},
	COMMAND_TEXT_ADD_MY_TG_CHANNEL: {
		arEG: "أضف قناتي",
		de:   "Meinen Kanal hinzufügen",
		en:   "Add my channel",
		es:   "Agrega mi canal",
		faIR: "کانال من را اضافه کنید",
		fr:   "Ajouter ma chaîne",
		id:   "Tambahkan saluran saya",
		it:   "Aggiungi il mio canale",
		jaJP: "私のチャンネルを追加する",
		koKR: "내 채널 추가",
		pl:   "Dodaj mój kanał",
		ptBR: "Adicione meu canal",
		ptPT: "Adicionar o meu canal",
		ru:   "Добавить мой канал",
		tr:   "Kanalımı ekle",
		ukUA: "Додати мій канал",
		uz:   "Mening kanalimni qo&#39;shing",
		zhCN: "添加我的频道"},
	MESSAGE_TEXT_DEBTUS_COMMANDS: {
		arEG: `<b>أوامر للبوت</b> 
🏠 /menu - عرض القائمة الرئيسية
🔙 /return - إرجاع الديون المسجلة سابقًا
📥 /got - تسجيل الأموال التي تلقيتها من الآخرين
📤 /gave - تسجيل الأموال التي أعطيتها للآخرين
📚 /history - أحدث المعاملات
🏁 /balance - عرض الرصيد الحالي
⚙ /settings - تعديل تفضيلاتك`,
		de: `<b>Befehle für den Bot</b> 
🏠 /menu – Hauptmenü anzeigen
🔙 /return – zuvor erfasste Schulden zurückgeben
📥 /got – Geld aufzeichnen, das Sie von anderen erhalten haben
📤 /gave – Geld aufzeichnen, das Sie anderen gegeben haben
📚 /history – letzte Transaktionen
🏁 /balance – aktuellen Kontostand anzeigen
⚙ /settings – Ihre Einstellungen anpassen`,
		en: `<b>Commands for the bot</b>
🏠 /menu - show main menu
🔙 /return - return previously recorded debt
📥 /got - record money you received from others
📤 /gave - record money you gave to others
📚 /history - latest transactions
🏁 /balance - display current balance
⚙ /settings - adjust your preferences`,
		es: `<b>Comandos para el bot</b> 
🏠 /menu - mostrar el menú principal
🔙 /return - devolver la deuda registrada previamente
📥 /got - registrar el dinero que recibiste de otros
📤 /gave - registrar el dinero que diste a otros
📚 /history - últimas transacciones
🏁 /balance - mostrar el saldo actual
⚙ /settings - ajustar tus preferencias`,
		faIR: //nolint:staticcheck // disable ST1018 for this line
		`<b>دستورات ربات</b> 
🏠 /menu - نمایش منوی اصلی 
🔙 /return - بازگرداندن بدهی ثبت شده قبلی 
📥 /got - ثبت پولی که از دیگران دریافت کرده‌اید 
📤 /gave - ثبت پولی که به دیگران داده‌اید 
📚 /history - آخرین تراکنش‌ها 
🏁 /balance - نمایش موجودی فعلی 
⚙ /settings - تنظیم تنظیمات برگزیده شما`,
		fr: `<b>Commandes pour le bot</b> 
🏠 /menu - afficher le menu principal
🔙 /return - restituer la dette précédemment enregistrée
📥 /got - enregistrer l&#39;argent que vous avez reçu des autres
📤 /gave - enregistrer l&#39;argent que vous avez donné aux autres
📚 /history - dernières transactions
🏁 /balance - afficher le solde actuel
⚙ /settings - ajuster vos préférences`,
		id: `<b>Perintah untuk bot</b> 
🏠 /menu - tampilkan menu utama
🔙 /return - kembalikan utang yang tercatat sebelumnya
📥 /got - catat uang yang Anda terima dari orang lain
📤 /gave - catat uang yang Anda berikan kepada orang lain
📚 /history - transaksi terkini
🏁 /balance - tampilkan saldo terkini
⚙ /settings - sesuaikan preferensi Anda`,
		it: `<b>Comandi per il bot</b> 
🏠 /menu - mostra il menu principale
🔙 /return - restituisci il debito registrato in precedenza
📥 /got - registra il denaro ricevuto da altri
📤 /gave - registra il denaro dato ad altri
📚 /history - ultime transazioni
🏁 /balance - mostra il saldo attuale
⚙ /settings - regola le tue preferenze`,
		jaJP: `<b>ボットのコマンド</b>
🏠 /menu - メインメニューを表示する
🔙 /return - 以前に記録した借金を返す
📥 /got - 他の人から受け取ったお金を記録する
📤 /gave - 他の人に与えたお金を記録する
📚 /history - 最新の取引
🏁 /balance - 現在の残高を表示する
⚙ /settings - 設定を調整する`,
		koKR: `<b>봇에 대한 명령어</b> 
🏠 /menu - 메인 메뉴 표시
🔙 /return - 이전에 기록된 부채 반환
📥 /got - 다른 사람에게서 받은 돈 기록
📤 /gave - 다른 사람에게 준 돈 기록
📚 /history - 최근 거래
🏁 /balance - 현재 잔액 표시
⚙ /settings - 기본 설정 조정`,
		pl: `<b>Polecenia dla bota</b> 
🏠 /menu - wyświetl menu główne
🔙 /return - zwróć poprzednio zarejestrowany dług
📥 /got - zapisz pieniądze otrzymane od innych
📤 /gave - zapisz pieniądze przekazane innym
📚 /history - ostatnie transakcje
🏁 /balance - wyświetl bieżące saldo
⚙ /settings - dostosuj swoje preferencje`,
		ptBR: `<b>Comandos para o bot</b> 
🏠 /menu - mostrar menu principal
🔙 /return - retornar dívida registrada anteriormente
📥 /got - registrar dinheiro que você recebeu de outros
📤 /gave - registrar dinheiro que você deu a outros
📚 /history - últimas transações
🏁 /balance - exibir saldo atual
⚙ /settings - ajustar suas preferências`,
		ptPT: `<b>Comandos para o bot</b> 
🏠 /menu - mostrar menu principal
🔙 /return - retornar dívida registada anteriormente
📥 /got - registar dinheiro que recebeu de outros
📤 /gave - registar dinheiro que deu a outros
📚 /history - últimas transações
🏁 /balance - exibir saldo atual
⚙ /settings - ajustar as suas preferências`,
		ru: `<b>Команды бота</b>
🏠 /menu - показать основное меню
🔙 /return - записать возврат долга
📥 /got - записать о взятии в долг
📤 /gave - записать о том что дали взаймы
📚 /history - последние транзакции
🏁 /balance - показать текущий баланс
⚙ /settings - настройки`,
		tr: `<b>Bot için komutlar</b> 
🏠 /menu - ana menüyü göster
🔙 /return - daha önce kaydedilen borcu döndür
📥 /got - başkalarından aldığınız parayı kaydedin
📤 /gave - başkalarına verdiğiniz parayı kaydedin
📚 /history - son işlemler
🏁 /balance - mevcut bakiyeyi göster
⚙ /settings - tercihlerinizi ayarlayın`,
		ukUA: `<b>Команди для бота</b> 
🏠 /menu - показати головне меню
🔙 /return - повернути раніше записаний борг
📥 /got - записати гроші, отримані від інших
📤 /gave - записати гроші, які ви дали іншим
📚 /history - останні транзакції
🏁 /balance - показати поточний баланс
⚙ /settings - налаштувати ваші налаштування`,
		uz: `<b>Bot uchun buyruqlar</b> 
🏠 /menyu - asosiy menyuni ko&#39;rsatish
🔙 /qaytish - ilgari qayd etilgan qarzni qaytarish
📥 /oldim - boshqalardan olgan pulingizni yozib oling
📤 /berdim - boshqalarga bergan pulingizni yozib oling
📚 /tarix - so&#39;nggi tranzaksiyalar
📚 /history - so&#39;nggi tranzaksiyalar
⚙}joriy balansingizni rostlash /n afzalliklar`,
		zhCN: `<b>机器人的命令</b>
🏠 /menu - 显示主菜单
🔙 /return - 返回之前记录的债务
📥 /got - 记录你从别人那里收到的钱
📤 /gave - 记录你给予他人的钱
📚 /history - 最新交易
🏁 /balance - 显示当前余额
⚙ /settings - 调整你的偏好`},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		arEG: `
إذا اقترضت من شخص ما لتسجيله، استخدم /got.
إذا أقرضت شخصًا ما لتسجيله، استخدم /gave.

أو استخدم القائمة في الأسفل.`,
		de: `
Wenn du dir was von jemanden geliehen hast, wähle /anleihen.
Wenn du was an jemanden verliehen hast, wähle /verleihen.

Oder benutzt das Menü unten.`,
		en: `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.`,
		es: `
		Si alguien te ha prestado usa el comando  /recibido.
		Si has prestado a alguien usa el comando /dado.

O usa el menú de abajo en la pantalla.`,
		faIR: `
اگر از کسی قرض گرفته اید برای ثبت آن از /قرض_گرفتن استفاده کنید.
اگر به کسی قرض داده اید برای ثبت آن از /قرض_دادن استفاده کنید.

یا از منوی پایین استفاده نمایید.`,
		fr: `
Si vous avez emprunté à quelqu&#39;un pour l&#39;enregistrer, utilisez /got.
Si vous avez prêté à quelqu&#39;un pour l&#39;enregistrer, utilisez /gave.

Ou utilisez le menu en bas.`,
		id: `
Jika Anda meminjam dari seseorang untuk merekamnya gunakan /got.
Jika Anda meminjamkan kepada seseorang untuk merekamnya gunakan /gave.

Atau gunakan menu di bagian bawah.`,
		it: `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.`,
		jaJP: `
誰かから借りて録音した場合は /got を使用します。
誰かに貸して録音した場合は /gave を使用します。

または、下部にあるメニューを使用します。`,
		koKR: `
누군가에게 빌려서 녹음했다면 /got을 사용하세요.
누군가에게 빌려서 녹음했다면 /gave를 사용하세요.

또는 하단의 메뉴를 사용하세요.`,
		pl: `
Jeśli pożyczyłeś coś od kogoś, aby to nagrać, użyj /got.
Jeśli pożyczyłeś coś komuś, aby to nagrać, użyj /gave.

Lub użyj menu na dole.`,
		ptBR: `
Se você pegou emprestado de alguém para gravar, use /got.
Se você emprestou para alguém para gravar, use /gave.

Ou use o menu na parte inferior.`,
		ptPT: `
Se pediu emprestado a alguém para gravar, use /got.
Se emprestou a alguém para gravar, use /gave.

Ou use o menu na parte inferior.`,
		ru: `
	Если вы дали в долг воспользуйтесь командой /дал.
	Если вы одолжили что-то - командой /взял.

	Или воспользуйтесь меню внизу экрана.`,
		tr: `
Kaydetmek için birinden ödünç aldıysanız /got kullanın.
Kaydetmek için birinden ödünç aldıysanız /gave kullanın.

Veya alttaki menüyü kullanın.`,
		ukUA: `
Якщо ви позичили у когось для запису, використовуйте /got.
Якщо ви позичили комусь для запису, використовуйте /gave.

Або скористайтеся меню внизу.`,
		uz: `
Agar uni yozib olish uchun kimdandir qarz olgan boʻlsangiz, /got.
Agar kimgadir yozib olish uchun qarz bergan boʻlsangiz, /gave-dan foydalaning.

Yoki pastki qismidagi menyudan foydalaning.`,
		zhCN: `
如果您向某人借钱来记录，请使用 /got。
如果您借钱给某人来记录，请使用 /gave。

或者使用底部的菜单。`},
	MESSAGE_TEXT_HISTORY_HEADER: {
		arEG: "تاريخ",
		de:   "Verlauf",
		en:   "History",
		es:   "Cronología",
		faIR: "سوابق",
		fr:   "Histoire",
		id:   "Sejarah",
		it:   "Cronologia",
		jaJP: "歴史",
		koKR: "역사",
		pl:   "Historia",
		ptBR: "História",
		ptPT: "História",
		ru:   "История",
		tr:   "Tarih",
		ukUA: "Історія",
		uz:   "Tarix",
		zhCN: "历史"},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		arEG: "ليس لديك أي سجلات حتى الآن.",
		de:   "Du hast noch nichts gespeichert.",
		en:   "You don't have any records yet.",
		es:   "Aún no tienes ninguna notificación.",
		faIR: "شما هنوز هیچ ثبت سابقه ای ندارید.",
		fr:   "Vous n&#39;avez pas encore d&#39;enregistrements.",
		id:   "Anda belum memiliki rekaman apa pun.",
		it:   "Non hai nulla memorizzato.",
		jaJP: "まだ記録はありません。",
		koKR: "아직 기록이 없습니다.",
		pl:   "Nie masz jeszcze żadnych rekordów.",
		ptBR: "Você ainda não tem nenhum registro.",
		ptPT: "Ainda não tem nenhum registo.",
		ru:   "У вас пока нет ни одной записи.",
		tr:   "Henüz bir kaydınız yok.",
		ukUA: "У вас ще немає жодних записів.",
		uz:   "Sizda hali hech qanday yozuv yo&#39;q.",
		zhCN: "您还没有任何记录。"},
	MESSAGE_TEXT_HISTORY_LIST: {
		arEG: `<b>%v</b> <i>(آخر %d):</i> 
───────────────
%v`,
		de: `<b>%v</b> <i>(bis %d):</i>
─────────────
%v`,
		en: `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		es: `<b>%v</b> <i>(últimos %d):</i>
─────────────
%v`,
		faIR: `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,
		fr: `<b>%v</b> <i>(dernier %d) :</i> 
─────────────
%v`,
		id: `<b>%v</b> <i>(%d terakhir):</i> 
──────────────
%v`,
		it: `<b>%v</b> <i>(ultimi %d):</i>

─────────────
%v`,
		jaJP: `<b>%v</b> <i>(過去 %d):</i> 
───────────────
%v`,
		koKR: `<b>%v</b> <i>(마지막 %d):</i> 
──────────────
%v`,
		pl: `<b>%v</b> <i>(ostatni %d):</i> 
─────────────
%v`,
		ptBR: `<b>%v</b> <i>(último %d):</i> 
─────────────
%v`,
		ptPT: `<b>%v</b> <i>(último %d):</i> 
─────────────
%v`,
		ru: `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
		tr: `<b>%v</b> <i>(son %d):</i> 
─────────────
%v`,
		ukUA: `<b>%v</b> <i>(останній %d):</i> 
──────────────
%v`,
		uz: `<b>%v</b> <i>(oxirgi %d):</i> 
─────────────
%v`,
		zhCN: `<b>%v</b> <i>（最后 %d）：</i> 
──────────────
%v`},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		arEG: "ليس لديك أي سجلات بشأن الديون الحالية.",
		de:   "Keine unbeglichenen Schulden.",
		en:   "You have no records on current debts.",
		es:   "No hay ninguna notificación de deudas actuales.",
		faIR: "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
		fr:   "Vous n’avez aucun dossier sur les dettes actuelles.",
		id:   "Anda tidak memiliki catatan mengenai hutang saat ini.",
		it:   "Non hai nulla memorizzato nel debito corrente.",
		jaJP: "現在の借金に関する記録はありません。",
		koKR: "현재 부채에 대한 기록이 없습니다.",
		pl:   "Nie posiadasz żadnych zapisów dotyczących bieżących długów.",
		ptBR: "Você não tem registros de dívidas atuais.",
		ptPT: "Não tem registos de dívidas atuais.",
		ru:   "Нет записей о текущих долгах.",
		tr:   "Cari borçlarınıza ilişkin kaydınız bulunmamaktadır.",
		ukUA: "У вас немає записів про поточні борги.",
		uz:   "Sizda joriy qarzlar bo&#39;yicha hech qanday yozuv yo&#39;q.",
		zhCN: "您没有当前债务记录。"},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		arEG: "المجموع",
		de:   "Insgesamt",
		en:   "Total",
		es:   "Total",
		faIR: "مجموع",
		fr:   "Total",
		id:   "Total",
		it:   "Totale",
		jaJP: "合計",
		koKR: "총",
		pl:   "Całkowity",
		ptBR: "Total",
		ptPT: "Total",
		ru:   "Всего",
		tr:   "Toplam",
		ukUA: "Всього",
		uz:   "Jami",
		zhCN: "全部的"},
	BT_OTHER_CURRENCY: {
		arEG: "عملة أخرى",
		de:   "Eine andere Währung",
		en:   "Another currency",
		es:   "Otra moneda",
		faIR: "ارز دیگر",
		fr:   "Une autre monnaie",
		id:   "Mata uang lainnya",
		it:   "Un'altra valuta",
		jaJP: "別の通貨",
		koKR: "또 다른 통화",
		pl:   "Inna waluta",
		ptBR: "Outra moeda",
		ptPT: "Outra moeda",
		ru:   "Другая валюта",
		tr:   "Başka bir para birimi",
		ukUA: "Інша валюта",
		uz:   "Boshqa valyuta",
		zhCN: "另一种货币"},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		arEG: "حسنًا، من الآن فصاعدًا سأستخدم &#39;%v&#39; كعملة أساسية.",
		de:   "OK, von nun an ist '%v' deine Hauptwährung.",
		en:   "OK, from now on I will use '%v' as a primary currency.",
		es:   "OK, ahora voy a usar '%v' como moneda principal. ",
		faIR: "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
		fr:   "OK, à partir de maintenant, j&#39;utiliserai « %v » comme devise principale.",
		id:   "Oke, mulai sekarang saya akan menggunakan &#39;%v&#39; sebagai mata uang utama.",
		it:   "OK, da ora in poi usero' '%v' come moneta principale.",
		jaJP: "OK、これからは「%v」を主要通貨として使用します。",
		koKR: "좋습니다. 이제부터 &#39;%v&#39;를 기본 통화로 사용하겠습니다.",
		pl:   "OK, od tej pory będę używał &#39;%v&#39; jako waluty podstawowej.",
		ptBR: "OK, de agora em diante usarei &#39;%v&#39; como moeda primária.",
		ptPT: "OK, a partir de agora vou usar &#39;%v&#39; como moeda primária.",
		ru:   "OK, теперь я буду использовать '%v' как основную валюту.",
		tr:   "Tamam, bundan sonra birincil para birimi olarak &#39;%v&#39; kullanacağım.",
		ukUA: "Добре, відтепер я використовуватиму &#39;%v&#39; як основну валюту.",
		uz:   "OK, bundan buyon men asosiy valyuta sifatida “%v” dan foydalanaman.",
		zhCN: "好的，从现在开始我将使用“%v”作为主要货币。"},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		arEG: "%v - مدين لك بـ %v",
		de:   "%v - schuldet dir %v",
		en:   "%v - owes to you %v",
		es:   "%v - te debe %v",
		faIR: "%v - %v به شما بدهکار است ",
		fr:   "%v - vous doit %v",
		id:   "%v - berutang padamu %v",
		it:   "%v - ti deve %v.",
		jaJP: "%v - あなたに借りがあります %v",
		koKR: "%v - 당신에게 빚진 것 %v",
		pl:   "%v - jest Ci winien %v",
		ptBR: "%v - deve a você %v",
		ptPT: "%v - deve-lhe %v",
		ru:   "%v - долг вам %v",
		tr:   "%v - sana borçludur %v",
		ukUA: "%v - винен вам %v",
		uz:   "%v - sizga %v qarzdor",
		zhCN: "%v - 欠你 %v"},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		arEG: "مدين لك بنسبة %v",
		de:   "Schuldet dir %v",
		en:   "Owes to you %v",
		es:   "Te debe %v",
		faIR: "%v به شما بدهکار است ",
		fr:   "Je te dois %v",
		id:   "Berutang padamu %v",
		it:   "Hai un credito di %v",
		jaJP: "あなたに借りがある %v",
		koKR: "당신에게 빚진 것 %v",
		pl:   "Jest Ci winien %v",
		ptBR: "Deve a você %v",
		ptPT: "Deve-lhe %v",
		ru:   "Вам должны %v",
		tr:   "Sana %v borçlu",
		ukUA: "Винен тобі %v",
		uz:   "Sizdan qarzdorman %v",
		zhCN: "欠你 %v"},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		arEG: "تهانينا! ليس عليك أي شيء آخر لـ <b>%v</b> .",
		de:   "Hurra, du bist jetzt quitt mit <b>%v</b>.",
		en:   "Congratulations! You don't owe anything more to <b>%v</b>.",
		es:   "Bravo! Has saldado tu deuda con <b>%v</b>.",
		faIR: "تبریک! شما دیگر چیزی به <b>%v</b> بدهکار نیستید .",
		fr:   "Félicitations ! Vous ne devez plus rien à <b>%v</b> .",
		id:   "Selamat! Anda tidak berutang apa pun lagi kepada <b>%v</b> .",
		it:   "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
		jaJP: "おめでとうございます！これで<b>%v</b>に対して借りがあることはありません。",
		koKR: "축하합니다! <b>%v</b> 에게 더 이상 빚진 게 없습니다.",
		pl:   "Gratulacje! Nie jesteś już nic winien <b>%v</b> .",
		ptBR: "Parabéns! Você não deve mais nada a <b>%v</b> .",
		ptPT: "Parabéns! Não deve mais nada a <b>%v</b> .",
		ru:   "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		tr:   "Tebrikler! <b>%v&#39;ye</b> artık hiçbir borcunuz yok.",
		ukUA: "Вітаємо! Ви більше нічого не винні <b>%v</b> .",
		uz:   "Tabriklaymiz! Siz <b>%v</b> dan ortiq qarzdor emassiz.",
		zhCN: "恭喜！你不再欠<b>%v</b>任何东西了。"},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		arEG: "لا يدين لك <b>%v</b> بأي شيء آخر.",
		de:   "Du bist jetzt mit <b>%v</b> quitt.",
		en:   "<b>%v</b> does not owe anything more to you.",
		es:   "<b>%v</b> nadie te debe nada ya.",
		faIR: "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
		fr:   "<b>%v</b> ne vous doit plus rien.",
		id:   "<b>%v</b> tidak berutang apa pun lagi kepada Anda.",
		it:   "<b>%v</b> ha saldato il suo debito verso di te.",
		jaJP: "<b>%v は</b>あなたに対してこれ以上何も借りがありません。",
		koKR: "<b>%v는</b> 당신에게 더 이상 빚진 것이 없습니다.",
		pl:   "<b>%v</b> nie jest Ci już nic winien.",
		ptBR: "<b>%v</b> não lhe deve mais nada.",
		ptPT: "<b>%v</b> não lhe deve mais nada.",
		ru:   "У <b>%v</b> больше не осталось долгов перед вами.",
		tr:   "<b>%v&#39;nin</b> sana daha fazla borcu yok.",
		ukUA: "<b>%v</b> більше нічого вам не винен.",
		uz:   "<b>%v</b> sizdan boshqa hech narsa qarzdor emas.",
		zhCN: "<b>%v</b>不再欠你任何东西。"},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		arEG: "أنت مدين بـ%v",
		de:   "Du schuldest %v",
		en:   "You owe %v",
		es:   "Tú debes %v",
		faIR: "شما %v بدهکار هستید ",
		fr:   "Vous devez %v",
		id:   "Anda berutang %v",
		it:   "Hai un debito di %v",
		jaJP: "あなたは%vを借りています",
		koKR: "당신은 %v를 빚지고 있습니다",
		pl:   "Jesteś winien %v",
		ptBR: "Você deve %v",
		ptPT: "Deve %v",
		ru:   "Вы должны %v",
		tr:   "%v borcunuz var",
		ukUA: "Ви винні %v",
		uz:   "%v qarzingiz bor",
		zhCN: "你欠 %v"},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		arEG: "%v - أنت مدين بـ %v",
		de:   "%v - du schuldest %v",
		en:   "%v - you owe %v",
		es:   "%v - tú debes %v",
		faIR: "%v - شما %v بدهکار هستید ",
		fr:   "%v - vous devez %v",
		id:   "%v - Anda berutang %v",
		it:   "%v - tu gli/le devi %v",
		jaJP: "%v - あなたは %v を借りています",
		koKR: "%v - 당신은 %v를 빚지고 있습니다",
		pl:   "%v - jesteś winien %v",
		ptBR: "%v - você deve %v",
		ptPT: "%v - deve %v",
		ru:   "%v - вы должны %v",
		tr:   "%v - %v borcunuz var",
		ukUA: "%v - ви винні %v",
		uz:   "%v - siz %v qarzdorsiz",
		zhCN: "%v – 你欠 %v"},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		arEG: "ما هي عملتك الأساسية؟",
		de:   "Was ist deine Hauptwährung?",
		en:   "What is your primary currency?",
		es:   "¿Cuál es tu moneda principal?",
		faIR: "واحد پولی اولیه شما چیست؟",
		fr:   "Quelle est votre devise principale ?",
		id:   "Apa mata uang utama Anda?",
		it:   "Qual'e' la tua valuta principale?",
		jaJP: "あなたの主な通貨は何ですか?",
		koKR: "귀하의 주요 통화는 무엇입니까?",
		pl:   "Jaka jest Twoja podstawowa waluta?",
		ptBR: "Qual é sua moeda principal?",
		ptPT: "Qual é a sua moeda principal?",
		ru:   "Какая валюта для вас основная?",
		tr:   "Birincil para biriminiz nedir?",
		ukUA: "Яка ваша основна валюта?",
		uz:   "Sizning asosiy valyutangiz nima?",
		zhCN: "您的主要货币是什么？"},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY_FOR_GROUP: {
		arEG: "ما هي العملة الأساسية لهذه المجموعة؟",
		de:   "Was ist die Hauptwährung für diese Gruppe?",
		en:   "What is a primary currency for this group?",
		es:   "¿Cuál es tu moneda principal?",
		faIR: "واحد پولی اولیه شما چیست؟",
		fr:   "Quelle est la devise principale de ce groupe ?",
		id:   "Apa mata uang utama untuk kelompok ini?",
		it:   "Qual'e' la tua valuta principale?",
		jaJP: "このグループの主要通貨は何ですか?",
		koKR: "이 그룹의 주요 통화는 무엇입니까?",
		pl:   "Jaka jest podstawowa waluta tej grupy?",
		ptBR: "Qual é a moeda principal deste grupo?",
		ptPT: "Qual é a moeda principal deste grupo?",
		ru:   "Какая валюта основная для этой группы?",
		tr:   "Bu grubun birincil para birimi nedir?",
		ukUA: "Яка основна валюта для цієї групи?",
		uz:   "Bu guruh uchun asosiy valyuta nima?",
		zhCN: "这个群体的主要货币是什么？"},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		arEG: "فشل حذف المستخدم: %v",
		de:   "Konnte den Benutzer nicht löschen: %v",
		en:   "Failed to delete user: %v",
		es:   "Error durante la cancelación del usuario: %v",
		faIR: "حذف کاربر ناموفق بود: %v",
		fr:   "Échec de la suppression de l'utilisateur: %v",
		id:   "Gagal menghapus pengguna: %v",
		it:   "Errore durante la cancellazione dell'utente: %v",
		jaJP: "ユーザーの削除に失敗しました: %v",
		koKR: "사용자 삭제 실패: %v",
		pl:   "Nie udało się usunąć użytkownika: %v",
		ptBR: "Falha ao excluir usuário: %v",
		ptPT: "Falha ao eliminar utilizador: %v",
		ru:   "Не удалось удалить данные пользователя: %v",
		tr:   "Kullanıcı silinemedi: %v",
		ukUA: "Не вдалося видалити користувача: %v",
		uz:   "Foydalanuvchini o'chirib bo'lmadi: %v",
		zhCN: "删除用户失败: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		arEG: "تم حذف بيانات المستخدم",
		de:   "Die Benutzerdaten wurden gelöscht.",
		en:   "User's data has been deleted",
		es:   "Los datos del usuario han sido eliminados",
		faIR: "اطلاعات کاربر حذف شد.",
		fr:   "Les données de l'utilisateur ont été supprimées",
		id:   "Data pengguna telah dihapus",
		it:   "I dati dell'utente sono stati cancellati",
		jaJP: "ユーザーのデータが削除されました",
		koKR: "사용자 데이터가 삭제되었습니다",
		pl:   "Dane użytkownika zostały usunięte",
		ptBR: "Os dados do usuário foram excluídos",
		ptPT: "Os dados do utilizador foram excluídos",
		ru:   "Данные пользователя удалены",
		tr:   "Kullanıcı verileri silindi",
		ukUA: "Дані користувача видалено",
		uz:   "Foydalanuvchi ma'lumotlari o'chirildi",
		zhCN: "用户数据已被删除",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		arEG: "الرجاء اختيار من أعاد الدين أو لمن أعادته.",
		de:   "Bitte wähle, wer die Schuld beglichen hat oder wem du sie zurückgezahlt hast.",
		en:   "Please choose who returned the debt or to who you returned it.",
		es:   "Por favor, elige quien ha devuelto o a quien ha sido devuelta la deuda ",
		faIR: "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
		fr:   "Veuillez choisir qui a remboursé la dette ou à qui vous l'avez remboursée.",
		id:   "Silakan pilih siapa yang mengembalikan hutang atau kepada siapa Anda mengembalikannya.",
		it:   "Scegli con chi hai sanato un credito o un debito.",
		jaJP: "誰が借金を返済したか、または誰に返済したかを選択してください。",
		koKR: "누가 부채를 반환했는지 또는 누구에게 반환했는지 선택하십시오.",
		pl:   "Wybierz, kto zwrócił dług lub komu go zwróciłeś.",
		ptBR: "Por favor, escolha quem devolveu a dívida ou para quem você a devolveu.",
		ptPT: "Por favor, escolha quem devolveu a dívida ou a quem a devolveu.",
		ru:   "Выберете кому вы вернули долг или кто вернул его вам.",
		tr:   "Lütfen borcu kimin iade ettiğini veya kime iade ettiğinizi seçin.",
		ukUA: "Будь ласка, виберіть, хто повернув борг або кому ви його повернули.",
		uz:   "Iltimos, qarzni kim qaytarganini yoki siz kimga qaytarganingizni tanlang.",
		zhCN: "请选择谁归还了债务或您归还给了谁。",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		arEG: "الرجاء اختيار الدين الذي تم سداده كليًا أو جزئيًا.",
		de:   "Bitte wähle, ob die Schuld vollständig oder teilweise beglichen wurde.",
		en:   "Please choose a debt that has been returned fully or partially.",
		es:   "Por favor, elige una deuda que ha sido devuelta total o parcialmente. ",
		faIR: "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
		fr:   "Veuillez choisir une dette qui a été remboursée entièrement ou partiellement.",
		id:   "Silakan pilih hutang yang telah dikembalikan sepenuhnya atau sebagian.",
		it:   "Scegli un debito che e' stato restituito completamente o parzialmente.",
		jaJP: "完全に、または部分的に返済された借金を選択してください。",
		koKR: "전액 또는 부분적으로 반환된 부채를 선택하십시오.",
		pl:   "Wybierz dług, który został zwrócony w całości lub częściowo.",
		ptBR: "Por favor, escolha uma dívida que foi devolvida total ou parcialmente.",
		ptPT: "Selecione uma dívida que tenha sido paga total ou parcialmente.",
		ru:   "Выберите долг который был возвращён целиком или частично.",
		tr:   "Lütfen tamamen veya kısmen iade edilmiş bir borç seçin.",
		ukUA: "Будь ласка, виберіть борг, який було повернуто повністю або частково.",
		uz:   "Iltimos, to'liq yoki qisman qaytarilgan qarzni tanlang.",
		zhCN: "请选择已全部或部分归还的债务。",
	},
	MESSAGE_TEXT_NO_DEBTS_TO_RETURN: {
		arEG: "ليس لديك سجلات للديون التي يمكن إرجاعها.",
		de:   "Du hast keine Aufzeichnungen über Schulden, die zurückgegeben werden können.",
		en:   "You have no records for debts that can be returned.",
		es:   "No tienes registros de deudas que puedan ser devueltas.",
		faIR: "شما هیچ سابقه ای از بدهی هایی که قابل بازگشت باشند ندارید.",
		fr:   "Vous n'avez aucun enregistrement de dettes qui peuvent être remboursées.",
		id:   "Anda tidak memiliki catatan untuk hutang yang dapat dikembalikan.",
		it:   "Non hai registrazioni di debiti che possono essere restituiti.",
		jaJP: "返済可能な借金の記録がありません。",
		koKR: "반환할 수 있는 부채에 대한 기록이 없습니다.",
		pl:   "Nie masz żadnych zapisów długów, które mogą zostać zwrócone.",
		ptBR: "Você não tem registros de dívidas que possam ser devolvidas.",
		ptPT: "Não tem registos de dívidas que podem ser devolvidas.",
		ru:   "У вас нет записей о догах для возврата.",
		tr:   "İade edilebilecek borçlar için hiçbir kaydınız yok.",
		ukUA: "У вас немає записів про борги, які можна повернути.",
		uz:   "Sizda qaytarilishi mumkin bo'lgan qarzlar uchun hech qanday yozuvlar yo'q.",
		zhCN: "您没有可以返还的债务记录。",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		arEG: "يرجى تأكيد أو رفض هذا التحويل.",
		de:   "Bitte stimme dem zu oder lehne es ab.",
		en:   "Please confirm or decline this transfer.",
		es:   "Por favor, confirma o rechaza la transacción.",
		faIR: "لطفاً این تراکنش را تایید یا رد نمایید.",
		fr:   "Veuillez confirmer ou refuser ce transfert.",
		id:   "Silakan konfirmasi atau tolak transfer ini.",
		it:   "Conferma o rifiuta questo debito/credito.",
		jaJP: "この転送を確認または拒否してください。",
		koKR: "이 전송을 확인하거나 거부하십시오.",
		pl:   "Proszę potwierdzić lub odrzucić ten transfer.",
		ptBR: "Por favor, confirme ou recuse esta transferência.",
		ptPT: "Por favor confirme ou recuse esta transferência.",
		ru:   "Пожалуйста подтвердите или отклоните эту транзакцию.",
		tr:   "Lütfen bu transferi onaylayın veya reddedin.",
		ukUA: "Будь ласка, підтвердіть або відхиліть цей переказ.",
		uz:   "Iltimos, ushbu o'tkazmani tasdiqlang yoki rad eting.",
		zhCN: "请确认或拒绝此转账。",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		arEG: "لقد تم قبول هذا النقل بالفعل.",
		de:   "Du hast dem bereits zugestimmt.",
		en:   "This transfer has been accepted already.",
		es:   "Esta transacción ya ha sido aceptada.",
		faIR: "این تراکنش قبلا قبول شده است.",
		fr:   "Ce transfert a déjà été accepté.",
		id:   "Transfer ini sudah diterima.",
		it:   "Questo debito/credito e' gia' stato accettato.",
		jaJP: "この転送はすでに承認されています。",
		koKR: "이 전송은 이미 수락되었습니다.",
		pl:   "Ten transfer został już zaakceptowany.",
		ptBR: "Esta transferência já foi aceita.",
		ptPT: "Esta transferência já foi aceite.",
		ru:   "Эта транзакция уже подтверждена.",
		tr:   "Bu transfer zaten kabul edildi.",
		ukUA: "Цей переказ вже підтверджено.",
		uz:   "Bu o'tkazma allaqachon qabul qilingan.",
		zhCN: "此转账已被接受。",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		arEG: "لقد تم رفض هذا النقل بالفعل.",
		de:   "Du hast dem bereits widersprochen.",
		en:   "This transfer has been declined already.",
		es:   "Esta transacción ya ha sido rechazada.",
		faIR: "این تراکنش قبلاً رد شده است.",
		fr:   "Ce transfert a déjà été refusé.",
		id:   "Transfer ini sudah ditolak.",
		it:   "Questo debito/credito e' gia' stato rifiutato.",
		jaJP: "この転送はすでに拒否されています。",
		koKR: "이 전송은 이미 거부되었습니다.",
		pl:   "Ten transfer został już odrzucony.",
		ptBR: "Esta transferência já foi recusada.",
		ptPT: "Esta transferência já foi recusada.",
		ru:   "Эта транзакция уже отклонена.",
		tr:   "Bu transfer zaten reddedildi.",
		ukUA: "Цей переказ вже відхилено.",
		uz:   "Bu o'tkazma allaqachon rad etilgan.",
		zhCN: "此转账已被拒绝。",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		arEG: "التفاصيل هنا: %v",
		de:   "Details hier: %v",
		en:   "Details here: %v",
		es:   "Detalles aquí: %v",
		faIR: "جزئیات: %v",
		fr:   "Détails ici: %v",
		id:   "Detail di sini: %v",
		it:   "Maggiori dettagli qui: %v",
		jaJP: "詳細はこちら: %v",
		koKR: "세부 정보: %v",
		pl:   "Szczegóły tutaj: %v",
		ptBR: "Detalhes aqui: %v",
		ptPT: "Detalhes aqui: %v",
		ru:   "Подробнее тут: %v",
		tr:   "Ayrıntılar burada: %v",
		ukUA: "Деталі тут: %v",
		uz:   "Tafsilotlar shu yerda: %v",
		zhCN: "详情在这里: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		arEG: "يرجى تقديم رقم الهاتف لـ <b>%v</b>",
		de:   "Bitte gib die Telefonnummer von <b>%v</b> an:",
		en:   "Please provide phone number for <b>%v</b>",
		es:   "Por favor escribe el número de teléfono de <b>%v</b>",
		faIR: "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
		fr:   "Veuillez fournir le numéro de téléphone pour <b>%v</b>",
		id:   "Silakan berikan nomor telepon untuk <b>%v</b>",
		it:   "Per favore fornisci il numero di telefono di <b>%v</b>",
		jaJP: "<b>%v</b>の電話番号を入力してください",
		koKR: "<b>%v</b>의 전화번호를 입력해주세요",
		pl:   "Proszę podać numer telefonu dla <b>%v</b>",
		ptBR: "Por favor, forneça o número de telefone para <b>%v</b>",
		ptPT: "Por favor, forneça o número de telefone de <b>%v</b>",
		ru:   "Пожалуйста напишите номер телефона <b>%v</b>.",
		tr:   "Lütfen <b>%v</b> için telefon numarası girin",
		ukUA: "Будь ласка, вкажіть номер телефону для <b>%v</b>",
		uz:   "Iltimos, <b>%v</b> uchun telefon raqamini kiriting",
		zhCN: "请提供<b>%v</b>的电话号码",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		arEG: "إذا كان رقم الهاتف موجودًا في دفتر العناوين الخاص بك، فيمكنك <b>استخدام الزر %v</b> لإرسال جهة الاتصال.",
		de:   "Wenn die Telefonnummer in deinem Adressbuch ist, kannst du den <b>%v Button benutzen</b>, um einen Kontakt zu senden.",
		en:   "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		es:   "Si el número está en tu agenda puedes <b>usar %v el botón</b> para enviar el contacto.",
		faIR: "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
		fr:   "Si le numéro de téléphone est dans votre carnet d'adresses, vous pouvez <b>utiliser le bouton %v</b> pour envoyer le contact.",
		id:   "Jika nomor telepon ada di buku alamat Anda, Anda dapat <b>menggunakan tombol %v</b> untuk mengirim kontak.",
		it:   "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
		jaJP: "電話番号がアドレス帳にある場合は、<b>%vボタンを使用</b>して連絡先を送信できます。",
		koKR: "전화번호가 주소록에 있는 경우 <b>%v 버튼을 사용</b>하여 연락처를 보낼 수 있습니다.",
		pl:   "Jeśli numer telefonu znajduje się w książce adresowej, możesz <b>użyć przycisku %v</b>, aby wysłać kontakt.",
		ptBR: "Se o número de telefone estiver em sua agenda, você pode <b>usar o botão %v</b> para enviar o contato.",
		ptPT: "Se o número de telefone estiver na sua lista de endereços, pode <b>utilizar o botão %v</b> para enviar o contacto.",
		ru:   "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		tr:   "Telefon numarası adres defterinizde varsa, kişiyi göndermek için <b>%v düğmesini kullanabilirsiniz</b>.",
		ukUA: "Якщо номер телефону є у вашій адресній книзі, ви можете <b>використати кнопку %v</b>, щоб надіслати контакт.",
		uz:   "Agar telefon raqami manzillar kitobingizda bo'lsa, kontaktni yuborish uchun <b>%v tugmasidan foydalanishingiz</b> mumkin.",
		zhCN: "如果电话号码在您的通讯录中，您可以<b>使用%v按钮</b>发送联系人。",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		arEG: "يجب أن يكون الرقم وفقًا للمعايير الدولية: * يبدأ بعلامة &quot;+&quot; متبوعًا برمز الدولة * يتكون من أرقام فقط * مثال: <b>+1</b> <code>999012345678</code>",
		de:   "Die Telefonnummer sollte dem internationalen Standard entsprechen:\n\t* Beginnend mit '+' gefolgt vom Ländercode (Deutschland +49)\n\t* Consist of numbers only\nExample: <b>+49</b><code>157123456</code>",
		en:   "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		es:   "El número debe tener formato internacional estándar:\n\t* Empezar con '+' seguido del código del país\n\t* formado solo por números\nEjemplo: <b>+1</b><code>999012345678</code>",
		faIR: "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <b>+1</b><code>999012345678</code>",
		fr:   "Le numéro doit être au format international:\n\t* Commence par '+' suivi du code du pays\n\t* Composé uniquement de chiffres\nExemple: <b>+33</b><code>612345678</code>",
		id:   "Nomor harus dalam standar internasional:\n\t* Dimulai dengan '+' diikuti oleh kode negara\n\t* Hanya terdiri dari angka\nContoh: <b>+62</b><code>812345678</code>",
		it:   "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <b>+39</b><code>34612345678</code>",
		jaJP: "番号は国際標準形式である必要があります:\n\t* '+'で始まり、国コードが続く\n\t* 数字のみで構成\n例: <b>+81</b><code>9012345678</code>",
		koKR: "번호는 국제 표준이어야 합니다:\n\t* '+'로 시작하여 국가 코드가 뒤따름\n\t* 숫자로만 구성\n예: <b>+82</b><code>1012345678</code>",
		pl:   "Numer powinien być w standardzie międzynarodowym:\n\t* Zaczyna się od '+' a następnie kod kraju\n\t* Składa się tylko z cyfr\nPrzykład: <b>+48</b><code>512345678</code>",
		ptBR: "O número deve estar no padrão internacional:\n\t* Começa com '+' seguido pelo código do país\n\t* Consiste apenas de números\nExemplo: <b>+55</b><code>11912345678</code>",
		ptPT: "O número deve estar no padrão internacional:\n\t* Começa por &#39;+&#39; seguido do indicativo do país\n\t* É constituído apenas por números\nExemplo: <b>+1</b> <code>999012345678</code>",
		ru:   "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <b>+7</b><code>999012345678</code>",
		tr:   "Numara uluslararası standartta olmalıdır:\n\t* '+' ile başlayıp ülke kodu ile devam eder\n\t* Sadece rakamlardan oluşur\nÖrnek: <b>+90</b><code>5301234567</code>",
		ukUA: "Номер повинен бути в міжнародному форматі:\n\t* Починатися зі знаку '+' та коду країни\n\t* Складатися тільки з цифр\nПриклад: <b>+380</b><code>991234567</code>",
		uz:   "Raqam xalqaro standartda bo'lishi kerak:\n\t* '+' bilan boshlanib, mamlakat kodi bilan davom etadi\n\t* Faqat raqamlardan iborat\nMisol: <b>+998</b><code>901234567</code>",
		zhCN: "号码应符合国际标准:\n\t* 以'+'开头，后跟国家代码\n\t* 仅由数字组成\n示例: <b>+86</b><code>13912345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		arEG: "سيتم إرسال رسالة نصية قصيرة إلى هذا الرقم:",
		de:   "Wir werden eine SMS an diese Nummer schicken:",
		en:   "Will send an SMS to this number:",
		es:   "Enviaremos una SMS a este número:",
		faIR: "یک پیام کوتاه به این شماره ارسال خواهد شد:",
		fr:   "Nous enverrons un SMS à ce numéro:",
		id:   "Akan mengirim SMS ke nomor ini:",
		it:   "Invieremo un SMS a questo numero:",
		jaJP: "このSMSをこの番号に送信します:",
		koKR: "이 번호로 SMS를 보낼 것입니다:",
		pl:   "Wyślemy SMS na ten numer:",
		ptBR: "Enviaremos um SMS para este número:",
		ptPT: "Enviarei um SMS para este número:",
		ru:   "На этот номер мы отправим SMS:",
		tr:   "Bu numaraya SMS gönderilecek:",
		ukUA: "Надішлемо SMS на цей номер:",
		uz:   "Bu raqamga SMS yuboriladi:",
		zhCN: "将向此号码发送短信:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		arEG: `<b>%v</b> مدين لك <b>بـ %v</b> .`,
		de:   `<b>%v</b> schuldet dir <b>%v</b>.`,
		en:   `<b>%v</b> owes to you <b>%v</b>.`,
		es:   `<b>%v</b> has prestado <b>%v</b>.`,
		faIR: `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
		fr:   `<b>%v</b> vous doit <b>%v</b>.`,
		id:   `<b>%v</b> berhutang kepada Anda <b>%v</b>.`,
		it:   `<b>%v</b> e' in debito di <b>%v</b>.`,
		jaJP: `<b>%v</b>はあなたに<b>%v</b>を借りています。`,
		koKR: `<b>%v</b>님이 당신에게 <b>%v</b>를 빚지고 있습니다.`,
		pl:   `<b>%v</b> jest ci winien <b>%v</b>.`,
		ptBR: `<b>%v</b> deve a você <b>%v</b>.`,
		ptPT: `<b>%v</b> deve-lhe <b>%v</b> .`,
		ru:   `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		tr:   `<b>%v</b> size <b>%v</b> borçlu.`,
		ukUA: `<b>%v</b> заборгував вам <b>%v</b>.`,
		uz:   `<b>%v</b> sizga <b>%v</b> qarzdir.`,
		zhCN: `<b>%v</b> 欠您 <b>%v</b>。`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		arEG: "أنت مدين لـ <b>%v</b> <b>%v</b> .",
		de:   "Du schuldest <b>%v</b> <b>%v</b>.",
		en:   "You owe to <b>%v</b> <b>%v</b>.",
		es:   "Te ha prestado <b>%v</b> <b>%v</b>.",
		faIR: "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
		fr:   "Vous devez <b>%v</b> à <b>%v</b>.",
		id:   "Anda berhutang kepada <b>%v</b> <b>%v</b>.",
		it:   `Tu devi dare a <b>%v</b> <b>%v</b>.`,
		jaJP: "あなたは<b>%v</b>に<b>%v</b>を借りています。",
		koKR: "당신은 <b>%v</b>에게 <b>%v</b>를 빚지고 있습니다.",
		pl:   "Jesteś winien <b>%v</b> <b>%v</b>.",
		ptBR: "Você deve a <b>%v</b> <b>%v</b>.",
		ptPT: "Deve a <b>%v</b> <b>%v</b> .",
		ru:   `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		tr:   "<b>%v</b>'e <b>%v</b> borçlusunuz.",
		ukUA: "Ви заборгували <b>%v</b> <b>%v</b>.",
		uz:   "Siz <b>%v</b>ga <b>%v</b> qarzdorsiz.",
		zhCN: "您欠<b>%v</b> <b>%v</b>。",
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {
		arEG: `هل تم سداد هذا الدين بالكامل؟

 <i>إذا كان جزئيًا، فيمكنك إدخال المبلغ على الفور.</i>`,
		de: `Wurde diese Schuld vollständig beglichen?

		<i>Falls nur teilweise, kann der Teilbetrag direkt eingegeben werden.</i>`,
		en: `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,
		es: `¿Ha sido totalmente devuelta esta deuda?

		<i>si ha sido devuelta parcialmente puedes introducir el importe.</i>`,
		faIR: `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,
		fr: `Cette dette a-t-elle été remboursée intégralement?

		<i>Si partiellement, vous pouvez saisir le montant tout de suite.</i>`,
		id: `Apakah hutang ini telah dikembalikan sepenuhnya?

		<i>Jika sebagian, Anda dapat memasukkan jumlah langsung.</i>`,
		it: `Il debito e' stato saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare da sottrarre, direttamente qui sotto.</i>`,
		jaJP: `この借金は全額返済されましたか？

		<i>部分的に返済された場合は、すぐに金額を入力できます。</i>`,
		koKR: `이 빚이 전액 상환되었습니까?

		<i>부분적으로 상환된 경우 금액을 바로 입력할 수 있습니다.</i>`,
		pl: `Czy ten dług został spłacony w całości?

		<i>Jeśli częściowo, możesz od razu wprowadzić kwotę.</i>`,
		ptBR: `Esta dívida foi devolvida integralmente?

		<i>Se parcialmente, você pode inserir o valor imediatamente.</i>`,
		ptPT: `Esta dívida foi liquidada na totalidade?

 <i>Se for parcialmente, pode introduzir o valor imediatamente.</i>`,
		ru: `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,
		tr: `Bu borç tamamen geri ödendi mi?

		<i>Kısmen ödendiyse, tutarı hemen girebilirsiniz.</i>`,
		ukUA: `Чи повернуто цей борг повністю?

		<i>Якщо частково, ви можете відразу ввести суму.</i>`,
		uz: `Bu qarz to'liq qaytarildimi?

		<i>Agar qisman bo'lsa, miqdorni darhol kiritishingiz mumkin.</i>`,
		zhCN: `这笔债务是否已全额归还？

		<i>如果部分归还，您可以立即输入金额。</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		arEG: `هذا البرنامج <b>مجاني</b> . يُرجى <a href="https://debtstracker.io/en/help-us">المساعدة</a> في تحسينه!`,
		de:   `Das Programm ist <b>kostenlos</b>. Bitte <a href="https://debtstracker.io/en/help-us">hilf</a> es besser zu machen!`,
		en:   `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		es:   `Este programa es <b>gratis</b>. Por favor <a href="https://debtstracker.io/en/help-us">ayúdanos</a> a mejorarlo!`,
		faIR: `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
		fr:   `Ce programme est <b>gratuit</b>. S'il vous plaît <a href="https://debtstracker.io/en/help-us">aidez-nous</a> à l'améliorer!`,
		id:   `Program ini <b>gratis untuk digunakan</b>. Silakan <a href="https://debtstracker.io/en/help-us">bantu</a> untuk membuatnya lebih baik!`,
		it:   `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
		jaJP: `このプログラムは<b>無料で使用できます</b>。より良くするために<a href="https://debtstracker.io/en/help-us">ご協力</a>ください！`,
		koKR: `이 프로그램은 <b>무료로 사용</b>할 수 있습니다. 더 나은 프로그램을 만들기 위해 <a href="https://debtstracker.io/en/help-us">도와주세요</a>!`,
		pl:   `Ten program jest <b>darmowy</b>. Proszę <a href="https://debtstracker.io/en/help-us">pomóż</a> uczynić go lepszym!`,
		ptBR: `Este programa é <b>gratuito para uso</b>. Por favor, <a href="https://debtstracker.io/en/help-us">ajude</a> a torná-lo melhor!`,
		ptPT: `Este programa é <b>gratuito</b> . <a href="https://debtstracker.io/en/help-us">Ajude</a> a melhorá-lo!`,
		ru:   `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,
		tr:   `Bu program <b>ücretsiz kullanım</b> içindir. Lütfen daha iyi hale getirmek için <a href="https://debtstracker.io/en/help-us">yardım edin</a>!`,
		ukUA: `Ця програма <b>безкоштовна</b>. Будь ласка, <a href="https://debtstracker.io/en/help-us">допоможіть</a> зробити її кращою!`,
		uz:   `Bu dastur <b>bepul foydalanish</b> uchun. Iltimos, uni yaxshilash uchun <a href="https://debtstracker.io/en/help-us">yordam bering</a>!`,
		zhCN: `此程序<b>免费使用</b>。请<a href="https://debtstracker.io/en/help-us">帮助</a>我们使其变得更好！`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		arEG: "%v | أنت مدين بـ: %v",
		de:   "%v | du schuldest: %v",
		en:   "%v | you owe: %v",
		es:   "%v | tú debes: %v",
		faIR: "%v | شما بدهکارید: %v",
		fr:   "%v | vous devez: %v",
		id:   "%v | Anda berhutang: %v",
		it:   "%v | tu devi: %v",
		jaJP: "%v | あなたの借り: %v",
		koKR: "%v | 당신이 빚진 금액: %v",
		pl:   "%v | jesteś winien: %v",
		ptBR: "%v | você deve: %v",
		ptPT: "%v | deve: %v",
		ru:   "%v | вы должны: %v",
		tr:   "%v | borçlusunuz: %v",
		ukUA: "%v | ви винні: %v",
		uz:   "%v | siz qarzdorsiz: %v",
		zhCN: "%v | 您欠: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		arEG: "%v | مدين لك: %v",
		de:   "%v | schuldet dir: %v",
		en:   "%v | owes to you: %v",
		es:   "%v | te debe: %v",
		faIR: "%v | به شما بدهکار است: %v",
		fr:   "%v | vous doit: %v",
		id:   "%v | berhutang kepada Anda: %v",
		it:   "%v | ti deve: %v",
		jaJP: "%v | あなたへの借り: %v",
		koKR: "%v | 당신에게 빚진 금액: %v",
		pl:   "%v | jest ci winien: %v",
		ptBR: "%v | deve a você: %v",
		ptPT: "%v | deve-lhe: %v",
		ru:   "%v | долг вам: %v",
		tr:   "%v | size borçlu: %v",
		ukUA: "%v | винен вам: %v",
		uz:   "%v | sizga qarzdor: %v",
		zhCN: "%v | 欠您: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		arEG: "نعم بالكامل",
		de:   "Ja, vollständig",
		en:   "Yes, fully",
		es:   "Sí, completamente",
		faIR: "بله، به صورت کامل",
		fr:   "Oui, entièrement",
		id:   "Ya, sepenuhnya",
		it:   "Si, completamente",
		jaJP: "はい、完全に",
		koKR: "예, 완전히",
		pl:   "Tak, w całości",
		ptBR: "Sim, totalmente",
		ptPT: "Sim, totalmente",
		ru:   "Да, целиком",
		tr:   "Evet, tamamen",
		ukUA: "Так, повністю",
		uz:   "Ha, to'liq",
		zhCN: "是的，完全",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		arEG: "لا، فقط جزئيا",
		de:   "Nein, nur teilweise",
		en:   "No, just partially",
		es:   "No, solo parcialmente",
		faIR: "خیر، تنها قسمتی",
		fr:   "Non, seulement partiellement",
		id:   "Tidak, hanya sebagian",
		it:   "No, parzialmente",
		jaJP: "いいえ、部分的に",
		koKR: "아니요, 부분적으로만",
		pl:   "Nie, tylko częściowo",
		ptBR: "Não, apenas parcialmente",
		ptPT: "Não, apenas parcialmente",
		ru:   "Нет, только часть",
		tr:   "Hayır, sadece kısmen",
		ukUA: "Ні, лише частково",
		uz:   "Yo'q, faqat qisman",
		zhCN: "不，仅部分",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		arEG: "لا يجب عليك استخدام دعوتك الخاصة ؛)",
		de:   "Du solltest dich nicht selber einladen ;)",
		en:   "You should not use your own invite ;)",
		es:   "No deberías invitarte a ti mismo ;)",
		faIR: "نباید از دعوت خود استفاده کنید ;)",
		fr:   "Vous ne devriez pas utiliser votre propre invitation ;)",
		id:   "Anda tidak boleh menggunakan undangan Anda sendiri ;)",
		it:   "Non dovresti usare il tuo codice invito con te stesso :)",
		jaJP: "自分の招待を使用すべきではありません ;)",
		koKR: "자신의 초대를 사용해서는 안 됩니다 ;)",
		pl:   "Nie powinieneś używać własnego zaproszenia ;)",
		ptBR: "Você não deve usar seu próprio convite ;)",
		ptPT: "Não deve usar o seu próprio convite ;)",
		ru:   "Хорошая попытка пригласить самого себя ;)",
		tr:   "Kendi davetinizi kullanmamalısınız ;)",
		ukUA: "Ви не повинні використовувати власне запрошення ;)",
		uz:   "Siz o'z taklifingizdan foydalanmasligingiz kerak ;)",
		zhCN: "您不应该使用自己的邀请 ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		arEG: "مرحبًا بك وشكراً لك على قبول الدعوة!",
		de:   "Willkommen. Schön, dass du der Einladung gefolgt bist!",
		en:   "Welcome and thanks for accepting the invite!",
		es:   "Bienvenido y gracias por aceptar la invitación",
		faIR: "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
		fr:   "Bienvenue et merci d'avoir accepté l'invitation!",
		id:   "Selamat datang dan terima kasih telah menerima undangan!",
		it:   "Benvenuto e grazie per aver accettato l'invito!",
		jaJP: "ようこそ、招待を受け入れていただきありがとうございます！",
		koKR: "환영합니다, 초대를 수락해 주셔서 감사합니다!",
		pl:   "Witamy i dziękujemy za przyjęcie zaproszenia!",
		ptBR: "Bem-vindo e obrigado por aceitar o convite!",
		ptPT: "Bem-vindo e obrigado por aceitar o convite!",
		ru:   "Спасибо за то что воспользовались приглашением!",
		tr:   "Hoş geldiniz ve daveti kabul ettiğiniz için teşekkürler!",
		ukUA: "Ласкаво просимо та дякуємо за прийняття запрошення!",
		uz:   "Xush kelibsiz va taklifni qabul qilganingiz uchun rahmat!",
		zhCN: "欢迎并感谢您接受邀请！",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		arEG: "هذا الإجراء لـ %v فقط.",
		de:   "Das darf nur %v.",
		en:   "This action for %v only.",
		es:   "Esta acción está disponible solo para%v.",
		faIR: "این عمل تنها برای %v می باشد.",
		fr:   "Cette action est uniquement pour %v.",
		id:   "Tindakan ini hanya untuk %v.",
		it:   "Questa azione e' disponibile solo per %v.",
		jaJP: "このアクションは%vのみです。",
		koKR: "이 작업은 %v만을 위한 것입니다.",
		pl:   "Ta akcja tylko dla %v.",
		ptBR: "Esta ação é apenas para %v.",
		ptPT: "Esta ação é válida apenas para %v.",
		ru:   "Это действие доступно только для %v.",
		tr:   "Bu işlem sadece %v için.",
		ukUA: "Ця дія тільки для %v.",
		uz:   "Bu harakat faqat %v uchun.",
		zhCN: "此操作仅适用于%v。",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		arEG: "إظهار تفاصيل الإيصال",
		de:   "Quittungsdetails anzeigen",
		en:   "Show receipt details",
		es:   "Mostrar detalles",
		faIR: "جزئیات رسید را نشان بده",
		fr:   "Afficher les détails du reçu",
		id:   "Tampilkan detail tanda terima",
		it:   "Mostra i dettagli del debito/credito",
		jaJP: "領収書の詳細を表示",
		koKR: "영수증 세부 정보 표시",
		pl:   "Pokaż szczegóły paragonu",
		ptBR: "Mostrar detalhes do recibo",
		ptPT: "Mostrar detalhes do recibo",
		ru:   "Показать детали",
		tr:   "Makbuz detaylarını göster",
		ukUA: "Показати деталі квитанції",
		uz:   "Kvitansiya tafsilotlarini ko'rsatish",
		zhCN: "显示收据详情",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		arEG: "لقد قمت باختيار دعوة صديق عبر البريد الإلكتروني.",
		de:   "Du hast gewählt, einen Freund per Mail einzuladen.",
		en:   "You've selected to invite friend by email.",
		es:   "Has decidido invitar a un amigo por e-mail.",
		faIR: "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
		fr:   "Vous avez choisi d'inviter un ami par e-mail.",
		id:   "Anda telah memilih untuk mengundang teman melalui email.",
		it:   "Hai scelto di invitare l'amico tramite email.",
		jaJP: "友達をメールで招待することを選択しました。",
		koKR: "이메일로 친구를 초대하기로 선택하셨습니다.",
		pl:   "Wybrałeś zaproszenie znajomego przez e-mail.",
		ptBR: "Você selecionou convidar um amigo por e-mail.",
		ptPT: "Selecionou convidar um amigo por e-mail.",
		ru:   "Вы решили пригласить друга через email.",
		tr:   "Arkadaşınızı e-posta ile davet etmeyi seçtiniz.",
		ukUA: "Ви вибрали запросити друга електронною поштою.",
		uz:   "Siz do'stingizni elektron pochta orqali taklif qilishni tanladingiz.",
		zhCN: "您已选择通过电子邮件邀请朋友。",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		arEG: "لقد قمت باختيار دعوة صديق عبر الرسائل النصية القصيرة.",
		de:   "Du hast gewählt, einen Freund per SMS einzuladen.",
		en:   "You've selected to invite friend by SMS.",
		es:   "Has decidido invitar a un amigo por SMS.",
		faIR: "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
		fr:   "Vous avez choisi d'inviter un ami par SMS.",
		id:   "Anda telah memilih untuk mengundang teman melalui SMS.",
		it:   "Hai scelto di invitare l'amico tramite SMS.",
		jaJP: "友達をSMSで招待することを選択しました。",
		koKR: "SMS로 친구를 초대하기로 선택하셨습니다.",
		pl:   "Wybrałeś zaproszenie znajomego przez SMS.",
		ptBR: "Você selecionou convidar um amigo por SMS.",
		ptPT: "Selecionou convidar um amigo por SMS.",
		ru:   "Вы решили пригласить друга через SMS.",
		tr:   "Arkadaşınızı SMS ile davet etmeyi seçtiniz.",
		ukUA: "Ви вибрали запросити друга через SMS.",
		uz:   "Siz do'stingizni SMS orqali taklif qilishni tanladingiz.",
		zhCN: "您已选择通过短信邀请朋友。",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		arEG: `كيف تريد تمرير رمز الدعوة؟`,
		de:   "Wie möchtest du den Code weitergeben?",
		en:   `How do you want to pass the invite code?`,
		es:   `¿Cómo quieres enviarle el código?`,
		faIR: `آیا میخواهید کد دعوت را ارسال کنید؟`,
		fr:   `Comment voulez-vous transmettre le code d'invitation?`,
		id:   `Bagaimana Anda ingin meneruskan kode undangan?`,
		it:   `Come vuoi inviargli il codice invito?`,
		jaJP: `招待コードをどのように渡しますか？`,
		koKR: `초대 코드를 어떻게 전달하시겠습니까?`,
		pl:   `Jak chcesz przekazać kod zaproszenia?`,
		ptBR: `Como você deseja passar o código de convite?`,
		ptPT: `Como pretende passar o código de convite?`,
		ru:   `Как вы хотите передать код приглашение?`,
		tr:   `Davet kodunu nasıl iletmek istiyorsunuz?`,
		ukUA: `Як ви хочете передати код запрошення?`,
		uz:   `Taklif kodini qanday o'tkazmoqchisiz?`,
		zhCN: `您想如何传递邀请码？`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		arEG: "%v تم حظر التذكيرات المتعلقة بالمعاملات بواسطة: %v",
		de:   "%v hat Erinnerungen über folgendes Anliegen blockiert: %v",
		en:   "%v blocked reminders about transactions by: %v",
		es:   "%v ha bloqueado las notificaciones de las transacciones por: %v",
		faIR: "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
		fr:   "%v a bloqué les rappels concernant les transactions par: %v",
		id:   "%v memblokir pengingat tentang transaksi oleh: %v",
		it:   "%v bloccato promemoria riguardo la transazione da: %v.",
		jaJP: "%vは%vによるトランザクションに関するリマインダーをブロックしました",
		koKR: "%v님이 %v의 거래에 대한 알림을 차단했습니다",
		pl:   "%v zablokował przypomnienia o transakcjach przez: %v",
		ptBR: "%v bloqueou lembretes sobre transações por: %v",
		ptPT: "%v lembretes bloqueados sobre transações por: %v",
		ru:   "%v заблокировал получение оповешений о транзакиях через: %v.",
		tr:   "%v, %v tarafından yapılan işlemler hakkındaki hatırlatıcıları engelledi",
		ukUA: "%v заблокував нагадування про транзакції від: %v",
		uz:   "%v %v tomonidan tranzaksiyalar haqida eslatmalarni blokladi",
		zhCN: "%v已阻止来自%v的交易提醒",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		arEG: "انتظر ثانية...",
		de:   "Warte eine Sekunde...",
		en:   "Wait a second...",
		es:   "Un segundo...",
		faIR: "یک ثانیه صبر کنید ...",
		fr:   "Attendez une seconde...",
		id:   "Tunggu sebentar...",
		it:   "Solo un attimo...",
		jaJP: "ちょっと待って...",
		koKR: "잠시만 기다려주세요...",
		pl:   "Poczekaj chwilę...",
		ptBR: "Espere um segundo...",
		ptPT: "Espere um segundo...",
		ru:   "Секундочку...",
		tr:   "Bir saniye bekleyin...",
		ukUA: "Зачекайте секунду...",
		uz:   "Bir soniya kuting...",
		zhCN: "稍等一下...",
	},
	HTML_USING_TELEGRAM: {
		arEG: "استخدام تطبيق تيليجرام ماسنجر",
		de:   "benutze Telegram messenger",
		en:   "using Telegram messenger",
		es:   "Usando Telegram",
		faIR: "استفاده از پیام رسان تلگرام",
		fr:   "utilisant Telegram messenger",
		id:   "menggunakan Telegram messenger",
		it:   "usa Telegram",
		jaJP: "Telegramメッセンジャーを使用",
		koKR: "텔레그램 메신저 사용",
		pl:   "używając komunikatora Telegram",
		ptBR: "usando o Telegram messenger",
		ptPT: "utilizando o mensageiro Telegram",
		ru:   "используя Telegram мессенджер",
		tr:   "Telegram messenger kullanarak",
		ukUA: "використовуючи Telegram messenger",
		uz:   "Telegram messenjeridan foydalanish",
		zhCN: "使用Telegram信使",
	},
	COMMAND_TEXT_ACCEPT: {
		arEG: "يقبل",
		de:   "Akzeptieren",
		en:   "Accept",
		es:   "Aceptar",
		faIR: "قبول",
		fr:   "Accepter",
		id:   "Terima",
		it:   "Accetta",
		jaJP: "承認",
		koKR: "수락",
		pl:   "Akceptuj",
		ptBR: "Aceitar",
		ptPT: "Aceitar",
		ru:   "Согласиться",
		tr:   "Kabul et",
		ukUA: "Прийняти",
		uz:   "Qabul qilish",
		zhCN: "接受",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	en: "Accept ",
	//	es: "Confirmar ",

	//	faIR: "قبول ",

	//  it:   "Accetta",
	//	ru:  "Подтвердить ",

	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	en: "Decline (using Telegram messenger)",
	//	es: "Rechazar (usandoTelegram)",

	//	faIR: "رد ( استفاده از پیام رسان تلگرام)",

	//  it:   "Rifiuta (usando Telegram)",
	//	ru:  "Отказаться (используя Telegram)",

	//},
	COMMAND_TEXT_DECLINE: {
		arEG: "انخفاض",
		de:   "Ablehnen",
		en:   "Decline",
		es:   "Rechazar",
		faIR: "رد",
		fr:   "Déclin",
		id:   "Menolak",
		it:   "Rifiuta",
		jaJP: "衰退",
		koKR: "감소",
		pl:   "Spadek",
		ptBR: "Declínio",
		ptPT: "Recusar",
		ru:   "Отклонить",
		tr:   "Reddetmek",
		ukUA: "Відхилення",
		uz:   "Rad etish",
		zhCN: "衰退"},
	FamilyMember: {
		arEG: "عضو العائلة",
		de:   "Familienmitglied",
		en:   "Family member",
		es:   "Miembro de la familia",
		faIR: "عضو خانواده",
		fr:   "Membre de la famille",
		id:   "Anggota keluarga",
		it:   "Membro della famiglia",
		jaJP: "家族の一員",
		koKR: "가족 구성원",
		pl:   "Członek rodziny",
		ptBR: "Membro da família",
		ptPT: "Membro da família",
		ru:   "Член семьи",
		tr:   "Aile üyesi",
		ukUA: "Член родини",
		uz:   "Oila aʼzosi",
		zhCN: "家庭成员",
	},
	UserHasNotJoinedSpaceYet: {
		arEG: "لم ينضم هذا الشخص إلى هذه المساحة بعد.",
		de:   "Dieser Kontakt ist diesem Bereich noch nicht beigetreten.",
		en:   "This contact has not joined this space yet.",
		es:   "Este contacto aún no se ha unido a este espacio.",
		faIR: "این مخاطب هنوز به این فضا نپیوسته است.",
		fr:   "Ce contact n'a pas encore rejoint cet espace.",
		id:   "Kontak ini belum bergabung dengan ruang ini.",
		it:   "Questo contatto non si è ancora unito a questo spazio.",
		jaJP: "この連絡先はまだこのスペースに参加していません。",
		koKR: "이 연락처는 아직 이 공간에 참여하지 않았습니다.",
		pl:   "Ten kontakt nie dołączył jeszcze do tej przestrzeni.",
		ptBR: "Este contato ainda não entrou neste espaço.",
		ptPT: "Este contacto ainda não aderiu a este espaço.",
		ru:   "Eще не присоединился к этому пространству.",
		tr:   "Bu kişi henüz bu alana katılmadı.",
		ukUA: "Цей контакт ще не приєднався до цього простору.",
		uz:   "Bu kontakt hali bu joyga qoʻshilmadi.",
		zhCN: "该联系人尚未加入此空间。",
	},
	UserHasNotJoinedFamilySpaceYet: {
		arEG: "لم ينضم إلى هذه المساحة العائلية بعد.",
		de:   "Ist diesem Familienbereich noch nicht beigetreten.",
		en:   "Has not joined this family space yet.",
		es:   "Aún no se ha unido a este espacio familiar.",
		faIR: "هنوز به این فضای خانوادگی نپیوسته است.",
		fr:   "N'a pas encore rejoint cet espace familial.",
		id:   "Belum bergabung dengan ruang keluarga ini.",
		it:   "Non si è ancora unito a questo spazio familiare.",
		jaJP: "このファミリースペースにまだ参加していません。",
		koKR: "아직 이 가족 공간에 참여하지 않았습니다.",
		pl:   "Nie dołączył jeszcze do tej przestrzeni rodzinnej.",
		ptBR: "Ainda não entrou neste espaço familiar.",
		ptPT: "Ainda não aderiu a este espaço familiar.",
		ru:   "Eще не присоединился к этому семейному пространству.",
		tr:   "Henüz bu aile alanına katılmadı.",
		ukUA: "Ще не приєднався до цього сімейного простору.",
		uz:   "Hali bu oila bo'shlig'iga qo'shilmagan.",
		zhCN: "尚未加入该家庭空间。",
	},
	BtnSendInviteByTelegram: {
		arEG: "إرسال دعوة عبر Telegram",
		de:   "Einladung über Telegram senden",
		en:   "Send invite over Telegram",
		es:   "Enviar invitación por Telegram",
		faIR: "ارسال دعوتنامه از طریق تلگرام",
		fr:   "Envoyer une invitation via Telegram",
		id:   "Kirim undangan melalui Telegram",
		it:   "Invia invito tramite Telegram",
		jaJP: "Telegramで招待を送る",
		koKR: "텔레그램으로 초대 보내기",
		pl:   "Wyślij zaproszenie przez Telegram",
		ptBR: "Enviar convite pelo Telegram",
		ptPT: "Enviar convite pelo Telegram",
		ru:   "Отправить приглашение через Telegram",
		tr:   "Telegram üzerinden davetiye gönder",
		ukUA: "Надіслати запрошення через Telegram",
		uz:   "Telegram orqali taklif yuboring",
		zhCN: "通过Telegram发送邀请",
	},
	BtnTextAcceptInvite: {
		arEG: "قبول الدعوة",
		de:   "Akzeptiere Einladung",
		en:   "Accept invite",
		es:   "Aceptar la invitación",
		faIR: "قبول دعوت",
		fr:   "Accepter l'invitation",
		id:   "Terima undangan",
		it:   "Accetta l'invito",
		jaJP: "招待を承認",
		koKR: "초대 수락",
		pl:   "Zaakceptuj zaproszenie",
		ptBR: "Aceitar convite",
		ptPT: "Aceitar convite",
		ru:   "Принять приглашение",
		tr:   "Daveti kabul et",
		ukUA: "Прийняти запрошення",
		uz:   "Taklifni qabul qilish",
		zhCN: "接受邀请",
	},
	BtnTextDeclineInvite: {
		arEG: "رفض الدعوة",
		de:   "Ablehnen Einladung",
		en:   "Decline invite",
		es:   "Rechazar la invitación",
		faIR: "رد دعوت",
		fr:   "Décliner l'invitation",
		id:   "Tolak undangan",
		it:   "Rifiuta l'invito",
		jaJP: "招待を拒否",
		koKR: "초대 거절",
		pl:   "Odrzuć zaproszenie",
		ptBR: "Recusar convite",
		ptPT: "Recusar convite",
		ru:   "Отклонить приглашение",
		tr:   "Daveti reddet",
		ukUA: "Відхилити запрошення",
		uz:   "Taklifni rad etish",
		zhCN: "拒绝邀请",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		arEG: "انظر تفاصيل الإيصال",
		de:   "Quittungsdetails anzeigen",
		en:   "See receipt details",
		es:   "Ver el recibo",
		faIR: "دیدن جزئیات رسید",
		fr:   "Voir les détails du reçu",
		id:   "Lihat detail tanda terima",
		it:   "Vedi dettagli",
		jaJP: "領収書の詳細を見る",
		koKR: "영수증 세부 정보 보기",
		pl:   "Zobacz szczegóły paragonu",
		ptBR: "Ver detalhes do recibo",
		ptPT: "Ver detalhes do recibo",
		ru:   "Посмотреть квитанцию",
		tr:   "Makbuz detaylarını görüntüle",
		ukUA: "Переглянути квитанцію",
		uz:   "Kvitansiya tafsilotlarini ko'rish",
		zhCN: "查看收据详情",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		arEG: "طرق أخرى لإرسال دعوة",
		de:   "Andere Wege, eine Einladung zu senden",
		en:   "Other ways to send an invite",
		es:   "Otras maneras para enviar la invitación",
		faIR: "سایر راههای ارسال دعوت",
		fr:   "Autres façons d'envoyer une invitation",
		id:   "Cara lain untuk mengirim undangan",
		it:   "Altri modi per inviare un invito",
		jaJP: "招待状を送信する他の方法",
		koKR: "초대장을 보내는 다른 방법",
		pl:   "Inne sposoby wysyłania zaproszenia",
		ptBR: "Outras formas de enviar um convite",
		ptPT: "Outras formas de enviar um convite",
		ru:   "Другие способы отправить приглашение",
		tr:   "Davetiye göndermenin diğer yolları",
		ukUA: "Інші способи надіслати запрошення",
		uz:   "Taklifnoma yuborishning boshqa usullari",
		zhCN: "发送邀请的其他方式",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		arEG: "أرسل رقم هاتفي",
		de:   "meine Telefonnummer senden",
		en:   "Send my phone number",
		es:   "Enviar mi número",
		faIR: "شماره تلفن مرا ارسال کنید",
		fr:   "Envoyer mon numéro de téléphone",
		id:   "Kirim nomor telepon saya",
		it:   "Invia il mio numero",
		jaJP: "私の電話番号を送信する",
		koKR: "내 전화번호 보내기",
		pl:   "Wyślij mój numer telefonu",
		ptBR: "Enviar meu número de telefone",
		ptPT: "Enviar o meu número de telefone",
		ru:   "Отправить мой номер",
		tr:   "Telefon numaramı gönder",
		ukUA: "Надіслати мій номер телефону",
		uz:   "Telefon raqamimni yuborish",
		zhCN: "发送我的电话号码",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		arEG: "عبر البريد الإلكتروني",
		de:   "per Mail",
		en:   "By Email",
		es:   "Vía e-mail",
		faIR: "بوسیله ی ایمیل",
		fr:   "Par e-mail",
		id:   "Melalui Email",
		it:   "Tramite email",
		jaJP: "メールで",
		koKR: "이메일로",
		pl:   "Przez e-mail",
		ptBR: "Por e-mail",
		ptPT: "Por e-mail",
		ru:   "Через Email",
		tr:   "E-posta ile",
		ukUA: "Через електронну пошту",
		uz:   "Email orqali",
		zhCN: "通过电子邮件",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		arEG: "عن طريق الرسائل القصيرة",
		de:   "per SMS",
		en:   "By SMS",
		es:   "Vía SMS",
		faIR: "بوسیله پیام کوتاه",
		fr:   "Par SMS",
		id:   "Melalui SMS",
		it:   "Tramite SMS",
		jaJP: "SMSで",
		koKR: "SMS로",
		pl:   "Przez SMS",
		ptBR: "Por SMS",
		ptPT: "Por SMS",
		ru:   "Через SMS",
		tr:   "SMS ile",
		ukUA: "Через SMS",
		uz:   "SMS orqali",
		zhCN: "通过短信",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		arEG: "دعوة عبر تيليجرام",
		de:   "Einladen per Telegram",
		en:   "Invite By Telegram",
		es:   "Invitar vía Telegram",
		faIR: "دعوت با تلگرام",
		fr:   "Inviter par Telegram",
		id:   "Undang Melalui Telegram",
		it:   "Tramite Telegram",
		jaJP: "Telegramで招待する",
		koKR: "텔레그램으로 초대하기",
		pl:   "Zaproś przez Telegram",
		ptBR: "Convidar pelo Telegram",
		ptPT: "Convite por Telegram",
		ru:   "Пригласить через Telegram",
		tr:   "Telegram ile davet et",
		ukUA: "Запросити через Telegram",
		uz:   "Telegram orqali taklif qilish",
		zhCN: "通过Telegram邀请",
	},
	MESSAGE_TEXT_INVITE_CREATED: {
		arEG: `لقد أرسلنا رمز دعوة إلى صديقك. (#%v)

بمجرد أن يقبل صديقك الدعوة، ستقوم بمشاركة الرصيد والتحويلات بينكما للتأكد من أنكما على نفس الصفحة مع الحد الأدنى من الجهود.`,
		de: `Wir haben deinen Freund einen Code geschickt. (#%v)

Sobald dein Freund die Einladung akzeptiert hat, könnt ihr das Geld, was ihr euch teit, mit Leichtigkeit managen.`,
		en: `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,
		es: `Hemos enviado el código de la invitación a tu amigo. (#%v)

Cuando tu amigo accepte la invitación vaís a tener transacciones y balance en común (solo entre vosotros). Todo eso os ayuda minimizar los esfuerzos para controlar la cuenta.`,
		faIR: `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,
		fr: `Nous avons envoyé un code d'invitation à votre ami. (#%v)

Une fois que votre ami accepte l'invitation, vous partagerez le solde et les transferts entre vous pour vous assurer que vous êtes tous les deux sur la même page avec un minimum d'efforts.`,
		id: `Kami telah mengirimkan kode undangan kepada teman Anda. (#%v)

Setelah teman Anda menerima undangan, Anda akan berbagi saldo & transfer antara Anda untuk memastikan Anda berdua berada di halaman yang sama dengan upaya minimal.`,
		it: `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,
		jaJP: `友達に招待コードを送信しました。(#%v)

友達が招待を受け入れると、最小限の労力で両方が同じページにいることを確認するために、あなたの間で残高と転送を共有します。`,
		koKR: `친구에게 초대 코드를 보냈습니다. (#%v)

친구가 초대를 수락하면 최소한의 노력으로 두 사람이 같은 페이지에 있는지 확인하기 위해 잔액과 이체를 공유하게 됩니다.`,
		pl: `Wysłaliśmy kod zaproszenia do Twojego znajomego. (#%v)

Gdy Twój znajomy zaakceptuje zaproszenie, będziecie dzielić saldo i przelewy między sobą, aby upewnić się, że oboje jesteście na tej samej stronie przy minimalnym wysiłku.`,
		ptBR: `Enviamos um código de convite para seu amigo. (#%v)

Quando seu amigo aceitar o convite, vocês compartilharão saldo e transferências entre si para garantir que ambos estejam na mesma página com o mínimo de esforço.`,
		ptPT: `Enviamos um código de convite ao seu amigo. (#%v)

Assim que o seu amigo aceitar o convite, irão dividir o saldo e as transferências entre si para garantir que ambos estão na mesma página com o mínimo de esforço.`,
		ru: `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,
		tr: `Arkadaşınıza bir davet kodu gönderdik. (#%v)

Arkadaşınız daveti kabul ettiğinde, minimum çabayla ikinizin de aynı sayfada olduğundan emin olmak için aranızda bakiye ve transferleri paylaşacaksınız.`,
		ukUA: `Ми надіслали код запрошення вашому другу. (#%v)

Коли ваш друг прийме запрошення, ви будете ділитися балансом і переказами між собою, щоб переконатися, що ви обидва на одній сторінці з мінімальними зусиллями.`,
		uz: `Do'stingizga taklifnoma kodini yubordik. (#%v)

Do'stingiz taklifni qabul qilgandan so'ng, minimal kuch bilan ikkalangiz ham bir xil sahifada ekanligingizga ishonch hosil qilish uchun o'zaro balans va o'tkazmalarni almashishingiz mumkin.`,
		zhCN: `我们已向您的朋友发送邀请码。(#%v)

一旦您的朋友接受邀请，您将共享余额和转账，以确保您们双方以最少的努力保持一致。`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		arEG: "الرجاء إدخال عنوان البريد الإلكتروني لصديقك حيث يجب أن نرسل رمز الدعوة إليه.",
		de:   "Bitte gib die eMail Adresse deines Freundes an, wohin wir den Code schicken:",
		en:   "Please enter email address of your friend where we should send an invite code.",
		es:   "Por favor, introduce el e-maill de tu amigo al cual enviaremos el código de la invitación.",
		faIR: "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
		fr:   "Veuillez entrer l'adresse e-mail de votre ami où nous devrions envoyer un code d'invitation.",
		id:   "Silakan masukkan alamat email teman Anda di mana kami harus mengirimkan kode undangan.",
		it:   "Inserisci l'email dell'amico al quale inviare il codice invito.",
		jaJP: "招待コードを送信する友達のメールアドレスを入力してください。",
		koKR: "초대 코드를 보낼 친구의 이메일 주소를 입력하세요.",
		pl:   "Proszę podać adres e-mail znajomego, na który powinniśmy wysłać kod zaproszenia.",
		ptBR: "Por favor, insira o endereço de e-mail do seu amigo para onde devemos enviar um código de convite.",
		ptPT: "Por favor, introduza o endereço de e-mail do seu amigo para onde devemos enviar um código de convite.",
		ru:   "Введите email вашего друга на который мы отправим код приглашения.",
		tr:   "Lütfen davet kodunu göndermemiz gereken arkadaşınızın e-posta adresini girin.",
		ukUA: "Будь ласка, введіть електронну адресу вашого друга, куди ми повинні надіслати код запрошення.",
		uz:   "Iltimos, taklifnoma kodini yuborishimiz kerak bo'lgan do'stingizning elektron pochta manzilini kiriting.",
		zhCN: "请输入您朋友的电子邮件地址，我们将向其发送邀请码。",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		arEG: "الرجاء إدخال عنوان البريد الإلكتروني لصديقك (%v) الذي يجب أن نرسل إليه الإيصال.",
		de:   "Bitte gib die eMail Adresse deines Freundes (%v), wohin wir die Quittung schicken:",
		en:   "Please enter email address of your friend (%v) where we should send the receipt.",
		es:   "Introduce el e-maill de tu amigo (%v) al cual enviaremos el recibo.",
		faIR: "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
		fr:   "Veuillez entrer l'adresse e-mail de votre ami (%v) où nous devrions envoyer le reçu.",
		id:   "Silakan masukkan alamat email teman Anda (%v) di mana kami harus mengirimkan tanda terima.",
		it:   "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
		jaJP: "領収書を送信する友達 (%v) のメールアドレスを入力してください。",
		koKR: "영수증을 보낼 친구 (%v)의 이메일 주소를 입력하세요.",
		pl:   "Proszę podać adres e-mail znajomego (%v), na który powinniśmy wysłać paragon.",
		ptBR: "Por favor, insira o endereço de e-mail do seu amigo (%v) para onde devemos enviar o recibo.",
		ptPT: "Por favor, introduza o endereço de e-mail do seu amigo (%v) para onde devemos enviar o recibo.",
		ru:   "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		tr:   "Lütfen makbuzu göndermemiz gereken arkadaşınızın (%v) e-posta adresini girin.",
		ukUA: "Будь ласка, введіть електронну адресу вашого друга (%v), куди ми повинні надіслати квитанцію.",
		uz:   "Iltimos, kvitansiyani yuborishimiz kerak bo'lgan do'stingizning (%v) elektron pochta manzilini kiriting.",
		zhCN: "请输入您朋友 (%v) 的电子邮件地址，我们将向其发送收据。",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		arEG: "الرجاء إدخال رقم صديقك الذي سنرسل له دعوة.",
		de:   "Bitte gib die Telefonnummer deines Freundes an oder teile einen Kontakt, wohin wir den Code schicken:",
		en:   "Please enter number of your friend where we'll send an invite.",
		es:   "Por favor, introduce el número del teléfono de tu amigo al cual enviaremos el código de la invitación.",
		faIR: "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
		fr:   "Veuillez entrer le numéro de votre ami où nous enverrons une invitation.",
		id:   "Silakan masukkan nomor teman Anda di mana kami akan mengirimkan undangan.",
		it:   "Condividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
		jaJP: "招待状を送信する友達の番号を入力してください。",
		koKR: "초대장을 보낼 친구의 번호를 입력하세요.",
		pl:   "Proszę podać numer telefonu znajomego, na który wyślemy zaproszenie.",
		ptBR: "Por favor, insira o número do seu amigo para onde enviaremos um convite.",
		ptPT: "Por favor, insira o número do seu amigo a quem enviaremos o convite.",
		ru:   "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		tr:   "Lütfen davet göndereceğimiz arkadaşınızın numarasını girin.",
		ukUA: "Будь ласка, введіть номер вашого друга, куди ми надішлемо запрошення.",
		uz:   "Iltimos, taklifnoma yuboradigan do'stingizning raqamini kiriting.",
		zhCN: "请输入我们将发送邀请的朋友的号码。",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		arEG: "يرجى مشاركة جهة اتصال صديقك الذي ترغب في إرسال رمز دعوة إليه:",
		de:   "Bitte wähle einen Kontakt, welchen du den Code schicken willst:",
		en:   "Please share a contact of your friend you wish to send an invite code:",
		es:   "Por favor, comparte el contacto de tu amigo al cual quieres enviar el código de la invitación.",
		faIR: "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
		fr:   "Veuillez partager un contact de votre ami à qui vous souhaitez envoyer un code d'invitation :",
		id:   "Silakan bagikan kontak teman Anda yang ingin Anda kirimi kode undangan:",
		it:   "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
		jaJP: "招待コードを送信したい友達の連絡先を共有してください：",
		koKR: "초대 코드를 보내고 싶은 친구의 연락처를 공유하세요:",
		pl:   "Proszę udostępnić kontakt znajomego, któremu chcesz wysłać kod zaproszenia:",
		ptBR: "Por favor, compartilhe um contato do seu amigo para quem deseja enviar um código de convite:",
		ptPT: "Por favor, partilhe o contacto de um amigo a quem pretende enviar um código de convite:",
		ru:   "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		tr:   "Lütfen davet kodu göndermek istediğiniz arkadaşınızın bir kişisini paylaşın:",
		ukUA: "Будь ласка, поділіться контактом вашого друга, якому ви бажаєте надіслати код запрошення:",
		uz:   "Iltimos, taklifnoma kodini yubormoqchi bo'lgan do'stingizning kontaktini ulashing:",
		zhCN: "请分享您希望发送邀请码的朋友的联系人：",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		arEG: "بريد إلكتروني غير صالح. تحقق وحاول مرة أخرى؟ /القائمة",
		de:   "Ungültige eMail Adresse. Versuche es erneut oder gehe ins Haupt /menu",
		en:   "Invalid email. Check and try it again? /menu",
		es:   "Email incorrecto. ¿Comprobarlo e intentalo de nuevo? /menú",
		faIR: "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟ /منو",
		fr:   "Email invalide. Vérifiez et réessayez ? /menu",
		id:   "Email tidak valid. Periksa dan coba lagi? /menu",
		it:   "Email scorretta. Controlla e riprova. /menu",
		jaJP: "無効なメール。確認して再試行しますか？ /menu",
		koKR: "잘못된 이메일입니다. 확인하고 다시 시도하시겠습니까? /menu",
		pl:   "Nieprawidłowy email. Sprawdź i spróbuj ponownie? /menu",
		ptBR: "Email inválido. Verificar e tentar novamente? /menu",
		ptPT: "E-mail inválido. Verificar e tentar novamente? /menu",
		ru:   "Неверный email. Проверить и попробовать ещё раз? /menu",
		tr:   "Geçersiz e-posta. Kontrol edip tekrar deneyin? /menu",
		ukUA: "Недійсна електронна пошта. Перевірити та спробувати ще раз? /menu",
		uz:   "Yaroqsiz elektron pochta. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		zhCN: "无效的电子邮件。检查并重试？ /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		arEG: "سنة غير صحيحة. يجب أن يتكون جزء السنة من رقمين أو أربعة أرقام ( <i>مثل ٢٠١٦ أو ١٦</i> ).",
		de:   "Ungültiges Jahr. Der Jahresangabe sollte 2 oder 4 Ziffern sein (<i>z.B. 2016 or 16</i>).",
		en:   "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		es:   "Año incorrecto. El año tiene que constar de 2 o 4 números (<i>ejemplo 2016 o 16</i>).",
		faIR: "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
		fr:   "Année invalide. L'année doit comporter 2 ou 4 chiffres (<i>par exemple 2016 ou 16</i>).",
		id:   "Tahun tidak valid. Bagian tahun harus berupa 2 atau 4 angka (<i>misalnya 2016 atau 16</i>).",
		it:   "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
		jaJP: "無効な年。年の部分は2桁または4桁の数字である必要があります (<i>例：2016または16</i>)。",
		koKR: "잘못된 연도입니다. 연도 부분은 2자리 또는 4자리 숫자여야 합니다 (<i>예: 2016 또는 16</i>).",
		pl:   "Nieprawidłowy rok. Część roku powinna składać się z 2 lub 4 cyfr (<i>np. 2016 lub 16</i>).",
		ptBR: "Ano inválido. A parte do ano deve ter 2 ou 4 números (<i>por exemplo, 2016 ou 16</i>).",
		ptPT: "Ano inválido. A parte do ano deve ter 2 ou 4 números ( <i>ex.: 2016 ou 16</i> ).",
		ru:   "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		tr:   "Geçersiz yıl. Yıl kısmı 2 veya 4 rakam olmalıdır (<i>örneğin 2016 veya 16</i>).",
		ukUA: "Недійсний рік. Частина року повинна складатися з 2 або 4 цифр (<i>наприклад, 2016 або 16</i>).",
		uz:   "Yaroqsiz yil. Yil qismi 2 yoki 4 raqamdan iborat bo'lishi kerak (<i>masalan, 2016 yoki 16</i>).",
		zhCN: "无效的年份。年份部分应为2位或4位数字 (<i>例如2016或16</i>)。",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		arEG: "شهر غير صالح. يجب أن يكون الشهر عددًا صحيحًا من ١ إلى ١٢.",
		de:   "Ungültiger Monat. Der Monatsangabe sollte eine Ganzzahl zwischen 1 und 12 sein.",
		en:   "Invalid month. Month should be an integer from 1 to 12.",
		es:   "El mes es incorrecto. El mes hay que introducirlo del 1 al 12.",
		faIR: "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
		fr:   "Mois invalide. Le mois doit être un nombre entier de 1 à 12.",
		id:   "Bulan tidak valid. Bulan harus berupa bilangan bulat dari 1 hingga 12.",
		it:   "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
		jaJP: "無効な月。月は1から12までの整数である必要があります。",
		koKR: "잘못된 월입니다. 월은 1에서 12 사이의 정수여야 합니다.",
		pl:   "Nieprawidłowy miesiąc. Miesiąc powinien być liczbą całkowitą od 1 do 12.",
		ptBR: "Mês inválido. O mês deve ser um número inteiro de 1 a 12.",
		ptPT: "Mês inválido. O mês deve ser um número inteiro de 1 a 12.",
		ru:   "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		tr:   "Geçersiz ay. Ay, 1'den 12'ye kadar bir tam sayı olmalıdır.",
		ukUA: "Недійсний місяць. Місяць повинен бути цілим числом від 1 до 12.",
		uz:   "Yaroqsiz oy. Oy 1 dan 12 gacha bo'lgan butun son bo'lishi kerak.",
		zhCN: "无效的月份。月份应为1到12之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		arEG: "يوم غير صالح. يجب أن يكون اليوم عددًا صحيحًا من ١ إلى ٣١.",
		de:   "Ungültiger Tag. Der Tagesangabe sollte eine Ganzzahl zwischen 1 und 31 sein.",
		en:   "Invalid day. The day should be an integer from 1 to 31.",
		es:   "El día es incorrecto. El día hay que introducirlo del 1 al 31.",
		faIR: "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
		fr:   "Jour invalide. Le jour doit être un nombre entier de 1 à 31.",
		id:   "Hari tidak valid. Hari harus berupa bilangan bulat dari 1 hingga 31.",
		it:   "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
		jaJP: "無効な日。日は1から31までの整数である必要があります。",
		koKR: "잘못된 일입니다. 일은 1에서 31 사이의 정수여야 합니다.",
		pl:   "Nieprawidłowy dzień. Dzień powinien być liczbą całkowitą od 1 do 31.",
		ptBR: "Dia inválido. O dia deve ser um número inteiro de 1 a 31.",
		ptPT: "Dia inválido. O dia deve ser um número inteiro de 1 a 31.",
		ru:   "Неверный день. День должен быть задан целым числом от 1 до 31.",
		tr:   "Geçersiz gün. Gün, 1'den 31'e kadar bir tam sayı olmalıdır.",
		ukUA: "Недійсний день. День повинен бути цілим числом від 1 до 31.",
		uz:   "Yaroqsiz kun. Kun 1 dan 31 gacha bo'lgan butun son bo'lishi kerak.",
		zhCN: "无效的日期。日期应为1到31之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		arEG: "تنسيق التاريخ غير صحيح. على سبيل المثال، لتاريخ ٢٠ فبراير ٢٠١٩، يُرجى إدخال: ٢٠ فبراير ٢٠١٩ أو ٢٠ فبراير ٢٠١٩",
		de:   "Ungültiges Datumsformat. Zum Beispiel für den 20. February 2019 sende: 20.02.2019 oder 20.02.19",
		en:   "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		es:   "El formato de la fecha no es correcto. Por ejemplo para el día 20 de Febrero de 2019 introduce: 20.02.2019 o 20.02.19",
		faIR: "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
		fr:   "Format de date invalide. Par exemple, pour le 20 février 2019, veuillez soumettre : 20.02.2019 ou 20.02.19",
		id:   "Format tanggal tidak valid. Misalnya untuk 20 Februari 2019 silakan kirimkan: 20.02.2019 atau 20.02.19",
		it:   "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
		jaJP: "無効な日付形式です。例えば、2019年2月20日の場合は、20.02.2019または20.02.19と入力してください。",
		koKR: "잘못된 날짜 형식입니다. 예를 들어 2019년 2월 20일의 경우 다음과 같이 제출하세요: 20.02.2019 또는 20.02.19",
		pl:   "Nieprawidłowy format daty. Na przykład dla 20 lutego 2019 proszę podać: 20.02.2019 lub 20.02.19",
		ptBR: "Formato de data inválido. Por exemplo, para 20 de fevereiro de 2019, envie: 20.02.2019 ou 20.02.19",
		ptPT: "Formato de data inválido. Por exemplo, para 20 de fevereiro de 2019, envie: 20/02/2019 ou 20/02/19.",
		ru:   "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		tr:   "Geçersiz tarih formatı. Örneğin 20 Şubat 2019 için lütfen şunu girin: 20.02.2019 veya 20.02.19",
		ukUA: "Недійсний формат дати. Наприклад, для 20 лютого 2019 року, будь ласка, введіть: 20.02.2019 або 20.02.19",
		uz:   "Yaroqsiz sana formati. Masalan, 2019 yil 20 fevral uchun quyidagilarni yuboring: 20.02.2019 yoki 20.02.19",
		zhCN: "无效的日期格式。例如，对于2019年2月20日，请提交：20.02.2019或20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		arEG: "رقم الهاتف غير صحيح. تحقق وحاول مرة أخرى؟ /القائمة",
		de:   "Ungültige Telefonnummer. Versuche es erneut oder gehe ins Haupt /menu",
		en:   "Invalid phone number. Check and try it again? /menu",
		es:   "El número del teléfono no es correcto. ¿Comprobarlo y intentarlo de nuevo? /menú",
		faIR: "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟ /منو",
		fr:   "Numéro de téléphone invalide. Vérifiez et réessayez ? /menu",
		id:   "Nomor telepon tidak valid. Periksa dan coba lagi? /menu",
		it:   "Numero di telefono invalido. Controlla e riprova. /menu",
		jaJP: "無効な電話番号です。確認して再試行しますか？ /menu",
		koKR: "잘못된 전화번호입니다. 확인하고 다시 시도하시겠습니까? /menu",
		pl:   "Nieprawidłowy numer telefonu. Sprawdź i spróbuj ponownie? /menu",
		ptBR: "Número de telefone inválido. Verificar e tentar novamente? /menu",
		ptPT: "Número de telefone inválido. Verificar e tentar novamente? /menu",
		ru:   "Неверный номер. Проверить и попробовать ещё раз? /menu",
		tr:   "Geçersiz telefon numarası. Kontrol edip tekrar deneyin? /menu",
		ukUA: "Недійсний номер телефону. Перевірити та спробувати ще раз? /menu",
		uz:   "Yaroqsiz telefon raqami. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		zhCN: "无效的电话号码。检查并重试？ /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		arEG: "هذا الرقم غير قادر على استقبال الرسائل النصية. هل تريد تجربة رقم آخر؟ /القائمة",
		de:   "Diese Telefonnummer kann von uns keine SMS empfangen. Versuche eine andere oder gehe ins Haupt /menu",
		en:   "This phone number not able to receive SMS. Try another number? /menu",
		es:   "Este número de teléfono no acepta SMS. ¿Intentar otro número? /menú",
		faIR: "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟ /منو",
		fr:   "Ce numéro de téléphone ne peut pas recevoir de SMS. Essayer un autre numéro ? /menu",
		id:   "Nomor telepon ini tidak dapat menerima SMS. Coba nomor lain? /menu",
		it:   "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
		jaJP: "この電話番号はSMSを受信できません。別の番号を試しますか？ /menu",
		koKR: "이 전화번호는 SMS를 수신할 수 없습니다. 다른 번호를 시도하시겠습니까? /menu",
		pl:   "Ten numer telefonu nie może odbierać SMS-ów. Spróbować innego numeru? /menu",
		ptBR: "Este número de telefone não pode receber SMS. Tentar outro número? /menu",
		ptPT: "Este número de telefone não aceita SMS. Tentar outro número? /menu",
		ru:   "Данный номер не принимает SMS. Попробовать другой номер? /menu",
		tr:   "Bu telefon numarası SMS alamıyor. Başka bir numara deneyin? /menu",
		ukUA: "Цей номер телефону не може отримувати SMS. Спробувати інший номер? /menu",
		uz:   "Bu telefon raqami SMS qabul qila olmaydi. Boshqa raqamni sinab ko'rasizmi? /menu",
		zhCN: "此电话号码无法接收短信。尝试其他号码？ /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		arEG: "لم نتلقَّ أي اتصالات. تعليمات كيفية القيام بذلك. /القائمة",
		de:   "Wir haben keine Kontakte empfangen. Versuche es erneut oder gehe ins Haupt /menu",
		en:   "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		es:   "No hemos recibido ningún contacto. LA INSTRUCCIÓN COMO HACERLO. /menú",
		faIR: "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار. /منو",
		fr:   "Nous n'avons reçu aucun contact. INSTRUCTION COMMENT LE FAIRE. /menu",
		id:   "Kami belum menerima kontak apa pun. INSTRUKSI CARA MELAKUKANNYA. /menu",
		it:   "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
		jaJP: "連絡先を受信していません。操作方法の説明。 /menu",
		koKR: "연락처를 받지 못했습니다. 방법 설명. /menu",
		pl:   "Nie otrzymaliśmy żadnych kontaktów. INSTRUKCJA JAK TO ZROBIĆ. /menu",
		ptBR: "Não recebemos nenhum contato. INSTRUÇÃO COMO FAZER. /menu",
		ptPT: "Não recebemos qualquer contacto. INSTRUÇÕES DE COMO FAZER. /menu",
		ru:   "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		tr:   "Herhangi bir kişi almadık. NASIL YAPILACAĞI TALİMATI. /menu",
		ukUA: "Ми не отримали жодних контактів. ІНСТРУКЦІЯ ЯК ЦЕ ЗРОБИТИ. /menu",
		uz:   "Biz hech qanday kontakt olmadik. BUNI QANDAY QILISH BO'YICHA KO'RSATMA. /menu",
		zhCN: "我们尚未收到任何联系人。操作指南。 /menu",
	},
	MESSAGE_TEXT_YOU_HAVE_NO_CONTACTS: {
		arEG: "لم تقم بإنشاء أي جهات اتصال بعد.",
		de:   "Du hast noch keine Kontakte hinzugefügt.",
		en:   "You have not created any contacts yet.",
		es:   "Todavía no has creado ningún contacto.",
		faIR: "هنوز هیچ مخاطبی را ایجاد نکرده اید",
		fr:   "Vous n'avez pas encore créé de contacts.",
		id:   "Anda belum membuat kontak apa pun.",
		it:   "Non hai ancora creato alcun contatto.",
		jaJP: "まだ連絡先を作成していません。",
		koKR: "아직 연락처를 만들지 않았습니다.",
		pl:   "Nie utworzyłeś jeszcze żadnych kontaktów.",
		ptBR: "Você ainda não criou nenhum contato.",
		ptPT: "Ainda não criou nenhum contacto.",
		ru:   "Вы ещё не создали контактов",
		tr:   "Henüz hiç kişi oluşturmadınız.",
		ukUA: "Ви ще не створили жодних контактів.",
		uz:   "Siz hali hech qanday kontakt yaratmagansiz.",
		zhCN: "您尚未创建任何联系人。",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		arEG: "لقد أدخلت أرقامًا فقط لاسم جهة الاتصال. يُرجى استخدام بعض الأحرف النصية.",
		de:   "Du kannst nicht nur Zahlen für einen Kontaktnamen eingeben. Bitte benutz ein paar Buchstaben - du kannst mir vertrauen.",
		en:   "You've entered just digits for a contact name. Please use some text characters.",
		es:   "Has introducido solo números para el nombre del contacto. Por favor usa algunas letras.",
		faIR: "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
		fr:   "Vous avez saisi uniquement des chiffres pour un nom de contact. Veuillez utiliser des caractères textuels.",
		id:   "Anda hanya memasukkan angka untuk nama kontak. Harap gunakan beberapa karakter teks.",
		it:   "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
		jaJP: "連絡先名に数字のみを入力しました。テキスト文字を使用してください。",
		koKR: "연락처 이름에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		pl:   "Wprowadziłeś same cyfry jako nazwę kontaktu. Proszę użyć znaków tekstowych.",
		ptBR: "Você digitou apenas dígitos para um nome de contato. Por favor, use alguns caracteres de texto.",
		ptPT: "Apenas introduziu dígitos para o nome do contacto. Por favor, utilize alguns caracteres de texto.",
		ru:   "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		tr:   "Bir kişi adı için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		ukUA: "Ви ввели лише цифри для імені контакту. Будь ласка, використовуйте текстові символи.",
		uz:   "Kontakt nomi uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		zhCN: "您只输入了数字作为联系人名称。请使用一些文本字符。",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		arEG: "لقد أدخلت أرقامًا فقط للعملة. يُرجى استخدام بعض الأحرف النصية.",
		de:   "Bei der Währung erwarte ich eigentlich keine Zahlen. Nimm ein paar Buchstaben hinzu, um Verwirrung zu vermeiden.",
		en:   "You've entered just digits for currency. Please use some text characters.",
		es:   "Has introducido solo números para la moneda. Por favor usa algunas letras.",
		faIR: "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
		fr:   "Vous avez saisi uniquement des chiffres pour la devise. Veuillez utiliser des caractères textuels.",
		id:   "Anda hanya memasukkan angka untuk mata uang. Harap gunakan beberapa karakter teks.",
		it:   "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
		jaJP: "通貨に数字のみを入力しました。テキスト文字を使用してください。",
		koKR: "통화에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		pl:   "Wprowadziłeś same cyfry dla waluty. Proszę użyć znaków tekstowych.",
		ptBR: "Você digitou apenas dígitos para moeda. Por favor, use alguns caracteres de texto.",
		ptPT: "Apenas introduziu dígitos para a moeda. Por favor, utilize alguns caracteres de texto.",
		ru:   "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		tr:   "Para birimi için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		ukUA: "Ви ввели лише цифри для валюти. Будь ласка, використовуйте текстові символи.",
		uz:   "Valyuta uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		zhCN: "您只输入了数字作为货币。请使用一些文本字符。",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		arEG: "%v - %s ⇒ لك: %s",
		de:   "%v - %s ⇒ dir: %s",
		en:   "%v - %s ⇒ to you: %s",
		es:   "%v - %s ⇒ a ti: %s",
		faIR: "%v - %s ⇒ به شما: %s",
		fr:   "%v - %s ⇒ à vous: %s",
		id:   "%v - %s ⇒ kepada Anda: %s",
		it:   "%v - %s ⇒ a te: %s",
		jaJP: "%v - %s ⇒ あなたへ: %s",
		koKR: "%v - %s ⇒ 당신에게: %s",
		pl:   "%v - %s ⇒ do Ciebie: %s",
		ptBR: "%v - %s ⇒ para você: %s",
		ptPT: "%v - %s ⇒ para si: %s",
		ru:   "%v - %s ⇒ Вам : %s",
		tr:   "%v - %s ⇒ size: %s",
		ukUA: "%v - %s ⇒ Вам: %s",
		uz:   "%v - %s ⇒ sizga: %s",
		zhCN: "%v - %s ⇒ 给您: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		arEG: "%v - أنت ⇒ %s : %s",
		de:   "%v - Du ⇒ %s : %s",
		en:   "%v - You ⇒ %s : %s",
		es:   "%v - Tú ⇒ %s : %s",
		faIR: "%v - شما ⇒ %s : %s",
		fr:   "%v - Vous ⇒ %s : %s",
		id:   "%v - Anda ⇒ %s : %s",
		it:   "%v - Tu ⇒ %s : %s",
		jaJP: "%v - あなた ⇒ %s : %s",
		koKR: "%v - 당신 ⇒ %s : %s",
		pl:   "%v - Ty ⇒ %s : %s",
		ptBR: "%v - Você ⇒ %s : %s",
		ptPT: "%v - Você ⇒ %s : %s",
		ru:   "%v - Вы ⇒ %s : %s",
		tr:   "%v - Siz ⇒ %s : %s",
		ukUA: "%v - Ви ⇒ %s : %s",
		uz:   "%v - Siz ⇒ %s : %s",
		zhCN: "%v - 您 ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		arEG: "دعونا نرسل الرسائل القصيرة",
		de:   "Lass uns eine SMS senden",
		en:   "Let's send SMS",
		es:   "Vamos a enviar un SMS",
		faIR: "پیام کوتاه ارسال کنید",
		fr:   "Envoyons un SMS",
		id:   "Mari kirim SMS",
		it:   "Inviamo un SMS",
		jaJP: "SMSを送信しましょう",
		koKR: "SMS를 보내봅시다",
		pl:   "Wyślijmy SMS",
		ptBR: "Vamos enviar SMS",
		ptPT: "Vamos enviar SMS",
		ru:   "Давайте отправим SMS",
		tr:   "SMS gönderelim",
		ukUA: "Давайте надішлемо SMS",
		uz:   "SMS yuboraylik",
		zhCN: "让我们发送短信",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		arEG: "انتظار إرسال الرسائل النصية القصيرة إلى الرقم %v...",
		de:   "Setze SMS in Sendewarteschlange für %v...",
		en:   "Queuing SMS for sending to number %v...",
		es:   "El SMS se está poniendo en la cola para enviar al número %v...",
		faIR: "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
		fr:   "Mise en file d'attente du SMS pour envoi au numéro %v...",
		id:   "Antrian SMS untuk dikirim ke nomor %v...",
		it:   "SMS in coda per l'invio al numero %v...",
		jaJP: "番号 %v に送信するためのSMSをキューに入れています...",
		koKR: "번호 %v로 보내기 위해 SMS를 대기열에 넣는 중...",
		pl:   "Kolejkowanie SMS do wysłania na numer %v...",
		ptBR: "Enfileirando SMS para envio ao número %v...",
		ptPT: "Enfileiramento de SMS para envio para o número %v...",
		ru:   "SMS ставится в очередь на отправку на номер %v...",
		tr:   "%v numarasına gönderilmek üzere SMS sıraya alınıyor...",
		ukUA: "Ставимо SMS у чергу для надсилання на номер %v...",
		uz:   "%v raqamiga yuborish uchun SMS navbatga qo'yilmoqda...",
		zhCN: "正在将短信排队发送到号码 %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		arEG: "تم وضع الرسائل القصيرة في قائمة الانتظار لإرسالها إلى الرقم %v",
		de:   "SMS in Sendewarteschlange für %v",
		en:   "SMS is queued for sending to number %v",
		es:   "El SMS está en la cola para enviar al número %v",
		faIR: "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
		fr:   "SMS mis en file d'attente pour envoi au numéro %v",
		id:   "SMS diantrekan untuk dikirim ke nomor %v",
		it:   "SMS inserito in coda per l'invio al numero %v",
		jaJP: "SMSは番号 %v に送信するためにキューに入れられています",
		koKR: "SMS가 번호 %v로 보내기 위해 대기열에 있습니다",
		pl:   "SMS jest w kolejce do wysłania na numer %v",
		ptBR: "SMS está na fila para envio ao número %v",
		ptPT: "O SMS está na fila para envio para o número %v",
		ru:   "SMS поставлена в очередь на отправку на номер %v",
		tr:   "SMS, %v numarasına gönderilmek üzere sıraya alındı",
		ukUA: "SMS поставлено в чергу для надсилання на номер %v",
		uz:   "SMS %v raqamiga yuborish uchun navbatga qo'yildi",
		zhCN: "短信已排队发送到号码 %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		arEG: "توازن",
		de:   "Ausstehend",
		en:   "Balance",
		es:   "Balance",
		faIR: "تراز",
		fr:   "Solde",
		id:   "Saldo",
		it:   "Bilancio",
		jaJP: "残高",
		koKR: "잔액",
		pl:   "Saldo",
		ptBR: "Saldo",
		ptPT: "Equilíbrio",
		ru:   "Баланс",
		tr:   "Bakiye",
		ukUA: "Баланс",
		uz:   "Balans",
		zhCN: "余额",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		arEG: "عذراً، في الوقت الحالي تتوفر هذه القنوات فقط لإرسال الإيصال:",
		de:   "Entschuldigung, im Moment funktionieren nur ein paar Kanäle für den Versand von Quittungen:",
		en:   "Sorry, at the moment just these channels are available for sending a receipt:",
		es:   "Disculpa, de momento solo estos canales están disponibles para enviar el recibo:",
		faIR: "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
		fr:   "Désolé, pour le moment seuls ces canaux sont disponibles pour l'envoi d'un reçu :",
		id:   "Maaf, saat ini hanya saluran ini yang tersedia untuk mengirim tanda terima:",
		it:   "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
		jaJP: "申し訳ありませんが、現時点では領収書の送信に利用できるのはこれらのチャンネルのみです：",
		koKR: "죄송합니다. 현재 영수증 발송에 사용할 수 있는 채널은 다음과 같습니다:",
		pl:   "Przepraszamy, w tej chwili tylko te kanały są dostępne do wysyłania paragonu:",
		ptBR: "Desculpe, no momento apenas estes canais estão disponíveis para enviar um recibo:",
		ptPT: "Lamentamos, de momento apenas estes canais estão disponíveis para envio de recibo:",
		ru:   "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		tr:   "Üzgünüz, şu anda makbuz göndermek için yalnızca bu kanallar kullanılabilir:",
		ukUA: "Вибачте, на даний момент для надсилання квитанції доступні лише ці канали:",
		uz:   "Kechirasiz, hozirda kvitansiya yuborish uchun faqat ushbu kanallar mavjud:",
		zhCN: "抱歉，目前只有这些渠道可用于发送收据：",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		arEG: "📤 تم إرسال الإيصال عبر Telegram.",
		de:   "📤 Quittung per Telegram verschickt.",
		en:   "📤 Receipt sent through Telegram.",
		es:   "📤 El recibo ha sido enviado vía Telegram.",
		faIR: "📤 رسید از طریق تلگرام ارسال شد.",
		fr:   "📤 Reçu envoyé via Telegram.",
		id:   "📤 Tanda terima dikirim melalui Telegram.",
		it:   "📤 Notifica inviata tramite Telegram.",
		jaJP: "📤 Telegramを通じて領収書が送信されました。",
		koKR: "📤 텔레그램을 통해 영수증이 전송되었습니다.",
		pl:   "📤 Paragon wysłany przez Telegram.",
		ptBR: "📤 Recibo enviado pelo Telegram.",
		ptPT: "📤 Recibo enviado via Telegram.",
		ru:   "📤 Квитанция отправлена через телеграм.",
		tr:   "📤 Makbuz Telegram üzerinden gönderildi.",
		ukUA: "📤 Квитанцію надіслано через Telegram.",
		uz:   "📤 Kvitansiya Telegram orqali yuborildi.",
		zhCN: "📤 收据已通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		arEG: "لم يتم إرسال الإيصال عبر Telegram حيث قام %v بحذف الدردشة مع الروبوت.",
		de:   "Quittung konnte nicht per Telegram gesendet werden, da %v den Chat mit dem Bot gelöscht hat.",
		en:   "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		es:   "El recibo NO ha sido enviado vía Telegram porque %v ha eliminado el chat del bot.",
		faIR: "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
		fr:   "Reçu NON envoyé via Telegram car %v a supprimé le chat avec le bot.",
		id:   "Tanda terima TIDAK dikirim melalui Telegram karena %v telah menghapus obrolan dengan bot.",
		it:   "Notifica NON inviata tramite Telegram a %v perche' ha cancellato la chat con il bot",
		jaJP: "%v がボットとのチャットを削除したため、Telegramを通じて領収書が送信されませんでした。",
		koKR: "%v가 봇과의 채팅을 삭제했기 때문에 텔레그램을 통해 영수증이 전송되지 않았습니다.",
		pl:   "Paragon NIE został wysłany przez Telegram, ponieważ %v usunął czat z botem.",
		ptBR: "Recibo NÃO enviado pelo Telegram, pois %v excluiu o chat com o bot.",
		ptPT: "Recibo NÃO enviado pelo Telegram porque %v apagou o chat com o bot.",
		ru:   "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		tr:   "%v botla sohbeti sildiği için makbuz Telegram üzerinden GÖNDERİLMEDİ.",
		ukUA: "Квитанцію НЕ надіслано через Telegram, оскільки %v видалив чат з ботом.",
		uz:   "%v bot bilan suhbatni o'chirganligi sababli kvitansiya Telegram orqali YUBORILMADI.",
		zhCN: "由于 %v 已删除与机器人的聊天，收据未通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		arEG: "تم إرسال الإيصال عبر البريد الإلكتروني. (المعرف: %v)",
		de:   "Quittung wurde per Mail gesendet. (id: %v)",
		en:   "Receipt sent through email. (id: %v)",
		es:   "El recibo ha sido enviado vía e-mail. (id: %v)",
		faIR: "رسید از طریق ایمیل ارسال شد. (id: %v)",
		fr:   "Reçu envoyé par e-mail. (id: %v)",
		id:   "Tanda terima dikirim melalui email. (id: %v)",
		it:   "Notifica inviata tramite email (id: %v)",
		jaJP: "領収書がメールで送信されました。(id: %v)",
		koKR: "이메일을 통해 영수증이 전송되었습니다. (id: %v)",
		pl:   "Paragon wysłany przez e-mail. (id: %v)",
		ptBR: "Recibo enviado por e-mail. (id: %v)",
		ptPT: "Recibo enviado por e-mail. (id: %v)",
		ru:   "Квитанция отправлена через email. (id: %v)",
		tr:   "Makbuz e-posta yoluyla gönderildi. (id: %v)",
		ukUA: "Квитанцію надіслано електронною поштою. (id: %v)",
		uz:   "Kvitansiya elektron pochta orqali yuborildi. (id: %v)",
		zhCN: "收据已通过电子邮件发送。(id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		arEG: "تم إرسال الإيصال عبر الرسائل القصيرة",
		de:   "Quittung wurde per SMS gesendet.",
		en:   "Receipt sent through SMS",
		es:   "El recibo ha sido enviado vía SMS.",
		faIR: "رسید از طریق پیام کوتاه ارسال شد.",
		fr:   "Reçu envoyé par SMS",
		id:   "Tanda terima dikirim melalui SMS",
		it:   "Notifica inviata tramite SMS",
		jaJP: "領収書がSMSで送信されました",
		koKR: "SMS를 통해 영수증이 전송되었습니다",
		pl:   "Paragon wysłany przez SMS",
		ptBR: "Recibo enviado por SMS",
		ptPT: "Recibo enviado por SMS",
		ru:   "Квитанция отправлена через SMS.",
		tr:   "Makbuz SMS yoluyla gönderildi",
		ukUA: "Квитанцію надіслано через SMS",
		uz:   "Kvitansiya SMS orqali yuborildi",
		zhCN: "收据已通过短信发送",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		arEG: "انتقل إلى الوضع الخاص لرؤية تفاصيل الإيصال.",
		de:   "Schalte in den Privatmodus, um die Quittungsdetails zu sehen",
		en:   "Switch to private mode to see receipt details.",
		es:   "Pasar al modo privado para ver el recibo.",
		faIR: "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
		fr:   "Passez en mode privé pour voir les détails du reçu.",
		id:   "Beralih ke mode pribadi untuk melihat detail tanda terima.",
		it:   "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
		jaJP: "領収書の詳細を確認するには、プライベートモードに切り替えてください。",
		koKR: "영수증 세부 정보를 보려면 개인 모드로 전환하세요.",
		pl:   "Przełącz na tryb prywatny, aby zobaczyć szczegóły paragonu.",
		ptBR: "Mude para o modo privado para ver os detalhes do recibo.",
		ptPT: "Mude para o modo privado para ver os detalhes do recibo.",
		ru:   "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		tr:   "Makbuz ayrıntılarını görmek için özel moda geçin.",
		ukUA: "Перейдіть у приватний режим, щоб побачити деталі квитанції.",
		uz:   "Kvitansiya tafsilotlarini ko'rish uchun shaxsiy rejimga o'ting.",
		zhCN: "切换到私人模式以查看收据详情。",
	},
	MESSAGE_TEXT_RECEIPT_ATTEMPT_TO_VIEW_OWN: {
		arEG: "محاولة لعرض الإيصال الخاص.",
		de:   "Ein Versuch, den eigenen Kassenbon einzusehen.",
		en:   "An attempt to view own receipt.",
		es:   "Un intento de ver su propio recibo.",
		faIR: "تلاشی برای مشاهده رسید شخصی.",
		fr:   "Une tentative de visualisation de son propre reçu.",
		id:   "Upaya untuk melihat tanda terima sendiri.",
		it:   "Tentativo di visualizzare la propria ricevuta.",
		jaJP: "自分の領収書を表示しようとしました。",
		koKR: "자신의 영수증을 보려는 시도입니다.",
		pl:   "Próba obejrzenia własnego rachunku.",
		ptBR: "Uma tentativa de visualizar o próprio recibo.",
		ptPT: "Uma tentativa de visualizar o próprio recibo.",
		ru:   "Попытка посмотреть свою собственную квитанцию.",
		tr:   "Kendi fişini görüntüleme girişimi.",
		ukUA: "Спроба переглянути власний чек.",
		uz:   "O&#39;z chekini ko&#39;rishga urinish.",
		zhCN: "尝试查看自己的收据。"},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		arEG: "👓 تم عرض الإيصال",
		de:   "👓 Quittung wurde von Gegenpartei gesichtet",
		en:   "👓 Receipt viewed",
		es:   "👓 El recibo ha sido visto",
		faIR: "👓 رسید رویت شد.",
		fr:   "👓 Reçu consulté",
		id:   "👓 Tanda terima dilihat",
		it:   "👓 Debiti visti",
		jaJP: "👓 領収書を確認しました",
		koKR: "👓 영수증 확인됨",
		pl:   "👓 Paragon obejrzany",
		ptBR: "👓 Recibo visualizado",
		ptPT: "👓 Recibo visualizado",
		ru:   "👓 Квитанция просмотрена",
		tr:   "👓 Fiş görüntülendi",
		ukUA: "👓 Квитанцію переглянуто",
		uz:   "👓 Kvitansiya koʻrib chiqildi",
		zhCN: "👓 已查看收据"},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		arEG: "بإمكانك عرض رقم هاتفك الخاص بالتنسيق الذي نتوقعه.",
		de:   "Du kannst deine eigene Nummer in dem Format anzeigen, welches wir erwarten.",
		en:   "You can view your own phone number in the format we are expecting.",
		es:   "Puedes ver tu número de teléfono en el formato previsto.",
		faIR: "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
		fr:   "Vous pouvez consulter votre propre numéro de téléphone dans le format que nous attendons.",
		id:   "Anda dapat melihat nomor telepon Anda sendiri dalam format yang kami harapkan.",
		it:   "Puoi visionare il tuo numero di telefono nel formato previsto.",
		jaJP: "ご自身の電話番号を、期待される形式で表示できます。",
		koKR: "당신은 우리가 기대하는 형식으로 자신의 전화번호를 볼 수 있습니다.",
		pl:   "Możesz wyświetlić swój numer telefonu w oczekiwanym przez nas formacie.",
		ptBR: "Você pode visualizar seu próprio número de telefone no formato que estamos esperando.",
		ptPT: "Pode visualizar o seu próprio número de telefone no formato que estamos à espera.",
		ru:   "Вы можете посмотреть свой номер телефона в ожидаемом нами формате.",
		tr:   "Beklediğimiz formatta kendi telefon numaranızı görüntüleyebilirsiniz.",
		ukUA: "Ви можете переглянути свій власний номер телефону у форматі, який ми очікуємо.",
		uz:   "Siz o&#39;z telefon raqamingizni biz kutgan formatda ko&#39;rishingiz mumkin.",
		zhCN: "您可以按照我们期望的格式查看您自己的电话号码。"},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		arEG: "عرض رقمي بالتنسيق المتوقع",
		de:   "Zeige meine Nummer im erwarteten Format.",
		en:   "View my number in the expectd format",
		es:   "Mostrar mi número de teléfono en el formato previsto",
		faIR: "رویت شماره خودم با فرمت مورد انتظار",
		fr:   "Afficher mon numéro au format attendu",
		id:   "Lihat nomor saya dalam format yang diharapkan",
		it:   "Mostra il mio numero nel formato previsto",
		jaJP: "期待される形式で自分の番号を表示する",
		koKR: "예상 형식으로 내 번호를 확인하세요",
		pl:   "Wyświetl mój numer w oczekiwanym formacie",
		ptBR: "Veja meu número no formato esperado",
		ptPT: "Veja o meu número no formato esperado",
		ru:   "Посмотреть мой номер в ожидаемоем формате",
		tr:   "Numaramı beklenen formatta görüntüle",
		ukUA: "Переглянути мій номер у належному форматі",
		uz:   "Mening raqamimni kutilgan formatda ko&#39;ring",
		zhCN: "以预期格式查看我的号码"},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		arEG: "عرض التاريخ الكامل",
		de:   "Lade den ganzen Verlauf",
		en:   "Show full history",
		es:   "Mostrar cronología completa",
		faIR: "نمایش کامل سوابق",
		fr:   "Afficher l&#39;historique complet",
		id:   "Tampilkan riwayat lengkap",
		it:   "Mostra cronologia completa",
		jaJP: "全履歴を表示",
		koKR: "전체 기록 보기",
		pl:   "Pokaż pełną historię",
		ptBR: "Mostrar histórico completo",
		ptPT: "Mostrar histórico completo",
		ru:   "Показать всю историю",
		tr:   "Tam geçmişi göster",
		ukUA: "Показати повну історію",
		uz:   "Toʻliq tarixni koʻrsatish",
		zhCN: "显示完整历史记录"},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		arEG: "إنه ليس رقمًا",
		de:   "Das ist keine Nummer",
		en:   "it is not a number",
		es:   "No es un número",
		faIR: "این یک شماره نیست",
		fr:   "ce n&#39;est pas un numéro",
		id:   "itu bukan angka",
		it:   "Non e' un numero",
		jaJP: "それは数字ではない",
		koKR: "그것은 숫자가 아니다",
		pl:   "to nie jest liczba",
		ptBR: "não é um número",
		ptPT: "não é um número",
		ru:   "Это не цифра",
		tr:   "bu bir sayı değil",
		ukUA: "це не число",
		uz:   "bu raqam emas",
		zhCN: "它不是一个数字"},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		arEG: "يجب أن يكون الرقم موجبًا ( <i>أكبر من 0</i> )",
		de:   "Die Nummer sollte positiv sein (<i>größer als 0</i>)",
		en:   "The number should be positive (<i>greater than 0</i>)",
		es:   "El número tiene que ser positivo (<i>más de  0</i>)",
		faIR: "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
		fr:   "Le nombre doit être positif ( <i>supérieur à 0</i> )",
		id:   "Angkanya harus positif ( <i>lebih besar dari 0</i> )",
		it:   "Il numero dovrebbe essere positivo (<i>maggiore di 0</i>)",
		jaJP: "数値は正の数（ <i>0より大きい</i>）である必要があります",
		koKR: "숫자는 양수( <i>0보다 큼</i> )여야 합니다.",
		pl:   "Liczba powinna być dodatnia ( <i>większa niż 0</i> )",
		ptBR: "O número deve ser positivo ( <i>maior que 0</i> )",
		ptPT: "O número deve ser positivo ( <i>maior que 0</i> )",
		ru:   "Цифра должна быть положительной (<i>больше нуля</i>)",
		tr:   "Sayı pozitif ( <i>0&#39;dan büyük</i> ) olmalıdır",
		ukUA: "Число має бути додатним ( <i>більшим за 0</i> )",
		uz:   "Raqam ijobiy bo&#39;lishi kerak ( <i>0 dan katta</i> )",
		zhCN: "该数字应为正数（<i>大于 0</i> ）"},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		arEG: "كم تم إرجاعه؟",
		de:   "Wie viel wurde beglichen?",
		en:   "How much have been returned?",
		es:   "¿Cuánto/s te han devuelto?",
		faIR: "چه مقدار بازپرداخت شده است؟",
		fr:   "Combien ont été retournés ?",
		id:   "Berapa banyak yang sudah dikembalikan?",
		it:   "Quanto ti e' stato restituito?",
		jaJP: "いくら返金されましたか？",
		koKR: "얼마나 반환되었나요?",
		pl:   "Ile zwrócono?",
		ptBR: "Quanto foi devolvido?",
		ptPT: "Quanto foi devolvido?",
		ru:   "Сколько было возвращено?",
		tr:   "Ne kadarı iade edildi?",
		ukUA: "Скільки вже повернули?",
		uz:   "Qancha pul qaytarildi?",
		zhCN: "退还了多少？"},
	MESSAGE_TEXT_RETURN_IS_TOO_BIG: {
		arEG: "لقد قررت إرجاع %v ولكن المبلغ المستحق هو %v فقط.\n\nالرجاء إدخال قيم تساوي %v أو أقل.",
		de:   "Sie haben entschieden, %v zurückzugeben, aber der ausstehende Betrag ist nur %v.\n\nBitte geben Sie Werte gleich %v oder weniger ein.",
		en:   "You decided to return %v but outstanding amount is just %v.\n\nPlease enter values equal to %v or less.",
		es:   "Decidiste devolver %v pero la cantidad pendiente es solo %v.\n\nPor favor ingrese valores iguales a %v o menos.",
		faIR: "شما تصمیم گرفتید %v را بازگردانید اما مقدار قابل توجهی فقط %v است.\n\nلطفا مقادیر برابر %v یا کمتر را وارد کنید",
		fr:   "Vous avez décidé de retourner %v mais le montant restant est de seulement %v.\n\nVeuillez saisir des valeurs égales à %v ou moins.",
		id:   "Anda memutuskan untuk mengembalikan %v tetapi jumlah yang terutang hanya %v.\n\nSilakan masukkan nilai yang sama dengan %v atau kurang.",
		it:   "Hai deciso di restituire %v ma la quantità in sospeso è solo %v.\n\nInserisci valori pari o uguali a %v o meno.",
		jaJP: "%v を返却することにしましたが、未払い金額は %v だけです。\n\n%v 以下の値を入力してください。",
		koKR: "%v를 반환하기로 결정했지만 미지불 금액이 %v에 불과합니다.\n\n%v 이하의 값을 입력하세요.",
		pl:   "Zdecydowałeś się zwrócić %v, ale kwota należności wynosi zaledwie %v.\n\nProszę wprowadzić wartości równe %v lub mniejsze.",
		ptBR: "Você decidiu devolver %v, mas o valor pendente é de apenas %v.\n\nInsira valores iguais ou menores que %v.",
		ptPT: "Decidiu devolver %v, mas o valor em dívida é apenas de %v. \n\nIntroduza valores iguais ou inferiores a %v.",
		ru:   "Вы решили вернуть %v, но непогашенная сумма равна %v. \n\n Пожалуйста, введите значение равное %v или меньше.",
		tr:   "%v iade etmeye karar verdiniz ancak bakiye sadece %v.\n\nLütfen %v veya daha düşük değerler girin.",
		ukUA: "Ви вирішили повернути %v, але сума непогашеної оплати становить лише %v.\n\nБудь ласка, введіть значення, що дорівнюють %v або менше.",
		uz:   "Siz %v ni qaytarishga qaror qildingiz, ammo to‘lanmagan summa faqat %v.\n\nIltimos, %v yoki undan kichikroq qiymatlarni kiriting.",
		zhCN: "您决定退回 %v，但未结金额仅为 %v。\n\n请输入等于或小于 %v 的值。"},
	MESSAGE_TEXT_HELP_ROOT: {
		arEG: "ما هو سؤالك؟ إذا فاتك أي شيء هنا، فلا تتردد في طرحه في @%v",
		de:   "Was hast du für eine Frage? Wenn irgendwas unklar ist, frag ruhig hier @%v",
		en:   "What is your question? If anything is missed here, feel free to ask in our @%v",
		es:   "¿Cuál es tu pregunta? Si algo se pierde aquí, siéntase libre de preguntar en nuestro @%v",
		faIR: "سوالت چیست؟ اگر چیزی در اینجا از دست رفته است، لطفا در @%v ما بپرسید",
		fr:   "Quelle est votre question ? Si vous avez oublié quelque chose, n&#39;hésitez pas à nous contacter via notre @%v",
		id:   "Apa pertanyaan Anda? Jika ada yang terlewat di sini, silakan bertanya di @%v kami",
		it:   "Qual è la tua domanda? Se qualche cosa è mancato qui, non esitate a chiedere al nostro @%v",
		jaJP: "ご質問は何ですか？何かご不明な点がございましたら、お気軽に@%vまでお問い合わせください。",
		koKR: "궁금한 점이 있으신가요? 궁금한 점이 있으시면 @%v로 편하게 질문해 주세요.",
		pl:   "Jakie jest Twoje pytanie? Jeśli coś jest tu pominięte, możesz zapytać na naszym @%v",
		ptBR: "Qual é a sua pergunta? Se algo estiver faltando, sinta-se à vontade para perguntar em @%v",
		ptPT: "Qual é a sua pergunta? Se faltar algo, sinta-se à vontade para perguntar em @%v",
		ru:   "Какой у вас вопрос? Если здесь нет ответа пожалуйста спросите в нашей группе @%v",
		tr:   "Sorunuz nedir? Burada bir şey atlanırsa, @%v adresinden sormaktan çekinmeyin.",
		ukUA: "Яке ваше запитання? Якщо тут щось пропущено, не соромтеся запитати в нашому @%v",
		uz:   "Savolingiz nima? Agar biror narsa o&#39;tkazib yuborilgan bo&#39;lsa, @%v orqali so&#39;rang",
		zhCN: "你有什么问题？如果这里还有什么遗漏，欢迎在我们的@%v 中提问。"},
	MESSAGE_TEXT_HELP_BACK_TO_ROOT: {
		arEG: "العودة إلى قائمة الأسئلة الشائعة",
		de:   "Zurück zur FAQ Liste",
		en:   "Back to FAQ list",
		es:   "Volver a la lista de preguntas frecuentes",
		faIR: "بازگشت به لیست سوالات متداول",
		fr:   "Retour à la liste des FAQ",
		id:   "Kembali ke daftar FAQ",
		it:   "Torna all'elenco delle FAQ",
		jaJP: "FAQリストに戻る",
		koKR: "FAQ 목록으로 돌아가기",
		pl:   "Powrót do listy FAQ",
		ptBR: "Voltar para a lista de perguntas frequentes",
		ptPT: "Voltar à lista de perguntas frequentes",
		ru:   "Назад к списку вопросов",
		tr:   "SSS listesine geri dön",
		ukUA: "Назад до списку поширених запитань",
		uz:   "FAQ roʻyxatiga qaytish",
		zhCN: "返回常见问题列表"},
	HELP_HOW_TO_CREATE_BILL_Q: {
		arEG: "كيفية إنشاء فاتورة جديدة؟",
		de:   "Wie erstellt man Rechnungen?",
		en:   "How to create new bill?",
		es:   "¿Cómo crear una nueva factura?",
		faIR: "چگونه برای ایجاد لایحه جدید؟",
		fr:   "Comment créer une nouvelle facture ?",
		id:   "Bagaimana cara membuat tagihan baru?",
		it:   "Come creare un nuovo conto?",
		jaJP: "新しい請求書を作成するには?",
		koKR: "새로운 청구서를 어떻게 생성하나요?",
		pl:   "Jak utworzyć nowy rachunek?",
		ptBR: "Como criar uma nova fatura?",
		ptPT: "Como criar uma nova fatura?",
		ru:   "Как создать новый счёт?",
		tr:   "Yeni fatura nasıl oluşturulur?",
		ukUA: "Як створити новий рахунок?",
		uz:   "Yangi qonun loyihasini qanday yaratish kerak?",
		zhCN: "如何创建新账单？"},
	HELP_HOW_TO_CREATE_BILL_A: {
		arEG: `<b>كيفية إنشاء فاتورة جديدة</b> {ن}<pre style=";text-align:right;direction:rtl"> الفاتورة - نفقات مشتركة بين شخصين أو أكثر.</pre> 

لذلك، من الأفضل <b>إنشاء فاتورة في دردشة تيليجرام بخطوتين فقط</b> : 
 <i>استخدم &quot;تقسيم الفاتورة مع مستخدمي تيليجرام&quot; للقيام بذلك بسرعة</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">أضفني إلى مجموعة تيليجرام</a> أو افتح دردشة مع صديق. 

 2. اكتب <code>@splitusbot {amount} {bill_name}</code> واختر النتيجة من القائمة. على سبيل المثال: 
<pre style=";text-align:right;direction:rtl"> @splitusbot 45.5 بيتزا</pre> 
 يمكن لأي عضو في المجموعة مشاركة الفاتورة بالضغط على زر <code>Join</code> .

 <b>أو</b> يمكنك إنشاء فاتورة مباشرةً في @{{.BotCode}}. ولكن بعد ذلك، ستحتاج إلى إضافة المشاركين يدويًا.`,
		de: `<b>Wie man eine Rechnung erstellt</b>
<pre>Rechnung — Kosten, die unter zwei oder mehr Personen geteilt werden.</pre>

Deswegen kannst du am besten direkt <b>im Telegram Chat eine Rechnung erstellen, in zwei Schritten</b>:
<i>Benutze die Funktion "Teile Rechnung mit Telegram Benutzer", um es schnell zu machen:</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Füg mich einer Gruppe hinzu</a> oder öffne einen Chat mit deinem Freund.

	2. Schreibe <code>@splitusbot {amount} {bill_name}</code> und wähle ein passendes Ergebnis. Zum Beispiel:
<pre>		@splitusbot 45.5 pizza</pre>
	   Und dann kann jeder, der die Rechnung teilen will mit <code>Join</code> einsteigen.

<b>Alternativ</b> kannst du die Rechnung auch direkt in @{{.BotCode}} erstellen. Aber dann musst du die Personen, mit denen du die Rechnung teilst, einzeln hinzufügen.`,
		en: `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
		es: `<b>Cómo crear una nueva factura</b> 
<pre> Factura: gasto compartido entre dos o más personas.</pre> Por eso, lo mejor es <b>crear una factura en el chat de Telegram en solo 2 pasos</b> :
 <i>Usa &quot;Dividir factura con usuarios de Telegram&quot; para hacerlo rápidamente.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Agrégame a un grupo de Telegram</a> o abre un chat con un amigo.

 2. Escribe <code>@splitusbot {amount} {bill_name}</code> y selecciona el resultado en el menú. Por ejemplo:
<pre> @splitusbot 45.5 pizza</pre> 
 Cualquier miembro del grupo puede compartir la cuenta pulsando el botón <code>Join</code> .

 <b>También</b> puedes crear una cuenta directamente en @{{.BotCode}}. En ese caso, tendrías que añadir participantes manualmente.`,
		faIR: //nolint:staticcheck // disable ST1018 for this line
		`<b>نحوه ایجاد یک صورتحساب جدید</b> 
<pre style=";text-align:right;direction:rtl"> صورتحساب - هزینه مشترک بین دو یا چند نفر.</pre> 

به همین دلیل بهترین کار این است که <b>در چت تلگرام فقط در ۲ مرحله صورتحساب ایجاد کنید</b> :
 <i>برای انجام سریع این کار از «تقسیم صورتحساب با کاربر(های) تلگرام» استفاده کنید</i> 

 ۱. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">من را به گروه تلگرام اضافه کنید</a> یا چت با یک دوست را باز کنید.

 ۲. عبارت <code>@splitusbot {amount} {bill_name}</code> را تایپ کنید و نتیجه را از منو انتخاب کنید. به عنوان مثال:
<pre style=";text-align:right;direction:rtl"> @splitusbot پیتزای ۴۵.۵</pre> 
 سپس هر یک از اعضای گروه می‌توانند با فشردن دکمه <code>Join</code> ، صورتحساب را به اشتراک بگذارند.

 <b>همچنین</b> می‌توانید مستقیماً در @{{.BotCode}} صورتحساب ایجاد کنید. اما در این صورت باید شرکت‌کنندگان را به صورت دستی اضافه کنید.`,
		fr: `<b>Comment créer une nouvelle facture</b> 
<pre> Facture — dépense partagée entre deux ou plusieurs personnes.</pre> 

C&#39;est pourquoi il est préférable de <b>créer une facture dans le chat Telegram en deux étapes</b> :
 <i>utilisez « Partager la facture avec un ou plusieurs utilisateurs Telegram » pour le faire rapidement.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Ajoutez-moi à un groupe Telegram</a> ou ouvrez un chat avec un ami.

 2. Tapez <code>@splitusbot {amount} {bill_name}</code> et sélectionnez le résultat dans le menu. Par exemple :
<pre> @splitusbot 45,5 pizzas</pre> 
 N&#39;importe quel membre du groupe peut partager la facture en appuyant sur le bouton <code>Join</code> .

 Vous pouvez <b>également</b> créer une facture directement dans le @{{.BotCode}}. Vous devrez alors ajouter manuellement les participants.`,
		id: `<b>Cara membuat tagihan baru</b> 
<pre> Tagihan — biaya yang dibagi antara dua orang atau lebih.</pre> 

Itulah sebabnya cara terbaik adalah <b>membuat tagihan di obrolan Telegram hanya dalam 2 langkah</b> :
 <i>gunakan &quot;Pisahkan tagihan dengan pengguna Telegram&quot; untuk melakukannya dengan cepat.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Tambahkan saya ke grup Telegram</a> atau buka obrolan dengan teman.

 2. Ketik <code>@splitusbot {amount} {bill_name}</code> dan pilih hasil dari menu. Misalnya:
<pre> @splitusbot 45,5 pizzanya</pre> 
 Kemudian setiap anggota grup dapat membagikan tagihan dengan menekan tombol <code>Join</code> .

 <b>Atau</b> Anda dapat membuat tagihan langsung di @{{.BotCode}}. Namun, Anda perlu menambahkan peserta secara manual.`,
		it: `<b>Come creare una nuova fattura</b> 
<pre> Bolletta: spesa condivisa tra due o più persone.</pre> 

Ecco perché la soluzione migliore è <b>creare una fattura nella chat di Telegram in soli 2 passaggi</b> :
 <i>usa &quot;Dividi la fattura con gli utenti Telegram&quot; per farlo rapidamente.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Aggiungimi al gruppo Telegram</a> o apri una chat con un amico.

 2. Digita <code>@splitusbot {amount} {bill_name}</code> e seleziona il risultato dal menu. Ad esempio:
<pre> @splitusbot 45.5 pizza</pre> 
 Quindi qualsiasi membro del gruppo può condividere la fattura premendo il pulsante <code>Join</code> .

 <b>In alternativa,</b> puoi creare una fattura direttamente in @{{.BotCode}}. In tal caso, però, dovrai aggiungere manualmente i partecipanti.`,
		jaJP: `<b>新しい請求書を作成する方法</b>
<pre>請求書 - 2 人以上で共同で支払う費用。</pre> 

そのため<b>、Telegramチャットで請求書を2ステップで作成する</b>のが最善です。
 <i>「Telegramユーザーと請求書を分割」を使用すると、簡単に作成できます</i>。

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Telegramグループに追加する</a>か、友達とチャットを開きます。

 2. <code>@splitusbot {amount} {bill_name}</code>と入力し、メニューから結果を選択します。例:
<pre> @splitusbot 45.5 ピザ</pre>
 グループのメンバーは誰でも<code>Join</code>ボタンを押すことで請求書を共有できます。

<b>または、</b> @{{.BotCode}} 内で請求書を作成することもできます。ただし、その場合は参加者を手動で追加する必要があります。`,
		koKR: `<b>새로운 청구서를 만드는 방법</b> 
<pre> 청구서 - 두 사람 이상이 공유하는 비용.</pre> 

그래서 가장 좋은 방법은 <b>텔레그램 채팅에서 두 단계만으로 청구서를 생성하는</b> 것입니다.
 <i>&quot;텔레그램 사용자와 청구서 분할&quot;을 사용하면 빠르게 청구서를 생성할 수 있습니다.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">텔레그램 그룹에 저를 추가하거나</a> 친구와 채팅을 시작합니다.

 2. <code>@splitusbot {amount} {bill_name}</code> 입력하고 메뉴에서 결과를 선택합니다. 예:
<pre> @splitusbot 45.5 피자</pre> 
 그룹의 모든 구성원이 <code>Join</code> 버튼을 눌러 청구서를 공유할 수 있습니다.

 <b>또는</b> @{{.BotCode}}에서 바로 청구서를 생성할 수도 있습니다. 하지만 이 경우 참여자를 직접 추가해야 합니다.`,
		pl: `<b>Jak utworzyć nowy rachunek</b> 
<pre> Rachunek — wydatek dzielony między dwie lub więcej osób.</pre> 

Dlatego najlepiej jest <b>utworzyć rachunek w czacie Telegram w zaledwie 2 krokach</b> :
 <i>użyj „Podziel rachunek z użytkownikiem(ami) Telegrama”, aby zrobić to szybko</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Dodaj mnie do grupy Telegrama</a> lub otwórz czat ze znajomym.

 2. Wpisz <code>@splitusbot {amount} {bill_name}</code> i wybierz wynik z menu. Na przykład:
<pre> @splitusbot 45,5 pizzy</pre> 
 Następnie każdy członek grupy może udostępnić rachunek, naciskając przycisk <code>Join</code> .

 <b>Alternatywnie</b> możesz utworzyć rachunek bezpośrednio w @{{.BotCode}}. Ale wtedy będziesz musiał ręcznie dodać uczestników.`,
		ptBR: `<b>Como criar uma nova fatura</b> 
<pre> Conta — despesa compartilhada entre duas ou mais pessoas.</pre> 

Por isso, a melhor maneira de <b>criar uma conta no chat do Telegram é em apenas 2 passos</b> :
 <i>use &quot;Dividir conta com usuário(s) do Telegram&quot; para fazer isso rapidamente.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Adicione-me ao grupo do Telegram</a> ou abra um chat com um amigo.

 2. Digite <code>@splitusbot {amount} {bill_name}</code> e selecione o resultado no menu. Por exemplo:
<pre> @splitusbot 45,5 pizza</pre> 
 Qualquer membro do grupo pode compartilhar a conta clicando no botão <code>Join</code> .

 <b>Como alternativa,</b> você pode criar uma conta diretamente no @{{.BotCode}}. Nesse caso, você precisaria adicionar os participantes manualmente.`,
		ptPT: `<b>Como criar uma nova fatura</b> 
<pre> Conta — despesa repartida entre duas ou mais pessoas.</pre> 

Por isso, a melhor forma de <b>criar uma conta no chat do Telegram é em apenas 2 passos</b> :
 <i>utilize &quot;Dividir conta com utilizador(es) do Telegram&quot; para o fazer rapidamente.</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Adicione-me ao grupo do Telegram</a> ou abra um chat com um amigo.

 2. Digite <code>@splitusbot {amount} {bill_name}</code> e seleccione o resultado no menu. Por exemplo:
<pre> @splitusbot 45,5 pizza</pre> 
 Qualquer membro do grupo pode partilhar a conta clicando no botão <code>Join</code> .

 <b>Em alternativa,</b> pode criar uma conta diretamente no @{{.BotCode}}. Nesse caso, precisaria de adicionar os participantes manualmente.`,
		ru: `<b>Как создать счёт*</b>
<pre>Счёт — совместная трата между несколькими людьми.</pre>

Вот почему это хорошая идея <b>создать счёт в Telegram чате всего за 2 шага</b>:
<i>используйте "Разделить счёт с Telegram пользователем" чтобы сделать это быстро</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Добавить бот в Telegram группу</a> или откройте чат с другом.

	2. Напишите <code>@splitusbot {amount} {bill_name}</code> и  выберите результат из меню. Например:
<pre>		@splitusbot 45.5 пицца</pre>
	   После этого любой член группы может разделить счёт нажав кнопку <code>Присоедениться</code>.

<b>Или</b> вы можете создать счёт прямо в @{{.BotCode}}. Но тогда надо будет вручную добавить участников.`,
		tr: `<b>Yeni bir fatura nasıl oluşturulur</b> 
<pre> Fatura — iki veya daha fazla kişi arasında paylaşılan masraf.</pre> 

Bu yüzden en iyisi <b>Telegram sohbetinde fatura oluşturmak sadece 2 adımdadır</b> :
 <i>Bunu hızlı bir şekilde yapmak için &quot;Faturayı Telegram kullanıcı(lar)ıyla böl&quot; seçeneğini kullanın</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Beni Telegram grubuna ekleyin</a> veya bir arkadaşınızla sohbeti açın.

 2. <code>@splitusbot {amount} {bill_name}</code> yazın ve menüden sonucu seçin. Örneğin:
<pre> @splitusbot 45.5 pizza</pre> 
 Grubun herhangi bir üyesi <code>Join</code> düğmesine basarak faturayı paylaşabilir.

 <b>Alternatif olarak,</b> doğrudan @{{.BotCode}}&#39;te bir fatura oluşturabilirsiniz. Ancak o zaman katılımcıları manuel olarak eklemeniz gerekir.`,
		ukUA: `<b>Як створити новий рахунок</b> 
<pre> Рахунок — розподіл витрат між двома або більше особами.</pre> 

Ось чому найкраще <b>створити рахунок у чаті Telegram лише за 2 кроки</b> :
 <i>скористайтеся опцією &quot;Розділити рахунок з користувачем(ами) Telegram&quot;, щоб зробити це швидко</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Додайте мене до групи Telegram</a> або відкрийте чат з другом.

 2. Введіть <code>@splitusbot {amount} {bill_name}</code> та виберіть результат з меню. Наприклад:
<pre> @splitusbot 45.5 піца</pre> 
 Тоді будь-який учасник групи може поділитися рахунком, натиснувши кнопку <code>Join</code> .

 <b>Або ж</b> ви можете створити рахунок безпосередньо в @{{.BotCode}}. Але тоді вам потрібно буде вручну додавати учасників.`,
		uz: `<b>Yangi hisobni qanday yaratish kerak</b> 
<pre> Bill - ikki yoki undan ortiq kishi o&#39;rtasidagi umumiy xarajat.</pre> 

Shuning uchun <b>Telegram chatida atigi 2 bosqichda hisob yaratish</b> eng yaxshisidir:
 <i>buni tezda amalga oshirish uchun “Telegram foydalanuvchi(lar)i bilan hisob-kitobni taqsimlash”dan foydalaning</i> 

 1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Meni Telegram guruhiga qoʻshing</a> yoki doʻstingiz bilan suhbatni oching.

 2. <code>@splitusbot {amount} {bill_name}</code> kiriting va menyudan natijani tanlang. Masalan:
<pre> @splitusbot 45,5 pizza</pre> 
 Guruhning istalgan aʼzosi <code>Join</code> tugmasini bosib hisobni baham koʻrishi mumkin.

 <b>Shu bilan bir qatorda,</b> toʻgʻridan-toʻgʻri @{{.BotCode}} orqali hisob yaratishingiz mumkin. Ammo keyin ishtirokchilarni qo&#39;lda kiritishingiz kerak bo&#39;ladi.`,
		zhCN: `<b>如何创建新账单</b>
<pre>账单 — 两个或两个以上的人共同承担的费用。</pre> 

因此，<b>在 Telegram 聊天中创建账单的最佳方式是只需两步</b>：
<i>使用“与 Telegram 用户分摊账单”快速完成</i>

 1.<a href="https://t.me/{{.BotCode}}?startgroup=from-help">将我添加到 Telegram 群组</a>或与好友开启聊天。

 2. 输入<code>@splitusbot {amount} {bill_name}</code>并从菜单中选择结果。例如：
<pre> @splitusbot 45.5 披萨</pre>
 之后，小组成员可以通过点击<code>Join</code>按钮来共享账单。

<b>或者，</b>您也可以直接在 @{{.BotCode}} 中创建账单。但您需要手动添加参与者。`},
	MESSAGE_TEXT_HELP: {
		arEG: "يرجى الإبلاغ عن أي مشكلة أو تقديم طلب ميزة على موقعنا.",
		de:   "Bitte melde jedes Problem und jeden Wunsch auf unserer Webseite.",
		en:   "Please report any issue or submit a feature request at our website.",
		es:   "Puedes informarnos de algún problema o proponernos cualquier mejora en nuestra web.",
		faIR: "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
		fr:   "Veuillez signaler tout problème ou soumettre une demande de fonctionnalité sur notre site Web.",
		id:   "Silakan laporkan masalah atau kirimkan permintaan fitur di situs web kami.",
		it:   "Segnala un problema o proponi un miglioramento sul nostro sito web.",
		jaJP: "問題がございましたら、弊社の Web サイトで報告するか、機能のリクエストを送信してください。",
		koKR: "문제가 있으면 당사 웹사이트에 보고해 주시고, 기능 요청도 제출해 주시기 바랍니다.",
		pl:   "Prosimy o zgłaszanie wszelkich problemów i przesyłanie próśb o dodanie funkcji na naszej stronie internetowej.",
		ptBR: "Por favor, reporte qualquer problema ou envie uma solicitação de recurso em nosso site.",
		ptPT: "Por favor, reporte qualquer problema ou envie um pedido de recurso no nosso site.",
		ru:   "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
		tr:   "Lütfen herhangi bir sorunu bildirin veya web sitemiz üzerinden bir özellik talebi gönderin.",
		ukUA: "Будь ласка, повідомте про будь-яку проблему або надішліть запит на функцію на нашому вебсайті.",
		uz:   "Iltimos, har qanday muammo haqida xabar bering yoki veb-saytimizga xususiyat so&#39;rovini yuboring.",
		zhCN: "请在我们的网站上报告任何问题或提交功能请求。"},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		arEG: "صفحة الدعم",
		de:   "Support Seite",
		en:   "Support page",
		es:   "La página de ayuda",
		faIR: "صفحه پشتیبانی",
		fr:   "Page d&#39;assistance",
		id:   "Halaman dukungan",
		it:   "Pagina d'aiuto",
		jaJP: "サポートページ",
		koKR: "지원 페이지",
		pl:   "Strona wsparcia",
		ptBR: "Página de suporte",
		ptPT: "Página de suporte",
		ru:   "Cтраница поддержки ",
		tr:   "Destek sayfası",
		ukUA: "Сторінка підтримки",
		uz:   "Qo&#39;llab-quvvatlash sahifasi",
		zhCN: "支持页面"},
	COMMAND_TEXT_REPORT_A_BUG: {
		arEG: "الإبلاغ عن خطأ",
		de:   "einen Fehler melden",
		en:   "Report a bug",
		es:   "Informar de un error",
		faIR: "گزارش یک باگ",
		fr:   "Signaler un bug",
		id:   "Laporkan bug",
		it:   "Segnala un bug",
		jaJP: "バグを報告する",
		koKR: "버그 신고",
		pl:   "Zgłoś błąd",
		ptBR: "Reportar um bug",
		ptPT: "Reportar um bug",
		ru:   "Сообщить об ошибке",
		tr:   "Bir hatayı bildirin",
		ukUA: "Повідомити про помилку",
		uz:   "Xato haqida xabar bering",
		zhCN: "报告错误"},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		arEG: "أضف فكرة",
		de:   "eine Idee äußern",
		en:   "Add an idea",
		es:   "Proponer una idea",
		faIR: "یک ایده اضافه کنید.",
		fr:   "Ajouter une idée",
		id:   "Tambahkan ide",
		it:   "Proponi un idea",
		jaJP: "アイデアを追加する",
		koKR: "아이디어를 추가하세요",
		pl:   "Dodaj pomysł",
		ptBR: "Adicione uma ideia",
		ptPT: "Adicione uma ideia",
		ru:   "Предложить идею",
		tr:   "Bir fikir ekle",
		ukUA: "Додати ідею",
		uz:   "Fikr qo&#39;shing",
		zhCN: "添加想法"},
	MESSAGE_TEXT_WELCOME: {
		arEG: `مرحبًا، أنا كوليكتيوس - محاسبك الشخصي وجامعك.

يمكنني تسجيل من هو المستحق لمن وتذكيرك بموعد استحقاق الإقرار.

ماذا تريد أن تفعل؟`,
		de: `Hallo, ich bin Collectius - dein persönlicher Buchhalter.

Ich kann dir sagen, wer wem was schuldet und wann die Schuld jeweils fällig ist.

Was würdest du gerne machen?`,
		en: `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,
		es: `Hola, me llamo Collectius, soy tu contable y asesor personal.

Puedo apuntar quién debe a quién y recordarte la fecha de devolución.

¿Qué te apetecería hacer?`,
		faIR: `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,
		fr: `Bonjour, je suis Collectius - votre comptable et collecteur personnel.

Je peux enregistrer qui doit de l&#39;argent à qui et vous rappeler quand la déclaration est due.

Que souhaitez-vous faire ?`,
		id: `Hai, saya Collectius - akuntan &amp; kolektor pribadi Anda.

Saya dapat mencatat siapa yang berutang kepada siapa dan mengingatkan kapan pengembalian jatuh tempo.

Apa yang ingin Anda lakukan?`,
		it: `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,
		jaJP: `こんにちは。Collectius です。あなたの専属会計士兼徴収人です。

誰が誰に借金をしているのかを記録し、返済期限が来たらリマインダーを送信できます。

何をご希望ですか?`,
		koKR: `안녕하세요, 저는 귀하의 개인 회계사 겸 수금 담당자인 콜렉티우스입니다.

누가 누구에게 빚이 있는지 기록하고, 신고 기한이 되면 알려드릴 수 있습니다.

무엇을 하시겠습니까?`,
		pl: `Cześć, nazywam się Collectius i jestem Twoim osobistym księgowym i windykatorem.

Mogę zarejestrować, komu i kiedy należy złożyć zeznanie, oraz przypomnieć o terminie jego złożenia.

Co chcesz zrobić?`,
		ptBR: `Olá, sou a Collectius - seu contador e cobrador pessoal.

Posso registrar quem deve a quem e lembrá-lo quando a declaração vence.

O que você gostaria de fazer?`,
		ptPT: `Olá, sou a Collectius - o seu contabilista e cobrador pessoal.

Posso registar quem deve a quem e lembrá-lo quando a declaração expira.

O que gostaria de fazer?`,
		ru: `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
		tr: `Merhaba, ben Collectius - kişisel muhasebeciniz ve tahsildarınız.

Kimin kime borcu olduğunu kaydedebilir ve beyannamenin ne zaman ödenmesi gerektiğini hatırlatabilirim.

Ne yapmak istersiniz?`,
		ukUA: `Привіт, я Collectius — ваш особистий бухгалтер та колектор.

Я можу фіксувати, хто кому винен, та нагадувати, коли потрібно повернути декларацію.

Що б ви хотіли зробити?`,
		uz: `Assalomu alaykum, men Collectius - shaxsiy hisobchingiz va inkassatoringiz.

Men kimning kimga qarzdorligini yozib olaman va qaytarish muddatini eslatib qo&#39;yaman.

Nima qilishni xohlaysiz?`,
		zhCN: `您好，我是 Collectius - 您的私人会计师和收款员。

我可以记录谁欠谁的钱，并提醒何时应该还款。

您想做什么？`},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		arEG: "أريد الحصول على دعوة",
		de:   "Ich hätte gerne einen Code",
		en:   "I want to get an invite",
		es:   "Me gustaría obtener una invitación",
		faIR: "می خواهم یک دعوت دریافت کنم.",
		fr:   "Je veux recevoir une invitation",
		id:   "Saya ingin mendapatkan undangan",
		it:   "Voglio ottenere un invito",
		jaJP: "招待を受けたい",
		koKR: "초대장을 받고 싶어요",
		pl:   "Chcę dostać zaproszenie",
		ptBR: "Quero receber um convite",
		ptPT: "Quero receber um convite",
		ru:   "Хочу получить приглашение",
		tr:   "Davetiye almak istiyorum",
		ukUA: "Я хочу отримати запрошення",
		uz:   "Men taklif olmoqchiman",
		zhCN: "我想获得邀请"},
	COMMAND_TEXT_I_HAVE_INVITE: {
		arEG: "لدي رمز الدعوة",
		de:   "Ich habe einen Code",
		en:   "I have the invitation code",
		es:   "Tengo el código de la invitación",
		faIR: "من کد دعوت را دارم.",
		fr:   "J&#39;ai le code d&#39;invitation",
		id:   "Saya punya kode undangan",
		it:   "Ho il codice invito",
		jaJP: "招待コードを持っています",
		koKR: "초대코드가 있어요",
		pl:   "Mam kod zaproszenia",
		ptBR: "Eu tenho o código de convite",
		ptPT: "Eu tenho o código de convite",
		ru:   "У меня есть код приглашения",
		tr:   "Davetiye kodum var",
		ukUA: "У мене є код запрошення",
		uz:   "Menda taklifnoma kodi bor",
		zhCN: "我有邀请码"},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		arEG: "لم أحصل على أي رسائل بريد إلكتروني",
		de:   "Ich habe noch keine Mails bekommen",
		en:   "I have not got any emails",
		es:   "No he recibido ningún e-mail",
		faIR: "من ایمیلی دریافت نکردم",
		fr:   "Je n&#39;ai reçu aucun e-mail",
		id:   "Saya tidak menerima email apa pun",
		it:   "Non ho ricevuto nessun email",
		jaJP: "メールが来ない",
		koKR: "아직 이메일이 오지 않았습니다",
		pl:   "Nie otrzymałem żadnych e-maili",
		ptBR: "Não recebi nenhum e-mail",
		ptPT: "Não recebi qualquer e-mail",
		ru:   "Я не получил письма на email",
		tr:   "Hiçbir e-posta almadım",
		ukUA: "У мене немає жодних електронних листів",
		uz:   "Hech qanday elektron pochta xabarim yo&#39;q",
		zhCN: "我没有收到任何电子邮件"},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {
		arEG: `<b>%v</b> ،

في الوقت الحالي، يتوفر الروبوت الخاص بنا فقط عن طريق دعوة من الأصدقاء.

إذا لم يكن لديك رمز، فيمكنك ترك تفاصيل الاتصال الخاصة بك وسنرسل لك دعوة بمجرد استحقاق قائمة الانتظار الخاصة بك.

نرسل 10 دعوات يوميًا إلى أولئك الموجودين في رأس قائمة الانتظار ودعوة واحدة عشوائيًا.`,
		de: `<b>%v</b>,

Im Moment ist der Bot leider nur durch Einladungen von Freunden zugänglich.

Wenn du keinen Code hast, lass deine Kontaktdaten da und wir senden dir einen Code sobald du dran bist.

Wir senden 10 Codes am Tag an die, die am längsten warten und einen zufällig.`,
		en: `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,
		es: `<b>%v</b>,

De momento nuestro bot está disponible solo por invitación de amigos.

Si no tienes el código puedes dejarnos tu contacto y te lo enviaremos cuando sea tu turno en la cola .

Enviamos 10 invitaciones por día a los primeros de la cola + 1 de modo casual.`,
		faIR: `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم.

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,
		fr: `<b>%v</b> ,

Pour le moment, notre bot est disponible uniquement sur invitation d&#39;amis.

Si vous n&#39;avez pas de code, vous pouvez laisser vos coordonnées et nous vous enverrons une invitation dès que votre file d&#39;attente sera terminée.

Nous envoyons 10 invitations par jour à ceux qui sont en tête de la file d&#39;attente et 1 au hasard.`,
		id: `<b>%v</b> ,

Saat ini bot kami tersedia hanya berdasarkan undangan dari teman-teman.

Jika Anda tidak mempunyai kode, Anda dapat meninggalkan detail kontak Anda dan kami akan mengirimkan undangan kepada Anda segera setelah antrean Anda berakhir.

Kami mengirimkan 10 undangan per hari kepada mereka yang berada di bagian depan antrean dan 1 secara acak.`,
		it: `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,
		jaJP: `<b>%v</b> 、

現時点では、私たちのボットは友達からの招待によってのみ利用可能です。

コードがない場合は、連絡先の詳細を残していただければ、キューの期限が切れ次第、招待状をお送りします。

キューの先頭にいる人には 1 日あたり 10 通の招待状が送信され、ランダムに 1 通が送信されます。`,
		koKR: `<b>%v</b> ,

현재 저희 봇은 친구의 초대를 통해서만 이용 가능합니다.

코드가 없으신 경우 연락처 정보를 남겨주시면 대기열이 마감되는 대로 초대장을 보내드리겠습니다.

대기열 선두에 있는 분들께는 하루에 초대장 10장을 보내드리고, 무작위로 1장을 보내드립니다.`,
		pl: `<b>%v</b> ,

W tej chwili nasz bot jest dostępny tylko na zaproszenie znajomych.

Jeśli nie masz kodu, możesz zostawić swoje dane kontaktowe, a my wyślemy Ci zaproszenie, gdy tylko skończy się kolejka.

Wysyłamy 10 zaproszeń dziennie do osób znajdujących się na początku kolejki i 1 losowo.`,
		ptBR: `<b>%v</b> ,

No momento, nosso bot está disponível apenas por convite de amigos.

Se você não tiver um código, pode deixar seus dados de contato e nós lhe enviaremos um convite assim que sua fila estiver pronta.

Enviamos 10 convites por dia para aqueles que estão no início da fila e 1 aleatoriamente.`,
		ptPT: `<b>%v</b> ,

Neste momento, o nosso bot só está disponível por convite de amigos.

Se não tiver um código, pode deixar os seus dados de contacto e nós enviar-lhe-emos um convite assim que a sua fila estiver pronta.

Enviamos 10 convites por dia para aqueles que estão no início da fila e 1 aleatoriamente.`,
		ru: `<b>%v</b>,

	На данный момент наш бот доступен только тем кто получил приглашение от друзей.

	Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

	Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
		tr: `<b>%v</b> ,

Şu anda botumuz sadece arkadaşlarınızdan gelen davetle kullanılabilir.

Eğer kodunuz yoksa iletişim bilgilerinizi bırakabilirsiniz ve sıranız geldiğinde size bir davet göndereceğiz.

Sıranın başındakilere günde 10 davet ve rastgele 1 davet gönderiyoruz.`,
		ukUA: `<b>%v</b> ,

Наразі наш бот доступний лише за запрошенням від друзів.

Якщо у вас немає коду, ви можете залишити свої контактні дані, і ми надішлемо вам запрошення, щойно ваша черга буде готова.

Ми надсилаємо 10 запрошень на день тим, хто на чолі черги, та 1 випадковим чином.`,
		uz: `<b>%v</b> ,

Hozirda bizning botimiz faqat doʻstlar taklifiga binoan mavjud.

Agar sizda kod boʻlmasa, kontakt maʼlumotlarini qoldirishingiz mumkin, biz sizga navbat tugashi bilan taklifnoma yuboramiz.

Kuni navbatda turganlarga 10 ta taklif yuboramiz, 1 ta tasodifiy.`,
		zhCN: `<b>%v</b> ,

目前我们的机器人仅接受好友邀请。

如果您没有代码，您可以留下您的联系方式，我们会在您的队列到期后立即向您发送邀请。

我们每天向队列头部的用户发送 10 个邀请，并随机发送 1 个。`},
	EMAIL_INVITE_SUBJ: {
		arEG: "دعوة من {{.FromName}} - الكود: {{.InviteCode}}",
		de:   "Eine Einladung von {{.FromName}} - Code: {{.InviteCode}}",
		en:   "An invite from {{.FromName}} - code: {{.InviteCode}}",
		es:   "La invitación de {{.FromName}} - el código: {{.InviteCode}}",
		faIR: "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
		fr:   "Une invitation de {{.FromName}} - code : {{.InviteCode}}",
		id:   "Undangan dari {{.FromName}} - kode: {{.InviteCode}}",
		it:   "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
		jaJP: "{{.FromName}} からの招待 - コード: {{.InviteCode}}",
		koKR: "{{.FromName}}님의 초대 - 코드: {{.InviteCode}}",
		pl:   "Zaproszenie od {{.FromName}} - kod: {{.InviteCode}}",
		ptBR: "Um convite de {{.FromName}} - código: {{.InviteCode}}",
		ptPT: "Um convite de {{.FromName}} - código: {{.InviteCode}}",
		ru:   "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
		tr:   "{{.FromName}}&#39;dan bir davet - kod: {{.InviteCode}}",
		ukUA: "Запрошення від {{.FromName}} - код: {{.InviteCode}}",
		uz:   "{{.FromName}} taklifi – kod: {{.InviteCode}}",
		zhCN: "来自 {{.FromName}} 的邀请 - 代码：{{.InviteCode}}"},
	SMS_INVITE_TEXT: {
		arEG: `مرحبًا {{.ToName}}، يدعوك {{.FromName}} لتجربة تطبيق تتبع الديون - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

رمز الدعوة الشخصي الخاص بك هو: {{.InviteCode}}`,
		de: `Hey {{.ToName}}, {{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,
		en: `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,
		es: `Hola {{.ToName}}, {{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,
		faIR: `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,
		fr: `Bonjour {{.ToName}}, {{.FromName}} vous invite à essayer l&#39;application de suivi des dettes - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Votre code d&#39;invitation personnel est : {{.InviteCode}}`,
		id: `Hai {{.ToName}}, {{.FromName}} mengundang Anda untuk mencoba aplikasi pelacakan utang - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Kode undangan pribadi Anda adalah: {{.InviteCode}}`,
		it: `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
		jaJP: `こんにちは {{.ToName}}、{{.FromName}} が借金追跡アプリを試すよう招待しています - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

あなたの個人招待コードは次のとおりです: {{.InviteCode}}`,
		koKR: `안녕하세요 {{.ToName}}님, {{.FromName}}님이 부채 추적 앱(https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

)을 사용해 보도록 초대합니다. 귀하의 개인 초대 코드는 {{.InviteCode}}입니다.`,
		pl: `Cześć {{.ToName}}, {{.FromName}} zaprasza Cię do wypróbowania aplikacji do śledzenia długów - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Twój osobisty kod zaproszeniowy to: {{.InviteCode}}`,
		ptBR: `Olá, {{.ToName}}, {{.FromName}} está convidando você para experimentar o aplicativo de rastreamento de dívidas - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Seu código de convite pessoal é: {{.InviteCode}}`,
		ptPT: `Olá, {{.ToName}}, {{.FromName}} está a convidar-vos a experimentar a aplicação de rastreio de dívidas - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

O seu código de convite pessoal é: {{.InviteCode}}`,
		ru: `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
		tr: `Merhaba {{.ToName}}, {{.FromName}} sizi borç takip uygulamasını denemeye davet ediyor - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Kişisel davet kodunuz: {{.InviteCode}}`,
		ukUA: `Вітаємо, {{.ToName}}, {{.FromName}} запрошує вас спробувати додаток для відстеження боргів - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Ваш персональний код запрошення: {{.InviteCode}}`,
		uz: `Salom {{.ToName}}, {{.FromName}} sizni qarzlarni kuzatish ilovasini sinab koʻrishga taklif qilmoqda - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Shaxsiy taklif kodingiz: {{.InviteCode}}`,
		zhCN: `您好{{.ToName}}，{{.FromName}}邀请您试用债务追踪应用程序 - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

您的个人邀请码是：{{.InviteCode}}`},
	EMAIL_INVITE_TEXT: {
		arEG: `مرحبًا {{.ToName}}، يدعوك 

{{.FromName}} لاستخدام تطبيق تتبع الديون - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

رمز الدعوة الخاص بك هو: {{.InviteCode}}`,
		de: `Hey {{.ToName}},

{{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,
		en: `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,
		es: `Hola {{.ToName}},

{{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,
		faIR: `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,
		fr: `Bonjour {{.ToName}},

{{.FromName}} vous invite à utiliser l&#39;application de suivi des dettes : https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Votre code d&#39;invitation est : {{.InviteCode}}`,
		id: `Hai {{.ToName}},

{{.FromName}} mengundang Anda untuk menggunakan aplikasi pelacakan utang - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Kode undangan Anda adalah: {{.InviteCode}}`,
		it: `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,
		jaJP: `こんにちは、{{.ToName}} さん、

{{.FromName}} が借金追跡アプリの使用にあなたを招待しています - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

あなたの招待コードは次のとおりです: {{.InviteCode}}`,
		koKR: `안녕하세요 {{.ToName}}님,

{{.FromName}}님이 부채 추적 앱(https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

)을 사용하도록 초대합니다. 초대 코드는 {{.InviteCode}}입니다.`,
		pl: `Cześć {{.ToName}},

{{.FromName}} zaprasza Cię do korzystania z aplikacji do śledzenia długów - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Twój kod zaproszeniowy to: {{.InviteCode}}`,
		ptBR: `Olá, {{.ToName}},

{{.FromName}} está convidando você para usar o aplicativo de rastreamento de dívidas - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Seu código de convite é: {{.InviteCode}}`,
		ptPT: `Olá, {{.ToName}},

{{.FromName}} está a convidar-vos a utilizar a aplicação de rastreio de dívidas - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

O seu código de convite é: {{.InviteCode}}`,
		ru: `Привет {{.ToName}},

	{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

	Ваш код приглашения: {{.InviteCode}}`,
		tr: `Merhaba {{.ToName}},

{{.FromName}} sizi borç izleme uygulamasını kullanmaya davet ediyor - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Davet kodunuz: {{.InviteCode}}`,
		ukUA: `Вітаємо, {{.ToName}},

{{.FromName}} запрошує вас скористатися додатком для відстеження боргів - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Ваш код запрошення: {{.InviteCode}}`,
		uz: `Salom {{.ToName}},

{{.FromName}} sizni qarzlarni kuzatish ilovasidan foydalanishga taklif qilmoqda - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

Sizning taklif kodi: {{.InviteCode}}`,
		zhCN: `您好{{.ToName}}，

{{.FromName}}邀请您使用债务追踪应用程序 - https://debtstracker.io/invite#id={{.InviteCode}}&amp;telegram-bot={{.TgBot}}&amp;{{.Utm}}

您的邀请码是：{{.InviteCode}}`},
	EMAIL_INVITE_HTML: {
		arEG: `<p style=";text-align:right;direction:rtl">مرحباً {{.ToName}}،</p> {ن}{ن}<p style=";text-align:right;direction:rtl"> {{.FromName}} يدعوك لتجربة <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">تطبيق تتبع الديون</a> .</p> {ن}{ن}<p style=";text-align:right;direction:rtl"> رمز الدعوة الخاص بك هو: <b>{{.InviteCode}}</b></p>`,
		de: `<p>Hey {{.ToName}}, </p>

<p>{{.FromName}} lädt dich ein <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">die neue Schuldentracker App auszuprobieren</a>.</p>

<p>Dein persönlicher Code lautet: <b>{{.InviteCode}}</b></p>`,
		en: `<p>Hi {{.ToName}}, </p>

<p>{{.FromName}} is inviting you to try <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,
		es: `<p>Hola {{.ToName}}, </p>

<p>{{.FromName}} te ha invitado a <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">probar la app para controlar tus deudas</a>.</p>

<p>El código de tu invitación es: <b>{{.InviteCode}}</b></p>`,
		faIR: `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,
		fr: `<p>Bonjour {{.ToName}},</p> 

<p> {{.FromName}} vous invite à essayer <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">l&#39;application de suivi des dettes</a> .</p> 

<p> Votre code d&#39;invitation est : <b>{{.InviteCode}}</b></p>`,
		id: `<p>Hai {{.ToName}},</p> 

<p> {{.FromName}} mengundang Anda untuk mencoba <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">aplikasi pelacakan utang</a> .</p> 

<p> Kode undangan Anda adalah: <b>{{.InviteCode}}</b></p>`,
		it: `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,
		jaJP: `<p>こんにちは、{{.ToName}}さん</p>

<p> {{.FromName}} が<a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">借金追跡アプリ</a>を試すよう招待しています。</p> 

<p>あなたの招待コード: <b>{{.InviteCode}}</b></p>`,
		koKR: `<p>안녕하세요 {{.ToName}}님,</p> 

<p> {{.FromName}}님이 <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">부채 추적 앱을</a> 사용해보라고 초대합니다.</p> 

<p> 귀하의 초대 코드는 <b>{{.InviteCode}}</b> 입니다.</p>`,
		pl: `<p>Cześć {{.ToName}},</p> n. n. n.<p> {{.FromName}} zaprasza Cię do wypróbowania <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">aplikacji do śledzenia długów</a> .</p> n. n. n.<p> Twój kod zaproszeniowy to: <b>{{.InviteCode}}</b></p>`,
		ptBR: `<p>Olá {{.ToName}},</p> 

<p> {{.FromName}} está convidando você a experimentar <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">o aplicativo de rastreamento de dívidas</a> .</p> 

<p> Seu código de convite é: <b>{{.InviteCode}}</b></p>`,
		ptPT: `<p>Olá {{.ToName}},</p> 

<p> {{.FromName}} está a convidá-lo a experimentar <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">a aplicação de rastreamento de dívidas</a> .</p> 

<p> O seu código de convite é: <b>{{.InviteCode}}</b></p>`,
		ru: `<p>Привет {{.ToName}}, </P

	<p>{{.FromName}} приглашает тебя <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

	<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
		tr: `<p>Merhaba {{.ToName}},</p> 

<p> {{.FromName}} sizi <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">borç takip uygulamasını</a> denemeye davet ediyor.</p> 

<p> Davetiye kodunuz: <b>{{.InviteCode}}</b></p>`,
		ukUA: `<p>Привіт, {{.ToName}},</p> 

<p> {{.FromName}} запрошує вас спробувати <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">додаток для відстеження боргів</a> .</p> 

<p> Ваш код запрошення: <b>{{.InviteCode}}</b></p>`,
		uz: `<p>Salom {{.ToName}},</p> 

<p> {{.FromName}} sizni <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">qarzlarni kuzatish ilovasini</a> sinashga taklif qilmoqda.</p> 

<p> Taklif kodingiz: <b>{{.InviteCode}}</b></p>`,
		zhCN: `<p>嗨 {{.ToName}}，</p> 

<p> {{.FromName}} 邀请您尝试<a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">债务追踪应用程序</a>。</p> 

<p>您的邀请码是： <b>{{.InviteCode}}</b></p>`},
	EMAIL_RECEIPT_SUBJ: {
		arEG: "سجل الديون - {{.FromName}}",
		de:   "Schuldschein - {{.FromName}}",
		en:   "Debt record - {{.FromName}}",
		es:   "La notificación de la deuda - {{.FromName}}",
		faIR: "سوابق بدهی - {{.FromName}}",
		fr:   "Historique de la dette - {{.FromName}}",
		id:   "Catatan hutang - {{.FromName}}",
		it:   "Debito - {{.FromName}}",
		jaJP: "債務記録 - {{.FromName}}",
		koKR: "부채 기록 - {{.FromName}}",
		pl:   "Rejestr długów - {{.FromName}}",
		ptBR: "Registro de dívida - {{.FromName}}",
		ptPT: "Registo de dívida - {{.FromName}}",
		ru:   "Запись о долге - {{.FromName}}",
		tr:   "Borç kaydı - {{.FromName}}",
		ukUA: "Історія боргу - {{.FromName}}",
		uz:   "Qarz rekordi - {{.FromName}}",
		zhCN: "债务记录 - {{.FromName}}"},
	EMAIL_RECEIPT_BODY_TEXT: {
		arEG: "{{.FromName}} أنشأ سجل ديون: {{.ReceiptURL}}",
		de:   "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		en:   "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		es:   "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		faIR: "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		fr:   "{{.FromName}} a créé un enregistrement de dette : {{.ReceiptURL}}",
		id:   "{{.FromName}} membuat catatan utang: {{.ReceiptURL}}",
		it:   "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
		jaJP: "{{.FromName}} が債務記録を作成しました: {{.ReceiptURL}}",
		koKR: "{{.FromName}}님이 부채 기록을 생성했습니다: {{.ReceiptURL}}",
		pl:   "{{.FromName}} utworzył rekord długu: {{.ReceiptURL}}",
		ptBR: "{{.FromName}} criou um registro de dívida: {{.ReceiptURL}}",
		ptPT: "{{.FromName}} criou um registo de dívida: {{.ReceiptURL}}",
		ru:   "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		tr:   "{{.FromName}} bir borç kaydı oluşturdu: {{.ReceiptURL}}",
		ukUA: "{{.FromName}} створив(-ла) борговий запис: {{.ReceiptURL}}",
		uz:   "{{.FromName}} qarz rekordini yaratdi: {{.ReceiptURL}}",
		zhCN: "{{.FromName}}创建了债务记录：{{.ReceiptURL}}"},
	MESSENGER_RECEIPT_TEXT: {
		arEG: "لقد قمت بإنشاء سجل ديون لك، راجع التفاصيل في {{.ReceiptURL}}",
		de:   "Ich habe online einen Schuldschein erstellt, für mehr Details siehe {{.ReceiptURL}}",
		en:   "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		es:   "He creado una notificación de la deuda, las detalles están aquí {{.ReceiptURL}}",
		faIR: "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
		fr:   "J&#39;ai créé un dossier de dette vous concernant, voir les détails sur {{.ReceiptURL}}",
		id:   "Saya telah membuat catatan utang mengenai Anda, lihat detailnya di {{.ReceiptURL}}",
		it:   "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
		jaJP: "債務記録を作成しました。詳細は{{.ReceiptURL}}をご覧ください。",
		koKR: "안녕하세요. 부채 기록을 생성했습니다. 자세한 내용은 {{.ReceiptURL}}에서 확인하세요.",
		pl:   "Utworzyłem dla Ciebie rekord należności, szczegóły znajdziesz na {{.ReceiptURL}}",
		ptBR: "Criei um registro de dívida para você, veja os detalhes em {{.ReceiptURL}}",
		ptPT: "Criei um registo de dívida para si, veja os detalhes em {{.ReceiptURL}}",
		ru:   "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
		tr:   "Sizinle ilgili bir borç kaydı oluşturdum, ayrıntıları {{.ReceiptURL}} adresinde görebilirsiniz",
		ukUA: "Я створив(-ла) для вас борговий довідник, див. деталі на {{.ReceiptURL}}",
		uz:   "Sizni hurmat qilish uchun qarz yozuvini tuzdim, tafsilotlarni {{.ReceiptURL}} sahifasida koʻring",
		zhCN: "我已经创建了您的债务记录，详情请参阅{{.ReceiptURL}}"},
	EMAIL_RECEIPT_BODY_HTML: {
		arEG: "{{.FromName}} أنشأ سجل ديون: {{.ReceiptURL}}",
		de:   "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		en:   "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		es:   "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		faIR: "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		fr:   "{{.FromName}} a créé un enregistrement de dette : {{.ReceiptURL}}",
		id:   "{{.FromName}} membuat catatan utang: {{.ReceiptURL}}",
		it:   "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
		jaJP: "{{.FromName}} が債務記録を作成しました: {{.ReceiptURL}}",
		koKR: "{{.FromName}}님이 부채 기록을 생성했습니다: {{.ReceiptURL}}",
		pl:   "{{.FromName}} utworzył rekord długu: {{.ReceiptURL}}",
		ptBR: "{{.FromName}} criou um registro de dívida: {{.ReceiptURL}}",
		ptPT: "{{.FromName}} criou um registo de dívida: {{.ReceiptURL}}",
		ru:   "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
		tr:   "{{.FromName}} bir borç kaydı oluşturdu: {{.ReceiptURL}}",
		ukUA: "{{.FromName}} створив(-ла) борговий запис: {{.ReceiptURL}}",
		uz:   "{{.FromName}} qarz rekordini yaratdi: {{.ReceiptURL}}",
		zhCN: "{{.FromName}}创建了债务记录：{{.ReceiptURL}}"},
	INLINE_RECEIPT_TITLE: {
		arEG: "الإيصال: %v",
		de:   "Empfänger: %v",
		en:   "Receipt: %v",
		es:   "El recibo: %v",
		faIR: "رسید: %v",
		fr:   "Reçu : %v",
		id:   "Penerimaan: %v",
		it:   "Debito/credito: %v",
		jaJP: "領収書: %v",
		koKR: "영수증: %v",
		pl:   "Paragon: %v",
		ptBR: "Recibo: %v",
		ptPT: "Recibo: %v",
		ru:   "Квитанция: %v",
		tr:   "Makbuz: %v",
		ukUA: "Квитанція: %v",
		uz:   "Kvitansiya: %v",
		zhCN: "收据：%v"},
	INLINE_RECEIPT_DESCRIPTION: {
		arEG: "انقر هنا لإرسال الإيصال",
		de:   "Klicken sie hier, um die Quittung zu sehen",
		en:   "Click here to send the receipt",
		es:   "Haz click aquí para enviar el recibo",
		faIR: "برای ارسال رسید اینجا کلیک کنید.",
		fr:   "Cliquez ici pour envoyer le reçu",
		id:   "Klik di sini untuk mengirim tanda terima",
		it:   "Clicca qui per inviare un debito/credito",
		jaJP: "領収書を送信するにはここをクリックしてください",
		koKR: "영수증을 보내려면 여기를 클릭하세요",
		pl:   "Kliknij tutaj, aby wysłać paragon",
		ptBR: "Clique aqui para enviar o recibo",
		ptPT: "Clique aqui para enviar o recibo",
		ru:   "Нажмите здесь чтобы отправить квитанцию",
		tr:   "Makbuzu göndermek için buraya tıklayın",
		ukUA: "Натисніть тут, щоб надіслати квитанцію",
		uz:   "Kvitansiyani yuborish uchun shu yerni bosing",
		zhCN: "点击此处发送收据"},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		arEG: "<b>يرجى اختيار اللغة لرؤية تفاصيل الديون</b> التي تم تسجيلها بواسطة {{.Creator}}.",
		de:   "<b>Bitte wählen Sie eine Sprache, um weitere Details zu sehen, die den Schuldschein betreffen</b>, der erstellt wurde von {{.Creator}}.",
		en:   "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		es:   "<b>Elige el idioma para ver los detalles de la deuda</b> que ha sido creada por {{.Creator}}.",
		faIR: "<b> لطفا برای رویت جزئیات بدهی که توسط </b>  {{.Creator}} ثبت شده است زبان را انتخاب کنید.",
		fr:   "<b>Veuillez choisir la langue pour voir les détails de la dette</b> qui a été enregistrée par {{.Creator}}.",
		id:   "<b>Silakan pilih bahasa untuk melihat rincian utang</b> yang telah dicatat oleh {{.Creator}}.",
		it:   "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
		jaJP: "{{.Creator}} によって記録された<b>債務の詳細を表示するには、言語を選択してください</b>。",
		koKR: "{{.Creator}}에 의해 기록된 <b>부채의 세부 정보를 보려면 언어를 선택하세요</b> .",
		pl:   "<b>Proszę wybrać język, aby zobaczyć szczegóły długu</b> zarejestrowanego przez {{.Creator}}.",
		ptBR: "<b>Selecione o idioma para ver detalhes da dívida</b> registrada por {{.Creator}}.",
		ptPT: "<b>Selecione o idioma para ver os detalhes da dívida</b> registada por {{.Creator}}.",
		ru:   "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
		tr:   "{{.Creator}} tarafından kaydedilen <b>borcun ayrıntılarını görmek için lütfen dili seçin</b> .",
		ukUA: "<b>Будь ласка, виберіть мову, щоб переглянути деталі боргу</b> , зареєстрованого користувачем {{.Creator}}.",
		uz:   "{{.Creator}} tomonidan qayd etilgan <b>qarz tafsilotlarini koʻrish uchun tilni tanlang</b> .",
		zhCN: "<b>请选择语言来查看{{.Creator}}记录的债务详情</b>。"},
	INLINE_RECEIPT_FOOTER: {
		arEG: `{{.SiteLink}} — سيساعدك تطبيق تتبع الديون على:

 - معرفة صافي ربحك دائمًا

 - سداد الديون في الوقت المحدد
 <i>(تذكير لك ولمدينيك)</i>`,
		de: `{{.SiteLink}} — eine App, die dir hilft Schulden zu überwachen:

  - Du weißt immer, wie viel du allen schuldest

  - Keine Fälligkeit wird verpasst
    <i>(erinnert dich und die Gläubiger)</i>`,
		en: `{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		es: `{{.SiteLink}} — la app para controlar tus deudas te ayuda a:

  - Saber siempre quién debe a quién

  - Devolver la deuda a tiempo
    <i>(recordatorio a ti y a tus deudores)</i>`,
		faIR: `{{.SiteLink}} — یک برنامه پیگیری بدهی است که به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی ها به موقع پرداخت شوند.
    <i>(با ارسال یادآوری به  شما و بدهکاران )</i>`,
		fr: `{{.SiteLink}} — une application de suivi des dettes vous aidera à :

 - Toujours connaître votre résultat net ;

 - Rembourser vos dettes à temps ;
 <i>(rappels pour vous et vos débiteurs)</i>`,
		id: `{{.SiteLink}} — aplikasi untuk melacak utang akan membantu Anda untuk:

 - Selalu mengetahui batas akhir Anda

 - Mengembalikan utang tepat waktu
 <i>(pengingat untuk Anda &amp; debitur Anda)</i>`,
		it: `{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
		jaJP: `{{.SiteLink}} — 借金追跡アプリは次のことに役立ちます:

 - 常に収支を把握する

 - 借金を期日通りに返済する
 <i>(あなたと債務者へのリマインダー)</i>`,
		koKR: `{{.SiteLink}} — 부채 추적 앱은 다음과 같은 데 도움이 됩니다.

 - 항상 최종 이익을 파악합니다.

 - 제때 부채를 갚습니다.
 <i>(본인과 채무자에게 알림)</i>`,
		pl: `{{.SiteLink}} — aplikacja do śledzenia długów pomoże Ci:

 - Zawsze znać swój wynik finansowy

 - Spłacać długi na czas
 <i>(przypomnienia dla Ciebie i Twoich dłużników)</i>`,
		ptBR: `{{.SiteLink}} — um aplicativo para controle de dívidas ajudará você a:

 - Sempre saber seu lucro

 - Pagar dívidas em dia
 <i>(lembretes para você e seus devedores)</i>`,
		ptPT: `{{.SiteLink}} — uma aplicação para controlo de dívidas irá ajudá-lo a:

 - Saber sempre o seu lucro

 - Pagar as dívidas a tempo
 <i>(lembretes para si e para os seus devedores)</i>`,
		ru: `{{.SiteLink}} — программа для учёта долгов поможет:

	  - Всегда знать кто кому сколько должен

	  - Незабыть вовремя отдать или востребовать долг
	    <i>(напоминания вам и вашим должникам)</i>`,
		tr: `{{.SiteLink}} — borç takibi için bir uygulama size şunlarda yardımcı olacaktır:

 - Her zaman alt sınırınızı bilin

 - Borçları zamanında ödeyin
 <i>(size ve borçlularınıza hatırlatmalar)</i>`,
		ukUA: `{{.SiteLink}} — додаток для відстеження боргів допоможе вам:

 - Завжди знати свій прибуток

 - Вчасно повертати борги
 <i>(нагадання вам та вашим боржникам)</i>`,
		uz: `{{.SiteLink}} — qarzlarni kuzatish uchun ilova sizga quyidagilarga yordam beradi:

 - Har doim oʻz daromadingizni bilish

 - Qarzlarni oʻz vaqtida qaytarish
 <i>(siz va qarzdorlaringizga eslatma)</i>`,
		zhCN: `{{.SiteLink}} — 一款债务追踪应用程序将帮助您：

 - 始终了解您的底线

 - 按时偿还债务
 <i>（提醒您和您的债务人）</i>`},
	INLINE_RECEIPT_GENERATING_MESSAGE: {
		arEG: `<b>{{.Creator}} سجل دينًا</b> مرتبطًا بك.

 &gt;&gt; <i>جاري إنشاء الإيصال...</i>`,
		de: `<b>{{.Creator}} erstellte online einen Schuldschein</b> der dich betrifft.

>> Generating receipt`,
		en: `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <i>Generating receipt...</i>`,
		es: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <i>Generating receipt...</i>`,
		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <i>Generating receipt...</i>`,
		fr: `<b>{{.Creator}} a enregistré une dette</b> associée à vous.

 &gt;&gt; <i>Génération du reçu...</i>`,
		id: `<b>{{.Creator}} mencatat utang</b> yang terkait dengan Anda.

 &gt;&gt; <i>Membuat tanda terima...</i>`,
		it: `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <i>Generating receipt...</i>`,
		jaJP: `<b>{{.Creator}} はあなたに関連する負債を記録しました</b>。

 &gt;&gt;<i>領収書を生成しています...</i>`,
		koKR:
		//	INLINE_RECEIPT_MESSAGE: {
		//
		//		en: `<b>{{.Creator}} recorded a debt</b> associated with you.
		//
		//`,
		//
		//		es: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.
		//
		//`,
		//
		//		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.
		//
		//`,
		//
		//		it:   `<b>{{.Creator}} ha registrato un debito</b> associato a te.
		//
		//`,
		//
		//		ru:  `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.
		//
		//`,
		//
		//	},
		`<b>{{.Creator}}님이 귀하와 관련된 부채를 기록했습니다</b> .

 &gt;&gt; <i>영수증 생성 중...</i>`,
		pl: `<b>{{.Creator}} zarejestrował dług</b> powiązany z Tobą.

 &gt;&gt; <i>Generowanie paragonu...</i>`,
		ptBR: `<b>{{.Creator}} registrou uma dívida</b> associada a você.

 &gt;&gt; <i>Gerando recibo...</i>`,
		ptPT: `<b>{{.Creator}} registou uma dívida</b> associada a si.

 &gt;&gt; <i>Gerando recibo...</i>`,
		ru: `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <i>Generating receipt...</i>`,
		tr: `<b>{{.Creator}} sizinle ilişkili bir borcu kaydetti</b> .

 &gt;&gt; <i>Makbuz oluşturuluyor...</i>`,
		ukUA: `<b>{{.Creator}} зареєстрував борг,</b> пов&#39;язаний з вами.

 &gt;&gt; <i>Генерування квитанції...</i>`,
		uz: `<b>{{.Creator}} siz bilan bog‘liq qarzni qayd etdi</b> .

 &gt;&gt; <i>Kvitansiya yaratilmoqda...</i>`,
		zhCN: `<b>{{.Creator}}记录了与您相关的债务</b>。

 &gt;&gt;<i>正在生成收据...</i>`},

	INLINE_RECEIPT_MESSAGE: {
		arEG: `<b>{{.Creator}} سجل دينًا</b> مرتبطًا بك.

 &gt;&gt; <a href="{{.ReceiptUrl}}">انقر هنا لعرض الإيصال</a>`,
		de: `<b>{{.Creator}} hat eine mit Ihnen verbundene Schuld erfasst</b> .

 &gt;&gt; <a href="{{.ReceiptUrl}}">Klicken Sie hier, um die Quittung anzuzeigen</a>`,
		en: `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		es: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		fr: `<b>{{.Creator}} a enregistré une dette</b> vous concernant.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Cliquez ici pour voir le reçu</a>`,
		id: `<b>{{.Creator}} mencatat utang</b> yang terkait dengan Anda.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Klik di sini untuk melihat tanda terima</a>`,
		it: `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		jaJP: `<b>{{.Creator}} はあなたに関連する負債を記録しました</b>。

 &gt;&gt;<a href="{{.ReceiptUrl}}">領収書を表示するにはここをクリックしてください</a>`,
		koKR: `<b>{{.Creator}}님이 귀하와 관련된 부채를 기록했습니다</b> .

 &gt;&gt; <a href="{{.ReceiptUrl}}">영수증을 보려면 여기를 클릭하세요.</a>`,
		pl: `<b>{{.Creator}} zarejestrował dług</b> związany z Tobą.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Kliknij tutaj, aby wyświetlić potwierdzenie</a>`,
		ptBR: `<b>{{.Creator}} registrou uma dívida</b> associada a você.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Clique aqui para ver o recibo</a>`,
		ptPT: `<b>{{.Creator}} registou uma dívida</b> associada a si.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Clique aqui para ver o recibo</a>`,
		ru: `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <a href="{{.ReceiptUrl}}">Посмотреть квитанцию</a>`,
		tr: `<b>{{.Creator}} sizinle ilişkili bir borcu kaydetti</b> .

 &gt;&gt; <a href="{{.ReceiptUrl}}">Makbuzu görüntülemek için buraya tıklayın</a>`,
		ukUA: `<b>{{.Creator}} зареєстрував борг,</b> пов&#39;язаний з вами.

 &gt;&gt; <a href="{{.ReceiptUrl}}">Натисніть тут, щоб переглянути квитанцію</a>`,
		uz: `<b>{{.Creator}} siz bilan bog‘liq qarzni qayd etdi</b> .

 &gt;&gt; <a href="{{.ReceiptUrl}}">Kvitansiyani ko‘rish uchun shu yerni bosing</a>`,
		zhCN: `<b>{{.Creator}}记录了与您相关的债务</b>。

 &gt;&gt;<a href="{{.ReceiptUrl}}">单击此处查看收据</a>`},
	InlineInviteToJoinFamilyTitle: {
		arEG: "دعوة للانضمام إلى العائلة في @%s",
		de:   "Einladung der Familie beizutreten bei @%s",
		en:   "Invitation to join family at @%s",
		es:   "Invitación para unirse a la familia en @%s",
		faIR: "دعوت به پیوستن به خانواده در @%s",
		fr:   "Invitation à rejoindre la famille à @%s",
		id:   "Undangan untuk bergabung dengan keluarga di @%s",
		it:   "Invito a unirsi alla famiglia a @%s",
		jaJP: "@%s のファミリーへの招待",
		koKR: "@%s 가족에 가입하라는 초대",
		pl:   "Zaproszenie do dołączenia do rodziny w @%s",
		ptBR: "Convite para se juntar à família em @%s",
		ptPT: "Convite para se juntar à família em @%s",
		ru:   "Приглашение в семью на @%s",
		tr:   "@%s ailesine katılma daveti",
		ukUA: "Запрошення приєднатися до родини в @%s",
		uz:   "@%s oilasiga qo&#39;shilishga taklif",
		zhCN: "邀请加入@%s 的家庭"},
	InlineInviteToJoinFamilyDescription: {
		arEG: "انقر هنا لإرسال دعوة",
		de:   "Klick hier, um eine Einladung zu versenden",
		en:   "Click here to send an invite",
		es:   "Haz click para enviar la invitación",
		faIR: "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
		fr:   "Cliquez ici pour envoyer une invitation",
		id:   "Klik di sini untuk mengirim undangan",
		it:   "Clicca qui per spedire un invito",
		jaJP: "招待を送信するにはここをクリックしてください",
		koKR: "여기를 클릭하여 초대장을 보내세요",
		pl:   "Kliknij tutaj, aby wysłać zaproszenie",
		ptBR: "Clique aqui para enviar um convite",
		ptPT: "Clique aqui para enviar um convite",
		ru:   "Нажмите здесь для отправки приглашения",
		tr:   "Davetiye göndermek için buraya tıklayın",
		ukUA: "Натисніть тут, щоб надіслати запрошення",
		uz:   "Taklif yuborish uchun shu yerni bosing",
		zhCN: "单击此处发送邀请"},
	YouAreInvitedToJoinFamilyMessage: {
		arEG: "لقد تمت دعوتك للانضمام إلى حساب العائلة على @{BOT_ID}.",
		de:   "Sie sind eingeladen, dem Familienkonto bei @{BOT_ID} beizutreten.",
		en:   "You are invited to join family account at @{BOT_ID}.",
		es:   "Se te invita a unirte a la cuenta familiar en @{BOT_ID}.",
		faIR: "شما دعوت شده\u200cاید که به حساب خانواده در @{BOT_ID} بپیوندید.",
		fr:   "Vous êtes invité à rejoindre le compte familial sur @{BOT_ID}.",
		id:   "Anda diundang untuk bergabung dengan akun keluarga di @{BOT_ID}.",
		it:   "Sei invitato a unirti al conto familiare su @{BOT_ID}.",
		jaJP: "@{BOT_ID}で家族アカウントに参加するように招待されています。",
		koKR: "@{BOT_ID}에서 가족 계정에 초대되었습니다.",
		pl:   "Zostałeś zaproszony do dołączenia do konta rodzinnego na @{BOT_ID}.",
		ptBR: "Você foi convidado a entrar na conta familiar em @{BOT_ID}.",
		ptPT: "Está convidado a aderir à conta familiar em @{BOT_ID}.",
		ru:   "Вас пригласили присоединиться к семейному аккаунту на @{BOT_ID}.",
		tr:   "@{BOT_ID} üzerinde aile hesabına katılmaya davetlisiniz.",
		ukUA: "Вас запросили приєднатися до сімейного акаунта на @{BOT_ID}.",
		uz:   "Sizni @{BOT_ID} oilaviy hisobiga qo‘shilishga taklif qilishdi.",
		zhCN: "您被邀请加入 @{BOT_ID} 的家庭账户。",
	},
	BtnViewFamilyAccount: {
		arEG: "عرض حساب العائلة",
		de:   "Familienkonto ansehen",
		en:   "View family account",
		es:   "Ver cuenta familiar",
		faIR: "مشاهده حساب خانواده",
		fr:   "Voir le compte familial",
		id:   "Lihat akun keluarga",
		it:   "Visualizza conto familiare",
		jaJP: "家族アカウントを表示",
		koKR: "가족 계정 보기",
		pl:   "Zobacz konto rodzinne",
		ptBR: "Ver conta familiar",
		ptPT: "Ver conta familiar",
		ru:   "Посмотреть семейный аккаунт",
		tr:   "Aile hesabını görüntüle",
		ukUA: "Переглянути сімейний акаунт",
		uz:   "Oilaviy hisobni ko‘rish",
		zhCN: "查看家庭账户",
	},
	SMS_RECEIPT_YOU_GOT: {
		arEG: "لقد حصلت على %v من %v.",
		de:   "Du hast dir %v von %v geliehen.",
		en:   "You've got %v from %v.",
		es:   "Has recibido %v de %v.",
		faIR: "شما دریافت کرده اید %v از %v.",
		fr:   "Vous avez obtenu %v de %v.",
		id:   "Anda mendapat %v dari %v.",
		it:   "Hai ricevuto %v da %v",
		jaJP: "%v から %v を取得しました。",
		koKR: "%v에서 %v를 얻었습니다.",
		pl:   "Masz %v z %v.",
		ptBR: "Você obteve %v de %v.",
		ptPT: "Obteve %v de %v.",
		ru:   "Вы получили %v от %v.",
		tr:   "%v&#39;den %v&#39;yi elde ettiniz.",
		ukUA: "Ви отримали %v від %v.",
		uz:   "Sizda %v dan %v bor.",
		zhCN: "您已从 %v 得到 %v。"},
	SMS_RECEIPT_YOU_GAVE: {
		arEG: "لقد أعطيت %v إلى %v.",
		de:   "Du hast %v an %v verliehen.",
		en:   "You've given %v to %v.",
		es:   "Has dado %v a %v.",
		faIR: "شما پرداخت کرده اید %v به %v.",
		fr:   "Vous avez donné %v à %v.",
		id:   "Anda telah memberikan %v ke %v.",
		it:   "Hai dato %v a %v",
		jaJP: "%v を %v に渡しました。",
		koKR: "%v에게 %v를 주었습니다.",
		pl:   "Przekazałeś %v %v.",
		ptBR: "Você deu %v para %v.",
		ptPT: "Você deu %v a %v.",
		ru:   "Вы дали %v - взял(а) %v.",
		tr:   "%v&#39;ye %v verdiniz.",
		ukUA: "Ви передали %v користувачеві %v.",
		uz:   "Siz %v ga %v berdingiz.",
		zhCN: "您已将 %v 给予 %v。"},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		arEG: "انقر فوق %v للتأكيد أو الرفض.",
		de:   "Klicke %v um zu akzeptieren oder abzulehnen.",
		en:   "Click %v to confirm or decline.",
		es:   "Haz click %v para confirmar o rechazar.",
		faIR: "کلیک کنید %v تا رد خود را تایید نمایید",
		fr:   "Cliquez sur %v pour confirmer ou refuser.",
		id:   "Klik %v untuk mengonfirmasi atau menolak.",
		it:   "Clicca %v per confermare o rifiutare",
		jaJP: "確認または拒否するには %v をクリックします。",
		koKR: "확인 또는 거부하려면 %v를 클릭하세요.",
		pl:   "Kliknij %v, aby potwierdzić lub odrzucić.",
		ptBR: "Clique em %v para confirmar ou recusar.",
		ptPT: "Clique em %v para confirmar ou recusar.",
		ru:   "Нажмите %v чтобы подтвердить или отвергнуть.",
		tr:   "Onaylamak veya reddetmek için %v&#39;ye tıklayın.",
		ukUA: "Натисніть %v, щоб підтвердити або відхилити.",
		uz:   "Tasdiqlash yoki rad etish uchun %v tugmasini bosing.",
		zhCN: "单击 %v 确认或拒绝。"},
	HTML_DATE: {
		arEG: "تاريخ",
		de:   "Datum",
		en:   "Date",
		es:   "Fecha",
		faIR: "تاریخ",
		fr:   "Date",
		id:   "Tanggal",
		it:   "Data",
		jaJP: "日付",
		koKR: "날짜",
		pl:   "Data",
		ptBR: "Data",
		ptPT: "Data",
		ru:   "Дата",
		tr:   "Tarih",
		ukUA: "Дата",
		uz:   "Sana",
		zhCN: "日期"},
	HTML_RECEIPT: {
		arEG: "إيصال",
		de:   "Empfänger",
		en:   "Receipt",
		es:   "Recibo",
		faIR: "رسید",
		fr:   "Reçu",
		id:   "Kuitansi",
		it:   "Scontrino",
		jaJP: "レシート",
		koKR: "영수증",
		pl:   "Paragon",
		ptBR: "Recibo",
		ptPT: "Recibo",
		ru:   "Квитанция",
		tr:   "Fiş",
		ukUA: "Квитанція",
		uz:   "Kvitansiya",
		zhCN: "收据"},
	HTML_AMOUNT: {
		arEG: "كمية",
		de:   "Betrag",
		en:   "Amount",
		es:   "Importe",
		faIR: "مقدار",
		fr:   "Montant",
		id:   "Jumlah",
		it:   "Totale",
		jaJP: "額",
		koKR: "양",
		pl:   "Kwota",
		ptBR: "Quantia",
		ptPT: "Montante",
		ru:   "Сумма",
		tr:   "Miktar",
		ukUA: "Сума",
		uz:   "Miqdori",
		zhCN: "数量"},
	HTML_FROM: {
		arEG: "من",
		de:   "Von",
		en:   "From",
		es:   "De",
		faIR: "از",
		fr:   "Depuis",
		id:   "Dari",
		it:   "Da",
		jaJP: "から",
		koKR: "에서",
		pl:   "Z",
		ptBR: "De",
		ptPT: "De",
		ru:   "Дал",
		tr:   "İtibaren",
		ukUA: "Від",
		uz:   "Kimdan",
		zhCN: "从"},
	HTML_TO: {
		arEG: "ل",
		de:   "An",
		en:   "To",
		es:   "Para",
		faIR: "به",
		fr:   "À",
		id:   "Ke",
		it:   "Per",
		jaJP: "に",
		koKR: "에게",
		pl:   "Do",
		ptBR: "Para",
		ptPT: "Para",
		ru:   "Получил",
		tr:   "İle",
		ukUA: "До",
		uz:   "Kimga",
		zhCN: "到"},
	NO_NAME: {
		arEG: "لا اسم",
		de:   "unbekannt",
		en:   "no name",
		es:   "sin nombre",
		faIR: "بدون نام",
		fr:   "pas de nom",
		id:   "tidak ada nama",
		it:   "senza nome",
		jaJP: "名前なし",
		koKR: "이름 없음",
		pl:   "bez nazwy",
		ptBR: "sem nome",
		ptPT: "sem nome",
		ru:   "без имени",
		tr:   "isim yok",
		ukUA: "без імені",
		uz:   "ism yo&#39;q",
		zhCN: "没有名字"},
	TELEGRAM_RECEIPT: {
		arEG: "{{.FromName}} أنشأ سجل ديون ({{.TransferCurrency}})",
		de:   "{{.FromName}} hat einen Schuldschein erstellt ({{.TransferCurrency}})",
		en:   "{{.FromName}} created a debt record ({{.TransferCurrency}})",
		es:   "{{.FromName}} ha creado una deuda ({{.TransferCurrency}})",
		faIR: "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
		fr:   "{{.FromName}} a créé un enregistrement de dette ({{.TransferCurrency}})",
		id:   "{{.FromName}} membuat catatan utang ({{.TransferCurrency}})",
		it:   "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
		jaJP: "{{.FromName}} が債務記録 ({{.TransferCurrency}}) を作成しました",
		koKR: "{{.FromName}}님이 부채 기록({{.TransferCurrency}})을 생성했습니다.",
		pl:   "{{.FromName}} utworzył rekord długu ({{.TransferCurrency}})",
		ptBR: "{{.FromName}} criou um registro de dívida ({{.TransferCurrency}})",
		ptPT: "{{.FromName}} criou um registo de dívida ({{.TransferCurrency}})",
		ru:   "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
		tr:   "{{.FromName}} bir borç kaydı oluşturdu ({{.TransferCurrency}})",
		ukUA: "{{.FromName}} створив(-ла) борговий запис ({{.TransferCurrency}})",
		uz:   "{{.FromName}} qarz rekordini yaratdi ({{.TransferCurrency}})",
		zhCN: "{{.FromName}} 创建了债务记录 ({{.TransferCurrency}})"},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		arEG: "الرجاء الاختيار من الخيارات المقدمة.",
		de:   "Bitte wähle aus den angezeigten Optionen.",
		en:   "Please choose from provided options.",
		es:   "Por favor, elige una de las siguientes opciones.",
		faIR: "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
		fr:   "Veuillez choisir parmi les options proposées.",
		id:   "Silakan pilih dari opsi yang disediakan.",
		it:   "Scegli tra le opzioni fornite.",
		jaJP: "提供されたオプションから選択してください。",
		koKR: "제공된 옵션 중에서 선택하십시오.",
		pl:   "Proszę wybrać z podanych opcji.",
		ptBR: "Por favor, escolha entre as opções fornecidas.",
		ptPT: "Escolha entre as opções fornecidas.",
		ru:   "Пожалуйста выберете из предоставленных опций.",
		tr:   "Lütfen sağlanan seçeneklerden birini seçin.",
		ukUA: "Будь ласка, виберіть із наданих опцій.",
		uz:   "Iltimos, taqdim etilgan variantlardan tanlang.",
		zhCN: "请从提供的选项中选择。",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		arEG: "<b>هل ترغب بإضافة ملاحظة أو تعليق؟</b> المذكرات هي سجلات خاصة بك. التعليقات متاحة للجميع ممن لديهم الإذن لعرض هذه المعاملة. %v %v",
		de:   "<b>Möchtest du eine Bemerkung oder Notiz hinzufügen?</b>\n%v Deine Notizen kannst nur du sehen.\n%v Eine Bemerkung wird quasi auf dem Schuldschein und der Quittung vermerkt und ist insofern für beide sichtbar.",
		en:   "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		es:   "<b>¿Quieres añadir una nota o comentario?</b>\n%v Las notas se graban de manera privada para tu propia información.\n%v Los comentarios son visibles para todos los autorizados a ver esta transacción.",
		faIR: "<b>آیا میخواهید یادداشت یا شرحی اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v شرح در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
		fr:   "<b>Voulez-vous ajouter une note ou un commentaire?</b>\n%v Les mémos sont des enregistrements privés pour votre propre référence.\n%v Les commentaires sont accessibles à tous ceux qui ont la permission de voir cette transaction.",
		id:   "<b>Apakah Anda ingin menambahkan catatan atau komentar?</b>\n%v Memo adalah catatan pribadi untuk referensi Anda sendiri.\n%v Komentar tersedia untuk semua orang yang memiliki izin untuk melihat transaksi ini.",
		it:   "<b>Vuoi aggiungere una nota o un commento?</b> \n%v I memo sono record privati per il riferimento di yoru.\n%v I commenti sono disponibili a tutti coloro che hanno l'autorizzazione a visualizzare questa transazione.",
		jaJP: "<b>メモやコメントを追加しますか？</b>\n%v メモはあなた自身の参照用のプライベートな記録です。\n%v コメントはこのトランザクションを閲覧する権限を持つすべての人が利用できます。",
		koKR: "<b>메모나 댓글을 추가하시겠습니까?</b>\n%v 메모는 자신의 참조를 위한 개인 기록입니다.\n%v 댓글은 이 거래를 볼 수 있는 권한이 있는 모든 사람이 이용할 수 있습니다.",
		pl:   "<b>Czy chcesz dodać notatkę lub komentarz?</b>\n%v Notatki są prywatnymi zapisami do własnego użytku.\n%v Komentarze są dostępne dla wszystkich, którzy mają uprawnienia do przeglądania tej transakcji.",
		ptBR: "<b>Deseja adicionar uma nota ou comentário?</b>\n%v Memorandos são registros privados para sua própria referência.\n%v Comentários estão disponíveis para todos que têm permissão para visualizar esta transação.",
		ptPT: "<b>Deseja adicionar uma nota ou comentário?</b> \n%v Os memorandos são registos privados para sua referência. \n%v Os comentários estão disponíveis para todos os que têm permissão para visualizar esta transação.",
		ru:   "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личного пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		tr:   "<b>Not veya yorum eklemek istiyor musunuz?</b>\n%v Notlar kendi referansınız için özel kayıtlardır.\n%v Yorumlar, bu işlemi görüntüleme izni olan herkes tarafından görülebilir.",
		ukUA: "<b>Хочете додати нотатку або коментар?</b>\n%v Нотатки зберігаються для вашого особистого користування.\n%v Коментар видно всім, кому дозволено перегляд цієї транзакції.",
		uz:   "<b>Eslatma yoki izoh qo'shmoqchimisiz?</b>\n%v Eslatmalar shaxsiy yozuvlar bo'lib, o'zingiz uchun ma'lumotdir.\n%v Izohlar ushbu tranzaksiyani ko'rish huquqiga ega bo'lgan har bir kishi uchun mavjud.",
		zhCN: "<b>您想添加备注或评论吗？</b>\n%v 备忘录是供您自己参考的私人记录。\n%v 评论对所有有权查看此交易的人可见。",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		arEG: "يرجى كتابة ملاحظة:",
		de:   "Bitte schreibe eine Notiz:",
		en:   "Please write a note:",
		es:   "Por favor, escribe una nota:",
		faIR: "لطفاً یک یادداشت بنویسید:",
		fr:   "Veuillez écrire une note :",
		id:   "Silakan tulis catatan:",
		it:   "Per favore scrivi un appunto:",
		jaJP: "メモを書いてください：",
		koKR: "메모를 작성해 주세요:",
		pl:   "Proszę napisać notatkę:",
		ptBR: "Por favor, escreva uma nota:",
		ptPT: "Por favor, escreva uma nota:",
		ru:   "Напишите заметку:",
		tr:   "Lütfen bir not yazın:",
		ukUA: "Напишіть нотатку:",
		uz:   "Iltimos, eslatma yozing:",
		zhCN: "请写一个备注：",
	},
	COMMAND_TEXT_MORE_ABOUT_INTEREST_COMMAND: {
		arEG: "المزيد عن الاهتمام",
		de:   "Mehr über Zinsen",
		en:   "More about interest",
		es:   "Más sobre intereses",
		faIR: "اطلاعات بیشتر درباره بهره",
		fr:   "Plus d'informations sur les intérêts",
		id:   "Lebih lanjut tentang bunga",
		it:   "Più informazioni sugli interessi",
		jaJP: "利息についての詳細",
		koKR: "이자에 대한 자세한 정보",
		pl:   "Więcej o odsetkach",
		ptBR: "Mais sobre juros",
		ptPT: "Mais sobre juros",
		ru:   "Подробнее о процентах",
		tr:   "Faiz hakkında daha fazla",
		ukUA: "Детальніше про відсотки",
		uz:   "Foiz haqida ko'proq",
		zhCN: "更多关于利息的信息",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_SHORT: {
		arEG: `<b>الفائدة والملاحظات</b> 

لتعيين معدل الفائدة والفترة، أرسل رسالة بالتنسيق التالي:
<pre style=";text-align:right;direction:rtl"> النسبة المئوية/الفترة المئوية/الفترة الدنيا/فترة النعمة:ملاحظة</pre>`,
		de: `<b>Zinsen und Kommentar</b>

Um den Zinssatz und den Zeitraum festzulegen, senden Sie eine Nachricht im folgenden Format:
		<pre>prozent/prozent_zeitraum/min_zeitraum/karenzzeit:notiz</pre>`,
		en: `<b>Interest & notes</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`,
		es: `<b>Interés y comentario</b>

Para establecer la tasa de interés y el período, envía un mensaje en el siguiente formato:
		<pre>porcentaje/período_porcentaje/período_mínimo/período_gracia:nota</pre>`,
		faIR: `<b>نرخ بهره و یادداشت</b>

برای تنظیم نرخ بهره و دوره، پیامی به فرمت زیر ارسال کنید:
		<pre>درصد/دوره_درصد/دوره_حداقل/دوره_تنفس:یادداشت</pre>`,
		fr: `<b>Intérêt et notes</b>

Pour définir le taux d'intérêt et la période, envoyez un message au format suivant:
		<pre>pourcentage/période_pourcentage/période_min/période_grâce:note</pre>`,
		id: `<b>Bunga & catatan</b>

Untuk mengatur suku bunga & periode kirim pesan dalam format berikut:
		<pre>persen/periode_persen/periode_min/masa_tenggang:catatan</pre>`,
		it: `<b>Interessi e note</b>

Per impostare il tasso di interesse e il periodo, invia un messaggio nel seguente formato:
		<pre>percentuale/periodo_percentuale/periodo_minimo/periodo_grazia:nota</pre>`,
		jaJP: `<b>利息とメモ</b>

金利と期間を設定するには、次の形式でメッセージを送信してください：
		<pre>パーセント/パーセント期間/最小期間/猶予期間:メモ</pre>`,
		koKR: `<b>이자 및 메모</b>

이자율 및 기간을 설정하려면 다음 형식으로 메시지를 보내십시오:
		<pre>퍼센트/퍼센트_기간/최소_기간/유예_기간:메모</pre>`,
		pl: `<b>Odsetki i notatki</b>

Aby ustawić stopę procentową i okres, wyślij wiadomość w następującym formacie:
		<pre>procent/okres_procentowy/min_okres/okres_karencji:notatka</pre>`,
		ptBR: `<b>Juros e notas</b>

Para definir a taxa de juros e o período, envie uma mensagem no seguinte formato:
		<pre>percentual/período_percentual/período_mínimo/período_carência:nota</pre>`,
		ptPT: `<b>Juros e notas</b> 

Para definir a taxa de juro e o período, envie uma mensagem no seguinte formato:
<pre> percentagem/período_percentagem/período_min/período_de_carência:nota</pre>`,
		ru: `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:
	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>`,
		tr: `<b>Faiz ve notlar</b>

Faiz oranını ve dönemi ayarlamak için aşağıdaki formatta bir mesaj gönderin:
		<pre>yüzde/yüzde_dönem/min_dönem/grace_dönem:not</pre>`,
		ukUA: `<b>Відсоток і коментар</b>

Щоб встановити відсоткову ставку і період, надішліть повідомлення в наступному форматі:
		<pre>відсоток/відсотковий_період/мінімальний_період/грейс_період:примітка</pre>`,
		uz: `<b>Foiz va eslatmalar</b>

Foiz stavkasi va davrni o'rnatish uchun quyidagi formatda xabar yuboring:
		<pre>foiz/foiz_davri/min_davr/imtiyoz_davri:eslatma</pre>`,
		zhCN: `<b>利息和备注</b>

要设置利率和期限，请按以下格式发送消息：
		<pre>百分比/百分比期限/最小期限/宽限期:备注</pre>`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_LONG: {
		arEG: `<b>الفائدة والملاحظات</b> 

لتعيين معدل الفائدة والفترة، أرسل رسالة بالتنسيق التالي:

<pre style=";text-align:right;direction:rtl"> النسبة المئوية/الفترة المئوية/الفترة الدنيا/فترة النعمة:ملاحظة</pre> 

حيث ( <i>أول معاملين مطلوبان</i> ): 
 * <code>percent</code> - حتى رقمين بعد الفاصلة. 
 * <code>min_period</code> - عدد أيام فترة الفائدة. 
 * <code>min_perdio</code> - الحد الأدنى لعدد الأيام لحساب الفائدة. افتراضيًا، يكون الرقم 1 ولا يمكن أن يكون أقل. 
 * <code>grace_period</code> - فترة بدون فوائد. في الوقت الحالي، لا يمكنك تعيين فترة سماح وفترة حد أدنى في نفس الوقت. 
 * <code>note</code> - أي نص توضيحي سيكون مرئيًا لك ولطرفك المقابل. 

يتم حساب الفائدة يوميًا ( <i>كل 24 ساعة</i> ) باستخدام صيغة <a href="https://www.investopedia.com/terms/s/simple_interest.asp#utm_source=DebtsTrackerBot&utm_medium=telegram&utm_campaign=new_debt_wizard">النسبة المئوية البسيطة</a> . 

 <b>الأمثلة</b> : 

 <code>2/7/5</code> - 2% أسبوعيًا، الحد الأدنى لمدة 5 أيام 
 <code>15/360</code> - 15%/سنة، ( <i>الحد الأدنى لمدة يوم واحد</i> ) 
 <code>3/30/0/10</code> - 3% شهريًا مع فترة سماح لمدة 10 أيام 

❗ وظيفة النسبة المئوية في مرحلة الاختبار التجريبي، يرجى إخبارنا في @DebtsTrackerGroup إذا كان أي شيء لا يعمل كما تتوقع.`,
		de: `<b>Zinsen & Notizen</b>

Um den Zinssatz und Zeitraum festzulegen, senden Sie eine Nachricht im folgenden Format:

		<pre>prozent/prozent_zeitraum/min_zeitraum/karenzzeit:notiz</pre>

Wobei (<i>die ersten 2 Parameter erforderlich sind</i>):
 * <code>prozent</code> - bis zu 2 Stellen nach dem Komma.
 * <code>prozent_zeitraum</code> - Anzahl der Tage für den Zinszeitraum.
 * <code>min_zeitraum</code> - Mindestanzahl von Tagen für die Zinsberechnung. Standardmäßig 1 und kann nicht weniger sein.
 * <code>karenzzeit</code> - zinsfreier Zeitraum. Derzeit können Sie nicht gleichzeitig einen Mindest- und einen Karenzzeit festlegen.
 * <code>notiz</code> - jeder erklärende Text, der für Sie und Ihren Geschäftspartner sichtbar ist.

Die Zinsen werden täglich (<i>alle 24 Stunden</i>) nach der Formel des <a href="https://de.wikipedia.org/wiki/Zinsrechnung#Einfache_Verzinsung">einfachen Zinses</a> berechnet.

<b>Beispiele</b>:

		<code>2/7/5</code> - 2% pro Woche, mindestens für 5 Tage
		<code>15/360</code> - 15% pro Jahr, (<i>mindestens für 1 Tag</i>)
		<code>3/30/0/10</code> - 3% pro Monat mit 10 Tagen Karenzzeit

❗ Die % Funktionalität befindet sich noch in der BETA-Testphase. Bitte teilen Sie uns in @DebtsTrackerGroup mit, wenn etwas nicht wie erwartet funktioniert.`,
		en: `<b>Interest & notes</b>

To set interest rate & period send a message in following format:

		<pre>percent/percent_period/min_period/grace_period:note</pre>

Where (<i>first 2 params are required</i>):
 * <code>percent</code> - up to 2 digits after comma.
 * <code>min_period</code> - number of days for interest period.
 * <code>min_perdio</code> - minimum number of days for interest calculation. Is 1 by default and can't be less.'.
 * <code>grace_period</code> - interest-less period. At the moment you can't set grace & min period at the same time'.
 * <code>note</code> - any explanatory text that will be visible to you and your counterparty.

Interest is calculated daily (<i>every 24 hours</i>) using <a href="https://www.investopedia.com/terms/s/simple_interest.asp#utm_source=DebtsTrackerBot&utm_medium=telegram&utm_campaign=new_debt_wizard">simple percent</a> formula.

<b>Examples</b>:

		<code>2/7/5</code> - 2% per week, minimum for 5 days
		<code>15/360</code> - 15%/year, (<i>minimum for 1 day</i>)
		<code>3/30/0/10</code> - 3% per month with 10 days grace period

❗ The % functionality is in BETA testing stage, please let us know in @DebtsTrackerGroup if anything works not as you would expect.`,
		es: `<b>Interés y notas</b>

Para establecer la tasa de interés y el período, envía un mensaje en el siguiente formato:

		<pre>porcentaje/período_porcentaje/período_mínimo/período_gracia:nota</pre>

Donde (<i>los primeros 2 parámetros son obligatorios</i>):
 * <code>porcentaje</code> - hasta 2 dígitos después de la coma.
 * <code>período_porcentaje</code> - número de días para el período de interés.
 * <code>período_mínimo</code> - número mínimo de días para el cálculo de intereses. Es 1 por defecto y no puede ser menor.
 * <code>período_gracia</code> - período sin intereses. Por el momento, no puedes establecer períodos de gracia y mínimo al mismo tiempo.
 * <code>nota</code> - cualquier texto explicativo que será visible para ti y tu contraparte.

El interés se calcula diariamente (<i>cada 24 horas</i>) utilizando la fórmula de <a href="https://es.wikipedia.org/wiki/Inter%C3%A9s_simple">interés simple</a>.

<b>Ejemplos</b>:

		<code>2/7/5</code> - 2% por semana, mínimo por 5 días
		<code>15/360</code> - 15% anual, (<i>mínimo por 1 día</i>)
		<code>3/30/0/10</code> - 3% por mes con 10 días de período de gracia

❗ La funcionalidad de % está en fase de prueba BETA, por favor háganoslo saber en @DebtsTrackerGroup si algo no funciona como esperabas.`,
		faIR: //nolint:staticcheck // disable ST1018 for this line
		`<b>نرخ بهره و یادداشت‌ها</b>

برای تنظیم نرخ بهره و دوره، پیامی به فرمت زیر ارسال کنید:

		<pre>درصد/دوره_درصد/دوره_حداقل/دوره_تنفس:یادداشت</pre>

جایی که (<i>2 پارامتر اول الزامی هستند</i>):
 * <code>درصد</code> - تا 2 رقم بعد از اعشار.
 * <code>دوره_درصد</code> - تعداد روزهای دوره بهره.
 * <code>دوره_حداقل</code> - حداقل تعداد روزها برای محاسبه بهره. به طور پیش‌فرض 1 است و نمی‌تواند کمتر باشد.
 * <code>دوره_تنفس</code> - دوره بدون بهره. در حال حاضر نمی‌توانید همزمان دوره حداقل و دوره تنفس را تعیین کنید.
 * <code>یادداشت</code> - هر متن توضیحی که برای شما و طرف مقابل قابل مشاهده خواهد بود.

بهره به صورت روزانه (<i>هر 24 ساعت</i>) با استفاده از فرمول <a href="https://fa.wikipedia.org/wiki/%D8%A8%D9%87%D8%B1%D9%87_%D8%B3%D8%A7%D8%AF%D9%87">بهره ساده</a> محاسبه می‌شود.

<b>مثال‌ها</b>:

		<code>2/7/5</code> - 2% در هفته، حداقل برای 5 روز
		<code>15/360</code> - 15% در سال، (<i>حداقل برای 1 روز</i>)
		<code>3/30/0/10</code> - 3% در ماه با 10 روز دوره تنفس

❗ قابلیت درصد در مرحله آزمایشی BETA است، لطفاً اگر چیزی مطابق انتظار شما کار نمی‌کند، به ما در @DebtsTrackerGroup اطلاع دهید.`,
		fr: `<b>Intérêts et notes</b>

Pour définir le taux d'intérêt et la période, envoyez un message au format suivant :

		<pre>pourcentage/période_pourcentage/période_min/période_grâce:note</pre>

Où (<i>les 2 premiers paramètres sont obligatoires</i>) :
 * <code>pourcentage</code> - jusqu'à 2 chiffres après la virgule.
 * <code>période_pourcentage</code> - nombre de jours pour la période d'intérêt.
 * <code>période_min</code> - nombre minimum de jours pour le calcul des intérêts. Est 1 par défaut et ne peut pas être inférieur.
 * <code>période_grâce</code> - période sans intérêt. Pour le moment, vous ne pouvez pas définir la période de grâce et la période minimale en même temps.
 * <code>note</code> - tout texte explicatif qui sera visible pour vous et votre contrepartie.

Les intérêts sont calculés quotidiennement (<i>toutes les 24 heures</i>) selon la formule de <a href="https://fr.wikipedia.org/wiki/Int%C3%A9r%C3%AAt_simple">l'intérêt simple</a>.

<b>Exemples</b> :

		<code>2/7/5</code> - 2% par semaine, minimum pour 5 jours
		<code>15/360</code> - 15% par an, (<i>minimum pour 1 jour</i>)
		<code>3/30/0/10</code> - 3% par mois avec 10 jours de période de grâce

❗ La fonctionnalité % est en phase de test BETA, veuillez nous informer dans @DebtsTrackerGroup si quelque chose ne fonctionne pas comme prévu.`,
		id: `<b>Bunga & catatan</b>

Untuk mengatur suku bunga & periode kirim pesan dalam format berikut:

		<pre>persen/periode_persen/periode_min/masa_tenggang:catatan</pre>

Dimana (<i>2 parameter pertama wajib diisi</i>):
 * <code>persen</code> - hingga 2 digit setelah koma.
 * <code>periode_persen</code> - jumlah hari untuk periode bunga.
 * <code>periode_min</code> - jumlah hari minimum untuk perhitungan bunga. Secara default adalah 1 dan tidak bisa kurang.
 * <code>masa_tenggang</code> - periode tanpa bunga. Saat ini Anda tidak dapat mengatur masa tenggang & periode minimum pada saat yang sama.
 * <code>catatan</code> - teks penjelasan apa pun yang akan terlihat oleh Anda dan pihak lawan Anda.

Bunga dihitung harian (<i>setiap 24 jam</i>) menggunakan rumus <a href="https://id.wikipedia.org/wiki/Bunga_sederhana">bunga sederhana</a>.

<b>Contoh</b>:

		<code>2/7/5</code> - 2% per minggu, minimal untuk 5 hari
		<code>15/360</code> - 15%/tahun, (<i>minimal untuk 1 hari</i>)
		<code>3/30/0/10</code> - 3% per bulan dengan masa tenggang 10 hari

❗ Fungsi % masih dalam tahap pengujian BETA, harap beri tahu kami di @DebtsTrackerGroup jika ada yang tidak berfungsi seperti yang Anda harapkan.`,
		it: `<b>Interessi e note</b>

Per impostare il tasso di interesse e il periodo, invia un messaggio nel seguente formato:

		<pre>percentuale/periodo_percentuale/periodo_minimo/periodo_grazia:nota</pre>

Dove (<i>i primi 2 parametri sono obbligatori</i>):
 * <code>percentuale</code> - fino a 2 cifre dopo la virgola.
 * <code>periodo_percentuale</code> - numero di giorni per il periodo di interesse.
 * <code>periodo_minimo</code> - numero minimo di giorni per il calcolo degli interessi. È 1 per impostazione predefinita e non può essere inferiore.
 * <code>periodo_grazia</code> - periodo senza interessi. Al momento non è possibile impostare contemporaneamente il periodo di grazia e il periodo minimo.
 * <code>nota</code> - qualsiasi testo esplicativo che sarà visibile a te e alla tua controparte.

Gli interessi vengono calcolati giornalmente (<i>ogni 24 ore</i>) utilizzando la formula dell'<a href="https://it.wikipedia.org/wiki/Interesse_semplice">interesse semplice</a>.

<b>Esempi</b>:

		<code>2/7/5</code> - 2% a settimana, minimo per 5 giorni
		<code>15/360</code> - 15% all'anno, (<i>minimo per 1 giorno</i>)
		<code>3/30/0/10</code> - 3% al mese con 10 giorni di periodo di grazia

❗ La funzionalità % è in fase di test BETA, ti preghiamo di farci sapere in @DebtsTrackerGroup se qualcosa non funziona come previsto.`,
		jaJP: `<b>利息とメモ</b>

金利と期間を設定するには、次の形式でメッセージを送信してください：

		<pre>パーセント/パーセント期間/最小期間/猶予期間:メモ</pre>

以下の説明（<i>最初の2つのパラメータは必須</i>）：
 * <code>パーセント</code> - 小数点以下2桁まで。
 * <code>パーセント期間</code> - 利息期間の日数。
 * <code>最小期間</code> - 利息計算のための最小日数。デフォルトは1で、それより少なくすることはできません。
 * <code>猶予期間</code> - 無利息期間。現時点では、猶予期間と最小期間を同時に設定することはできません。
 * <code>メモ</code> - あなたと相手に表示される説明テキスト。

利息は毎日（<i>24時間ごと</i>）<a href="https://ja.wikipedia.org/wiki/%E5%8D%98%E5%88%A9">単利</a>の計算式を使用して計算されます。

<b>例</b>：

		<code>2/7/5</code> - 週2%、最低5日間
		<code>15/360</code> - 年15%、（<i>最低1日</i>）
		<code>3/30/0/10</code> - 月3%、10日間の猶予期間付き

❗ %機能はBETAテスト段階にあります。期待通りに動作しない場合は@DebtsTrackerGroupでお知らせください。`,
		koKR: `<b>이자 및 메모</b>

이자율 및 기간을 설정하려면 다음 형식으로 메시지를 보내십시오:

		<pre>퍼센트/퍼센트_기간/최소_기간/유예_기간:메모</pre>

여기서 (<i>처음 2개 매개변수는 필수</i>):
 * <code>퍼센트</code> - 소수점 이하 최대 2자리까지.
 * <code>퍼센트_기간</code> - 이자 기간의 일수.
 * <code>최소_기간</code> - 이자 계산을 위한 최소 일수. 기본값은 1이며 더 적을 수 없습니다.
 * <code>유예_기간</code> - 무이자 기간. 현재는 유예 기간과 최소 기간을 동시에 설정할 수 없습니다.
 * <code>메모</code> - 귀하와 상대방에게 표시될 설명 텍스트.

이자는 매일 (<i>24시간마다</i>) <a href="https://ko.wikipedia.org/wiki/%EB%8B%A8%EB%A6%AC">단리</a> 공식을 사용하여 계산됩니다.

<b>예시</b>:

		<code>2/7/5</code> - 주 2%, 최소 5일
		<code>15/360</code> - 연 15%, (<i>최소 1일</i>)
		<code>3/30/0/10</code> - 월 3%, 10일 유예 기간 포함

❗ % 기능은 BETA 테스트 단계에 있으므로, 예상대로 작동하지 않는 경우 @DebtsTrackerGroup에 알려주시기 바랍니다.`,
		pl: `<b>Odsetki i notatki</b>

Aby ustawić stopę procentową i okres, wyślij wiadomość w następującym formacie:

		<pre>procent/okres_procentowy/min_okres/okres_karencji:notatka</pre>

Gdzie (<i>pierwsze 2 parametry są wymagane</i>):
 * <code>procent</code> - do 2 cyfr po przecinku.
 * <code>okres_procentowy</code> - liczba dni dla okresu odsetkowego.
 * <code>min_okres</code> - minimalna liczba dni do obliczenia odsetek. Domyślnie wynosi 1 i nie może być mniejsza.
 * <code>okres_karencji</code> - okres bez odsetek. W tej chwili nie można ustawić jednocześnie okresu karencji i minimalnego okresu.
 * <code>notatka</code> - dowolny tekst wyjaśniający, który będzie widoczny dla Ciebie i Twojego kontrahenta.

Odsetki są obliczane codziennie (<i>co 24 godziny</i>) przy użyciu formuły <a href="https://pl.wikipedia.org/wiki/Odsetki_proste">odsetek prostych</a>.

<b>Przykłady</b>:

		<code>2/7/5</code> - 2% tygodniowo, minimum na 5 dni
		<code>15/360</code> - 15% rocznie, (<i>minimum na 1 dzień</i>)
		<code>3/30/0/10</code> - 3% miesięcznie z 10-dniowym okresem karencji

❗ Funkcjonalność % jest w fazie testów BETA, prosimy o informację w @DebtsTrackerGroup, jeśli coś nie działa zgodnie z oczekiwaniami.`,
		ptBR: `<b>Juros e notas</b>

Para definir a taxa de juros e o período, envie uma mensagem no seguinte formato:

		<pre>percentual/período_percentual/período_mínimo/período_carência:nota</pre>

Onde (<i>os primeiros 2 parâmetros são obrigatórios</i>):
 * <code>percentual</code> - até 2 dígitos após a vírgula.
 * <code>período_percentual</code> - número de dias para o período de juros.
 * <code>período_mínimo</code> - número mínimo de dias para o cálculo de juros. É 1 por padrão e não pode ser menor.
 * <code>período_carência</code> - período sem juros. No momento, você não pode definir o período de carência e o período mínimo ao mesmo tempo.
 * <code>nota</code> - qualquer texto explicativo que será visível para você e sua contraparte.

Os juros são calculados diariamente (<i>a cada 24 horas</i>) usando a fórmula de <a href="https://pt.wikipedia.org/wiki/Juros_simples">juros simples</a>.

<b>Exemplos</b>:

		<code>2/7/5</code> - 2% por semana, mínimo de 5 dias
		<code>15/360</code> - 15% ao ano, (<i>mínimo de 1 dia</i>)
		<code>3/30/0/10</code> - 3% ao mês com período de carência de 10 dias

❗ A funcionalidade de % está em fase de teste BETA, por favor nos informe em @DebtsTrackerGroup se algo não funcionar como esperado.`,
		ptPT: `<b>Juros e notas</b> 

Para definir a taxa de juro e o período, envie uma mensagem no seguinte formato:

<pre> percentagem/período_percentagem/período_min/período_de_carência:nota</pre> 

Onde ( <i>os 2 primeiros parâmetros são obrigatórios</i> ):
 * <code>percent</code> - até 2 dígitos após a vírgula.
 * <code>min_period</code> - número de dias para o período de juros.
 * <code>min_perdio</code> - número mínimo de dias para cálculo de juros. É 1 por defeito e não pode ser inferior.&#39;.
 * <code>grace_period</code> - período sem juros. Neste momento, não pode definir o período de carência e o período mínimo ao mesmo tempo. 
 * <code>note</code> : qualquer texto explicativo que seja visível para si e para a sua contraparte. 

Os juros são calculados diariamente ( <i>de 24 em 24 horas</i> ) através de uma fórmula <a href="https://www.investopedia.com/terms/s/simple_interest.asp#utm_source=DebtsTrackerBot&utm_medium=telegram&utm_campaign=new_debt_wizard">de percentagem simples</a> . 

 <b>Exemplos</b> : 

 <code>2/7/5</code> - 2% por semana, no mínimo durante 5 dias. 
 <code>15/360</code> - 15%/ano ( <i>mínimo durante 1 dia</i> ). 
 <code>3/30/0/10</code> - 3% ao mês com um período de carência de 10 dias. 

❗ A funcionalidade % está em fase de teste BETA. Informe-nos em @DebtsTrackerGroup se algo não funcionar como espera.`,
		ru: `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:

	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>

Где (<i>первые 2 параметра обязательны</i>):
 * <code>процент</code> - возможно до 2х знаков после запятой.
 * <code>процентный_период</code> - количество дней за которые начисляется процент указанный предыдущим числом.
 * <code>минимальный_период</code> - минимальное количество дней за которые сразу начисляются проценты. По умолчанию 1 и не может быть меньше.
 * <code>грэйс_период</code> - безпроцентный период. Пока нельзя задавать и минимальный процент и безпроцентный период одновременно.
 * <code>комментарий</code> - любой поясняющий текст видимый вам и другому участнику сделки.

Начисление ежедневно (<i>каждые 24 часа</i>) по формуле <a href="https://ru.wikipedia.org/wiki/%D0%9F%D1%80%D0%BE%D1%81%D1%82%D1%8B%D0%B5_%D0%BF%D1%80%D0%BE%D1%86%D0%B5%D0%BD%D1%82%D1%8B">простого процента</a>.

<b>Примеры</b>:

		<code>2/7/5</code> - под 2% в неделю, минимум на 5 дней
		<code>15/360</code> - под 15% годовых
		<code>3/30/0/10</code> - под 3% в месяц с безпроцентным периодом в 10 дней

❗ Функционал % ещё тестируется, пожалуйста сообщите в @DebtsTrackerGroup если что-то пошло не так.`,
		tr: `<b>Faiz ve notlar</b>

Faiz oranını ve dönemi ayarlamak için aşağıdaki formatta bir mesaj gönderin:

		<pre>yüzde/yüzde_dönem/min_dönem/grace_dönem:not</pre>

Burada (<i>ilk 2 parametre zorunludur</i>):
 * <code>yüzde</code> - virgülden sonra en fazla 2 basamak.
 * <code>yüzde_dönem</code> - faiz dönemi için gün sayısı.
 * <code>min_dönem</code> - faiz hesaplaması için minimum gün sayısı. Varsayılan olarak 1'dir ve daha az olamaz.
 * <code>grace_dönem</code> - faizsiz dönem. Şu anda aynı anda hem grace hem de minimum dönemi ayarlayamazsınız.
 * <code>not</code> - size ve karşı tarafınıza görünecek herhangi bir açıklayıcı metin.

Faiz, <a href="https://tr.wikipedia.org/wiki/Basit_faiz">basit faiz</a> formülü kullanılarak günlük olarak (<i>her 24 saatte bir</i>) hesaplanır.

<b>Örnekler</b>:

		<code>2/7/5</code> - haftada %2, minimum 5 gün için
		<code>15/360</code> - yıllık %15, (<i>minimum 1 gün için</i>)
		<code>3/30/0/10</code> - 10 günlük grace dönemi ile ayda %3

❗ % işlevselliği BETA test aşamasındadır, beklendiği gibi çalışmayan bir şey varsa lütfen @DebtsTrackerGroup'ta bize bildirin.`,
		ukUA: `<b>Відсоток і коментар</b>

Щоб встановити відсоткову ставку і період, надішліть повідомлення в наступному форматі:

		<pre>відсоток/відсотковий_період/мінімальний_період/грейс_період:примітка</pre>

Де (<i>перші 2 параметри обов'язкові</i>):
 * <code>відсоток</code> - можливо до 2-х знаків після коми.
 * <code>відсотковий_період</code> - кількість днів за які нараховується відсоток, вказаний попереднім числом.
 * <code>мінімальний_період</code> - мінімальна кількість днів за які відразу нараховуються відсотки. За замовчуванням 1 і не може бути менше.
 * <code>грейс_період</code> - безвідсотковий період. Поки не можна задавати і мінімальний відсоток і безвідсотковий період одночасно.
 * <code>примітка</code> - будь-який пояснювальний текст видимий вам і іншому учаснику угоди.

Нарахування щодня (<i>кожні 24 години</i>) за формулою <a href="https://uk.wikipedia.org/wiki/%D0%9F%D1%80%D0%BE%D1%81%D1%82%D1%96_%D0%B2%D1%96%D0%B4%D1%81%D0%BE%D1%82%D0%BA%D0%B8">простого відсотка</a>.

<b>Приклади</b>:

		<code>2/7/5</code> - під 2% на тиждень, мінімум на 5 днів
		<code>15/360</code> - під 15% річних
		<code>3/30/0/10</code> - під 3% на місяць з безвідсотковим періодом в 10 днів

❗ Функціонал % ще тестується, будь ласка, повідомте в @DebtsTrackerGroup якщо щось пішло не так.`,
		uz: `<b>Foiz va eslatmalar</b>

Foiz stavkasi va davrni o'rnatish uchun quyidagi formatda xabar yuboring:

		<pre>foiz/foiz_davri/min_davr/imtiyoz_davri:eslatma</pre>

Quyidagilar (<i>birinchi 2 ta parametr talab qilinadi</i>):
 * <code>foiz</code> - verguldan keyin 2 ta raqamgacha.
 * <code>foiz_davri</code> - foiz davri uchun kunlar soni.
 * <code>min_davr</code> - foiz hisoblanishi uchun minimal kunlar soni. Sukut bo'yicha 1 va bundan kam bo'lishi mumkin emas.
 * <code>imtiyoz_davri</code> - foizsiz davr. Hozirda siz bir vaqtning o'zida ham imtiyoz, ham minimal davrni o'rnatishingiz mumkin emas.
 * <code>eslatma</code> - siz va hamkoringiz uchun ko'rinadigan har qanday tushuntiruvchi matn.

Foiz kunlik (<i>har 24 soatda</i>) <a href="https://uz.wikipedia.org/wiki/Oddiy_foiz">oddiy foiz</a> formulasi yordamida hisoblanadi.

<b>Misollar</b>:

		<code>2/7/5</code> - haftasiga 2%, minimal 5 kun uchun
		<code>15/360</code> - yillik 15%, (<i>minimal 1 kun uchun</i>)
		<code>3/30/0/10</code> - oyiga 3%, 10 kunlik imtiyoz davri bilan

❗ % funksiyasi BETA sinov bosqichida, agar biror narsa kutganingizdek ishlamasa, @DebtsTrackerGroup'da bizga xabar bering.`,
		zhCN: `<b>利息和备注</b>

要设置利率和期限，请按以下格式发送消息：

		<pre>百分比/百分比期限/最小期限/宽限期:备注</pre>

其中（<i>前2个参数为必填</i>）：
 * <code>百分比</code> - 小数点后最多2位数字。
 * <code>百分比期限</code> - 利息期限的天数。
 * <code>最小期限</code> - 计算利息的最小天数。默认为1，不能更少。
 * <code>宽限期</code> - 无息期。目前您不能同时设置宽限期和最小期限。
 * <code>备注</code> - 任何对您和您的交易对手可见的解释性文本。

利息每日计算（<i>每24小时</i>），使用<a href="https://zh.wikipedia.org/wiki/%E5%8D%95%E5%88%A9">单利</a>公式。

<b>示例</b>：

		<code>2/7/5</code> - 每周2%，最少5天
		<code>15/360</code> - 年利率15%，（<i>最少1天</i>）
		<code>3/30/0/10</code> - 每月3%，10天宽限期

❗ %功能正处于BETA测试阶段，如果有任何不符合您预期的情况，请在@DebtsTrackerGroup告知我们。`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		arEG: `إذا كنت تريد إضافة تعليق فقط أرسل رسالة نصية الآن.`,
		de:   "Falls du eine Bemerkung auf den Schuldschein schreiben willst, schick mir jetzt den Text.",
		en:   `If you want to add a comment just send a text now.`,
		es:   `Si quieres añadir un comentario simplemente envía un texto ahora.`,
		faIR: `شما می توانید یک شرح اضافه کنید. تنها کافیست یک متن ارسال کنید.`,
		fr:   `Si vous souhaitez ajouter un commentaire, envoyez simplement un texte maintenant.`,
		id:   `Jika Anda ingin menambahkan komentar, cukup kirim teks sekarang.`,
		it:   `Se vuoi aggiungere un commento invia del testo ora.`,
		jaJP: `コメントを追加したい場合は、今すぐテキストを送信してください。`,
		koKR: `댓글을 추가하려면 지금 텍스트를 보내세요.`,
		pl:   `Jeśli chcesz dodać komentarz, po prostu wyślij tekst teraz.`,
		ptBR: `Se você quiser adicionar um comentário, basta enviar um texto agora.`,
		ptPT: `Se quiser adicionar um comentário, basta enviar uma mensagem agora.`,
		ru:   `Если хотите добавить комментарий просто отправьте текст.`,
		tr:   `Yorum eklemek istiyorsanız şimdi bir metin gönderin.`,
		ukUA: `Якщо ви хочете додати коментар, просто надішліть текст зараз.`,
		uz:   `Izoh qo'shmoqchi bo'lsangiz, hozir matn yuboring.`,
		zhCN: `如果您想添加评论，请立即发送文本。`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		arEG: "مرئية لك &amp; %v",
		de:   "sichtbar für dich & %v",
		en:   "visible to you & %v",
		es:   "visible solo para ti & %v",
		faIR: "قابل مشاهده برای شما & %v",
		fr:   "visible pour vous et %v",
		id:   "terlihat oleh Anda &amp; %v",
		it:   "visibile solo a te e %v",
		jaJP: "あなたに表示されます &amp; %v",
		koKR: "당신과 %v에게 보여짐",
		pl:   "widoczne dla Ciebie i %v",
		ptBR: "visível para você e %v",
		ptPT: "visível para si e %v",
		ru:   "виден вам и %v",
		tr:   "sizin tarafınızdan görülebilir &amp; %v",
		ukUA: "видиме для вас і %v",
		uz:   "sizga va %v",
		zhCN: "您可见 &amp; %v"},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		arEG: "الرجاء كتابة التعليق:",
		de:   "Schreibe nun eine Bemerkung auf den Schuldschein:",
		en:   "Please write the comment:",
		es:   "Por favor, escribe un comentario:",
		faIR: "لطفاً شرح را ثبت کنید:",
		fr:   "S&#39;il vous plaît écrire le commentaire:",
		id:   "Silakan tulis komentar:",
		it:   "Per favore scrivi un commento:",
		jaJP: "コメントを書いてください:",
		koKR: "댓글을 적어주세요:",
		pl:   "Proszę napisać komentarz:",
		ptBR: "Por favor escreva o comentário:",
		ptPT: "Por favor escreva o comentário:",
		ru:   "Напишите комментарий:",
		tr:   "Lütfen yorumu yazın:",
		ukUA: "Будь ласка, напишіть коментар:",
		uz:   "Iltimos, sharh yozing:",
		zhCN: "请写下评论："},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		arEG: "تمت إضافة المذكرات. هل ترغب بكتابة تعليق؟",
		de:   "Deine Notiz wurde hinzugefügt, möchtest du noch eine Bemerkung auf den Schuldschein schreiben?",
		en:   "Memo have been added. Do you want to write a comment?",
		es:   "La nota está añadida. ¿Quieres escribir un comentario?",
		faIR: "یادداشت اضافه شد. آیا میخواهید یک شرح ثبت کنید؟",
		fr:   "Des mémos ont été ajoutés. Souhaitez-vous laisser un commentaire ?",
		id:   "Memo telah ditambahkan. Apakah Anda ingin menulis komentar?",
		it:   "Promemoria aggiunto. Vuoi scrivere un commento?",
		jaJP: "メモを追加しました。コメントを書き込みますか？",
		koKR: "메모가 추가되었습니다. 댓글을 남기시겠습니까?",
		pl:   "Dodano notatki. Czy chcesz napisać komentarz?",
		ptBR: "Um memorando foi adicionado. Deseja escrever um comentário?",
		ptPT: "Um memorando foi adicionado. Deseja escrever um comentário?",
		ru:   "Заметка добавлена. Хотите написать комментарий?",
		tr:   "Not eklendi. Yorum yazmak ister misiniz?",
		ukUA: "Додано нотатки. Бажаєте залишити коментар?",
		uz:   "Eslatma qo&#39;shildi. Fikr yozmoqchimisiz?",
		zhCN: "已添加备忘录。您想发表评论吗？"},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		arEG: "تمت إضافة تعليق. هل ترغب بكتابة ملاحظة؟",
		de:   "Die Bemerkung wurde vermerkt. Möchtest du noch eine Notiz für dich hinzufügen?",
		en:   "Comment have been added. Do you want to write a note?",
		es:   "El comentario está añadido. ¿Quieres escribir una nota?",
		faIR: "شرح موردنظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
		fr:   "Des commentaires ont été ajoutés. Voulez-vous écrire un message ?",
		id:   "Komentar telah ditambahkan. Apakah Anda ingin menulis catatan?",
		it:   "Commento aggiunto. Vuoi scrivere un appunto?",
		jaJP: "コメントが追加されました。メモを残しますか？",
		koKR: "댓글이 추가되었습니다. 메모를 작성하시겠습니까?",
		pl:   "Komentarze zostały dodane. Czy chcesz napisać notatkę?",
		ptBR: "Comentários foram adicionados. Deseja escrever uma nota?",
		ptPT: "Comentários foram adicionados. Deseja escrever uma nota?",
		ru:   "Комментарий добавлен. Хотите написать заметку?",
		tr:   "Yorum eklendi. Not yazmak ister misiniz?",
		ukUA: "Коментарі додано. Бажаєте написати нотатку?",
		uz:   "Izoh qo&#39;shildi. Eslatma yozmoqchimisiz?",
		zhCN: "已添加评论。您想写一条注释吗？"},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		arEG: "بدون ملاحظات أو تعليقات",
		de:   "Keine Notizen oder Bemerkungen",
		en:   "Without notes or comments",
		es:   "Sin notas ni comentarios",
		faIR: "بدون یادداشت یا شرح",
		fr:   "Sans notes ni commentaires",
		id:   "Tanpa catatan atau komentar",
		it:   "Senza appunti o commenti",
		jaJP: "メモやコメントなし",
		koKR: "메모나 코멘트 없이",
		pl:   "Bez notatek i komentarzy",
		ptBR: "Sem notas ou comentários",
		ptPT: "Sem notas ou comentários",
		ru:   "Без заметок и комментариев",
		tr:   "Not veya yorum olmadan",
		ukUA: "Без приміток та коментарів",
		uz:   "Eslatmalar va izohlarsiz",
		zhCN: "没有注释或评论"},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		arEG: "لا توجد تعليقات",
		de:   "Keine Bemerkungen",
		en:   "No comments",
		es:   "Sin comentarios",
		faIR: "بدون شرح",
		fr:   "Sans commentaires",
		id:   "Tidak ada komentar",
		it:   "Nessun commento",
		jaJP: "コメントはありません",
		koKR: "댓글 없음",
		pl:   "Brak komentarzy",
		ptBR: "Sem comentários",
		ptPT: "Sem comentários",
		ru:   "Без комментариев",
		tr:   "Yorum yok",
		ukUA: "Без коментарів",
		uz:   "Izohlarsiz",
		zhCN: "暂无评论"},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		arEG: "بدون ملاحظات",
		de:   "Keine Notizen",
		en:   "Without notes",
		es:   "Sin notas",
		faIR: "بدون یادداشت",
		fr:   "Sans notes",
		id:   "Tanpa catatan",
		it:   "Senza appunti/note",
		jaJP: "メモなし",
		koKR: "노트 없이",
		pl:   "Bez notatek",
		ptBR: "Sem notas",
		ptPT: "Sem notas",
		ru:   "Без заметок",
		tr:   "Notlar olmadan",
		ukUA: "Без приміток",
		uz:   "Eslatmalarsiz",
		zhCN: "无注释"},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		arEG: "أضف ملاحظة (خاصة)",
		de:   "Notiz hinzufügen (privat)",
		en:   "Add a note (private)",
		es:   "Añadir una nota (privada)",
		faIR: "یک یادداشت اضافه کنید (خصوصی)",
		fr:   "Ajouter une note (privé)",
		id:   "Tambahkan catatan (pribadi)",
		it:   "Aggiungi una nota (privata)",
		jaJP: "メモを追加（非公開）",
		koKR: "메모 추가(비공개)",
		pl:   "Dodaj notatkę (prywatną)",
		ptBR: "Adicionar uma nota (privada)",
		ptPT: "Adicionar uma nota (privada)",
		ru:   "Добавить заметку",
		tr:   "Not ekle (özel)",
		ukUA: "Додати нотатку (приватну)",
		uz:   "Eslatma qo&#39;shing (shaxsiy)",
		zhCN: "添加注释（私人）"},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		arEG: "أضف تعليقًا (عامًا)",
		de:   "Bemerkung hinzufügen (öffentlich)",
		en:   "Add a comment (public)",
		es:   "Añadir un comentario (público)",
		faIR: "یک شرح اضافه کنید (عمومی)",
		fr:   "Ajouter un commentaire (public)",
		id:   "Tambahkan komentar (publik)",
		it:   "Aggiungi un commento (pubblico)",
		jaJP: "コメントを追加する（公開）",
		koKR: "댓글 추가 (공개)",
		pl:   "Dodaj komentarz (publiczny)",
		ptBR: "Adicionar um comentário (público)",
		ptPT: "Adicionar um comentário (público)",
		ru:   "Добавить комментарий",
		tr:   "Yorum ekle (genel)",
		ukUA: "Додати коментар (публічний)",
		uz:   "Fikr qo&#39;shing (ommaga ochiq)",
		zhCN: "添加评论（公开）"},
	DUE_IN_NOW: {
		arEG: "الآن",
		de:   "jetzt",
		en:   "now",
		es:   "ahora",
		faIR: "حالا",
		fr:   "maintenant",
		id:   "Sekarang",
		it:   "adesso",
		jaJP: "今",
		koKR: "지금",
		pl:   "Teraz",
		ptBR: "agora",
		ptPT: "agora",
		ru:   "прямо сейчас",
		tr:   "Şimdi",
		ukUA: "зараз",
		uz:   "hozir",
		zhCN: "现在"},
	DUE_IN_A_MINUTE: {
		arEG: "في دقيقة واحدة",
		de:   "in einer Minute",
		en:   "in a minute",
		es:   "en un minuto",
		faIR: "ظرف یک دقیقه",
		fr:   "dans une minute",
		id:   "dalam semenit",
		it:   "in un minuto",
		jaJP: "1分後",
		koKR: "1분 안에",
		pl:   "za minutę",
		ptBR: "em um minuto",
		ptPT: "num minuto",
		ru:   "через минуту",
		tr:   "bir dakika içinde",
		ukUA: "за хвилину",
		uz:   "bir daqiqada",
		zhCN: "一分钟后"},
	DUE_IN_X_MINUTES: {
		arEG: "في %v دقيقة",
		de:   "in %v Minuten",
		en:   "in %v minutes",
		es:   "en %v minutos",
		faIR: "در %v دقیقه",
		fr:   "en %v minutes",
		id:   "dalam %v menit",
		it:   "in %v minuti/o",
		jaJP: "%v分後",
		koKR: "%v분 안에",
		pl:   "w %v minut",
		ptBR: "em %v minutos",
		ptPT: "em %v minutos",
		ru:   "через %v минут(ы)",
		tr:   "%v dakika içinde",
		ukUA: "у %v хвилинах",
		uz:   "%v daqiqada",
		zhCN: "%v 分钟后"},
	DUE_IN_AN_HOUR: {
		arEG: "في ساعة واحدة",
		de:   "in einer Stunde",
		en:   "in an hour",
		es:   "en  una hora",
		faIR: "ظرف یک ساعت",
		fr:   "dans une heure",
		id:   "sejam lagi",
		it:   "in un ora",
		jaJP: "1時間後",
		koKR: "한 시간 안에",
		pl:   "za godzinę",
		ptBR: "em uma hora",
		ptPT: "em uma hora",
		ru:   "через час",
		tr:   "bir saat içinde",
		ukUA: "за годину",
		uz:   "bir soat ichida",
		zhCN: "一小时内"},
	DUE_IN_X_HOURS: {
		arEG: "في %v ساعات",
		de:   "in %v Stunde",
		en:   "in %v hours",
		es:   "en %v horas",
		faIR: "در %v ساعت",
		fr:   "en %v heures",
		id:   "dalam %v jam",
		it:   "in %v ore/a",
		jaJP: "%v時間以内",
		koKR: "%v시간 안에",
		pl:   "w %v godzinach",
		ptBR: "em %v horas",
		ptPT: "em %v horas",
		ru:   "через %v часа/часов",
		tr:   "%v saat içinde",
		ukUA: "у %v годинах",
		uz:   "%v soat ichida",
		zhCN: "%v 小时后"},
	DUE_IN_X_DAYS: {
		arEG: "في %v أيام",
		de:   "in %v Tagen",
		en:   "in %v days",
		es:   "en %v días",
		faIR: "در %v روز",
		fr:   "dans %v jours",
		id:   "dalam %v hari",
		it:   "in %v giorni/o",
		jaJP: "%v日後",
		koKR: "%v일 후",
		pl:   "za %v dni",
		ptBR: "em %v dias",
		ptPT: "em %v dias",
		ru:   "через %v дня/дней",
		tr:   "%v gün içinde",
		ukUA: "за %v днів",
		uz:   "%v kun ichida",
		zhCN: "%v 天后"},
	WS_ALEX_T: {
		arEG: "ألكسندر تراخيمينوك",
		de:   "Alexander Trakhimenok",
		en:   "Alexander Trakhimenok",
		es:   "Alexander Trakhimenok",
		faIR: "الکساندر تراخیمِنوک",
		fr:   "Alexandre Trakhimenok",
		id:   "Alexander Trakhimenok",
		it:   "Alexander Trakhimenok",
		jaJP: "アレクサンダー・トラヒメノク",
		koKR: "알렉산더 트라키메녹",
		pl:   "Aleksander Trachimenok",
		ptBR: "Alexandre Trakhimenok",
		ptPT: "Alexandre Trakhimenok",
		ru:   "Александр Трахимёнок",
		tr:   "Aleksandr Trakhimenok",
		ukUA: "Олександр Трахіменок",
		uz:   "Aleksandr Traximenok",
		zhCN: "亚历山大·特拉希梅诺克"},

	WS_INDEX_TITLE: {
		arEG: "DebtsTracker.io - تطبيق IOU لتتبع ديونك وأصولك الشخصية",
		de:   "DebtsTracker.io - eine App, um Ihre persönlichen Schulden zu verfolgen",
		en:   "DebtsTracker.io - an IOU app to track your personal debts & assets",
		es:   "DebtsTracker.io es una aplicación para el control de sus deudas personales",
		faIR: "DebtsTracker.io - برنامه ای برای ردیابی بدهی ها و دارایی های شما",
		fr:   "DebtsTracker.io - une application pour suivre vos dettes personnelles",
		id:   "DebtsTracker.io - aplikasi IOU untuk melacak utang &amp; aset pribadi Anda",
		it:   "DebtsTracker.io - un app per monitorare i tuoi debiti personali",
		jaJP: "DebtsTracker.io は - アプリはあなたの個人的な借金を追跡します",
		koKR: "DebtsTracker.io 은 - 앱이 사용자의 개인 채무를 추적",
		pl:   "DebtsTracker.io - aplikacja do śledzenia osobistych długów",
		ptBR: "DebtsTracker.io - um aplicativo de IOU para rastrear suas dívidas e ativos pessoais",
		ptPT: "DebtsTracker.io - um aplicativo para controlar suas dívidas pessoais",
		ru:   "DebtsTracker.io - программа для учёта личных долгов и активов",
		tr:   "DebtsTracker.io - Kişisel borçlarınızı ve varlıklarınızı takip etmek için bir IOU uygulaması",
		ukUA: "DebtsTracker.io — додаток для відстеження боргів та активів",
		uz:   "DebtsTracker.io - shaxsiy qarzlaringiz va aktivlaringizni kuzatish uchun IOU ilovasi",
		zhCN: "DebtsTracker.io - 一个应用程序来跟踪你的个人债务",
	},
	WS_LIVE_DEMO: {
		arEG: "مظاهرة",
		de:   "Live-Demo",
		en:   "Demostración",
		es:   "Demo en vivo",
		faIR: "نسخه ی نمایشی زنده",
		fr:   "Démo en direct",
		id:   "Demonstrasi",
		it:   "Demo online",
		jaJP: "ライブデモ",
		koKR: "실시간 데모",
		pl:   "Demo na żywo",
		ptBR: "Demonstração",
		ptPT: "Demonstração ao vivo",
		ru:   "Демо версия online",
		tr:   "Gösterim",
		ukUA: "Демонстрація",
		uz:   "Namoyish",
		zhCN: "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		arEG: "بوت دردشة لتطبيق تيليجرام",
		de:   "Chat-Bot für Telegram",
		en:   "Chat bot for Telegram messenger",
		es:   "Chat bot para Telegram",
		faIR: "ربات چت برای پیام رسان تلگرام",
		fr:   "bot Chat for Telegram messenger",
		id:   "Bot obrolan untuk messenger Telegram",
		it:   "Bot Chat per Telegram",
		jaJP: "電報メッセンジャーのためのチャットボット",
		koKR: "전보 메신저 채팅 봇",
		pl:   "Chat bot do telegramu posłańca",
		ptBR: "Bot de bate-papo para o mensageiro Telegram",
		ptPT: "bot de bate-papo para Telegram messenger",
		ru:   "Бот для Telegram",
		tr:   "Telegram messenger için sohbet botu",
		ukUA: "Чат-бот для месенджера Telegram",
		uz:   "Telegram messenjeri uchun chat-bot",
		zhCN: "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		arEG: "افتح في تيليجرام 🚀",
		de:   "Öffnen in Telegram &#x1F680;",
		en:   "Open in Telegram &#x1F680;",
		es:   "Abrir en Telegram &#x1F680;",
		faIR: "بازکردن در تلگرام &#x1F680;",
		fr:   "Ouvrir dans Telegram 🚀",
		id:   "Buka di Telegram 🚀",
		it:   "Apri su Telegram &#x1F680;",
		jaJP: "電報 &#x1F680; で開きます。",
		koKR: "전보 &#x1F680; 에서 열기;",
		pl:   "Otwórz w telegramu &#x1F680;",
		ptBR: "Abrir no Telegram 🚀",
		ptPT: "Abrir no Telegram 🚀",
		ru:   "Открыть в Телеграмме &#x1F680;",
		tr:   "Telegram&#39;da aç 🚀",
		ukUA: "Відкрити в Телеграмі 🚀",
		uz:   "Telegramda oching 🚀",
		zhCN: "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		arEG: "في الوقت الحالي برنامجنا متاح فقط على <a href='https://telegram.org/'>تطبيق Telegram messenger</a>",
		de:   "Im Moment ist unser Programm nur auf <a href='https://telegram.org/'>Telegram verfügbar</a>",
		en:   "At the moment our program is available just on <a href='https://telegram.org/'>Telegram messenger</a>",
		es:   "De momento nuestro programa está disponible sólo en <a href='https://telegram.org/'>Telegrama mensajero </a>",
		faIR: "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>Телеграм</a>تلگرام",
		fr:   "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'>Telegram messager</a>",
		id:   "Saat ini program kami hanya tersedia di <a href='https://telegram.org/'>Telegram messenger</a>",
		it:   "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'>Telegram</a>",
		jaJP: "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'>Telegram</a>電報のメッセンジャーで提供されています",
		koKR: "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>Telegram</a> 의 <b> 전보 </b>을 메신저 를 볼 수 있습니다",
		pl:   "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'>Telegram messenger</a>",
		ptBR: "No momento nosso programa está disponível apenas no <a href='https://telegram.org/'>mensageiro Telegram</a>",
		ptPT: "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'>Telegram messenger</a>",
		ru:   "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		tr:   "Şu anda programımız sadece <a href='https://telegram.org/'>Telegram messenger&#39;da</a> mevcuttur",
		ukUA: "Наразі наша програма доступна лише в <a href='https://telegram.org/'>месенджері Telegram.</a>",
		uz:   "Hozirda dasturimiz faqat <a href='https://telegram.org/'>Telegram messenjerida</a> mavjud",
		zhCN: "目前我们的计划是只提供在<a href='https://telegram.org/'>Telegram</a>电报的使者",
	},
	WS_MOTTO: {
		arEG: "اعرف الحد الأدنى الخاص بك ولا تفوت أبدًا سداد الديون!",
		de:   "Wissen, wem man wie viel schuldet!",
		en:   "Know your bottom line & never miss a debt payment!",
		es:   "¡Controle sus pagos y deudas!",
		faIR: "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
		fr:   "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		id:   "Ketahui batas akhir Anda dan jangan pernah melewatkan pembayaran utang!",
		it:   "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		jaJP: "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		koKR: "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		pl:   "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		ptBR: "Conheça seus limites e nunca deixe de pagar uma dívida!",
		ptPT: "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		ru:   "Платежи по долгам целиком и вовремя!",
		tr:   "Alt sınırınızı bilin ve hiçbir borç ödemesini kaçırmayın!",
		ukUA: "Знайте свій прибуток і ніколи не пропускайте виплати боргів!",
		uz:   "Xulosangizni biling va hech qachon qarzni to&#39;lashni o&#39;tkazib yubormang!",
		zhCN: "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		arEG: "DebtsTracker.io هو تطبيق جوال لسندات الدين وخدمة تذكير تساعدك على تتبع ديونك وائتماناتك وأصولك. يرسل رسائل تذكير آلية عبر البريد الإلكتروني والرسائل النصية القصيرة إلى مدينيك.",
		de:   "DebtsTracker.io ist eine mobile App, die beim Verwalten von persönlichen Schulden hilft - egal ob Sie Geld verleihen oder welches leihen. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner und Gläubiger.",
		en:   "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		es:   "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorios que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		faIR: "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		fr:   "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		id:   "DebtsTracker.io adalah aplikasi IOU seluler &amp; layanan pengingat yang membantu melacak utang, kredit &amp; aset Anda. Mengirimkan pengingat email &amp; SMS otomatis kepada debitur Anda.",
		it:   "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		jaJP: "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		koKR: "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		pl:   "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		ptBR: "DebtsTracker.io é um aplicativo de IOU móvel e um serviço de lembretes que ajuda a rastrear suas dívidas, créditos e ativos. Envia lembretes automáticos por e-mail e SMS para seus devedores.",
		ptPT: "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		ru:   "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		tr:   "DebtsTracker.io, borçlarınızı, kredilerinizi ve varlıklarınızı takip etmenize yardımcı olan bir mobil IOU uygulaması ve hatırlatma hizmetidir. Borçlularınıza otomatik e-posta ve SMS hatırlatıcıları gönderir.",
		ukUA: "DebtsTracker.io — це мобільний додаток для відстеження боргів та сервіс нагадувань, який допомагає відстежувати ваші борги, кредити та активи. Надсилає автоматичні нагадування електронною поштою та SMS вашим боржникам.",
		uz:   "DebtsTracker.io - bu sizning qarzlaringiz, kreditlaringiz va aktivlaringizni kuzatishga yordam beruvchi mobil IOU ilovasi va eslatma xizmati. Qarzdorlaringizga avtomatik elektron pochta va SMS eslatmalarini yuboradi.",
		zhCN: "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		arEG: "كيف يمكنك المساعدة في مشروع DebtsTracker.io",
		de:   "Wie kann man beim DebtsTracker.io Projekt helfen kann",
		en:   "How you can help to DebtsTracker.io project",
		es:   "Como puedes ayudar a DebtsTracker.io project",
		faIR: "چگونه می توانید به پروژه  DebtsTracker.io کمک کنید.",
		fr:   "Comment vous pouvez contribuer au projet DebtsTracker.io",
		id:   "Bagaimana Anda dapat membantu proyek DebtsTracker.io",
		it:   "Come potete aiutare il progetto DebtsTracker.io",
		jaJP: "DebtsTracker.ioプロジェクトへの協力方法",
		koKR: "DebtsTracker.io 프로젝트에 도움을 줄 수 있는 방법",
		pl:   "Jak możesz pomóc projektowi DebtsTracker.io",
		ptBR: "Como você pode ajudar o projeto DebtsTracker.io",
		ptPT: "Como pode ajudar o projeto DebtsTracker.io",
		ru:   "Как вы можете помочь проекту DebtsTracker.io",
		tr:   "DebtsTracker.io projesine nasıl yardımcı olabilirsiniz?",
		ukUA: "Як ви можете допомогти проєкту DebtsTracker.io",
		uz:   "DebtsTracker.io loyihasiga qanday yordam berishingiz mumkin",
		zhCN: "如何帮助 DebtsTracker.io 项目"},
	WS_ADS_TITLE: {
		arEG: "إعلانات @ DebtsTracker.IO",
		de:   "Werbung @ DebtsTracker.IO",
		en:   "Ads @ DebtsTracker.IO",
		es:   "Anuncio @ DebtsTracker.IO",
		faIR: "تبلیغات @ DebtsTracker.IO",
		fr:   "Annonces sur DebtsTracker.IO",
		id:   "Iklan @ DebtsTracker.IO",
		it:   "Pubblicita' @ DebtsTracker.IO",
		jaJP: "広告 @ DebtsTracker.IO",
		koKR: "DebtsTracker.IO 광고",
		pl:   "Reklamy @ DebtsTracker.IO",
		ptBR: "Anúncios @ DebtsTracker.IO",
		ptPT: "Anúncios @ DebtsTracker.IO",
		ru:   "Реклама на DebtsTracker.IO",
		tr:   "Reklamlar @ DebtsTracker.IO",
		ukUA: "Реклама @ DebtsTracker.IO",
		uz:   "Reklamalar @ DebtsTracker.IO",
		zhCN: "广告@DebtsTracker.IO"},
	WS_ADS_CONTENT: {
		arEG: `لإضافة إعلانات في تطبيقنا، يرجى إرسال بريد إلكتروني إلينا على <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		de:   `Um Werbung in unserer App zu schalten, schick uns eine Mail an <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		en:   `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		es:   `Para publicar un anuncio en nuestra app escríbenos un e-mail a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		faIR: `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		fr:   `Pour placer des annonces dans notre application, veuillez nous envoyer un e-mail à <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		id:   `Untuk memasang iklan di aplikasi kami, silakan kirim email ke <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		it:   `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		jaJP: `アプリに広告を掲載するには、 <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>までメールをお送りください。`,
		koKR: `앱에 광고를 게재하려면 <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> 로 이메일을 보내주세요.`,
		pl:   `Aby umieścić reklamy w naszej aplikacji, wyślij wiadomość e-mail na <a href="mailto:ads@debtstracker.io">adres ads@debtstracker.io</a> .`,
		ptBR: `Para colocar anúncios em nosso aplicativo, envie um e-mail para <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		ptPT: `Para colocar anúncios na nossa aplicação, envie um e-mail para <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		ru:   `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		tr:   `Uygulamamıza reklam yerleştirmek için lütfen <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> adresine bir e-posta gönderin.`,
		ukUA: `Щоб розмістити рекламу в нашому додатку, надішліть нам електронного листа на адресу <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> .`,
		uz:   `Ilovamizga reklama joylashtirish uchun <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> manziliga elektron pochta xabarini yuboring.`,
		zhCN: `要在我们的应用中投放广告，请发送电子邮件至<a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a> 。`},
	FB_MAKE_RECORD_HEADER: {
		arEG: "سجل الديون",
		de:   "Schuldschein",
		en:   "Record debt",
		es:   "Registrar la deuda",
		faIR: "ثبت بدهی",
		fr:   "Une dette record",
		id:   "Rekor utang",
		it:   "Registra il debito",
		jaJP: "記録的な負債",
		koKR: "기록적인 부채",
		pl:   "Rekordowy dług",
		ptBR: "Dívida recorde",
		ptPT: "Dívida recorde",
		ru:   "Записать долг",
		tr:   "Rekor borç",
		ukUA: "Рекордний борг",
		uz:   "Rekord qarz",
		zhCN: "创纪录的债务"},
	FB_MAKE_RECORD_HEADLINE: {
		arEG: "قم بالتمرير إلى اليسار لرؤية الخيارات الأخرى.",
		de:   "Scroll nach links, um weitere Optionen zu sehen.",
		en:   "Scroll left to see other options.",
		es:   "Desliza a la izquierda para ver otras opciones",
		faIR: "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
		fr:   "Faites défiler vers la gauche pour voir d’autres options.",
		id:   "Gulir ke kiri untuk melihat pilihan lainnya.",
		it:   "Scrolla a sinistra per vedere altre opzioni",
		jaJP: "他のオプションを表示するには左にスクロールしてください。",
		koKR: "다른 옵션을 보려면 왼쪽으로 스크롤하세요.",
		pl:   "Przewiń w lewo, aby zobaczyć inne opcje.",
		ptBR: "Role para a esquerda para ver outras opções.",
		ptPT: "Percorra para a esquerda para ver outras opções.",
		ru:   "Пролистайте карточки влево чтобы увидеть другие опции.",
		tr:   "Diğer seçenekleri görmek için sola kaydırın.",
		ukUA: "Прокрутіть ліворуч, щоб побачити інші опції.",
		uz:   "Boshqa variantlarni ko&#39;rish uchun chapga aylantiring.",
		zhCN: "向左滚动即可查看其他选项。"},
	FB_HOW_ARE_THINGS_HEADER: {
		arEG: "كيف حالك؟",
		de:   "Wie geht es dir?",
		en:   "How are you doing?",
		es:   "¿Cómo va todo?",
		faIR: "حال شما چطوره؟",
		fr:   "Comment allez-vous?",
		id:   "Apa kabarmu?",
		it:   "Come te la passi?",
		jaJP: "お元気ですか？",
		koKR: "어떻게 지내세요?",
		pl:   "Jak się masz?",
		ptBR: "Como vai?",
		ptPT: "Como estás?",
		ru:   "Как идут дела?",
		tr:   "Nasılsın?",
		ukUA: "Як справи?",
		uz:   "Ahvoling yaxshimi?",
		zhCN: "你好吗？"},
	SNEATBOT_MSG_TXT_START: {
		arEG: `<b>من مُنشئ البوت</b> : مرحبًا %s!

 @SneatBot يُساعدك على إدارة حياتك العائلية اليومية. أو يُمكنك إنشاء مساحة لإدارة مجموعتك/فريقك/مجتمعك.

لقد بذلتُ جهدًا كبيرًا لجعل هذا البوت مفيدًا وسريعًا وموثوقًا. آمل أن يُعجبك.

يمكنك التعرّف على الميزات الجديدة للبوت في قناة @SneatApp حيث <a href="https://t.me/StarGiveaways_EN">نُقدّم 500 🌟 شهريًا</a> .
`,
		de: `<b>Vom Ersteller des Bots</b>: Hallo %s!

@SneatBot hilft Ihnen, Ihren Familienalltag zu organisieren. Oder Sie können einen Bereich erstellen, in dem Sie Ihre Gruppe/Ihr Team/Ihre Community verwalten können.

Ich habe viel Zeit darauf verwendet, diesen Bot nützlich, schnell und zuverlässig zu machen. Ich hoffe, er gefällt Ihnen.

Sie können sich über neue Funktionen des Bots im @SneatApp-Kanal informieren, wo wir <a href="https://t.me/StarGiveaways_EN">JEDEN Monat 500 🌟 verschenken</a>.`,
		en: `
<b>From bot's creator</b>: Hi %s!

@SneatBot helps to manage your day-to-day family life. Or you can create a space to manage your group/team/community.

I've spend lot's of time to make this bot useful, quick & reliable.I hope you'll like it.

You can learn about new features of the bot in @SneatApp channel where <a href="https://t.me/StarGiveaways_EN">we giveaway 500 🌟 EVERY month</a>.
`,
		es: `<b>Del creador del bot:</b> ¡Hola, %s!

@SneatBot te ayuda a gestionar tu vida familiar diaria. O puedes crear un espacio para gestionar tu grupo/equipo/comunidad.

He dedicado mucho tiempo a hacer que este bot sea útil, rápido y fiable. Espero que te guste.

Puedes conocer las nuevas funciones del bot en el canal de @SneatApp, donde <a href="https://t.me/StarGiveaways_EN">regalamos 500 🌟 CADA mes</a>.`,
		faIR: //nolint:staticcheck // disable ST1018 for this line
		`
 <b>از طرف سازنده ربات</b> : سلام %s!

@SneatBot به شما کمک می‌کند تا زندگی روزمره خانوادگی خود را مدیریت کنید. یا می‌توانید فضایی برای مدیریت گروه/تیم/انجمن خود ایجاد کنید.

من زمان زیادی را صرف کرده‌ام تا این ربات را مفید، سریع و قابل اعتماد کنم. امیدوارم از آن خوشتان بیاید.

می‌توانید در کانال @SneatApp با ویژگی‌های جدید ربات آشنا شوید، جایی که <a href="https://t.me/StarGiveaways_EN">ما هر ماه ۵۰۰ 🌟 جایزه می‌دهیم</a> .
`,
		fr: `<b>De la part du créateur du bot</b> : Bonjour %s !

@SneatBot vous aide à gérer votre vie de famille au quotidien. Vous pouvez également créer un espace pour gérer votre groupe/équipe/communauté.

J'ai passé beaucoup de temps à rendre ce bot utile, rapide et fiable. J'espère qu'il vous plaira.

Vous pouvez en apprendre davantage sur les nouvelles fonctionnalités du bot sur le canal @SneatApp où <a href="https://t.me/StarGiveaways_FR">nous offrons 500 🌟 CHAQUE mois</a>.`,
		id: `
 <b>Dari pembuat bot</b> : Hai %s!

@SneatBot membantu mengelola kehidupan keluarga Anda sehari-hari. Atau Anda dapat membuat ruang untuk mengelola grup/tim/komunitas Anda.

Saya telah menghabiskan banyak waktu untuk membuat bot ini berguna, cepat &amp; andal. Saya harap Anda menyukainya.

Anda dapat mempelajari tentang fitur-fitur baru bot di saluran @SneatApp tempat <a href="https://t.me/StarGiveaways_EN">kami memberikan hadiah 500 🌟 SETIAP bulan</a> .
`,
		it: `<b>Dal creatore del bot</b>: Ciao %s!

@SneatBot ti aiuta a gestire la tua vita familiare quotidiana. Oppure puoi creare uno spazio per gestire il tuo gruppo/team/community.

Ho dedicato molto tempo a rendere questo bot utile, veloce e affidabile. Spero che ti piaccia.

Puoi scoprire le nuove funzionalità del bot nel canale @SneatApp dove <a href="https://t.me/StarGiveaways_EN">regaliamo 500 🌟 OGNI mese</a>.`,
		jaJP: `<b>ボットの作成者より</b>: こんにちは、%s!

@SneatBot は、日々の家族生活の管理に役立ちます。または、グループ/チーム/コミュニティを管理するスペースを作成することもできます。

このボットを便利で、迅速で、信頼できるものにするために、多くの時間を費やしました。気に入っていただければ幸いです。

ボットの新機能については、@SneatApp チャンネルで確認できます。ここでは、<a href="https://t.me/StarGiveaways_EN">毎月 500 🌟 をプレゼント</a>しています。`,
		koKR: `<b>봇 제작자로부터</b>: 안녕하세요 %s!

@SneatBot은 일상의 가족 생활을 관리하는 데 도움이 됩니다. 또는 그룹/팀/커뮤니티를 관리할 공간을 만들 수 있습니다.

저는 이 봇을 유용하고 빠르고 안정적으로 만들기 위해 많은 시간을 투자했습니다. 마음에 들어 하시기를 바랍니다.

@SneatApp 채널에서 봇의 새로운 기능에 대해 알아볼 수 있습니다. <a href="https://t.me/StarGiveaways_EN">매월 500🌟를 경품으로 드립니다</a>.`,
		pl: `<b>Od twórcy bota</b>: Cześć %s!

@SneatBot pomaga zarządzać codziennym życiem rodzinnym. Możesz też utworzyć przestrzeń do zarządzania swoją grupą/zespołem/społecznością.

Poświęciłem dużo czasu, aby ten bot był użyteczny, szybki i niezawodny. Mam nadzieję, że Ci się spodoba.

Możesz dowiedzieć się o nowych funkcjach bota na kanale @SneatApp, gdzie <a href="https://t.me/StarGiveaways_EN">rozdajemy 500 🌟 CO MIESIĄC</a>.`,
		ptBR: `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerenciar sua vida familiar cotidiana. Ou você pode criar um espaço para gerenciar seu grupo/equipe/comunidade.

Eu gastei muito tempo para tornar este bot útil, rápido e confiável. Espero que você goste.

Você pode aprender sobre os novos recursos do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">nós doamos 500 🌟 TODO mês</a>.`,
		ptPT: `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerir a sua vida familiar quotidiana. Ou pode criar um espaço para gerir o seu grupo/equipa/comunidade.

Gastei muito tempo para tornar este bot útil, rápido e fiável.

Pode conhecer as novas funcionalidades do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">distribuímos 500 🌟 TODOS os meses</a>.`,
		ru: `
<b>От создателя бота:</b> Привет %s!

@SneatBot помогает организовать вашу семейную жизнь.
Так же можно создать пространство для управления группой/командой/сообществом.

Вы можете узнавать о новых возможностях бота в канале @SneatApp_ru где <a href="https://t.me/StarGiveaways_RU">мы разыгрываем 500 🌟 КАЖДЫЙ месяц</a>.
`,
		tr: `
 <b>Botun yaratıcısından</b> : Merhaba %s!

@SneatBot günlük aile hayatınızı yönetmenize yardımcı olur. Ya da grubunuzu/ekibinizi/topluluğunuzu yönetmek için bir alan oluşturabilirsiniz.

Bu botu kullanışlı, hızlı ve güvenilir hale getirmek için çok zaman harcadım. Umarım beğenirsiniz.

 <a href="https://t.me/StarGiveaways_EN">Her ay 500 🌟 hediye ettiğimiz</a> @SneatApp kanalında botun yeni özelliklerini öğrenebilirsiniz.
`,
		ukUA: `<b>Від творця бота</b>: Привіт, %s!

@SneatBot допомагає керувати повсякденним сімейним життям. Або ви можете створити простір для керування своєю групою/командою/спільнотою.

Я витратив багато часу, щоб зробити цього бота корисним, швидким і надійним. Сподіваюся, він вам сподобається.

Ви можете дізнатися про нові функції бота на каналі @SneatApp, де <a href="https://t.me/StarGiveaways_EN">ми роздаємо 500 🌟 КОЖНОГО місяця</a>.`,
		uz: `
 <b>Bot yaratuvchisidan</b> : Salom %s!

@SneatBot kundalik oilaviy hayotingizni boshqarishga yordam beradi. Yoki guruhingizni/jamoangizni/hamjamiyatingizni boshqarish uchun joy yaratishingiz mumkin.

Men bu botni foydali, tez va ishonchli qilish uchun ko‘p vaqt sarfladim. U sizga yoqadi degan umiddaman.

 <a href="https://t.me/StarGiveaways_EN">Har oyda 500 🌟 sovg‘a qiladigan</a> @SneatApp kanalida botning yangi funksiyalari haqida bilib olishingiz mumkin.
`,
		zhCN: `<b>来自机器人的创建者</b>：嗨 %s！

@SneatBot 帮助您管理日常家庭生活。或者您可以创建一个空间来管理您的群组/团队/社区。

我花了很多时间让这个机器人变得有用、快速和可靠。希望你会喜欢它。

您可以在 @SneatApp 频道了解该机器人的新功能，<a href="https://t.me/StarGiveaways_EN">我们每月赠送 500 🌟</a>。`,
	},
	SpaceCmdText: {
		arEG: "المساحة الحالية: %s <b>%s</b>",
		de:   "Aktueller Raum: %s <b>%s</b>",
		en:   "Current space: %s <b>%s</b>",
		es:   "Espacio actual: %s <b>%s</b>",
		faIR: "فضای فعلی: %s <b>%s</b>",
		fr:   "Espace actuel: %s <b>%s</b>",
		id:   "Ruang saat ini: %s <b>%s</b>",
		it:   "Spazio corrente: %s <b>%s</b>",
		jaJP: "現在のスペース: %s <b>%s</b>",
		koKR: "현재 공간: %s <b>%s</b>",
		pl:   "Aktualna przestrzeń: %s <b>%s</b>",
		ptBR: "Espaço atual: %s <b>%s</b>",
		ptPT: "Espaço atual: %s <b>%s</b>",
		ru:   "Текущее пространство: %s <b>%s</b>",
		tr:   "Mevcut alan: %s <b>%s</b>",
		ukUA: "Поточний простір: %s <b>%s</b>",
		uz:   "Joriy joy: %s <b>%s</b>",
		zhCN: "当前空间: %s <b>%s</b>",
	},
	SpaceCmdBtnContacts: {
		arEG: "جهات الاتصال",
		de:   "Kontakte",
		en:   "Contacts",
		es:   "Contactos",
		faIR: "مخاطبین",
		fr:   "Contacts",
		id:   "Kontak",
		it:   "Contatti",
		jaJP: "連絡先",
		koKR: "연락처",
		pl:   "Kontakty",
		ptBR: "Contatos",
		ptPT: "Contatos",
		ru:   "Контакты",
		tr:   "Kişiler",
		ukUA: "Контакти",
		uz:   "Kontaktlar",
		zhCN: "联系人",
	},
	SpaceCmdBtnMembers: {
		arEG: "أعضاء",
		de:   "Mitglieder",
		en:   "Members",
		es:   "Miembros",
		faIR: "اعضا",
		fr:   "Membres",
		id:   "Anggota",
		it:   "Membri",
		jaJP: "メンバー",
		koKR: "회원",
		pl:   "Członkowie",
		ptBR: "Membros",
		ptPT: "Membros",
		ru:   "Участники",
		tr:   "Üyeler",
		ukUA: "Учасники",
		uz:   "A&#39;zolar",
		zhCN: "成员",
	},
	FamilyMembers: {
		arEG: "أفراد العائلة",
		de:   "Familienmitglieder",
		en:   "Family members",
		es:   "Miembros de la familia",
		faIR: "اعضای خانواده",
		fr:   "Membres de la famille",
		id:   "Anggota keluarga",
		it:   "Membri della famiglia",
		jaJP: "家族のメンバー",
		koKR: "가족 구성원",
		pl:   "Członkowie rodziny",
		ptBR: "Membros da família",
		ptPT: "Membros da família",
		ru:   "Члены семьи",
		tr:   "Aile üyeleri",
		ukUA: "Члени сім'ї",
		uz:   "Oila a&#39;zolari",
		zhCN: "家庭成员",
	},
	SpaceMembers: {
		arEG: "أعضاء الفضاء",
		de:   "Raummitglieder",
		en:   "Space members",
		es:   "Miembros del espacio",
		faIR: "اعضای فضا",
		fr:   "Membres de l'espace",
		id:   "Anggota ruang",
		it:   "Membri dello spazio",
		jaJP: "スペースメンバー",
		koKR: "공간 멤버",
		pl:   "Członkowie przestrzeni",
		ptBR: "Membros do espaço",
		ptPT: "Membros do espaço",
		ru:   "Участники пространства",
		tr:   "Alan üyeleri",
		ukUA: "Члени простору",
		uz:   "Kosmik a&#39;zolar",
		zhCN: "空间成员",
	},
	SpaceCmdBtnLists: {
		arEG: "القوائم",
		de:   "Listen",
		en:   "Lists",
		es:   "Listas",
		faIR: "لیست ها",
		fr:   "Listes",
		id:   "Daftar",
		it:   "Elenchi",
		jaJP: "リスト",
		koKR: "목록",
		pl:   "Listy",
		ptBR: "Listas",
		ptPT: "Listas",
		ru:   "Списки",
		tr:   "Listeler",
		ukUA: "Списки",
		uz:   "Ro&#39;yxatlar",
		zhCN: "清单",
	},
	SpaceCmdBtnAssets: {
		arEG: "أصول",
		de:   "Vermögenswerte",
		en:   "Assets",
		es:   "Activos",
		faIR: "دارایی ها",
		fr:   "Actifs",
		id:   "Aset",
		it:   "Attività",
		jaJP: "資産",
		koKR: "자산",
		pl:   "Aktywa",
		ptBR: "Ativos",
		ptPT: "Ativos",
		ru:   "Активы",
		tr:   "Varlıklar",
		ukUA: "Активи",
		uz:   "Aktivlar",
		zhCN: "资产",
	},
	SpaceCmdBtnBudget: {
		arEG: "ميزانية",
		de:   "Budget",
		en:   "Budget",
		es:   "Presupuesto",
		faIR: "بودجه",
		fr:   "Budget",
		id:   "Anggaran",
		it:   "Bilancio",
		jaJP: "予算",
		koKR: "예산",
		pl:   "Budżet",
		ptBR: "Orçamento",
		ptPT: "Orçamento",
		ru:   "Бюджет",
		tr:   "Bütçe",
		ukUA: "Бюджет",
		uz:   "Byudjet",
		zhCN: "预算",
	},
	SpaceCmdBtnDebts: {
		arEG: "الديون",
		de:   "Schulden",
		en:   "Debts",
		es:   "Deudas",
		faIR: "بدهی ها",
		fr:   "Dettes",
		id:   "Hutang",
		it:   "Debiti",
		jaJP: "借金",
		koKR: "부채",
		pl:   "Długi",
		ptBR: "Dívidas",
		ptPT: "Dívidas",
		ru:   "Долги",
		tr:   "Borçlar",
		ukUA: "Борги",
		uz:   "Qarzlar",
		zhCN: "债务",
	},
	SpaceCmdBtnCalendar: {
		arEG: "تقويم",
		de:   "Kalender",
		en:   "Calendar",
		es:   "Calendario",
		faIR: "تقویم",
		fr:   "Calendrier",
		id:   "Kalender",
		it:   "Calendario",
		jaJP: "カレンダー",
		koKR: "달력",
		pl:   "Kalendarz",
		ptBR: "Calendário",
		ptPT: "Calendário",
		ru:   "Календарь",
		tr:   "Takvim",
		ukUA: "Календар",
		uz:   "Kalendar",
		zhCN: "日历",
	},
	SpaceCmdBtnTrackers: {
		arEG: "المتتبعون",
		de:   "Tracker",
		en:   "Trackers",
		es:   "Rastreadores",
		faIR: "ردیاب\u200cها",
		fr:   "Traqueurs",
		id:   "Pelacak",
		it:   "Tracker",
		jaJP: "トラッカー",
		koKR: "추적기",
		pl:   "Monitorujące",
		ptBR: "Rastreadores",
		ptPT: "Rastreadores",
		ru:   "Трекеры",
		tr:   "İzleyiciler",
		ukUA: "Трекери",
		uz:   "Kuzatuvchilar",
		zhCN: "追踪器",
	},
	BtnSpaces: {
		arEG: "المساحات",
		de:   "Räume",
		en:   "Spaces",
		es:   "Espacios",
		faIR: "فضاها",
		fr:   "Espaces",
		id:   "Ruang",
		it:   "Spazi",
		jaJP: "スペース",
		koKR: "공간",
		pl:   "Przestrzenie",
		ptBR: "Espaços",
		ptPT: "Espaços",
		ru:   "Пространства",
		tr:   "Mekanlar",
		ukUA: "Простори",
		uz:   "Bo&#39;shliqlar",
		zhCN: "空间",
	},
	SpaceCmdBtnSettings: {
		arEG: "إعدادات",
		de:   "Einstellungen",
		en:   "Settings",
		es:   "Ajustes",
		faIR: "تنظیمات",
		fr:   "Paramètres",
		id:   "Pengaturan",
		it:   "Impostazioni",
		jaJP: "設定",
		koKR: "설정",
		pl:   "Ustawienia",
		ptBR: "Configurações",
		ptPT: "Configurações",
		ru:   "Настройки",
		tr:   "Ayarlar",
		ukUA: "Налаштування",
		uz:   "Sozlamalar",
		zhCN: "设置",
	},
	LIST_CMD_BUY: {
		arEG: "يشتري",
		de:   "Kaufen",
		en:   "Buy",
		es:   "Comprar",
		faIR: "خرید",
		fr:   "Acheter",
		id:   "Beli",
		it:   "Acquista",
		jaJP: "購入",
		koKR: "사다",
		pl:   "Kup",
		ptBR: "Comprar",
		ptPT: "Comprar",
		ru:   "Купить",
		tr:   "Satın al",
		ukUA: "Купити",
		uz:   "Sotib olish",
		zhCN: "购买",
	},
	LIST_CMD_TODO: {
		arEG: "المهام",
		de:   "Aufgaben",
		en:   "ToDo",
		es:   "Tareas",
		faIR: "وظایف",
		fr:   "Tâches",
		id:   "Tugas",
		it:   "Compiti",
		jaJP: "タスク",
		koKR: "할 일",
		pl:   "Zadania",
		ptBR: "Tarefas",
		ptPT: "Tarefas",
		ru:   "Задачи",
		tr:   "Görevler",
		ukUA: "Завдання",
		uz:   "ToDo",
		zhCN: "任务",
	},
	LIST_CMD_WATCH: {
		arEG: "يشاهد",
		de:   "Ansehen",
		en:   "Watch",
		es:   "Ver",
		faIR: "تماشا کنید",
		fr:   "Regarder",
		id:   "Menonton",
		it:   "Guarda",
		jaJP: "見る",
		koKR: "보기",
		pl:   "Oglądaj",
		ptBR: "Assistir",
		ptPT: "Assistir",
		ru:   "Смотреть",
		tr:   "İzle",
		ukUA: "Дивитися",
		uz:   "Tomosha qiling",
		zhCN: "观看",
	},
	LIST_CMD_READ: {
		arEG: "يقرأ",
		de:   "Lesen",
		en:   "Read",
		es:   "Leer",
		faIR: "خواندن",
		fr:   "Lire",
		id:   "Baca",
		it:   "Leggi",
		jaJP: "読む",
		koKR: "읽기",
		pl:   "Czytaj",
		ptBR: "Ler",
		ptPT: "Ler",
		ru:   "Читать",
		tr:   "Oku",
		ukUA: "Читати",
		uz:   "O&#39;qing",
		zhCN: "读",
	},
	ListCmdBtnToRead: {
		arEG: "للقراءة",
		de:   "Zu lesen",
		en:   "To read",
		es:   "Leer",
		faIR: "برای خواندن",
		fr:   "À lire",
		id:   "Untuk dibaca",
		it:   "Da leggere",
		jaJP: "読む",
		koKR: "읽을 것",
		pl:   "Do przeczytania",
		ptBR: "Para ler",
		ptPT: "Para ler",
		ru:   "Прочитать",
		tr:   "Okunacaklar",
		ukUA: "Читати",
		uz:   "O&#39;qish uchun",
		zhCN: "阅读",
	},
	ListCmdBtnToWatch: {
		arEG: "للمشاهدة",
		de:   "Ansehen",
		en:   "To watch",
		es:   "Ver",
		faIR: "برای تماشا",
		fr:   "À regarder",
		id:   "Untuk ditonton",
		it:   "Da guardare",
		jaJP: "見る",
		koKR: "볼 것",
		pl:   "Do obejrzenia",
		ptBR: "Para assistir",
		ptPT: "Para assistir",
		ru:   "Посмотреть",
		tr:   "İzlenecekler",
		ukUA: "Дивитися",
		uz:   "Tomosha qilish uchun",
		zhCN: "观看",
	},
	ListCmdBtnToBuy: {
		arEG: "للشراء",
		de:   "Zu kaufen",
		en:   "To buy",
		es:   "Comprar",
		faIR: "برای خرید",
		fr:   "À acheter",
		id:   "Untuk dibeli",
		it:   "Da comprare",
		jaJP: "購入",
		koKR: "구매할 것",
		pl:   "Do kupienia",
		ptBR: "Para comprar",
		ptPT: "Para comprar",
		ru:   "Купить",
		tr:   "Alınacaklar",
		ukUA: "Купити",
		uz:   "Sotib olish uchun",
		zhCN: "购买",
	},
	ListCmdBtnToDo: {
		arEG: "للقيام به",
		de:   "Zu tun",
		en:   "To do",
		es:   "Hacer",
		faIR: "انجام دادن",
		fr:   "À faire",
		id:   "Untuk dilakukan",
		it:   "Da fare",
		jaJP: "やる",
		koKR: "할 일",
		pl:   "Do zrobienia",
		ptBR: "Para fazer",
		ptPT: "Para fazer",
		ru:   "Сделать",
		tr:   "Yapılacaklar",
		ukUA: "Зробити",
		uz:   "Qilish uchun",
		zhCN: "做",
	},
	Groceries: {
		arEG: "البقالة",
		de:   "Lebensmittel",
		en:   "Groceries",
		es:   "Comestibles",
		faIR: "خوار و بار",
		fr:   "Épicerie",
		id:   "Bahan makanan",
		it:   "Generi alimentari",
		jaJP: "食料品",
		koKR: "식료품",
		pl:   "Artykuły spożywcze",
		ptBR: "Comestíveis",
		ptPT: "Comestíveis",
		ru:   "Продукты",
		tr:   "Bakkaliye",
		ukUA: "Продукти",
		uz:   "Oziq-ovqat",
		zhCN: "杂货",
	},
	Books: {
		arEG: "الكتب",
		de:   "Bücher",
		en:   "Books",
		es:   "Libros",
		faIR: "کتاب ها",
		fr:   "Livres",
		id:   "Buku",
		it:   "Libri",
		jaJP: "本",
		koKR: "책",
		pl:   "Książki",
		ptBR: "Livros",
		ptPT: "Livros",
		ru:   "Книги",
		tr:   "Kitaplar",
		ukUA: "Книги",
		uz:   "Kitoblar",
		zhCN: "书籍",
	},
	Movies: {
		arEG: "أفلام",
		de:   "Filme",
		en:   "Movies",
		es:   "Películas",
		faIR: "فیلم ها",
		fr:   "Films",
		id:   "Film",
		it:   "Film",
		jaJP: "映画",
		koKR: "영화",
		pl:   "Filmy",
		ptBR: "Filmes",
		ptPT: "Filmes",
		ru:   "Фильмы",
		tr:   "Filmler",
		ukUA: "Фільми",
		uz:   "Filmlar",
		zhCN: "电影",
	},
	Tasks: {
		arEG: "المهام",
		de:   "Aufgaben",
		en:   "Tasks",
		es:   "Tareas",
		faIR: "وظایف",
		fr:   "Tâches",
		id:   "Tugas",
		it:   "Compiti",
		jaJP: "タスク",
		koKR: "할 일",
		pl:   "Zadania",
		ptBR: "Tarefas",
		ptPT: "Tarefas",
		ru:   "Задачи",
		tr:   "Görevler",
		ukUA: "Завдання",
		uz:   "Vazifalar",
		zhCN: "任务",
	},
	ListsOfFamily: {
		arEG: "قوائم العائلة",
		de:   "Familienlisten",
		en:   "Family lists",
		es:   "Listas familiares",
		faIR: "لیست های خانواده",
		fr:   "Listes familiales",
		id:   "Daftar keluarga",
		it:   "Liste familiari",
		jaJP: "家族のリスト",
		koKR: "가족 목록",
		pl:   "Listy rodzinne",
		ptBR: "Listas familiares",
		ptPT: "Listas familiares",
		ru:   "Семейные списки",
		tr:   "Aile listeleri",
		ukUA: "Сімейні списки",
		uz:   "Oila ro&#39;yxati",
		zhCN: "家庭清单",
	},
	ListsOfPrivate: {
		arEG: "القوائم الخاصة",
		de:   "Private Listen",
		en:   "Private lists",
		es:   "Listas privadas",
		faIR: "لیست های خصوصی",
		fr:   "Listes privées",
		id:   "Daftar pribadi",
		it:   "Liste private",
		jaJP: "プライベートリスト",
		koKR: "개인 목록",
		pl:   "Prywatne listy",
		ptBR: "Listas privadas",
		ptPT: "Listas privadas",
		ru:   "Личные списки",
		tr:   "Özel listeler",
		ukUA: "Особисті списки",
		uz:   "Shaxsiy ro&#39;yxatlar",
		zhCN: "私人清单",
	},
	ListsOfSpace: {
		arEG: "القوائم @ %s",
		de:   "Listen @ %s",
		en:   "Lists @ %s",
		es:   "Listas @ %s",
		faIR: "لیست ها @ %s",
		fr:   "Listes @ %s",
		id:   "Daftar @ %s",
		it:   "Liste @ %s",
		jaJP: "リスト @ %s",
		koKR: "목록 @ %s",
		pl:   "Listy @ %s",
		ptBR: "Listas @ %s",
		ptPT: "Listas @ %s",
		ru:   "Списки @ %s",
		tr:   "Listeler @ %s",
		ukUA: "Списки @ %s",
		uz:   "Ro'yxatlar @ %s",
		zhCN: "清单 @ %s",
	},
	BtnBackToSpace: {
		arEG: "العودة إلى الفضاء",
		de:   "Zurück zum Bereich",
		en:   "Back to space",
		es:   "Volver al espacio",
		faIR: "بازگشت به فضا",
		fr:   "Retour à l'espace",
		id:   "Kembali ke ruang",
		it:   "Torna allo spazio",
		jaJP: "スペースに戻る",
		koKR: "공간으로 돌아가기",
		pl:   "Powrót do przestrzeni",
		ptBR: "Voltar ao espaço",
		ptPT: "Voltar ao espaço",
		ru:   "К пространству",
		tr:   "Alana geri dön",
		ukUA: "Назад до простору",
		uz:   "Kosmosga qaytish",
		zhCN: "返回空间",
	},
	BtnPrivate: {
		arEG: "خاص",
		de:   "Privat",
		en:   "Private",
		es:   "Privado",
		faIR: "شخصی",
		fr:   "Privé",
		id:   "Pribadi",
		it:   "Privato",
		jaJP: "個人",
		koKR: "개인적인",
		pl:   "Prywatne",
		ptBR: "Privado",
		ptPT: "Privado",
		ru:   "Личное",
		tr:   "Özel",
		ukUA: "Особисте",
		uz:   "Shaxsiy",
		zhCN: "私人",
	},
	BtnFamily: {
		arEG: "عائلة",
		de:   "Familie",
		en:   "Family",
		es:   "Familia",
		faIR: "خانواده",
		fr:   "Famille",
		id:   "Keluarga",
		it:   "Famiglia",
		jaJP: "家族",
		koKR: "가족",
		pl:   "Rodzina",
		ptBR: "Família",
		ptPT: "Família",
		ru:   "Семья",
		tr:   "Aile",
		ukUA: "Сім'я",
		uz:   "Oila",
		zhCN: "家庭",
	},
	TrackerPushUps: {
		arEG: "تمارين الضغط",
		de:   "Liegestütze",
		en:   "Push-ups",
		es:   "Flexiones",
		faIR: "پرسه",
		fr:   "Pompes",
		id:   "Push-up",
		it:   "Flessioni",
		jaJP: "腕立て伏せ",
		koKR: "푸시업",
		pl:   "Pompki",
		ptBR: "Flexões",
		ptPT: "Flexões",
		ru:   "Отжимания",
		tr:   "Şınav",
		ukUA: "Віджимання",
		uz:   "Otjimaniye&quot; mashqi",
		zhCN: "俯卧撑",
	},
	TrackerPullUps: {
		arEG: "عمليات السحب",
		de:   "Klimmzüge",
		en:   "Pull-ups",
		es:   "Dominadas",
		faIR: "کشیدن",
		fr:   "Tirages",
		id:   "Pull-up",
		it:   "Trazioni",
		jaJP: "引き上げ",
		koKR: "풀업",
		pl:   "Podciągania",
		ptBR: "Flexões",
		ptPT: "Flexões",
		ru:   "Подтягивания",
		tr:   "Çekme",
		ukUA: "Підтягування",
		uz:   "Pull-uplar",
		zhCN: "引体向上",
	},
	TrackerSquats: {
		arEG: "القرفصاء",
		de:   "Kniebeugen",
		en:   "Squats",
		es:   "Sentadillas",
		faIR: "کرنچ",
		fr:   "Squats",
		id:   "Squat",
		it:   "Squat",
		jaJP: "スクワット",
		koKR: "스쿼트",
		pl:   "Przysiady",
		ptBR: "Agachamentos",
		ptPT: "Agachamentos",
		ru:   "Приседания",
		tr:   "Squat",
		ukUA: "Присідання",
		uz:   "Squats",
		zhCN: "深蹲",
	},
	TrackerJumpingJacks: {
		arEG: "القفزات",
		de:   "Hampelmänner",
		en:   "Jumping jacks",
		es:   "Saltos de tijera",
		faIR: "حرکت پروانه",
		fr:   "Sauts avec écart",
		id:   "Lompat bintang",
		it:   "Salti con apertura",
		jaJP: "ジャンピングジャック",
		koKR: "팔 벌려 뛰기",
		pl:   "Pajacyki",
		ptBR: "Polichinelos",
		ptPT: "Saltos de estrela",
		ru:   "Прыжки на месте",
		tr:   "Zıplama hareketi",
		ukUA: "Стрибки з розмахом рук",
		uz:   "Sakrash jaklari",
		zhCN: "开合跳",
	},
	TrackerFuelCost: {
		arEG: "تكلفة الوقود",
		de:   "Kraftstoffkosten",
		en:   "Fuel Cost",
		es:   "Costo del combustible",
		faIR: "هزینه سوخت",
		fr:   "Coût du carburant",
		id:   "Biaya bahan bakar",
		it:   "Costo del carburante",
		jaJP: "燃料費",
		koKR: "연료 비용",
		pl:   "Koszt paliwa",
		ptBR: "Custo do combustível",
		ptPT: "Custo do combustível",
		ru:   "Стоимость топлива",
		tr:   "Yakıt Maliyeti",
		ukUA: "Вартість палива",
		uz:   "Yoqilg&#39;i narxi",
		zhCN: "燃料成本",
	},
	TrackerFuelVolume: {
		arEG: "حجم الوقود",
		de:   "Kraftstoffvolumen",
		en:   "Fuel Volume",
		es:   "Volumen de combustible",
		faIR: "حجم سوخت",
		fr:   "Volume de carburant",
		id:   "Volume bahan bakar",
		it:   "Volume di carburante",
		jaJP: "燃料量",
		koKR: "연료 양",
		pl:   "Objętość paliwa",
		ptBR: "Volume de combustível",
		ptPT: "Volume de combustível",
		ru:   "Объем топлива",
		tr:   "Yakıt Hacmi",
		ukUA: "Об'єм палива",
		uz:   "Yoqilg&#39;i hajmi",
		zhCN: "燃料容量",
	},
	TrackerMileage: {
		arEG: "المسافة المقطوعة",
		de:   "Kilometerstand",
		en:   "Mileage",
		es:   "Kilometraje",
		faIR: "مسافت پیموده شده",
		fr:   "Kilométrage",
		id:   "Jarak tempuh",
		it:   "Chilometraggio",
		jaJP: "走行距離",
		koKR: "주행 거리",
		pl:   "Przebieg",
		ptBR: "Quilometragem",
		ptPT: "Quilometragem",
		ru:   "Пробег",
		tr:   "Kilometre",
		ukUA: "Пробіг",
		uz:   "Kilometr",
		zhCN: "里程",
	},
	TrackerHeight: {
		arEG: "ارتفاع",
		de:   "Höhe",
		en:   "Height",
		es:   "Altura",
		faIR: "قد",
		fr:   "Taille",
		id:   "Tinggi",
		it:   "Altezza",
		jaJP: "身長",
		koKR: "키",
		pl:   "Wzrost",
		ptBR: "Altura",
		ptPT: "Altura",
		ru:   "Рост",
		tr:   "Boy",
		ukUA: "Зріст",
		uz:   "Balandligi",
		zhCN: "身高",
	},
	TrackerWeight: {
		arEG: "وزن",
		de:   "Gewicht",
		en:   "Weight",
		es:   "Peso",
		faIR: "وزن",
		fr:   "Poids",
		id:   "Berat",
		it:   "Peso",
		jaJP: "体重",
		koKR: "몸무게",
		pl:   "Waga",
		ptBR: "Peso",
		ptPT: "Peso",
		ru:   "Вес",
		tr:   "Ağırlık",
		ukUA: "Вага",
		uz:   "Og&#39;irligi",
		zhCN: "体重",
	},
	TrackerSpending: {
		arEG: "الإنفاق",
		de:   "Ausgaben",
		en:   "Spending",
		es:   "Gastos",
		faIR: "هزینه\u200cها",
		fr:   "Dépenses",
		id:   "Pengeluaran",
		it:   "Spese",
		jaJP: "支出",
		koKR: "지출",
		pl:   "Wydatki",
		ptBR: "Despesas",
		ptPT: "Despesas",
		ru:   "Расходы",
		tr:   "Harcamalar",
		ukUA: "Витрати",
		uz:   "Sarflash",
		zhCN: "支出",
	},
	TrackerBloodPressure: {
		arEG: "ضغط الدم",
		de:   "Blutdruck",
		en:   "Blood Pressure",
		es:   "Presión arterial",
		faIR: "فشار خون",
		fr:   "Pression artérielle",
		id:   "Tekanan darah",
		it:   "Pressione sanguigna",
		jaJP: "血圧",
		koKR: "혈압",
		pl:   "Ciśnienie krwi",
		ptBR: "Pressão arterial",
		ptPT: "Pressão arterial",
		ru:   "Кровяное давление",
		tr:   "Kan basıncı",
		ukUA: "Кров'яний тиск",
		uz:   "Qon bosimi",
		zhCN: "血压",
	},
	TrackerCategoryHealth: {
		arEG: "صحة",
		de:   "Gesundheit",
		en:   "Health",
		es:   "Salud",
		faIR: "سلامتی",
		fr:   "Santé",
		id:   "Kesehatan",
		it:   "Salute",
		jaJP: "健康",
		koKR: "건강",
		pl:   "Zdrowie",
		ptBR: "Saúde",
		ptPT: "Saúde",
		ru:   "Здоровье",
		tr:   "Sağlık",
		ukUA: "Здоров'я",
		uz:   "Salomatlik",
		zhCN: "健康",
	},
	TrackerCategoryFitness: {
		arEG: "لياقة بدنية",
		de:   "Fitness",
		en:   "Fitness",
		es:   "Aptitud física",
		faIR: "تناسب اندام",
		fr:   "Forme physique",
		id:   "Kebugaran fisik",
		it:   "Fitness",
		jaJP: "フィットネス",
		koKR: "피트니스",
		pl:   "Zdatność",
		ptBR: "Aptidão física",
		ptPT: "Aptidão física",
		ru:   "Фитнес",
		tr:   "Uygunluk",
		ukUA: "Фітнес",
		uz:   "Fitnes",
		zhCN: "健身",
	},
	TrackerCategoryVehicle: {
		arEG: "عربة",
		de:   "Fahrzeug",
		en:   "Vehicle",
		es:   "Vehículo",
		faIR: "وسیله نقلیه",
		fr:   "Véhicule",
		id:   "Kendaraan",
		it:   "Veicolo",
		jaJP: "車両",
		koKR: "차량",
		pl:   "Pojazd",
		ptBR: "Veículo",
		ptPT: "Veículo",
		ru:   "Транспортное средство",
		tr:   "Araç",
		ukUA: "Транспортний засіб",
		uz:   "Avtomobil",
		zhCN: "车辆",
	},
	CustomTracker: {
		arEG: "متتبع مخصص",
		de:   "Benutzerdefinierter Tracker",
		en:   "Custom tracker",
		es:   "Rastreador personalizado",
		faIR: "ردیاب سفارشی",
		fr:   "Tracker personnalisé",
		id:   "Pelacak kustom",
		it:   "Tracker personalizzato",
		jaJP: "カスタムトラッカー",
		koKR: "사용자 정의 추적기",
		pl:   "Niestandardowy tracker",
		ptBR: "Rastreador personalizado",
		ptPT: "Rastreador personalizado",
		ru:   "Свой трекер",
		tr:   "Özel izleyici",
		ukUA: "Свій трекер",
		uz:   "Maxsus kuzatuvchi",
		zhCN: "自定义追踪器",
	},
	BackToTrackers: {
		arEG: "العودة إلى المتتبعين",
		de:   "Zurück zu Trackern",
		en:   "Back to trackers",
		es:   "Volver a los rastreadores",
		faIR: "بازگشت به ردیاب\u200cها",
		fr:   "Retour aux trackers",
		id:   "Kembali ke pelacak",
		it:   "Torna ai tracker",
		jaJP: "トラッカーに戻る",
		koKR: "트래커로 돌아가기",
		pl:   "Powrót do trackerów",
		ptBR: "Voltar aos rastreadores",
		ptPT: "Voltar aos rastreadores",
		ru:   "Назад к трекерам",
		tr:   "İzleyicilere geri dön",
		ukUA: "Повернутись до трекерів",
		uz:   "Kuzatuvchilar sahifasiga qaytish",
		zhCN: "返回追踪器",
	},
	AddTracker: {
		arEG: "إضافة متتبع",
		de:   "Tracker hinzufügen",
		en:   "Add tracker",
		es:   "Añadir rastreador",
		faIR: "افزودن ردیاب",
		fr:   "Ajouter un tracker",
		id:   "Tambahkan pelacak",
		it:   "Aggiungi tracker",
		jaJP: "トラッカーを追加",
		koKR: "추적기 추가",
		pl:   "Dodaj tracker",
		ptBR: "Adicionar rastreador",
		ptPT: "Adicionar rastreador",
		ru:   "Добавить трекер",
		tr:   "İzleyici ekle",
		ukUA: "Додати трекер",
		uz:   "Kuzatuvchi qo&#39;shing",
		zhCN: "添加追踪器",
	},
	ShareTracker: {
		arEG: "متتبع المشاركة",
		de:   "Tracker teilen",
		en:   "Share tracker",
		es:   "Compartir rastreador",
		faIR: "اشتراک\u200cگذاری ردیاب",
		fr:   "Partager le tracker",
		id:   "Bagikan pelacak",
		it:   "Condividi il tracker",
		jaJP: "トラッカーを共有",
		koKR: "추적기 공유",
		pl:   "Udostępnij tracker",
		ptBR: "Compartilhar rastreador",
		ptPT: "Partilhar rastreador",
		ru:   "Поделиться трекером",
		tr:   "İzleyiciyi paylaş",
		ukUA: "Поділитись трекером",
		uz:   "Ulashish kuzatuvchisi",
		zhCN: "分享追踪器",
	},
	RecordTrackerValue: {
		arEG: "قيمة السجل",
		de:   "Wert aufzeichnen",
		en:   "Record value",
		es:   "Registrar valor",
		faIR: "ثبت مقدار",
		fr:   "Enregistrer la valeur",
		id:   "Catat nilai",
		it:   "Registra valore",
		jaJP: "値を記録",
		koKR: "값 기록",
		pl:   "Zapisz wartość",
		ptBR: "Registrar valor",
		ptPT: "Registrar valor",
		ru:   "Записать значение",
		tr:   "Değeri kaydet",
		ukUA: "Записати значення",
		uz:   "Rekord qiymati",
		zhCN: "记录值",
	},
	FamilyTrackers: {
		arEG: "متتبعات العائلة",
		de:   "Familientracker",
		en:   "Family trackers",
		es:   "Rastreadores familiares",
		faIR: "ردیاب\u200cهای خانواده",
		fr:   "Trackers familiaux",
		id:   "Pelacak keluarga",
		it:   "Tracker familiari",
		jaJP: "家族のトラッカー",
		koKR: "가족 추적기",
		pl:   "Trackery rodzinne",
		ptBR: "Rastreadores de família",
		ptPT: "Rastreadores de família",
		ru:   "Семейные трекеры",
		tr:   "Aile izleyicileri",
		ukUA: "Сімейні трекери",
		uz:   "Oila kuzatuvchilari",
		zhCN: "家庭追踪器",
	},
	NoActiveTrackers: {
		arEG: "ليس لديك أي أدوات تتبع نشطة حاليًا. أضف بعض الأدوات لبدء تتبع تقدمك.",
		de:   "Derzeit haben Sie keine aktiven Tracker. Fügen Sie einige hinzu, um Ihren Fortschritt zu verfolgen.",
		en:   "Currently you do not have any active trackers. Add some to start tracking your progress.",
		es:   "Actualmente, no tienes rastreadores activos. Agrega algunos para comenzar a rastrear tu progreso.",
		faIR: "در حال حاضر هیچ ردیابی فعال ندارید. چند مورد اضافه کنید تا پیشرفت خود را پیگیری کنید.",
		fr:   "Vous n'avez actuellement aucun tracker actif. Ajoutez-en pour commencer à suivre vos progrès.",
		id:   "Saat ini Anda tidak memiliki pelacak aktif. Tambahkan beberapa untuk mulai melacak kemajuan Anda.",
		it:   "Attualmente non hai tracker attivi. Aggiungine alcuni per iniziare a monitorare i tuoi progressi.",
		jaJP: "現在、有効なトラッカーはありません。進捗を追跡するには、いくつか追加してください。",
		koKR: "현재 활성화된 추적기가 없습니다. 진행 상황을 추적하려면 몇 가지를 추가하세요.",
		pl:   "Obecnie nie masz żadnych aktywnych trackerów. Dodaj kilka, aby zacząć śledzić swoje postępy.",
		ptBR: "Atualmente, você não tem rastreadores ativos. Adicione alguns para começar a rastrear seu progresso.",
		ptPT: "Atualmente, não tens rastreadores ativos. Adiciona alguns para começar a monitorizar o teu progresso.",
		ru:   "В настоящее время у вас нет активных трекеров. Добавьте их, чтобы начать отслеживать ваш прогресс.",
		tr:   "Şu anda aktif izleyiciniz yok. İlerlemenizi izlemek için birkaç tane ekleyin.",
		ukUA: "Наразі у вас немає активних трекерів. Додайте кілька, щоб почати відстежувати свій прогрес.",
		uz:   "Hozirda sizda faol kuzatuvchilar yo&#39;q. Taraqqiyotingizni kuzatishni boshlash uchun biroz qo&#39;shing.",
		zhCN: "目前你还没有任何活跃的追踪器。添加一些以开始追踪你的进度。",
	},
	OneActiveTracker: {
		arEG: "لديك مُتتبّع واحد نشط. اختره لتسجيل التقدم. أو أضف مُتتبّعًا جديدًا لتتبع المزيد.",
		de:   "Sie haben einen aktiven Tracker. Wählen Sie ihn aus, um Fortschritte aufzuzeichnen, oder fügen Sie einen neuen hinzu, um mehr zu verfolgen.",
		en:   "You have one active tracker. Select it to record progress. Or add a new one to track more.",
		es:   "Tienes un rastreador activo. Selecciónalo para registrar el progreso o agrega uno nuevo para rastrear más.",
		faIR: "شما یک ردیاب فعال دارید. آن را انتخاب کنید تا پیشرفت خود را ثبت کنید یا یک ردیاب جدید اضافه کنید تا موارد بیشتری را پیگیری کنید.",
		fr:   "Vous avez un tracker actif. Sélectionnez-le pour enregistrer vos progrès ou ajoutez-en un nouveau pour en suivre davantage.",
		id:   "Anda memiliki satu pelacak aktif. Pilih untuk mencatat kemajuan atau tambahkan yang baru untuk melacak lebih banyak.",
		it:   "Hai un tracker attivo. Selezionalo per registrare i progressi o aggiungi uno nuovo per monitorare di più.",
		jaJP: "アクティブなトラッカーが1つあります。選択して進捗を記録するか、新しいものを追加してさらに追跡してください。",
		koKR: "활성화된 추적기 하나가 있습니다. 선택하여 진행 상황을 기록하거나 더 많이 추적하려면 새로 추가하세요.",
		pl:   "Masz jeden aktywny tracker. Wybierz go, aby zarejestrować postępy, lub dodaj nowy, aby śledzić więcej.",
		ptBR: "Você tem um rastreador ativo. Selecione-o para registrar o progresso ou adicione um novo para rastrear mais.",
		ptPT: "Tens um rastreador ativo. Seleciona-o para registar o progresso, ou adiciona um novo para monitorizar mais.",
		ru:   "У вас есть один активный трекер. Выберите его, чтобы записать прогресс, или добавьте новый для отслеживания большего.",
		tr:   "Bir adet aktif izleyiciniz var. İlerlemenizi kaydetmek için bunu seçin veya daha fazla şey izlemek için yeni bir tane ekleyin.",
		ukUA: "У вас є один активний трекер. Виберіть його, щоб записати прогрес, або додайте новий для відстеження більшого.",
		uz:   "Sizda bitta faol kuzatuvchi bor. Jarayonni yozib olish uchun uni tanlang. Yoki ko&#39;proq kuzatish uchun yangisini qo&#39;shing.",
		zhCN: "你有一个活跃的追踪器。选择它来记录你的进度，或者添加一个新的来追踪更多。",
	},
	NActiveTrackers: {
		arEG: "لديك عدد قليل من أجهزة التتبع النشطة. اختر واحدة لتسجيل التقدم.",
		de:   "Sie haben einige aktive Tracker. Wählen Sie einen aus, um Fortschritte aufzuzeichnen.",
		en:   "You have few active tracker. Select one to record progress.",
		es:   "Tienes varios rastreadores activos. Selecciona uno para registrar el progreso.",
		faIR: "چند ردیاب فعال دارید. یکی را انتخاب کنید تا پیشرفت خود را ثبت کنید.",
		fr:   "Vous avez plusieurs trackers actifs. Sélectionnez-en un pour enregistrer vos progrès.",
		id:   "Anda memiliki beberapa pelacak aktif. Pilih salah satu untuk mencatat kemajuan.",
		it:   "Hai alcuni tracker attivi. Selezionane uno per registrare i progressi.",
		jaJP: "複数のアクティブなトラッカーがあります。1つ選択して進捗を記録してください。",
		koKR: "활성화된 추적기 몇 개가 있습니다. 하나를 선택하여 진행 상황을 기록하세요.",
		pl:   "Masz kilka aktywnych trackerów. Wybierz jeden, aby zarejestrować postępy.",
		ptBR: "Você tem vários rastreadores ativos. Selecione um para registrar o progresso.",
		ptPT: "Tens vários rastreadores ativos. Seleciona um para registar o progresso.",
		ru:   "У вас несколько активных трекеров. Выберите один, чтобы записать прогресс.",
		tr:   "Birkaç aktif izleyiciniz var. İlerlemenizi kaydetmek için birini seçin.",
		ukUA: "У вас є кілька активних трекерів. Виберіть один, щоб записати прогрес.",
		uz:   "Sizda bir nechta faol kuzatuvchi bor. Rivojlanishni qayd etish uchun birini tanlang.",
		zhCN: "你有几个活跃的追踪器。选择一个来记录你的进度。",
	},
	HintForIntTracker: {
		arEG: "أرسل رقمًا صحيحًا لتسجيله في جهاز التتبع",
		de:   "Senden Sie eine Ganzzahl, um sie im Tracker zu speichern",
		en:   "Send an integer number to record it to the tracker",
		es:   "Envía un número entero para registrarlo en el rastreador",
		faIR: "یک عدد صحیح ارسال کنید تا آن را در ردیاب ثبت کنید",
		fr:   "Envoyez un nombre entier pour l'enregistrer dans le tracker",
		id:   "Kirim bilangan bulat untuk mencatatnya ke pelacak",
		it:   "Invia un numero intero per registrarlo nel tracker",
		jaJP: "整数値を送信してトラッカーに記録してください",
		koKR: "정수 값을 추적기에 기록하려면 전송하세요",
		pl:   "Wyślij liczbę całkowitą, aby zapisać ją w trackerze",
		ptBR: "Envie um número inteiro para registrá-lo no rastreador",
		ptPT: "Envia um número inteiro para o registar no rastreador",
		ru:   "Отправьте целое число, чтобы записать его в трекер",
		tr:   "Bir tamsayı göndererek bunu izleyiciye kaydedin",
		ukUA: "Надішліть ціле число, щоб записати його в трекер",
		uz:   "Butun sonni kuzatuvchiga yozib olish uchun yuboring",
		zhCN: "发送一个整数以将其记录到追踪器中",
	},
	Tracker: {
		arEG: "متتبع",
		de:   "Tracker",
		en:   "Tracker",
		es:   "Rastreador",
		faIR: "ردیاب",
		fr:   "Traqueur",
		id:   "Pelacak",
		it:   "Tracciatore",
		jaJP: "トラッカー",
		koKR: "추적기",
		pl:   "Abstrakt",
		ptBR: "Rastreador",
		ptPT: "Rastreador",
		ru:   "Трекер",
		tr:   "İzleyici",
		ukUA: "Трекер",
		uz:   "Kuzatuvchi",
		zhCN: "追踪器",
	},
	HintToShareTracker: {
		arEG: `يمكنك مشاركة هذا المتعقب إما 
 🤫 مع صديق ( <i>شخص واحد سيكون قادرًا على القبول</i> ) 
 🌍 أو بشكل عام ( <i>يمكن لأي شخص لديه رابط القبول، ويمكنك إلغاؤه في أي وقت</i> )`,
		de: `\n\nSie können diesen Tracker entweder
🤫 mit einem Freund teilen (<i>nur eine Person kann ihn akzeptieren</i>)
🌍 oder öffentlich (<i>jeder mit dem Link kann ihn akzeptieren, Sie können es jederzeit abbrechen</i>)`,
		en: `\n\nYou can share this tracker either
	🤫 with a friend (<i>a single person would be able to accept</i>)
	🌍 or publicly (<i>anyone with a link can accept, you can cancel it at any time</i>)`,
		es: `\n\nPuedes compartir este rastreador
🤫 con un amigo (<i>una sola persona puede aceptarlo</i>)
🌍 o públicamente (<i>cualquiera con el enlace puede aceptarlo, puedes cancelarlo en cualquier momento</i>)`,
		faIR: `\n\nشما می\u200cتوانید این ردیاب را به دو صورت به اشتراک بگذارید
🤫 با یک دوست (<i>فقط یک نفر می\u200cتواند بپذیرد</i>)
🌍 یا به صورت عمومی (<i>هرکسی با لینک می\u200cتواند بپذیرد، شما می\u200cتوانید در هر زمان آن را لغو کنید</i>)`,
		fr: `\n\nVous pouvez partager ce tracker
🤫 avec un ami (<i>une seule personne pourra l'accepter</i>)
🌍 ou publiquement (<i>n'importe qui avec un lien peut l'accepter, vous pouvez l'annuler à tout moment</i>)`,
		id: `\n\nAnda dapat membagikan pelacak ini
🤫 kepada seorang teman (<i>hanya satu orang yang dapat menerima</i>)
🌍 atau secara publik (<i>siapa saja yang memiliki tautan dapat menerimanya, Anda dapat membatalkannya kapan saja</i>)`,
		it: `\n\nPuoi condividere questo tracker
🤫 con un amico (<i>una sola persona potrà accettarlo</i>)
🌍 oppure pubblicamente (<i>chiunque con il link potrà accettarlo, puoi annullarlo in qualsiasi momento</i>)`,
		jaJP: `\n\nこのトラッカーを以下の方法で共有できます
🤫 友人と (<i>受け取れるのは1人だけ</i>)
🌍 または公開で (<i>リンクを持っている誰でも受け取れます。いつでもキャンセルできます</i>)`,
		koKR: `\n\n이 추적기를 공유할 수 있습니다
🤫 친구와 (<i>단 한 사람만 수락 가능</i>)
🌍 또는 공개적으로 (<i>링크를 가진 모든 사람이 수락할 수 있으며, 언제든 취소할 수 있습니다</i>)`,
		pl: `\n\nMożesz udostępnić ten tracker
🤫 znajomemu (<i>tylko jedna osoba będzie mogła przyjąć</i>)
🌍 lub publicznie (<i>każdy z linkiem może przyjąć, możesz to anulować w dowolnym momencie</i>)`,
		ptBR: `\n\nVocê pode compartilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa pode aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, você pode cancelar a qualquer momento</i>)`,
		ptPT: `\n\nPodes partilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa poderá aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, podes cancelá-lo a qualquer momento</i>)`,
		ru: `\n\nВы можете поделиться этим трекером либо
🤫 с другом (<i>только один человек сможет принять</i>)
🌍 или публично (<i>любой с ссылкой может принять, вы можете отменить это в любое время</i>)`,
		tr: `\n\nBu izleyiciyi şu şekilde paylaşabilirsiniz
🤫 bir arkadaşla (<i>yalnızca bir kişi kabul edebilir</i>)
🌍 veya genel olarak (<i>bağlantıya sahip herkes kabul edebilir, istediğiniz zaman iptal edebilirsiniz</i>)`,
		ukUA: `\n\nВи можете поділитися цим трекером
🤫 з другом (<i>лише одна людина зможе прийняти</i>)
🌍 або публічно (<i>кожен, хто має посилання, може прийняти, ви можете скасувати це в будь-який час</i>)`,
		uz: `\n\nSiz ushbu trekerni 
 🤫 doʻstingiz bilan baham koʻrishingiz mumkin ( <i>bir kishi qabul qilishi mumkin</i> )
 🌍 yoki hammaga ochiq ( <i>havolasi boʻlgan har kim qabul qilishi mumkin, uni istalgan vaqtda bekor qilishingiz mumkin</i> )`,
		zhCN: `\n\n您可以通过以下方式共享此追踪器
🤫 与朋友共享 (<i>只有一个人可以接受</i>)
🌍 或公开共享 (<i>任何拥有链接的人都可以接受，您可以随时取消</i>)`,
	},
	YourSpaces: {
		arEG: "مساحاتك",
		de:   "Ihre Bereiche",
		en:   "Your spaces",
		es:   "Tus espacios",
		faIR: "فضاهای شما",
		fr:   "Vos espaces",
		id:   "Ruang Anda",
		it:   "I tuoi spazi",
		jaJP: "あなたのスペース",
		koKR: "귀하의 공간",
		pl:   "Twoje przestrzenie",
		ptBR: "Seus espaços",
		ptPT: "Os teus espaços",
		ru:   "Ваши пространства",
		tr:   "Alanlarınız",
		ukUA: "Ваші простори",
		uz:   "Ваши пространства",
		zhCN: "您的空间",
	},
	CurrentSpace: {
		arEG: "المساحة الحالية",
		de:   "Aktueller Bereich",
		en:   "Current space",
		es:   "Espacio actual",
		faIR: "فضای کنونی",
		fr:   "Espace actuel",
		id:   "Ruang saat ini",
		it:   "Spazio corrente",
		jaJP: "現在のスペース",
		koKR: "현재 공간",
		pl:   "Bieżąca przestrzeń",
		ptBR: "Espaço atual",
		ptPT: "Espaço atual",
		ru:   "Текущее пространство",
		tr:   "Mevcut alan",
		ukUA: "Поточний простір",
		uz:   "Joriy bo&#39;sh joy",
		zhCN: "当前空间",
	},
	ClickToSwitchCurrentSpace: {
		arEG: "انقر للتبديل إلى المساحة الحالية",
		de:   "Klicken Sie hier, um den aktuellen Bereich zu wechseln",
		en:   "Click to switch current space",
		es:   "Haz clic para cambiar el espacio actual",
		faIR: "برای تغییر فضای کنونی کلیک کنید",
		fr:   "Cliquez pour changer l'espace actuel",
		id:   "Klik untuk mengganti ruang saat ini",
		it:   "Clicca per cambiare spazio corrente",
		jaJP: "クリックで現在のスペースを切り替えます",
		koKR: "현재 공간을 전환하려면 클릭하세요",
		pl:   "Kliknij, aby zmienić bieżącą przestrzeń",
		ptBR: "Clique para alterar o espaço atual",
		ptPT: "Clique para mudar o espaço atual",
		ru:   "Нажмите, чтобы переключить текущее пространство",
		tr:   "Mevcut alanı değiştirmek için tıklayın",
		ukUA: "Натисніть, щоб змінити поточний простір",
		uz:   "Joriy joyni almashtirish uchun bosing",
		zhCN: "点击切换当前空间",
	},
	Family: {
		arEG: "عائلة",
		de:   "Familie",
		en:   "Family",
		es:   "Familia",
		faIR: "خانواده",
		fr:   "Famille",
		id:   "Keluarga",
		it:   "Famiglia",
		jaJP: "家族",
		koKR: "가족",
		pl:   "Rodzina",
		ptBR: "Família",
		ptPT: "Família",
		ru:   "Семья",
		tr:   "Aile",
		ukUA: "Родина",
		uz:   "Oila",
		zhCN: "家庭",
	},
	Private: {
		arEG: "خاص",
		de:   "Privat",
		en:   "Private",
		es:   "Privado",
		faIR: "خصوصی",
		fr:   "Privé",
		id:   "Pribadi",
		it:   "Privato",
		jaJP: "プライベート",
		koKR: "개인",
		pl:   "Prywatne",
		ptBR: "Privado",
		ptPT: "Privado",
		ru:   "Личное",
		tr:   "Özel",
		ukUA: "Приватне",
		uz:   "Shaxsiy",
		zhCN: "私人",
	},
	ChooseSpace: {
		arEG: "اختر المساحة لجعلها نشطة للأوامر الأخرى.",
		de:   "Wählen Sie einen Bereich, um ihn für andere Befehle zu aktivieren.",
		en:   "Choose space to make it active for other commands.",
		es:   "Elija un espacio para activarlo para otros comandos.",
		faIR: "یک فضا را انتخاب کنید تا برای دستورات دیگر فعال شود.",
		fr:   "Choisissez un espace pour le rendre actif pour d'autres commandes.",
		id:   "Pilih ruang untuk menjadikannya aktif untuk perintah lainnya.",
		it:   "Scegli uno spazio per renderlo attivo per altri comandi.",
		jaJP: "他のコマンド用にアクティブにするスペースを選択してください。",
		koKR: "다른 명령에 활성화하려면 공간을 선택하세요.",
		pl:   "Wybierz przestrzeń, aby była aktywna dla innych poleceń.",
		ptBR: "Escolha um espaço para torná-lo ativo para outros comandos.",
		ptPT: "Escolha um espaço para torná-lo ativo para outros comandos.",
		ru:   "Выберите пространство, чтобы сделать его активным для других команд.",
		tr:   "Diğer komutlar için aktif yapmak üzere bir alan seçin.",
		ukUA: "Виберіть простір, щоб зробити його активним для інших команд.",
		uz:   "Uni boshqa buyruqlar uchun faol qilish uchun bo&#39;sh joyni tanlang.",
		zhCN: "选择一个空间以使其对其他命令有效。",
	},
	FamilyContacts: {
		arEG: "اتصالات العائلة",
		de:   "Familienkontakte",
		en:   "Family contacts",
		es:   "Contactos familiares",
		faIR: "مخاطبین خانواده",
		fr:   "Contacts familiaux",
		id:   "Kontak keluarga",
		it:   "Contatti familiari",
		jaJP: "家族の連絡先",
		koKR: "가족 연락처",
		pl:   "Kontakty rodzinne",
		ptBR: "Contatos familiares",
		ptPT: "Contactos familiares",
		ru:   "Семейные контакты",
		tr:   "Aile kişileriniz",
		ukUA: "Сімейні контакти",
		uz:   "Oilaviy aloqalar",
		zhCN: "家庭联系人",
	},
	YourContacts: {
		arEG: "جهات الاتصال الخاصة بك",
		de:   "Ihre Kontakte",
		en:   "Your contacts",
		es:   "Tus contactos",
		faIR: "مخاطبین شما",
		fr:   "Vos contacts",
		id:   "Kontak Anda",
		it:   "I tuoi contatti",
		jaJP: "あなたの連絡先",
		koKR: "당신의 연락처",
		pl:   "Twoje kontakty",
		ptBR: "Seus contatos",
		ptPT: "Os seus contactos",
		ru:   "Ваши контакты",
		tr:   "Kişileriniz",
		ukUA: "Ваші контакти",
		uz:   "Sizning kontaktlaringiz",
		zhCN: "您的联系人",
	},
	ContactsTitle: {
		arEG: "جهات الاتصال",
		de:   "Kontakte",
		en:   "Contacts",
		es:   "Contactos",
		faIR: "مخاطبین",
		fr:   "Contacts",
		id:   "Kontak",
		it:   "Contatti",
		jaJP: "連絡先",
		koKR: "연락처",
		pl:   "Kontakty",
		ptBR: "Contatos",
		ptPT: "Contactos",
		ru:   "Контакты",
		tr:   "Kişiler",
		ukUA: "Контакти",
		uz:   "Kontaktlar",
		zhCN: "联系人",
	},
	OpenInApp: {
		arEG: "افتح في التطبيق",
		de:   "In der App öffnen",
		en:   "Open in app",
		es:   "Abrir en la aplicación",
		faIR: "باز کردن در برنامه",
		fr:   "Ouvrir dans l'application",
		id:   "Buka di aplikasi",
		it:   "Apri nell'app",
		jaJP: "アプリで開く",
		koKR: "앱에서 열기",
		pl:   "Otwórz w aplikacji",
		ptBR: "Abrir no aplicativo",
		ptPT: "Abrir na aplicação",
		ru:   "Открыть в приложении",
		tr:   "Uygulamada aç",
		ukUA: "Відкрити у додатку",
		uz:   "Ilovada oching",
		zhCN: "在应用中打开",
	},
	ManageInApp: {
		arEG: "إدارة في التطبيق",
		de:   "In der App verwalten",
		en:   "Manage in app",
		es:   "Gestionar en la aplicación",
		faIR: "مدیریت در برنامه",
		fr:   "Gérer dans l'application",
		id:   "Kelola di aplikasi",
		it:   "Gestisci nell'app",
		jaJP: "アプリで管理",
		koKR: "앱에서 관리",
		pl:   "Zarządzaj w aplikacji",
		ptBR: "Gerenciar no aplicativo",
		ptPT: "Gerir na aplicação",
		ru:   "Управление в приложении",
		tr:   "Uygulamada yönet",
		ukUA: "Керуйте в додатку",
		uz:   "Ilovada boshqarish",
		zhCN: "在应用中管理",
	},
	AddContact: {
		arEG: "إضافة جهة اتصال",
		de:   "Kontakt hinzufügen",
		en:   "Add contact",
		es:   "Agregar contacto",
		faIR: "اضافه کردن مخاطب",
		fr:   "Ajouter un contact",
		id:   "Tambah kontak",
		it:   "Aggiungi contatto",
		jaJP: "連絡先を追加",
		koKR: "연락처 추가",
		pl:   "Dodaj kontakt",
		ptBR: "Adicionar contato",
		ptPT: "Adicionar contacto",
		ru:   "Добавить контакт",
		tr:   "Kişi ekle",
		ukUA: "Додати контакт",
		uz:   "Kontakt qo&#39;shing",
		zhCN: "添加联系人",
	},
	AddMember: {
		arEG: "إضافة عضو",
		de:   "Mitglied hinzufügen",
		en:   "Add member",
		es:   "Agregar miembro",
		faIR: "اضافه کردن عضو",
		fr:   "Ajouter un membre",
		id:   "Tambah anggota",
		it:   "Aggiungi membro",
		jaJP: "メンバーを追加",
		koKR: "구성원 추가",
		pl:   "Dodaj członka",
		ptBR: "Adicionar membro",
		ptPT: "Adicionar membro",
		ru:   "Добавить участника",
		tr:   "Üye ekle",
		ukUA: "Додати учасника",
		uz:   "A'zo qo'shish",
		zhCN: "添加成员",
	},
	MyContact: {
		arEG: "جهة اتصالي",
		de:   "Mein Kontakt",
		en:   "My contact",
		es:   "Mi contacto",
		faIR: "مخاطب من",
		fr:   "Mon contact",
		id:   "Kontak saya",
		it:   "Il mio contatto",
		jaJP: "私の連絡先",
		koKR: "내 연락처",
		pl:   "Mój kontakt",
		ptBR: "Meu contato",
		ptPT: "O meu contacto",
		ru:   "Мой контакт",
		tr:   "Benim kişim",
		ukUA: "Мій контакт",
		uz:   "Mening kontaktim",
		zhCN: "我的联系人",
	},
	GenderTitled: {
		arEG: "جنس",
		de:   "Geschlecht",
		en:   "Gender",
		es:   "Género",
		faIR: "جنسیت",
		fr:   "Genre",
		id:   "Jenis Kelamin",
		it:   "Genere",
		jaJP: "性別",
		koKR: "성별",
		pl:   "Płeć",
		ptBR: "Gênero",
		ptPT: "Género",
		ru:   "Пол",
		tr:   "Cinsiyet",
		ukUA: "Стать",
		uz:   "Jins",
		zhCN: "性别",
	},
	MaleTitled: {
		arEG: "ذكر",
		de:   "Männlich",
		en:   "Male",
		es:   "Masculino",
		faIR: "مرد",
		fr:   "Homme",
		id:   "Pria",
		it:   "Maschio",
		jaJP: "男性",
		koKR: "남성",
		pl:   "Mężczyzna",
		ptBR: "Masculino",
		ptPT: "Masculino",
		ru:   "Мужской",
		tr:   "Erkek",
		ukUA: "Чоловік",
		uz:   "Erkak",
		zhCN: "男性",
	},
	FemaleTitled: {
		arEG: "أنثى",
		de:   "Weiblich",
		en:   "Female",
		es:   "Femenino",
		faIR: "زن",
		fr:   "Femme",
		id:   "Wanita",
		it:   "Femmina",
		jaJP: "女性",
		koKR: "여성",
		pl:   "Kobieta",
		ptBR: "Feminino",
		ptPT: "Feminino",
		ru:   "Женский",
		tr:   "Kadın",
		ukUA: "Жінка",
		uz:   "Ayol",
		zhCN: "女性",
	},
	UnknownTitled: {
		arEG: "مجهول",
		de:   "Unbekannt",
		en:   "Unknown",
		es:   "Desconocido",
		faIR: "ناشناس",
		fr:   "Inconnu",
		id:   "Tidak Diketahui",
		it:   "Sconosciuto",
		jaJP: "不明",
		koKR: "알 수 없음",
		pl:   "Nieznany",
		ptBR: "Desconhecido",
		ptPT: "Desconhecido",
		ru:   "Неизвестно",
		tr:   "Bilinmiyor",
		ukUA: "Невідомо",
		uz:   "Noma&#39;lum",
		zhCN: "未知",
	},
	UndisclosedTitled: {
		arEG: "غير معلن",
		de:   "Nicht offengelegt",
		en:   "Undisclosed",
		es:   "No revelado",
		faIR: "فاش نشده",
		fr:   "Non divulgué",
		id:   "Tidak Diungkapkan",
		it:   "Non divulgato",
		jaJP: "未公開",
		koKR: "공개되지 않음",
		pl:   "Nieujawnione",
		ptBR: "Não divulgado",
		ptPT: "Não divulgado",
		ru:   "Не разглашено",
		tr:   "Açıklanmamış",
		ukUA: "Не розголошено",
		uz:   "Ochilmagan",
		zhCN: "未披露",
	},
	InviteToJoinSpace: {
		arEG: "دعوة للانضمام",
		de:   "Einladen beizutreten",
		en:   "Invite to join",
		es:   "Invitar a unirse",
		faIR: "دعوت به پیوستن",
		fr:   "Inviter à rejoindre",
		id:   "Undang untuk bergabung",
		it:   "Invita a unirti",
		jaJP: "参加に招待",
		koKR: "가입 초대",
		pl:   "Zaproś do dołączenia",
		ptBR: "Convidar para participar",
		ptPT: "Convidar para juntar-se",
		ru:   "Пригласить присоедениться",
		tr:   "Katılmaya davet et",
		ukUA: "Запросити приєднатися",
		uz:   "Qo&#39;shilishga taklif qiling",
		zhCN: "邀请加入",
	},
	BackToContacts: {
		arEG: "العودة إلى جهات الاتصال",
		de:   "Zurück zu Kontakten",
		en:   "Back to contacts",
		es:   "Volver a los contactos",
		faIR: "بازگشت به مخاطبین",
		fr:   "Retour aux contacts",
		id:   "Kembali ke kontak",
		it:   "Torna ai contatti",
		jaJP: "連絡先に戻る",
		koKR: "연락처로 돌아가기",
		pl:   "Powrót do kontaktów",
		ptBR: "Voltar para os contatos",
		ptPT: "Voltar aos contatos",
		ru:   "Вернуться к контактам",
		tr:   "Kişilere geri dön",
		ukUA: "Повернутися до контактів",
		uz:   "Kontaktlarga qaytish",
		zhCN: "返回联系人",
	},
	BackToMembers: {
		arEG: "العودة إلى الأعضاء",
		de:   "Zurück zu Mitgliedern",
		en:   "Back to members",
		es:   "Volver a los miembros",
		faIR: "بازگشت به اعضا",
		fr:   "Retour aux membres",
		id:   "Kembali ke anggota",
		it:   "Torna ai membri",
		jaJP: "メンバーに戻る",
		koKR: "멤버로 돌아가기",
		pl:   "Powrót do członków",
		ptBR: "Voltar para os membros",
		ptPT: "Voltar aos membros",
		ru:   "Вернуться к участникам",
		tr:   "Üyelere geri dön",
		ukUA: "Повернутися до учасників",
		uz:   "A&#39;zolarga qaytish",
		zhCN: "返回成员",
	},
	BackBtnTitle: {
		arEG: "خلف",
		de:   "Zurück",
		en:   "Back",
		es:   "Atrás",
		faIR: "بازگشت",
		fr:   "Retour",
		id:   "Kembali",
		it:   "Indietro",
		jaJP: "戻る",
		koKR: "뒤로",
		pl:   "Wstecz",
		ptBR: "Voltar",
		ptPT: "Voltar",
		ru:   "Назад",
		tr:   "Geri",
		ukUA: "Назад",
		uz:   "Orqaga",
		zhCN: "返回",
	},
	SelectTrackerToAdd: {
		arEG: "حدد متتبعًا لإضافته",
		de:   "Wählen Sie einen Tracker zum Hinzufügen aus",
		en:   "Select a tracker to add",
		es:   "Seleccionar un rastreador para agregar",
		faIR: "یک ردیاب برای افزودن انتخاب کنید",
		fr:   "Sélectionner un tracker à ajouter",
		id:   "Pilih pelacak untuk ditambahkan",
		it:   "Seleziona un tracker da aggiungere",
		jaJP: "追加するトラッカーを選択",
		koKR: "추가할 트래커를 선택하세요",
		pl:   "Wybierz tracker do dodania",
		ptBR: "Selecionar um rastreador para adicionar",
		ptPT: "Selecionar um rastreador para adicionar",
		ru:   "Выберите трекер для добавления",
		tr:   "Eklemek için bir izleyici seçin",
		ukUA: "Виберіть трекер для додавання",
		uz:   "Qo&#39;shish uchun kuzatuvchini tanlang",
		zhCN: "选择要添加的跟踪器",
	},
	SelectCategoryForNewTracker: {
		arEG: "حدد فئة للمتتبع الجديد",
		de:   "Wählen Sie eine Kategorie für den neuen Tracker aus",
		en:   "Select a category for the new tracker",
		es:   "Seleccionar una categoría para el nuevo rastreador",
		faIR: "یک دسته برای ردیاب جدید انتخاب کنید",
		fr:   "Sélectionnez une catégorie pour le nouveau traqueur",
		id:   "Pilih kategori untuk pelacak baru",
		it:   "Seleziona una categoria per il nuovo tracker",
		jaJP: "新しいトラッカーのカテゴリーを選択",
		koKR: "새 트래커의 카테고리를 선택하세요",
		pl:   "Wybierz kategorię dla nowego trackera",
		ptBR: "Selecionar uma categoria para o novo rastreador",
		ptPT: "Selecionar uma categoria para o novo rastreador",
		ru:   "Выберите категорию для нового трекера",
		tr:   "Yeni izleyici için bir kategori seçin",
		ukUA: "Виберіть категорію для нового трекера",
		uz:   "Yangi treker uchun toifani tanlang",
		zhCN: "为新跟踪器选择一个类别",
	},
	BackToLists: {
		arEG: "العودة إلى القوائم",
		de:   "Zurück zu den Listen",
		en:   "Back to lists",
		es:   "Volver a las listas",
		faIR: "بازگشت به لیست\u200cها",
		fr:   "Retour aux listes",
		id:   "Kembali ke daftar-daftar",
		it:   "Torna alle liste",
		jaJP: "リストに戻る",
		koKR: "목록으로 돌아가기",
		pl:   "Powrót do list",
		ptBR: "Voltar para as listas",
		ptPT: "Voltar às listas",
		ru:   "Вернуться к спискам",
		tr:   "Listelere geri dön",
		ukUA: "Повернутися до списків",
		uz:   "Roʻyxatlarga qaytish",
		zhCN: "返回列表",
	},
	BackToList: {
		arEG: "العودة إلى القائمة",
		de:   "Zurück zur Liste",
		en:   "Back to list",
		es:   "Volver a la lista",
		faIR: "بازگشت به لیست",
		fr:   "Retour à la liste",
		id:   "Kembali ke daftar",
		it:   "Torna alla lista",
		jaJP: "リストに戻る",
		koKR: "목록으로 돌아가기",
		pl:   "Powrót do listy",
		ptBR: "Voltar para a lista",
		ptPT: "Voltar à lista",
		ru:   "Вернуться к списку",
		tr:   "Listeye geri dön",
		ukUA: "Повернутися до списку",
		uz:   "Roʻyxatga qaytish",
		zhCN: "返回列表",
	},
	ClearList: {
		arEG: "مسح القائمة",
		de:   "Liste löschen",
		en:   "Clear list",
		es:   "Limpiar lista",
		faIR: "پاک کردن لیست",
		fr:   "Effacer la liste",
		id:   "Bersihkan daftar",
		it:   "Svuota lista",
		jaJP: "リストをクリア",
		koKR: "목록 지우기",
		pl:   "Wyczyść listę",
		ptBR: "Limpar lista",
		ptPT: "Limpar lista",
		ru:   "Очистить список",
		tr:   "Listeyi temizle",
		ukUA: "Очистити список",
		uz:   "Roʻyxatni tozalash",
		zhCN: "清空列表",
	},
	MarkAsDone: {
		arEG: "وضع علامة على أنه تم",
		de:   "Als erledigt markieren",
		en:   "Mark as done",
		es:   "Marcar como hecho",
		faIR: "علامت\u200cگذاری به عنوان انجام\u200cشده",
		fr:   "Marquer comme fait",
		id:   "Tandai sebagai selesai",
		it:   "Segna come completato",
		jaJP: "完了としてマークする",
		koKR: "완료로 표시",
		pl:   "Oznacz jako zrobione",
		ptBR: "Marcar como concluído",
		ptPT: "Marcar como feito",
		ru:   "Отметить как выполненное",
		tr:   "Tamamlandı olarak işaretle",
		ukUA: "Позначити як виконане",
		uz:   "Bajarildi deb belgilang",
		zhCN: "标记为完成",
	},
	AddStandard: {
		arEG: "إضافة معيار",
		de:   "Standard hinzufügen",
		en:   "Add standard",
		es:   "Agregar estándar",
		faIR: "افزودن استاندارد",
		fr:   "Ajouter une norme",
		id:   "Tambahkan standar",
		it:   "Aggiungere standard",
		jaJP: "標準を追加",
		koKR: "표준 추가",
		pl:   "Dodaj standard",
		ptBR: "Adicionar padrão",
		ptPT: "Adicionar padrão",
		ru:   "Добавить стандартный",
		tr:   "Standart ekle",
		ukUA: "Додати стандарт",
		uz:   "Standart qo&#39;shing",
		zhCN: "添加标准",
	},
	DeleteItems: {
		arEG: "حذف العناصر",
		de:   "Elemente löschen",
		en:   "Delete items",
		es:   "Eliminar elementos",
		faIR: "حذف موارد",
		fr:   "Supprimer les éléments",
		id:   "Hapus item",
		it:   "Elimina elementi",
		jaJP: "項目を削除",
		koKR: "항목 삭제",
		pl:   "Usuń elementy",
		ptBR: "Excluir itens",
		ptPT: "Eliminar itens",
		ru:   "Удалить элементы",
		tr:   "Öğeleri sil",
		ukUA: "Видалити елементи",
		uz:   "Elementlarni o&#39;chirish",
		zhCN: "删除项目",
	},
	YouCanAddItemBySendingMessage: {
		arEG: "يمكنك إضافة العناصر إلى هذه القائمة عن طريق إرسال رسالة لي.",
		de:   "Sie können Elemente zu dieser Liste hinzufügen, indem Sie mir eine Nachricht senden.",
		en:   "You can add items to this list by sending a message to me.",
		es:   "Puedes agregar elementos a esta lista enviándome un mensaje.",
		faIR: "می\u200cتوانید با ارسال یک پیام به من، موارد را به این لیست اضافه کنید.",
		fr:   "Vous pouvez ajouter des éléments à cette liste en m'envoyant un message.",
		id:   "Anda dapat menambahkan item ke daftar ini dengan mengirimkan pesan kepada saya.",
		it:   "Puoi aggiungere elementi a questa lista inviandomi un messaggio.",
		jaJP: "私にメッセージを送ることで、このリストに項目を追加できます。",
		koKR: "저에게 메시지를 보내서 이 목록에 항목들을 추가할 수 있습니다.",
		pl:   "Możesz dodać elementy do tej listy, wysyłając mi wiadomość.",
		ptBR: "Você pode adicionar itens a esta lista enviando uma mensagem para mim.",
		ptPT: "Pode adicionar items a esta lista enviando-me uma mensagem.",
		ru:   "Вы можете добавлять элементы в этот список, отправив мне сообщение.",
		tr:   "Bu listeye bir mesaj göndererek öğeler ekleyebilirsiniz.",
		ukUA: "Ви можете додати елементи до цього списку, надіславши мені повідомлення.",
		uz:   "Menga xabar yuborish orqali ushbu ro&#39;yxatga narsalarni qo&#39;shishingiz mumkin.",
		zhCN: "您可以通过向我发送消息将项目添加到此列表。",
	},
	NoItemsInTheListYet: {
		arEG: "لا يوجد عناصر في القائمة حتى الآن.",
		de:   "Noch keine Elemente in der Liste.",
		en:   "No items in the list yet.",
		es:   "Todavía no hay elementos en la lista.",
		faIR: "هنوز هیچ آیتمی در لیست وجود ندارد.",
		fr:   "Pas encore d'éléments dans la liste.",
		id:   "Belum ada item dalam daftar.",
		it:   "Non ci sono ancora elementi nella lista.",
		jaJP: "リストに項目がまだありません。",
		koKR: "목록에 아직 항목이 없습니다.",
		pl:   "Jeszcze brak elementów na liście.",
		ptBR: "Ainda não há itens na lista.",
		ptPT: "Ainda não existem itens na lista.",
		ru:   "Пока нет элементов в списке.",
		tr:   "Henüz listede öğe yok.",
		ukUA: "Ще немає елементів у списку.",
		uz:   "Roʻyxatda hali biror narsa yoʻq.",
		zhCN: "列表中尚无项目。",
	},
	FamilyList: {
		arEG: "قائمة العائلة",
		de:   "Familienliste",
		en:   "Family list",
		es:   "Lista familiar",
		faIR: "لیست خانواده",
		fr:   "Liste familiale",
		id:   "Daftar keluarga",
		it:   "Elenco familiare",
		jaJP: "家族リスト",
		koKR: "가족 목록",
		pl:   "Lista rodzinna",
		ptBR: "Lista da família",
		ptPT: "Lista familiar",
		ru:   "Семейный список",
		tr:   "Aile listesi",
		ukUA: "Сімейний список",
		uz:   "Oila ro&#39;yxati",
		zhCN: "家庭清单",
	},
	PrivateList: {
		arEG: "قائمة خاصة",
		de:   "Private Liste",
		en:   "Private list",
		es:   "Lista privada",
		faIR: "لیست خصوصی",
		fr:   "Liste privée",
		id:   "Daftar pribadi",
		it:   "Lista personale",
		jaJP: "プライベートリスト",
		koKR: "개인 목록",
		pl:   "Lista prywatna",
		ptBR: "Lista privada",
		ptPT: "Lista privada",
		ru:   "Личный список",
		tr:   "Özel liste",
		ukUA: "Особистий список",
		uz:   "Shaxsiy ro&#39;yxat",
		zhCN: "私人清单",
	},
	Refresh: {
		arEG: "ينعش",
		de:   "Aktualisieren",
		en:   "Refresh",
		es:   "Refrescar",
		faIR: "بروزرسانی",
		fr:   "Rafraîchir",
		id:   "Segarkan",
		it:   "Aggiorna",
		jaJP: "リフレッシュ",
		koKR: "새로고침",
		pl:   "Odśwież",
		ptBR: "Atualizar",
		ptPT: "Atualizar",
		ru:   "Обновить",
		tr:   "Yenile",
		ukUA: "Оновити",
		uz:   "Yangilash",
		zhCN: "刷新",
	},
	AdviseToUseTelegramForTgUsers: {
		arEG: `إذا كان الشخص الذي تريد إضافته يستخدم Telegram، فننصحك بتحديد &quot;اختيار مستخدم Telegram&quot;. وإلا، فيمكنك إضافته يدويًا في Sneat.app.`,
		de: `Wenn die Person, die Sie hinzufügen möchten, Telegram verwendet, empfehlen wir, "Telegram-Benutzer auswählen" auszuwählen.
		Andernfalls können Sie sie manuell in Sneat.app hinzufügen.`,
		en: `If the person you want to add uses telegram we advise to select "Choose Telegram User".
Otherwise you can add them manually in Sneat.app.`,
		es: `Si la persona que deseas agregar usa Telegram, te recomendamos seleccionar "Elegir usuario de Telegram".
		De lo contrario, puedes agregarlos manualmente en Sneat.app.`,
		faIR: `اگر شخصی که می\u200cخواهید اضافه کنید از تلگرام استفاده می\u200cکند، 
ما توصیه می\u200cکنیم گزینه "انتخاب کاربر تلگرام" را انتخاب کنید. 
در غیر این صورت می\u200cتوانید آنها را به صورت دستی در Sneat.app اضافه کنید.`,
		fr: `Si la personne que vous souhaitez ajouter utilise Telegram, nous vous conseillons de sélectionner "Choisir un utilisateur Telegram".
		Sinon, vous pouvez les ajouter manuellement sur Sneat.app.`,
		id: `Jika orang yang ingin Anda tambahkan menggunakan Telegram, kami sarankan untuk memilih "Pilih Pengguna Telegram".
		Jika tidak, Anda dapat menambahkannya secara manual di Sneat.app.`,
		it: `Se la persona che vuoi aggiungere utilizza telegram, ti consigliamo di selezionare "Scegli utente Telegram".
		In caso contrario, puoi aggiungerli manualmente in Sneat.app.`,
		jaJP: `追加したい相手がTelegramを使用している場合、「Telegram ユーザーを選択」を選択することをお勧めします。
		それ以外の場合は、Sneat.app で手動で追加することができます。`,
		koKR: `추가하려는 사람이 Telegram을 사용한다면 "Telegram 사용자 선택"을 선택하시길 권장드립니다.
		그렇지 않을 경우 Sneat.app에서 수동으로 추가할 수 있습니다.`,
		pl: `Jeśli osoba, którą chcesz dodać, korzysta z Telegrama, radzimy wybrać "Wybierz użytkownika Telegrama".
		W przeciwnym razie możesz dodać ich ręcznie w aplikacji Sneat.app.`,
		ptBR: `Se a pessoa que você deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente no Sneat.app.`,
		ptPT: `Se a pessoa que deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente na Sneat.app.`,
		ru: `Если человек, которого вы хотите добавить, использует Telegram,
		мы рекомендуем выбрать "Выбрать пользователя Telegram". 
		В противном случае вы можете добавить их вручную в Sneat.app.`,
		tr: `Eklemek istediğiniz kişi Telegram kullanıyorsa, "Telegram Kullanıcısını Seç" seçeneğini seçmenizi öneririz.
		Aksi takdirde Sneat.app'te manuel olarak ekleyebilirsiniz.`,
		ukUA: `Якщо людина, яку ви хочете додати, використовує Telegram,
		ми радимо вибрати "Вибрати користувача Telegram".
		Інакше, ви можете додати їх вручну на Sneat.app.`,
		uz: `Agar siz qoʻshmoqchi boʻlgan shaxs telegramdan foydalansa, “Telegram foydalanuvchisini tanlash” bandini tanlashni maslahat beramiz.
Aks holda ularni Sneat.app ilovasida qoʻlda qoʻshishingiz mumkin.`,
		zhCN: `如果您想添加的人使用 Telegram，我们建议选择“选择 Telegram 用户”。
		否则，您可以在 Sneat.app 中手动添加他们。`,
	},
	ChooseTelegramUser: {
		arEG: "اختر مستخدم تيليجرام",
		de:   "Telegram-Benutzer auswählen",
		en:   "Choose Telegram user",
		es:   "Elegir usuario de Telegram",
		faIR: "انتخاب کاربر تلگرام",
		fr:   "Choisir un utilisateur Telegram",
		id:   "Pilih Pengguna Telegram",
		it:   "Scegli utente Telegram",
		jaJP: "Telegram ユーザーを選択",
		koKR: "Telegram 사용자 선택",
		pl:   "Wybierz użytkownika Telegrama",
		ptBR: "Escolher usuário do Telegram",
		ptPT: "Escolher usuário do Telegram",
		ru:   "Выбрать пользователя Telegram",
		tr:   "Telegram Kullanıcısını Seç",
		ukUA: "Вибрати користувача Telegram",
		uz:   "Telegram foydalanuvchisini tanlang",
		zhCN: "选择 Telegram 用户",
	},
	AddManuallyInSneatApp: {
		arEG: "أضف يدويًا إلى Sneat.app",
		de:   "Manuell in Sneat.app hinzufügen",
		en:   "Add manually in Sneat.app",
		es:   "Agregar manualmente en Sneat.app",
		faIR: "به صورت دستی در Sneat.app اضافه کنید",
		fr:   "Ajouter manuellement dans Sneat.app",
		id:   "Tambahkan secara manual di Sneat.app",
		it:   "Aggiungi manualmente in Sneat.app",
		jaJP: "Sneat.app で手動で追加",
		koKR: "Sneat.app에서 수동으로 추가",
		pl:   "Dodaj ręcznie w Sneat.app",
		ptBR: "Adicionar manualmente no Sneat.app",
		ptPT: "Adicionar manualmente no Sneat.app",
		ru:   "Добавить вручную в Sneat.app",
		tr:   "Sneat.app'te manuel olarak ekle",
		ukUA: "Додати вручну на Sneat.app",
		uz:   "Sneat.app-da qo&#39;lda qo&#39;shing",
		zhCN: "在 Sneat.app 中手动添加",
	},
	CancelAddingMember: {
		arEG: "إلغاء إضافة العضو",
		de:   "Hinzufügen eines Mitglieds abbrechen",
		en:   "Cancel adding member",
		es:   "Cancelar la adición de miembro",
		faIR: "لغو افزودن عضو",
		fr:   "Annuler l'ajout d'un membre",
		id:   "Batalkan menambahkan anggota",
		it:   "Annulla l'aggiunta di un membro",
		jaJP: "メンバー追加をキャンセル",
		koKR: "멤버 추가 취소",
		pl:   "Anuluj dodawanie członka",
		ptBR: "Cancelar a adição de membro",
		ptPT: "Cancelar a adição de membro",
		ru:   "Отменить добавление участника",
		tr:   "Üye eklemeyi iptal et",
		ukUA: "Скасувати додавання учасника",
		uz:   "Aʼzo qoʻshishni bekor qilish",
		zhCN: "取消添加成员",
	},
	CancelAddingContact: {
		arEG: "إلغاء إضافة جهة اتصال",
		de:   "Hinzufügen eines Kontakts abbrechen",
		en:   "Cancel adding contact",
		es:   "Cancelar la adición de contacto",
		faIR: "لغو افزودن مخاطب",
		fr:   "Annuler l'ajout d'un contact",
		id:   "Batalkan menambahkan kontak",
		it:   "Annulla l'aggiunta di un contatto",
		jaJP: "連絡先追加をキャンセル",
		koKR: "연락처 추가 취소",
		pl:   "Anuluj dodawanie kontaktu",
		ptBR: "Cancelar a adição de contato",
		ptPT: "Cancelar a adição de contato",
		ru:   "Отменить добавление контакта",
		tr:   "Kişi eklemeyi iptal et",
		ukUA: "Скасувати додавання контакту",
		uz:   "Kontakt qo‘shishni bekor qiling",
		zhCN: "取消添加联系人",
	},
	FamilyDebts: {
		arEG: "ديون العائلة",
		de:   "Familienschulden",
		en:   "Family debts",
		es:   "Deudas familiares",
		faIR: "بدهی های خانواده",
		fr:   "Dettes familiales",
		id:   "Hutang keluarga",
		it:   "Debiti familiari",
		jaJP: "家族の借金",
		koKR: "가족 부채",
		pl:   "Długi rodzinne",
		ptBR: "Dívidas familiares",
		ptPT: "Dívidas da família",
		ru:   "Семейные долги",
		tr:   "Aile borçları",
		ukUA: "Сімейні борги",
		uz:   "Oilaviy qarzlar",
		zhCN: "家庭债务",
	},
	DebtsRelatedContacts: {
		arEG: "جهات الاتصال المتعلقة بالديون",
		de:   "Schuldenbezogene Kontakte",
		en:   "Debts related contacts",
		es:   "Contactos relacionados con deudas",
		faIR: "ارتباطات مربوط به بدهی\u200cها",
		fr:   "Contacts liés aux dettes",
		id:   "Kontak yang terkait dengan hutang",
		it:   "Contatti legati ai debiti",
		jaJP: "借金関連の連絡先",
		koKR: "부채 관련 연락처",
		pl:   "Kontakty związane z długami",
		ptBR: "Contatos relacionados a dívidas",
		ptPT: "Contactos relacionados com dívidas",
		ru:   "Контакты, связанные с долгами",
		tr:   "Borçlarla ilgili kişiler",
		ukUA: "Контакти, пов’язані з боргами",
		uz:   "Qarz bilan bog&#39;liq aloqalar",
		zhCN: "债务相关联系人",
	},
	BackToDebtsMenu: {
		arEG: "العودة إلى قائمة الديون",
		de:   "Zurück zum Schuldenmenü",
		en:   "Back to debts menu",
		es:   "Volver al menú de deudas",
		faIR: "بازگشت به منوی بدهی\u200cها",
		fr:   "Retour au menu des dettes",
		id:   "Kembali ke menu hutang",
		it:   "Torna al menu dei debiti",
		jaJP: "借金メニューに戻る",
		koKR: "부채 메뉴로 돌아가기",
		pl:   "Powrót do menu długów",
		ptBR: "Voltar ao menu de dívidas",
		ptPT: "Voltar ao menu das dívidas",
		ru:   "Вернуться в меню долгов",
		tr:   "Borçlar menüsüne geri dön",
		ukUA: "Повернутися до меню боргів",
		uz:   "Qarz menyusiga qaytish",
		zhCN: "返回债务菜单",
	},
	NewEventTitle: {
		arEG: "حدث جديد",
		de:   "Neue Veranstaltung",
		en:   "New Event",
		es:   "Nuevo Evento",
		faIR: "رویداد جدید",
		fr:   "Nouvel Événement",
		id:   "Acara Baru",
		it:   "Nuovo Evento",
		jaJP: "新しいイベント",
		koKR: "새 이벤트",
		pl:   "Nowe Wydarzenie",
		ptBR: "Novo Evento",
		ptPT: "Novo Evento",
		ru:   "Новое Событие",
		tr:   "Yeni Etkinlik",
		ukUA: "Нова Подія",
		uz:   "Yangi Tadbir",
		zhCN: "新事件",
	},
	NewEventText: {
		arEG: "يساعد @{BOT_CODE} في اختيار التاريخ والوقت والمكان المناسب لجميع المشاركين.",
		de:   "@{BOT_CODE} hilft dabei, Datum, Uhrzeit und Ort zu wählen, die für alle Teilnehmer am besten funktionieren.",
		en:   "@{BOT_CODE} helps to choose date, time & place that works best for all participants.",
		es:   "@{BOT_CODE} ayuda a elegir fecha, hora y lugar que funcione mejor para todos los participantes.",
		faIR: "@{BOT_CODE} به انتخاب تاریخ، زمان و مکانی که برای همه شرکت\u200Cکنندگان مناسب است کمک می\u200Cکند.",
		fr:   "@{BOT_CODE} aide à choisir la date, l'heure et le lieu qui conviennent le mieux à tous les participants.",
		id:   "@{BOT_CODE} membantu memilih tanggal, waktu & tempat yang paling cocok untuk semua peserta.",
		it:   "@{BOT_CODE} aiuta a scegliere data, ora e luogo che funzionano meglio per tutti i partecipanti.",
		jaJP: "@{BOT_CODE} は、すべての参加者に最適な日付、時間、場所を選ぶのに役立ちます。",
		koKR: "@{BOT_CODE}는 모든 참가자에게 가장 적합한 날짜, 시간 및 장소를 선택하는 데 도움을 줍니다.",
		pl:   "@{BOT_CODE} pomaga wybrać datę, godzinę i miejsce, które najlepiej pasują wszystkim uczestników.",
		ptBR: "@{BOT_CODE} ajuda a escolher data, hora e local que funcionam melhor para todos os participantes.",
		ptPT: "@{BOT_CODE} ajuda a escolher a data, a hora e o local que melhor funcionam para todos os participantes.",
		ru:   "@{BOT_CODE} помогает выбрать дату, время и место, которые лучше всего подходят для всех участников.",
		tr:   "@{BOT_CODE} tüm katılımcılar için en uygun tarih, saat ve yeri seçmeye yardımcı olur.",
		ukUA: "@{BOT_CODE} допомагає обрати дату, час і місце, які найкраще підходять для всіх учасників.",
		uz:   "@{BOT_CODE} barcha ishtirokchilar uchun eng mos sana, vaqt va joyni tanlashda yordam beradi.",
		zhCN: "@{BOT_CODE} 帮助选择最适合所有参与者的日期、时间和地点。",
	},
	NewEventHint: {
		arEG: "أدخل عنوان الحدث الجديد الخاص بك:",
		de:   "Geben Sie den Titel Ihrer neuen Veranstaltung ein:",
		en:   "Enter title of your new event:",
		es:   "Ingrese el título de su nuevo evento:",
		faIR: "عنوان رویداد جدید خود را وارد کنید:",
		fr:   "Entrez le titre de votre nouvel événement:",
		id:   "Masukkan judul acara baru Anda:",
		it:   "Inserisci il titolo del tuo nuovo evento:",
		jaJP: "新しいイベントのタイトルを入力してください:",
		koKR: "새 이벤트의 제목을 입력하세요:",
		pl:   "Wprowadź tytuł swojego nowego wydarzenia:",
		ptBR: "Digite o título do seu novo evento:",
		ptPT: "Introduza o título do seu novo evento:",
		ru:   "Введите название вашего нового события:",
		tr:   "Yeni etkinliğinizin başlığını girin:",
		ukUA: "Введіть назву вашої нової події:",
		uz:   "Yangi tadbiringizning sarlavhasini kiriting:",
		zhCN: "输入您的新事件标题:",
	},
	TodayButtonText: {
		arEG: "🕒 اليوم — {DATE}",
		de:   "🕒 Heute — {DATE}",
		en:   "🕒 Today — {DATE}",
		es:   "🕒 Hoy — {DATE}",
		faIR: "🕒 امروز — {DATE}",
		fr:   "🕒 Aujourd'hui — {DATE}",
		id:   "🕒 Hari ini — {DATE}",
		it:   "🕒 Oggi — {DATE}",
		jaJP: "🕒 今日 — {DATE}",
		koKR: "🕒 오늘 — {DATE}",
		pl:   "🕒 Dzisiaj — {DATE}",
		ptBR: "🕒 Hoje — {DATE}",
		ptPT: "🕒 Hoje — {DATE}",
		ru:   "🕒 Сегодня — {DATE}",
		tr:   "🕒 Bugün — {DATE}",
		ukUA: "🕒 Сьогодні — {DATE}",
		uz:   "🕒 Bugun — {DATE}",
		zhCN: "🕒 今天 — {DATE}",
	},
	TomorrowButtonText: {
		arEG: "🌅 غدًا — {DATE}",
		de:   "🌅 Morgen — {DATE}",
		en:   "🌅 Tomorrow —  {DATE}",
		es:   "🌅 Mañana — {DATE}",
		faIR: "🌅 فردا — {DATE}",
		fr:   "🌅 Demain — {DATE}",
		id:   "🌅 Besok — {DATE}",
		it:   "🌅 Domani — {DATE}",
		jaJP: "🌅 明日 — {DATE}",
		koKR: "🌅 내일 — {DATE}",
		pl:   "🌅 Jutro — {DATE}",
		ptBR: "🌅 Amanhã — {DATE}",
		ptPT: "🌅 Amanhã — {DATE}",
		ru:   "🌅 Завтра — {DATE}",
		tr:   "🌅 Yarın — {DATE}",
		ukUA: "🌅 Завтра — {DATE}",
		uz:   "🌅 Ertaga — {DATE}",
		zhCN: "🌅 明天 — {DATE}",
	},
	SpotGoingToDoActivities: {
		arEG: "سأفعل: {ACTIVITIES}",
		de:   "Vorhaben: {ACTIVITIES}",
		en:   "Going to do: {ACTIVITIES}",
		es:   "Voy a hacer: {ACTIVITIES}",
		faIR: "قصد انجام: {ACTIVITIES}",
		fr:   "Va faire: {ACTIVITIES}",
		id:   "Akan melakukan: {ACTIVITIES}",
		it:   "Intenzione di fare: {ACTIVITIES}",
		jaJP: "する予定: {ACTIVITIES}",
		koKR: "할 일: {ACTIVITIES}",
		pl:   "Zamierzam zrobić: {ACTIVITIES}",
		ptBR: "Vai fazer: {ACTIVITIES}",
		ptPT: "Vai fazer: {ACTIVITIES}",
		ru:   "Собираюсь делать: {ACTIVITIES}",
		tr:   "Yapacağım: {ACTIVITIES}",
		ukUA: "Збираюся робити: {ACTIVITIES}",
		uz:   "Qilmoqchi: {ACTIVITIES}",
		zhCN: "打算做: {ACTIVITIES}",
	},
	ChooseSpotToRSVP: {
		arEG: "اختر مكانًا للرد",
		de:   "Wählen Sie einen Platz zum Zusagen",
		en:   "Choose a spot to RSVP",
		es:   "Elige un lugar para confirmar asistencia",
		faIR: "مکانی برای تایید حضور انتخاب کنید",
		fr:   "Choisissez un lieu pour confirmer votre présence",
		id:   "Pilih tempat untuk konfirmasi kehadiran",
		it:   "Scegli un posto per confermare la presenza",
		jaJP: "出席返事をする場所を選択してください",
		koKR: "참석 응답할 장소를 선택하세요",
		pl:   "Wybierz miejsce, aby potwierdzić obecność",
		ptBR: "Escolha um local para confirmar presença",
		ptPT: "Escolha um local para confirmar a presença",
		ru:   "Выберите место для подтверждения участия",
		tr:   "Katılımı onaylamak için bir yer seçin",
		ukUA: "Оберіть місце для підтвердження участі",
		uz:   "Ishtirok etishni tasdiqlash uchun joy tanlang",
		zhCN: "选择一个地点来回复邀请",
	},
	TogdIntentPublished: {
		arEG: "لقد قمت بنشر نيتك بنجاح.",
		de:   "Sie haben Ihre Absicht erfolgreich veröffentlicht.",
		en:   "You've successfully published your intention.",
		es:   "Has publicado tu intención exitosamente.",
		faIR: "شما با موفقیت قصد خود را منتشر کردید.",
		fr:   "Vous avez publié votre intention avec succès.",
		id:   "Anda telah berhasil menerbitkan niat Anda.",
		it:   "Hai pubblicato con successo la tua intenzione.",
		jaJP: "意図を正常に公開しました。",
		koKR: "의도를 성공적으로 게시했습니다.",
		pl:   "Pomyślnie opublikowałeś swoją intencję.",
		ptBR: "Você publicou sua intenção com sucesso.",
		ptPT: "Publicou a sua intenção com sucesso.",
		ru:   "Вы успешно опубликовали свое намерение.",
		tr:   "Niyetinizi başarıyla yayınladınız.",
		ukUA: "Ви успішно опублікували свій намір.",
		uz:   "Siz o'z niyatingizni muvaffaqiyatli e'lon qildingiz.",
		zhCN: "您已成功发布您的意图。",
	},
	TogdBackToActivities: {
		arEG: "🔙 العودة إلى الأنشطة",
		de:   "🔙 Zurück zu Aktivitäten",
		en:   "🔙 Back to AllActivityCodes",
		es:   "🔙 Volver a Actividades",
		faIR: "🔙 بازگشت به فعالیت\u200cها",
		fr:   "🔙 Retour aux Activités",
		id:   "🔙 Kembali ke Aktivitas",
		it:   "🔙 Torna alle Attività",
		jaJP: "🔙 アクティビティに戻る",
		koKR: "🔙 활동으로 돌아가기",
		pl:   "🔙 Powrót do Aktywności",
		ptBR: "🔙 Voltar às Atividades",
		ptPT: "🔙 Voltar às atividades",
		ru:   "🔙 Назад к Активностям",
		tr:   "🔙 Aktivitelere Dön",
		ukUA: "🔙 Назад до Активностей",
		uz:   "🔙 Faoliyatlarga qaytish",
		zhCN: "🔙 返回活动",
	},
	TogdPlansButtonText: {
		arEG: "📝 الخطط",
		de:   "📝 Pläne",
		en:   "📝 Plans",
		es:   "📝 Planes",
		faIR: "📝 برنامه\u200cها",
		fr:   "📝 Plans",
		id:   "📝 Rencana",
		it:   "📝 Piani",
		jaJP: "📝 プラン",
		koKR: "📝 계획",
		pl:   "📝 Plany",
		ptBR: "📝 Planos",
		ptPT: "📝 Planos",
		ru:   "📝 Планы",
		tr:   "📝 Planlar",
		ukUA: "📝 Плани",
		uz:   "📝 Rejalar",
		zhCN: "📝 计划",
	},
	TogdSpotsButtonText: {
		arEG: "📍 أماكن",
		de:   "📍 Orte",
		en:   "📍 Spots",
		es:   "📍 Lugares",
		faIR: "📍 مکان\u200cها",
		fr:   "📍 Lieux",
		id:   "📍 Tempat",
		it:   "📍 Luoghi",
		jaJP: "📍 スポット",
		koKR: "📍 장소",
		pl:   "📍 Miejsca",
		ptBR: "📍 Locais",
		ptPT: "📍 Locais",
		ru:   "📍 Места",
		tr:   "📍 Yerler",
		ukUA: "📍 Місця",
		uz:   "📍 Joylar",
		zhCN: "📍 地点",
	},
	RsvpQuestionOnWhatDate: {
		arEG: "في أي يوم سوف تحضر؟",
		de:   "An welchem Tag werden Sie teilnehmen?",
		en:   "On what day are you going to attend?",
		es:   "¿Qué día vas a asistir?",
		faIR: "چه روزی قصد شرکت دارید؟",
		fr:   "Quel jour allez-vous assister ?",
		id:   "Pada hari apa Anda akan hadir?",
		it:   "In che giorno parteciperai?",
		jaJP: "何日に参加する予定ですか？",
		koKR: "어느 날에 참석할 예정인가요?",
		pl:   "W którym dniu zamierzasz uczestniczyć?",
		ptBR: "Em que dia você vai comparecer?",
		ptPT: "Em que dia vai comparecer?",
		ru:   "В какой день вы собираетесь присутствовать?",
		tr:   "Hangi gün katılacaksınız?",
		ukUA: "У який день ви збираєтеся відвідати?",
		uz:   "Qaysi kuni qatnashmoqchisiz?",
		zhCN: "您打算哪一天参加？",
	},
	RsvpQuestionAtWhatTime: {
		arEG: "متى سوف تصل؟",
		de:   "Um wie viel Uhr werden Sie ankommen?",
		en:   "At what time are you going to arrive?",
		es:   "¿A qué hora vas a llegar?",
		faIR: "در چه ساعتی خواهید رسید؟",
		fr:   "À quelle heure allez-vous arriver ?",
		id:   "Pada jam berapa Anda akan tiba?",
		it:   "A che ora arriverai?",
		jaJP: "何時に到着予定ですか？",
		koKR: "몇 시에 도착할 예정인가요?",
		pl:   "O której godzinie zamierzasz przyjść?",
		ptBR: "A que horas você vai chegar?",
		ptPT: "A que horas vai chegar?",
		ru:   "Во сколько вы собираетесь прибыть?",
		tr:   "Saat kaçta geleceksiniz?",
		ukUA: "О котрій годині ви збираєтеся прибути?",
		uz:   "Soat nechada kelasiz?",
		zhCN: "您几点到达？",
	},
	RsvpQuestionWhatActivity: {
		en: "What you going to do?",
		ru: "Что вы собираетесь делать?",
	},
	RsvpTimeIsChangeable: {
		arEG: "سيكون بإمكانك تغيير الدقائق إذا لزم الأمر لاحقًا.",
		de:   "Sie können die Minuten bei Bedarf später ändern.",
		en:   "You would be able to change minutes if needed later.",
		es:   "Podrás cambiar los minutos si es necesario más tarde.",
		faIR: "در صورت نیاز بعداً می\u200cتوانید دقایق را تغییر دهید.",
		fr:   "Vous pourrez modifier les minutes plus tard si nécessaire.",
		id:   "Anda dapat mengubah menit jika diperlukan nanti.",
		it:   "Potrai cambiare i minuti se necessario più tardi.",
		jaJP: "必要に応じて後で分数を変更できます。",
		koKR: "필요하면 나중에 분을 변경할 수 있습니다.",
		pl:   "W razie potrzeby będziesz mógł później zmienić minuty.",
		ptBR: "Você poderá alterar os minutos se necessário mais tarde.",
		ptPT: "Poderá alterar os minutos mais tarde, se necessário.",
		ru:   "При необходимости вы сможете изменить минуты позже.",
		tr:   "Gerekirse daha sonra dakikaları değiştirebileceksiniz.",
		ukUA: "За потреби ви зможете змінити хвилини пізніше.",
		uz:   "Kerak bo'lsa keyinroq daqiqalarni o'zgartirishingiz mumkin.",
		zhCN: "如果需要，您稍后可以更改分钟数。",
	},
	RsvpResponse100Percent: {
		arEG: "سأكون هناك 💯٪",
		de:   "Ich werde da sein 💯%",
		en:   "I'll be there 💯%",
		es:   "Estaré allí 💯%",
		faIR: "💯% آنجا خواهم بود",
		fr:   "Je serai là 💯%",
		id:   "Saya akan ada di sana 💯%",
		it:   "Ci sarò 💯%",
		jaJP: "💯% そこにいます",
		koKR: "💯% 거기에 있을게요",
		pl:   "Będę tam 💯%",
		ptBR: "Estarei lá 💯%",
		ptPT: "Estarei lá 💯%",
		ru:   "Я буду там 💯%",
		tr:   "Orada olacağım 💯%",
		ukUA: "Я буду там 💯%",
		uz:   "Men u yerda bo'laman 💯%",
		zhCN: "我会在那里 💯%",
	},
	RsvpResponseNotAttending: {
		arEG: "عدم الحضور",
		de:   "Nicht teilnehmend",
		en:   "Not attending",
		es:   "No asistiré",
		faIR: "شرکت نمی\u200cکنم",
		fr:   "Ne participe pas",
		id:   "Tidak hadir",
		it:   "Non partecipo",
		jaJP: "参加しません",
		koKR: "참석하지 않음",
		pl:   "Nie biorę udziału",
		ptBR: "Não vou participar",
		ptPT: "Não comparecendo",
		ru:   "Не участвую",
		tr:   "Katılmıyorum",
		ukUA: "Не беру участь",
		uz:   "Qatnashmayman",
		zhCN: "不参加",
	},
	RsvpResponseMostLikely: {
		arEG: "على الأرجح",
		de:   "Höchstwahrscheinlich",
		en:   "Most likely",
		es:   "Muy probable",
		faIR: "خیلی محتمل",
		fr:   "Très probablement",
		id:   "Kemungkinan besar",
		it:   "Molto probabilmente",
		jaJP: "おそらく",
		koKR: "아마도",
		pl:   "Najprawdopodobniej",
		ptBR: "Muito provável",
		ptPT: "Provavelmente",
		ru:   "Скорее всего",
		tr:   "Büyük ihtimalle",
		ukUA: "Найімовірніше",
		uz:   "Katta ehtimol bilan",
		zhCN: "很可能",
	},
	RsvpResponseMaybe: {
		arEG: "ربما",
		de:   "Vielleicht",
		en:   "Maybe",
		es:   "Tal vez",
		faIR: "شاید",
		fr:   "Peut-être",
		id:   "Mungkin",
		it:   "Forse",
		jaJP: "たぶん",
		koKR: "아마",
		pl:   "Może",
		ptBR: "Talvez",
		ptPT: "Talvez",
		ru:   "Возможно",
		tr:   "Belki",
		ukUA: "Можливо",
		uz:   "Balki",
		zhCN: "也许",
	},
	RsvpResponseUnlikely: {
		arEG: "من غير المحتمل",
		de:   "Unwahrscheinlich",
		en:   "Unlikely",
		es:   "Poco probable",
		faIR: "بعید",
		fr:   "Peu probable",
		id:   "Tidak mungkin",
		it:   "Improbabile",
		jaJP: "ありそうにない",
		koKR: "가능성 낮음",
		pl:   "Mało prawdopodobne",
		ptBR: "Improvável",
		ptPT: "Improvável",
		ru:   "Маловероятно",
		tr:   "Pek olası değil",
		ukUA: "Малоймовірно",
		uz:   "Kam ehtimol",
		zhCN: "不太可能",
	},
	RsvpHowLikelyQuestion: {
		arEG: "ما مدى احتمالية تواجدك هناك؟",
		de:   "Wie wahrscheinlich ist es, dass Sie dort sein werden?",
		en:   "How likely is it you are going to be there?",
		es:   "¿Qué tan probable es que vayas a estar allí?",
		faIR: "چقدر احتمال دارد که آنجا باشید؟",
		fr:   "Quelle est la probabilité que vous y soyez ?",
		id:   "Seberapa besar kemungkinan Anda akan ada di sana?",
		it:   "Quanto è probabile che tu ci sia?",
		jaJP: "そこにいる可能性はどのくらいですか？",
		koKR: "거기에 있을 가능성이 얼마나 됩니까?",
		pl:   "Jak prawdopodobne jest, że tam będziesz?",
		ptBR: "Qual a probabilidade de você estar lá?",
		ptPT: "Qual a probabilidade de lá estar?",
		ru:   "Насколько вероятно, что вы будете там?",
		tr:   "Orada olma olasılığınız ne kadar?",
		ukUA: "Наскільки ймовірно, що ви будете там?",
		uz:   "U yerda bo'lish ehtimoli qanchalik?",
		zhCN: "您去那里的可能性有多大？",
	},
	EventTileWithLabelAndEmoji: {
		en: "Event: {EMOJI} {TITLE}",
		ru: "Мероприятие: {EMOJI} {TITLE}",
	},
	SpotTileWithLabelAndEmoji: {
		arEG: "الموقع: {EMOJI} {TITLE}",
		de:   "Stelle: {EMOJI} {TITLE}",
		en:   "Spot: {EMOJI} {TITLE}",
		es:   "Punto: {EMOJI} {TITLE}",
		faIR: "نقطه: {EMOJI} {TITLE}",
		fr:   "Spot : {EMOJI} {TITLE}",
		id:   "Titik: {EMOJI} {TITLE}",
		it:   "Punto: {EMOJI} {TITLE}",
		jaJP: "スポット: {EMOJI} {TITLE}",
		koKR: "스팟: {EMOJI} {TITLE}",
		pl:   "Miejsce: {EMOJI} {TITLE}",
		ptBR: "Local: {EMOJI} {TITLE}",
		ptPT: "Localização: {EMOJI} {TITLE}",
		ru:   "Место: {EMOJI} {TITLE}",
		tr:   "Nokta: {EMOJI} {TITLE}",
		ukUA: "Місце: {EMOJI} {TITLE}",
		uz:   "Joy: {EMOJI} {TITLE}",
		zhCN: "现货：{EMOJI} {TITLE}"},
	SpotTitleWithLocation: {
		arEG: "الموقع: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		de:   "Ort: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		en:   "Spot: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		es:   "Lugar: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		faIR: "مکان: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		fr:   "Lieu: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		id:   "Tempat: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		it:   "Posto: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		jaJP: "スポット: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		koKR: "장소: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		pl:   "Miejsce: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		ptBR: "Local: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		ptPT: "Localização: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		ru:   "Место: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		tr:   "Nokta: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		ukUA: "Місце: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		uz:   "Joy: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		zhCN: "地点: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
	},
	ActivitiesTitle: {
		arEG: "أنشطة",
		de:   "Aktivitäten",
		en:   "AllActivityCodes",
		es:   "Actividades",
		faIR: "فعالیت\u200cها",
		fr:   "Activités",
		id:   "Aktivitas",
		it:   "Attività",
		jaJP: "アクティビティ",
		koKR: "활동",
		pl:   "Aktywności",
		ptBR: "Atividades",
		ptPT: "Atividades",
		ru:   "Активности",
		tr:   "Etkinlikler",
		ukUA: "Активності",
		uz:   "Faoliyatlar",
		zhCN: "活动",
	},
	SpotButtonText: {
		arEG: "الموقع: {SPOT_TITLE}",
		de:   "Platz: {SPOT_TITLE}",
		en:   "Spot: {SPOT_TITLE}",
		es:   "Lugar: {SPOT_TITLE}",
		faIR: "مکان: {SPOT_TITLE}",
		fr:   "Lieu: {SPOT_TITLE}",
		id:   "Tempat: {SPOT_TITLE}",
		it:   "Posto: {SPOT_TITLE}",
		jaJP: "スポット: {SPOT_TITLE}",
		koKR: "장소: {SPOT_TITLE}",
		pl:   "Miejsce: {SPOT_TITLE}",
		ptBR: "Local: {SPOT_TITLE}",
		ptPT: "Localização: {SPOT_TITLE}",
		ru:   "Место: {SPOT_TITLE}",
		tr:   "Yer: {SPOT_TITLE}",
		ukUA: "Місце: {SPOT_TITLE}",
		uz:   "Joy: {SPOT_TITLE}",
		zhCN: "地点: {SPOT_TITLE}",
	},
	ShareActivityButtonText: {
		arEG: "📤 شارك المكان",
		de:   "📤 Spot teilen",
		en:   "📤 Share spot",
		es:   "📤 Comparte lugar",
		faIR: "📤 نقطه اشتراک گذاری",
		fr:   "📤 Partager l&#39;espace",
		id:   "📤 Bagikan tempat",
		it:   "📤 Condividi il posto",
		jaJP: "📤 スポットを共有する",
		koKR: "📤 장소 공유",
		pl:   "📤 Udostępnij miejsce",
		ptBR: "📤 Compartilhe o lugar",
		ptPT: "📤 Partilhe o lugar",
		ru:   "📤 Поделиться активностью",
		tr:   "📤 Paylaşım noktası",
		ukUA: "📤 Поділитися місцем",
		uz:   "📤 Joyni baham ko&#39;ring",
		zhCN: "📤 分享地点"},
	ShareEventButtonText: {
		en: "📤 Share spot",
		ru: "📤 Пригласить",
	},
	ShareSpotButtonText: {
		arEG: "📤 شارك المكان",
		de:   "📤 Ort teilen",
		en:   "📤 Share spot",
		es:   "📤 Compartir lugar",
		faIR: "📤 اشتراک\u200cگذاری مکان",
		fr:   "📤 Partager lieu",
		id:   "📤 Bagikan tempat",
		it:   "📤 Condividi posto",
		jaJP: "📤 スポットを共有",
		koKR: "📤 장소 공유",
		pl:   "📤 Udostępnij miejsce",
		ptBR: "📤 Compartilhar local",
		ptPT: "📤 Partilhe o lugar",
		ru:   "📤 Поделиться местом",
		tr:   "📤 Noktayı paylaş",
		ukUA: "📤 Поділитися місцем",
		uz:   "📤 Joyni baham ko'rish",
		zhCN: "📤 分享地点",
	},
	ShareUserProfileButtonText: {
		arEG: "📤مشاركة ملف تعريف المستخدم",
		de:   "📤 Benutzerprofil teilen",
		en:   "📤 Share user profile",
		es:   "📤 Compartir perfil de usuario",
		faIR: "📤 اشتراک‌گذاری پروفایل کاربری",
		fr:   "📤 Partager le profil utilisateur",
		id:   "📤 Bagikan profil pengguna",
		it:   "📤 Condividi il profilo utente",
		jaJP: "📤 ユーザープロフィールを共有する",
		koKR: "📤 사용자 프로필 공유",
		pl:   "📤 Udostępnij profil użytkownika",
		ptBR: "📤 Compartilhe o perfil do usuário",
		ptPT: "📤 Partilhar o perfil do utilizador",
		ru:   "📤 Поделиться профилем",
		tr:   "📤 Kullanıcı profilini paylaş",
		ukUA: "📤 Поділитися профілем користувача",
		uz:   "📤 Foydalanuvchi profilini ulashing",
		zhCN: "📤 分享用户资料"},
	SeeFullUserProfileInlineButtonText: {
		arEG: "👤🔍 شاهد الملف الشخصي الكامل للمستخدم",
		de:   "👤🔍 Vollständiges Benutzerprofil anzeigen",
		en:   "👤🔍 See full user profile",
		es:   "👤🔍 Ver perfil de usuario completo",
		faIR: "👤🔍 مشاهده پروفایل کامل کاربر",
		fr:   "👤🔍 Voir le profil utilisateur complet",
		id:   "👤🔍 Lihat profil pengguna lengkap",
		it:   "👤🔍 Vedi il profilo utente completo",
		jaJP: "👤🔍 完全なユーザープロフィールを見る",
		koKR: "👤🔍 전체 사용자 프로필 보기",
		pl:   "👤🔍 Zobacz pełny profil użytkownika",
		ptBR: "👤🔍 Veja o perfil completo do usuário",
		ptPT: "👤🔍 Ver perfil completo do utilizador",
		ru:   "👤🔍 Полный профиль пользователя",
		tr:   "👤🔍 Tam kullanıcı profilini görün",
		ukUA: "👤🔍 Дивитися повний профіль користувача",
		uz:   "👤🔍 Toʻliq foydalanuvchi profilini koʻring",
		zhCN: "👤🔍 查看完整用户资料"},
	UserProfileInlineArticleDescription: {
		arEG: "{LANGUAGE}: النقر على هذا سيؤدي إلى مشاركة ملف تعريف المستخدم في الدردشة المحددة.",
		de:   "{LANGUAGE}: Wenn Sie hierauf klicken, wird das Benutzerprofil mit dem ausgewählten Chat geteilt.",
		en:   "{LANGUAGE}: Clicking to this will share user's profile to the selected chat.",
		es:   "{LANGUAGE}: Al hacer clic aquí, se compartirá el perfil del usuario en el chat seleccionado.",
		faIR: "{LANGUAGE}: با کلیک روی این گزینه، پروفایل کاربر در گفتگوی انتخاب شده به اشتراک گذاشته می‌شود.",
		fr:   "{LANGUAGE} : Cliquer ici partagera le profil de l&#39;utilisateur avec le chat sélectionné.",
		id:   "{LANGUAGE}: Mengklik ini akan membagikan profil pengguna ke obrolan yang dipilih.",
		it:   "{LANGUAGE}: Facendo clic qui, il profilo dell&#39;utente verrà condiviso nella chat selezionata.",
		jaJP: "{LANGUAGE}: これをクリックすると、ユーザーのプロフィールが選択したチャットに共有されます。",
		koKR: "{LANGUAGE}: 이것을 클릭하면 사용자 프로필이 선택된 채팅에 공유됩니다.",
		pl:   "{LANGUAGE}: Kliknięcie tego przycisku spowoduje udostępnienie profilu użytkownika wybranemu czacie.",
		ptBR: "{LANGUAGE}: Clicar aqui compartilhará o perfil do usuário no chat selecionado.",
		ptPT: "{LANGUAGE}: Clicar aqui irá partilhar o perfil do utilizador no chat selecionado.",
		ru:   "{LANGUAGE}: Выберите этот вариант чтобы поделиться профилем пользователя в выбранно чате.",
		tr:   "{LANGUAGE}: Buraya tıklandığında kullanıcının profili seçili sohbette paylaşılacaktır.",
		ukUA: "{LANGUAGE}: Натискання цієї кнопки надішле профіль користувача у вибраний чат.",
		uz:   "{LANGUAGE}: Buni bosish tanlangan chatga foydalanuvchi profilini ulashadi.",
		zhCN: "{LANGUAGE}：单击此处可将用户的个人资料分享到选定的聊天中。"},
	PlanEventButtonText: {
		arEG: "🎯تخطيط الحدث",
		de:   "🎯 Event planen",
		en:   "🎯 Plan event",
		es:   "🎯 Planificar evento",
		faIR: "🎯 برنامه\u200cریزی رویداد",
		fr:   "🎯 Planifier événement",
		id:   "🎯 Rencanakan acara",
		it:   "🎯 Pianifica evento",
		jaJP: "🎯 イベントを計画",
		koKR: "🎯 이벤트 계획",
		pl:   "🎯 Zaplanuj wydarzenie",
		ptBR: "🎯 Planejar evento",
		ptPT: "🎯 Planear evento",
		ru:   "🎯 Запланировать событие",
		tr:   "🎯 Etkinlik planla",
		ukUA: "🎯 Запланувати подію",
		uz:   "🎯 Tadbir rejalashtirish",
		zhCN: "🎯 计划活动",
	},
	RemoveFromSpots: {
		arEG: "💔 إزالة من البقع",
		de:   "💔 Aus Favoriten entfernen",
		en:   "💔 Remove from spots",
		es:   "💔 Quitar de favoritos",
		faIR: "💔 حذف از مکان\u200cهای منتخب",
		fr:   "💔 Supprimer des favoris",
		id:   "💔 Hapus dari tempat favorit",
		it:   "💔 Rimuovi dai preferiti",
		jaJP: "💔 お気に入りから削除",
		koKR: "💔 즐겨찾기에서 제거",
		pl:   "💔 Usuń z ulubionych",
		ptBR: "💔 Remover dos favoritos",
		ptPT: "💔 Remover dos favoritos",
		ru:   "💔 Удалить из избранного",
		tr:   "💔 Favorilerden kaldır",
		ukUA: "💔 Видалити з обраного",
		uz:   "💔 Sevimlilardan olib tashlash",
		zhCN: "💔 从收藏中移除",
	},
	AddToMySpots: {
		arEG: "💛 أضف إلى مواقعي",
		de:   "💛 Zu meinen Favoriten hinzufügen",
		en:   "💛 Add to my spots",
		es:   "💛 Añadir a mis favoritos",
		faIR: "💛 افزودن به مکان\u200cهای من",
		fr:   "💛 Ajouter à mes favoris",
		id:   "💛 Tambah ke tempat favorit saya",
		it:   "💛 Aggiungi ai miei preferiti",
		jaJP: "💛 お気に入りに追加",
		koKR: "💛 내 즐겨찾기에 추가",
		pl:   "💛 Dodaj do moich ulubionych",
		ptBR: "💛 Adicionar aos meus favoritos",
		ptPT: "💛 Adicionar aos meus favoritos",
		ru:   "💛 Добавить в мои места",
		tr:   "💛 Favorilerime ekle",
		ukUA: "💛 Додати до моїх місць",
		uz:   "💛 Mening sevimlilarimga qo'shish",
		zhCN: "💛 添加到我的收藏",
	},
	FollowsButtonText: {
		arEG: "👀 متابعين",
		de:   "👀 Folgt",
		en:   "👀 Follows",
		es:   "👀 Sigue",
		faIR: "👀 دنبال می‌کند",
		fr:   "👀 Abonnés",
		id:   "👀 Mengikuti",
		it:   "👀 Segue",
		jaJP: "👀フォロー",
		koKR: "👀 팔로우",
		pl:   "👀 Obserwuje",
		ptBR: "👀 Segue",
		ptPT: "👀 Segue",
		ru:   "👀 Подписки",
		tr:   "👀 Takip eder",
		ukUA: "👀 Підписники",
		uz:   "👀 Kuzatadi",
		zhCN: "👀 关注"},
	FollowersButtonText: {
		arEG: "🕵️‍♂️ متابعين",
		de:   "🕵️‍♂️ Follower",
		en:   "🕵️‍♂️ Followers",
		es:   "🕵️‍♂️ Seguidores",
		faIR: "🕵️‍♂️ دنبال‌کنندگان",
		fr:   "🕵️‍♂️ Abonnés",
		id:   "🕵️‍♂️ Pengikut",
		it:   "🕵️‍♂️ Seguaci",
		jaJP: "🕵️‍♂️ フォロワー",
		koKR: "🕵️‍♂️ 팔로워",
		pl:   "🕵️‍♂️ Obserwujący",
		ptBR: "🕵️‍♂️ Seguidores",
		ptPT: "🕵️‍♂️ Seguidores",
		ru:   "🕵️‍♂️ Подписчики",
		tr:   "🕵️‍♂️ Takipçiler",
		ukUA: "🕵️‍♂️ Підписники",
		uz:   "🕵️‍♂️ Obunachilar",
		zhCN: "🕵️‍♂️ 关注者"},
	FavoriteSpotsTitle: {
		arEG: "الأماكن المفضلة",
		de:   "Lieblingsorte",
		en:   "Favorite Spots",
		es:   "Lugares favoritos",
		faIR: "نقاط مورد علاقه",
		fr:   "Endroits préférés",
		id:   "Tempat Favorit",
		it:   "Luoghi preferiti",
		jaJP: "お気に入りのスポット",
		koKR: "좋아하는 장소",
		pl:   "Ulubione miejsca",
		ptBR: "Lugares favoritos",
		ptPT: "Lugares favoritos",
		ru:   "Любимые места",
		tr:   "Favori Noktalar",
		ukUA: "Улюблені місця",
		uz:   "Sevimli joylar",
		zhCN: "最喜欢的地方"},
	FavoriteActivitiesTitle: {
		arEG: "الأنشطة المفضلة",
		de:   "Lieblingsbeschäftigungen",
		en:   "Favorite AllActivityCodes",
		es:   "Actividades favoritas",
		faIR: "فعالیت‌های مورد علاقه",
		fr:   "Activités préférées",
		id:   "Kegiatan Favorit",
		it:   "Attività preferite",
		jaJP: "好きなアクティビティ",
		koKR: "좋아하는 활동",
		pl:   "Ulubione zajęcia",
		ptBR: "Atividades favoritas",
		ptPT: "Atividades favoritas",
		ru:   "Любимые активности",
		tr:   "Favori Aktiviteler",
		ukUA: "Улюблені заняття",
		uz:   "Sevimli tadbirlar",
		zhCN: "最喜欢的活动"},
	NoFavoriteSpotsYet: {
		arEG: "لم يتم إضافة أي مكان إلى المفضلة بعد.",
		de:   "Es wurden noch keine Spots zu den Favoriten hinzugefügt.",
		en:   "No spots have been added to favorites yet.",
		es:   "Aún no se han añadido lugares a favoritos.",
		faIR: "هنوز هیچ مکانی به علاقه‌مندی‌ها اضافه نشده است.",
		fr:   "Aucun spot n&#39;a encore été ajouté aux favoris.",
		id:   "Belum ada tempat yang ditambahkan ke favorit.",
		it:   "Nessun posto è stato ancora aggiunto ai preferiti.",
		jaJP: "お気に入りにはまだスポットが追加されていません。",
		koKR: "아직 즐겨찾기에 추가된 장소가 없습니다.",
		pl:   "Żadne miejsce nie zostało jeszcze dodane do ulubionych.",
		ptBR: "Nenhum lugar foi adicionado aos favoritos ainda.",
		ptPT: "Ainda não foi adicionado nenhum lugar aos favoritos.",
		ru:   "Пока ещё нет мест добавленных в избранное.",
		tr:   "Henüz favorilere eklenen bir yer yok.",
		ukUA: "До обраного ще не додано жодних місць.",
		uz:   "Sevimlilarga hali hech qanday joy qo‘shilmagan.",
		zhCN: "尚未将任何景点添加至收藏夹。"},
	NoFavoriteActivitiesYet: {
		arEG: "ليس هناك أي نشاط إضافي بشكل واضح.",
		de:   "Es gibt keine aktive Wirkung im Bild.",
		en:   "Пока ещё нет активностей добавленных в избранное.",
		es:   "Poka ещё нет активностей добавленных в избранное.",
		faIR: "Пока ещё нет активностей добавленных во انتخابное.",
		fr:   "Il ne s&#39;agit peut-être pas d&#39;un moyen actif pour l&#39;entreprise.",
		id:   "Anda tidak perlu melakukan tindakan apa pun dalam hal ini.",
		it:   "Non c&#39;è nulla di attivo nell&#39;ambiente.",
		jaJP: "Пока ещё нет активностей добавленных в избранное.",
		koKR: "당신은 избранное에서 добавленных нет активностей.",
		pl:   "Пока ещё нет активностей добавленных в избранное.",
		ptBR: "Não há nenhuma atividade ativa na operação.",
		ptPT: "Não existe qualquer atividade ativa na operação.",
		ru:   "Пока ещё нет активностей добавленных в избранное.",
		tr:   "Ama asla aktif olarak hareket edemiyorum.",
		ukUA: "Пока ще немає активностей, доданих у вибране.",
		uz:   "Poka eschyo net aktivnostey dobavlennyh v izbrannoe.",
		zhCN: "Пока ещё нет активностей добавленных в избранное。"},
	SubscribeFollowUserButtonText: {
		arEG: "🔔 اشترك - تابع المستخدم",
		de:   "🔔 Abonnieren - Benutzer folgen",
		en:   "🔔 Subscribe - follow user",
		es:   "🔔 Suscríbete - sigue al usuario",
		faIR: "🔔 مشترک شوید - کاربر را دنبال کنید",
		fr:   "🔔 S&#39;abonner - suivre l&#39;utilisateur",
		id:   "🔔 Berlangganan - ikuti pengguna",
		it:   "🔔 Iscriviti - segui l&#39;utente",
		jaJP: "🔔 購読 - ユーザーをフォロー",
		koKR: "🔔 구독 - 사용자 팔로우",
		pl:   "🔔 Subskrybuj - obserwuj użytkownika",
		ptBR: "🔔 Inscreva-se - siga o usuário",
		ptPT: "🔔 Subscrever - seguir o utilizador",
		ru:   "🔔 Подписаться на пользователя",
		tr:   "🔔 Abone ol - kullanıcıyı takip et",
		ukUA: "🔔 Підпишіться - слідкуйте за користувачем",
		uz:   "🔔 Obuna bo&#39;ling - foydalanuvchini kuzatib boring",
		zhCN: "🔔 订阅 - 关注用户"},
	UnsubscribeUnfollowUserButtonText: {
		arEG: "إلغاء الاشتراك - إلغاء متابعة المستخدم",
		de:   "Abmelden – Benutzer nicht mehr folgen",
		en:   "Unsubscribe - unfollow user",
		es:   "Darse de baja - dejar de seguir al usuario",
		faIR: "لغو اشتراک - لغو دنبال کردن کاربر",
		fr:   "Se désabonner - ne plus suivre l&#39;utilisateur",
		id:   "Berhenti berlangganan - berhenti mengikuti pengguna",
		it:   "Annulla iscrizione - non seguire più l&#39;utente",
		jaJP: "登録解除 - ユーザーのフォローを解除",
		koKR: "구독 취소 - 사용자 팔로우 취소",
		pl:   "Wypisz się – przestań obserwować użytkownika",
		ptBR: "Cancelar inscrição - deixar de seguir o usuário",
		ptPT: "Cancelar subscrição - deixar de seguir o utilizador",
		ru:   "Отписаться от пользователя",
		tr:   "Aboneliği iptal et - kullanıcıyı takipten çık",
		ukUA: "Відписатися - скасувати підписку на користувача",
		uz:   "Obunani bekor qilish - foydalanuvchini kuzatishni bekor qilish",
		zhCN: "取消订阅 - 取消关注用户"},
	FollowButtonText: {
		arEG: "👀 تابع",
		de:   "👀 Folgen",
		en:   "👀 Follow",
		es:   "👀 Seguir",
		faIR: "👀 دنبال کردن",
		fr:   "👀 Suivre",
		id:   "👀 Ikuti",
		it:   "👀 Segui",
		jaJP: "👀 フォロー",
		koKR: "👀 팔로우",
		pl:   "👀 Obserwuj",
		ptBR: "👀 Seguir",
		ptPT: "👀 Seguir",
		ru:   "👀 Подписаться",
		tr:   "👀 Takip et",
		ukUA: "👀 Слідкувати",
		uz:   "👀 Kuzatish",
		zhCN: "👀 关注",
	},
	UnfollowButtonText: {
		arEG: "👀 إلغاء المتابعة",
		de:   "👀 Nicht mehr folgen",
		en:   "👀 Unfollow",
		es:   "👀 Dejar de seguir",
		faIR: "👀 لغو دنبال کردن",
		fr:   "👀 Ne plus suivre",
		id:   "👀 Berhenti mengikuti",
		it:   "👀 Non seguire più",
		jaJP: "👀 フォロー解除",
		koKR: "👀 언팔로우",
		pl:   "👀 Przestań obserwować",
		ptBR: "👀 Deixar de seguir",
		ptPT: "👀 Deixar de seguir",
		ru:   "👀 Отписаться",
		tr:   "👀 Takibi bırak",
		ukUA: "👀 Перестати слідкувати",
		uz:   "👀 Kuzatishni to'xtatish",
		zhCN: "👀 取消关注",
	},
	CanceledEventTitle: {
		en: "Canceled Event",
		ru: "Отменённое Мероприятие",
	},
	EventTitle: {
		arEG: "حدث",
		de:   "Veranstaltung",
		en:   "Event",
		es:   "Evento",
		faIR: "رویداد",
		fr:   "Événement",
		id:   "Acara",
		it:   "Evento",
		jaJP: "イベント",
		koKR: "이벤트",
		pl:   "Wydarzenie",
		ptBR: "Evento",
		ptPT: "Evento",
		ru:   "Мероприятие",
		tr:   "Etkinlik",
		ukUA: "Подія",
		uz:   "Tadbir",
		zhCN: "活动",
	},
	EventCreated: {
		arEG: "تم إنشاء الحدث",
		de:   "Veranstaltung erstellt",
		en:   "Event created",
		es:   "Evento creado",
		faIR: "رویداد ایجاد شد",
		fr:   "Événement créé",
		id:   "Acara dibuat",
		it:   "Evento creato",
		jaJP: "イベントが作成されました",
		koKR: "이벤트 생성됨",
		pl:   "Wydarzenie utworzone",
		ptBR: "Evento criado",
		ptPT: "Evento criado",
		ru:   "Событие создано",
		tr:   "Etkinlik oluşturuldu",
		ukUA: "Подію створено",
		uz:   "Tadbir yaratildi",
		zhCN: "活动已创建",
	},
	CancelEvent: {
		arEG: "إلغاء الحدث",
		de:   "Veranstaltung absagen",
		en:   "Cancel event",
		es:   "Cancelar evento",
		faIR: "لغو رویداد",
		fr:   "Annuler l'événement",
		id:   "Batalkan acara",
		it:   "Annulla evento",
		jaJP: "イベントをキャンセル",
		koKR: "이벤트 취소",
		pl:   "Anuluj wydarzenie",
		ptBR: "Cancelar evento",
		ptPT: "Cancelar evento",
		ru:   "Отменить событие",
		tr:   "Etkinliği iptal et",
		ukUA: "Скасувати подію",
		uz:   "Tadbirni bekor qilish",
		zhCN: "取消活动",
	},
	BackToEvents: {
		arEG: "العودة إلى الأحداث",
		de:   "Zurück zu Veranstaltungen",
		en:   "Back to events",
		es:   "Volver a eventos",
		faIR: "بازگشت به رویدادها",
		fr:   "Retour aux événements",
		id:   "Kembali ke acara",
		it:   "Torna agli eventi",
		jaJP: "イベントに戻る",
		koKR: "이벤트로 돌아가기",
		pl:   "Powrót do wydarzeń",
		ptBR: "Voltar aos eventos",
		ptPT: "Voltar aos eventos",
		ru:   "Назад к событиям",
		tr:   "Etkinliklere dön",
		ukUA: "Назад до подій",
		uz:   "Tadbirlarga qaytish",
		zhCN: "返回活动",
	},
	EventOptionsButton: {
		arEG: "خيارات",
		de:   "Optionen",
		en:   "Options",
		es:   "Opciones",
		faIR: "گزینه\u200cها",
		fr:   "Options",
		id:   "Opsi",
		it:   "Opzioni",
		jaJP: "オプション",
		koKR: "옵션",
		pl:   "Opcje",
		ptBR: "Opções",
		ptPT: "Opções",
		ru:   "Варианты",
		tr:   "Seçenekler",
		ukUA: "Варіанти",
		uz:   "Variantlar",
		zhCN: "选项",
	},
	NewEventOptionButton: {
		arEG: "إضافة خيار",
		de:   "Option hinzufügen",
		en:   "Add option",
		es:   "Añadir opción",
		faIR: "افزودن گزینه",
		fr:   "Ajouter une option",
		id:   "Tambah opsi",
		it:   "Aggiungi opzione",
		jaJP: "オプションを追加",
		koKR: "옵션 추가",
		pl:   "Dodaj opcję",
		ptBR: "Adicionar opção",
		ptPT: "Adicionar opção",
		ru:   "Добавить вариант",
		tr:   "Seçenek ekle",
		ukUA: "Додати варіант",
		uz:   "Variant qo'shish",
		zhCN: "添加选项",
	},
	EventVisibility: {
		arEG: "الرؤية: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		de:   "Sichtbarkeit: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		en:   "Visibility: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		es:   "Visibilidad: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		faIR: "نمایان\u200cبودن: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		fr:   "Visibilité: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		id:   "Visibilitas: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		it:   "Visibilità: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		jaJP: "表示設定: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		koKR: "공개 범위: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		pl:   "Widoczność: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		ptBR: "Visibilidade: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		ptPT: "Visibilidade: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		ru:   "Видимость: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		tr:   "Görünürlük: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		ukUA: "Видимість: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		uz:   "Ko'rinish: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		zhCN: "可见性: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
	},
	EventStatus: {
		arEG: "الحالة: {EVENT_STATUS}",
		de:   "Status: {EVENT_STATUS}",
		en:   "Status: {EVENT_STATUS}",
		es:   "Estado: {EVENT_STATUS}",
		faIR: "وضعیت: {EVENT_STATUS}",
		fr:   "Statut: {EVENT_STATUS}",
		id:   "Status: {EVENT_STATUS}",
		it:   "Stato: {EVENT_STATUS}",
		jaJP: "ステータス: {EVENT_STATUS}",
		koKR: "상태: {EVENT_STATUS}",
		pl:   "Status: {EVENT_STATUS}",
		ptBR: "Status: {EVENT_STATUS}",
		ptPT: "Estado: {EVENT_STATUS}",
		ru:   "Статус: {EVENT_STATUS}",
		tr:   "Durum: {EVENT_STATUS}",
		ukUA: "Статус: {EVENT_STATUS}",
		uz:   "Holat: {EVENT_STATUS}",
		zhCN: "状态: {EVENT_STATUS}",
	},
	TogdMyProfile: {
		arEG: "ملفي الشخصي",
		de:   "Mein Profil",
		en:   "My profile",
		es:   "Mi perfil",
		faIR: "پروفایل من",
		fr:   "Mon profil",
		id:   "Profil saya",
		it:   "Il mio profilo",
		jaJP: "マイプロフィール",
		koKR: "내 프로필",
		pl:   "Mój profil",
		ptBR: "Meu perfil",
		ptPT: "O meu perfil",
		ru:   "Мой профиль",
		tr:   "Benim profilim",
		ukUA: "Мій профіль",
		uz:   "Mening profilim",
		zhCN: "我的资料",
	},
	TogdMyFavoriteSpotsTitle: {
		arEG: "أماكني المفضلة",
		de:   "Meine Lieblingsorte",
		en:   "My Favorite Spots",
		es:   "Mis lugares favoritos",
		faIR: "نقاط مورد علاقه من",
		fr:   "Mes endroits préférés",
		id:   "Tempat Favoritku",
		it:   "I miei posti preferiti",
		jaJP: "私のお気に入りのスポット",
		koKR: "내가 가장 좋아하는 장소",
		pl:   "Moje ulubione miejsca",
		ptBR: "Meus lugares favoritos",
		ptPT: "Os meus lugares favoritos",
		ru:   "Мои избранные места",
		tr:   "En Sevdiğim Yerler",
		ukUA: "Мої улюблені місця",
		uz:   "Mening sevimli joylarim",
		zhCN: "我最喜欢的地方"},
	TogdFavoriteSpotsOfUserTitle: {
		arEG: "الأماكن المفضلة لـ {USER_NAME}",
		de:   "Lieblingsorte von {USER_NAME}",
		en:   "Favorite Spots of {USER_NAME}",
		es:   "Lugares favoritos de {USER_NAME}",
		faIR: "مکان‌های مورد علاقه‌ی {USER_NAME}",
		fr:   "Endroits préférés de {USER_NAME}",
		id:   "Tempat Favorit {USER_NAME}",
		it:   "Luoghi preferiti di {USER_NAME}",
		jaJP: "{USER_NAME}のお気に入りスポット",
		koKR: "{USER_NAME}의 인기 명소",
		pl:   "Ulubione miejsca {USER_NAME}",
		ptBR: "Lugares favoritos de {USER_NAME}",
		ptPT: "Locais favoritos de {USER_NAME}",
		ru:   "Избранные места: {USER_NAME}",
		tr:   "{USER_NAME} Favori Noktaları",
		ukUA: "Улюблені місця міста {USER_NAME}",
		uz:   "{USER_NAME}ning sevimli joylari",
		zhCN: "{USER_NAME} 最喜欢的景点"},
	TogdYouHaveNoFavoriteSpotsYet: {
		arEG: "لم تقم بإضافة أو إنشاء أي أماكن مفضلة بعد.",
		de:   "Sie haben noch keine Lieblingsorte hinzugefügt oder erstellt.",
		en:   "You have no added or created any favorite spots yet.",
		es:   "Aún no has añadido ni creado ningún lugar favorito.",
		faIR: "شما هنوز هیچ مکان مورد علاقه‌ای اضافه یا ایجاد نکرده‌اید.",
		fr:   "Vous n&#39;avez pas encore ajouté ou créé de lieux favoris.",
		id:   "Anda belum menambahkan atau membuat tempat favorit.",
		it:   "Non hai ancora aggiunto o creato alcun posto preferito.",
		jaJP: "お気に入りのスポットはまだ追加または作成されていません。",
		koKR: "아직 추가하거나 만든 즐겨찾는 장소가 없습니다.",
		pl:   "Nie dodałeś ani nie utworzyłeś jeszcze żadnych ulubionych miejsc.",
		ptBR: "Você ainda não adicionou ou criou nenhum lugar favorito.",
		ptPT: "Ainda não adicionou ou criou nenhum lugar favorito.",
		ru:   "Вы ещё не добавили ни одно места проведения активностей в избранное.",
		tr:   "Henüz favori mekan eklemediniz veya oluşturmadınız.",
		ukUA: "Ви ще не додали та не створили жодних улюблених місць.",
		uz:   "Siz hali hech qanday sevimli joy qo‘shmagansiz yoki yaratmagansiz.",
		zhCN: "您尚未添加或创建任何喜爱的地点。"},
	TogdMySpots: {
		arEG: "بقعتي",
		de:   "Meine Orte",
		en:   "My spots",
		es:   "Mis lugares",
		faIR: "مکان\u200cهای من",
		fr:   "Mes endroits",
		id:   "Tempat saya",
		it:   "I miei posti",
		jaJP: "マイスポット",
		koKR: "내 장소들",
		pl:   "Moje miejsca",
		ptBR: "Meus locais",
		ptPT: "As minhas manchas",
		ru:   "Мои места",
		tr:   "Benim yerlerim",
		ukUA: "Мої місця",
		uz:   "Mening joylarim",
		zhCN: "我的地点",
	},

	TogdMyEvents: {
		arEG: "أحداثي",
		de:   "Meine Events",
		en:   "My events",
		es:   "Mis eventos",
		faIR: "رویدادهای من",
		fr:   "Mes événements",
		id:   "Acara saya",
		it:   "I miei eventi",
		jaJP: "マイイベント",
		koKR: "내 이벤트들",
		pl:   "Moje wydarzenia",
		ptBR: "Meus eventos",
		ptPT: "Os meus eventos",
		ru:   "Мои события",
		tr:   "Benim etkinliklerim",
		ukUA: "Мої події",
		uz:   "Mening tadbirlarim",
		zhCN: "我的活动",
	},
	TogdMyActivity: {
		arEG: "نشاطي",
		de:   "Meine Aktivität",
		en:   "My activity",
		es:   "Mi actividad",
		faIR: "فعالیت من",
		fr:   "Mon activité",
		id:   "Aktivitas saya",
		it:   "La mia attività",
		jaJP: "私の活動",
		koKR: "내 활동",
		pl:   "Moja aktywność",
		ptBR: "Minha atividade",
		ptPT: "A minha atividade",
		ru:   "Моя активность",
		tr:   "Benim etkinliğim",
		ukUA: "Моя активність",
		uz:   "Mening faoliyatim",
		zhCN: "我的活动"},
	TogdMyActivities: {
		arEG: "أنشطتي",
		de:   "Meine Aktivitäten",
		en:   "My activities",
		es:   "Mis actividades",
		faIR: "فعالیت\u200cهای من",
		fr:   "Mes activités",
		id:   "Aktivitas saya",
		it:   "Le mie attività",
		jaJP: "マイアクティビティ",
		koKR: "내 활동들",
		pl:   "Moje aktywności",
		ptBR: "Minhas atividades",
		ptPT: "As minhas atividades",
		ru:   "Мои активности",
		tr:   "Benim aktivitelerim",
		ukUA: "Мої активності",
		uz:   "Mening faoliyatlarim",
		zhCN: "我的活动",
	},
	TogdMyPlans: {
		arEG: "خططي",
		de:   "Meine Pläne",
		en:   "My plans",
		es:   "Mis planes",
		faIR: "برنامه\u200cهای من",
		fr:   "Mes projets",
		id:   "Rencana saya",
		it:   "I miei piani",
		jaJP: "マイプラン",
		koKR: "내 계획들",
		pl:   "Moje plany",
		ptBR: "Meus planos",
		ptPT: "Meus planos",
		ru:   "Мои планы",
		tr:   "Benim planlarım",
		ukUA: "Мої плани",
		uz:   "Mening rejalarim",
		zhCN: "我的计划",
	},
	TogdUserProfile: {
		arEG: "ملف تعريف المستخدم",
		de:   "Benutzerprofil",
		en:   "User profile",
		es:   "Perfil de usuario",
		faIR: "پروفایل کاربر",
		fr:   "Profil utilisateur",
		id:   "Profil pengguna",
		it:   "Profilo utente",
		jaJP: "ユーザープロフィール",
		koKR: "사용자 프로필",
		pl:   "Profil użytkownika",
		ptBR: "Perfil do usuário",
		ptPT: "Perfil do utilizador",
		ru:   "Профиль пользователя",
		tr:   "Kullanıcı profili",
		ukUA: "Профіль користувача",
		uz:   "Foydalanuvchi profili",
		zhCN: "用户资料",
	},

	TogdActivitiesOfUser: {
		arEG: "أنشطة {USER_NAME}",
		de:   "Aktivitäten von {USER_NAME}",
		en:   "AllActivityCodes of {USER_NAME}",
		es:   "Actividades de {USER_NAME}",
		faIR: "فعالیت\u200cهای {USER_NAME}",
		fr:   "Activités de {USER_NAME}",
		id:   "Aktivitas {USER_NAME}",
		it:   "Attività di {USER_NAME}",
		jaJP: "{USER_NAME}のアクティビティ",
		koKR: "{USER_NAME}의 활동",
		pl:   "Aktywności użytkownika {USER_NAME}",
		ptBR: "Atividades de {USER_NAME}",
		ptPT: "Atividades de {USER_NAME}",
		ru:   "Активности пользователя {USER_NAME}",
		tr:   "{USER_NAME} etkinlikleri",
		ukUA: "Активності користувача {USER_NAME}",
		uz:   "{USER_NAME} faoliyatlari",
		zhCN: "{USER_NAME}的活动",
	},

	YouHaveNoFavoriteActivities: {
		arEG: "ليس لديك أي أنشطة مفضلة",
		de:   "Sie haben keine Lieblingsaktivitäten",
		en:   "You have no favorite activities",
		es:   "No tienes actividades favoritas",
		faIR: "شما هیچ فعالیت مورد علاقه\u200cای ندارید",
		fr:   "Vous n'avez aucune activité favorite",
		id:   "Anda tidak memiliki aktivitas favorit",
		it:   "Non hai attività preferite",
		jaJP: "お気に入りのアクティビティがありません",
		koKR: "좋아하는 활동이 없습니다",
		pl:   "Nie masz ulubionych aktywności",
		ptBR: "Você não tem atividades favoritas",
		ptPT: "Não tem atividades favoritas",
		ru:   "У вас нет любимых активностей",
		tr:   "Favori etkinliğiniz yok",
		ukUA: "У вас немає улюблених активностей",
		uz:   "Sizda sevimli faoliyatlar yo'q",
		zhCN: "您没有喜欢的活动",
	},

	TogdYouHaveNoActiveEvents: {
		en: "You have no active events.",
		ru: "У вас нет активных событий.",
	},
	TogdUserHasNoActiveEvents: {
		en: "User has no active events.",
		ru: "У пользователя нет активных событий.",
	},
	InstructionHowToAddActivityInBot: {
		arEG: "لإضافة الأنشطة قم بإرسالها مفصولة بفاصلة.",
		de:   "Um Aktivitäten hinzuzufügen, senden Sie sie durch Kommas getrennt.",
		en:   "To add activities send them separated by comma.",
		es:   "Para agregar actividades envíalas separadas por comas.",
		faIR: "برای افزودن فعالیت\u200cها آنها را با کاما جدا کرده ارسال کنید.",
		fr:   "Pour ajouter des activités, envoyez-les séparées par des virgules.",
		id:   "Untuk menambahkan aktivitas, kirim mereka dipisahkan dengan koma.",
		it:   "Per aggiungere attività, inviale separate da virgole.",
		jaJP: "アクティビティを追加するには、カンマで区切って送信してください。",
		koKR: "활동을 추가하려면 쉼표로 구분해서 보내세요.",
		pl:   "Aby dodać aktywności, wyślij je oddzielone przecinkami.",
		ptBR: "Para adicionar atividades, envie-as separadas por vírgulas.",
		ptPT: "Para adicionar atividades envie-as separadas por vírgula.",
		ru:   "Чтобы добавить активности, отправьте их через запятую.",
		tr:   "Etkinlik eklemek için virgülle ayırarak gönderin.",
		ukUA: "Щоб додати активності, надсилайте їх через кому.",
		uz:   "Faoliyatlarni qo'shish uchun ularni vergul bilan ajratib yuboring.",
		zhCN: "要添加活动，请用逗号分隔发送。",
	},
	TogdUserActivity: {
		arEG: "نشاط {USER_NAME}",
		de:   "Aktivität von {USER_NAME}",
		en:   "Activity of {USER_NAME}",
		es:   "Actividad de {USER_NAME}",
		faIR: "فعالیت {USER_NAME}",
		fr:   "Activité de {USER_NAME}",
		id:   "Aktivitas {USER_NAME}",
		it:   "Attività di {USER_NAME}",
		jaJP: "{USER_NAME}のアクティビティ",
		koKR: "{USER_NAME}의 활동",
		pl:   "Aktywność {USER_NAME}",
		ptBR: "Atividade de {USER_NAME}",
		ptPT: "Atividade de {USER_NAME}",
		ru:   "Активность {USER_NAME}",
		tr:   "{USER_NAME} etkinliği",
		ukUA: "Активність користувача {USER_NAME}",
		uz:   "{USER_NAME} faolligi",
		zhCN: "{USER_NAME} 的活动"},
	TogdUserActivities: {
		arEG: "أنشطة {USER_NAME}",
		de:   "Aktivitäten von {USER_NAME}",
		en:   "AllActivityCodes of {USER_NAME}",
		es:   "Actividades del usuario",
		faIR: "فعالیت\u200cهای کاربر",
		fr:   "Activités utilisateur",
		id:   "Aktivitas pengguna",
		it:   "Attività utente",
		jaJP: "ユーザーアクティビティ",
		koKR: "사용자 활동",
		pl:   "Aktywności użytkownika",
		ptBR: "Atividades do usuário",
		ptPT: "Atividades de {USER_NAME}",
		ru:   "Активности пользователя",
		tr:   "Kullanıcı aktiviteleri",
		ukUA: "Активності користувача",
		uz:   "Foydalanuvchi faoliyatlari",
		zhCN: "用户活动",
	},
	TogdUserEvents: {
		arEG: "أحداث المستخدم",
		de:   "Benutzerereignisse",
		en:   "User events",
		es:   "Eventos del usuario",
		faIR: "رویدادهای کاربر",
		fr:   "Événements utilisateur",
		id:   "Acara pengguna",
		it:   "Eventi utente",
		jaJP: "ユーザーイベント",
		koKR: "사용자 이벤트",
		pl:   "Wydarzenia użytkownika",
		ptBR: "Eventos do usuário",
		ptPT: "Eventos do utilizador",
		ru:   "События пользователя",
		tr:   "Kullanıcı etkinlikleri",
		ukUA: "Події користувача",
		uz:   "Foydalanuvchi tadbirlari",
		zhCN: "用户活动",
	},

	TogdUserPlans: {
		arEG: "خطط المستخدم",
		de:   "Benutzerpläne",
		en:   "User plans",
		es:   "Planes del usuario",
		faIR: "برنامه\u200cهای کاربر",
		fr:   "Plans utilisateur",
		id:   "Rencana pengguna",
		it:   "Piani utente",
		jaJP: "ユーザープラン",
		koKR: "사용자 계획",
		pl:   "Plany użytkownika",
		ptBR: "Planos do usuário",
		ptPT: "Planos de utilizador",
		ru:   "Планы пользователя",
		tr:   "Kullanıcı planları",
		ukUA: "Плани користувача",
		uz:   "Foydalanuvchi rejalari",
		zhCN: "用户计划",
	},

	TogdUserSpots: {
		arEG: "نقاط المستخدم",
		de:   "Benutzerorte",
		en:   "User spots",
		es:   "Lugares del usuario",
		faIR: "مکان\u200cهای کاربر",
		fr:   "Lieux utilisateur",
		id:   "Tempat pengguna",
		it:   "Luoghi utente",
		jaJP: "ユーザースポット",
		koKR: "사용자 장소",
		pl:   "Miejsca użytkownika",
		ptBR: "Locais do usuário",
		ptPT: "Pontos de utilizador",
		ru:   "Места пользователя",
		tr:   "Kullanıcı mekanları",
		ukUA: "Місця користувача",
		uz:   "Foydalanuvchi joylari",
		zhCN: "用户地点",
	},
	TogdBotWelcome: {
		arEG: `مرحبًا بك في @ToGetheredBot — مساعدك البسيط والذكي في التخطيط للقاء الأصدقاء وتنظيم الأنشطة الجماعية أو مجرد إخبار الآخرين بمكان تواجدك. سواء كان الأمر يتعلق بالتزلج الشراعي على الشاطئ أو لعب كرة السلة في الشارع أو التخطيط للقاء غير رسمي، فإن ToGethered يجعل التنسيق أمرًا في غاية السهولة.

يوفر البوت ميزتين رئيسيتين:

 1. <b>مشاركة الخطة</b> - أخبر الآخرين <b>بمكان ووقت تواجدك في مكان ما</b> . يمكنك تحديد موقع ونطاق زمني، حتى يتمكن الآخرون من رؤية خططك والانضمام إذا كانوا متاحين أيضًا.

 2. <b>تنسيق الأحداث</b> - نظم الأنشطة الجماعية بسهولة من خلال اقتراح خيارات متعددة للوقت والمكان. يجمع البوت الأصوات من المشاركين ويعرض التركيبات التي تعمل بشكل أفضل، مما يساعد المجموعة على الاتفاق على خطة دون الحاجة إلى سلاسل محادثات طويلة.

مع @ToGetheredBot، يصبح التخطيط اجتماعيًا وواضحًا وسهلًا - مثاليًا للجلسات العفوية أو التجمعات المنظمة.`,
		de: `Willkommen bei @ToGetheredBot — Ihrem einfachen und intelligenten Planungsassistenten für Treffen mit Freunden, die Organisation von Gruppenaktivitäten oder einfach um anderen mitzuteilen, wo Sie sein werden. Ob Kitesurfen am Strand, Streetball spielen oder ein lockeres Treffen planen, ToGethered macht die Koordination mühelos.

Der Bot bietet zwei Hauptfunktionen:

	1.	<b>Plan-Sharing</b> – Lassen Sie andere wissen, <b>wo und wann Sie irgendwo sein möchten</b>. Sie können einen Ort und einen Zeitrahmen angeben, sodass andere Ihre Pläne sehen und mitmachen können, wenn sie auch verfügbar sind.

	2.	<b>Event-Koordination</b> – Organisieren Sie einfach Gruppenaktivitäten, indem Sie mehrere Zeit- und Ortsoptionen vorschlagen. Der Bot sammelt Stimmen von Teilnehmern und zeigt, welche Kombinationen am besten funktionieren, damit sich die Gruppe auf einen Plan einigen kann, ohne lange Chat-Threads.

Mit @ToGetheredBot wird Planung sozial, sichtbar und reibungslos — perfekt für spontane Sessions oder organisierte Zusammenkünfte.`,
		en: `Welcome to @ToGetheredBot — your simple and smart planning assistant for meeting up with friends, organising group activities, or just letting others know where you'll be. Whether it's kitesurfing at the beach, playing street basketball, or planning a casual meetup, ToGethered makes coordination effortless.

The bot offers two main features:

	1.	<b>Plan Sharing</b> – Let others know <b>where and when you're planning to be somewhere</b>. You can specify a location and a time range, so others can see your plans and join in if they're available too.

	2.	<b>Event Coordination</b> – Easily organise group activities by suggesting multiple time and place options. The bot collects votes from participants and shows which combinations work best, helping the group agree on a plan without long chat threads.

With @ToGetheredBot, planning becomes social, visible, and frictionless — perfect for spontaneous sessions or organised gatherings.`,
		es: `Bienvenido a @ToGetheredBot — tu asistente de planificación simple e inteligente para quedar con amigos, organizar actividades grupales o simplemente avisar a otros dónde estarás. Ya sea kitesurf en la playa, jugar baloncesto callejero o planificar una reunión casual, ToGethered hace que la coordinación sea sin esfuerzo.

El bot ofrece dos características principales:

	1.	<b>Compartir Planes</b> – Haz saber a otros <b>dónde y cuándo planeas estar en algún lugar</b>. Puedes especificar una ubicación y un rango de tiempo, para que otros puedan ver tus planes y unirse si también están disponibles.

	2.	<b>Coordinación de Eventos</b> – Organiza fácilmente actividades grupales sugiriendo múltiples opciones de tiempo y lugar. El bot recopila votos de los participantes y muestra qué combinaciones funcionan mejor, ayudando al grupo a acordar un plan sin largos hilos de chat.

Con @ToGetheredBot, la planificación se vuelve social, visible y sin fricciones — perfecto para sesiones espontáneas o reuniones organizadas.`,
		faIR: `به @ToGetheredBot خوش آمدید — دستیار برنامه\u200cریزی ساده و هوشمند شما برای ملاقات با دوستان، سازماندهی فعالیت\u200cهای گروهی، یا فقط اطلاع دادن به دیگران که کجا خواهید بود. چه کایت\u200cسرفینگ در ساحل، بازی بسکتبال خیابانی، یا برنامه\u200cریزی برای یک ملاقات غیررسمی، ToGethered هماهنگی را بدون زحمت می\u200cکند.

ربات دو ویژگی اصلی ارائه می\u200cدهد:

	1.	<b>اشتراک\u200cگذاری برنامه</b> – به دیگران اطلاع دهید <b>کجا و چه زمانی قصد حضور در جایی را دارید</b>. می\u200cتوانید مکان و بازه زمانی مشخص کنید تا دیگران بتوانند برنامه\u200cهای شما را ببینند و در صورت در دسترس بودن به شما بپیوندند.

	2.	<b>هماهنگی رویداد</b> – به راحتی فعالیت\u200cهای گروهی را با پیشنهاد گزینه\u200cهای متعدد زمان و مکان سازماندهی کنید. ربات رای\u200cهای شرکت\u200cکنندگان را جمع\u200cآوری می\u200cکند و نشان می\u200cدهد کدام ترکیبات بهترین عملکرد را دارند، و به گروه کمک می\u200cکند بدون رشته\u200cهای چت طولانی روی یک برنامه توافق کنند.

با @ToGetheredBot، برنامه\u200cریزی اجتماعی، قابل مشاهده و بدون اصطکاک می\u200cشود — عالی برای جلسات خودجوش یا گردهمایی\u200cهای سازمان\u200cیافته.`,
		fr: `Bienvenue sur @ToGetheredBot — votre assistant de planification simple et intelligent pour rencontrer des amis, organiser des activités de groupe, ou simplement faire savoir aux autres où vous serez. Que ce soit du kitesurf à la plage, jouer au basket de rue, ou planifier une rencontre décontractée, ToGethered rend la coordination sans effort.

Le bot offre deux fonctionnalités principales :

	1.	<b>Partage de Plans</b> – Faites savoir aux autres <b>où et quand vous prévoyez d'être quelque part</b>. Vous pouvez spécifier un lieu et une plage horaire, afin que les autres puissent voir vos plans et se joindre s'ils sont également disponibles.

	2.	<b>Coordination d'Événements</b> – Organisez facilement des activités de groupe en suggérant plusieurs options de temps et de lieu. Le bot collecte les votes des participants et montre quelles combinaisons fonctionnent le mieux, aidant le groupe à s'accorder sur un plan sans longs fils de discussion.

Avec @ToGetheredBot, la planification devient sociale, visible et sans friction — parfait pour des sessions spontanées ou des rassemblements organisés.`,
		id: `Selamat datang di @ToGetheredBot — asisten perencanaan sederhana dan cerdas Anda untuk bertemu dengan teman, mengorganisir kegiatan grup, atau sekadar memberi tahu orang lain di mana Anda akan berada. Baik itu kitesurfing di pantai, bermain basket jalanan, atau merencanakan pertemuan santai, ToGethered membuat koordinasi menjadi mudah.

Bot ini menawarkan dua fitur utama:

	1.	<b>Berbagi Rencana</b> – Beri tahu orang lain <b>di mana dan kapan Anda berencana berada di suatu tempat</b>. Anda dapat menentukan lokasi dan rentang waktu, sehingga orang lain dapat melihat rencana Anda dan bergabung jika mereka juga tersedia.

	2.	<b>Koordinasi Acara</b> – Dengan mudah mengorganisir kegiatan grup dengan menyarankan beberapa pilihan waktu dan tempat. Bot mengumpulkan suara dari peserta dan menunjukkan kombinasi mana yang paling berhasil, membantu grup menyepakati rencana tanpa thread chat yang panjang.

Dengan @ToGetheredBot, perencanaan menjadi sosial, terlihat, dan tanpa gesekan — sempurna untuk sesi spontan atau pertemuan yang terorganisir.`,
		it: `Benvenuto su @ToGetheredBot — il tuo assistente di pianificazione semplice e intelligente per incontrare amici, organizzare attività di gruppo, o semplicemente far sapere agli altri dove sarai. Che si tratti di kitesurf in spiaggia, giocare a basket di strada, o pianificare un incontro informale, ToGethered rende il coordinamento senza sforzo.

Il bot offre due funzionalità principali:

	1.	<b>Condivisione Piani</b> – Fai sapere agli altri <b>dove e quando hai intenzione di essere da qualche parte</b>. Puoi specificare una posizione e un intervallo di tempo, così gli altri possono vedere i tuoi piani e unirsi se sono disponibili anche loro.

	2.	<b>Coordinamento Eventi</b> – Organizza facilmente attività di gruppo suggerendo multiple opzioni di tempo e luogo. Il bot raccoglie i voti dai partecipanti e mostra quali combinazioni funzionano meglio, aiutando il gruppo a concordare su un piano senza lunghe discussioni in chat.

Con @ToGetheredBot, la pianificazione diventa sociale, visibile e senza attriti — perfetto per sessioni spontanee o raduni organizzati.`,
		jaJP: `@ToGetheredBotへようこそ — 友達との待ち合わせ、グループ活動の企画、または単に自分がどこにいるかを他の人に知らせるためのシンプルでスマートな計画アシスタントです。ビーチでのカイトサーフィン、ストリートバスケットボール、またはカジュアルな集まりの計画など、ToGetheredは調整を簡単にします。

ボットは2つの主要機能を提供します：

	1.	<b>プラン共有</b> – <b>どこでいつ何かをする予定かを</b>他の人に知らせます。場所と時間範囲を指定できるので、他の人があなたの計画を見て、利用可能であれば参加できます。

	2.	<b>イベント調整</b> – 複数の時間と場所のオプションを提案してグループ活動を簡単に企画します。ボットは参加者からの投票を収集し、どの組み合わせが最適かを表示し、長いチャットスレッドなしでグループが計画に合意するのを助けます。

@ToGetheredBotで、計画はソーシャルで、見える化され、摩擦のないものになります — 自発的なセッションや企画された集まりに最適です。`,
		koKR: `@ToGetheredBot에 오신 것을 환영합니다 — 친구들과의 만남, 그룹 활동 조직, 또는 단순히 다른 사람들에게 당신이 어디에 있을지 알려주는 간단하고 스마트한 계획 도우미입니다. 해변에서의 카이트서핑, 길거리 농구, 또는 캐주얼한 만남 계획 등 무엇이든, ToGethered는 조정을 쉽게 만듭니다.

봇은 두 가지 주요 기능을 제공합니다:

	1.	<b>계획 공유</b> – 다른 사람들에게 <b>언제 어디에 있을 계획인지</b> 알려주세요. 위치와 시간 범위를 지정할 수 있어서 다른 사람들이 당신의 계획을 보고 그들도 가능하다면 참여할 수 있습니다.

	2.	<b>이벤트 조정</b> – 여러 시간과 장소 옵션을 제안하여 그룹 활동을 쉽게 조직하세요. 봇은 참가자들의 투표를 수집하고 어떤 조합이 가장 잘 작동하는지 보여주어, 긴 채팅 스레드 없이 그룹이 계획에 동의할 수 있도록 도와줍니다.

@ToGetheredBot과 함께라면 계획이 사회적이고, 가시적이며, 마찰 없이 이루어집니다 — 자발적인 세션이나 조직된 모임에 완벽합니다.`,
		pl: `Witaj w @ToGetheredBot — Twoim prostym i inteligentnym asystentem planowania spotkań z przyjaciółmi, organizowania działań grupowych lub po prostu informowania innych, gdzie będziesz. Czy to kitesurfing na plaży, granie w koszykówkę uliczną, czy planowanie nieformalnego spotkania, ToGethered sprawia, że koordynacja staje się bezwysiłkowa.

Bot oferuje dwie główne funkcje:

	1.	<b>Udostępnianie Planów</b> – Poinformuj innych <b>gdzie i kiedy planujesz gdzieś być</b>. Możesz określić lokalizację i zakres czasowy, aby inni mogli zobaczyć Twoje plany i dołączyć, jeśli są również dostępni.

	2.	<b>Koordynacja Wydarzeń</b> – Łatwo organizuj działania grupowe, proponując wiele opcji czasu i miejsca. Bot zbiera głosy od uczestników i pokazuje, które kombinacje działają najlepiej, pomagając grupie uzgodnić plan bez długich wątków czatu.

Z @ToGetheredBot planowanie staje się społeczne, widoczne i bezproblemowe — idealne dla spontanicznych sesji lub zorganizowanych spotkań.`,
		ptBR: `Bem-vindo ao @ToGetheredBot — seu assistente de planejamento simples e inteligente para encontrar amigos, organizar atividades em grupo, ou simplesmente avisar outros onde você estará. Seja kitesurf na praia, jogar basquete de rua, ou planejar um encontro casual, ToGethered torna a coordenação sem esforço.

O bot oferece duas funcionalidades principais:

	1.	<b>Compartilhamento de Planos</b> – Deixe outros saberem <b>onde e quando você planeja estar em algum lugar</b>. Você pode especificar uma localização e um intervalo de tempo, para que outros possam ver seus planos e participar se também estiverem disponíveis.

	2.	<b>Coordenação de Eventos</b> – Organize facilmente atividades em grupo sugerindo múltiplas opções de tempo e local. O bot coleta votos dos participantes e mostra quais combinações funcionam melhor, ajudando o grupo a concordar com um plano sem longas discussões no chat.

Com @ToGetheredBot, o planejamento se torna social, visível e sem atritos — perfeito para sessões espontâneas ou encontros organizados.`,
		ptPT: `Bem-vindo ao @ToGetheredBot — o seu assistente de planeamento simples e inteligente para encontrar amigos, organizar atividades de grupo ou simplesmente dizer-lhe onde vai estar. Seja a praticar kitesurf na praia, a jogar basquetebol de rua ou a planear um encontro casual, o ToGethered facilita a coordenação.

O bot oferece duas características principais:

 1. <b>Partilha de planos</b> – Avise outras pessoas <b>onde e quando planeia estar em algum lugar</b> . Pode especificar um local e um intervalo de tempo para que outras pessoas vejam os seus planos e participem, se também estiverem disponíveis.

 2. <b>Coordenação de eventos</b> – Organize facilmente atividades de grupo sugerindo várias opções de horário e local. O bot recolhe os votos dos participantes e mostra quais as combinações que funcionam melhor, ajudando o grupo a concordar com um plano sem conversas longas.

Com o @ToGetheredBot, o planeamento torna-se social, visível e sem atritos — perfeito para sessões espontâneas ou reuniões organizadas.`,
		ru: `Добро пожаловать в @ToGetheredBot — ваш простой и умный помощник для планирования встреч с друзьями, организации групповых мероприятий или просто для того, чтобы сообщить другим, где вы будете. Будь то кайтсёрфинг на пляже, игра в уличный баскетбол или планирование неформальной встречи, ToGethered делает координацию лёгкой.

Бот предлагает две основные функции:

	1.	<b>Обмен планами</b> – Дайте другим знать, <b>где и когда вы планируете где-то быть</b>. Вы можете указать место и временной диапазон, чтобы другие могли видеть ваши планы и присоединиться, если они тоже свободны.

	2.	<b>Координация событий</b> – Легко организуйте групповые мероприятия, предлагая несколько вариантов времени и места. Бот собирает голоса участников и показывает, какие комбинации работают лучше всего, помогая группе договориться о плане без длинных цепочек сообщений.

С @ToGetheredBot планирование становится социальным, видимым и беспрепятственным — идеально подходит для спонтанных сессий или организованных встреч.`,
		tr: `@ToGetheredBot'a hoş geldiniz — arkadaşlarınızla buluşmak, grup etkinlikleri düzenlemek veya sadece başkalarına nerede olacağınızı bildirmek için basit ve akıllı planlama asistanınız. İster sahilde rüzgar sörfü, ister sokak basketbolu oynamak, ister gündelik bir buluşma planlamak olsun, ToGethered koordinasyonu zahmetsiz hale getirir.

Bot iki ana özellik sunar:

	1.	<b>Plan Paylaşımı</b> – Başkalarına <b>nerede ve ne zaman bir yerde olmayı planladığınızı</b> bildirin. Bir konum ve zaman aralığı belirtebilirsiniz, böylece diğerleri planlarınızı görebilir ve onlar da müsaitse katılabilir.

	2.	<b>Etkinlik Koordinasyonu</b> – Birden fazla zaman ve yer seçeneği önererek grup etkinliklerini kolayca düzenleyin. Bot katılımcılardan oyları toplar ve hangi kombinasyonların en iyi çalıştığını gösterir, grubun uzun sohbet zincirleri olmadan bir plan üzerinde anlaşmasına yardımcı olur.

@ToGetheredBot ile planlama sosyal, görünür ve sürtünmesiz hale gelir — spontane oturumlar veya organize toplantılar için mükemmel.`,
		ukUA: `Ласкаво просимо до @ToGetheredBot — вашого простого та розумного помічника для планування зустрічей з друзями, організації групових заходів або просто повідомлення іншим, де ви будете. Чи то кайтсерфінг на пляжі, гра в вуличний баскетбол, чи планування неформальної зустрічі, ToGethered робить координацію легкою.

Бот пропонує дві основні функції:

	1.	<b>Обмін планами</b> – Дайте іншим знати, <b>де і коли ви плануєте десь бути</b>. Ви можете вказати місце та часовий діапазон, щоб інші могли бачити ваші плани та приєднатися, якщо вони теж вільні.

	2.	<b>Координація подій</b> – Легко організовуйте групові заходи, пропонуючи декілька варіантів часу та місця. Бот збирає голоси учасників і показує, які комбінації працюють найкраще, допомагаючи групі домовитися про план без довгих ланцюжків повідомлень.

З @ToGetheredBot планування стає соціальним, видимим та безперешкодним — ідеально підходить для спонтанних сесій або організованих зустрічей.`,
		uz: `@ToGetheredBot'ga xush kelibsiz — do'stlaringiz bilan uchrashish, guruh faoliyatlarini tashkil qilish yoki boshqalarga qayerda bo'lishingizni bildirish uchun oddiy va aqlli rejalashtirish yordamchingiz. Plyajda kite surfing, ko'cha basketboli o'ynash yoki oddiy uchrashuv rejalashtirish bo'ladimi, ToGethered muvofiqlashtirish jarayonini osonlashtiradi.

Bot ikkita asosiy xususiyatni taklif qiladi:

	1.	<b>Reja Almashish</b> – Boshqalarga <b>qayerda va qachon biror joyda bo'lishni rejalashtirayotganingizni</b> bildiring. Joylashuv va vaqt oralig'ini belgilashingiz mumkin, shunda boshqalar rejalaringizni ko'rib, agar ular ham bo'sh bo'lsa qo'shilishlari mumkin.

	2.	<b>Tadbir Muvofiqlashtirish</b> – Bir nechta vaqt va joy variantlarini taklif qilib guruh faoliyatlarini osonlikcha tashkil qiling. Bot ishtirokchilardan ovozlarni to'playdi va qaysi kombinatsiyalar eng yaxshi ishlashini ko'rsatadi, guruhga uzun chat zanjirlarisiz reja bo'yicha kelishishga yordam beradi.

@ToGetheredBot bilan rejalashtirish ijtimoiy, ko'rinadigan va muammosiz bo'ladi — o'z-o'zidan yuzaga keladigan sessiyalar yoki tashkil etilgan uchrashuvlar uchun mukammal.`,
		zhCN: `欢迎使用@ToGetheredBot — 您简单智能的规划助手，用于与朋友见面、组织团体活动，或只是让其他人知道您将在哪里。无论是在海滩风筝冲浪、打街头篮球，还是规划休闲聚会，ToGethered让协调变得轻松。

机器人提供两个主要功能：

	1.	<b>计划分享</b> – 让其他人知道<b>您计划在何时何地出现</b>。您可以指定地点和时间范围，这样其他人就能看到您的计划，如果他们也有空就可以加入。

	2.	<b>活动协调</b> – 通过建议多个时间和地点选项轻松组织团体活动。机器人收集参与者的投票并显示哪些组合效果最好，帮助团队在不需要长聊天记录的情况下达成计划共识。

使用@ToGetheredBot，规划变得社交化、可视化且无摩擦 — 非常适合自发的聚会或有组织的集会。`,
	},
	TogdMainMenuText: {
		en: `
To share your plans choose a spot or activity and hit the <b>{RSVP}</b> button.

To organize an event create a <b>{NEW_EVENT}</b> from "My Events".

To learn more about the bot use /about command.`,

		ru: `
		Чтобы поделиться своими планами, выберите место или активность и нажмите кнопку <b>{RSVP}</b>.

		Чтобы организовать событие, создайте <b>{NEW_EVENT}</b> в разделе "Мои события".

		Чтобы узнать больше используйте команду /about
			`},
	RsvpButtonText: {
		arEG: "من فضلك، أرسل رد",
		de:   "Anmelden",
		en:   "RSVP",
		es:   "Confirmar",
		faIR: "ثبت نام",
		fr:   "S'inscrire",
		id:   "Daftar",
		it:   "Iscriviti",
		jaJP: "参加登録",
		koKR: "등록하기",
		pl:   "Zapisz się",
		ptBR: "Inscrever-se",
		ptPT: "Inscrever-se",
		ru:   "Отметиться",
		tr:   "Kayıt ol",
		ukUA: "Зареєструватися",
		uz:   "Ro'yxatdan o'tish",
		zhCN: "报名",
	},
	NewEventButtonText: {
		arEG: "حدث جديد",
		de:   "Neues Event",
		en:   "New Event",
		es:   "Nuevo evento",
		faIR: "رویداد جدید",
		fr:   "Nouvel événement",
		id:   "Acara baru",
		it:   "Nuovo evento",
		jaJP: "新しいイベント",
		koKR: "새 이벤트",
		pl:   "Nowe wydarzenie",
		ptBR: "Novo evento",
		ptPT: "Novo Evento",
		ru:   "Новое событие",
		tr:   "Yeni etkinlik",
		ukUA: "Нова подія",
		uz:   "Yangi tadbir",
		zhCN: "新活动",
	},
	CalendarButtonText: {
		arEG: "🗓️️ التقويم",
		de:   "🗓️️ Kalender",
		en:   "🗓️️ Calendar",
		es:   "🗓️️ Calendario",
		faIR: "🗓️️ تقویم",
		fr:   "🗓️️ Calendrier",
		id:   "🗓️️ Kalender",
		it:   "🗓️️ Calendario",
		jaJP: "🗓️️ カレンダー",
		koKR: "🗓️️ 달력",
		pl:   "🗓️️ Kalendarz",
		ptBR: "🗓️️ Calendário",
		ptPT: "🗓️️ Calendário",
		ru:   "🗓️️ Календарь",
		tr:   "🗓️️ Takvim",
		ukUA: "🗓️️ Календар",
		uz:   "🗓️️ Taqvim",
		zhCN: "🗓️️ 日历",
	},
	ButtonPrivacy: {
		arEG: "خصوصية",
		de:   "Datenschutz",
		en:   "Privacy",
		es:   "Privacidad",
		faIR: "حریم خصوصی",
		fr:   "Confidentialité",
		id:   "Pribadi",
		it:   "Riservatezza",
		jaJP: "プライバシー",
		koKR: "은둔",
		pl:   "Prywatność",
		ptBR: "Privacidade",
		ptPT: "Privacidade",
		ru:   "Конфиденциальность",
		tr:   "Mahremiyet",
		ukUA: "Конфіденційність",
		uz:   "Maxfiylik",
		zhCN: "隐私",
	},
	ButtonIGotIt: {
		arEG: "حصلت عليه!",
		de:   "Ich habe es!",
		en:   "I got it!",
		es:   "¡Lo tengo!",
		faIR: "گرفتمش!",
		fr:   "J&#39;ai compris!",
		id:   "Saya dapat!",
		it:   "Capito!",
		jaJP: "わかった！",
		koKR: "알겠어요!",
		pl:   "Zrozumiałem!",
		ptBR: "Eu entendi!",
		ptPT: "Eu entendi!",
		ru:   "Я понял!",
		tr:   "Anladım!",
		ukUA: "Я зрозумів!",
		uz:   "Men tushundim!",
		zhCN: "我得到了它！"},
	PrivacyCommandTitle: {
		arEG: "خصوصيتك مهمة بالنسبة لنا",
		de:   "Ihre Privatsphäre ist uns wichtig",
		en:   "Your privacy matters to us",
		es:   "Su privacidad nos importa",
		faIR: "حریم خصوصی شما برای ما مهم است",
		fr:   "Votre vie privée est importante pour nous",
		id:   "Privasi Anda penting bagi kami",
		it:   "La tua privacy è importante per noi",
		jaJP: "あなたのプライバシーは私たちにとって重要です",
		koKR: "귀하의 개인 정보는 저희에게 중요합니다",
		pl:   "Twoja prywatność jest dla nas ważna",
		ptBR: "Sua privacidade é importante para nós",
		ptPT: "A sua privacidade é importante para nós",
		ru:   "Ваша конфиденциальность для нас важна",
		tr:   "Gizliliğiniz bizim için önemlidir",
		ukUA: "Ваша конфіденційність важлива для нас",
		uz:   "Sizning maxfiyligingiz biz uchun muhim",
		zhCN: "我们重视您的隐私"},
	PrivacyCommandText: {
		arEG: `نحن لا نبيع بياناتك الشخصية أبدًا. لا يجوز مشاركة المعلومات إلا بناءً على طلبك المباشر - بموافقتك وتحت سيطرتك. نهدف إلى أن نكون شفافين وموثوقين قدر الإمكان.`,
		de: `Wir verkaufen Ihre personenbezogenen Daten nicht – niemals.
Informationen dürfen nur auf Ihren direkten Wunsch weitergegeben werden – mit Ihrer Zustimmung und unter Ihrer Kontrolle.
Wir möchten so transparent und vertrauenswürdig wie möglich sein.`,
		en: `We don't sell your personal data - ever.
Information may only be shared at your direct request — with your consent and under your control.
We aim to be as transparent and trustworthy as possible.`,
		es: `No vendemos sus datos personales, nunca.
La información solo puede compartirse si la solicita directamente, con su consentimiento y bajo su control.
Nuestro objetivo es ser lo más transparentes y confiables posible.`,
		faIR: `ما هرگز اطلاعات شخصی شما را نمی‌فروشیم.
اطلاعات فقط با درخواست مستقیم شما - با رضایت و تحت کنترل شما - به اشتراک گذاشته می‌شود.
هدف ما این است که تا حد امکان شفاف و قابل اعتماد باشیم.`,
		fr: `Nous ne vendons jamais vos données personnelles. Les informations ne peuvent être partagées qu&#39;à votre demande directe, avec votre consentement et sous votre contrôle. Nous visons à être aussi transparents et dignes de confiance que possible.`,
		id: `Kami tidak akan pernah menjual data pribadi Anda. Informasi hanya boleh dibagikan atas permintaan langsung Anda — dengan persetujuan dan di bawah kendali Anda. Kami berupaya untuk bersikap setransparan dan sedapat mungkin dapat dipercaya.`,
		it: `Non vendiamo mai i tuoi dati personali.
Le informazioni possono essere condivise solo su tua richiesta diretta, con il tuo consenso e sotto il tuo controllo.
Miriamo a essere il più trasparenti e affidabili possibile.`,
		jaJP: `弊社がお客様の個人情報を販売することは一切ありません。
情報は、お客様からの直接のリクエストがあった場合にのみ、お客様の同意と管理のもとで共有されます。
弊社は、可能な限り透明性と信頼性を高めることを目指しています。`,
		koKR: `당사는 귀하의 개인 정보를 절대 판매하지 않습니다.
 귀하의 동의 하에 귀하의 직접적인 요청에 따라서만 정보가 공유될 수 있으며, 귀하의 통제 하에 있습니다.
 당사는 최대한 투명하고 신뢰할 수 있는 서비스를 제공하기 위해 노력합니다.`,
		pl: `Nigdy nie sprzedajemy Twoich danych osobowych.
Informacje mogą być udostępniane wyłącznie na Twoje bezpośrednie żądanie — za Twoją zgodą i pod Twoją kontrolą.
Staramy się być tak transparentni i godni zaufania, jak to tylko możliwe.`,
		ptBR: `Não vendemos seus dados pessoais, nunca. As informações só podem ser compartilhadas mediante solicitação direta sua, com seu consentimento e sob seu controle. Nosso objetivo é ser o mais transparente e confiável possível.`,
		ptPT: `Não vendemos os seus dados pessoais, nunca. As informações só podem ser partilhadas mediante solicitação direta sua, com o seu consentimento e sob o seu controlo. O nosso objetivo é ser o mais transparente e fiável possível.`,
		ru: `Мы никогда не продаём ваши личные данные.
Информация может быть передана только по вашему прямому запросу — с вашего согласия и по вашему выбору.
Мы стараемся быть максимально прозрачными и надёжными.`,
		tr: `Kişisel verilerinizi asla satmayız.
Bilgiler yalnızca doğrudan talebiniz üzerine, sizin izninizle ve sizin kontrolünüz altında paylaşılabilir.
Mümkün olduğunca şeffaf ve güvenilir olmayı hedefliyoruz.`,
		ukUA: `Ми ніколи не продаємо ваші персональні дані.
Інформація може бути передана лише за вашим прямим запитом — з вашої згоди та під вашим контролем.
Ми прагнемо бути максимально прозорими та надійними.`,
		uz: `Biz sizning shaxsiy maʼlumotlaringizni hech qachon sotmaymiz.
Maʼlumotlar faqat sizning toʻgʻridan-toʻgʻri soʻrovingiz boʻyicha – sizning roziligingiz va nazoratingiz ostida boʻlishish mumkin.
Biz imkon qadar shaffof va ishonchli boʻlishni maqsad qilganmiz.`,
		zhCN: `我们永远不会出售您的个人数据。
我们只会根据您的直接要求共享信息 - 并征得您的同意并在您的控制之下。
我们的目标是尽可能透明和值得信赖。`},
	NumberOfFollowers: {
		arEG: "{N} متابع",
		de:   "{N} Follower",
		en:   "{N} followers",
		es:   "{N} seguidores",
		faIR: "{N} دنبال‌کننده",
		fr:   "{N} abonnés",
		id:   "{N} pengikut",
		it:   "{N} follower",
		jaJP: "フォロワー数: {N}",
		koKR: "팔로워 {N}명",
		pl:   "{N} obserwujących",
		ptBR: "{N} seguidores",
		ptPT: "{N} seguidores",
		ru:   "{N} подписчиков",
		tr:   "{N} takipçi",
		ukUA: "{N} підписників",
		uz:   "{N} ta obunachi",
		zhCN: "{N} 位关注者"},
	NumberOfFavorites: {
		arEG: "تمت إضافة {N} مستخدمًا إلى المفضلة",
		de:   "{N} Benutzer zu Favoriten hinzugefügt",
		en:   "{N} users added to favorites",
		es:   "{N} usuarios añadidos a favoritos",
		faIR: "{N} کاربر به علاقه‌مندی‌ها اضافه شدند",
		fr:   "{N} utilisateurs ajoutés aux favoris",
		id:   "{N} pengguna ditambahkan ke favorit",
		it:   "{N} utenti aggiunti ai preferiti",
		jaJP: "{N} 人のユーザーがお気に入りに追加されました",
		koKR: "{N}명의 사용자가 즐겨찾기에 추가되었습니다.",
		pl:   "{N} użytkowników dodanych do ulubionych",
		ptBR: "{N} usuários adicionados aos favoritos",
		ptPT: "{N} utilizadores adicionados aos favoritos",
		ru:   "{N} пользователей добавили в избранное",
		tr:   "{N} kullanıcı favorilere eklendi",
		ukUA: "Користувачів додано до обраного: {N}",
		uz:   "{N} ta foydalanuvchi sevimlilarga qoʻshildi",
		zhCN: "已将 {N} 位用户添加到收藏夹"},
	RsvpLetOthersKnow: {
		arEG: "الرد - أخبر الآخرين",
		de:   "RSVP – andere informieren",
		en:   "RSVP - let others know",
		es:   "RSVP - avisa a los demás",
		faIR: "RSVP - به دیگران اطلاع دهید",
		fr:   "RSVP - faites-le savoir aux autres",
		id:   "RSVP - beri tahu orang lain",
		it:   "RSVP - fai sapere agli altri",
		jaJP: "RSVP - 他の人に知らせる",
		koKR: "RSVP - 다른 사람들에게 알리세요",
		pl:   "RSVP – daj znać innym",
		ptBR: "RSVP - avise os outros",
		ptPT: "RSVP - avise os outros",
		ru:   "Ответить — дайте знать другим",
		tr:   "RSVP - diğerlerine bildirin",
		ukUA: "RSVP – повідомте інших",
		uz:   "RSVP - boshqalarga xabar bering",
		zhCN: "RSVP - 让其他人知道"},
	SeeFullSpotProfile: {
		arEG: "🔍 الملف الكامل للمكان",
		de:   "🔍 Vollständiges Spotprofil",
		en:   "🔍 Full spot profile",
		es:   "🔍 Perfil completo del spot",
		faIR: "🔍 پروفایل کامل اسپات",
		fr:   "🔍 Profil complet du spot",
		id:   "🔍 Profil spot lengkap",
		it:   "🔍 Profilo spot completo",
		jaJP: "🔍 スポットの完全なプロフィール",
		koKR: "🔍 전체 스팟 프로필",
		pl:   "🔍 Pełny profil spotu",
		ptBR: "🔍 Perfil completo do local",
		ptPT: "🔍 Perfil completo do local",
		ru:   "🔍 Полный профиль места",
		tr:   "🔍 Tam spot profili",
		ukUA: "🔍 Повний профіль місця",
		uz:   "🔍 To&#39;liq joy profili",
		zhCN: "🔍 完整的现场资料"},
	ShareEventInlineDescription: {
		en: "Click this to share event link to the current chat.",
		ru: "Нажмите тут чтобы поделиться ссылкой на событие в текущем чате.",
	},
	EventIsCanceled: {
		en: "Event is cancelled",
		ru: "Мероприятие отменено",
	},
	DoYouWantToCancelThisEvent: {
		en: "Do you want to cancel this event?",
		ru: "Вы хотите отменить это событие?",
	},
	ShareSpotInlineDescription: {
		arEG: "انقر هنا لمشاركة رابط الموقع للدردشة الحالية.",
		de:   "Klicken Sie hier, um den Spot-Link zum aktuellen Chat freizugeben.",
		en:   "Click this to share spot link to the current chat.",
		es:   "Haga clic aquí para compartir el enlace al chat actual.",
		faIR: "برای اشتراک‌گذاری لینک گفتگوی فعلی، روی این کلیک کنید.",
		fr:   "Cliquez ici pour partager le lien du spot vers la discussion en cours.",
		id:   "Klik ini untuk membagikan tautan spot ke obrolan saat ini.",
		it:   "Fai clic qui per condividere il collegamento alla chat corrente.",
		jaJP: "現在のチャットへのスポットリンクを共有するには、これをクリックします。",
		koKR: "현재 채팅에 스팟 링크를 공유하려면 여기를 클릭하세요.",
		pl:   "Kliknij tutaj, aby udostępnić link do bieżącego czatu.",
		ptBR: "Clique aqui para compartilhar o link do chat atual.",
		ptPT: "Clique aqui para partilhar o link do chat atual.",
		ru:   "Нажмите тут чтобы поделиться ссылкой на спот в текущем чате.",
		tr:   "Mevcut sohbete ait spot bağlantıyı paylaşmak için buraya tıklayın.",
		ukUA: "Натисніть це, щоб поділитися посиланням на поточний чат.",
		uz:   "Joriy chatga joylashish havolasini baham ko&#39;rish uchun buni bosing.",
		zhCN: "单击此处可将现场链接分享到当前聊天。"},
	ProfileOfUser: {
		arEG: "الملف الشخصي لـ {USER_NAME}",
		de:   "Profil von {USER_NAME}",
		en:   "Profile of {USER_NAME}",
		es:   "Perfil de {USER_NAME}",
		faIR: "پروفایل {USER_NAME}",
		fr:   "Profil de {USER_NAME}",
		id:   "Profil {USER_NAME}",
		it:   "Profilo di {USER_NAME}",
		jaJP: "{USER_NAME}のプロフィール",
		koKR: "{USER_NAME}의 프로필",
		pl:   "Profil {USER_NAME}",
		ptBR: "Perfil de {USER_NAME}",
		ptPT: "Perfil de {USER_NAME}",
		ru:   "Профиль {USER_NAME}",
		tr:   "{USER_NAME} Profili",
		ukUA: "Профіль користувача {USER_NAME}",
		uz:   "{USER_NAME} profili",
		zhCN: "{USER_NAME} 的个人资料"},
	ActivityHiking: {
		arEG: "جولة على الأقدام",
		de:   "Wandern",
		en:   "Hiking",
		es:   "Senderismo",
		faIR: "پیاده‌روی",
		fr:   "Randonnée",
		id:   "Lintas alam",
		it:   "Escursionismo",
		jaJP: "ハイキング",
		koKR: "하이킹",
		pl:   "Turystyka piesza",
		ptBR: "Caminhada",
		ptPT: "Pedestrianismo",
		ru:   "Поход",
		tr:   "Doğa yürüyüşü",
		ukUA: "Піші прогулянки",
		uz:   "Piyoda yurish",
		zhCN: "远足"},
	ActivityBasketball: {
		arEG: "كرة السلة",
		de:   "Basketball",
		en:   "Basketball",
		es:   "Baloncesto",
		faIR: "بسکتبال",
		fr:   "Basket-ball",
		id:   "Bola basket",
		it:   "Pallacanestro",
		jaJP: "バスケットボール",
		koKR: "농구",
		pl:   "Koszykówka",
		ptBR: "Basquetebol",
		ptPT: "Basquetebol",
		ru:   "Баскетбол",
		tr:   "Basketbol",
		ukUA: "Баскетбол",
		uz:   "Basketbol",
		zhCN: "篮球"},
	ActivitySoccer: {
		arEG: "كرة القدم",
		de:   "Fußball",
		en:   "Football",
		es:   "Fútbol americano",
		faIR: "فوتبال",
		fr:   "Football",
		id:   "Sepak bola",
		it:   "Calcio",
		jaJP: "フットボール",
		koKR: "축구",
		pl:   "Piłkarski",
		ptBR: "Futebol",
		ptPT: "Futebol",
		ru:   "Футбол",
		tr:   "Futbol",
		ukUA: "Футбол",
		uz:   "Futbol",
		zhCN: "足球"},
	ActivityPingPong: {
		arEG: "كرة الطاولة",
		de:   "Tischtennis",
		en:   "Table Tennis",
		es:   "Tenis de mesa",
		faIR: "تنیس روی میز",
		fr:   "Tennis de table",
		id:   "Tenis meja",
		it:   "Tennis da tavolo",
		jaJP: "卓球",
		koKR: "탁구",
		pl:   "Tenis stołowy",
		ptBR: "Tênis de mesa",
		ptPT: "Tênis de mesa",
		ru:   "Настольный теннис",
		tr:   "Masa tenisi",
		ukUA: "Настільний теніс",
		uz:   "Stol tennisi",
		zhCN: "乒乓球"},
	ActivityTennis: {
		arEG: "التنس",
		de:   "Tennis",
		en:   "Tennis",
		es:   "Tenis",
		faIR: "تنیس",
		fr:   "Tennis",
		id:   "Tenis",
		it:   "Tennis",
		jaJP: "テニス",
		koKR: "테니스",
		pl:   "Tenis",
		ptBR: "Tênis",
		ptPT: "Tênis",
		ru:   "Теннис",
		tr:   "Tenis",
		ukUA: "Теніс",
		uz:   "Tennis",
		zhCN: "网球"},
	ActivitySurfing: {
		arEG: "ركوب الأمواج",
		de:   "Surfen",
		en:   "Surfing",
		es:   "Surf",
		faIR: "موج‌سواری",
		fr:   "Surf",
		id:   "Berselancar",
		it:   "Surf",
		jaJP: "サーフィン",
		koKR: "서핑",
		pl:   "Surfing",
		ptBR: "Surfe",
		ptPT: "Surfe",
		ru:   "Сёрфинг",
		tr:   "Sörf",
		ukUA: "Серфінг",
		uz:   "Serfing",
		zhCN: "冲浪"},
	ActivityFishing: {
		en: "Fishing",
		ru: "Рыбалка",
	},
	ActivityKitesurfing: {
		arEG: "ركوب الأمواج الشراعية",
		de:   "Kitesurfen",
		en:   "Kitesurfing",
		es:   "Kitesurf",
		faIR: "کایت سرفینگ",
		fr:   "Kitesurf",
		id:   "Selancar Layang",
		it:   "Kitesurf",
		jaJP: "カイトサーフィン",
		koKR: "카이트서핑",
		pl:   "Kitesurfing",
		ptBR: "Kitesurf",
		ptPT: "Kitesurf",
		ru:   "Кайт-сёрфинг",
		tr:   "Uçurtma sörfü",
		ukUA: "Кайтсерфінг",
		uz:   "Kaytsörfing",
		zhCN: "风筝冲浪"},
	ActivityRunning: {
		arEG: "جري",
		de:   "Läuft",
		en:   "Running",
		es:   "Correr",
		faIR: "دویدن",
		fr:   "En cours d&#39;exécution",
		id:   "Berlari",
		it:   "Corsa",
		jaJP: "ランニング",
		koKR: "달리기",
		pl:   "Działanie",
		ptBR: "Correndo",
		ptPT: "Corrida",
		ru:   "Бег",
		tr:   "Koşma",
		ukUA: "Біг",
		uz:   "Yugurish",
		zhCN: "跑步"},
	ActivityCycling: {
		arEG: "ركوب الدراجات",
		de:   "Radfahren",
		en:   "Cycling",
		es:   "Ciclismo",
		faIR: "دوچرخه‌سواری",
		fr:   "Vélo",
		id:   "Bersepeda",
		it:   "Ciclismo",
		jaJP: "サイクリング",
		koKR: "사이클링",
		pl:   "Kolarstwo",
		ptBR: "Ciclismo",
		ptPT: "Ciclismo",
		ru:   "Велоспорт",
		tr:   "Bisikletçilik",
		ukUA: "Велоспорт",
		uz:   "Velosport",
		zhCN: "骑自行车"},
	ActivitySkateboarding: {
		arEG: "التزلج على الألواح",
		de:   "Skateboarding",
		en:   "Skateboarding",
		es:   "Patineta",
		faIR: "اسکیت‌بورد",
		fr:   "Skateboard",
		id:   "Skateboarding",
		it:   "Skateboard",
		jaJP: "スケートボード",
		koKR: "스케이트보딩",
		pl:   "Jazda na deskorolce",
		ptBR: "Skate",
		ptPT: "Skate",
		ru:   "Скейтбординг",
		tr:   "Kaykay",
		ukUA: "Скейтбординг",
		uz:   "Skeytbording",
		zhCN: "滑板"},
	ActivityVolleyball: {
		arEG: "الكرة الطائرة",
		de:   "Volleyball",
		en:   "Volleyball",
		es:   "Voleibol",
		faIR: "والیبال",
		fr:   "Volley-ball",
		id:   "Bola voli",
		it:   "Pallavolo",
		jaJP: "バレーボール",
		koKR: "배구",
		pl:   "Siatkówka",
		ptBR: "Voleibol",
		ptPT: "Voleibol",
		ru:   "Волейбол",
		tr:   "Voleybol",
		ukUA: "Волейбол",
		uz:   "Voleybol",
		zhCN: "排球"},
	ActivitySwimming: {
		arEG: "سباحة",
		de:   "Baden",
		en:   "Swimming",
		es:   "Nadar",
		faIR: "شنا",
		fr:   "Natation",
		id:   "Renang",
		it:   "Nuoto",
		jaJP: "水泳",
		koKR: "수영",
		pl:   "Pływacki",
		ptBR: "Natação",
		ptPT: "Natação",
		ru:   "Плавание",
		tr:   "Yüzme",
		ukUA: "Плавання",
		uz:   "Suzish",
		zhCN: "游泳"},
	ActivityYoga: {
		arEG: "اليوغا",
		de:   "Yoga",
		en:   "Yoga",
		es:   "Yoga",
		faIR: "یوگا",
		fr:   "Yoga",
		id:   "Berjoget",
		it:   "Yoga",
		jaJP: "ヨガ",
		koKR: "요가",
		pl:   "Joga",
		ptBR: "Ioga",
		ptPT: "Ioga",
		ru:   "Йога",
		tr:   "Yoga",
		ukUA: "Йога",
		uz:   "Yoga",
		zhCN: "瑜伽"},
	ActivityClimbing: {
		arEG: "التسلق",
		de:   "Klettern",
		en:   "Climbing",
		es:   "Escalada",
		faIR: "کوهنوردی",
		fr:   "Escalade",
		id:   "Pendakian",
		it:   "Arrampicata",
		jaJP: "クライミング",
		koKR: "등반",
		pl:   "Wspinaczka",
		ptBR: "Escalando",
		ptPT: "Escalando",
		ru:   "Скалолазание",
		tr:   "Tırmanma",
		ukUA: "Скелелазіння",
		uz:   "Toqqa chiqish",
		zhCN: "攀登"},
	ActivityGym: {
		arEG: "نادي رياضي",
		de:   "Fitnessstudio",
		en:   "Gym",
		es:   "Gimnasia",
		faIR: "باشگاه ورزشی",
		fr:   "Salle de sport",
		id:   "Pusat kebugaran",
		it:   "Palestra",
		jaJP: "ジム",
		koKR: "체육관",
		pl:   "Sala gimnastyczna",
		ptBR: "Academia",
		ptPT: "Ginásio",
		ru:   "Тренажёрный зал",
		tr:   "Spor salonu",
		ukUA: "Тренажерний зал",
		uz:   "sportzal",
		zhCN: "健身房"},
	ActivityBookClub: {
		arEG: "نادي الكتاب",
		de:   "Buchclub",
		en:   "Book Club",
		es:   "Círculo de lectores",
		faIR: "باشگاه کتاب",
		fr:   "Club de lecture",
		id:   "Klub Buku",
		it:   "Club del libro",
		jaJP: "読書クラブ",
		koKR: "독서 동아리",
		pl:   "Klub Książki",
		ptBR: "Clube do Livro",
		ptPT: "Clube do Livro",
		ru:   "Книжный клуб",
		tr:   "Kitap Kulübü",
		ukUA: "Книжковий клуб",
		uz:   "Kitob klubi",
		zhCN: "读书俱乐部"},
	ActivityCoffeeMeetup: {
		arEG: "لقاء القهوة",
		de:   "Kaffee-Treffen",
		en:   "Coffee Meetup",
		es:   "Reunión de café",
		faIR: "گردهمایی قهوه",
		fr:   "Rencontre café",
		id:   "Pertemuan Kopi",
		it:   "Incontro al caffè",
		jaJP: "コーヒーミートアップ",
		koKR: "커피 미팅",
		pl:   "Spotkanie przy kawie",
		ptBR: "Encontro de Café",
		ptPT: "Encontro de Café",
		ru:   "Встреча за кофе",
		tr:   "Kahve Buluşması",
		ukUA: "Зустріч за кавою",
		uz:   "Kofe uchrashuvi",
		zhCN: "咖啡聚会"},
	ActivityGameNight: {
		arEG: "ليلة اللعبة",
		de:   "Spieleabend",
		en:   "Game Night",
		es:   "Noche de juegos",
		faIR: "شب بازی",
		fr:   "Soirée jeux",
		id:   "Malam Permainan",
		it:   "Serata di giochi",
		jaJP: "ゲームナイト",
		koKR: "게임의 밤",
		pl:   "Wieczór gier",
		ptBR: "Noite de jogos",
		ptPT: "Noite de jogos",
		ru:   "Игровой вечер",
		tr:   "Oyun Gecesi",
		ukUA: "Ігровий вечір",
		uz:   "O&#39;yin kechasi",
		zhCN: "游戏之夜"},
	ActivityMovieNight: {
		arEG: "ليلة الفيلم",
		de:   "Filmabend",
		en:   "Movie Night",
		es:   "Noche de cine",
		faIR: "شب فیلم",
		fr:   "Soirée cinéma",
		id:   "Malam Film",
		it:   "Serata cinema",
		jaJP: "映画の夜",
		koKR: "영화의 밤",
		pl:   "Wieczór filmowy",
		ptBR: "Noite de cinema",
		ptPT: "Noite de cinema",
		ru:   "Киновечер",
		tr:   "Film Gecesi",
		ukUA: "Кіновечір",
		uz:   "Kino kechasi",
		zhCN: "电影之夜"},
	ActivityTriviaNight: {
		arEG: "ليلة التوافه",
		de:   "Quizabend",
		en:   "Trivia Night",
		es:   "Noche de trivia",
		faIR: "شب چیزهای بی اهمیت",
		fr:   "Soirée quiz",
		id:   "Malam Trivia",
		it:   "Serata dei quiz",
		jaJP: "トリビアナイト",
		koKR: "퀴즈 나이트",
		pl:   "Wieczór wiedzy",
		ptBR: "Noite de curiosidades",
		ptPT: "Noite de curiosidades",
		ru:   "Викторина",
		tr:   "Bilgi Yarışması Gecesi",
		ukUA: "Вечір вікторин",
		uz:   "Trivia kechasi",
		zhCN: "琐事之夜"},
	ActivityPotluck: {
		arEG: "وجبة مشتركة",
		de:   "Potluck",
		en:   "Potluck",
		es:   "Probabilidad",
		faIR: "پاتلاک",
		fr:   "Repas-partage",
		id:   "Makan bersama",
		it:   "Potluck",
		jaJP: "ポットラック",
		koKR: "포틀럭",
		pl:   "Cośkolwiek do jedzenia",
		ptBR: "Potluck",
		ptPT: "Potluck",
		ru:   "Обед в складчину",
		tr:   "Yemek paylaşımı",
		ukUA: "Спільний обід",
		uz:   "Potluck",
		zhCN: "便饭"},
	ActivityPicnic: {
		arEG: "نزهه",
		de:   "Picknick",
		en:   "Picnic",
		es:   "Picnic",
		faIR: "پیک نیک",
		fr:   "Pique-nique",
		id:   "Piknik",
		it:   "Picnic",
		jaJP: "ピクニック",
		koKR: "피크닉",
		pl:   "Piknik",
		ptBR: "Piquenique",
		ptPT: "Piquenique",
		ru:   "Пикник",
		tr:   "Piknik",
		ukUA: "Пікнік",
		uz:   "Piknik",
		zhCN: "野餐"},
	ActivityBarbecue: {
		arEG: "الشواء",
		de:   "Grill",
		en:   "Barbecue",
		es:   "Parilla",
		faIR: "باربیکیو",
		fr:   "Barbecue",
		id:   "Barbekyu",
		it:   "Barbecue",
		jaJP: "バーベキュー",
		koKR: "야외 파티",
		pl:   "Rożen",
		ptBR: "Churrasco",
		ptPT: "Churrasco",
		ru:   "Барбекю",
		tr:   "Barbekü",
		ukUA: "Барбекю",
		uz:   "Barbekyu",
		zhCN: "烧烤"},
	ActivityCrafting: {
		arEG: "صناعة",
		de:   "Basteln",
		en:   "Crafting",
		es:   "Elaboración",
		faIR: "صنایع دستی",
		fr:   "Artisanat",
		id:   "Kerajinan",
		it:   "Artigianato",
		jaJP: "クラフト",
		koKR: "크래프팅",
		pl:   "Rzemiosło",
		ptBR: "Artesanato",
		ptPT: "Artesanato",
		ru:   "Рукоделие",
		tr:   "El sanatları",
		ukUA: "Крафтінг",
		uz:   "Hunarmandchilik",
		zhCN: "制作"},
	ActivityKaraoke: {
		arEG: "كاريوكي",
		de:   "Karaoke",
		en:   "Karaoke",
		es:   "Karaoke",
		faIR: "کارائوکه",
		fr:   "Karaoké",
		id:   "Karaoke",
		it:   "Karaoke",
		jaJP: "カラオケ",
		koKR: "노래방",
		pl:   "Karaoke",
		ptBR: "Karaokê",
		ptPT: "Karaoke",
		ru:   "Караоке",
		tr:   "Karaoke",
		ukUA: "Караоке",
		uz:   "Karaoke",
		zhCN: "卡拉OK"},
	ActivityMusicJam: {
		arEG: "مربى الموسيقى",
		de:   "Musik Jam",
		en:   "Music Jam",
		es:   "Jam musical",
		faIR: "موسیقی جم",
		fr:   "Musique Jam",
		id:   "Jam Musik",
		it:   "Jam musicale",
		jaJP: "ミュージックジャム",
		koKR: "뮤직잼",
		pl:   "Muzyczny dżem",
		ptBR: "Jam de música",
		ptPT: "Jam de música",
		ru:   "Музыкальный джем",
		tr:   "Müzik Sıkışıklığı",
		ukUA: "Музичний джем",
		uz:   "Musiqa Jam",
		zhCN: "音乐果酱"},
	ActivityBoardGames: {
		arEG: "ألعاب الطاولة",
		de:   "Brettspiele",
		en:   "Board Games",
		es:   "Juegos de mesa",
		faIR: "بازی‌های تخته‌ای",
		fr:   "Jeux de société",
		id:   "Permainan Papan",
		it:   "Giochi da tavolo",
		jaJP: "ボードゲーム",
		koKR: "보드 게임",
		pl:   "Gry planszowe",
		ptBR: "Jogos de tabuleiro",
		ptPT: "Jogos de tabuleiro",
		ru:   "Настольные игры",
		tr:   "Masa Oyunları",
		ukUA: "Настільні ігри",
		uz:   "Stol oʻyinlari",
		zhCN: "棋盘游戏"},
	ActivityArtJam: {
		arEG: "مربى الفن",
		de:   "Art Jam",
		en:   "Art Jam",
		es:   "Art Jam",
		faIR: "جام هنر",
		fr:   "Art Jam",
		id:   "kemacetan seni",
		it:   "Art Jam",
		jaJP: "アートジャム",
		koKR: "아트잼",
		pl:   "Jam artystyczny",
		ptBR: "Geleia de Arte",
		ptPT: "Geleia de Arte",
		ru:   "Творческий джем",
		tr:   "Sanat Sıkışmaları",
		ukUA: "Арт Джем",
		uz:   "Art Jam",
		zhCN: "艺术果酱"},
	ActivityParkPlaydate: {
		arEG: "موعد اللعب في الحديقة",
		de:   "Spieltreffen im Park",
		en:   "Park Playdate",
		es:   "Cita para jugar en el parque",
		faIR: "پارک بازی",
		fr:   "Rendez-vous de jeu au parc",
		id:   "Taman Bermain",
		it:   "Appuntamento al parco",
		jaJP: "公園でのプレイデート",
		koKR: "파크 플레이데이트",
		pl:   "Spotkanie w parku",
		ptBR: "Encontro no parque",
		ptPT: "Encontro no parque",
		ru:   "Встреча в парке",
		tr:   "Park Oyun Randevusu",
		ukUA: "Ігрове побачення в парку",
		uz:   "Park o&#39;yin sanasi",
		zhCN: "公园玩耍约会"},
	ActivityStoryTime: {
		arEG: "وقت القصة",
		de:   "Märchenstunde",
		en:   "Story Time",
		es:   "La hora del cuento",
		faIR: "زمان داستان",
		fr:   "L&#39;heure du conte",
		id:   "Waktu Cerita",
		it:   "L&#39;ora della storia",
		jaJP: "ストーリータイム",
		koKR: "스토리 타임",
		pl:   "Czas na opowieść",
		ptBR: "Hora da história",
		ptPT: "Hora da história",
		ru:   "Чтение сказок",
		tr:   "Hikaye Zamanı",
		ukUA: "Час історій",
		uz:   "Hikoya vaqti",
		zhCN: "故事时间"},
	ActivityToySwap: {
		arEG: "تبادل الألعاب",
		de:   "Spielzeugtausch",
		en:   "Toy Swap",
		es:   "Intercambio de juguetes",
		faIR: "مبادله اسباب بازی",
		fr:   "Échange de jouets",
		id:   "Tukar Mainan",
		it:   "Scambio di giocattoli",
		jaJP: "おもちゃ交換",
		koKR: "장난감 교환",
		pl:   "Wymiana zabawek",
		ptBR: "Troca de brinquedos",
		ptPT: "Troca de brinquedos",
		ru:   "Обмен игрушками",
		tr:   "Oyuncak Takası",
		ukUA: "Обмін іграшками",
		uz:   "O&#39;yinchoq almashish",
		zhCN: "玩具交换"},
	ActivityHomeworkHelp: {
		arEG: "مساعدة في الواجبات المنزلية",
		de:   "Hausaufgabenhilfe",
		en:   "Homework Help",
		es:   "Ayuda con las tareas",
		faIR: "کمک در انجام تکالیف",
		fr:   "Aide aux devoirs",
		id:   "Bantuan Pekerjaan Rumah",
		it:   "Aiuto con i compiti",
		jaJP: "宿題ヘルプ",
		koKR: "숙제 도움",
		pl:   "Pomoc w odrabianiu zadań domowych",
		ptBR: "Ajuda com a lição de casa",
		ptPT: "Ajuda com os trabalhos de casa",
		ru:   "Помощь с домашкой",
		tr:   "Ödev Yardımı",
		ukUA: "Допомога з домашнім завданням",
		uz:   "Uy vazifasiga yordam",
		zhCN: "家庭作业帮助"},
	ActivityBirthdayParty: {
		arEG: "حفلة عيد ميلاد",
		de:   "Geburtstagsfeier",
		en:   "Birthday Party",
		es:   "Fiesta de cumpleaños",
		faIR: "جشن تولد",
		fr:   "Fête d&#39;anniversaire",
		id:   "Pesta Ulang Tahun",
		it:   "Festa di compleanno",
		jaJP: "誕生日パーティー",
		koKR: "생일 파티",
		pl:   "Urodziny",
		ptBR: "Festa de aniversário",
		ptPT: "Festa de aniversário",
		ru:   "День рождения",
		tr:   "Doğum günü partisi",
		ukUA: "Вечірка на день народження",
		uz:   "Tug&#39;ilgan kun partiyasi",
		zhCN: "生日聚会"},
	ActivityBabysittingSwap: {
		arEG: "تبادل مجالسة الأطفال",
		de:   "Babysitter-Tausch",
		en:   "Babysitting Swap",
		es:   "Intercambio de niñeras",
		faIR: "تعویض نگهداری از کودک",
		fr:   "Échange de baby-sitting",
		id:   "Pertukaran Pengasuh Bayi",
		it:   "Scambio di babysitter",
		jaJP: "ベビーシッター交換",
		koKR: "베이비시팅 스왑",
		pl:   "Zastępstwo Opieki nad Dziećmi",
		ptBR: "Troca de babá",
		ptPT: "Troca de babysitting",
		ru:   "Обмен няней",
		tr:   "Bebek Bakıcılığı Takası",
		ukUA: "Обмін послугами няні",
		uz:   "Chaqaloqlarni almashtirish",
		zhCN: "保姆互换"},
	ActivityStudyGroup: {
		arEG: "مجموعة الدراسة",
		de:   "Studiengruppe",
		en:   "Study Group",
		es:   "Grupo de estudio",
		faIR: "گروه مطالعه",
		fr:   "Groupe d&#39;étude",
		id:   "Belajar kelompok",
		it:   "Gruppo di studio",
		jaJP: "研究グループ",
		koKR: "스터디 그룹",
		pl:   "Grupa studyjna",
		ptBR: "Grupo de Estudo",
		ptPT: "Grupo de Estudo",
		ru:   "Учебная группа",
		tr:   "Çalışma Grubu",
		ukUA: "Навчальна група",
		uz:   "O&#39;quv guruhi",
		zhCN: "学习小组"},
	ActivityGaming: {
		arEG: "الألعاب",
		de:   "Spiele",
		en:   "Gaming",
		es:   "Juego de azar",
		faIR: "بازی",
		fr:   "Jeux vidéo",
		id:   "Permainan",
		it:   "Gioco d&#39;azzardo",
		jaJP: "ゲーム",
		koKR: "노름",
		pl:   "Hazard",
		ptBR: "Jogos",
		ptPT: "Jogos",
		ru:   "Игры",
		tr:   "Oyun",
		ukUA: "Ігри",
		uz:   "Oʻyin",
		zhCN: "赌博"},
	ActivitySkating: {
		arEG: "التزلج",
		de:   "Skaten",
		en:   "Skating",
		es:   "Patinaje",
		faIR: "اسکیت",
		fr:   "Patinage",
		id:   "Berseluncur",
		it:   "Pattinaggio",
		jaJP: "スケート",
		koKR: "스케이팅",
		pl:   "Łyżwiarstwo",
		ptBR: "Patinação",
		ptPT: "Patinação",
		ru:   "Катание на коньках",
		tr:   "Paten kaymak",
		ukUA: "Катання на ковзанах",
		uz:   "Konkida uchish",
		zhCN: "溜冰"},
	ActivityDanceClass: {
		arEG: "صف الرقص",
		de:   "Tanzkurs",
		en:   "Dance Class",
		es:   "Clase de baile",
		faIR: "کلاس رقص",
		fr:   "Cours de danse",
		id:   "Kelas Tari",
		it:   "Lezione di danza",
		jaJP: "ダンスクラス",
		koKR: "댄스 수업",
		pl:   "Klasa taneczna",
		ptBR:

		//ButtonSoundsGood: buttonSoundsGood(),
		"Aula de dança",
		ptPT: "Aula de dança",
		ru:   "Танцевальный класс",
		tr:   "Dans Dersi",
		ukUA: "Урок танців",
		uz:   "Raqs darsi",
		zhCN: "舞蹈课"},
	ActivityHangout: {
		arEG: "يتسكع",
		de:   "Abhängen",
		en:   "Hangout",
		es:   "Pasar el rato",
		faIR: "پاتوق",
		fr:   "Traîner",
		id:   "Nongkrong bareng",
		it:   "Uscire",
		jaJP: "遊ぶ",
		koKR: "놀다",
		pl:   "Spędzać czas",
		ptBR: "Passar tempo junto",
		ptPT: "Passar o tempo",
		ru:   "Тусовка",
		tr:   "Oyalanmak",
		ukUA: "Тусовка",
		uz:   "Vaqtni chog &#39;o&#39;tkazish",
		zhCN: "挂出"},
	ActivityEscapeRoom: {
		arEG: "غرفة الهروب",
		de:   "Fluchtraum",
		en:   "Escape Room",
		es:   "Sala de escape",
		faIR: "اتاق فرار",
		fr:   "Escape Room",
		id:   "Ruang Pelarian",
		it:   "Stanza di fuga",
		jaJP: "脱出ゲーム",
		koKR: "탈출실",
		pl:   "Pokój zagadek",
		ptBR: "Sala de fuga",
		ptPT: "Sala de fuga",
		ru:   "Квест-комната",
		tr:   "Kaçış Odası",
		ukUA: "Квест-кімната",
		uz:   "Qochish xonasi",
		zhCN: "密室逃脱",
	},
	TogetheredBotDescription: {
		en: `@ToGetheredBot offers two main features:

	1. 📍🗓️🕒📣 Plans Sharing – let others know where and when you're planning to be somewhere and know who is in - before you head out.

	2. 🗓️🕒📍🗳️ Events Coordination – easily organise group activities by suggesting multiple time and place options. The bot collects votes from participants and shows which combinations work best, helping the group agree on a plan without long chat threads.`,
		ru: `@ToGetheredBot предлагает две основные функции:

	1. 📍🗓️🕒📣 Обмен планами — сообщите другим, где и когда вы планируете быть, и узнайте, кто там будет, — прежде чем вы отправитесь.

	2. 🗓️🕒📍🗳️ Координация мероприятий — легко организуйте групповые мероприятия, предлагая несколько вариантов времени и места. Бот собирает голоса участников и показывает, какие комбинации работают лучше всего, помогая группе согласовать планы без длинных чатов.`,
	},
	TogetheredBotShortDescription: {
		en: "📍🗓️🕒📣🗳 Turns intentions into gatherings. Helps to organize events. Know who is in - before you head out.",
		ru: "📍🗓️🕒📣🗳 Превращает намерения во встречи. Помогает организовывать мероприятия. Показывает кто собирается учавствовать.",
	},
	AboutCommandDescription: {
		en: "About our bot",
		ru: "О нашем боте",
	},
	TogetheredBotCommandDescription: {
		en: "Main menu",
		ru: "Главное меню",
	},
	SettingsCommandDescription: {
		en: "Check & adjust bot settings",
		ru: "Проверить и поменять настройки бота",
	},
	TermsCommandDescription: {
		en: "Terms & Conditions",
		ru: "Пользовательское соглашение",
	},
	TogetheredBotCommandSpots: {
		en: "My places of events",
		ru: "Мои места проведения событий",
	},
	TogetheredBotCommandPlans: {
		en: "My plans",
		ru: "Мои планы",
	},
	AdvertisingCommandDescription: {
		en: "Advertise at @{BOT_CODE}",
		ru: "Реклама в @{BOT_CODE}",
	},
	SupportUsCommandDescription: {
		en: "Support development of @{BOT_CODE}",
		ru: "Поддержать разработку @{BOT_CODE}",
	},
	MembershipCommandDescription: {
		en: "Use for free or get PRO account",
		ru: "Используй бесплатно или выбери ПРО акккаунт",
	},
	WalletCommandDescription: {
		en: "Your balance & latest transactions",
		ru: "Ваш баланс и последние транзакции",
	},
	WalletTitle: {
		en: "My balance @{BOT_CODE}",
		ru: "Мой баланс @{BOT_CODE}",
	},
	TopUpWithStarts: {
		en: "➕ Top up with ⭐",
		ru: "➕ Пополнить кошелёк ⭐",
	},
	StarsBalance: {
		en: "⭐ %d stars",
		ru: "⭐ %d звёзд",
	},
	RemoveActivityButtonText: {
		en: "❌ Remove activity",
		ru: "❌ Убрать активность",
	},
	MyFollowersTitle: {
		en: "🕵️‍♂ My follower",
		ru: "🕵️‍♂ Мои подписчики",
	},
	FollowersOfUserTitle: {
		en: "🕵️‍♂ Follower of %s",
		ru: "🕵️‍♂ Подписчики %s",
	},
	MyFollowsTitle: {
		en: "👀 My follows",
		ru: "👀 Мои подписки",
	},
	UserFollowingTitle: {
		en: "👀‍ %s follows",
		ru: "👀 Подписчики %s",
	},
	IntentNotificationMessageUserIsGoingTo: {
		en: "{USER_NAME} is going to",
		ru: "{USER_NAME} планирует быть",
	},
	IntentNotificationPlannedActivities: {
		en: "Planned activities",
		ru: "Запланированные активности",
	},
	WouldYouJoinRsvpNow: {
		en: "Would you join? RSVP now!",
		ru: "Вы присоеденитесь? Дайте знать!",
	},
	RsvpWizardTitle: {
		en: "Let others know where & where you going to be.",
		ru: "Сообщи другим, где и когда ты собираешься быть.",
	},
	AboutBotTitle: {
		en: "About @%s",
		ru: "Про @%s",
	},
	TechStackTitle: {
		en: "Our tech stack",
		ru: "Наш технологический стек",
	},
	TechStackText: {
		en: `
- written in <a href='https://go.dev/'>Go language</a>
- hosted on Google <a href='https://cloud.google.com/appengine'>App Engine</a>
- data stored in Google <a href='https://firebase.google.com/products/firestore'>Firestore DB</a>
- uses https://github.com/bots-go-framework/
- uses <a href='https://github.com/bots-go-framework/bots-api-telegram'>bots-api-telegram</a> to communicate with Telegram
uses https://github.com/dal-go/ as Database Abstraction Layer

If you want to know more about how we develop subscribe to @SneatDevDiaries!
`,
		ru: `
- написан на языке программирования <a href='https://go.dev/'>Go</a>
- работает на Google <a href='https://cloud.google.com/appengine'>App Engine</a>
- данные хранятся в Google <a href='https://firebase.google.com/products/firestore'>Firestore DB</a>
- использует https://github.com/bots-go-framework/
- использует <a href='https://github.com/bots-go-framework/bots-api-telegram'>bots-api-telegram</a> для коммуникции с Telegram
- использует https://github.com/dal-go/ как <code>Database Abstraction Layer</code>

Если хотите узнать боьлше про наш процесс разработки подписывайтесь на @SneatDevDiaries!
`,
	},
	WouldYouSupportUsTitle: {
		en: "Would you support us?",
		ru: "Поддержите нас?",
	},
	WouldYouSupportUsText: {
		en: `
If you found @{BOT_CODE} to be useful to you please consider taking a minute to:")

	- <a href='https://t.me/{BOT_CODE}?start=donate'>donate</a> few Telegram stars to speedup development

	- give a star to our open source projects - <a href='https://github.com/bots-go-framework/'>bots-go-framework</a>, <a href='https://github.com/dal-go/'>DALgo</a>")

<i>Support from our users boosts our motivation to do more cool stuff on the bot!</i>
`,
		ru: `
Если @{BOT_CODE} оказался полезным для вам подумайте о поддержке нашего проекта. Вы можете:

	- <a href='https://t.me/{BOT_CODE}?start=donate'>задонатить</a> Телеграм звёзды чтобы ускороить разработку

	- поставить звезду одному из наших проектов с открытым исходным кодом: <a href='https://github.com/bots-go-framework/'>bots-go-framework</a>, <a href='https://github.com/dal-go/'>DALgo</a>

<i>Поддердка наших пользователей сильно повышает нашу мотивацию по улучшению бота!</i>
`,
	},
	PremiumAccountTitle: {
		en: "Premium Account",
		ru: "Премиальный аккаунт",
	},
	PremiumAccountText: {
		en: `
- Monthly plan - 100⭐ per month
- Yearly plan - <u>save 25%%</u> - 75⭐ per month`,
		ru: `
- Ежемесячный план - 100⭐ в месяц
- Годовой план - <u>экономия 25%%</u> - 75⭐ в месяц`,
	},
	GetMonthlyPlanButtonText: {
		en: "I choose monthly plan",
		ru: "Я выбираю ежемесячный план",
	},
	GetYearlyPlanButtonText: {
		en: "I choose yearly plan",
		ru: "Я выбираю ежегодный план",
	},
	WhyGoPremiumTitle: {
		en: "Why go premium?",
		ru: "Почему стоит выбрать премиум?",
	},
	WhyGoPremiumText: {
		en: `
You always can use our bot for free .

But ads in a bot hardly cover even hosting expenses. By choosing a paid plan you:

- Support improvements & new features
- Hide ads (<i>except the local ads linked to a specific spot/venue</i>)
`,
		ru: `
Вы всегда сможете использовать наш бот беслатно.

Но реклама в боте едва покрывает затраты на хостинг. Выбирая платный тариф вы:

- Поддержите разработку улучшениий
- Скроете рекламу (<i>кроме локальной - привязанной к определённому месту проведения событий</i>)`,
	},
	InvoiceMonthlyPlanTitle: {
		en: "Monthly payment for @%s",
		ru: "Ежемесячный платёж для @%s",
	},
	InvoiceYearlyPlanTitle: {
		en: "Yearly payment for @%s",
		ru: "Ежегодный платёж для @%s",
	},
	StarsPerMonth: {
		en: "%d⭐ per month",
		ru: "%d⭐ за месяц",
	},
	DonationToBotInvoiceTitle: {
		en: "Donation to @%s",
		ru: "Пожертвование для @%s",
	},
	DonationToBotInvoiceDescription: {
		en: "Donation to @%s",
		ru: "Пожертвование для @%s",
	},
	DonationInvoiceLabelCoffee: {
		en: "☕️ Buy us a coffee ~ €%d",
		ru: "☕️ Купите нам кофе ~ €%d",
	},
	DonationInvoiceLabelBeer: {
		en: "🍺 Buy us a beer ~ €%d",
		ru: "🍺 Купите нам пива ~ €%d",
	},
	DonationInvoiceLabelLunch: {
		en: "🍽️ Buy us a lunch ~ €%d",
		ru: "🍽️ Купите нам обед ~ €%d",
	},
	DonationInvoiceLabelAwesomeness: {
		en: "🙌 Is this bot that awesome? ~ €%d",
		ru: "🙌 Является ли этот бот настолько потрясаюшим? ~ €%d",
	},
	DonationButtonCoffee: {
		en: "☕️ Buy us a coffee %d⭐ ~ €%d",
		ru: "☕️ Купите нам кофе %d⭐ ~ €%d",
	},
	DonationButtonBeer: {
		en: "🍺 Buy us a beer %d⭐ ~ €%d",
		ru: "🍺 Купите нам пива %d⭐ ~ €%d",
	},
	DonationButtonLunch: {
		en: "🍽️ Buy us a lunch %d⭐ ~ €%d",
		ru: "🍽️ Купите нам обед %d⭐ ~ €%d",
	},
	DonationButtonAwesomeness: {
		en: "🚀🔥🙌 You are awesome! - %s🌟 ~ €%d",
		ru: "🚀🔥🙌 Вы супер! - %s🌟 ~ €%d",
	},
	SupportDevelopmentOfBotTitle: {
		en: "Support development of @%s",
		ru: "Поддержите развитие @%s",
	},
	SupportDevelopmentOfBotText: {
		en: `
By giving us a 1-time donation you are helping to accelerate development of the bot.

Also it would be cool if you can <a href="https://t.me/{BOT_CODE}?start=get_pro_account">subscribe to our PRO membership</a>.

If monetary donation is not an option at the moment and you have an account at GitHub we'd appreciate if you grant a star to our open source projects:
- <a href="https://github.com/bots-go-framework/bots-fw">bots-go-framework</a>
- <a href="https://github.com/dal-go/dalgo">DALgo</a> - Data-access Abstraction Layer 
`,
		ru: `
Сделав разовое пожертвование, вы поможете ускорить разработку бота.

Так же будет здорово если вы можете <a href="https://t.me/{BOT_CODE}?start=get_pro_account">подписаться на наш премиальный тариф</a>.

Если денежное пожертвование в текущем моменте не вариант и у вас есть аккаунт на GitHub мы будем признательны если поставите звезду нашим проектам с открытым исходным кодом:
- <a href="https://github.com/bots-go-framework/bots-fw">bots-go-framework</a>
- <a href="https://github.com/dal-go/dalgo">DALgo</a> - Data-access Abstraction Layer
`,
	},
	AdvertiseAtBotTitle: {
		en: "Advertising at @%s",
		ru: "Реклама на @%s",
	},
	AdvertiseAtBotText: {
		en: `
You can place an ad for activity from here (<i>with optional targeting by city/country</i>).

If you want to create a local add for a spot do it from the spot menu.`,
		ru: `
Вы можете разместить объявление о мероприятии отсюда (<i>с необязательным таргетингом по городу/стране</i>).

Если вы хотите создать локальное объявление для спота, сделайте это через меню спота.`,
	},
	MyAdsButtonText: {
		en: "🛍️ My ads",
		ru: "🛍️ Мои объявления",
	},
	NewAdButtonText: {
		en: "➕ New ad",
		ru: "➕ Новое объявление",
	},
	AdsBillingButtonText: {
		en: "🧾 Billing",
		ru: "🧾 Биллинг",
	},
	AdvertiseWithUsButtonText: {
		en: "🛍️ Advertise with us",
		ru: "🛍️ Реклама в боте",
	},
	GetProAccountButtonText: {
		en: "🚀 Get PREMIUM account",
		ru: "🚀 Премиальный аккаунт",
	},
	SupportUsButtonText: {
		en: "❤️ Support us",
		ru: "❤️ Поддержать",
	},
	VewEventDetailsButtonText: {
		en: "Go to event",
		ru: "Перейти на мероприятие",
	},
	RsvpTpEventButtonText: {
		en: "RSVP",
		ru: "Подтвердить участие",
	},
	PromoTitle: {
		en: "Promo",
		ru: "Промо",
	},
	TogdAd1: {
		en: `You can <a href="https://t.me/{BOT_CODE}?start=advertising">place an ad</a> or <a href="https://t.me/{BOT_CODE}?start=get_pro_account">get <b>PRO account</b></a> to <a href="https://t.me/{BOT_CODE}?start=support_us">support us</a>.`,
		ru: `Вы можете <a href="https://t.me/{BOT_CODE}?start=advertising">разместить рекламу</a> или <a href="https://t.me/{BOT_CODE}?start=get_pro_account">выбрать <b>ПРО аккаунт</b></a> чтобы <a href="https://t.me/{BOT_CODE}?start=support_us">поддержать нас</a>.`,
	},
	TogdAd2: {
		en: `Get <a href="https://t.me/{BOT_CODE}?start=get_pro_account"><b>PRO</b> account</a> to get rid of <a href="https://t.me/{BOT_CODE}?start=advertising">advertising</a> & to <a href="https://t.me/{BOT_CODE}?start=support_us">support us</a>.`,
		ru: `Выберите <a href="https://t.me/{BOT_CODE}?start=get_pro_account"><b>ПРО</b> аккаунт</a> чтобы убрать <a href="https://t.me/{BOT_CODE}?start=advertising">рекламу</a> а также <a href="https://t.me/{BOT_CODE}?start=support_us">поддержать нас</a>.`,
	},
	BeforeStartChooseTimezoneTitle: {
		en: "Before we start please choose your timezone",
		ru: "Перед тем как начать выберите вашу временную зону",
	},
	BeforeStartChooseTimezoneText: {
		en: `
This is needed to:
 - Correctly display times (we use UTC internally)
 - Avoid sending notifications during night hours
 - Save your time when creating spots and events (<i>your timezone will be used by default</i>)

The easiest way to set your timezone is to send your location.
But you also can choose from the list provided.`,
		ru: `
Это необходимо чтобы:
- Правильно отоброжать время (<i>мы храним врямя как UTC</i>) 
- Избежать отравки оповещений когда у вас ночь
- Ускороить создание событий и места их проведения (<i>вашу временная зона будет исользвоана по умолчанию</i>.)`,
	},
	HintToSendLocationForTimezoneDetection: {
		en: `The easiest and fastest way to set your timezone is to <b>send your location</b> - works only mobile devices. But you also can choose from the list provided.`,
		ru: `Самый простой и быстрый способ задать вашу временную зону это '<b>отправить ваше месторасположение</b> - работает только на мобильных устройствах. Но вы можете просто выбрать временную зону из списка.`,
	},
	YourTimezoneIsSetTo: {
		en: "Your time zone is set to %s",
		ru: "Ваша временная зона установлена: %s",
	},
	SendLocationForTimezoneDetection: {
		en: "Detect timezone by location",
		ru: "Определить временную зону",
	},
}
