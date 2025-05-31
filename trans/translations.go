package trans

const adsCommandTitle = "\xE2\xAD\x90\xE2\xAD\x90\xE2\xAD\x90"

/*
Proper order of locale keys in var TRANS, use it as a reference for all values:
- "de-DE":
- "en-UK":
- "en-US":
- "es-ES":
- "fa-IR":
- "fr-FR":
- "id-ID":
- "it-IT":
- "ja-JP":
- "ko-KO":
- "pl-PL":
- "pt-BR":
- "ru-RU":
- "tr-TR":
- "ua-UA":
- "uz-UZ":
- "zh-CN":
*/

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		"de-DE": "BEISPIEL",
		"en-UK": "SAMPLE",
		"en-US": "SAMPLE",
		"es-ES": "EJEMPLO",
		"fa-IR": "نمونه",
		"fr-FR": "EXEMPLE",
		"id-ID": "CONTOH",
		"it-IT": "ESEMPIO",
		"ja-JP": "例",
		"ko-KO": "예",
		"pl-PL": "PRZYKŁAD",
		"pt-BR": "EXEMPLO",
		"ru-RU": "ПРИМЕР",
		"tr-TR": "ÖRNEK",
		"ua-UA": "ПРИКЛАД",
		"uz-UZ": "MISOL",
		"zh-CN": "例子",
	},

	HowdyUser: {
		"de-DE": "Hallo {USER_NAME}!",
		"en-UK": "Howdy {USER_NAME}!",
		"en-US": "Howdy {USER_NAME}!",
		"es-ES": "¡Hola {USER_NAME}!",
		"fa-IR": "سلام {USER_NAME}!",
		"fr-FR": "Salut {USER_NAME} !",
		"id-ID": "Halo {USER_NAME}!",
		"it-IT": "Ciao {USER_NAME}!",
		"ja-JP": "やあ {USER_NAME}!",
		"ko-KO": "안녕 {USER_NAME}!",
		"pl-PL": "Cześć {USER_NAME}!",
		"pt-BR": "Olá {USER_NAME}!",
		"ru-RU": "Привет {USER_NAME}!",
		"tr-TR": "Merhaba {USER_NAME}!",
		"ua-UA": "Привіт {USER_NAME}!",
		"uz-UZ": "Salom {USER_NAME}!",
		"zh-CN": "你好 {USER_NAME}！",
	},

	ButtonAdd: {
		"de-DE": "Hinzufügen",
		"en-UK": "Add",
		"en-US": "Add",
		"es-ES": "Añadir",
		"fa-IR": "افزودن",
		"fr-FR": "Ajouter",
		"id-ID": "Tambah",
		"it-IT": "Aggiungi",
		"ja-JP": "追加",
		"ko-KO": "추가",
		"pl-PL": "Dodaj",
		"pt-BR": "Adicionar",
		"ru-RU": "Добавить",
		"tr-TR": "Ekle",
		"ua-UA": "Додати",
		"uz-UZ": "Qo'shish",
		"zh-CN": "添加",
	},
	ButtonRemove: {
		"de-DE": "Entfernen",
		"en-UK": "Remove",
		"en-US": "Remove",
		"es-ES": "Eliminar",
		"fa-IR": "حذف",
		"fr-FR": "Supprimer",
		"id-ID": "Hapus",
		"it-IT": "Rimuovi",
		"ja-JP": "削除",
		"ko-KO": "제거",
		"pl-PL": "Usuń",
		"pt-BR": "Remover",
		"ru-RU": "Удалить",
		"tr-TR": "Kaldır",
		"ua-UA": "Видалити",
		"uz-UZ": "Olib tashlash",
		"zh-CN": "移除",
	},

	"Jan": {
		"de-DE": "Jan",
		"en-UK": "Jan",
		"en-US": "Jan",
		"es-ES": "Ene",
		"fa-IR": "ژانویه",
		"fr-FR": "Jan",
		"id-ID": "Jan",
		"it-IT": "Gen",
		"ja-JP": "1月",
		"ko-KO": "1월",
		"pl-PL": "Sty",
		"pt-BR": "Jan",
		"ru-RU": "Янв.",
		"tr-TR": "Oca",
		"ua-UA": "Січ",
		"uz-UZ": "Yan",
		"zh-CN": "一月",
	},

	"Feb": {
		"de-DE": "Feb",
		"en-UK": "Feb",
		"en-US": "Feb",
		"es-ES": "Feb",
		"fa-IR": "فوریه",
		"fr-FR": "Fév",
		"id-ID": "Feb",
		"it-IT": "Feb",
		"ja-JP": "2月",
		"ko-KO": "2월",
		"pl-PL": "Lut",
		"pt-BR": "Fev",
		"ru-RU": "Фев.",
		"tr-TR": "Şub",
		"ua-UA": "Лют",
		"uz-UZ": "Fev",
		"zh-CN": "二月",
	},

	"Mar": {
		"de-DE": "Mär",
		"en-UK": "Mar",
		"en-US": "Mar",
		"es-ES": "Mar",
		"fa-IR": "مارس",
		"fr-FR": "Mars",
		"id-ID": "Mar",
		"it-IT": "Mar",
		"ja-JP": "3月",
		"ko-KO": "3월",
		"pl-PL": "Mar",
		"pt-BR": "Março",
		"ru-RU": "Мврт",
		"tr-TR": "Mar",
		"ua-UA": "Бер",
		"uz-UZ": "Mar",
		"zh-CN": "三月",
	},

	"Apr": {
		"de-DE": "Apr",
		"en-UK": "Apr",
		"en-US": "Apr",
		"es-ES": "Abr",
		"fa-IR": "آوریل",
		"fr-FR": "Avril",
		"id-ID": "Apr",
		"it-IT": "Apr",
		"ja-JP": "4月",
		"ko-KO": "4월",
		"pl-PL": "Kwi",
		"pt-BR": "Abr",
		"ru-RU": "Апр.",
		"tr-TR": "Nis",
		"ua-UA": "Кві",
		"uz-UZ": "Apr",
		"zh-CN": "四月",
	},

	"May": {
		"de-DE": "Mai",
		"en-UK": "May",
		"en-US": "May",
		"es-ES": "May",
		"fa-IR": "مه",
		"fr-FR": "Mai",
		"id-ID": "Mei",
		"it-IT": "Mag",
		"ja-JP": "5月",
		"ko-KO": "5월",
		"pl-PL": "Maj",
		"pt-BR": "Maio",
		"ru-RU": "Май",
		"tr-TR": "May",
		"ua-UA": "Тра",
		"uz-UZ": "May",
		"zh-CN": "五月",
	},

	"Jun": {
		"de-DE": "Jun",
		"en-UK": "June",
		"en-US": "June",
		"es-ES": "Jun",
		"fa-IR": "ژوئن",
		"fr-FR": "Juin",
		"id-ID": "Jun",
		"it-IT": "Giu",
		"ja-JP": "6月",
		"ko-KO": "6월",
		"pl-PL": "Cze",
		"pt-BR": "Junho",
		"ru-RU": "Июнь",
		"tr-TR": "Haz",
		"ua-UA": "Чер",
		"uz-UZ": "Iyun",
		"zh-CN": "六月",
	},

	"Jul": {
		"de-DE": "Jul",
		"en-UK": "July",
		"en-US": "July",
		"es-ES": "Jul",
		"fa-IR": "ژوئیه",
		"fr-FR": "Juil",
		"id-ID": "Jul",
		"it-IT": "Lug",
		"ja-JP": "7月",
		"ko-KO": "7월",
		"pl-PL": "Lip",
		"pt-BR": "Julho",
		"ru-RU": "Июль",
		"tr-TR": "Tem",
		"ua-UA": "Лип",
		"uz-UZ": "Iyul",
		"zh-CN": "七月",
	},

	"Aug": {
		"de-DE": "Aug",
		"en-UK": "Aug",
		"en-US": "Aug",
		"es-ES": "Ago",
		"fa-IR": "اوت",
		"fr-FR": "Août",
		"id-ID": "Agu",
		"it-IT": "Ago",
		"ja-JP": "8月",
		"ko-KO": "8월",
		"pl-PL": "Sie",
		"pt-BR": "Ago",
		"ru-RU": "Авг.",
		"tr-TR": "Ağu",
		"ua-UA": "Сер",
		"uz-UZ": "Avg",
		"zh-CN": "八月",
	},

	"Sep": {
		"de-DE": "Sep",
		"en-UK": "Sep",
		"en-US": "Sep",
		"es-ES": "Sep",
		"fa-IR": "سپتامبر",
		"fr-FR": "Sep",
		"id-ID": "Sep",
		"it-IT": "Sett",
		"ja-JP": "9月",
		"ko-KO": "9월",
		"pl-PL": "Wrz",
		"pt-BR": "Set",
		"ru-RU": "Сен.",
		"tr-TR": "Eyl",
		"ua-UA": "Вер",
		"uz-UZ": "Sen",
		"zh-CN": "九月",
	},

	"Oct": {
		"de-DE": "Okt",
		"en-UK": "Oct",
		"en-US": "Oct",
		"es-ES": "Oct",
		"fa-IR": "اکتبر",
		"fr-FR": "Oct",
		"id-ID": "Okt",
		"it-IT": "Ott",
		"ja-JP": "10月",
		"ko-KO": "10월",
		"pl-PL": "Paź",
		"pt-BR": "Out",
		"ru-RU": "Окт.",
		"tr-TR": "Eki",
		"ua-UA": "Жов",
		"uz-UZ": "Okt",
		"zh-CN": "十月",
	},

	"Nov": {
		"de-DE": "Nov",
		"en-UK": "Nov",
		"en-US": "Nov",
		"es-ES": "Nov",
		"fa-IR": "نوامبر",
		"fr-FR": "Nov",
		"id-ID": "Nov",
		"it-IT": "Nov",
		"ja-JP": "11月",
		"ko-KO": "11월",
		"pl-PL": "Lis",
		"pt-BR": "Nov",
		"ru-RU": "Нбр.",
		"tr-TR": "Kas",
		"ua-UA": "Лис",
		"uz-UZ": "Noy",
		"zh-CN": "十一月",
	},

	"Dec": {
		"de-DE": "Dez",
		"en-UK": "Dec",
		"en-US": "Dec",
		"es-ES": "Dic",
		"fa-IR": "دسامبر",
		"fr-FR": "Déc",
		"id-ID": "Des",
		"it-IT": "Dic",
		"ja-JP": "12月",
		"ko-KO": "12월",
		"pl-PL": "Gru",
		"pt-BR": "Dez",
		"ru-RU": "Дек.",
		"tr-TR": "Ara",
		"ua-UA": "Гру",
		"uz-UZ": "Dek",
		"zh-CN": "十二月",
	},
	COMMAND_START: {
		"de-DE": "start",
		"en-UK": "start",
		"en-US": "start",
		"es-ES": "inicio",
		"fa-IR": "شروع",
		"fr-FR": "démarrer",
		"id-ID": "mulai",
		"it-IT": "inizio",
		"ja-JP": "スタート",
		"ko-KO": "시작",
		"pl-PL": "start",
		"pt-BR": "iniciar",
		"ru-RU": "старт",
		"tr-TR": "başlat",
		"ua-UA": "старт",
		"uz-UZ": "boshlash",
		"zh-CN": "开始",
	},
	COMMAND_MENU: {
		"de-DE": "menu",
		"en-UK": "menu",
		"en-US": "menu",
		"es-ES": "menú",
		"fa-IR": "منو",
		"fr-FR": "menu",
		"id-ID": "menu",
		"it-IT": "menu",
		"ja-JP": "メニュー",
		"ko-KO": "메뉴",
		"pl-PL": "menu",
		"pt-BR": "menu",
		"ru-RU": "меню",
		"tr-TR": "menü",
		"ua-UA": "меню",
		"uz-UZ": "menyu",
		"zh-CN": "菜单",
	},
	COMMAND_GAVE: {
		"de-DE": "verleihen",
		"en-UK": "gave",
		"en-US": "gave",
		"es-ES": "prestado_a_ti",
		"fa-IR": "قرض_دادن",
		"fr-FR": "donné",
		"id-ID": "memberi",
		"it-IT": "debito",
		"ja-JP": "貸した",
		"ko-KO": "주었다",
		"pl-PL": "dał",
		"pt-BR": "emprestou",
		"ru-RU": "дать",
		"tr-TR": "verdi",
		"ua-UA": "дати",
		"uz-UZ": "berdi",
		"zh-CN": "给予",
	},
	COMMAND_GOT: {
		"de-DE": "anleihen",
		"en-UK": "got",
		"en-US": "got",
		"es-ES": "prestado_por_ti",
		"fa-IR": "قرض_گرفتن",
		"fr-FR": "reçu",
		"id-ID": "mendapat",
		"it-IT": "credito",
		"ja-JP": "借りた",
		"ko-KO": "받았다",
		"pl-PL": "dostał",
		"pt-BR": "recebeu",
		"ru-RU": "взять",
		"tr-TR": "aldı",
		"ua-UA": "взяти",
		"uz-UZ": "oldi",
		"zh-CN": "得到",
	},
	COMMAND_RETURNED: {
		"de-DE": "beglichen",
		"en-UK": "return",
		"en-US": "return",
		"es-ES": "devuelto",
		"fa-IR": "بازگردانده_شده",
		"fr-FR": "retourné",
		"id-ID": "kembali",
		"it-IT": "rientra",
		"ja-JP": "返済",
		"ko-KO": "반환",
		"pl-PL": "zwrócił",
		"pt-BR": "devolveu",
		"ru-RU": "вернуть",
		"tr-TR": "iade",
		"ua-UA": "повернути",
		"uz-UZ": "qaytardi",
		"zh-CN": "归还",
	},
	COMMAND_BALANCE: {
		"de-DE": "ausstehend",
		"en-UK": "balance",
		"en-US": "balance",
		"es-ES": "balance",
		"fa-IR": "تراز",
		"fr-FR": "solde",
		"id-ID": "saldo",
		"it-IT": "bilancio",
		"ja-JP": "残高",
		"ko-KO": "잔액",
		"pl-PL": "saldo",
		"pt-BR": "saldo",
		"ru-RU": "баланс",
		"tr-TR": "bakiye",
		"ua-UA": "баланс",
		"uz-UZ": "balans",
		"zh-CN": "余额",
	},
	COMMAND_HISTORY: {
		"de-DE": "verlauf",
		"en-UK": "history",
		"en-US": "history",
		"es-ES": "cronología",
		"fa-IR": "سوابق",
		"fr-FR": "historique",
		"id-ID": "riwayat",
		"it-IT": "cronologia",
		"ja-JP": "履歴",
		"ko-KO": "기록",
		"pl-PL": "historia",
		"pt-BR": "histórico",
		"ru-RU": "история",
		"tr-TR": "geçmiş",
		"ua-UA": "історія",
		"uz-UZ": "tarix",
		"zh-CN": "历史",
	},
	COMMAND_SETTINGS: {
		"de-DE": "einstellungen",
		"en-UK": "settings",
		"en-US": "settings",
		"es-ES": "ajustes",
		"fa-IR": "تنظیمات",
		"fr-FR": "paramètres",
		"id-ID": "pengaturan",
		"it-IT": "impostazioni",
		"ja-JP": "設定",
		"ko-KO": "설정",
		"pl-PL": "ustawienia",
		"pt-BR": "configurações",
		"ru-RU": "настройки",
		"tr-TR": "ayarlar",
		"ua-UA": "налаштування",
		"uz-UZ": "sozlamalar",
		"zh-CN": "设置",
	},
	COMMAND_HELP: {
		"de-DE": "hilfe",
		"en-UK": "help",
		"en-US": "help",
		"es-ES": "ayuda",
		"fa-IR": "کمک",
		"fr-FR": "aide",
		"id-ID": "bantuan",
		"it-IT": "aiuto",
		"ja-JP": "ヘルプ",
		"ko-KO": "도움말",
		"pl-PL": "pomoc",
		"pt-BR": "ajuda",
		"ru-RU": "помощь",
		"tr-TR": "yardım",
		"ua-UA": "допомога",
		"uz-UZ": "yordam",
		"zh-CN": "帮助",
	},
	COMMAND_CANCEL: {
		"de-DE": "abbrechen",
		"en-UK": "cancel",
		"en-US": "cancel",
		"es-ES": "cancelar",
		"fa-IR": "کنسل",
		"fr-FR": "annuler",
		"id-ID": "batal",
		"it-IT": "annulla",
		"ja-JP": "キャンセル",
		"ko-KO": "취소",
		"pl-PL": "anuluj",
		"pt-BR": "cancelar",
		"ru-RU": "/отменить",
		"tr-TR": "iptal",
		"ua-UA": "скасувати",
		"uz-UZ": "bekor",
		"zh-CN": "取消",
	},
	COMMAND_CLEAR: {
		"de-DE": "leeren",
		"en-UK": "clear",
		"en-US": "clear",
		"es-ES": "borrar",
		"fa-IR": "پاک_کردن",
		"fr-FR": "effacer",
		"id-ID": "bersihkan",
		"it-IT": "chiaro",
		"ja-JP": "クリア",
		"ko-KO": "지우기",
		"pl-PL": "wyczyść",
		"pt-BR": "limpar",
		"ru-RU": "очистить",
		"tr-TR": "temizle",
		"ua-UA": "очистити",
		"uz-UZ": "tozalash",
		"zh-CN": "清除",
	},
	adsCommandTitle: {
		"de-DE": adsCommandTitle,
		"en-UK": adsCommandTitle,
		"en-US": adsCommandTitle,
		"es-ES": adsCommandTitle,
		"fa-IR": adsCommandTitle,
		"fr-FR": adsCommandTitle,
		"id-ID": adsCommandTitle,
		"it-IT": adsCommandTitle,
		"ja-JP": adsCommandTitle,
		"ko-KO": adsCommandTitle,
		"pl-PL": adsCommandTitle,
		"pt-BR": adsCommandTitle,
		"ru-RU": adsCommandTitle,
		"tr-TR": adsCommandTitle,
		"ua-UA": adsCommandTitle,
		"uz-UZ": adsCommandTitle,
		"zh-CN": adsCommandTitle,
	},
	" and ": {
		"de-DE": " und ",
		"en-UK": " and ",
		"en-US": " and ",
		"es-ES": " y ",
		"fa-IR": " و ",
		"fr-FR": " et ",
		"id-ID": " dan ",
		"it-IT": " e ",
		"ja-JP": " と ",
		"ko-KO": " 그리고 ",
		"pl-PL": " i ",
		"pt-BR": " e ",
		"ru-RU": " и ",
		"tr-TR": " ve ",
		"ua-UA": " і ",
		"uz-UZ": " va ",
		"zh-CN": " 和 ",
	},
	"MessageTextOopsSomethingWentWrong": {
		"de-DE": "Ups, etwas ist schiefgelaufen... \xF0\x9F\x98\xB3",
		"en-UK": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"en-US": "Oops, something went wrong... \xF0\x9F\x98\xB3",
		"es-ES": "Ops,  algo ha salido mal... \xF0\x9F\x98\xB3",
		"fa-IR": "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
		"fr-FR": "Oups, quelque chose s'est mal passé... \xF0\x9F\x98\xB3",
		"id-ID": "Ups, ada yang salah... \xF0\x9F\x98\xB3",
		"it-IT": "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
		"ja-JP": "おっと、何か問題が発生しました... \xF0\x9F\x98\xB3",
		"ko-KO": "이런, 문제가 발생했습니다... \xF0\x9F\x98\xB3",
		"pl-PL": "Ups, coś poszło nie tak... \xF0\x9F\x98\xB3",
		"pt-BR": "Ops, algo deu errado... \xF0\x9F\x98\xB3",
		"ru-RU": "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		"tr-TR": "Hata, bir şeyler yanlış gitti... \xF0\x9F\x98\xB3",
		"ua-UA": "Упс, щось пішло не так... \xF0\x9F\x98\xB3",
		"uz-UZ": "Xatolik, nimadir noto'g'ri ketdi... \xF0\x9F\x98\xB3",
		"zh-CN": "哎呀，出了点问题... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		"de-DE": "Wann ist die Schuld fällig?",
		"en-UK": "When is the due date?",
		"en-US": "When is the due date?",
		"es-ES": "¿Cuándo es la fecha de devolución?",
		"fa-IR": "سررسید چه زمانی است؟",
		"fr-FR": "Quelle est la date d'échéance?",
		"id-ID": "Kapan tanggal jatuh tempo?",
		"it-IT": "Data di scadenza?",
		"ja-JP": "期日はいつですか？",
		"ko-KO": "만기일은 언제입니까?",
		"pl-PL": "Kiedy jest termin płatności?",
		"pt-BR": "Qual é a data de vencimento?",
		"ru-RU": "Когда дата возврата?",
		"tr-TR": "Vade tarihi ne zaman?",
		"ua-UA": "Коли дата повернення?",
		"uz-UZ": "To'lov muddati qachon?",
		"zh-CN": "到期日期是什么时候？",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		"de-DE": `Sende mir das Datum, an welches du <b>erneut</b> erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		"en-UK": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"en-US": `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha recordatoria escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
    <i>20.01.2017</i>`,

		"fa-IR": `لطفاً برای تنظیم یادآور بعدی آنرا با متنی با این فرمت ارسال نمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 15 خرداد 1396 ثبت کنید:
    <i>15.03.1396</i>`,

		"fr-FR": `Pour définir la date du prochain rappel, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
    <i>20.01.2017</i>`,

		"id-ID": `Untuk mengatur tanggal pengingat berikutnya, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
    <i>20.01.2017</i>`,

		"it-IT": `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
    <i>20.01.2017</i>`,

		"ja-JP": `次のリマインダーの日付を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
    <i>20.01.2017</i>`,

		"ko-KO": `다음 알림의 날짜를 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
    <i>20.01.2017</i>`,

		"pl-PL": `Aby ustawić datę następnego przypomnienia, wyślij ją jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
    <i>20.01.2017</i>`,

		"pt-BR": `Para definir a data do próximo lembrete, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
    <i>20.01.2017</i>`,

		"ru-RU": `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		"tr-TR": `Bir sonraki hatırlatma için tarih belirlemek üzere <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
    <i>20.01.2017</i>`,

		"ua-UA": `Щоб встановити дату наступного нагадування, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
    <i>20.01.2017</i>`,

		"uz-UZ": `Keyingi eslatma uchun sanani belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
    <i>20.01.2017</i>`,

		"zh-CN": `要设置下一次提醒的日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		"de-DE": `Sende mir das Datum, an welches du erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		"en-UK": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
<i>20.01.2017</i>`,

		"en-US": `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
<i>20.01.2017</i>`,

		"es-ES": `Para establecer la fecha de devolución escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
<i>20.01.2017</i>`,

		"fa-IR": `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید.<i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
<i>20.01.2017</i>`,

		"fr-FR": `Pour définir la date d'échéance, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
<i>20.01.2017</i>`,

		"id-ID": `Untuk mengatur tanggal jatuh tempo, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
<i>20.01.2017</i>`,

		"it-IT": `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
<i>20.01.2017</i>`,

		"ja-JP": `期日を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
<i>20.01.2017</i>`,

		"ko-KO": `만기일을 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
<i>20.01.2017</i>`,

		"pl-PL": `Aby ustawić termin płatności, wyślij go jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
<i>20.01.2017</i>`,

		"pt-BR": `Para definir a data de vencimento, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
<i>20.01.2017</i>`,

		"ru-RU": `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г.отправьте:
<i>20.01.2017</i>`,

		"tr-TR": `Vade tarihini ayarlamak için <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
<i>20.01.2017</i>`,

		"ua-UA": `Щоб встановити дату повернення, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
<i>20.01.2017</i>`,

		"uz-UZ": `To'lov muddatini belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
<i>20.01.2017</i>`,

		"zh-CN": `要设置到期日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
<i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		"de-DE": "Entschuldigung, aber mit diesem Datum stimmt etwas nicht.",
		"en-UK": "Sorry, there is something wrong with the date you've provided.",
		"en-US": "Sorry, there is something wrong with the date you've provided.",
		"es-ES": "Lo siento, algo no es correcto con la fecha que has puesto",
		"fa-IR": "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
		"fr-FR": "Désolé, il y a un problème avec la date que vous avez fournie.",
		"id-ID": "Maaf, ada yang salah dengan tanggal yang Anda berikan.",
		"it-IT": "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
		"ja-JP": "申し訳ありませんが、提供された日付に問題があります。",
		"ko-KO": "죄송합니다. 제공하신 날짜에 문제가 있습니다.",
		"pl-PL": "Przepraszamy, ale z podaną datą jest coś nie tak.",
		"pt-BR": "Desculpe, há algo errado com a data que você forneceu.",
		"ru-RU": "Извините, что-то не так с датой которую вы отправили.",
		"tr-TR": "Üzgünüz, verdiğiniz tarihte bir sorun var.",
		"ua-UA": "Вибачте, але з датою щось не так.",
		"uz-UZ": "Kechirasiz, siz taqdim etgan sana bilan bog'liq muammo bor.",
		"zh-CN": "抱歉，您提供的日期有问题。",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		"de-DE": "Nicht erinnern",
		"en-UK": "No reminder",
		"en-US": "No reminder",
		"es-ES": "No recordar",
		"fa-IR": "بدون یادآور",
		"fr-FR": "Pas de rappel",
		"id-ID": "Tidak ada pengingat",
		"it-IT": "Nessun promemoria",
		"ja-JP": "リマインダーなし",
		"ko-KO": "알림 없음",
		"pl-PL": "Bez przypomnienia",
		"pt-BR": "Sem lembrete",
		"ru-RU": "Не напоминать",
		"tr-TR": "Hatırlatma yok",
		"ua-UA": "Не нагадувати",
		"uz-UZ": "Eslatma yo'q",
		"zh-CN": "无提醒",
	},
	COMMAND_TEXT_TOMORROW: {
		"de-DE": "Morgen",
		"en-UK": "Tomorrow",
		"en-US": "Tomorrow",
		"es-ES": "Mañana",
		"fa-IR": "فردا",
		"fr-FR": "Demain",
		"id-ID": "Besok",
		"it-IT": "Domani",
		"ja-JP": "明日",
		"ko-KO": "내일",
		"pl-PL": "Jutro",
		"pt-BR": "Amanhã",
		"ru-RU": "Завтра",
		"tr-TR": "Yarın",
		"ua-UA": "Завтра",
		"uz-UZ": "Ertaga",
		"zh-CN": "明天",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		"de-DE": "Übermorgen",
		"en-UK": "Day after tomorrow",
		"en-US": "Day after tomorrow",
		"es-ES": "Pasada mañana",
		"fa-IR": "پس فردا",
		"fr-FR": "Après-demain",
		"id-ID": "Lusa",
		"it-IT": "Dopo domani",
		"ja-JP": "明後日",
		"ko-KO": "모레",
		"pl-PL": "Pojutrze",
		"pt-BR": "Depois de amanhã",
		"ru-RU": "Послезавтра",
		"tr-TR": "Öbür gün",
		"ua-UA": "Післязавтра",
		"uz-UZ": "Indinga",
		"zh-CN": "后天",
	},
	COMMAND_TEXT_THIS_WEEK: {
		"de-DE": "Diese Woche",
		"en-UK": "This week",
		"en-US": "This week",
		"es-ES": "Esta semana",
		"fa-IR": "این هفته",
		"fr-FR": "Cette semaine",
		"id-ID": "Minggu ini",
		"it-IT": "Questa settimana",
		"ja-JP": "今週",
		"ko-KO": "이번 주",
		"pl-PL": "W tym tygodniu",
		"pt-BR": "Esta semana",
		"ru-RU": "На этой неделе",
		"tr-TR": "Bu hafta",
		"ua-UA": "Цього тижня",
		"uz-UZ": "Shu hafta",
		"zh-CN": "本周",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		"de-DE": "Ja, es hat eine Frist!",
		"en-UK": "Yes, it has a deadline!",
		"en-US": "Yes, it has a deadline!", //TODO(US)
		"es-ES": "Sí, hay una fecha de devolución!",
		"fa-IR": "بله، دارای آخرین فرصت می باشد!",
		"fr-FR": "Yes, it has a deadline!", //TODO(FR)
		"id-ID": "Yes, it has a deadline!", //TODO(ID)
		"it-IT": "Si, c'e' una data di scadenza",
		"ja-JP": "Yes, it has a deadline!", //TODO(JP)
		"ko-KO": "Yes, it has a deadline!", //TODO(KO)
		"pl-PL": "Yes, it has a deadline!", //TODO(PL)
		"pt-BR": "Yes, it has a deadline!", //TODO(BR)
		"ru-RU": "Да, есть срок возврата!",
		"tr-TR": "Yes, it has a deadline!", //TODO(TR)
		"ua-UA": "Yes, it has a deadline!", //TODO(UA)
		"uz-UZ": "Yes, it has a deadline!", //TODO(UZ)
		"zh-CN": "Yes, it has a deadline!", //TODO(CN)
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		"de-DE": "Nein, sobald möglich.",
		"en-UK": "No, whenever is fine.",
		"en-US": "No, whenever is fine.", //TODO(US)
		"es-ES": "No, sin fecha límite.",
		"fa-IR": "خیر، هر زمانی مناسب است.",
		"fr-FR": "No, whenever is fine.", //TODO(FR)
		"id-ID": "No, whenever is fine.", //TODO(ID)
		"it-IT": "No, nessuna scadenza",
		"ja-JP": "No, whenever is fine.", //TODO(JP)
		"ko-KO": "No, whenever is fine.", //TODO(KO)
		"pl-PL": "No, whenever is fine.", //TODO(PL)
		"pt-BR": "No, whenever is fine.", //TODO(BR)
		"ru-RU": "Нет, срок возврата не важен.",
		"tr-TR": "No, whenever is fine.", //TODO(TR)
		"ua-UA": "No, whenever is fine.", //TODO(UA)
		"uz-UZ": "No, whenever is fine.", //TODO(UZ)
		"zh-CN": "No, whenever is fine.", //TODO(CN)
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		"de-DE": "Unbefristet",
		"en-UK": "Whenever is fine",
		"en-US": "Whenever is fine", //TODO(US)
		"es-ES": "Cualquier día",
		"fa-IR": "هر زمانی مناسب است.",
		"fr-FR": "Whenever is fine", //TODO(FR)
		"id-ID": "Whenever is fine", //TODO(ID)
		"it-IT": "No, Nessuna scadenza",
		"ja-JP": "Whenever is fine", //TODO(JP)
		"ko-KO": "Whenever is fine", //TODO(KO)
		"pl-PL": "Whenever is fine", //TODO(PL)
		"pt-BR": "Whenever is fine", //TODO(BR)
		"ru-RU": "Когда-нибудь",
		"tr-TR": "Whenever is fine", //TODO(TR)
		"ua-UA": "Whenever is fine", //TODO(UA)
		"uz-UZ": "Whenever is fine", //TODO(UZ)
		"zh-CN": "Whenever is fine", //TODO(CN)
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		"de-DE": "In wenigen Minuten",
		"en-UK": "In few minutes",
		"en-US": "In few minutes", //TODO(US)
		"es-ES": "En unos minutos",
		"fa-IR": "در چند دقیقه",
		"fr-FR": "In few minutes", //TODO(FR)
		"id-ID": "In few minutes", //TODO(ID)
		"it-IT": "Fra alcuni minuti",
		"ja-JP": "In few minutes", //TODO(JP)
		"ko-KO": "In few minutes", //TODO(KO)
		"pl-PL": "In few minutes", //TODO(PL)
		"pt-BR": "In few minutes", //TODO(BR)
		"ru-RU": "Через минуту",
		"tr-TR": "In few minutes", //TODO(TR)
		"ua-UA": "In few minutes", //TODO(UA)
		"uz-UZ": "In few minutes", //TODO(UZ)
		"zh-CN": "In few minutes", //TODO(CN)
	},
	COMMAND_TEXT_IN_1_WEEK: {
		"de-DE": "In einer Woche",
		"en-UK": "In 1 week",
		"en-US": "In 1 week", //TODO(US)
		"es-ES": "En una semana",
		"fa-IR": "ظرف یک هفته",
		"fr-FR": "In 1 week", //TODO(FR)
		"id-ID": "In 1 week", //TODO(ID)
		"it-IT": "Fra una settimana",
		"ja-JP": "In 1 week", //TODO(JP)
		"ko-KO": "In 1 week", //TODO(KO)
		"pl-PL": "In 1 week", //TODO(PL)
		"pt-BR": "In 1 week", //TODO(BR)
		"ru-RU": "Через неделю",
		"tr-TR": "In 1 week", //TODO(TR)
		"ua-UA": "In 1 week", //TODO(UA)
		"uz-UZ": "In 1 week", //TODO(UZ)
		"zh-CN": "In 1 week", //TODO(CN)
	},
	COMMAND_TEXT_IN_1_MONTH: {
		"de-DE": "In einem Monat",
		"en-UK": "In 1 month",
		"en-US": "In 1 month", //TODO(US)
		"es-ES": "En un mes",
		"fa-IR": "ظرف یک ماه",
		"fr-FR": "In 1 month", //TODO(FR)
		"id-ID": "In 1 month", //TODO(ID)
		"it-IT": "Fra un mese",
		"ja-JP": "In 1 month", //TODO(JP)
		"ko-KO": "In 1 month", //TODO(KO)
		"pl-PL": "In 1 month", //TODO(PL)
		"pt-BR": "In 1 month", //TODO(BR)
		"ru-RU": "Через месяц",
		"tr-TR": "In 1 month", //TODO(TR)
		"ua-UA": "In 1 month", //TODO(UA)
		"uz-UZ": "In 1 month", //TODO(UZ)
		"zh-CN": "In 1 month", //TODO(CN)
	},
	COMMAND_TEXT_SET_DATE: {
		"de-DE": "Datum setzen",
		"en-UK": "Set date",
		"en-US": "Set date", //TODO(US)
		"es-ES": "Establecer la fecha",
		"fa-IR": "تاریخ را تنظیم کنید",
		"fr-FR": "Set date", //TODO(FR)
		"id-ID": "Set date", //TODO(ID)
		"it-IT": "Imposta la data",
		"ja-JP": "Set date", //TODO(JP)
		"ko-KO": "Set date", //TODO(KO)
		"pl-PL": "Set date", //TODO(PL)
		"pt-BR": "Set date", //TODO(BR)
		"ru-RU": "Задать дату",
		"tr-TR": "Set date", //TODO(TR)
		"ua-UA": "Set date", //TODO(UA)
		"uz-UZ": "Set date", //TODO(UZ)
		"zh-CN": "Set date", //TODO(CN)
	},
	COMMAND_TEXT_MONDAY: {
		"de-DE": "Montag",
		"en-UK": "Monday",
		"en-US": "Monday", //TODO(US)
		"es-ES": "Lunes",
		"fa-IR": "دوشنبه",
		"fr-FR": "Monday", //TODO(FR)
		"id-ID": "Monday", //TODO(ID)
		"it-IT": "Lunedi'",
		"ja-JP": "Monday", //TODO(JP)
		"ko-KO": "Monday", //TODO(KO)
		"pl-PL": "Monday", //TODO(PL)
		"pt-BR": "Monday", //TODO(BR)
		"ru-RU": "Понедельник",
		"tr-TR": "Monday", //TODO(TR)
		"ua-UA": "Monday", //TODO(UA)
		"uz-UZ": "Monday", //TODO(UZ)
		"zh-CN": "Monday", //TODO(CN)
	},
	COMMAND_TEXT_TUESDAY: {
		"de-DE": "Dienstag",
		"en-UK": "Tuesday",
		"en-US": "Tuesday", //TODO(US)
		"es-ES": "Martes",
		"fa-IR": "سه شنبه",
		"fr-FR": "Tuesday", //TODO(FR)
		"id-ID": "Tuesday", //TODO(ID)
		"it-IT": "Martedi'",
		"ja-JP": "Tuesday", //TODO(JP)
		"ko-KO": "Tuesday", //TODO(KO)
		"pl-PL": "Tuesday", //TODO(PL)
		"pt-BR": "Tuesday", //TODO(BR)
		"ru-RU": "Вторник",
		"tr-TR": "Tuesday", //TODO(TR)
		"ua-UA": "Tuesday", //TODO(UA)
		"uz-UZ": "Tuesday", //TODO(UZ)
		"zh-CN": "Tuesday", //TODO(CN)
	},
	COMMAND_TEXT_WEDNESDAY: {
		"de-DE": "Mittwoch",
		"en-UK": "Wednesday",
		"en-US": "Wednesday", //TODO(US)
		"es-ES": "Miercoles",
		"fa-IR": "چهارشنبه",
		"fr-FR": "Wednesday", //TODO(FR)
		"id-ID": "Wednesday", //TODO(ID)
		"it-IT": "Mercoledi'",
		"ja-JP": "Wednesday", //TODO(JP)
		"ko-KO": "Wednesday", //TODO(KO)
		"pl-PL": "Wednesday", //TODO(PL)
		"pt-BR": "Wednesday", //TODO(BR)
		"ru-RU": "Среда",
		"tr-TR": "Wednesday", //TODO(TR)
		"ua-UA": "Wednesday", //TODO(UA)
		"uz-UZ": "Wednesday", //TODO(UZ)
		"zh-CN": "Wednesday", //TODO(CN)
	},
	COMMAND_TEXT_THURSDAY: {
		"de-DE": "Donnerstag",
		"en-UK": "Thursday",
		"en-US": "Thursday", //TODO(US)
		"es-ES": "Jueves",
		"fa-IR": "پنج شنبه",
		"fr-FR": "Thursday", //TODO(FR)
		"id-ID": "Thursday", //TODO(ID)
		"it-IT": "Giovedi'",
		"ja-JP": "Thursday", //TODO(JP)
		"ko-KO": "Thursday", //TODO(KO)
		"pl-PL": "Thursday", //TODO(PL)
		"pt-BR": "Thursday", //TODO(BR)
		"ru-RU": "Четверг",
		"tr-TR": "Thursday", //TODO(TR)
		"ua-UA": "Thursday", //TODO(UA)
		"uz-UZ": "Thursday", //TODO(UZ)
		"zh-CN": "Thursday", //TODO(CN)
	},
	COMMAND_TEXT_FRIDAY: {
		"de-DE": "Freitag",
		"en-UK": "Friday",
		"en-US": "Friday", //TODO(US)
		"es-ES": "Viernes",
		"fa-IR": "جمعه",
		"fr-FR": "Friday", //TODO(FR)
		"id-ID": "Friday", //TODO(ID)
		"it-IT": "Venerdi'",
		"ja-JP": "Friday", //TODO(JP)
		"ko-KO": "Friday", //TODO(KO)
		"pl-PL": "Friday", //TODO(PL)
		"pt-BR": "Friday", //TODO(BR)
		"ru-RU": "Пятница",
		"tr-TR": "Friday", //TODO(TR)
		"ua-UA": "Friday", //TODO(UA)
		"uz-UZ": "Friday", //TODO(UZ)
		"zh-CN": "Friday", //TODO(CN)
	},
	COMMAND_TEXT_SATURDAY: {
		"de-DE": "Samstag",
		"en-UK": "Saturday",
		"en-US": "Saturday", //TODO(US)
		"es-ES": "Sabado",
		"fa-IR": "شنبه",
		"fr-FR": "Saturday", //TODO(FR)
		"id-ID": "Saturday", //TODO(ID)
		"it-IT": "Sabato",
		"ja-JP": "Saturday", //TODO(JP)
		"ko-KO": "Saturday", //TODO(KO)
		"pl-PL": "Saturday", //TODO(PL)
		"pt-BR": "Saturday", //TODO(BR)
		"ru-RU": "Суббота",
		"tr-TR": "Saturday", //TODO(TR)
		"ua-UA": "Saturday", //TODO(UA)
		"uz-UZ": "Saturday", //TODO(UZ)
		"zh-CN": "Saturday", //TODO(CN)
	},
	COMMAND_TEXT_SUNDAY: {
		"de-DE": "Sonntag",
		"en-UK": "Sunday",
		"en-US": "Sunday", //TODO(US)
		"es-ES": "Domingo",
		"fa-IR": "یکشنبه",
		"fr-FR": "Sunday", //TODO(FR)
		"id-ID": "Sunday", //TODO(ID)
		"it-IT": "Domenica",
		"ja-JP": "Sunday", //TODO(JP)
		"ko-KO": "Sunday", //TODO(KO)
		"pl-PL": "Sunday", //TODO(PL)
		"pt-BR": "Sunday", //TODO(BR)
		"ru-RU": "Воскресенье",
		"tr-TR": "Sunday", //TODO(TR)
		"ua-UA": "Sunday", //TODO(UA)
		"uz-UZ": "Sunday", //TODO(UZ)
		"zh-CN": "Sunday", //TODO(CN)
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		"de-DE": "Keine Quittung schicken",
		"en-UK": "Do not send the receipt",
		"en-US": "Do not send the receipt", //TODO(US)
		"es-ES": "No enviar el recibo",
		"fa-IR": "رسید را ارسال نکنید",
		"fr-FR": "Do not send the receipt", //TODO(FR)
		"id-ID": "Do not send the receipt", //TODO(ID)
		"it-IT": "Non inviare la ricevuta",
		"ja-JP": "Do not send the receipt", //TODO(JP)
		"ko-KO": "Do not send the receipt", //TODO(KO)
		"pl-PL": "Do not send the receipt", //TODO(PL)
		"pt-BR": "Do not send the receipt", //TODO(BR)
		"ru-RU": "Не отправлять квитанцию",
		"tr-TR": "Do not send the receipt", //TODO(TR)
		"ua-UA": "Do not send the receipt", //TODO(UA)
		"uz-UZ": "Do not send the receipt", //TODO(UZ)
		"zh-CN": "Do not send the receipt", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		"de-DE": "Du hast dich gegen eine Quittung entschieden.",
		"en-UK": "You've decided not to send the receipt.",
		"en-US": "You've decided not to send the receipt.", //TODO(US)
		"es-ES": "Has decidido no enviar el recibo",
		"fa-IR": "شما تصمیم گرفتید که رسید را ارسال نکنید.",
		"fr-FR": "You've decided not to send the receipt.", //TODO(FR)
		"id-ID": "You've decided not to send the receipt.", //TODO(ID)
		"it-IT": "Hai scelto di non inviare la ricevuta",
		"ja-JP": "You've decided not to send the receipt.", //TODO(JP)
		"ko-KO": "You've decided not to send the receipt.", //TODO(KO)
		"pl-PL": "You've decided not to send the receipt.", //TODO(PL)
		"pt-BR": "You've decided not to send the receipt.", //TODO(BR)
		"ru-RU": "Вы решили не отправлять квитанцию.",
		"tr-TR": "You've decided not to send the receipt.", //TODO(TR)
		"ua-UA": "You've decided not to send the receipt.", //TODO(UA)
		"uz-UZ": "You've decided not to send the receipt.", //TODO(UZ)
		"zh-CN": "You've decided not to send the receipt.", //TODO(CN)
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		"de-DE": "Ich habe meine Meinung geändert",
		"en-UK": "I've changed my mind",
		"en-US": "I've changed my mind", //TODO(US)
		"es-ES": "He cambiado de opinion",
		"fa-IR": "نظرم را عوض کردم.",
		"fr-FR": "I've changed my mind", //TODO(FR)
		"id-ID": "I've changed my mind", //TODO(ID)
		"it-IT": "Ho cambiato idea",
		"ja-JP": "I've changed my mind", //TODO(JP)
		"ko-KO": "I've changed my mind", //TODO(KO)
		"pl-PL": "I've changed my mind", //TODO(PL)
		"pt-BR": "I've changed my mind", //TODO(BR)
		"ru-RU": "Я передумал(а)",
		"tr-TR": "I've changed my mind", //TODO(TR)
		"ua-UA": "I've changed my mind", //TODO(UA)
		"uz-UZ": "I've changed my mind", //TODO(UZ)
		"zh-CN": "I've changed my mind", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		"de-DE": "Senden per Telegram",
		"en-UK": "Send by Telegram",
		"en-US": "Send by Telegram", //TODO(US)
		"es-ES": "Enviar a través de Telegram",
		"fa-IR": "با تلگرام ارسال شود",
		"fr-FR": "Send by Telegram", //TODO(FR)
		"id-ID": "Send by Telegram", //TODO(ID)
		"it-IT": "Invia tramite Telegram",
		"ja-JP": "Send by Telegram", //TODO(JP)
		"ko-KO": "Send by Telegram", //TODO(KO)
		"pl-PL": "Send by Telegram", //TODO(PL)
		"pt-BR": "Send by Telegram", //TODO(BR)
		"ru-RU": "Отправить через Telelgram",
		"tr-TR": "Send by Telegram", //TODO(TR)
		"ua-UA": "Send by Telegram", //TODO(UA)
		"uz-UZ": "Send by Telegram", //TODO(UZ)
		"zh-CN": "Send by Telegram", //TODO(CN)
	},
	COMMAND_TEXT_GET_LINK_FOR_RECEIPT_IN_TELEGRAM: {
		"de-DE": "Erhalten sie einen link für eine quittung in Telegram", // TODO(DE) verify
		"en-UK": "Get link for a receipt in Telegram",
		"en-US": "Get link for a receipt in Telegram",        //TODO(US)
		"es-ES": "Obtener enlace para recibirlo en Telegram", // TODO(ES) verify
		"fa-IR": "دریافت پیوند برای دریافت در Telegram",      // TODO(FA) verify
		"fr-FR": "Get link for a receipt in Telegram",        //TODO(FR)
		"id-ID": "Get link for a receipt in Telegram",        //TODO(ID)
		"it-IT": "Link per la ricevuta nel Telegram",         // TODO(IT)
		"ja-JP": "Get link for a receipt in Telegram",        //TODO(JP)
		"ko-KO": "Get link for a receipt in Telegram",        //TODO(KO)
		"pl-PL": "Get link for a receipt in Telegram",        //TODO(PL)
		"pt-BR": "Get link for a receipt in Telegram",        //TODO(BR)
		"ru-RU": "Ссылка для квитанции в Телеграмме",
		"tr-TR": "Get link for a receipt in Telegram", //TODO(TR)
		"ua-UA": "Get link for a receipt in Telegram", //TODO(UA)
		"uz-UZ": "Get link for a receipt in Telegram", //TODO(UZ)
		"zh-CN": "Get link for a receipt in Telegram", //TODO(CN)
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		"de-DE": "Senden per FB, WhatsApp, Viber, etc.",
		"en-UK": "Send by FB, WhatsApp, Viber, etc.",
		"en-US": "Send by FB, WhatsApp, Viber, etc.", //TODO(US)
		"es-ES": "Enviar a través de FB, WhatsApp, Viber, etc.",
		"fa-IR": "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
		"fr-FR": "Send by FB, WhatsApp, Viber, etc.", //TODO(FR)
		"id-ID": "Send by FB, WhatsApp, Viber, etc.", //TODO(ID)
		"it-IT": "Invia con FB, WhatsCrap, Viber, etc.",
		"ja-JP": "Send by FB, WhatsApp, Viber, etc.", //TODO(JP)
		"ko-KO": "Send by FB, WhatsApp, Viber, etc.", //TODO(KO)
		"pl-PL": "Send by FB, WhatsApp, Viber, etc.", //TODO(PL)
		"pt-BR": "Send by FB, WhatsApp, Viber, etc.", //TODO(BR)
		"ru-RU": "Отправить через Viber, VK, FB, ...",
		"tr-TR": "Send by FB, WhatsApp, Viber, etc.", //TODO(TR)
		"ua-UA": "Send by FB, WhatsApp, Viber, etc.", //TODO(UA)
		"uz-UZ": "Send by FB, WhatsApp, Viber, etc.", //TODO(UZ)
		"zh-CN": "Send by FB, WhatsApp, Viber, etc.", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		"de-DE": "Senden per SMS",
		"en-UK": "Send by SMS",
		"en-US": "Send by SMS", //TODO(US)
		"es-ES": "Enviar a través de SMS",
		"fa-IR": "با پیام کوتاه ارسال شود",
		"fr-FR": "Send by SMS", //TODO(FR)
		"id-ID": "Send by SMS", //TODO(ID)
		"it-IT": "Invia tramite SMS",
		"ja-JP": "Send by SMS", //TODO(JP)
		"ko-KO": "Send by SMS", //TODO(KO)
		"pl-PL": "Send by SMS", //TODO(PL)
		"pt-BR": "Send by SMS", //TODO(BR)
		"ru-RU": "Отправить через SMS",
		"tr-TR": "Send by SMS", //TODO(TR)
		"ua-UA": "Send by SMS", //TODO(UA)
		"uz-UZ": "Send by SMS", //TODO(UZ)
		"zh-CN": "Send by SMS", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		"de-DE": "Senden per VK.com",
		"en-UK": "Send throw VK.com",
		"en-US": "Send throw VK.com", //TODO(US)
		"es-ES": "Enviar vía VK.com",
		"fa-IR": "ارسال شود VK.com از طریق ",
		"fr-FR": "Send throw VK.com", //TODO(FR)
		"id-ID": "Send throw VK.com", //TODO(ID)
		"it-IT": "Invia tramite VK.com (Facebook russo)",
		"ja-JP": "Send throw VK.com", //TODO(JP)
		"ko-KO": "Send throw VK.com", //TODO(KO)
		"pl-PL": "Send throw VK.com", //TODO(PL)
		"pt-BR": "Send throw VK.com", //TODO(BR)
		"ru-RU": "Отправить через ВКонтакте",
		"tr-TR": "Send throw VK.com", //TODO(TR)
		"ua-UA": "Send throw VK.com", //TODO(UA)
		"uz-UZ": "Send throw VK.com", //TODO(UZ)
		"zh-CN": "Send throw VK.com", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		"de-DE": "Senden per OK",
		"en-UK": "Send throw OK",
		"en-US": "Send throw OK", //TODO(US)
		"es-ES": "Enviar a través de OK",
		"fa-IR": "ارسال شود OK از طریق ",
		"fr-FR": "Send throw OK", //TODO(FR)
		"id-ID": "Send throw OK", //TODO(ID)
		"it-IT": "Invia tramite OK",
		"ja-JP": "Send throw OK", //TODO(JP)
		"ko-KO": "Send throw OK", //TODO(KO)
		"pl-PL": "Send throw OK", //TODO(PL)
		"pt-BR": "Send throw OK", //TODO(BR)
		"ru-RU": "Отправить через Одноклассники",
		"tr-TR": "Send throw OK", //TODO(TR)
		"ua-UA": "Send throw OK", //TODO(UA)
		"uz-UZ": "Send throw OK", //TODO(UZ)
		"zh-CN": "Send throw OK", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		"de-DE": "Senden per Facebook",
		"en-UK": "Send throw Facebook",
		"en-US": "Send throw Facebook", //TODO(US)
		"es-ES": "Enviar a través de Facebook",
		"fa-IR": "از طریق فیسبوک ارسال شود.",
		"fr-FR": "Send throw Facebook", //TODO(FR)
		"id-ID": "Send throw Facebook", //TODO(ID)
		"it-IT": "Invia tramite Facebook",
		"ja-JP": "Send throw Facebook", //TODO(JP)
		"ko-KO": "Send throw Facebook", //TODO(KO)
		"pl-PL": "Send throw Facebook", //TODO(PL)
		"pt-BR": "Send throw Facebook", //TODO(BR)
		"ru-RU": "Отправить через Facebook",
		"tr-TR": "Send throw Facebook", //TODO(TR)
		"ua-UA": "Send throw Facebook", //TODO(UA)
		"uz-UZ": "Send throw Facebook", //TODO(UZ)
		"zh-CN": "Send throw Facebook", //TODO(CN)
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		"de-DE": "Senden per Twitter",
		"en-UK": "Send throw Twitter",
		"en-US": "Send throw Twitter", //TODO(US)
		"es-ES": "Enviar a través de Twitter",
		"fa-IR": "از طریق توئیتر ارسال شود.",
		"fr-FR": "Send throw Twitter", //TODO(FR)
		"id-ID": "Send throw Twitter", //TODO(ID)
		"it-IT": "Invia tramite Twitter",
		"ja-JP": "Send throw Twitter", //TODO(JP)
		"ko-KO": "Send throw Twitter", //TODO(KO)
		"pl-PL": "Send throw Twitter", //TODO(PL)
		"pt-BR": "Send throw Twitter", //TODO(BR)
		"ru-RU": "Отправить через Twitter",
		"tr-TR": "Send throw Twitter", //TODO(TR)
		"ua-UA": "Send throw Twitter", //TODO(UA)
		"uz-UZ": "Send throw Twitter", //TODO(UZ)
		"zh-CN": "Send throw Twitter", //TODO(CN)
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		"de-DE": "Sendung der Quittung per Telegram abbrechen",
		"en-UK": "Cancel sending receipt by Telegram",
		"en-US": "Cancel sending receipt by Telegram", //TODO(US)
		"es-ES": "Cancelar el envío a través de Telegram",
		"fa-IR": "ارسال رسید با تلگرام کنسل شود",
		"fr-FR": "Cancel sending receipt by Telegram", //TODO(FR)
		"id-ID": "Cancel sending receipt by Telegram", //TODO(ID)
		"it-IT": "Annulla l'invio tramite Telegram",
		"ja-JP": "Cancel sending receipt by Telegram", //TODO(JP)
		"ko-KO": "Cancel sending receipt by Telegram", //TODO(KO)
		"pl-PL": "Cancel sending receipt by Telegram", //TODO(PL)
		"pt-BR": "Cancel sending receipt by Telegram", //TODO(BR)
		"ru-RU": "Отменить отправку через Telegram",
		"tr-TR": "Cancel sending receipt by Telegram", //TODO(TR)
		"ua-UA": "Cancel sending receipt by Telegram", //TODO(UA)
		"uz-UZ": "Cancel sending receipt by Telegram", //TODO(UZ)
		"zh-CN": "Cancel sending receipt by Telegram", //TODO(CN)
	},
	MAIN_MENU: {
		"de-DE": "Hauptmenü",
		"en-UK": "Main menu",
		"en-US": "Main menu",
		"es-ES": "Menú principal",
		"fa-IR": "منوی اصلی",
		"fr-FR": "Menu principal",
		"id-ID": "Menu utama",
		"it-IT": "Menu principale",
		"ja-JP": "メインメニュー",
		"ko-KO": "메인 메뉴",
		"pl-PL": "Menu główne",
		"pt-BR": "Menu principal",
		"ru-RU": "Главное меню",
		"tr-TR": "Ana menü",
		"ua-UA": "Головне меню",
		"uz-UZ": "Asosiy menyu",
		"zh-CN": "主菜单",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		"de-DE": "Menü /menu",
		"en-UK": "Main /menu",
		"en-US": "Main /menu", //TODO(US)
		"es-ES": "Inicio /menú",
		"fa-IR": "/منو ی اصلی ",
		"fr-FR": "Main /menu", //TODO(FR)
		"id-ID": "Main /menu", //TODO(ID)
		"it-IT": "Menu' /menu",
		"ja-JP": "Main /menu", //TODO(JP)
		"ko-KO": "Main /menu", //TODO(KO)
		"pl-PL": "Main /menu", //TODO(PL)
		"pt-BR": "Main /menu", //TODO(BR)
		"ru-RU": "Главное /меню",
		"tr-TR": "Main /menu", //TODO(TR)
		"ua-UA": "Main /menu", //TODO(UA)
		"uz-UZ": "Main /menu", //TODO(UZ)
		"zh-CN": "Main /menu", //TODO(CN)
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		"de-DE": "Nichts zum abbrechen.",
		"en-UK": "Nothing to cancel.",
		"en-US": "Nothing to cancel.", //TODO(US)
		"es-ES": "No hay nada que anular.",
		"fa-IR": "چیزی برای کنسل شدن وجود ندارد",
		"fr-FR": "Nothing to cancel.", //TODO(FR)
		"id-ID": "Nothing to cancel.", //TODO(ID)
		"it-IT": "Nulla da annullare.",
		"ja-JP": "Nothing to cancel.", //TODO(JP)
		"ko-KO": "Nothing to cancel.", //TODO(KO)
		"pl-PL": "Nothing to cancel.", //TODO(PL)
		"pt-BR": "Nothing to cancel.", //TODO(BR)
		"ru-RU": "Нечего отменять.",
		"tr-TR": "Nothing to cancel.", //TODO(TR)
		"ua-UA": "Nothing to cancel.", //TODO(UA)
		"uz-UZ": "Nothing to cancel.", //TODO(UZ)
		"zh-CN": "Nothing to cancel.", //TODO(CN)
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		"de-DE": "Erstellung des Schuldscheins abgebrochen.",
		"en-UK": "Creation of debt record has been canceled.",
		"en-US": "Creation of debt record has been canceled.", //TODO(US)
		"es-ES": "La creación del recordatorio se ha cancelado.",
		"fa-IR": "ایجاد سابقه بدهی کنسل شد.",
		"fr-FR": "Creation of debt record has been canceled.", //TODO(FR)
		"id-ID": "Creation of debt record has been canceled.", //TODO(ID)
		"it-IT": "Creazione record annullata",
		"ja-JP": "Creation of debt record has been canceled.", //TODO(JP)
		"ko-KO": "Creation of debt record has been canceled.", //TODO(KO)
		"pl-PL": "Creation of debt record has been canceled.", //TODO(PL)
		"pt-BR": "Creation of debt record has been canceled.", //TODO(BR)
		"ru-RU": "Создание записи о долге отменено.",
		"tr-TR": "Creation of debt record has been canceled.", //TODO(TR)
		"ua-UA": "Creation of debt record has been canceled.", //TODO(UA)
		"uz-UZ": "Creation of debt record has been canceled.", //TODO(UZ)
		"zh-CN": "Creation of debt record has been canceled.", //TODO(CN)
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		"de-DE": "Zeige alle...",
		"en-UK": "Show all...",
		"en-US": "Show all...", //TODO(US)
		"es-ES": "Mostrar todo...",
		"fa-IR": "نمایش تمام موارد ...",
		"fr-FR": "Show all...", //TODO(FR)
		"id-ID": "Show all...", //TODO(ID)
		"it-IT": "Mostra tutto...",
		"ja-JP": "Show all...", //TODO(JP)
		"ko-KO": "Show all...", //TODO(KO)
		"pl-PL": "Show all...", //TODO(PL)
		"pt-BR": "Show all...", //TODO(BR)
		"ru-RU": "Показать все...",
		"tr-TR": "Show all...", //TODO(TR)
		"ua-UA": "Show all...", //TODO(UA)
		"uz-UZ": "Show all...", //TODO(UZ)
		"zh-CN": "Show all...", //TODO(CN)
	},
	COMMAND_TEXT_CONTACTS: {
		"de-DE": "Kontakte",
		"en-UK": "Contacts",
		"en-US": "Contacts", //TODO(US)
		"es-ES": "Contactos",
		"fa-IR": "لیست تماس",
		"fr-FR": "Contacts", //TODO(FR)
		"id-ID": "Contacts", //TODO(ID)
		"it-IT": "Сontatti",
		"ja-JP": "Contacts", //TODO(JP)
		"ko-KO": "Contacts", //TODO(KO)
		"pl-PL": "Contacts", //TODO(PL)
		"pt-BR": "Contacts", //TODO(BR)
		"ru-RU": "Контакты",
		"tr-TR": "Contacts", //TODO(TR)
		"ua-UA": "Contacts", //TODO(UA)
		"uz-UZ": "Contacts", //TODO(UZ)
		"zh-CN": "Contacts", //TODO(CN)
	},
	COMMAND_TEXT_REFRESH: {
		"de-DE": "Aktualisieren",
		"en-UK": "Refresh",
		"en-US": "Refresh", //TODO(US)
		"es-ES": "Recargar",
		"fa-IR": "تازه کردن",
		"fr-FR": "Refresh", //TODO(FR)
		"id-ID": "Refresh", //TODO(ID)
		"it-IT": "Ricaricare",
		"ja-JP": "Refresh", //TODO(JP)
		"ko-KO": "Refresh", //TODO(KO)
		"pl-PL": "Refresh", //TODO(PL)
		"pt-BR": "Refresh", //TODO(BR)
		"ru-RU": "Обновить",
		"tr-TR": "Refresh", //TODO(TR)
		"ua-UA": "Refresh", //TODO(UA)
		"uz-UZ": "Refresh", //TODO(UZ)
		"zh-CN": "Refresh", //TODO(CN)
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		"de-DE": "Etwas anderes",
		"en-UK": "Something else",
		"en-US": "Something else", //TODO(US)
		"es-ES": "Otra cosa",
		"fa-IR": "چیزی دیگر",
		"fr-FR": "Something else", //TODO(FR)
		"id-ID": "Something else", //TODO(ID)
		"it-IT": "Qualcos'altro",
		"ja-JP": "Something else", //TODO(JP)
		"ko-KO": "Something else", //TODO(KO)
		"pl-PL": "Something else", //TODO(PL)
		"pt-BR": "Something else", //TODO(BR)
		"ru-RU": "Что-то другое",
		"tr-TR": "Something else", //TODO(TR)
		"ua-UA": "Something else", //TODO(UA)
		"uz-UZ": "Something else", //TODO(UZ)
		"zh-CN": "Something else", //TODO(CN)
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		"de-DE": "Wurde diese Schuld beglichen?",
		"en-UK": "Have this debt been returned?",
		"en-US": "Have this debt been returned?", //TODO(US)
		"es-ES": "¿Se ha devuelto esta deuda?",
		"fa-IR": "آیا این بدهی بازپرداخت شده است؟",
		"fr-FR": "Have this debt been returned?", //TODO(FR)
		"id-ID": "Have this debt been returned?", //TODO(ID)
		"it-IT": "Questo debito e' stato saldato?",
		"ja-JP": "Have this debt been returned?", //TODO(JP)
		"ko-KO": "Have this debt been returned?", //TODO(KO)
		"pl-PL": "Have this debt been returned?", //TODO(PL)
		"pt-BR": "Have this debt been returned?", //TODO(BR)
		"ru-RU": "Был ли этот долг возвращён?",
		"tr-TR": "Have this debt been returned?", //TODO(TR)
		"ua-UA": "Have this debt been returned?", //TODO(UA)
		"uz-UZ": "Have this debt been returned?", //TODO(UZ)
		"zh-CN": "Have this debt been returned?", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		"de-DE": "Wann willst du wieder an diese Schuld erinnert werden?",
		"en-UK": "When should we remind you about this debt again?",
		"en-US": "When should we remind you about this debt again?", //TODO(US)
		"es-ES": "¿Cuándo recordarte de esta deuda otra vez?",
		"fa-IR": "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
		"fr-FR": "When should we remind you about this debt again?", //TODO(FR)
		"id-ID": "When should we remind you about this debt again?", //TODO(ID)
		"it-IT": "Quando devo ricordarti di questo debito?",
		"ja-JP": "When should we remind you about this debt again?", //TODO(JP)
		"ko-KO": "When should we remind you about this debt again?", //TODO(KO)
		"pl-PL": "When should we remind you about this debt again?", //TODO(PL)
		"pt-BR": "When should we remind you about this debt again?", //TODO(BR)
		"ru-RU": "Когда вам напомнить об этом долге ещё раз?",
		"tr-TR": "When should we remind you about this debt again?", //TODO(TR)
		"ua-UA": "When should we remind you about this debt again?", //TODO(UA)
		"uz-UZ": "When should we remind you about this debt again?", //TODO(UZ)
		"zh-CN": "When should we remind you about this debt again?", //TODO(CN)
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		"de-DE": "Du hast angegeben, dass diese Schuld vollständig beglichen ist.",
		"en-UK": "You've replied back that debt has been returned fully.",
		"en-US": "You've replied back that debt has been returned fully.", //TODO(US)
		"es-ES": "Has confirmado que la deuda se ha saldado totalmente",
		"fa-IR": "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
		"fr-FR": "You've replied back that debt has been returned fully.", //TODO(FR)
		"id-ID": "You've replied back that debt has been returned fully.", //TODO(ID)
		"it-IT": "Hai confermato che il debito e' stato saldato.",
		"ja-JP": "You've replied back that debt has been returned fully.", //TODO(JP)
		"ko-KO": "You've replied back that debt has been returned fully.", //TODO(KO)
		"pl-PL": "You've replied back that debt has been returned fully.", //TODO(PL)
		"pt-BR": "You've replied back that debt has been returned fully.", //TODO(BR)
		"ru-RU": "Вы ответили что долг возвращён полностью.",
		"tr-TR": "You've replied back that debt has been returned fully.", //TODO(TR)
		"ua-UA": "You've replied back that debt has been returned fully.", //TODO(UA)
		"uz-UZ": "You've replied back that debt has been returned fully.", //TODO(UZ)
		"zh-CN": "You've replied back that debt has been returned fully.", //TODO(CN)
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		"de-DE": "Die Schuld ist vollständig beglichen.",
		"en-UK": "The debt has been returned fully.",
		"en-US": "The debt has been returned fully.", //TODO(US)
		"es-ES": "La deuda se ha saldado totalmente",
		"fa-IR": "بدهی به صورت کامل بازپرداخت شده است",
		"fr-FR": "The debt has been returned fully.", //TODO(FR)
		"id-ID": "The debt has been returned fully.", //TODO(ID)
		"it-IT": "Il debito e' stato saldato.",
		"ja-JP": "The debt has been returned fully.", //TODO(JP)
		"ko-KO": "The debt has been returned fully.", //TODO(KO)
		"pl-PL": "The debt has been returned fully.", //TODO(PL)
		"pt-BR": "The debt has been returned fully.", //TODO(BR)
		"ru-RU": "Долг возвращён полностью.",
		"tr-TR": "The debt has been returned fully.", //TODO(TR)
		"ua-UA": "The debt has been returned fully.", //TODO(UA)
		"uz-UZ": "The debt has been returned fully.", //TODO(UZ)
		"zh-CN": "The debt has been returned fully.", //TODO(CN)
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		"de-DE": "Details hier: %v",
		"en-UK": "Details here: %v",
		"en-US": "Details here: %v", //TODO(US)
		"es-ES": "Detalles aquí: %v",
		"fa-IR": "جزئیات در اینجا: %v",
		"fr-FR": "Details here: %v", //TODO(FR)
		"id-ID": "Details here: %v", //TODO(ID)
		"it-IT": "Dettagli qui: %v",
		"ja-JP": "Details here: %v", //TODO(JP)
		"ko-KO": "Details here: %v", //TODO(KO)
		"pl-PL": "Details here: %v", //TODO(PL)
		"pt-BR": "Details here: %v", //TODO(BR)
		"ru-RU": "Подробности тут: %v",
		"tr-TR": "Details here: %v", //TODO(TR)
		"ua-UA": "Details here: %v", //TODO(UA)
		"uz-UZ": "Details here: %v", //TODO(UZ)
		"zh-CN": "Details here: %v", //TODO(CN)
	},
	MESSAGE_TEXT_REMINDER: {
		"de-DE": "Erinnerung",
		"en-UK": "Reminder",
		"en-US": "Reminder", //TODO(US)
		"es-ES": "Recordatorio",
		"fa-IR": "یادآور",
		"fr-FR": "Reminder", //TODO(FR)
		"id-ID": "Reminder", //TODO(ID)
		"it-IT": "Promemoria",
		"ja-JP": "Reminder", //TODO(JP)
		"ko-KO": "Reminder", //TODO(KO)
		"pl-PL": "Reminder", //TODO(PL)
		"pt-BR": "Reminder", //TODO(BR)
		"ru-RU": "Напоминание",
		"tr-TR": "Reminder", //TODO(TR)
		"ua-UA": "Reminder", //TODO(UA)
		"uz-UZ": "Reminder", //TODO(UZ)
		"zh-CN": "Reminder", //TODO(CN)
	},
	MESSAGE_TEXT_REMINDER_SET: {
		"de-DE": "Erinnerung am: %v",
		"en-UK": "Reminder set for: %v",
		"en-US": "Reminder set for: %v", //TODO(US)
		"es-ES": "Recordatorio establecito para: %v",
		"fa-IR": "یادآور تنظیم شده است برای: %v",
		"fr-FR": "Reminder set for: %v", //TODO(FR)
		"id-ID": "Reminder set for: %v", //TODO(ID)
		"it-IT": "Imposta promemoria per: %v",
		"ja-JP": "Reminder set for: %v", //TODO(JP)
		"ko-KO": "Reminder set for: %v", //TODO(KO)
		"pl-PL": "Reminder set for: %v", //TODO(PL)
		"pt-BR": "Reminder set for: %v", //TODO(BR)
		"ru-RU": "Напоминание установлено на: %v",
		"tr-TR": "Reminder set for: %v", //TODO(TR)
		"ua-UA": "Reminder set for: %v", //TODO(UA)
		"uz-UZ": "Reminder set for: %v", //TODO(UZ)
		"zh-CN": "Reminder set for: %v", //TODO(CN)
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		"de-DE": "Du hast die Erinnerung an diese Schuld deaktiviert.",
		"en-UK": "You've disabled reminders for this debt.",
		"en-US": "You've disabled reminders for this debt.", //TODO(US)
		"es-ES": "Recordatorio para esta deuda se ha deshabilitado.",
		"fa-IR": "شما یادآور را برای این بدهی غیرفعال نموده اید.",
		"fr-FR": "You've disabled reminders for this debt.", //TODO(FR)
		"id-ID": "You've disabled reminders for this debt.", //TODO(ID)
		"it-IT": "Hai disabilitato il promemoria per questo debito.",
		"ja-JP": "You've disabled reminders for this debt.", //TODO(JP)
		"ko-KO": "You've disabled reminders for this debt.", //TODO(KO)
		"pl-PL": "You've disabled reminders for this debt.", //TODO(PL)
		"pt-BR": "You've disabled reminders for this debt.", //TODO(BR)
		"ru-RU": "Напоминания об этом долге отключены.",
		"tr-TR": "You've disabled reminders for this debt.", //TODO(TR)
		"ua-UA": "You've disabled reminders for this debt.", //TODO(UA)
		"uz-UZ": "You've disabled reminders for this debt.", //TODO(UZ)
		"zh-CN": "You've disabled reminders for this debt.", //TODO(CN)
	},
	COMMAND_TEXT_REMINDER_ENABLE: {
		"de-DE": "Erinnerung aktivieren",
		"en-UK": "Turn-on reminder",
		"en-US": "Turn-on reminder",          //TODO(US)
		"es-ES": "Recordatorio de encendido", // TODO(es) verify
		"fa-IR": "یادآوری روشن",              // TODO(fa) verify
		"fr-FR": "Turn-on reminder",          //TODO(FR)
		"id-ID": "Turn-on reminder",          //TODO(ID)
		"it-IT": "Ricordo promozionale",      // TODO(it) verify
		"ja-JP": "Turn-on reminder",          //TODO(JP)
		"ko-KO": "Turn-on reminder",          //TODO(KO)
		"pl-PL": "Turn-on reminder",          //TODO(PL)
		"pt-BR": "Turn-on reminder",          //TODO(BR)
		"ru-RU": "Включить напоминание",
		"tr-TR": "Turn-on reminder", //TODO(TR)
		"ua-UA": "Turn-on reminder", //TODO(UA)
		"uz-UZ": "Turn-on reminder", //TODO(UZ)
		"zh-CN": "Turn-on reminder", //TODO(CN)
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		"de-DE": "Du wirst bereits erneut erinnert.",
		"en-UK": "You've already rescheduled this reminder.",
		"en-US": "You've already rescheduled this reminder.", //TODO(US)
		"es-ES": "Recordatorio para esta deuda se ha reprogramado ya.",
		"fa-IR": "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
		"fr-FR": "You've already rescheduled this reminder.", //TODO(FR)
		"id-ID": "You've already rescheduled this reminder.", //TODO(ID)
		"it-IT": "Hai gia' impostato questo promemoria",
		"ja-JP": "You've already rescheduled this reminder.", //TODO(JP)
		"ko-KO": "You've already rescheduled this reminder.", //TODO(KO)
		"pl-PL": "You've already rescheduled this reminder.", //TODO(PL)
		"pt-BR": "You've already rescheduled this reminder.", //TODO(BR)
		"ru-RU": "Напоминание об этом долге уже перенесено.",
		"tr-TR": "You've already rescheduled this reminder.", //TODO(TR)
		"ua-UA": "You've already rescheduled this reminder.", //TODO(UA)
		"uz-UZ": "You've already rescheduled this reminder.", //TODO(UZ)
		"zh-CN": "You've already rescheduled this reminder.", //TODO(CN)
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		"de-DE": "Ja, vollständig beglichen",
		"en-UK": "Yes, returne in full",
		"en-US": "Yes, returne in full", //TODO(US)
		"es-ES": "Sí, devuelto totalmente",
		"fa-IR": "بله، بازپرداخت به صورت کامل",
		"fr-FR": "Yes, returne in full", //TODO(FR)
		"id-ID": "Yes, returne in full", //TODO(ID)
		"it-IT": "Si, completamento saldato",
		"ja-JP": "Yes, returne in full", //TODO(JP)
		"ko-KO": "Yes, returne in full", //TODO(KO)
		"pl-PL": "Yes, returne in full", //TODO(PL)
		"pt-BR": "Yes, returne in full", //TODO(BR)
		"ru-RU": "Да, возвращено полностью",
		"tr-TR": "Yes, returne in full", //TODO(TR)
		"ua-UA": "Yes, returne in full", //TODO(UA)
		"uz-UZ": "Yes, returne in full", //TODO(UZ)
		"zh-CN": "Yes, returne in full", //TODO(CN)
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		"de-DE": "Teilweise beglichen",
		"en-UK": "Returned partially",
		"en-US": "Returned partially", //TODO(US)
		"es-ES": "Devuelto parcialmente",
		"fa-IR": "تا اندازه ای بازپرداخت شده است",
		"fr-FR": "Returned partially", //TODO(FR)
		"id-ID": "Returned partially", //TODO(ID)
		"it-IT": "Parzialmente saldato",
		"ja-JP": "Returned partially", //TODO(JP)
		"ko-KO": "Returned partially", //TODO(KO)
		"pl-PL": "Returned partially", //TODO(PL)
		"pt-BR": "Returned partially", //TODO(BR)
		"ru-RU": "Возврашено частично",
		"tr-TR": "Returned partially", //TODO(TR)
		"ua-UA": "Returned partially", //TODO(UA)
		"uz-UZ": "Returned partially", //TODO(UZ)
		"zh-CN": "Returned partially", //TODO(CN)
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		"de-DE": "Nicht beglichen",
		"en-UK": "Not returned",
		"en-US": "Not returned", //TODO(US)
		"es-ES": "No devuelto",
		"fa-IR": "بازپرداخت نشده است",
		"fr-FR": "Not returned", //TODO(FR)
		"id-ID": "Not returned", //TODO(ID)
		"it-IT": "Debito non saldato",
		"ja-JP": "Not returned", //TODO(JP)
		"ko-KO": "Not returned", //TODO(KO)
		"pl-PL": "Not returned", //TODO(PL)
		"pt-BR": "Not returned", //TODO(BR)
		"ru-RU": "Не возвращено",
		"tr-TR": "Not returned", //TODO(TR)
		"ua-UA": "Not returned", //TODO(UA)
		"uz-UZ": "Not returned", //TODO(UZ)
		"zh-CN": "Not returned", //TODO(CN)
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		"de-DE": "Beantwortet: %v",
		"en-UK": "You've replied: %v",
		"es-ES": "Has respondido: %v",
		"fa-IR": "شما پاسخ داده اید: %v",
		"it-IT": "Hai risposto: %v",
		"ru-RU": "Вы ответили: %v",
	},
	"book": {
		"de-DE": "buchen",
		"en-UK": "book",
		"es-ES": "libro",
		"fa-IR": "کتاب",
		"it-IT": "libro",
		"ru-RU": "книгу",
	},
	"MessageTextBotDidNotUnderstandTheCommand": {
		"de-DE": "\xF0\x9F\x98\xB3 Entschuldigung, aber ich habe deinen Befehl nicht verstanden. Vielleicht bin ich ein bisschen dumm...\n\nDu kannst zurück ins /menu",
		"en-UK": "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		"es-ES": "\xF0\x9F\x98\xB3 Disculpa, no he entendido tu orden. Tal vez soy un poco tonto...\n\nPuedes volver al Menu principal /menu",
		"fa-IR": "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		"it-IT": "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
		"ru-RU": "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
	},
	COMMAND_TEXT_LANGUAGE: {
		"de-DE": "Sprache / Language",
		"en-UK": "Bot language",
		"es-ES": "Idioma / Language",
		"fa-IR": "زبان",
		"it-IT": "Lingua / Language",
		"ru-RU": "Язык / Language",
	},
	"/start": {
		"de-DE": "/start",
		"en-UK": "/start",
		"es-ES": "/comienzo",
		"fa-IR": "/شروع",
		"it-IT": "/start",
		"ru-RU": "/старт",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		"de-DE": "Fällige Schulden",
		"en-UK": "Due returns",
		"es-ES": "Devoluciones",
		"fa-IR": "بازپرداخت بدهی",
		"it-IT": "Debiti",
		"ru-RU": "Предстоящие платежи",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		"de-DE": "<b>Überfällige Schulden:</b>",
		"en-UK": "<b>Overdue debts:</b>",
		"es-ES": "<b>Deudas atrasadas:</b>",
		"fa-IR": "<b>بدهی های معوق:</b>",
		"it-IT": "<b>Debiti in ritardo:</b>",
		"ru-RU": "<b>Просроченные долги:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		"de-DE": "<b>Bald fällige Schulden:</b>",
		"en-UK": "<b>Closest debts to return:</b>",
		"es-ES": "<b>Deudas más cercanos que pagar:</b>",
		"fa-IR": "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
		"it-IT": "<b>Debiti in scadenza:</b>",
		"ru-RU": "<b>Ближайшие долги к возрату:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		"de-DE": "%v bekommt %v von dir, spätestens in %v",
		"en-UK": "%v expects %v from you in %v",
		"es-ES": "%v espera %v que devuelvas en %v",
		"it-IT": "%v aspetta %v da te entro il %v",
		"fa-IR": "%v انتظار دارد %v از شما در %v",
		"ru-RU": "%v ожидает от вас возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		"de-DE": "%v gibt dir %v, spätestens in %v",
		"en-UK": "You expect %v to return %v to you in %v",
		"es-ES": "Estas esperando de %v que devuelva %v a ti en %v",
		"fa-IR": "شما انتظار دارید %v بازگرداند %v به شما در %v",
		"it-IT": "Stai aspettando %v che ti dia %v entro il %v",
		"ru-RU": "Вы ожидаете от %v возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		"de-DE": "Du hast keine Schulden mit Fälligkeitsdatum.",
		"en-UK": "You have no debts with set due date.",
		"es-ES": "No tienes deudas con la fecha señalada para devolver. ",
		"fa-IR": "شما بدهی ای با ثبت سررسید ندارید.",
		"it-IT": "Non hai debiti con una data di scadenza.",
		"ru-RU": "У вас нет долгов с назначеным сроком к возврату.",
	},
	COMMAND_TEXT_GAVE: {
		"de-DE": "Verleihen",
		"en-UK": "Gave",
		"es-ES": "Prestado por ti",
		"fa-IR": "قرض_دادن",
		"it-IT": "Credito",
		"ru-RU": "Дать",
	},
	COMMAND_TEXT_GOT: {
		"de-DE": "Anleihen",
		"en-UK": "Got",
		"es-ES": "Prestado a ti",
		"fa-IR": "قرض_گرفتن",
		"it-IT": "Debito",
		"ru-RU": "Взять",
	},
	COMMAND_TEXT_RETURN: {
		"de-DE": "Beglichen",
		"en-UK": "Return",
		"es-ES": "Devuelto",
		"fa-IR": "بازگشت",
		"it-IT": "Rientra",
		"ru-RU": "Вернуть",
	},
	COMMAND_TEXT_BALANCE: {
		"de-DE": "Ausstehend",
		"en-UK": "Balance",
		"es-ES": "Balance",
		"fa-IR": "تراز",
		"it-IT": "Bilancio",
		"ru-RU": "Баланс",
	},
	COMMAND_TEXT_SETTING: {
		"de-DE": "Einstellungen",
		"en-UK": "Settings",
		"es-ES": "Ajustes",
		"fa-IR": "تنظیمات",
		"it-IT": "Settaggi",
		"ru-RU": "Настройки",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		"de-DE": "Gib mir Fünf!",
		"en-UK": "High five!",
		"en-US": "High five!",
		"es-ES": "¡Choca esos cinco!",
		"fa-IR": "بزن قدش!",
		"fr-FR": "Tape m'en cinq !",
		"id-ID": "High five!",
		"it-IT": "Batti 5 bro!",
		"ja-JP": "ハイタッチ！",
		"ko-KO": "하이 파이브!",
		"pl-PL": "Przybij piątkę!",
		"pt-BR": "High five!",
		"ru-RU": "Дать пять!",
		"tr-TR": "Çak bir beşlik!",
		"ua-UA": "Дай п'ять!",
		"uz-UZ": "Besh qo'l!",
		"zh-CN": "击掌！",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		"de-DE": "Sprache",
		"en-UK": "Language",
		"en-US": "Language",
		"es-ES": "Idioma",
		"fa-IR": "زبان",
		"fr-FR": "Langue",
		"id-ID": "Bahasa",
		"it-IT": "Lingua",
		"ja-JP": "言語",
		"ko-KO": "언어",
		"pl-PL": "Język",
		"pt-BR": "Idioma",
		"ru-RU": "Язык",
		"tr-TR": "Dil",
		"ua-UA": "Мова",
		"uz-UZ": "Til",
		"zh-CN": "语言",
	},
	COMMAND_TEXT_HELP: {
		"de-DE": "Hilfe",
		"en-UK": "Help",
		"en-US": "Help",
		"es-ES": "Ayuda",
		"fa-IR": "کمک",
		"fr-FR": "Aide",
		"id-ID": "Bantuan",
		"it-IT": "Aiuto",
		"ja-JP": "ヘルプ",
		"ko-KO": "도움말",
		"pl-PL": "Pomoc",
		"pt-BR": "Ajuda",
		"ru-RU": "Помощь",
		"tr-TR": "Yardım",
		"ua-UA": "Допомога",
		"uz-UZ": "Yordam",
		"zh-CN": "帮助",
	},
	COMMAND_TEXT_HISTORY: {
		"de-DE": "Verlauf",
		"en-UK": "History",
		"en-US": "History",
		"es-ES": "Cronología",
		"fa-IR": "پیشینه",
		"fr-FR": "Historique",
		"id-ID": "Riwayat",
		"it-IT": "Cronologia",
		"ja-JP": "履歴",
		"ko-KO": "기록",
		"pl-PL": "Historia",
		"pt-BR": "Histórico",
		"ru-RU": "История",
		"tr-TR": "Geçmiş",
		"ua-UA": "Історія",
		"uz-UZ": "Tarix",
		"zh-CN": "历史",
	},
	COMMAND_TEXT_CANCEL: {
		"de-DE": "Abbrechen",
		"en-UK": "Cancel",
		"en-US": "Cancel",
		"es-ES": "Cancelar",
		"fa-IR": "کنسل",
		"fr-FR": "Annuler",
		"id-ID": "Batal",
		"it-IT": "Annulla",
		"ja-JP": "キャンセル",
		"ko-KO": "취소",
		"pl-PL": "Anuluj",
		"pt-BR": "Cancelar",
		"ru-RU": "Отменить",
		"tr-TR": "İptal",
		"ua-UA": "Скасувати",
		"uz-UZ": "Bekor qilish",
		"zh-CN": "取消",
	},
	COMMAND_TEXT_REFERRERS: {
		"de-DE": "Empfehlungen",
		"en-UK": "Referrers",
		"en-US": "Referrers",
		"es-ES": "Referentes",
		"fa-IR": "معرف‌ها",
		"fr-FR": "Référents",
		"id-ID": "Referensi",
		"it-IT": "Referenti",
		"ja-JP": "紹介者",
		"ko-KO": "추천인",
		"pl-PL": "Polecający",
		"pt-BR": "Referências",
		"ru-RU": "Нас рекомендуют",
		"tr-TR": "Referanslar",
		"ua-UA": "Нас рекомендують",
		"uz-UZ": "Tavsiya qiluvchilar",
		"zh-CN": "推荐人",
	},
	MESSAGE_TEXT_HOW_TO_ADD_TG_CHANNEL: {
		"de-DE": `Um deinen Kanal zur Liste hinzuzufügen, schreibe einfach über uns mit einem Link wie %v <code>&lt;-</code> ersetze <code>YOUR_CHANNEL</code> durch deinen eigenen Kanal.

Es ist besser, wenn du den Link in HTML versteckst als:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Dies sollte von Telegram-Clients so dargestellt werden: <a href="%v">@%v</a>

Die Top 5 Empfehlungen der letzten 100 neuen Benutzer werden hier angezeigt.`,

		"en-UK": `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,

		"en-US": `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,

		"es-ES": `Para añadir tu canal a la lista, simplemente escribe sobre nosotros con un enlace como %v <code>&lt;-</code> reemplaza <code>YOUR_CHANNEL</code> con tu propio canal.

Es mejor si ocultas el enlace en HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Esto debería ser renderizado por los clientes de Telegram como: <a href="%v">@%v</a>

Los 5 principales referentes de los últimos 100 nuevos usuarios se mostrarán aquí.`,

		"fa-IR": `برای اضافه کردن کانال خود به لیست، فقط درباره ما با لینکی مانند %v <code>&lt;-</code> بنویسید و <code>YOUR_CHANNEL</code> را با کانال خود جایگزین کنید.

بهتر است اگر لینک را در HTML به این صورت پنهان کنید:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

این باید توسط کلاینت‌های تلگرام به این صورت نمایش داده شود: <a href="%v">@%v</a>

5 معرف برتر برای 100 کاربر جدید آخر در اینجا نشان داده خواهد شد.`,

		"fr-FR": `Pour ajouter votre chaîne à la liste, écrivez simplement à propos de nous avec un lien comme %v <code>&lt;-</code> remplacez <code>YOUR_CHANNEL</code> par votre propre chaîne.

C'est mieux si vous cachez le lien en HTML comme:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Cela devrait être rendu par les clients Telegram comme: <a href="%v">@%v</a>

Les 5 principaux référents pour les 100 derniers nouveaux utilisateurs seront affichés ici.`,

		"id-ID": `Untuk menambahkan saluran Anda ke daftar, cukup tulis tentang kami dengan tautan sebagai %v <code>&lt;-</code> ganti <code>YOUR_CHANNEL</code> dengan saluran Anda sendiri.

Lebih baik jika Anda menyembunyikan tautan dalam HTML sebagai:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Ini harus dirender oleh klien Telegram sebagai: <a href="%v">@%v</a>

5 referensi teratas untuk 100 pengguna baru terakhir akan ditampilkan di sini.`,

		"it-IT": `Per aggiungere il tuo canale all'elenco, scrivi semplicemente di noi con un link come %v <code>&lt;-</code> sostituisci <code>YOUR_CHANNEL</code> con il tuo canale.

È meglio se nascondi il link in HTML come:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Questo dovrebbe essere visualizzato dai client Telegram come: <a href="%v">@%v</a>

I primi 5 referenti per gli ultimi 100 nuovi utenti saranno mostrati qui.`,

		"ja-JP": `あなたのチャンネルをリストに追加するには、%v <code>&lt;-</code> のようなリンクで私たちについて書くだけです。<code>YOUR_CHANNEL</code>をあなた自身のチャンネルに置き換えてください。

HTMLでリンクを隠すとより良いでしょう：

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

これはTelegramクライアントによって次のように表示されるはずです： <a href="%v">@%v</a>

最後の100人の新規ユーザーのトップ5紹介者がここに表示されます。`,

		"ko-KO": `채널을 목록에 추가하려면 %v <code>&lt;-</code>와 같은 링크로 우리에 대해 작성하세요. <code>YOUR_CHANNEL</code>을 자신의 채널로 바꾸세요.

HTML에서 링크를 숨기는 것이 좋습니다:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

이것은 Telegram 클라이언트에서 다음과 같이 렌더링되어야 합니다: <a href="%v">@%v</a>

최근 100명의 새 사용자에 대한 상위 5개 추천인이 여기에 표시됩니다.`,

		"pl-PL": `Aby dodać swój kanał do listy, po prostu napisz o nas z linkiem jako %v <code>&lt;-</code> zastąp <code>YOUR_CHANNEL</code> swoim własnym kanałem.

Lepiej, jeśli ukryjesz link w HTML jako:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Powinno to być renderowane przez klientów Telegram jako: <a href="%v">@%v</a>

Tutaj zostanie wyświetlonych 5 najlepszych polecających dla ostatnich 100 nowych użytkowników.`,

		"pt-BR": `Para adicionar seu canal à lista, basta escrever sobre nós com um link como %v <code>&lt;-</code> substitua <code>YOUR_CHANNEL</code> pelo seu próprio canal.

É melhor se você ocultar o link em HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Isso deve ser renderizado pelos clientes do Telegram como: <a href="%v">@%v</a>

Os 5 principais referenciadores para os últimos 100 novos usuários serão mostrados aqui.`,

		"ru-RU": `Чтобы добавить ваш канал в этот список просто напишите об этом боте использую ссылку вида %v <code>&lt;-</code> замените <code>YOUR_CHANNEL</code> на ваш канал.

Будет лучше  если вы спрячете её в HTML как:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Такой код должен отобразиться в Телеграмме как: <a href="%v">@%v</a>

Топ-5 источников последних 100 пользователей будут показаны здесь.`,

		"tr-TR": `Kanalınızı listeye eklemek için sadece %v <code>&lt;-</code> gibi bir bağlantı ile hakkımızda yazın, <code>YOUR_CHANNEL</code> yerine kendi kanalınızı yazın.

Bağlantıyı HTML'de şu şekilde gizlerseniz daha iyi olur:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Bu, Telegram istemcileri tarafından şöyle görüntülenmelidir: <a href="%v">@%v</a>

Son 100 yeni kullanıcı için en iyi 5 referans burada gösterilecektir.`,

		"ua-UA": `Щоб додати свій канал до списку, просто напишіть про нас із посиланням як %v <code>&lt;-</code> замініть <code>YOUR_CHANNEL</code> на свій власний канал.

Краще, якщо ви приховаєте посилання в HTML як:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Це має відображатися клієнтами Telegram як: <a href="%v">@%v</a>

Тут буде показано 5 найкращих рекомендацій для останніх 100 нових користувачів.`,

		"uz-UZ": `Kanalingizni ro'yxatga qo'shish uchun shunchaki %v <code>&lt;-</code> kabi havola bilan biz haqimizda yozing, <code>YOUR_CHANNEL</code> o'rniga o'z kanalingizni yozing.

Havolani HTML-da quyidagicha yashirsangiz yaxshiroq bo'ladi:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Bu Telegram mijozlari tomonidan quyidagicha ko'rsatilishi kerak: <a href="%v">@%v</a>

Oxirgi 100 ta yangi foydalanuvchi uchun eng yaxshi 5 ta tavsiya qiluvchi bu yerda ko'rsatiladi.`,

		"zh-CN": `要将您的频道添加到列表中，只需使用链接 %v <code>&lt;-</code> 写下关于我们的信息，将 <code>YOUR_CHANNEL</code> 替换为您自己的频道。

如果您在HTML中隐藏链接会更好：

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

这应该由Telegram客户端呈现为： <a href="%v">@%v</a>

这里将显示最近100个新用户的前5名推荐人。`,
	},
	ButtonTextCancel: {
		"de-DE": "↩ Zurück",
		"en-UK": "↩ Cancel",
		"en-US": "↩ Cancel",
		"es-ES": "↩ Cancelar",
		"fa-IR": "↪ کنسل",
		"fr-FR": "↩ Annuler",
		"id-ID": "↩ Batal",
		"it-IT": "↩ Annulla",
		"ja-JP": "↩ キャンセル",
		"ko-KO": "↩ 취소",
		"pl-PL": "↩ Anuluj",
		"pt-BR": "↩ Cancelar",
		"ru-RU": "↩ Отменить",
		"tr-TR": "↩ İptal",
		"ua-UA": "↩ Скасувати",
		"uz-UZ": "↩ Bekor qilish",
		"zh-CN": "↩ 取消",
	},
	BUTTON_TEXT_MAIN_MENU: {
		"de-DE": "↩ Hauptmenü",
		"en-UK": "↩ Main menu",
		"en-US": "↩ Main menu",
		"es-ES": "↩ Menú principal",
		"fa-IR": "↪ منوی اصلی",
		"fr-FR": "↩ Menu principal",
		"id-ID": "↩ Menu utama",
		"it-IT": "↩ Menu principale",
		"ja-JP": "↩ メインメニュー",
		"ko-KO": "↩ 메인 메뉴",
		"pl-PL": "↩ Menu główne",
		"pt-BR": "↩ Menu principal",
		"ru-RU": "↩ Главное меню",
		"tr-TR": "↩ Ana menü",
		"ua-UA": "↩ Головне меню",
		"uz-UZ": "↩ Asosiy menyu",
		"zh-CN": "↩ 主菜单",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		"de-DE": "Hauptwährung",
		"en-UK": "Primary currency",
		"en-US": "Primary currency",
		"es-ES": "Moneda principal",
		"fa-IR": "واحد پول اولیه",
		"fr-FR": "Devise principale",
		"id-ID": "Mata uang utama",
		"it-IT": "Valuta principale",
		"ja-JP": "主要通貨",
		"ko-KO": "기본 통화",
		"pl-PL": "Główna waluta",
		"pt-BR": "Moeda principal",
		"ru-RU": "Основная валюта",
		"tr-TR": "Ana para birimi",
		"ua-UA": "Основна валюта",
		"uz-UZ": "Asosiy valyuta",
		"zh-CN": "主要货币",
	},
	COMMAND_TEXT_ADD_GROUP: {
		"de-DE": "Gruppe hinzufügen",
		"en-UK": "Add group",
		"en-US": "Add group",
		"es-ES": "Añadir grupo",
		"fa-IR": "اضافه کردن گروه",
		"fr-FR": "Ajouter un groupe",
		"id-ID": "Tambahkan grup",
		"it-IT": "Aggiungi gruppo",
		"ja-JP": "グループを追加",
		"ko-KO": "그룹 추가",
		"pl-PL": "Dodaj grupę",
		"pt-BR": "Adicionar grupo",
		"ru-RU": "Добавить группу",
		"tr-TR": "Grup ekle",
		"ua-UA": "Додати групу",
		"uz-UZ": "Guruh qo'shish",
		"zh-CN": "添加群组",
	},
	COMMAND_TEXT_GROUPS: {
		"de-DE": "Gruppen",
		"en-UK": "Groups",
		"en-US": "Groups",
		"es-ES": "Grupos",
		"fa-IR": "گروه‌ها",
		"fr-FR": "Groupes",
		"id-ID": "Grup",
		"it-IT": "Gruppi",
		"ja-JP": "グループ",
		"ko-KO": "그룹",
		"pl-PL": "Grupy",
		"pt-BR": "Grupos",
		"ru-RU": "Группы",
		"tr-TR": "Gruplar",
		"ua-UA": "Групи",
		"uz-UZ": "Guruhlar",
		"zh-CN": "群组",
	},
	COMMAND_TEXT_BILLS: {
		"de-DE": "Rechnungen",
		"en-UK": "Bills",
		"en-US": "Bills",
		"es-ES": "Facturas",
		"fa-IR": "صورتحساب‌ها",
		"fr-FR": "Factures",
		"id-ID": "Tagihan",
		"it-IT": "Fatture",
		"ja-JP": "請求書",
		"ko-KO": "청구서",
		"pl-PL": "Rachunki",
		"pt-BR": "Contas",
		"ru-RU": "Счета",
		"tr-TR": "Faturalar",
		"ua-UA": "Рахунки",
		"uz-UZ": "Hisob-fakturalar",
		"zh-CN": "账单",
	},
	COMMAND_TEXT_SETTLE_BILL: {
		"de-DE": "Rechnung begleichen",
		"en-UK": "Settle bill",
		"en-US": "Settle bill",
		"es-ES": "Liquidar factura",
		"fa-IR": "تسویه صورتحساب",
		"fr-FR": "Régler la facture",
		"id-ID": "Selesaikan tagihan",
		"it-IT": "Saldare il conto",
		"ja-JP": "請求書を決済する",
		"ko-KO": "청구서 정산",
		"pl-PL": "Rozlicz rachunek",
		"pt-BR": "Quitar conta",
		"ru-RU": "Оплатить счёт",
		"tr-TR": "Fatura öde",
		"ua-UA": "Оплатити рахунок",
		"uz-UZ": "Hisob-fakturani to'lash",
		"zh-CN": "结算账单",
	},
	COMMAND_TEXT_SETTLE_BILLS: {
		"de-DE": "Rechnungen begleichen",
		"en-UK": "Settle bills",
		"en-US": "Settle bills",
		"es-ES": "Liquidar facturas",
		"fa-IR": "تسویه صورتحساب‌ها",
		"fr-FR": "Régler les factures",
		"id-ID": "Selesaikan tagihan",
		"it-IT": "Saldare i conti",
		"ja-JP": "請求書を決済する",
		"ko-KO": "청구서 정산",
		"pl-PL": "Rozlicz rachunki",
		"pt-BR": "Quitar contas",
		"ru-RU": "Закрыть счета",
		"tr-TR": "Faturaları öde",
		"ua-UA": "Оплатити рахунки",
		"uz-UZ": "Hisob-fakturalarni to'lash",
		"zh-CN": "结算账单",
	},
	COMMAND_TEXT_INVITE_FIREND: {
		"de-DE": "Freund einladen",
		"en-UK": "Invite friend",
		"en-US": "Invite friend",
		"es-ES": "Invitar a un amigo",
		"fa-IR": "دوستی دعوت کن",
		"fr-FR": "Inviter un ami",
		"id-ID": "Undang teman",
		"it-IT": "Invita un amico",
		"ja-JP": "友達を招待",
		"ko-KO": "친구 초대",
		"pl-PL": "Zaproś przyjaciela",
		"pt-BR": "Convidar amigo",
		"ru-RU": "Пригласить друга",
		"tr-TR": "Arkadaş davet et",
		"ua-UA": "Запросити друга",
		"uz-UZ": "Do'stni taklif qilish",
		"zh-CN": "邀请朋友",
	},
	COMMAND_TEXT_INVITE_MEMBER: {
		"de-DE": "Mitglied einladen",
		"en-UK": "Invite member",
		"en-US": "Invite member",
		"es-ES": "Invitar miembro",
		"fa-IR": "دعوت از اعضا",
		"fr-FR": "Inviter un membre",
		"id-ID": "Undang anggota",
		"it-IT": "Invita membro",
		"ja-JP": "メンバーを招待",
		"ko-KO": "멤버 초대",
		"pl-PL": "Zaproś członka",
		"pt-BR": "Convidar membro",
		"ru-RU": "Пригласить участника",
		"tr-TR": "Üye davet et",
		"ua-UA": "Запросити учасника",
		"uz-UZ": "A'zoni taklif qilish",
		"zh-CN": "邀请成员",
	},
	COMMAND_TEXT_NEW_BILL: {
		"de-DE": "Neue Rechnung",
		"en-UK": "New bill",
		"en-US": "New bill",
		"es-ES": "Nueva factura",
		"fa-IR": "صورتحساب جدید",
		"fr-FR": "Nouvelle facture",
		"id-ID": "Tagihan baru",
		"it-IT": "Nuova fattura",
		"ja-JP": "新しい請求書",
		"ko-KO": "새 청구서",
		"pl-PL": "Nowy rachunek",
		"pt-BR": "Nova conta",
		"ru-RU": "Новый счёт",
		"tr-TR": "Yeni fatura",
		"ua-UA": "Новий рахунок",
		"uz-UZ": "Yangi hisob-faktura",
		"zh-CN": "新账单",
	},
	COMMAND_TEXT_NEW_FUNDRAISING: {
		"de-DE": "Neue Spendensammlung",
		"en-UK": "New fundraising",
		"en-US": "New fundraising",
		"es-ES": "Nueva recaudación de fondos",
		"fa-IR": "جمع آوری پول جدید",
		"fr-FR": "Nouvelle collecte de fonds",
		"id-ID": "Penggalangan dana baru",
		"it-IT": "Nuova raccolta fondi",
		"ja-JP": "新しい資金調達",
		"ko-KO": "새 모금",
		"pl-PL": "Nowa zbiórka pieniędzy",
		"pt-BR": "Nova arrecadação de fundos",
		"ru-RU": "Новый сбор средств",
		"tr-TR": "Yeni bağış toplama",
		"ua-UA": "Новий збір коштів",
		"uz-UZ": "Yangi mablag' yig'ish",
		"zh-CN": "新筹款",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		"de-DE": "neuer Kontakt",
		"en-UK": "Add new",
		"en-US": "Add new",
		"es-ES": "Añadir",
		"fa-IR": "اضافه کردن مورد جدید",
		"fr-FR": "Ajouter nouveau",
		"id-ID": "Tambah baru",
		"it-IT": "Aggiungi nuovo",
		"ja-JP": "新規追加",
		"ko-KO": "새로 추가",
		"pl-PL": "Dodaj nowy",
		"pt-BR": "Adicionar novo",
		"ru-RU": "Добавить",
		"tr-TR": "Yeni ekle",
		"ua-UA": "Додати новий",
		"uz-UZ": "Yangi qo'shish",
		"zh-CN": "添加新的",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		"de-DE": "Dein Code um dich an der App anzumelden: <b>%v</b>",
		"en-UK": "Your code for signing in to app: <b>%v</b>",
		"en-US": "Your code for signing in to app: <b>%v</b>",
		"es-ES": "Tu código para entrar en la app: <b>%v</b>",
		"fa-IR": "کد شما برای ورود به برنامه: <b>%v</b>",
		"fr-FR": "Votre code pour vous connecter à l'application: <b>%v</b>",
		"id-ID": "Kode Anda untuk masuk ke aplikasi: <b>%v</b>",
		"it-IT": "Il tuo codice per accedere all'app e': <b>%v</b>",
		"ja-JP": "アプリにサインインするためのコード: <b>%v</b>",
		"ko-KO": "앱에 로그인하기 위한 코드: <b>%v</b>",
		"pl-PL": "Twój kod do logowania do aplikacji: <b>%v</b>",
		"pt-BR": "Seu código para entrar no aplicativo: <b>%v</b>",
		"ru-RU": "Ваш код для входа в приложение: <b>%v</b>",
		"tr-TR": "Uygulamaya giriş yapma kodunuz: <b>%v</b>",
		"ua-UA": "Ваш код для входу в додаток: <b>%v</b>",
		"uz-UZ": "Ilovaga kirish uchun kodingiz: <b>%v</b>",
		"zh-CN": "您登录应用的验证码: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		"de-DE": `Bitte gib einen Namen für den neuen Kontakt ein:
		Du kannst in eintippen oder aus deinem Adressbuch wählen (<i>mit dem "Büroklammer"-Symbol und dann Kontakt</i>).

		<i>Send '.' to cancel</i>`,

		"en-UK": `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		"en-US": `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		"es-ES": `Escribe un nombre para el nuevo contacto:
		Puedes escribirlo o elegirlo de tus contactos (<i>a traves del icono "clip"</i>).

		<i>Enviar '.' para anular</i>`,

		"fa-IR": `لطفا برای مخاطب جدید یک نام وارد کنید:
		میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

		<i>Send '.' برای کنسل کردن</i>`,

		"fr-FR": `Veuillez entrer un nom pour le nouveau contact:
		Vous pouvez taper manuellement ou choisir dans votre carnet d'adresses (<i>via l'icône "trombone"</i>).

		<i>Envoyez '.' pour annuler</i>`,

		"id-ID": `Silakan masukkan nama untuk kontak baru:
		Anda dapat mengetik secara manual atau memilih dari buku alamat Anda (<i>melalui ikon "klip"</i>).

		<i>Kirim '.' untuk membatalkan</i>`,

		"it-IT": `Inserisci un nome per il nuovo contatto:
		Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

		<i>Digita '.' ed invia per annullare</i>`,

		"ja-JP": `新しい連絡先の名前を入力してください:
		手動で入力するか、アドレス帳から選択できます（<i>「クリップ」アイコンを通じて</i>）。

		<i>キャンセルするには '.' を送信してください</i>`,

		"ko-KO": `새 연락처의 이름을 입력하세요:
		수동으로 입력하거나 주소록에서 선택할 수 있습니다 (<i>"클립" 아이콘을 통해</i>).

		<i>취소하려면 '.'를 보내세요</i>`,

		"pl-PL": `Wprowadź nazwę dla nowego kontaktu:
		Możesz wpisać ręcznie lub wybrać z książki adresowej (<i>przez ikonę "spinacza"</i>).

		<i>Wyślij '.' aby anulować</i>`,

		"pt-BR": `Por favor, digite um nome para o novo contato:
		Você pode digitar manualmente ou escolher do seu livro de endereços (<i>através do ícone "clipe"</i>).

		<i>Envie '.' para cancelar</i>`,

		"ru-RU": `<b>Имя для нового контакта</b>
		Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).
		<i>Отправьте '.' для отмены</i>`,

		"tr-TR": `Yeni kişi için bir isim girin:
		Manuel olarak yazabilir veya adres defterinizden seçebilirsiniz (<i>"ataç" simgesi aracılığıyla</i>).

		<i>İptal etmek için '.' gönderin</i>`,

		"ua-UA": `Будь ласка, введіть ім'я для нового контакту:
		Ви можете ввести вручну або вибрати з адресної книги (<i>через іконку "скріпка"</i>).

		<i>Надішліть '.' для скасування</i>`,

		"uz-UZ": `Yangi kontakt uchun ism kiriting:
		Siz qo'lda yozishingiz yoki manzillar kitobingizdan tanlashingiz mumkin (<i>"qisqich" belgisi orqali</i>).

		<i>Bekor qilish uchun '.' yuboring</i>`,

		"zh-CN": `请输入新联系人的名称:
		您可以手动输入或从通讯录中选择（<i>通过"回形针"图标</i>）。

		<i>发送 '.' 取消</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		"de-DE": "Transferiere...",
		"en-UK": "Creating transfer...",
		"en-US": "Creating transfer...",
		"es-ES": "Estoy creando la nueva nota...",
		"fa-IR": "ایجاد انتقال ...",
		"fr-FR": "Création du transfert...",
		"id-ID": "Membuat transfer...",
		"it-IT": "Sto creando la nuova voce...",
		"ja-JP": "転送を作成中...",
		"ko-KO": "전송 생성 중...",
		"pl-PL": "Tworzenie transferu...",
		"pt-BR": "Criando transferência...",
		"ru-RU": "Создаю запись...",
		"tr-TR": "Transfer oluşturuluyor...",
		"ua-UA": "Створення переказу...",
		"uz-UZ": "O'tkazma yaratilmoqda...",
		"zh-CN": "创建转账中...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		"de-DE": "Bitte warten",
		"en-UK": "Please wait",
		"en-US": "Please wait",
		"es-ES": "Espera, por favor",
		"fa-IR": "لطفا صبر کنید",
		"fr-FR": "Veuillez patienter",
		"id-ID": "Mohon tunggu",
		"it-IT": "Aspetta per favore",
		"ja-JP": "お待ちください",
		"ko-KO": "잠시만 기다려주세요",
		"pl-PL": "Proszę czekać",
		"pt-BR": "Por favor, aguarde",
		"ru-RU": "Пожалуйста подождите",
		"tr-TR": "Lütfen bekleyin",
		"ua-UA": "Будь ласка, зачекайте",
		"uz-UZ": "Iltimos kuting",
		"zh-CN": "请稍等",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		"de-DE": "Bitte warten...",
		"en-UK": "Please wait...",
		"en-US": "Please wait...",
		"es-ES": "Espera, por favor...",
		"fa-IR": "لطفا صبر کنید ...",
		"fr-FR": "Veuillez patienter...",
		"id-ID": "Mohon tunggu...",
		"it-IT": "Aspetta per favore...",
		"ja-JP": "お待ちください...",
		"ko-KO": "잠시만 기다려주세요...",
		"pl-PL": "Proszę czekać...",
		"pt-BR": "Por favor, aguarde...",
		"ru-RU": "Пожалуйста подождите...",
		"tr-TR": "Lütfen bekleyin...",
		"ua-UA": "Будь ласка, зачекайте...",
		"uz-UZ": "Iltimos kuting...",
		"zh-CN": "请稍等...",
	},
	MESAGE_TEXT_CREATING_BILL: {
		"de-DE": "Rechnung erstellen",
		"en-UK": "Creating bill",
		"en-US": "Creating bill",
		"es-ES": "Creando factura",
		"fa-IR": "ایجاد صورتحساب",
		"fr-FR": "Création de facture",
		"id-ID": "Membuat tagihan",
		"it-IT": "Creazione di fattura",
		"ja-JP": "請求書を作成中",
		"ko-KO": "청구서 생성 중",
		"pl-PL": "Tworzenie rachunku",
		"pt-BR": "Criando conta",
		"ru-RU": "Создаётся счёт",
		"tr-TR": "Fatura oluşturuluyor",
		"ua-UA": "Створення рахунку",
		"uz-UZ": "Hisob-faktura yaratilmoqda",
		"zh-CN": "创建账单中",
	},
	MESSAGE_TEXT_ASK_BILL_CURRENCY: {
		"de-DE": "In welcher Währung ist die Rechnung?",
		"en-UK": "What currency this bill in?",
		"en-US": "What currency is this bill in?",
		"es-ES": "¿En qué moneda está esta factura?",
		"fa-IR": "این صورتحساب به چه ارزی است؟",
		"fr-FR": "Dans quelle devise est cette facture?",
		"id-ID": "Dalam mata uang apa tagihan ini?",
		"it-IT": "In quale valuta è questa fattura?",
		"ja-JP": "この請求書の通貨は何ですか？",
		"ko-KO": "이 청구서의 통화는 무엇입니까?",
		"pl-PL": "W jakiej walucie jest ten rachunek?",
		"pt-BR": "Em qual moeda está esta conta?",
		"ru-RU": "В какой валюте этот счёт?",
		"tr-TR": "Bu fatura hangi para biriminde?",
		"ua-UA": "У якій валюті цей рахунок?",
		"uz-UZ": "Bu hisob-faktura qaysi valyutada?",
		"zh-CN": "这个账单使用什么货币？",
	},
	MESSAGE_TEXT_ASK_BILL_PAYER: {
		"de-DE": "Wer hat die Rechnung bezahlt?",
		"en-UK": "Who paid for the bill?",
		"en-US": "Who paid for the bill?",
		"es-ES": "¿Quién pagó la factura?",
		"fa-IR": "چه کسی صورتحساب را پرداخت کرد؟",
		"fr-FR": "Qui a payé la facture?",
		"id-ID": "Siapa yang membayar tagihan?",
		"it-IT": "Chi ha pagato il conto?",
		"ja-JP": "誰が請求書を支払いましたか？",
		"ko-KO": "누가 청구서를 지불했습니까?",
		"pl-PL": "Kto zapłacił rachunek?",
		"pt-BR": "Quem pagou a conta?",
		"ru-RU": "Кто оплатил счёт?",
		"tr-TR": "Faturayı kim ödedi?",
		"ua-UA": "Хто оплатив рахунок?",
		"uz-UZ": "Hisob-fakturani kim to'ladi?",
		"zh-CN": "谁支付了账单？",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		"de-DE": "%v muss dem zustimmen",
		"en-UK": "Acknowledgement is expected from %v",
		"en-US": "Acknowledgement is expected from %v",
		"es-ES": "Se espera la confirmación de %v",
		"fa-IR": "انتظار تصدیق می رود از %v",
		"fr-FR": "Confirmation attendue de %v",
		"id-ID": "Pengakuan diharapkan dari %v",
		"it-IT": "Conferma in attesa da %v",
		"ja-JP": "%vからの確認が必要です",
		"ko-KO": "%v의 확인이 필요합니다",
		"pl-PL": "Oczekiwane potwierdzenie od %v",
		"pt-BR": "Confirmação esperada de %v",
		"ru-RU": "Подтверждение ожидается от %v",
		"tr-TR": "%v'den onay bekleniyor",
		"ua-UA": "Очікується підтвердження від %v",
		"uz-UZ": "%v dan tasdiqlash kutilmoqda",
		"zh-CN": "等待%v的确认",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		"de-DE": "Du hast dem zugestimmt.",
		"en-UK": "You've accepted this transaction.",
		"en-US": "You've accepted this transaction.",
		"es-ES": "Has confirmado esta transacción",
		"fa-IR": ".شما این تراکنش را قبول کردید ",
		"fr-FR": "Vous avez accepté cette transaction.",
		"id-ID": "Anda telah menerima transaksi ini.",
		"it-IT": "Hai accettato il debito/credito.",
		"ja-JP": "このトランザクションを承認しました。",
		"ko-KO": "이 거래를 수락했습니다.",
		"pl-PL": "Zaakceptowałeś tę transakcję.",
		"pt-BR": "Você aceitou esta transação.",
		"ru-RU": "Вы подтвердили эту транзакцию.",
		"tr-TR": "Bu işlemi kabul ettiniz.",
		"ua-UA": "Ви підтвердили цю транзакцію.",
		"uz-UZ": "Siz ushbu tranzaksiyani qabul qildingiz.",
		"zh-CN": "您已接受此交易。",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		"de-DE": `Du hast dieser Anfrage nicht zugestimmt.
		Der Vorgang wird zurückgestellt und die Gegenpartei benachrichtigt.`,

		"en-UK": `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,

		"en-US": `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,

		"es-ES": `No estas de acuerdo con la transacción.
		La transacción NO será cancelada, pero el creador será notificado.`,

		"fa-IR": `شما با این تراکنش موافق نیستید.
		تراکنش حذف نخواهد شد اما به ایجاد کننده اطلاع داده خواهد شد.`,

		"fr-FR": `Vous n'êtes pas d'accord avec cette transaction.
		La transaction ne sera pas supprimée mais le créateur sera notifié.`,

		"id-ID": `Anda tidak setuju dengan transaksi ini.
		Transaksi tidak akan dihapus tetapi pembuat akan diberi tahu.`,

		"it-IT": `Hai rifiutato il debito/credito.
		La transazione non sarà eliminata ma il creatore sarà avvisato.`,

		"ja-JP": `このトランザクションに同意しません。
		トランザクションは削除されませんが、作成者に通知されます。`,

		"ko-KO": `이 거래에 동의하지 않습니다.
		거래는 삭제되지 않지만 생성자에게 알림이 갑니다.`,

		"pl-PL": `Nie zgadzasz się z tą transakcją.
		Transakcja nie zostanie usunięta, ale twórca zostanie powiadomiony.`,

		"pt-BR": `Você não concorda com esta transação.
		A transação não será excluída, mas o criador será notificado.`,

		"ru-RU": `Вы НЕ согласны с этой транзакцией.

Сама транзакция НЕ будет отменена, но создатель будет оповещён.`,

		"tr-TR": `Bu işleme katılmıyorsunuz.
		İşlem silinmeyecek ancak oluşturucu bilgilendirilecek.`,

		"ua-UA": `Ви не згодні з цією транзакцією.
		Транзакція не буде видалена, але творець буде повідомлений.`,

		"uz-UZ": `Siz ushbu tranzaksiyaga rozi emassiz.
		Tranzaksiya o'chirilmaydi, lekin yaratuvchi xabardor qilinadi.`,

		"zh-CN": `您不同意此交易。
		交易不会被删除，但创建者会收到通知。`,
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		"de-DE": "%v hat deiner Anfrage <b>zugestimmt</b>:",
		"en-UK": "%v accepted your transaction:",
		"en-US": "%v accepted your transaction:",
		"es-ES": "%v ha aceptado tu transacción",
		"fa-IR": ": تراکنش شمارا تایید کرد %v ",
		"fr-FR": "%v a accepté votre transaction:",
		"id-ID": "%v menerima transaksi Anda:",
		"it-IT": "%v ha accettato il tuo credito/debito:",
		"ja-JP": "%vがあなたの取引を承認しました:",
		"ko-KO": "%v님이 귀하의 거래를 수락했습니다:",
		"pl-PL": "%v zaakceptował(a) twoją transakcję:",
		"pt-BR": "%v aceitou sua transação:",
		"ru-RU": "%v подтвердил(a) вашу транзакцию:",
		"tr-TR": "%v işleminizi kabul etti:",
		"ua-UA": "%v підтвердив(ла) вашу транзакцію:",
		"uz-UZ": "%v sizning tranzaksiyangizni qabul qildi:",
		"zh-CN": "%v接受了您的交易：",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		"de-DE": "%v hat deine Anfrage <b>abgelehnt</b>. Wenn die Sache besprochen ist, kann die Anfrage erneut gesendet werden.",
		"en-UK": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(FA)
		"en-US": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(US)
		"es-ES": "%v ha confirmado tu transacción.  La transacción no ha sido cancelada, igual mejor comentarlo.",          //TODO(ES)
		"fa-IR": "تراکنش شما را رد کرد  %v declined your transaction.",
		"fr-FR": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(FR)
		"id-ID": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(ID)
		"it-IT": "%v ha rifiutato il tuo credito/debito.  The transaction is not canceled but you may want to discuss it.", //TODO(IT)
		"ja-JP": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(JP)
		"ko-KO": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(KO)
		"pl-PL": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(PL)
		"pt-BR": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(BR)
		"ru-RU": "%v <b>НЕ</b> подтвердил(a) вашу транзакцию. Транзакция не отменена, но возможно вам стоит это обсудить.",
		"tr-TR": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(TR)
		"ua-UA": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(UA)
		"uz-UZ": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(UZ)
		"zh-CN": "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.", //TODO(CN)
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		"de-DE": "Ich will die App!",
		"en-UK": "I want the app!",
		"en-US": "I want the app!", //TODO(US)
		"es-ES": "¡Quiero la aplicación!",
		"fa-IR": "!من برنامه را می خواهم",
		"fr-FR": "I want the app!", //TODO(FR)
		"id-ID": "I want the app!", //TODO(ID)
		"it-IT": "Voglio l'aplicazione!",
		"ja-JP": "I want the app!", //TODO(JP)
		"ko-KO": "I want the app!", //TODO(KO)
		"pl-PL": "I want the app!", //TODO(PL)
		"pt-BR": "I want the app!", //TODO(BR)
		"ru-RU": "Хочу приложение!",
		"tr-TR": "I want the app!", //TODO(TR)
		"ua-UA": "I want the app!", //TODO(UA)
		"uz-UZ": "I want the app!", //TODO(UZ)
		"zh-CN": "I want the app!", //TODO(CN)
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		"de-DE": "Mir reicht der Bot!",
		"en-UK": "I'm fine with just the bot!",
		"en-US": "I'm fine with just the bot!", //TODO(US)
		"es-ES": "¡Estoy satisfecho con este bot!",
		"fa-IR": "! ربات به تنهایی برای من کافی است",
		"fr-FR": "I'm fine with just the bot!", //TODO(FR)
		"id-ID": "I'm fine with just the bot!", //TODO(ID)
		"it-IT": "Mi accontento del bot per ora!",
		"ja-JP": "I'm fine with just the bot!", //TODO(JP)
		"ko-KO": "I'm fine with just the bot!", //TODO(KO)
		"pl-PL": "I'm fine with just the bot!", //TODO(PL)
		"pt-BR": "I'm fine with just the bot!", //TODO(BR)
		"ru-RU": "Меня вполне устраивает бот!",
		"tr-TR": "I'm fine with just the bot!", //TODO(TR)
		"ua-UA": "I'm fine with just the bot!", //TODO(UA)
		"uz-UZ": "I'm fine with just the bot!", //TODO(UZ)
		"zh-CN": "I'm fine with just the bot!", //TODO(CN)
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		"de-DE": "Du wirst darüber informiert, wenn die App zum Download zur Verfügung steht.",
		"en-UK": "We'll let you know once the app is available for download.",
		"en-US": "We'll let you know once the app is available for download.", //TODO(US)
		"es-ES": "Te avisamos cuando la aplicación esté disponible para descargarla",
		"fa-IR": ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
		"fr-FR": "We'll let you know once the app is available for download.", //TODO(FR)
		"id-ID": "We'll let you know once the app is available for download.", //TODO(ID)
		"it-IT": "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
		"ja-JP": "We'll let you know once the app is available for download.", //TODO(JP)
		"ko-KO": "We'll let you know once the app is available for download.", //TODO(KO)
		"pl-PL": "We'll let you know once the app is available for download.", //TODO(PL)
		"pt-BR": "We'll let you know once the app is available for download.", //TODO(BR)
		"ru-RU": "Мы сообщим вам когда приложение будет доступно для загруки.",
		"tr-TR": "We'll let you know once the app is available for download.", //TODO(TR)
		"ua-UA": "We'll let you know once the app is available for download.", //TODO(UA)
		"uz-UZ": "We'll let you know once the app is available for download.", //TODO(UZ)
		"zh-CN": "We'll let you know once the app is available for download.", //TODO(CN)
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		"de-DE": "Gut, wir sind froh, dass dir der Bot reicht und wir uns mit der App nicht beeilen müssen.",
		"en-UK": "Well, we are happy our bot is good enough and there is no need to download an app.",
		"en-US": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(US)
		"es-ES": "Bueno, estamos contentos de que te haya gustado nuestro bot y no hace falta descargar ninguna otra aplicación",
		"fa-IR": ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
		"fr-FR": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(FR)
		"id-ID": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(ID)
		"it-IT": "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
		"ja-JP": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(JP)
		"ko-KO": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(KO)
		"pl-PL": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(PL)
		"pt-BR": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(BR)
		"ru-RU": "Что ж, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		"tr-TR": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(TR)
		"ua-UA": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(UA)
		"uz-UZ": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(UZ)
		"zh-CN": "Well, we are happy our bot is good enough and there is no need to download an app.", //TODO(CN)
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		"de-DE": "Hier könnte <a href>ihre Werbung</a> stehen",
		"en-UK": "You can <a href>advertise here</a>",
		"en-US": "You can <a href>advertise here</a>", //TODO(US)
		"es-ES": "Aquí se puede <a href>publicar un anuncio</a>",
		"fa-IR": "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
		"fr-FR": "You can <a href>advertise here</a>", //TODO(FR)
		"id-ID": "You can <a href>advertise here</a>", //TODO(ID)
		"it-IT": "Puoi <a href>pubblicizzare qui</a>",
		"ja-JP": "You can <a href>advertise here</a>", //TODO(JP)
		"ko-KO": "You can <a href>advertise here</a>", //TODO(KO)
		"pl-PL": "You can <a href>advertise here</a>", //TODO(PL)
		"pt-BR": "You can <a href>advertise here</a>", //TODO(BR)
		"ru-RU": "Здесь можно <a href>разместить рекламу</a>",
		"tr-TR": "You can <a href>advertise here</a>", //TODO(TR)
		"ua-UA": "You can <a href>advertise here</a>", //TODO(UA)
		"uz-UZ": "You can <a href>advertise here</a>", //TODO(UZ)
		"zh-CN": "You can <a href>advertise here</a>", //TODO(CN)
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		"de-DE": `🤖: Ich hin ein guter Roboter - klar. Aber manchmal kommt es besser eine eigene App für etwas zu haben. Es ist noch nicht ganz fertig, aber falls du schonmal reinschauen willst: <a href="https://debtstracker.io/de/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Möchtest du daran erinnert werden, wenn die App rauskommt?`,
		"en-UK": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,
		"en-US": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`, //TODO(US)

		"es-ES": `🤖: Claro que soy un robot encantador, pero más comodo usar una aplicación especial.No esta disponible ya pero se puede ver como será: <a href = "https://debtstracker.io/es/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	¿Quieres que te avisemos cuando esté lista?`,
		"fa-IR": `🤖: مطمئناً من روبات خوبی هستم. اما بعضی وقت هاساده تر و مناسب تر است که از یک برنامه به خوبی تخصصی شده استفاده شود، این برنامه هنوز برای استفاده عموم آماده نیست ولی می توانید چک کنید که چگونه به نظر خواهد رسید: <a href="https://debtstracker.io/fa/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	آیا می خواهید وقتی منتشر شد دعوتنامه ای دریافت کنید؟`,
		"fr-FR": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/fr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/fr/</a>

	Do you want to get an invite when it gets released?`, //TODO(FR)
		"id-ID": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/id/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/id/</a>

	Do you want to get an invite when it gets released?`, //TODO(ID)

		"it-IT": `🤖: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href="https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

	Vuoi essere invitato non appena viene rilasciata?`,
		"ja-JP": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ja/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ja/</a>

	Do you want to get an invite when it gets released?`, //TODO(JP)
		"ko-KO": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ko/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ko/</a>

	Do you want to get an invite when it gets released?`, //TODO(KO)
		"pl-PL": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pl/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pl/</a>

	Do you want to get an invite when it gets released?`, //TODO(PL)
		"pt-BR": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pt/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pt/</a>

	Do you want to get an invite when it gets released?`, //TODO(BR)

		"ru-RU": `🤖: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

		Хотите получить оповещение когда оно выйдет?`,
		"tr-TR": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/tr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/tr/</a>

	Do you want to get an invite when it gets released?`, //TODO(TR)
		"ua-UA": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ua/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ua/</a>

	Do you want to get an invite when it gets released?`, //TODO(UA)
		"uz-UZ": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/uz/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/uz/</a>

	Do you want to get an invite when it gets released?`, //TODO(UZ)
		"zh-CN": `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/zh/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/zh/</a>

	Do you want to get an invite when it gets released?`, //TODO(CN)
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		"de-DE": "Entschuldigung, aber du kannst nur Zahlen für Menge oder Wert wählen (<i>mit zwei Nachkommastellen</i>).",
		"en-UK": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		"en-US": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(US)
		"es-ES": "Lo siento, solo puedes utilizar numeros como importe/cantidad (<i>con un maximo de 2 dígitos despues de la coma</i>).",
		"fa-IR": "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
		"fr-FR": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(FR)
		"id-ID": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(ID)
		"it-IT": "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
		"ja-JP": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(JP)
		"ko-KO": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(KO)
		"pl-PL": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(PL)
		"pt-BR": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(BR)
		"ru-RU": "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаков после запятой</i>).",
		"tr-TR": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(TR)
		"ua-UA": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(UA)
		"uz-UZ": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(UZ)
		"zh-CN": "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		"de-DE": "<b>Was hast du jemanden geliehen?</b>",
		"en-UK": "<b>What did you lend to someone?</b>",
		"en-US": "<b>What did you lend to someone?</b>", //TODO(US)
		"es-ES": "<b>¿Qué has prestado?</b>",
		"fa-IR": "<b> چه چیزی به کسی قرض داده اید؟</b>",
		"fr-FR": "<b>What did you lend to someone?</b>", //TODO(FR)
		"id-ID": "<b>What did you lend to someone?</b>", //TODO(ID)
		"it-IT": "<b>Cos'hai prestato?</b>",
		"ja-JP": "<b>What did you lend to someone?</b>", //TODO(JP)
		"ko-KO": "<b>What did you lend to someone?</b>", //TODO(KO)
		"pl-PL": "<b>What did you lend to someone?</b>", //TODO(PL)
		"pt-BR": "<b>What did you lend to someone?</b>", //TODO(BR)
		"ru-RU": "<b>Что вы дали в долг?</b>",
		"tr-TR": "<b>What did you lend to someone?</b>", //TODO(TR)
		"ua-UA": "<b>What did you lend to someone?</b>", //TODO(UA)
		"uz-UZ": "<b>What did you lend to someone?</b>", //TODO(UZ)
		"zh-CN": "<b>What did you lend to someone?</b>", //TODO(CN)
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {

		"de-DE": `Bitte wähle <a>eine Währung aus der Liste</a>.

	Falls die Standardoptionen nicht reichen, sende mir einen Text. Zum Beispiel: <i>Äpfel</i>".`,

		"en-UK": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		"en-US": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(US)

		"es-ES": `Elige del menú abajo de la pantalla o <a>selecciona la moneda de la lista</a>.

	Si no encuentras la opción correcta simplemente envía un texto. Por ejemplo: "<i>manzana</i>".`,

		"fa-IR": `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

	اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال:. "<i>سیب</i>".`,

		"fr-FR": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(FR)

		"id-ID": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(ID)

		"it-IT": `Scegli dalle opzioni qui sotto o <a>seleziona una valuta dalla lista</a>.

	Se le opzioni standard non bastano semplicemente invia un testo.Per esempio: "<i>mele</i>".`,

		"ja-JP": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(JP)

		"ko-KO": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(KO)

		"pl-PL": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(PL)

		"pt-BR": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(BR)

		"ru-RU": `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

	Если ни один из стандартных вариантов не подходит просто напишите текстом.Например: "<i>яблоко</i>".`,

		"tr-TR": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(TR)

		"ua-UA": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(UA)

		"uz-UZ": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(UZ)

		"zh-CN": `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`, //TODO(CN)
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		"de-DE": "Wie viel <b>%v</b> hast du verliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		"en-US": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(US)
		"es-ES": "Cuanto(s) <b>%v</b> has prestado\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"fr-FR": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(FR)
		"id-ID": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(ID)
		"it-IT": "Quanti <b>%v</b> hai prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ja-JP": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(JP)
		"ko-KO": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(KO)
		"pl-PL": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(PL)
		"pt-BR": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(BR)
		"ru-RU": "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"tr-TR": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(TR)
		"ua-UA": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(UA)
		"uz-UZ": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(UZ)
		"zh-CN": "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		"de-DE": "Wer hat sich <b>%v</b> von dir geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"en-US": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(US)
		"es-ES": "A quién has prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه کسی از شما <b>%v</b> قرض گرفته است؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"fr-FR": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(FR)
		"id-ID": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(ID)
		"it-IT": "Chi e' in debito di <b>%v</b> con te?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ja-JP": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(JP)
		"ko-KO": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(KO)
		"pl-PL": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(PL)
		"pt-BR": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(BR)
		"ru-RU": "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"tr-TR": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(TR)
		"ua-UA": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(UA)
		"uz-UZ": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(UZ)
		"zh-CN": "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		"de-DE": "Was hast du dir geliehen?",
		"en-UK": "What did you lend?",
		"en-US": "What did you lend?", //TODO(US)
		"es-ES": "¿Qué te han prestado?",
		"fa-IR": "چه چیزی قرض گرفته اید؟",
		"fr-FR": "What did you lend?", //TODO(FR)
		"id-ID": "What did you lend?", //TODO(ID)
		"it-IT": "Cosa ti hanno prestato?",
		"ja-JP": "What did you lend?", //TODO(JP)
		"ko-KO": "What did you lend?", //TODO(KO)
		"pl-PL": "What did you lend?", //TODO(PL)
		"pt-BR": "What did you lend?", //TODO(BR)
		"ru-RU": "Что вы взяли в долг?",
		"tr-TR": "What did you lend?", //TODO(TR)
		"ua-UA": "What did you lend?", //TODO(UA)
		"uz-UZ": "What did you lend?", //TODO(UZ)
		"zh-CN": "What did you lend?", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		"de-DE": "Wie viel <b>%v</b> hast du geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		"en-US": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(US)
		"es-ES": "¿Cuánto <b>%v</b> has prestado?\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		"fr-FR": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(FR)
		"id-ID": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(ID)
		"it-IT": "Quanti <b>%v</b> ti hanno prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ja-JP": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(JP)
		"ko-KO": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(KO)
		"pl-PL": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(PL)
		"pt-BR": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(BR)
		"ru-RU": "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		"tr-TR": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(TR)
		"ua-UA": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(UA)
		"uz-UZ": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(UZ)
		"zh-CN": "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)", //TODO(CN)
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		"de-DE": "Wer hat dir <b>%v</b> geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		"en-UK": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		"en-US": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(US)
		"es-ES": "¿Quién te ha prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		"fa-IR": "چه کسی به شما <b>%v</b> قرض داده است؟ \n(<i>ارسال '.' برای کنسل کردن</i>)",
		"fr-FR": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(FR)
		"id-ID": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(ID)
		"it-IT": "Chi ti ha prestato <b>%v</b>?\n(<i>Digita '.' ed invia per annullare</i>)",
		"ja-JP": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(JP)
		"ko-KO": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(KO)
		"pl-PL": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(PL)
		"pt-BR": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(BR)
		"ru-RU": "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		"tr-TR": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(TR)
		"ua-UA": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(UA)
		"uz-UZ": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(UZ)
		"zh-CN": "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)", //TODO(CN)
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		"de-DE": "Soll eine <a receipt>Quittung</a> an <a counterparty>%v</a> gesendet werden?",
		"en-UK": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		"en-US": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(US)
		"es-ES": "¿Debo enviar <a receipt> el recibo</a> a <a counterparty>%v</a>?",
		"fa-IR": "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
		"fr-FR": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(FR)
		"id-ID": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(ID)
		"it-IT": "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
		"ja-JP": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(JP)
		"ko-KO": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(KO)
		"pl-PL": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(PL)
		"pt-BR": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(BR)
		"ru-RU": "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		"tr-TR": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(TR)
		"ua-UA": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(UA)
		"uz-UZ": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(UZ)
		"zh-CN": "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?", //TODO(CN)
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		"de-DE": "Entschuldigung, aber eine Quittung selber per SMS zu schicken ist im Moment noch nicht möglich. Aber dafür geht es mit %v.",
		"en-UK": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		"en-US": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(US)
		"es-ES": "Lo siento, el envio del recibo a ti mismo a través de SMS en este momento está desactivado. Pero lo puedes enviar a %v.",
		"fa-IR": "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
		"fr-FR": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(FR)
		"id-ID": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(ID)
		"it-IT": "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
		"ja-JP": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(JP)
		"ko-KO": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(KO)
		"pl-PL": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(PL)
		"pt-BR": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(BR)
		"ru-RU": "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		"tr-TR": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(TR)
		"ua-UA": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(UA)
		"uz-UZ": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(UZ)
		"zh-CN": "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		"de-DE": "Die Quittung wird %v per Telegram geschickt...",
		"en-UK": "We are sending receipt to %v by Telegram...",
		"en-US": "We are sending receipt to %v by Telegram...", //TODO(US)
		"es-ES": "El recibo está enviando a%v a través de Telegram…",
		"fa-IR": "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
		"fr-FR": "We are sending receipt to %v by Telegram...", //TODO(FR)
		"id-ID": "We are sending receipt to %v by Telegram...", //TODO(ID)
		"it-IT": "Sto inviando la notifica a %v tramite Telegram...",
		"ja-JP": "We are sending receipt to %v by Telegram...", //TODO(JP)
		"ko-KO": "We are sending receipt to %v by Telegram...", //TODO(KO)
		"pl-PL": "We are sending receipt to %v by Telegram...", //TODO(PL)
		"pt-BR": "We are sending receipt to %v by Telegram...", //TODO(BR)
		"ru-RU": "Отправляем для %v извещение через Telegram...",
		"tr-TR": "We are sending receipt to %v by Telegram...", //TODO(TR)
		"ua-UA": "We are sending receipt to %v by Telegram...", //TODO(UA)
		"uz-UZ": "We are sending receipt to %v by Telegram...", //TODO(UZ)
		"zh-CN": "We are sending receipt to %v by Telegram...", //TODO(CN)
	},
	DAY: {
		"de-DE": "%v day", //TODO(DE)
		"en-UK": "%v day",
		"en-US": "%v day", //TODO(US)
		"es-ES": "%v day", //TODO(ES)
		"fa-IR": "%v day", //TODO(FA)
		"fr-FR": "%v day", //TODO(FR)
		"id-ID": "%v day", //TODO(ID)
		"it-IT": "%v day", //TODO(IT)
		"ja-JP": "%v day", //TODO(JP)
		"ko-KO": "%v day", //TODO(KO)
		"pl-PL": "%v day", //TODO(PL)
		"pt-BR": "%v day", //TODO(BR)
		"ru-RU": "%v день",
		"tr-TR": "%v day", //TODO(TR)
		"ua-UA": "%v day", //TODO(UA)
		"uz-UZ": "%v day", //TODO(UZ)
		"zh-CN": "%v day", //TODO(CN)
	},
	DAYS_234: {
		"de-DE": "%v days", //TODO(DE)
		"en-UK": "%v days",
		"en-US": "%v days", //TODO(US)
		"es-ES": "%v days", //TODO(ES)
		"fa-IR": "%v days", //TODO(FA)
		"fr-FR": "%v days", //TODO(FR)
		"id-ID": "%v days", //TODO(ID)
		"it-IT": "%v days", //TODO(IT)
		"ja-JP": "%v days", //TODO(JP)
		"ko-KO": "%v days", //TODO(KO)
		"pl-PL": "%v days", //TODO(PL)
		"pt-BR": "%v days", //TODO(BR)
		"ru-RU": "%v дня",
		"tr-TR": "%v days", //TODO(TR)
		"ua-UA": "%v days", //TODO(UA)
		"uz-UZ": "%v days", //TODO(UZ)
		"zh-CN": "%v days", //TODO(CN)
	},
	DAYS: {
		"de-DE": "%v days", //TODO(DE)
		"en-UK": "%v days",
		"en-US": "%v days", //TODO(US)
		"es-ES": "%v days", //TODO(ES)
		"fa-IR": "%v days", //TODO(FA)
		"fr-FR": "%v days", //TODO(FR)
		"id-ID": "%v days", //TODO(ID)
		"it-IT": "%v days", //TODO(IT)
		"ja-JP": "%v days", //TODO(JP)
		"ko-KO": "%v days", //TODO(KO)
		"pl-PL": "%v days", //TODO(PL)
		"pt-BR": "%v days", //TODO(BR)
		"ru-RU": "%v дней",
		"tr-TR": "%v days", //TODO(TR)
		"ua-UA": "%v days", //TODO(UA)
		"uz-UZ": "%v days", //TODO(UZ)
		"zh-CN": "%v days", //TODO(CN)
	},
	MESSAGE_TEXT_INTEREST_PLEASE_SPECIFY_PERIOD: {
		"de-DE": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(DE)
		"en-UK": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		"en-US": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(US)
		"es-ES": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(ES)
		"fa-IR": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(FA)
		"fr-FR": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(FR)
		"id-ID": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(ID)
		"it-IT": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(IT)
		"ja-JP": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(JP)
		"ko-KO": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(KO)
		"pl-PL": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(PL)
		"pt-BR": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(BR)
		"ru-RU": "Пожалуйста укажите также процентный период, т.е. уточните %%v%% это процент за какое количество дней?",
		"tr-TR": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(TR)
		"ua-UA": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(UA)
		"uz-UZ": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(UZ)
		"zh-CN": "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?", //TODO(CN)
	},
	MESSAGE_TEXT_INTEREST: {
		"de-DE": "<b>Interest</b>: %v%% per %v", //TODO(DE)
		"en-UK": "<b>Interest</b>: %v%% per %v",
		"en-US": "<b>Interest</b>: %v%% per %v", //TODO(US)
		"es-ES": "<b>Interest</b>: %v%% per %v", //TODO(ES)
		"fa-IR": "<b>Interest</b>: %v%% per %v", //TODO(FA)
		"fr-FR": "<b>Interest</b>: %v%% per %v", //TODO(FR)
		"id-ID": "<b>Interest</b>: %v%% per %v", //TODO(ID)
		"it-IT": "<b>Interest</b>: %v%% per %v", //TODO(IT)
		"ja-JP": "<b>Interest</b>: %v%% per %v", //TODO(JP)
		"ko-KO": "<b>Interest</b>: %v%% per %v", //TODO(KO)
		"pl-PL": "<b>Interest</b>: %v%% per %v", //TODO(PL)
		"pt-BR": "<b>Interest</b>: %v%% per %v", //TODO(BR)
		"ru-RU": "<b>Ставка</b>: %v%% за %v",
		"tr-TR": "<b>Interest</b>: %v%% per %v", //TODO(TR)
		"ua-UA": "<b>Interest</b>: %v%% per %v", //TODO(UA)
		"uz-UZ": "<b>Interest</b>: %v%% per %v", //TODO(UZ)
		"zh-CN": "<b>Interest</b>: %v%% per %v", //TODO(CN)
	},
	MESSAGE_TEXT_INTEREST_MIN_PERIOD: {
		"de-DE": "minimum period %v", //TODO(DE)
		"en-UK": "minimum period %v",
		"en-US": "minimum period %v", //TODO(US)
		"es-ES": "minimum period %v", //TODO(ES)
		"fa-IR": "minimum period %v", //TODO(FA)
		"fr-FR": "minimum period %v", //TODO(FR)
		"id-ID": "minimum period %v", //TODO(ID)
		"it-IT": "minimum period %v", //TODO(IT)
		"ja-JP": "minimum period %v", //TODO(JP)
		"ko-KO": "minimum period %v", //TODO(KO)
		"pl-PL": "minimum period %v", //TODO(PL)
		"pt-BR": "minimum period %v", //TODO(BR)
		"ru-RU": "минимальный период %v",
		"tr-TR": "minimum period %v", //TODO(TR)
		"ua-UA": "minimum period %v", //TODO(UA)
		"uz-UZ": "minimum period %v", //TODO(UZ)
		"zh-CN": "minimum period %v", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		"de-DE": "{{.Counterparty}} schuldet dir {{.Amount}} .",
		"en-UK": "{{.Counterparty}} borrowed from you {{.Amount}}.",
		"en-US": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(US)
		"es-ES": "{{.Counterparty}} prestado por tí {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} از شما {{.Amount}} قرض گرفته است .",
		"fr-FR": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(FR)
		"id-ID": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(ID)
		//"it-IT": "{{.Counterparty}} ha preso in prestito da te {{.Amount}}.",
		"it-IT": "{{.Counterparty}} e' in debito di {{.Amount}} con te.",
		"ja-JP": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(JP)
		"ko-KO": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(KO)
		"pl-PL": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(PL)
		"pt-BR": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(BR)
		"ru-RU": "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		"tr-TR": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(TR)
		"ua-UA": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(UA)
		"uz-UZ": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(UZ)
		"zh-CN": "{{.Counterparty}} borrowed from you {{.Amount}}.", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		"de-DE": "{{.Counterparty}} hat dir {{.Amount}} geliehen.",
		"en-UK": "{{.Counterparty}} lended to you {{.Amount}}.",
		"en-US": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(US)
		"es-ES": "{{.Counterparty}} prestado a mí {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما {{.Amount}} قرض داده است .",
		"fr-FR": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(FR)
		"id-ID": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(ID)
		"it-IT": "{{.Counterparty}} ti ha prestato {{.Amount}}.",
		"ja-JP": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(JP)
		"ko-KO": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(KO)
		"pl-PL": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(PL)
		"pt-BR": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(BR)
		"ru-RU": "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		"tr-TR": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(TR)
		"ua-UA": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(UA)
		"uz-UZ": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(UZ)
		"zh-CN": "{{.Counterparty}} lended to you {{.Amount}}.", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		"de-DE": "Du hast {{.Amount}} an {{.Counterparty}} beglichen.",
		"en-UK": "You returned {{.Amount}} to {{.Counterparty}}.",
		"en-US": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(US)
		"es-ES": "Has devuelto {{.Amount}} a {{.Counterparty}}.",
		"fa-IR": "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
		"fr-FR": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(FR)
		"id-ID": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(ID)
		"it-IT": "Hai ridato {{.Amount}} a {{.Counterparty}}.",
		"ja-JP": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(JP)
		"ko-KO": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(KO)
		"pl-PL": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(PL)
		"pt-BR": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(BR)
		"ru-RU": "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		"tr-TR": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(TR)
		"ua-UA": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(UA)
		"uz-UZ": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(UZ)
		"zh-CN": "You returned {{.Amount}} to {{.Counterparty}}.", //TODO(CN)
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		"de-DE": "{{.Counterparty}} hat dir {{.Amount}} beglichen.",
		"en-UK": "{{.Counterparty}} returned to you {{.Amount}}.",
		"en-US": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(US)
		"es-ES": "{{.Counterparty}} te ha devuelto {{.Amount}}.",
		"fa-IR": "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
		"fr-FR": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(FR)
		"id-ID": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(ID)
		"it-IT": "{{.Counterparty}} ti ha ridato {{.Amount}}.",
		"ja-JP": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(JP)
		"ko-KO": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(KO)
		"pl-PL": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(PL)
		"pt-BR": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(BR)
		"ru-RU": "{{.Counterparty}} вернул вам {{.Amount}}.",
		"tr-TR": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(TR)
		"ua-UA": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(UA)
		"uz-UZ": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(UZ)
		"zh-CN": "{{.Counterparty}} returned to you {{.Amount}}.", //TODO(CN)
	},
	MESSAGE_TEXT_TRANSFER_ALREADY_FULLY_RETURNED: {
		"de-DE": "Diese Schuld ist bereits vollständig beglichen.",
		"en-UK": "This debts is already fully returned.",
		"es-ES": "Esta deuda se ha devuelta totalmente.",
		"it-IT": "Questi debiti sono già completamente restituiti.", // TODO(IT) verify
		"fa-IR": "این بدهی ها در حال حاضر به طور کامل بازگشته است.", // TODO(FA) verify
		"ru-RU": "Этот долг уже полностью возвращён.",
	},
	MESSAGE_TEXT_RECEIPT_ALREADY_RETURNED_AMOUNT: {
		"de-DE": "Bereits beglichen: {{.Amount}}.",
		"en-UK": "Already returned: {{.Amount}}.",
		"es-ES": "Se ha devuelto ya: {{.Amount}}.",
		"fa-IR": "قبلا برگشت: {{.Amount}}.",     // TODO(FA) - verify
		"it-IT": "Già restituito: {{.Amount}}.", // TODO(IT) - verify
		"ru-RU": "Уже возвращено: {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_OUTSTANDING_AMOUNT: {
		"de-DE": "Ausstehend: {{.Amount}}.",
		"en-UK": "Outstanding: {{.Amount}}.",
		"es-ES": "Falta devolver: {{.Amount}}.",
		"fa-IR": "برجسته: {{.Amount}}",  // TODO(FA) verify
		"it-IT": "Inevaso: {{.Amount}}", // TODO(IT) verify
		"ru-RU": "Осталось вернуть: {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		"de-DE": "<b>Fällig am</b>: %v",
		"en-UK": "<b>Return till</b>: %v",
		"es-ES": "<b>Devolver hasta</b>: %v",
		"fa-IR": "<b>بازگردانده شود تا</b>: %v",
		"it-IT": "<b>Scadenza</b>: %v",
		"ru-RU": "<b>Вернуть до</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		"de-DE": "Notiz",
		"en-UK": "Note",
		"es-ES": "Nota",
		"fa-IR": "نکته",
		"it-IT": "Nota",
		"ru-RU": "Заметка",
	},
	MESSAGE_TEXT_COMMENT: {
		"de-DE": "Bemerkung",
		"en-UK": "Comment",
		"es-ES": "Comentario",
		"fa-IR": "شرح",
		"it-IT": "Commenti",
		"ru-RU": "Комментарий",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		"de-DE": `<a>Hier klicken</a>, um sich an der Wep-App anzumelden.`,
		"en-UK": `Click to <a>sign in</a> to web-app.`,
		"es-ES": `Haz click para <a>acceder</a>la web-app.`,
		"it-IT": "Fai clic per <a>accedi</a> per app web.",
		"fa-IR": `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
		"ru-RU": `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		"de-DE": "Magst du @{{bot}}?",
		"en-UK": "Do you like @{{bot}}?",
		"es-ES": "¿Te gusta @{{bot}}?",
		"fa-IR": "آیا می پسندید @{{bot}}?",
		"it-IT": "Divertito con @{{bot}}?",
		"ru-RU": "Вам нравится @{{bot}}?",
	},
	COMMAND_TEXT_YES_EXCLAMATION: {
		"de-DE": "%v Ja!",
		"en-UK": "%v Yes!",
		"es-ES": "%v ¡Sí!",
		"fa-IR": "بله! %v",
		"it-IT": "%v Si!",
		"ru-RU": "%v Да!",
	},
	COMMAND_TEXT_YES: {
		"de-DE": "%v Ja",
		"en-UK": "%v Yes",
		"es-ES": "%v Sí",
		"it-IT": "%v Si",
		"fa-IR": "بله %v",
		"ru-RU": "%v Да",
	},
	COMMAND_TEXT_NO: {
		"de-DE": "%v Nein",
		"en-UK": "%v No",
		"es-ES": "%v No",
		"it-IT": "%v No",
		"fa-IR": "خیر %v",
		"ru-RU": "%v Нет",
	},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		"de-DE": "%v Nicht so sehr",
		"en-UK": "%v Not too much",
		"es-ES": "%vNo mucho",
		"it-IT": "%v Non troppo",
		"fa-IR": "نه خیلی زیاد %v",
		"ru-RU": "%v Не очень",
	},
	COMMAND_TEXT_FEEDBACK: {
		"de-DE": "/Bewertung",
		"en-UK": "/Feedback",
		"es-ES": "/Respuesta",
		"it-IT": "/Risposta",
		"fa-IR": "/بازخورد",
		"ru-RU": "/Отзыв",
	},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		"de-DE": "%v Bewertung schreiben",
		"en-UK": "%v Write feedback",
		"es-ES": "%v Escribir un comentario",
		"it-IT": "%v Scrivi commenti",
		"fa-IR": "ارسال بازخورد %v",
		"ru-RU": "%v Написать отзыв",
	},
	MESSAGE_TEXT_THANKS: {
		"de-DE": "🙏 Danke!",
		"en-UK": "🙏 Thanks!",
		"es-ES": "🙏 ¡Gracias!",
		"fa-IR": "🙏 تشکر!",
		"it-IT": "🙏 Grazie!",
		"ru-RU": "🙏 Спасибо!",
	},
	MESSGE_TEXT_DEBT_ERROR_FIXED_START_OVER: {
		"de-DE": "🙏 Entschuldigung, da gab es einen Fehler. Er wird bald behoben, aber du musst nochmal neu anfangen.",
		"en-UK": "🙏 Sorry, there was an error. It has been fixed but please re-enter your data for this debt.",
		"es-ES": "🙏 Lo siento, ha salido un error. Lo ha arreglado, pero para esta deuda hay que introducir los datos de nuevo. ",
		"ru-RU": "🙏 Извините, у нас была ошибка. Она была исправлено, но потребуется внести данные для этого долга заново.",
	},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		"de-DE": "Bitte senden sie einen Text.",
		"en-UK": "Please send text.",
		"es-ES": "Por favor, envia el texto.",
		"fa-IR": "لطفاً متن ارسال کنید.",
		"it-IT": "Si prega di inviare il testo.",
		"ru-RU": "Пожалуйста отправьте текст.",
	},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {

		"de-DE": `🤖 Kannst du mich im Store Bot hoch bewerten und eine gute Bewertung schreiben?
		Es wird dich weniger als eine Minute kosten! 😇`,
		"en-UK": `🤖 Can you rate it high and write a good review in bots catalog Store Bot?
		It will take less than a minute of your time! 😇`,

		"es-ES": `🤖 Puedes valolarlo con una buena nota y una buena opinión en el catálogo Store Bot?
		Te costará menos de un minuto! 😇`,

		"fa-IR": `🤖  آیا می توانید در کاتالوگ روباتها در استور بوت امتیاز بالایی داده و اظهار نظر خوبی در مورد این روبات ثبت کنید؟
		این کار کمتر از یک دقیقه از وقت شما را می گیرد! 😇`,

		"it-IT": `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo Bot Store?
		Ci vorrà meno di un minuto del tuo tempo! 😇`,

		"ru-RU": `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?
		Это займет меньше минуты вашего времени! 😇`,
	},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		"de-DE": "Schreibe (auf Englisch oder Russisch) uns, was man am Bot besser machen kann:",
		"en-UK": "Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		"es-ES": "Comparte tus pensamientos (en Inglés o Ruso) sobre qué podemos hacer para que el bot sea mejor:",
		"fa-IR": "نظرات خود را (به انگلیسی و روسی ) در مورد اینکه چه کاری می توان انجام داد تا این ربات بهتر شود، با ما به اشتراک بگذارید:",
		"it-IT": "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
		"ru-RU": "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
	},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {

		"de-DE": `<b>Wie man bewertet - in 3 einfachen Schritten:</b>

	1. Klick auf diesen Link, um eine Bewertung abzugeben:
	https://t.me/storebot?start={{bot}}

	2. Wähle "⭐️⭐️⭐️⭐️⭐️" 

	3. Schreib etwas Nettes auf Englisch oder wähle "Skip this step"

	Wirklich vielen Dank! Dank deiner Bewertung werden vielleicht mehr Leute auf diesen Bot aufmerksam. Das ist gut für die Motivation der Entwickler dieses Bots! 😎`,

		"en-UK": `<b>How to rate in 3 simple steps:</b>

	1. Click on this link to rate and review:
	https://t.me/storebot?start={{bot}}

	2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

	3. Write your message or press "Skip this step" button

	Thank you very much! As a result of your actions, even more people will learn about the bot.All this will serve as the additional motivation for the developers! 😎`,

		"es-ES": `<b>Como valorar en 3 simples pasos:</b>

	1. Click este link para valorar y dejar tu opinión:
	https://t.me/storebot?start={{bot}}

	2. Click en "⭐️⭐️⭐️⭐️⭐️" botón

	3. Escribe tu mensage o apreta "Skip this step" botón

	¡Muchas gracias! Merced a tus acciones más gente conocerá a bot. Todo eso sirve para una motivación adicional a los creadores! 😎`,

		"it-IT": `<b>Come valutare in 3 semplici passaggi:</b>
	1. Clicca su questo link per votare e lasciare una recensione:
	https://t.me/storebot?start={{bot}}

	2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

	3. Scrivi il tuo messaggio o premi "Salta questo step"

	Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot.Dando anche un motivo in più per continuare ai developers! 😎`,

		"fa-IR": `<b>چگونگی امتیازدهی در سه گام ساده :</b>

	1. برای امتیازدهی و ثبت نظرات بر روی لینگ زیر کلیک کنید
	https://t.me/storebot?start={{bot}}

	2. بر روی دکمه "⭐️⭐️⭐️⭐️⭐️" کلیک کنید

	3. پیام خودرا ثبت کنید یا روی دکمه "پرش از این مرحله" کلیک کنید

	بسیار سپاسگزاریم! عمل شما باعث می شود افراد بیشتری در مورد bot.All بیاموزند. این امر انگیزه مضاعفی به توسعه دهندگان این ربات می دهد ! 😎`,

		"ru-RU": `<b>Как поставить оценку в три простых шага:</b>

	1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
	https://t.me/storebot?start={{bot}}

	2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

	3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

	Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBACK: {
		"de-DE": "Über ein kleines Feedback wie der Bot so ist, würden wir uns freuen. Es dauert nur ein paar Sekunden.",
		"en-UK": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"en-US": "We would appreciate if tell us how we doing. It takes just few seconds.",
		"es-ES": "Te agredecemos si valoras el funccionamiento de nuestro applicación. Te costará solo unos segundos.",
		"fa-IR": "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
		"fr-FR": "Nous apprécierions si vous nous disiez comment nous allons. Cela ne prend que quelques secondes.",
		"id-ID": "Kami akan menghargai jika memberi tahu kami bagaimana kami melakukannya. Ini hanya membutuhkan beberapa detik.",
		"it-IT": "Ci farebbe piacere se lasciassi un voto per il nostro lavoro. Ti bastano solo alcuni secondi.",
		"ja-JP": "私たちの仕事ぶりを教えていただければ幸いです。ほんの数秒で済みます。",
		"ko-KO": "우리가 어떻게 하고 있는지 알려주시면 감사하겠습니다. 몇 초 밖에 걸리지 않습니다.",
		"pl-PL": "Bylibyśmy wdzięczni, gdybyś powiedział nam, jak sobie radzimy. To zajmie tylko kilka sekund.",
		"pt-BR": "Agradecemos se nos disser como estamos indo. Leva apenas alguns segundos.",
		"ru-RU": "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		"tr-TR": "Nasıl yaptığımızı bize söylerseniz memnun oluruz. Sadece birkaç saniye sürer.",
		"ua-UA": "Ми будемо вдячні, якщо ви розкажете нам, як ми працюємо. Це займе лише кілька секунд.",
		"uz-UZ": "Qanday qilayotganimizni bizga aytsangiz, minnatdor bo'lardik. Bu atigi bir necha soniya vaqt oladi.",
		"zh-CN": "如果您告诉我们我们的表现如何，我们将不胜感激。这只需要几秒钟。",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		"de-DE": "Bewerte diesen Bot",
		"en-UK": "Rate this bot",
		"en-US": "Rate this bot",
		"es-ES": "Valora a bot",
		"fa-IR": "به این ربات امتیاز بدهید",
		"fr-FR": "Évaluer ce bot",
		"id-ID": "Nilai bot ini",
		"it-IT": "Vota questo bot",
		"ja-JP": "このボットを評価する",
		"ko-KO": "이 봇 평가하기",
		"pl-PL": "Oceń tego bota",
		"pt-BR": "Avalie este bot",
		"ru-RU": "Оценить приложение",
		"tr-TR": "Bu botu değerlendir",
		"ua-UA": "Оцінити цього бота",
		"uz-UZ": "Bu botni baholang",
		"zh-CN": "评价这个机器人",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		"de-DE": "Mache eine Bewertung auf @Storebot",
		"en-UK": "Leave rating at @Storebot",
		"en-US": "Leave rating at @Storebot",
		"es-ES": "Valorar en @Storebot",
		"fa-IR": "امتیاز خود را اینجا وارد کنید @Storebot",
		"fr-FR": "Laisser une évaluation sur @Storebot",
		"id-ID": "Tinggalkan peringkat di @Storebot",
		"it-IT": "Lascia un voto a @Storebot",
		"ja-JP": "@Storebotで評価を残す",
		"ko-KO": "@Storebot에서 평가 남기기",
		"pl-PL": "Zostaw ocenę na @Storebot",
		"pt-BR": "Deixe sua avaliação no @Storebot",
		"ru-RU": "Оценить на  @Storebot",
		"tr-TR": "@Storebot'ta değerlendirme bırakın",
		"ua-UA": "Залиште оцінку на @Storebot",
		"uz-UZ": "@Storebot'da baho qoldiring",
		"zh-CN": "在 @Storebot 上留下评分",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		"de-DE": `Okay, vielleicht werden wir wann anders bewertet.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ansonsten freuen wir uns immer zu hören, was man besser machen kann.
	`,
		"en-UK": `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		"en-US": `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		"es-ES": `OK, Quizás puedas valorar en otro momento.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	También te agredeceríamos si ofrecieras alguna mejora.
	`,
		"fa-IR": `بسیار خوب، ممکن است شما بتوانید زمان دیگری به ما امتیاز بدهید.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	همچنین سپاسگزار خواهیم بود اگر شما هرگونه امکان بهبود را با ما در میان بگذارید.
	`,
		"fr-FR": `OK, peut-être que vous pourrez nous évaluer une autre fois.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous suggérez des améliorations.
	`,
		"id-ID": `OK, mungkin Anda dapat menilai kami lain kali.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda menyarankan perbaikan apa pun.
	`,
		"it-IT": `OK, forse ci puoi valutare un'altra volta.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Apprezzeremo anche se suggeriamo qualche miglioramento.
	`,
		"ja-JP": `OK、また今度評価してください。

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	改善点を提案していただければ幸いです。
	`,
		"ko-KO": `네, 다음에 평가해 주실 수 있을 것입니다.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	개선 사항을 제안해 주시면 감사하겠습니다.
	`,
		"pl-PL": `OK, może ocenisz nas innym razem.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli zaproponujesz jakieś ulepszenia.
	`,
		"pt-BR": `OK, talvez você possa nos avaliar em outra ocasião.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você sugerir melhorias.
	`,
		"ru-RU": `ОК, возможно вы смоежете поставить оценку в другой раз.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы предложите любые улучшения.
	`,
		"tr-TR": `Tamam, belki başka bir zaman bizi değerlendirebilirsiniz.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Herhangi bir iyileştirme önerirseniz de memnun oluruz.
	`,
		"ua-UA": `Гаразд, можливо, ви зможете оцінити нас іншим разом.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви запропонуєте будь-які покращення.
	`,
		"uz-UZ": `OK, balki boshqa safar bizni baholashingiz mumkin.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz biron bir yaxshilanishni taklif qilsangiz, biz ham minnatdor bo'lamiz.
	`,
		"zh-CN": `好的，也许您可以在另一个时间给我们评分。

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您提出任何改进建议，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		"de-DE": `Danke, wir arbeiten hart dran!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Wir freuen uns auch immer über <a suggest-idea>neue Ideen</a>.
	`,
		"en-UK": `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		"en-US": `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		"es-ES": `Gracias, hemos trabajado duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Te agradeceríamos si <a suggest-idea>propusieras alguna mejora</a>.
	`,
		"fa-IR": `ممنونیم، ما سخت کارکرده ایم!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	ما همچنین سپاسگزار خواهیم بود اگر شما<a suggest-idea> هرگونه امکان بهبود را با ما در میان بگذارید </a>را.
	`,
		"fr-FR": `Merci, nous avons travaillé dur !

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous <a suggest-idea>suggérez des améliorations</a>.
	`,
		"id-ID": `Terima kasih, kami bekerja keras!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda <a suggest-idea>menyarankan perbaikan</a>.
	`,
		"it-IT": `GRAZIE MILLE, abbiamo lavoro duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Sarebbe ancora piu' apprezzatto se ci <a suggest-idea>suggerisci qualche miglioramento</a>.
	`,
		"ja-JP": `ありがとうございます、私たちは一生懸命働きました！

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>改善を提案</a>していただければ幸いです。
	`,
		"ko-KO": `감사합니다, 우리는 열심히 일했습니다!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>개선 사항을 제안</a>해 주시면 감사하겠습니다.
	`,
		"pl-PL": `Dziękujemy, ciężko pracowaliśmy!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli <a suggest-idea>zaproponujesz ulepszenia</a>.
	`,
		"pt-BR": `Obrigado, trabalhamos duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você <a suggest-idea>sugerir melhorias</a>.
	`,
		"ru-RU": `Спасибо, мы очень старались!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
	`,
		"tr-TR": `Teşekkürler, çok çalıştık!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>İyileştirmeler önerirseniz</a> de memnun oluruz.
	`,
		"ua-UA": `Дякуємо, ми старанно працювали!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви <a suggest-idea>запропонуєте покращення</a>.
	`,
		"uz-UZ": `Rahmat, biz qattiq ishladik!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz <a suggest-idea>yaxshilanishlarni taklif qilsangiz</a>, biz ham minnatdor bo'lamiz.
	`,
		"zh-CN": `谢谢，我们努力工作！

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您<a suggest-idea>提出改进建议</a>，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		"de-DE": `
Es gibt viele Wege uns zu helfen:

* Gib uns 5⭐ im <a storebot>Verzeichnis der Bots</a>.

* Erzähl am besten all deinen Freunde davon.
Du könntest es auf <a share-fb>Facebook</a> posten oder auf <a share-twitter>Twitter</a> twittern.

* Ansonsten auch gerne eine kleine Spende - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		"en-UK": `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		"en-US": `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		"es-ES": `
Nos ayudarías mucho si:

* Nos pusieras 5⭐ en <a storebot>directory of bots</a>.

* Hablaras del bot a tus amigos.
Por ejemplo <a share-fb>Facebook</a> o <a share-twitter>Twitter</a>.

* Apoyaras el siguiente trabajo - <a href = "https://goo.gl/Qhh0yL">€2 vía PayPal</a> (<i>about $2.2</i>)
`,
		"fa-IR": `:شما به ما کمک بسیاری می کنید اگر

* ثبت بازخورد مثبت در <a storebot>دایرکتوری روبات ها</a>.

* به دوستان خود در مورد برنامه اطلاع رسانی کنید.
برای مثال در <a share-fb>فیسبوک</a> یا <a share-twitter>توئیتر</a>.

* از توسعه بیشتر برنامه پشتیبانی کنید - <a href = "https://goo.gl/Qhh0yL">€2 از طریق پی پال</a> (<i>$2.2 حدود </i>)`,
		"fr-FR": `
Vous pouvez nous aider beaucoup si vous:

* Nous donnez 5⭐ au <a storebot>répertoire des bots</a>.

* Parlez de l'application à vos amis.
Par exemple sur <a share-fb>Facebook</a> ou <a share-twitter>Twitter</a>.

* Soutenez le développement ultérieur - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>environ $2.2</i>)
`,
		"id-ID": `
Anda dapat membantu kami banyak jika Anda:

* Beri kami 5⭐ di <a storebot>direktori bot</a>.

* Ceritakan tentang aplikasi kepada teman-teman Anda.
Misalnya di <a share-fb>Facebook</a> atau <a share-twitter>Twitter</a>.

* Dukung pengembangan lebih lanjut - <a href = "https://goo.gl/Qhh0yL">€2 melalui PayPal</a> (<i>sekitar $2.2</i>)
`,
		"it-IT": `
Ci aiuteresti moltissimo se:

* Lasci un feedback positivo 5⭐ alla <a storebot>pagina del bot</a>.

* Raccontare dell'app ai tuoi amici.
Per esempio su <a share-fb>Facebook</a> o su <a share-twitter>Twitter</a>.

* Supporta ulteriormente lo sviluppo del bot - <a href="https://goo.gl/Qhh0yL">2€ tramite PayPal</a> (<i>circa $2.2</i>)
`,
		"ja-JP": `
あなたが以下のことをしてくれると、私たちにとって大きな助けになります：

* <a storebot>ボットディレクトリ</a>で5⭐を付けてください。

* アプリについて友達に教えてください。
例えば<a share-fb>Facebook</a>や<a share-twitter>Twitter</a>で。

* さらなる開発をサポートする - <a href = "https://goo.gl/Qhh0yL">PayPalで€2</a> (<i>約$2.2</i>)
`,
		"ko-KO": `
다음과 같은 방법으로 많은 도움을 주실 수 있습니다:

* <a storebot>봇 디렉토리</a>에서 5⭐을 주세요.

* 앱에 대해 친구들에게 알려주세요.
예를 들어 <a share-fb>Facebook</a> 또는 <a share-twitter>Twitter</a>에서.

* 추가 개발 지원 - <a href = "https://goo.gl/Qhh0yL">PayPal을 통해 €2</a> (<i>약 $2.2</i>)
`,
		"pl-PL": `
Możesz nam bardzo pomóc, jeśli:

* Dasz nam 5⭐ w <a storebot>katalogu botów</a>.

* Opowiesz o aplikacji swoim znajomym.
Na przykład na <a share-fb>Facebooku</a> lub <a share-twitter>Twitterze</a>.

* Wesprzyj dalszy rozwój - <a href = "https://goo.gl/Qhh0yL">€2 przez PayPal</a> (<i>około $2.2</i>)
`,
		"pt-BR": `
Você pode nos ajudar muito se:

* Nos der 5⭐ no <a storebot>diretório de bots</a>.

* Contar sobre o aplicativo para seus amigos.
Por exemplo, no <a share-fb>Facebook</a> ou <a share-twitter>Twitter</a>.

* Apoiar o desenvolvimento adicional - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>cerca de $2.2</i>)
`,
		"ru-RU": `
	Вы нам очень поможете если:

	* Поставите нам 5⭐ в <a storebot>каталоге ботов</a>.

	* Расскажите о боте своим друзьям.
	Например в <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

* Поддержите дальнейшую разработку - <a href ="https://goo.gl/Qhh0yL">€2 через PayPal</a>
`,
		"tr-TR": `
Şunları yaparsanız bize çok yardımcı olabilirsiniz:

* <a storebot>bot dizininde</a> bize 5⭐ verin.

* Uygulama hakkında arkadaşlarınıza anlatın.
Örneğin <a share-fb>Facebook</a> veya <a share-twitter>Twitter</a>'da.

* Daha fazla geliştirmeyi destekleyin - <a href = "https://goo.gl/Qhh0yL">PayPal üzerinden €2</a> (<i>yaklaşık $2.2</i>)
`,
		"ua-UA": `
Ви можете дуже допомогти нам, якщо:

* Дасте нам 5⭐ у <a storebot>каталозі ботів</a>.

* Розкажете про додаток своїм друзям.
Наприклад, у <a share-fb>Facebook</a> або <a share-twitter>Twitter</a>.

* Підтримаєте подальший розвиток - <a href = "https://goo.gl/Qhh0yL">€2 через PayPal</a> (<i>приблизно $2.2</i>)
`,
		"uz-UZ": `
Agar quyidagilarni qilsangiz, bizga juda yordam bergan bo'lasiz:

* <a storebot>botlar katalogida</a> bizga 5⭐ bering.

* Ilova haqida do'stlaringizga ayting.
Masalan, <a share-fb>Facebook</a> yoki <a share-twitter>Twitter</a>da.

* Keyingi rivojlanishni qo'llab-quvvatlang - <a href = "https://goo.gl/Qhh0yL">PayPal orqali €2</a> (<i>taxminan $2.2</i>)
`,
		"zh-CN": `
如果您：

* 在<a storebot>机器人目录</a>中给我们5⭐。

* 向您的朋友介绍这个应用程序。
例如在<a share-fb>Facebook</a>或<a share-twitter>Twitter</a>上。

* 支持进一步开发 - <a href = "https://goo.gl/Qhh0yL">通过PayPal支付€2</a> (<i>约$2.2</i>)
`,
	},
	MESSAGE_TEXT_COUNTERPARTY_HAS_EMPTY_BALANCE: {
		"de-DE": `Du bist quitt mit %v`,
		"en-UK": `Balance is empty for %v`,
		"en-US": `Balance is empty for %v`,
		"es-ES": `El balance es cero para %v`,
		"fa-IR": `تراز خالی است برای %v`,
		"fr-FR": `Le solde est vide pour %v`,
		"id-ID": `Saldo kosong untuk %v`,
		"it-IT": `Non hai alcun credito o debito con %v`,
		"ja-JP": `%vの残高は空です`,
		"ko-KO": `%v에 대한 잔액이 비어 있습니다`,
		"pl-PL": `Saldo jest puste dla %v`,
		"pt-BR": `Saldo está vazio para %v`,
		"ru-RU": `Нулевой баланс для %v`,
		"tr-TR": `%v için bakiye boş`,
		"ua-UA": `Баланс порожній для %v`,
		"uz-UZ": `%v uchun balans bo'sh`,
		"zh-CN": `%v的余额为空`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		"de-DE": `Möchtest du den Bot in einer anderen Sprache? Du kannst beim <a>Übersetzen helfen</a>.`,
		"en-UK": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		"en-US": `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		"es-ES": `¿Te gustaría que nuestro bot hablara en otro idioma? Puedes <a>ayudar con traducción</a>.`,
		"fa-IR": `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,
		"fr-FR": `Voulez-vous que notre bot parle dans une autre langue? Vous pouvez <a>aider à la traduction</a>.`,
		"id-ID": `Apakah Anda ingin bot kami berbicara dalam bahasa lain? Anda dapat <a>membantu dengan terjemahan</a>.`,
		"it-IT": `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
		"ja-JP": `私たちのボットが他の言語で話すことを望みますか？あなたは<a>翻訳を手伝う</a>ことができます。`,
		"ko-KO": `우리 봇이 다른 언어로 말하기를 원하십니까? <a>번역을 도울 수 있습니다</a>.`,
		"pl-PL": `Czy chcesz, aby nasz bot mówił w innym języku? Możesz <a>pomóc w tłumaczeniu</a>.`,
		"pt-BR": `Você quer que nosso bot fale em outro idioma? Você pode <a>ajudar com a tradução</a>.`,
		"ru-RU": `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		"tr-TR": `Botumuzun başka bir dilde konuşmasını ister misiniz? <a>Çeviriye yardımcı olabilirsiniz</a>.`,
		"ua-UA": `Хочете, щоб наш бот розмовляв іншою мовою? Ви можете <a>допомогти з перекладом</a>.`,
		"uz-UZ": `Botimiz boshqa tilda gaplashishini xohlaysizmi? Siz <a>tarjima qilishda yordam</a> berishingiz mumkin.`,
		"zh-CN": `您想要我们的机器人用其他语言交谈吗？您可以<a>帮助翻译</a>。`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		"de-DE": `Gut, wir geben uns Mühe. Deine Rückmeldung wird an die Entwickler weitergeleitet.

Vielleicht willst du <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		"en-UK": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		"en-US": `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		"es-ES": `Bueno, hemos trabajado duro. Tu opinión se pasará a los creadores.

Quizás puedas <a submit-bug>informarnos de algún problema</a> o <a suggest-idea>proponernos qué podríamos mejorar</a>?`,
		"fa-IR": `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		"fr-FR": `Eh bien, nous avons travaillé dur. Vos commentaires seront transmis aux développeurs.

Peut-être pouvez-vous <a submit-bug>signaler votre problème</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		"id-ID": `Yah, kami bekerja keras. Umpan balik Anda akan diteruskan ke pengembang.

Mungkin Anda dapat <a submit-bug>melaporkan masalah Anda</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		"it-IT": `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
		"ja-JP": `まあ、私たちは一生懸命働きました。あなたのフィードバックは開発者に伝えられます。

<a submit-bug>問題を報告</a>したり、<a suggest-idea>改善方法を提案</a>したりすることもできますか？`,
		"ko-KO": `음, 우리는 열심히 일했습니다. 귀하의 피드백은 개발자에게 전달됩니다.

<a submit-bug>문제를 보고</a>하거나 <a suggest-idea>개선 방법을 제안</a>할 수 있습니까?`,
		"pl-PL": `Cóż, ciężko pracowaliśmy. Twoja opinia zostanie przekazana deweloperom.

Może możesz <a submit-bug>zgłosić swój problem</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		"pt-BR": `Bem, trabalhamos duro. Seu feedback será passado para os desenvolvedores.

Talvez você possa <a submit-bug>relatar seu problema</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		"ru-RU": `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		"tr-TR": `Eh, çok çalıştık. Geri bildiriminiz geliştiricilere iletilecektir.

Belki <a submit-bug>sorununuzu bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		"ua-UA": `Ну, ми старанно працювали. Ваш відгук буде переданий розробникам.

Можливо, ви можете <a submit-bug>повідомити про свою проблему</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		"uz-UZ": `Xo'sh, biz qattiq ishladik. Sizning fikr-mulohazalaringiz ishlab chiqaruvchilarga yetkaziladi.

Balki siz <a submit-bug>muammoingizni xabar qilishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		"zh-CN": `嗯，我们努力工作。您的反馈将传递给开发人员。

也许您可以<a submit-bug>报告您的问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		"de-DE": `Das tut uns sehr leid. Vielleicht willst du uns <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		"en-UK": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		"en-US": `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		"es-ES": `Lo sentimos mucho. Igual podrías <a submit-bug>decirnos qué no funcciona bien</a> o <a suggest-idea>proponernos cómo podemos mejorarlo</a>?`,
		"fa-IR": `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		"fr-FR": `Nous sommes vraiment désolés. Peut-être pouvez-vous <a submit-bug>nous faire savoir ce qui ne va pas</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		"id-ID": `Kami sangat menyesal. Mungkin Anda dapat <a submit-bug>memberi tahu kami apa yang salah</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		"it-IT": `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
		"ja-JP": `大変申し訳ありません。<a submit-bug>何が問題かを教えていただく</a>か、<a suggest-idea>改善方法を提案していただく</a>ことはできますか？`,
		"ko-KO": `매우 죄송합니다. <a submit-bug>무엇이 잘못되었는지 알려주시거나</a> <a suggest-idea>개선 방법을 제안</a>해 주실 수 있습니까?`,
		"pl-PL": `Bardzo nam przykro. Może możesz <a submit-bug>dać nam znać, co jest nie tak</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		"pt-BR": `Lamentamos muito. Talvez você possa <a submit-bug>nos informar o que está errado</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		"ru-RU": `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		"tr-TR": `Çok üzgünüz. Belki <a submit-bug>neyin yanlış olduğunu bize bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		"ua-UA": `Нам дуже шкода. Можливо, ви можете <a submit-bug>повідомити нам, що не так</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		"uz-UZ": `Juda uzr so'raymiz. Balki siz <a submit-bug>nima noto'g'ri ekanligini bizga aytib berishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		"zh-CN": `我们非常抱歉。也许您可以<a submit-bug>告诉我们哪里出了问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		"de-DE": "Bitte bewerte unsere App",
		"en-UK": "Please rate our app",
		"en-US": "Please rate our app",
		"es-ES": "Por favor valora nuestro app",
		"fa-IR": "لطفاً به برنامه ما امتیاز دهید",
		"fr-FR": "Veuillez évaluer notre application",
		"id-ID": "Silakan nilai aplikasi kami",
		"it-IT": "Per favore vota il nostro bot",
		"ja-JP": "アプリを評価してください",
		"ko-KO": "앱을 평가해 주세요",
		"pl-PL": "Oceń naszą aplikację",
		"pt-BR": "Por favor, avalie nosso aplicativo",
		"ru-RU": "Оцените наше приложение?",
		"tr-TR": "Lütfen uygulamamızı değerlendirin",
		"ua-UA": "Будь ласка, оцініть наш додаток",
		"uz-UZ": "Iltimos, ilovamizni baholang",
		"zh-CN": "请评价我们的应用",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		"de-DE": "Ja, es ist eine tolle App!",
		"en-UK": "Yes, it's a great app!",
		"en-US": "Yes, it's a great app!",
		"es-ES": "¡Sí, es una app fantástica!",
		"fa-IR": "بله، این برنامه عالی است",
		"fr-FR": "Oui, c'est une excellente application !",
		"id-ID": "Ya, ini aplikasi yang bagus!",
		"it-IT": "Si, e' un app fantastica!",
		"ja-JP": "はい、素晴らしいアプリです！",
		"ko-KO": "네, 훌륭한 앱입니다!",
		"pl-PL": "Tak, to świetna aplikacja!",
		"pt-BR": "Sim, é um ótimo aplicativo!",
		"ru-RU": "Да, отличное приложение!",
		"tr-TR": "Evet, harika bir uygulama!",
		"ua-UA": "Так, це чудовий додаток!",
		"uz-UZ": "Ha, bu ajoyib ilova!",
		"zh-CN": "是的，这是一个很棒的应用！",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		"de-DE": "Nicht schlecht, aber könnte besser sein",
		"en-UK": "Not bad but can be better.",
		"en-US": "Not bad but can be better.",
		"es-ES": "No está mal, pero podría ser mejor.",
		"fa-IR": "بد نیست ولی می تواند بهتر باشد.",
		"fr-FR": "Pas mal mais peut être amélioré.",
		"id-ID": "Tidak buruk tapi bisa lebih baik.",
		"it-IT": "Non male ma potrebbe esser migliore.",
		"ja-JP": "悪くないですが、もっと良くなる可能性があります。",
		"ko-KO": "나쁘지 않지만 더 좋아질 수 있습니다.",
		"pl-PL": "Niezły, ale może być lepszy.",
		"pt-BR": "Não é ruim, mas pode ser melhor.",
		"ru-RU": "Неплохо, но можно лучше.",
		"tr-TR": "Fena değil ama daha iyi olabilir.",
		"ua-UA": "Непогано, але може бути краще.",
		"uz-UZ": "Yomon emas, lekin yaxshiroq bo'lishi mumkin.",
		"zh-CN": "不错，但可以更好。",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		"de-DE": "Ich mag sie nicht",
		"en-UK": "Don't like it",
		"en-US": "Don't like it",
		"es-ES": "No me gusta",
		"fa-IR": "از این برنامه را نمی پسندم",
		"fr-FR": "Je n'aime pas",
		"id-ID": "Tidak suka",
		"it-IT": "Non mi piace",
		"ja-JP": "好きではない",
		"ko-KO": "마음에 들지 않음",
		"pl-PL": "Nie lubię tego",
		"pt-BR": "Não gosto",
		"ru-RU": "Не нравится",
		"tr-TR": "Beğenmedim",
		"ua-UA": "Не подобається",
		"uz-UZ": "Yoqmadi",
		"zh-CN": "不喜欢",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		"de-DE": "Ich bin mir noch unsicher",
		"en-UK": "Not decided yet",
		"en-US": "Not decided yet",
		"es-ES": "Estoy aún indeciso",
		"fa-IR": "هنوز تصمیم نگرفته ام.",
		"fr-FR": "Pas encore décidé",
		"id-ID": "Belum memutuskan",
		"it-IT": "Sono indeciso",
		"ja-JP": "まだ決めていません",
		"ko-KO": "아직 결정하지 않았습니다",
		"pl-PL": "Jeszcze nie zdecydowałem",
		"pt-BR": "Ainda não decidi",
		"ru-RU": "Пока не понятно",
		"tr-TR": "Henüz karar vermedim",
		"ua-UA": "Ще не вирішив",
		"uz-UZ": "Hali qaror qabul qilinmagan",
		"zh-CN": "尚未决定",
	},
	MESSAGE_TEXT_SETTINGS: {
		"de-DE": "Was willst du ändern?",
		"en-UK": "What do you want to adjust?",
		"en-US": "What do you want to adjust?",
		"es-ES": "¿Qué te gustaría modificar?",
		"fa-IR": "می خواهید چه چیزی را تنظیم کنید؟",
		"fr-FR": "Que voulez-vous ajuster ?",
		"id-ID": "Apa yang ingin Anda sesuaikan?",
		"it-IT": "Cosa vuoi modificare?",
		"ja-JP": "何を調整したいですか？",
		"ko-KO": "무엇을 조정하고 싶으신가요?",
		"pl-PL": "Co chcesz dostosować?",
		"pt-BR": "O que você deseja ajustar?",
		"ru-RU": "Что будем настраивать?",
		"tr-TR": "Neyi ayarlamak istiyorsunuz?",
		"ua-UA": "Що ви хочете налаштувати?",
		"uz-UZ": "Nimani sozlamoqchisiz?",
		"zh-CN": "您想调整什么？",
	},
	MT_NO_OUTSTANDING_TRANSFERS: {
		"de-DE": `Sie versuchen, einen Rückgabedatensatz zu erstellen, aber es gibt keine ausstehenden Schulden.

Wenn Sie glauben, dass dies ein Fehler ist, teilen Sie uns dies bitte in @DebtsTrackerGroup mit.`,
		"en-UK": `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		"en-US": `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		"es-ES": `Estás intentando crear un registro de devolución pero no hay deudas pendientes.

Si crees que esto es un error, háganoslo saber en @DebtsTrackerGroup.`,
		"fa-IR": `شما در حال تلاش برای ایجاد سابقه بازگشت هستید اما هیچ بدهی معوقه ای وجود ندارد.

اگر فکر می کنید این یک اشتباه است، لطفاً به ما در @DebtsTrackerGroup اطلاع دهید.`,
		"fr-FR": `Vous essayez de créer un enregistrement de retour mais il n'y a pas de dettes en cours.

Si vous pensez qu'il s'agit d'une erreur, veuillez nous en informer dans @DebtsTrackerGroup.`,
		"id-ID": `Anda mencoba membuat catatan pengembalian tetapi tidak ada hutang yang belum dibayar.

Jika Anda yakin ini adalah kesalahan, beri tahu kami di @DebtsTrackerGroup.`,
		"it-IT": `Stai cercando di creare un record di restituzione ma non ci sono debiti in sospeso.

Se ritieni che si tratti di un errore, ti preghiamo di farcelo sapere in @DebtsTrackerGroup.`,
		"ja-JP": `返却記録を作成しようとしていますが、未払いの借金はありません。

これが間違いだと思われる場合は、@DebtsTrackerGroupでお知らせください。`,
		"ko-KO": `반환 기록을 생성하려고 하지만 미결제 부채가 없습니다.

이것이 실수라고 생각되면 @DebtsTrackerGroup에서 알려주십시오.`,
		"pl-PL": `Próbujesz utworzyć rekord zwrotu, ale nie ma zaległych długów.

Jeśli uważasz, że to pomyłka, daj nam znać w @DebtsTrackerGroup.`,
		"pt-BR": `Você está tentando criar um registro de devolução, mas não há dívidas pendentes.

Se você acredita que isso é um erro, por favor, nos avise em @DebtsTrackerGroup.`,
		"ru-RU": `Вы пытаетесь создать запись о возврате долга, но мы не нашли не закрытых долгов.

Если вы считаете что это ошибка пожалуйста дайте нам знать в @DebtsTrackerGroup.`,
		"tr-TR": `İade kaydı oluşturmaya çalışıyorsunuz ancak bekleyen borç bulunmamaktadır.

Bunun bir hata olduğunu düşünüyorsanız, lütfen @DebtsTrackerGroup'ta bize bildirin.`,
		"ua-UA": `Ви намагаєтеся створити запис про повернення, але немає непогашених боргів.

Якщо ви вважаєте, що це помилка, будь ласка, повідомте нам у @DebtsTrackerGroup.`,
		"uz-UZ": `Siz qaytarish yozuvini yaratmoqchi bo'lyapsiz, lekin to'lanmagan qarzlar yo'q.

Agar bu xato deb o'ylasangiz, iltimos, bizga @DebtsTrackerGroup'da xabar bering.`,
		"zh-CN": `您正在尝试创建退款记录，但没有未偿还的债务。

如果您认为这是一个错误，请在 @DebtsTrackerGroup 中告诉我们。`,
	},
	MT_ATTEMPT_TO_CREATE_DEBT_WITH_INTEREST_AFFECTING_OUTSTANDING: {
		"de-DE": "Sie versuchen, eine Schuld mit Zinsen zu erstellen, die sich auf ausstehende Überweisungen auswirken wird. Bitte schließen Sie diese zuerst.",
		"en-UK": "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		"en-US": "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		"es-ES": "Estás intentando crear una deuda con intereses que afectará a las transferencias pendientes. Por favor, ciérralas primero.",
		"fa-IR": "شما در حال تلاش برای ایجاد بدهی با بهره هستید که بر انتقال های معوق تأثیر می گذارد. لطفا ابتدا آنها را ببندید.",
		"fr-FR": "Vous essayez de créer une dette avec intérêt qui affectera les transferts en cours. Veuillez les fermer d'abord.",
		"id-ID": "Anda mencoba membuat hutang dengan bunga yang akan memengaruhi transfer yang belum diselesaikan. Harap tutup terlebih dahulu.",
		"it-IT": "Stai cercando di creare un debito con interessi che influenzerà i trasferimenti in sospeso. Per favore, chiudili prima.",
		"ja-JP": "未決済の送金に影響する利息付きの借金を作成しようとしています。まずそれらを閉じてください。",
		"ko-KO": "미결제 송금에 영향을 미치는 이자가 있는 부채를 만들려고 합니다. 먼저 그것들을 닫으십시오.",
		"pl-PL": "Próbujesz utworzyć dług z odsetkami, który wpłynie na nierozliczone przelewy. Proszę najpierw je zamknąć.",
		"pt-BR": "Você está tentando criar uma dívida com juros que afetará transferências pendentes. Por favor, feche-as primeiro.",
		"ru-RU": "Вы пытаетесь создать запись о долге с процентами которая затронет незакрытые долги. Пожалуйста закройте их сначала.",
		"tr-TR": "Bekleyen transferleri etkileyecek faizli bir borç oluşturmaya çalışıyorsunuz. Lütfen önce onları kapatın.",
		"ua-UA": "Ви намагаєтеся створити борг з відсотками, який вплине на невиконані перекази. Будь ласка, спочатку закрийте їх.",
		"uz-UZ": "Siz to'lanmagan o'tkazmalarga ta'sir qiladigan foizli qarz yaratmoqchisiz. Iltimos, avval ularni yoping.",
		"zh-CN": "您正在尝试创建一个带有利息的债务，这将影响未结清的转账。请先关闭它们。",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		"de-DE": "Entschuldigung, diese Funktion ist noch nicht fertig programmiert.",
		"en-UK": "Sorry, this functionality is not implemented yet.",
		"en-US": "Sorry, this functionality is not implemented yet.",
		"es-ES": "Lo sentimos, esta función no está activa aún.",
		"fa-IR": "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
		"fr-FR": "Désolé, cette fonctionnalité n'est pas encore implémentée.",
		"id-ID": "Maaf, fungsi ini belum diimplementasikan.",
		"it-IT": "Spiacenti ma questa funzionalita' non e' ancora attiva.",
		"ja-JP": "申し訳ありませんが、この機能はまだ実装されていません。",
		"ko-KO": "죄송합니다. 이 기능은 아직 구현되지 않았습니다.",
		"pl-PL": "Przepraszamy, ta funkcjonalność nie jest jeszcze zaimplementowana.",
		"pt-BR": "Desculpe, esta funcionalidade ainda não foi implementada.",
		"ru-RU": "Извините, данный функционал ещё не реализован.",
		"tr-TR": "Üzgünüz, bu işlevsellik henüz uygulanmadı.",
		"ua-UA": "Вибачте, ця функціональність ще не реалізована.",
		"uz-UZ": "Kechirasiz, bu funksionallik hali amalga oshirilmagan.",
		"zh-CN": "抱歉，此功能尚未实现。",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		"de-DE": "Worüber möchtest du eingeladen werden?",
		"en-UK": "How do you want to get an invite?",
		"en-US": "How do you want to get an invite?",
		"es-ES": "¿Comó prefieres recibir la invitación?",
		"fa-IR": "می خواهید چگونه دعوت شوید؟",
		"fr-FR": "Comment voulez-vous recevoir une invitation ?",
		"id-ID": "Bagaimana Anda ingin mendapatkan undangan?",
		"it-IT": "Come vuoi ricevere l'invito?",
		"ja-JP": "招待状をどのように受け取りますか？",
		"ko-KO": "초대를 어떻게 받고 싶으신가요?",
		"pl-PL": "Jak chcesz otrzymać zaproszenie?",
		"pt-BR": "Como você deseja receber um convite?",
		"ru-RU": "Как вы хотите получить код приглашения?",
		"tr-TR": "Daveti nasıl almak istersiniz?",
		"ua-UA": "Як ви хочете отримати запрошення?",
		"uz-UZ": "Taklifnomani qanday olishni xohlaysiz?",
		"zh-CN": "您想如何获得邀请？",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		"de-DE": "Bitte gib den Bestätigungs-Code ein:",
		"en-UK": "Please enter an invite code:",
		"en-US": "Please enter an invite code:",
		"es-ES": "Introduce el código de la invitación",
		"fa-IR": "لطفاً یک کد دعوت وارد کنید:",
		"fr-FR": "Veuillez entrer un code d'invitation :",
		"id-ID": "Silakan masukkan kode undangan:",
		"it-IT": "Inserisci un codice invito:",
		"ja-JP": "招待コードを入力してください：",
		"ko-KO": "초대 코드를 입력하세요:",
		"pl-PL": "Proszę wprowadzić kod zaproszenia:",
		"pt-BR": "Por favor, insira um código de convite:",
		"ru-RU": "Пожалуйста введите код приглашения:",
		"tr-TR": "Lütfen bir davet kodu girin:",
		"ua-UA": "Будь ласка, введіть код запрошення:",
		"uz-UZ": "Iltimos, taklif kodini kiriting:",
		"zh-CN": "请输入邀请码：",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		"de-DE": "Wir haben eine Nachricht an %v gesendet.\n\nBitte öffne die Nachricht und klick auf den Link, um deine Mail-Adresse zu bestätigen.",
		"en-UK": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"en-US": "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		"es-ES": "Hemos enviado un mensage a %v.\n\nPor favor, abre tu e-mail y haz click en el link para confirmar tu e-mail.",
		"fa-IR": "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
		"fr-FR": "Nous avons envoyé un message à %v.\n\nVeuillez ouvrir l'e-mail et cliquer sur un lien pour confirmer votre adresse e-mail.",
		"id-ID": "Kami telah mengirim pesan ke %v.\n\nSilakan buka email dan klik tautan untuk mengonfirmasi alamat email Anda.",
		"it-IT": "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
		"ja-JP": "%vにメッセージを送信しました。\n\nメールを開き、リンクをクリックしてメールアドレスを確認してください。",
		"ko-KO": "%v에 메시지를 보냈습니다.\n\n이메일을 열고 링크를 클릭하여 이메일 주소를 확인하십시오.",
		"pl-PL": "Wysłaliśmy wiadomość do %v.\n\nProszę otworzyć e-mail i kliknąć link, aby potwierdzić swój adres e-mail.",
		"pt-BR": "Enviamos uma mensagem para %v.\n\nPor favor, abra o e-mail e clique em um link para confirmar seu endereço de e-mail.",
		"ru-RU": "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		"tr-TR": "%v adresine bir mesaj gönderdik.\n\nLütfen e-postayı açın ve e-posta adresinizi onaylamak için bir bağlantıya tıklayın.",
		"ua-UA": "Ми надіслали повідомлення на %v.\n\nБудь ласка, відкрийте електронний лист і натисніть посилання, щоб підтвердити свою електронну адресу.",
		"uz-UZ": "Biz %v ga xabar yubordik.\n\nIltimos, elektron pochtani oching va elektron pochta manzilingizni tasdiqlash uchun havolani bosing.",
		"zh-CN": "我们已向 %v 发送了一条消息。\n\n请打开电子邮件并点击链接确认您的电子邮件地址。",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		"de-DE": "Wenn Telegram öffnet, drücke auf <b>Start</b>.",
		"en-UK": "Once Telegram pop ups click the <b>Start</b> button.",
		"en-US": "Once Telegram pop ups click the <b>Start</b> button.",
		"es-ES": "Después de abrir Telegram aprieta el <b>Start</b> botón.",
		"fa-IR": "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
		"fr-FR": "Une fois que Telegram apparaît, cliquez sur le bouton <b>Start</b>.",
		"id-ID": "Setelah Telegram muncul, klik tombol <b>Start</b>.",
		"it-IT": "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
		"ja-JP": "Telegramがポップアップしたら、<b>Start</b>ボタンをクリックしてください。",
		"ko-KO": "Telegram이 팝업되면 <b>Start</b> 버튼을 클릭하세요.",
		"pl-PL": "Gdy pojawi się Telegram, kliknij przycisk <b>Start</b>.",
		"pt-BR": "Quando o Telegram aparecer, clique no botão <b>Start</b>.",
		"ru-RU": "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		"tr-TR": "Telegram açıldığında <b>Start</b> düğmesine tıklayın.",
		"ua-UA": "Коли з'явиться Telegram, натисніть кнопку <b>Start</b>.",
		"uz-UZ": "Telegram ochilganda, <b>Start</b> tugmasini bosing.",
		"zh-CN": "Telegram 弹出后，点击 <b>Start</b> 按钮。",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		"de-DE": "Danke, du bist in der Warteschlange für eine Einladung.\n\nEs dauert etwa zwei bis drei Tage.\n\nAber du könntest den Code noch heute bekommen, wenn du einen Link auf Facebook teilst.",
		"en-UK": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"en-US": "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		"es-ES": "Gracias, ya estás inscrito en la cola para conseguir la invitación.\n\nTiempo de espera 2-3 días.\n\nPuedes conseguirlo hoy si compartes el link de nuestro bot en Facebook.",
		"fa-IR": "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
		"fr-FR": "Merci, vous avez été mis en file d'attente pour une invitation.\n\nLe temps d'attente actuel est de 2 à 3 jours.\n\nVous pouvez obtenir un code d'invitation aujourd'hui en partageant un lien vers le bot sur Facebook.",
		"id-ID": "Terima kasih, Anda telah antri untuk undangan.\n\nWaktu menunggu saat ini adalah 2-3 hari.\n\nAnda dapat memperoleh kode undangan hari ini dengan membagikan tautan ke bot di Facebook.",
		"it-IT": "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
		"ja-JP": "ありがとうございます。招待のために列に並んでいます。\n\n現在の待ち時間は2〜3日です。\n\nFacebookでボットへのリンクを共有することで、今日招待コードを取得できます。",
		"ko-KO": "감사합니다. 초대를 위해 대기열에 올랐습니다.\n\n현재 대기 시간은 2-3일입니다.\n\nFacebook에서 봇에 대한 링크를 공유하여 오늘 초대 코드를 받을 수 있습니다.",
		"pl-PL": "Dziękujemy, zostałeś umieszczony w kolejce do zaproszenia.\n\nObecny czas oczekiwania to 2-3 dni.\n\nMożesz otrzymać kod zaproszenia już dziś, udostępniając link do bota na Facebooku.",
		"pt-BR": "Obrigado, você foi colocado na fila para um convite.\n\nO tempo de espera atual é de 2 a 3 dias.\n\nVocê pode obter um código de convite hoje compartilhando um link para o bot no Facebook.",
		"ru-RU": "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		"tr-TR": "Teşekkürler, bir davetiye için sıraya alındınız.\n\nMevcut bekleme süresi 2-3 gündür.\n\nBotun bağlantısını Facebook'ta paylaşarak bugün bir davet kodu alabilirsiniz.",
		"ua-UA": "Дякуємо, вас поставлено в чергу на запрошення.\n\nПоточний час очікування - 2-3 дні.\n\nВи можете отримати код запрошення сьогодні, поділившись посиланням на бота у Facebook.",
		"uz-UZ": "Rahmat, siz taklif uchun navbatga qo'yildingiz.\n\nHozirgi kutish vaqti 2-3 kun.\n\nFacebookda botga havola ulashish orqali bugun taklif kodini olishingiz mumkin.",
		"zh-CN": "谢谢，您已排队等候邀请。\n\n当前等待时间为2-3天。\n\n您可以通过在Facebook上分享机器人链接，今天就获得邀请码。",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		"de-DE": "Bitte gib deine e-Mail Adresse an:",
		"en-UK": "Please provide your email address",
		"en-US": "Please provide your email address",
		"es-ES": "Por favor, esctibe tu e-mail",
		"fa-IR": "لطفاً آدرس ایمیل خود را وارد کنید.",
		"fr-FR": "Veuillez fournir votre adresse e-mail",
		"id-ID": "Silakan berikan alamat email Anda",
		"it-IT": "Inserisci il tuo indirizzo email:",
		"ja-JP": "メールアドレスを入力してください",
		"ko-KO": "이메일 주소를 입력해 주세요",
		"pl-PL": "Proszę podać swój adres e-mail",
		"pt-BR": "Por favor, forneça seu endereço de e-mail",
		"ru-RU": "Пожалуйста напишите ваш email адрес:",
		"tr-TR": "Lütfen e-posta adresinizi girin",
		"ua-UA": "Будь ласка, вкажіть вашу електронну адресу",
		"uz-UZ": "Iltimos, elektron pochta manzilingizni kiriting",
		"zh-CN": "请提供您的电子邮件地址",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		"de-DE": "Bitte gib deine Telefonnummer an:",
		"en-UK": "Please provide your phone number",
		"en-US": "Please provide your phone number",
		"es-ES": "Por favor, esctibe tu número de teléfono",
		"fa-IR": "لطفاً شماره تلفن خود را وارد نمایید.",
		"fr-FR": "Veuillez fournir votre numéro de téléphone",
		"id-ID": "Silakan berikan nomor telepon Anda",
		"it-IT": "Inserisci il tuo numero di telefono:",
		"ja-JP": "電話番号を入力してください",
		"ko-KO": "전화번호를 입력해 주세요",
		"pl-PL": "Proszę podać swój numer telefonu",
		"pt-BR": "Por favor, forneça seu número de telefone",
		"ru-RU": "Пожалуйста напишите номер вашего телефона:",
		"tr-TR": "Lütfen telefon numaranızı girin",
		"ua-UA": "Будь ласка, вкажіть ваш номер телефону",
		"uz-UZ": "Iltimos, telefon raqamingizni kiriting",
		"zh-CN": "请提供您的电话号码",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		"de-DE": "Ungültiger Code: %v",
		"en-UK": "Wrong invite code: %v",
		"en-US": "Wrong invite code: %v",
		"es-ES": "El código de invitación no es correcto: %v",
		"fa-IR": "کد دعوت اشتباه است %v",
		"fr-FR": "Code d'invitation incorrect : %v",
		"id-ID": "Kode undangan salah: %v",
		"it-IT": "Codice invito: %v errato",
		"ja-JP": "招待コードが間違っています：%v",
		"ko-KO": "초대 코드가 잘못되었습니다: %v",
		"pl-PL": "Nieprawidłowy kod zaproszenia: %v",
		"pt-BR": "Código de convite errado: %v",
		"ru-RU": "Неправильный код приглашения: %v",
		"tr-TR": "Yanlış davet kodu: %v",
		"ua-UA": "Неправильний код запрошення: %v",
		"uz-UZ": "Noto'g'ri taklif kodi: %v",
		"zh-CN": "邀请码错误：%v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		"de-DE": "Ungültige e-Mail Adresse.",
		"en-UK": "Wrong email address.",
		"en-US": "Wrong email address.",
		"es-ES": "El e-mail no es correcto.",
		"fa-IR": "آدرس ایمیل اشتباه است.",
		"fr-FR": "Adresse e-mail incorrecte.",
		"id-ID": "Alamat email salah.",
		"it-IT": "L'email inserita e' sbagliata.",
		"ja-JP": "メールアドレスが間違っています。",
		"ko-KO": "이메일 주소가 잘못되었습니다.",
		"pl-PL": "Nieprawidłowy adres e-mail.",
		"pt-BR": "Endereço de e-mail incorreto.",
		"ru-RU": "Неправильный email адрес.",
		"tr-TR": "Yanlış e-posta adresi.",
		"ua-UA": "Неправильна електронна адреса.",
		"uz-UZ": "Noto'g'ri elektron pochta manzili.",
		"zh-CN": "电子邮件地址错误。",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		"de-DE": "Ungültige Telefonnummer.",
		"en-UK": "Wrong phone number.",
		"en-US": "Wrong phone number.",
		"es-ES": "El número de telefono no es correcto.",
		"fa-IR": "شماره تلفن اشتباه است",
		"fr-FR": "Numéro de téléphone incorrect.",
		"id-ID": "Nomor telepon salah.",
		"it-IT": "Il numero inserito e' sbagliato.",
		"ja-JP": "電話番号が間違っています。",
		"ko-KO": "전화번호가 잘못되었습니다.",
		"pl-PL": "Nieprawidłowy numer telefonu.",
		"pt-BR": "Número de telefone incorreto.",
		"ru-RU": "Неправильный номер телефона.",
		"tr-TR": "Yanlış telefon numarası.",
		"ua-UA": "Неправильний номер телефону.",
		"uz-UZ": "Noto'g'ri telefon raqami.",
		"zh-CN": "电话号码错误。",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		"de-DE": "Ok, bitte versuche es erneut.",
		"en-UK": "Ok, please try again.",
		"es-ES": "Ok, inténtalo de nuevo.",
		"fa-IR": "بسیار خوب، لطفا مجدداً سعی کنید.",
		"it-IT": "Ok, prova di nuovo.",
		"ru-RU": "Хорошо, попробуйте ещё раз.",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		"de-DE": "Ich habe mich vertippt.",
		"en-UK": "I've mistyped, will try again.",
		"es-ES": "Me he equivocado, lo intentaré otra vez",
		"fa-IR": "اشتباه تایپ کردم، دوباره امتحان می کنم.",
		"it-IT": "Ho sbagliato, riprovo.",
		"ru-RU": "Я опечатался, попробую ещё раз.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		"de-DE": "Erzähl mir mehr über diese Codes",
		"en-UK": "Tell me more about the codes",
		"es-ES": "Dime más de los códigos",
		"fa-IR": "در مورد کدها بیشتر برای من توضیح دهید.",
		"it-IT": "Ulteriori informazioni riguardo il codice invito.",
		"ru-RU": "Расскажите ка мне об этих кодах",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		"de-DE": "Sende mir einen Code per Mail",
		"en-UK": "Send me invite code by email",
		"es-ES": "Envíame el código de invitación por e-mail",
		"fa-IR": "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite email",
		"ru-RU": "Хочу код приглашения на email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		"de-DE": "Sende mir einen Code per SMS",
		"en-UK": "Send me invite code by SMS",
		"es-ES": "Envíame el código de invitación a través de SMS",
		"fa-IR": "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
		"it-IT": "Inviami il codice invito tramite SMS",
		"ru-RU": "Хочу код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		"de-DE": "Sende mir nochmal einen Code per Mail",
		"en-UK": "Send me new invite code by email",
		"es-ES": "Envíame un nuevo código de invitación por e-mail",
		"fa-IR": "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite email",
		"ru-RU": "Новый код приглашения на email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		"de-DE": "Sende mir nochmal einen Code per SMS",
		"en-UK": "Send me new invite code by SMS",
		"es-ES": "Envíame un nuevo código de invitación a través de SMS",
		"fa-IR": "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite SMS",
		"ru-RU": "Новый код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		"de-DE": "Sende mir nochmal einen Code per Telegram",
		"en-UK": "Send me new invite by Telegram",
		"es-ES": "Envíame un nuevo código de invitación a través de Telegram",
		"fa-IR": "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
		"it-IT": "Inviami il nuovo codice invito tramite Telegram",
		"ru-RU": "Получить приграшение в Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		"de-DE": "Unbekannte Sprache. Bitte wähle eine aus der Liste:",
		"en-UK": "Unknown language. Please choose 1 from the options:",
		"es-ES": "Idioma desconocido. Por favor elige 1 de las opciones:",
		"fa-IR": "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
		"it-IT": "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
		"ru-RU": "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		"de-DE": "Unbekannte Gegnerpartei. Bitte wähle aus der Liste oder wähle <b>neuer Kontakt</b>.",
		"en-UK": "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		"es-ES": "Contacto desconocido. Por favor elige <b>Añadir</b> si es un contacto nuevo.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
		"it-IT": "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		"de-DE": "Unbekannte Gegnerpartei. Bitte wähle aus der Liste.",
		"en-UK": "Unknown counterparty. Please choose from the list.",
		"es-ES": "Contacto desconocido. Por favor elige de la lista.",
		"fa-IR": "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
		"it-IT": "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
		"ru-RU": "Неизвестный контакт. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		"de-DE": "Unbekannter Schuldschein. Bitte wähle aus der Liste.",
		"en-UK": "Unknown debt. Please choose from the list.",
		"es-ES": "Deuda desconocida. Por favor elige de la lista.",
		"fa-IR": "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
		"it-IT": "Debito sconosciuto. Scegli dalla lista qui sotto.",
		"ru-RU": "Неизвестный долг. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_BILL_CARD_HEADER: {
		"de-DE": `<b>Rechnung</b>: <code>%v</code> — %v`,
		"en-UK": `<b>Bill</b>: <code>%v</code> — %v`,
		"es-ES": `<b>Cuenta</b>: <code>%v</code> — %v`,
		"fa-IR": "<b>Bill</b>: <code>%v</code> — %v", // TODO(FA)
		"it-IT": "<b>Bill</b>: <code>%v</code> — %v", // TODO(IT)
		"ru-RU": `<b>Cчёт</b>: <code>%v</code> — %v`,
	},
	MESSAGE_TEXT_MEMBERS_TITLE: {
		"de-DE": "Mitglieder",
		"en-UK": "Members",
		"es-ES": "Miembros",
		"fa-IR": "اعضا",   // TODO(FA) verify
		"it-IT": "Membri", // TODO(IT)
		"ru-RU": "Участники",
	},
	MT_TEXT_MEMBERS_COUNT: {
		"en-UK": "<b>Members</b>: %d",
		"es-ES": "<b>Miembros</b>: %d",
		"fa-IR": "اعضا %d",           // TODO(FA) verify - missin <b></b>
		"it-IT": "<b>Membri</b>: %d", // TODO(IT)
		"ru-RU": "<b>Участников</b>: %d",
	},
	ALERT_TEXT_NOTHING_CHANGED: {
		"de-DE": "Nichts geändert",
		"en-UK": "Nothing changed",
		"es-ES": "Nada ha cambiado",     // TODO(ES)
		"fa-IR": "چیزی تغییر نکرده است", // TODO(FA) verify
		"it-IT": "Niente è cambiato",    // TODO(IT) verify
		"ru-RU": "Ничего не изменилось",
	},
	ALERT_TEXT_YOU_ARE_ALREADY_MEMBER_OF_THE_GROUP: {
		"de-DE": "Du bist schon Mitglied beim Teilen dieser Rechnung.",
		"en-UK": "You are already a member of this bill splitting group.",
		"es-ES": "Ya es miembro de este grupo de división de facturas.",  // TODO(ES)
		"fa-IR": "شما قبلا عضو این گروه تقسیم لایحه هستید.",              // TODO(FA)
		"it-IT": "Sei già membro di questo gruppo di divisione fatture.", // TODO(IT)
		"ru-RU": "Вы уже участник этой группы по совместной оплате счетов.",
	},
	MESSAGE_TEXT_YOUR_BILL_SPLITTING_GROUPS: {
		"de-DE": "Gruppen, mit denen du Rechnungen teilst",
		"en-UK": "Your bills splitting groups",
		"es-ES": "Ya es miembro de este grupo de división de facturas.",  // TODO(ES) verify
		"fa-IR": "شما قبلا عضو این گروه تقسیم لایحه هستید.",              // TODO(FA) verify
		"it-IT": "Sei già membro di questo gruppo di divisione fatture.", // TODO(IT) verify
		"ru-RU": "Ваши группы совметсной оплаты",
	},
	MESSAGE_TEXT_USE_ARROWS_TO_SELECT_GROUP: {
		"en-UK": "Use ⬅️ & ➡️ to select group",
		"es-ES": "Usa ⬅️ y ➡️ para seleccionar el grupo",    // TODO(ES) verify
		"fa-IR": "برای انتخاب گروه از ⬅️ & ❢️ استفاده کنید", // TODO(FA) verify
		"it-IT": "Usare ⬅️ & ➡️ per selezionare il gruppo",  // TODO(IT) verify
		"ru-RU": "Используйте ⬅️ и ➡️ чтобы выбрать группу.",
	},
	MESSAGE_TEXT_NO_GROUPS: {
		"de-DE": "Du gehörst zu keiner Gruppe, die sich Rechnungen teilt.",
		"en-UK": "You are not a participant of any bill splitting group.",
		"es-ES": "Usted no es participante de ningún grupo de división de facturas.", // TODO(ES) verify
		"fa-IR": "شما شرکت کننده در هر گروه تقسیم لایحه نیستید.",                     // TODO(FA) verify
		"it-IT": "Non sei un partecipante a qualsiasi gruppo di divisione fatture.",  // TODO(IT) verify
		"ru-RU": "Вы не состоите в группах совместной оплаты.",
	},
	MESSAGE_TEXT_USER_JOINED_GROUP: {
		"de-DE": `Hi %v, du bist der Gruppe, die sich Rechnungen teilt, beigetreten.`,
		"en-UK": `Hi %v, you joined this bill splitting group.`,
		"fa-IR": "سلام %v، شما به گروه تقسیم این لایحه پیوستید",              // TODO(FA) verify
		"it-IT": "Hi %v, sei entrato in questo gruppo di divisione fatture.", // TODO(IT) verify
		"ru-RU": `Привет %v, вы присоеденились к этой группе по совместной оплате счетов.
		`,
	},
	MESSAGE_TEXT_MEMBERS_CARD_TITLE: {
		"de-DE": "<b>Wer sich die Rechnung teilt</b> (%d)",
		"en-UK": "<b>Bills splitting members</b> (%d)",
		"fa-IR": "(%d) <b>نقض تقسیم اعضا</b>",                     // TODO(FA) verify
		"it-IT": "<b>Membri di divisione delle bollette</b> (%d)", // TODO(IT) verify
		"ru-RU": "<b>Участники совместных оплат</b> (%d)",
	},
	MESSAGE_TEXT_SPLIT_MODE: {
		"de-DE": "<b>Teilen</b>: %v",
		"en-UK": "<b>Split</b>: %v",
		"es-ES": "<b>División</b>: %v", // TODO(ES) verify
		"fa-IR": "<b>شکاف</b>: %v",     // TODO(FA) verify
		"it-IT": "<b>Diviso</b>: %v",   // TODO(IT) verify
		"ru-RU": "<b>Делить</b>: %v",
	},
	MESSAGE_TEXT_ASK_HOW_TO_SPLIT_IN_GROP: {
		"de-DE": "In welchem Verhältnis teilt ihr in dieser Gruppe eure Rechnungen?",
		"en-UK": "In what proportion do you split bills in this group?",
		"es-ES": "¿En qué proporción divide las facturas en este grupo?",     // TODO(ES) verify
		"fa-IR": "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟", // TODO(FA) verify
		"it-IT": "In quale percentuale dividi le fatture in questo gruppo?",  // TODO(IT) verify
		"ru-RU": "В какой пропорции вы делите счета в этой группе?",
	},
	MESSAGE_TEXT_MEMBERS_CARD_FOOTER: {
		"de-DE": "Klick <code>Join</code>, um auch Rechnungen zu teilen.",
		"en-UK": "Click <code>Join</code> to participate in bills splitting.",
		"es-ES": "¿En qué proporción divide las facturas en este grupo?",     // TODO(ES) verify
		"fa-IR": "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟", // TODO(FA) verify
		"it-IT": "In quale percentuale dividi le fatture in questo gruppo?",  // TODO(IT) verify
		"ru-RU": "Жмите <code>Присоедениться</code> чтобы учавствовать.",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBER_TITLE: {
		"en-UK": "{{.N}}. {{.MemberName}}",
		"es-ES": "{{.N}}. {{.MemberName}}",
		"fa-IR": "{{.N}}. {{.MemberName}}", // TODO(FA) verify
		"it-IT": "{{.N}}. {{.MemberName}}",
		"ru-RU": "{{.N}}. {{.MemberName}}",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW: {
		"de-DE": "<i>{{.Percent}}%</i>",
		"en-UK": "<i>{{.Percent}}%</i>",
		"es-ES": "<i>{{.Percent}}%</i>",
		"fa-IR": "<i>{{.Percent}}%</i>", // TODO(FA) verify
		"it-IT": "<i>{{.Percent}}%</i>",
		"ru-RU": "<i>{{.Percent}}%</i>",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_OWES: {
		"de-DE": "\n   <i>schuldet {{.Owes}}</i>",
		"en-UK": "\n   <i>owes {{.Owes}}</i>",
		"es-ES": "\n   <i>debo {{.Owes}}</i>",
		"ru-RU": "\n   <i>должен {{.Owes}}</i>",
		"fa-IR": "\n   <i>بدهکار است {{.Owes}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PAID: {
		"de-DE": "\n   <i>bezahlte {{.Paid}}</i>",
		"en-UK": "\n   <i>paid {{.Paid}}</i>",
		"es-ES": "\n   <i>he pagado {{.Paid}}</i>",
		"ru-RU": "\n   <i>заплатил {{.Paid}}</i>",
		"fa-IR": "\n   <i>پرداخت شده {{.Paid}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PART_PAID: {
		"de-DE": "\n<i>bezahlte {{.Paid}}, schuldet noch {{.Owes}}</i>",
		"en-UK": "\n<i>paid {{.Paid}}, owes {{.Owes}}</i>",
		"es-ES": "\n<i>he pagado {{.Paid}}, debo {{.Owes}}</i>",
		"ru-RU": "\n<i>заплатил {{.Paid}}, должен {{.Owes}}</i>",
		"fa-IR": "\n<i>پرداخت شده {{.Paid}}, بدهکار است {{.Owes}}</i>", // TODO(FA) verify
	},
	MESSAGE_TEXT_BILL_ASK_WHO_PAID: {
		"de-DE": "Bitte wähle, wer die Rechnung gezahlt hat:",
		"en-UK": "Please choose who paid for the bill:",
		"es-ES": "Por favor, elige quien ha pagado la cuenta:",
		"fa-IR": "لطفا انتخاب کنید که چه کسانی برای این لایحه پرداخت کرده اند:", // TODO(FA) verify
		"it-IT": "Scegliere chi ha pagato la fattura:",                          // TODO(IT) verify
		"ru-RU": "Пожалуйста выберите кто заплатил по счёту:",
	},
	MESSAGE_TEXT_STATUS: {
		"de-DE": "Status: %v",
		"en-UK": "Status: %v",
		"es-ES": "Estado: %v",
		"fa-IR": "وضعیت: %v", // TODO(FA) verify
		"it-IT": "Stato: %v", // TODO(IT) verify
		"ru-RU": "Статус: %v",
	},
	BUTTON_TEXT_ADD_MEMBER: {
		"de-DE": "Partei hinzufügen",
		"en-UK": "Add participant",
		"es-ES": "Añadir participante",
		"fa-IR": "افزودن مشارکت کننده",   // TODO(FA) verify
		"it-IT": "Aggiungi partecipante", // TODO(IT) verify
		"ru-RU": "Добавить участника",
	},
	BUTTON_TEXT_FINALIZE_BILL: {
		"de-DE": "🔓 Rechnung abschließen",
		"en-UK": "🔓 Lock the bill",
		"es-ES": "🔓 Cerrar la cuenta",
		"fa-IR": "🔓 لایحه را قفل کنید", // TODO(FA) verify
		"it-IT": "🔓 Bloccare il conto", // TODO(IT) verify
		"ru-RU": "🔓 Закрыть счёт",
	},
	BUTTON_TEXT_EDIT_BILL: {
		"de-DE": "✏️ Bearbeiten",
		"en-UK": "✏️ Edit",
		"es-ES": "✏️ Editar",
		"fa-IR": "✏️ ویرایش",   // TODO(FA) verify
		"it-IT": "✏️ Modifica", // TODO(IT) verify
		"ru-RU": "✏️ Изменить",
	},
	BUTTON_TEXT_SPLIT_MODE: {
		"de-DE": "➗ Teilen: %v",
		"en-UK": "➗ Split: %v",
		"es-ES": "➗ Dividir: %v",
		"fa-IR": "➗ تقسیم: %v", // TODO(FA) verify
		"it-IT": "➗ Split: %v", // TODO(IT) verify
		"ru-RU": "➗ Делить: %v",
	},
	MESSAGE_TEXT_SPLIT_LABEL_WITH_VALUE: {
		"de-DE": "Teilen: %v",
		"en-UK": "Split: %v",
		"es-ES": "Dividir: %v",
		"fa-IR": "تقسیم: %v", // TODO(FA) verify
		"it-IT": "Split: %v", // TODO(IT) verify
		"ru-RU": "Делить: %v",
	},
	STATUS_DRAFT: {
		"de-DE": "Entwurf",
		"en-UK": "draft",
		"es-ES": "borrador",
		"fa-IR": "پیش نویس", // TODO(FA) verify
		"it-IT": "bozza",    // TODO(IT) verify
		"ru-RU": "черновик",
	},
	SPLIT_MODE_EQUALLY: {
		"de-DE": "Gleichverteilt",
		"en-UK": "Equally",
		"es-ES": "A partes iguales",
		"ru-RU": "Поровну",
		"fa-IR": "به همان اندازه", // TODO(FA) verify
		"it-IT": "Ugualmente",     // TODO(IT) verify
	},
	SPLIT_MODE_PERCENTAGE: {
		"de-DE": "Prozentual",
		"en-UK": "Percentage",
		"es-ES": "Porcentaje",
		"fa-IR": "درصد",        // TODO(FA)
		"it-IT": "Percentuale", // TODO(IT) verify
		"ru-RU": "В процентах",
	},
	SPLIT_MODE_EXACT_AMOUNT: {
		"de-DE": "Exakte Summen",
		"en-UK": "Exact amounts",
		"es-ES": "Importes exactos",
		"fa-IR": "مقادیر دقیق",     // TODO(FA) verify
		"it-IT": "Quantità esatte", // TODO(IT) verify
		"ru-RU": "Точные суммы",
	},
	SPLIT_MODE_SHARES: {
		"de-DE": "Teilen",
		"en-UK": "Shares",
		"es-ES": "Proporciones",
		"fa-IR": "سهام",   // TODO(FA) verify
		"it-IT": "Azioni", // TODO(IT) verify
		"ru-RU": "В долях",
	},
	BUTTON_TEXT_JOIN: {
		"de-DE": "➕ Beitreten",
		"en-UK": "➕ Join",
		"es-ES": "➕ Adherirse",
		"fa-IR": "➕ عضویت", // TODO(FA) verify
		"it-IT": "➕ Join",  // TODO(IT) verify
		"ru-RU": "➕ Присоедениться",
	},
	BUTTON_TEXT_LEAVE: {
		"de-DE": "Verlassen",
		"en-UK": "Leave",
		"es-ES": "Salir",    // TODO(ES) verify
		"fa-IR": "ترک کردن", // TODO(FA) verify
		"it-IT": "Partire",  // TODO(IT) verify
		"ru-RU": "Покинуть",
	},
	BUTTON_TEXT_DUE: {
		"de-DE": "📅 Fällig: %v",
		"en-UK": "📅 Due: %v",
		"es-ES": "📅 Hasta: %v",
		"fa-IR": "📅 مورد: %v",   // TODO(FA) verify
		"it-IT": "📅 Dovuto: %v", // TODO(IT) verify
		"ru-RU": "📅 До: %v",
	},
	NOT_SET: {
		"de-DE": "nicht gesetzt",
		"en-UK": "not set",
		"es-ES": "no establecido",
		"fa-IR": "تنظیم نشده",    // TODO(FA) verify
		"it-IT": "non impostato", // TODO(IT) verify
		"ru-RU": "не задано",
	},
	BUTTON_TEXT_MANAGE_MEMBERS: {
		"de-DE": "👫 Parteien",
		"en-UK": "👫 Participants",
		"es-ES": "👫 Participantes",
		"fa-IR": "👫 شرکت کنندگان", // TODO(FA) verify
		"it-IT": "👫 Partecipanti", // TODO(IT) verify
		"ru-RU": "👫 Участники",
	},
	BUTTON_TEXT_CHANGE_BILL_PAYER: {
		"de-DE": "🔀 Bezahler ändern",
		"en-UK": "🔀 Change payer",
		"es-ES": "🔀 Cambiar el pagador",
		"fa-IR": "🔀 تغییر پرداخت کننده", // TODO(FA) verify
		"it-IT": "🔀 Cambia il pagatore", // TODO(IT) verify
		"ru-RU": "🔀 Сменить плательщика",
	},
	COMMAND_TEXT_I_PAID: {
		"de-DE": "Ich habe bezahlt",
		"en-UK": "I paid",
		"es-ES": "he pagado",
		"fa-IR": "پرداخت کردم", // TODO(FA) verify
		"it-IT": "Ho pagato",   // TODO(IT) verify
		"ru-RU": "Я заплатил",
	},
	COMMAND_TEXT_I_OWE: {
		"de-DE": "Ich schulde",
		"en-UK": "I owe",
		"es-ES": "Yo debo",
		"fa-IR": "من بدهکارم",     // TODO(FA) verify
		"it-IT": "Sono in debito", // TODO(IT) verify
		"ru-RU": "Я должен",
	},
	COMMAND_TEXT_OWED_TO_ME: {
		"de-DE": "schuldet mir",
		"en-UK": "Owed to me",
		"es-ES": "Me deben",
		"fa-IR": "به من تعلق دارد", // TODO(FA) verify
		"it-IT": "È dovuto a me",   // TODO(IT) verify
		"ru-RU": "Должны мне",
	},
	MESSAGE_TEXT_BILL_HEADER: {
		"de-DE": "Rechnung: %v",
		"en-UK": "Bill: %v",
		"es-ES": "Cuenta: %v",
		"fa-IR": "بیل :%v",  // TODO(FA) verify
		"it-IT": "Bill: %v", // TODO(IT) verify
		"ru-RU": "Cчёт: %v",
	},
	MESSAGE_TEXT_NEW_DEBT_HEADER: {
		"de-DE": "Rechnung: %v",
		"en-UK": "Bill: %v",
		"es-ES": "Cuenta: %v",
		"fa-IR": "بیل: %v",  // TODO(FA) verify
		"it-IT": "Bill: %v", // TODO(IT) verify
		"ru-RU": "Cчёт: %v",
	},
	MESSAGE_TEXT_GROUPS_ONLY_COMMAND: {
		"de-DE": "",
		"en-UK": "This command is available in group chats only for now.",
		"es-ES": "",
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": "Эта команда пока что доступна только в групповых чатах",
	},
	MESSAGE_TEXT_ALREADY_HAS_CONTACT_WITH_SUCH_NAME: {
		"de-DE": "", // TODO(DE)
		"en-UK": "You already have contact with name: %v",
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": "У вас уже есть контакт с таким именем: %v",
	},
	MESSAGE_TEXT_ALREADY_BILL_MEMBER: {
		"de-DE": "%v, du teilst diese Rechnung bereits.",
		"en-UK": "%v, you are sharing this bill already.",
		"es-ES": "%v, estás compartiendo esta cuenta ya.",
		"fa-IR": "%v، شما قبلا این لایحه را به اشتراک می گذارید",      // TODO(FA) verify
		"it-IT": "%v, stai già condividendo questo disegno di legge.", // TODO(IT) verify
		"ru-RU": "%v, вы уже входите в состав участников.",
	},
	MESSAGE_TEXT_USER_JOINED_BILL: {
		"de-DE": "%v ist dem Teilen der Rechnung beigetreten.",
		"en-UK": "%v joined to bill sharing.",
		"es-ES": "%v pagar conjuntamente.",
		"fa-IR": "%v به اشتراک گذاری لایحه پیوست.",        // TODO(FA) verify
		"it-IT": "%v unito alla condivisione di fatture.", // TODO(IT) verify
		"ru-RU": "%v присоеденился к совместной оплате.",
	},
	BUTTON_TEXT_I_PAID_FOR_THE_BILL: {
		"de-DE": "Die Rechnung wurde von mir bezahlt.",
		"en-UK": "The bill was paid by me.",
		"es-ES": "La factura fue pagada por mí.",  // TODO(ES) verify
		"fa-IR": "این لایحه توسط من پرداخت شد",    // TODO(FA) verify
		"it-IT": "Il conto è stato pagato da me.", // TODO(IT) verify
		"ru-RU": "Этот счёт оплатил(а) я.",
	},
	BUTTON_TEXT_I_OWE_FOR_THE_BILL: {
		"de-DE": "Ich muss noch was dabeigeben",
		"en-UK": "I owe for this bill",
		"es-ES": "Debo esta factura",                // TODO(ES) verify
		"fa-IR": "من برای این لایحه بدهکار هستم",    // TODO(FA) verify
		"it-IT": "Devo per questo disegno di legge", // TODO(IT) verify
		"ru-RU": "Я должен по этому счёту",
	},
	BUTTON_TEXT_I_DO_NOT_SHARE_THIS_BILL: {
		"de-DE": "Ich habe damit nichts zutun",
		"en-UK": "I don't share this bill",
		"es-ES": "No comparto esta cuenta",               // TODO(ES) verify
		"fa-IR": "من این لایحه را به اشتراک نمی گذارم",   // TODO(FA) verify
		"it-IT": "Non condivido questo disegno di legge", // TODO(IT) verify
		"ru-RU": "Я не учавствую в этой покупке",
	},
	MESSAGE_TEXT_YOU_JOINED_BILL: {
		"de-DE": "Du bist dem Teilen der Rechnung beigetreten.",
		"en-UK": "You've joined to bill sharing.",
		"es-ES": "Te has agregado para pagar conjuntamente .",
		"fa-IR": "شما به اشتراک گذاشتن لایحه پیوستید",          // TODO(FA) verify
		"it-IT": "Sei entrato a far parte della fatturazione.", // TODO(IT) verify
		"ru-RU": "Вы присоеденились к совместной оплате.",
	},
	ARTICLE_TITLE_SPLIT_BILL: {
		"de-DE": "eine Rechnung teilen",
		"en-UK": "Split bill/purchase",
		"es-ES": "Compartir la cuenta/compra",
		"fa-IR": "لایحه / خرید تقسیم شده",    // TODO(FA) verify
		"it-IT": "Bolletta Split / acquisto", // TODO(IT) verify
		"ru-RU": "Разделить счёт/покупку",
	},
	ARTICLE_SUBTITLE_SPLIT_BILL: {
		"de-DE": "Wert: %v\nTeile deine Kosten mit Freunden und verfolge deren Rückzahlungen.",
		"en-UK": "Amount: %v\nShares expenses with friends & track paybacks",
		"es-ES": "Importe: %v\nCompartir los gastos entre amigos y seguir las devoluciones",  // TODO(ES): Have to be shorter
		"fa-IR": "مقدار: %v" + "\n" + "هزینه ها را با دوستان و بازپرداخت پیگیری می کند",      // TODO(FA) verify
		"it-IT": "Importo: %v\nDisponi i costi con gli amici e le retribuzioni delle tracce", // TODO(IT) verify
		"ru-RU": "Сумма: %v\nПоделить траты между друзьями и отследить возвраты",
	},

	ARTICLE_NEW_DEBT_TITLE: {
		"de-DE": "Neuer Schuldschein",
		"en-UK": "New debt",
		"es-ES": "Nueva deuda",
		"fa-IR": "بدهی جدید",    // TODO(FA): Verify
		"it-IT": "Nuovo debito", // TODO(IT): Verify
		"ru-RU": "Новый долг",
	},
	ARTICLE_NEW_DEBT_SUBTITLE: {
		"de-DE": "Wert: %v\nZur Fälligkeit wird eine Benachrichtigung geschickt, falls so eingestellt",
		"en-UK": "Amount: %v\nSends notifications on due date if set",
		"es-ES": "Importe: %v\nEnviar las notificaciones el día de vencimiento",
		"fa-IR": "مقدار: %v" + "\n" + "اگر تنظیم شود، اطلاعیه ها را در تاریخ تعیین شده ارسال می کند", // TODO(FA):  verify
		"it-IT": "Importo: %v\nSend le notifiche alla data di scadenza se impostato",                 // TODO(IT) verify
		"ru-RU": "Сумма: %v\nЗапись долга и рассылка оповещений в день возврата.",
	},
	SPLITUS_PLEASE_JOIN_IF_NOT_ON_THE_LIST: {
		"de-DE": `Bitte tritt zuerst bei, falls dein Name nicht auf der Liste ist.`,
		"en-UK": `Please join if your name is not on the list above.`,
		"fa-IR": `اگر نام شما در لیست بالا نیست، لطفا پیوست شوید.`,                   // TODO(FA) verify
		"it-IT": `Si prega di unirti se il tuo nome non è nell'elenco di cui sopra.`, // TODO(IT) verify
		"ru-RU": `Пожалуйста присоеденяйтесь если ваше не в списке.`,                 // TODO(RU)
	},
	SPLITUS_TEXT_HI_IN_GROUP: {
		"de-DE": `Ich bin <b>Splitus</b>. Danke fürs hinzufügen!`,
		"en-UK": `I'm <b>Splitus</b>. Thanks for adding me!`,
		"es-ES": `Soy <b>Splitus</b>. ¡Gracias por agregarme!`,      // TODO(ES) verify
		"fa-IR": `من <b>Splitus</b> با تشکر برای اضافه کردن من!`,    // TODO(FA) verify
		"it-IT": `Sono <b>Splitus</b>. Grazie per averci aggiunto!`, // TODO(IT) verify
		"ru-RU": `Меня зовут <b>Сплитус</b>. Спасибо что добавили!`,
	},
	COLLECTUS_TEXT_HI_IN_GROUP: {
		"en-UK": `I'm <b>Collectus.</b> Thanks for adding me!`,
		"es-ES": `Soy <b>Collectus.</b> ¡Gracias por agregarme!`,      // TODO(ES) verify
		"fa-IR": `من <b>Collectus</b> با تشکر برای اضافه کردن من!`,    // TODO(FA) verify
		"it-IT": `Sono <b>Collectus.</b> Grazie per averci aggiunto!`, // TODO(IT) verify
		"ru-RU": `Меня зовут <b>Коллектус.</b> Спасибо что добавили!`,
	},
	MT_GROUP_LABEL: {
		"en-UK": `<b>Group</b>: %v`,
		"es-ES": `<b>Group</b>: %v`, // TODO(ES)
		"fa-IR": `<b>Group</b>: %v`, // TODO(FA)
		"it-IT": `<b>Group</b>: %v`, // TODO(IT)
		"ru-RU": `<b>Группа</b>: %v`,
	},
	MT_SPONSORS_HEADER: {
		"en-UK": `<b>Sponsors</b>:`,
		"es-ES": `<b>Patrocinadores</b>:`, // TODO(ES)
		"fa-IR": `<b>حامیان</b>:`,         // TODO(FA) verify
		"it-IT": `<b>Sponsors</b>:`,       // TODO(IT)
		"ru-RU": `<b>Спонсоры</b>:`,
	},
	MT_DEBTORS_HEADER: {
		"en-UK": `<b>Debtors</b>:`,
		"es-ES": `<b>Deudores</b>:`, // TODO(ES) verify
		"fa-IR": `<b>Debtors</b>:`,  // TODO(FA) verify
		"it-IT": `<b>بدهکار</b>:`,   // TODO(IT) verify
		"ru-RU": `<b>Должники</b>:`,
	},
	BT_DEFAULT_CURRENCY: {
		"en-UK": `Currency: %v`,
		"es-ES": `Moneda: %v`,   // TODO(ES) verify
		"fa-IR": `واحد پول: %v`, // TODO(FA) verify
		"it-IT": `Moneta: %v`,   // TODO(IT) verify
		"ru-RU": `Валюта: %v`,
	},
	MESSAGE_TEXT_ASK_LANG: {
		"de-DE": `Welche Sprache wird hier gesprochen?`,
		"en-UK": `What language should I use in this group?`,
		"es-ES": `¿Qué idioma debería usar en este grupo?`,      // TODO(ES) verify
		"fa-IR": `کدام زبان باید در این گروه استفاده کنم؟`,      // TODO(FA) verify
		"it-IT": `Che lingua devo utilizzare in questo gruppo?`, // TODO(IT) verify
		"ru-RU": `Какой язык я должен использовать в этой группе?`,
	},
	MESSAGE_TEXT_HI_IN_GROUP_LANG_SET: {
		"en-UK": `Great, I'll be using English.`,
		"es-ES": `Genial, usaré español.`, // TODO(ES) verify
		"de-DE": `Kein Problem, dann schreibe ich auf Deutsch.`,
		"fa-IR": `عالی، من از فارسی استفاده خواهم کرد.`, // TODO(FA) verify
		"it-IT": `Ottimo, userò l'italiano.`,            // TODO(IT) verify
		"ru-RU": `Отлично, я буду использовать русский`,
	},
	SPLITUS_TEXT_ABOUT_ME_AND_CO: {
		"de-DE": `Ich kann helfen, <b>Rechnungen zu teilen</b>. Mein Freund @DebtsTrackerBot passt darauf auf, dass alle Schulden zurückgezahlt werden.`,
		"en-UK": `I help to <b>split bills</b>. My friend @DebtsTrackerBot is tracking paybacks & debts.`,
		"es-ES": `Ayudo a <b>dividir billetes</b>. Mi amigo @DebtsTrackerBot está rastreando pagos y deudas.`,                 // TODO(ES) verify
		"fa-IR": `من به <b>تقسیم صورتحساب </b> کمک میکنم دوست منDebtsTrackerBot ردیابی بازپرداخت و بدهی است.`,                 // TODO(FA) verify
		"it-IT": `Aiuto a <b>dividere le bollette</b>. Il mio amico @DebtsTrackerBot sta monitorando i pagamenti e i debiti.`, // TODO(IT) verify
		"ru-RU": `Я помогаю делить счета. Мой друг @DebtsTrackerRuBot отслеживает платежи и долги.`,
	},
	COLLECTUS_TEXT_ABOUT_ME_AND_CO: {
		"de-DE": `Ich kann dabei helfen <b>Geld zu sammeln</b>. Zum Beispiel für ein Geburtstagsgeschenk. 🎉

Mein Freund @DebtsTrackerBot überwacht Schulden und deren Rückzahlungen.

Wenn du Geld für irgendwas sammeln willst, kann @SplitusBot dir dabei helfen.`,

		"en-UK": `I help to <b>collect money</b> for a good cause. For example for a birthday present. 🎉

My buddy @DebtsTrackerBot is tracking debts & paybacks.

And if you do collective purchases and want to split bills @SplitusBot is here to help.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Я помогаю <b>собирать деньги</b> на что нибудь. Например для подарка на день рождение. 🎉

Мой друг @DebtsTrackerRuBot отслеживает долги и платежи.

А если хотите вести учёт совместных покупок и удобно делить счета вам поможет @SplitusBot.`,
	},
	SPLITUS_TEXT_HI: {
		"de-DE": `Ich bin <b>Splitus</b>.`,
		"en-UK": `I'm <b>Splitus</b>.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Меня зовут <b>Сплитус</b>.`,
	},
	COLLECTUS_TEXT_HI: {
		"de-DE": `Ich bin <b>Collectus</b>.`,
		"en-UK": `I'm <b>Collectus</b>.`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Меня зовут <b>Коллектус</b>.`,
	},
	SPLITUS_TG_COMMANDS: {
		"de-DE": ``,
		"en-UK": `<b>Bot commands:</b>
	/groups - List of groups
	/bills - List of outstanding bills
	/menu - Main menu
	/settings - Settings
	/help - Learn how to use bot, report issues, ask questions`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `<b>Команды для бота:</b>
	/groups - Список групп
	/bills - Список незакрытых платежей
	/menu - Главное меню
	/settings - Настройки
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	COLLECTUS_TG_COMMANDS: {
		"en-UK": `<b>Bot commands:</b>

	/groups - List of groups
	/fundraisings - List of active fundraisings
	/help - Learn how to use bot, report issues, ask questions`,
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT)
		"ru-RU": `<b>Команды для бота:</b>
	/groups - Список групп
	/fundraisings - Список активных сборов
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	MESSAGE_TEXT_SEND_HELP_COMMAND_FOR_HELP: { // This is the same for all languages.
		"de-DE": `Sende /help für Hilfe über den Umgang mit diesen Bot.`,
		"en-UK": `Send /help for details on how to use this bot.`,
		"es-ES": ``,
		"fa-IR": ``,
		"it-IT": ``,
		"ru-RU": `Отправьте /help для справки по использованию бота.`,
	},
	MESSAGE_TEXT_HI: { // This is the same for all languages.
		"de-DE": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"en-UK": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"es-ES": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"fa-IR": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"it-IT": `¡Hola! Hi! Привет! سلام! Hallo!`,
		"ru-RU": `¡Hola! Hi! Привет! سلام! Hallo!`,
	},
	MESSAGE_TEXT_HI_USERNAME: { // This is the same for all languages.
		"de-DE": `Hi %v!`,
		"en-UK": `Hi %v!`,
		"es-ES": `¡Hola %v!`,
		"fa-IR": ``,
		"it-IT": ``, // TODO(IT)
		"ru-RU": `Привет %v!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		"de-DE": `Du kannst zurück zum Haupt /menu`,
		"en-UK": `You can go back to main /menu`,
		"es-ES": `Puedes volver al inicio /menú`,
		"fa-IR": `شما میتوانید به /منو ی اصلی مراجعه کنید.`,
		"it-IT": `Puoi tornare al menu' principale tramite /menu`,
		"ru-RU": `Можно вернуться назад в главное /меню`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		"de-DE": `Bevorzugte Sprache: %s`,
		"en-UK": `Preferred bot language: %s`,
		"es-ES": `Idioma favorito del bot: %s`,
		"fa-IR": `زبان برنامه: %s`,
		"it-IT": `Lingua del bot preferita: %s`,
		"ru-RU": `Выбранный язык бота: %s`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		"de-DE": `Welche ist deine bevorzugte Sprache?`,
		"en-UK": `What is your preferred language?`,
		"es-ES": `¿cuál es tu idioma preferida?`,
		"fa-IR": `شما چه زبانی را ترجیح می دهید؟`,
		"it-IT": `Qual'e' la tua lingua madre?`,
		"ru-RU": `На каком языке вы хотели бы общаться?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		"de-DE": "Auf welcher Sprache würdest du dich gerne mit mir unterhalten?",
		"en-UK": "Which language you would like to talk to me?",
		"es-ES": "¿En cuál idioma preferirías hablar conmigo?",
		"fa-IR": "دوست دارید به چه زبانی با من صحبت کنید؟",
		"it-IT": "In quale lingua preferisci parlarmi?",
		"ru-RU": "На каком языке вы хотели бы общаться со мной?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		"de-DE": "Du hast die Sprache geändert auf %v",
		"en-UK": "You've switched language to %v",
		"es-ES": "Has cambiado el idioma a %v",
		"fa-IR": "شما زبان را به %v تغییر دادید. ",
		"it-IT": "Hai cambiato lingua in %v",
		"ru-RU": "Вы поменяли язык на %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		"de-DE": "Was als nächstes?",
		"en-UK": "What's next?",
		"es-ES": "¿Algo más?",
		"fa-IR": "اقدام بعدی چیست؟",
		"it-IT": "Cosa posso fare ora?",
		"ru-RU": "Что будем делать дальше?",
	},
	MESSAGE_TEXT_REFERRERS_TITLE: {
		"en-UK": "Our friends:",
		"ru-RU": "Наши друзья:",
	},
	COMMAND_TEXT_ADD_MY_TG_CHANNEL: {
		"en-UK": "Add my channel",
		"ru-RU": "Добавить мой канал",
	},
	MESSAGE_TEXT_DEBTUS_COMMANDS: {
		"en-UK": `<b>Commands for the bot</b>
🏠 /menu - show main menu
🔙 /return - return previously recorded debt
📥 /got - record money you received from others
📤 /gave - record money you gave to others
📚 /history - latest transactions
🏁 /balance - display current balance
⚙ /settings - adjust your preferences`,

		"ru-RU": `<b>Команды бота</b>
🏠 /menu - показать основное меню
🔙 /return - записать возврат долга
📥 /got - записать о взятии в долг
📤 /gave - записать о том что дали взаймы
📚 /history - последние транзакции
🏁 /balance - показать текущий баланс
⚙ /settings - настройки`,
	},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		"de-DE": `
Wenn du dir was von jemanden geliehen hast, wähle /anleihen.
Wenn du was an jemanden verliehen hast, wähle /verleihen.

Oder benutzt das Menü unten.`,
		"en-UK": `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.`,

		"es-ES": `
		Si alguien te ha prestado usa el comando  /recibido.
		Si has prestado a alguien usa el comando /dado.

O usa el menú de abajo en la pantalla.`,

		"fa-IR": `
اگر از کسی قرض گرفته اید برای ثبت آن از /قرض_گرفتن استفاده کنید.
اگر به کسی قرض داده اید برای ثبت آن از /قرض_دادن استفاده کنید.

یا از منوی پایین استفاده نمایید.`,

		"it-IT": `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.`,

		"ru-RU": `
	Если вы дали в долг воспользуйтесь командой /дал.
	Если вы одолжили что-то - командой /взял.

	Или воспользуйтесь меню внизу экрана.`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		"de-DE": "Verlauf",
		"en-UK": "History",
		"es-ES": "Cronología",
		"fa-IR": "سوابق",
		"it-IT": "Cronologia",
		"ru-RU": "История",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		"de-DE": "Du hast noch nichts gespeichert.",
		"en-UK": "You don't have any records yet.",
		"es-ES": "Aún no tienes ninguna notificación.",
		"fa-IR": "شما هنوز هیچ ثبت سابقه ای ندارید.",
		"it-IT": "Non hai nulla memorizzato.",
		"ru-RU": "У вас пока нет ни одной записи.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {

		"de-DE": `<b>%v</b> <i>(bis %d):</i>
─────────────
%v`,
		"en-UK": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		"en-US": `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		"es-ES": `<b>%v</b> <i>(últimos %d):</i>
─────────────
%v`,

		"fa-IR": `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,
		"fr-FR": "",
		"id-ID": "",
		"it-IT": `<b>%v</b> <i>(ultimi %d):</i>

─────────────
%v`,
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		"de-DE": "Keine unbeglichenen Schulden.",
		"en-UK": "You have no records on current debts.",
		"en-US": "You have no records on current debts.",
		"es-ES": "No hay ninguna notificación de deudas actuales.",
		"fa-IR": "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Non hai nulla memorizzato nel debito corrente.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Нет записей о текущих долгах.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		"de-DE": "Insgesamt",
		"en-UK": "Total",
		"en-US": "Total",
		"es-ES": "Total",
		"fa-IR": "مجموع",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Totale",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Всего",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	BT_OTHER_CURRENCY: {
		"de-DE": "",
		"en-UK": "Another currency",
		"en-US": "Another currency",
		"es-ES": "Otra moneda", // TODO(es) verify
		"fa-IR": "ارز دیگر",    // TODO(fa) verify
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Un'altra valuta", // TODO(it) verify
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Другая валюта",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		"de-DE": "OK, von nun an ist '%v' deine Hauptwährung.",
		"en-UK": "OK, from now on I will use '%v' as a primary currency.",
		"en-US": "OK, from now on I will use '%v' as a primary currency.",
		"es-ES": "OK, ahora voy a usar '%v' como moneda principal. ",
		"fa-IR": "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "OK, da ora in poi usero' '%v' come moneta principale.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "OK, теперь я буду использовать '%v' как основную валюту.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		"de-DE": "%v - schuldet dir %v",
		"en-UK": "%v - owes you %v",
		"en-US": "%v - owes to you %v",
		"es-ES": "%v - te debe %v",
		"fa-IR": "%v - %v به شما بدهکار است ",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "%v - ti deve %v.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "%v - долг вам %v",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		"de-DE": "Schuldet dir %v",
		"en-UK": "Owes to you %v",
		"en-US": "Owes to you %v",
		"es-ES": "Te debe %v",
		"fa-IR": "%v به شما بدهکار است ",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Hai un credito di %v",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Вам должны %v",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		"de-DE": "Hurra, du bist jetzt quitt mit <b>%v</b>.",
		"en-UK": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"en-US": "Congratulations! You don't owe anything more to <b>%v</b>.",
		"es-ES": "Bravo! Has saldado tu deuda con <b>%v</b>.",
		"fa-IR": "تبریک! شما دیگر چیزی به <b>%v</b> بدهکار نیستید .",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		"de-DE": "Du bist jetzt mit <b>%v</b> quitt.",
		"en-UK": "<b>%v</b> does not owe anything more to you.",
		"en-US": "<b>%v</b> does not owe anything more to you.",
		"es-ES": "<b>%v</b> nadie te debe nada ya.",
		"fa-IR": "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "<b>%v</b> ha saldato il suo debito verso di te.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "У <b>%v</b> больше не осталось долгов перед вами.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		"de-DE": "Du schuldest %v",
		"en-UK": "You owe %v",
		"en-US": "You owe %v",
		"es-ES": "Tú debes %v",
		"fa-IR": "شما %v بدهکار هستید ",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Hai un debito di %v",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Вы должны %v",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		"de-DE": "%v - du schuldest %v",
		"en-UK": "%v - you owe %v",
		"en-US": "%v - you owe %v",
		"es-ES": "%v - tú debes %v",
		"fa-IR": "%v - شما %v بدهکار هستید ",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "%v - tu gli/le devi %v",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "%v - вы должны %v",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		"de-DE": "Was ist deine Hauptwährung?",
		"en-UK": "What is your primary currency?",
		"en-US": "What is your primary currency?",
		"es-ES": "¿Cuál es tu moneda principal?",
		"fa-IR": "واحد پولی اولیه شما چیست؟",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Qual'e' la tua valuta principale?",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Какая валюта для вас основная?",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY_FOR_GROUP: {
		"de-DE": "",
		"en-UK": "What is a primary currency for this group?",
		"en-US": "What is a primary currency for this group?",
		"es-ES": "¿Cuál es tu moneda principal?", //TODO(ES)
		"fa-IR": "واحد پولی اولیه شما چیست؟",     //TODO(FA)
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Qual'e' la tua valuta principale?", //TODO(IT)
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Какая валюта основная для этой группы?",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		"de-DE": "Konnte den Benutzer nicht löschen: %v",
		"en-UK": "Failed to delete user: %v",
		"en-US": "Failed to delete user: %v",
		"es-ES": "Error durante la cancelación del usuario: %v",
		"fa-IR": "حذف کاربر ناموفق بود: %v",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Errore durante la cancellazione dell'utente: %v",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Не удалось удалить данные пользователя: %v",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_USER_DELETED: {
		"de-DE": "Die Benutzerdaten wurden gelöscht.",
		"en-UK": "User's data has been deleted",
		"en-US": "User's data has been deleted",
		"es-ES": "Los datos del usuario han sido eliminados",
		"fa-IR": "اطلاعات کاربر حذف شد.",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "I dati dell'utente sono stati cancellati",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Данные пользователя удалены",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		"de-DE": "Bitte wähle, wer die Schuld beglichen hat oder wem du sie zurückgezahlt hast.",
		"en-UK": "Please choose who returned the debt or to who you returned it.",
		"en-US": "Please choose who returned the debt or to who you returned it.",
		"es-ES": "Por favor, elige quien ha devuelto o a quien ha sido devuelta la deuda ",
		"fa-IR": "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Scegli con chi hai sanato un credito o un debito.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Выберете кому вы вернули долг или кто вернул его вам.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		"de-DE": "Bitte wähle, ob die Schuld vollständig oder teilweise beglichen wurde.",
		"en-UK": "Please choose a debt that has been returned fully or partially.",
		"en-US": "Please choose a debt that has been returned fully or partially.",
		"es-ES": "Por favor, elige una deuda que ha sido devuelta total o parcialmente. ",
		"fa-IR": "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "Scegli un debito che e' stato restituito completamente o parzialmente.",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "Выберите долг который был возвращён целиком или частично.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_NO_DEBTS_TO_RETURN: {
		"de-DE": "You have no records for debts that can be returned.",
		"en-UK": "",
		"en-US": "",
		"es-ES": "",
		"fa-IR": "",
		"fr-FR": "",
		"id-ID": "",
		"it-IT": "",
		"ja-JP": "",
		"ko-KO": "",
		"pl-PL": "",
		"pt-BR": "",
		"ru-RU": "У вас нет записей о догах для возврата.",
		"tr-TR": "",
		"ua-UA": "",
		"uz-UZ": "",
		"zh-CN": "",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		"de-DE": "Bitte stimme dem zu oder lehne es ab.",
		"en-UK": "Please confirm or decline this transfer.",
		"en-US": "Please confirm or decline this transfer.",
		"es-ES": "Por favor, confirma o rechaza la transacción.",
		"fa-IR": "لطفاً این تراکنش را تایید یا رد نمایید.",
		"fr-FR": "Veuillez confirmer ou refuser ce transfert.",
		"id-ID": "Silakan konfirmasi atau tolak transfer ini.",
		"it-IT": "Conferma o rifiuta questo debito/credito.",
		"ja-JP": "この転送を確認または拒否してください。",
		"ko-KO": "이 전송을 확인하거나 거부하십시오.",
		"pl-PL": "Proszę potwierdzić lub odrzucić ten transfer.",
		"pt-BR": "Por favor, confirme ou recuse esta transferência.",
		"ru-RU": "Пожалуйста подтвердите или отклоните эту транзакцию.",
		"tr-TR": "Lütfen bu transferi onaylayın veya reddedin.",
		"ua-UA": "Будь ласка, підтвердіть або відхиліть цей переказ.",
		"uz-UZ": "Iltimos, ushbu o'tkazmani tasdiqlang yoki rad eting.",
		"zh-CN": "请确认或拒绝此转账。",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		"de-DE": "Du hast dem bereits zugestimmt.",
		"en-UK": "This transfer has been accepted already.",
		"en-US": "This transfer has been accepted already.",
		"es-ES": "Esta transacción ya ha sido aceptada.",
		"fa-IR": "این تراکنش قبلا قبول شده است.",
		"fr-FR": "Ce transfert a déjà été accepté.",
		"id-ID": "Transfer ini sudah diterima.",
		"it-IT": "Questo debito/credito e' gia' stato accettato.",
		"ja-JP": "この転送はすでに承認されています。",
		"ko-KO": "이 전송은 이미 수락되었습니다.",
		"pl-PL": "Ten transfer został już zaakceptowany.",
		"pt-BR": "Esta transferência já foi aceita.",
		"ru-RU": "Эта транзакция уже подтверждена.",
		"tr-TR": "Bu transfer zaten kabul edildi.",
		"ua-UA": "Цей переказ вже підтверджено.",
		"uz-UZ": "Bu o'tkazma allaqachon qabul qilingan.",
		"zh-CN": "此转账已被接受。",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		"de-DE": "Du hast dem bereits widersprochen.",
		"en-UK": "This transfer has been declined already.",
		"en-US": "This transfer has been declined already.",
		"es-ES": "Esta transacción ya ha sido rechazada.",
		"fa-IR": "این تراکنش قبلاً رد شده است.",
		"fr-FR": "Ce transfert a déjà été refusé.",
		"id-ID": "Transfer ini sudah ditolak.",
		"it-IT": "Questo debito/credito e' gia' stato rifiutato.",
		"ja-JP": "この転送はすでに拒否されています。",
		"ko-KO": "이 전송은 이미 거부되었습니다.",
		"pl-PL": "Ten transfer został już odrzucony.",
		"pt-BR": "Esta transferência já foi recusada.",
		"ru-RU": "Эта транзакция уже отклонена.",
		"tr-TR": "Bu transfer zaten reddedildi.",
		"ua-UA": "Цей переказ вже відхилено.",
		"uz-UZ": "Bu o'tkazma allaqachon rad etilgan.",
		"zh-CN": "此转账已被拒绝。",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		"de-DE": "Details hier: %v",
		"en-UK": "Details here: %v",
		"en-US": "Details here: %v",
		"es-ES": "Detalles aquí: %v",
		"fa-IR": "جزئیات: %v",
		"fr-FR": "Détails ici: %v",
		"id-ID": "Detail di sini: %v",
		"it-IT": "Maggiori dettagli qui: %v",
		"ja-JP": "詳細はこちら: %v",
		"ko-KO": "세부 정보: %v",
		"pl-PL": "Szczegóły tutaj: %v",
		"pt-BR": "Detalhes aqui: %v",
		"ru-RU": "Подробнее тут: %v",
		"tr-TR": "Ayrıntılar burada: %v",
		"ua-UA": "Деталі тут: %v",
		"uz-UZ": "Tafsilotlar shu yerda: %v",
		"zh-CN": "详情在这里: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		"de-DE": "Bitte gib die Telefonnummer von <b>%v</b> an:",
		"en-UK": "Please provide phone number for <b>%v</b>",
		"en-US": "Please provide phone number for <b>%v</b>",
		"es-ES": "Por favor escribe el número de teléfono de <b>%v</b>",
		"fa-IR": "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
		"fr-FR": "Veuillez fournir le numéro de téléphone pour <b>%v</b>",
		"id-ID": "Silakan berikan nomor telepon untuk <b>%v</b>",
		"it-IT": "Per favore fornisci il numero di telefono di <b>%v</b>",
		"ja-JP": "<b>%v</b>の電話番号を入力してください",
		"ko-KO": "<b>%v</b>의 전화번호를 입력해주세요",
		"pl-PL": "Proszę podać numer telefonu dla <b>%v</b>",
		"pt-BR": "Por favor, forneça o número de telefone para <b>%v</b>",
		"ru-RU": "Пожалуйста напишите номер телефона <b>%v</b>.",
		"tr-TR": "Lütfen <b>%v</b> için telefon numarası girin",
		"ua-UA": "Будь ласка, вкажіть номер телефону для <b>%v</b>",
		"uz-UZ": "Iltimos, <b>%v</b> uchun telefon raqamini kiriting",
		"zh-CN": "请提供<b>%v</b>的电话号码",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		"de-DE": "Wenn die Telefonnummer in deinem Adressbuch ist, kannst du den <b>%v Button benutzen</b>, um einen Kontakt zu senden.",
		"en-UK": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"en-US": "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		"es-ES": "Si el número está en tu agenda puedes <b>usar %v el botón</b> para enviar el contacto.",
		"fa-IR": "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
		"fr-FR": "Si le numéro de téléphone est dans votre carnet d'adresses, vous pouvez <b>utiliser le bouton %v</b> pour envoyer le contact.",
		"id-ID": "Jika nomor telepon ada di buku alamat Anda, Anda dapat <b>menggunakan tombol %v</b> untuk mengirim kontak.",
		"it-IT": "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
		"ja-JP": "電話番号がアドレス帳にある場合は、<b>%vボタンを使用</b>して連絡先を送信できます。",
		"ko-KO": "전화번호가 주소록에 있는 경우 <b>%v 버튼을 사용</b>하여 연락처를 보낼 수 있습니다.",
		"pl-PL": "Jeśli numer telefonu znajduje się w książce adresowej, możesz <b>użyć przycisku %v</b>, aby wysłać kontakt.",
		"pt-BR": "Se o número de telefone estiver em sua agenda, você pode <b>usar o botão %v</b> para enviar o contato.",
		"ru-RU": "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		"tr-TR": "Telefon numarası adres defterinizde varsa, kişiyi göndermek için <b>%v düğmesini kullanabilirsiniz</b>.",
		"ua-UA": "Якщо номер телефону є у вашій адресній книзі, ви можете <b>використати кнопку %v</b>, щоб надіслати контакт.",
		"uz-UZ": "Agar telefon raqami manzillar kitobingizda bo'lsa, kontaktni yuborish uchun <b>%v tugmasidan foydalanishingiz</b> mumkin.",
		"zh-CN": "如果电话号码在您的通讯录中，您可以<b>使用%v按钮</b>发送联系人。",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		"de-DE": "Die Telefonnummer sollte dem internationalen Standard entsprechen:\n\t* Beginnend mit '+' gefolgt vom Ländercode (Deutschland +49)\n\t* Consist of numbers only\nExample: <b>+49</b><code>157123456</code>",
		"en-UK": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		"en-US": "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		"es-ES": "El número debe tener formato internacional estándar:\n\t* Empezar con '+' seguido del código del país\n\t* formado solo por números\nEjemplo: <b>+1</b><code>999012345678</code>",
		"fa-IR": "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <b>+1</b><code>999012345678</code>",
		"fr-FR": "Le numéro doit être au format international:\n\t* Commence par '+' suivi du code du pays\n\t* Composé uniquement de chiffres\nExemple: <b>+33</b><code>612345678</code>",
		"id-ID": "Nomor harus dalam standar internasional:\n\t* Dimulai dengan '+' diikuti oleh kode negara\n\t* Hanya terdiri dari angka\nContoh: <b>+62</b><code>812345678</code>",
		"it-IT": "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <b>+39</b><code>34612345678</code>",
		"ja-JP": "番号は国際標準形式である必要があります:\n\t* '+'で始まり、国コードが続く\n\t* 数字のみで構成\n例: <b>+81</b><code>9012345678</code>",
		"ko-KO": "번호는 국제 표준이어야 합니다:\n\t* '+'로 시작하여 국가 코드가 뒤따름\n\t* 숫자로만 구성\n예: <b>+82</b><code>1012345678</code>",
		"pl-PL": "Numer powinien być w standardzie międzynarodowym:\n\t* Zaczyna się od '+' a następnie kod kraju\n\t* Składa się tylko z cyfr\nPrzykład: <b>+48</b><code>512345678</code>",
		"pt-BR": "O número deve estar no padrão internacional:\n\t* Começa com '+' seguido pelo código do país\n\t* Consiste apenas de números\nExemplo: <b>+55</b><code>11912345678</code>",
		"ru-RU": "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <b>+7</b><code>999012345678</code>",
		"tr-TR": "Numara uluslararası standartta olmalıdır:\n\t* '+' ile başlayıp ülke kodu ile devam eder\n\t* Sadece rakamlardan oluşur\nÖrnek: <b>+90</b><code>5301234567</code>",
		"ua-UA": "Номер повинен бути в міжнародному форматі:\n\t* Починатися зі знаку '+' та коду країни\n\t* Складатися тільки з цифр\nПриклад: <b>+380</b><code>991234567</code>",
		"uz-UZ": "Raqam xalqaro standartda bo'lishi kerak:\n\t* '+' bilan boshlanib, mamlakat kodi bilan davom etadi\n\t* Faqat raqamlardan iborat\nMisol: <b>+998</b><code>901234567</code>",
		"zh-CN": "号码应符合国际标准:\n\t* 以'+'开头，后跟国家代码\n\t* 仅由数字组成\n示例: <b>+86</b><code>13912345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		"de-DE": "Wir werden eine SMS an diese Nummer schicken:",
		"en-UK": "Will send an SMS to this number:",
		"en-US": "Will send an SMS to this number:",
		"es-ES": "Enviaremos una SMS a este número:",
		"fa-IR": "یک پیام کوتاه به این شماره ارسال خواهد شد:",
		"fr-FR": "Nous enverrons un SMS à ce numéro:",
		"id-ID": "Akan mengirim SMS ke nomor ini:",
		"it-IT": "Invieremo un SMS a questo numero:",
		"ja-JP": "このSMSをこの番号に送信します:",
		"ko-KO": "이 번호로 SMS를 보낼 것입니다:",
		"pl-PL": "Wyślemy SMS na ten numer:",
		"pt-BR": "Enviaremos um SMS para este número:",
		"ru-RU": "На этот номер мы отправим SMS:",
		"tr-TR": "Bu numaraya SMS gönderilecek:",
		"ua-UA": "Надішлемо SMS на цей номер:",
		"uz-UZ": "Bu raqamga SMS yuboriladi:",
		"zh-CN": "将向此号码发送短信:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		"de-DE": `<b>%v</b> schuldet dir <b>%v</b>.`,
		"en-UK": `<b>%v</b> owes to you <b>%v</b>.`,
		"en-US": `<b>%v</b> owes to you <b>%v</b>.`,
		"es-ES": `<b>%v</b> has prestado <b>%v</b>.`,
		"fa-IR": `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
		"fr-FR": `<b>%v</b> vous doit <b>%v</b>.`,
		"id-ID": `<b>%v</b> berhutang kepada Anda <b>%v</b>.`,
		"it-IT": `<b>%v</b> e' in debito di <b>%v</b>.`,
		"ja-JP": `<b>%v</b>はあなたに<b>%v</b>を借りています。`,
		"ko-KO": `<b>%v</b>님이 당신에게 <b>%v</b>를 빚지고 있습니다.`,
		"pl-PL": `<b>%v</b> jest ci winien <b>%v</b>.`,
		"pt-BR": `<b>%v</b> deve a você <b>%v</b>.`,
		"ru-RU": `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		"tr-TR": `<b>%v</b> size <b>%v</b> borçlu.`,
		"ua-UA": `<b>%v</b> заборгував вам <b>%v</b>.`,
		"uz-UZ": `<b>%v</b> sizga <b>%v</b> qarzdir.`,
		"zh-CN": `<b>%v</b> 欠您 <b>%v</b>。`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		"de-DE": "Du schuldest <b>%v</b> <b>%v</b>.",
		"en-UK": "You owe to <b>%v</b> <b>%v</b>.",
		"en-US": "You owe to <b>%v</b> <b>%v</b>.",
		"es-ES": "Te ha prestado <b>%v</b> <b>%v</b>.",
		"fa-IR": "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
		"fr-FR": "Vous devez <b>%v</b> à <b>%v</b>.",
		"id-ID": "Anda berhutang kepada <b>%v</b> <b>%v</b>.",
		"it-IT": `Tu devi dare a <b>%v</b> <b>%v</b>.`,
		"ja-JP": "あなたは<b>%v</b>に<b>%v</b>を借りています。",
		"ko-KO": "당신은 <b>%v</b>에게 <b>%v</b>를 빚지고 있습니다.",
		"pl-PL": "Jesteś winien <b>%v</b> <b>%v</b>.",
		"pt-BR": "Você deve a <b>%v</b> <b>%v</b>.",
		"ru-RU": `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		"tr-TR": "<b>%v</b>'e <b>%v</b> borçlusunuz.",
		"ua-UA": "Ви заборгували <b>%v</b> <b>%v</b>.",
		"uz-UZ": "Siz <b>%v</b>ga <b>%v</b> qarzdorsiz.",
		"zh-CN": "您欠<b>%v</b> <b>%v</b>。",
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {

		"de-DE": `Wurde diese Schuld vollständig beglichen?

		<i>Falls nur teilweise, kann der Teilbetrag direkt eingegeben werden.</i>`,
		"en-UK": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,
		"en-US": `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,

		"es-ES": `¿Ha sido totalmente devuelta esta deuda?

		<i>si ha sido devuelta parcialmente puedes introducir el importe.</i>`,

		"fa-IR": `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,

		"fr-FR": `Cette dette a-t-elle été remboursée intégralement?

		<i>Si partiellement, vous pouvez saisir le montant tout de suite.</i>`,

		"id-ID": `Apakah hutang ini telah dikembalikan sepenuhnya?

		<i>Jika sebagian, Anda dapat memasukkan jumlah langsung.</i>`,

		"it-IT": `Il debito e' stato saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare da sottrarre, direttamente qui sotto.</i>`,

		"ja-JP": `この借金は全額返済されましたか？

		<i>部分的に返済された場合は、すぐに金額を入力できます。</i>`,

		"ko-KO": `이 빚이 전액 상환되었습니까?

		<i>부분적으로 상환된 경우 금액을 바로 입력할 수 있습니다.</i>`,

		"pl-PL": `Czy ten dług został spłacony w całości?

		<i>Jeśli częściowo, możesz od razu wprowadzić kwotę.</i>`,

		"pt-BR": `Esta dívida foi devolvida integralmente?

		<i>Se parcialmente, você pode inserir o valor imediatamente.</i>`,

		"ru-RU": `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		"tr-TR": `Bu borç tamamen geri ödendi mi?

		<i>Kısmen ödendiyse, tutarı hemen girebilirsiniz.</i>`,

		"ua-UA": `Чи повернуто цей борг повністю?

		<i>Якщо частково, ви можете відразу ввести суму.</i>`,

		"uz-UZ": `Bu qarz to'liq qaytarildimi?

		<i>Agar qisman bo'lsa, miqdorni darhol kiritishingiz mumkin.</i>`,

		"zh-CN": `这笔债务是否已全额归还？

		<i>如果部分归还，您可以立即输入金额。</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		"de-DE": `Das Programm ist <b>kostenlos</b>. Bitte <a href="https://debtstracker.io/en/help-us">hilf</a> es besser zu machen!`,
		"en-UK": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		"en-US": `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		"es-ES": `Este programa es <b>gratis</b>. Por favor <a href="https://debtstracker.io/en/help-us">ayúdanos</a> a mejorarlo!`,
		"fa-IR": `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
		"fr-FR": `Ce programme est <b>gratuit</b>. S'il vous plaît <a href="https://debtstracker.io/en/help-us">aidez-nous</a> à l'améliorer!`,
		"id-ID": `Program ini <b>gratis untuk digunakan</b>. Silakan <a href="https://debtstracker.io/en/help-us">bantu</a> untuk membuatnya lebih baik!`,
		"it-IT": `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
		"ja-JP": `このプログラムは<b>無料で使用できます</b>。より良くするために<a href="https://debtstracker.io/en/help-us">ご協力</a>ください！`,
		"ko-KO": `이 프로그램은 <b>무료로 사용</b>할 수 있습니다. 더 나은 프로그램을 만들기 위해 <a href="https://debtstracker.io/en/help-us">도와주세요</a>!`,
		"pl-PL": `Ten program jest <b>darmowy</b>. Proszę <a href="https://debtstracker.io/en/help-us">pomóż</a> uczynić go lepszym!`,
		"pt-BR": `Este programa é <b>gratuito para uso</b>. Por favor, <a href="https://debtstracker.io/en/help-us">ajude</a> a torná-lo melhor!`,
		"ru-RU": `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,
		"tr-TR": `Bu program <b>ücretsiz kullanım</b> içindir. Lütfen daha iyi hale getirmek için <a href="https://debtstracker.io/en/help-us">yardım edin</a>!`,
		"ua-UA": `Ця програма <b>безкоштовна</b>. Будь ласка, <a href="https://debtstracker.io/en/help-us">допоможіть</a> зробити її кращою!`,
		"uz-UZ": `Bu dastur <b>bepul foydalanish</b> uchun. Iltimos, uni yaxshilash uchun <a href="https://debtstracker.io/en/help-us">yordam bering</a>!`,
		"zh-CN": `此程序<b>免费使用</b>。请<a href="https://debtstracker.io/en/help-us">帮助</a>我们使其变得更好！`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		"de-DE": "%v | du schuldest: %v",
		"en-UK": "%v | you owe: %v",
		"en-US": "%v | you owe: %v",
		"es-ES": "%v | tú debes: %v",
		"fa-IR": "%v | شما بدهکارید: %v",
		"fr-FR": "%v | vous devez: %v",
		"id-ID": "%v | Anda berhutang: %v",
		"it-IT": "%v | tu devi: %v",
		"ja-JP": "%v | あなたの借り: %v",
		"ko-KO": "%v | 당신이 빚진 금액: %v",
		"pl-PL": "%v | jesteś winien: %v",
		"pt-BR": "%v | você deve: %v",
		"ru-RU": "%v | вы должны: %v",
		"tr-TR": "%v | borçlusunuz: %v",
		"ua-UA": "%v | ви винні: %v",
		"uz-UZ": "%v | siz qarzdorsiz: %v",
		"zh-CN": "%v | 您欠: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		"de-DE": "%v | schuldet dir: %v",
		"en-UK": "%v | owes to you: %v",
		"en-US": "%v | owes to you: %v",
		"es-ES": "%v | te debe: %v",
		"fa-IR": "%v | به شما بدهکار است: %v",
		"fr-FR": "%v | vous doit: %v",
		"id-ID": "%v | berhutang kepada Anda: %v",
		"it-IT": "%v | ti deve: %v",
		"ja-JP": "%v | あなたへの借り: %v",
		"ko-KO": "%v | 당신에게 빚진 금액: %v",
		"pl-PL": "%v | jest ci winien: %v",
		"pt-BR": "%v | deve a você: %v",
		"ru-RU": "%v | долг вам: %v",
		"tr-TR": "%v | size borçlu: %v",
		"ua-UA": "%v | винен вам: %v",
		"uz-UZ": "%v | sizga qarzdor: %v",
		"zh-CN": "%v | 欠您: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		"de-DE": "Ja, vollständig",
		"en-UK": "Yes, fully",
		"en-US": "Yes, fully",
		"es-ES": "Sí, completamente",
		"fa-IR": "بله، به صورت کامل",
		"fr-FR": "Oui, entièrement",
		"id-ID": "Ya, sepenuhnya",
		"it-IT": "Si, completamente",
		"ja-JP": "はい、完全に",
		"ko-KO": "예, 완전히",
		"pl-PL": "Tak, w całości",
		"pt-BR": "Sim, totalmente",
		"ru-RU": "Да, целиком",
		"tr-TR": "Evet, tamamen",
		"ua-UA": "Так, повністю",
		"uz-UZ": "Ha, to'liq",
		"zh-CN": "是的，完全",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		"de-DE": "Nein, nur teilweise",
		"en-UK": "No, just partially",
		"en-US": "No, just partially",
		"es-ES": "No, solo parcialmente",
		"fa-IR": "خیر، تنها قسمتی",
		"fr-FR": "Non, seulement partiellement",
		"id-ID": "Tidak, hanya sebagian",
		"it-IT": "No, parzialmente",
		"ja-JP": "いいえ、部分的に",
		"ko-KO": "아니요, 부분적으로만",
		"pl-PL": "Nie, tylko częściowo",
		"pt-BR": "Não, apenas parcialmente",
		"ru-RU": "Нет, только часть",
		"tr-TR": "Hayır, sadece kısmen",
		"ua-UA": "Ні, лише частково",
		"uz-UZ": "Yo'q, faqat qisman",
		"zh-CN": "不，仅部分",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		"de-DE": "Du solltest dich nicht selber einladen ;)",
		"en-UK": "You should not use your own invite ;)",
		"en-US": "You should not use your own invite ;)",
		"es-ES": "No deberías invitarte a ti mismo ;)",
		"fa-IR": "نباید از دعوت خود استفاده کنید ;)",
		"fr-FR": "Vous ne devriez pas utiliser votre propre invitation ;)",
		"id-ID": "Anda tidak boleh menggunakan undangan Anda sendiri ;)",
		"it-IT": "Non dovresti usare il tuo codice invito con te stesso :)",
		"ja-JP": "自分の招待を使用すべきではありません ;)",
		"ko-KO": "자신의 초대를 사용해서는 안 됩니다 ;)",
		"pl-PL": "Nie powinieneś używać własnego zaproszenia ;)",
		"pt-BR": "Você não deve usar seu próprio convite ;)",
		"ru-RU": "Хорошая попытка пригласить самого себя ;)",
		"tr-TR": "Kendi davetinizi kullanmamalısınız ;)",
		"ua-UA": "Ви не повинні використовувати власне запрошення ;)",
		"uz-UZ": "Siz o'z taklifingizdan foydalanmasligingiz kerak ;)",
		"zh-CN": "您不应该使用自己的邀请 ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		"de-DE": "Willkommen. Schön, dass du der Einladung gefolgt bist!",
		"en-UK": "Welcome and thanks for accepting the invite!",
		"en-US": "Welcome and thanks for accepting the invite!",
		"es-ES": "Bienvenido y gracias por aceptar la invitación",
		"fa-IR": "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
		"fr-FR": "Bienvenue et merci d'avoir accepté l'invitation!",
		"id-ID": "Selamat datang dan terima kasih telah menerima undangan!",
		"it-IT": "Benvenuto e grazie per aver accettato l'invito!",
		"ja-JP": "ようこそ、招待を受け入れていただきありがとうございます！",
		"ko-KO": "환영합니다, 초대를 수락해 주셔서 감사합니다!",
		"pl-PL": "Witamy i dziękujemy za przyjęcie zaproszenia!",
		"pt-BR": "Bem-vindo e obrigado por aceitar o convite!",
		"ru-RU": "Спасибо за то что воспользовались приглашением!",
		"tr-TR": "Hoş geldiniz ve daveti kabul ettiğiniz için teşekkürler!",
		"ua-UA": "Ласкаво просимо та дякуємо за прийняття запрошення!",
		"uz-UZ": "Xush kelibsiz va taklifni qabul qilganingiz uchun rahmat!",
		"zh-CN": "欢迎并感谢您接受邀请！",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		"de-DE": "Das darf nur %v.",
		"en-UK": "This action for %v only.",
		"en-US": "This action for %v only.",
		"es-ES": "Esta acción está disponible solo para%v.",
		"fa-IR": "این عمل تنها برای %v می باشد.",
		"fr-FR": "Cette action est uniquement pour %v.",
		"id-ID": "Tindakan ini hanya untuk %v.",
		"it-IT": "Questa azione e' disponibile solo per %v.",
		"ja-JP": "このアクションは%vのみです。",
		"ko-KO": "이 작업은 %v만을 위한 것입니다.",
		"pl-PL": "Ta akcja tylko dla %v.",
		"pt-BR": "Esta ação é apenas para %v.",
		"ru-RU": "Это действие доступно только для %v.",
		"tr-TR": "Bu işlem sadece %v için.",
		"ua-UA": "Ця дія тільки для %v.",
		"uz-UZ": "Bu harakat faqat %v uchun.",
		"zh-CN": "此操作仅适用于%v。",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		"de-DE": "Quittungsdetails anzeigen",
		"en-UK": "Show receipt details",
		"en-US": "Show receipt details",
		"es-ES": "Mostrar detalles",
		"fa-IR": "جزئیات رسید را نشان بده",
		"fr-FR": "Afficher les détails du reçu",
		"id-ID": "Tampilkan detail tanda terima",
		"it-IT": "Mostra i dettagli del debito/credito",
		"ja-JP": "領収書の詳細を表示",
		"ko-KO": "영수증 세부 정보 표시",
		"pl-PL": "Pokaż szczegóły paragonu",
		"pt-BR": "Mostrar detalhes do recibo",
		"ru-RU": "Показать детали",
		"tr-TR": "Makbuz detaylarını göster",
		"ua-UA": "Показати деталі квитанції",
		"uz-UZ": "Kvitansiya tafsilotlarini ko'rsatish",
		"zh-CN": "显示收据详情",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		"de-DE": "Du hast gewählt, einen Freund per Mail einzuladen.",
		"en-UK": "You've selected to invite friend by email.",
		"en-US": "You've selected to invite friend by email.",
		"es-ES": "Has decidido invitar a un amigo por e-mail.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
		"fr-FR": "Vous avez choisi d'inviter un ami par e-mail.",
		"id-ID": "Anda telah memilih untuk mengundang teman melalui email.",
		"it-IT": "Hai scelto di invitare l'amico tramite email.",
		"ja-JP": "友達をメールで招待することを選択しました。",
		"ko-KO": "이메일로 친구를 초대하기로 선택하셨습니다.",
		"pl-PL": "Wybrałeś zaproszenie znajomego przez e-mail.",
		"pt-BR": "Você selecionou convidar um amigo por e-mail.",
		"ru-RU": "Вы решили пригласить друга через email.",
		"tr-TR": "Arkadaşınızı e-posta ile davet etmeyi seçtiniz.",
		"ua-UA": "Ви вибрали запросити друга електронною поштою.",
		"uz-UZ": "Siz do'stingizni elektron pochta orqali taklif qilishni tanladingiz.",
		"zh-CN": "您已选择通过电子邮件邀请朋友。",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		"de-DE": "Du hast gewählt, einen Freund per SMS einzuladen.",
		"en-UK": "You've selected to invite friend by SMS.",
		"en-US": "You've selected to invite friend by SMS.",
		"es-ES": "Has decidido invitar a un amigo por SMS.",
		"fa-IR": "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
		"fr-FR": "Vous avez choisi d'inviter un ami par SMS.",
		"id-ID": "Anda telah memilih untuk mengundang teman melalui SMS.",
		"it-IT": "Hai scelto di invitare l'amico tramite SMS.",
		"ja-JP": "友達をSMSで招待することを選択しました。",
		"ko-KO": "SMS로 친구를 초대하기로 선택하셨습니다.",
		"pl-PL": "Wybrałeś zaproszenie znajomego przez SMS.",
		"pt-BR": "Você selecionou convidar um amigo por SMS.",
		"ru-RU": "Вы решили пригласить друга через SMS.",
		"tr-TR": "Arkadaşınızı SMS ile davet etmeyi seçtiniz.",
		"ua-UA": "Ви вибрали запросити друга через SMS.",
		"uz-UZ": "Siz do'stingizni SMS orqali taklif qilishni tanladingiz.",
		"zh-CN": "您已选择通过短信邀请朋友。",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		"de-DE": "Wie möchtest du den Code weitergeben?",

		"en-UK": `How do you want to pass the invite code?`,

		"en-US": `How do you want to pass the invite code?`,

		"es-ES": `¿Cómo quieres enviarle el código?`,

		"fa-IR": `آیا میخواهید کد دعوت را ارسال کنید؟`,

		"fr-FR": `Comment voulez-vous transmettre le code d'invitation?`,

		"id-ID": `Bagaimana Anda ingin meneruskan kode undangan?`,

		"it-IT": `Come vuoi inviargli il codice invito?`,

		"ja-JP": `招待コードをどのように渡しますか？`,

		"ko-KO": `초대 코드를 어떻게 전달하시겠습니까?`,

		"pl-PL": `Jak chcesz przekazać kod zaproszenia?`,

		"pt-BR": `Como você deseja passar o código de convite?`,

		"ru-RU": `Как вы хотите передать код приглашение?`,

		"tr-TR": `Davet kodunu nasıl iletmek istiyorsunuz?`,

		"ua-UA": `Як ви хочете передати код запрошення?`,

		"uz-UZ": `Taklif kodini qanday o'tkazmoqchisiz?`,

		"zh-CN": `您想如何传递邀请码？`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		"de-DE": "%v hat Erinnerungen über folgendes Anliegen blockiert: %v",
		"en-UK": "%v blocked reminders about transactions by: %v",
		"en-US": "%v blocked reminders about transactions by: %v",
		"es-ES": "%v ha bloqueado las notificaciones de las transacciones por: %v",
		"fa-IR": "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
		"fr-FR": "%v a bloqué les rappels concernant les transactions par: %v",
		"id-ID": "%v memblokir pengingat tentang transaksi oleh: %v",
		"it-IT": "%v bloccato promemoria riguardo la transazione da: %v.",
		"ja-JP": "%vは%vによるトランザクションに関するリマインダーをブロックしました",
		"ko-KO": "%v님이 %v의 거래에 대한 알림을 차단했습니다",
		"pl-PL": "%v zablokował przypomnienia o transakcjach przez: %v",
		"pt-BR": "%v bloqueou lembretes sobre transações por: %v",
		"ru-RU": "%v заблокировал получение оповешений о транзакиях через: %v.",
		"tr-TR": "%v, %v tarafından yapılan işlemler hakkındaki hatırlatıcıları engelledi",
		"ua-UA": "%v заблокував нагадування про транзакції від: %v",
		"uz-UZ": "%v %v tomonidan tranzaksiyalar haqida eslatmalarni blokladi",
		"zh-CN": "%v已阻止来自%v的交易提醒",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		"de-DE": "Warte eine Sekunde...",
		"en-UK": "Wait a second...",
		"en-US": "Wait a second...",
		"es-ES": "Un segundo...",
		"fa-IR": "یک ثانیه صبر کنید ...",
		"fr-FR": "Attendez une seconde...",
		"id-ID": "Tunggu sebentar...",
		"it-IT": "Solo un attimo...",
		"ja-JP": "ちょっと待って...",
		"ko-KO": "잠시만 기다려주세요...",
		"pl-PL": "Poczekaj chwilę...",
		"pt-BR": "Espere um segundo...",
		"ru-RU": "Секундочку...",
		"tr-TR": "Bir saniye bekleyin...",
		"ua-UA": "Зачекайте секунду...",
		"uz-UZ": "Bir soniya kuting...",
		"zh-CN": "稍等一下...",
	},
	HTML_USING_TELEGRAM: {
		"de-DE": "benutze Telegram messenger",
		"en-UK": "using Telegram messenger",
		"en-US": "using Telegram messenger",
		"es-ES": "Usando Telegram",
		"fa-IR": "استفاده از پیام رسان تلگرام",
		"fr-FR": "utilisant Telegram messenger",
		"id-ID": "menggunakan Telegram messenger",
		"it-IT": "usa Telegram",
		"ja-JP": "Telegramメッセンジャーを使用",
		"ko-KO": "텔레그램 메신저 사용",
		"pl-PL": "używając komunikatora Telegram",
		"pt-BR": "usando o Telegram messenger",
		"ru-RU": "используя Telegram мессенджер",
		"tr-TR": "Telegram messenger kullanarak",
		"ua-UA": "використовуючи Telegram messenger",
		"uz-UZ": "Telegram messenjeridan foydalanish",
		"zh-CN": "使用Telegram信使",
	},
	COMMAND_TEXT_ACCEPT: {
		"de-DE": "Akzeptieren",
		"en-UK": "Accept",
		"en-US": "Accept",
		"es-ES": "Aceptar",
		"fa-IR": "قبول",
		"fr-FR": "Accepter",
		"id-ID": "Terima",
		"it-IT": "Accetta",
		"ja-JP": "承認",
		"ko-KO": "수락",
		"pl-PL": "Akceptuj",
		"pt-BR": "Aceitar",
		"ru-RU": "Согласиться",
		"tr-TR": "Kabul et",
		"ua-UA": "Прийняти",
		"uz-UZ": "Qabul qilish",
		"zh-CN": "接受",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	"en-UK": "Accept ",
	//	"es-ES": "Confirmar ",

	//	"fa-IR": "قبول ",

	//  "it-IT": "Accetta",
	//	"ru-RU": "Подтвердить ",

	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	"en-UK": "Decline (using Telegram messenger)",
	//	"es-ES": "Rechazar (usandoTelegram)",

	//	"fa-IR": "رد ( استفاده از پیام رسان تلگرام)",

	//  "it-IT": "Rifiuta (usando Telegram)",
	//	"ru-RU": "Отказаться (используя Telegram)",

	//},
	COMMAND_TEXT_DECLINE: {
		"de-DE": "Ablehnen",
		"en-UK": "Decline",
		"es-ES": "Rechazar",
		"fa-IR": "رد",
		"it-IT": "Rifiuta",
		"ru-RU": "Отклонить",
	},
	FamilyMember: {
		"de-DE": "Familienmitglied",
		"en-UK": "Family member",
		"en-US": "Family member", // Placeholder
		"es-ES": "Miembro de la familia",
		"fa-IR": "عضو خانواده",
		"fr-FR": "Membre de la famille", // Placeholder
		"id-ID": "Anggota keluarga",     // Placeholder
		"it-IT": "Membro della famiglia",
		"ja-JP": "家族の一員",             // Placeholder
		"ko-KO": "가족 구성원",            // Placeholder
		"pl-PL": "Członek rodziny",   // Placeholder
		"pt-BR": "Membro da família", // Placeholder
		"ru-RU": "Член семьи",
		"tr-TR": "Aile üyesi",  // Placeholder
		"ua-UA": "Член родини", // Placeholder
		"uz-UZ": "Oila aʼzosi", // Placeholder
		"zh-CN": "家庭成员",        // Placeholder
	},
	UserHasNotJoinedSpaceYet: {
		"de-DE": "Dieser Kontakt ist diesem Bereich noch nicht beigetreten.",
		"en-UK": "This contact has not joined this space yet.",
		"es-ES": "Este contacto aún no se ha unido a este espacio.",
		"fa-IR": "این مخاطب هنوز به این فضا نپیوسته است.",
		"fr-FR": "Ce contact n'a pas encore rejoint cet espace.",
		"id-ID": "Kontak ini belum bergabung dengan ruang ini.",
		"it-IT": "Questo contatto non si è ancora unito a questo spazio.",
		"ja-JP": "この連絡先はまだこのスペースに参加していません。",
		"ko-KO": "이 연락처는 아직 이 공간에 참여하지 않았습니다.",
		"pl-PL": "Ten kontakt nie dołączył jeszcze do tej przestrzeni.",
		"pt-BR": "Este contato ainda não entrou neste espaço.",
		"ru-RU": "Eще не присоединился к этому пространству.",
		"tr-TR": "Bu kişi henüz bu alana katılmadı.",
		"ua-UA": "Цей контакт ще не приєднався до цього простору.",
		"uz-UZ": "Bu kontakt hali bu joyga qoʻshilmadi.",
		"zh-CN": "该联系人尚未加入此空间。",
	},
	UserHasNotJoinedFamilySpaceYet: {
		"de-DE": "Ist diesem Familienbereich noch nicht beigetreten.",
		"en-UK": "Has not joined this family space yet.",
		"es-ES": "Aún no se ha unido a este espacio familiar.",
		"fa-IR": "هنوز به این فضای خانوادگی نپیوسته است.",
		"fr-FR": "N'a pas encore rejoint cet espace familial.",
		"id-ID": "Belum bergabung dengan ruang keluarga ini.",
		"it-IT": "Non si è ancora unito a questo spazio familiare.",
		"ja-JP": "このファミリースペースにまだ参加していません。",
		"ko-KO": "아직 이 가족 공간에 참여하지 않았습니다.",
		"pl-PL": "Nie dołączył jeszcze do tej przestrzeni rodzinnej.",
		"pt-BR": "Ainda não entrou neste espaço familiar.",
		"ru-RU": "Eще не присоединился к этому семейному пространству.",
		"tr-TR": "Henüz bu aile alanına katılmadı.",
		"ua-UA": "Ще не приєднався до цього сімейного простору.",
		"uz-UZ": "Hali bu oila bo'shlig'iga qo'shilmagan.",
		"zh-CN": "尚未加入该家庭空间。",
	},
	BtnSendInviteByTelegram: {
		"de-DE": "Einladung über Telegram senden",
		"en-UK": "Send invite over Telegram",
		"en-US": "Send invite over Telegram",
		"es-ES": "Enviar invitación por Telegram",
		"fa-IR": "ارسال دعوتنامه از طریق تلگرام",
		"fr-FR": "Envoyer une invitation via Telegram",
		"id-ID": "Kirim undangan melalui Telegram",
		"it-IT": "Invia invito tramite Telegram",
		"ja-JP": "Telegramで招待を送る",
		"ko-KO": "텔레그램으로 초대 보내기",
		"pl-PL": "Wyślij zaproszenie przez Telegram",
		"pt-BR": "Enviar convite pelo Telegram",
		"ru-RU": "Отправить приглашение через Telegram",
		"tr-TR": "Telegram üzerinden davetiye gönder",
		"ua-UA": "Надіслати запрошення через Telegram",
		"uz-UZ": "Telegram orqali taklif yuboring",
		"zh-CN": "通过Telegram发送邀请",
	},
	BtnTextAcceptInvite: {
		"de-DE": "Akzeptiere Einladung",
		"en-UK": "Accept invite",
		"en-US": "Accept invite",
		"es-ES": "Aceptar la invitación",
		"fa-IR": "قبول دعوت",
		"fr-FR": "Accepter l'invitation",
		"id-ID": "Terima undangan",
		"it-IT": "Accetta l'invito",
		"ja-JP": "招待を承認",
		"ko-KO": "초대 수락",
		"pl-PL": "Zaakceptuj zaproszenie",
		"pt-BR": "Aceitar convite",
		"ru-RU": "Принять приглашение",
		"tr-TR": "Daveti kabul et",
		"ua-UA": "Прийняти запрошення",
		"uz-UZ": "Taklifni qabul qilish",
		"zh-CN": "接受邀请",
	},
	BtnTextDeclineInvite: {
		"de-DE": "Ablehnen Einladung",
		"en-UK": "Decline invite",
		"en-US": "Decline invite",
		"es-ES": "Rechazar la invitación",
		"fa-IR": "رد دعوت",
		"fr-FR": "Décliner l'invitation",
		"id-ID": "Tolak undangan",
		"it-IT": "Rifiuta l'invito",
		"ja-JP": "招待を拒否",
		"ko-KO": "초대 거절",
		"pl-PL": "Odrzuć zaproszenie",
		"pt-BR": "Recusar convite",
		"ru-RU": "Отклонить приглашение",
		"tr-TR": "Daveti reddet",
		"ua-UA": "Відхилити запрошення",
		"uz-UZ": "Taklifni rad etish",
		"zh-CN": "拒绝邀请",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		"de-DE": "Quittungsdetails anzeigen",
		"en-UK": "See receipt details",
		"en-US": "See receipt details",
		"es-ES": "Ver el recibo",
		"fa-IR": "دیدن جزئیات رسید",
		"fr-FR": "Voir les détails du reçu",
		"id-ID": "Lihat detail tanda terima",
		"it-IT": "Vedi dettagli",
		"ja-JP": "領収書の詳細を見る",
		"ko-KO": "영수증 세부 정보 보기",
		"pl-PL": "Zobacz szczegóły paragonu",
		"pt-BR": "Ver detalhes do recibo",
		"ru-RU": "Посмотреть квитанцию",
		"tr-TR": "Makbuz detaylarını görüntüle",
		"ua-UA": "Переглянути квитанцію",
		"uz-UZ": "Kvitansiya tafsilotlarini ko'rish",
		"zh-CN": "查看收据详情",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		"de-DE": "Andere Wege, eine Einladung zu senden",
		"en-UK": "Other ways to send an invite",
		"en-US": "Other ways to send an invite",
		"es-ES": "Otras maneras para enviar la invitación",
		"fa-IR": "سایر راههای ارسال دعوت",
		"fr-FR": "Autres façons d'envoyer une invitation",
		"id-ID": "Cara lain untuk mengirim undangan",
		"it-IT": "Altri modi per inviare un invito",
		"ja-JP": "招待状を送信する他の方法",
		"ko-KO": "초대장을 보내는 다른 방법",
		"pl-PL": "Inne sposoby wysyłania zaproszenia",
		"pt-BR": "Outras formas de enviar um convite",
		"ru-RU": "Другие способы отправить приглашение",
		"tr-TR": "Davetiye göndermenin diğer yolları",
		"ua-UA": "Інші способи надіслати запрошення",
		"uz-UZ": "Taklifnoma yuborishning boshqa usullari",
		"zh-CN": "发送邀请的其他方式",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		"de-DE": "meine Telefonnummer senden",
		"en-UK": "Send my phone number",
		"en-US": "Send my phone number",
		"es-ES": "Enviar mi número",
		"fa-IR": "شماره تلفن مرا ارسال کنید",
		"fr-FR": "Envoyer mon numéro de téléphone",
		"id-ID": "Kirim nomor telepon saya",
		"it-IT": "Invia il mio numero",
		"ja-JP": "私の電話番号を送信する",
		"ko-KO": "내 전화번호 보내기",
		"pl-PL": "Wyślij mój numer telefonu",
		"pt-BR": "Enviar meu número de telefone",
		"ru-RU": "Отправить мой номер",
		"tr-TR": "Telefon numaramı gönder",
		"ua-UA": "Надіслати мій номер телефону",
		"uz-UZ": "Telefon raqamimni yuborish",
		"zh-CN": "发送我的电话号码",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		"de-DE": "per Mail",
		"en-UK": "By Email",
		"en-US": "By Email",
		"es-ES": "Vía e-mail",
		"fa-IR": "بوسیله ی ایمیل",
		"fr-FR": "Par e-mail",
		"id-ID": "Melalui Email",
		"it-IT": "Tramite email",
		"ja-JP": "メールで",
		"ko-KO": "이메일로",
		"pl-PL": "Przez e-mail",
		"pt-BR": "Por e-mail",
		"ru-RU": "Через Email",
		"tr-TR": "E-posta ile",
		"ua-UA": "Через електронну пошту",
		"uz-UZ": "Email orqali",
		"zh-CN": "通过电子邮件",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		"de-DE": "per SMS",
		"en-UK": "By SMS",
		"en-US": "By SMS",
		"es-ES": "Vía SMS",
		"fa-IR": "بوسیله پیام کوتاه",
		"fr-FR": "Par SMS",
		"id-ID": "Melalui SMS",
		"it-IT": "Tramite SMS",
		"ja-JP": "SMSで",
		"ko-KO": "SMS로",
		"pl-PL": "Przez SMS",
		"pt-BR": "Por SMS",
		"ru-RU": "Через SMS",
		"tr-TR": "SMS ile",
		"ua-UA": "Через SMS",
		"uz-UZ": "SMS orqali",
		"zh-CN": "通过短信",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		"de-DE": "Einladen per Telegram",
		"en-UK": "Invite By Telegram",
		"en-US": "Invite By Telegram",
		"es-ES": "Invitar vía Telegram",
		"fa-IR": "دعوت با تلگرام",
		"fr-FR": "Inviter par Telegram",
		"id-ID": "Undang Melalui Telegram",
		"it-IT": "Tramite Telegram",
		"ja-JP": "Telegramで招待する",
		"ko-KO": "텔레그램으로 초대하기",
		"pl-PL": "Zaproś przez Telegram",
		"pt-BR": "Convidar pelo Telegram",
		"ru-RU": "Пригласить через Telegram",
		"tr-TR": "Telegram ile davet et",
		"ua-UA": "Запросити через Telegram",
		"uz-UZ": "Telegram orqali taklif qilish",
		"zh-CN": "通过Telegram邀请",
	},
	MESSAGE_TEXT_INVITE_CREATED: {

		"de-DE": `Wir haben deinen Freund einen Code geschickt. (#%v)

Sobald dein Freund die Einladung akzeptiert hat, könnt ihr das Geld, was ihr euch teit, mit Leichtigkeit managen.`,

		"en-UK": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		"en-US": `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		"es-ES": `Hemos enviado el código de la invitación a tu amigo. (#%v)

Cuando tu amigo accepte la invitación vaís a tener transacciones y balance en común (solo entre vosotros). Todo eso os ayuda minimizar los esfuerzos para controlar la cuenta.`,

		"fa-IR": `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,

		"fr-FR": `Nous avons envoyé un code d'invitation à votre ami. (#%v)

Une fois que votre ami accepte l'invitation, vous partagerez le solde et les transferts entre vous pour vous assurer que vous êtes tous les deux sur la même page avec un minimum d'efforts.`,

		"id-ID": `Kami telah mengirimkan kode undangan kepada teman Anda. (#%v)

Setelah teman Anda menerima undangan, Anda akan berbagi saldo & transfer antara Anda untuk memastikan Anda berdua berada di halaman yang sama dengan upaya minimal.`,

		"it-IT": `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,

		"ja-JP": `友達に招待コードを送信しました。(#%v)

友達が招待を受け入れると、最小限の労力で両方が同じページにいることを確認するために、あなたの間で残高と転送を共有します。`,

		"ko-KO": `친구에게 초대 코드를 보냈습니다. (#%v)

친구가 초대를 수락하면 최소한의 노력으로 두 사람이 같은 페이지에 있는지 확인하기 위해 잔액과 이체를 공유하게 됩니다.`,

		"pl-PL": `Wysłaliśmy kod zaproszenia do Twojego znajomego. (#%v)

Gdy Twój znajomy zaakceptuje zaproszenie, będziecie dzielić saldo i przelewy między sobą, aby upewnić się, że oboje jesteście na tej samej stronie przy minimalnym wysiłku.`,

		"pt-BR": `Enviamos um código de convite para seu amigo. (#%v)

Quando seu amigo aceitar o convite, vocês compartilharão saldo e transferências entre si para garantir que ambos estejam na mesma página com o mínimo de esforço.`,

		"ru-RU": `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		"tr-TR": `Arkadaşınıza bir davet kodu gönderdik. (#%v)

Arkadaşınız daveti kabul ettiğinde, minimum çabayla ikinizin de aynı sayfada olduğundan emin olmak için aranızda bakiye ve transferleri paylaşacaksınız.`,

		"ua-UA": `Ми надіслали код запрошення вашому другу. (#%v)

Коли ваш друг прийме запрошення, ви будете ділитися балансом і переказами між собою, щоб переконатися, що ви обидва на одній сторінці з мінімальними зусиллями.`,

		"uz-UZ": `Do'stingizga taklifnoma kodini yubordik. (#%v)

Do'stingiz taklifni qabul qilgandan so'ng, minimal kuch bilan ikkalangiz ham bir xil sahifada ekanligingizga ishonch hosil qilish uchun o'zaro balans va o'tkazmalarni almashishingiz mumkin.`,

		"zh-CN": `我们已向您的朋友发送邀请码。(#%v)

一旦您的朋友接受邀请，您将共享余额和转账，以确保您们双方以最少的努力保持一致。`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		"de-DE": "Bitte gib die eMail Adresse deines Freundes an, wohin wir den Code schicken:",
		"en-UK": "Please enter emaill address of your friend where we should send an invite code.",
		"en-US": "Please enter email address of your friend where we should send an invite code.",
		"es-ES": "Por favor, introduce el e-maill de tu amigo al cual enviaremos el código de la invitación.",
		"fa-IR": "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
		"fr-FR": "Veuillez entrer l'adresse e-mail de votre ami où nous devrions envoyer un code d'invitation.",
		"id-ID": "Silakan masukkan alamat email teman Anda di mana kami harus mengirimkan kode undangan.",
		"it-IT": "Inserisci l'email dell'amico al quale inviare il codice invito.",
		"ja-JP": "招待コードを送信する友達のメールアドレスを入力してください。",
		"ko-KO": "초대 코드를 보낼 친구의 이메일 주소를 입력하세요.",
		"pl-PL": "Proszę podać adres e-mail znajomego, na który powinniśmy wysłać kod zaproszenia.",
		"pt-BR": "Por favor, insira o endereço de e-mail do seu amigo para onde devemos enviar um código de convite.",
		"ru-RU": "Введите email вашего друга на который мы отправим код приглашения.",
		"tr-TR": "Lütfen davet kodunu göndermemiz gereken arkadaşınızın e-posta adresini girin.",
		"ua-UA": "Будь ласка, введіть електронну адресу вашого друга, куди ми повинні надіслати код запрошення.",
		"uz-UZ": "Iltimos, taklifnoma kodini yuborishimiz kerak bo'lgan do'stingizning elektron pochta manzilini kiriting.",
		"zh-CN": "请输入您朋友的电子邮件地址，我们将向其发送邀请码。",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		"de-DE": "Bitte gib die eMail Adresse deines Freundes (%v), wohin wir die Quittung schicken:", // TODO(DE) - verify
		"en-UK": "Please enter emaill address of your friend (%v) where we should send the receipt.",
		"en-US": "Please enter email address of your friend (%v) where we should send the receipt.",
		"es-ES": "Introduce el e-maill de tu amigo (%v) al cual enviaremos el recibo.",
		"fa-IR": "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
		"fr-FR": "Veuillez entrer l'adresse e-mail de votre ami (%v) où nous devrions envoyer le reçu.",
		"id-ID": "Silakan masukkan alamat email teman Anda (%v) di mana kami harus mengirimkan tanda terima.",
		"it-IT": "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
		"ja-JP": "領収書を送信する友達 (%v) のメールアドレスを入力してください。",
		"ko-KO": "영수증을 보낼 친구 (%v)의 이메일 주소를 입력하세요.",
		"pl-PL": "Proszę podać adres e-mail znajomego (%v), na który powinniśmy wysłać paragon.",
		"pt-BR": "Por favor, insira o endereço de e-mail do seu amigo (%v) para onde devemos enviar o recibo.",
		"ru-RU": "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		"tr-TR": "Lütfen makbuzu göndermemiz gereken arkadaşınızın (%v) e-posta adresini girin.",
		"ua-UA": "Будь ласка, введіть електронну адресу вашого друга (%v), куди ми повинні надіслати квитанцію.",
		"uz-UZ": "Iltimos, kvitansiyani yuborishimiz kerak bo'lgan do'stingizning (%v) elektron pochta manzilini kiriting.",
		"zh-CN": "请输入您朋友 (%v) 的电子邮件地址，我们将向其发送收据。",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		"de-DE": "Bitte gib die Telefonnummer deines Freundes an oder teile einen Kontakt, wohin wir den Code schicken:",
		"en-UK": "Please enter number of yoru friend where we'll send an invite.",
		"en-US": "Please enter number of your friend where we'll send an invite.",
		"es-ES": "Por favor, introduce el número del teléfono de tu amigo al cual enviaremos el código de la invitación.",
		"fa-IR": "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
		"fr-FR": "Veuillez entrer le numéro de votre ami où nous enverrons une invitation.",
		"id-ID": "Silakan masukkan nomor teman Anda di mana kami akan mengirimkan undangan.",
		"it-IT": "Condividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
		"ja-JP": "招待状を送信する友達の番号を入力してください。",
		"ko-KO": "초대장을 보낼 친구의 번호를 입력하세요.",
		"pl-PL": "Proszę podać numer telefonu znajomego, na który wyślemy zaproszenie.",
		"pt-BR": "Por favor, insira o número do seu amigo para onde enviaremos um convite.",
		"ru-RU": "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		"tr-TR": "Lütfen davet göndereceğimiz arkadaşınızın numarasını girin.",
		"ua-UA": "Будь ласка, введіть номер вашого друга, куди ми надішлемо запрошення.",
		"uz-UZ": "Iltimos, taklifnoma yuboradigan do'stingizning raqamini kiriting.",
		"zh-CN": "请输入我们将发送邀请的朋友的号码。",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		"de-DE": "Bitte wähle einen Kontakt, welchen du den Code schicken willst:",
		"en-UK": "Please share a contact of your friend you wish to send an invite code:",
		"en-US": "Please share a contact of your friend you wish to send an invite code:",
		"es-ES": "Por favor, comparte el contacto de tu amigo al cual quieres enviar el código de la invitación.",
		"fa-IR": "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
		"fr-FR": "Veuillez partager un contact de votre ami à qui vous souhaitez envoyer un code d'invitation :",
		"id-ID": "Silakan bagikan kontak teman Anda yang ingin Anda kirimi kode undangan:",
		"it-IT": "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
		"ja-JP": "招待コードを送信したい友達の連絡先を共有してください：",
		"ko-KO": "초대 코드를 보내고 싶은 친구의 연락처를 공유하세요:",
		"pl-PL": "Proszę udostępnić kontakt znajomego, któremu chcesz wysłać kod zaproszenia:",
		"pt-BR": "Por favor, compartilhe um contato do seu amigo para quem deseja enviar um código de convite:",
		"ru-RU": "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		"tr-TR": "Lütfen davet kodu göndermek istediğiniz arkadaşınızın bir kişisini paylaşın:",
		"ua-UA": "Будь ласка, поділіться контактом вашого друга, якому ви бажаєте надіслати код запрошення:",
		"uz-UZ": "Iltimos, taklifnoma kodini yubormoqchi bo'lgan do'stingizning kontaktini ulashing:",
		"zh-CN": "请分享您希望发送邀请码的朋友的联系人：",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		"de-DE": "Ungültige eMail Adresse. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "Invalid email. Check and try it again? /menu",
		"en-US": "Invalid email. Check and try it again? /menu",
		"es-ES": "Email incorrecto. ¿Comprobarlo e intentalo de nuevo? /menú",
		"fa-IR": "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟ /منو",
		"fr-FR": "Email invalide. Vérifiez et réessayez ? /menu",
		"id-ID": "Email tidak valid. Periksa dan coba lagi? /menu",
		"it-IT": "Email scorretta. Controlla e riprova. /menu",
		"ja-JP": "無効なメール。確認して再試行しますか？ /menu",
		"ko-KO": "잘못된 이메일입니다. 확인하고 다시 시도하시겠습니까? /menu",
		"pl-PL": "Nieprawidłowy email. Sprawdź i spróbuj ponownie? /menu",
		"pt-BR": "Email inválido. Verificar e tentar novamente? /menu",
		"ru-RU": "Неверный email. Проверить и попробовать ещё раз? /menu",
		"tr-TR": "Geçersiz e-posta. Kontrol edip tekrar deneyin? /menu",
		"ua-UA": "Недійсна електронна пошта. Перевірити та спробувати ще раз? /menu",
		"uz-UZ": "Yaroqsiz elektron pochta. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		"zh-CN": "无效的电子邮件。检查并重试？ /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		"de-DE": "Ungültiges Jahr. Der Jahresangabe sollte 2 oder 4 Ziffern sein (<i>z.B. 2016 or 16</i>).",
		"en-UK": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"en-US": "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		"es-ES": "Año incorrecto. El año tiene que constar de 2 o 4 números (<i>ejemplo 2016 o 16</i>).",
		"fa-IR": "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
		"fr-FR": "Année invalide. L'année doit comporter 2 ou 4 chiffres (<i>par exemple 2016 ou 16</i>).",
		"id-ID": "Tahun tidak valid. Bagian tahun harus berupa 2 atau 4 angka (<i>misalnya 2016 atau 16</i>).",
		"it-IT": "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
		"ja-JP": "無効な年。年の部分は2桁または4桁の数字である必要があります (<i>例：2016または16</i>)。",
		"ko-KO": "잘못된 연도입니다. 연도 부분은 2자리 또는 4자리 숫자여야 합니다 (<i>예: 2016 또는 16</i>).",
		"pl-PL": "Nieprawidłowy rok. Część roku powinna składać się z 2 lub 4 cyfr (<i>np. 2016 lub 16</i>).",
		"pt-BR": "Ano inválido. A parte do ano deve ter 2 ou 4 números (<i>por exemplo, 2016 ou 16</i>).",
		"ru-RU": "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		"tr-TR": "Geçersiz yıl. Yıl kısmı 2 veya 4 rakam olmalıdır (<i>örneğin 2016 veya 16</i>).",
		"ua-UA": "Недійсний рік. Частина року повинна складатися з 2 або 4 цифр (<i>наприклад, 2016 або 16</i>).",
		"uz-UZ": "Yaroqsiz yil. Yil qismi 2 yoki 4 raqamdan iborat bo'lishi kerak (<i>masalan, 2016 yoki 16</i>).",
		"zh-CN": "无效的年份。年份部分应为2位或4位数字 (<i>例如2016或16</i>)。",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		"de-DE": "Ungültiger Monat. Der Monatsangabe sollte eine Ganzzahl zwischen 1 und 12 sein.",
		"en-UK": "Invalid month. Month should be an integer from 1 to 12.",
		"en-US": "Invalid month. Month should be an integer from 1 to 12.",
		"es-ES": "El mes es incorrecto. El mes hay que introducirlo del 1 al 12.",
		"fa-IR": "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
		"fr-FR": "Mois invalide. Le mois doit être un nombre entier de 1 à 12.",
		"id-ID": "Bulan tidak valid. Bulan harus berupa bilangan bulat dari 1 hingga 12.",
		"it-IT": "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
		"ja-JP": "無効な月。月は1から12までの整数である必要があります。",
		"ko-KO": "잘못된 월입니다. 월은 1에서 12 사이의 정수여야 합니다.",
		"pl-PL": "Nieprawidłowy miesiąc. Miesiąc powinien być liczbą całkowitą od 1 do 12.",
		"pt-BR": "Mês inválido. O mês deve ser um número inteiro de 1 a 12.",
		"ru-RU": "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		"tr-TR": "Geçersiz ay. Ay, 1'den 12'ye kadar bir tam sayı olmalıdır.",
		"ua-UA": "Недійсний місяць. Місяць повинен бути цілим числом від 1 до 12.",
		"uz-UZ": "Yaroqsiz oy. Oy 1 dan 12 gacha bo'lgan butun son bo'lishi kerak.",
		"zh-CN": "无效的月份。月份应为1到12之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		"de-DE": "Ungültiger Tag. Der Tagesangabe sollte eine Ganzzahl zwischen 1 und 31 sein.",
		"en-UK": "Invalid day. The day should be an integer from 1 to 31.",
		"en-US": "Invalid day. The day should be an integer from 1 to 31.",
		"es-ES": "El día es incorrecto. El día hay que introducirlo del 1 al 31.",
		"fa-IR": "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
		"fr-FR": "Jour invalide. Le jour doit être un nombre entier de 1 à 31.",
		"id-ID": "Hari tidak valid. Hari harus berupa bilangan bulat dari 1 hingga 31.",
		"it-IT": "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
		"ja-JP": "無効な日。日は1から31までの整数である必要があります。",
		"ko-KO": "잘못된 일입니다. 일은 1에서 31 사이의 정수여야 합니다.",
		"pl-PL": "Nieprawidłowy dzień. Dzień powinien być liczbą całkowitą od 1 do 31.",
		"pt-BR": "Dia inválido. O dia deve ser um número inteiro de 1 a 31.",
		"ru-RU": "Неверный день. День должен быть задан целым числом от 1 до 31.",
		"tr-TR": "Geçersiz gün. Gün, 1'den 31'e kadar bir tam sayı olmalıdır.",
		"ua-UA": "Недійсний день. День повинен бути цілим числом від 1 до 31.",
		"uz-UZ": "Yaroqsiz kun. Kun 1 dan 31 gacha bo'lgan butun son bo'lishi kerak.",
		"zh-CN": "无效的日期。日期应为1到31之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		"de-DE": "Ungültiges Datumsformat. Zum Beispiel für den 20. February 2019 sende: 20.02.2019 oder 20.02.19",
		"en-UK": "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"en-US": "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		"es-ES": "El formato de la fecha no es correcto. Por ejemplo para el día 20 de Febrero de 2019 introduce: 20.02.2019 o 20.02.19",
		"fa-IR": "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
		"fr-FR": "Format de date invalide. Par exemple, pour le 20 février 2019, veuillez soumettre : 20.02.2019 ou 20.02.19",
		"id-ID": "Format tanggal tidak valid. Misalnya untuk 20 Februari 2019 silakan kirimkan: 20.02.2019 atau 20.02.19",
		"it-IT": "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
		"ja-JP": "無効な日付形式です。例えば、2019年2月20日の場合は、20.02.2019または20.02.19と入力してください。",
		"ko-KO": "잘못된 날짜 형식입니다. 예를 들어 2019년 2월 20일의 경우 다음과 같이 제출하세요: 20.02.2019 또는 20.02.19",
		"pl-PL": "Nieprawidłowy format daty. Na przykład dla 20 lutego 2019 proszę podać: 20.02.2019 lub 20.02.19",
		"pt-BR": "Formato de data inválido. Por exemplo, para 20 de fevereiro de 2019, envie: 20.02.2019 ou 20.02.19",
		"ru-RU": "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		"tr-TR": "Geçersiz tarih formatı. Örneğin 20 Şubat 2019 için lütfen şunu girin: 20.02.2019 veya 20.02.19",
		"ua-UA": "Недійсний формат дати. Наприклад, для 20 лютого 2019 року, будь ласка, введіть: 20.02.2019 або 20.02.19",
		"uz-UZ": "Yaroqsiz sana formati. Masalan, 2019 yil 20 fevral uchun quyidagilarni yuboring: 20.02.2019 yoki 20.02.19",
		"zh-CN": "无效的日期格式。例如，对于2019年2月20日，请提交：20.02.2019或20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		"de-DE": "Ungültige Telefonnummer. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "Invalid phone number. Check and try it again? /menu",
		"en-US": "Invalid phone number. Check and try it again? /menu",
		"es-ES": "El número del teléfono no es correcto. ¿Comprobarlo y intentarlo de nuevo? /menú",
		"fa-IR": "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟ /منو",
		"fr-FR": "Numéro de téléphone invalide. Vérifiez et réessayez ? /menu",
		"id-ID": "Nomor telepon tidak valid. Periksa dan coba lagi? /menu",
		"it-IT": "Numero di telefono invalido. Controlla e riprova. /menu",
		"ja-JP": "無効な電話番号です。確認して再試行しますか？ /menu",
		"ko-KO": "잘못된 전화번호입니다. 확인하고 다시 시도하시겠습니까? /menu",
		"pl-PL": "Nieprawidłowy numer telefonu. Sprawdź i spróbuj ponownie? /menu",
		"pt-BR": "Número de telefone inválido. Verificar e tentar novamente? /menu",
		"ru-RU": "Неверный номер. Проверить и попробовать ещё раз? /menu",
		"tr-TR": "Geçersiz telefon numarası. Kontrol edip tekrar deneyin? /menu",
		"ua-UA": "Недійсний номер телефону. Перевірити та спробувати ще раз? /menu",
		"uz-UZ": "Yaroqsiz telefon raqami. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		"zh-CN": "无效的电话号码。检查并重试？ /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		"de-DE": "Diese Telefonnummer kann von uns keine SMS empfangen. Versuche eine andere oder gehe ins Haupt /menu",
		"en-UK": "This phone number not able to receive SMS. Try another number? /menu",
		"en-US": "This phone number not able to receive SMS. Try another number? /menu",
		"es-ES": "Este número de teléfono no acepta SMS. ¿Intentar otro número? /menú",
		"fa-IR": "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟ /منو",
		"fr-FR": "Ce numéro de téléphone ne peut pas recevoir de SMS. Essayer un autre numéro ? /menu",
		"id-ID": "Nomor telepon ini tidak dapat menerima SMS. Coba nomor lain? /menu",
		"it-IT": "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
		"ja-JP": "この電話番号はSMSを受信できません。別の番号を試しますか？ /menu",
		"ko-KO": "이 전화번호는 SMS를 수신할 수 없습니다. 다른 번호를 시도하시겠습니까? /menu",
		"pl-PL": "Ten numer telefonu nie może odbierać SMS-ów. Spróbować innego numeru? /menu",
		"pt-BR": "Este número de telefone não pode receber SMS. Tentar outro número? /menu",
		"ru-RU": "Данный номер не принимает SMS. Попробовать другой номер? /menu",
		"tr-TR": "Bu telefon numarası SMS alamıyor. Başka bir numara deneyin? /menu",
		"ua-UA": "Цей номер телефону не може отримувати SMS. Спробувати інший номер? /menu",
		"uz-UZ": "Bu telefon raqami SMS qabul qila olmaydi. Boshqa raqamni sinab ko'rasizmi? /menu",
		"zh-CN": "此电话号码无法接收短信。尝试其他号码？ /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		"de-DE": "Wir haben keine Kontakte empfangen. Versuche es erneut oder gehe ins Haupt /menu",
		"en-UK": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"en-US": "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		"es-ES": "No hemos recibido ningún contacto. LA INSTRUCCIÓN COMO HACERLO. /menú",
		"fa-IR": "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار. /منو",
		"fr-FR": "Nous n'avons reçu aucun contact. INSTRUCTION COMMENT LE FAIRE. /menu",
		"id-ID": "Kami belum menerima kontak apa pun. INSTRUKSI CARA MELAKUKANNYA. /menu",
		"it-IT": "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
		"ja-JP": "連絡先を受信していません。操作方法の説明。 /menu",
		"ko-KO": "연락처를 받지 못했습니다. 방법 설명. /menu",
		"pl-PL": "Nie otrzymaliśmy żadnych kontaktów. INSTRUKCJA JAK TO ZROBIĆ. /menu",
		"pt-BR": "Não recebemos nenhum contato. INSTRUÇÃO COMO FAZER. /menu",
		"ru-RU": "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		"tr-TR": "Herhangi bir kişi almadık. NASIL YAPILACAĞI TALİMATI. /menu",
		"ua-UA": "Ми не отримали жодних контактів. ІНСТРУКЦІЯ ЯК ЦЕ ЗРОБИТИ. /menu",
		"uz-UZ": "Biz hech qanday kontakt olmadik. BUNI QANDAY QILISH BO'YICHA KO'RSATMA. /menu",
		"zh-CN": "我们尚未收到任何联系人。操作指南。 /menu",
	},
	MESSAGE_TEXT_YOU_HAVE_NO_CONTACTS: {
		"de-DE": "Du hast noch keine Kontakte hinzugefügt.",
		"en-UK": "You have not created any contacts yet.",
		"en-US": "You have not created any contacts yet.",
		"es-ES": "Todavía no has creado ningún contacto.", //TODO:es - verify
		"fa-IR": "هنوز هیچ مخاطبی را ایجاد نکرده اید",     //TODO:fa - verify
		"fr-FR": "Vous n'avez pas encore créé de contacts.",
		"id-ID": "Anda belum membuat kontak apa pun.",
		"it-IT": "Non hai ancora creato alcun contatto.", //TODO:it - verify
		"ja-JP": "まだ連絡先を作成していません。",
		"ko-KO": "아직 연락처를 만들지 않았습니다.",
		"pl-PL": "Nie utworzyłeś jeszcze żadnych kontaktów.",
		"pt-BR": "Você ainda não criou nenhum contato.",
		"ru-RU": "Вы ещё не создали контактов",
		"tr-TR": "Henüz hiç kişi oluşturmadınız.",
		"ua-UA": "Ви ще не створили жодних контактів.",
		"uz-UZ": "Siz hali hech qanday kontakt yaratmagansiz.",
		"zh-CN": "您尚未创建任何联系人。",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		"de-DE": "Du kannst nicht nur Zahlen für einen Kontaktnamen eingeben. Bitte benutz ein paar Buchstaben - du kannst mir vertrauen.",
		"en-UK": "You've entered just digits for a contact name. Please use some text characters.",
		"en-US": "You've entered just digits for a contact name. Please use some text characters.",
		"es-ES": "Has introducido solo números para el nombre del contacto. Por favor usa algunas letras.",
		"fa-IR": "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
		"fr-FR": "Vous avez saisi uniquement des chiffres pour un nom de contact. Veuillez utiliser des caractères textuels.",
		"id-ID": "Anda hanya memasukkan angka untuk nama kontak. Harap gunakan beberapa karakter teks.",
		"it-IT": "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
		"ja-JP": "連絡先名に数字のみを入力しました。テキスト文字を使用してください。",
		"ko-KO": "연락처 이름에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		"pl-PL": "Wprowadziłeś same cyfry jako nazwę kontaktu. Proszę użyć znaków tekstowych.",
		"pt-BR": "Você digitou apenas dígitos para um nome de contato. Por favor, use alguns caracteres de texto.",
		"ru-RU": "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		"tr-TR": "Bir kişi adı için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		"ua-UA": "Ви ввели лише цифри для імені контакту. Будь ласка, використовуйте текстові символи.",
		"uz-UZ": "Kontakt nomi uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		"zh-CN": "您只输入了数字作为联系人名称。请使用一些文本字符。",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		"de-DE": "Bei der Währung erwarte ich eigentlich keine Zahlen. Nimm ein paar Buchstaben hinzu, um Verwirrung zu vermeiden.",
		"en-UK": "You've entered just digits for currency. Please use some text characters.",
		"en-US": "You've entered just digits for currency. Please use some text characters.",
		"es-ES": "Has introducido solo números para la moneda. Por favor usa algunas letras.",
		"fa-IR": "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
		"fr-FR": "Vous avez saisi uniquement des chiffres pour la devise. Veuillez utiliser des caractères textuels.",
		"id-ID": "Anda hanya memasukkan angka untuk mata uang. Harap gunakan beberapa karakter teks.",
		"it-IT": "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
		"ja-JP": "通貨に数字のみを入力しました。テキスト文字を使用してください。",
		"ko-KO": "통화에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		"pl-PL": "Wprowadziłeś same cyfry dla waluty. Proszę użyć znaków tekstowych.",
		"pt-BR": "Você digitou apenas dígitos para moeda. Por favor, use alguns caracteres de texto.",
		"ru-RU": "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		"tr-TR": "Para birimi için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		"ua-UA": "Ви ввели лише цифри для валюти. Будь ласка, використовуйте текстові символи.",
		"uz-UZ": "Valyuta uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		"zh-CN": "您只输入了数字作为货币。请使用一些文本字符。",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		"de-DE": "%v - %s ⇒ dir: %s",
		"en-UK": "%v - %s ⇒ to you: %s",
		"en-US": "%v - %s ⇒ to you: %s",
		"es-ES": "%v - %s ⇒ a ti: %s",
		"fa-IR": "%v - %s ⇒ به شما: %s",
		"fr-FR": "%v - %s ⇒ à vous: %s",
		"id-ID": "%v - %s ⇒ kepada Anda: %s",
		"it-IT": "%v - %s ⇒ a te: %s",
		"ja-JP": "%v - %s ⇒ あなたへ: %s",
		"ko-KO": "%v - %s ⇒ 당신에게: %s",
		"pl-PL": "%v - %s ⇒ do Ciebie: %s",
		"pt-BR": "%v - %s ⇒ para você: %s",
		"ru-RU": "%v - %s ⇒ Вам : %s",
		"tr-TR": "%v - %s ⇒ size: %s",
		"ua-UA": "%v - %s ⇒ Вам: %s",
		"uz-UZ": "%v - %s ⇒ sizga: %s",
		"zh-CN": "%v - %s ⇒ 给您: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		"de-DE": "%v - Du ⇒ %s : %s",
		"en-UK": "%v - You ⇒ %s : %s",
		"en-US": "%v - You ⇒ %s : %s",
		"es-ES": "%v - Tú ⇒ %s : %s",
		"fa-IR": "%v - شما ⇒ %s : %s",
		"fr-FR": "%v - Vous ⇒ %s : %s",
		"id-ID": "%v - Anda ⇒ %s : %s",
		"it-IT": "%v - Tu ⇒ %s : %s",
		"ja-JP": "%v - あなた ⇒ %s : %s",
		"ko-KO": "%v - 당신 ⇒ %s : %s",
		"pl-PL": "%v - Ty ⇒ %s : %s",
		"pt-BR": "%v - Você ⇒ %s : %s",
		"ru-RU": "%v - Вы ⇒ %s : %s",
		"tr-TR": "%v - Siz ⇒ %s : %s",
		"ua-UA": "%v - Ви ⇒ %s : %s",
		"uz-UZ": "%v - Siz ⇒ %s : %s",
		"zh-CN": "%v - 您 ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		"de-DE": "Lass uns eine SMS senden",
		"en-UK": "Let's send SMS",
		"en-US": "Let's send SMS",
		"es-ES": "Vamos a enviar un SMS",
		"fa-IR": "پیام کوتاه ارسال کنید",
		"fr-FR": "Envoyons un SMS",
		"id-ID": "Mari kirim SMS",
		"it-IT": "Inviamo un SMS",
		"ja-JP": "SMSを送信しましょう",
		"ko-KO": "SMS를 보내봅시다",
		"pl-PL": "Wyślijmy SMS",
		"pt-BR": "Vamos enviar SMS",
		"ru-RU": "Давайте отправим SMS",
		"tr-TR": "SMS gönderelim",
		"ua-UA": "Давайте надішлемо SMS",
		"uz-UZ": "SMS yuboraylik",
		"zh-CN": "让我们发送短信",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		"de-DE": "Setze SMS in Sendewarteschlange für %v...",
		"en-UK": "Queuing SMS for sending to number %v...",
		"en-US": "Queuing SMS for sending to number %v...",
		"es-ES": "El SMS se está poniendo en la cola para enviar al número %v...",
		"fa-IR": "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
		"fr-FR": "Mise en file d'attente du SMS pour envoi au numéro %v...",
		"id-ID": "Antrian SMS untuk dikirim ke nomor %v...",
		"it-IT": "SMS in coda per l'invio al numero %v...",
		"ja-JP": "番号 %v に送信するためのSMSをキューに入れています...",
		"ko-KO": "번호 %v로 보내기 위해 SMS를 대기열에 넣는 중...",
		"pl-PL": "Kolejkowanie SMS do wysłania na numer %v...",
		"pt-BR": "Enfileirando SMS para envio ao número %v...",
		"ru-RU": "SMS ставится в очередь на отправку на номер %v...",
		"tr-TR": "%v numarasına gönderilmek üzere SMS sıraya alınıyor...",
		"ua-UA": "Ставимо SMS у чергу для надсилання на номер %v...",
		"uz-UZ": "%v raqamiga yuborish uchun SMS navbatga qo'yilmoqda...",
		"zh-CN": "正在将短信排队发送到号码 %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		"de-DE": "SMS in Sendewarteschlange für %v",
		"en-UK": "SMS is queued for sending to number %v",
		"en-US": "SMS is queued for sending to number %v",
		"es-ES": "El SMS está en la cola para enviar al número %v",
		"fa-IR": "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
		"fr-FR": "SMS mis en file d'attente pour envoi au numéro %v",
		"id-ID": "SMS diantrekan untuk dikirim ke nomor %v",
		"it-IT": "SMS inserito in coda per l'invio al numero %v",
		"ja-JP": "SMSは番号 %v に送信するためにキューに入れられています",
		"ko-KO": "SMS가 번호 %v로 보내기 위해 대기열에 있습니다",
		"pl-PL": "SMS jest w kolejce do wysłania na numer %v",
		"pt-BR": "SMS está na fila para envio ao número %v",
		"ru-RU": "SMS поставлена в очередь на отправку на номер %v",
		"tr-TR": "SMS, %v numarasına gönderilmek üzere sıraya alındı",
		"ua-UA": "SMS поставлено в чергу для надсилання на номер %v",
		"uz-UZ": "SMS %v raqamiga yuborish uchun navbatga qo'yildi",
		"zh-CN": "短信已排队发送到号码 %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		"de-DE": "Ausstehend",
		"en-UK": "Balance",
		"en-US": "Balance",
		"es-ES": "Balance",
		"fa-IR": "تراز",
		"fr-FR": "Solde",
		"id-ID": "Saldo",
		"it-IT": "Bilancio",
		"ja-JP": "残高",
		"ko-KO": "잔액",
		"pl-PL": "Saldo",
		"pt-BR": "Saldo",
		"ru-RU": "Баланс",
		"tr-TR": "Bakiye",
		"ua-UA": "Баланс",
		"uz-UZ": "Balans",
		"zh-CN": "余额",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		"de-DE": "Entschuldigung, im Moment funktionieren nur ein paar Kanäle für den Versand von Quittungen:",
		"en-UK": "Sorry, at the moment just this channels are available for sending a receipt:",
		"en-US": "Sorry, at the moment just these channels are available for sending a receipt:",
		"es-ES": "Disculpa, de momento solo estos canales están disponibles para enviar el recibo:",
		"fa-IR": "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
		"fr-FR": "Désolé, pour le moment seuls ces canaux sont disponibles pour l'envoi d'un reçu :",
		"id-ID": "Maaf, saat ini hanya saluran ini yang tersedia untuk mengirim tanda terima:",
		"it-IT": "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
		"ja-JP": "申し訳ありませんが、現時点では領収書の送信に利用できるのはこれらのチャンネルのみです：",
		"ko-KO": "죄송합니다. 현재 영수증 발송에 사용할 수 있는 채널은 다음과 같습니다:",
		"pl-PL": "Przepraszamy, w tej chwili tylko te kanały są dostępne do wysyłania paragonu:",
		"pt-BR": "Desculpe, no momento apenas estes canais estão disponíveis para enviar um recibo:",
		"ru-RU": "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		"tr-TR": "Üzgünüz, şu anda makbuz göndermek için yalnızca bu kanallar kullanılabilir:",
		"ua-UA": "Вибачте, на даний момент для надсилання квитанції доступні лише ці канали:",
		"uz-UZ": "Kechirasiz, hozirda kvitansiya yuborish uchun faqat ushbu kanallar mavjud:",
		"zh-CN": "抱歉，目前只有这些渠道可用于发送收据：",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		"de-DE": "📤 Quittung per Telegram verschickt.",
		"en-UK": "📤 Receipt sent through Telegram.",
		"en-US": "📤 Receipt sent through Telegram.",
		"es-ES": "📤 El recibo ha sido enviado vía Telegram.",
		"fa-IR": "📤 رسید از طریق تلگرام ارسال شد.",
		"fr-FR": "📤 Reçu envoyé via Telegram.",
		"id-ID": "📤 Tanda terima dikirim melalui Telegram.",
		"it-IT": "📤 Notifica inviata tramite Telegram.",
		"ja-JP": "📤 Telegramを通じて領収書が送信されました。",
		"ko-KO": "📤 텔레그램을 통해 영수증이 전송되었습니다.",
		"pl-PL": "📤 Paragon wysłany przez Telegram.",
		"pt-BR": "📤 Recibo enviado pelo Telegram.",
		"ru-RU": "📤 Квитанция отправлена через телеграм.",
		"tr-TR": "📤 Makbuz Telegram üzerinden gönderildi.",
		"ua-UA": "📤 Квитанцію надіслано через Telegram.",
		"uz-UZ": "📤 Kvitansiya Telegram orqali yuborildi.",
		"zh-CN": "📤 收据已通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		"de-DE": "Quittung konnte nicht per Telegram gesendet werden, da %v den Chat mit dem Bot gelöscht hat.",
		"en-UK": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"en-US": "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		"es-ES": "El recibo NO ha sido enviado vía Telegram porque %v ha eliminado el chat del bot.",
		"fa-IR": "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
		"fr-FR": "Reçu NON envoyé via Telegram car %v a supprimé le chat avec le bot.",
		"id-ID": "Tanda terima TIDAK dikirim melalui Telegram karena %v telah menghapus obrolan dengan bot.",
		"it-IT": "Notifica NON inviata tramite Telegram a %v perche' ha cancellato la chat con il bot",
		"ja-JP": "%v がボットとのチャットを削除したため、Telegramを通じて領収書が送信されませんでした。",
		"ko-KO": "%v가 봇과의 채팅을 삭제했기 때문에 텔레그램을 통해 영수증이 전송되지 않았습니다.",
		"pl-PL": "Paragon NIE został wysłany przez Telegram, ponieważ %v usunął czat z botem.",
		"pt-BR": "Recibo NÃO enviado pelo Telegram, pois %v excluiu o chat com o bot.",
		"ru-RU": "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		"tr-TR": "%v botla sohbeti sildiği için makbuz Telegram üzerinden GÖNDERİLMEDİ.",
		"ua-UA": "Квитанцію НЕ надіслано через Telegram, оскільки %v видалив чат з ботом.",
		"uz-UZ": "%v bot bilan suhbatni o'chirganligi sababli kvitansiya Telegram orqali YUBORILMADI.",
		"zh-CN": "由于 %v 已删除与机器人的聊天，收据未通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		"de-DE": "Quittung wurde per Mail gesendet. (id: %v)",
		"en-UK": "Receipt sent through email. (id: %v)",
		"en-US": "Receipt sent through email. (id: %v)",
		"es-ES": "El recibo ha sido enviado vía e-mail. (id: %v)",
		"fa-IR": "رسید از طریق ایمیل ارسال شد. (id: %v)",
		"fr-FR": "Reçu envoyé par e-mail. (id: %v)",
		"id-ID": "Tanda terima dikirim melalui email. (id: %v)",
		"it-IT": "Notifica inviata tramite email (id: %v)",
		"ja-JP": "領収書がメールで送信されました。(id: %v)",
		"ko-KO": "이메일을 통해 영수증이 전송되었습니다. (id: %v)",
		"pl-PL": "Paragon wysłany przez e-mail. (id: %v)",
		"pt-BR": "Recibo enviado por e-mail. (id: %v)",
		"ru-RU": "Квитанция отправлена через email. (id: %v)",
		"tr-TR": "Makbuz e-posta yoluyla gönderildi. (id: %v)",
		"ua-UA": "Квитанцію надіслано електронною поштою. (id: %v)",
		"uz-UZ": "Kvitansiya elektron pochta orqali yuborildi. (id: %v)",
		"zh-CN": "收据已通过电子邮件发送。(id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		"de-DE": "Quittung wurde per SMS gesendet.",
		"en-UK": "Receipt sent through SMS",
		"en-US": "Receipt sent through SMS",
		"es-ES": "El recibo ha sido enviado vía SMS.",
		"fa-IR": "رسید از طریق پیام کوتاه ارسال شد.",
		"fr-FR": "Reçu envoyé par SMS",
		"id-ID": "Tanda terima dikirim melalui SMS",
		"it-IT": "Notifica inviata tramite SMS",
		"ja-JP": "領収書がSMSで送信されました",
		"ko-KO": "SMS를 통해 영수증이 전송되었습니다",
		"pl-PL": "Paragon wysłany przez SMS",
		"pt-BR": "Recibo enviado por SMS",
		"ru-RU": "Квитанция отправлена через SMS.",
		"tr-TR": "Makbuz SMS yoluyla gönderildi",
		"ua-UA": "Квитанцію надіслано через SMS",
		"uz-UZ": "Kvitansiya SMS orqali yuborildi",
		"zh-CN": "收据已通过短信发送",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		"de-DE": "Schalte in den Privatmodus, um die Quittungsdetails zu sehen",
		"en-UK": "Switch to private mode to see receipt details.",
		"en-US": "Switch to private mode to see receipt details.",
		"es-ES": "Pasar al modo privado para ver el recibo.",
		"fa-IR": "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
		"fr-FR": "Passez en mode privé pour voir les détails du reçu.",
		"id-ID": "Beralih ke mode pribadi untuk melihat detail tanda terima.",
		"it-IT": "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
		"ja-JP": "領収書の詳細を確認するには、プライベートモードに切り替えてください。",
		"ko-KO": "영수증 세부 정보를 보려면 개인 모드로 전환하세요.",
		"pl-PL": "Przełącz na tryb prywatny, aby zobaczyć szczegóły paragonu.",
		"pt-BR": "Mude para o modo privado para ver os detalhes do recibo.",
		"ru-RU": "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		"tr-TR": "Makbuz ayrıntılarını görmek için özel moda geçin.",
		"ua-UA": "Перейдіть у приватний режим, щоб побачити деталі квитанції.",
		"uz-UZ": "Kvitansiya tafsilotlarini ko'rish uchun shaxsiy rejimga o'ting.",
		"zh-CN": "切换到私人模式以查看收据详情。",
	},
	MESSAGE_TEXT_RECEIPT_ATTEMPT_TO_VIEW_OWN: {
		"en-UK": "An attempt to view own receipt.",
		"ru-RU": "Попытка посмотреть свою собственную квитанцию.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		"de-DE": "👓 Quittung wurde von Gegenpartei gesichtet",
		"en-UK": "👓 Receipt viewed",
		"es-ES": "👓 El recibo ha sido visto",
		"fa-IR": "👓 رسید رویت شد.",
		"it-IT": "👓 Debiti visti",
		"ru-RU": "👓 Квитанция просмотрена",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"de-DE": "Du kannst deine eigene Nummer in dem Format anzeigen, welches wir erwarten.",
		"en-UK": "You can view your own phone number in the format we are expecting.",
		"es-ES": "Puedes ver tu número de teléfono en el formato previsto.",
		"fa-IR": "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
		"it-IT": "Puoi visionare il tuo numero di telefono nel formato previsto.",
		"ru-RU": "Вы можете посмотреть свой номер телефона в ожидаемом нами формате.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		"de-DE": "Zeige meine Nummer im erwarteten Format.",
		"en-UK": "View my number in the expectd format",
		"es-ES": "Mostrar mi número de teléfono en el formato previsto",
		"fa-IR": "رویت شماره خودم با فرمت مورد انتظار",
		"it-IT": "Mostra il mio numero nel formato previsto",
		"ru-RU": "Посмотреть мой номер в ожидаемоем формате",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		"de-DE": "Lade den ganzen Verlauf",
		"en-UK": "Show full history",
		"es-ES": "Mostrar cronología completa",
		"fa-IR": "نمایش کامل سوابق",
		"it-IT": "Mostra cronologia completa",
		"ru-RU": "Показать всю историю",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		"de-DE": "Das ist keine Nummer",
		"en-UK": "it is not a number",
		"es-ES": "No es un número",
		"fa-IR": "این یک شماره نیست",
		"it-IT": "Non e' un numero",
		"ru-RU": "Это не цифра",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		"de-DE": "Die Nummer sollte positiv sein (<i>größer als 0</i>)",
		"en-UK": "The number should be positive (<i>greater than 0</i>)",
		"es-ES": "El número tiene que ser positivo (<i>más de  0</i>)",
		"fa-IR": "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
		"it-IT": "Il numero dovrebbe essere positivo (<i>maggiore di 0</i>)",
		"ru-RU": "Цифра должна быть положительной (<i>больше нуля</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		"de-DE": "Wie viel wurde beglichen?",
		"en-UK": "How much have been returned?",
		"es-ES": "¿Cuánto/s te han devuelto?",
		"fa-IR": "چه مقدار بازپرداخت شده است؟",
		"it-IT": "Quanto ti e' stato restituito?",
		"ru-RU": "Сколько было возвращено?",
	},
	MESSAGE_TEXT_RETURN_IS_TOO_BIG: {
		"de-DE": "Sie haben entschieden, %v zurückzugeben, aber der ausstehende Betrag ist nur %v.\n\nBitte geben Sie Werte gleich %v oder weniger ein.", // TODO(DE) verify
		"en-UK": "You decided to return %v but outstanding amount is just %v.\n\nPlease enter values equal to %v or less.",
		"es-ES": "Decidiste devolver %v pero la cantidad pendiente es solo %v.\n\nPor favor ingrese valores iguales a %v o menos.", // TODO(ES) verify
		"fa-IR": "شما تصمیم گرفتید %v را بازگردانید اما مقدار قابل توجهی فقط %v است.\n\nلطفا مقادیر برابر %v یا کمتر را وارد کنید", // TODO(FA) verify
		"it-IT": "Hai deciso di restituire %v ma la quantità in sospeso è solo %v.\n\nInserisci valori pari o uguali a %v o meno.", // TODO(IT) verify
		"ru-RU": "Вы решили вернуть %v, но непогашенная сумма равна %v. \n\n Пожалуйста, введите значение равное %v или меньше.",
	},
	MESSAGE_TEXT_HELP_ROOT: {
		"de-DE": "Was hast du für eine Frage? Wenn irgendwas unklar ist, frag ruhig hier @%v",
		"en-UK": "What is your question? If anything is missed here, feel free to ask in our @%v",
		"es-ES": "¿Cuál es tu pregunta? Si algo se pierde aquí, siéntase libre de preguntar en nuestro @%v",   // TODO(es) verify
		"fa-IR": "سوالت چیست؟ اگر چیزی در اینجا از دست رفته است، لطفا در @%v ما بپرسید",                       // TODO(fa) verify
		"it-IT": "Qual è la tua domanda? Se qualche cosa è mancato qui, non esitate a chiedere al nostro @%v", // TODO(it) verify
		"ru-RU": "Какой у вас вопрос? Если здесь нет ответа пожалуйста спросите в нашей группе @%v",
	},
	MESSAGE_TEXT_HELP_BACK_TO_ROOT: {
		"de-DE": "Zurück zur FAQ Liste",
		"en-UK": "Back to FAQ list",
		"es-ES": "Volver a la lista de preguntas frecuentes", // TODO(es) verify
		"fa-IR": "بازگشت به لیست سوالات متداول",              // TODO(fa) verify
		"it-IT": "Torna all'elenco delle FAQ",                // TODO(it) verify
		"ru-RU": "Назад к списку вопросов",
	},
	HELP_HOW_TO_CREATE_BILL_Q: {
		"de-DE": "Wie erstellt man Rechnungen?",
		"en-UK": "How to create new bill?",
		"es-ES": "¿Cómo crear una nueva factura?", // TODO(es) verify
		"fa-IR": "چگونه برای ایجاد لایحه جدید؟",   // TODO(fa) verify
		"it-IT": "Come creare un nuovo conto?",    // TODO(it) verify
		"ru-RU": "Как создать новый счёт?",
	},
	HELP_HOW_TO_CREATE_BILL_A: {
		"de-DE": `<b>Wie man eine Rechnung erstellt</b>
<pre>Rechnung — Kosten, die unter zwei oder mehr Personen geteilt werden.</pre>

Deswegen kannst du am besten direkt <b>im Telegram Chat eine Rechnung erstellen, in zwei Schritten</b>:
<i>Benutze die Funktion "Teile Rechnung mit Telegram Benutzer", um es schnell zu machen:</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Füg mich einer Gruppe hinzu</a> oder öffne einen Chat mit deinem Freund.

	2. Schreibe <code>@splitusbot {amount} {bill_name}</code> und wähle ein passendes Ergebnis. Zum Beispiel:
<pre>		@splitusbot 45.5 pizza</pre>
	   Und dann kann jeder, der die Rechnung teilen will mit <code>Join</code> einsteigen.

<b>Alternativ</b> kannst du die Rechnung auch direkt in @{{.BotCode}} erstellen. Aber dann musst du die Personen, mit denen du die Rechnung teilst, einzeln hinzufügen.`,
		"en-UK": `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
		"es-ES": "", // TODO(ES)
		"fa-IR": "", // TODO(FA)
		"it-IT": "", // TODO(IT)
		"ru-RU": `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
	},
	MESSAGE_TEXT_HELP: {
		"de-DE": "Bitte melde jedes Problem und jeden Wunsch auf unserer Webseite.",
		"en-UK": "Please report any issue or submit a feature request at our website.",
		"es-ES": "Puedes informarnos de algún problema o proponernos cualquier mejora en nuestra web.",
		"fa-IR": "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
		"it-IT": "Segnala un problema o proponi un miglioramento sul nostro sito web.",
		"ru-RU": "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		"de-DE": "Support Seite",
		"en-UK": "Support page",
		"es-ES": "La página de ayuda",
		"fa-IR": "صفحه پشتیبانی",
		"it-IT": "Pagina d'aiuto",
		"ru-RU": "Cтраница поддержки ",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		"de-DE": "einen Fehler melden",
		"en-UK": "Report a bug",
		"es-ES": "Informar de un error",
		"fa-IR": "گزارش یک باگ",
		"it-IT": "Segnala un bug",
		"ru-RU": "Сообщить об ошибке",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		"de-DE": "eine Idee äußern",
		"en-UK": "Add an idea",
		"es-ES": "Proponer una idea",
		"fa-IR": "یک ایده اضافه کنید.",
		"it-IT": "Proponi un idea",
		"ru-RU": "Предложить идею",
	},
	MESSAGE_TEXT_WELCOME: {
		"de-DE": `Hallo, ich bin Collectius - dein persönlicher Buchhalter.

Ich kann dir sagen, wer wem was schuldet und wann die Schuld jeweils fällig ist.

Was würdest du gerne machen?`,
		"en-UK": `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,

		"es-ES": `Hola, me llamo Collectius, soy tu contable y asesor personal.

Puedo apuntar quién debe a quién y recordarte la fecha de devolución.

¿Qué te apetecería hacer?`,

		"fa-IR": `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,

		"it-IT": `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,

		"ru-RU": `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		"de-DE": "Ich hätte gerne einen Code",
		"en-UK": "I want to get an invite",
		"es-ES": "Me gustaría obtener una invitación",
		"fa-IR": "می خواهم یک دعوت دریافت کنم.",
		"it-IT": "Voglio ottenere un invito",
		"ru-RU": "Хочу получить приглашение",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		"de-DE": "Ich habe einen Code",
		"en-UK": "I have the invitation code",
		"es-ES": "Tengo el código de la invitación",
		"fa-IR": "من کد دعوت را دارم.",
		"it-IT": "Ho il codice invito",
		"ru-RU": "У меня есть код приглашения",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		"de-DE": "Ich habe noch keine Mails bekommen",
		"en-UK": "I have not got any emails",
		"es-ES": "No he recibido ningún e-mail",
		"fa-IR": "من ایمیلی دریافت نکردم",
		"it-IT": "Non ho ricevuto nessun email",
		"ru-RU": "Я не получил письма на email",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {

		"de-DE": `<b>%v</b>,

Im Moment ist der Bot leider nur durch Einladungen von Freunden zugänglich.

Wenn du keinen Code hast, lass deine Kontaktdaten da und wir senden dir einen Code sobald du dran bist.

Wir senden 10 Codes am Tag an die, die am längsten warten und einen zufällig.`,

		"en-UK": `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,

		"es-ES": `<b>%v</b>,

De momento nuestro bot está disponible solo por invitación de amigos.

Si no tienes el código puedes dejarnos tu contacto y te lo enviaremos cuando sea tu turno en la cola .

Enviamos 10 invitaciones por día a los primeros de la cola + 1 de modo casual.`,

		"fa-IR": `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم.

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,

		"it-IT": `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,

		"ru-RU": `<b>%v</b>,

	На данный момент наш бот доступен только тем кто получил приглашение от друзей.

	Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

	Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
	},
	EMAIL_INVITE_SUBJ: {
		"de-DE": "Eine Einladung von {{.FromName}} - Code: {{.InviteCode}}",
		"en-UK": "An invite from {{.FromName}} - code: {{.InviteCode}}",
		"es-ES": "La invitación de {{.FromName}} - el código: {{.InviteCode}}",
		"fa-IR": "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
		"it-IT": "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
		"ru-RU": "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {

		"de-DE": `Hey {{.ToName}}, {{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		"en-UK": `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,

		"es-ES": `Hola {{.ToName}}, {{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		"ru-RU": `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		"de-DE": `Hey {{.ToName}}, 

{{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		"en-UK": `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,

		"es-ES": `Hola {{.ToName}},

{{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		"fa-IR": `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		"it-IT": `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		"ru-RU": `Привет {{.ToName}},

	{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

	Ваш код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		"de-DE": `<p>Hey {{.ToName}}, </p>

<p>{{.FromName}} lädt dich ein <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">die neue Schuldentracker App auszuprobieren</a>.</p>

<p>Dein persönlicher Code lautet: <b>{{.InviteCode}}</b></p>`,
		"en-UK": `<p>Hi {{.ToName}}, </p>

<p>{{.FromName}} is inviting you to try <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,

		"es-ES": `<p>Hola {{.ToName}}, </p>

<p>{{.FromName}} te ha invitado a <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">probar la app para controlar tus deudas</a>.</p>

<p>El código de tu invitación es: <b>{{.InviteCode}}</b></p>`,

		"fa-IR": `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,

		"it-IT": `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,

		"ru-RU": `<p>Привет {{.ToName}}, </P

	<p>{{.FromName}} приглашает тебя <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

	<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		"de-DE": "Schuldschein - {{.FromName}}",
		"en-UK": "Debt record - {{.FromName}}",
		"es-ES": "La notificación de la deuda - {{.FromName}}",
		"fa-IR": "سوابق بدهی - {{.FromName}}",
		"it-IT": "Debito - {{.FromName}}",
		"ru-RU": "Запись о долге - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		"en-UK": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"es-ES": "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	MESSENGER_RECEIPT_TEXT: {
		"de-DE": "Ich habe online einen Schuldschein erstellt, für mehr Details siehe {{.ReceiptURL}}",
		"en-UK": "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		"es-ES": "He creado una notificación de la deuda, las detalles están aquí {{.ReceiptURL}}",
		"fa-IR": "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
		"it-IT": "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
		"ru-RU": "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		"en-UK": "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		"es-ES": "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		"fa-IR": "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		"it-IT": "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
		"ru-RU": "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		"de-DE": "Empfänger: %v",
		"en-UK": "Receipt: %v",
		"es-ES": "El recibo: %v",
		"fa-IR": "رسید: %v",
		"it-IT": "Debito/credito: %v",
		"ru-RU": "Квитанция: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		"de-DE": "Klicken sie hier, um die Quittung zu sehen",
		"en-UK": "Click here to send the receipt",
		"es-ES": "Haz click aquí para enviar el recibo",
		"fa-IR": "برای ارسال رسید اینجا کلیک کنید.",
		"it-IT": "Clicca qui per inviare un debito/credito",
		"ru-RU": "Нажмите здесь чтобы отправить квитанцию",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		"de-DE": "<b>Bitte wählen Sie eine Sprache, um weitere Details zu sehen, die den Schuldschein betreffen</b>, der erstellt wurde von {{.Creator}}.",
		"en-UK": "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		"es-ES": "<b>Elige el idioma para ver los detalles de la deuda</b> que ha sido creada por {{.Creator}}.",
		"fa-IR": "<b> لطفا برای رویت جزئیات بدهی که توسط </b>  {{.Creator}} ثبت شده است زبان را انتخاب کنید.",
		"it-IT": "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
		"ru-RU": "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
	},
	INLINE_RECEIPT_FOOTER: {
		//-------------------------------------------------------
		"de-DE": `{{.SiteLink}} — eine App, die dir hilft Schulden zu überwachen:

  - Du weißt immer, wie viel du allen schuldest

  - Keine Fälligkeit wird verpasst
    <i>(erinnert dich und die Gläubiger)</i>`,
		//-------------------------------------------------------
		"en-UK": `{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		//-------------------------------------------------------
		"es-ES": `{{.SiteLink}} — la app para controlar tus deudas te ayuda a:

  - Saber siempre quién debe a quién

  - Devolver la deuda a tiempo
    <i>(recordatorio a ti y a tus deudores)</i>`,
		//-------------------------------------------------------
		"fa-IR": `{{.SiteLink}} — یک برنامه پیگیری بدهی است که به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی ها به موقع پرداخت شوند.
    <i>(با ارسال یادآوری به  شما و بدهکاران )</i>`,
		//-------------------------------------------------------
		"it-IT": `{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
		//-------------------------------------------------------
		"ru-RU": `{{.SiteLink}} — программа для учёта долгов поможет:

	  - Всегда знать кто кому сколько должен

	  - Незабыть вовремя отдать или востребовать долг
	    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
	},
	INLINE_RECEIPT_GENERATING_MESSAGE: {
		"de-DE": `<b>{{.Creator}} erstellte online einen Schuldschein</b> der dich betrifft.

>> Generating receipt`, // TODO(DE)
		//-------------------------------------------------------
		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <i>Generating receipt...</i>`, // TODO(ES)
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <i>Generating receipt...</i>`, // TODO(FA)
		//-------------------------------------------------------
		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <i>Generating receipt...</i>`, // TODO(IT)
		//-------------------------------------------------------
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
	},
	//	INLINE_RECEIPT_MESSAGE: {
	//		//-------------------------------------------------------
	//		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.
	//
	//`,
	//		//-------------------------------------------------------
	//		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.
	//
	//`,
	//		//-------------------------------------------------------
	//		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.
	//
	//`,
	//		//-------------------------------------------------------
	//		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.
	//
	//`,
	//		//-------------------------------------------------------
	//		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.
	//
	//`,
	//		//-------------------------------------------------------
	//	},
	INLINE_RECEIPT_MESSAGE: {
		//-------------------------------------------------------
		"en-UK": `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"es-ES": `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"fa-IR": `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"it-IT": `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		"ru-RU": `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <a href="{{.ReceiptUrl}}">Посмотреть квитанцию</a>`,
		//-------------------------------------------------------
	},
	InlineInviteToJoinFamilyTitle: {
		"en-UK": "Invitation to join family at @%s",
		"ru-RU": "Приглашение в семью на @%s",
		"de-DE": "Einladung der Familie beizutreten bei @%s",
		"es-ES": "Invitación para unirse a la familia en @%s",
		"fr-FR": "Invitation à rejoindre la famille à @%s",
		"it-IT": "Invito a unirsi alla famiglia a @%s",
		"pt-BR": "Convite para se juntar à família em @%s",
		"fa-IR": "دعوت به پیوستن به خانواده در @%s",
	},
	InlineInviteToJoinFamilyDescription: {
		"en-UK": "Click here to send an invite",
		"ru-RU": "Нажмите здесь для отправки приглашения",
		"de-DE": "Klick hier, um eine Einladung zu versenden",
		"es-ES": "Haz click para enviar la invitación",
		"fr-FR": "Cliquez ici pour envoyer une invitation",
		"it-IT": "Clicca qui per spedire un invito",
		"pt-BR": "Clique aqui para enviar um convite",
		"fa-IR": "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
	},
	YouAreInvitedToJoinFamilyMessage: {
		"de-DE": "Sie sind eingeladen, dem Familienkonto bei @{BOT_ID} beizutreten.",
		"en-UK": "You are invited to join family account at @{BOT_ID}.",
		"es-ES": "Se te invita a unirte a la cuenta familiar en @{BOT_ID}.",
		"fa-IR": "شما دعوت شده\u200cاید که به حساب خانواده در @{BOT_ID} بپیوندید.",
		"fr-FR": "Vous êtes invité à rejoindre le compte familial sur @{BOT_ID}.",
		"id-ID": "Anda diundang untuk bergabung dengan akun keluarga di @{BOT_ID}.",
		"it-IT": "Sei invitato a unirti al conto familiare su @{BOT_ID}.",
		"ja-JP": "@{BOT_ID}で家族アカウントに参加するように招待されています。",
		"ko-KO": "@{BOT_ID}에서 가족 계정에 초대되었습니다.",
		"pl-PL": "Zostałeś zaproszony do dołączenia do konta rodzinnego na @{BOT_ID}.",
		"pt-BR": "Você foi convidado a entrar na conta familiar em @{BOT_ID}.",
		"ru-RU": "Вас пригласили присоединиться к семейному аккаунту на @{BOT_ID}.",
		"tr-TR": "@{BOT_ID} üzerinde aile hesabına katılmaya davetlisiniz.",
		"ua-UA": "Вас запросили приєднатися до сімейного акаунта на @{BOT_ID}.",
		"uz-UZ": "Sizni @{BOT_ID} oilaviy hisobiga qo‘shilishga taklif qilishdi.",
		"zh-CN": "您被邀请加入 @{BOT_ID} 的家庭账户。",
	},
	BtnViewFamilyAccount: {
		"de-DE": "Familienkonto ansehen",
		"en-UK": "View family account",
		"es-ES": "Ver cuenta familiar",
		"fa-IR": "مشاهده حساب خانواده",
		"fr-FR": "Voir le compte familial",
		"id-ID": "Lihat akun keluarga",
		"it-IT": "Visualizza conto familiare",
		"ja-JP": "家族アカウントを表示",
		"ko-KO": "가족 계정 보기",
		"pl-PL": "Zobacz konto rodzinne",
		"pt-BR": "Ver conta familiar",
		"ru-RU": "Посмотреть семейный аккаунт",
		"tr-TR": "Aile hesabını görüntüle",
		"ua-UA": "Переглянути сімейний акаунт",
		"uz-UZ": "Oilaviy hisobni ko‘rish",
		"zh-CN": "查看家庭账户",
	},
	SMS_RECEIPT_YOU_GOT: {
		"de-DE": "Du hast dir %v von %v geliehen.",
		"en-UK": "You've got %v from %v.",
		"es-ES": "Has recibido %v de %v.",
		"fa-IR": "شما دریافت کرده اید %v از %v.",
		"it-IT": "Hai ricevuto %v da %v",
		"ru-RU": "Вы получили %v от %v.",
	},
	SMS_RECEIPT_YOU_GAVE: {
		"de-DE": "Du hast %v an %v verliehen.",
		"en-UK": "You've given %v to %v.",
		"es-ES": "Has dado %v a %v.",
		"fa-IR": "شما پرداخت کرده اید %v به %v.",
		"it-IT": "Hai dato %v a %v",
		"ru-RU": "Вы дали %v - взял(а) %v.",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		"de-DE": "Klicke %v um zu akzeptieren oder abzulehnen.",
		"en-UK": "Click %v to confirm or decline.",
		"es-ES": "Haz click %v para confirmar o rechazar.",
		"fa-IR": "کلیک کنید %v تا رد خود را تایید نمایید",
		"it-IT": "Clicca %v per confermare o rifiutare",
		"ru-RU": "Нажмите %v чтобы подтвердить или отвергнуть.",
	},
	HTML_DATE: {
		"de-DE": "Datum",
		"en-UK": "Date",
		"es-ES": "Fecha",
		"fa-IR": "تاریخ",
		"it-IT": "Data",
		"ru-RU": "Дата",
	},
	HTML_RECEIPT: {
		"de-DE": "Empfänger",
		"en-UK": "Receipt",
		"es-ES": "Recibo",
		"fa-IR": "رسید",
		"it-IT": "Scontrino", //To upgrade, not the best translation from Russian
		"ru-RU": "Квитанция",
	},
	HTML_AMOUNT: {
		"de-DE": "Betrag",
		"en-UK": "Amount",
		"es-ES": "Importe",
		"fa-IR": "مقدار",
		"it-IT": "Totale",
		"ru-RU": "Сумма",
	},
	HTML_FROM: {
		"de-DE": "Von",
		"en-UK": "From",
		"es-ES": "De",
		"fa-IR": "از",
		"it-IT": "Da",
		"ru-RU": "Дал",
	},
	HTML_TO: {
		"de-DE": "An",
		"en-UK": "To",
		"es-ES": "Para",
		"fa-IR": "به",
		"it-IT": "Per",
		"ru-RU": "Получил",
	},
	NO_NAME: {
		"de-DE": "unbekannt",
		"en-UK": "no name",
		"es-ES": "sin nombre", // TODO(es) verify
		"fa-IR": "بدون نام",   // TODO(fa) verify
		"it-IT": "senza nome", // TODO(it) verify
		"ru-RU": "без имени",
	},
	TELEGRAM_RECEIPT: {
		"de-DE": "{{.FromName}} hat einen Schuldschein erstellt ({{.TransferCurrency}})",
		"en-UK": "{{.FromName}} created a debt record ({{.TransferCurrency}})",
		"es-ES": "{{.FromName}} ha creado una deuda ({{.TransferCurrency}})",
		"fa-IR": "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
		"it-IT": "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
		"ru-RU": "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		"de-DE": "Bitte wähle aus den angezeigten Optionen.",
		"en-UK": "Please choose from provided options.",
		"es-ES": "Por favor, elige una de las siguientes opciones.",
		"fa-IR": "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
		"it-IT": "Scegli tra le opzioni fornite.",
		"ru-RU": "Пожалуйста выберете из предоставленных опций.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		"de-DE": "<b>Möchtest du eine Bemerkung oder Notiz hinzufügen?</b>\n%v Deine Notizen kannst nur du sehen.\n%v Eine Bemerkung wird quasi auf dem Schuldschein und der Quittung vermerkt und ist insofern für beide sichtbar.",
		"en-UK": "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		"es-ES": "<b>¿Quieres añadir una nota o comentario?</b>\n%v Las notas se graban de manera privada para tu propia información.\n%v Los comentarios son visibles para todos los autorizados a ver esta transacción.",
		"fa-IR": "<b>آیا میخواهید یادداشت یا شرحی اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v شرح در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
		"it-IT": "<b>Vuoi aggiungere una nota o un commento?</b> \n%v I memo sono record privati per il riferimento di yoru.\n%v I commenti sono disponibili a tutti coloro che hanno l'autorizzazione a visualizzare questa transazione.",
		"ru-RU": "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личного пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		"de-DE": "Bitte schreibe eine Notiz:",
		"en-UK": "Please write a note:",
		"es-ES": "Por favor, escribe una nota:",
		"fa-IR": "لطفاً یک یادداشت بنویسید:",
		"it-IT": "Per favore scrivi un appunto:",
		"ru-RU": "Напишите заметку:",
	},
	COMMAND_TEXT_MORE_ABOUT_INTEREST_COMMAND: {
		"de-DE": "Mehr über Prozentsätze", // TODO(DE)
		"en-UK": "More about interest",
		"es-ES": "Más sobre interés",   // TODO(ES)
		"fa-IR": "بیشتر در مورد علاقه", // TODO(FA)
		"it-IT": "Dimmi di più",        // TODO(IT)
		"ru-RU": "Подробнее о процентах",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_SHORT: {
		"de-DE": `<b>Prozent und Kommentar</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(DE)
		"en-UK": `<b>Interest & notes</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`,
		"es-ES": `<b>Porcentaje y comentario</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(ES)
		"fa-IR": `<b>درصد و نظر</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(FA)
		"it-IT": `<b>Percentuale e commento</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`, // TODO(IT
		"ru-RU": `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:
	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_LONG: {
		"de-DE": ``, // TODO(DE)
		"en-UK": `<b>Interest & notes</b>

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

❗ The % functionality is in BETA testing stage, please let us know in @DebtsTrackerGroup if anything works not as you would expect.`, // TODO - replace link!
		"es-ES": ``, // TODO(ES)
		"fa-IR": ``, // TODO(FA)
		"it-IT": ``, // TODO(IT
		"ru-RU": `<b>Процент и комментарий</b>

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
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT_ONLY: {
		"de-DE": "Falls du eine Bemerkung auf den Schuldschein schreiben willst, schick mir jetzt den Text.",
		"en-UK": `If you want to add a comment just send a text now.`,
		"es-ES": `si quieres añadir un comentario simplemente envia un texto.`,
		"fa-IR": `شما می توانید یک شرح اضافه کنید. تنها کافیست یک متن ارسال کنید.`,
		"it-IT": `Se vuoi aggiungere un commento invia del testo ora.`,
		"ru-RU": `Если хотите добавить комментарий просто отправьте текст.`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		"de-DE": "sichtbar für dich & %v",
		"en-UK": "visible to you & %v",
		"es-ES": "visible solo para ti & %v",
		"fa-IR": "قابل مشاهده برای شما & %v",
		"it-IT": "visibile solo a te e %v",
		"ru-RU": "виден вам и %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		"de-DE": "Schreibe nun eine Bemerkung auf den Schuldschein:",
		"en-UK": "Please write the comment:",
		"es-ES": "Por favor, escribe un comentario:",
		"fa-IR": "لطفاً شرح را ثبت کنید:",
		"it-IT": "Per favore scrivi un commento:",
		"ru-RU": "Напишите комментарий:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		"de-DE": "Deine Notiz wurde hinzugefügt, möchtest du noch eine Bemerkung auf den Schuldschein schreiben?",
		"en-UK": "Memo have been added. Do you want to write a comment?",
		"es-ES": "La nota está añadida. ¿Quieres escribir un comentario?",
		"fa-IR": "یادداشت اضافه شد. آیا میخواهید یک شرح ثبت کنید؟",
		"it-IT": "Promemoria aggiunto. Vuoi scrivere un commento?",
		"ru-RU": "Заметка добавлена. Хотите написать комментарий?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		"de-DE": "Die Bemerkung wurde vermerkt. Möchtest du noch eine Notiz für dich hinzufügen?",
		"en-UK": "Comment have been added. Do you want to write a note?",
		"es-ES": "El comentario está añadido. ¿Quieres escribir una nota?",
		"fa-IR": "شرح موردنظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
		"it-IT": "Commento aggiunto. Vuoi scrivere un appunto?",
		"ru-RU": "Комментарий добавлен. Хотите написать заметку?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		"de-DE": "Keine Notizen oder Bemerkungen",
		"en-UK": "Without notes or comments",
		"es-ES": "Sin notas ni comentarios",
		"fa-IR": "بدون یادداشت یا شرح",
		"it-IT": "Senza appunti o commenti",
		"ru-RU": "Без заметок и комментариев",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		"de-DE": "Keine Bemerkungen",
		"en-UK": "No comments",
		"es-ES": "Sin comentarios",
		"fa-IR": "بدون شرح",
		"it-IT": "Nessun commento",
		"ru-RU": "Без комментариев",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		"de-DE": "Keine Notizen",
		"en-UK": "Without notes",
		"es-ES": "Sin notas",
		"fa-IR": "بدون یادداشت",
		"it-IT": "Senza appunti/note",
		"ru-RU": "Без заметок",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		"de-DE": "Notiz hinzufügen (privat)",
		"en-UK": "Add a note (private)",
		"es-ES": "Añadir una nota (privada)",
		"fa-IR": "یک یادداشت اضافه کنید (خصوصی)",
		"it-IT": "Aggiungi una nota (privata)",
		"ru-RU": "Добавить заметку",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		"de-DE": "Bemerkung hinzufügen (öffentlich)",
		"en-UK": "Add a comment (public)",
		"es-ES": "Añadir un comentario (público)",
		"fa-IR": "یک شرح اضافه کنید (عمومی)",
		"it-IT": "Aggiungi un commento (pubblico)",
		"ru-RU": "Добавить комментарий",
	},
	DUE_IN_NOW: {
		"de-DE": "jetzt",
		"en-UK": "now",
		"es-ES": "ahora",
		"fa-IR": "حالا",
		"it-IT": "adesso",
		"ru-RU": "прямо сейчас",
	},
	DUE_IN_A_MINUTE: {
		"de-DE": "in einer Minute",
		"en-UK": "in a minute",
		"es-ES": "en un minuto",
		"fa-IR": "ظرف یک دقیقه",
		"it-IT": "in un minuto",
		"ru-RU": "через минуту",
	},
	DUE_IN_X_MINUTES: {
		"de-DE": "in %v Minuten",
		"en-UK": "in %v minutes",
		"es-ES": "en %v minutos",
		"fa-IR": "در %v دقیقه",
		"it-IT": "in %v minuti/o",
		"ru-RU": "через %v минут(ы)",
	},
	DUE_IN_AN_HOUR: {
		"de-DE": "in einer Stunde",
		"en-UK": "in an hour",
		"es-ES": "en  una hora",
		"fa-IR": "ظرف یک ساعت",
		"it-IT": "in un ora",
		"ru-RU": "через час",
	},
	DUE_IN_X_HOURS: {
		"de-DE": "in %v Stunde",
		"en-UK": "in %v hours",
		"es-ES": "en %v horas",
		"fa-IR": "در %v ساعت",
		"it-IT": "in %v ore/a",
		"ru-RU": "через %v часа/часов",
	},
	DUE_IN_X_DAYS: {
		"de-DE": "in %v Tagen",
		"en-UK": "in %v days",
		"es-ES": "en %v días",
		"fa-IR": "در %v روز",
		"it-IT": "in %v giorni/o",
		"ru-RU": "через %v дня/дней",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		"de-DE": "Alexander Trakhimenok",
		"en-UK": "Alexander Trakhimenok",
		"es-ES": "Alexander Trakhimenok",
		"fa-IR": "الکساندر تراخیمِنوک",
		"it-IT": "Alexander Trakhimenok",
		"ru-RU": "Александр Трахимёнок",
	},

	WS_INDEX_TITLE: {
		"ru-RU": "DebtsTracker.io - программа для учёта личных долгов и активов",
		"en-UK": "DebtsTracker.io - an IOU app to track your personal debts & assets",
		"es-ES": "DebtsTracker.io es una aplicación para el control de sus deudas personales",
		"fa-IR": "DebtsTracker.io - برنامه ای برای ردیابی بدهی ها و دارایی های شما",
		"pl-PL": "DebtsTracker.io - aplikacja do śledzenia osobistych długów",
		"pt-PT": "DebtsTracker.io - um aplicativo para controlar suas dívidas pessoais",
		"de-DE": "DebtsTracker.io - eine App, um Ihre persönlichen Schulden zu verfolgen",
		"fr-FR": "DebtsTracker.io - une application pour suivre vos dettes personnelles",
		"it-IT": "DebtsTracker.io - un app per monitorare i tuoi debiti personali",
		"ko-KO": "DebtsTracker.io 은 - 앱이 사용자의 개인 채무를 추적",
		"ja-JP": "DebtsTracker.io は - アプリはあなたの個人的な借金を追跡します",
		"zh-CN": "DebtsTracker.io - 一个应用程序来跟踪你的个人债务",
	},
	WS_LIVE_DEMO: {
		"ru-RU": "Демо версия online",
		"en-UK": "Demostración",
		"es-ES": "Demo en vivo",
		"fa-IR": "نسخه ی نمایشی زنده",
		"pl-PL": "Demo na żywo",
		"pt-PT": "Demonstração ao vivo",
		"de-DE": "Live-Demo",
		"fr-FR": "Démo en direct",
		"it-IT": "Demo online",
		"ko-KO": "실시간 데모",
		"ja-JP": "ライブデモ",
		"zh-CN": "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		"ru-RU": "Бот для Telegram",
		"en-UK": "Chat bot for Telegram messenger",
		"es-ES": "Chat bot para Telegram",
		"fa-IR": "ربات چت برای پیام رسان تلگرام",
		"pl-PL": "Chat bot do telegramu posłańca",
		"pt-PT": "bot de bate-papo para Telegram messenger",
		"de-DE": "Chat-Bot für Telegram",
		"fr-FR": "bot Chat for Telegram messenger",
		"it-IT": "Bot Chat per Telegram",
		"ko-KO": "전보 메신저 채팅 봇",
		"ja-JP": "電報メッセンジャーのためのチャットボット",
		"zh-CN": "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		"ru-RU": "Открыть в Телеграмме &#x1F680;",
		"en-UK": "Open in Telegram &#x1F680;",
		"es-ES": "Abrir en Telegram &#x1F680;",
		"fa-IR": "بازکردن در تلگرام &#x1F680;",
		"pl-PL": "Otwórz w telegramu &#x1F680;",
		"pt-PT": "Open in Telegram &#x1F680;",
		"de-DE": "Öffnen in Telegram &#x1F680;",
		"fr-FR": "Open in Telegram &#x1F680;",
		"it-IT": "Apri su Telegram &#x1F680;",
		"ko-KO": "전보 &#x1F680; 에서 열기;",
		"ja-JP": "電報 &#x1F680; で開きます。",
		"zh-CN": "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		"ru-RU": "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		"en-UK": "At the moment our program is available just on <a href='https://telegram.org/'>Telegram messenger</a>",
		"es-ES": "De momento nuestro programa está disponible sólo en <a href='https://telegram.org/'>Telegrama mensajero </a>",
		"fa-IR": "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>Телеграм</a>تلگرام",
		"pl-PL": "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'>Telegram messenger</a>",
		"pt-PT": "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'>Telegram messenger</a>",
		"de-DE": "Im Moment ist unser Programm nur auf <a href='https://telegram.org/'>Telegram verfügbar</a>",
		"fr-FR": "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'>Telegram messager</a>",
		"it-IT": "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'>Telegram</a>",
		"ko-KO": "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>Telegram</a> 의 <b> 전보 </b>을 메신저 를 볼 수 있습니다",
		"ja-JP": "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'>Telegram</a>電報のメッセンジャーで提供されています",
		"zh-CN": "目前我们的计划是只提供在<a href='https://telegram.org/'>Telegram</a>电报的使者",
	},
	WS_MOTTO: {
		"ru-RU": "Платежи по долгам целиком и вовремя!",
		"en-UK": "Know your bottom line & never miss a debt payment!",
		"es-ES": "¡Controle sus pagos y deudas!",
		"fa-IR": "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
		"pl-PL": "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		"pt-PT": "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		"de-DE": "Wissen, wem man wie viel schuldet!",
		"fr-FR": "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		"it-IT": "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		"ko-KO": "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		"ja-JP": "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		"zh-CN": "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		"ru-RU": "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		"en-UK": "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		"es-ES": "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorios que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		"fa-IR": "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		"pl-PL": "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		"pt-PT": "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		"de-DE": "DebtsTracker.io ist eine mobile App, die beim Verwalten von persönlichen Schulden hilft - egal ob Sie Geld verleihen oder welches leihen. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner und Gläubiger.",
		"fr-FR": "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		"it-IT": "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		"ko-KO": "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		"ja-JP": "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		"zh-CN": "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		"de-DE": "Wie kann man beim DebtsTracker.io Projekt helfen kann",
		"en-UK": "How you can help to DebtsTracker.io project",
		"es-ES": "Como puedes ayudar a DebtsTracker.io project",
		"fa-IR": "چگونه می توانید به پروژه  DebtsTracker.io کمک کنید.",
		"it-IT": "Come potete aiutare il progetto DebtsTracker.io", // TODO(IT): Google translated
		"ru-RU": "Как вы можете помочь проекту DebtsTracker.io",
	},
	WS_ADS_TITLE: {
		"de-DE": "Werbung @ DebtsTracker.IO",
		"en-UK": "Ads @ DebtsTracker.IO",
		"es-ES": "Anuncio @ DebtsTracker.IO",
		"fa-IR": "تبلیغات @ DebtsTracker.IO",
		"it-IT": "Pubblicita' @ DebtsTracker.IO",
		"ru-RU": "Реклама на DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		"de-DE": `Um Werbung in unserer App zu schalten, schick uns eine Mail an <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"en-UK": `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"es-ES": `Para publicar un anuncio en nuestra app escríbenos un e-mail a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"fa-IR": `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"it-IT": `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		"ru-RU": `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		"de-DE": "Schuldschein",
		"en-UK": "Record debt",
		"es-ES": "Registrar la deuda",
		"fa-IR": "ثبت بدهی",
		"it-IT": "Registra il debito",
		"ru-RU": "Записать долг",
	},
	FB_MAKE_RECORD_HEADLINE: {
		"de-DE": "Scroll nach links, um weitere Optionen zu sehen.",
		"en-UK": "Scroll left to see other options.",
		"es-ES": "Desliza a la izquierda para ver otras opciones",
		"fa-IR": "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
		"it-IT": "Scrolla a sinistra per vedere altre opzioni",
		"ru-RU": "Пролистайте карточки влево чтобы увидеть другие опции.",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		"de-DE": "Wie geht es dir?",
		"en-UK": "How are you doing?",
		"es-ES": "¿Cómo va todo?",
		"fa-IR": "حال شما چطوره؟",
		"it-IT": "Come te la passi?",
		"ru-RU": "Как идут дела?",
	},
	SNEATBOT_MSG_TXT_START: {
		"en-UK": `
<b>From bot's creator</b>: Hi %s!

@SneatBot helps to manage your day-to-day family life. Or you can create a space to manage your group/team/community.

I've spend lot's of time to make this bot useful, quick & reliable.I hope you'll like it.

You can learn about new features of the bot in @SneatApp channel where <a href="https://t.me/StarGiveaways_EN">we giveaway 500 🌟 EVERY month</a>.
`,
		"ru-RU": `
<b>От создателя бота:</b> Привет %s!

@SneatBot помогает организовать вашу семейную жизнь.
Так же можно создать пространство для управления группой/командой/сообществом.

Вы можете узнавать о новых возможностях бота в канале @SneatApp_ru где <a href="https://t.me/StarGiveaways_RU">мы разыгрываем 500 🌟 КАЖДЫЙ месяц</a>.
`,
		"ko-KO": `<b>봇 제작자로부터</b>: 안녕하세요 %s!

@SneatBot은 일상의 가족 생활을 관리하는 데 도움이 됩니다. 또는 그룹/팀/커뮤니티를 관리할 공간을 만들 수 있습니다.

저는 이 봇을 유용하고 빠르고 안정적으로 만들기 위해 많은 시간을 투자했습니다. 마음에 들어 하시기를 바랍니다.

@SneatApp 채널에서 봇의 새로운 기능에 대해 알아볼 수 있습니다. <a href="https://t.me/StarGiveaways_EN">매월 500🌟를 경품으로 드립니다</a>.`,
		"es-ES": `<b>Del creador del bot:</b> ¡Hola, %s!

@SneatBot te ayuda a gestionar tu vida familiar diaria. O puedes crear un espacio para gestionar tu grupo/equipo/comunidad.

He dedicado mucho tiempo a hacer que este bot sea útil, rápido y fiable. Espero que te guste.

Puedes conocer las nuevas funciones del bot en el canal de @SneatApp, donde <a href="https://t.me/StarGiveaways_EN">regalamos 500 🌟 CADA mes</a>.`,
		"fr-FR": `<b>De la part du créateur du bot</b> : Bonjour %s !

@SneatBot vous aide à gérer votre vie de famille au quotidien. Vous pouvez également créer un espace pour gérer votre groupe/équipe/communauté.

J'ai passé beaucoup de temps à rendre ce bot utile, rapide et fiable. J'espère qu'il vous plaira.

Vous pouvez en apprendre davantage sur les nouvelles fonctionnalités du bot sur le canal @SneatApp où <a href="https://t.me/StarGiveaways_FR">nous offrons 500 🌟 CHAQUE mois</a>.`,
		"it-IT": `<b>Dal creatore del bot</b>: Ciao %s!

@SneatBot ti aiuta a gestire la tua vita familiare quotidiana. Oppure puoi creare uno spazio per gestire il tuo gruppo/team/community.

Ho dedicato molto tempo a rendere questo bot utile, veloce e affidabile. Spero che ti piaccia.

Puoi scoprire le nuove funzionalità del bot nel canale @SneatApp dove <a href="https://t.me/StarGiveaways_EN">regaliamo 500 🌟 OGNI mese</a>.`,
		"ja-JP": `<b>ボットの作成者より</b>: こんにちは、%s!

@SneatBot は、日々の家族生活の管理に役立ちます。または、グループ/チーム/コミュニティを管理するスペースを作成することもできます。

このボットを便利で、迅速で、信頼できるものにするために、多くの時間を費やしました。気に入っていただければ幸いです。

ボットの新機能については、@SneatApp チャンネルで確認できます。ここでは、<a href="https://t.me/StarGiveaways_EN">毎月 500 🌟 をプレゼント</a>しています。`,
		"zh-CN": `<b>来自机器人的创建者</b>：嗨 %s！

@SneatBot 帮助您管理日常家庭生活。或者您可以创建一个空间来管理您的群组/团队/社区。

我花了很多时间让这个机器人变得有用、快速和可靠。希望你会喜欢它。

您可以在 @SneatApp 频道了解该机器人的新功能，<a href="https://t.me/StarGiveaways_EN">我们每月赠送 500 🌟</a>。`,
		"de-DE": `<b>Vom Ersteller des Bots</b>: Hallo %s!

@SneatBot hilft Ihnen, Ihren Familienalltag zu organisieren. Oder Sie können einen Bereich erstellen, in dem Sie Ihre Gruppe/Ihr Team/Ihre Community verwalten können.

Ich habe viel Zeit darauf verwendet, diesen Bot nützlich, schnell und zuverlässig zu machen. Ich hoffe, er gefällt Ihnen.

Sie können sich über neue Funktionen des Bots im @SneatApp-Kanal informieren, wo wir <a href="https://t.me/StarGiveaways_EN">JEDEN Monat 500 🌟 verschenken</a>.`,
		"pt-PT": `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerir a sua vida familiar quotidiana. Ou pode criar um espaço para gerir o seu grupo/equipa/comunidade.

Gastei muito tempo para tornar este bot útil, rápido e fiável.

Pode conhecer as novas funcionalidades do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">distribuímos 500 🌟 TODOS os meses</a>.`,
		"pt-BR": `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerenciar sua vida familiar cotidiana. Ou você pode criar um espaço para gerenciar seu grupo/equipe/comunidade.

Eu gastei muito tempo para tornar este bot útil, rápido e confiável. Espero que você goste.

Você pode aprender sobre os novos recursos do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">nós doamos 500 🌟 TODO mês</a>.`,
		"pl-PL": `<b>Od twórcy bota</b>: Cześć %s!

@SneatBot pomaga zarządzać codziennym życiem rodzinnym. Możesz też utworzyć przestrzeń do zarządzania swoją grupą/zespołem/społecznością.

Poświęciłem dużo czasu, aby ten bot był użyteczny, szybki i niezawodny. Mam nadzieję, że Ci się spodoba.

Możesz dowiedzieć się o nowych funkcjach bota na kanale @SneatApp, gdzie <a href="https://t.me/StarGiveaways_EN">rozdajemy 500 🌟 CO MIESIĄC</a>.`,
		"ua-UA": `<b>Від творця бота</b>: Привіт, %s!

@SneatBot допомагає керувати повсякденним сімейним життям. Або ви можете створити простір для керування своєю групою/командою/спільнотою.

Я витратив багато часу, щоб зробити цього бота корисним, швидким і надійним. Сподіваюся, він вам сподобається.

Ви можете дізнатися про нові функції бота на каналі @SneatApp, де <a href="https://t.me/StarGiveaways_EN">ми роздаємо 500 🌟 КОЖНОГО місяця</a>.`,
	},
	SpaceCmdText: {
		"en-UK": "Current space: %s <b>%s</b>",
		"ru-RU": "Текущее пространство: %s <b>%s</b>",
		"es-ES": "Espacio actual: %s <b>%s</b>",
		"fa-IR": "فضای فعلی: %s <b>%s</b>",
		"it-IT": "Spazio corrente: %s <b>%s</b>",
		"de-DE": "Aktueller Raum: %s <b>%s</b>",
		"fr-FR": "Espace actuel: %s <b>%s</b>",
		"pl-PL": "Aktualna przestrzeń: %s <b>%s</b>",
		"pt-PT": "Espaço atual: %s <b>%s</b>",
		"ko-KO": "현재 공간: %s <b>%s</b>",
		"ja-JP": "現在のスペース: %s <b>%s</b>",
		"zh-CN": "当前空间: %s <b>%s</b>",
		"ua-UA": "Поточний простір: %s <b>%s</b>",
		"pt-BR": "Espaço atual: %s <b>%s</b>",
		"tr-TR": "Mevcut alan: %s <b>%s</b>",
		"id-ID": "Ruang saat ini: %s <b>%s</b>",
	},
	SpaceCmdBtnContacts: {
		"en-UK": "Contacts",
		"ru-RU": "Контакты",
		"es-ES": "Contactos",
		"fa-IR": "مخاطبین",
		"it-IT": "Contatti",
		"de-DE": "Kontakte",
		"fr-FR": "Contacts",
		"pl-PL": "Kontakty",
		"pt-PT": "Contatos",
		"ko-KO": "연락처",
		"ja-JP": "連絡先",
		"zh-CN": "联系人",
		"ua-UA": "Контакти",
		"pt-BR": "Contatos",
		"tr-TR": "Kişiler",
		"id-ID": "Kontak",
	},
	SpaceCmdBtnMembers: {
		"en-UK": "Members",
		"ru-RU": "Участники",
		"es-ES": "Miembros",
		"fa-IR": "اعضا",
		"it-IT": "Membri",
		"de-DE": "Mitglieder",
		"fr-FR": "Membres",
		"pl-PL": "Członkowie",
		"pt-PT": "Membros",
		"ko-KO": "회원",
		"ja-JP": "メンバー",
		"zh-CN": "成员",
		"ua-UA": "Учасники",
		"pt-BR": "Membros",
		"tr-TR": "Üyeler",
		"id-ID": "Anggota",
	},
	FamilyMembers: {
		"en-UK": "Family members",
		"ru-RU": "Члены семьи",
		"es-ES": "Miembros de la familia",
		"fa-IR": "اعضای خانواده",
		"it-IT": "Membri della famiglia",
		"de-DE": "Familienmitglieder",
		"fr-FR": "Membres de la famille",
		"pl-PL": "Członkowie rodziny",
		"pt-PT": "Membros da família",
		"ko-KO": "가족 구성원",
		"ja-JP": "家族のメンバー",
		"zh-CN": "家庭成员",
		"ua-UA": "Члени сім'ї",
		"pt-BR": "Membros da família",
		"tr-TR": "Aile üyeleri",
		"id-ID": "Anggota keluarga",
	},
	SpaceMembers: {
		"en-UK": "Space members",
		"ru-RU": "Участники пространства",
		"es-ES": "Miembros del espacio",
		"fa-IR": "اعضای فضا",
		"it-IT": "Membri dello spazio",
		"de-DE": "Raummitglieder",
		"fr-FR": "Membres de l'espace",
		"pl-PL": "Członkowie przestrzeni",
		"pt-PT": "Membros do espaço",
		"ko-KO": "공간 멤버",
		"ja-JP": "スペースメンバー",
		"zh-CN": "空间成员",
		"ua-UA": "Члени простору",
		"pt-BR": "Membros do espaço",
		"tr-TR": "Alan üyeleri",
		"id-ID": "Anggota ruang",
	},
	SpaceCmdBtnLists: {
		"en-UK": "Lists",
		"ru-RU": "Списки",
		"es-ES": "Listas",
		"fa-IR": "لیست ها",
		"it-IT": "Elenchi",
		"de-DE": "Listen",
		"fr-FR": "Listes",
		"pl-PL": "Listy",
		"pt-PT": "Listas",
		"ko-KO": "목록",
		"ja-JP": "リスト",
		"zh-CN": "清单",
		"ua-UA": "Списки",
		"pt-BR": "Listas",
		"tr-TR": "Listeler",
		"id-ID": "Daftar",
	},
	SpaceCmdBtnAssets: {
		"en-UK": "Assets",
		"ru-RU": "Активы",
		"es-ES": "Activos",
		"fa-IR": "دارایی ها",
		"it-IT": "Attività",
		"de-DE": "Vermögenswerte",
		"fr-FR": "Actifs",
		"pl-PL": "Aktywa",
		"pt-PT": "Ativos",
		"ko-KO": "자산",
		"ja-JP": "資産",
		"zh-CN": "资产",
		"ua-UA": "Активи",
		"pt-BR": "Ativos",
		"tr-TR": "Varlıklar",
		"id-ID": "Aset",
	},
	SpaceCmdBtnBudget: {
		"en-UK": "Budget",
		"ru-RU": "Бюджет",
		"es-ES": "Presupuesto",
		"fa-IR": "بودجه",
		"it-IT": "Budget",
		"de-DE": "Budget",
		"fr-FR": "Budget",
		"pl-PL": "Budżet",
		"pt-PT": "Orçamento",
		"ko-KO": "예산",
		"ja-JP": "予算",
		"zh-CN": "预算",
		"ua-UA": "Бюджет",
		"pt-BR": "Orçamento",
		"tr-TR": "Bütçe",
		"id-ID": "Anggaran",
	},
	SpaceCmdBtnDebts: {
		"en-UK": "Debts",
		"ru-RU": "Долги",
		"es-ES": "Deudas",
		"fa-IR": "بدهی ها",
		"it-IT": "Debiti",
		"de-DE": "Schulden",
		"fr-FR": "Dettes",
		"pl-PL": "Długi",
		"pt-PT": "Dívidas",
		"ko-KO": "부채",
		"ja-JP": "借金",
		"zh-CN": "债务",
		"ua-UA": "Борги",
		"pt-BR": "Dívidas",
		"tr-TR": "Borçlar",
		"id-ID": "Hutang",
	},
	SpaceCmdBtnCalendar: {
		"en-UK": "Calendar",
		"ru-RU": "Календарь",
		"es-ES": "Calendario",
		"fa-IR": "تقویم",
		"it-IT": "Calendario",
		"de-DE": "Kalender",
		"fr-FR": "Calendrier",
		"pl-PL": "Kalendarz",
		"pt-PT": "Calendário",
		"ko-KO": "달력",
		"ja-JP": "カレンダー",
		"zh-CN": "日历",
		"ua-UA": "Календар",
		"pt-BR": "Calendário",
		"tr-TR": "Takvim",
		"id-ID": "Kalender",
	},
	SpaceCmdBtnTrackers: {
		"de-DE": "Tracker", // Placeholder
		"en-UK": "Trackers",
		"en-US": "Trackers",      // Placeholder
		"es-ES": "Rastreadores",  // Placeholder
		"fa-IR": "ردیاب\u200cها", // Placeholder
		"fr-FR": "Traqueurs",     // Placeholder
		"id-ID": "Pelacak",       // Placeholder
		"it-IT": "Tracker",       // Placeholder
		"ja-JP": "トラッカー",         // Placeholder
		"ko-KO": "추적기",           // Placeholder
		"pl-PL": "Monitorujące",  // Placeholder
		"pt-BR": "Rastreadores",  // Placeholder
		"ru-RU": "Трекеры",
		"tr-TR": "İzleyiciler",   // Placeholder
		"ua-UA": "Трекери",       // Placeholder
		"uz-UZ": "Kuzatuvchilar", // Placeholder
		"zh-CN": "追踪器",           // Placeholder
	},
	BtnSpaces: {
		"en-UK": "Spaces",
		"ru-RU": "Пространства",
		"es-ES": "Espacios",
		"fa-IR": "فضاها",
		"it-IT": "Spazi",
		"de-DE": "Räume",
		"fr-FR": "Espaces",
		"pl-PL": "Przestrzenie",
		"pt-PT": "Espaços",
		"ko-KO": "공간",
		"ja-JP": "スペース",
		"zh-CN": "空间",
		"ua-UA": "Простори",
		"pt-BR": "Espaços",
		"tr-TR": "Mekanlar",
		"id-ID": "Ruang",
	},
	SpaceCmdBtnSettings: {
		"en-UK": "Settings",
		"ru-RU": "Настройки",
		"es-ES": "Ajustes",
		"fa-IR": "تنظیمات",
		"it-IT": "Impostazioni",
		"de-DE": "Einstellungen",
		"fr-FR": "Paramètres",
		"pl-PL": "Ustawienia",
		"pt-PT": "Configurações",
		"ko-KO": "설정",
		"ja-JP": "設定",
		"zh-CN": "设置",
		"ua-UA": "Налаштування",
		"pt-BR": "Configurações",
		"tr-TR": "Ayarlar",
		"id-ID": "Pengaturan",
	},
	LIST_CMD_BUY: {
		"en-UK": "Buy",
		"ru-RU": "Купить",
		"es-ES": "Comprar",
		"fa-IR": "خرید",
		"it-IT": "Acquista",
		"de-DE": "Kaufen",
		"fr-FR": "Acheter",
		"pl-PL": "Kup",
		"pt-PT": "Comprar",
		"ko-KO": "사다",
		"ja-JP": "購入",
		"zh-CN": "购买",
		"ua-UA": "Купити",
		"pt-BR": "Comprar",
		"tr-TR": "Satın al",
		"id-ID": "Beli",
	},
	LIST_CMD_TODO: {
		"en-UK": "ToDo",
		"ru-RU": "Задачи",
		"es-ES": "Tareas",
		"fa-IR": "وظایف",
		"it-IT": "Compiti",
		"de-DE": "Aufgaben",
		"fr-FR": "Tâches",
		"pl-PL": "Zadania",
		"pt-PT": "Tarefas",
		"ko-KO": "할 일",
		"ja-JP": "タスク",
		"zh-CN": "任务",
		"ua-UA": "Завдання",
		"pt-BR": "Tarefas",
		"tr-TR": "Görevler",
		"id-ID": "Tugas",
	},
	LIST_CMD_WATCH: {
		"en-UK": "Watch",
		"ru-RU": "Смотреть",
		"es-ES": "Ver",
		"fa-IR": "تماشا کنید",
		"it-IT": "Guarda",
		"de-DE": "Ansehen",
		"fr-FR": "Regarder",
		"pl-PL": "Oglądaj",
		"pt-PT": "Assistir",
		"ko-KO": "보기",
		"ja-JP": "見る",
		"zh-CN": "观看",
		"ua-UA": "Дивитися",
		"pt-BR": "Assistir",
		"tr-TR": "İzle",
		"id-ID": "Menonton",
	},
	LIST_CMD_READ: {
		"en-UK": "Read",
		"ru-RU": "Читать",
		"es-ES": "Leer",
		"fa-IR": "خواندن",
		"it-IT": "Leggi",
		"de-DE": "Lesen",
		"fr-FR": "Lire",
		"pl-PL": "Czytaj",
		"pt-PT": "Ler",
		"ko-KO": "읽기",
		"ja-JP": "読む",
		"zh-CN": "读",
		"ua-UA": "Читати",
		"pt-BR": "Ler",
		"tr-TR": "Oku",
		"id-ID": "Baca",
	},
	ListCmdBtnToRead: {
		"en-UK": "To read",
		"ru-RU": "Прочитать",
		"es-ES": "Leer",
		"fa-IR": "برای خواندن",
		"it-IT": "Da leggere",
		"de-DE": "Zu lesen",
		"fr-FR": "À lire",
		"pl-PL": "Do przeczytania",
		"pt-PT": "Para ler",
		"ko-KO": "읽을 것",
		"ja-JP": "読む",
		"zh-CN": "阅读",
		"ua-UA": "Читати",
		"pt-BR": "Para ler",
		"tr-TR": "Okunacaklar",
		"id-ID": "Untuk dibaca",
	},
	ListCmdBtnToWatch: {
		"en-UK": "To watch",
		"ru-RU": "Посмотреть",
		"es-ES": "Ver",
		"fa-IR": "برای تماشا",
		"it-IT": "Da guardare",
		"de-DE": "Ansehen",
		"fr-FR": "À regarder",
		"pl-PL": "Do obejrzenia",
		"pt-PT": "Para assistir",
		"ko-KO": "볼 것",
		"ja-JP": "見る",
		"zh-CN": "观看",
		"ua-UA": "Дивитися",
		"pt-BR": "Para assistir",
		"tr-TR": "İzlenecekler",
		"id-ID": "Untuk ditonton",
	},
	ListCmdBtnToBuy: {
		"en-UK": "To buy",
		"ru-RU": "Купить",
		"es-ES": "Comprar",
		"fa-IR": "برای خرید",
		"it-IT": "Da comprare",
		"de-DE": "Zu kaufen",
		"fr-FR": "À acheter",
		"pl-PL": "Do kupienia",
		"pt-PT": "Para comprar",
		"ko-KO": "구매할 것",
		"ja-JP": "購入",
		"zh-CN": "购买",
		"ua-UA": "Купити",
		"pt-BR": "Para comprar",
		"tr-TR": "Alınacaklar",
		"id-ID": "Untuk dibeli",
	},
	ListCmdBtnToDo: {
		"en-UK": "To do",
		"ru-RU": "Сделать",
		"es-ES": "Hacer",
		"fa-IR": "انجام دادن",
		"it-IT": "Da fare",
		"de-DE": "Zu tun",
		"fr-FR": "À faire",
		"pl-PL": "Do zrobienia",
		"pt-PT": "Para fazer",
		"ko-KO": "할 일",
		"ja-JP": "やる",
		"zh-CN": "做",
		"ua-UA": "Зробити",
		"pt-BR": "Para fazer",
		"tr-TR": "Yapılacaklar",
		"id-ID": "Untuk dilakukan",
	},
	Groceries: {
		"en-UK": "Groceries",
		"ru-RU": "Продукты",
		"es-ES": "Comestibles",
		"fa-IR": "خوار و بار",
		"it-IT": "Generi alimentari",
		"de-DE": "Lebensmittel",
		"fr-FR": "Épicerie",
		"pl-PL": "Artykuły spożywcze",
		"pt-PT": "Comestíveis",
		"ko-KO": "식료품",
		"ja-JP": "食料品",
		"zh-CN": "杂货",
		"ua-UA": "Продукти",
		"pt-BR": "Comestíveis",
		"tr-TR": "Bakkaliye",
		"id-ID": "Bahan makanan",
	},
	Books: {
		"en-UK": "Books",
		"ru-RU": "Книги",
		"es-ES": "Libros",
		"fa-IR": "کتاب ها",
		"it-IT": "Libri",
		"de-DE": "Bücher",
		"fr-FR": "Livres",
		"pl-PL": "Książki",
		"pt-PT": "Livros",
		"ko-KO": "책",
		"ja-JP": "本",
		"zh-CN": "书籍",
		"ua-UA": "Книги",
		"pt-BR": "Livros",
		"tr-TR": "Kitaplar",
		"id-ID": "Buku",
	},
	Movies: {
		"en-UK": "Movies",
		"ru-RU": "Фильмы",
		"es-ES": "Películas",
		"fa-IR": "فیلم ها",
		"it-IT": "Film",
		"de-DE": "Filme",
		"fr-FR": "Films",
		"pl-PL": "Filmy",
		"pt-PT": "Filmes",
		"ko-KO": "영화",
		"ja-JP": "映画",
		"zh-CN": "电影",
		"ua-UA": "Фільми",
		"pt-BR": "Filmes",
		"tr-TR": "Filmler",
		"id-ID": "Film",
	},
	Tasks: {
		"en-UK": "Tasks",
		"ru-RU": "Задачи",
		"es-ES": "Tareas",
		"fa-IR": "وظایف",
		"it-IT": "Compiti",
		"de-DE": "Aufgaben",
		"fr-FR": "Tâches",
		"pl-PL": "Zadania",
		"pt-PT": "Tarefas",
		"ko-KO": "할 일",
		"ja-JP": "タスク",
		"zh-CN": "任务",
		"ua-UA": "Завдання",
		"pt-BR": "Tarefas",
		"tr-TR": "Görevler",
		"id-ID": "Tugas",
	},
	ListsOfFamily: {
		"en-UK": "Family lists",
		"ru-RU": "Семейные списки",
		"es-ES": "Listas familiares",
		"fa-IR": "لیست های خانواده",
		"it-IT": "Liste familiari",
		"de-DE": "Familienlisten",
		"fr-FR": "Listes familiales",
		"pl-PL": "Listy rodzinne",
		"pt-PT": "Listas familiares",
		"ko-KO": "가족 목록",
		"ja-JP": "家族のリスト",
		"zh-CN": "家庭清单",
		"ua-UA": "Сімейні списки",
		"pt-BR": "Listas familiares",
		"tr-TR": "Aile listeleri",
		"id-ID": "Daftar keluarga",
	},
	ListsOfPrivate: {
		"en-UK": "Private lists",
		"ru-RU": "Личные списки",
		"es-ES": "Listas privadas",
		"fa-IR": "لیست های خصوصی",
		"it-IT": "Liste private",
		"de-DE": "Private Listen",
		"fr-FR": "Listes privées",
		"pl-PL": "Prywatne listy",
		"pt-PT": "Listas privadas",
		"ko-KO": "개인 목록",
		"ja-JP": "プライベートリスト",
		"zh-CN": "私人清单",
		"ua-UA": "Особисті списки",
		"pt-BR": "Listas privadas",
		"tr-TR": "Özel listeler",
		"id-ID": "Daftar pribadi",
	},
	ListsOfSpace: {
		"en-UK": "Lists @ %s",
		"ru-RU": "Списки @ %s",
		"es-ES": "Listas @ %s",
		"fa-IR": "لیست ها @ %s",
		"it-IT": "Liste @ %s",
		"de-DE": "Listen @ %s",
		"fr-FR": "Listes @ %s",
		"pl-PL": "Listy @ %s",
		"pt-PT": "Listas @ %s",
		"ko-KO": "목록 @ %s",
		"ja-JP": "リスト @ %s",
		"zh-CN": "清单 @ %s",
		"ua-UA": "Списки @ %s",
		"pt-BR": "Listas @ %s",
		"tr-TR": "Listeler @ %s",
		"id-ID": "Daftar @ %s",
	},
	BtnBackToSpace: {
		"en-UK": "Back to space",
		"ru-RU": "К пространству",
		"es-ES": "Volver al espacio",
		"fr-FR": "Retour à l'espace",
		"de-DE": "Zurück zum Bereich",
		"it-IT": "Torna allo spazio",
		"pt-PT": "Voltar ao espaço",
		"pt-BR": "Voltar ao espaço",
		"zh-CN": "返回空间",
		"ja-JP": "スペースに戻る",
		"ko-KO": "공간으로 돌아가기",
		"pl-PL": "Powrót do przestrzeni",
		"ua-UA": "Назад до простору",
		"tr-TR": "Alana geri dön",
		"id-ID": "Kembali ke ruang",
		"fa-IR": "بازگشت به فضا",
	},
	BtnPrivate: {
		"en-UK": "Private",
		"ru-RU": "Личное",
		"es-ES": "Privado",
		"fa-IR": "شخصی",
		"it-IT": "Privato",
		"de-DE": "Privat",
		"fr-FR": "Privé",
		"pl-PL": "Prywatne",
		"pt-PT": "Privado",
		"ko-KO": "개인적인",
		"ja-JP": "個人",
		"zh-CN": "私人",
		"ua-UA": "Особисте",
		"tr-TR": "Özel",
		"id-ID": "Pribadi",
	},
	BtnFamily: {
		"en-UK": "Family",
		"ru-RU": "Семья",
		"es-ES": "Familia",
		"fa-IR": "خانواده",
		"it-IT": "Famiglia",
		"de-DE": "Familie",
		"fr-FR": "Famille",
		"pl-PL": "Rodzina",
		"pt-PT": "Família",
		"ko-KO": "가족",
		"ja-JP": "家族",
		"zh-CN": "家庭",
		"ua-UA": "Сім'я",
		"pt-BR": "Família",
		"tr-TR": "Aile",
		"id-ID": "Keluarga",
	},
	TrackerPushUps: {
		"en-UK": "Push-ups",
		"ru-RU": "Отжимания",
		"es-ES": "Flexiones",
		"fa-IR": "پرسه",
		"it-IT": "Flessioni",
		"de-DE": "Liegestütze",
		"fr-FR": "Pompes",
		"pl-PL": "Pompki",
		"pt-PT": "Flexões",
		"ko-KO": "푸시업",
		"ja-JP": "腕立て伏せ",
		"zh-CN": "俯卧撑",
		"ua-UA": "Віджимання",
		"pt-BR": "Flexões",
		"tr-TR": "Şınav",
		"id-ID": "Push-up",
	},
	TrackerPullUps: {
		"en-UK": "Pull-ups",
		"ru-RU": "Подтягивания",
		"es-ES": "Dominadas",
		"fa-IR": "کشیدن",
		"it-IT": "Trazioni",
		"de-DE": "Klimmzüge",
		"fr-FR": "Tirages",
		"pl-PL": "Podciągania",
		"pt-PT": "Pull-ups",
		"ko-KO": "풀업",
		"ja-JP": "引き上げ",
		"zh-CN": "引体向上",
		"ua-UA": "Підтягування",
		"pt-BR": "Pull-ups",
		"tr-TR": "Çekme",
		"id-ID": "Pull-up",
	},
	TrackerSquats: {
		"en-UK": "Squats",
		"ru-RU": "Приседания",
		"es-ES": "Sentadillas",
		"fa-IR": "کرنچ",
		"it-IT": "Squat",
		"de-DE": "Kniebeugen",
		"fr-FR": "Squats",
		"pl-PL": "Przysiady",
		"pt-PT": "Agachamentos",
		"ko-KO": "스쿼트",
		"ja-JP": "スクワット",
		"zh-CN": "深蹲",
		"ua-UA": "Присідання",
		"pt-BR": "Agachamentos",
		"tr-TR": "Squat",
		"id-ID": "Squat",
	},
	TrackerJumpingJacks: {
		"en-UK": "Jumping jacks",
		"ru-RU": "Прыжки на месте",
		"es-ES": "Saltos de tijera",
		"fa-IR": "حرکت پروانه",
		"it-IT": "Salti con apertura",
		"de-DE": "Hampelmänner",
		"fr-FR": "Jumping jacks",
		"pl-PL": "Pajacyki",
		"pt-PT": "Saltos de estrela",
		"ko-KO": "팔 벌려 뛰기",
		"ja-JP": "ジャンピングジャック",
		"zh-CN": "开合跳",
		"ua-UA": "Стрибки з розмахом рук",
		"pt-BR": "Polichinelos",
		"tr-TR": "Zıplama hareketi",
		"id-ID": "Lompat bintang",
	},
	TrackerFuelCost: {
		"en-UK": "Fuel Cost",
		"ru-RU": "Стоимость топлива",
		"es-ES": "Costo del combustible",
		"fa-IR": "هزینه سوخت",
		"it-IT": "Costo del carburante",
		"de-DE": "Kraftstoffkosten",
		"fr-FR": "Coût du carburant",
		"pl-PL": "Koszt paliwa",
		"pt-PT": "Custo do combustível",
		"ko-KO": "연료 비용",
		"ja-JP": "燃料費",
		"zh-CN": "燃料成本",
		"ua-UA": "Вартість палива",
		"pt-BR": "Custo do combustível",
		"tr-TR": "Yakıt Maliyeti",
		"id-ID": "Biaya bahan bakar",
	},
	TrackerFuelVolume: {
		"en-UK": "Fuel Volume",
		"ru-RU": "Объем топлива",
		"es-ES": "Volumen de combustible",
		"fa-IR": "حجم سوخت",
		"it-IT": "Volume di carburante",
		"de-DE": "Kraftstoffvolumen",
		"fr-FR": "Volume de carburant",
		"pl-PL": "Objętość paliwa",
		"pt-PT": "Volume de combustível",
		"ko-KO": "연료 양",
		"ja-JP": "燃料量",
		"zh-CN": "燃料容量",
		"ua-UA": "Об'єм палива",
		"pt-BR": "Volume de combustível",
		"tr-TR": "Yakıt Hacmi",
		"id-ID": "Volume bahan bakar",
	},
	TrackerMileage: {
		"en-UK": "Mileage",
		"ru-RU": "Пробег",
		"es-ES": "Kilometraje",
		"fa-IR": "مسافت پیموده شده",
		"it-IT": "Chilometraggio",
		"de-DE": "Kilometerstand",
		"fr-FR": "Kilométrage",
		"pl-PL": "Przebieg",
		"pt-PT": "Quilometragem",
		"ko-KO": "주행 거리",
		"ja-JP": "走行距離",
		"zh-CN": "里程",
		"ua-UA": "Пробіг",
		"pt-BR": "Quilometragem",
		"tr-TR": "Kilometre",
		"id-ID": "Jarak tempuh",
	},
	TrackerHeight: {
		"en-UK": "Height",
		"ru-RU": "Рост",
		"es-ES": "Altura",
		"fa-IR": "قد",
		"it-IT": "Altezza",
		"de-DE": "Höhe",
		"fr-FR": "Taille",
		"pl-PL": "Wzrost",
		"pt-PT": "Altura",
		"ko-KO": "키",
		"ja-JP": "身長",
		"zh-CN": "身高",
		"ua-UA": "Зріст",
		"pt-BR": "Altura",
		"tr-TR": "Boy",
		"id-ID": "Tinggi",
	},
	TrackerWeight: {
		"en-UK": "Weight",
		"ru-RU": "Вес",
		"es-ES": "Peso",
		"fa-IR": "وزن",
		"it-IT": "Peso",
		"de-DE": "Gewicht",
		"fr-FR": "Poids",
		"pl-PL": "Waga",
		"pt-PT": "Peso",
		"ko-KO": "몸무게",
		"ja-JP": "体重",
		"zh-CN": "体重",
		"ua-UA": "Вага",
		"pt-BR": "Peso",
		"tr-TR": "Ağırlık",
		"id-ID": "Berat",
	},
	TrackerSpending: {
		"en-UK": "Spending",
		"ru-RU": "Расходы",
		"es-ES": "Gastos",
		"fa-IR": "هزینه\u200cها",
		"it-IT": "Spese",
		"de-DE": "Ausgaben",
		"fr-FR": "Dépenses",
		"pl-PL": "Wydatki",
		"pt-PT": "Despesas",
		"ko-KO": "지출",
		"ja-JP": "支出",
		"zh-CN": "支出",
		"ua-UA": "Витрати",
		"pt-BR": "Despesas",
		"tr-TR": "Harcamalar",
		"id-ID": "Pengeluaran",
	},
	TrackerBloodPressure: {
		"en-UK": "Blood Pressure",
		"ru-RU": "Кровяное давление",
		"es-ES": "Presión arterial",
		"fa-IR": "فشار خون",
		"it-IT": "Pressione sanguigna",
		"de-DE": "Blutdruck",
		"fr-FR": "Pression artérielle",
		"pl-PL": "Ciśnienie krwi",
		"pt-PT": "Pressão arterial",
		"ko-KO": "혈압",
		"ja-JP": "血圧",
		"zh-CN": "血压",
		"ua-UA": "Кров'яний тиск",
		"pt-BR": "Pressão arterial",
		"tr-TR": "Kan basıncı",
		"id-ID": "Tekanan darah",
	},
	TrackerCategoryHealth: {

		"en-UK": "Health",
		"ru-RU": "Здоровье",
		"es-ES": "Salud",
		"fa-IR": "سلامتی",
		"it-IT": "Salute",
		"de-DE": "Gesundheit",
		"fr-FR": "Santé",
		"pl-PL": "Zdrowie",
		"pt-PT": "Saúde",
		"ko-KO": "건강",
		"ja-JP": "健康",
		"zh-CN": "健康",
		"ua-UA": "Здоров'я",
		"pt-BR": "Saúde",
		"tr-TR": "Sağlık",
		"id-ID": "Kesehatan",
	},
	TrackerCategoryFitness: {
		"en-UK": "Fitness",
		"ru-RU": "Фитнес",
		"es-ES": "Aptitud física",
		"fa-IR": "تناسب اندام",
		"it-IT": "Fitness",
		"de-DE": "Fitness",
		"fr-FR": "Forme physique",
		"pl-PL": "Fitness",
		"pt-PT": "Aptidão física",
		"ko-KO": "피트니스",
		"ja-JP": "フィットネス",
		"zh-CN": "健身",
		"ua-UA": "Фітнес",
		"pt-BR": "Aptidão física",
		"tr-TR": "Fitness",
		"id-ID": "Kebugaran fisik",
	},
	TrackerCategoryVehicle: {
		"en-UK": "Vehicle",
		"ru-RU": "Транспортное средство",
		"es-ES": "Vehículo",
		"fa-IR": "وسیله نقلیه",
		"it-IT": "Veicolo",
		"de-DE": "Fahrzeug",
		"fr-FR": "Véhicule",
		"pl-PL": "Pojazd",
		"pt-PT": "Veículo",
		"ko-KO": "차량",
		"ja-JP": "車両",
		"zh-CN": "车辆",
		"ua-UA": "Транспортний засіб",
		"pt-BR": "Veículo",
		"tr-TR": "Araç",
		"id-ID": "Kendaraan",
	},
	CustomTracker: {
		"en-UK": "Custom tracker",
		"ru-RU": "Свой трекер",
		"es-ES": "Rastreador personalizado",
		"fa-IR": "ردیاب سفارشی",
		"it-IT": "Tracker personalizzato",
		"de-DE": "Benutzerdefinierter Tracker",
		"fr-FR": "Tracker personnalisé",
		"pl-PL": "Niestandardowy tracker",
		"pt-PT": "Rastreador personalizado",
		"ko-KO": "사용자 정의 추적기",
		"ja-JP": "カスタムトラッカー",
		"zh-CN": "自定义追踪器",
		"ua-UA": "Свій трекер",
		"pt-BR": "Rastreador personalizado",
		"tr-TR": "Özel izleyici",
		"id-ID": "Pelacak kustom",
	},
	BackToTrackers: {
		"en-UK": "Back to trackers",
		"ru-RU": "Назад к трекерам",
		"es-ES": "Volver a los rastreadores",
		"fa-IR": "بازگشت به ردیاب\u200cها",
		"it-IT": "Torna ai tracker",
		"de-DE": "Zurück zu Trackern",
		"fr-FR": "Retour aux trackers",
		"pl-PL": "Powrót do trackerów",
		"pt-PT": "Voltar aos rastreadores",
		"ko-KO": "트래커로 돌아가기",
		"ja-JP": "トラッカーに戻る",
		"zh-CN": "返回追踪器",
		"ua-UA": "Повернутись до трекерів",
		"pt-BR": "Voltar aos rastreadores",
		"tr-TR": "İzleyicilere geri dön",
		"id-ID": "Kembali ke pelacak",
	},
	AddTracker: {
		"en-UK": "Add tracker",
		"ru-RU": "Добавить трекер",
		"es-ES": "Añadir rastreador",
		"fa-IR": "افزودن ردیاب",
		"it-IT": "Aggiungi tracker",
		"de-DE": "Tracker hinzufügen",
		"fr-FR": "Ajouter un tracker",
		"pl-PL": "Dodaj tracker",
		"pt-PT": "Adicionar rastreador",
		"ko-KO": "추적기 추가",
		"ja-JP": "トラッカーを追加",
		"zh-CN": "添加追踪器",
		"ua-UA": "Додати трекер",
		"pt-BR": "Adicionar rastreador",
		"tr-TR": "İzleyici ekle",
		"id-ID": "Tambahkan pelacak",
	},
	ShareTracker: {
		"en-UK": "Share tracker",
		"ru-RU": "Поделиться трекером",
		"es-ES": "Compartir rastreador",
		"fa-IR": "اشتراک\u200cگذاری ردیاب",
		"it-IT": "Condividi il tracker",
		"de-DE": "Tracker teilen",
		"fr-FR": "Partager le tracker",
		"pl-PL": "Udostępnij tracker",
		"pt-PT": "Partilhar rastreador",
		"ko-KO": "추적기 공유",
		"ja-JP": "トラッカーを共有",
		"zh-CN": "分享追踪器",
		"ua-UA": "Поділитись трекером",
		"pt-BR": "Compartilhar rastreador",
		"tr-TR": "İzleyiciyi paylaş",
		"id-ID": "Bagikan pelacak",
	},
	RecordTrackerValue: {
		"en-UK": "Record value",
		"ru-RU": "Записать значение",
		"es-ES": "Registrar valor",
		"fa-IR": "ثبت مقدار",
		"it-IT": "Registra valore",
		"de-DE": "Wert aufzeichnen",
		"fr-FR": "Enregistrer la valeur",
		"pl-PL": "Zapisz wartość",
		"pt-PT": "Registrar valor",
		"ko-KO": "값 기록",
		"ja-JP": "値を記録",
		"zh-CN": "记录值",
		"ua-UA": "Записати значення",
		"pt-BR": "Registrar valor",
		"tr-TR": "Değeri kaydet",
		"id-ID": "Catat nilai",
	},
	FamilyTrackers: {
		"en-UK": "Family trackers",
		"ru-RU": "Семейные трекеры",
		"es-ES": "Rastreadores familiares",
		"fa-IR": "ردیاب\u200cهای خانواده",
		"it-IT": "Tracker familiari",
		"de-DE": "Familientracker",
		"fr-FR": "Trackers familiaux",
		"pl-PL": "Trackery rodzinne",
		"pt-PT": "Rastreadores de família",
		"ko-KO": "가족 추적기",
		"ja-JP": "家族のトラッカー",
		"zh-CN": "家庭追踪器",
		"ua-UA": "Сімейні трекери",
		"pt-BR": "Rastreadores de família",
		"tr-TR": "Aile izleyicileri",
		"id-ID": "Pelacak keluarga",
	},
	NoActiveTrackers: {
		"en-UK": "Currently you do not have any active trackers. Add some to start tracking your progress.",
		"ru-RU": "В настоящее время у вас нет активных трекеров. Добавьте их, чтобы начать отслеживать ваш прогресс.",
		"es-ES": "Actualmente, no tienes rastreadores activos. Agrega algunos para comenzar a rastrear tu progreso.",
		"fa-IR": "در حال حاضر هیچ ردیابی فعال ندارید. چند مورد اضافه کنید تا پیشرفت خود را پیگیری کنید.",
		"it-IT": "Attualmente non hai tracker attivi. Aggiungine alcuni per iniziare a monitorare i tuoi progressi.",
		"de-DE": "Derzeit haben Sie keine aktiven Tracker. Fügen Sie einige hinzu, um Ihren Fortschritt zu verfolgen.",
		"fr-FR": "Vous n'avez actuellement aucun tracker actif. Ajoutez-en pour commencer à suivre vos progrès.",
		"pl-PL": "Obecnie nie masz żadnych aktywnych trackerów. Dodaj kilka, aby zacząć śledzić swoje postępy.",
		"pt-PT": "Atualmente, não tens rastreadores ativos. Adiciona alguns para começar a monitorizar o teu progresso.",
		"ko-KO": "현재 활성화된 추적기가 없습니다. 진행 상황을 추적하려면 몇 가지를 추가하세요.",
		"ja-JP": "現在、有効なトラッカーはありません。進捗を追跡するには、いくつか追加してください。",
		"zh-CN": "目前你还没有任何活跃的追踪器。添加一些以开始追踪你的进度。",
		"ua-UA": "Наразі у вас немає активних трекерів. Додайте кілька, щоб почати відстежувати свій прогрес.",
		"pt-BR": "Atualmente, você não tem rastreadores ativos. Adicione alguns para começar a rastrear seu progresso.",
		"tr-TR": "Şu anda aktif izleyiciniz yok. İlerlemenizi izlemek için birkaç tane ekleyin.",
		"id-ID": "Saat ini Anda tidak memiliki pelacak aktif. Tambahkan beberapa untuk mulai melacak kemajuan Anda.",
	},
	OneActiveTracker: {
		"en-UK": "You have one active tracker. Select it to record progress. Or add a new one to track more.",
		"ru-RU": "У вас есть один активный трекер. Выберите его, чтобы записать прогресс, или добавьте новый для отслеживания большего.",
		"es-ES": "Tienes un rastreador activo. Selecciónalo para registrar el progreso o agrega uno nuevo para rastrear más.",
		"fa-IR": "شما یک ردیاب فعال دارید. آن را انتخاب کنید تا پیشرفت خود را ثبت کنید یا یک ردیاب جدید اضافه کنید تا موارد بیشتری را پیگیری کنید.",
		"it-IT": "Hai un tracker attivo. Selezionalo per registrare i progressi o aggiungi uno nuovo per monitorare di più.",
		"de-DE": "Sie haben einen aktiven Tracker. Wählen Sie ihn aus, um Fortschritte aufzuzeichnen, oder fügen Sie einen neuen hinzu, um mehr zu verfolgen.",
		"fr-FR": "Vous avez un tracker actif. Sélectionnez-le pour enregistrer vos progrès ou ajoutez-en un nouveau pour en suivre davantage.",
		"pl-PL": "Masz jeden aktywny tracker. Wybierz go, aby zarejestrować postępy, lub dodaj nowy, aby śledzić więcej.",
		"pt-PT": "Tens um rastreador ativo. Seleciona-o para registar o progresso, ou adiciona um novo para monitorizar mais.",
		"ko-KO": "활성화된 추적기 하나가 있습니다. 선택하여 진행 상황을 기록하거나 더 많이 추적하려면 새로 추가하세요.",
		"ja-JP": "アクティブなトラッカーが1つあります。選択して進捗を記録するか、新しいものを追加してさらに追跡してください。",
		"zh-CN": "你有一个活跃的追踪器。选择它来记录你的进度，或者添加一个新的来追踪更多。",
		"ua-UA": "У вас є один активний трекер. Виберіть його, щоб записати прогрес, або додайте новий для відстеження більшого.",
		"pt-BR": "Você tem um rastreador ativo. Selecione-o para registrar o progresso ou adicione um novo para rastrear mais.",
		"tr-TR": "Bir adet aktif izleyiciniz var. İlerlemenizi kaydetmek için bunu seçin veya daha fazla şey izlemek için yeni bir tane ekleyin.",
		"id-ID": "Anda memiliki satu pelacak aktif. Pilih untuk mencatat kemajuan atau tambahkan yang baru untuk melacak lebih banyak.",
	},
	NActiveTrackers: {
		"en-UK": "You have few active tracker. Select one to record progress.",
		"ru-RU": "У вас несколько активных трекеров. Выберите один, чтобы записать прогресс.",
		"es-ES": "Tienes varios rastreadores activos. Selecciona uno para registrar el progreso.",
		"fa-IR": "چند ردیاب فعال دارید. یکی را انتخاب کنید تا پیشرفت خود را ثبت کنید.",
		"it-IT": "Hai alcuni tracker attivi. Selezionane uno per registrare i progressi.",
		"de-DE": "Sie haben einige aktive Tracker. Wählen Sie einen aus, um Fortschritte aufzuzeichnen.",
		"fr-FR": "Vous avez plusieurs trackers actifs. Sélectionnez-en un pour enregistrer vos progrès.",
		"pl-PL": "Masz kilka aktywnych trackerów. Wybierz jeden, aby zarejestrować postępy.",
		"pt-PT": "Tens vários rastreadores ativos. Seleciona um para registar o progresso.",
		"ko-KO": "활성화된 추적기 몇 개가 있습니다. 하나를 선택하여 진행 상황을 기록하세요.",
		"ja-JP": "複数のアクティブなトラッカーがあります。1つ選択して進捗を記録してください。",
		"zh-CN": "你有几个活跃的追踪器。选择一个来记录你的进度。",
		"ua-UA": "У вас є кілька активних трекерів. Виберіть один, щоб записати прогрес.",
		"pt-BR": "Você tem vários rastreadores ativos. Selecione um para registrar o progresso.",
		"tr-TR": "Birkaç aktif izleyiciniz var. İlerlemenizi kaydetmek için birini seçin.",
		"id-ID": "Anda memiliki beberapa pelacak aktif. Pilih salah satu untuk mencatat kemajuan.",
	},
	HintForIntTracker: {
		"en-UK": "Send an integer number to record it to the tracker",
		"ru-RU": "Отправьте целое число, чтобы записать его в трекер",
		"es-ES": "Envía un número entero para registrarlo en el rastreador",
		"fa-IR": "یک عدد صحیح ارسال کنید تا آن را در ردیاب ثبت کنید",
		"it-IT": "Invia un numero intero per registrarlo nel tracker",
		"de-DE": "Senden Sie eine Ganzzahl, um sie im Tracker zu speichern",
		"fr-FR": "Envoyez un nombre entier pour l'enregistrer dans le tracker",
		"pl-PL": "Wyślij liczbę całkowitą, aby zapisać ją w trackerze",
		"pt-PT": "Envia um número inteiro para o registar no rastreador",
		"ko-KO": "정수 값을 추적기에 기록하려면 전송하세요",
		"ja-JP": "整数値を送信してトラッカーに記録してください",
		"zh-CN": "发送一个整数以将其记录到追踪器中",
		"ua-UA": "Надішліть ціле число, щоб записати його в трекер",
		"pt-BR": "Envie um número inteiro para registrá-lo no rastreador",
		"tr-TR": "Bir tamsayı göndererek bunu izleyiciye kaydedin",
		"id-ID": "Kirim bilangan bulat untuk mencatatnya ke pelacak",
	},
	Tracker: {
		"en-UK": "Tracker",
		"ru-RU": "Трекер",
		"es-ES": "Rastreador",
		"fa-IR": "ردیاب",
		"it-IT": "Tracker",
		"de-DE": "Tracker",
		"fr-FR": "Tracker",
		"pl-PL": "Tracker",
		"pt-PT": "Rastreador",
		"ko-KO": "추적기",
		"ja-JP": "トラッカー",
		"zh-CN": "追踪器",
		"ua-UA": "Трекер",
		"pt-BR": "Rastreador",
		"tr-TR": "İzleyici",
		"id-ID": "Pelacak",
	},
	HintToShareTracker: {
		"en-UK": `\n\nYou can share this tracker either
	🤫 with a friend (<i>a single person would be able to accept</i>)
	🌍 or publicly (<i>anyone with a link can accept, you can cancel it at any time</i>)`,

		"ru-RU": `\n\nВы можете поделиться этим трекером либо
🤫 с другом (<i>только один человек сможет принять</i>)
🌍 или публично (<i>любой с ссылкой может принять, вы можете отменить это в любое время</i>)`,
		"es-ES": `\n\nPuedes compartir este rastreador
🤫 con un amigo (<i>una sola persona puede aceptarlo</i>)
🌍 o públicamente (<i>cualquiera con el enlace puede aceptarlo, puedes cancelarlo en cualquier momento</i>)`,
		"fa-IR": `\n\nشما می\u200cتوانید این ردیاب را به دو صورت به اشتراک بگذارید
🤫 با یک دوست (<i>فقط یک نفر می\u200cتواند بپذیرد</i>)
🌍 یا به صورت عمومی (<i>هرکسی با لینک می\u200cتواند بپذیرد، شما می\u200cتوانید در هر زمان آن را لغو کنید</i>)`,
		"it-IT": `\n\nPuoi condividere questo tracker
🤫 con un amico (<i>una sola persona potrà accettarlo</i>)
🌍 oppure pubblicamente (<i>chiunque con il link potrà accettarlo, puoi annullarlo in qualsiasi momento</i>)`,
		"de-DE": `\n\nSie können diesen Tracker entweder
🤫 mit einem Freund teilen (<i>nur eine Person kann ihn akzeptieren</i>)
🌍 oder öffentlich (<i>jeder mit dem Link kann ihn akzeptieren, Sie können es jederzeit abbrechen</i>)`,
		"fr-FR": `\n\nVous pouvez partager ce tracker
🤫 avec un ami (<i>une seule personne pourra l'accepter</i>)
🌍 ou publiquement (<i>n'importe qui avec un lien peut l'accepter, vous pouvez l'annuler à tout moment</i>)`,
		"pl-PL": `\n\nMożesz udostępnić ten tracker
🤫 znajomemu (<i>tylko jedna osoba będzie mogła przyjąć</i>)
🌍 lub publicznie (<i>każdy z linkiem może przyjąć, możesz to anulować w dowolnym momencie</i>)`,
		"pt-PT": `\n\nPodes partilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa poderá aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, podes cancelá-lo a qualquer momento</i>)`,
		"ko-KO": `\n\n이 추적기를 공유할 수 있습니다
🤫 친구와 (<i>단 한 사람만 수락 가능</i>)
🌍 또는 공개적으로 (<i>링크를 가진 모든 사람이 수락할 수 있으며, 언제든 취소할 수 있습니다</i>)`,
		"ja-JP": `\n\nこのトラッカーを以下の方法で共有できます
🤫 友人と (<i>受け取れるのは1人だけ</i>)
🌍 または公開で (<i>リンクを持っている誰でも受け取れます。いつでもキャンセルできます</i>)`,
		"zh-CN": `\n\n您可以通过以下方式共享此追踪器
🤫 与朋友共享 (<i>只有一个人可以接受</i>)
🌍 或公开共享 (<i>任何拥有链接的人都可以接受，您可以随时取消</i>)`,
		"ua-UA": `\n\nВи можете поділитися цим трекером
🤫 з другом (<i>лише одна людина зможе прийняти</i>)
🌍 або публічно (<i>кожен, хто має посилання, може прийняти, ви можете скасувати це в будь-який час</i>)`,
		"pt-BR": `\n\nVocê pode compartilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa pode aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, você pode cancelar a qualquer momento</i>)`,
		"tr-TR": `\n\nBu izleyiciyi şu şekilde paylaşabilirsiniz
🤫 bir arkadaşla (<i>yalnızca bir kişi kabul edebilir</i>)
🌍 veya genel olarak (<i>bağlantıya sahip herkes kabul edebilir, istediğiniz zaman iptal edebilirsiniz</i>)`,
		"id-ID": `\n\nAnda dapat membagikan pelacak ini
🤫 kepada seorang teman (<i>hanya satu orang yang dapat menerima</i>)
🌍 atau secara publik (<i>siapa saja yang memiliki tautan dapat menerimanya, Anda dapat membatalkannya kapan saja</i>)`,
	},
	YourSpaces: {
		"en-UK": "Your spaces",
		"ru-RU": "Ваши пространства",
		"es-ES": "Tus espacios",
		"fa-IR": "فضاهای شما",
		"it-IT": "I tuoi spazi",
		"de-DE": "Ihre Bereiche",
		"fr-FR": "Vos espaces",
		"pl-PL": "Twoje przestrzenie",
		"pt-PT": "Os teus espaços",
		"ko-KO": "귀하의 공간",
		"ja-JP": "あなたのスペース",
		"zh-CN": "您的空间",
		"ua-UA": "Ваші простори",
		"pt-BR": "Seus espaços",
		"tr-TR": "Alanlarınız",
		"id-ID": "Ruang Anda",
	},
	CurrentSpace: {
		"en-UK": "Current space",
		"ru-RU": "Текущее пространство",
		"es-ES": "Espacio actual",
		"fa-IR": "فضای کنونی",
		"it-IT": "Spazio corrente",
		"de-DE": "Aktueller Bereich",
		"fr-FR": "Espace actuel",
		"pl-PL": "Bieżąca przestrzeń",
		"pt-PT": "Espaço atual",
		"ko-KO": "현재 공간",
		"ja-JP": "現在のスペース",
		"zh-CN": "当前空间",
		"ua-UA": "Поточний простір",
		"pt-BR": "Espaço atual",
		"tr-TR": "Mevcut alan",
		"id-ID": "Ruang saat ini",
	},
	ClickToSwitchCurrentSpace: {
		"en-UK": "Click to switch current space",
		"ru-RU": "Нажмите, чтобы переключить текущее пространство",
		"es-ES": "Haz clic para cambiar el espacio actual",
		"fa-IR": "برای تغییر فضای کنونی کلیک کنید",
		"it-IT": "Clicca per cambiare spazio corrente",
		"de-DE": "Klicken Sie hier, um den aktuellen Bereich zu wechseln",
		"fr-FR": "Cliquez pour changer l'espace actuel",
		"pl-PL": "Kliknij, aby zmienić bieżącą przestrzeń",
		"pt-PT": "Clique para mudar o espaço atual",
		"ko-KO": "현재 공간을 전환하려면 클릭하세요",
		"ja-JP": "クリックで現在のスペースを切り替えます",
		"zh-CN": "点击切换当前空间",
		"ua-UA": "Натисніть, щоб змінити поточний простір",
		"pt-BR": "Clique para alterar o espaço atual",
		"tr-TR": "Mevcut alanı değiştirmek için tıklayın",
		"id-ID": "Klik untuk mengganti ruang saat ini",
	},
	Family: {
		"en-UK": "Family",
		"ru-RU": "Семья",
		"es-ES": "Familia",
		"fa-IR": "خانواده",
		"it-IT": "Famiglia",
		"de-DE": "Familie",
		"fr-FR": "Famille",
		"pl-PL": "Rodzina",
		"pt-PT": "Família",
		"ko-KO": "가족",
		"ja-JP": "家族",
		"zh-CN": "家庭",
		"ua-UA": "Родина",
		"pt-BR": "Família",
		"tr-TR": "Aile",
		"id-ID": "Keluarga",
	},
	Private: {
		"en-UK": "Private",
		"ru-RU": "Личное",
		"es-ES": "Privado",
		"fa-IR": "خصوصی",
		"it-IT": "Privato",
		"de-DE": "Privat",
		"fr-FR": "Privé",
		"pl-PL": "Prywatne",
		"pt-PT": "Privado",
		"ko-KO": "개인",
		"ja-JP": "プライベート",
		"zh-CN": "私人",
		"ua-UA": "Приватне",
		"pt-BR": "Privado",
		"tr-TR": "Özel",
		"id-ID": "Pribadi",
	},
	ChooseSpace: {
		"en-UK": "Choose space to make it active for other commands.",
		"ru-RU": "Выберите пространство, чтобы сделать его активным для других команд.",
		"es-ES": "Elija un espacio para activarlo para otros comandos.",
		"fa-IR": "یک فضا را انتخاب کنید تا برای دستورات دیگر فعال شود.",
		"it-IT": "Scegli uno spazio per renderlo attivo per altri comandi.",
		"de-DE": "Wählen Sie einen Bereich, um ihn für andere Befehle zu aktivieren.",
		"fr-FR": "Choisissez un espace pour le rendre actif pour d'autres commandes.",
		"pl-PL": "Wybierz przestrzeń, aby była aktywna dla innych poleceń.",
		"pt-PT": "Escolha um espaço para torná-lo ativo para outros comandos.",
		"ko-KO": "다른 명령에 활성화하려면 공간을 선택하세요.",
		"ja-JP": "他のコマンド用にアクティブにするスペースを選択してください。",
		"zh-CN": "选择一个空间以使其对其他命令有效。",
		"ua-UA": "Виберіть простір, щоб зробити його активним для інших команд.",
		"pt-BR": "Escolha um espaço para torná-lo ativo para outros comandos.",
		"tr-TR": "Diğer komutlar için aktif yapmak üzere bir alan seçin.",
		"id-ID": "Pilih ruang untuk menjadikannya aktif untuk perintah lainnya.",
	},
	FamilyContacts: {
		"en-UK": "Family contacts",
		"ru-RU": "Семейные контакты",
		"es-ES": "Contactos familiares",
		"fa-IR": "مخاطبین خانواده",
		"it-IT": "Contatti familiari",
		"de-DE": "Familienkontakte",
		"fr-FR": "Contacts familiaux",
		"pl-PL": "Kontakty rodzinne",
		"pt-PT": "Contactos familiares",
		"ko-KO": "가족 연락처",
		"ja-JP": "家族の連絡先",
		"zh-CN": "家庭联系人",
		"ua-UA": "Сімейні контакти",
		"pt-BR": "Contatos familiares",
		"tr-TR": "Aile kişileriniz",
		"id-ID": "Kontak keluarga",
	},
	YourContacts: {
		"en-UK": "Your contacts",
		"ru-RU": "Ваши контакты",
		"es-ES": "Tus contactos",
		"fa-IR": "مخاطبین شما",
		"it-IT": "I tuoi contatti",
		"de-DE": "Ihre Kontakte",
		"fr-FR": "Vos contacts",
		"pl-PL": "Twoje kontakty",
		"pt-PT": "Os seus contactos",
		"ko-KO": "당신의 연락처",
		"ja-JP": "あなたの連絡先",
		"zh-CN": "您的联系人",
		"ua-UA": "Ваші контакти",
		"pt-BR": "Seus contatos",
		"tr-TR": "Kişileriniz",
		"id-ID": "Kontak Anda",
	},
	ContactsTitle: {
		"en-UK": "Contacts",
		"ru-RU": "Контакты",
		"es-ES": "Contactos",
		"fa-IR": "مخاطبین",
		"it-IT": "Contatti",
		"de-DE": "Kontakte",
		"fr-FR": "Contacts",
		"pl-PL": "Kontakty",
		"pt-PT": "Contactos",
		"ko-KO": "연락처",
		"ja-JP": "連絡先",
		"zh-CN": "联系人",
		"ua-UA": "Контакти",
		"pt-BR": "Contatos",
		"tr-TR": "Kişiler",
		"id-ID": "Kontak",
	},
	OpenInApp: {
		"en-UK": "Open in app",
		"ru-RU": "Открыть в приложении",
		"fa-IR": "باز کردن در برنامه",
		"es-ES": "Abrir en la aplicación",
		"it-IT": "Apri nell'app",
		"de-DE": "In der App öffnen",
		"fr-FR": "Ouvrir dans l'application",
		"pl-PL": "Otwórz w aplikacji",
		"pt-PT": "Abrir na aplicação",
		"ko-KO": "앱에서 열기",
		"ja-JP": "アプリで開く",
		"zh-CN": "在应用中打开",
		"ua-UA": "Відкрити у додатку",
		"pt-BR": "Abrir no aplicativo",
		"tr-TR": "Uygulamada aç",
		"id-ID": "Buka di aplikasi",
	},
	ManageInApp: {
		"en-UK": "Manage in app",
		"ru-RU": "Управление в приложении",
		"es-ES": "Gestionar en la aplicación",
		"fa-IR": "مدیریت در برنامه",
		"it-IT": "Gestisci nell'app",
		"de-DE": "In der App verwalten",
		"fr-FR": "Gérer dans l'application",
		"pl-PL": "Zarządzaj w aplikacji",
		"pt-PT": "Gerir na aplicação",
		"ko-KO": "앱에서 관리",
		"ja-JP": "アプリで管理",
		"zh-CN": "在应用中管理",
		"ua-UA": "Керуйте в додатку",
		"pt-BR": "Gerenciar no aplicativo",
		"tr-TR": "Uygulamada yönet",
		"id-ID": "Kelola di aplikasi",
	},
	AddContact: {
		"en-UK": "Add contact",
		"ru-RU": "Добавить контакт",
		"es-ES": "Agregar contacto",
		"fa-IR": "اضافه کردن مخاطب",
		"it-IT": "Aggiungi contatto",
		"de-DE": "Kontakt hinzufügen",
		"fr-FR": "Ajouter un contact",
		"pl-PL": "Dodaj kontakt",
		"pt-PT": "Adicionar contacto",
		"ko-KO": "연락처 추가",
		"ja-JP": "連絡先を追加",
		"zh-CN": "添加联系人",
		"ua-UA": "Додати контакт",
		"pt-BR": "Adicionar contato",
		"tr-TR": "Kişi ekle",
		"id-ID": "Tambah kontak",
	},
	AddMember: {
		"en-UK": "Add member",
		"ru-RU": "Добавить участника",
		"es-ES": "Agregar miembro",
		"fa-IR": "اضافه کردن عضو",
		"it-IT": "Aggiungi membro",
		"de-DE": "Mitglied hinzufügen",
		"fr-FR": "Ajouter un membre",
		"pl-PL": "Dodaj członka",
		"pt-PT": "Adicionar membro",
		"ko-KO": "구성원 추가",
		"ja-JP": "メンバーを追加",
		"zh-CN": "添加成员",
		"ua-UA": "Додати учасника",
		"pt-BR": "Adicionar membro",
		"tr-TR": "Üye ekle",
		"id-ID": "Tambah anggota",
	},
	MyContact: {
		"en-UK": "My contact",
		"ru-RU": "Мой контакт",
		"es-ES": "Mi contacto",
		"fa-IR": "مخاطب من",
		"it-IT": "Il mio contatto",
		"de-DE": "Mein Kontakt",
		"fr-FR": "Mon contact",
		"pl-PL": "Mój kontakt",
		"pt-PT": "O meu contacto",
		"ko-KO": "내 연락처",
		"ja-JP": "私の連絡先",
		"zh-CN": "我的联系人",
		"ua-UA": "Мій контакт",
		"pt-BR": "Meu contato",
		"tr-TR": "Benim kişim",
		"id-ID": "Kontak saya",
	},
	GenderTitled: {
		"en-UK": "Gender",
		"ru-RU": "Пол",
		"es-ES": "Género",
		"fa-IR": "جنسیت",
		"it-IT": "Genere",
		"de-DE": "Geschlecht",
		"fr-FR": "Genre",
		"pl-PL": "Płeć",
		"pt-PT": "Género",
		"ko-KO": "성별",
		"ja-JP": "性別",
		"zh-CN": "性别",
		"ua-UA": "Стать",
		"pt-BR": "Gênero",
		"tr-TR": "Cinsiyet",
		"id-ID": "Jenis Kelamin",
	},
	MaleTitled: {
		"en-UK": "Male",
		"ru-RU": "Мужской",
		"es-ES": "Masculino",
		"fa-IR": "مرد",
		"it-IT": "Maschio",
		"de-DE": "Männlich",
		"fr-FR": "Homme",
		"pl-PL": "Mężczyzna",
		"pt-PT": "Masculino",
		"ko-KO": "남성",
		"ja-JP": "男性",
		"zh-CN": "男性",
		"ua-UA": "Чоловік",
		"pt-BR": "Masculino",
		"tr-TR": "Erkek",
		"id-ID": "Pria",
	},
	FemaleTitled: {
		"en-UK": "Female",
		"ru-RU": "Женский",
		"es-ES": "Femenino",
		"fa-IR": "زن",
		"it-IT": "Femmina",
		"de-DE": "Weiblich",
		"fr-FR": "Femme",
		"pl-PL": "Kobieta",
		"pt-PT": "Feminino",
		"ko-KO": "여성",
		"ja-JP": "女性",
		"zh-CN": "女性",
		"ua-UA": "Жінка",
		"pt-BR": "Feminino",
		"tr-TR": "Kadın",
		"id-ID": "Wanita",
	},
	UnknownTitled: {
		"en-UK": "Unknown",
		"ru-RU": "Неизвестно",
		"es-ES": "Desconocido",
		"fa-IR": "ناشناس",
		"it-IT": "Sconosciuto",
		"de-DE": "Unbekannt",
		"fr-FR": "Inconnu",
		"pl-PL": "Nieznany",
		"pt-PT": "Desconhecido",
		"ko-KO": "알 수 없음",
		"ja-JP": "不明",
		"zh-CN": "未知",
		"ua-UA": "Невідомо",
		"pt-BR": "Desconhecido",
		"tr-TR": "Bilinmiyor",
		"id-ID": "Tidak Diketahui",
	},
	UndisclosedTitled: {
		"en-UK": "Undisclosed",
		"ru-RU": "Не разглашено",
		"es-ES": "No revelado",
		"fa-IR": "فاش نشده",
		"it-IT": "Non divulgato",
		"de-DE": "Nicht offengelegt",
		"fr-FR": "Non divulgué",
		"pl-PL": "Nieujawnione",
		"pt-PT": "Não divulgado",
		"ko-KO": "공개되지 않음",
		"ja-JP": "未公開",
		"zh-CN": "未披露",
		"ua-UA": "Не розголошено",
		"pt-BR": "Não divulgado",
		"tr-TR": "Açıklanmamış",
		"id-ID": "Tidak Diungkapkan",
	},
	InviteToJoinSpace: {
		"en-UK": "Invite to join",
		"ru-RU": "Пригласить присоедениться",
		"es-ES": "Invitar a unirse",
		"fa-IR": "دعوت به پیوستن",
		"it-IT": "Invita a unirti",
		"de-DE": "Einladen beizutreten",
		"fr-FR": "Inviter à rejoindre",
		"pl-PL": "Zaproś do dołączenia",
		"pt-PT": "Convidar para juntar-se",
		"ko-KO": "가입 초대",
		"ja-JP": "参加に招待",
		"zh-CN": "邀请加入",
		"ua-UA": "Запросити приєднатися",
		"pt-BR": "Convidar para participar",
		"tr-TR": "Katılmaya davet et",
		"id-ID": "Undang untuk bergabung",
	},
	BackToContacts: {
		"en-UK": "Back to contacts",
		"ru-RU": "Вернуться к контактам",
		"es-ES": "Volver a los contactos",
		"fa-IR": "بازگشت به مخاطبین",
		"it-IT": "Torna ai contatti",
		"de-DE": "Zurück zu Kontakten",
		"fr-FR": "Retour aux contacts",
		"pl-PL": "Powrót do kontaktów",
		"pt-PT": "Voltar aos contatos",
		"ko-KO": "연락처로 돌아가기",
		"ja-JP": "連絡先に戻る",
		"zh-CN": "返回联系人",
		"ua-UA": "Повернутися до контактів",
		"pt-BR": "Voltar para os contatos",
		"tr-TR": "Kişilere geri dön",
		"id-ID": "Kembali ke kontak",
	},
	BackToMembers: {
		"en-UK": "Back to members",
		"ru-RU": "Вернуться к участникам",
		"es-ES": "Volver a los miembros",
		"fa-IR": "بازگشت به اعضا",
		"it-IT": "Torna ai membri",
		"de-DE": "Zurück zu Mitgliedern",
		"fr-FR": "Retour aux membres",
		"pl-PL": "Powrót do członków",
		"pt-PT": "Voltar aos membros",
		"ko-KO": "멤버로 돌아가기",
		"ja-JP": "メンバーに戻る",
		"zh-CN": "返回成员",
		"ua-UA": "Повернутися до учасників",
		"pt-BR": "Voltar para os membros",
		"tr-TR": "Üyelere geri dön",
		"id-ID": "Kembali ke anggota",
	},
	BackBtnTitle: {
		"en-UK": "Back",
		"ru-RU": "Назад",
		"es-ES": "Atrás",
		"fa-IR": "بازگشت",
		"it-IT": "Indietro",
		"de-DE": "Zurück",
		"fr-FR": "Retour",
		"pl-PL": "Wstecz",
		"pt-PT": "Voltar",
		"ko-KO": "뒤로",
		"ja-JP": "戻る",
		"zh-CN": "返回",
		"ua-UA": "Назад",
		"pt-BR": "Voltar",
		"tr-TR": "Geri",
		"id-ID": "Kembali",
	},
	SelectTrackerToAdd: {
		"en-UK": "Select a tracker to add",
		"ru-RU": "Выберите трекер для добавления",
		"es-ES": "Seleccionar un rastreador para agregar",
		"fa-IR": "یک ردیاب برای افزودن انتخاب کنید",
		"it-IT": "Seleziona un tracker da aggiungere",
		"de-DE": "Wählen Sie einen Tracker zum Hinzufügen aus",
		"fr-FR": "Sélectionner un tracker à ajouter",
		"pl-PL": "Wybierz tracker do dodania",
		"pt-PT": "Selecionar um rastreador para adicionar",
		"ko-KO": "추가할 트래커를 선택하세요",
		"ja-JP": "追加するトラッカーを選択",
		"zh-CN": "选择要添加的跟踪器",
		"ua-UA": "Виберіть трекер для додавання",
		"pt-BR": "Selecionar um rastreador para adicionar",
		"tr-TR": "Eklemek için bir izleyici seçin",
		"id-ID": "Pilih pelacak untuk ditambahkan",
	},
	SelectCategoryForNewTracker: {
		"en-UK": "Select a category for the new tracker",
		"ru-RU": "Выберите категорию для нового трекера",
		"es-ES": "Seleccionar una categoría para el nuevo rastreador",
		"fa-IR": "یک دسته برای ردیاب جدید انتخاب کنید",
		"it-IT": "Seleziona una categoria per il nuovo tracker",
		"de-DE": "Wählen Sie eine Kategorie für den neuen Tracker aus",
		"fr-FR": "Sélectionnez une catégorie pour le nouveau traqueur",
		"pl-PL": "Wybierz kategorię dla nowego trackera",
		"pt-PT": "Selecionar uma categoria para o novo rastreador",
		"ko-KO": "새 트래커의 카테고리를 선택하세요",
		"ja-JP": "新しいトラッカーのカテゴリーを選択",
		"zh-CN": "为新跟踪器选择一个类别",
		"ua-UA": "Виберіть категорію для нового трекера",
		"pt-BR": "Selecionar uma categoria para o novo rastreador",
		"tr-TR": "Yeni izleyici için bir kategori seçin",
		"id-ID": "Pilih kategori untuk pelacak baru",
	},
	BackToLists: {
		"en-UK": "Back to lists",
		"ru-RU": "Вернуться к спискам",
		"es-ES": "Volver a las listas",
		"fa-IR": "بازگشت به لیست\u200cها",
		"it-IT": "Torna alle liste",
		"de-DE": "Zurück zu den Listen",
		"fr-FR": "Retour aux listes",
		"pl-PL": "Powrót do list",
		"pt-PT": "Voltar às listas",
		"ko-KO": "목록으로 돌아가기",
		"ja-JP": "リストに戻る",
		"zh-CN": "返回列表",
		"ua-UA": "Повернутися до списків",
		"pt-BR": "Voltar para as listas",
		"tr-TR": "Listelere geri dön",
		"id-ID": "Kembali ke daftar-daftar",
	},
	BackToList: {
		"en-UK": "Back to list",
		"ru-RU": "Вернуться к списку",
		"es-ES": "Volver a la lista",
		"fa-IR": "بازگشت به لیست",
		"it-IT": "Torna alla lista",
		"de-DE": "Zurück zur Liste",
		"fr-FR": "Retour à la liste",
		"pl-PL": "Powrót do listy",
		"pt-PT": "Voltar à lista",
		"ko-KO": "목록으로 돌아가기",
		"ja-JP": "リストに戻る",
		"zh-CN": "返回列表",
		"ua-UA": "Повернутися до списку",
		"pt-BR": "Voltar para a lista",
		"tr-TR": "Listeye geri dön",
		"id-ID": "Kembali ke daftar",
	},
	ClearList: {
		"en-UK": "Clear list",
		"ru-RU": "Очистить список",
		"es-ES": "Limpiar lista",
		"fa-IR": "پاک کردن لیست",
		"it-IT": "Svuota lista",
		"de-DE": "Liste löschen",
		"fr-FR": "Effacer la liste",
		"pl-PL": "Wyczyść listę",
		"pt-PT": "Limpar lista",
		"ko-KO": "목록 지우기",
		"ja-JP": "リストをクリア",
		"zh-CN": "清空列表",
		"ua-UA": "Очистити список",
		"pt-BR": "Limpar lista",
		"tr-TR": "Listeyi temizle",
		"id-ID": "Bersihkan daftar",
	},
	MarkAsDone: {
		"en-UK": "Mark as done",
		"ru-RU": "Отметить как выполненное",
		"es-ES": "Marcar como hecho",
		"fa-IR": "علامت\u200cگذاری به عنوان انجام\u200cشده",
		"it-IT": "Segna come completato",
		"de-DE": "Als erledigt markieren",
		"fr-FR": "Marquer comme fait",
		"pl-PL": "Oznacz jako zrobione",
		"pt-PT": "Marcar como feito",
		"ko-KO": "완료로 표시",
		"ja-JP": "完了としてマークする",
		"zh-CN": "标记为完成",
		"ua-UA": "Позначити як виконане",
		"pt-BR": "Marcar como concluído",
		"tr-TR": "Tamamlandı olarak işaretle",
		"id-ID": "Tandai sebagai selesai",
	},
	AddStandard: {
		"en-UK": "Add standard",
		"ru-RU": "Добавить стандартный",
		"es-ES": "Agregar estándar",
		"fa-IR": "افزودن استاندارد",
		"it-IT": "Aggiungere standard",
		"de-DE": "Standard hinzufügen",
		"fr-FR": "Ajouter une norme",
		"pl-PL": "Dodaj standard",
		"pt-PT": "Adicionar padrão",
		"ko-KO": "표준 추가",
		"ja-JP": "標準を追加",
		"zh-CN": "添加标准",
		"ua-UA": "Додати стандарт",
		"pt-BR": "Adicionar padrão",
		"tr-TR": "Standart ekle",
		"id-ID": "Tambahkan standar",
	},
	DeleteItems: {
		"en-UK": "Delete items",
		"ru-RU": "Удалить элементы",
		"fa-IR": "حذف موارد",
		"it-IT": "Elimina elementi",
		"de-DE": "Elemente löschen",
		"fr-FR": "Supprimer les éléments",
		"pl-PL": "Usuń elementy",
		"pt-PT": "Eliminar itens",
		"ko-KO": "항목 삭제",
		"ja-JP": "項目を削除",
		"zh-CN": "删除项目",
		"ua-UA": "Видалити елементи",
		"pt-BR": "Excluir itens",
		"tr-TR": "Öğeleri sil",
		"id-ID": "Hapus item",
	},
	YouCanAddItemBySendingMessage: {
		"en-UK": "You can add items to this list by sending a message to me.",
		"ru-RU": "Вы можете добавлять элементы в этот список, отправив мне сообщение.",
		"es-ES": "Puedes agregar elementos a esta lista enviándome un mensaje.",
		"fa-IR": "می\u200cتوانید با ارسال یک پیام به من، موارد را به این لیست اضافه کنید.",
		"it-IT": "Puoi aggiungere elementi a questa lista inviandomi un messaggio.",
		"de-DE": "Sie können Elemente zu dieser Liste hinzufügen, indem Sie mir eine Nachricht senden.",
		"fr-FR": "Vous pouvez ajouter des éléments à cette liste en m'envoyant un message.",
		"pl-PL": "Możesz dodać elementy do tej listy, wysyłając mi wiadomość.",
		"pt-PT": "Pode adicionar items a esta lista enviando-me uma mensagem.",
		"ko-KO": "저에게 메시지를 보내서 이 목록에 항목들을 추가할 수 있습니다.",
		"ja-JP": "私にメッセージを送ることで、このリストに項目を追加できます。",
		"zh-CN": "您可以通过向我发送消息将项目添加到此列表。",
		"ua-UA": "Ви можете додати елементи до цього списку, надіславши мені повідомлення.",
		"pt-BR": "Você pode adicionar itens a esta lista enviando uma mensagem para mim.",
		"tr-TR": "Bu listeye bir mesaj göndererek öğeler ekleyebilirsiniz.",
		"id-ID": "Anda dapat menambahkan item ke daftar ini dengan mengirimkan pesan kepada saya.",
	},
	NoItemsInTheListYet: {
		"en-UK": "No items in the list yet.",
		"ru-RU": "Пока нет элементов в списке.",
		"es-ES": "Todavía no hay elementos en la lista.",
		"fa-IR": "هنوز هیچ آیتمی در لیست وجود ندارد.",
		"it-IT": "Non ci sono ancora elementi nella lista.",
		"de-DE": "Noch keine Elemente in der Liste.",
		"fr-FR": "Pas encore d'éléments dans la liste.",
		"pl-PL": "Jeszcze brak elementów na liście.",
		"pt-PT": "Ainda não existem itens na lista.",
		"ko-KO": "목록에 아직 항목이 없습니다.",
		"ja-JP": "リストに項目がまだありません。",
		"zh-CN": "列表中尚无项目。",
		"ua-UA": "Ще немає елементів у списку.",
		"pt-BR": "Ainda não há itens na lista.",
		"tr-TR": "Henüz listede öğe yok.",
		"id-ID": "Belum ada item dalam daftar.",
	},
	FamilyList: {
		"en-UK": "Family list",
		"ru-RU": "Семейный список",
		"es-ES": "Lista familiar",
		"fa-IR": "لیست خانواده",
		"it-IT": "Elenco familiare",
		"de-DE": "Familienliste",
		"fr-FR": "Liste familiale",
		"pl-PL": "Lista rodzinna",
		"pt-PT": "Lista familiar",
		"ko-KO": "가족 목록",
		"ja-JP": "家族リスト",
		"zh-CN": "家庭清单",
		"ua-UA": "Сімейний список",
		"pt-BR": "Lista da família",
		"tr-TR": "Aile listesi",
		"id-ID": "Daftar keluarga",
	},
	PrivateList: {
		"en-UK": "Private list",
		"ru-RU": "Личный список",
		"es-ES": "Lista privada",
		"fa-IR": "لیست خصوصی",
		"it-IT": "Lista personale",
		"de-DE": "Private Liste",
		"fr-FR": "Liste privée",
		"pl-PL": "Lista prywatna",
		"pt-PT": "Lista privada",
		"ko-KO": "개인 목록",
		"ja-JP": "プライベートリスト",
		"zh-CN": "私人清单",
		"ua-UA": "Особистий список",
		"pt-BR": "Lista privada",
		"tr-TR": "Özel liste",
		"id-ID": "Daftar pribadi",
	},
	Refresh: {
		"en-UK": "Refresh",
		"ru-RU": "Обновить",
		"es-ES": "Refrescar",
		"fa-IR": "بروزرسانی",
		"it-IT": "Aggiorna",
		"de-DE": "Aktualisieren",
		"fr-FR": "Rafraîchir",
		"pl-PL": "Odśwież",
		"pt-PT": "Atualizar",
		"ko-KO": "새로고침",
		"ja-JP": "リフレッシュ",
		"zh-CN": "刷新",
		"ua-UA": "Оновити",
		"pt-BR": "Atualizar",
		"tr-TR": "Yenile",
		"id-ID": "Segarkan",
	},
	AdviseToUseTelegramForTgUsers: {
		"en-UK": `If the person you want to add uses telegram we advise to select "Choose Telegram User".
Otherwise you can add them manually in Sneat.app.`,
		"ru-RU": `Если человек, которого вы хотите добавить, использует Telegram, 
		мы рекомендуем выбрать "Выбрать пользователя Telegram". 
		В противном случае вы можете добавить их вручную в Sneat.app.`,
		"es-ES": `Si la persona que deseas agregar usa Telegram, te recomendamos seleccionar "Elegir usuario de Telegram".
		De lo contrario, puedes agregarlos manualmente en Sneat.app.`,
		"fa-IR": `اگر شخصی که می\u200cخواهید اضافه کنید از تلگرام استفاده می\u200cکند، 
ما توصیه می\u200cکنیم گزینه "انتخاب کاربر تلگرام" را انتخاب کنید. 
در غیر این صورت می\u200cتوانید آنها را به صورت دستی در Sneat.app اضافه کنید.`,
		"it-IT": `Se la persona che vuoi aggiungere utilizza telegram, ti consigliamo di selezionare "Scegli utente Telegram".
		In caso contrario, puoi aggiungerli manualmente in Sneat.app.`,
		"de-DE": `Wenn die Person, die Sie hinzufügen möchten, Telegram verwendet, empfehlen wir, "Telegram-Benutzer auswählen" auszuwählen.
		Andernfalls können Sie sie manuell in Sneat.app hinzufügen.`,
		"fr-FR": `Si la personne que vous souhaitez ajouter utilise Telegram, nous vous conseillons de sélectionner "Choisir un utilisateur Telegram".
		Sinon, vous pouvez les ajouter manuellement sur Sneat.app.`,
		"pl-PL": `Jeśli osoba, którą chcesz dodać, korzysta z Telegrama, radzimy wybrać "Wybierz użytkownika Telegrama".
		W przeciwnym razie możesz dodać ich ręcznie w aplikacji Sneat.app.`,
		"pt-PT": `Se a pessoa que deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente na Sneat.app.`,
		"ko-KO": `추가하려는 사람이 Telegram을 사용한다면 "Telegram 사용자 선택"을 선택하시길 권장드립니다.
		그렇지 않을 경우 Sneat.app에서 수동으로 추가할 수 있습니다.`,
		"ja-JP": `追加したい相手がTelegramを使用している場合、「Telegram ユーザーを選択」を選択することをお勧めします。
		それ以外の場合は、Sneat.app で手動で追加することができます。`,
		"zh-CN": `如果您想添加的人使用 Telegram，我们建议选择“选择 Telegram 用户”。
		否则，您可以在 Sneat.app 中手动添加他们。`,
		"ua-UA": `Якщо людина, яку ви хочете додати, використовує Telegram, 
		ми радимо вибрати "Вибрати користувача Telegram". 
		Інакше, ви можете додати їх вручну на Sneat.app.`,
		"pt-BR": `Se a pessoa que você deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente no Sneat.app.`,
		"tr-TR": `Eklemek istediğiniz kişi Telegram kullanıyorsa, "Telegram Kullanıcısını Seç" seçeneğini seçmenizi öneririz.
		Aksi takdirde Sneat.app'te manuel olarak ekleyebilirsiniz.`,
		"id-ID": `Jika orang yang ingin Anda tambahkan menggunakan Telegram, kami sarankan untuk memilih "Pilih Pengguna Telegram".
		Jika tidak, Anda dapat menambahkannya secara manual di Sneat.app.`,
	},
	ChooseTelegramUser: {
		"en-UK": "Choose Telegram user",
		"ru-RU": "Выбрать пользователя Telegram",
		"es-ES": "Elegir usuario de Telegram",
		"fa-IR": "انتخاب کاربر تلگرام",
		"it-IT": "Scegli utente Telegram",
		"de-DE": "Telegram-Benutzer auswählen",
		"fr-FR": "Choisir un utilisateur Telegram",
		"pl-PL": "Wybierz użytkownika Telegrama",
		"pt-PT": "Escolher usuário do Telegram",
		"ko-KO": "Telegram 사용자 선택",
		"ja-JP": "Telegram ユーザーを選択",
		"zh-CN": "选择 Telegram 用户",
		"ua-UA": "Вибрати користувача Telegram",
		"pt-BR": "Escolher usuário do Telegram",
		"tr-TR": "Telegram Kullanıcısını Seç",
		"id-ID": "Pilih Pengguna Telegram",
	},
	AddManuallyInSneatApp: {
		"en-UK": "Add manually in Sneat.app",
		"ru-RU": "Добавить вручную в Sneat.app",
		"es-ES": "Agregar manualmente en Sneat.app",
		"fa-IR": "به صورت دستی در Sneat.app اضافه کنید",
		"it-IT": "Aggiungi manualmente in Sneat.app",
		"de-DE": "Manuell in Sneat.app hinzufügen",
		"fr-FR": "Ajouter manuellement dans Sneat.app",
		"pl-PL": "Dodaj ręcznie w Sneat.app",
		"pt-PT": "Adicionar manualmente no Sneat.app",
		"ko-KO": "Sneat.app에서 수동으로 추가",
		"ja-JP": "Sneat.app で手動で追加",
		"zh-CN": "在 Sneat.app 中手动添加",
		"ua-UA": "Додати вручну на Sneat.app",
		"pt-BR": "Adicionar manualmente no Sneat.app",
		"tr-TR": "Sneat.app'te manuel olarak ekle",
		"id-ID": "Tambahkan secara manual di Sneat.app",
	},
	CancelAddingMember: {
		"en-UK": "Cancel adding member",
		"ru-RU": "Отменить добавление участника",
		"es-ES": "Cancelar la adición de miembro",
		"fa-IR": "لغو افزودن عضو",
		"it-IT": "Annulla l'aggiunta di un membro",
		"de-DE": "Hinzufügen eines Mitglieds abbrechen",
		"fr-FR": "Annuler l'ajout d'un membre",
		"pl-PL": "Anuluj dodawanie członka",
		"pt-PT": "Cancelar a adição de membro",
		"ko-KO": "멤버 추가 취소",
		"ja-JP": "メンバー追加をキャンセル",
		"zh-CN": "取消添加成员",
		"ua-UA": "Скасувати додавання учасника",
		"pt-BR": "Cancelar a adição de membro",
		"tr-TR": "Üye eklemeyi iptal et",
		"id-ID": "Batalkan menambahkan anggota",
	},
	CancelAddingContact: {
		"en-UK": "Cancel adding contact",
		"ru-RU": "Отменить добавление контакта",
		"es-ES": "Cancelar la adición de contacto",
		"fa-IR": "لغو افزودن مخاطب",
		"it-IT": "Annulla l'aggiunta di un contatto",
		"de-DE": "Hinzufügen eines Kontakts abbrechen",
		"fr-FR": "Annuler l'ajout d'un contact",
		"pl-PL": "Anuluj dodawanie kontaktu",
		"pt-PT": "Cancelar a adição de contato",
		"ko-KO": "연락처 추가 취소",
		"ja-JP": "連絡先追加をキャンセル",
		"zh-CN": "取消添加联系人",
		"ua-UA": "Скасувати додавання контакту",
		"pt-BR": "Cancelar a adição de contato",
		"tr-TR": "Kişi eklemeyi iptal et",
		"id-ID": "Batalkan menambahkan kontak",
	},
	FamilyDebts: {
		"en-UK": "Family debts",
		"ru-RU": "Семейные долги",
		"es-ES": "Deudas familiares",
		"fa-IR": "بدهی های خانواده",
		"it-IT": "Debiti familiari",
		"de-DE": "Familienschulden",
		"fr-FR": "Dettes familiales",
		"pl-PL": "Długi rodzinne",
		"pt-PT": "Dívidas da família",
		"ko-KO": "가족 부채",
		"ja-JP": "家族の借金",
		"zh-CN": "家庭债务",
		"ua-UA": "Сімейні борги",
		"pt-BR": "Dívidas familiares",
		"tr-TR": "Aile borçları",
		"id-ID": "Hutang keluarga",
	},
	DebtsRelatedContacts: {
		"en-UK": "Debts related contacts",
		"ru-RU": "Контакты, связанные с долгами",
		"es-ES": "Contactos relacionados con deudas",
		"fa-IR": "ارتباطات مربوط به بدهی\u200cها",
		"it-IT": "Contatti legati ai debiti",
		"de-DE": "Schuldenbezogene Kontakte",
		"fr-FR": "Contacts liés aux dettes",
		"pl-PL": "Kontakty związane z długami",
		"pt-PT": "Contactos relacionados com dívidas",
		"ko-KO": "부채 관련 연락처",
		"ja-JP": "借金関連の連絡先",
		"zh-CN": "债务相关联系人",
		"ua-UA": "Контакти, пов’язані з боргами",
		"pt-BR": "Contatos relacionados a dívidas",
		"tr-TR": "Borçlarla ilgili kişiler",
		"id-ID": "Kontak yang terkait dengan hutang",
	},
	BackToDebtsMenu: {
		"en-UK": "Back to debts menu",
		"ru-RU": "Вернуться в меню долгов",
		"es-ES": "Volver al menú de deudas",
		"fa-IR": "بازگشت به منوی بدهی\u200cها",
		"it-IT": "Torna al menu dei debiti",
		"de-DE": "Zurück zum Schuldenmenü",
		"fr-FR": "Retour au menu des dettes",
		"pl-PL": "Powrót do menu długów",
		"pt-PT": "Voltar ao menu das dívidas",
		"ko-KO": "부채 메뉴로 돌아가기",
		"ja-JP": "借金メニューに戻る",
		"zh-CN": "返回债务菜单",
		"ua-UA": "Повернутися до меню боргів",
		"pt-BR": "Voltar ao menu de dívidas",
		"tr-TR": "Borçlar menüsüne geri dön",
		"id-ID": "Kembali ke menu hutang",
	},
}
