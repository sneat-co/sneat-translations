package trans

/*
Expected IDs in proper order of locale keys in var TRANS, use it as a reference for all values:
- arEG:
- deDE
- enUK:
- enUS:
- esES:
- faIR:
- frFR:
- idID:
- itIT:
- jaJP:
- koKR:
- plPL:
- ptBR:
- ruRU:
- trTR:
- ukUA:
- uzUZ:
- zhCN:
*/

// TRANS - translation string
var TRANS = map[string]map[string]string{
	"EXAMPLE": {
		deDE: "BEISPIEL",
		enUK: "SAMPLE",
		enUS: "SAMPLE",
		esES: "EJEMPLO",
		faIR: "نمونه",
		frFR: "EXEMPLE",
		idID: "CONTOH",
		itIT: "ESEMPIO",
		jaJP: "例",
		koKR: "예",
		plPL: "PRZYKŁAD",
		ptBR: "EXEMPLO",
		ruRU: "ПРИМЕР",
		trTR: "ÖRNEK",
		ukUA: "ПРИКЛАД",
		uzUZ: "MISOL",
		zhCN: "例子",
	},

	HowdyUser: {
		deDE: "Hallo {USER_NAME}!",
		enUK: "Howdy {USER_NAME}!",
		enUS: "Howdy {USER_NAME}!",
		esES: "¡Hola {USER_NAME}!",
		faIR: "سلام {USER_NAME}!",
		frFR: "Salut {USER_NAME} !",
		idID: "Halo {USER_NAME}!",
		itIT: "Ciao {USER_NAME}!",
		jaJP: "やあ {USER_NAME}!",
		koKR: "안녕 {USER_NAME}!",
		plPL: "Cześć {USER_NAME}!",
		ptBR: "Olá {USER_NAME}!",
		ptPT: "Olá {USER_NAME}!",
		ruRU: "Привет {USER_NAME}!",
		trTR: "Merhaba {USER_NAME}!",
		ukUA: "Привіт {USER_NAME}!",
		uzUZ: "Salom {USER_NAME}!",
		zhCN: "你好 {USER_NAME}！",
	},

	ButtonAdd: {
		deDE: "Hinzufügen",
		enUK: "Add",
		enUS: "Add",
		esES: "Añadir",
		faIR: "افزودن",
		frFR: "Ajouter",
		idID: "Tambah",
		itIT: "Aggiungi",
		jaJP: "追加",
		koKR: "추가",
		plPL: "Dodaj",
		ptBR: "Adicionar",
		ruRU: "Добавить",
		trTR: "Ekle",
		ukUA: "Додати",
		uzUZ: "Qo'shish",
		zhCN: "添加",
	},
	ButtonRemove: {
		deDE: "Entfernen",
		enUK: "Remove",
		enUS: "Remove",
		esES: "Eliminar",
		faIR: "حذف",
		frFR: "Supprimer",
		idID: "Hapus",
		itIT: "Rimuovi",
		jaJP: "削除",
		koKR: "제거",
		plPL: "Usuń",
		ptBR: "Remover",
		ptPT: "Remover",
		ruRU: "Удалить",
		trTR: "Kaldır",
		ukUA: "Видалити",
		uzUZ: "Olib tashlash",
		zhCN: "移除",
	},

	"Jan": {
		deDE: "Jan",
		enUK: "Jan",
		enUS: "Jan",
		esES: "Ene",
		faIR: "ژانویه",
		frFR: "Jan",
		idID: "Jan",
		itIT: "Gen",
		jaJP: "1月",
		koKR: "1월",
		plPL: "Sty",
		ptBR: "Jan",
		ptPT: "Jan",
		ruRU: "Янв.",
		trTR: "Oca",
		ukUA: "Січ",
		uzUZ: "Yan",
		zhCN: "一月",
	},

	"Feb": {
		deDE: "Feb",
		enUK: "Feb",
		enUS: "Feb",
		esES: "Feb",
		faIR: "فوریه",
		frFR: "Fév",
		idID: "Feb",
		itIT: "Feb",
		jaJP: "2月",
		koKR: "2월",
		plPL: "Lut",
		ptBR: "Fev",
		ptPT: "Fev",
		ruRU: "Фев.",
		trTR: "Şub",
		ukUA: "Лют",
		uzUZ: "Fev",
		zhCN: "二月",
	},

	"Mar": {
		deDE: "Mär",
		enUK: "Mar",
		enUS: "Mar",
		esES: "Mar",
		faIR: "مارس",
		frFR: "Mars",
		idID: "Mar",
		itIT: "Mar",
		jaJP: "3月",
		koKR: "3월",
		plPL: "Mar",
		ptBR: "Março",
		ptPT: "Março",
		ruRU: "Мврт",
		trTR: "Mar",
		ukUA: "Бер",
		uzUZ: "Mar",
		zhCN: "三月",
	},

	"Apr": {
		deDE: "Apr",
		enUK: "Apr",
		enUS: "Apr",
		esES: "Abr",
		faIR: "آوریل",
		frFR: "Avril",
		idID: "Apr",
		itIT: "Apr",
		jaJP: "4月",
		koKR: "4월",
		plPL: "Kwi",
		ptBR: "Abr",
		ptPT: "Abr",
		ruRU: "Апр.",
		trTR: "Nis",
		ukUA: "Кві",
		uzUZ: "Apr",
		zhCN: "四月",
	},

	"May": {
		deDE: "Mai",
		enUK: "May",
		enUS: "May",
		esES: "May",
		faIR: "مه",
		frFR: "Mai",
		idID: "Mei",
		itIT: "Mag",
		jaJP: "5月",
		koKR: "5월",
		plPL: "Maj",
		ptBR: "Maio",
		ptPT: "Maio",
		ruRU: "Май",
		trTR: "May",
		ukUA: "Тра",
		uzUZ: "May",
		zhCN: "五月",
	},

	"Jun": {
		deDE: "Jun",
		enUK: "June",
		enUS: "June",
		esES: "Jun",
		faIR: "ژوئن",
		frFR: "Juin",
		idID: "Jun",
		itIT: "Giu",
		jaJP: "6月",
		koKR: "6월",
		plPL: "Cze",
		ptBR: "Junho",
		ptPT: "Junho",
		ruRU: "Июнь",
		trTR: "Haz",
		ukUA: "Чер",
		uzUZ: "Iyun",
		zhCN: "六月",
	},

	"Jul": {
		deDE: "Jul",
		enUK: "July",
		enUS: "July",
		esES: "Jul",
		faIR: "ژوئیه",
		frFR: "Juil",
		idID: "Jul",
		itIT: "Lug",
		jaJP: "7月",
		koKR: "7월",
		plPL: "Lip",
		ptBR: "Julho",
		ptPT: "Julho",
		ruRU: "Июль",
		trTR: "Tem",
		ukUA: "Лип",
		uzUZ: "Iyul",
		zhCN: "七月",
	},

	"Aug": {
		deDE: "Aug",
		enUK: "Aug",
		enUS: "Aug",
		esES: "Ago",
		faIR: "اوت",
		frFR: "Août",
		idID: "Agu",
		itIT: "Ago",
		jaJP: "8月",
		koKR: "8월",
		plPL: "Sie",
		ptBR: "Ago",
		ptPT: "Ago",
		ruRU: "Авг.",
		trTR: "Ağu",
		ukUA: "Сер",
		uzUZ: "Avg",
		zhCN: "八月",
	},

	"Sep": {
		deDE: "Sep",
		enUK: "Sep",
		enUS: "Sep",
		esES: "Sep",
		faIR: "سپتامبر",
		frFR: "Sep",
		idID: "Sep",
		itIT: "Sett",
		jaJP: "9月",
		koKR: "9월",
		plPL: "Wrz",
		ptBR: "Set",
		ptPT: "Set",
		ruRU: "Сен.",
		trTR: "Eyl",
		ukUA: "Вер",
		uzUZ: "Sen",
		zhCN: "九月",
	},

	"Oct": {
		deDE: "Okt",
		enUK: "Oct",
		enUS: "Oct",
		esES: "Oct",
		faIR: "اکتبر",
		frFR: "Oct",
		idID: "Okt",
		itIT: "Ott",
		jaJP: "10月",
		koKR: "10월",
		plPL: "Paź",
		ptBR: "Out",
		ptPT: "Out",
		ruRU: "Окт.",
		trTR: "Eki",
		ukUA: "Жов",
		uzUZ: "Okt",
		zhCN: "十月",
	},

	"Nov": {
		deDE: "Nov",
		enUK: "Nov",
		enUS: "Nov",
		esES: "Nov",
		faIR: "نوامبر",
		frFR: "Nov",
		idID: "Nov",
		itIT: "Nov",
		jaJP: "11月",
		koKR: "11월",
		plPL: "Lis",
		ptBR: "Nov",
		ptPT: "Nov",
		ruRU: "Нбр.",
		trTR: "Kas",
		ukUA: "Лис",
		uzUZ: "Noy",
		zhCN: "十一月",
	},

	"Dec": {
		deDE: "Dez",
		enUK: "Dec",
		enUS: "Dec",
		esES: "Dic",
		faIR: "دسامبر",
		frFR: "Déc",
		idID: "Des",
		itIT: "Dic",
		jaJP: "12月",
		koKR: "12월",
		plPL: "Gru",
		ptBR: "Dez",
		ptPT: "Dez",
		ruRU: "Дек.",
		trTR: "Ara",
		ukUA: "Гру",
		uzUZ: "Dek",
		zhCN: "十二月",
	},
	COMMAND_START: {
		deDE: "start",
		enUK: "start",
		enUS: "start",
		esES: "inicio",
		faIR: "شروع",
		frFR: "démarrer",
		idID: "mulai",
		itIT: "inizio",
		jaJP: "スタート",
		koKR: "시작",
		plPL: "start",
		ptBR: "iniciar",
		ruRU: "старт",
		trTR: "başlat",
		ukUA: "старт",
		uzUZ: "boshlash",
		zhCN: "开始",
	},
	COMMAND_MENU: {
		deDE: "menu",
		enUK: "menu",
		enUS: "menu",
		esES: "menú",
		faIR: "منو",
		frFR: "menu",
		idID: "menu",
		itIT: "menu",
		jaJP: "メニュー",
		koKR: "메뉴",
		plPL: "menu",
		ptBR: "menu",
		ruRU: "меню",
		trTR: "menü",
		ukUA: "меню",
		uzUZ: "menyu",
		zhCN: "菜单",
	},
	COMMAND_GAVE: {
		deDE: "verleihen",
		enUK: "gave",
		enUS: "gave",
		esES: "prestado_a_ti",
		faIR: "قرض_دادن",
		frFR: "donné",
		idID: "memberi",
		itIT: "debito",
		jaJP: "貸した",
		koKR: "주었다",
		plPL: "dał",
		ptBR: "emprestou",
		ruRU: "дать",
		trTR: "verdi",
		ukUA: "дати",
		uzUZ: "berdi",
		zhCN: "给予",
	},
	COMMAND_GOT: {
		deDE: "anleihen",
		enUK: "got",
		enUS: "got",
		esES: "prestado_por_ti",
		faIR: "قرض_گرفتن",
		frFR: "reçu",
		idID: "mendapat",
		itIT: "credito",
		jaJP: "借りた",
		koKR: "받았다",
		plPL: "dostał",
		ptBR: "recebeu",
		ruRU: "взять",
		trTR: "aldı",
		ukUA: "взяти",
		uzUZ: "oldi",
		zhCN: "得到",
	},
	COMMAND_RETURNED: {
		deDE: "beglichen",
		enUK: "return",
		enUS: "return",
		esES: "devuelto",
		faIR: "بازگردانده_شده",
		frFR: "retourné",
		idID: "kembali",
		itIT: "rientra",
		jaJP: "返済",
		koKR: "반환",
		plPL: "zwrócił",
		ptBR: "devolveu",
		ruRU: "вернуть",
		trTR: "iade",
		ukUA: "повернути",
		uzUZ: "qaytardi",
		zhCN: "归还",
	},
	COMMAND_BALANCE: {
		deDE: "ausstehend",
		enUK: "balance",
		enUS: "balance",
		esES: "balance",
		faIR: "تراز",
		frFR: "solde",
		idID: "saldo",
		itIT: "bilancio",
		jaJP: "残高",
		koKR: "잔액",
		plPL: "saldo",
		ptBR: "saldo",
		ruRU: "баланс",
		trTR: "bakiye",
		ukUA: "баланс",
		uzUZ: "balans",
		zhCN: "余额",
	},
	COMMAND_HISTORY: {
		deDE: "verlauf",
		enUK: "history",
		enUS: "history",
		esES: "cronología",
		faIR: "سوابق",
		frFR: "historique",
		idID: "riwayat",
		itIT: "cronologia",
		jaJP: "履歴",
		koKR: "기록",
		plPL: "historia",
		ptBR: "histórico",
		ptPT: "histórico",
		ruRU: "история",
		trTR: "geçmiş",
		ukUA: "історія",
		uzUZ: "tarix",
		zhCN: "历史",
	},
	COMMAND_SETTINGS: {
		deDE: "einstellungen",
		enUK: "settings",
		enUS: "settings",
		esES: "ajustes",
		faIR: "تنظیمات",
		frFR: "paramètres",
		idID: "pengaturan",
		itIT: "impostazioni",
		jaJP: "設定",
		koKR: "설정",
		plPL: "ustawienia",
		ptBR: "configurações",
		ruRU: "настройки",
		trTR: "ayarlar",
		ukUA: "налаштування",
		uzUZ: "sozlamalar",
		zhCN: "设置",
	},
	COMMAND_HELP: {
		deDE: "hilfe",
		enUK: "help",
		enUS: "help",
		esES: "ayuda",
		faIR: "کمک",
		frFR: "aide",
		idID: "bantuan",
		itIT: "aiuto",
		jaJP: "ヘルプ",
		koKR: "도움말",
		plPL: "pomoc",
		ptBR: "ajuda",
		ruRU: "помощь",
		trTR: "yardım",
		ukUA: "допомога",
		uzUZ: "yordam",
		zhCN: "帮助",
	},
	COMMAND_CANCEL: {
		deDE: "abbrechen",
		enUK: "cancel",
		enUS: "cancel",
		esES: "cancelar",
		faIR: "کنسل",
		frFR: "annuler",
		idID: "batal",
		itIT: "annulla",
		jaJP: "キャンセル",
		koKR: "취소",
		plPL: "anuluj",
		ptBR: "cancelar",
		ruRU: "/отменить",
		trTR: "iptal",
		ukUA: "скасувати",
		uzUZ: "bekor",
		zhCN: "取消",
	},
	COMMAND_CLEAR: {
		deDE: "leeren",
		enUK: "clear",
		enUS: "clear",
		esES: "borrar",
		faIR: "پاک_کردن",
		frFR: "effacer",
		idID: "bersihkan",
		itIT: "chiaro",
		jaJP: "クリア",
		koKR: "지우기",
		plPL: "wyczyść",
		ptBR: "limpar",
		ruRU: "очистить",
		trTR: "temizle",
		ukUA: "очистити",
		uzUZ: "tozalash",
		zhCN: "清除",
	},
	"Ads": {
		enUK: "Ads",
		ruRU: "Реклама",
	},
	" and ": {
		deDE: " und ",
		enUK: " and ",
		enUS: " and ",
		esES: " y ",
		faIR: " و ",
		frFR: " et ",
		idID: " dan ",
		itIT: " e ",
		jaJP: " と ",
		koKR: " 그리고 ",
		plPL: " i ",
		ptBR: " e ",
		ruRU: " и ",
		trTR: " ve ",
		ukUA: " і ",
		uzUZ: " va ",
		zhCN: " 和 ",
	},
	"MessageTextOopsSomethingWentWrong": {
		deDE: "Ups, etwas ist schiefgelaufen... \xF0\x9F\x98\xB3",
		enUK: "Oops, something went wrong... \xF0\x9F\x98\xB3",
		enUS: "Oops, something went wrong... \xF0\x9F\x98\xB3",
		esES: "Ops,  algo ha salido mal... \xF0\x9F\x98\xB3",
		faIR: "اوه، یک جای کار مشکل دارد ...  \xF0\x9F\x98\xB3",
		frFR: "Oups, quelque chose s'est mal passé... \xF0\x9F\x98\xB3",
		idID: "Ups, ada yang salah... \xF0\x9F\x98\xB3",
		itIT: "Ops, qualcosa e' andato storto... \xF0\x9F\x98\xB3",
		jaJP: "おっと、何か問題が発生しました... \xF0\x9F\x98\xB3",
		koKR: "이런, 문제가 발생했습니다... \xF0\x9F\x98\xB3",
		plPL: "Ups, coś poszło nie tak... \xF0\x9F\x98\xB3",
		ptBR: "Ops, algo deu errado... \xF0\x9F\x98\xB3",
		ruRU: "Упс, что-то пошло не так... \xF0\x9F\x98\xB3",
		trTR: "Hata, bir şeyler yanlış gitti... \xF0\x9F\x98\xB3",
		ukUA: "Упс, щось пішло не так... \xF0\x9F\x98\xB3",
		uzUZ: "Xatolik, nimadir noto'g'ri ketdi... \xF0\x9F\x98\xB3",
		zhCN: "哎呀，出了点问题... \xF0\x9F\x98\xB3",
	},
	MESSAGE_TEXT_ASK_DUE: {
		deDE: "Wann ist die Schuld fällig?",
		enUK: "When is the due date?",
		enUS: "When is the due date?",
		esES: "¿Cuándo es la fecha de devolución?",
		faIR: "سررسید چه زمانی است؟",
		frFR: "Quelle est la date d'échéance?",
		idID: "Kapan tanggal jatuh tempo?",
		itIT: "Data di scadenza?",
		jaJP: "期日はいつですか？",
		koKR: "만기일은 언제입니까?",
		plPL: "Kiedy jest termin płatności?",
		ptBR: "Qual é a data de vencimento?",
		ruRU: "Когда дата возврата?",
		trTR: "Vade tarihi ne zaman?",
		ukUA: "Коли дата повернення?",
		uzUZ: "To'lov muddati qachon?",
		zhCN: "到期日期是什么时候？",
	},
	MESSAGE_TEXT_ASK_DATE_TO_REMIND: {
		deDE: `Sende mir das Datum, an welches du <b>erneut</b> erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		enUK: `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		enUS: `To set date for next reminder please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
    <i>20.01.2017</i>`,

		esES: `Para establecer la fecha recordatoria escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
    <i>20.01.2017</i>`,

		faIR: `لطفاً برای تنظیم یادآور بعدی آنرا با متنی با این فرمت ارسال نمایید. <i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 15 خرداد 1396 ثبت کنید:
    <i>15.03.1396</i>`,

		frFR: `Pour définir la date du prochain rappel, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
    <i>20.01.2017</i>`,

		idID: `Untuk mengatur tanggal pengingat berikutnya, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
    <i>20.01.2017</i>`,

		itIT: `Per impostare la data per il promemoria successivo invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
    <i>20.01.2017</i>`,

		jaJP: `次のリマインダーの日付を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
    <i>20.01.2017</i>`,

		koKR: `다음 알림의 날짜를 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
    <i>20.01.2017</i>`,

		plPL: `Aby ustawić datę następnego przypomnienia, wyślij ją jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
    <i>20.01.2017</i>`,

		ptBR: `Para definir a data do próximo lembrete, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
    <i>20.01.2017</i>`,

		ruRU: `Чтобы задать дату напопинания напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г. отправьте:
    <i>20.01.2017</i>`,

		trTR: `Bir sonraki hatırlatma için tarih belirlemek üzere <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
    <i>20.01.2017</i>`,

		ukUA: `Щоб встановити дату наступного нагадування, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
    <i>20.01.2017</i>`,

		uzUZ: `Keyingi eslatma uchun sanani belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
    <i>20.01.2017</i>`,

		zhCN: `要设置下一次提醒的日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
    <i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_ASK_DUE_DATE: {
		deDE: `Sende mir das Datum, an welches du erinnert werden möchtest, in der Form <i>DD.MM.YEAR</i>.
<b>Zum Beispiel</b> für den 20. Januar 2017, schreibe:
    <i>20.01.2017</i>`,

		enUK: `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
<i>20.01.2017</i>`,

		enUS: `To set due date please send it as a text in format of <i>DD.MM.YEAR</i>.
<b>For example</b> for 20th of January 2017 submit:
<i>20.01.2017</i>`,

		esES: `Para establecer la fecha de devolución escribela en el siguiente formato <i>DD.MM.AÑO</i>.
<b>For example</b> para 20 de Enero 2017 envia:
<i>20.01.2017</i>`,

		faIR: `لطفاً برای تنظیم تاریخ سررسید این فرمت را رعایت فرمایید.<i>روز.ماه.سال</i>.
<b>برای مثال</b> برای 20 ژانویه 2017 ثبت کنید:
<i>20.01.2017</i>`,

		frFR: `Pour définir la date d'échéance, veuillez l'envoyer sous forme de texte au format <i>JJ.MM.ANNÉE</i>.
<b>Par exemple</b> pour le 20 janvier 2017, soumettez:
<i>20.01.2017</i>`,

		idID: `Untuk mengatur tanggal jatuh tempo, silakan kirim dalam format teks <i>DD.MM.YEAR</i>.
<b>Misalnya</b> untuk 20 Januari 2017 kirim:
<i>20.01.2017</i>`,

		itIT: `Per impostare la data di scadenza invia il messaggio con la data nel seguente formato <i>GG.MM.ANNO</i>.
<b>Esempio</b> per indicare la data 20 Gennaio 2017 inserisci:
<i>20.01.2017</i>`,

		jaJP: `期日を設定するには、<i>DD.MM.YEAR</i>の形式でテキストとして送信してください。
<b>例えば</b>2017年1月20日の場合は次のように送信します:
<i>20.01.2017</i>`,

		koKR: `만기일을 설정하려면 <i>DD.MM.YEAR</i> 형식의 텍스트로 보내주세요.
<b>예를 들어</b> 2017년 1월 20일의 경우 다음과 같이 제출하세요:
<i>20.01.2017</i>`,

		plPL: `Aby ustawić termin płatności, wyślij go jako tekst w formacie <i>DD.MM.YEAR</i>.
<b>Na przykład</b> dla 20 stycznia 2017 r. wyślij:
<i>20.01.2017</i>`,

		ptBR: `Para definir a data de vencimento, envie-a como texto no formato <i>DD.MM.YEAR</i>.
<b>Por exemplo</b> para 20 de janeiro de 2017, envie:
<i>20.01.2017</i>`,

		ruRU: `Чтобы задать дату возврата напишите её в формате <i>ДД.MM.ГОД</i>.
<b>Например</b> для 20 января 2017 г.отправьте:
<i>20.01.2017</i>`,

		trTR: `Vade tarihini ayarlamak için <i>GG.AA.YIL</i> formatında metin olarak gönderin.
<b>Örneğin</b> 20 Ocak 2017 için şunu gönderin:
<i>20.01.2017</i>`,

		ukUA: `Щоб встановити дату повернення, надішліть її у форматі <i>ДД.ММ.РІК</i>.
<b>Наприклад</b> для 20 січня 2017 року надішліть:
<i>20.01.2017</i>`,

		uzUZ: `To'lov muddatini belgilash uchun uni <i>KK.OO.YIL</i> formatida matn sifatida yuboring.
<b>Masalan</b> 2017 yil 20 yanvar uchun quyidagini yuboring:
<i>20.01.2017</i>`,

		zhCN: `要设置到期日期，请以<i>DD.MM.YEAR</i>格式发送文本。
<b>例如</b>对于2017年1月20日，提交:
<i>20.01.2017</i>`,
	},
	MESSAGE_TEXT_WRONG_DATE: {
		deDE: "Entschuldigung, aber mit diesem Datum stimmt etwas nicht.",
		enUK: "Sorry, there is something wrong with the date you've provided.",
		enUS: "Sorry, there is something wrong with the date you've provided.",
		esES: "Lo siento, algo no es correcto con la fecha que has puesto",
		faIR: "متاسفم، در تاریخی که وارد نمودید مشکلی وجود دارد.",
		frFR: "Désolé, il y a un problème avec la date que vous avez fournie.",
		idID: "Maaf, ada yang salah dengan tanggal yang Anda berikan.",
		itIT: "Mi spiace, ma c'e' qualcosa di sbagliato nella data che hai inserito.",
		jaJP: "申し訳ありませんが、提供された日付に問題があります。",
		koKR: "죄송합니다. 제공하신 날짜에 문제가 있습니다.",
		plPL: "Przepraszamy, ale z podaną datą jest coś nie tak.",
		ptBR: "Desculpe, há algo errado com a data que você forneceu.",
		ruRU: "Извините, что-то не так с датой которую вы отправили.",
		trTR: "Üzgünüz, verdiğiniz tarihte bir sorun var.",
		ukUA: "Вибачте, але з датою щось не так.",
		uzUZ: "Kechirasiz, siz taqdim etgan sana bilan bog'liq muammo bor.",
		zhCN: "抱歉，您提供的日期有问题。",
	},
	COMMAND_TEXT_DISABLE_REMINDER: {
		deDE: "Nicht erinnern",
		enUK: "No reminder",
		enUS: "No reminder",
		esES: "No recordar",
		faIR: "بدون یادآور",
		frFR: "Pas de rappel",
		idID: "Tidak ada pengingat",
		itIT: "Nessun promemoria",
		jaJP: "リマインダーなし",
		koKR: "알림 없음",
		plPL: "Bez przypomnienia",
		ptBR: "Sem lembrete",
		ruRU: "Не напоминать",
		trTR: "Hatırlatma yok",
		ukUA: "Не нагадувати",
		uzUZ: "Eslatma yo'q",
		zhCN: "无提醒",
	},
	COMMAND_TEXT_TOMORROW: {
		deDE: "Morgen",
		enUK: "Tomorrow",
		enUS: "Tomorrow",
		esES: "Mañana",
		faIR: "فردا",
		frFR: "Demain",
		idID: "Besok",
		itIT: "Domani",
		jaJP: "明日",
		koKR: "내일",
		plPL: "Jutro",
		ptBR: "Amanhã",
		ruRU: "Завтра",
		trTR: "Yarın",
		ukUA: "Завтра",
		uzUZ: "Ertaga",
		zhCN: "明天",
	},
	COMMAND_TEXT_DAY_AFTER_TOMORROW: {
		deDE: "Übermorgen",
		enUK: "Day after tomorrow",
		enUS: "Day after tomorrow",
		esES: "Pasada mañana",
		faIR: "پس فردا",
		frFR: "Après-demain",
		idID: "Lusa",
		itIT: "Dopo domani",
		jaJP: "明後日",
		koKR: "모레",
		plPL: "Pojutrze",
		ptBR: "Depois de amanhã",
		ruRU: "Послезавтра",
		trTR: "Öbür gün",
		ukUA: "Післязавтра",
		uzUZ: "Indinga",
		zhCN: "后天",
	},
	COMMAND_TEXT_THIS_WEEK: {
		deDE: "Diese Woche",
		enUK: "This week",
		enUS: "This week",
		esES: "Esta semana",
		faIR: "این هفته",
		frFR: "Cette semaine",
		idID: "Minggu ini",
		itIT: "Questa settimana",
		jaJP: "今週",
		koKR: "이번 주",
		plPL: "W tym tygodniu",
		ptBR: "Esta semana",
		ruRU: "На этой неделе",
		trTR: "Bu hafta",
		ukUA: "Цього тижня",
		uzUZ: "Shu hafta",
		zhCN: "本周",
	},
	COMMAND_TEXT_YES_IT_HAS_RETURN_DEADLINE: {
		deDE: "Ja, es hat eine Frist!",
		enUK: "Yes, it has a deadline!",
		enUS: "Yes, it has a deadline!",
		esES: "Sí, hay una fecha de devolución!",
		faIR: "بله، دارای آخرین فرصت می باشد!",
		frFR: "Yes, it has a deadline!",
		idID: "Yes, it has a deadline!",
		itIT: "Si, c'e' una data di scadenza",
		jaJP: "Yes, it has a deadline!",
		koKR: "Yes, it has a deadline!",
		plPL: "Yes, it has a deadline!",
		ptBR: "Yes, it has a deadline!",
		ruRU: "Да, есть срок возврата!",
		trTR: "Yes, it has a deadline!",
		ukUA: "Yes, it has a deadline!",
		uzUZ: "Yes, it has a deadline!",
		zhCN: "Yes, it has a deadline!",
	},
	COMMAND_TEXT_NO_IT_CAN_BE_RETURNED_ANYTIME: {
		deDE: "Nein, sobald möglich.",
		enUK: "No, whenever is fine.",
		enUS: "No, whenever is fine.",
		esES: "No, sin fecha límite.",
		faIR: "خیر، هر زمانی مناسب است.",
		frFR: "No, whenever is fine.",
		idID: "No, whenever is fine.",
		itIT: "No, nessuna scadenza",
		jaJP: "No, whenever is fine.",
		koKR: "No, whenever is fine.",
		plPL: "No, whenever is fine.",
		ptBR: "No, whenever is fine.",
		ruRU: "Нет, срок возврата не важен.",
		trTR: "No, whenever is fine.",
		ukUA: "No, whenever is fine.",
		uzUZ: "No, whenever is fine.",
		zhCN: "No, whenever is fine.",
	},
	COMMAND_TEXT_IT_CAN_BE_RETURNED_ANYTIME: {
		deDE: "Unbefristet",
		enUK: "Whenever is fine",
		enUS: "Whenever is fine",
		esES: "Cualquier día",
		faIR: "هر زمانی مناسب است.",
		frFR: "Whenever is fine",
		idID: "Whenever is fine",
		itIT: "No, Nessuna scadenza",
		jaJP: "Whenever is fine",
		koKR: "Whenever is fine",
		plPL: "Whenever is fine",
		ptBR: "Whenever is fine",
		ruRU: "Когда-нибудь",
		trTR: "Whenever is fine",
		ukUA: "Whenever is fine",
		uzUZ: "Whenever is fine",
		zhCN: "Whenever is fine",
	},
	COMMAND_TEXT_IN_FEW_MINUTES: {
		deDE: "In wenigen Minuten",
		enUK: "In few minutes",
		enUS: "In few minutes",
		esES: "En unos minutos",
		faIR: "در چند دقیقه",
		frFR: "In few minutes",
		idID: "In few minutes",
		itIT: "Fra alcuni minuti",
		jaJP: "In few minutes",
		koKR: "In few minutes",
		plPL: "In few minutes",
		ptBR: "In few minutes",
		ruRU: "Через минуту",
		trTR: "In few minutes",
		ukUA: "In few minutes",
		uzUZ: "In few minutes",
		zhCN: "In few minutes",
	},
	COMMAND_TEXT_IN_1_WEEK: {
		deDE: "In einer Woche",
		enUK: "In 1 week",
		enUS: "In 1 week",
		esES: "En una semana",
		faIR: "ظرف یک هفته",
		frFR: "In 1 week",
		idID: "In 1 week",
		itIT: "Fra una settimana",
		jaJP: "In 1 week",
		koKR: "In 1 week",
		plPL: "In 1 week",
		ptBR: "In 1 week",
		ruRU: "Через неделю",
		trTR: "In 1 week",
		ukUA: "In 1 week",
		uzUZ: "In 1 week",
		zhCN: "In 1 week",
	},
	COMMAND_TEXT_IN_1_MONTH: {
		deDE: "In einem Monat",
		enUK: "In 1 month",
		enUS: "In 1 month",
		esES: "En un mes",
		faIR: "ظرف یک ماه",
		frFR: "In 1 month",
		idID: "In 1 month",
		itIT: "Fra un mese",
		jaJP: "In 1 month",
		koKR: "In 1 month",
		plPL: "In 1 month",
		ptBR: "In 1 month",
		ruRU: "Через месяц",
		trTR: "In 1 month",
		ukUA: "In 1 month",
		uzUZ: "In 1 month",
		zhCN: "In 1 month",
	},
	COMMAND_TEXT_SET_DATE: {
		deDE: "Datum setzen",
		enUK: "Set date",
		enUS: "Set date",
		esES: "Establecer la fecha",
		faIR: "تاریخ را تنظیم کنید",
		frFR: "Set date",
		idID: "Set date",
		itIT: "Imposta la data",
		jaJP: "Set date",
		koKR: "Set date",
		plPL: "Set date",
		ptBR: "Set date",
		ruRU: "Задать дату",
		trTR: "Set date",
		ukUA: "Set date",
		uzUZ: "Set date",
		zhCN: "Set date",
	},
	COMMAND_TEXT_MONDAY: {
		deDE: "Montag",
		enUK: "Monday",
		enUS: "Monday",
		esES: "Lunes",
		faIR: "دوشنبه",
		frFR: "Monday",
		idID: "Monday",
		itIT: "Lunedi'",
		jaJP: "Monday",
		koKR: "Monday",
		plPL: "Monday",
		ptBR: "Monday",
		ruRU: "Понедельник",
		trTR: "Monday",
		ukUA: "Monday",
		uzUZ: "Monday",
		zhCN: "Monday",
	},
	COMMAND_TEXT_TUESDAY: {
		deDE: "Dienstag",
		enUK: "Tuesday",
		enUS: "Tuesday",
		esES: "Martes",
		faIR: "سه شنبه",
		frFR: "Tuesday",
		idID: "Tuesday",
		itIT: "Martedi'",
		jaJP: "Tuesday",
		koKR: "Tuesday",
		plPL: "Tuesday",
		ptBR: "Tuesday",
		ruRU: "Вторник",
		trTR: "Tuesday",
		ukUA: "Tuesday",
		uzUZ: "Tuesday",
		zhCN: "Tuesday",
	},
	COMMAND_TEXT_WEDNESDAY: {
		deDE: "Mittwoch",
		enUK: "Wednesday",
		enUS: "Wednesday",
		esES: "Miercoles",
		faIR: "چهارشنبه",
		frFR: "Wednesday",
		idID: "Wednesday",
		itIT: "Mercoledi'",
		jaJP: "Wednesday",
		koKR: "Wednesday",
		plPL: "Wednesday",
		ptBR: "Wednesday",
		ruRU: "Среда",
		trTR: "Wednesday",
		ukUA: "Wednesday",
		uzUZ: "Wednesday",
		zhCN: "Wednesday",
	},
	COMMAND_TEXT_THURSDAY: {
		deDE: "Donnerstag",
		enUK: "Thursday",
		enUS: "Thursday",
		esES: "Jueves",
		faIR: "پنج شنبه",
		frFR: "Thursday",
		idID: "Thursday",
		itIT: "Giovedi'",
		jaJP: "Thursday",
		koKR: "Thursday",
		plPL: "Thursday",
		ptBR: "Thursday",
		ruRU: "Четверг",
		trTR: "Thursday",
		ukUA: "Thursday",
		uzUZ: "Thursday",
		zhCN: "Thursday",
	},
	COMMAND_TEXT_FRIDAY: {
		deDE: "Freitag",
		enUK: "Friday",
		enUS: "Friday",
		esES: "Viernes",
		faIR: "جمعه",
		frFR: "Friday",
		idID: "Friday",
		itIT: "Venerdi'",
		jaJP: "Friday",
		koKR: "Friday",
		plPL: "Friday",
		ptBR: "Friday",
		ruRU: "Пятница",
		trTR: "Friday",
		ukUA: "Friday",
		uzUZ: "Friday",
		zhCN: "Friday",
	},
	COMMAND_TEXT_SATURDAY: {
		deDE: "Samstag",
		enUK: "Saturday",
		enUS: "Saturday",
		esES: "Sabado",
		faIR: "شنبه",
		frFR: "Saturday",
		idID: "Saturday",
		itIT: "Sabato",
		jaJP: "Saturday",
		koKR: "Saturday",
		plPL: "Saturday",
		ptBR: "Saturday",
		ruRU: "Суббота",
		trTR: "Saturday",
		ukUA: "Saturday",
		uzUZ: "Saturday",
		zhCN: "Saturday",
	},
	COMMAND_TEXT_SUNDAY: {
		deDE: "Sonntag",
		enUK: "Sunday",
		enUS: "Sunday",
		esES: "Domingo",
		faIR: "یکشنبه",
		frFR: "Sunday",
		idID: "Sunday",
		itIT: "Domenica",
		jaJP: "Sunday",
		koKR: "Sunday",
		plPL: "Sunday",
		ptBR: "Sunday",
		ruRU: "Воскресенье",
		trTR: "Sunday",
		ukUA: "Sunday",
		uzUZ: "Sunday",
		zhCN: "Sunday",
	},
	COMMAND_TEXT_DO_NOT_SEND_RECEIPT: {
		deDE: "Keine Quittung schicken",
		enUK: "Do not send the receipt",
		enUS: "Do not send the receipt",
		esES: "No enviar el recibo",
		faIR: "رسید را ارسال نکنید",
		frFR: "Do not send the receipt",
		idID: "Do not send the receipt",
		itIT: "Non inviare la ricevuta",
		jaJP: "Do not send the receipt",
		koKR: "Do not send the receipt",
		plPL: "Do not send the receipt",
		ptBR: "Do not send the receipt",
		ruRU: "Не отправлять квитанцию",
		trTR: "Do not send the receipt",
		ukUA: "Do not send the receipt",
		uzUZ: "Do not send the receipt",
		zhCN: "Do not send the receipt",
	},
	MESSAGE_TEXT_RECEIPT_WILL_NOT_BE_SENT: {
		deDE: "Du hast dich gegen eine Quittung entschieden.",
		enUK: "You've decided not to send the receipt.",
		enUS: "You've decided not to send the receipt.",
		esES: "Has decidido no enviar el recibo",
		faIR: "شما تصمیم گرفتید که رسید را ارسال نکنید.",
		frFR: "You've decided not to send the receipt.",
		idID: "You've decided not to send the receipt.",
		itIT: "Hai scelto di non inviare la ricevuta",
		jaJP: "You've decided not to send the receipt.",
		koKR: "You've decided not to send the receipt.",
		plPL: "You've decided not to send the receipt.",
		ptBR: "You've decided not to send the receipt.",
		ruRU: "Вы решили не отправлять квитанцию.",
		trTR: "You've decided not to send the receipt.",
		ukUA: "You've decided not to send the receipt.",
		uzUZ: "You've decided not to send the receipt.",
		zhCN: "You've decided not to send the receipt.",
	},
	COMMAND_TEXT_I_HAVE_CHANGED_MY_MIND: {
		deDE: "Ich habe meine Meinung geändert",
		enUK: "I've changed my mind",
		enUS: "I've changed my mind",
		esES: "He cambiado de opinion",
		faIR: "نظرم را عوض کردم.",
		frFR: "I've changed my mind",
		idID: "I've changed my mind",
		itIT: "Ho cambiato idea",
		jaJP: "I've changed my mind",
		koKR: "I've changed my mind",
		plPL: "I've changed my mind",
		ptBR: "I've changed my mind",
		ruRU: "Я передумал(а)",
		trTR: "I've changed my mind",
		ukUA: "I've changed my mind",
		uzUZ: "I've changed my mind",
		zhCN: "I've changed my mind",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TELEGRAM: {
		deDE: "Senden per Telegram",
		enUK: "Send by Telegram",
		enUS: "Send by Telegram",
		esES: "Enviar a través de Telegram",
		faIR: "با تلگرام ارسال شود",
		frFR: "Send by Telegram",
		idID: "Send by Telegram",
		itIT: "Invia tramite Telegram",
		jaJP: "Send by Telegram",
		koKR: "Send by Telegram",
		plPL: "Send by Telegram",
		ptBR: "Send by Telegram",
		ruRU: "Отправить через Telelgram",
		trTR: "Send by Telegram",
		ukUA: "Send by Telegram",
		uzUZ: "Send by Telegram",
		zhCN: "Send by Telegram",
	},
	COMMAND_TEXT_GET_LINK_FOR_RECEIPT_IN_TELEGRAM: {
		deDE: "Erhalten sie einen link für eine quittung in Telegram",
		enUK: "Get link for a receipt in Telegram",
		enUS: "Get link for a receipt in Telegram",
		esES: "Obtener enlace para recibirlo en Telegram",
		faIR: "دریافت پیوند برای دریافت در Telegram",
		frFR: "Get link for a receipt in Telegram",
		idID: "Get link for a receipt in Telegram",
		itIT: "Link per la ricevuta nel Telegram",
		jaJP: "Get link for a receipt in Telegram",
		koKR: "Get link for a receipt in Telegram",
		plPL: "Get link for a receipt in Telegram",
		ptBR: "Get link for a receipt in Telegram",
		ruRU: "Ссылка для квитанции в Телеграмме",
		trTR: "Get link for a receipt in Telegram",
		ukUA: "Get link for a receipt in Telegram",
		uzUZ: "Get link for a receipt in Telegram",
		zhCN: "Get link for a receipt in Telegram",
	},
	COMMAND_TEXT_COUNTERPARTY_HAS_NO_TELEGRAM: {
		deDE: "Senden per FB, WhatsApp, Viber, etc.",
		enUK: "Send by FB, WhatsApp, Viber, etc.",
		enUS: "Send by FB, WhatsApp, Viber, etc.",
		esES: "Enviar a través de FB, WhatsApp, Viber, etc.",
		faIR: "با فیسبوک، واتس آپ، وایبر و ... ارسال شود.",
		frFR: "Send by FB, WhatsApp, Viber, etc.",
		idID: "Send by FB, WhatsApp, Viber, etc.",
		itIT: "Invia con FB, WhatsCrap, Viber, etc.",
		jaJP: "Send by FB, WhatsApp, Viber, etc.",
		koKR: "Send by FB, WhatsApp, Viber, etc.",
		plPL: "Send by FB, WhatsApp, Viber, etc.",
		ptBR: "Send by FB, WhatsApp, Viber, etc.",
		ruRU: "Отправить через Viber, VK, FB, ...",
		trTR: "Send by FB, WhatsApp, Viber, etc.",
		ukUA: "Send by FB, WhatsApp, Viber, etc.",
		uzUZ: "Send by FB, WhatsApp, Viber, etc.",
		zhCN: "Send by FB, WhatsApp, Viber, etc.",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_SMS: {
		deDE: "Senden per SMS",
		enUK: "Send by SMS",
		enUS: "Send by SMS",
		esES: "Enviar a través de SMS",
		faIR: "با پیام کوتاه ارسال شود",
		frFR: "Send by SMS",
		idID: "Send by SMS",
		itIT: "Invia tramite SMS",
		jaJP: "Send by SMS",
		koKR: "Send by SMS",
		plPL: "Send by SMS",
		ptBR: "Send by SMS",
		ruRU: "Отправить через SMS",
		trTR: "Send by SMS",
		ukUA: "Send by SMS",
		uzUZ: "Send by SMS",
		zhCN: "Send by SMS",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_VK: {
		deDE: "Senden per VK.com",
		enUK: "Send throw VK.com",
		enUS: "Send throw VK.com",
		esES: "Enviar vía VK.com",
		faIR: "ارسال شود VK.com از طریق ",
		frFR: "Send throw VK.com",
		idID: "Send throw VK.com",
		itIT: "Invia tramite VK.com (Facebook russo)",
		jaJP: "Send throw VK.com",
		koKR: "Send throw VK.com",
		plPL: "Send throw VK.com",
		ptBR: "Send throw VK.com",
		ruRU: "Отправить через ВКонтакте",
		trTR: "Send throw VK.com",
		ukUA: "Send throw VK.com",
		uzUZ: "Send throw VK.com",
		zhCN: "Send throw VK.com",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_OK: {
		deDE: "Senden per OK",
		enUK: "Send throw OK",
		enUS: "Send throw OK",
		esES: "Enviar a través de OK",
		faIR: "ارسال شود OK از طریق ",
		frFR: "Send throw OK",
		idID: "Send throw OK",
		itIT: "Invia tramite OK",
		jaJP: "Send throw OK",
		koKR: "Send throw OK",
		plPL: "Send throw OK",
		ptBR: "Send throw OK",
		ruRU: "Отправить через Одноклассники",
		trTR: "Send throw OK",
		ukUA: "Send throw OK",
		uzUZ: "Send throw OK",
		zhCN: "Send throw OK",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_FB: {
		deDE: "Senden per Facebook",
		enUK: "Send throw Facebook",
		enUS: "Send throw Facebook",
		esES: "Enviar a través de Facebook",
		faIR: "از طریق فیسبوک ارسال شود.",
		frFR: "Send throw Facebook",
		idID: "Send throw Facebook",
		itIT: "Invia tramite Facebook",
		jaJP: "Send throw Facebook",
		koKR: "Send throw Facebook",
		plPL: "Send throw Facebook",
		ptBR: "Send throw Facebook",
		ruRU: "Отправить через Facebook",
		trTR: "Send throw Facebook",
		ukUA: "Send throw Facebook",
		uzUZ: "Send throw Facebook",
		zhCN: "Send throw Facebook",
	},
	COMMAND_TEXT_SEND_RECEIPT_BY_TWT: {
		deDE: "Senden per Twitter",
		enUK: "Send throw Twitter",
		enUS: "Send throw Twitter",
		esES: "Enviar a través de Twitter",
		faIR: "از طریق توئیتر ارسال شود.",
		frFR: "Send throw Twitter",
		idID: "Send throw Twitter",
		itIT: "Invia tramite Twitter",
		jaJP: "Send throw Twitter",
		koKR: "Send throw Twitter",
		plPL: "Send throw Twitter",
		ptBR: "Send throw Twitter",
		ruRU: "Отправить через Twitter",
		trTR: "Send throw Twitter",
		ukUA: "Send throw Twitter",
		uzUZ: "Send throw Twitter",
		zhCN: "Send throw Twitter",
	},
	COMMAND_TEXT_CANCEL_SENDING_RECEIPT_BY_TELEGRAM: {
		deDE: "Sendung der Quittung per Telegram abbrechen",
		enUK: "Cancel sending receipt by Telegram",
		enUS: "Cancel sending receipt by Telegram",
		esES: "Cancelar el envío a través de Telegram",
		faIR: "ارسال رسید با تلگرام کنسل شود",
		frFR: "Cancel sending receipt by Telegram",
		idID: "Cancel sending receipt by Telegram",
		itIT: "Annulla l'invio tramite Telegram",
		jaJP: "Cancel sending receipt by Telegram",
		koKR: "Cancel sending receipt by Telegram",
		plPL: "Cancel sending receipt by Telegram",
		ptBR: "Cancel sending receipt by Telegram",
		ruRU: "Отменить отправку через Telegram",
		trTR: "Cancel sending receipt by Telegram",
		ukUA: "Cancel sending receipt by Telegram",
		uzUZ: "Cancel sending receipt by Telegram",
		zhCN: "Cancel sending receipt by Telegram",
	},
	MAIN_MENU: {
		deDE: "Hauptmenü",
		enUK: "Main menu",
		enUS: "Main menu",
		esES: "Menú principal",
		faIR: "منوی اصلی",
		frFR: "Menu principal",
		idID: "Menu utama",
		itIT: "Menu principale",
		jaJP: "メインメニュー",
		koKR: "메인 메뉴",
		plPL: "Menu główne",
		ptBR: "Menu principal",
		ruRU: "Главное меню",
		trTR: "Ana menü",
		ukUA: "Головне меню",
		uzUZ: "Asosiy menyu",
		zhCN: "主菜单",
	},
	COMMAND_TEXT_MAIN_MENU_TITLE: {
		deDE: "Menü /menu",
		enUK: "Main /menu",
		enUS: "Main /menu",
		esES: "Inicio /menú",
		faIR: "/منو ی اصلی ",
		frFR: "Main /menu",
		idID: "Main /menu",
		itIT: "Menu' /menu",
		jaJP: "Main /menu",
		koKR: "Main /menu",
		plPL: "Main /menu",
		ptBR: "Main /menu",
		ruRU: "Главное /меню",
		trTR: "Main /menu",
		ukUA: "Main /menu",
		uzUZ: "Main /menu",
		zhCN: "Main /menu",
	},
	MESSAGE_TEXT_NOTHING_TO_CANCEL: {
		deDE: "Nichts zum abbrechen.",
		enUK: "Nothing to cancel.",
		enUS: "Nothing to cancel.",
		esES: "No hay nada que anular.",
		faIR: "چیزی برای کنسل شدن وجود ندارد",
		frFR: "Nothing to cancel.",
		idID: "Nothing to cancel.",
		itIT: "Nulla da annullare.",
		jaJP: "Nothing to cancel.",
		koKR: "Nothing to cancel.",
		plPL: "Nothing to cancel.",
		ptBR: "Nothing to cancel.",
		ruRU: "Нечего отменять.",
		trTR: "Nothing to cancel.",
		ukUA: "Nothing to cancel.",
		uzUZ: "Nothing to cancel.",
		zhCN: "Nothing to cancel.",
	},
	MESSAGE_TEXT_TRANSFER_CREATION_CANCELED: {
		deDE: "Erstellung des Schuldscheins abgebrochen.",
		enUK: "Creation of debt record has been canceled.",
		enUS: "Creation of debt record has been canceled.",
		esES: "La creación del recordatorio se ha cancelado.",
		faIR: "ایجاد سابقه بدهی کنسل شد.",
		frFR: "Creation of debt record has been canceled.",
		idID: "Creation of debt record has been canceled.",
		itIT: "Creazione record annullata",
		jaJP: "Creation of debt record has been canceled.",
		koKR: "Creation of debt record has been canceled.",
		plPL: "Creation of debt record has been canceled.",
		ptBR: "Creation of debt record has been canceled.",
		ruRU: "Создание записи о долге отменено.",
		trTR: "Creation of debt record has been canceled.",
		ukUA: "Creation of debt record has been canceled.",
		uzUZ: "Creation of debt record has been canceled.",
		zhCN: "Creation of debt record has been canceled.",
	},
	COMMAND_TEXT_SHOW_ALL_CONTACTS: {
		deDE: "Zeige alle...",
		enUK: "Show all...",
		enUS: "Show all...",
		esES: "Mostrar todo...",
		faIR: "نمایش تمام موارد ...",
		frFR: "Show all...",
		idID: "Show all...",
		itIT: "Mostra tutto...",
		jaJP: "Show all...",
		koKR: "Show all...",
		plPL: "Show all...",
		ptBR: "Show all...",
		ruRU: "Показать все...",
		trTR: "Show all...",
		ukUA: "Show all...",
		uzUZ: "Show all...",
		zhCN: "Show all...",
	},
	COMMAND_TEXT_CONTACTS: {
		deDE: "Kontakte",
		enUK: "Contacts",
		enUS: "Contacts",
		esES: "Contactos",
		faIR: "لیست تماس",
		frFR: "Contacts",
		idID: "Contacts",
		itIT: "Сontatti",
		jaJP: "Contacts",
		koKR: "Contacts",
		plPL: "Contacts",
		ptBR: "Contacts",
		ruRU: "Контакты",
		trTR: "Contacts",
		ukUA: "Contacts",
		uzUZ: "Contacts",
		zhCN: "Contacts",
	},
	COMMAND_TEXT_REFRESH: {
		deDE: "Aktualisieren",
		enUK: "Refresh",
		enUS: "Refresh",
		esES: "Recargar",
		faIR: "تازه کردن",
		frFR: "Refresh",
		idID: "Refresh",
		itIT: "Ricaricare",
		jaJP: "Refresh",
		koKR: "Refresh",
		plPL: "Refresh",
		ptBR: "Refresh",
		ruRU: "Обновить",
		trTR: "Refresh",
		ukUA: "Refresh",
		uzUZ: "Refresh",
		zhCN: "Refresh",
	},
	COMMAND_TEXT_ADD_YOUR_OWN_OPTION: {
		deDE: "Etwas anderes",
		enUK: "Something else",
		enUS: "Something else",
		esES: "Otra cosa",
		faIR: "چیزی دیگر",
		frFR: "Something else",
		idID: "Something else",
		itIT: "Qualcos'altro",
		jaJP: "Something else",
		koKR: "Something else",
		plPL: "Something else",
		ptBR: "Something else",
		ruRU: "Что-то другое",
		trTR: "Something else",
		ukUA: "Something else",
		uzUZ: "Something else",
		zhCN: "Something else",
	},
	MESSAGE_TEXT_REMINDER_ASK_IF_RETURNED: {
		deDE: "Wurde diese Schuld beglichen?",
		enUK: "Have this debt been returned?",
		enUS: "Have this debt been returned?",
		esES: "¿Se ha devuelto esta deuda?",
		faIR: "آیا این بدهی بازپرداخت شده است؟",
		frFR: "Have this debt been returned?",
		idID: "Have this debt been returned?",
		itIT: "Questo debito e' stato saldato?",
		jaJP: "Have this debt been returned?",
		koKR: "Have this debt been returned?",
		plPL: "Have this debt been returned?",
		ptBR: "Have this debt been returned?",
		ruRU: "Был ли этот долг возвращён?",
		trTR: "Have this debt been returned?",
		ukUA: "Have this debt been returned?",
		uzUZ: "Have this debt been returned?",
		zhCN: "Have this debt been returned?",
	},
	MESSAGE_TEXT_ASK_WHEN_TO_REMIND_AGAIN: {
		deDE: "Wann willst du wieder an diese Schuld erinnert werden?",
		enUK: "When should we remind you about this debt again?",
		enUS: "When should we remind you about this debt again?",
		esES: "¿Cuándo recordarte de esta deuda otra vez?",
		faIR: "چه زمانی لازم است مجدداً در مورد این بدهی به شما یادآوری نماییم؟",
		frFR: "When should we remind you about this debt again?",
		idID: "When should we remind you about this debt again?",
		itIT: "Quando devo ricordarti di questo debito?",
		jaJP: "When should we remind you about this debt again?",
		koKR: "When should we remind you about this debt again?",
		plPL: "When should we remind you about this debt again?",
		ptBR: "When should we remind you about this debt again?",
		ruRU: "Когда вам напомнить об этом долге ещё раз?",
		trTR: "When should we remind you about this debt again?",
		ukUA: "When should we remind you about this debt again?",
		uzUZ: "When should we remind you about this debt again?",
		zhCN: "When should we remind you about this debt again?",
	},
	MESSAGE_TEXT_REPLIED_DEBT_RETURNED_FULLY: {
		deDE: "Du hast angegeben, dass diese Schuld vollständig beglichen ist.",
		enUK: "You've replied back that debt has been returned fully.",
		enUS: "You've replied back that debt has been returned fully.",
		esES: "Has confirmado que la deuda se ha saldado totalmente",
		faIR: "شما پاسخ داده اید که بدهی به صورت کامل بازپرداخت شده است.",
		frFR: "You've replied back that debt has been returned fully.",
		idID: "You've replied back that debt has been returned fully.",
		itIT: "Hai confermato che il debito e' stato saldato.",
		jaJP: "You've replied back that debt has been returned fully.",
		koKR: "You've replied back that debt has been returned fully.",
		plPL: "You've replied back that debt has been returned fully.",
		ptBR: "You've replied back that debt has been returned fully.",
		ruRU: "Вы ответили что долг возвращён полностью.",
		trTR: "You've replied back that debt has been returned fully.",
		ukUA: "You've replied back that debt has been returned fully.",
		uzUZ: "You've replied back that debt has been returned fully.",
		zhCN: "You've replied back that debt has been returned fully.",
	},
	MESSAGE_TEXT_DEBT_IS_RETURNED: {
		deDE: "Die Schuld ist vollständig beglichen.",
		enUK: "The debt has been returned fully.",
		enUS: "The debt has been returned fully.",
		esES: "La deuda se ha saldado totalmente",
		faIR: "بدهی به صورت کامل بازپرداخت شده است",
		frFR: "The debt has been returned fully.",
		idID: "The debt has been returned fully.",
		itIT: "Il debito e' stato saldato.",
		jaJP: "The debt has been returned fully.",
		koKR: "The debt has been returned fully.",
		plPL: "The debt has been returned fully.",
		ptBR: "The debt has been returned fully.",
		ruRU: "Долг возвращён полностью.",
		trTR: "The debt has been returned fully.",
		ukUA: "The debt has been returned fully.",
		uzUZ: "The debt has been returned fully.",
		zhCN: "The debt has been returned fully.",
	},
	MESSAGE_TEXT_DETAILS_ARE_HERE: {
		deDE: "Details hier: %v",
		enUK: "Details here: %v",
		enUS: "Details here: %v",
		esES: "Detalles aquí: %v",
		faIR: "جزئیات در اینجا: %v",
		frFR: "Details here: %v",
		idID: "Details here: %v",
		itIT: "Dettagli qui: %v",
		jaJP: "Details here: %v",
		koKR: "Details here: %v",
		plPL: "Details here: %v",
		ptBR: "Details here: %v",
		ruRU: "Подробности тут: %v",
		trTR: "Details here: %v",
		ukUA: "Details here: %v",
		uzUZ: "Details here: %v",
		zhCN: "Details here: %v",
	},
	MESSAGE_TEXT_REMINDER: {
		deDE: "Erinnerung",
		enUK: "Reminder",
		enUS: "Reminder",
		esES: "Recordatorio",
		faIR: "یادآور",
		frFR: "Reminder",
		idID: "Reminder",
		itIT: "Promemoria",
		jaJP: "Reminder",
		koKR: "Reminder",
		plPL: "Reminder",
		ptBR: "Reminder",
		ruRU: "Напоминание",
		trTR: "Reminder",
		ukUA: "Reminder",
		uzUZ: "Reminder",
		zhCN: "Reminder",
	},
	MESSAGE_TEXT_REMINDER_SET: {
		deDE: "Erinnerung am: %v",
		enUK: "Reminder set for: %v",
		enUS: "Reminder set for: %v",
		esES: "Recordatorio establecito para: %v",
		faIR: "یادآور تنظیم شده است برای: %v",
		frFR: "Reminder set for: %v",
		idID: "Reminder set for: %v",
		itIT: "Imposta promemoria per: %v",
		jaJP: "Reminder set for: %v",
		koKR: "Reminder set for: %v",
		plPL: "Reminder set for: %v",
		ptBR: "Reminder set for: %v",
		ruRU: "Напоминание установлено на: %v",
		trTR: "Reminder set for: %v",
		ukUA: "Reminder set for: %v",
		uzUZ: "Reminder set for: %v",
		zhCN: "Reminder set for: %v",
	},
	MESSAGE_TEXT_REMINDER_DISABLED: {
		deDE: "Du hast die Erinnerung an diese Schuld deaktiviert.",
		enUK: "You've disabled reminders for this debt.",
		enUS: "You've disabled reminders for this debt.",
		esES: "Recordatorio para esta deuda se ha deshabilitado.",
		faIR: "شما یادآور را برای این بدهی غیرفعال نموده اید.",
		frFR: "You've disabled reminders for this debt.",
		idID: "You've disabled reminders for this debt.",
		itIT: "Hai disabilitato il promemoria per questo debito.",
		jaJP: "You've disabled reminders for this debt.",
		koKR: "You've disabled reminders for this debt.",
		plPL: "You've disabled reminders for this debt.",
		ptBR: "You've disabled reminders for this debt.",
		ruRU: "Напоминания об этом долге отключены.",
		trTR: "You've disabled reminders for this debt.",
		ukUA: "You've disabled reminders for this debt.",
		uzUZ: "You've disabled reminders for this debt.",
		zhCN: "You've disabled reminders for this debt.",
	},
	COMMAND_TEXT_REMINDER_ENABLE: {
		deDE: "Erinnerung aktivieren",
		enUK: "Turn-on reminder",
		enUS: "Turn-on reminder",
		esES: "Recordatorio de encendido",
		faIR: "یادآوری روشن",
		frFR: "Turn-on reminder",
		idID: "Turn-on reminder",
		itIT: "Ricordo promozionale",
		jaJP: "Turn-on reminder",
		koKR: "Turn-on reminder",
		plPL: "Turn-on reminder",
		ptBR: "Turn-on reminder",
		ruRU: "Включить напоминание",
		trTR: "Turn-on reminder",
		ukUA: "Turn-on reminder",
		uzUZ: "Turn-on reminder",
		zhCN: "Turn-on reminder",
	},
	MESSAGE_TEXT_REMINDER_ALREADY_RESCHEDULED: {
		deDE: "Du wirst bereits erneut erinnert.",
		enUK: "You've already rescheduled this reminder.",
		enUS: "You've already rescheduled this reminder.",
		esES: "Recordatorio para esta deuda se ha reprogramado ya.",
		faIR: "شما قبلا به صورت مجدد این یادآور را زمانبندی نموده اید.",
		frFR: "You've already rescheduled this reminder.",
		idID: "You've already rescheduled this reminder.",
		itIT: "Hai gia' impostato questo promemoria",
		jaJP: "You've already rescheduled this reminder.",
		koKR: "You've already rescheduled this reminder.",
		plPL: "You've already rescheduled this reminder.",
		ptBR: "You've already rescheduled this reminder.",
		ruRU: "Напоминание об этом долге уже перенесено.",
		trTR: "You've already rescheduled this reminder.",
		ukUA: "You've already rescheduled this reminder.",
		uzUZ: "You've already rescheduled this reminder.",
		zhCN: "You've already rescheduled this reminder.",
	},
	COMMAND_TEXT_REMINDER_RETURNED_IN_FULL: {
		deDE: "Ja, vollständig beglichen",
		enUK: "Yes, returne in full",
		enUS: "Yes, returne in full",
		esES: "Sí, devuelto totalmente",
		faIR: "بله، بازپرداخت به صورت کامل",
		frFR: "Yes, returne in full",
		idID: "Yes, returne in full",
		itIT: "Si, completamento saldato",
		jaJP: "Yes, returne in full",
		koKR: "Yes, returne in full",
		plPL: "Yes, returne in full",
		ptBR: "Yes, returne in full",
		ruRU: "Да, возвращено полностью",
		trTR: "Yes, returne in full",
		ukUA: "Yes, returne in full",
		uzUZ: "Yes, returne in full",
		zhCN: "Yes, returne in full",
	},
	COMMAND_TEXT_REMINDER_RETURNED_PARTIALLY: {
		deDE: "Teilweise beglichen",
		enUK: "Returned partially",
		enUS: "Returned partially",
		esES: "Devuelto parcialmente",
		faIR: "تا اندازه ای بازپرداخت شده است",
		frFR: "Returned partially",
		idID: "Returned partially",
		itIT: "Parzialmente saldato",
		jaJP: "Returned partially",
		koKR: "Returned partially",
		plPL: "Returned partially",
		ptBR: "Returned partially",
		ruRU: "Возврашено частично",
		trTR: "Returned partially",
		ukUA: "Returned partially",
		uzUZ: "Returned partially",
		zhCN: "Returned partially",
	},
	COMMAND_TEXT_REMINDER_NOT_RETURNED: {
		deDE: "Nicht beglichen",
		enUK: "Not returned",
		enUS: "Not returned",
		esES: "No devuelto",
		faIR: "بازپرداخت نشده است",
		frFR: "Not returned",
		idID: "Not returned",
		itIT: "Debito non saldato",
		jaJP: "Not returned",
		koKR: "Not returned",
		plPL: "Not returned",
		ptBR: "Not returned",
		ruRU: "Не возвращено",
		trTR: "Not returned",
		ukUA: "Not returned",
		uzUZ: "Not returned",
		zhCN: "Not returned",
	},
	MESSAGE_TEXT_YOU_REPLIED: {
		deDE: "Beantwortet: %v",
		enUK: "You've replied: %v",
		esES: "Has respondido: %v",
		faIR: "شما پاسخ داده اید: %v",
		itIT: "Hai risposto: %v",
		ruRU: "Вы ответили: %v",
	},
	"book": {
		deDE: "buchen",
		enUK: "book",
		esES: "libro",
		faIR: "کتاب",
		itIT: "libro",
		ruRU: "книгу",
	},
	"MessageTextBotDidNotUnderstandTheCommand": {
		deDE: "\xF0\x9F\x98\xB3 Entschuldigung, aber ich habe deinen Befehl nicht verstanden. Vielleicht bin ich ein bisschen dumm...\n\nDu kannst zurück ins /menu",
		enUK: "\xF0\x9F\x98\xB3 Sorry, I did not understand your order. May be I'm a little bit dumb...\n\nYou can return to main /menu",
		esES: "\xF0\x9F\x98\xB3 Disculpa, no he entendido tu orden. Tal vez soy un poco tonto...\n\nPuedes volver al Menu principal /menu",
		faIR: "\xF0\x9F\x98\xB3 ببخشید، من دستور شما را نفهمیدم. احتمالا کمی کند ذهن هستم...\n\nشما میتوانید به /منو ی اصلی بازگردید",
		itIT: "\xF0\x9F\x98\xB3 Scusami ma non ho capito cosa vuoi. Sono ancora un po' sciocco...\n\nPuoi ritornare al Menu con /menu",
		ruRU: "\xF0\x9F\x98\xB3 Извините, я не понял вашу команду. Возможно я немного туповат...\n\nЧтобы начать сначала нажмите /menu",
	},
	COMMAND_TEXT_LANGUAGE: {
		deDE: "Sprache / Language",
		enUK: "Bot language",
		esES: "Idioma / Language",
		faIR: "زبان",
		itIT: "Lingua / Language",
		ruRU: "Язык / Language",
	},
	"/start": {
		deDE: "/start",
		enUK: "/start",
		esES: "/comienzo",
		faIR: "/شروع",
		itIT: "/start",
		ruRU: "/старт",
	},
	COMMAND_TEXT_DUE_RETURNS: {
		deDE: "Fällige Schulden",
		enUK: "Due returns",
		esES: "Devoluciones",
		faIR: "بازپرداخت بدهی",
		itIT: "Debiti",
		ruRU: "Предстоящие платежи",
	},
	MESSAGE_TEXT_OVERDUE_RETURNS_HEADER: {
		deDE: "<b>Überfällige Schulden:</b>",
		enUK: "<b>Overdue debts:</b>",
		esES: "<b>Deudas atrasadas:</b>",
		faIR: "<b>بدهی های معوق:</b>",
		itIT: "<b>Debiti in ritardo:</b>",
		ruRU: "<b>Просроченные долги:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_HEADER: {
		deDE: "<b>Bald fällige Schulden:</b>",
		enUK: "<b>Closest debts to return:</b>",
		esES: "<b>Deudas más cercanos que pagar:</b>",
		faIR: "<b>نزدیک ترین بدهی برای بازپرداخت:</b>",
		itIT: "<b>Debiti in scadenza:</b>",
		ruRU: "<b>Ближайшие долги к возрату:</b>",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_USER: {
		deDE: "%v bekommt %v von dir, spätestens in %v",
		enUK: "%v expects %v from you in %v",
		esES: "%v espera %v que devuelvas en %v",
		itIT: "%v aspetta %v da te entro il %v",
		faIR: "%v انتظار دارد %v از شما در %v",
		ruRU: "%v ожидает от вас возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_ROW_BY_COUNTERPARTY: {
		deDE: "%v gibt dir %v, spätestens in %v",
		enUK: "You expect %v to return %v to you in %v",
		esES: "Estas esperando de %v que devuelva %v a ti en %v",
		faIR: "شما انتظار دارید %v بازگرداند %v به شما در %v",
		itIT: "Stai aspettando %v che ti dia %v entro il %v",
		ruRU: "Вы ожидаете от %v возврата %v через %v",
	},
	MESSAGE_TEXT_DUE_RETURNS_EMPTY: {
		deDE: "Du hast keine Schulden mit Fälligkeitsdatum.",
		enUK: "You have no debts with set due date.",
		esES: "No tienes deudas con la fecha señalada para devolver. ",
		faIR: "شما بدهی ای با ثبت سررسید ندارید.",
		itIT: "Non hai debiti con una data di scadenza.",
		ruRU: "У вас нет долгов с назначеным сроком к возврату.",
	},
	COMMAND_TEXT_GAVE: {
		deDE: "Verleihen",
		enUK: "Gave",
		esES: "Prestado por ti",
		faIR: "قرض_دادن",
		itIT: "Credito",
		ruRU: "Дать",
	},
	COMMAND_TEXT_GOT: {
		deDE: "Anleihen",
		enUK: "Got",
		esES: "Prestado a ti",
		faIR: "قرض_گرفتن",
		itIT: "Debito",
		ruRU: "Взять",
	},
	COMMAND_TEXT_RETURN: {
		deDE: "Beglichen",
		enUK: "Return",
		esES: "Devuelto",
		faIR: "بازگشت",
		itIT: "Rientra",
		ruRU: "Вернуть",
	},
	COMMAND_TEXT_BALANCE: {
		deDE: "Ausstehend",
		enUK: "Balance",
		esES: "Balance",
		faIR: "تراز",
		itIT: "Bilancio",
		ruRU: "Баланс",
	},
	COMMAND_TEXT_SETTING: {
		deDE: "Einstellungen",
		enUK: "Settings",
		esES: "Ajustes",
		faIR: "تنظیمات",
		itIT: "Settaggi",
		ruRU: "Настройки",
	},
	COMMAND_TEXT_HIGH_FIVE: {
		deDE: "Gib mir Fünf!",
		enUK: "High five!",
		enUS: "High five!",
		esES: "¡Choca esos cinco!",
		faIR: "بزن قدش!",
		frFR: "Tape m'en cinq !",
		idID: "High five!",
		itIT: "Batti 5 bro!",
		jaJP: "ハイタッチ！",
		koKR: "하이 파이브!",
		plPL: "Przybij piątkę!",
		ptBR: "High five!",
		ruRU: "Дать пять!",
		trTR: "Çak bir beşlik!",
		ukUA: "Дай п'ять!",
		uzUZ: "Besh qo'l!",
		zhCN: "击掌！",
	},
	COMMAND_TEXT_CHANGE_LANG: {
		deDE: "Sprache",
		enUK: "Language",
		enUS: "Language",
		esES: "Idioma",
		faIR: "زبان",
		frFR: "Langue",
		idID: "Bahasa",
		itIT: "Lingua",
		jaJP: "言語",
		koKR: "언어",
		plPL: "Język",
		ptBR: "Idioma",
		ruRU: "Язык",
		trTR: "Dil",
		ukUA: "Мова",
		uzUZ: "Til",
		zhCN: "语言",
	},
	COMMAND_TEXT_HELP: {
		deDE: "Hilfe",
		enUK: "Help",
		enUS: "Help",
		esES: "Ayuda",
		faIR: "کمک",
		frFR: "Aide",
		idID: "Bantuan",
		itIT: "Aiuto",
		jaJP: "ヘルプ",
		koKR: "도움말",
		plPL: "Pomoc",
		ptBR: "Ajuda",
		ruRU: "Помощь",
		trTR: "Yardım",
		ukUA: "Допомога",
		uzUZ: "Yordam",
		zhCN: "帮助",
	},
	COMMAND_TEXT_HISTORY: {
		deDE: "Verlauf",
		enUK: "History",
		enUS: "History",
		esES: "Cronología",
		faIR: "پیشینه",
		frFR: "Historique",
		idID: "Riwayat",
		itIT: "Cronologia",
		jaJP: "履歴",
		koKR: "기록",
		plPL: "Historia",
		ptBR: "Histórico",
		ruRU: "История",
		trTR: "Geçmiş",
		ukUA: "Історія",
		uzUZ: "Tarix",
		zhCN: "历史",
	},
	COMMAND_TEXT_CANCEL: {
		deDE: "Abbrechen",
		enUK: "Cancel",
		enUS: "Cancel",
		esES: "Cancelar",
		faIR: "کنسل",
		frFR: "Annuler",
		idID: "Batal",
		itIT: "Annulla",
		jaJP: "キャンセル",
		koKR: "취소",
		plPL: "Anuluj",
		ptBR: "Cancelar",
		ruRU: "Отменить",
		trTR: "İptal",
		ukUA: "Скасувати",
		uzUZ: "Bekor qilish",
		zhCN: "取消",
	},
	COMMAND_TEXT_REFERRERS: {
		deDE: "Empfehlungen",
		enUK: "Referrers",
		enUS: "Referrers",
		esES: "Referentes",
		faIR: "معرف\u200cها",
		frFR: "Référents",
		idID: "Referensi",
		itIT: "Referenti",
		jaJP: "紹介者",
		koKR: "추천인",
		plPL: "Polecający",
		ptBR: "Referências",
		ruRU: "Нас рекомендуют",
		trTR: "Referanslar",
		ukUA: "Нас рекомендують",
		uzUZ: "Tavsiya qiluvchilar",
		zhCN: "推荐人",
	},
	MESSAGE_TEXT_HOW_TO_ADD_TG_CHANNEL: {
		deDE: `Um deinen Kanal zur Liste hinzuzufügen, schreibe einfach über uns mit einem Link wie %v <code>&lt;-</code> ersetze <code>YOUR_CHANNEL</code> durch deinen eigenen Kanal.

Es ist besser, wenn du den Link in HTML versteckst als:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Dies sollte von Telegram-Clients so dargestellt werden: <a href="%v">@%v</a>

Die Top 5 Empfehlungen der letzten 100 neuen Benutzer werden hier angezeigt.`,

		enUK: `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,

		enUS: `To add your channel to the list just write about us with a link as %v <code>&lt;-</code> replace <code>YOUR_CHANNEL</code> with your own channel.

It's better if you hide the link in HTML as:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

This should be rendered by Telegram clients as: <a href="%v">@%v</a>

Top 5 referrers for the last 100 new users will be shown here.`,

		esES: `Para añadir tu canal a la lista, simplemente escribe sobre nosotros con un enlace como %v <code>&lt;-</code> reemplaza <code>YOUR_CHANNEL</code> con tu propio canal.

Es mejor si ocultas el enlace en HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Esto debería ser renderizado por los clientes de Telegram como: <a href="%v">@%v</a>

Los 5 principales referentes de los últimos 100 nuevos usuarios se mostrarán aquí.`,

		faIR: `برای اضافه کردن کانال خود به لیست، فقط درباره ما با لینکی مانند %v <code>&lt;-</code> بنویسید و <code>YOUR_CHANNEL</code> را با کانال خود جایگزین کنید.

بهتر است اگر لینک را در HTML به این صورت پنهان کنید:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

این باید توسط کلاینت\u200cهای تلگرام به این صورت نمایش داده شود: <a href="%v">@%v</a>

5 معرف برتر برای 100 کاربر جدید آخر در اینجا نشان داده خواهد شد.`,

		frFR: `Pour ajouter votre chaîne à la liste, écrivez simplement à propos de nous avec un lien comme %v <code>&lt;-</code> remplacez <code>YOUR_CHANNEL</code> par votre propre chaîne.

C'est mieux si vous cachez le lien en HTML comme:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Cela devrait être rendu par les clients Telegram comme: <a href="%v">@%v</a>

Les 5 principaux référents pour les 100 derniers nouveaux utilisateurs seront affichés ici.`,

		idID: `Untuk menambahkan saluran Anda ke daftar, cukup tulis tentang kami dengan tautan sebagai %v <code>&lt;-</code> ganti <code>YOUR_CHANNEL</code> dengan saluran Anda sendiri.

Lebih baik jika Anda menyembunyikan tautan dalam HTML sebagai:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Ini harus dirender oleh klien Telegram sebagai: <a href="%v">@%v</a>

5 referensi teratas untuk 100 pengguna baru terakhir akan ditampilkan di sini.`,

		itIT: `Per aggiungere il tuo canale all'elenco, scrivi semplicemente di noi con un link come %v <code>&lt;-</code> sostituisci <code>YOUR_CHANNEL</code> con il tuo canale.

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

		plPL: `Aby dodać swój kanał do listy, po prostu napisz o nas z linkiem jako %v <code>&lt;-</code> zastąp <code>YOUR_CHANNEL</code> swoim własnym kanałem.

Lepiej, jeśli ukryjesz link w HTML jako:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Powinno to być renderowane przez klientów Telegram jako: <a href="%v">@%v</a>

Tutaj zostanie wyświetlonych 5 najlepszych polecających dla ostatnich 100 nowych użytkowników.`,

		ptBR: `Para adicionar seu canal à lista, basta escrever sobre nós com um link como %v <code>&lt;-</code> substitua <code>YOUR_CHANNEL</code> pelo seu próprio canal.

É melhor se você ocultar o link em HTML como:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Isso deve ser renderizado pelos clientes do Telegram como: <a href="%v">@%v</a>

Os 5 principais referenciadores para os últimos 100 novos usuários serão mostrados aqui.`,

		ruRU: `Чтобы добавить ваш канал в этот список просто напишите об этом боте использую ссылку вида %v <code>&lt;-</code> замените <code>YOUR_CHANNEL</code> на ваш канал.

Будет лучше  если вы спрячете её в HTML как:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Такой код должен отобразиться в Телеграмме как: <a href="%v">@%v</a>

Топ-5 источников последних 100 пользователей будут показаны здесь.`,

		trTR: `Kanalınızı listeye eklemek için sadece %v <code>&lt;-</code> gibi bir bağlantı ile hakkımızda yazın, <code>YOUR_CHANNEL</code> yerine kendi kanalınızı yazın.

Bağlantıyı HTML'de şu şekilde gizlerseniz daha iyi olur:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Bu, Telegram istemcileri tarafından şöyle görüntülenmelidir: <a href="%v">@%v</a>

Son 100 yeni kullanıcı için en iyi 5 referans burada gösterilecektir.`,

		ukUA: `Щоб додати свій канал до списку, просто напишіть про нас із посиланням як %v <code>&lt;-</code> замініть <code>YOUR_CHANNEL</code> на свій власний канал.

Краще, якщо ви приховаєте посилання в HTML як:

<pre>&lt;a href="%v"&gt;@%v&lt;/a&gt;</pre>

Це має відображатися клієнтами Telegram як: <a href="%v">@%v</a>

Тут буде показано 5 найкращих рекомендацій для останніх 100 нових користувачів.`,

		uzUZ: `Kanalingizni ro'yxatga qo'shish uchun shunchaki %v <code>&lt;-</code> kabi havola bilan biz haqimizda yozing, <code>YOUR_CHANNEL</code> o'rniga o'z kanalingizni yozing.

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
		deDE: "↩ Zurück",
		enUK: "↩ Cancel",
		enUS: "↩ Cancel",
		esES: "↩ Cancelar",
		faIR: "↪ کنسل",
		frFR: "↩ Annuler",
		idID: "↩ Batal",
		itIT: "↩ Annulla",
		jaJP: "↩ キャンセル",
		koKR: "↩ 취소",
		plPL: "↩ Anuluj",
		ptBR: "↩ Cancelar",
		ruRU: "↩ Отменить",
		trTR: "↩ İptal",
		ukUA: "↩ Скасувати",
		uzUZ: "↩ Bekor qilish",
		zhCN: "↩ 取消",
	},
	BUTTON_TEXT_MAIN_MENU: {
		deDE: "↩ Hauptmenü",
		enUK: "↩ Main menu",
		enUS: "↩ Main menu",
		esES: "↩ Menú principal",
		faIR: "↪ منوی اصلی",
		frFR: "↩ Menu principal",
		idID: "↩ Menu utama",
		itIT: "↩ Menu principale",
		jaJP: "↩ メインメニュー",
		koKR: "↩ 메인 메뉴",
		plPL: "↩ Menu główne",
		ptBR: "↩ Menu principal",
		ruRU: "↩ Главное меню",
		trTR: "↩ Ana menü",
		ukUA: "↩ Головне меню",
		uzUZ: "↩ Asosiy menyu",
		zhCN: "↩ 主菜单",
	},
	COMMAND_TEXT_SETTINGS_PRIMARY_CURRENCY: {
		deDE: "Hauptwährung",
		enUK: "Primary currency",
		enUS: "Primary currency",
		esES: "Moneda principal",
		faIR: "واحد پول اولیه",
		frFR: "Devise principale",
		idID: "Mata uang utama",
		itIT: "Valuta principale",
		jaJP: "主要通貨",
		koKR: "기본 통화",
		plPL: "Główna waluta",
		ptBR: "Moeda principal",
		ruRU: "Основная валюта",
		trTR: "Ana para birimi",
		ukUA: "Основна валюта",
		uzUZ: "Asosiy valyuta",
		zhCN: "主要货币",
	},
	COMMAND_TEXT_ADD_GROUP: {
		deDE: "Gruppe hinzufügen",
		enUK: "Add group",
		enUS: "Add group",
		esES: "Añadir grupo",
		faIR: "اضافه کردن گروه",
		frFR: "Ajouter un groupe",
		idID: "Tambahkan grup",
		itIT: "Aggiungi gruppo",
		jaJP: "グループを追加",
		koKR: "그룹 추가",
		plPL: "Dodaj grupę",
		ptBR: "Adicionar grupo",
		ruRU: "Добавить группу",
		trTR: "Grup ekle",
		ukUA: "Додати групу",
		uzUZ: "Guruh qo'shish",
		zhCN: "添加群组",
	},
	COMMAND_TEXT_GROUPS: {
		deDE: "Gruppen",
		enUK: "Groups",
		enUS: "Groups",
		esES: "Grupos",
		faIR: "گروه\u200cها",
		frFR: "Groupes",
		idID: "Grup",
		itIT: "Gruppi",
		jaJP: "グループ",
		koKR: "그룹",
		plPL: "Grupy",
		ptBR: "Grupos",
		ruRU: "Группы",
		trTR: "Gruplar",
		ukUA: "Групи",
		uzUZ: "Guruhlar",
		zhCN: "群组",
	},
	COMMAND_TEXT_BILLS: {
		deDE: "Rechnungen",
		enUK: "Bills",
		enUS: "Bills",
		esES: "Facturas",
		faIR: "صورتحساب\u200cها",
		frFR: "Factures",
		idID: "Tagihan",
		itIT: "Fatture",
		jaJP: "請求書",
		koKR: "청구서",
		plPL: "Rachunki",
		ptBR: "Contas",
		ruRU: "Счета",
		trTR: "Faturalar",
		ukUA: "Рахунки",
		uzUZ: "Hisob-fakturalar",
		zhCN: "账单",
	},
	COMMAND_TEXT_SETTLE_BILL: {
		deDE: "Rechnung begleichen",
		enUK: "Settle bill",
		enUS: "Settle bill",
		esES: "Liquidar factura",
		faIR: "تسویه صورتحساب",
		frFR: "Régler la facture",
		idID: "Selesaikan tagihan",
		itIT: "Saldare il conto",
		jaJP: "請求書を決済する",
		koKR: "청구서 정산",
		plPL: "Rozlicz rachunek",
		ptBR: "Quitar conta",
		ruRU: "Оплатить счёт",
		trTR: "Fatura öde",
		ukUA: "Оплатити рахунок",
		uzUZ: "Hisob-fakturani to'lash",
		zhCN: "结算账单",
	},
	COMMAND_TEXT_SETTLE_BILLS: {
		deDE: "Rechnungen begleichen",
		enUK: "Settle bills",
		enUS: "Settle bills",
		esES: "Liquidar facturas",
		faIR: "تسویه صورتحساب\u200cها",
		frFR: "Régler les factures",
		idID: "Selesaikan tagihan",
		itIT: "Saldare i conti",
		jaJP: "請求書を決済する",
		koKR: "청구서 정산",
		plPL: "Rozlicz rachunki",
		ptBR: "Quitar contas",
		ruRU: "Закрыть счета",
		trTR: "Faturaları öde",
		ukUA: "Оплатити рахунки",
		uzUZ: "Hisob-fakturalarni to'lash",
		zhCN: "结算账单",
	},
	COMMAND_TEXT_INVITE_FIREND: {
		deDE: "Freund einladen",
		enUK: "Invite friend",
		enUS: "Invite friend",
		esES: "Invitar a un amigo",
		faIR: "دوستی دعوت کن",
		frFR: "Inviter un ami",
		idID: "Undang teman",
		itIT: "Invita un amico",
		jaJP: "友達を招待",
		koKR: "친구 초대",
		plPL: "Zaproś przyjaciela",
		ptBR: "Convidar amigo",
		ruRU: "Пригласить друга",
		trTR: "Arkadaş davet et",
		ukUA: "Запросити друга",
		uzUZ: "Do'stni taklif qilish",
		zhCN: "邀请朋友",
	},
	COMMAND_TEXT_INVITE_MEMBER: {
		deDE: "Mitglied einladen",
		enUK: "Invite member",
		enUS: "Invite member",
		esES: "Invitar miembro",
		faIR: "دعوت از اعضا",
		frFR: "Inviter un membre",
		idID: "Undang anggota",
		itIT: "Invita membro",
		jaJP: "メンバーを招待",
		koKR: "멤버 초대",
		plPL: "Zaproś członka",
		ptBR: "Convidar membro",
		ruRU: "Пригласить участника",
		trTR: "Üye davet et",
		ukUA: "Запросити учасника",
		uzUZ: "A'zoni taklif qilish",
		zhCN: "邀请成员",
	},
	COMMAND_TEXT_NEW_BILL: {
		deDE: "Neue Rechnung",
		enUK: "New bill",
		enUS: "New bill",
		esES: "Nueva factura",
		faIR: "صورتحساب جدید",
		frFR: "Nouvelle facture",
		idID: "Tagihan baru",
		itIT: "Nuova fattura",
		jaJP: "新しい請求書",
		koKR: "새 청구서",
		plPL: "Nowy rachunek",
		ptBR: "Nova conta",
		ptPT: "Nova conta",
		ruRU: "Новый счёт",
		trTR: "Yeni fatura",
		ukUA: "Новий рахунок",
		uzUZ: "Yangi hisob-faktura",
		zhCN: "新账单",
	},
	COMMAND_TEXT_NEW_FUNDRAISING: {
		deDE: "Neue Spendensammlung",
		enUK: "New fundraising",
		enUS: "New fundraising",
		esES: "Nueva recaudación de fondos",
		faIR: "جمع آوری پول جدید",
		frFR: "Nouvelle collecte de fonds",
		idID: "Penggalangan dana baru",
		itIT: "Nuova raccolta fondi",
		jaJP: "新しい資金調達",
		koKR: "새 모금",
		plPL: "Nowa zbiórka pieniędzy",
		ptBR: "Nova arrecadação de fundos",
		ruRU: "Новый сбор средств",
		trTR: "Yeni bağış toplama",
		ukUA: "Новий збір коштів",
		uzUZ: "Yangi mablag' yig'ish",
		zhCN: "新筹款",
	},
	COMMAND_TEXT_NEW_COUNTERPARTY: {
		deDE: "neuer Kontakt",
		enUK: "Add new",
		enUS: "Add new",
		esES: "Añadir",
		faIR: "اضافه کردن مورد جدید",
		frFR: "Ajouter nouveau",
		idID: "Tambah baru",
		itIT: "Aggiungi nuovo",
		jaJP: "新規追加",
		koKR: "새로 추가",
		plPL: "Dodaj nowy",
		ptBR: "Adicionar novo",
		ruRU: "Добавить",
		trTR: "Yeni ekle",
		ukUA: "Додати новий",
		uzUZ: "Yangi qo'shish",
		zhCN: "添加新的",
	},
	MESSAGE_TEXT_LOGIN_CODE: {
		deDE: "Dein Code um dich an der App anzumelden: <b>%v</b>",
		enUK: "Your code for signing in to app: <b>%v</b>",
		enUS: "Your code for signing in to app: <b>%v</b>",
		esES: "Tu código para entrar en la app: <b>%v</b>",
		faIR: "کد شما برای ورود به برنامه: <b>%v</b>",
		frFR: "Votre code pour vous connecter à l'application: <b>%v</b>",
		idID: "Kode Anda untuk masuk ke aplikasi: <b>%v</b>",
		itIT: "Il tuo codice per accedere all'app e': <b>%v</b>",
		jaJP: "アプリにサインインするためのコード: <b>%v</b>",
		koKR: "앱에 로그인하기 위한 코드: <b>%v</b>",
		plPL: "Twój kod do logowania do aplikacji: <b>%v</b>",
		ptBR: "Seu código para entrar no aplicativo: <b>%v</b>",
		ruRU: "Ваш код для входа в приложение: <b>%v</b>",
		trTR: "Uygulamaya giriş yapma kodunuz: <b>%v</b>",
		ukUA: "Ваш код для входу в додаток: <b>%v</b>",
		uzUZ: "Ilovaga kirish uchun kodingiz: <b>%v</b>",
		zhCN: "您登录应用的验证码: <b>%v</b>",
	},
	MESSAGE_TEXT_ASK_NEW_COUNTERPARTY_NAME: {
		deDE: `Bitte gib einen Namen für den neuen Kontakt ein:
		Du kannst in eintippen oder aus deinem Adressbuch wählen (<i>mit dem "Büroklammer"-Symbol und dann Kontakt</i>).

		<i>Send '.' to cancel</i>`,

		enUK: `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		enUS: `Please enter a name for the new contact:
		You can type manually or choose from your address book (<i>through "clip" icon</i>).

		<i>Send '.' to cancel</i>`,

		esES: `Escribe un nombre para el nuevo contacto:
		Puedes escribirlo o elegirlo de tus contactos (<i>a traves del icono "clip"</i>).

		<i>Enviar '.' para anular</i>`,

		faIR: `لطفا برای مخاطب جدید یک نام وارد کنید:
		میتوانید به صورت دستی تایپ نموده یا از لیست مخاطبین خود انتخاب نمایید (<i>throw "clip" icon</i>).

		<i>Send '.' برای کنسل کردن</i>`,

		frFR: `Veuillez entrer un nom pour le nouveau contact:
		Vous pouvez taper manuellement ou choisir dans votre carnet d'adresses (<i>via l'icône "trombone"</i>).

		<i>Envoyez '.' pour annuler</i>`,

		idID: `Silakan masukkan nama untuk kontak baru:
		Anda dapat mengetik secara manual atau memilih dari buku alamat Anda (<i>melalui ikon "klip"</i>).

		<i>Kirim '.' untuk membatalkan</i>`,

		itIT: `Inserisci un nome per il nuovo contatto:
		Puoi digitarlo o sceglierlo dalla tua rubrica (<i>attraverso l'icona "clip"</i>).

		<i>Digita '.' ed invia per annullare</i>`,

		jaJP: `新しい連絡先の名前を入力してください:
		手動で入力するか、アドレス帳から選択できます（<i>「クリップ」アイコンを通じて</i>）。

		<i>キャンセルするには '.' を送信してください</i>`,

		koKR: `새 연락처의 이름을 입력하세요:
		수동으로 입력하거나 주소록에서 선택할 수 있습니다 (<i>"클립" 아이콘을 통해</i>).

		<i>취소하려면 '.'를 보내세요</i>`,

		plPL: `Wprowadź nazwę dla nowego kontaktu:
		Możesz wpisać ręcznie lub wybrać z książki adresowej (<i>przez ikonę "spinacza"</i>).

		<i>Wyślij '.' aby anulować</i>`,

		ptBR: `Por favor, digite um nome para o novo contato:
		Você pode digitar manualmente ou escolher do seu livro de endereços (<i>através do ícone "clipe"</i>).

		<i>Envie '.' para cancelar</i>`,

		ruRU: `<b>Имя для нового контакта</b>
		Напишите сами или выберите из своей адресной книги (<i>через иконку "скрепка"</i>).
		<i>Отправьте '.' для отмены</i>`,

		trTR: `Yeni kişi için bir isim girin:
		Manuel olarak yazabilir veya adres defterinizden seçebilirsiniz (<i>"ataç" simgesi aracılığıyla</i>).

		<i>İptal etmek için '.' gönderin</i>`,

		ukUA: `Будь ласка, введіть ім'я для нового контакту:
		Ви можете ввести вручну або вибрати з адресної книги (<i>через іконку "скріпка"</i>).

		<i>Надішліть '.' для скасування</i>`,

		uzUZ: `Yangi kontakt uchun ism kiriting:
		Siz qo'lda yozishingiz yoki manzillar kitobingizdan tanlashingiz mumkin (<i>"qisqich" belgisi orqali</i>).

		<i>Bekor qilish uchun '.' yuboring</i>`,

		zhCN: `请输入新联系人的名称:
		您可以手动输入或从通讯录中选择（<i>通过"回形针"图标</i>）。

		<i>发送 '.' 取消</i>`,
	},
	MESSAGE_TEXT_TRANSFER_IS_CREATING: {
		deDE: "Transferiere...",
		enUK: "Creating transfer...",
		enUS: "Creating transfer...",
		esES: "Estoy creando la nueva nota...",
		faIR: "ایجاد انتقال ...",
		frFR: "Création du transfert...",
		idID: "Membuat transfer...",
		itIT: "Sto creando la nuova voce...",
		jaJP: "転送を作成中...",
		koKR: "전송 생성 중...",
		plPL: "Tworzenie transferu...",
		ptBR: "Criando transferência...",
		ruRU: "Создаю запись...",
		trTR: "Transfer oluşturuluyor...",
		ukUA: "Створення переказу...",
		uzUZ: "O'tkazma yaratilmoqda...",
		zhCN: "创建转账中...",
	},
	COMMAND_TEXT_PLEASE_WAIT: {
		deDE: "Bitte warten",
		enUK: "Please wait",
		enUS: "Please wait",
		esES: "Espera, por favor",
		faIR: "لطفا صبر کنید",
		frFR: "Veuillez patienter",
		idID: "Mohon tunggu",
		itIT: "Aspetta per favore",
		jaJP: "お待ちください",
		koKR: "잠시만 기다려주세요",
		plPL: "Proszę czekać",
		ptBR: "Por favor, aguarde",
		ruRU: "Пожалуйста подождите",
		trTR: "Lütfen bekleyin",
		ukUA: "Будь ласка, зачекайте",
		uzUZ: "Iltimos kuting",
		zhCN: "请稍等",
	},
	MESSAGE_TEXT_PLEASE_WAIT: {
		deDE: "Bitte warten...",
		enUK: "Please wait...",
		enUS: "Please wait...",
		esES: "Espera, por favor...",
		faIR: "لطفا صبر کنید ...",
		frFR: "Veuillez patienter...",
		idID: "Mohon tunggu...",
		itIT: "Aspetta per favore...",
		jaJP: "お待ちください...",
		koKR: "잠시만 기다려주세요...",
		plPL: "Proszę czekać...",
		ptBR: "Por favor, aguarde...",
		ruRU: "Пожалуйста подождите...",
		trTR: "Lütfen bekleyin...",
		ukUA: "Будь ласка, зачекайте...",
		uzUZ: "Iltimos kuting...",
		zhCN: "请稍等...",
	},
	MESAGE_TEXT_CREATING_BILL: {
		deDE: "Rechnung erstellen",
		enUK: "Creating bill",
		enUS: "Creating bill",
		esES: "Creando factura",
		faIR: "ایجاد صورتحساب",
		frFR: "Création de facture",
		idID: "Membuat tagihan",
		itIT: "Creazione di fattura",
		jaJP: "請求書を作成中",
		koKR: "청구서 생성 중",
		plPL: "Tworzenie rachunku",
		ptBR: "Criando conta",
		ruRU: "Создаётся счёт",
		trTR: "Fatura oluşturuluyor",
		ukUA: "Створення рахунку",
		uzUZ: "Hisob-faktura yaratilmoqda",
		zhCN: "创建账单中",
	},
	MESSAGE_TEXT_ASK_BILL_CURRENCY: {
		deDE: "In welcher Währung ist die Rechnung?",
		enUK: "What currency this bill in?",
		enUS: "What currency is this bill in?",
		esES: "¿En qué moneda está esta factura?",
		faIR: "این صورتحساب به چه ارزی است؟",
		frFR: "Dans quelle devise est cette facture?",
		idID: "Dalam mata uang apa tagihan ini?",
		itIT: "In quale valuta è questa fattura?",
		jaJP: "この請求書の通貨は何ですか？",
		koKR: "이 청구서의 통화는 무엇입니까?",
		plPL: "W jakiej walucie jest ten rachunek?",
		ptBR: "Em qual moeda está esta conta?",
		ruRU: "В какой валюте этот счёт?",
		trTR: "Bu fatura hangi para biriminde?",
		ukUA: "У якій валюті цей рахунок?",
		uzUZ: "Bu hisob-faktura qaysi valyutada?",
		zhCN: "这个账单使用什么货币？",
	},
	MESSAGE_TEXT_ASK_BILL_PAYER: {
		deDE: "Wer hat die Rechnung bezahlt?",
		enUK: "Who paid for the bill?",
		enUS: "Who paid for the bill?",
		esES: "¿Quién pagó la factura?",
		faIR: "چه کسی صورتحساب را پرداخت کرد؟",
		frFR: "Qui a payé la facture?",
		idID: "Siapa yang membayar tagihan?",
		itIT: "Chi ha pagato il conto?",
		jaJP: "誰が請求書を支払いましたか？",
		koKR: "누가 청구서를 지불했습니까?",
		plPL: "Kto zapłacił rachunek?",
		ptBR: "Quem pagou a conta?",
		ruRU: "Кто оплатил счёт?",
		trTR: "Faturayı kim ödedi?",
		ukUA: "Хто оплатив рахунок?",
		uzUZ: "Hisob-fakturani kim to'ladi?",
		zhCN: "谁支付了账单？",
	},
	MESSAGE_TEXT_SELF_ACKNOWLEDGEMENT: {
		deDE: "%v muss dem zustimmen",
		enUK: "Acknowledgement is expected from %v",
		enUS: "Acknowledgement is expected from %v",
		esES: "Se espera la confirmación de %v",
		faIR: "انتظار تصدیق می رود از %v",
		frFR: "Confirmation attendue de %v",
		idID: "Pengakuan diharapkan dari %v",
		itIT: "Conferma in attesa da %v",
		jaJP: "%vからの確認が必要です",
		koKR: "%v의 확인이 필요합니다",
		plPL: "Oczekiwane potwierdzenie od %v",
		ptBR: "Confirmação esperada de %v",
		ruRU: "Подтверждение ожидается от %v",
		trTR: "%v'den onay bekleniyor",
		ukUA: "Очікується підтвердження від %v",
		uzUZ: "%v dan tasdiqlash kutilmoqda",
		zhCN: "等待%v的确认",
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_YOU: {
		deDE: "Du hast dem zugestimmt.",
		enUK: "You've accepted this transaction.",
		enUS: "You've accepted this transaction.",
		esES: "Has confirmado esta transacción",
		faIR: ".شما این تراکنش را قبول کردید ",
		frFR: "Vous avez accepté cette transaction.",
		idID: "Anda telah menerima transaksi ini.",
		itIT: "Hai accettato il debito/credito.",
		jaJP: "このトランザクションを承認しました。",
		koKR: "이 거래를 수락했습니다.",
		plPL: "Zaakceptowałeś tę transakcję.",
		ptBR: "Você aceitou esta transação.",
		ruRU: "Вы подтвердили эту транзакцию.",
		trTR: "Bu işlemi kabul ettiniz.",
		ukUA: "Ви підтвердили цю транзакцію.",
		uzUZ: "Siz ushbu tranzaksiyani qabul qildingiz.",
		zhCN: "您已接受此交易。",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_YOU: {
		deDE: `Du hast dieser Anfrage nicht zugestimmt.
		Der Vorgang wird zurückgestellt und die Gegenpartei benachrichtigt.`,

		enUK: `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,

		enUS: `You do not agree with this transaction.
                The transaction will not be deleted but the creator will be notified.`,

		esES: `No estas de acuerdo con la transacción.
		La transacción NO será cancelada, pero el creador será notificado.`,

		faIR: `شما با این تراکنش موافق نیستید.
		تراکنش حذف نخواهد شد اما به ایجاد کننده اطلاع داده خواهد شد.`,

		frFR: `Vous n'êtes pas d'accord avec cette transaction.
		La transaction ne sera pas supprimée mais le créateur sera notifié.`,

		idID: `Anda tidak setuju dengan transaksi ini.
		Transaksi tidak akan dihapus tetapi pembuat akan diberi tahu.`,

		itIT: `Hai rifiutato il debito/credito.
		La transazione non sarà eliminata ma il creatore sarà avvisato.`,

		jaJP: `このトランザクションに同意しません。
		トランザクションは削除されませんが、作成者に通知されます。`,

		koKR: `이 거래에 동의하지 않습니다.
		거래는 삭제되지 않지만 생성자에게 알림이 갑니다.`,

		plPL: `Nie zgadzasz się z tą transakcją.
		Transakcja nie zostanie usunięta, ale twórca zostanie powiadomiony.`,

		ptBR: `Você não concorda com esta transação.
		A transação não será excluída, mas o criador será notificado.`,

		ruRU: `Вы НЕ согласны с этой транзакцией.

Сама транзакция НЕ будет отменена, но создатель будет оповещён.`,

		trTR: `Bu işleme katılmıyorsunuz.
		İşlem silinmeyecek ancak oluşturucu bilgilendirilecek.`,

		ukUA: `Ви не згодні з цією транзакцією.
		Транзакція не буде видалена, але творець буде повідомлений.`,

		uzUZ: `Siz ushbu tranzaksiyaga rozi emassiz.
		Tranzaksiya o'chirilmaydi, lekin yaratuvchi xabardor qilinadi.`,

		zhCN: `您不同意此交易。
		交易不会被删除，但创建者会收到通知。`,
	},
	MESSAGE_TEXT_TRANSFER_ACCEPTED_BY_COUNTERPARTY: {
		deDE: "%v hat deiner Anfrage <b>zugestimmt</b>:",
		enUK: "%v accepted your transaction:",
		enUS: "%v accepted your transaction:",
		esES: "%v ha aceptado tu transacción",
		faIR: ": تراکنش شمارا تایید کرد %v ",
		frFR: "%v a accepté votre transaction:",
		idID: "%v menerima transaksi Anda:",
		itIT: "%v ha accettato il tuo credito/debito:",
		jaJP: "%vがあなたの取引を承認しました:",
		koKR: "%v님이 귀하의 거래를 수락했습니다:",
		plPL: "%v zaakceptował(a) twoją transakcję:",
		ptBR: "%v aceitou sua transação:",
		ruRU: "%v подтвердил(a) вашу транзакцию:",
		trTR: "%v işleminizi kabul etti:",
		ukUA: "%v підтвердив(ла) вашу транзакцію:",
		uzUZ: "%v sizning tranzaksiyangizni qabul qildi:",
		zhCN: "%v接受了您的交易：",
	},
	MESSAGE_TEXT_TRANSFER_DECLINED_BY_COUNTERPARTY: {
		deDE: "%v hat deine Anfrage <b>abgelehnt</b>. Wenn die Sache besprochen ist, kann die Anfrage erneut gesendet werden.",
		enUK: "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.",
		enUS: "%v did not agree with your transaction. The transaction is not canceled but you may want to discuss it.",
		esES: "%v no está de acuerdo con tu transacción. La transacción no ha sido cancelada, pero quizás deberías discutirlo.",
		faIR: "%v با تراکنش شما موافقت نکرد. تراکنش لغو نشده است اما ممکن است بخواهید در مورد آن صحبت کنید.",
		frFR: "%v n'est pas d'accord avec votre transaction. La transaction n'est pas annulée mais vous voudrez peut-être en discuter.",
		idID: "%v tidak setuju dengan transaksi Anda. Transaksi tidak dibatalkan tetapi Anda mungkin ingin mendiskusikannya.",
		itIT: "%v non è d'accordo con la tua transazione. La transazione non è annullata ma potresti voler discuterne.",
		jaJP: "%vはあなたの取引に同意しませんでした。取引はキャンセルされていませんが、話し合いたいかもしれません。",
		koKR: "%v님이 귀하의 거래에 동의하지 않았습니다. 거래는 취소되지 않았지만 논의하고 싶을 수 있습니다.",
		plPL: "%v nie zgadza się z twoją transakcją. Transakcja nie jest anulowana, ale możesz chcieć to omówić.",
		ptBR: "%v não concordou com sua transação. A transação não foi cancelada, mas você pode querer discutir isso.",
		ruRU: "%v <b>НЕ</b> подтвердил(a) вашу транзакцию. Транзакция не отменена, но возможно вам стоит это обсудить.",
		trTR: "%v işleminizle aynı fikirde değil. İşlem iptal edilmedi ancak bunu tartışmak isteyebilirsiniz.",
		ukUA: "%v не погодився з вашою транзакцією. Транзакція не скасована, але, можливо, ви захочете це обговорити.",
		uzUZ: "%v sizning tranzaksiyangizga rozi bo'lmadi. Tranzaksiya bekor qilinmadi, lekin siz bu haqda muhokama qilishni xohlashingiz mumkin.",
		zhCN: "%v不同意您的交易。交易未取消，但您可能想讨论一下。",
	},
	COMMAND_TEXT_SUBSCRIBE_TO_APP: {
		deDE: "Ich will die App!",
		enUK: "I want the app!",
		enUS: "I want the app!",
		esES: "¡Quiero la aplicación!",
		faIR: "!من برنامه را می خواهم",
		frFR: "Je veux l'application !",
		idID: "Saya ingin aplikasinya!",
		itIT: "Voglio l'aplicazione!",
		jaJP: "アプリが欲しいです！",
		koKR: "앱이 필요합니다!",
		plPL: "Chcę aplikację!",
		ptBR: "Eu quero o aplicativo!",
		ruRU: "Хочу приложение!",
		trTR: "Uygulamayı istiyorum!",
		ukUA: "Я хочу додаток!",
		uzUZ: "Men ilovani xohlayman!",
		zhCN: "我想要应用程序！",
	},
	COMMAND_TEXT_I_AM_FINE_WITH_BOT: {
		deDE: "Mir reicht der Bot!",
		enUK: "I'm fine with just the bot!",
		enUS: "I'm fine with just the bot!",
		esES: "¡Estoy satisfecho con este bot!",
		faIR: "! ربات به تنهایی برای من کافی است",
		frFR: "Le bot me suffit !",
		idID: "Saya cukup dengan bot saja!",
		itIT: "Mi accontento del bot per ora!",
		jaJP: "ボットだけで大丈夫です！",
		koKR: "봇만으로도 괜찮습니다!",
		plPL: "Wystarczy mi sam bot!",
		ptBR: "Estou bem apenas com o bot!",
		ruRU: "Меня вполне устраивает бот!",
		trTR: "Sadece bot ile iyiyim!",
		ukUA: "Мені достатньо лише бота!",
		uzUZ: "Men faqat bot bilan yaxshiman!",
		zhCN: "我对只使用机器人很满意！",
	},
	MESSAGE_TEXT_SUBSCRIBED_TO_APP: {
		deDE: "Du wirst darüber informiert, wenn die App zum Download zur Verfügung steht.",
		enUK: "We'll let you know once the app is available for download.",
		enUS: "We'll let you know once the app is available for download.",
		esES: "Te avisamos cuando la aplicación esté disponible para descargarla",
		faIR: ".وقتی برنامه برای دانلود دردسترس بود به شما اطلاع می دهیم",
		frFR: "Nous vous informerons dès que l'application sera disponible au téléchargement.",
		idID: "Kami akan memberi tahu Anda setelah aplikasi tersedia untuk diunduh.",
		itIT: "Ti faremo sapere non appena l'applicazione sara' disponibile al download.",
		jaJP: "アプリがダウンロード可能になり次第お知らせします。",
		koKR: "앱이 다운로드 가능해지면 알려드리겠습니다.",
		plPL: "Poinformujemy Cię, gdy aplikacja będzie dostępna do pobrania.",
		ptBR: "Avisaremos você assim que o aplicativo estiver disponível para download.",
		ruRU: "Мы сообщим вам когда приложение будет доступно для загруки.",
		trTR: "Uygulama indirilebilir olduğunda size haber vereceğiz.",
		ukUA: "Ми повідомимо вас, коли додаток буде доступний для завантаження.",
		uzUZ: "Ilova yuklab olish uchun mavjud bo'lganda sizga xabar beramiz.",
		zhCN: "一旦应用程序可供下载，我们会通知您。",
	},
	MESSAGE_TEXT_NOT_INTERESTED_IN_APP: {
		deDE: "Gut, wir sind froh, dass dir der Bot reicht und wir uns mit der App nicht beeilen müssen.",
		enUK: "Well, we are happy our bot is good enough and there is no need to download an app.",
		enUS: "Well, we are happy our bot is good enough and there is no need to download an app.",
		esES: "Bueno, estamos contentos de que te haya gustado nuestro bot y no hace falta descargar ninguna otra aplicación",
		faIR: ".خب، ما خوشحال هستیم که ربات برای شما کافی است و نیازی به دانلود برنامه نیست",
		frFR: "Eh bien, nous sommes heureux que notre bot soit suffisant et qu'il n'y ait pas besoin de télécharger une application.",
		idID: "Baiklah, kami senang bot kami cukup baik dan tidak perlu mengunduh aplikasi.",
		itIT: "Bene, siamo contenti che il nostro bot sia di tuo gradimento e non hai bisogno di scaricare l'applicazione.",
		jaJP: "私たちのボットで十分であり、アプリをダウンロードする必要がないことを嬉しく思います。",
		koKR: "네, 저희 봇이 충분히 좋아서 앱을 다운로드할 필요가 없다니 기쁩니다.",
		plPL: "Cóż, cieszymy się, że nasz bot jest wystarczająco dobry i nie ma potrzeby pobierania aplikacji.",
		ptBR: "Bem, estamos felizes que nosso bot seja bom o suficiente e não há necessidade de baixar um aplicativo.",
		ruRU: "Что ж, мы рады что вас устраивает наш бот и нет необходимости загружать приложение.",
		trTR: "Peki, botumuzun yeterince iyi olduğuna ve bir uygulama indirmeye gerek olmadığına sevindik.",
		ukUA: "Що ж, ми раді, що наш бот достатньо хороший і немає потреби завантажувати додаток.",
		uzUZ: "Yaxshi, botimiz yetarlicha yaxshi ekanligi va ilovani yuklab olish kerak emasligi bizni xursand qiladi.",
		zhCN: "好的，我们很高兴我们的机器人已经足够好，不需要下载应用程序。",
	},
	MESSAGE_TEXT_YOUR_AD_COULD_BE_HERE: {
		deDE: "Hier könnte <a href>ihre Werbung</a> stehen",
		enUK: "You can <a href>advertise here</a>",
		enUS: "You can <a href>advertise here</a>",
		esES: "Aquí se puede <a href>publicar un anuncio</a>",
		faIR: "شما میتوانید <a href>در اینجا تبلیغ کنید</a>",
		frFR: "Vous pouvez <a href>faire de la publicité ici</a>",
		idID: "Anda dapat <a href>beriklan di sini</a>",
		itIT: "Puoi <a href>pubblicizzare qui</a>",
		jaJP: "ここに<a href>広告を掲載</a>できます",
		koKR: "여기에 <a href>광고할 수 있습니다</a>",
		plPL: "Możesz <a href>reklamować się tutaj</a>",
		ptBR: "Você pode <a href>anunciar aqui</a>",
		ruRU: "Здесь можно <a href>разместить рекламу</a>",
		trTR: "Burada <a href>reklam verebilirsiniz</a>",
		ukUA: "Ви можете <a href>розмістити рекламу тут</a>",
		uzUZ: "Siz bu yerda <a href>reklama bera olasiz</a>",
		zhCN: "您可以在这里<a href>做广告</a>",
	},
	MESSAGE_TEXT_YOUR_ABOUT_ADS: {
		deDE: `🤖: Ich hin ein guter Roboter - klar. Aber manchmal kommt es besser eine eigene App für etwas zu haben. Es ist noch nicht ganz fertig, aber falls du schonmal reinschauen willst: <a href="https://debtstracker.io/de/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Möchtest du daran erinnert werden, wenn die App rauskommt?`,
		enUK: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,
		enUS: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/en/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	Do you want to get an invite when it gets released?`,

		esES: `🤖: Claro que soy un robot encantador, pero más comodo usar una aplicación especial.No esta disponible ya pero se puede ver como será: <a href = "https://debtstracker.io/es/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	¿Quieres que te avisemos cuando esté lista?`,
		faIR: `🤖: مطمئناً من روبات خوبی هستم. اما بعضی وقت هاساده تر و مناسب تر است که از یک برنامه به خوبی تخصصی شده استفاده شود، این برنامه هنوز برای استفاده عموم آماده نیست ولی می توانید چک کنید که چگونه به نظر خواهد رسید: <a href="https://debtstracker.io/fa/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/en/</a>

	آیا می خواهید وقتی منتشر شد دعوتنامه ای دریافت کنید؟`,
		frFR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/fr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/fr/</a>

	Do you want to get an invite when it gets released?`,
		idID: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/id/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/id/</a>

	Do you want to get an invite when it gets released?`,

		itIT: `🤖: Di sicuro son un bravo bot, ma alcune volte e' piu' conveniente usare un'applicazione specializzata. Non e' ancora pronta per la pubblicazione ma puoi controllare l'avanzamento a questo indirizzo: <a href="https://debtstracker.io/it/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/it/</a>

	Vuoi essere invitato non appena viene rilasciata?`,
		jaJP: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ja/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ja/</a>

	Do you want to get an invite when it gets released?`,
		koKR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ko/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ko/</a>

	Do you want to get an invite when it gets released?`,
		plPL: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pl/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pl/</a>

	Do you want to get an invite when it gets released?`,
		ptBR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/pt/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/pt/</a>

	Do you want to get an invite when it gets released?`,

		ruRU: `🤖: Я конечно обоятельный робот, но пользоваться специализированным приложением бывает удобнее. Оно ещё не готово для общего доступа, но уже сейчас можно посмотреть как будет выглядеть: <a href="https://debtstracker.io/ru/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ru/</a>

		Хотите получить оповещение когда оно выйдет?`,
		trTR: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/tr/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/tr/</a>

	Do you want to get an invite when it gets released?`,
		ukUA: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/ua/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/ua/</a>

	Do you want to get an invite when it gets released?`,
		uzUZ: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/uz/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/uz/</a>

	Do you want to get an invite when it gets released?`,
		zhCN: `🤖: I'm a good robot, for sure.But sometimes it is more convinient to use a nice specialized app.It's not ready for public use yet but you can check how it is going to looks: <a href="https://debtstracker.io/zh/#utm_source=telegram&utm_campaign=ads_screen">https://debtstracker.io/zh/</a>

	Do you want to get an invite when it gets released?`,
	},
	MESSAGE_TEXT_INVALID_FLOAT: {
		deDE: "Entschuldigung, aber du kannst nur Zahlen für Menge oder Wert wählen (<i>mit zwei Nachkommastellen</i>).",
		enUK: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		enUS: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		esES: "Lo siento, solo puedes utilizar numeros como importe/cantidad (<i>con un maximo de 2 dígitos despues de la coma</i>).",
		faIR: "ببخشید، اما شما تنها میتوانید از اعداد بعنوان مقادیر / اندازه ها استفاده کنید (<i>با دو رقم اعشار</i>).",
		frFR: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		idID: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		itIT: "Spiacente, puoi utilizzare solo numeri (<i>con un massimo di 2 numeri dopo il punto</i>).",
		jaJP: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		koKR: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		plPL: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		ptBR: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		ruRU: "Извините, но вы можете использовать только числа в качестве суммы/количества (<i>до 2х знаков после запятой</i>).",
		trTR: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		ukUA: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		uzUZ: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
		zhCN: "Sorry, but you can use just numbers as amount/quantity (<i>with up to 2 digits after point</i>).",
	},
	MESSAGE_TEXT_ASK_LENDING_TYPE: {
		deDE: "<b>Was hast du jemanden geliehen?</b>",
		enUK: "<b>What did you lend to someone?</b>",
		enUS: "<b>What did you lend to someone?</b>",
		esES: "<b>¿Qué has prestado?</b>",
		faIR: "<b> چه چیزی به کسی قرض داده اید؟</b>",
		frFR: "<b>What did you lend to someone?</b>",
		idID: "<b>What did you lend to someone?</b>",
		itIT: "<b>Cos'hai prestato?</b>",
		jaJP: "<b>What did you lend to someone?</b>",
		koKR: "<b>What did you lend to someone?</b>",
		plPL: "<b>What did you lend to someone?</b>",
		ptBR: "<b>What did you lend to someone?</b>",
		ruRU: "<b>Что вы дали в долг?</b>",
		trTR: "<b>What did you lend to someone?</b>",
		ukUA: "<b>What did you lend to someone?</b>",
		uzUZ: "<b>What did you lend to someone?</b>",
		zhCN: "<b>What did you lend to someone?</b>",
	},
	MESSAGE_TEXT_CHOOSE_CURRENCY: {

		deDE: `Bitte wähle <a>eine Währung aus der Liste</a>.

	Falls die Standardoptionen nicht reichen, sende mir einen Text. Zum Beispiel: <i>Äpfel</i>".`,

		enUK: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		enUS: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		esES: `Elige del menú abajo de la pantalla o <a>selecciona la moneda de la lista</a>.

	Si no encuentras la opción correcta simplemente envía un texto. Por ejemplo: "<i>manzana</i>".`,

		faIR: `لطفا از بین گزینه های زیر انتخاب کنید یا <a>یک واحد پولی از لیست انتخاب کنید</a>.

	اگر گزینه های استاندارد کافی نبودند به سادگی یک متن بفرستید ، برای مثال:. "<i>سیب</i>".`,

		frFR: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		idID: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		itIT: `Scegli dalle opzioni qui sotto o <a>seleziona una valuta dalla lista</a>.

	Se le opzioni standard non bastano semplicemente invia un testo.Per esempio: "<i>mele</i>".`,

		jaJP: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		koKR: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		plPL: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		ptBR: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		ruRU: `Выберите из меню внизу экрана или <a>выберите валюту из списка</a>.

	Если ни один из стандартных вариантов не подходит просто напишите текстом.Например: "<i>яблоко</i>".`,

		trTR: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		ukUA: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		uzUZ: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,

		zhCN: `Please choose from the options below or <a>select a currency from the list</a>.

	If standard options are not enough simply send a text.For example: "<i>apple</i>".`,
	},
	MESSAGE_TEXT_ASK_LENDING_AMOUNT: {
		deDE: "Wie viel <b>%v</b> hast du verliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		enUK: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		enUS: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		esES: "Cuanto(s) <b>%v</b> has prestado\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه مقدار <b>%v</b> قرض داده اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		frFR: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		idID: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		itIT: "Quanti <b>%v</b> hai prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		koKR: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		plPL: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		ptBR: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		ruRU: "Сколько <b>%v</b> вы дали в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		trTR: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		ukUA: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		uzUZ: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
		zhCN: "How much <b>%v</b> did you lend?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_LENDING_COUNTERPARTY: {
		deDE: "Wer hat sich <b>%v</b> von dir geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		enUK: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		enUS: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		esES: "A quién has prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه کسی از شما <b>%v</b> قرض گرفته است؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		frFR: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		idID: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		itIT: "Chi e' in debito di <b>%v</b> con te?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		koKR: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		plPL: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ptBR: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ruRU: "Кому вы дали в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		trTR: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ukUA: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		uzUZ: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		zhCN: "Who borrowed from you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_TYPE: {
		deDE: "Was hast du dir geliehen?",
		enUK: "What did you lend?",
		enUS: "What did you lend?",
		esES: "¿Qué te han prestado?",
		faIR: "چه چیزی قرض گرفته اید؟",
		frFR: "What did you lend?",
		idID: "What did you lend?",
		itIT: "Cosa ti hanno prestato?",
		jaJP: "What did you lend?",
		koKR: "What did you lend?",
		plPL: "What did you lend?",
		ptBR: "What did you lend?",
		ruRU: "Что вы взяли в долг?",
		trTR: "What did you lend?",
		ukUA: "What did you lend?",
		uzUZ: "What did you lend?",
		zhCN: "What did you lend?",
	},
	MESSAGE_TEXT_ASK_BORROWING_AMOUNT: {
		deDE: "Wie viel <b>%v</b> hast du geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		enUK: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		enUS: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		esES: "¿Cuánto <b>%v</b> has prestado?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه مقدار <b>%v</b> قرض گرفته اید؟\n(<i>ارسال '.' برای کنسل کردن</i>)",
		frFR: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		idID: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		itIT: "Quanti <b>%v</b> ti hanno prestato?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		koKR: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		plPL: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		ptBR: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		ruRU: "Сколько <b>%v</b> вы взяли в долг?\n(<i>отправьте '.' чтобы отменить</i>)",
		trTR: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		ukUA: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		uzUZ: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
		zhCN: "How much <b>%v</b> did you borrow?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_ASK_BORROWING_COUNTERPARTY: {
		deDE: "Wer hat dir <b>%v</b> geliehen?\n(<i>Sende '.' zum Abbrechen</i>)",
		enUK: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		enUS: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		esES: "¿Quién te ha prestado <b>%v</b>?\n(<i>enviar '.' para cancelar</i>)",
		faIR: "چه کسی به شما <b>%v</b> قرض داده است؟ \n(<i>ارسال '.' برای کنسل کردن</i>)",
		frFR: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		idID: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		itIT: "Chi ti ha prestato <b>%v</b>?\n(<i>Digita '.' ed invia per annullare</i>)",
		jaJP: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		koKR: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		plPL: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ptBR: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ruRU: "У кого вы взяли в долг <b>%v</b>?\n(<i>отправьте '.' чтобы отменить</i>)",
		trTR: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		ukUA: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		uzUZ: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
		zhCN: "Who lended to you <b>%v</b>?\n(<i>send '.' to cancel</i>)",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT: {
		deDE: "Soll eine <a receipt>Quittung</a> an <a counterparty>%v</a> gesendet werden?",
		enUK: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		enUS: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		esES: "¿Debo enviar <a receipt> el recibo</a> a <a counterparty>%v</a>?",
		faIR: "آیا لازم است ماارسال کنیم یک <a receipt>رسید</a> به <a counterparty>%v</a>?",
		frFR: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		idID: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		itIT: "Devo inviare una <a receipt>notifica</a> a <a counterparty>%v</a>?",
		jaJP: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		koKR: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		plPL: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		ptBR: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		ruRU: "Отправить <a receipt>квитанцию</a> для <a counterparty>%v</a>?",
		trTR: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		ukUA: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		uzUZ: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
		zhCN: "Should we send a <a receipt>receipt</a> to <a counterparty>%v</a>?",
	},
	MESSAGE_TEXT_YOU_CAN_SEND_RECEIPT_TO_YOURSELF_BY_SMS: {
		deDE: "Entschuldigung, aber eine Quittung selber per SMS zu schicken ist im Moment noch nicht möglich. Aber dafür geht es mit %v.",
		enUK: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		enUS: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		esES: "Lo siento, el envio del recibo a ti mismo a través de SMS en este momento está desactivado. Pero lo puedes enviar a %v.",
		faIR: "متاسفم، درحال حاضرارسال یک رسید به خودتان بوسیله پیام کوتاه امکان پذیر نیست. شما میتوانید آنرا ارسال کنید به  %v از طریق.",
		frFR: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		idID: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		itIT: "Spiacente ma inviarsi da soli una notifica tramite SMS non e' al momento disponibile. Pero' puoi inviarla a %v.",
		jaJP: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		koKR: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		plPL: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		ptBR: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		ruRU: "К сожалению отправка квитанцию себе по СМС в данный момент отключена. Но вы можете отправить её для %v.",
		trTR: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		ukUA: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		uzUZ: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
		zhCN: "Sorry, sending a receipt to yourself by SMS is not available at the moment. You can send it to %v though.",
	},
	MESSAGE_TEXT_RECEIPT_IS_SENDING_BY_TELEGRAM: {
		deDE: "Die Quittung wird %v per Telegram geschickt...",
		enUK: "We are sending receipt to %v by Telegram...",
		enUS: "We are sending receipt to %v by Telegram...",
		esES: "El recibo está enviando a%v a través de Telegram…",
		faIR: "مادرحال ارسال رسید به %v از طریق تلگرام هستیم...",
		frFR: "We are sending receipt to %v by Telegram...",
		idID: "We are sending receipt to %v by Telegram...",
		itIT: "Sto inviando la notifica a %v tramite Telegram...",
		jaJP: "We are sending receipt to %v by Telegram...",
		koKR: "We are sending receipt to %v by Telegram...",
		plPL: "We are sending receipt to %v by Telegram...",
		ptBR: "We are sending receipt to %v by Telegram...",
		ruRU: "Отправляем для %v извещение через Telegram...",
		trTR: "We are sending receipt to %v by Telegram...",
		ukUA: "We are sending receipt to %v by Telegram...",
		uzUZ: "We are sending receipt to %v by Telegram...",
		zhCN: "We are sending receipt to %v by Telegram...",
	},
	DAY: {
		deDE: "%v day",
		enUK: "%v day",
		enUS: "%v day",
		esES: "%v day",
		faIR: "%v day",
		frFR: "%v day",
		idID: "%v day",
		itIT: "%v day",
		jaJP: "%v day",
		koKR: "%v day",
		plPL: "%v day",
		ptBR: "%v day",
		ruRU: "%v день",
		trTR: "%v day",
		ukUA: "%v day",
		uzUZ: "%v day",
		zhCN: "%v day",
	},
	DAYS_234: {
		deDE: "%v days",
		enUK: "%v days",
		enUS: "%v days",
		esES: "%v days",
		faIR: "%v days",
		frFR: "%v days",
		idID: "%v days",
		itIT: "%v days",
		jaJP: "%v days",
		koKR: "%v days",
		plPL: "%v days",
		ptBR: "%v days",
		ruRU: "%v дня",
		trTR: "%v days",
		ukUA: "%v days",
		uzUZ: "%v days",
		zhCN: "%v days",
	},
	DAYS: {
		deDE: "%v days",
		enUK: "%v days",
		enUS: "%v days",
		esES: "%v days",
		faIR: "%v days",
		frFR: "%v days",
		idID: "%v days",
		itIT: "%v days",
		jaJP: "%v days",
		koKR: "%v days",
		plPL: "%v days",
		ptBR: "%v days",
		ruRU: "%v дней",
		trTR: "%v days",
		ukUA: "%v days",
		uzUZ: "%v days",
		zhCN: "%v days",
	},
	MESSAGE_TEXT_INTEREST_PLEASE_SPECIFY_PERIOD: {
		deDE: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		enUK: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		enUS: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		esES: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		faIR: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		frFR: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		idID: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		itIT: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		jaJP: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		koKR: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		plPL: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		ptBR: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		ruRU: "Пожалуйста укажите также процентный период, т.е. уточните %%v%% это процент за какое количество дней?",
		trTR: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		ukUA: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		uzUZ: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
		zhCN: "Please also specify interest period, e.g. is it %v%% for per year, month, week, some number of days?",
	},
	MESSAGE_TEXT_INTEREST: {
		deDE: "<b>Interest</b>: %v%% per %v",
		enUK: "<b>Interest</b>: %v%% per %v",
		enUS: "<b>Interest</b>: %v%% per %v",
		esES: "<b>Interest</b>: %v%% per %v",
		faIR: "<b>Interest</b>: %v%% per %v",
		frFR: "<b>Interest</b>: %v%% per %v",
		idID: "<b>Interest</b>: %v%% per %v",
		itIT: "<b>Interest</b>: %v%% per %v",
		jaJP: "<b>Interest</b>: %v%% per %v",
		koKR: "<b>Interest</b>: %v%% per %v",
		plPL: "<b>Interest</b>: %v%% per %v",
		ptBR: "<b>Interest</b>: %v%% per %v",
		ruRU: "<b>Ставка</b>: %v%% за %v",
		trTR: "<b>Interest</b>: %v%% per %v",
		ukUA: "<b>Interest</b>: %v%% per %v",
		uzUZ: "<b>Interest</b>: %v%% per %v",
		zhCN: "<b>Interest</b>: %v%% per %v",
	},
	MESSAGE_TEXT_INTEREST_MIN_PERIOD: {
		deDE: "minimum period %v",
		enUK: "minimum period %v",
		enUS: "minimum period %v",
		esES: "minimum period %v",
		faIR: "minimum period %v",
		frFR: "minimum period %v",
		idID: "minimum period %v",
		itIT: "minimum period %v",
		jaJP: "minimum period %v",
		koKR: "minimum period %v",
		plPL: "minimum period %v",
		ptBR: "minimum period %v",
		ruRU: "минимальный период %v",
		trTR: "minimum period %v",
		ukUA: "minimum period %v",
		uzUZ: "minimum period %v",
		zhCN: "minimum period %v",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_FROM_USER: {
		deDE: "{{.Counterparty}} schuldet dir {{.Amount}} .",
		enUK: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		enUS: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		esES: "{{.Counterparty}} prestado por tí {{.Amount}}.",
		faIR: "{{.Counterparty}} از شما {{.Amount}} قرض گرفته است .",
		frFR: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		idID: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		//itIT:   "{{.Counterparty}} ha preso in prestito da te {{.Amount}}.",
		itIT: "{{.Counterparty}} e' in debito di {{.Amount}} con te.",
		jaJP: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		koKR: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		plPL: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		ptBR: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		ruRU: "{{.Counterparty}} взял(а) в долг {{.Amount}}.",
		trTR: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		ukUA: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		uzUZ: "{{.Counterparty}} borrowed from you {{.Amount}}.",
		zhCN: "{{.Counterparty}} borrowed from you {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_NEW_DEBT_TO_USER: {
		deDE: "{{.Counterparty}} hat dir {{.Amount}} geliehen.",
		enUK: "{{.Counterparty}} lended to you {{.Amount}}.",
		enUS: "{{.Counterparty}} lended to you {{.Amount}}.",
		esES: "{{.Counterparty}} prestado a mí {{.Amount}}.",
		faIR: "{{.Counterparty}} به شما {{.Amount}} قرض داده است .",
		frFR: "{{.Counterparty}} lended to you {{.Amount}}.",
		idID: "{{.Counterparty}} lended to you {{.Amount}}.",
		itIT: "{{.Counterparty}} ti ha prestato {{.Amount}}.",
		jaJP: "{{.Counterparty}} lended to you {{.Amount}}.",
		koKR: "{{.Counterparty}} lended to you {{.Amount}}.",
		plPL: "{{.Counterparty}} lended to you {{.Amount}}.",
		ptBR: "{{.Counterparty}} lended to you {{.Amount}}.",
		ruRU: "{{.Counterparty}} дал(а) вам в долг {{.Amount}}.",
		trTR: "{{.Counterparty}} lended to you {{.Amount}}.",
		ukUA: "{{.Counterparty}} lended to you {{.Amount}}.",
		uzUZ: "{{.Counterparty}} lended to you {{.Amount}}.",
		zhCN: "{{.Counterparty}} lended to you {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_FROM_USER: {
		deDE: "Du hast {{.Amount}} an {{.Counterparty}} beglichen.",
		enUK: "You returned {{.Amount}} to {{.Counterparty}}.",
		enUS: "You returned {{.Amount}} to {{.Counterparty}}.",
		esES: "Has devuelto {{.Amount}} a {{.Counterparty}}.",
		faIR: "شما بازگردانده اید {{.Amount}} به {{.Counterparty}}.",
		frFR: "You returned {{.Amount}} to {{.Counterparty}}.",
		idID: "You returned {{.Amount}} to {{.Counterparty}}.",
		itIT: "Hai ridato {{.Amount}} a {{.Counterparty}}.",
		jaJP: "You returned {{.Amount}} to {{.Counterparty}}.",
		koKR: "You returned {{.Amount}} to {{.Counterparty}}.",
		plPL: "You returned {{.Amount}} to {{.Counterparty}}.",
		ptBR: "You returned {{.Amount}} to {{.Counterparty}}.",
		ruRU: "Вы вернули долг - {{.Counterparty}} получил(а) {{.Amount}}.",
		trTR: "You returned {{.Amount}} to {{.Counterparty}}.",
		ukUA: "You returned {{.Amount}} to {{.Counterparty}}.",
		uzUZ: "You returned {{.Amount}} to {{.Counterparty}}.",
		zhCN: "You returned {{.Amount}} to {{.Counterparty}}.",
	},
	MESSAGE_TEXT_RECEIPT_RETURN_TO_USER: {
		deDE: "{{.Counterparty}} hat dir {{.Amount}} beglichen.",
		enUK: "{{.Counterparty}} returned to you {{.Amount}}.",
		enUS: "{{.Counterparty}} returned to you {{.Amount}}.",
		esES: "{{.Counterparty}} te ha devuelto {{.Amount}}.",
		faIR: "{{.Counterparty}} به شما بازپرداخت کرده است {{.Amount}}.",
		frFR: "{{.Counterparty}} returned to you {{.Amount}}.",
		idID: "{{.Counterparty}} returned to you {{.Amount}}.",
		itIT: "{{.Counterparty}} ti ha ridato {{.Amount}}.",
		jaJP: "{{.Counterparty}} returned to you {{.Amount}}.",
		koKR: "{{.Counterparty}} returned to you {{.Amount}}.",
		plPL: "{{.Counterparty}} returned to you {{.Amount}}.",
		ptBR: "{{.Counterparty}} returned to you {{.Amount}}.",
		ruRU: "{{.Counterparty}} вернул вам {{.Amount}}.",
		trTR: "{{.Counterparty}} returned to you {{.Amount}}.",
		ukUA: "{{.Counterparty}} returned to you {{.Amount}}.",
		uzUZ: "{{.Counterparty}} returned to you {{.Amount}}.",
		zhCN: "{{.Counterparty}} returned to you {{.Amount}}.",
	},
	MESSAGE_TEXT_TRANSFER_ALREADY_FULLY_RETURNED: {
		deDE: "Diese Schuld ist bereits vollständig beglichen.",
		enUK: "This debts is already fully returned.",
		esES: "Esta deuda se ha devuelta totalmente.",
		itIT: "Questi debiti sono già completamente restituiti.",
		faIR: "این بدهی ها در حال حاضر به طور کامل بازگشته است.",
		ruRU: "Этот долг уже полностью возвращён.",
	},
	MESSAGE_TEXT_RECEIPT_ALREADY_RETURNED_AMOUNT: {
		deDE: "Bereits beglichen: {{.Amount}}.",
		enUK: "Already returned: {{.Amount}}.",
		esES: "Se ha devuelto ya: {{.Amount}}.",
		faIR: "قبلا برگشت: {{.Amount}}.",
		itIT: "Già restituito: {{.Amount}}.",
		ruRU: "Уже возвращено: {{.Amount}}.",
	},
	MESSAGE_TEXT_RECEIPT_OUTSTANDING_AMOUNT: {
		deDE: "Ausstehend: {{.Amount}}.",
		enUK: "Outstanding: {{.Amount}}.",
		esES: "Falta devolver: {{.Amount}}.",
		faIR: "برجسته: {{.Amount}}",
		itIT: "Inevaso: {{.Amount}}",
		ruRU: "Осталось вернуть: {{.Amount}}.",
	},
	MESSAGE_TEXT_DUE_ON: {
		deDE: "<b>Fällig am</b>: %v",
		enUK: "<b>Return till</b>: %v",
		esES: "<b>Devolver hasta</b>: %v",
		faIR: "<b>بازگردانده شود تا</b>: %v",
		itIT: "<b>Scadenza</b>: %v",
		ruRU: "<b>Вернуть до</b>: %v",
	},
	MESSAGE_TEXT_NOTE: {
		deDE: "Notiz",
		enUK: "Note",
		esES: "Nota",
		faIR: "نکته",
		itIT: "Nota",
		ruRU: "Заметка",
	},
	MESSAGE_TEXT_COMMENT: {
		deDE: "Bemerkung",
		enUK: "Comment",
		esES: "Comentario",
		faIR: "شرح",
		itIT: "Commenti",
		ruRU: "Комментарий",
	},
	MESSAGE_TEXT_LOGIN_TO_WEB_APP: {
		deDE: `<a>Hier klicken</a>, um sich an der Wep-App anzumelden.`,
		enUK: `Click to <a>sign in</a> to web-app.`,
		esES: `Haz click para <a>acceder</a>la web-app.`,
		itIT: "Fai clic per <a>accedi</a> per app web.",
		faIR: `کلیک کنید تا <a>وارد شوید</a> برنامه وب.`,
		ruRU: `Перейдите по <a>ссылке</a> чтобы запустить web-приложение.`,
	},
	MESSAGE_TEXT_DO_YOU_LIKE_OUR_BOT: {
		deDE: "Magst du @{{bot}}?",
		enUK: "Do you like @{{bot}}?",
		esES: "¿Te gusta @{{bot}}?",
		faIR: "آیا می پسندید @{{bot}}?",
		itIT: "Divertito con @{{bot}}?",
		ruRU: "Вам нравится @{{bot}}?",
	},
	COMMAND_TEXT_YES_EXCLAMATION: {
		deDE: "%v Ja!",
		enUK: "%v Yes!",
		esES: "%v ¡Sí!",
		faIR: "بله! %v",
		itIT: "%v Si!",
		ruRU: "%v Да!",
	},
	COMMAND_TEXT_YES: {
		deDE: "%v Ja",
		enUK: "%v Yes",
		esES: "%v Sí",
		itIT: "%v Si",
		faIR: "بله %v",
		ruRU: "%v Да",
	},
	COMMAND_TEXT_NO: {
		deDE: "%v Nein",
		enUK: "%v No",
		esES: "%v No",
		itIT: "%v No",
		faIR: "خیر %v",
		ruRU: "%v Нет",
	},
	COMMAND_TEXT_NOT_TOO_MUCH: {
		deDE: "%v Nicht so sehr",
		enUK: "%v Not too much",
		esES: "%vNo mucho",
		itIT: "%v Non troppo",
		faIR: "نه خیلی زیاد %v",
		ruRU: "%v Не очень",
	},
	COMMAND_TEXT_FEEDBACK: {
		deDE: "/Bewertung",
		enUK: "/Feedback",
		esES: "/Respuesta",
		itIT: "/Risposta",
		faIR: "/بازخورد",
		ruRU: "/Отзыв",
	},
	COMMAND_TEXT_WRITE_FEEDBACK: {
		deDE: "%v Bewertung schreiben",
		enUK: "%v Write feedback",
		esES: "%v Escribir un comentario",
		itIT: "%v Scrivi commenti",
		faIR: "ارسال بازخورد %v",
		ruRU: "%v Написать отзыв",
	},
	MESSAGE_TEXT_THANKS: {
		deDE: "🙏 Danke!",
		enUK: "🙏 Thanks!",
		enUS: "🙏 Thanks!",
		esES: "🙏 ¡Gracias!",
		faIR: "🙏 تشکر!",
		frFR: "🙏 Merci!",
		idID: "🙏 Terima kasih!",
		itIT: "🙏 Grazie!",
		jaJP: "🙏 ありがとう!",
		koKR: "🙏 감사합니다!",
		plPL: "🙏 Dziękuję!",
		ptBR: "🙏 Obrigado!",
		ptPT: "🙏 Obrigado!",
		ruRU: "🙏 Спасибо!",
		trTR: "🙏 Teşekkürler!",
		ukUA: "🙏 Дякую!",
		uzUZ: "🙏 Rahmat!",
		zhCN: "🙏 谢谢!",
	},
	MESSGE_TEXT_DEBT_ERROR_FIXED_START_OVER: {
		deDE: "🙏 Entschuldigung, da gab es einen Fehler. Er wird bald behoben, aber du musst nochmal neu anfangen.",
		enUK: "🙏 Sorry, there was an error. It has been fixed but please re-enter your data for this debt.",
		esES: "🙏 Lo siento, ha salido un error. Lo ha arreglado, pero para esta deuda hay que introducir los datos de nuevo. ",
		ruRU: "🙏 Извините, у нас была ошибка. Она была исправлено, но потребуется внести данные для этого долга заново.",
	},
	MESSAGE_TEXT_PLEASE_SEND_TEXT: {
		deDE: "Bitte senden sie einen Text.",
		enUK: "Please send text.",
		esES: "Por favor, envia el texto.",
		faIR: "لطفاً متن ارسال کنید.",
		itIT: "Si prega di inviare il testo.",
		ruRU: "Пожалуйста отправьте текст.",
	},
	MESSAGE_TEXT_CAN_YOU_RATE_AT_STOREBOT: {

		deDE: `🤖 Kannst du mich im Store Bot hoch bewerten und eine gute Bewertung schreiben?
		Es wird dich weniger als eine Minute kosten! 😇`,
		enUK: `🤖 Can you rate it high and write a good review in bots catalog Store Bot?
		It will take less than a minute of your time! 😇`,

		esES: `🤖 Puedes valolarlo con una buena nota y una buena opinión en el catálogo Store Bot?
		Te costará menos de un minuto! 😇`,

		faIR: `🤖  آیا می توانید در کاتالوگ روباتها در استور بوت امتیاز بالایی داده و اظهار نظر خوبی در مورد این روبات ثبت کنید؟
		این کار کمتر از یک دقیقه از وقت شما را می گیرد! 😇`,

		itIT: `🤖 Puoi votarlo in alto e scrivere una buona revisione nel catalogo Bot Store?
		Ci vorrà meno di un minuto del tuo tempo! 😇`,

		ruRU: `🤖 Можете поставить ему высокую оценку и хороший отзыв в каталоге ботов Store Bot?
		Это займет меньше минуты вашего времени! 😇`,
	},
	MESSAGE_TEXT_ASK_TO_WRITE_FEEDBACK_WITHIN_MESSENGER: {
		deDE: "Schreibe (auf Englisch oder Russisch) uns, was man am Bot besser machen kann:",
		enUK: "Share your thoughts (in English or Russian) about what could be done to make the bot better:",
		esES: "Comparte tus pensamientos (en Inglés o Ruso) sobre qué podemos hacer para que el bot sea mejor:",
		faIR: "نظرات خود را (به انگلیسی و روسی ) در مورد اینکه چه کاری می توان انجام داد تا این ربات بهتر شود، با ما به اشتراک بگذارید:",
		itIT: "Condividi i tuoi pensieri (in Inglese o Russo) su come sarebbe migliore secondo te il bot:",
		ruRU: "Поделитесь вашими мыслями (на русском или английском) о том, что нужно сделать, чтобы бот стал лучше:",
	},
	MESSAGE_TEXT_HOW_TO_RATE_AT_STOREBOT: {

		deDE: `<b>Wie man bewertet - in 3 einfachen Schritten:</b>

	1. Klick auf diesen Link, um eine Bewertung abzugeben:
	https://t.me/storebot?start={{bot}}

	2. Wähle "⭐️⭐️⭐️⭐️⭐️"

	3. Schreib etwas Nettes auf Englisch oder wähle "Skip this step"

	Wirklich vielen Dank! Dank deiner Bewertung werden vielleicht mehr Leute auf diesen Bot aufmerksam. Das ist gut für die Motivation der Entwickler dieses Bots! 😎`,

		enUK: `<b>How to rate in 3 simple steps:</b>

	1. Click on this link to rate and review:
	https://t.me/storebot?start={{bot}}

	2. Click on the "⭐️⭐️⭐️⭐️⭐️" button

	3. Write your message or press "Skip this step" button

	Thank you very much! As a result of your actions, even more people will learn about the bot.All this will serve as the additional motivation for the developers! 😎`,

		esES: `<b>Como valorar en 3 simples pasos:</b>

	1. Click este link para valorar y dejar tu opinión:
	https://t.me/storebot?start={{bot}}

	2. Click en "⭐️⭐️⭐️⭐️⭐️" botón

	3. Escribe tu mensage o apreta "Skip this step" botón

	¡Muchas gracias! Merced a tus acciones más gente conocerá a bot. Todo eso sirve para una motivación adicional a los creadores! 😎`,

		itIT: `<b>Come valutare in 3 semplici passaggi:</b>
	1. Clicca su questo link per votare e lasciare una recensione:
	https://t.me/storebot?start={{bot}}

	2. Clicca sul "⭐️⭐️⭐️⭐️⭐️" bottone

	3. Scrivi il tuo messaggio o premi "Salta questo step"

	Grazie infinitamente! Come risultato delle tue azioni, altre persone guarderanno il bot.Dando anche un motivo in più per continuare ai developers! 😎`,

		faIR: `<b>چگونگی امتیازدهی در سه گام ساده :</b>

	1. برای امتیازدهی و ثبت نظرات بر روی لینگ زیر کلیک کنید
	https://t.me/storebot?start={{bot}}

	2. بر روی دکمه "⭐️⭐️⭐️⭐️⭐️" کلیک کنید

	3. پیام خودرا ثبت کنید یا روی دکمه "پرش از این مرحله" کلیک کنید

	بسیار سپاسگزاریم! عمل شما باعث می شود افراد بیشتری در مورد bot.All بیاموزند. این امر انگیزه مضاعفی به توسعه دهندگان این ربات می دهد ! 😎`,

		ruRU: `<b>Как поставить оценку в три простых шага:</b>

	1. Перейдите по ссылке, чтобы оставить оценку и отзыв:
	https://t.me/storebot?start={{bot}}

	2. Нажмите на кнопку "⭐️⭐️⭐️⭐️⭐️"

	3. Напишите сообщение или нажмите кнопку "Пропустить этот шаг"

	Спасибо вам большое! Благодаря этому о боте узнает больше людей — это служит дополнительной мотивацией для разработчиков! 😎`,
	},
	MESSAGE_TEXT_ASK_FOR_FEEDBACK: {
		deDE: "Über ein kleines Feedback wie der Bot so ist, würden wir uns freuen. Es dauert nur ein paar Sekunden.",
		enUK: "We would appreciate if tell us how we doing. It takes just few seconds.",
		enUS: "We would appreciate if tell us how we doing. It takes just few seconds.",
		esES: "Te agredecemos si valoras el funccionamiento de nuestro applicación. Te costará solo unos segundos.",
		faIR: "سپاسگزار خواهیم بود اگر به ما بگویید کارمان چطور بوده است. این تنها چند ثانیه زمان میبرد.",
		frFR: "Nous apprécierions si vous nous disiez comment nous allons. Cela ne prend que quelques secondes.",
		idID: "Kami akan menghargai jika memberi tahu kami bagaimana kami melakukannya. Ini hanya membutuhkan beberapa detik.",
		itIT: "Ci farebbe piacere se lasciassi un voto per il nostro lavoro. Ti bastano solo alcuni secondi.",
		jaJP: "私たちの仕事ぶりを教えていただければ幸いです。ほんの数秒で済みます。",
		koKR: "우리가 어떻게 하고 있는지 알려주시면 감사하겠습니다. 몇 초 밖에 걸리지 않습니다.",
		plPL: "Bylibyśmy wdzięczni, gdybyś powiedział nam, jak sobie radzimy. To zajmie tylko kilka sekund.",
		ptBR: "Agradecemos se nos disser como estamos indo. Leva apenas alguns segundos.",
		ruRU: "Будем признетельны если вы оцените работу нашего приложения. Это займёт всего несколько секунд.",
		trTR: "Nasıl yaptığımızı bize söylerseniz memnun oluruz. Sadece birkaç saniye sürer.",
		ukUA: "Ми будемо вдячні, якщо ви розкажете нам, як ми працюємо. Це займе лише кілька секунд.",
		uzUZ: "Qanday qilayotganimizni bizga aytsangiz, minnatdor bo'lardik. Bu atigi bir necha soniya vaqt oladi.",
		zhCN: "如果您告诉我们我们的表现如何，我们将不胜感激。这只需要几秒钟。",
	},
	COMMAND_TEXT_GIVE_FEEDBACK: {
		deDE: "Bewerte diesen Bot",
		enUK: "Rate this bot",
		enUS: "Rate this bot",
		esES: "Valora a bot",
		faIR: "به این ربات امتیاز بدهید",
		frFR: "Évaluer ce bot",
		idID: "Nilai bot ini",
		itIT: "Vota questo bot",
		jaJP: "このボットを評価する",
		koKR: "이 봇 평가하기",
		plPL: "Oceń tego bota",
		ptBR: "Avalie este bot",
		ruRU: "Оценить приложение",
		trTR: "Bu botu değerlendir",
		ukUA: "Оцінити цього бота",
		uzUZ: "Bu botni baholang",
		zhCN: "评价这个机器人",
	},
	COMMAND_TEXT_OPEN_STOREBOT_FOR_FEEDBACK: {
		deDE: "Mache eine Bewertung auf @Storebot",
		enUK: "Leave rating at @Storebot",
		enUS: "Leave rating at @Storebot",
		esES: "Valorar en @Storebot",
		faIR: "امتیاز خود را اینجا وارد کنید @Storebot",
		frFR: "Laisser une évaluation sur @Storebot",
		idID: "Tinggalkan peringkat di @Storebot",
		itIT: "Lascia un voto a @Storebot",
		jaJP: "@Storebotで評価を残す",
		koKR: "@Storebot에서 평가 남기기",
		plPL: "Zostaw ocenę na @Storebot",
		ptBR: "Deixe sua avaliação no @Storebot",
		ruRU: "Оценить на  @Storebot",
		trTR: "@Storebot'ta değerlendirme bırakın",
		ukUA: "Залиште оцінку на @Storebot",
		uzUZ: "@Storebot'da baho qoldiring",
		zhCN: "在 @Storebot 上留下评分",
	},
	MESSAGE_TEXT_ON_REFUSED_TO_RATE: {
		deDE: `Okay, vielleicht werden wir wann anders bewertet.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ansonsten freuen wir uns immer zu hören, was man besser machen kann.
	`,
		enUK: `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		enUS: `OK, maybe you can rate us another time.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you suggest any improvements.
	`,
		esES: `OK, Quizás puedas valorar en otro momento.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	También te agredeceríamos si ofrecieras alguna mejora.
	`,
		faIR: `بسیار خوب، ممکن است شما بتوانید زمان دیگری به ما امتیاز بدهید.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	همچنین سپاسگزار خواهیم بود اگر شما هرگونه امکان بهبود را با ما در میان بگذارید.
	`,
		frFR: `OK, peut-être que vous pourrez nous évaluer une autre fois.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous suggérez des améliorations.
	`,
		idID: `OK, mungkin Anda dapat menilai kami lain kali.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda menyarankan perbaikan apa pun.
	`,
		itIT: `OK, forse ci puoi valutare un'altra volta.

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
		plPL: `OK, może ocenisz nas innym razem.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli zaproponujesz jakieś ulepszenia.
	`,
		ptBR: `OK, talvez você possa nos avaliar em outra ocasião.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você sugerir melhorias.
	`,
		ruRU: `ОК, возможно вы смоежете поставить оценку в другой раз.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы предложите любые улучшения.
	`,
		trTR: `Tamam, belki başka bir zaman bizi değerlendirebilirsiniz.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Herhangi bir iyileştirme önerirseniz de memnun oluruz.
	`,
		ukUA: `Гаразд, можливо, ви зможете оцінити нас іншим разом.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви запропонуєте будь-які покращення.
	`,
		uzUZ: `OK, balki boshqa safar bizni baholashingiz mumkin.

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz biron bir yaxshilanishni taklif qilsangiz, biz ham minnatdor bo'lamiz.
	`,
		zhCN: `好的，也许您可以在另一个时间给我们评分。

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您提出任何改进建议，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_POSITIVE: {
		deDE: `Danke, wir arbeiten hart dran!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Wir freuen uns auch immer über <a suggest-idea>neue Ideen</a>.
	`,
		enUK: `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		enUS: `Thanks, we worked hard!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	We also will appreciate if you <a suggest-idea>suggest improvements</a>.
	`,
		esES: `Gracias, hemos trabajado duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Te agradeceríamos si <a suggest-idea>propusieras alguna mejora</a>.
	`,
		faIR: `ممنونیم، ما سخت کارکرده ایم!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	ما همچنین سپاسگزار خواهیم بود اگر شما<a suggest-idea> هرگونه امکان بهبود را با ما در میان بگذارید </a>را.
	`,
		frFR: `Merci, nous avons travaillé dur !

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Nous apprécierons également si vous <a suggest-idea>suggérez des améliorations</a>.
	`,
		idID: `Terima kasih, kami bekerja keras!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Kami juga akan menghargai jika Anda <a suggest-idea>menyarankan perbaikan</a>.
	`,
		itIT: `GRAZIE MILLE, abbiamo lavoro duro!

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
		plPL: `Dziękujemy, ciężko pracowaliśmy!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Będziemy również wdzięczni, jeśli <a suggest-idea>zaproponujesz ulepszenia</a>.
	`,
		ptBR: `Obrigado, trabalhamos duro!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Também agradeceremos se você <a suggest-idea>sugerir melhorias</a>.
	`,
		ruRU: `Спасибо, мы очень старались!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Так же будем признательны если вы <a suggest-idea>предложите улучшения</a>.
	`,
		trTR: `Teşekkürler, çok çalıştık!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	<a suggest-idea>İyileştirmeler önerirseniz</a> de memnun oluruz.
	`,
		ukUA: `Дякуємо, ми старанно працювали!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Ми також будемо вдячні, якщо ви <a suggest-idea>запропонуєте покращення</a>.
	`,
		uzUZ: `Rahmat, biz qattiq ishladik!

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	Agar siz <a suggest-idea>yaxshilanishlarni taklif qilsangiz</a>, biz ham minnatdor bo'lamiz.
	`,
		zhCN: `谢谢，我们努力工作！

	{{MESSAGE_TEXT_YOU_CAN_HELP_BY}}

	如果您<a suggest-idea>提出改进建议</a>，我们也将不胜感激。
	`,
	},
	MESSAGE_TEXT_YOU_CAN_HELP_BY: {
		deDE: `
Es gibt viele Wege uns zu helfen:

* Gib uns 5⭐ im <a storebot>Verzeichnis der Bots</a>.

* Erzähl am besten all deinen Freunde davon.
Du könntest es auf <a share-fb>Facebook</a> posten oder auf <a share-twitter>Twitter</a> twittern.

* Ansonsten auch gerne eine kleine Spende - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		enUK: `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		enUS: `
You can help us a lot if you:

* Give us 5⭐ at <a storebot>directory of bots</a>.

* Tell about the app to your friends.
For example at <a share-fb>Facebook</a> or <a share-twitter>Twitter</a>.

* Support further development - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>about $2.2</i>)
`,
		esES: `
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
		frFR: `
Vous pouvez nous aider beaucoup si vous:

* Nous donnez 5⭐ au <a storebot>répertoire des bots</a>.

* Parlez de l'application à vos amis.
Par exemple sur <a share-fb>Facebook</a> ou <a share-twitter>Twitter</a>.

* Soutenez le développement ultérieur - <a href = "https://goo.gl/Qhh0yL">€2 via PayPal</a> (<i>environ $2.2</i>)
`,
		idID: `
Anda dapat membantu kami banyak jika Anda:

* Beri kami 5⭐ di <a storebot>direktori bot</a>.

* Ceritakan tentang aplikasi kepada teman-teman Anda.
Misalnya di <a share-fb>Facebook</a> atau <a share-twitter>Twitter</a>.

* Dukung pengembangan lebih lanjut - <a href = "https://goo.gl/Qhh0yL">€2 melalui PayPal</a> (<i>sekitar $2.2</i>)
`,
		itIT: `
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
		plPL: `
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
		ruRU: `
	Вы нам очень поможете если:

	* Поставите нам 5⭐ в <a storebot>каталоге ботов</a>.

	* Расскажите о боте своим друзьям.
	Например в <a share-vk>ВКонтакте</a>, <a share-fb>Facebook</a> или <a share-twitter>Twitter</a>.

* Поддержите дальнейшую разработку - <a href ="https://goo.gl/Qhh0yL">€2 через PayPal</a>
`,
		trTR: `
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
		uzUZ: `
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
		deDE: `Du bist quitt mit %v`,
		enUK: `Balance is empty for %v`,
		enUS: `Balance is empty for %v`,
		esES: `El balance es cero para %v`,
		faIR: `تراز خالی است برای %v`,
		frFR: `Le solde est vide pour %v`,
		idID: `Saldo kosong untuk %v`,
		itIT: `Non hai alcun credito o debito con %v`,
		jaJP: `%vの残高は空です`,
		koKR: `%v에 대한 잔액이 비어 있습니다`,
		plPL: `Saldo jest puste dla %v`,
		ptBR: `Saldo está vazio para %v`,
		ruRU: `Нулевой баланс для %v`,
		trTR: `%v için bakiye boş`,
		ukUA: `Баланс порожній для %v`,
		uzUZ: `%v uchun balans bo'sh`,
		zhCN: `%v的余额为空`,
	},
	MESSAGE_TEXT_ASK_TO_TRANSLATE: {
		deDE: `Möchtest du den Bot in einer anderen Sprache? Du kannst beim <a>Übersetzen helfen</a>.`,
		enUK: `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		enUS: `Do you want our bot to talk in other language? You can <a>help with translation</a>.`,
		esES: `¿Te gustaría que nuestro bot hablara en otro idioma? Puedes <a>ayudar con traducción</a>.`,
		faIR: `آیا می خواهید ربات ما به زبان دیگری صحبت کند؟ شما می توانید <a>با ترجمه به ما کمک کنید</a>.`,
		frFR: `Voulez-vous que notre bot parle dans une autre langue? Vous pouvez <a>aider à la traduction</a>.`,
		idID: `Apakah Anda ingin bot kami berbicara dalam bahasa lain? Anda dapat <a>membantu dengan terjemahan</a>.`,
		itIT: `Vuoi che il nostro bot parli altre lingue? Aiuta con la <a>traduzione</a>.`,
		jaJP: `私たちのボットが他の言語で話すことを望みますか？あなたは<a>翻訳を手伝う</a>ことができます。`,
		koKR: `우리 봇이 다른 언어로 말하기를 원하십니까? <a>번역을 도울 수 있습니다</a>.`,
		plPL: `Czy chcesz, aby nasz bot mówił w innym języku? Możesz <a>pomóc w tłumaczeniu</a>.`,
		ptBR: `Você quer que nosso bot fale em outro idioma? Você pode <a>ajudar com a tradução</a>.`,
		ruRU: `Хотите чтобы наш бот разговаривал на другом языке? Вы можете <a>помочь с переводом</a>.`,
		trTR: `Botumuzun başka bir dilde konuşmasını ister misiniz? <a>Çeviriye yardımcı olabilirsiniz</a>.`,
		ukUA: `Хочете, щоб наш бот розмовляв іншою мовою? Ви можете <a>допомогти з перекладом</a>.`,
		uzUZ: `Botimiz boshqa tilda gaplashishini xohlaysizmi? Siz <a>tarjima qilishda yordam</a> berishingiz mumkin.`,
		zhCN: `您想要我们的机器人用其他语言交谈吗？您可以<a>帮助翻译</a>。`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEUTRAL: {
		deDE: `Gut, wir geben uns Mühe. Deine Rückmeldung wird an die Entwickler weitergeleitet.

Vielleicht willst du <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		enUK: `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		enUS: `Well, we worked hard. You feedback will be passed to developers.

Maybe you can <a submit-bug>report your issue</a> or <a suggest-idea>suggest how we can improve</a>?`,
		esES: `Bueno, hemos trabajado duro. Tu opinión se pasará a los creadores.

Quizás puedas <a submit-bug>informarnos de algún problema</a> o <a suggest-idea>proponernos qué podríamos mejorar</a>?`,
		faIR: `خب، ما سخت کارکردیم. بازخورد شما به توسعه دهندگان برنامه منعکس می شود.

شما می توانید <a submit-bug>مشکل خود را گزارش دهید</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		frFR: `Eh bien, nous avons travaillé dur. Vos commentaires seront transmis aux développeurs.

Peut-être pouvez-vous <a submit-bug>signaler votre problème</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		idID: `Yah, kami bekerja keras. Umpan balik Anda akan diteruskan ke pengembang.

Mungkin Anda dapat <a submit-bug>melaporkan masalah Anda</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		itIT: `Bene, il nostro lavoro non e' andato in vano. Il tuo feedback sara' inoltrato agli sviluppatori.

Per caso vuoi anche <a submit-bug>segnalare un problema</a> oppure <a suggest-idea>suggerire come possiamo migliorare</a>?`,
		jaJP: `まあ、私たちは一生懸命働きました。あなたのフィードバックは開発者に伝えられます。

<a submit-bug>問題を報告</a>したり、<a suggest-idea>改善方法を提案</a>したりすることもできますか？`,
		koKR: `음, 우리는 열심히 일했습니다. 귀하의 피드백은 개발자에게 전달됩니다.

<a submit-bug>문제를 보고</a>하거나 <a suggest-idea>개선 방법을 제안</a>할 수 있습니까?`,
		plPL: `Cóż, ciężko pracowaliśmy. Twoja opinia zostanie przekazana deweloperom.

Może możesz <a submit-bug>zgłosić swój problem</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		ptBR: `Bem, trabalhamos duro. Seu feedback será passado para os desenvolvedores.

Talvez você possa <a submit-bug>relatar seu problema</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		ruRU: `Чтож, мы очень старались. Ваша оценка будет передана разработчикам.

Может быть вы <a submit-bug>сообщите нам что не работает</a> или подскажите <a suggest-idea>как можно улучшить</a>?`,
		trTR: `Eh, çok çalıştık. Geri bildiriminiz geliştiricilere iletilecektir.

Belki <a submit-bug>sorununuzu bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		ukUA: `Ну, ми старанно працювали. Ваш відгук буде переданий розробникам.

Можливо, ви можете <a submit-bug>повідомити про свою проблему</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		uzUZ: `Xo'sh, biz qattiq ishladik. Sizning fikr-mulohazalaringiz ishlab chiqaruvchilarga yetkaziladi.

Balki siz <a submit-bug>muammoingizni xabar qilishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		zhCN: `嗯，我们努力工作。您的反馈将传递给开发人员。

也许您可以<a submit-bug>报告您的问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	MESSAGE_TEXT_ON_FEEDBACK_NEGATIVE: {
		deDE: `Das tut uns sehr leid. Vielleicht willst du uns <a submit-bug>einen Fehler melden</a> oder <a suggest-idea>eine Verbesserung vorschlagen</a>?`,
		enUK: `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		enUS: `We are very sorry. Maybe you can <a submit-bug>let us know what is wrong</a> or <a suggest-idea>suggest how we can improve</a>?`,
		esES: `Lo sentimos mucho. Igual podrías <a submit-bug>decirnos qué no funcciona bien</a> o <a suggest-idea>proponernos cómo podemos mejorarlo</a>?`,
		faIR: `ما بسیار متاسفیم. شما می توانید <a submit-bug>به ما بگویید مشکلتان چیست</a> یا <a suggest-idea>پیشنهاد دهید چطور میتوانیم بهبود ایجاد کنیم</a>?`,
		frFR: `Nous sommes vraiment désolés. Peut-être pouvez-vous <a submit-bug>nous faire savoir ce qui ne va pas</a> ou <a suggest-idea>suggérer comment nous pouvons nous améliorer</a>?`,
		idID: `Kami sangat menyesal. Mungkin Anda dapat <a submit-bug>memberi tahu kami apa yang salah</a> atau <a suggest-idea>menyarankan bagaimana kami dapat meningkatkan</a>?`,
		itIT: `Ci dispiace molto. Potresti farci sapere<a submit-bug>cosa non ti e' piaciuto</a> oppure <a suggest-idea>suggerirci come possiamo migliorare</a>?`,
		jaJP: `大変申し訳ありません。<a submit-bug>何が問題かを教えていただく</a>か、<a suggest-idea>改善方法を提案していただく</a>ことはできますか？`,
		koKR: `매우 죄송합니다. <a submit-bug>무엇이 잘못되었는지 알려주시거나</a> <a suggest-idea>개선 방법을 제안</a>해 주실 수 있습니까?`,
		plPL: `Bardzo nam przykro. Może możesz <a submit-bug>dać nam znać, co jest nie tak</a> lub <a suggest-idea>zasugerować, jak możemy się poprawić</a>?`,
		ptBR: `Lamentamos muito. Talvez você possa <a submit-bug>nos informar o que está errado</a> ou <a suggest-idea>sugerir como podemos melhorar</a>?`,
		ruRU: `Нам очень стыдно. Может быть вы <a submit-bug>подскажите что не так</a> или <a suggest-idea>предложите усовершенствования</a>?`,
		trTR: `Çok üzgünüz. Belki <a submit-bug>neyin yanlış olduğunu bize bildirebilir</a> veya <a suggest-idea>nasıl geliştirebileceğimizi önerebilirsiniz</a>?`,
		ukUA: `Нам дуже шкода. Можливо, ви можете <a submit-bug>повідомити нам, що не так</a> або <a suggest-idea>запропонувати, як ми можемо покращити</a>?`,
		uzUZ: `Juda uzr so'raymiz. Balki siz <a submit-bug>nima noto'g'ri ekanligini bizga aytib berishingiz</a> yoki <a suggest-idea>qanday yaxshilashimiz mumkinligini taklif qilishingiz</a> mumkin?`,
		zhCN: `我们非常抱歉。也许您可以<a submit-bug>告诉我们哪里出了问题</a>或<a suggest-idea>建议我们如何改进</a>？`,
	},
	COMMAND_TEXT_ASK_FOR_FEEDBACK: {
		deDE: "Bitte bewerte unsere App",
		enUK: "Please rate our app",
		enUS: "Please rate our app",
		esES: "Por favor valora nuestro app",
		faIR: "لطفاً به برنامه ما امتیاز دهید",
		frFR: "Veuillez évaluer notre application",
		idID: "Silakan nilai aplikasi kami",
		itIT: "Per favore vota il nostro bot",
		jaJP: "アプリを評価してください",
		koKR: "앱을 평가해 주세요",
		plPL: "Oceń naszą aplikację",
		ptBR: "Por favor, avalie nosso aplicativo",
		ruRU: "Оцените наше приложение?",
		trTR: "Lütfen uygulamamızı değerlendirin",
		ukUA: "Будь ласка, оцініть наш додаток",
		uzUZ: "Iltimos, ilovamizni baholang",
		zhCN: "请评价我们的应用",
	},
	COMMAND_TEXT_FEEDBACK_POSITIVE: {
		deDE: "Ja, es ist eine tolle App!",
		enUK: "Yes, it's a great app!",
		enUS: "Yes, it's a great app!",
		esES: "¡Sí, es una app fantástica!",
		faIR: "بله، این برنامه عالی است",
		frFR: "Oui, c'est une excellente application !",
		idID: "Ya, ini aplikasi yang bagus!",
		itIT: "Si, e' un app fantastica!",
		jaJP: "はい、素晴らしいアプリです！",
		koKR: "네, 훌륭한 앱입니다!",
		plPL: "Tak, to świetna aplikacja!",
		ptBR: "Sim, é um ótimo aplicativo!",
		ruRU: "Да, отличное приложение!",
		trTR: "Evet, harika bir uygulama!",
		ukUA: "Так, це чудовий додаток!",
		uzUZ: "Ha, bu ajoyib ilova!",
		zhCN: "是的，这是一个很棒的应用！",
	},
	COMMAND_TEXT_FEEDBACK_NEUTRAL: {
		deDE: "Nicht schlecht, aber könnte besser sein",
		enUK: "Not bad but can be better.",
		enUS: "Not bad but can be better.",
		esES: "No está mal, pero podría ser mejor.",
		faIR: "بد نیست ولی می تواند بهتر باشد.",
		frFR: "Pas mal mais peut être amélioré.",
		idID: "Tidak buruk tapi bisa lebih baik.",
		itIT: "Non male ma potrebbe esser migliore.",
		jaJP: "悪くないですが、もっと良くなる可能性があります。",
		koKR: "나쁘지 않지만 더 좋아질 수 있습니다.",
		plPL: "Niezły, ale może być lepszy.",
		ptBR: "Não é ruim, mas pode ser melhor.",
		ruRU: "Неплохо, но можно лучше.",
		trTR: "Fena değil ama daha iyi olabilir.",
		ukUA: "Непогано, але може бути краще.",
		uzUZ: "Yomon emas, lekin yaxshiroq bo'lishi mumkin.",
		zhCN: "不错，但可以更好。",
	},
	COMMAND_TEXT_FEEDBACK_NEGATIVE: {
		deDE: "Ich mag sie nicht",
		enUK: "Don't like it",
		enUS: "Don't like it",
		esES: "No me gusta",
		faIR: "از این برنامه را نمی پسندم",
		frFR: "Je n'aime pas",
		idID: "Tidak suka",
		itIT: "Non mi piace",
		jaJP: "好きではない",
		koKR: "마음에 들지 않음",
		plPL: "Nie lubię tego",
		ptBR: "Não gosto",
		ruRU: "Не нравится",
		trTR: "Beğenmedim",
		ukUA: "Не подобається",
		uzUZ: "Yoqmadi",
		zhCN: "不喜欢",
	},
	COMMAND_TEXT_FEEDBACK_NOT_READY: {
		deDE: "Ich bin mir noch unsicher",
		enUK: "Not decided yet",
		enUS: "Not decided yet",
		esES: "Estoy aún indeciso",
		faIR: "هنوز تصمیم نگرفته ام.",
		frFR: "Pas encore décidé",
		idID: "Belum memutuskan",
		itIT: "Sono indeciso",
		jaJP: "まだ決めていません",
		koKR: "아직 결정하지 않았습니다",
		plPL: "Jeszcze nie zdecydowałem",
		ptBR: "Ainda não decidi",
		ruRU: "Пока не понятно",
		trTR: "Henüz karar vermedim",
		ukUA: "Ще не вирішив",
		uzUZ: "Hali qaror qabul qilinmagan",
		zhCN: "尚未决定",
	},
	MESSAGE_TEXT_SETTINGS: {
		deDE: "Was willst du ändern?",
		enUK: "What do you want to adjust?",
		enUS: "What do you want to adjust?",
		esES: "¿Qué te gustaría modificar?",
		faIR: "می خواهید چه چیزی را تنظیم کنید؟",
		frFR: "Que voulez-vous ajuster ?",
		idID: "Apa yang ingin Anda sesuaikan?",
		itIT: "Cosa vuoi modificare?",
		jaJP: "何を調整したいですか？",
		koKR: "무엇을 조정하고 싶으신가요?",
		plPL: "Co chcesz dostosować?",
		ptBR: "O que você deseja ajustar?",
		ruRU: "Что будем настраивать?",
		trTR: "Neyi ayarlamak istiyorsunuz?",
		ukUA: "Що ви хочете налаштувати?",
		uzUZ: "Nimani sozlamoqchisiz?",
		zhCN: "您想调整什么？",
	},
	MT_NO_OUTSTANDING_TRANSFERS: {
		deDE: `Sie versuchen, einen Rückgabedatensatz zu erstellen, aber es gibt keine ausstehenden Schulden.

Wenn Sie glauben, dass dies ein Fehler ist, teilen Sie uns dies bitte in @DebtsTrackerGroup mit.`,
		enUK: `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		enUS: `You are trying to create return record but there are no outstanding debts.

If you believe this is a mistale please let us know in @DebtsTrackerGroup.`,
		esES: `Estás intentando crear un registro de devolución pero no hay deudas pendientes.

Si crees que esto es un error, háganoslo saber en @DebtsTrackerGroup.`,
		faIR: `شما در حال تلاش برای ایجاد سابقه بازگشت هستید اما هیچ بدهی معوقه ای وجود ندارد.

اگر فکر می کنید این یک اشتباه است، لطفاً به ما در @DebtsTrackerGroup اطلاع دهید.`,
		frFR: `Vous essayez de créer un enregistrement de retour mais il n'y a pas de dettes en cours.

Si vous pensez qu'il s'agit d'une erreur, veuillez nous en informer dans @DebtsTrackerGroup.`,
		idID: `Anda mencoba membuat catatan pengembalian tetapi tidak ada hutang yang belum dibayar.

Jika Anda yakin ini adalah kesalahan, beri tahu kami di @DebtsTrackerGroup.`,
		itIT: `Stai cercando di creare un record di restituzione ma non ci sono debiti in sospeso.

Se ritieni che si tratti di un errore, ti preghiamo di farcelo sapere in @DebtsTrackerGroup.`,
		jaJP: `返却記録を作成しようとしていますが、未払いの借金はありません。

これが間違いだと思われる場合は、@DebtsTrackerGroupでお知らせください。`,
		koKR: `반환 기록을 생성하려고 하지만 미결제 부채가 없습니다.

이것이 실수라고 생각되면 @DebtsTrackerGroup에서 알려주십시오.`,
		plPL: `Próbujesz utworzyć rekord zwrotu, ale nie ma zaległych długów.

Jeśli uważasz, że to pomyłka, daj nam znać w @DebtsTrackerGroup.`,
		ptBR: `Você está tentando criar um registro de devolução, mas não há dívidas pendentes.

Se você acredita que isso é um erro, por favor, nos avise em @DebtsTrackerGroup.`,
		ruRU: `Вы пытаетесь создать запись о возврате долга, но мы не нашли не закрытых долгов.

Если вы считаете что это ошибка пожалуйста дайте нам знать в @DebtsTrackerGroup.`,
		trTR: `İade kaydı oluşturmaya çalışıyorsunuz ancak bekleyen borç bulunmamaktadır.

Bunun bir hata olduğunu düşünüyorsanız, lütfen @DebtsTrackerGroup'ta bize bildirin.`,
		ukUA: `Ви намагаєтеся створити запис про повернення, але немає непогашених боргів.

Якщо ви вважаєте, що це помилка, будь ласка, повідомте нам у @DebtsTrackerGroup.`,
		uzUZ: `Siz qaytarish yozuvini yaratmoqchi bo'lyapsiz, lekin to'lanmagan qarzlar yo'q.

Agar bu xato deb o'ylasangiz, iltimos, bizga @DebtsTrackerGroup'da xabar bering.`,
		zhCN: `您正在尝试创建退款记录，但没有未偿还的债务。

如果您认为这是一个错误，请在 @DebtsTrackerGroup 中告诉我们。`,
	},
	MT_ATTEMPT_TO_CREATE_DEBT_WITH_INTEREST_AFFECTING_OUTSTANDING: {
		deDE: "Sie versuchen, eine Schuld mit Zinsen zu erstellen, die sich auf ausstehende Überweisungen auswirken wird. Bitte schließen Sie diese zuerst.",
		enUK: "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		enUS: "You are trying to create a debt with interest that will affect outstanding transfers. Please close them first.",
		esES: "Estás intentando crear una deuda con intereses que afectará a las transferencias pendientes. Por favor, ciérralas primero.",
		faIR: "شما در حال تلاش برای ایجاد بدهی با بهره هستید که بر انتقال های معوق تأثیر می گذارد. لطفا ابتدا آنها را ببندید.",
		frFR: "Vous essayez de créer une dette avec intérêt qui affectera les transferts en cours. Veuillez les fermer d'abord.",
		idID: "Anda mencoba membuat hutang dengan bunga yang akan memengaruhi transfer yang belum diselesaikan. Harap tutup terlebih dahulu.",
		itIT: "Stai cercando di creare un debito con interessi che influenzerà i trasferimenti in sospeso. Per favore, chiudili prima.",
		jaJP: "未決済の送金に影響する利息付きの借金を作成しようとしています。まずそれらを閉じてください。",
		koKR: "미결제 송금에 영향을 미치는 이자가 있는 부채를 만들려고 합니다. 먼저 그것들을 닫으십시오.",
		plPL: "Próbujesz utworzyć dług z odsetkami, który wpłynie na nierozliczone przelewy. Proszę najpierw je zamknąć.",
		ptBR: "Você está tentando criar uma dívida com juros que afetará transferências pendentes. Por favor, feche-as primeiro.",
		ruRU: "Вы пытаетесь создать запись о долге с процентами которая затронет незакрытые долги. Пожалуйста закройте их сначала.",
		trTR: "Bekleyen transferleri etkileyecek faizli bir borç oluşturmaya çalışıyorsunuz. Lütfen önce onları kapatın.",
		ukUA: "Ви намагаєтеся створити борг з відсотками, який вплине на невиконані перекази. Будь ласка, спочатку закрийте їх.",
		uzUZ: "Siz to'lanmagan o'tkazmalarga ta'sir qiladigan foizli qarz yaratmoqchisiz. Iltimos, avval ularni yoping.",
		zhCN: "您正在尝试创建一个带有利息的债务，这将影响未结清的转账。请先关闭它们。",
	},
	MESSAGE_TEXT_NOT_IMPLEMENTED_YET: {
		deDE: "Entschuldigung, diese Funktion ist noch nicht fertig programmiert.",
		enUK: "Sorry, this functionality is not implemented yet.",
		enUS: "Sorry, this functionality is not implemented yet.",
		esES: "Lo sentimos, esta función no está activa aún.",
		faIR: "متاسفم، این عملکرد هنوز پیاده سازی نشده است.",
		frFR: "Désolé, cette fonctionnalité n'est pas encore implémentée.",
		idID: "Maaf, fungsi ini belum diimplementasikan.",
		itIT: "Spiacenti ma questa funzionalita' non e' ancora attiva.",
		jaJP: "申し訳ありませんが、この機能はまだ実装されていません。",
		koKR: "죄송합니다. 이 기능은 아직 구현되지 않았습니다.",
		plPL: "Przepraszamy, ta funkcjonalność nie jest jeszcze zaimplementowana.",
		ptBR: "Desculpe, esta funcionalidade ainda não foi implementada.",
		ruRU: "Извините, данный функционал ещё не реализован.",
		trTR: "Üzgünüz, bu işlevsellik henüz uygulanmadı.",
		ukUA: "Вибачте, ця функціональність ще не реалізована.",
		uzUZ: "Kechirasiz, bu funksionallik hali amalga oshirilmagan.",
		zhCN: "抱歉，此功能尚未实现。",
	},
	MESSAGE_TEXT_ASK_INVITE_CHANNEL: {
		deDE: "Worüber möchtest du eingeladen werden?",
		enUK: "How do you want to get an invite?",
		enUS: "How do you want to get an invite?",
		esES: "¿Comó prefieres recibir la invitación?",
		faIR: "می خواهید چگونه دعوت شوید؟",
		frFR: "Comment voulez-vous recevoir une invitation ?",
		idID: "Bagaimana Anda ingin mendapatkan undangan?",
		itIT: "Come vuoi ricevere l'invito?",
		jaJP: "招待状をどのように受け取りますか？",
		koKR: "초대를 어떻게 받고 싶으신가요?",
		plPL: "Jak chcesz otrzymać zaproszenie?",
		ptBR: "Como você deseja receber um convite?",
		ruRU: "Как вы хотите получить код приглашения?",
		trTR: "Daveti nasıl almak istersiniz?",
		ukUA: "Як ви хочете отримати запрошення?",
		uzUZ: "Taklifnomani qanday olishni xohlaysiz?",
		zhCN: "您想如何获得邀请？",
	},
	MESSAGE_TEXT_PLEASE_ENTER_INVITE_CODE: {
		deDE: "Bitte gib den Bestätigungs-Code ein:",
		enUK: "Please enter an invite code:",
		enUS: "Please enter an invite code:",
		esES: "Introduce el código de la invitación",
		faIR: "لطفاً یک کد دعوت وارد کنید:",
		frFR: "Veuillez entrer un code d'invitation :",
		idID: "Silakan masukkan kode undangan:",
		itIT: "Inserisci un codice invito:",
		jaJP: "招待コードを入力してください：",
		koKR: "초대 코드를 입력하세요:",
		plPL: "Proszę wprowadzić kod zaproszenia:",
		ptBR: "Por favor, insira um código de convite:",
		ruRU: "Пожалуйста введите код приглашения:",
		trTR: "Lütfen bir davet kodu girin:",
		ukUA: "Будь ласка, введіть код запрошення:",
		uzUZ: "Iltimos, taklif kodini kiriting:",
		zhCN: "请输入邀请码：",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_RECEIVED: {
		deDE: "Wir haben eine Nachricht an %v gesendet.\n\nBitte öffne die Nachricht und klick auf den Link, um deine Mail-Adresse zu bestätigen.",
		enUK: "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		enUS: "We've sent a message to %v.\n\nPlease open the email and click a link to confirm your email address.",
		esES: "Hemos enviado un mensage a %v.\n\nPor favor, abre tu e-mail y haz click en el link para confirmar tu e-mail.",
		faIR: "ما یک پیام ارسال کردیم به %v.\n\nلطفاً ایمیل خود را باز کرده و روی لینک کلیک کنید تا آدرس ایمیل شما تایید شود.",
		frFR: "Nous avons envoyé un message à %v.\n\nVeuillez ouvrir l'e-mail et cliquer sur un lien pour confirmer votre adresse e-mail.",
		idID: "Kami telah mengirim pesan ke %v.\n\nSilakan buka email dan klik tautan untuk mengonfirmasi alamat email Anda.",
		itIT: "Abbiamo inviato un messaggio a %v.\n\nPer favore apri l'email e clicca sul link per confermare il tuo indirizzo email",
		jaJP: "%vにメッセージを送信しました。\n\nメールを開き、リンクをクリックしてメールアドレスを確認してください。",
		koKR: "%v에 메시지를 보냈습니다.\n\n이메일을 열고 링크를 클릭하여 이메일 주소를 확인하십시오.",
		plPL: "Wysłaliśmy wiadomość do %v.\n\nProszę otworzyć e-mail i kliknąć link, aby potwierdzić swój adres e-mail.",
		ptBR: "Enviamos uma mensagem para %v.\n\nPor favor, abra o e-mail e clique em um link para confirmar seu endereço de e-mail.",
		ruRU: "Мы отправили письмо на %v.\n\nПожалуйста откройте его и кликните на ссылку для подтверждения адреса.",
		trTR: "%v adresine bir mesaj gönderdik.\n\nLütfen e-postayı açın ve e-posta adresinizi onaylamak için bir bağlantıya tıklayın.",
		ukUA: "Ми надіслали повідомлення на %v.\n\nБудь ласка, відкрийте електронний лист і натисніть посилання, щоб підтвердити свою електронну адресу.",
		uzUZ: "Biz %v ga xabar yubordik.\n\nIltimos, elektron pochtani oching va elektron pochta manzilingizni tasdiqlash uchun havolani bosing.",
		zhCN: "我们已向 %v 发送了一条消息。\n\n请打开电子邮件并点击链接确认您的电子邮件地址。",
	},
	MESSAGE_TEXT_USER_EMAIL_FOR_INVITE_SENT_TELEGRAM: {
		deDE: "Wenn Telegram öffnet, drücke auf <b>Start</b>.",
		enUK: "Once Telegram pop ups click the <b>Start</b> button.",
		enUS: "Once Telegram pop ups click the <b>Start</b> button.",
		esES: "Después de abrir Telegram aprieta el <b>Start</b> botón.",
		faIR: "وقتی تلگرام اجرا شد برروی دکمه  <b>شروع</b> کلیک کنید.",
		frFR: "Une fois que Telegram apparaît, cliquez sur le bouton <b>Start</b>.",
		idID: "Setelah Telegram muncul, klik tombol <b>Start</b>.",
		itIT: "Una volta aperto il bot su telegram clicca su <b>Start</b>.",
		jaJP: "Telegramがポップアップしたら、<b>Start</b>ボタンをクリックしてください。",
		koKR: "Telegram이 팝업되면 <b>Start</b> 버튼을 클릭하세요.",
		plPL: "Gdy pojawi się Telegram, kliknij przycisk <b>Start</b>.",
		ptBR: "Quando o Telegram aparecer, clique no botão <b>Start</b>.",
		ruRU: "После того как откроется Telegram нажмите кнопку <b>Start</b>.",
		trTR: "Telegram açıldığında <b>Start</b> düğmesine tıklayın.",
		ukUA: "Коли з'явиться Telegram, натисніть кнопку <b>Start</b>.",
		uzUZ: "Telegram ochilganda, <b>Start</b> tugmasini bosing.",
		zhCN: "Telegram 弹出后，点击 <b>Start</b> 按钮。",
	},
	MESSAGE_TEXT_USER_CONTACT_FOR_INVITE_RECEIVED: {
		deDE: "Danke, du bist in der Warteschlange für eine Einladung.\n\nEs dauert etwa zwei bis drei Tage.\n\nAber du könntest den Code noch heute bekommen, wenn du einen Link auf Facebook teilst.",
		enUK: "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		enUS: "Thanks, you've been queued for an invite.\n\nCurrent awaiting time is 2-3 days.\n\nYou can get an invite code today by sharing a link to the bot on Facebook.",
		esES: "Gracias, ya estás inscrito en la cola para conseguir la invitación.\n\nTiempo de espera 2-3 días.\n\nPuedes conseguirlo hoy si compartes el link de nuestro bot en Facebook.",
		faIR: "سپاسگزاریم، شما در نوبت دعوت قرار گرفتید\n\nزمان انتظار شما در حال حاضر 2-3 روز می باشد.\n\n شما می توانید با به اشتراک گذاری لینک روبات در فیسبوک امروز یک کد دعوت دریافت کنید. ",
		frFR: "Merci, vous avez été mis en file d'attente pour une invitation.\n\nLe temps d'attente actuel est de 2 à 3 jours.\n\nVous pouvez obtenir un code d'invitation aujourd'hui en partageant un lien vers le bot sur Facebook.",
		idID: "Terima kasih, Anda telah antri untuk undangan.\n\nWaktu menunggu saat ini adalah 2-3 hari.\n\nAnda dapat memperoleh kode undangan hari ini dengan membagikan tautan ke bot di Facebook.",
		itIT: "Grazie, ora sei in coda per un codice invito.\n\nTempo di attesa medio 2-3 giorni.\n\nPuoi ottenere un codice invito subito condividendo il link al bot su Facebook.",
		jaJP: "ありがとうございます。招待のために列に並んでいます。\n\n現在の待ち時間は2〜3日です。\n\nFacebookでボットへのリンクを共有することで、今日招待コードを取得できます。",
		koKR: "감사합니다. 초대를 위해 대기열에 올랐습니다.\n\n현재 대기 시간은 2-3일입니다.\n\nFacebook에서 봇에 대한 링크를 공유하여 오늘 초대 코드를 받을 수 있습니다.",
		plPL: "Dziękujemy, zostałeś umieszczony w kolejce do zaproszenia.\n\nObecny czas oczekiwania to 2-3 dni.\n\nMożesz otrzymać kod zaproszenia już dziś, udostępniając link do bota na Facebooku.",
		ptBR: "Obrigado, você foi colocado na fila para um convite.\n\nO tempo de espera atual é de 2 a 3 dias.\n\nVocê pode obter um código de convite hoje compartilhando um link para o bot no Facebook.",
		ruRU: "Спасибо, вы записаны в очередь на получение приглашения.\n\nТекущее время ожидания 2-3 дня.\n\nВы можете получить приглашение сегодня если расскажите о нашем боте на Facebook.",
		trTR: "Teşekkürler, bir davetiye için sıraya alındınız.\n\nMevcut bekleme süresi 2-3 gündür.\n\nBotun bağlantısını Facebook'ta paylaşarak bugün bir davet kodu alabilirsiniz.",
		ukUA: "Дякуємо, вас поставлено в чергу на запрошення.\n\nПоточний час очікування - 2-3 дні.\n\nВи можете отримати код запрошення сьогодні, поділившись посиланням на бота у Facebook.",
		uzUZ: "Rahmat, siz taklif uchun navbatga qo'yildingiz.\n\nHozirgi kutish vaqti 2-3 kun.\n\nFacebookda botga havola ulashish orqali bugun taklif kodini olishingiz mumkin.",
		zhCN: "谢谢，您已排队等候邀请。\n\n当前等待时间为2-3天。\n\n您可以通过在Facebook上分享机器人链接，今天就获得邀请码。",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_EMAIL: {
		deDE: "Bitte gib deine e-Mail Adresse an:",
		enUK: "Please provide your email address",
		enUS: "Please provide your email address",
		esES: "Por favor, esctibe tu e-mail",
		faIR: "لطفاً آدرس ایمیل خود را وارد کنید.",
		frFR: "Veuillez fournir votre adresse e-mail",
		idID: "Silakan berikan alamat email Anda",
		itIT: "Inserisci il tuo indirizzo email:",
		jaJP: "メールアドレスを入力してください",
		koKR: "이메일 주소를 입력해 주세요",
		plPL: "Proszę podać swój adres e-mail",
		ptBR: "Por favor, forneça seu endereço de e-mail",
		ruRU: "Пожалуйста напишите ваш email адрес:",
		trTR: "Lütfen e-posta adresinizi girin",
		ukUA: "Будь ласка, вкажіть вашу електронну адресу",
		uzUZ: "Iltimos, elektron pochta manzilingizni kiriting",
		zhCN: "请提供您的电子邮件地址",
	},
	MESSAGE_TEXT_PLEASE_PROVIDE_YOUR_PHONE_NUMBER: {
		deDE: "Bitte gib deine Telefonnummer an:",
		enUK: "Please provide your phone number",
		enUS: "Please provide your phone number",
		esES: "Por favor, esctibe tu número de teléfono",
		faIR: "لطفاً شماره تلفن خود را وارد نمایید.",
		frFR: "Veuillez fournir votre numéro de téléphone",
		idID: "Silakan berikan nomor telepon Anda",
		itIT: "Inserisci il tuo numero di telefono:",
		jaJP: "電話番号を入力してください",
		koKR: "전화번호를 입력해 주세요",
		plPL: "Proszę podać swój numer telefonu",
		ptBR: "Por favor, forneça seu número de telefone",
		ruRU: "Пожалуйста напишите номер вашего телефона:",
		trTR: "Lütfen telefon numaranızı girin",
		ukUA: "Будь ласка, вкажіть ваш номер телефону",
		uzUZ: "Iltimos, telefon raqamingizni kiriting",
		zhCN: "请提供您的电话号码",
	},
	MESSAGE_TEXT_WRONG_INVITE_CODE: {
		deDE: "Ungültiger Code: %v",
		enUK: "Wrong invite code: %v",
		enUS: "Wrong invite code: %v",
		esES: "El código de invitación no es correcto: %v",
		faIR: "کد دعوت اشتباه است %v",
		frFR: "Code d'invitation incorrect : %v",
		idID: "Kode undangan salah: %v",
		itIT: "Codice invito: %v errato",
		jaJP: "招待コードが間違っています：%v",
		koKR: "초대 코드가 잘못되었습니다: %v",
		plPL: "Nieprawidłowy kod zaproszenia: %v",
		ptBR: "Código de convite errado: %v",
		ruRU: "Неправильный код приглашения: %v",
		trTR: "Yanlış davet kodu: %v",
		ukUA: "Неправильний код запрошення: %v",
		uzUZ: "Noto'g'ri taklif kodi: %v",
		zhCN: "邀请码错误：%v",
	},
	MESSAGE_TEXT_WRONG_EMAIL: {
		deDE: "Ungültige e-Mail Adresse.",
		enUK: "Wrong email address.",
		enUS: "Wrong email address.",
		esES: "El e-mail no es correcto.",
		faIR: "آدرس ایمیل اشتباه است.",
		frFR: "Adresse e-mail incorrecte.",
		idID: "Alamat email salah.",
		itIT: "L'email inserita e' sbagliata.",
		jaJP: "メールアドレスが間違っています。",
		koKR: "이메일 주소가 잘못되었습니다.",
		plPL: "Nieprawidłowy adres e-mail.",
		ptBR: "Endereço de e-mail incorreto.",
		ruRU: "Неправильный email адрес.",
		trTR: "Yanlış e-posta adresi.",
		ukUA: "Неправильна електронна адреса.",
		uzUZ: "Noto'g'ri elektron pochta manzili.",
		zhCN: "电子邮件地址错误。",
	},
	MESSAGE_TEXT_WRONG_PHONE_NUMBER: {
		deDE: "Ungültige Telefonnummer.",
		enUK: "Wrong phone number.",
		enUS: "Wrong phone number.",
		esES: "El número de telefono no es correcto.",
		faIR: "شماره تلفن اشتباه است",
		frFR: "Numéro de téléphone incorrect.",
		idID: "Nomor telepon salah.",
		itIT: "Il numero inserito e' sbagliato.",
		jaJP: "電話番号が間違っています。",
		koKR: "전화번호가 잘못되었습니다.",
		plPL: "Nieprawidłowy numer telefonu.",
		ptBR: "Número de telefone incorreto.",
		ruRU: "Неправильный номер телефона.",
		trTR: "Yanlış telefon numarası.",
		ukUA: "Неправильний номер телефону.",
		uzUZ: "Noto'g'ri telefon raqami.",
		zhCN: "电话号码错误。",
	},
	MESSAGE_TEXT_OK_PLEASE_TRY_AGAIN: {
		deDE: "Ok, bitte versuche es erneut.",
		enUK: "Ok, please try again.",
		enUS: "Ok, please try again.",
		esES: "Ok, inténtalo de nuevo.",
		faIR: "بسیار خوب، لطفا مجدداً سعی کنید.",
		frFR: "Ok, veuillez réessayer.",
		idID: "Ok, silakan coba lagi.",
		itIT: "Ok, prova di nuovo.",
		jaJP: "はい、もう一度お試しください。",
		koKR: "네, 다시 시도해 주세요.",
		plPL: "Ok, spróbuj ponownie.",
		ptBR: "Ok, tente novamente.",
		ptPT: "Ok, tente novamente.",
		ruRU: "Хорошо, попробуйте ещё раз.",
		trTR: "Tamam, lütfen tekrar deneyin.",
		ukUA: "Добре, спробуйте ще раз.",
		uzUZ: "Mayli, qaytadan urinib ko'ring.",
		zhCN: "好的，请再试一次。",
	},

	COMMAND_TEXT_MISTYPE_WILL_TRY_AGAIN: {
		deDE: "Ich habe mich vertippt.",
		enUK: "I've mistyped, will try again.",
		esES: "Me he equivocado, lo intentaré otra vez",
		faIR: "اشتباه تایپ کردم، دوباره امتحان می کنم.",
		itIT: "Ho sbagliato, riprovo.",
		ruRU: "Я опечатался, попробую ещё раз.",
	},
	COMMAND_TEXT_TELL_ME_MORE_ABOUT_INVITES: {
		deDE: "Erzähl mir mehr über diese Codes",
		enUK: "Tell me more about the codes",
		esES: "Dime más de los códigos",
		faIR: "در مورد کدها بیشتر برای من توضیح دهید.",
		itIT: "Ulteriori informazioni riguardo il codice invito.",
		ruRU: "Расскажите ка мне об этих кодах",
	},
	COMMAND_TEXT_INVITE_ME_BY_EMAIL: {
		deDE: "Sende mir einen Code per Mail",
		enUK: "Send me invite code by email",
		esES: "Envíame el código de invitación por e-mail",
		faIR: "کد دعوت را برای من از طریق ایمیل ارسال کنید.",
		itIT: "Inviami il codice invito tramite email",
		ruRU: "Хочу код приглашения на email",
	},
	COMMAND_TEXT_INVITE_ME_BY_SMS: {
		deDE: "Sende mir einen Code per SMS",
		enUK: "Send me invite code by SMS",
		esES: "Envíame el código de invitación a través de SMS",
		faIR: "کد دعوت را برای من از طریق پیام کوتاه ارسال کنید.",
		itIT: "Inviami il codice invito tramite SMS",
		ruRU: "Хочу код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_EMAIL: {
		deDE: "Sende mir nochmal einen Code per Mail",
		enUK: "Send me new invite code by email",
		esES: "Envíame un nuevo código de invitación por e-mail",
		faIR: "یک کد دعوت جدیداز طریق ایمیل برای من  ارسال کنید.",
		itIT: "Inviami il nuovo codice invito tramite email",
		ruRU: "Новый код приглашения на email",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_SMS: {
		deDE: "Sende mir nochmal einen Code per SMS",
		enUK: "Send me new invite code by SMS",
		esES: "Envíame un nuevo código de invitación a través de SMS",
		faIR: "یک کد دعوت جدید از طریق پیام کوتاه برای من ارسال کنید.",
		itIT: "Inviami il nuovo codice invito tramite SMS",
		ruRU: "Новый код приглашения по SMS",
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE_BY_TELEGRAM: {
		deDE: "Sende mir nochmal einen Code per Telegram",
		enUK: "Send me new invite by Telegram",
		esES: "Envíame un nuevo código de invitación a través de Telegram",
		faIR: "یک کد دعوت جدید از طریق تلگرام برای من ارسال کنید.",
		itIT: "Inviami il nuovo codice invito tramite Telegram",
		ruRU: "Получить приграшение в Telegram",
	},
	MESSAGE_TEXT_UNKNOWN_LANGUAGE: {
		deDE: "Unbekannte Sprache. Bitte wähle eine aus der Liste:",
		enUK: "Unknown language. Please choose 1 from the options:",
		esES: "Idioma desconocido. Por favor elige 1 de las opciones:",
		faIR: "زبان ناشناخته. لطفاً یکی از گزینه ها را انتخاب کنید:",
		itIT: "Lingua socnosciuta. Per favore scegline una tra le opzioni:",
		ruRU: "Незнакомый язык. Пожалуйста выберете один из предоставленных:",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY: {
		deDE: "Unbekannte Gegnerpartei. Bitte wähle aus der Liste oder wähle <b>neuer Kontakt</b>.",
		enUK: "Unknown counterparty. Please choose <b>Add new</b> if it's a new contact.",
		esES: "Contacto desconocido. Por favor elige <b>Añadir</b> si es un contacto nuevo.",
		faIR: "مخاطب تراکنش شناسایی نشد. <b>یک مورد جدید اضافه کنید</b> اگر این ایشان یک مخاطب جدید هستند.",
		itIT: "Destinatario sconosciuto. Scegli <b>Aggiugni nuovo</b> se e' un nuovo contatto.",
		ruRU: "Неизвестный контакт. Пожалуйста выберите <b>Добавить</b> если это новый контакт.",
	},
	MESSAGE_TEXT_UNKNOWN_COUNTERPARTY_ON_RETURN: {
		deDE: "Unbekannte Gegnerpartei. Bitte wähle aus der Liste.",
		enUK: "Unknown counterparty. Please choose from the list.",
		esES: "Contacto desconocido. Por favor elige de la lista.",
		faIR: "مخاطب تراکنش شناسایی نشد. لطفاً از فهرست انتخاب کنید.",
		itIT: "Destinatario sconosciuto. Scegli dalla lista qui sotto.",
		ruRU: "Неизвестный контакт. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_UNKNOWN_DEBT: {
		deDE: "Unbekannter Schuldschein. Bitte wähle aus der Liste.",
		enUK: "Unknown debt. Please choose from the list.",
		esES: "Deuda desconocida. Por favor elige de la lista.",
		faIR: "بدهی ناشناخته است. لطفا از فهرست انتخاب کنید.",
		itIT: "Debito sconosciuto. Scegli dalla lista qui sotto.",
		ruRU: "Неизвестный долг. Пожалуйста выберите из списка.",
	},
	MESSAGE_TEXT_BILL_CARD_HEADER: {
		deDE: `<b>Rechnung</b>: <code>%v</code> — %v`,
		enUK: `<b>Bill</b>: <code>%v</code> — %v`,
		esES: `<b>Cuenta</b>: <code>%v</code> — %v`,
		faIR: "<b>Bill</b>: <code>%v</code> — %v",
		itIT: "<b>Bill</b>: <code>%v</code> — %v",
		ruRU: `<b>Cчёт</b>: <code>%v</code> — %v`,
	},
	MESSAGE_TEXT_MEMBERS_TITLE: {
		deDE: "Mitglieder",
		enUK: "Members",
		enUS: "Members",
		esES: "Miembros",
		faIR: "اعضا",
		frFR: "Membres",
		idID: "Anggota",
		itIT: "Membri",
		jaJP: "メンバー",
		koKR: "멤버",
		plPL: "Członkowie",
		ptBR: "Membros",
		ruRU: "Участники",
		trTR: "Üyeler",
		ukUA: "Учасники",
		uzUZ: "A'zolar",
		zhCN: "成员",
	},
	ALERT_TEXT_NOTHING_CHANGED: {
		deDE: "Nichts geändert",
		enUK: "Nothing changed",
		esES: "Nada ha cambiado",
		faIR: "چیزی تغییر نکرده است",
		itIT: "Niente è cambiato",
		ruRU: "Ничего не изменилось",
	},
	ALERT_TEXT_YOU_ARE_ALREADY_MEMBER_OF_THE_GROUP: {
		deDE: "Du bist schon Mitglied beim Teilen dieser Rechnung.",
		enUK: "You are already a member of this bill splitting group.",
		esES: "Ya es miembro de este grupo de división de facturas.",
		faIR: "شما قبلا عضو این گروه تقسیم لایحه هستید.",
		itIT: "Sei già membro di questo gruppo di divisione fatture.",
		ruRU: "Вы уже участник этой группы по совместной оплате счетов.",
	},
	MESSAGE_TEXT_YOUR_BILL_SPLITTING_GROUPS: {
		deDE: "Gruppen, mit denen du Rechnungen teilst",
		enUK: "Your bills splitting groups",
		enUS: "Your bills splitting groups",
		esES: "Ya es miembro de este grupo de división de facturas.",
		faIR: "شما قبلا عضو این گروه تقسیم لایحه هستید.",
		frFR: "Vos groupes de partage de factures",
		idID: "Grup pembagian tagihan Anda",
		itIT: "Sei già membro di questo gruppo di divisione fatture.",
		jaJP: "あなたの請求書分割グループ",
		koKR: "귀하의 청구서 분할 그룹",
		plPL: "Twoje grupy dzielenia rachunków",
		ptBR: "Seus grupos de divisão de contas",
		ptPT: "Seus grupos de divisão de contas",
		ruRU: "Ваши группы совметсной оплаты",
		trTR: "Fatura bölme gruplarınız",
		ukUA: "Ваші групи спільної оплати",
		uzUZ: "Hisob-kitoblarni bo'lish guruhlari",
		zhCN: "您的账单分割组",
	},
	MESSAGE_TEXT_USE_ARROWS_TO_SELECT_GROUP: {
		enUK: "Use ⬅️ & ➡️ to select group",
		esES: "Usa ⬅️ y ➡️ para seleccionar el grupo",
		faIR: "برای انتخاب گروه از ⬅️ & ❢️ استفاده کنید",
		itIT: "Usare ⬅️ & ➡️ per selezionare il gruppo",
		ruRU: "Используйте ⬅️ и ➡️ чтобы выбрать группу.",
	},
	MESSAGE_TEXT_NO_GROUPS: {
		deDE: "Du gehörst zu keiner Gruppe, die sich Rechnungen teilt.",
		enUK: "You are not a participant of any bill splitting group.",
		esES: "Usted no es participante de ningún grupo de división de facturas.",
		faIR: "شما شرکت کننده در هر گروه تقسیم لایحه نیستید.",
		itIT: "Non sei un partecipante a qualsiasi gruppo di divisione fatture.",
		ruRU: "Вы не состоите в группах совместной оплаты.",
	},
	MESSAGE_TEXT_USER_JOINED_GROUP: {
		deDE: `Hi %v, du bist der Gruppe, die sich Rechnungen teilt, beigetreten.`,
		enUK: `Hi %v, you joined this bill splitting group.`,
		faIR: "سلام %v، شما به گروه تقسیم این لایحه پیوستید",
		itIT: "Hi %v, sei entrato in questo gruppo di divisione fatture.",
		ruRU: `Привет %v, вы присоеденились к этой группе по совместной оплате счетов.
		`,
	},
	MESSAGE_TEXT_MEMBERS_CARD_TITLE: {
		deDE: "<b>Wer sich die Rechnung teilt</b> (%d)",
		enUK: "<b>Bills splitting members</b> (%d)",
		faIR: "(%d) <b>نقض تقسیم اعضا</b>",
		itIT: "<b>Membri di divisione delle bollette</b> (%d)",
		ruRU: "<b>Участники совместных оплат</b> (%d)",
	},
	MESSAGE_TEXT_SPLIT_MODE: {
		deDE: "<b>Teilen</b>: %v",
		enUK: "<b>Split</b>: %v",
		esES: "<b>División</b>: %v",
		faIR: "<b>شکاف</b>: %v",
		itIT: "<b>Diviso</b>: %v",
		ruRU: "<b>Делить</b>: %v",
	},
	MESSAGE_TEXT_ASK_HOW_TO_SPLIT_IN_GROP: {
		deDE: "In welchem Verhältnis teilt ihr in dieser Gruppe eure Rechnungen?",
		enUK: "In what proportion do you split bills in this group?",
		esES: "¿En qué proporción divide las facturas en este grupo?",
		faIR: "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟",
		itIT: "In quale percentuale dividi le fatture in questo gruppo?",
		ruRU: "В какой пропорции вы делите счета в этой группе?",
	},
	MESSAGE_TEXT_MEMBERS_CARD_FOOTER: {
		deDE: "Klick <code>Join</code>, um auch Rechnungen zu teilen.",
		enUK: "Click <code>Join</code> to participate in bills splitting.",
		esES: "¿En qué proporción divide las facturas en este grupo?",
		faIR: "در این سهم، آیا شما در این گروه حساب ها را تقسیم می کنید؟",
		itIT: "In quale percentuale dividi le fatture in questo gruppo?",
		ruRU: "Жмите <code>Присоедениться</code> чтобы учавствовать.",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBER_TITLE: {
		enUK: "{{.N}}. {{.MemberName}}",
		esES: "{{.N}}. {{.MemberName}}",
		faIR: "{{.N}}. {{.MemberName}}",
		itIT: "{{.N}}. {{.MemberName}}",
		ruRU: "{{.N}}. {{.MemberName}}",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW: {
		deDE: "<i>{{.Percent}}%</i>",
		enUK: "<i>{{.Percent}}%</i>",
		esES: "<i>{{.Percent}}%</i>",
		faIR: "<i>{{.Percent}}%</i>",
		itIT: "<i>{{.Percent}}%</i>",
		ruRU: "<i>{{.Percent}}%</i>",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_OWES: {
		deDE: "\n   <i>schuldet {{.Owes}}</i>",
		enUK: "\n   <i>owes {{.Owes}}</i>",
		esES: "\n   <i>debo {{.Owes}}</i>",
		ruRU: "\n   <i>должен {{.Owes}}</i>",
		faIR: "\n   <i>بدهکار است {{.Owes}}</i>",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PAID: {
		deDE: "\n   <i>bezahlte {{.Paid}}</i>",
		enUK: "\n   <i>paid {{.Paid}}</i>",
		esES: "\n   <i>he pagado {{.Paid}}</i>",
		ruRU: "\n   <i>заплатил {{.Paid}}</i>",
		faIR: "\n   <i>پرداخت شده {{.Paid}}</i>",
	},
	MESSAGE_TEXT_BILL_CARD_MEMBERS_ROW_PART_PAID: {
		deDE: "\n<i>bezahlte {{.Paid}}, schuldet noch {{.Owes}}</i>",
		enUK: "\n<i>paid {{.Paid}}, owes {{.Owes}}</i>",
		esES: "\n<i>he pagado {{.Paid}}, debo {{.Owes}}</i>",
		ruRU: "\n<i>заплатил {{.Paid}}, должен {{.Owes}}</i>",
		faIR: "\n<i>پرداخت شده {{.Paid}}, بدهکار است {{.Owes}}</i>",
	},
	MESSAGE_TEXT_BILL_ASK_WHO_PAID: {
		deDE: "Bitte wähle, wer die Rechnung gezahlt hat:",
		enUK: "Please choose who paid for the bill:",
		esES: "Por favor, elige quien ha pagado la cuenta:",
		faIR: "لطفا انتخاب کنید که چه کسانی برای این لایحه پرداخت کرده اند:",
		itIT: "Scegliere chi ha pagato la fattura:",
		ruRU: "Пожалуйста выберите кто заплатил по счёту:",
	},
	MESSAGE_TEXT_STATUS: {
		deDE: "Status: %v",
		enUK: "Status: %v",
		esES: "Estado: %v",
		faIR: "وضعیت: %v",
		itIT: "Stato: %v",
		ruRU: "Статус: %v",
	},
	BUTTON_TEXT_ADD_MEMBER: {
		deDE: "Partei hinzufügen",
		enUK: "Add participant",
		esES: "Añadir participante",
		faIR: "افزودن مشارکت کننده",
		itIT: "Aggiungi partecipante",
		ruRU: "Добавить участника",
	},
	BUTTON_TEXT_FINALIZE_BILL: {
		deDE: "🔓 Rechnung abschließen",
		enUK: "🔓 Lock the bill",
		esES: "🔓 Cerrar la cuenta",
		faIR: "🔓 لایحه را قفل کنید",
		itIT: "🔓 Bloccare il conto",
		ruRU: "🔓 Закрыть счёт",
	},
	BUTTON_TEXT_EDIT_BILL: {
		deDE: "✏️ Bearbeiten",
		enUK: "✏️ Edit",
		esES: "✏️ Editar",
		faIR: "✏️ ویرایش",
		itIT: "✏️ Modifica",
		ruRU: "✏️ Изменить",
	},
	BUTTON_TEXT_SPLIT_MODE: {
		deDE: "➗ Teilen: %v",
		enUK: "➗ Split: %v",
		esES: "➗ Dividir: %v",
		faIR: "➗ تقسیم: %v",
		itIT: "➗ Split: %v",
		ruRU: "➗ Делить: %v",
	},
	MESSAGE_TEXT_SPLIT_LABEL_WITH_VALUE: {
		deDE: "Teilen: %v",
		enUK: "Split: %v",
		esES: "Dividir: %v",
		faIR: "تقسیم: %v",
		itIT: "Split: %v",
		ruRU: "Делить: %v",
	},
	STATUS_DRAFT: {
		deDE: "Entwurf",
		enUK: "draft",
		esES: "borrador",
		faIR: "پیش نویس",
		itIT: "bozza",
		ruRU: "черновик",
	},
	SPLIT_MODE_EQUALLY: {
		deDE: "Gleichverteilt",
		enUK: "Equally",
		esES: "A partes iguales",
		ruRU: "Поровну",
		faIR: "به همان اندازه",
		itIT: "Ugualmente",
	},
	SPLIT_MODE_PERCENTAGE: {
		deDE: "Prozentual",
		enUK: "Percentage",
		esES: "Porcentaje",
		faIR: "درصد",
		itIT: "Percentuale",
		ruRU: "В процентах",
	},
	SPLIT_MODE_EXACT_AMOUNT: {
		deDE: "Exakte Summen",
		enUK: "Exact amounts",
		esES: "Importes exactos",
		faIR: "مقادیر دقیق",
		itIT: "Quantità esatte",
		ruRU: "Точные суммы",
	},
	SPLIT_MODE_SHARES: {
		deDE: "Teilen",
		enUK: "Shares",
		esES: "Proporciones",
		faIR: "سهام",
		itIT: "Azioni",
		ruRU: "В долях",
	},
	BUTTON_TEXT_JOIN: {
		deDE: "➕ Beitreten",
		enUK: "➕ Join",
		esES: "➕ Adherirse",
		faIR: "➕ عضویت",
		itIT: "➕ Join",
		ruRU: "➕ Присоедениться",
	},
	BUTTON_TEXT_LEAVE: {
		deDE: "Verlassen",
		enUK: "Leave",
		esES: "Salir",
		faIR: "ترک کردن",
		itIT: "Partire",
		ruRU: "Покинуть",
	},
	BUTTON_TEXT_DUE: {
		deDE: "📅 Fällig: %v",
		enUK: "📅 Due: %v",
		enUS: "📅 Due: %v",
		esES: "📅 Hasta: %v",
		faIR: "📅 مورد: %v",
		frFR: "📅 Échéance: %v",
		idID: "📅 Jatuh tempo: %v",
		itIT: "📅 Dovuto: %v",
		jaJP: "📅 期限: %v",
		koKR: "📅 기한: %v",
		plPL: "📅 Termin: %v",
		ptBR: "📅 Vencimento: %v",
		ptPT: "📅 Vencimento: %v",
		ruRU: "📅 До: %v",
		trTR: "📅 Vade: %v",
		ukUA: "📅 До: %v",
		uzUZ: "📅 Muddati: %v",
		zhCN: "📅 到期: %v",
	},
	NOT_SET: {
		deDE: "nicht gesetzt",
		enUK: "not set",
		esES: "no establecido",
		faIR: "تنظیم نشده",
		itIT: "non impostato",
		ruRU: "не задано",
	},
	BUTTON_TEXT_MANAGE_MEMBERS: {
		deDE: "👫 Parteien",
		enUK: "👫 Participants",
		esES: "👫 Participantes",
		faIR: "👫 شرکت کنندگان",
		itIT: "👫 Partecipanti",
		ruRU: "👫 Участники",
	},
	BUTTON_TEXT_CHANGE_BILL_PAYER: {
		deDE: "🔀 Bezahler ändern",
		enUK: "🔀 Change payer",
		esES: "🔀 Cambiar el pagador",
		faIR: "🔀 تغییر پرداخت کننده",
		itIT: "🔀 Cambia il pagatore",
		ruRU: "🔀 Сменить плательщика",
	},
	COMMAND_TEXT_I_PAID: {
		deDE: "Ich habe bezahlt",
		enUK: "I paid",
		esES: "he pagado",
		faIR: "پرداخت کردم",
		itIT: "Ho pagato",
		ruRU: "Я заплатил",
	},
	COMMAND_TEXT_I_OWE: {
		deDE: "Ich schulde",
		enUK: "I owe",
		esES: "Yo debo",
		faIR: "من بدهکارم",
		itIT: "Sono in debito",
		ruRU: "Я должен",
	},
	COMMAND_TEXT_OWED_TO_ME: {
		deDE: "schuldet mir",
		enUK: "Owed to me",
		esES: "Me deben",
		faIR: "به من تعلق دارد",
		itIT: "È dovuto a me",
		ruRU: "Должны мне",
	},
	MESSAGE_TEXT_BILL_HEADER: {
		deDE: "Rechnung: %v",
		enUK: "Bill: %v",
		esES: "Cuenta: %v",
		faIR: "بیل :%v",
		itIT: "Bill: %v",
		ruRU: "Cчёт: %v",
	},
	MESSAGE_TEXT_NEW_DEBT_HEADER: {
		deDE: "Rechnung: %v",
		enUK: "Bill: %v",
		esES: "Cuenta: %v",
		faIR: "بیل: %v",
		itIT: "Bill: %v",
		ruRU: "Cчёт: %v",
	},
	MESSAGE_TEXT_GROUPS_ONLY_COMMAND: {
		deDE: "",
		enUK: "This command is available in group chats only for now.",
		esES: "",
		faIR: "",
		itIT: "",
		ruRU: "Эта команда пока что доступна только в групповых чатах",
	},
	MESSAGE_TEXT_ALREADY_HAS_CONTACT_WITH_SUCH_NAME: {
		deDE: "",
		enUK: "You already have contact with name: %v",
		esES: "",
		faIR: "",
		itIT: "",
		ruRU: "У вас уже есть контакт с таким именем: %v",
	},
	MESSAGE_TEXT_ALREADY_BILL_MEMBER: {
		deDE: "%v, du teilst diese Rechnung bereits.",
		enUK: "%v, you are sharing this bill already.",
		esES: "%v, estás compartiendo esta cuenta ya.",
		faIR: "%v، شما قبلا این لایحه را به اشتراک می گذارید",
		itIT: "%v, stai già condividendo questo disegno di legge.",
		ruRU: "%v, вы уже входите в состав участников.",
	},
	MESSAGE_TEXT_USER_JOINED_BILL: {
		deDE: "%v ist dem Teilen der Rechnung beigetreten.",
		enUK: "%v joined to bill sharing.",
		esES: "%v pagar conjuntamente.",
		faIR: "%v به اشتراک گذاری لایحه پیوست.",
		itIT: "%v unito alla condivisione di fatture.",
		ruRU: "%v присоеденился к совместной оплате.",
	},
	BUTTON_TEXT_I_PAID_FOR_THE_BILL: {
		deDE: "Die Rechnung wurde von mir bezahlt.",
		enUK: "The bill was paid by me.",
		enUS: "The bill was paid by me.",
		esES: "La factura fue pagada por mí.",
		faIR: "این لایحه توسط من پرداخت شد",
		frFR: "La facture a été payée par moi.",
		idID: "Tagihan dibayar oleh saya.",
		itIT: "Il conto è stato pagato da me.",
		jaJP: "請求書は私が支払いました。",
		koKR: "청구서는 내가 지불했습니다.",
		plPL: "Rachunek został zapłacony przeze mnie.",
		ptBR: "A conta foi paga por mim.",
		ptPT: "A conta foi paga por mim.",
		ruRU: "Этот счёт оплатил(а) я.",
		trTR: "Fatura benim tarafımdan ödendi.",
		ukUA: "Цей рахунок оплатив(ла) я.",
		uzUZ: "Hisob men tomonimdan to'landi.",
		zhCN: "账单由我支付。",
	},
	BUTTON_TEXT_I_OWE_FOR_THE_BILL: {
		deDE: "Ich muss noch was dabeigeben",
		enUK: "I owe for this bill",
		enUS: "I owe for this bill",
		esES: "Debo esta factura",
		faIR: "من برای این لایحه بدهکار هستم",
		frFR: "Je dois pour cette facture",
		idID: "Saya berhutang untuk tagihan ini",
		itIT: "Devo per questo disegno di legge",
		jaJP: "この請求書の支払いが必要です",
		koKR: "이 청구서에 대한 금액을 지불해야 합니다",
		plPL: "Jestem winien za ten rachunek",
		ptBR: "Eu devo por esta conta",
		ptPT: "Eu devo por esta conta",
		ruRU: "Я должен по этому счёту",
		trTR: "Bu fatura için borçluyum",
		ukUA: "Я винен за цим рахунком",
		uzUZ: "Men bu hisob uchun qarzdorman",
		zhCN: "我欠这个账单",
	},
	BUTTON_TEXT_I_DO_NOT_SHARE_THIS_BILL: {
		deDE: "Ich habe damit nichts zutun",
		enUK: "I don't share this bill",
		enUS: "I don't share this bill",
		esES: "No comparto esta cuenta",
		faIR: "من این لایحه را به اشتراک نمی گذارم",
		frFR: "Je ne partage pas cette facture",
		idID: "Saya tidak berbagi tagihan ini",
		itIT: "Non condivido questo disegno di legge",
		jaJP: "この請求書を共有していません",
		koKR: "이 청구서를 공유하지 않습니다",
		plPL: "Nie dzielę się tym rachunkiem",
		ptBR: "Eu não compartilho esta conta",
		ptPT: "Eu não compartilho esta conta",
		ruRU: "Я не учавствую в этой покупке",
		trTR: "Bu faturayı paylaşmıyorum",
		ukUA: "Я не беру участі в цій покупці",
		uzUZ: "Men bu hisobni baham ko'rmayman",
		zhCN: "我不分享这个账单",
	},
	MESSAGE_TEXT_YOU_JOINED_BILL: {
		deDE: "Du bist dem Teilen der Rechnung beigetreten.",
		enUK: "You've joined to bill sharing.",
		enUS: "You've joined to bill sharing.",
		esES: "Te has agregado para pagar conjuntamente .",
		faIR: "شما به اشتراک گذاشتن لایحه پیوستید",
		frFR: "Vous avez rejoint le partage de facture.",
		idID: "Anda telah bergabung dengan berbagi tagihan.",
		itIT: "Sei entrato a far parte della fatturazione.",
		jaJP: "請求書の共有に参加しました。",
		koKR: "청구서 공유에 참여하셨습니다.",
		plPL: "Dołączyłeś do dzielenia rachunku.",
		ptBR: "Você se juntou ao compartilhamento de contas.",
		ptPT: "Você se juntou ao compartilhamento de contas.",
		ruRU: "Вы присоеденились к совместной оплате.",
		trTR: "Fatura paylaşımına katıldınız.",
		ukUA: "Ви приєдналися до спільної оплати.",
		uzUZ: "Siz hisob-kitob ulashishga qo'shildingiz.",
		zhCN: "您已加入账单共享。",
	},
	ARTICLE_TITLE_SPLIT_BILL: {
		deDE: "eine Rechnung teilen",
		enUK: "Split bill/purchase",
		enUS: "Split bill/purchase",
		esES: "Compartir la cuenta/compra",
		faIR: "لایحه / خرید تقسیم شده",
		frFR: "Partager la facture/l'achat",
		idID: "Bagi tagihan/pembelian",
		itIT: "Bolletta Split / acquisto",
		jaJP: "請求書/購入を分割",
		koKR: "청구서/구매 분할",
		plPL: "Podziel rachunek/zakup",
		ptBR: "Dividir conta/compra",
		ptPT: "Dividir conta/compra",
		ruRU: "Разделить счёт/покупку",
		trTR: "Fatura/satın alma bölünmesi",
		ukUA: "Розділити рахунок/покупку",
		uzUZ: "Hisob/xaridni bo'lish",
		zhCN: "分摊账单/购买",
	},
	ARTICLE_SUBTITLE_SPLIT_BILL: {
		deDE: "Wert: %v\nTeile deine Kosten mit Freunden und verfolge deren Rückzahlungen.",
		enUK: "Amount: %v\nShares expenses with friends & track paybacks",
		enUS: "Amount: %v\nShares expenses with friends & track paybacks",
		esES: "Importe: %v\nCompartir los gastos entre amigos y seguir las devoluciones",
		faIR: "مقدار: %v" + "\n" + "هزینه ها را با دوستان و بازپرداخت پیگیری می کند",
		frFR: "Montant: %v\nPartagez les dépenses avec des amis et suivez les remboursements",
		idID: "Jumlah: %v\nBerbagi pengeluaran dengan teman & lacak pembayaran kembali",
		itIT: "Importo: %v\nDisponi i costi con gli amici e le retribuzioni delle tracce",
		jaJP: "金額: %v\n友達と費用を分担し、返済を追跡",
		koKR: "금액: %v\n친구와 비용을 공유하고 상환을 추적",
		plPL: "Kwota: %v\nDziel wydatki z przyjaciółmi i śledź spłaty",
		ptBR: "Valor: %v\nCompartilhe despesas com amigos e acompanhe reembolsos",
		ptPT: "Valor: %v\nCompartilhe despesas com amigos e acompanhe reembolsos",
		ruRU: "Сумма: %v\nПоделить траты между друзьями и отследить возвраты",
		trTR: "Tutar: %v\nArkadaşlarınızla masrafları paylaşın ve geri ödemeleri takip edin",
		ukUA: "Сума: %v\nПоділити витрати між друзями та відстежити повернення",
		uzUZ: "Miqdori: %v\nDo'stlar bilan xarajatlarni baham ko'ring va to'lovlarni kuzating",
		zhCN: "金额: %v\n与朋友分摊费用并跟踪还款",
	},

	ARTICLE_NEW_DEBT_TITLE: {
		deDE: "Neuer Schuldschein",
		enUK: "New debt",
		enUS: "New debt",
		esES: "Nueva deuda",
		faIR: "بدهی جدید",
		frFR: "Nouvelle dette",
		idID: "Hutang baru",
		itIT: "Nuovo debito",
		jaJP: "新しい借金",
		koKR: "새로운 부채",
		plPL: "Nowy dług",
		ptBR: "Nova dívida",
		ptPT: "Nova dívida",
		ruRU: "Новый долг",
		trTR: "Yeni borç",
		ukUA: "Новий борг",
		uzUZ: "Yangi qarz",
		zhCN: "新债务",
	},
	ARTICLE_NEW_DEBT_SUBTITLE: {
		deDE: "Wert: %v\nZur Fälligkeit wird eine Benachrichtigung geschickt, falls so eingestellt",
		enUK: "Amount: %v\nSends notifications on due date if set",
		enUS: "Amount: %v\nSends notifications on due date if set",
		esES: "Importe: %v\nEnviar las notificaciones el día de vencimiento",
		faIR: "مقدار: %v" + "\n" + "اگر تنظیم شود، اطلاعیه ها را در تاریخ تعیین شده ارسال می کند",
		frFR: "Montant: %v\nEnvoie des notifications à la date d'échéance si défini",
		idID: "Jumlah: %v\nMengirim notifikasi pada tanggal jatuh tempo jika diatur",
		itIT: "Importo: %v\nSend le notifiche alla data di scadenza se impostato",
		jaJP: "金額: %v\n設定されている場合、期日に通知を送信",
		koKR: "금액: %v\n설정된 경우 만기일에 알림 전송",
		plPL: "Kwota: %v\nWysyła powiadomienia w terminie płatności, jeśli ustawiono",
		ptBR: "Valor: %v\nEnvia notificações na data de vencimento, se definido",
		ptPT: "Valor: %v\nEnvia notificações na data de vencimento, se definido",
		ruRU: "Сумма: %v\nЗапись долга и рассылка оповещений в день возврата.",
		trTR: "Tutar: %v\nAyarlanmışsa vade tarihinde bildirim gönderir",
		ukUA: "Сума: %v\nНадсилає сповіщення в день повернення, якщо встановлено",
		uzUZ: "Miqdori: %v\nAgar o'rnatilgan bo'lsa, belgilangan sanada bildirishnomalar yuboradi",
		zhCN: "金额: %v\n如果设置，将在到期日发送通知",
	},
	SPLITUS_PLEASE_JOIN_IF_NOT_ON_THE_LIST: {
		deDE: `Bitte tritt zuerst bei, falls dein Name nicht auf der Liste ist.`,
		enUK: `Please join if your name is not on the list above.`,
		enUS: `Please join if your name is not on the list above.`,
		esES: `Por favor únete si tu nombre no está en la lista anterior.`,
		faIR: `اگر نام شما در لیست بالا نیست، لطفا پیوست شوید.`,
		frFR: `Veuillez vous joindre si votre nom ne figure pas dans la liste ci-dessus.`,
		idID: `Silakan bergabung jika nama Anda tidak ada dalam daftar di atas.`,
		itIT: `Si prega di unirti se il tuo nome non è nell'elenco di cui sopra.`,
		jaJP: `あなたの名前が上記のリストにない場合は、参加してください。`,
		koKR: `이름이 위 목록에 없으면 참여하세요.`,
		plPL: `Dołącz, jeśli Twojego imienia nie ma na powyższej liście.`,
		ptBR: `Por favor, junte-se se o seu nome não estiver na lista acima.`,
		ptPT: `Por favor, junte-se se o seu nome não estiver na lista acima.`,
		ruRU: `Пожалуйста присоеденяйтесь если ваше не в списке.`,
		trTR: `Adınız yukarıdaki listede yoksa lütfen katılın.`,
		ukUA: `Будь ласка, приєднуйтесь, якщо вашого імені немає у списку вище.`,
		uzUZ: `Agar ismingiz yuqoridagi ro'yxatda bo'lmasa, iltimos, qo'shiling.`,
		zhCN: `如果您的名字不在上面的列表中，请加入。`,
	},
	SPLITUS_TEXT_HI_IN_GROUP: {
		deDE: `Ich bin <b>Splitus</b>. Danke fürs hinzufügen!`,
		enUK: `I'm <b>Splitus</b>. Thanks for adding me!`,
		enUS: `I'm <b>Splitus</b>. Thanks for adding me!`,
		esES: `Soy <b>Splitus</b>. ¡Gracias por agregarme!`,
		faIR: `من <b>Splitus</b> با تشکر برای اضافه کردن من!`,
		frFR: `Je suis <b>Splitus</b>. Merci de m'avoir ajouté!`,
		idID: `Saya <b>Splitus</b>. Terima kasih telah menambahkan saya!`,
		itIT: `Sono <b>Splitus</b>. Grazie per averci aggiunto!`,
		jaJP: `私は<b>Splitus</b>です。追加してくれてありがとう！`,
		koKR: `저는 <b>Splitus</b>입니다. 저를 추가해 주셔서 감사합니다!`,
		plPL: `Jestem <b>Splitus</b>. Dziękuję za dodanie mnie!`,
		ptBR: `Eu sou <b>Splitus</b>. Obrigado por me adicionar!`,
		ptPT: `Eu sou <b>Splitus</b>. Obrigado por me adicionar!`,
		ruRU: `Меня зовут <b>Сплитус</b>. Спасибо что добавили!`,
		trTR: `Ben <b>Splitus</b>. Beni eklediğiniz için teşekkürler!`,
		ukUA: `Мене звати <b>Сплітус</b>. Дякую, що додали!`,
		uzUZ: `Men <b>Splitus</b>man. Meni qo'shganingiz uchun rahmat!`,
		zhCN: `我是<b>Splitus</b>。感谢您添加我！`,
	},
	COLLECTUS_TEXT_HI_IN_GROUP: {
		deDE: `Ich bin <b>Collectus.</b> Danke fürs hinzufügen!`,
		enUK: `I'm <b>Collectus.</b> Thanks for adding me!`,
		enUS: `I'm <b>Collectus.</b> Thanks for adding me!`,
		esES: `Soy <b>Collectus.</b> ¡Gracias por agregarme!`,
		faIR: `من <b>Collectus</b> با تشکر برای اضافه کردن من!`,
		frFR: `Je suis <b>Collectus.</b> Merci de m'avoir ajouté!`,
		idID: `Saya <b>Collectus.</b> Terima kasih telah menambahkan saya!`,
		itIT: `Sono <b>Collectus.</b> Grazie per averci aggiunto!`,
		jaJP: `私は<b>Collectus.</b>です。追加してくれてありがとう！`,
		koKR: `저는 <b>Collectus.</b>입니다. 저를 추가해 주셔서 감사합니다!`,
		plPL: `Jestem <b>Collectus.</b> Dziękuję za dodanie mnie!`,
		ptBR: `Eu sou <b>Collectus.</b> Obrigado por me adicionar!`,
		ptPT: `Eu sou <b>Collectus.</b> Obrigado por me adicionar!`,
		ruRU: `Меня зовут <b>Коллектус.</b> Спасибо что добавили!`,
		trTR: `Ben <b>Collectus.</b> Beni eklediğiniz için teşekkürler!`,
		ukUA: `Мене звати <b>Коллектус.</b> Дякую, що додали!`,
		uzUZ: `Men <b>Collectus.</b>man. Meni qo'shganingiz uchun rahmat!`,
		zhCN: `我是<b>Collectus.</b>感谢您添加我！`,
	},
	MT_GROUP_LABEL: {
		deDE: `<b>Gruppe</b>: %v`,
		enUK: `<b>Group</b>: %v`,
		enUS: `<b>Group</b>: %v`,
		esES: `<b>Grupo</b>: %v`,
		faIR: `<b>گروه</b>: %v`,
		frFR: `<b>Groupe</b>: %v`,
		idID: `<b>Grup</b>: %v`,
		itIT: `<b>Gruppo</b>: %v`,
		jaJP: `<b>グループ</b>: %v`,
		koKR: `<b>그룹</b>: %v`,
		plPL: `<b>Grupa</b>: %v`,
		ptBR: `<b>Grupo</b>: %v`,
		ptPT: `<b>Grupo</b>: %v`,
		ruRU: `<b>Группа</b>: %v`,
		trTR: `<b>Grup</b>: %v`,
		ukUA: `<b>Група</b>: %v`,
		uzUZ: `<b>Guruh</b>: %v`,
		zhCN: `<b>组</b>: %v`,
	},
	MT_SPONSORS_HEADER: {
		deDE: `<b>Sponsoren</b>:`,
		enUK: `<b>Sponsors</b>:`,
		enUS: `<b>Sponsors</b>:`,
		esES: `<b>Patrocinadores</b>:`,
		faIR: `<b>حامیان</b>:`,
		frFR: `<b>Sponsors</b>:`,
		idID: `<b>Sponsor</b>:`,
		itIT: `<b>Sponsors</b>:`,
		jaJP: `<b>スポンサー</b>:`,
		koKR: `<b>스폰서</b>:`,
		plPL: `<b>Sponsorzy</b>:`,
		ptBR: `<b>Patrocinadores</b>:`,
		ptPT: `<b>Patrocinadores</b>:`,
		ruRU: `<b>Спонсоры</b>:`,
		trTR: `<b>Sponsorlar</b>:`,
		ukUA: `<b>Спонсори</b>:`,
		uzUZ: `<b>Homiylar</b>:`,
		zhCN: `<b>赞助商</b>:`,
	},
	MT_DEBTORS_HEADER: {
		deDE: `<b>Schuldner</b>:`,
		enUK: `<b>Debtors</b>:`,
		enUS: `<b>Debtors</b>:`,
		esES: `<b>Deudores</b>:`,
		faIR: `<b>بدهکاران</b>:`,
		frFR: `<b>Débiteurs</b>:`,
		idID: `<b>Debitur</b>:`,
		itIT: `<b>Debitori</b>:`,
		jaJP: `<b>債務者</b>:`,
		koKR: `<b>채무자</b>:`,
		plPL: `<b>Dłużnicy</b>:`,
		ptBR: `<b>Devedores</b>:`,
		ptPT: `<b>Devedores</b>:`,
		ruRU: `<b>Должники</b>:`,
		trTR: `<b>Borçlular</b>:`,
		ukUA: `<b>Боржники</b>:`,
		uzUZ: `<b>Qarzdorlar</b>:`,
		zhCN: `<b>债务人</b>:`,
	},
	BT_DEFAULT_CURRENCY: {
		deDE: `Währung: %v`,
		enUK: `Currency: %v`,
		enUS: `Currency: %v`,
		esES: `Moneda: %v`,
		faIR: `واحد پول: %v`,
		frFR: `Devise: %v`,
		idID: `Mata uang: %v`,
		itIT: `Moneta: %v`,
		jaJP: `通貨: %v`,
		koKR: `통화: %v`,
		plPL: `Waluta: %v`,
		ptBR: `Moeda: %v`,
		ptPT: `Moeda: %v`,
		ruRU: `Валюта: %v`,
		trTR: `Para birimi: %v`,
		ukUA: `Валюта: %v`,
		uzUZ: `Valyuta: %v`,
		zhCN: `货币: %v`,
	},
	MESSAGE_TEXT_ASK_LANG: {
		deDE: `Welche Sprache wird hier gesprochen?`,
		enUK: `What language should I use in this group?`,
		enUS: `What language should I use in this group?`,
		esES: `¿Qué idioma debería usar en este grupo?`,
		faIR: `کدام زبان باید در این گروه استفاده کنم؟`,
		frFR: `Quelle langue dois-je utiliser dans ce groupe?`,
		idID: `Bahasa apa yang harus saya gunakan dalam grup ini?`,
		itIT: `Che lingua devo utilizzare in questo gruppo?`,
		jaJP: `このグループではどの言語を使用すべきですか？`,
		koKR: `이 그룹에서 어떤 언어를 사용해야 합니까?`,
		plPL: `Jakiego języka powinienem używać w tej grupie?`,
		ptBR: `Que idioma devo usar neste grupo?`,
		ptPT: `Que idioma devo usar neste grupo?`,
		ruRU: `Какой язык я должен использовать в этой группе?`,
		trTR: `Bu grupta hangi dili kullanmalıyım?`,
		ukUA: `Якою мовою мені спілкуватися в цій групі?`,
		uzUZ: `Bu guruhda qaysi tildan foydalanishim kerak?`,
		zhCN: `我应该在这个群组中使用什么语言？`,
	},
	MESSAGE_TEXT_HI_IN_GROUP_LANG_SET: {
		deDE: `Kein Problem, dann schreibe ich auf Deutsch.`,
		enUK: `Great, I'll be using English.`,
		enUS: `Great, I'll be using English.`,
		esES: `Genial, usaré español.`,
		faIR: `عالی، من از فارسی استفاده خواهم کرد.`,
		frFR: `Super, j'utiliserai le français.`,
		idID: `Bagus, saya akan menggunakan bahasa Indonesia.`,
		itIT: `Ottimo, userò l'italiano.`,
		jaJP: `素晴らしい、日本語を使用します。`,
		koKR: `좋아요, 한국어를 사용하겠습니다.`,
		plPL: `Świetnie, będę używać języka polskiego.`,
		ptBR: `Ótimo, vou usar o português.`,
		ptPT: `Ótimo, vou usar o português.`,
		ruRU: `Отлично, я буду использовать русский`,
		trTR: `Harika, Türkçe kullanacağım.`,
		ukUA: `Чудово, я буду використовувати українську мову.`,
		uzUZ: `Ajoyib, men o'zbek tilidan foydalanaman.`,
		zhCN: `太好了，我将使用中文。`,
	},
	SPLITUS_TEXT_ABOUT_ME_AND_CO: {
		deDE: `Ich kann helfen, <b>Rechnungen zu teilen</b>. Mein Freund @DebtsTrackerBot passt darauf auf, dass alle Schulden zurückgezahlt werden.`,
		enUK: `I help to <b>split bills</b>. My friend @DebtsTrackerBot is tracking paybacks & debts.`,
		esES: `Ayudo a <b>dividir billetes</b>. Mi amigo @DebtsTrackerBot está rastreando pagos y deudas.`,
		faIR: `من به <b>تقسیم صورتحساب </b> کمک میکنم دوست منDebtsTrackerBot ردیابی بازپرداخت و بدهی است.`,
		itIT: `Aiuto a <b>dividere le bollette</b>. Il mio amico @DebtsTrackerBot sta monitorando i pagamenti e i debiti.`,
		ruRU: `Я помогаю делить счета. Мой друг @DebtsTrackerRuBot отслеживает платежи и долги.`,
	},
	COLLECTUS_TEXT_ABOUT_ME_AND_CO: {
		deDE: `Ich kann dabei helfen <b>Geld zu sammeln</b>. Zum Beispiel für ein Geburtstagsgeschenk. 🎉

Mein Freund @DebtsTrackerBot überwacht Schulden und deren Rückzahlungen.

Wenn du Geld für irgendwas sammeln willst, kann @SplitusBot dir dabei helfen.`,

		enUK: `I help to <b>collect money</b> for a good cause. For example for a birthday present. 🎉

My buddy @DebtsTrackerBot is tracking debts & paybacks.

And if you do collective purchases and want to split bills @SplitusBot is here to help.`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `Я помогаю <b>собирать деньги</b> на что нибудь. Например для подарка на день рождение. 🎉

Мой друг @DebtsTrackerRuBot отслеживает долги и платежи.

А если хотите вести учёт совместных покупок и удобно делить счета вам поможет @SplitusBot.`,
	},
	SPLITUS_TEXT_HI: {
		deDE: `Ich bin <b>Splitus</b>.`,
		enUK: `I'm <b>Splitus</b>.`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `Меня зовут <b>Сплитус</b>.`,
	},
	COLLECTUS_TEXT_HI: {
		deDE: `Ich bin <b>Collectus</b>.`,
		enUK: `I'm <b>Collectus</b>.`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `Меня зовут <b>Коллектус</b>.`,
	},
	SPLITUS_TG_COMMANDS: {
		deDE: ``,
		enUK: `<b>Bot commands:</b>
	/groups - List of groups
	/bills - List of outstanding bills
	/menu - Main menu
	/settings - Settings
	/help - Learn how to use bot, report issues, ask questions`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `<b>Команды для бота:</b>
	/groups - Список групп
	/bills - Список незакрытых платежей
	/menu - Главное меню
	/settings - Настройки
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	COLLECTUS_TG_COMMANDS: {
		enUK: `<b>Bot commands:</b>

	/groups - List of groups
	/fundraisings - List of active fundraisings
	/help - Learn how to use bot, report issues, ask questions`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `<b>Команды для бота:</b>
	/groups - Список групп
	/fundraisings - Список активных сборов
	/help - Узнать как использовать, сообщить о проблеме, задать вопрос`,
	},
	MESSAGE_TEXT_SEND_HELP_COMMAND_FOR_HELP: { // This is the same for all languages.
		deDE: `Sende /help für Hilfe über den Umgang mit diesen Bot.`,
		enUK: `Send /help for details on how to use this bot.`,
		esES: ``,
		faIR: ``,
		itIT: ``,
		ruRU: `Отправьте /help для справки по использованию бота.`,
	},
	MESSAGE_TEXT_HI: { // This is the same for all languages.
		deDE: `¡Hola! Hi! Привет! سلام! Hallo!`,
		enUK: `¡Hola! Hi! Привет! سلام! Hallo!`,
		esES: `¡Hola! Hi! Привет! سلام! Hallo!`,
		faIR: `¡Hola! Hi! Привет! سلام! Hallo!`,
		itIT: `¡Hola! Hi! Привет! سلام! Hallo!`,
		ruRU: `¡Hola! Hi! Привет! سلام! Hallo!`,
	},
	MESSAGE_TEXT_HI_USERNAME: { // This is the same for all languages.
		deDE: `Hi %v!`,
		enUK: `Hi %v!`,
		esES: `¡Hola %v!`,
		faIR: ``,
		itIT: ``,
		ruRU: `Привет %v!`,
	},
	MESSAGE_TEXT_BACK_TO_MAIN_MENU: {
		deDE: `Du kannst zurück zum Haupt /menu`,
		enUK: `You can go back to main /menu`,
		esES: `Puedes volver al inicio /menú`,
		faIR: `شما میتوانید به /منو ی اصلی مراجعه کنید.`,
		itIT: `Puoi tornare al menu' principale tramite /menu`,
		ruRU: `Можно вернуться назад в главное /меню`,
	},
	MESSAGE_TEXT_YOUR_SELECTED_PREFERRED_LANGUAGE: { // This is the same for all languages.
		deDE: `Bevorzugte Sprache: %s`,
		enUK: `Preferred bot language: %s`,
		esES: `Idioma favorito del bot: %s`,
		faIR: `زبان برنامه: %s`,
		itIT: `Lingua del bot preferita: %s`,
		ruRU: `Выбранный язык бота: %s`,
	},
	MESSAGE_TEXT_ONBOARDING_ASK_TO_CHOOSE_LANGUAGE: {
		deDE: `Welche ist deine bevorzugte Sprache?`,
		enUK: `What is your preferred language?`,
		esES: `¿cuál es tu idioma preferida?`,
		faIR: `شما چه زبانی را ترجیح می دهید؟`,
		itIT: `Qual'e' la tua lingua madre?`,
		ruRU: `На каком языке вы хотели бы общаться?`,
	},
	MESSAGE_TEXT_CHOOSE_UI_LANGUAGE: {
		deDE: "Auf welcher Sprache würdest du dich gerne mit mir unterhalten?",
		enUK: "Which language you would like to talk to me?",
		esES: "¿En cuál idioma preferirías hablar conmigo?",
		faIR: "دوست دارید به چه زبانی با من صحبت کنید؟",
		itIT: "In quale lingua preferisci parlarmi?",
		ruRU: "На каком языке вы хотели бы общаться со мной?",
	},
	MESSAGE_TEXT_LOCALE_CHANGED: {
		deDE: "Du hast die Sprache geändert auf %v",
		enUK: "You've switched language to %v",
		esES: "Has cambiado el idioma a %v",
		faIR: "شما زبان را به %v تغییر دادید. ",
		itIT: "Hai cambiato lingua in %v",
		ruRU: "Вы поменяли язык на %v",
	},
	MESSAGE_TEXT_WHATS_NEXT: {
		deDE: "Was als nächstes?",
		enUK: "What's next?",
		esES: "¿Algo más?",
		faIR: "اقدام بعدی چیست؟",
		itIT: "Cosa posso fare ora?",
		ruRU: "Что будем делать дальше?",
	},
	MESSAGE_TEXT_REFERRERS_TITLE: {
		enUK: "Our friends:",
		ruRU: "Наши друзья:",
	},
	COMMAND_TEXT_ADD_MY_TG_CHANNEL: {
		enUK: "Add my channel",
		ruRU: "Добавить мой канал",
	},
	MESSAGE_TEXT_DEBTUS_COMMANDS: {
		enUK: `<b>Commands for the bot</b>
🏠 /menu - show main menu
🔙 /return - return previously recorded debt
📥 /got - record money you received from others
📤 /gave - record money you gave to others
📚 /history - latest transactions
🏁 /balance - display current balance
⚙ /settings - adjust your preferences`,

		ruRU: `<b>Команды бота</b>
🏠 /menu - показать основное меню
🔙 /return - записать возврат долга
📥 /got - записать о взятии в долг
📤 /gave - записать о том что дали взаймы
📚 /history - последние транзакции
🏁 /balance - показать текущий баланс
⚙ /settings - настройки`,
	},
	MESSAGE_TEXT_WHATS_NEXT_HINT: {
		deDE: `
Wenn du dir was von jemanden geliehen hast, wähle /anleihen.
Wenn du was an jemanden verliehen hast, wähle /verleihen.

Oder benutzt das Menü unten.`,
		enUK: `
If you borrowed from someone to record it use /got.
If you lent to someone to record it use /gave.

Or use menu at the bottom.`,

		esES: `
		Si alguien te ha prestado usa el comando  /recibido.
		Si has prestado a alguien usa el comando /dado.

O usa el menú de abajo en la pantalla.`,

		faIR: `
اگر از کسی قرض گرفته اید برای ثبت آن از /قرض_گرفتن استفاده کنید.
اگر به کسی قرض داده اید برای ثبت آن از /قرض_دادن استفاده کنید.

یا از منوی پایین استفاده نمایید.`,

		itIT: `
Se qualcuno ti ha prestato qualcosa per memorizzarlo usa /got.
Se hai prestato qualcosa a qualcuno per memorizzarlo usa /gave.

O usa il menu' qui sotto.`,

		ruRU: `
	Если вы дали в долг воспользуйтесь командой /дал.
	Если вы одолжили что-то - командой /взял.

	Или воспользуйтесь меню внизу экрана.`,
	},
	MESSAGE_TEXT_HISTORY_HEADER: {
		deDE: "Verlauf",
		enUK: "History",
		esES: "Cronología",
		faIR: "سوابق",
		itIT: "Cronologia",
		ruRU: "История",
	},
	MESSAGE_TEXT_HISTORY_NO_RECORDS: {
		deDE: "Du hast noch nichts gespeichert.",
		enUK: "You don't have any records yet.",
		esES: "Aún no tienes ninguna notificación.",
		faIR: "شما هنوز هیچ ثبت سابقه ای ندارید.",
		itIT: "Non hai nulla memorizzato.",
		ruRU: "У вас пока нет ни одной записи.",
	},
	MESSAGE_TEXT_HISTORY_LIST: {

		deDE: `<b>%v</b> <i>(bis %d):</i>
─────────────
%v`,
		enUK: `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		enUS: `<b>%v</b> <i>(last %d):</i>
─────────────
%v`,
		esES: `<b>%v</b> <i>(últimos %d):</i>
─────────────
%v`,

		faIR: `<b>%v</b> <i>(آخرین %d):</i>
─────────────
%v`,
		frFR: "",
		idID: "",
		itIT: `<b>%v</b> <i>(ultimi %d):</i>

─────────────
%v`,
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: `<b>%v</b> <i>(%d последних)</i>
─────────────
%v`,
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_IS_ZERO: {
		deDE: "Keine unbeglichenen Schulden.",
		enUK: "You have no records on current debts.",
		enUS: "You have no records on current debts.",
		esES: "No hay ninguna notificación de deudas actuales.",
		faIR: "شما در خصوص بدهی های اخیر ثبت سابقه ای ندارید.",
		frFR: "",
		idID: "",
		itIT: "Non hai nulla memorizzato nel debito corrente.",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Нет записей о текущих долгах.",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_TOTAL_INTRO: {
		deDE: "Insgesamt",
		enUK: "Total",
		enUS: "Total",
		esES: "Total",
		faIR: "مجموع",
		frFR: "",
		idID: "",
		itIT: "Totale",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Всего",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	BT_OTHER_CURRENCY: {
		deDE: "",
		enUK: "Another currency",
		enUS: "Another currency",
		esES: "Otra moneda",
		faIR: "ارز دیگر",
		frFR: "",
		idID: "",
		itIT: "Un'altra valuta",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Другая валюта",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_PRIMARY_CURRENCY_IS_SET_TO: {
		deDE: "OK, von nun an ist '%v' deine Hauptwährung.",
		enUK: "OK, from now on I will use '%v' as a primary currency.",
		enUS: "OK, from now on I will use '%v' as a primary currency.",
		esES: "OK, ahora voy a usar '%v' como moneda principal. ",
		faIR: "بسیار خوب، از الان من از '%v' بعنوان واحد پول اولیه استفاده می کنم",
		frFR: "",
		idID: "",
		itIT: "OK, da ora in poi usero' '%v' come moneta principale.",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "OK, теперь я буду использовать '%v' как основную валюту.",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_TO_USER: {
		deDE: "%v - schuldet dir %v",
		enUK: "%v - owes you %v",
		enUS: "%v - owes to you %v",
		esES: "%v - te debe %v",
		faIR: "%v - %v به شما بدهکار است ",
		frFR: "",
		idID: "",
		itIT: "%v - ti deve %v.",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "%v - долг вам %v",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_TO_USER: {
		deDE: "Schuldet dir %v",
		enUK: "Owes to you %v",
		enUS: "Owes to you %v",
		esES: "Te debe %v",
		faIR: "%v به شما بدهکار است ",
		frFR: "",
		idID: "",
		itIT: "Hai un credito di %v",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Вам должны %v",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_ON_RETURN_USER_DOES_NOT_OWE_ANYTHING_TO_COUNTERPARTY_ANYMORE: {
		deDE: "Hurra, du bist jetzt quitt mit <b>%v</b>.",
		enUK: "Congratulations! You don't owe anything more to <b>%v</b>.",
		enUS: "Congratulations! You don't owe anything more to <b>%v</b>.",
		esES: "Bravo! Has saldado tu deuda con <b>%v</b>.",
		faIR: "تبریک! شما دیگر چیزی به <b>%v</b> بدهکار نیستید .",
		frFR: "",
		idID: "",
		itIT: "Bravo! Hai saldato il tuo debito con <b>%v</b>.",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Поздравляем! У вас не осталось долгов перед <b>%v</b>.",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_ON_RETURN_COUNTERPARTY_DOES_NOT_OWE_ANYTHING_TO_USER_ANYMORE: {
		deDE: "Du bist jetzt mit <b>%v</b> quitt.",
		enUK: "<b>%v</b> does not owe anything more to you.",
		enUS: "<b>%v</b> does not owe anything more to you.",
		esES: "<b>%v</b> nadie te debe nada ya.",
		faIR: "<b>%v</b> دیگر چیزی به شما بدهکار نیست",
		frFR: "",
		idID: "",
		itIT: "<b>%v</b> ha saldato il suo debito verso di te.",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "У <b>%v</b> больше не осталось долгов перед вами.",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_CURRENCY_ROW_DEBT_BY_USER: {
		deDE: "Du schuldest %v",
		enUK: "You owe %v",
		enUS: "You owe %v",
		esES: "Tú debes %v",
		faIR: "شما %v بدهکار هستید ",
		frFR: "",
		idID: "",
		itIT: "Hai un debito di %v",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Вы должны %v",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_BALANCE_SINGLE_CURRENCY_COUNTERPARTY_DEBT_BY_USER: {
		deDE: "%v - du schuldest %v",
		enUK: "%v - you owe %v",
		enUS: "%v - you owe %v",
		esES: "%v - tú debes %v",
		faIR: "%v - شما %v بدهکار هستید ",
		frFR: "",
		idID: "",
		itIT: "%v - tu gli/le devi %v",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "%v - вы должны %v",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY: {
		deDE: "Was ist deine Hauptwährung?",
		enUK: "What is your primary currency?",
		enUS: "What is your primary currency?",
		esES: "¿Cuál es tu moneda principal?",
		faIR: "واحد پولی اولیه شما چیست؟",
		frFR: "",
		idID: "",
		itIT: "Qual'e' la tua valuta principale?",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Какая валюта для вас основная?",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_ASK_PRIMARY_CURRENCY_FOR_GROUP: {
		deDE: "",
		enUK: "What is a primary currency for this group?",
		enUS: "What is a primary currency for this group?",
		esES: "¿Cuál es tu moneda principal?",
		faIR: "واحد پولی اولیه شما چیست؟",
		frFR: "",
		idID: "",
		itIT: "Qual'e' la tua valuta principale?",
		jaJP: "",
		koKR: "",
		plPL: "",
		ptBR: "",
		ruRU: "Какая валюта основная для этой группы?",
		trTR: "",
		ukUA: "",
		uzUZ: "",
		zhCN: "",
	},
	MESSAGE_TEXT_FAILED_TO_DELETE_USER: {
		deDE: "Konnte den Benutzer nicht löschen: %v",
		enUK: "Failed to delete user: %v",
		enUS: "Failed to delete user: %v",
		esES: "Error durante la cancelación del usuario: %v",
		faIR: "حذف کاربر ناموفق بود: %v",
		frFR: "Échec de la suppression de l'utilisateur: %v",
		idID: "Gagal menghapus pengguna: %v",
		itIT: "Errore durante la cancellazione dell'utente: %v",
		jaJP: "ユーザーの削除に失敗しました: %v",
		koKR: "사용자 삭제 실패: %v",
		plPL: "Nie udało się usunąć użytkownika: %v",
		ptBR: "Falha ao excluir usuário: %v",
		ruRU: "Не удалось удалить данные пользователя: %v",
		trTR: "Kullanıcı silinemedi: %v",
		ukUA: "Не вдалося видалити користувача: %v",
		uzUZ: "Foydalanuvchini o'chirib bo'lmadi: %v",
		zhCN: "删除用户失败: %v",
	},
	MESSAGE_TEXT_USER_DELETED: {
		deDE: "Die Benutzerdaten wurden gelöscht.",
		enUK: "User's data has been deleted",
		enUS: "User's data has been deleted",
		esES: "Los datos del usuario han sido eliminados",
		faIR: "اطلاعات کاربر حذف شد.",
		frFR: "Les données de l'utilisateur ont été supprimées",
		idID: "Data pengguna telah dihapus",
		itIT: "I dati dell'utente sono stati cancellati",
		jaJP: "ユーザーのデータが削除されました",
		koKR: "사용자 데이터가 삭제되었습니다",
		plPL: "Dane użytkownika zostały usunięte",
		ptBR: "Os dados do usuário foram excluídos",
		ruRU: "Данные пользователя удалены",
		trTR: "Kullanıcı verileri silindi",
		ukUA: "Дані користувача видалено",
		uzUZ: "Foydalanuvchi ma'lumotlari o'chirildi",
		zhCN: "用户数据已被删除",
	},
	MESSAGE_TEXT_RETURN_ASK_TO_CHOOSE_COUNTERPARTY: {
		deDE: "Bitte wähle, wer die Schuld beglichen hat oder wem du sie zurückgezahlt hast.",
		enUK: "Please choose who returned the debt or to who you returned it.",
		enUS: "Please choose who returned the debt or to who you returned it.",
		esES: "Por favor, elige quien ha devuelto o a quien ha sido devuelta la deuda ",
		faIR: "لطفاً انتخاب کنید چه کسی بدهی اش را به شما پرداخت کرده یا شما بدهیتان را به چه کسی بازپرداخت نموده اید.",
		frFR: "Veuillez choisir qui a remboursé la dette ou à qui vous l'avez remboursée.",
		idID: "Silakan pilih siapa yang mengembalikan hutang atau kepada siapa Anda mengembalikannya.",
		itIT: "Scegli con chi hai sanato un credito o un debito.",
		jaJP: "誰が借金を返済したか、または誰に返済したかを選択してください。",
		koKR: "누가 부채를 반환했는지 또는 누구에게 반환했는지 선택하십시오.",
		plPL: "Wybierz, kto zwrócił dług lub komu go zwróciłeś.",
		ptBR: "Por favor, escolha quem devolveu a dívida ou para quem você a devolveu.",
		ruRU: "Выберете кому вы вернули долг или кто вернул его вам.",
		trTR: "Lütfen borcu kimin iade ettiğini veya kime iade ettiğinizi seçin.",
		ukUA: "Будь ласка, виберіть, хто повернув борг або кому ви його повернули.",
		uzUZ: "Iltimos, qarzni kim qaytarganini yoki siz kimga qaytarganingizni tanlang.",
		zhCN: "请选择谁归还了债务或您归还给了谁。",
	},
	MESSAGE_TEXT_CHOOSE_DEBT_THAT_HAS_BEEN_RETURNED: {
		deDE: "Bitte wähle, ob die Schuld vollständig oder teilweise beglichen wurde.",
		enUK: "Please choose a debt that has been returned fully or partially.",
		enUS: "Please choose a debt that has been returned fully or partially.",
		esES: "Por favor, elige una deuda que ha sido devuelta total o parcialmente. ",
		faIR: "لطفاً انتخاب کنید تمام یا بخشی از کدام بدهی پرداخت شده است.",
		frFR: "Veuillez choisir une dette qui a été remboursée entièrement ou partiellement.",
		idID: "Silakan pilih hutang yang telah dikembalikan sepenuhnya atau sebagian.",
		itIT: "Scegli un debito che e' stato restituito completamente o parzialmente.",
		jaJP: "完全に、または部分的に返済された借金を選択してください。",
		koKR: "전액 또는 부분적으로 반환된 부채를 선택하십시오.",
		plPL: "Wybierz dług, który został zwrócony w całości lub częściowo.",
		ptBR: "Por favor, escolha uma dívida que foi devolvida total ou parcialmente.",
		ruRU: "Выберите долг который был возвращён целиком или частично.",
		trTR: "Lütfen tamamen veya kısmen iade edilmiş bir borç seçin.",
		ukUA: "Будь ласка, виберіть борг, який було повернуто повністю або частково.",
		uzUZ: "Iltimos, to'liq yoki qisman qaytarilgan qarzni tanlang.",
		zhCN: "请选择已全部或部分归还的债务。",
	},
	MESSAGE_TEXT_NO_DEBTS_TO_RETURN: {
		deDE: "Du hast keine Aufzeichnungen über Schulden, die zurückgegeben werden können.",
		enUK: "You have no records for debts that can be returned.",
		enUS: "You have no records for debts that can be returned.",
		esES: "No tienes registros de deudas que puedan ser devueltas.",
		faIR: "شما هیچ سابقه ای از بدهی هایی که قابل بازگشت باشند ندارید.",
		frFR: "Vous n'avez aucun enregistrement de dettes qui peuvent être remboursées.",
		idID: "Anda tidak memiliki catatan untuk hutang yang dapat dikembalikan.",
		itIT: "Non hai registrazioni di debiti che possono essere restituiti.",
		jaJP: "返済可能な借金の記録がありません。",
		koKR: "반환할 수 있는 부채에 대한 기록이 없습니다.",
		plPL: "Nie masz żadnych zapisów długów, które mogą zostać zwrócone.",
		ptBR: "Você não tem registros de dívidas que possam ser devolvidas.",
		ruRU: "У вас нет записей о догах для возврата.",
		trTR: "İade edilebilecek borçlar için hiçbir kaydınız yok.",
		ukUA: "У вас немає записів про борги, які можна повернути.",
		uzUZ: "Sizda qaytarilishi mumkin bo'lgan qarzlar uchun hech qanday yozuvlar yo'q.",
		zhCN: "您没有可以返还的债务记录。",
	},
	MESSAGE_TEXT_PLEASE_ACKNOWLEDGE_TRANSFER: {
		deDE: "Bitte stimme dem zu oder lehne es ab.",
		enUK: "Please confirm or decline this transfer.",
		enUS: "Please confirm or decline this transfer.",
		esES: "Por favor, confirma o rechaza la transacción.",
		faIR: "لطفاً این تراکنش را تایید یا رد نمایید.",
		frFR: "Veuillez confirmer ou refuser ce transfert.",
		idID: "Silakan konfirmasi atau tolak transfer ini.",
		itIT: "Conferma o rifiuta questo debito/credito.",
		jaJP: "この転送を確認または拒否してください。",
		koKR: "이 전송을 확인하거나 거부하십시오.",
		plPL: "Proszę potwierdzić lub odrzucić ten transfer.",
		ptBR: "Por favor, confirme ou recuse esta transferência.",
		ruRU: "Пожалуйста подтвердите или отклоните эту транзакцию.",
		trTR: "Lütfen bu transferi onaylayın veya reddedin.",
		ukUA: "Будь ласка, підтвердіть або відхиліть цей переказ.",
		uzUZ: "Iltimos, ushbu o'tkazmani tasdiqlang yoki rad eting.",
		zhCN: "请确认或拒绝此转账。",
	},
	MESSAGE_TEXT_ALREADY_ACCEPTED_TRANSFER: {
		deDE: "Du hast dem bereits zugestimmt.",
		enUK: "This transfer has been accepted already.",
		enUS: "This transfer has been accepted already.",
		esES: "Esta transacción ya ha sido aceptada.",
		faIR: "این تراکنش قبلا قبول شده است.",
		frFR: "Ce transfert a déjà été accepté.",
		idID: "Transfer ini sudah diterima.",
		itIT: "Questo debito/credito e' gia' stato accettato.",
		jaJP: "この転送はすでに承認されています。",
		koKR: "이 전송은 이미 수락되었습니다.",
		plPL: "Ten transfer został już zaakceptowany.",
		ptBR: "Esta transferência já foi aceita.",
		ruRU: "Эта транзакция уже подтверждена.",
		trTR: "Bu transfer zaten kabul edildi.",
		ukUA: "Цей переказ вже підтверджено.",
		uzUZ: "Bu o'tkazma allaqachon qabul qilingan.",
		zhCN: "此转账已被接受。",
	},
	MESSAGE_TEXT_ALREADY_DECLINED_TRANSFER: {
		deDE: "Du hast dem bereits widersprochen.",
		enUK: "This transfer has been declined already.",
		enUS: "This transfer has been declined already.",
		esES: "Esta transacción ya ha sido rechazada.",
		faIR: "این تراکنش قبلاً رد شده است.",
		frFR: "Ce transfert a déjà été refusé.",
		idID: "Transfer ini sudah ditolak.",
		itIT: "Questo debito/credito e' gia' stato rifiutato.",
		jaJP: "この転送はすでに拒否されています。",
		koKR: "이 전송은 이미 거부되었습니다.",
		plPL: "Ten transfer został już odrzucony.",
		ptBR: "Esta transferência já foi recusada.",
		ruRU: "Эта транзакция уже отклонена.",
		trTR: "Bu transfer zaten reddedildi.",
		ukUA: "Цей переказ вже відхилено.",
		uzUZ: "Bu o'tkazma allaqachon rad etilgan.",
		zhCN: "此转账已被拒绝。",
	},
	MESSAGE_TEXT_RECEIPT_LINK: {
		deDE: "Details hier: %v",
		enUK: "Details here: %v",
		enUS: "Details here: %v",
		esES: "Detalles aquí: %v",
		faIR: "جزئیات: %v",
		frFR: "Détails ici: %v",
		idID: "Detail di sini: %v",
		itIT: "Maggiori dettagli qui: %v",
		jaJP: "詳細はこちら: %v",
		koKR: "세부 정보: %v",
		plPL: "Szczegóły tutaj: %v",
		ptBR: "Detalhes aqui: %v",
		ruRU: "Подробнее тут: %v",
		trTR: "Ayrıntılar burada: %v",
		ukUA: "Деталі тут: %v",
		uzUZ: "Tafsilotlar shu yerda: %v",
		zhCN: "详情在这里: %v",
	},
	MESSAGE_TEXT_ASK_PHONE_NUMBER_OF_COUNTERPARTY: {
		deDE: "Bitte gib die Telefonnummer von <b>%v</b> an:",
		enUK: "Please provide phone number for <b>%v</b>",
		enUS: "Please provide phone number for <b>%v</b>",
		esES: "Por favor escribe el número de teléfono de <b>%v</b>",
		faIR: "لطفا شماره تلفن ایشان را وارد کنید <b>%v</b>",
		frFR: "Veuillez fournir le numéro de téléphone pour <b>%v</b>",
		idID: "Silakan berikan nomor telepon untuk <b>%v</b>",
		itIT: "Per favore fornisci il numero di telefono di <b>%v</b>",
		jaJP: "<b>%v</b>の電話番号を入力してください",
		koKR: "<b>%v</b>의 전화번호를 입력해주세요",
		plPL: "Proszę podać numer telefonu dla <b>%v</b>",
		ptBR: "Por favor, forneça o número de telefone para <b>%v</b>",
		ruRU: "Пожалуйста напишите номер телефона <b>%v</b>.",
		trTR: "Lütfen <b>%v</b> için telefon numarası girin",
		ukUA: "Будь ласка, вкажіть номер телефону для <b>%v</b>",
		uzUZ: "Iltimos, <b>%v</b> uchun telefon raqamini kiriting",
		zhCN: "请提供<b>%v</b>的电话号码",
	},
	MESSAGE_TEXT_USE_CONTACT_TO_SEND_PHONE_NUMBER: {
		deDE: "Wenn die Telefonnummer in deinem Adressbuch ist, kannst du den <b>%v Button benutzen</b>, um einen Kontakt zu senden.",
		enUK: "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		enUS: "If phone number is in your address book you can <b>use %v button</b> to send the contact.",
		esES: "Si el número está en tu agenda puedes <b>usar %v el botón</b> para enviar el contacto.",
		faIR: "اگر شماره تلفن در فهرست مخاطبین شما وجود دارد شما می توانید <b> با استفاده از این %v دکمه</b> تماس را ارسال نمایید.",
		frFR: "Si le numéro de téléphone est dans votre carnet d'adresses, vous pouvez <b>utiliser le bouton %v</b> pour envoyer le contact.",
		idID: "Jika nomor telepon ada di buku alamat Anda, Anda dapat <b>menggunakan tombol %v</b> untuk mengirim kontak.",
		itIT: "Se il numero e' nella tua rubrica, puoi <b> usare il pulsante %v</b> per inviare il contatto.",
		jaJP: "電話番号がアドレス帳にある場合は、<b>%vボタンを使用</b>して連絡先を送信できます。",
		koKR: "전화번호가 주소록에 있는 경우 <b>%v 버튼을 사용</b>하여 연락처를 보낼 수 있습니다.",
		plPL: "Jeśli numer telefonu znajduje się w książce adresowej, możesz <b>użyć przycisku %v</b>, aby wysłać kontakt.",
		ptBR: "Se o número de telefone estiver em sua agenda, você pode <b>usar o botão %v</b> para enviar o contato.",
		ruRU: "Если номер телефона есть в записной книжке <b>воспользуйтесь кнопкой %v</b> (скрепка) чтобы отправить контакт.",
		trTR: "Telefon numarası adres defterinizde varsa, kişiyi göndermek için <b>%v düğmesini kullanabilirsiniz</b>.",
		ukUA: "Якщо номер телефону є у вашій адресній книзі, ви можете <b>використати кнопку %v</b>, щоб надіслати контакт.",
		uzUZ: "Agar telefon raqami manzillar kitobingizda bo'lsa, kontaktni yuborish uchun <b>%v tugmasidan foydalanishingiz</b> mumkin.",
		zhCN: "如果电话号码在您的通讯录中，您可以<b>使用%v按钮</b>发送联系人。",
	},
	MESSAGE_TEXT_ABOUT_PHONE_NUMBER_FORMAT: {
		deDE: "Die Telefonnummer sollte dem internationalen Standard entsprechen:\n\t* Beginnend mit '+' gefolgt vom Ländercode (Deutschland +49)\n\t* Consist of numbers only\nExample: <b>+49</b><code>157123456</code>",
		enUK: "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		enUS: "The number should be in international standard:\n\t* Starts with '+' following by country code\n\t* Consist of numbers only\nExample: <b>+1</b><code>999012345678</code>",
		esES: "El número debe tener formato internacional estándar:\n\t* Empezar con '+' seguido del código del país\n\t* formado solo por números\nEjemplo: <b>+1</b><code>999012345678</code>",
		faIR: "شماره باید به صورت استاندارد بین المللی باشد\n\t* با '+' شروع شده و بدنبال آن کد کشور وارد شود\n\t* تنها شامل اعداد باشد\nمثال: <b>+1</b><code>999012345678</code>",
		frFR: "Le numéro doit être au format international:\n\t* Commence par '+' suivi du code du pays\n\t* Composé uniquement de chiffres\nExemple: <b>+33</b><code>612345678</code>",
		idID: "Nomor harus dalam standar internasional:\n\t* Dimulai dengan '+' diikuti oleh kode negara\n\t* Hanya terdiri dari angka\nContoh: <b>+62</b><code>812345678</code>",
		itIT: "Il numero deve essere in formato internazionale:\n\t* Inizia con '+' seguito dal codice del paese (Italia +39)\n\t* \nEsempio: <b>+39</b><code>34612345678</code>",
		jaJP: "番号は国際標準形式である必要があります:\n\t* '+'で始まり、国コードが続く\n\t* 数字のみで構成\n例: <b>+81</b><code>9012345678</code>",
		koKR: "번호는 국제 표준이어야 합니다:\n\t* '+'로 시작하여 국가 코드가 뒤따름\n\t* 숫자로만 구성\n예: <b>+82</b><code>1012345678</code>",
		plPL: "Numer powinien być w standardzie międzynarodowym:\n\t* Zaczyna się od '+' a następnie kod kraju\n\t* Składa się tylko z cyfr\nPrzykład: <b>+48</b><code>512345678</code>",
		ptBR: "O número deve estar no padrão internacional:\n\t* Começa com '+' seguido pelo código do país\n\t* Consiste apenas de números\nExemplo: <b>+55</b><code>11912345678</code>",
		ruRU: "Номер должен быть в международном формате:\n\t* Начинаться со знака '+' и кода страны\n\t* Состоять только из цифр\nПример: <b>+7</b><code>999012345678</code>",
		trTR: "Numara uluslararası standartta olmalıdır:\n\t* '+' ile başlayıp ülke kodu ile devam eder\n\t* Sadece rakamlardan oluşur\nÖrnek: <b>+90</b><code>5301234567</code>",
		ukUA: "Номер повинен бути в міжнародному форматі:\n\t* Починатися зі знаку '+' та коду країни\n\t* Складатися тільки з цифр\nПриклад: <b>+380</b><code>991234567</code>",
		uzUZ: "Raqam xalqaro standartda bo'lishi kerak:\n\t* '+' bilan boshlanib, mamlakat kodi bilan davom etadi\n\t* Faqat raqamlardan iborat\nMisol: <b>+998</b><code>901234567</code>",
		zhCN: "号码应符合国际标准:\n\t* 以'+'开头，后跟国家代码\n\t* 仅由数字组成\n示例: <b>+86</b><code>13912345678</code>",
	},
	MESSAGE_TEXT_THIS_NUMBER_WILL_BE_USED_TO_SEND_RECEIPT: {
		deDE: "Wir werden eine SMS an diese Nummer schicken:",
		enUK: "Will send an SMS to this number:",
		enUS: "Will send an SMS to this number:",
		esES: "Enviaremos una SMS a este número:",
		faIR: "یک پیام کوتاه به این شماره ارسال خواهد شد:",
		frFR: "Nous enverrons un SMS à ce numéro:",
		idID: "Akan mengirim SMS ke nomor ini:",
		itIT: "Invieremo un SMS a questo numero:",
		jaJP: "このSMSをこの番号に送信します:",
		koKR: "이 번호로 SMS를 보낼 것입니다:",
		plPL: "Wyślemy SMS na ten numer:",
		ptBR: "Enviaremos um SMS para este número:",
		ruRU: "На этот номер мы отправим SMS:",
		trTR: "Bu numaraya SMS gönderilecek:",
		ukUA: "Надішлемо SMS на цей номер:",
		uzUZ: "Bu raqamga SMS yuboriladi:",
		zhCN: "将向此号码发送短信:",
	},
	MESSAGE_TEXT_COUNTERPARTY_OWES_YOU_SINGLE_DEBT: {
		deDE: `<b>%v</b> schuldet dir <b>%v</b>.`,
		enUK: `<b>%v</b> owes to you <b>%v</b>.`,
		enUS: `<b>%v</b> owes to you <b>%v</b>.`,
		esES: `<b>%v</b> has prestado <b>%v</b>.`,
		faIR: `<b>%v</b> به شما بدهکار بوده <b>%v</b>.`,
		frFR: `<b>%v</b> vous doit <b>%v</b>.`,
		idID: `<b>%v</b> berhutang kepada Anda <b>%v</b>.`,
		itIT: `<b>%v</b> e' in debito di <b>%v</b>.`,
		jaJP: `<b>%v</b>はあなたに<b>%v</b>を借りています。`,
		koKR: `<b>%v</b>님이 당신에게 <b>%v</b>를 빚지고 있습니다.`,
		plPL: `<b>%v</b> jest ci winien <b>%v</b>.`,
		ptBR: `<b>%v</b> deve a você <b>%v</b>.`,
		ruRU: `<b>%v</b> одалживал(а) у вас <b>%v</b>.`,
		trTR: `<b>%v</b> size <b>%v</b> borçlu.`,
		ukUA: `<b>%v</b> заборгував вам <b>%v</b>.`,
		uzUZ: `<b>%v</b> sizga <b>%v</b> qarzdir.`,
		zhCN: `<b>%v</b> 欠您 <b>%v</b>。`,
	},
	MESSAGE_TEXT_YOU_OWE_TO_COUNTERPARTY_SINGLE_DEBT: {
		deDE: "Du schuldest <b>%v</b> <b>%v</b>.",
		enUK: "You owe to <b>%v</b> <b>%v</b>.",
		enUS: "You owe to <b>%v</b> <b>%v</b>.",
		esES: "Te ha prestado <b>%v</b> <b>%v</b>.",
		faIR: "شما بدهکار هستید به <b>%v</b> <b>%v</b>.",
		frFR: "Vous devez <b>%v</b> à <b>%v</b>.",
		idID: "Anda berhutang kepada <b>%v</b> <b>%v</b>.",
		itIT: `Tu devi dare a <b>%v</b> <b>%v</b>.`,
		jaJP: "あなたは<b>%v</b>に<b>%v</b>を借りています。",
		koKR: "당신은 <b>%v</b>에게 <b>%v</b>를 빚지고 있습니다.",
		plPL: "Jesteś winien <b>%v</b> <b>%v</b>.",
		ptBR: "Você deve a <b>%v</b> <b>%v</b>.",
		ruRU: `<b>%v</b> одалживал(а) вам <b>%v</b>.`,
		trTR: "<b>%v</b>'e <b>%v</b> borçlusunuz.",
		ukUA: "Ви заборгували <b>%v</b> <b>%v</b>.",
		uzUZ: "Siz <b>%v</b>ga <b>%v</b> qarzdorsiz.",
		zhCN: "您欠<b>%v</b> <b>%v</b>。",
	},
	MESSAGE_TEXT_IS_IT_RETURNED_IN_FULL: {

		deDE: `Wurde diese Schuld vollständig beglichen?

		<i>Falls nur teilweise, kann der Teilbetrag direkt eingegeben werden.</i>`,
		enUK: `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,
		enUS: `Has this debt been returned in full?

		<i>If partially you can enter amount right away.</i>`,

		esES: `¿Ha sido totalmente devuelta esta deuda?

		<i>si ha sido devuelta parcialmente puedes introducir el importe.</i>`,

		faIR: `آیا این بدهی بصورت کامل بازپرداخت شده است؟

		<i>اگر بخشی از بدهی پرداخت شده است شما میتوانید مبلغ را وارد کنید.</i>`,

		frFR: `Cette dette a-t-elle été remboursée intégralement?

		<i>Si partiellement, vous pouvez saisir le montant tout de suite.</i>`,

		idID: `Apakah hutang ini telah dikembalikan sepenuhnya?

		<i>Jika sebagian, Anda dapat memasukkan jumlah langsung.</i>`,

		itIT: `Il debito e' stato saldato?

		<i>Se la risposta e' NO puoi inserire l'ammontare da sottrarre, direttamente qui sotto.</i>`,

		jaJP: `この借金は全額返済されましたか？

		<i>部分的に返済された場合は、すぐに金額を入力できます。</i>`,

		koKR: `이 빚이 전액 상환되었습니까?

		<i>부분적으로 상환된 경우 금액을 바로 입력할 수 있습니다.</i>`,

		plPL: `Czy ten dług został spłacony w całości?

		<i>Jeśli częściowo, możesz od razu wprowadzić kwotę.</i>`,

		ptBR: `Esta dívida foi devolvida integralmente?

		<i>Se parcialmente, você pode inserir o valor imediatamente.</i>`,

		ruRU: `Возвращено полностью?

		<i>Если частично можете сразу написать сумму.</i>`,

		trTR: `Bu borç tamamen geri ödendi mi?

		<i>Kısmen ödendiyse, tutarı hemen girebilirsiniz.</i>`,

		ukUA: `Чи повернуто цей борг повністю?

		<i>Якщо частково, ви можете відразу ввести суму.</i>`,

		uzUZ: `Bu qarz to'liq qaytarildimi?

		<i>Agar qisman bo'lsa, miqdorni darhol kiritishingiz mumkin.</i>`,

		zhCN: `这笔债务是否已全额归还？

		<i>如果部分归还，您可以立即输入金额。</i>`,
	},
	MESSAGE_TEXT_PLEASE_HELP_MAKE_IT_BETTER: {
		deDE: `Das Programm ist <b>kostenlos</b>. Bitte <a href="https://debtstracker.io/en/help-us">hilf</a> es besser zu machen!`,
		enUK: `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		enUS: `This program is <b>free to use</b>. Please <a href="https://debtstracker.io/en/help-us">help</a> to make it better!`,
		esES: `Este programa es <b>gratis</b>. Por favor <a href="https://debtstracker.io/en/help-us">ayúdanos</a> a mejorarlo!`,
		faIR: `این برنامه <b>رایگان می باشد</b>.لطفاً <a href="https://debtstracker.io/">به ما کمک کنید</a>تا آنرا بهبود دهیم!`,
		frFR: `Ce programme est <b>gratuit</b>. S'il vous plaît <a href="https://debtstracker.io/en/help-us">aidez-nous</a> à l'améliorer!`,
		idID: `Program ini <b>gratis untuk digunakan</b>. Silakan <a href="https://debtstracker.io/en/help-us">bantu</a> untuk membuatnya lebih baik!`,
		itIT: `Questo programma e' <b> completamente gratis</b>. Per favore <a href="https://debtstracker.io/">aiuta</a> a migliorarlo!`,
		jaJP: `このプログラムは<b>無料で使用できます</b>。より良くするために<a href="https://debtstracker.io/en/help-us">ご協力</a>ください！`,
		koKR: `이 프로그램은 <b>무료로 사용</b>할 수 있습니다. 더 나은 프로그램을 만들기 위해 <a href="https://debtstracker.io/en/help-us">도와주세요</a>!`,
		plPL: `Ten program jest <b>darmowy</b>. Proszę <a href="https://debtstracker.io/en/help-us">pomóż</a> uczynić go lepszym!`,
		ptBR: `Este programa é <b>gratuito para uso</b>. Por favor, <a href="https://debtstracker.io/en/help-us">ajude</a> a torná-lo melhor!`,
		ruRU: `Эта программа <b>бесплатна</b>. <a href="https://debtstracker.io/ru/help-us">Помогите</a> сделать её лучше!`,
		trTR: `Bu program <b>ücretsiz kullanım</b> içindir. Lütfen daha iyi hale getirmek için <a href="https://debtstracker.io/en/help-us">yardım edin</a>!`,
		ukUA: `Ця програма <b>безкоштовна</b>. Будь ласка, <a href="https://debtstracker.io/en/help-us">допоможіть</a> зробити її кращою!`,
		uzUZ: `Bu dastur <b>bepul foydalanish</b> uchun. Iltimos, uni yaxshilash uchun <a href="https://debtstracker.io/en/help-us">yordam bering</a>!`,
		zhCN: `此程序<b>免费使用</b>。请<a href="https://debtstracker.io/en/help-us">帮助</a>我们使其变得更好！`,
	},
	BUTTON_TEXT_YOU_OWE_AMOUNT_TO_SOMEONE: {
		deDE: "%v | du schuldest: %v",
		enUK: "%v | you owe: %v",
		enUS: "%v | you owe: %v",
		esES: "%v | tú debes: %v",
		faIR: "%v | شما بدهکارید: %v",
		frFR: "%v | vous devez: %v",
		idID: "%v | Anda berhutang: %v",
		itIT: "%v | tu devi: %v",
		jaJP: "%v | あなたの借り: %v",
		koKR: "%v | 당신이 빚진 금액: %v",
		plPL: "%v | jesteś winien: %v",
		ptBR: "%v | você deve: %v",
		ruRU: "%v | вы должны: %v",
		trTR: "%v | borçlusunuz: %v",
		ukUA: "%v | ви винні: %v",
		uzUZ: "%v | siz qarzdorsiz: %v",
		zhCN: "%v | 您欠: %v",
	},
	BUTTON_TEXT_SOMEONE_OWES_TO_YOU_AMOUNT: {
		deDE: "%v | schuldet dir: %v",
		enUK: "%v | owes to you: %v",
		enUS: "%v | owes to you: %v",
		esES: "%v | te debe: %v",
		faIR: "%v | به شما بدهکار است: %v",
		frFR: "%v | vous doit: %v",
		idID: "%v | berhutang kepada Anda: %v",
		itIT: "%v | ti deve: %v",
		jaJP: "%v | あなたへの借り: %v",
		koKR: "%v | 당신에게 빚진 금액: %v",
		plPL: "%v | jest ci winien: %v",
		ptBR: "%v | deve a você: %v",
		ruRU: "%v | долг вам: %v",
		trTR: "%v | size borçlu: %v",
		ukUA: "%v | винен вам: %v",
		uzUZ: "%v | sizga qarzdor: %v",
		zhCN: "%v | 欠您: %v",
	},
	BUTTON_TEXT_DEBT_RETURNED_FULLY: {
		deDE: "Ja, vollständig",
		enUK: "Yes, fully",
		enUS: "Yes, fully",
		esES: "Sí, completamente",
		faIR: "بله، به صورت کامل",
		frFR: "Oui, entièrement",
		idID: "Ya, sepenuhnya",
		itIT: "Si, completamente",
		jaJP: "はい、完全に",
		koKR: "예, 완전히",
		plPL: "Tak, w całości",
		ptBR: "Sim, totalmente",
		ruRU: "Да, целиком",
		trTR: "Evet, tamamen",
		ukUA: "Так, повністю",
		uzUZ: "Ha, to'liq",
		zhCN: "是的，完全",
	},
	BUTTON_TEXT_DEBT_RETURNED_PARTIALLY: {
		deDE: "Nein, nur teilweise",
		enUK: "No, just partially",
		enUS: "No, just partially",
		esES: "No, solo parcialmente",
		faIR: "خیر، تنها قسمتی",
		frFR: "Non, seulement partiellement",
		idID: "Tidak, hanya sebagian",
		itIT: "No, parzialmente",
		jaJP: "いいえ、部分的に",
		koKR: "아니요, 부분적으로만",
		plPL: "Nie, tylko częściowo",
		ptBR: "Não, apenas parcialmente",
		ruRU: "Нет, только часть",
		trTR: "Hayır, sadece kısmen",
		ukUA: "Ні, лише частково",
		uzUZ: "Yo'q, faqat qisman",
		zhCN: "不，仅部分",
	},
	MESSAGE_TEXT_ATTEMPT_TO_USE_OWN_INVITE: {
		deDE: "Du solltest dich nicht selber einladen ;)",
		enUK: "You should not use your own invite ;)",
		enUS: "You should not use your own invite ;)",
		esES: "No deberías invitarte a ti mismo ;)",
		faIR: "نباید از دعوت خود استفاده کنید ;)",
		frFR: "Vous ne devriez pas utiliser votre propre invitation ;)",
		idID: "Anda tidak boleh menggunakan undangan Anda sendiri ;)",
		itIT: "Non dovresti usare il tuo codice invito con te stesso :)",
		jaJP: "自分の招待を使用すべきではありません ;)",
		koKR: "자신의 초대를 사용해서는 안 됩니다 ;)",
		plPL: "Nie powinieneś używać własnego zaproszenia ;)",
		ptBR: "Você não deve usar seu próprio convite ;)",
		ruRU: "Хорошая попытка пригласить самого себя ;)",
		trTR: "Kendi davetinizi kullanmamalısınız ;)",
		ukUA: "Ви не повинні використовувати власне запрошення ;)",
		uzUZ: "Siz o'z taklifingizdan foydalanmasligingiz kerak ;)",
		zhCN: "您不应该使用自己的邀请 ;)",
	},
	MESSAGE_TEXT_WELCOME_ONBOARDING_INVITE_ACCEPTED: {
		deDE: "Willkommen. Schön, dass du der Einladung gefolgt bist!",
		enUK: "Welcome and thanks for accepting the invite!",
		enUS: "Welcome and thanks for accepting the invite!",
		esES: "Bienvenido y gracias por aceptar la invitación",
		faIR: "خوش آمدید و سپاسگزاریم که دعوت را پذیرفتید!",
		frFR: "Bienvenue et merci d'avoir accepté l'invitation!",
		idID: "Selamat datang dan terima kasih telah menerima undangan!",
		itIT: "Benvenuto e grazie per aver accettato l'invito!",
		jaJP: "ようこそ、招待を受け入れていただきありがとうございます！",
		koKR: "환영합니다, 초대를 수락해 주셔서 감사합니다!",
		plPL: "Witamy i dziękujemy za przyjęcie zaproszenia!",
		ptBR: "Bem-vindo e obrigado por aceitar o convite!",
		ruRU: "Спасибо за то что воспользовались приглашением!",
		trTR: "Hoş geldiniz ve daveti kabul ettiğiniz için teşekkürler!",
		ukUA: "Ласкаво просимо та дякуємо за прийняття запрошення!",
		uzUZ: "Xush kelibsiz va taklifni qabul qilganingiz uchun rahmat!",
		zhCN: "欢迎并感谢您接受邀请！",
	},
	MESSAGE_TEXT_FOR_COUNTERPARTY_ONLY: {
		deDE: "Das darf nur %v.",
		enUK: "This action for %v only.",
		enUS: "This action for %v only.",
		esES: "Esta acción está disponible solo para%v.",
		faIR: "این عمل تنها برای %v می باشد.",
		frFR: "Cette action est uniquement pour %v.",
		idID: "Tindakan ini hanya untuk %v.",
		itIT: "Questa azione e' disponibile solo per %v.",
		jaJP: "このアクションは%vのみです。",
		koKR: "이 작업은 %v만을 위한 것입니다.",
		plPL: "Ta akcja tylko dla %v.",
		ptBR: "Esta ação é apenas para %v.",
		ruRU: "Это действие доступно только для %v.",
		trTR: "Bu işlem sadece %v için.",
		ukUA: "Ця дія тільки для %v.",
		uzUZ: "Bu harakat faqat %v uchun.",
		zhCN: "此操作仅适用于%v。",
	},
	BUTTON_TEXT_SEE_RECEIPT_DETAILS: {
		deDE: "Quittungsdetails anzeigen",
		enUK: "Show receipt details",
		enUS: "Show receipt details",
		esES: "Mostrar detalles",
		faIR: "جزئیات رسید را نشان بده",
		frFR: "Afficher les détails du reçu",
		idID: "Tampilkan detail tanda terima",
		itIT: "Mostra i dettagli del debito/credito",
		jaJP: "領収書の詳細を表示",
		koKR: "영수증 세부 정보 표시",
		plPL: "Pokaż szczegóły paragonu",
		ptBR: "Mostrar detalhes do recibo",
		ruRU: "Показать детали",
		trTR: "Makbuz detaylarını göster",
		ukUA: "Показати деталі квитанції",
		uzUZ: "Kvitansiya tafsilotlarini ko'rsatish",
		zhCN: "显示收据详情",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_EMAIL: {
		deDE: "Du hast gewählt, einen Freund per Mail einzuladen.",
		enUK: "You've selected to invite friend by email.",
		enUS: "You've selected to invite friend by email.",
		esES: "Has decidido invitar a un amigo por e-mail.",
		faIR: "شما انتخاب کردید که یک دوست را بوسیله ایمیل دعوت کنید.",
		frFR: "Vous avez choisi d'inviter un ami par e-mail.",
		idID: "Anda telah memilih untuk mengundang teman melalui email.",
		itIT: "Hai scelto di invitare l'amico tramite email.",
		jaJP: "友達をメールで招待することを選択しました。",
		koKR: "이메일로 친구를 초대하기로 선택하셨습니다.",
		plPL: "Wybrałeś zaproszenie znajomego przez e-mail.",
		ptBR: "Você selecionou convidar um amigo por e-mail.",
		ruRU: "Вы решили пригласить друга через email.",
		trTR: "Arkadaşınızı e-posta ile davet etmeyi seçtiniz.",
		ukUA: "Ви вибрали запросити друга електронною поштою.",
		uzUZ: "Siz do'stingizni elektron pochta orqali taklif qilishni tanladingiz.",
		zhCN: "您已选择通过电子邮件邀请朋友。",
	},
	MESSAGE_TEXT_YOU_SELECTED_INVITE_BY_SMS: {
		deDE: "Du hast gewählt, einen Freund per SMS einzuladen.",
		enUK: "You've selected to invite friend by SMS.",
		enUS: "You've selected to invite friend by SMS.",
		esES: "Has decidido invitar a un amigo por SMS.",
		faIR: "شما انتخاب کردید که یک دوست را بوسیله پیام کوتاه دعوت کنید",
		frFR: "Vous avez choisi d'inviter un ami par SMS.",
		idID: "Anda telah memilih untuk mengundang teman melalui SMS.",
		itIT: "Hai scelto di invitare l'amico tramite SMS.",
		jaJP: "友達をSMSで招待することを選択しました。",
		koKR: "SMS로 친구를 초대하기로 선택하셨습니다.",
		plPL: "Wybrałeś zaproszenie znajomego przez SMS.",
		ptBR: "Você selecionou convidar um amigo por SMS.",
		ruRU: "Вы решили пригласить друга через SMS.",
		trTR: "Arkadaşınızı SMS ile davet etmeyi seçtiniz.",
		ukUA: "Ви вибрали запросити друга через SMS.",
		uzUZ: "Siz do'stingizni SMS orqali taklif qilishni tanladingiz.",
		zhCN: "您已选择通过短信邀请朋友。",
	},
	MESSAGE_TEXT_ABOUT_INVITES: {
		deDE: "Wie möchtest du den Code weitergeben?",

		enUK: `How do you want to pass the invite code?`,

		enUS: `How do you want to pass the invite code?`,

		esES: `¿Cómo quieres enviarle el código?`,

		faIR: `آیا میخواهید کد دعوت را ارسال کنید؟`,

		frFR: `Comment voulez-vous transmettre le code d'invitation?`,

		idID: `Bagaimana Anda ingin meneruskan kode undangan?`,

		itIT: `Come vuoi inviargli il codice invito?`,

		jaJP: `招待コードをどのように渡しますか？`,

		koKR: `초대 코드를 어떻게 전달하시겠습니까?`,

		plPL: `Jak chcesz przekazać kod zaproszenia?`,

		ptBR: `Como você deseja passar o código de convite?`,

		ruRU: `Как вы хотите передать код приглашение?`,

		trTR: `Davet kodunu nasıl iletmek istiyorsunuz?`,

		ukUA: `Як ви хочете передати код запрошення?`,

		uzUZ: `Taklif kodini qanday o'tkazmoqchisiz?`,

		zhCN: `您想如何传递邀请码？`,
	},
	MESSAGE_TEXT_USER_BLOCKED_TRANSFER_NOTIFICATIONS_BY: {
		deDE: "%v hat Erinnerungen über folgendes Anliegen blockiert: %v",
		enUK: "%v blocked reminders about transactions by: %v",
		enUS: "%v blocked reminders about transactions by: %v",
		esES: "%v ha bloqueado las notificaciones de las transacciones por: %v",
		faIR: "%v یادآور تراکنش مسدود شده است بوسیله ی: %v",
		frFR: "%v a bloqué les rappels concernant les transactions par: %v",
		idID: "%v memblokir pengingat tentang transaksi oleh: %v",
		itIT: "%v bloccato promemoria riguardo la transazione da: %v.",
		jaJP: "%vは%vによるトランザクションに関するリマインダーをブロックしました",
		koKR: "%v님이 %v의 거래에 대한 알림을 차단했습니다",
		plPL: "%v zablokował przypomnienia o transakcjach przez: %v",
		ptBR: "%v bloqueou lembretes sobre transações por: %v",
		ruRU: "%v заблокировал получение оповешений о транзакиях через: %v.",
		trTR: "%v, %v tarafından yapılan işlemler hakkındaki hatırlatıcıları engelledi",
		ukUA: "%v заблокував нагадування про транзакції від: %v",
		uzUZ: "%v %v tomonidan tranzaksiyalar haqida eslatmalarni blokladi",
		zhCN: "%v已阻止来自%v的交易提醒",
	},
	COMMAND_TEXT_WAIT_A_SECOND: {
		deDE: "Warte eine Sekunde...",
		enUK: "Wait a second...",
		enUS: "Wait a second...",
		esES: "Un segundo...",
		faIR: "یک ثانیه صبر کنید ...",
		frFR: "Attendez une seconde...",
		idID: "Tunggu sebentar...",
		itIT: "Solo un attimo...",
		jaJP: "ちょっと待って...",
		koKR: "잠시만 기다려주세요...",
		plPL: "Poczekaj chwilę...",
		ptBR: "Espere um segundo...",
		ruRU: "Секундочку...",
		trTR: "Bir saniye bekleyin...",
		ukUA: "Зачекайте секунду...",
		uzUZ: "Bir soniya kuting...",
		zhCN: "稍等一下...",
	},
	HTML_USING_TELEGRAM: {
		deDE: "benutze Telegram messenger",
		enUK: "using Telegram messenger",
		enUS: "using Telegram messenger",
		esES: "Usando Telegram",
		faIR: "استفاده از پیام رسان تلگرام",
		frFR: "utilisant Telegram messenger",
		idID: "menggunakan Telegram messenger",
		itIT: "usa Telegram",
		jaJP: "Telegramメッセンジャーを使用",
		koKR: "텔레그램 메신저 사용",
		plPL: "używając komunikatora Telegram",
		ptBR: "usando o Telegram messenger",
		ruRU: "используя Telegram мессенджер",
		trTR: "Telegram messenger kullanarak",
		ukUA: "використовуючи Telegram messenger",
		uzUZ: "Telegram messenjeridan foydalanish",
		zhCN: "使用Telegram信使",
	},
	COMMAND_TEXT_ACCEPT: {
		deDE: "Akzeptieren",
		enUK: "Accept",
		enUS: "Accept",
		esES: "Aceptar",
		faIR: "قبول",
		frFR: "Accepter",
		idID: "Terima",
		itIT: "Accetta",
		jaJP: "承認",
		koKR: "수락",
		plPL: "Akceptuj",
		ptBR: "Aceitar",
		ruRU: "Согласиться",
		trTR: "Kabul et",
		ukUA: "Прийняти",
		uzUZ: "Qabul qilish",
		zhCN: "接受",
	},
	//BUTTON_TEXT_ACCEPT_TRANSFER_USING_TELEGRAM:{
	//	enUK: "Accept ",
	//	esES: "Confirmar ",

	//	faIR: "قبول ",

	//  itIT:   "Accetta",
	//	ruRU:  "Подтвердить ",

	//},
	//BUTTON_TEXT_DECLINE_TRANSFER_USING_TELEGRAM:{
	//	enUK: "Decline (using Telegram messenger)",
	//	esES: "Rechazar (usandoTelegram)",

	//	faIR: "رد ( استفاده از پیام رسان تلگرام)",

	//  itIT:   "Rifiuta (usando Telegram)",
	//	ruRU:  "Отказаться (используя Telegram)",

	//},
	COMMAND_TEXT_DECLINE: {
		deDE: "Ablehnen",
		enUK: "Decline",
		esES: "Rechazar",
		faIR: "رد",
		itIT: "Rifiuta",
		ruRU: "Отклонить",
	},
	FamilyMember: {
		deDE: "Familienmitglied",
		enUK: "Family member",
		enUS: "Family member", // Placeholder
		esES: "Miembro de la familia",
		faIR: "عضو خانواده",
		frFR: "Membre de la famille", // Placeholder
		idID: "Anggota keluarga",     // Placeholder
		itIT: "Membro della famiglia",
		jaJP: "家族の一員",             // Placeholder
		koKR: "가족 구성원",            // Placeholder
		plPL: "Członek rodziny",   // Placeholder
		ptBR: "Membro da família", // Placeholder
		ruRU: "Член семьи",
		trTR: "Aile üyesi",  // Placeholder
		ukUA: "Член родини", // Placeholder
		uzUZ: "Oila aʼzosi", // Placeholder
		zhCN: "家庭成员",        // Placeholder
	},
	UserHasNotJoinedSpaceYet: {
		deDE: "Dieser Kontakt ist diesem Bereich noch nicht beigetreten.",
		enUK: "This contact has not joined this space yet.",
		enUS: "This contact has not joined this space yet.",
		esES: "Este contacto aún no se ha unido a este espacio.",
		faIR: "این مخاطب هنوز به این فضا نپیوسته است.",
		frFR: "Ce contact n'a pas encore rejoint cet espace.",
		idID: "Kontak ini belum bergabung dengan ruang ini.",
		itIT: "Questo contatto non si è ancora unito a questo spazio.",
		jaJP: "この連絡先はまだこのスペースに参加していません。",
		koKR: "이 연락처는 아직 이 공간에 참여하지 않았습니다.",
		plPL: "Ten kontakt nie dołączył jeszcze do tej przestrzeni.",
		ptBR: "Este contato ainda não entrou neste espaço.",
		ptPT: "Este contacto ainda não aderiu a este espaço.",
		ruRU: "Eще не присоединился к этому пространству.",
		trTR: "Bu kişi henüz bu alana katılmadı.",
		ukUA: "Цей контакт ще не приєднався до цього простору.",
		uzUZ: "Bu kontakt hali bu joyga qoʻshilmadi.",
		zhCN: "该联系人尚未加入此空间。",
	},
	UserHasNotJoinedFamilySpaceYet: {
		deDE: "Ist diesem Familienbereich noch nicht beigetreten.",
		enUK: "Has not joined this family space yet.",
		enUS: "Has not joined this family space yet.",
		esES: "Aún no se ha unido a este espacio familiar.",
		faIR: "هنوز به این فضای خانوادگی نپیوسته است.",
		frFR: "N'a pas encore rejoint cet espace familial.",
		idID: "Belum bergabung dengan ruang keluarga ini.",
		itIT: "Non si è ancora unito a questo spazio familiare.",
		jaJP: "このファミリースペースにまだ参加していません。",
		koKR: "아직 이 가족 공간에 참여하지 않았습니다.",
		plPL: "Nie dołączył jeszcze do tej przestrzeni rodzinnej.",
		ptBR: "Ainda não entrou neste espaço familiar.",
		ptPT: "Ainda não aderiu a este espaço familiar.",
		ruRU: "Eще не присоединился к этому семейному пространству.",
		trTR: "Henüz bu aile alanına katılmadı.",
		ukUA: "Ще не приєднався до цього сімейного простору.",
		uzUZ: "Hali bu oila bo'shlig'iga qo'shilmagan.",
		zhCN: "尚未加入该家庭空间。",
	},
	BtnSendInviteByTelegram: {
		deDE: "Einladung über Telegram senden",
		enUK: "Send invite over Telegram",
		enUS: "Send invite over Telegram",
		esES: "Enviar invitación por Telegram",
		faIR: "ارسال دعوتنامه از طریق تلگرام",
		frFR: "Envoyer une invitation via Telegram",
		idID: "Kirim undangan melalui Telegram",
		itIT: "Invia invito tramite Telegram",
		jaJP: "Telegramで招待を送る",
		koKR: "텔레그램으로 초대 보내기",
		plPL: "Wyślij zaproszenie przez Telegram",
		ptBR: "Enviar convite pelo Telegram",
		ruRU: "Отправить приглашение через Telegram",
		trTR: "Telegram üzerinden davetiye gönder",
		ukUA: "Надіслати запрошення через Telegram",
		uzUZ: "Telegram orqali taklif yuboring",
		zhCN: "通过Telegram发送邀请",
	},
	BtnTextAcceptInvite: {
		deDE: "Akzeptiere Einladung",
		enUK: "Accept invite",
		enUS: "Accept invite",
		esES: "Aceptar la invitación",
		faIR: "قبول دعوت",
		frFR: "Accepter l'invitation",
		idID: "Terima undangan",
		itIT: "Accetta l'invito",
		jaJP: "招待を承認",
		koKR: "초대 수락",
		plPL: "Zaakceptuj zaproszenie",
		ptBR: "Aceitar convite",
		ruRU: "Принять приглашение",
		trTR: "Daveti kabul et",
		ukUA: "Прийняти запрошення",
		uzUZ: "Taklifni qabul qilish",
		zhCN: "接受邀请",
	},
	BtnTextDeclineInvite: {
		deDE: "Ablehnen Einladung",
		enUK: "Decline invite",
		enUS: "Decline invite",
		esES: "Rechazar la invitación",
		faIR: "رد دعوت",
		frFR: "Décliner l'invitation",
		idID: "Tolak undangan",
		itIT: "Rifiuta l'invito",
		jaJP: "招待を拒否",
		koKR: "초대 거절",
		plPL: "Odrzuć zaproszenie",
		ptBR: "Recusar convite",
		ruRU: "Отклонить приглашение",
		trTR: "Daveti reddet",
		ukUA: "Відхилити запрошення",
		uzUZ: "Taklifni rad etish",
		zhCN: "拒绝邀请",
	},
	COMMAND_TEXT_VIEW_RECEIPT_DETAILS: {
		deDE: "Quittungsdetails anzeigen",
		enUK: "See receipt details",
		enUS: "See receipt details",
		esES: "Ver el recibo",
		faIR: "دیدن جزئیات رسید",
		frFR: "Voir les détails du reçu",
		idID: "Lihat detail tanda terima",
		itIT: "Vedi dettagli",
		jaJP: "領収書の詳細を見る",
		koKR: "영수증 세부 정보 보기",
		plPL: "Zobacz szczegóły paragonu",
		ptBR: "Ver detalhes do recibo",
		ruRU: "Посмотреть квитанцию",
		trTR: "Makbuz detaylarını görüntüle",
		ukUA: "Переглянути квитанцію",
		uzUZ: "Kvitansiya tafsilotlarini ko'rish",
		zhCN: "查看收据详情",
	},
	COMMAND_TEXT_OTHER_WAYS_TO_SEND_INVITE: {
		deDE: "Andere Wege, eine Einladung zu senden",
		enUK: "Other ways to send an invite",
		enUS: "Other ways to send an invite",
		esES: "Otras maneras para enviar la invitación",
		faIR: "سایر راههای ارسال دعوت",
		frFR: "Autres façons d'envoyer une invitation",
		idID: "Cara lain untuk mengirim undangan",
		itIT: "Altri modi per inviare un invito",
		jaJP: "招待状を送信する他の方法",
		koKR: "초대장을 보내는 다른 방법",
		plPL: "Inne sposoby wysyłania zaproszenia",
		ptBR: "Outras formas de enviar um convite",
		ruRU: "Другие способы отправить приглашение",
		trTR: "Davetiye göndermenin diğer yolları",
		ukUA: "Інші способи надіслати запрошення",
		uzUZ: "Taklifnoma yuborishning boshqa usullari",
		zhCN: "发送邀请的其他方式",
	},
	COMMAND_TEXT_SEND_MY_PHONE_NUMBER: {
		deDE: "meine Telefonnummer senden",
		enUK: "Send my phone number",
		enUS: "Send my phone number",
		esES: "Enviar mi número",
		faIR: "شماره تلفن مرا ارسال کنید",
		frFR: "Envoyer mon numéro de téléphone",
		idID: "Kirim nomor telepon saya",
		itIT: "Invia il mio numero",
		jaJP: "私の電話番号を送信する",
		koKR: "내 전화번호 보내기",
		plPL: "Wyślij mój numer telefonu",
		ptBR: "Enviar meu número de telefone",
		ruRU: "Отправить мой номер",
		trTR: "Telefon numaramı gönder",
		ukUA: "Надіслати мій номер телефону",
		uzUZ: "Telefon raqamimni yuborish",
		zhCN: "发送我的电话号码",
	},
	COMMAND_TEXT_SEND_BY_EMAIL: {
		deDE: "per Mail",
		enUK: "By Email",
		enUS: "By Email",
		esES: "Vía e-mail",
		faIR: "بوسیله ی ایمیل",
		frFR: "Par e-mail",
		idID: "Melalui Email",
		itIT: "Tramite email",
		jaJP: "メールで",
		koKR: "이메일로",
		plPL: "Przez e-mail",
		ptBR: "Por e-mail",
		ruRU: "Через Email",
		trTR: "E-posta ile",
		ukUA: "Через електронну пошту",
		uzUZ: "Email orqali",
		zhCN: "通过电子邮件",
	},
	COMMAND_TEXT_SEND_BY_SMS: {
		deDE: "per SMS",
		enUK: "By SMS",
		enUS: "By SMS",
		esES: "Vía SMS",
		faIR: "بوسیله پیام کوتاه",
		frFR: "Par SMS",
		idID: "Melalui SMS",
		itIT: "Tramite SMS",
		jaJP: "SMSで",
		koKR: "SMS로",
		plPL: "Przez SMS",
		ptBR: "Por SMS",
		ruRU: "Через SMS",
		trTR: "SMS ile",
		ukUA: "Через SMS",
		uzUZ: "SMS orqali",
		zhCN: "通过短信",
	},
	COMMAND_TEXT_INVITE_BY_TELEGRAM: {
		deDE: "Einladen per Telegram",
		enUK: "Invite By Telegram",
		enUS: "Invite By Telegram",
		esES: "Invitar vía Telegram",
		faIR: "دعوت با تلگرام",
		frFR: "Inviter par Telegram",
		idID: "Undang Melalui Telegram",
		itIT: "Tramite Telegram",
		jaJP: "Telegramで招待する",
		koKR: "텔레그램으로 초대하기",
		plPL: "Zaproś przez Telegram",
		ptBR: "Convidar pelo Telegram",
		ruRU: "Пригласить через Telegram",
		trTR: "Telegram ile davet et",
		ukUA: "Запросити через Telegram",
		uzUZ: "Telegram orqali taklif qilish",
		zhCN: "通过Telegram邀请",
	},
	MESSAGE_TEXT_INVITE_CREATED: {

		deDE: `Wir haben deinen Freund einen Code geschickt. (#%v)

Sobald dein Freund die Einladung akzeptiert hat, könnt ihr das Geld, was ihr euch teit, mit Leichtigkeit managen.`,

		enUK: `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		enUS: `We've sent an invite code to your friend. (#%v)

Once your friend accepts invitation you'll share balance & transfers between you to make sure you both are on the same page with minimum efforts.`,

		esES: `Hemos enviado el código de la invitación a tu amigo. (#%v)

Cuando tu amigo accepte la invitación vaís a tener transacciones y balance en común (solo entre vosotros). Todo eso os ayuda minimizar los esfuerzos para controlar la cuenta.`,

		faIR: `ما برای دوست شما یک  پیام دعوت ارسال کردیم. (#%v)

وقتی دوست شما دعوت را بپذیرد شما تراز و مبادلات بین خود را به اشتراک می گذارید تا با کمترین تلاش از درک مشترک میان خود اطمینان حاصل کنید. `,

		frFR: `Nous avons envoyé un code d'invitation à votre ami. (#%v)

Une fois que votre ami accepte l'invitation, vous partagerez le solde et les transferts entre vous pour vous assurer que vous êtes tous les deux sur la même page avec un minimum d'efforts.`,

		idID: `Kami telah mengirimkan kode undangan kepada teman Anda. (#%v)

Setelah teman Anda menerima undangan, Anda akan berbagi saldo & transfer antara Anda untuk memastikan Anda berdua berada di halaman yang sama dengan upaya minimal.`,

		itIT: `Abbiamo inviato il codice invito al tuo amico. (#%v)

Una volta che il tuo amico accetta l'invito potrete condividere i bilanci ed i trasferimenti con il minimo sforzo.`,

		jaJP: `友達に招待コードを送信しました。(#%v)

友達が招待を受け入れると、最小限の労力で両方が同じページにいることを確認するために、あなたの間で残高と転送を共有します。`,

		koKR: `친구에게 초대 코드를 보냈습니다. (#%v)

친구가 초대를 수락하면 최소한의 노력으로 두 사람이 같은 페이지에 있는지 확인하기 위해 잔액과 이체를 공유하게 됩니다.`,

		plPL: `Wysłaliśmy kod zaproszenia do Twojego znajomego. (#%v)

Gdy Twój znajomy zaakceptuje zaproszenie, będziecie dzielić saldo i przelewy między sobą, aby upewnić się, że oboje jesteście na tej samej stronie przy minimalnym wysiłku.`,

		ptBR: `Enviamos um código de convite para seu amigo. (#%v)

Quando seu amigo aceitar o convite, vocês compartilharão saldo e transferências entre si para garantir que ambos estejam na mesma página com o mínimo de esforço.`,

		ruRU: `Мы отправили код приглашения на указынный вами адрес. (#%v)

Когда ваш друг потдвердит приглашение у вас будут общий баланс и транзакции (только между вами). Это поможет вам минимизировать усилия по ведению учёта.`,

		trTR: `Arkadaşınıza bir davet kodu gönderdik. (#%v)

Arkadaşınız daveti kabul ettiğinde, minimum çabayla ikinizin de aynı sayfada olduğundan emin olmak için aranızda bakiye ve transferleri paylaşacaksınız.`,

		ukUA: `Ми надіслали код запрошення вашому другу. (#%v)

Коли ваш друг прийме запрошення, ви будете ділитися балансом і переказами між собою, щоб переконатися, що ви обидва на одній сторінці з мінімальними зусиллями.`,

		uzUZ: `Do'stingizga taklifnoma kodini yubordik. (#%v)

Do'stingiz taklifni qabul qilgandan so'ng, minimal kuch bilan ikkalangiz ham bir xil sahifada ekanligingizga ishonch hosil qilish uchun o'zaro balans va o'tkazmalarni almashishingiz mumkin.`,

		zhCN: `我们已向您的朋友发送邀请码。(#%v)

一旦您的朋友接受邀请，您将共享余额和转账，以确保您们双方以最少的努力保持一致。`,
	},
	MESSAGE_TEXT_INVITE_BY_EMAIL: {
		deDE: "Bitte gib die eMail Adresse deines Freundes an, wohin wir den Code schicken:",
		enUK: "Please enter emaill address of your friend where we should send an invite code.",
		enUS: "Please enter email address of your friend where we should send an invite code.",
		esES: "Por favor, introduce el e-maill de tu amigo al cual enviaremos el código de la invitación.",
		faIR: "لطفاً آدرس ایمیل دوست خود را وارد کنید تا ما از آن طریق کد دعوت را ارسال نماییم.",
		frFR: "Veuillez entrer l'adresse e-mail de votre ami où nous devrions envoyer un code d'invitation.",
		idID: "Silakan masukkan alamat email teman Anda di mana kami harus mengirimkan kode undangan.",
		itIT: "Inserisci l'email dell'amico al quale inviare il codice invito.",
		jaJP: "招待コードを送信する友達のメールアドレスを入力してください。",
		koKR: "초대 코드를 보낼 친구의 이메일 주소를 입력하세요.",
		plPL: "Proszę podać adres e-mail znajomego, na który powinniśmy wysłać kod zaproszenia.",
		ptBR: "Por favor, insira o endereço de e-mail do seu amigo para onde devemos enviar um código de convite.",
		ruRU: "Введите email вашего друга на который мы отправим код приглашения.",
		trTR: "Lütfen davet kodunu göndermemiz gereken arkadaşınızın e-posta adresini girin.",
		ukUA: "Будь ласка, введіть електронну адресу вашого друга, куди ми повинні надіслати код запрошення.",
		uzUZ: "Iltimos, taklifnoma kodini yuborishimiz kerak bo'lgan do'stingizning elektron pochta manzilini kiriting.",
		zhCN: "请输入您朋友的电子邮件地址，我们将向其发送邀请码。",
	},
	MESSAGE_TEXT_INVITE_ASK_EMAIL_FOR_RECEIPT: {
		deDE: "Bitte gib die eMail Adresse deines Freundes (%v), wohin wir die Quittung schicken:",
		enUK: "Please enter emaill address of your friend (%v) where we should send the receipt.",
		enUS: "Please enter email address of your friend (%v) where we should send the receipt.",
		esES: "Introduce el e-maill de tu amigo (%v) al cual enviaremos el recibo.",
		faIR: "لطفا آدرس ایمیل دوست خود را وارد کنید (%v) تا ما از آن طریق کد دعوت را ارسال نماییم.",
		frFR: "Veuillez entrer l'adresse e-mail de votre ami (%v) où nous devrions envoyer le reçu.",
		idID: "Silakan masukkan alamat email teman Anda (%v) di mana kami harus mengirimkan tanda terima.",
		itIT: "Inserisci l'email del tuo amico (%v) alla quale potremo inviare il debito/credito",
		jaJP: "領収書を送信する友達 (%v) のメールアドレスを入力してください。",
		koKR: "영수증을 보낼 친구 (%v)의 이메일 주소를 입력하세요.",
		plPL: "Proszę podać adres e-mail znajomego (%v), na który powinniśmy wysłać paragon.",
		ptBR: "Por favor, insira o endereço de e-mail do seu amigo (%v) para onde devemos enviar o recibo.",
		ruRU: "Введите email вашего друга (%v) на который мы отправим квитанцию подтверждения.",
		trTR: "Lütfen makbuzu göndermemiz gereken arkadaşınızın (%v) e-posta adresini girin.",
		ukUA: "Будь ласка, введіть електронну адресу вашого друга (%v), куди ми повинні надіслати квитанцію.",
		uzUZ: "Iltimos, kvitansiyani yuborishimiz kerak bo'lgan do'stingizning (%v) elektron pochta manzilini kiriting.",
		zhCN: "请输入您朋友 (%v) 的电子邮件地址，我们将向其发送收据。",
	},
	MESSAGE_TEXT_INVITE_BY_SMS: {
		deDE: "Bitte gib die Telefonnummer deines Freundes an oder teile einen Kontakt, wohin wir den Code schicken:",
		enUK: "Please enter number of yoru friend where we'll send an invite.",
		enUS: "Please enter number of your friend where we'll send an invite.",
		esES: "Por favor, introduce el número del teléfono de tu amigo al cual enviaremos el código de la invitación.",
		faIR: "لطفا شماره دوستان را که می خواهید برای او کد دعوت بفرستیم از لیست مخاطبین به اشتراک گذاشته یا آن را وارد کنید.",
		frFR: "Veuillez entrer le numéro de votre ami où nous enverrons une invitation.",
		idID: "Silakan masukkan nomor teman Anda di mana kami akan mengirimkan undangan.",
		itIT: "Condividi il contatto o inserisci il numero di telefono del tuo amico al quale invieremo il codice invito.",
		jaJP: "招待状を送信する友達の番号を入力してください。",
		koKR: "초대장을 보낼 친구의 번호를 입력하세요.",
		plPL: "Proszę podać numer telefonu znajomego, na który wyślemy zaproszenie.",
		ptBR: "Por favor, insira o número do seu amigo para onde enviaremos um convite.",
		ruRU: "Введите номер телефона вашего друга на который мы отправим код приглашения.",
		trTR: "Lütfen davet göndereceğimiz arkadaşınızın numarasını girin.",
		ukUA: "Будь ласка, введіть номер вашого друга, куди ми надішлемо запрошення.",
		uzUZ: "Iltimos, taklifnoma yuboradigan do'stingizning raqamini kiriting.",
		zhCN: "请输入我们将发送邀请的朋友的号码。",
	},
	MESSAGE_TEXT_INVITE_BY_TELEGRAM: {
		deDE: "Bitte wähle einen Kontakt, welchen du den Code schicken willst:",
		enUK: "Please share a contact of your friend you wish to send an invite code:",
		enUS: "Please share a contact of your friend you wish to send an invite code:",
		esES: "Por favor, comparte el contacto de tu amigo al cual quieres enviar el código de la invitación.",
		faIR: "لطفا اطلاعات تماس دوستتان را که میخواهید برایشان کد دعوت ارسال شود به اشتراک بگذارید.",
		frFR: "Veuillez partager un contact de votre ami à qui vous souhaitez envoyer un code d'invitation :",
		idID: "Silakan bagikan kontak teman Anda yang ingin Anda kirimi kode undangan:",
		itIT: "Condividi il contatto di un amico al quale desideri inviare il codice invito.",
		jaJP: "招待コードを送信したい友達の連絡先を共有してください：",
		koKR: "초대 코드를 보내고 싶은 친구의 연락처를 공유하세요:",
		plPL: "Proszę udostępnić kontakt znajomego, któremu chcesz wysłać kod zaproszenia:",
		ptBR: "Por favor, compartilhe um contato do seu amigo para quem deseja enviar um código de convite:",
		ruRU: "Вставьте в чат контакт вашего друга которому вы хотите отправить приглашение.",
		trTR: "Lütfen davet kodu göndermek istediğiniz arkadaşınızın bir kişisini paylaşın:",
		ukUA: "Будь ласка, поділіться контактом вашого друга, якому ви бажаєте надіслати код запрошення:",
		uzUZ: "Iltimos, taklifnoma kodini yubormoqchi bo'lgan do'stingizning kontaktini ulashing:",
		zhCN: "请分享您希望发送邀请码的朋友的联系人：",
	},
	MESSAGE_TEXT_INVALID_EMAIL: {
		deDE: "Ungültige eMail Adresse. Versuche es erneut oder gehe ins Haupt /menu",
		enUK: "Invalid email. Check and try it again? /menu",
		enUS: "Invalid email. Check and try it again? /menu",
		esES: "Email incorrecto. ¿Comprobarlo e intentalo de nuevo? /menú",
		faIR: "ایمیل غیر معتبر می باشد. آیا بررسی نموده، دوباره سعی می کنید؟ /منو",
		frFR: "Email invalide. Vérifiez et réessayez ? /menu",
		idID: "Email tidak valid. Periksa dan coba lagi? /menu",
		itIT: "Email scorretta. Controlla e riprova. /menu",
		jaJP: "無効なメール。確認して再試行しますか？ /menu",
		koKR: "잘못된 이메일입니다. 확인하고 다시 시도하시겠습니까? /menu",
		plPL: "Nieprawidłowy email. Sprawdź i spróbuj ponownie? /menu",
		ptBR: "Email inválido. Verificar e tentar novamente? /menu",
		ruRU: "Неверный email. Проверить и попробовать ещё раз? /menu",
		trTR: "Geçersiz e-posta. Kontrol edip tekrar deneyin? /menu",
		ukUA: "Недійсна електронна пошта. Перевірити та спробувати ще раз? /menu",
		uzUZ: "Yaroqsiz elektron pochta. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		zhCN: "无效的电子邮件。检查并重试？ /menu",
	},
	MESSAGE_TEXT_INVALID_YEAR: {
		deDE: "Ungültiges Jahr. Der Jahresangabe sollte 2 oder 4 Ziffern sein (<i>z.B. 2016 or 16</i>).",
		enUK: "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		enUS: "Invalid year. The year part should be 2 or 4 numbers (<i>e.g. 2016 or 16</i>).",
		esES: "Año incorrecto. El año tiene que constar de 2 o 4 números (<i>ejemplo 2016 o 16</i>).",
		faIR: "سال غیرمعتبر می باشد. سال باید به صورت 2 یا 4 رقمی وارد شود (<i>برای مثال 16 یا 2016</i>).",
		frFR: "Année invalide. L'année doit comporter 2 ou 4 chiffres (<i>par exemple 2016 ou 16</i>).",
		idID: "Tahun tidak valid. Bagian tahun harus berupa 2 atau 4 angka (<i>misalnya 2016 atau 16</i>).",
		itIT: "Anno scorretto. L'anno dev;essere composta da 2 o 4 numeri (<i>esempio 2017 oppure 17</i>)",
		jaJP: "無効な年。年の部分は2桁または4桁の数字である必要があります (<i>例：2016または16</i>)。",
		koKR: "잘못된 연도입니다. 연도 부분은 2자리 또는 4자리 숫자여야 합니다 (<i>예: 2016 또는 16</i>).",
		plPL: "Nieprawidłowy rok. Część roku powinna składać się z 2 lub 4 cyfr (<i>np. 2016 lub 16</i>).",
		ptBR: "Ano inválido. A parte do ano deve ter 2 ou 4 números (<i>por exemplo, 2016 ou 16</i>).",
		ruRU: "Неверный год. Год должен быть 2 или 4 цифры (<i>например 2016 или 16)</i>).",
		trTR: "Geçersiz yıl. Yıl kısmı 2 veya 4 rakam olmalıdır (<i>örneğin 2016 veya 16</i>).",
		ukUA: "Недійсний рік. Частина року повинна складатися з 2 або 4 цифр (<i>наприклад, 2016 або 16</i>).",
		uzUZ: "Yaroqsiz yil. Yil qismi 2 yoki 4 raqamdan iborat bo'lishi kerak (<i>masalan, 2016 yoki 16</i>).",
		zhCN: "无效的年份。年份部分应为2位或4位数字 (<i>例如2016或16</i>)。",
	},
	MESSAGE_TEXT_INVALID_MONTH: {
		deDE: "Ungültiger Monat. Der Monatsangabe sollte eine Ganzzahl zwischen 1 und 12 sein.",
		enUK: "Invalid month. Month should be an integer from 1 to 12.",
		enUS: "Invalid month. Month should be an integer from 1 to 12.",
		esES: "El mes es incorrecto. El mes hay que introducirlo del 1 al 12.",
		faIR: "ماه غیر معتبر می باشد. ماه باید به صورت عددی صحیح بین 1 تا 12 باشد.",
		frFR: "Mois invalide. Le mois doit être un nombre entier de 1 à 12.",
		idID: "Bulan tidak valid. Bulan harus berupa bilangan bulat dari 1 hingga 12.",
		itIT: "Mese scorretto. Il mese dovrebbe essere un numero da 1 a 12.",
		jaJP: "無効な月。月は1から12までの整数である必要があります。",
		koKR: "잘못된 월입니다. 월은 1에서 12 사이의 정수여야 합니다.",
		plPL: "Nieprawidłowy miesiąc. Miesiąc powinien być liczbą całkowitą od 1 do 12.",
		ptBR: "Mês inválido. O mês deve ser um número inteiro de 1 a 12.",
		ruRU: "Неверный месяц. Месяц должен быть задан целым числом от 1 до 12.",
		trTR: "Geçersiz ay. Ay, 1'den 12'ye kadar bir tam sayı olmalıdır.",
		ukUA: "Недійсний місяць. Місяць повинен бути цілим числом від 1 до 12.",
		uzUZ: "Yaroqsiz oy. Oy 1 dan 12 gacha bo'lgan butun son bo'lishi kerak.",
		zhCN: "无效的月份。月份应为1到12之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DAY: {
		deDE: "Ungültiger Tag. Der Tagesangabe sollte eine Ganzzahl zwischen 1 und 31 sein.",
		enUK: "Invalid day. The day should be an integer from 1 to 31.",
		enUS: "Invalid day. The day should be an integer from 1 to 31.",
		esES: "El día es incorrecto. El día hay que introducirlo del 1 al 31.",
		faIR: "روز غیر معتبر می باشد. روز باید عددی صحیح بین 1 تا 31 باشد.",
		frFR: "Jour invalide. Le jour doit être un nombre entier de 1 à 31.",
		idID: "Hari tidak valid. Hari harus berupa bilangan bulat dari 1 hingga 31.",
		itIT: "Giorno scorretto. Il giorno dovrebbe essere un numero da 1 a 31.",
		jaJP: "無効な日。日は1から31までの整数である必要があります。",
		koKR: "잘못된 일입니다. 일은 1에서 31 사이의 정수여야 합니다.",
		plPL: "Nieprawidłowy dzień. Dzień powinien być liczbą całkowitą od 1 do 31.",
		ptBR: "Dia inválido. O dia deve ser um número inteiro de 1 a 31.",
		ruRU: "Неверный день. День должен быть задан целым числом от 1 до 31.",
		trTR: "Geçersiz gün. Gün, 1'den 31'e kadar bir tam sayı olmalıdır.",
		ukUA: "Недійсний день. День повинен бути цілим числом від 1 до 31.",
		uzUZ: "Yaroqsiz kun. Kun 1 dan 31 gacha bo'lgan butun son bo'lishi kerak.",
		zhCN: "无效的日期。日期应为1到31之间的整数。",
	},
	MESSAGE_TEXT_INVALID_DATE: {
		deDE: "Ungültiges Datumsformat. Zum Beispiel für den 20. February 2019 sende: 20.02.2019 oder 20.02.19",
		enUK: "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		enUS: "Invalid date format. For example for 20 February 2019 please submit: 20.02.2019 or 20.02.19",
		esES: "El formato de la fecha no es correcto. Por ejemplo para el día 20 de Febrero de 2019 introduce: 20.02.2019 o 20.02.19",
		faIR: "فرمت تاریخ غیر معتبر می باشد. برای مثال برای 20 فوریه 2019 لطفا اینگونه وارد کنید: 20.02.2019 یا 20.02.19",
		frFR: "Format de date invalide. Par exemple, pour le 20 février 2019, veuillez soumettre : 20.02.2019 ou 20.02.19",
		idID: "Format tanggal tidak valid. Misalnya untuk 20 Februari 2019 silakan kirimkan: 20.02.2019 atau 20.02.19",
		itIT: "Formato data sbagliato. Esempio: per il 20 Febbraio 2019 inserisci: 20.02.2019 oppure 20.02.19",
		jaJP: "無効な日付形式です。例えば、2019年2月20日の場合は、20.02.2019または20.02.19と入力してください。",
		koKR: "잘못된 날짜 형식입니다. 예를 들어 2019년 2월 20일의 경우 다음과 같이 제출하세요: 20.02.2019 또는 20.02.19",
		plPL: "Nieprawidłowy format daty. Na przykład dla 20 lutego 2019 proszę podać: 20.02.2019 lub 20.02.19",
		ptBR: "Formato de data inválido. Por exemplo, para 20 de fevereiro de 2019, envie: 20.02.2019 ou 20.02.19",
		ruRU: "Неверный формат даты. Например для 20 февраля 2019 года надо ввести: 20.02.2019 или 20.02.19",
		trTR: "Geçersiz tarih formatı. Örneğin 20 Şubat 2019 için lütfen şunu girin: 20.02.2019 veya 20.02.19",
		ukUA: "Недійсний формат дати. Наприклад, для 20 лютого 2019 року, будь ласка, введіть: 20.02.2019 або 20.02.19",
		uzUZ: "Yaroqsiz sana formati. Masalan, 2019 yil 20 fevral uchun quyidagilarni yuboring: 20.02.2019 yoki 20.02.19",
		zhCN: "无效的日期格式。例如，对于2019年2月20日，请提交：20.02.2019或20.02.19",
	},
	MESSAGE_TEXT_INVALID_PHONE_NUMBER: {
		deDE: "Ungültige Telefonnummer. Versuche es erneut oder gehe ins Haupt /menu",
		enUK: "Invalid phone number. Check and try it again? /menu",
		enUS: "Invalid phone number. Check and try it again? /menu",
		esES: "El número del teléfono no es correcto. ¿Comprobarlo y intentarlo de nuevo? /menú",
		faIR: "شماره تلفن غیر معتبر می باشد. آیا بررسی نموده، مجدداً سعی می کنید؟ /منو",
		frFR: "Numéro de téléphone invalide. Vérifiez et réessayez ? /menu",
		idID: "Nomor telepon tidak valid. Periksa dan coba lagi? /menu",
		itIT: "Numero di telefono invalido. Controlla e riprova. /menu",
		jaJP: "無効な電話番号です。確認して再試行しますか？ /menu",
		koKR: "잘못된 전화번호입니다. 확인하고 다시 시도하시겠습니까? /menu",
		plPL: "Nieprawidłowy numer telefonu. Sprawdź i spróbuj ponownie? /menu",
		ptBR: "Número de telefone inválido. Verificar e tentar novamente? /menu",
		ruRU: "Неверный номер. Проверить и попробовать ещё раз? /menu",
		trTR: "Geçersiz telefon numarası. Kontrol edip tekrar deneyin? /menu",
		ukUA: "Недійсний номер телефону. Перевірити та спробувати ще раз? /menu",
		uzUZ: "Yaroqsiz telefon raqami. Tekshirib ko'ring va qayta urinib ko'ring? /menu",
		zhCN: "无效的电话号码。检查并重试？ /menu",
	},
	MESSAGE_TEXT_PHONE_NUMBER_IS_NOT_SMS_CAPABLE: {
		deDE: "Diese Telefonnummer kann von uns keine SMS empfangen. Versuche eine andere oder gehe ins Haupt /menu",
		enUK: "This phone number not able to receive SMS. Try another number? /menu",
		enUS: "This phone number not able to receive SMS. Try another number? /menu",
		esES: "Este número de teléfono no acepta SMS. ¿Intentar otro número? /menú",
		faIR: "این شماره تلفن قادر به دریافت پیام کوتاه نمی باشد. آیا شماره دیگری را امتحان میکنید؟ /منو",
		frFR: "Ce numéro de téléphone ne peut pas recevoir de SMS. Essayer un autre numéro ? /menu",
		idID: "Nomor telepon ini tidak dapat menerima SMS. Coba nomor lain? /menu",
		itIT: "Questo numero di telefono non e' abilitato a ricevere SMS. Vuoi provare un altro numero? /menu",
		jaJP: "この電話番号はSMSを受信できません。別の番号を試しますか？ /menu",
		koKR: "이 전화번호는 SMS를 수신할 수 없습니다. 다른 번호를 시도하시겠습니까? /menu",
		plPL: "Ten numer telefonu nie może odbierać SMS-ów. Spróbować innego numeru? /menu",
		ptBR: "Este número de telefone não pode receber SMS. Tentar outro número? /menu",
		ruRU: "Данный номер не принимает SMS. Попробовать другой номер? /menu",
		trTR: "Bu telefon numarası SMS alamıyor. Başka bir numara deneyin? /menu",
		ukUA: "Цей номер телефону не може отримувати SMS. Спробувати інший номер? /menu",
		uzUZ: "Bu telefon raqami SMS qabul qila olmaydi. Boshqa raqamni sinab ko'rasizmi? /menu",
		zhCN: "此电话号码无法接收短信。尝试其他号码？ /menu",
	},
	MESSAGE_TEXT_NO_CONTACT_RECEIVED: {
		deDE: "Wir haben keine Kontakte empfangen. Versuche es erneut oder gehe ins Haupt /menu",
		enUK: "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		enUS: "We have not received any contacts. INSTRUCTION HOW TO DO IT. /menu",
		esES: "No hemos recibido ningún contacto. LA INSTRUCCIÓN COMO HACERLO. /menú",
		faIR: "ما هیچ اطلاعات تماسی دریافت نکردیم. دستورالعمل چگونگی انجام این کار. /منو",
		frFR: "Nous n'avons reçu aucun contact. INSTRUCTION COMMENT LE FAIRE. /menu",
		idID: "Kami belum menerima kontak apa pun. INSTRUKSI CARA MELAKUKANNYA. /menu",
		itIT: "Non abbiamo ricevuto nesusn contatto. ISTRUZIONI SU COME FARE. /menu",
		jaJP: "連絡先を受信していません。操作方法の説明。 /menu",
		koKR: "연락처를 받지 못했습니다. 방법 설명. /menu",
		plPL: "Nie otrzymaliśmy żadnych kontaktów. INSTRUKCJA JAK TO ZROBIĆ. /menu",
		ptBR: "Não recebemos nenhum contato. INSTRUÇÃO COMO FAZER. /menu",
		ruRU: "Мы не получили контакта. Тут инструкция как это сделать. /menu",
		trTR: "Herhangi bir kişi almadık. NASIL YAPILACAĞI TALİMATI. /menu",
		ukUA: "Ми не отримали жодних контактів. ІНСТРУКЦІЯ ЯК ЦЕ ЗРОБИТИ. /menu",
		uzUZ: "Biz hech qanday kontakt olmadik. BUNI QANDAY QILISH BO'YICHA KO'RSATMA. /menu",
		zhCN: "我们尚未收到任何联系人。操作指南。 /menu",
	},
	MESSAGE_TEXT_YOU_HAVE_NO_CONTACTS: {
		deDE: "Du hast noch keine Kontakte hinzugefügt.",
		enUK: "You have not created any contacts yet.",
		enUS: "You have not created any contacts yet.",
		esES: "Todavía no has creado ningún contacto.", //TODO:es - verify
		faIR: "هنوز هیچ مخاطبی را ایجاد نکرده اید",     //TODO:fa - verify
		frFR: "Vous n'avez pas encore créé de contacts.",
		idID: "Anda belum membuat kontak apa pun.",
		itIT: "Non hai ancora creato alcun contatto.", //TODO:it - verify
		jaJP: "まだ連絡先を作成していません。",
		koKR: "아직 연락처를 만들지 않았습니다.",
		plPL: "Nie utworzyłeś jeszcze żadnych kontaktów.",
		ptBR: "Você ainda não criou nenhum contato.",
		ruRU: "Вы ещё не создали контактов",
		trTR: "Henüz hiç kişi oluşturmadınız.",
		ukUA: "Ви ще не створили жодних контактів.",
		uzUZ: "Siz hali hech qanday kontakt yaratmagansiz.",
		zhCN: "您尚未创建任何联系人。",
	},
	MESSAGE_TEXT_CONTACT_NAME_IS_NUMBER: {
		deDE: "Du kannst nicht nur Zahlen für einen Kontaktnamen eingeben. Bitte benutz ein paar Buchstaben - du kannst mir vertrauen.",
		enUK: "You've entered just digits for a contact name. Please use some text characters.",
		enUS: "You've entered just digits for a contact name. Please use some text characters.",
		esES: "Has introducido solo números para el nombre del contacto. Por favor usa algunas letras.",
		faIR: "شما تنها اعداد را برای نام مخاطب وارد کرده اید. لطفا کاراکتر های متنی وارد کنید.",
		frFR: "Vous avez saisi uniquement des chiffres pour un nom de contact. Veuillez utiliser des caractères textuels.",
		idID: "Anda hanya memasukkan angka untuk nama kontak. Harap gunakan beberapa karakter teks.",
		itIT: "Hai inserito solamente numeri per un nome contatto. Usa anche alcune lettere.",
		jaJP: "連絡先名に数字のみを入力しました。テキスト文字を使用してください。",
		koKR: "연락처 이름에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		plPL: "Wprowadziłeś same cyfry jako nazwę kontaktu. Proszę użyć znaków tekstowych.",
		ptBR: "Você digitou apenas dígitos para um nome de contato. Por favor, use alguns caracteres de texto.",
		ruRU: "Вы ввели только цифры в качестве имени контакта. Пожалуйста используйте текстовые символы.",
		trTR: "Bir kişi adı için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		ukUA: "Ви ввели лише цифри для імені контакту. Будь ласка, використовуйте текстові символи.",
		uzUZ: "Kontakt nomi uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		zhCN: "您只输入了数字作为联系人名称。请使用一些文本字符。",
	},
	MESSAGE_TEXT_CURRENCY_NAME_IS_NUMBER: {
		deDE: "Bei der Währung erwarte ich eigentlich keine Zahlen. Nimm ein paar Buchstaben hinzu, um Verwirrung zu vermeiden.",
		enUK: "You've entered just digits for currency. Please use some text characters.",
		enUS: "You've entered just digits for currency. Please use some text characters.",
		esES: "Has introducido solo números para la moneda. Por favor usa algunas letras.",
		faIR: "شما تنها اعداد را برای واحد پولی وارد کرده اید. لطفا کاراکترهای متنی وارد کنید.",
		frFR: "Vous avez saisi uniquement des chiffres pour la devise. Veuillez utiliser des caractères textuels.",
		idID: "Anda hanya memasukkan angka untuk mata uang. Harap gunakan beberapa karakter teks.",
		itIT: "Hai inserito solamente numeri per la valuta. Usa anche alcune lettere.",
		jaJP: "通貨に数字のみを入力しました。テキスト文字を使用してください。",
		koKR: "통화에 숫자만 입력했습니다. 텍스트 문자를 사용해 주세요.",
		plPL: "Wprowadziłeś same cyfry dla waluty. Proszę użyć znaków tekstowych.",
		ptBR: "Você digitou apenas dígitos para moeda. Por favor, use alguns caracteres de texto.",
		ruRU: "Вы ввели только цифры в качестве номинала. Пожалуйста используйте текстовые символы.",
		trTR: "Para birimi için sadece rakamlar girdiniz. Lütfen bazı metin karakterleri kullanın.",
		ukUA: "Ви ввели лише цифри для валюти. Будь ласка, використовуйте текстові символи.",
		uzUZ: "Valyuta uchun faqat raqamlarni kiritdingiz. Iltimos, ba'zi matn belgilaridan foydalaning.",
		zhCN: "您只输入了数字作为货币。请使用一些文本字符。",
	},
	MESSAGE_TEXT_HISTORY_ROW_TO_USER_WITH_NAME: {
		deDE: "%v - %s ⇒ dir: %s",
		enUK: "%v - %s ⇒ to you: %s",
		enUS: "%v - %s ⇒ to you: %s",
		esES: "%v - %s ⇒ a ti: %s",
		faIR: "%v - %s ⇒ به شما: %s",
		frFR: "%v - %s ⇒ à vous: %s",
		idID: "%v - %s ⇒ kepada Anda: %s",
		itIT: "%v - %s ⇒ a te: %s",
		jaJP: "%v - %s ⇒ あなたへ: %s",
		koKR: "%v - %s ⇒ 당신에게: %s",
		plPL: "%v - %s ⇒ do Ciebie: %s",
		ptBR: "%v - %s ⇒ para você: %s",
		ruRU: "%v - %s ⇒ Вам : %s",
		trTR: "%v - %s ⇒ size: %s",
		ukUA: "%v - %s ⇒ Вам: %s",
		uzUZ: "%v - %s ⇒ sizga: %s",
		zhCN: "%v - %s ⇒ 给您: %s",
	},
	MESSAGE_TEXT_HISTORY_ROW_FROM_USER_WITH_NAME: {
		deDE: "%v - Du ⇒ %s : %s",
		enUK: "%v - You ⇒ %s : %s",
		enUS: "%v - You ⇒ %s : %s",
		esES: "%v - Tú ⇒ %s : %s",
		faIR: "%v - شما ⇒ %s : %s",
		frFR: "%v - Vous ⇒ %s : %s",
		idID: "%v - Anda ⇒ %s : %s",
		itIT: "%v - Tu ⇒ %s : %s",
		jaJP: "%v - あなた ⇒ %s : %s",
		koKR: "%v - 당신 ⇒ %s : %s",
		plPL: "%v - Ty ⇒ %s : %s",
		ptBR: "%v - Você ⇒ %s : %s",
		ruRU: "%v - Вы ⇒ %s : %s",
		trTR: "%v - Siz ⇒ %s : %s",
		ukUA: "%v - Ви ⇒ %s : %s",
		uzUZ: "%v - Siz ⇒ %s : %s",
		zhCN: "%v - 您 ⇒ %s : %s",
	},
	MESSAGE_TEXT_LETS_SEND_SMS: {
		deDE: "Lass uns eine SMS senden",
		enUK: "Let's send SMS",
		enUS: "Let's send SMS",
		esES: "Vamos a enviar un SMS",
		faIR: "پیام کوتاه ارسال کنید",
		frFR: "Envoyons un SMS",
		idID: "Mari kirim SMS",
		itIT: "Inviamo un SMS",
		jaJP: "SMSを送信しましょう",
		koKR: "SMS를 보내봅시다",
		plPL: "Wyślijmy SMS",
		ptBR: "Vamos enviar SMS",
		ruRU: "Давайте отправим SMS",
		trTR: "SMS gönderelim",
		ukUA: "Давайте надішлемо SMS",
		uzUZ: "SMS yuboraylik",
		zhCN: "让我们发送短信",
	},
	MESSAGE_TEXT_SMS_QUEUING_FOR_SENDING: {
		deDE: "Setze SMS in Sendewarteschlange für %v...",
		enUK: "Queuing SMS for sending to number %v...",
		enUS: "Queuing SMS for sending to number %v...",
		esES: "El SMS se está poniendo en la cola para enviar al número %v...",
		faIR: "پیام کوتاه برای ارسال به شماره مقابل در حال قرارگیری در نوبت ارسال می باشد %v...",
		frFR: "Mise en file d'attente du SMS pour envoi au numéro %v...",
		idID: "Antrian SMS untuk dikirim ke nomor %v...",
		itIT: "SMS in coda per l'invio al numero %v...",
		jaJP: "番号 %v に送信するためのSMSをキューに入れています...",
		koKR: "번호 %v로 보내기 위해 SMS를 대기열에 넣는 중...",
		plPL: "Kolejkowanie SMS do wysłania na numer %v...",
		ptBR: "Enfileirando SMS para envio ao número %v...",
		ruRU: "SMS ставится в очередь на отправку на номер %v...",
		trTR: "%v numarasına gönderilmek üzere SMS sıraya alınıyor...",
		ukUA: "Ставимо SMS у чергу для надсилання на номер %v...",
		uzUZ: "%v raqamiga yuborish uchun SMS navbatga qo'yilmoqda...",
		zhCN: "正在将短信排队发送到号码 %v...",
	},
	MESSAGE_TEXT_SMS_QUEUED_FOR_SENDING: {
		deDE: "SMS in Sendewarteschlange für %v",
		enUK: "SMS is queued for sending to number %v",
		enUS: "SMS is queued for sending to number %v",
		esES: "El SMS está en la cola para enviar al número %v",
		faIR: "پیام کوتاه برای شماره مقابل در نوبت ارسال قرار گرفت %v",
		frFR: "SMS mis en file d'attente pour envoi au numéro %v",
		idID: "SMS diantrekan untuk dikirim ke nomor %v",
		itIT: "SMS inserito in coda per l'invio al numero %v",
		jaJP: "SMSは番号 %v に送信するためにキューに入れられています",
		koKR: "SMS가 번호 %v로 보내기 위해 대기열에 있습니다",
		plPL: "SMS jest w kolejce do wysłania na numer %v",
		ptBR: "SMS está na fila para envio ao número %v",
		ruRU: "SMS поставлена в очередь на отправку на номер %v",
		trTR: "SMS, %v numarasına gönderilmek üzere sıraya alındı",
		ukUA: "SMS поставлено в чергу для надсилання на номер %v",
		uzUZ: "SMS %v raqamiga yuborish uchun navbatga qo'yildi",
		zhCN: "短信已排队发送到号码 %v",
	},
	MESSAGE_TEXT_BALANCE_HEADER: {
		deDE: "Ausstehend",
		enUK: "Balance",
		enUS: "Balance",
		esES: "Balance",
		faIR: "تراز",
		frFR: "Solde",
		idID: "Saldo",
		itIT: "Bilancio",
		jaJP: "残高",
		koKR: "잔액",
		plPL: "Saldo",
		ptBR: "Saldo",
		ruRU: "Баланс",
		trTR: "Bakiye",
		ukUA: "Баланс",
		uzUZ: "Balans",
		zhCN: "余额",
	},
	MESSAGE_TEXT_RECEIPT_AVAILABLE_CHANNELS: {
		deDE: "Entschuldigung, im Moment funktionieren nur ein paar Kanäle für den Versand von Quittungen:",
		enUK: "Sorry, at the moment just this channels are available for sending a receipt:",
		enUS: "Sorry, at the moment just these channels are available for sending a receipt:",
		esES: "Disculpa, de momento solo estos canales están disponibles para enviar el recibo:",
		faIR: "متاسفانه، در حال حاضر تنها این کانالها برای ارسال رسید در دسترس می باشند.",
		frFR: "Désolé, pour le moment seuls ces canaux sont disponibles pour l'envoi d'un reçu :",
		idID: "Maaf, saat ini hanya saluran ini yang tersedia untuk mengirim tanda terima:",
		itIT: "Spiacenti ma al momento solo questi canali sono disponibili per inviare debiti/crediti:",
		jaJP: "申し訳ありませんが、現時点では領収書の送信に利用できるのはこれらのチャンネルのみです：",
		koKR: "죄송합니다. 현재 영수증 발송에 사용할 수 있는 채널은 다음과 같습니다:",
		plPL: "Przepraszamy, w tej chwili tylko te kanały są dostępne do wysyłania paragonu:",
		ptBR: "Desculpe, no momento apenas estes canais estão disponíveis para enviar um recibo:",
		ruRU: "Извините, в данный момент доступны только эти каналы для отправки квитанции:",
		trTR: "Üzgünüz, şu anda makbuz göndermek için yalnızca bu kanallar kullanılabilir:",
		ukUA: "Вибачте, на даний момент для надсилання квитанції доступні лише ці канали:",
		uzUZ: "Kechirasiz, hozirda kvitansiya yuborish uchun faqat ushbu kanallar mavjud:",
		zhCN: "抱歉，目前只有这些渠道可用于发送收据：",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_TELEGRAM: {
		deDE: "📤 Quittung per Telegram verschickt.",
		enUK: "📤 Receipt sent through Telegram.",
		enUS: "📤 Receipt sent through Telegram.",
		esES: "📤 El recibo ha sido enviado vía Telegram.",
		faIR: "📤 رسید از طریق تلگرام ارسال شد.",
		frFR: "📤 Reçu envoyé via Telegram.",
		idID: "📤 Tanda terima dikirim melalui Telegram.",
		itIT: "📤 Notifica inviata tramite Telegram.",
		jaJP: "📤 Telegramを通じて領収書が送信されました。",
		koKR: "📤 텔레그램을 통해 영수증이 전송되었습니다.",
		plPL: "📤 Paragon wysłany przez Telegram.",
		ptBR: "📤 Recibo enviado pelo Telegram.",
		ruRU: "📤 Квитанция отправлена через телеграм.",
		trTR: "📤 Makbuz Telegram üzerinden gönderildi.",
		ukUA: "📤 Квитанцію надіслано через Telegram.",
		uzUZ: "📤 Kvitansiya Telegram orqali yuborildi.",
		zhCN: "📤 收据已通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_NOT_SENT_AS_COUNTERPARTY_HAS_DISABLED_TG_BOT: {
		deDE: "Quittung konnte nicht per Telegram gesendet werden, da %v den Chat mit dem Bot gelöscht hat.",
		enUK: "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		enUS: "Receipt NOT sent through Telegram as %v has deleted chat with the bot.",
		esES: "El recibo NO ha sido enviado vía Telegram porque %v ha eliminado el chat del bot.",
		faIR: "از آنجایی که %v چت انجام شده با روبات را حذف کرده است رسید از طریق تلگرام ارسال نشد.",
		frFR: "Reçu NON envoyé via Telegram car %v a supprimé le chat avec le bot.",
		idID: "Tanda terima TIDAK dikirim melalui Telegram karena %v telah menghapus obrolan dengan bot.",
		itIT: "Notifica NON inviata tramite Telegram a %v perche' ha cancellato la chat con il bot",
		jaJP: "%v がボットとのチャットを削除したため、Telegramを通じて領収書が送信されませんでした。",
		koKR: "%v가 봇과의 채팅을 삭제했기 때문에 텔레그램을 통해 영수증이 전송되지 않았습니다.",
		plPL: "Paragon NIE został wysłany przez Telegram, ponieważ %v usunął czat z botem.",
		ptBR: "Recibo NÃO enviado pelo Telegram, pois %v excluiu o chat com o bot.",
		ruRU: "Квитанция НЕ отправлена через телеграм так как %v удалил чат с ботом.",
		trTR: "%v botla sohbeti sildiği için makbuz Telegram üzerinden GÖNDERİLMEDİ.",
		ukUA: "Квитанцію НЕ надіслано через Telegram, оскільки %v видалив чат з ботом.",
		uzUZ: "%v bot bilan suhbatni o'chirganligi sababli kvitansiya Telegram orqali YUBORILMADI.",
		zhCN: "由于 %v 已删除与机器人的聊天，收据未通过Telegram发送。",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_EMAIL: {
		deDE: "Quittung wurde per Mail gesendet. (id: %v)",
		enUK: "Receipt sent through email. (id: %v)",
		enUS: "Receipt sent through email. (id: %v)",
		esES: "El recibo ha sido enviado vía e-mail. (id: %v)",
		faIR: "رسید از طریق ایمیل ارسال شد. (id: %v)",
		frFR: "Reçu envoyé par e-mail. (id: %v)",
		idID: "Tanda terima dikirim melalui email. (id: %v)",
		itIT: "Notifica inviata tramite email (id: %v)",
		jaJP: "領収書がメールで送信されました。(id: %v)",
		koKR: "이메일을 통해 영수증이 전송되었습니다. (id: %v)",
		plPL: "Paragon wysłany przez e-mail. (id: %v)",
		ptBR: "Recibo enviado por e-mail. (id: %v)",
		ruRU: "Квитанция отправлена через email. (id: %v)",
		trTR: "Makbuz e-posta yoluyla gönderildi. (id: %v)",
		ukUA: "Квитанцію надіслано електронною поштою. (id: %v)",
		uzUZ: "Kvitansiya elektron pochta orqali yuborildi. (id: %v)",
		zhCN: "收据已通过电子邮件发送。(id: %v)",
	},
	MESSAGE_TEXT_RECEIPT_SENT_THROW_SMS: {
		deDE: "Quittung wurde per SMS gesendet.",
		enUK: "Receipt sent through SMS",
		enUS: "Receipt sent through SMS",
		esES: "El recibo ha sido enviado vía SMS.",
		faIR: "رسید از طریق پیام کوتاه ارسال شد.",
		frFR: "Reçu envoyé par SMS",
		idID: "Tanda terima dikirim melalui SMS",
		itIT: "Notifica inviata tramite SMS",
		jaJP: "領収書がSMSで送信されました",
		koKR: "SMS를 통해 영수증이 전송되었습니다",
		plPL: "Paragon wysłany przez SMS",
		ptBR: "Recibo enviado por SMS",
		ruRU: "Квитанция отправлена через SMS.",
		trTR: "Makbuz SMS yoluyla gönderildi",
		ukUA: "Квитанцію надіслано через SMS",
		uzUZ: "Kvitansiya SMS orqali yuborildi",
		zhCN: "收据已通过短信发送",
	},
	MESSAGE_TEXT_SWITCH_TO_PM_TO_VIEW_RECEIPT: {
		deDE: "Schalte in den Privatmodus, um die Quittungsdetails zu sehen",
		enUK: "Switch to private mode to see receipt details.",
		enUS: "Switch to private mode to see receipt details.",
		esES: "Pasar al modo privado para ver el recibo.",
		faIR: "انتقال به حالت خصوصی جهت رویت جزئیات رسید.",
		frFR: "Passez en mode privé pour voir les détails du reçu.",
		idID: "Beralih ke mode pribadi untuk melihat detail tanda terima.",
		itIT: "Passa alla modalita' privata per vedere i dettagli dei tuoi crediti/debiti.",
		jaJP: "領収書の詳細を確認するには、プライベートモードに切り替えてください。",
		koKR: "영수증 세부 정보를 보려면 개인 모드로 전환하세요.",
		plPL: "Przełącz na tryb prywatny, aby zobaczyć szczegóły paragonu.",
		ptBR: "Mude para o modo privado para ver os detalhes do recibo.",
		ruRU: "Переключитьсь на чат с ботом чтобы посмотреть квитанцию",
		trTR: "Makbuz ayrıntılarını görmek için özel moda geçin.",
		ukUA: "Перейдіть у приватний режим, щоб побачити деталі квитанції.",
		uzUZ: "Kvitansiya tafsilotlarini ko'rish uchun shaxsiy rejimga o'ting.",
		zhCN: "切换到私人模式以查看收据详情。",
	},
	MESSAGE_TEXT_RECEIPT_ATTEMPT_TO_VIEW_OWN: {
		enUK: "An attempt to view own receipt.",
		ruRU: "Попытка посмотреть свою собственную квитанцию.",
	},
	MESSAGE_TEXT_RECEIPT_VIEWED_BY_COUNTERPARTY: {
		deDE: "👓 Quittung wurde von Gegenpartei gesichtet",
		enUK: "👓 Receipt viewed",
		esES: "👓 El recibo ha sido visto",
		faIR: "👓 رسید رویت شد.",
		itIT: "👓 Debiti visti",
		ruRU: "👓 Квитанция просмотрена",
	},
	MESSAGE_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		deDE: "Du kannst deine eigene Nummer in dem Format anzeigen, welches wir erwarten.",
		enUK: "You can view your own phone number in the format we are expecting.",
		esES: "Puedes ver tu número de teléfono en el formato previsto.",
		faIR: "شما می توانید شماره تلفن خود را با فرمتی که ما انتظار داریم ببینید.",
		itIT: "Puoi visionare il tuo numero di telefono nel formato previsto.",
		ruRU: "Вы можете посмотреть свой номер телефона в ожидаемом нами формате.",
	},
	COMMAND_TEXT_VIEW_MY_NUMBER_IN_INTERNATIONAL_FORMAT: {
		deDE: "Zeige meine Nummer im erwarteten Format.",
		enUK: "View my number in the expectd format",
		esES: "Mostrar mi número de teléfono en el formato previsto",
		faIR: "رویت شماره خودم با فرمت مورد انتظار",
		itIT: "Mostra il mio numero nel formato previsto",
		ruRU: "Посмотреть мой номер в ожидаемоем формате",
	},
	INLINE_BUTTON_SHOW_FULL_HISTORY: {
		deDE: "Lade den ganzen Verlauf",
		enUK: "Show full history",
		esES: "Mostrar cronología completa",
		faIR: "نمایش کامل سوابق",
		itIT: "Mostra cronologia completa",
		ruRU: "Показать всю историю",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_NOT_A_NUMBER: {
		deDE: "Das ist keine Nummer",
		enUK: "it is not a number",
		esES: "No es un número",
		faIR: "این یک شماره نیست",
		itIT: "Non e' un numero",
		ruRU: "Это не цифра",
	},
	MESSAGE_TEXT_INCORRECT_VALUE_IS_NEGATIVE: {
		deDE: "Die Nummer sollte positiv sein (<i>größer als 0</i>)",
		enUK: "The number should be positive (<i>greater than 0</i>)",
		esES: "El número tiene que ser positivo (<i>más de  0</i>)",
		faIR: "شماره باید مثبت باشد (<i>بزرگتر از 0</i>)",
		itIT: "Il numero dovrebbe essere positivo (<i>maggiore di 0</i>)",
		ruRU: "Цифра должна быть положительной (<i>больше нуля</i>)",
	},
	MESSAGE_TEXT_ASK_HOW_MUCH_HAS_BEEN_RETURNED: {
		deDE: "Wie viel wurde beglichen?",
		enUK: "How much have been returned?",
		esES: "¿Cuánto/s te han devuelto?",
		faIR: "چه مقدار بازپرداخت شده است؟",
		itIT: "Quanto ti e' stato restituito?",
		ruRU: "Сколько было возвращено?",
	},
	MESSAGE_TEXT_RETURN_IS_TOO_BIG: {
		deDE: "Sie haben entschieden, %v zurückzugeben, aber der ausstehende Betrag ist nur %v.\n\nBitte geben Sie Werte gleich %v oder weniger ein.",
		enUK: "You decided to return %v but outstanding amount is just %v.\n\nPlease enter values equal to %v or less.",
		esES: "Decidiste devolver %v pero la cantidad pendiente es solo %v.\n\nPor favor ingrese valores iguales a %v o menos.",
		faIR: "شما تصمیم گرفتید %v را بازگردانید اما مقدار قابل توجهی فقط %v است.\n\nلطفا مقادیر برابر %v یا کمتر را وارد کنید",
		itIT: "Hai deciso di restituire %v ma la quantità in sospeso è solo %v.\n\nInserisci valori pari o uguali a %v o meno.",
		ruRU: "Вы решили вернуть %v, но непогашенная сумма равна %v. \n\n Пожалуйста, введите значение равное %v или меньше.",
	},
	MESSAGE_TEXT_HELP_ROOT: {
		deDE: "Was hast du für eine Frage? Wenn irgendwas unklar ist, frag ruhig hier @%v",
		enUK: "What is your question? If anything is missed here, feel free to ask in our @%v",
		esES: "¿Cuál es tu pregunta? Si algo se pierde aquí, siéntase libre de preguntar en nuestro @%v",
		faIR: "سوالت چیست؟ اگر چیزی در اینجا از دست رفته است، لطفا در @%v ما بپرسید",
		itIT: "Qual è la tua domanda? Se qualche cosa è mancato qui, non esitate a chiedere al nostro @%v",
		ruRU: "Какой у вас вопрос? Если здесь нет ответа пожалуйста спросите в нашей группе @%v",
	},
	MESSAGE_TEXT_HELP_BACK_TO_ROOT: {
		deDE: "Zurück zur FAQ Liste",
		enUK: "Back to FAQ list",
		esES: "Volver a la lista de preguntas frecuentes",
		faIR: "بازگشت به لیست سوالات متداول",
		itIT: "Torna all'elenco delle FAQ",
		ruRU: "Назад к списку вопросов",
	},
	HELP_HOW_TO_CREATE_BILL_Q: {
		deDE: "Wie erstellt man Rechnungen?",
		enUK: "How to create new bill?",
		esES: "¿Cómo crear una nueva factura?",
		faIR: "چگونه برای ایجاد لایحه جدید؟",
		itIT: "Come creare un nuovo conto?",
		ruRU: "Как создать новый счёт?",
	},
	HELP_HOW_TO_CREATE_BILL_A: {
		deDE: `<b>Wie man eine Rechnung erstellt</b>
<pre>Rechnung — Kosten, die unter zwei oder mehr Personen geteilt werden.</pre>

Deswegen kannst du am besten direkt <b>im Telegram Chat eine Rechnung erstellen, in zwei Schritten</b>:
<i>Benutze die Funktion "Teile Rechnung mit Telegram Benutzer", um es schnell zu machen:</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Füg mich einer Gruppe hinzu</a> oder öffne einen Chat mit deinem Freund.

	2. Schreibe <code>@splitusbot {amount} {bill_name}</code> und wähle ein passendes Ergebnis. Zum Beispiel:
<pre>		@splitusbot 45.5 pizza</pre>
	   Und dann kann jeder, der die Rechnung teilen will mit <code>Join</code> einsteigen.

<b>Alternativ</b> kannst du die Rechnung auch direkt in @{{.BotCode}} erstellen. Aber dann musst du die Personen, mit denen du die Rechnung teilst, einzeln hinzufügen.`,
		enUK: `<b>How to create a new bill</b>
<pre>Bill — shared expense between two or more people.</pre>

That is why the best is to <b>create bill in Telegram chat just in 2 steps</b>:
<i>use "Split bill with Telegram user(s)" to do it quickly</i>

	1. <a href="https://t.me/{{.BotCode}}?startgroup=from-help">Add me to Telegram group</a> or open chat with a friend.

	2. Type <code>@splitusbot {amount} {bill_name}</code> and select result from menu. For example:
<pre>		@splitusbot 45.5 pizza</pre>
	   Than any member of the group can share the bill by pressing <code>Join</code> button.

<b>Alternatively</b> you can create a bill right in the @{{.BotCode}}. But then you would need manually to add participants.`,
		esES: "",
		faIR: "",
		itIT: "",
		ruRU: `<b>How to create a new bill</b>
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
		deDE: "Bitte melde jedes Problem und jeden Wunsch auf unserer Webseite.",
		enUK: "Please report any issue or submit a feature request at our website.",
		esES: "Puedes informarnos de algún problema o proponernos cualquier mejora en nuestra web.",
		faIR: "لطفاً در وب سایت ما هرگونه مسئله ای را اعلام فرموده یا درخواست خود را مطرح نمایید.",
		itIT: "Segnala un problema o proponi un miglioramento sul nostro sito web.",
		ruRU: "Вы можете сообщить о проблеме или сделать предложения по улучшению программы на нашем сайте.",
	},
	COMMAND_TEXT_OPEN_USER_REPORT: {
		deDE: "Support Seite",
		enUK: "Support page",
		esES: "La página de ayuda",
		faIR: "صفحه پشتیبانی",
		itIT: "Pagina d'aiuto",
		ruRU: "Cтраница поддержки ",
	},
	COMMAND_TEXT_REPORT_A_BUG: {
		deDE: "einen Fehler melden",
		enUK: "Report a bug",
		esES: "Informar de un error",
		faIR: "گزارش یک باگ",
		itIT: "Segnala un bug",
		ruRU: "Сообщить об ошибке",
	},
	COMMAND_TEXT_SUBMIT_AN_IDEA: {
		deDE: "eine Idee äußern",
		enUK: "Add an idea",
		esES: "Proponer una idea",
		faIR: "یک ایده اضافه کنید.",
		itIT: "Proponi un idea",
		ruRU: "Предложить идею",
	},
	MESSAGE_TEXT_WELCOME: {
		deDE: `Hallo, ich bin Collectius - dein persönlicher Buchhalter.

Ich kann dir sagen, wer wem was schuldet und wann die Schuld jeweils fällig ist.

Was würdest du gerne machen?`,
		enUK: `Hi, I'm Collectius - your personal accountant & collector.

I can record who is owing to whom and remind when the return is due.

What would you like to do?`,

		esES: `Hola, me llamo Collectius, soy tu contable y asesor personal.

Puedo apuntar quién debe a quién y recordarte la fecha de devolución.

¿Qué te apetecería hacer?`,

		faIR: `سلام، من کالکتیوس هستم - حسابدار شخصی و مامور وصول شما

من میتوانم اینکه چه کسی به چه کسی بدهکار است را ثبت کرده و زمان بازپرداخت را یادآوری کنم.

دوست دارید چکار کنید؟`,

		itIT: `Ciao, sono Collectius - il tuo contabile ed esattore.

Posso annotare chi deve soldi a chi e ricordarti la data di scadenza.

Cosa vorresti fare ora?`,

		ruRU: `Привет, я Коллектиус - Ваш персональный счетовод и коллектор.

Могу записать кто кому чего должен и, и при необходимости, напомнить когда должок пора возвращать.

Чего изволит новый хозяин?`,
	},
	COMMAND_TEXT_SEND_ME_NEW_INVITE: {
		deDE: "Ich hätte gerne einen Code",
		enUK: "I want to get an invite",
		esES: "Me gustaría obtener una invitación",
		faIR: "می خواهم یک دعوت دریافت کنم.",
		itIT: "Voglio ottenere un invito",
		ruRU: "Хочу получить приглашение",
	},
	COMMAND_TEXT_I_HAVE_INVITE: {
		deDE: "Ich habe einen Code",
		enUK: "I have the invitation code",
		esES: "Tengo el código de la invitación",
		faIR: "من کد دعوت را دارم.",
		itIT: "Ho il codice invito",
		ruRU: "У меня есть код приглашения",
	},
	COMMAND_TEXT_I_HAVE_NOT_GOT_EMAIL: {
		deDE: "Ich habe noch keine Mails bekommen",
		enUK: "I have not got any emails",
		esES: "No he recibido ningún e-mail",
		faIR: "من ایمیلی دریافت نکردم",
		itIT: "Non ho ricevuto nessun email",
		ruRU: "Я не получил письма на email",
	},
	MESSAGE_TEXT_ONBOARDING_TELL_ABOUT_INVITES: {

		deDE: `<b>%v</b>,

Im Moment ist der Bot leider nur durch Einladungen von Freunden zugänglich.

Wenn du keinen Code hast, lass deine Kontaktdaten da und wir senden dir einen Code sobald du dran bist.

Wir senden 10 Codes am Tag an die, die am längsten warten und einen zufällig.`,

		enUK: `<b>%v</b>,

At the moment our bot is available just by invitation from friends.

If you have no code you can leave your contact details and we'll send you an invite as soon as your queue is due.

We send 10 invites per day to those in the head of the queue and 1 randomly.`,

		esES: `<b>%v</b>,

De momento nuestro bot está disponible solo por invitación de amigos.

Si no tienes el código puedes dejarnos tu contacto y te lo enviaremos cuando sea tu turno en la cola .

Enviamos 10 invitaciones por día a los primeros de la cola + 1 de modo casual.`,

		faIR: `<b>%v</b>,

درحال حاضر ربات ما تنها با دریافت دعوت از دوستان در دسترس می باشد.

اگر شما کدی در اختیار ندارید می توانید اطلاعات تماس خود را برای من وارد نموده و من به محض اینکه نوبت شما فرارسید یک دعوتنامه برایتان ارسال می کنم.

ما روزانه 10 دعوتنامه برای نفرات اول لیست انتظار و همچنین یک دعوتنامه تصادفی ارسال میکنیم.`,

		itIT: `<b>%v</b>,

Al momento il nostro bot e' disponibile solo tramite invito da amici.

Se non hai un codice puoi lasciarci il tuo contatto e ti manderemo un invito non appena sara' il tuo turno.

Inviamo 10 inviti al giorno ai primi 10 della lista d'attesa ed 1 in modo casuale.`,

		ruRU: `<b>%v</b>,

	На данный момент наш бот доступен только тем кто получил приглашение от друзей.

	Если у вас нет кода можете оставить свои контакты и мы вышлем вам приглашение как только наступит ваша очередь.

	Мы высылаем по 10 кодов в день первоочередникам + 1 случайным образом.`,
	},
	EMAIL_INVITE_SUBJ: {
		deDE: "Eine Einladung von {{.FromName}} - Code: {{.InviteCode}}",
		enUK: "An invite from {{.FromName}} - code: {{.InviteCode}}",
		esES: "La invitación de {{.FromName}} - el código: {{.InviteCode}}",
		faIR: "دعوت از طرف {{.FromName}} - کد: {{.InviteCode}}",
		itIT: "Hai ricevuto un codice invito da {{.FromName}} - codice: {{.InviteCode}}",
		ruRU: "Приглашение от {{.FromName}} - код: {{.InviteCode}}",
	},
	SMS_INVITE_TEXT: {

		deDE: `Hey {{.ToName}}, {{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		enUK: `Hi {{.ToName}}, {{.FromName}} is inviting you to try debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Your personal invitation code is: {{.InviteCode}}`,

		esES: `Hola {{.ToName}}, {{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		faIR: `سلام{{.ToName}}, {{.FromName}} شما را دعوت کرده تا برنامه ردیابی بدهی ها را امتحان کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		itIT: `Ciao {{.ToName}}, {{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		ruRU: `Привет {{.ToName}}, {{.FromName}} рекомендует приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_TEXT: {
		deDE: `Hey {{.ToName}},

{{.FromName}} lädt dich ein die neue Schuldentracker App auszuprobieren - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Dein persönlicher Code lautet: {{.InviteCode}}`,

		enUK: `Hi {{.ToName}},

{{.FromName}} is inviting you to use debts tracking app - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

You invitation code is: {{.InviteCode}}`,

		esES: `Hola {{.ToName}},

{{.FromName}} te ha invitado a probar la aplicación para controlar tus deudas - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

El código de tu invitación es: {{.InviteCode}}`,

		faIR: `سلام{{.ToName}},

{{.FromName}} شما را دعوت کرده تا از برنامه ردیابی بدهی ها استفاده کنید.- https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

کد دعوت شما: {{.InviteCode}}`,

		itIT: `Ciao {{.ToName}},

{{.FromName}} ti ha invitato a provare 'debts tracking app' - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

Il tuo codice di invito personale e': {{.InviteCode}}`,

		ruRU: `Привет {{.ToName}},

	{{.FromName}} приглашает тебя попробовать приложение для учёта долгов - https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}

	Ваш код приглашения: {{.InviteCode}}`,
	},
	EMAIL_INVITE_HTML: {
		deDE: `<p>Hey {{.ToName}}, </p>

<p>{{.FromName}} lädt dich ein <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">die neue Schuldentracker App auszuprobieren</a>.</p>

<p>Dein persönlicher Code lautet: <b>{{.InviteCode}}</b></p>`,
		enUK: `<p>Hi {{.ToName}}, </p>

<p>{{.FromName}} is inviting you to try <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Your invitation code is: <b>{{.InviteCode}}</b></p>`,

		esES: `<p>Hola {{.ToName}}, </p>

<p>{{.FromName}} te ha invitado a <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">probar la app para controlar tus deudas</a>.</p>

<p>El código de tu invitación es: <b>{{.InviteCode}}</b></p>`,

		faIR: `<p>سلام{{.ToName}},</p>

<p>{{.FromName}} п شما را دعوت کرده به <a href="https://debtstracker.io/invite#id={{.InviteCode}}&ربات تلگرام={{.TgBot}}&{{.Utm}}"> امتحان برنامه ردیابی بدهی ها</a>.</p>

<p>You invitation code is: <b>{{.InviteCode}}</b></p>`,

		itIT: `<p>Ciao {{.ToName}},</p>

<p>{{.FromName}} ti ha invitato a provare <a href="https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">debts tracking app</a>.</p>

<p>Il tuo codice di invito personale e': <b>{{.InviteCode}}</b></p>`,

		ruRU: `<p>Привет {{.ToName}}, </P

	<p>{{.FromName}} приглашает тебя <a href = "https://debtstracker.io/invite#id={{.InviteCode}}&telegram-bot={{.TgBot}}&{{.Utm}}">попробовать приложение для учёта долгов</a>.</p>

	<p>Ваш код приглашения: <b>{{.InviteCode}}</b></p>`,
	},
	EMAIL_RECEIPT_SUBJ: {
		deDE: "Schuldschein - {{.FromName}}",
		enUK: "Debt record - {{.FromName}}",
		esES: "La notificación de la deuda - {{.FromName}}",
		faIR: "سوابق بدهی - {{.FromName}}",
		itIT: "Debito - {{.FromName}}",
		ruRU: "Запись о долге - {{.FromName}}",
	},
	EMAIL_RECEIPT_BODY_TEXT: {
		deDE: "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		enUK: "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		esES: "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		faIR: "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		itIT: "{{.FromName}} ha creato un debito: {{.ReceiptURL}} ",
		ruRU: "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	MESSENGER_RECEIPT_TEXT: {
		deDE: "Ich habe online einen Schuldschein erstellt, für mehr Details siehe {{.ReceiptURL}}",
		enUK: "I've created a debt record regards you, see details at {{.ReceiptURL}}",
		esES: "He creado una notificación de la deuda, las detalles están aquí {{.ReceiptURL}}",
		faIR: "من یک سابقه بدهی برای شما ایجاد کرده ام، لطفا جزئیات را ملاحظه فرمایید در {{.ReceiptURL}}",
		itIT: "Ho creato un debito a tuo nome, controlla i dettagli qui - {{.ReceiptURL}}",
		ruRU: "Я создал(а) запись о долге, подробности тут - {{.ReceiptURL}}",
	},
	EMAIL_RECEIPT_BODY_HTML: {
		deDE: "{{.FromName}} hat einen Schuldschein erstellt: {{.ReceiptURL}}",
		enUK: "{{.FromName}} created a debt record: {{.ReceiptURL}}",
		esES: "{{.FromName}} ha creado una notificación de la deuda: {{.ReceiptURL}}",
		faIR: "{{.FromName}} یک سابقه بدهی ایجاد کرده است: {{.ReceiptURL}}",
		itIT: "{{.FromName}} ha creato un debito: {{.ReceiptURL}}",
		ruRU: "{{.FromName}} создал(а) запись о долге: {{.ReceiptURL}}",
	},
	INLINE_RECEIPT_TITLE: {
		deDE: "Empfänger: %v",
		enUK: "Receipt: %v",
		esES: "El recibo: %v",
		faIR: "رسید: %v",
		itIT: "Debito/credito: %v",
		ruRU: "Квитанция: %v",
	},
	INLINE_RECEIPT_DESCRIPTION: {
		deDE: "Klicken sie hier, um die Quittung zu sehen",
		enUK: "Click here to send the receipt",
		esES: "Haz click aquí para enviar el recibo",
		faIR: "برای ارسال رسید اینجا کلیک کنید.",
		itIT: "Clicca qui per inviare un debito/credito",
		ruRU: "Нажмите здесь чтобы отправить квитанцию",
	},
	INLINE_RECEIPT_CHOOSE_LANGUAGE: {
		deDE: "<b>Bitte wählen Sie eine Sprache, um weitere Details zu sehen, die den Schuldschein betreffen</b>, der erstellt wurde von {{.Creator}}.",
		enUK: "<b>Please choose language to see details of the debt</b> that has been recorded by {{.Creator}}.",
		esES: "<b>Elige el idioma para ver los detalles de la deuda</b> que ha sido creada por {{.Creator}}.",
		faIR: "<b> لطفا برای رویت جزئیات بدهی که توسط </b>  {{.Creator}} ثبت شده است زبان را انتخاب کنید.",
		itIT: "<b>Scegli la lingua per vedere i dettagli del debito</b> registrato da {{.Creator}}.",
		ruRU: "<b>Выберите язык чтобы посмотреть подробности записи о долге</b> которую создал(а) {{.Creator}}.",
	},
	INLINE_RECEIPT_FOOTER: {
		//-------------------------------------------------------
		deDE: `{{.SiteLink}} — eine App, die dir hilft Schulden zu überwachen:

  - Du weißt immer, wie viel du allen schuldest

  - Keine Fälligkeit wird verpasst
    <i>(erinnert dich und die Gläubiger)</i>`,
		//-------------------------------------------------------
		enUK: `{{.SiteLink}} — an app for debts tracking will help you to:

  - Always know your bottom line

  - Return debts on time
    <i>(reminders to you & your debtors)</i>`,
		//-------------------------------------------------------
		esES: `{{.SiteLink}} — la app para controlar tus deudas te ayuda a:

  - Saber siempre quién debe a quién

  - Devolver la deuda a tiempo
    <i>(recordatorio a ti y a tus deudores)</i>`,
		//-------------------------------------------------------
		faIR: `{{.SiteLink}} — یک برنامه پیگیری بدهی است که به شما کمک می کند تا:

  - همیشه از سود و زیان خود مطلع باشید.

  - بدهی ها به موقع پرداخت شوند.
    <i>(با ارسال یادآوری به  شما و بدهکاران )</i>`,
		//-------------------------------------------------------
		itIT: `{{.SiteLink}} — un app per i debiti che ti consento di:

  - Sapere sempre chi deve soldi a chi

  - Restituire i soldi in tempo
    <i>(lo ricorda a te ed al tuo debitore)</i>`,
		//-------------------------------------------------------
		ruRU: `{{.SiteLink}} — программа для учёта долгов поможет:

	  - Всегда знать кто кому сколько должен

	  - Незабыть вовремя отдать или востребовать долг
	    <i>(напоминания вам и вашим должникам)</i>`,
		//-------------------------------------------------------
	},
	INLINE_RECEIPT_GENERATING_MESSAGE: {
		deDE: `<b>{{.Creator}} erstellte online einen Schuldschein</b> der dich betrifft.

>> Generating receipt`,
		//-------------------------------------------------------
		enUK: `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		esES: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		itIT: `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
		ruRU: `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <i>Generating receipt...</i>`,
		//-------------------------------------------------------
	},
	//	INLINE_RECEIPT_MESSAGE: {
	//		//-------------------------------------------------------
	//		enUK: `<b>{{.Creator}} recorded a debt</b> associated with you.
	//
	//`,
	//		//-------------------------------------------------------
	//		esES: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.
	//
	//`,
	//		//-------------------------------------------------------
	//		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.
	//
	//`,
	//		//-------------------------------------------------------
	//		itIT:   `<b>{{.Creator}} ha registrato un debito</b> associato a te.
	//
	//`,
	//		//-------------------------------------------------------
	//		ruRU:  `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.
	//
	//`,
	//		//-------------------------------------------------------
	//	},
	INLINE_RECEIPT_MESSAGE: {
		//-------------------------------------------------------
		enUK: `<b>{{.Creator}} recorded a debt</b> associated with you.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		esES: `<b>{{.Creator}} ha creado una deuda</b> asociada a ti.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		faIR: `<b>{{.Creator}} یک بدهی </b> مرتبط با شما ثبت نموده است.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		itIT: `<b>{{.Creator}} ha registrato un debito</b> associato a te.

  >> <a href="{{.ReceiptUrl}}">Click here to view receipt</a>`,
		//-------------------------------------------------------
		ruRU: `<b>{{.Creator}} создал(а) запись о долге</b> касающегося Вас.

  >> <a href="{{.ReceiptUrl}}">Посмотреть квитанцию</a>`,
		//-------------------------------------------------------
	},
	InlineInviteToJoinFamilyTitle: {
		enUK: "Invitation to join family at @%s",
		ruRU: "Приглашение в семью на @%s",
		deDE: "Einladung der Familie beizutreten bei @%s",
		esES: "Invitación para unirse a la familia en @%s",
		frFR: "Invitation à rejoindre la famille à @%s",
		itIT: "Invito a unirsi alla famiglia a @%s",
		ptBR: "Convite para se juntar à família em @%s",
		faIR: "دعوت به پیوستن به خانواده در @%s",
	},
	InlineInviteToJoinFamilyDescription: {
		enUK: "Click here to send an invite",
		ruRU: "Нажмите здесь для отправки приглашения",
		deDE: "Klick hier, um eine Einladung zu versenden",
		esES: "Haz click para enviar la invitación",
		frFR: "Cliquez ici pour envoyer une invitation",
		itIT: "Clicca qui per spedire un invito",
		ptBR: "Clique aqui para enviar um convite",
		faIR: "برای ارسال یک دعوتنامه اینجا کلیک کنید.",
	},
	YouAreInvitedToJoinFamilyMessage: {
		deDE: "Sie sind eingeladen, dem Familienkonto bei @{BOT_ID} beizutreten.",
		enUK: "You are invited to join family account at @{BOT_ID}.",
		esES: "Se te invita a unirte a la cuenta familiar en @{BOT_ID}.",
		faIR: "شما دعوت شده\u200cاید که به حساب خانواده در @{BOT_ID} بپیوندید.",
		frFR: "Vous êtes invité à rejoindre le compte familial sur @{BOT_ID}.",
		idID: "Anda diundang untuk bergabung dengan akun keluarga di @{BOT_ID}.",
		itIT: "Sei invitato a unirti al conto familiare su @{BOT_ID}.",
		jaJP: "@{BOT_ID}で家族アカウントに参加するように招待されています。",
		koKR: "@{BOT_ID}에서 가족 계정에 초대되었습니다.",
		plPL: "Zostałeś zaproszony do dołączenia do konta rodzinnego na @{BOT_ID}.",
		ptBR: "Você foi convidado a entrar na conta familiar em @{BOT_ID}.",
		ruRU: "Вас пригласили присоединиться к семейному аккаунту на @{BOT_ID}.",
		trTR: "@{BOT_ID} üzerinde aile hesabına katılmaya davetlisiniz.",
		ukUA: "Вас запросили приєднатися до сімейного акаунта на @{BOT_ID}.",
		uzUZ: "Sizni @{BOT_ID} oilaviy hisobiga qo‘shilishga taklif qilishdi.",
		zhCN: "您被邀请加入 @{BOT_ID} 的家庭账户。",
	},
	BtnViewFamilyAccount: {
		deDE: "Familienkonto ansehen",
		enUK: "View family account",
		esES: "Ver cuenta familiar",
		faIR: "مشاهده حساب خانواده",
		frFR: "Voir le compte familial",
		idID: "Lihat akun keluarga",
		itIT: "Visualizza conto familiare",
		jaJP: "家族アカウントを表示",
		koKR: "가족 계정 보기",
		plPL: "Zobacz konto rodzinne",
		ptBR: "Ver conta familiar",
		ruRU: "Посмотреть семейный аккаунт",
		trTR: "Aile hesabını görüntüle",
		ukUA: "Переглянути сімейний акаунт",
		uzUZ: "Oilaviy hisobni ko‘rish",
		zhCN: "查看家庭账户",
	},
	SMS_RECEIPT_YOU_GOT: {
		deDE: "Du hast dir %v von %v geliehen.",
		enUK: "You've got %v from %v.",
		esES: "Has recibido %v de %v.",
		faIR: "شما دریافت کرده اید %v از %v.",
		itIT: "Hai ricevuto %v da %v",
		ruRU: "Вы получили %v от %v.",
	},
	SMS_RECEIPT_YOU_GAVE: {
		deDE: "Du hast %v an %v verliehen.",
		enUK: "You've given %v to %v.",
		esES: "Has dado %v a %v.",
		faIR: "شما پرداخت کرده اید %v به %v.",
		itIT: "Hai dato %v a %v",
		ruRU: "Вы дали %v - взял(а) %v.",
	},
	SMS_CLICK_TO_CONFIRM_OR_DECLINE: {
		deDE: "Klicke %v um zu akzeptieren oder abzulehnen.",
		enUK: "Click %v to confirm or decline.",
		esES: "Haz click %v para confirmar o rechazar.",
		faIR: "کلیک کنید %v تا رد خود را تایید نمایید",
		itIT: "Clicca %v per confermare o rifiutare",
		ruRU: "Нажмите %v чтобы подтвердить или отвергнуть.",
	},
	HTML_DATE: {
		deDE: "Datum",
		enUK: "Date",
		esES: "Fecha",
		faIR: "تاریخ",
		itIT: "Data",
		ruRU: "Дата",
	},
	HTML_RECEIPT: {
		deDE: "Empfänger",
		enUK: "Receipt",
		esES: "Recibo",
		faIR: "رسید",
		itIT: "Scontrino", //To upgrade, not the best translation from Russian
		ruRU: "Квитанция",
	},
	HTML_AMOUNT: {
		deDE: "Betrag",
		enUK: "Amount",
		esES: "Importe",
		faIR: "مقدار",
		itIT: "Totale",
		ruRU: "Сумма",
	},
	HTML_FROM: {
		deDE: "Von",
		enUK: "From",
		esES: "De",
		faIR: "از",
		itIT: "Da",
		ruRU: "Дал",
	},
	HTML_TO: {
		deDE: "An",
		enUK: "To",
		esES: "Para",
		faIR: "به",
		itIT: "Per",
		ruRU: "Получил",
	},
	NO_NAME: {
		deDE: "unbekannt",
		enUK: "no name",
		esES: "sin nombre",
		faIR: "بدون نام",
		itIT: "senza nome",
		ruRU: "без имени",
	},
	TELEGRAM_RECEIPT: {
		deDE: "{{.FromName}} hat einen Schuldschein erstellt ({{.TransferCurrency}})",
		enUK: "{{.FromName}} created a debt record ({{.TransferCurrency}})",
		esES: "{{.FromName}} ha creado una deuda ({{.TransferCurrency}})",
		faIR: "{{.FromName}} ایجاد یک سابقه بدهی ({{.TransferCurrency}})",
		itIT: "{{.FromName}} ha registrato un debito ({{.TransferCurrency}})",
		ruRU: "{{.FromName}} создал запись о долге ({{.TransferCurrency}})",
	},
	MESSAGE_TEXT_PLEASE_CHOOSE_FROM_OPTIONS_PROVIDED: {
		deDE: "Bitte wähle aus den angezeigten Optionen.",
		enUK: "Please choose from provided options.",
		enUS: "Please choose from provided options.",
		esES: "Por favor, elige una de las siguientes opciones.",
		faIR: "لطفاً از گزینه های ارائه شده انتخاب نمایید.",
		frFR: "Veuillez choisir parmi les options proposées.",
		idID: "Silakan pilih dari opsi yang disediakan.",
		itIT: "Scegli tra le opzioni fornite.",
		jaJP: "提供されたオプションから選択してください。",
		koKR: "제공된 옵션 중에서 선택하십시오.",
		plPL: "Proszę wybrać z podanych opcji.",
		ptBR: "Por favor, escolha entre as opções fornecidas.",
		ruRU: "Пожалуйста выберете из предоставленных опций.",
		trTR: "Lütfen sağlanan seçeneklerden birini seçin.",
		ukUA: "Будь ласка, виберіть із наданих опцій.",
		uzUZ: "Iltimos, taqdim etilgan variantlardan tanlang.",
		zhCN: "请从提供的选项中选择。",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE_OR_COMMENT: {
		deDE: "<b>Möchtest du eine Bemerkung oder Notiz hinzufügen?</b>\n%v Deine Notizen kannst nur du sehen.\n%v Eine Bemerkung wird quasi auf dem Schuldschein und der Quittung vermerkt und ist insofern für beide sichtbar.",
		enUK: "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for yoru own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		enUS: "<b>Do you want to add a note or comment?</b>\n%v Memos are private records for your own reference.\n%v Comments are available to everyone who has permission to view this transaction.",
		esES: "<b>¿Quieres añadir una nota o comentario?</b>\n%v Las notas se graban de manera privada para tu propia información.\n%v Los comentarios son visibles para todos los autorizados a ver esta transacción.",
		faIR: "<b>آیا میخواهید یادداشت یا شرحی اضافه کنید؟</b>\n%v یادداشت ها نوشته های خصوصی برای مراجعه خود شما هستند.\n%v شرح در دسترس تمام کسانی که مجاز رویت این تراکنش هستند میباشد.",
		frFR: "<b>Voulez-vous ajouter une note ou un commentaire?</b>\n%v Les mémos sont des enregistrements privés pour votre propre référence.\n%v Les commentaires sont accessibles à tous ceux qui ont la permission de voir cette transaction.",
		idID: "<b>Apakah Anda ingin menambahkan catatan atau komentar?</b>\n%v Memo adalah catatan pribadi untuk referensi Anda sendiri.\n%v Komentar tersedia untuk semua orang yang memiliki izin untuk melihat transaksi ini.",
		itIT: "<b>Vuoi aggiungere una nota o un commento?</b> \n%v I memo sono record privati per il riferimento di yoru.\n%v I commenti sono disponibili a tutti coloro che hanno l'autorizzazione a visualizzare questa transazione.",
		jaJP: "<b>メモやコメントを追加しますか？</b>\n%v メモはあなた自身の参照用のプライベートな記録です。\n%v コメントはこのトランザクションを閲覧する権限を持つすべての人が利用できます。",
		koKR: "<b>메모나 댓글을 추가하시겠습니까?</b>\n%v 메모는 자신의 참조를 위한 개인 기록입니다.\n%v 댓글은 이 거래를 볼 수 있는 권한이 있는 모든 사람이 이용할 수 있습니다.",
		plPL: "<b>Czy chcesz dodać notatkę lub komentarz?</b>\n%v Notatki są prywatnymi zapisami do własnego użytku.\n%v Komentarze są dostępne dla wszystkich, którzy mają uprawnienia do przeglądania tej transakcji.",
		ptBR: "<b>Deseja adicionar uma nota ou comentário?</b>\n%v Memorandos são registros privados para sua própria referência.\n%v Comentários estão disponíveis para todos que têm permissão para visualizar esta transação.",
		ruRU: "<b>Хотите добавить заметку или комментарий?</b>\n%v Заметки хранятся для вашего личного пользования.\n%v Комментарий виден всем кому разрешён просмотр этой транзакции.",
		trTR: "<b>Not veya yorum eklemek istiyor musunuz?</b>\n%v Notlar kendi referansınız için özel kayıtlardır.\n%v Yorumlar, bu işlemi görüntüleme izni olan herkes tarafından görülebilir.",
		ukUA: "<b>Хочете додати нотатку або коментар?</b>\n%v Нотатки зберігаються для вашого особистого користування.\n%v Коментар видно всім, кому дозволено перегляд цієї транзакції.",
		uzUZ: "<b>Eslatma yoki izoh qo'shmoqchimisiz?</b>\n%v Eslatmalar shaxsiy yozuvlar bo'lib, o'zingiz uchun ma'lumotdir.\n%v Izohlar ushbu tranzaksiyani ko'rish huquqiga ega bo'lgan har bir kishi uchun mavjud.",
		zhCN: "<b>您想添加备注或评论吗？</b>\n%v 备忘录是供您自己参考的私人记录。\n%v 评论对所有有权查看此交易的人可见。",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_NOTE: {
		deDE: "Bitte schreibe eine Notiz:",
		enUK: "Please write a note:",
		enUS: "Please write a note:",
		esES: "Por favor, escribe una nota:",
		faIR: "لطفاً یک یادداشت بنویسید:",
		frFR: "Veuillez écrire une note :",
		idID: "Silakan tulis catatan:",
		itIT: "Per favore scrivi un appunto:",
		jaJP: "メモを書いてください：",
		koKR: "메모를 작성해 주세요:",
		plPL: "Proszę napisać notatkę:",
		ptBR: "Por favor, escreva uma nota:",
		ruRU: "Напишите заметку:",
		trTR: "Lütfen bir not yazın:",
		ukUA: "Напишіть нотатку:",
		uzUZ: "Iltimos, eslatma yozing:",
		zhCN: "请写一个备注：",
	},
	COMMAND_TEXT_MORE_ABOUT_INTEREST_COMMAND: {
		deDE: "Mehr über Zinsen", // Updated from TODO
		enUK: "More about interest",
		enUS: "More about interest",
		esES: "Más sobre intereses",       // Updated from TODO
		faIR: "اطلاعات بیشتر درباره بهره", // Updated from TODO
		frFR: "Plus d'informations sur les intérêts",
		idID: "Lebih lanjut tentang bunga",
		itIT: "Più informazioni sugli interessi", // Updated from TODO
		jaJP: "利息についての詳細",
		koKR: "이자에 대한 자세한 정보",
		plPL: "Więcej o odsetkach",
		ptBR: "Mais sobre juros",
		ruRU: "Подробнее о процентах",
		trTR: "Faiz hakkında daha fazla",
		ukUA: "Детальніше про відсотки",
		uzUZ: "Foiz haqida ko'proq",
		zhCN: "更多关于利息的信息",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_SHORT: {
		deDE: `<b>Zinsen und Kommentar</b>

Um den Zinssatz und den Zeitraum festzulegen, senden Sie eine Nachricht im folgenden Format:
		<pre>prozent/prozent_zeitraum/min_zeitraum/karenzzeit:notiz</pre>`,
		enUK: `<b>Interest & notes</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`,
		enUS: `<b>Interest & notes</b>

To set interest rate & period send a message in following format:
		<pre>percent/percent_period/min_period/grace_period:note</pre>`,
		esES: `<b>Interés y comentario</b>

Para establecer la tasa de interés y el período, envía un mensaje en el siguiente formato:
		<pre>porcentaje/período_porcentaje/período_mínimo/período_gracia:nota</pre>`,
		faIR: `<b>نرخ بهره و یادداشت</b>

برای تنظیم نرخ بهره و دوره، پیامی به فرمت زیر ارسال کنید:
		<pre>درصد/دوره_درصد/دوره_حداقل/دوره_تنفس:یادداشت</pre>`,
		frFR: `<b>Intérêt et notes</b>

Pour définir le taux d'intérêt et la période, envoyez un message au format suivant:
		<pre>pourcentage/période_pourcentage/période_min/période_grâce:note</pre>`,
		idID: `<b>Bunga & catatan</b>

Untuk mengatur suku bunga & periode kirim pesan dalam format berikut:
		<pre>persen/periode_persen/periode_min/masa_tenggang:catatan</pre>`,
		itIT: `<b>Interessi e note</b>

Per impostare il tasso di interesse e il periodo, invia un messaggio nel seguente formato:
		<pre>percentuale/periodo_percentuale/periodo_minimo/periodo_grazia:nota</pre>`,
		jaJP: `<b>利息とメモ</b>

金利と期間を設定するには、次の形式でメッセージを送信してください：
		<pre>パーセント/パーセント期間/最小期間/猶予期間:メモ</pre>`,
		koKR: `<b>이자 및 메모</b>

이자율 및 기간을 설정하려면 다음 형식으로 메시지를 보내십시오:
		<pre>퍼센트/퍼센트_기간/최소_기간/유예_기간:메모</pre>`,
		plPL: `<b>Odsetki i notatki</b>

Aby ustawić stopę procentową i okres, wyślij wiadomość w następującym formacie:
		<pre>procent/okres_procentowy/min_okres/okres_karencji:notatka</pre>`,
		ptBR: `<b>Juros e notas</b>

Para definir a taxa de juros e o período, envie uma mensagem no seguinte formato:
		<pre>percentual/período_percentual/período_mínimo/período_carência:nota</pre>`,
		ruRU: `<b>Процент и комментарий</b>

Чтобы задать процент по долгу отправьте сообщение в следующем формате:
	<pre>процент/процентный_период/минимальный_период/грэйс_период:комментарий</pre>`,
		trTR: `<b>Faiz ve notlar</b>

Faiz oranını ve dönemi ayarlamak için aşağıdaki formatta bir mesaj gönderin:
		<pre>yüzde/yüzde_dönem/min_dönem/grace_dönem:not</pre>`,
		ukUA: `<b>Відсоток і коментар</b>

Щоб встановити відсоткову ставку і період, надішліть повідомлення в наступному форматі:
		<pre>відсоток/відсотковий_період/мінімальний_період/грейс_період:примітка</pre>`,
		uzUZ: `<b>Foiz va eslatmalar</b>

Foiz stavkasi va davrni o'rnatish uchun quyidagi formatda xabar yuboring:
		<pre>foiz/foiz_davri/min_davr/imtiyoz_davri:eslatma</pre>`,
		zhCN: `<b>利息和备注</b>

要设置利率和期限，请按以下格式发送消息：
		<pre>百分比/百分比期限/最小期限/宽限期:备注</pre>`,
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_INTEREST_LONG: {
		deDE: `<b>Zinsen & Notizen</b>

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
		enUK: `<b>Interest & notes</b>

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
		enUS: `<b>Interest & notes</b>

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
		esES: `<b>Interés y notas</b>

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
		faIR: `<b>نرخ بهره و یادداشت‌ها</b>

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
		frFR: `<b>Intérêts et notes</b>

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
		idID: `<b>Bunga & catatan</b>

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
		itIT: `<b>Interessi e note</b>

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
		plPL: `<b>Odsetki i notatki</b>

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
		ruRU: `<b>Процент и комментарий</b>

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
		trTR: `<b>Faiz ve notlar</b>

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
		uzUZ: `<b>Foiz va eslatmalar</b>

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
		deDE: "Falls du eine Bemerkung auf den Schuldschein schreiben willst, schick mir jetzt den Text.",
		enUK: `If you want to add a comment just send a text now.`,
		enUS: `If you want to add a comment just send a text now.`,
		esES: `Si quieres añadir un comentario simplemente envía un texto ahora.`,
		faIR: `شما می توانید یک شرح اضافه کنید. تنها کافیست یک متن ارسال کنید.`,
		frFR: `Si vous souhaitez ajouter un commentaire, envoyez simplement un texte maintenant.`,
		idID: `Jika Anda ingin menambahkan komentar, cukup kirim teks sekarang.`,
		itIT: `Se vuoi aggiungere un commento invia del testo ora.`,
		jaJP: `コメントを追加したい場合は、今すぐテキストを送信してください。`,
		koKR: `댓글을 추가하려면 지금 텍스트를 보내세요.`,
		plPL: `Jeśli chcesz dodać komentarz, po prostu wyślij tekst teraz.`,
		ptBR: `Se você quiser adicionar um comentário, basta enviar um texto agora.`,
		ruRU: `Если хотите добавить комментарий просто отправьте текст.`,
		trTR: `Yorum eklemek istiyorsanız şimdi bir metin gönderin.`,
		ukUA: `Якщо ви хочете додати коментар, просто надішліть текст зараз.`,
		uzUZ: `Izoh qo'shmoqchi bo'lsangiz, hozir matn yuboring.`,
		zhCN: `如果您想添加评论，请立即发送文本。`,
	},
	MESSAGE_TEXT_VISIBLE_TO_YOU_AND_COUNTERPARTY: {
		deDE: "sichtbar für dich & %v",
		enUK: "visible to you & %v",
		esES: "visible solo para ti & %v",
		faIR: "قابل مشاهده برای شما & %v",
		itIT: "visibile solo a te e %v",
		ruRU: "виден вам и %v",
	},
	MESSAGE_TEXT_TRANSFER_ASK_FOR_COMMENT: {
		deDE: "Schreibe nun eine Bemerkung auf den Schuldschein:",
		enUK: "Please write the comment:",
		esES: "Por favor, escribe un comentario:",
		faIR: "لطفاً شرح را ثبت کنید:",
		itIT: "Per favore scrivi un commento:",
		ruRU: "Напишите комментарий:",
	},
	MESSAGE_TEXT_TRANSFER_NOTE_ADDED_ASK_FOR_COMMENT: {
		deDE: "Deine Notiz wurde hinzugefügt, möchtest du noch eine Bemerkung auf den Schuldschein schreiben?",
		enUK: "Memo have been added. Do you want to write a comment?",
		esES: "La nota está añadida. ¿Quieres escribir un comentario?",
		faIR: "یادداشت اضافه شد. آیا میخواهید یک شرح ثبت کنید؟",
		itIT: "Promemoria aggiunto. Vuoi scrivere un commento?",
		ruRU: "Заметка добавлена. Хотите написать комментарий?",
	},
	MESSAGE_TEXT_TRANSFER_COMMENT_ADDED_ASK_FOR_NOTE: {
		deDE: "Die Bemerkung wurde vermerkt. Möchtest du noch eine Notiz für dich hinzufügen?",
		enUK: "Comment have been added. Do you want to write a note?",
		esES: "El comentario está añadido. ¿Quieres escribir una nota?",
		faIR: "شرح موردنظر شما ثبت شد. آیا می خواهید یک یادداشت بنویسید؟",
		itIT: "Commento aggiunto. Vuoi scrivere un appunto?",
		ruRU: "Комментарий добавлен. Хотите написать заметку?",
	},
	COMMAND_TEXT_NO_COMMENT_OR_NOTE_FOR_TRANSFER: {
		deDE: "Keine Notizen oder Bemerkungen",
		enUK: "Without notes or comments",
		esES: "Sin notas ni comentarios",
		faIR: "بدون یادداشت یا شرح",
		itIT: "Senza appunti o commenti",
		ruRU: "Без заметок и комментариев",
	},
	COMMAND_TEXT_NO_COMMENT_FOR_TRANSFER: {
		deDE: "Keine Bemerkungen",
		enUK: "No comments",
		esES: "Sin comentarios",
		faIR: "بدون شرح",
		itIT: "Nessun commento",
		ruRU: "Без комментариев",
	},
	COMMAND_TEXT_NO_NOTE_FOR_TRANSFER: {
		deDE: "Keine Notizen",
		enUK: "Without notes",
		esES: "Sin notas",
		faIR: "بدون یادداشت",
		itIT: "Senza appunti/note",
		ruRU: "Без заметок",
	},
	COMMAND_TEXT_ADD_NOTE_TO_TRANSFER: {
		deDE: "Notiz hinzufügen (privat)",
		enUK: "Add a note (private)",
		esES: "Añadir una nota (privada)",
		faIR: "یک یادداشت اضافه کنید (خصوصی)",
		itIT: "Aggiungi una nota (privata)",
		ruRU: "Добавить заметку",
	},
	COMMAND_TEXT_ADD_COMMENT_TO_TRANSFER: {
		deDE: "Bemerkung hinzufügen (öffentlich)",
		enUK: "Add a comment (public)",
		esES: "Añadir un comentario (público)",
		faIR: "یک شرح اضافه کنید (عمومی)",
		itIT: "Aggiungi un commento (pubblico)",
		ruRU: "Добавить комментарий",
	},
	DUE_IN_NOW: {
		deDE: "jetzt",
		enUK: "now",
		esES: "ahora",
		faIR: "حالا",
		itIT: "adesso",
		ruRU: "прямо сейчас",
	},
	DUE_IN_A_MINUTE: {
		deDE: "in einer Minute",
		enUK: "in a minute",
		esES: "en un minuto",
		faIR: "ظرف یک دقیقه",
		itIT: "in un minuto",
		ruRU: "через минуту",
	},
	DUE_IN_X_MINUTES: {
		deDE: "in %v Minuten",
		enUK: "in %v minutes",
		esES: "en %v minutos",
		faIR: "در %v دقیقه",
		itIT: "in %v minuti/o",
		ruRU: "через %v минут(ы)",
	},
	DUE_IN_AN_HOUR: {
		deDE: "in einer Stunde",
		enUK: "in an hour",
		esES: "en  una hora",
		faIR: "ظرف یک ساعت",
		itIT: "in un ora",
		ruRU: "через час",
	},
	DUE_IN_X_HOURS: {
		deDE: "in %v Stunde",
		enUK: "in %v hours",
		esES: "en %v horas",
		faIR: "در %v ساعت",
		itIT: "in %v ore/a",
		ruRU: "через %v часа/часов",
	},
	DUE_IN_X_DAYS: {
		deDE: "in %v Tagen",
		enUK: "in %v days",
		esES: "en %v días",
		faIR: "در %v روز",
		itIT: "in %v giorni/o",
		ruRU: "через %v дня/дней",
	},
	//-------------------------------------------------------------------------------------------------------------------
	WS_ALEX_T: {
		deDE: "Alexander Trakhimenok",
		enUK: "Alexander Trakhimenok",
		esES: "Alexander Trakhimenok",
		faIR: "الکساندر تراخیمِنوک",
		itIT: "Alexander Trakhimenok",
		ruRU: "Александр Трахимёнок",
	},

	WS_INDEX_TITLE: {
		ruRU: "DebtsTracker.io - программа для учёта личных долгов и активов",
		enUK: "DebtsTracker.io - an IOU app to track your personal debts & assets",
		esES: "DebtsTracker.io es una aplicación para el control de sus deudas personales",
		faIR: "DebtsTracker.io - برنامه ای برای ردیابی بدهی ها و دارایی های شما",
		plPL: "DebtsTracker.io - aplikacja do śledzenia osobistych długów",
		ptPT: "DebtsTracker.io - um aplicativo para controlar suas dívidas pessoais",
		deDE: "DebtsTracker.io - eine App, um Ihre persönlichen Schulden zu verfolgen",
		frFR: "DebtsTracker.io - une application pour suivre vos dettes personnelles",
		itIT: "DebtsTracker.io - un app per monitorare i tuoi debiti personali",
		koKR: "DebtsTracker.io 은 - 앱이 사용자의 개인 채무를 추적",
		jaJP: "DebtsTracker.io は - アプリはあなたの個人的な借金を追跡します",
		zhCN: "DebtsTracker.io - 一个应用程序来跟踪你的个人债务",
	},
	WS_LIVE_DEMO: {
		ruRU: "Демо версия online",
		enUK: "Demostración",
		esES: "Demo en vivo",
		faIR: "نسخه ی نمایشی زنده",
		plPL: "Demo na żywo",
		ptPT: "Demonstração ao vivo",
		deDE: "Live-Demo",
		frFR: "Démo en direct",
		itIT: "Demo online",
		koKR: "실시간 데모",
		jaJP: "ライブデモ",
		zhCN: "现场演示",
	},
	WS_INDEX_TG_BOT_H2: {
		ruRU: "Бот для Telegram",
		enUK: "Chat bot for Telegram messenger",
		esES: "Chat bot para Telegram",
		faIR: "ربات چت برای پیام رسان تلگرام",
		plPL: "Chat bot do telegramu posłańca",
		ptPT: "bot de bate-papo para Telegram messenger",
		deDE: "Chat-Bot für Telegram",
		frFR: "bot Chat for Telegram messenger",
		itIT: "Bot Chat per Telegram",
		koKR: "전보 메신저 채팅 봇",
		jaJP: "電報メッセンジャーのためのチャットボット",
		zhCN: "聊天机器人的电报使者",
	},
	WS_INDEX_TG_BOT_OPEN: {
		ruRU: "Открыть в Телеграмме &#x1F680;",
		enUK: "Open in Telegram &#x1F680;",
		esES: "Abrir en Telegram &#x1F680;",
		faIR: "بازکردن در تلگرام &#x1F680;",
		plPL: "Otwórz w telegramu &#x1F680;",
		ptPT: "Open in Telegram &#x1F680;",
		deDE: "Öffnen in Telegram &#x1F680;",
		frFR: "Open in Telegram &#x1F680;",
		itIT: "Apri su Telegram &#x1F680;",
		koKR: "전보 &#x1F680; 에서 열기;",
		jaJP: "電報 &#x1F680; で開きます。",
		zhCN: "打开在电报 &#x1F680;",
	},

	WS_INDEX_TG_BOT_P: {
		ruRU: "В настоящий момент наша программа доступна в мессенджере <a href='https://telegram.org/'>Телеграм</a>.",
		enUK: "At the moment our program is available just on <a href='https://telegram.org/'>Telegram messenger</a>",
		esES: "De momento nuestro programa está disponible sólo en <a href='https://telegram.org/'>Telegrama mensajero </a>",
		faIR: "درحال حاضر برنامه ما فقط در دسترس است در <a href='https://telegram.org/'>Телеграм</a>تلگرام",
		plPL: "W tej chwili nasz program jest dostępny tylko na <a href='https://telegram.org/'>Telegram messenger</a>",
		ptPT: "No momento em que o nosso programa está disponível apenas na <a href='https://telegram.org/'>Telegram messenger</a>",
		deDE: "Im Moment ist unser Programm nur auf <a href='https://telegram.org/'>Telegram verfügbar</a>",
		frFR: "Au moment de notre programme est disponible seulement sur <a href='https://telegram.org/'>Telegram messager</a>",
		itIT: "Al momento il nostro programma è disponibile solo su <a href='https://telegram.org/'>Telegram</a>",
		koKR: "지금이 순간 우리의 프로그램은 단지에 <a href='https://telegram.org/'>Telegram</a> 의 <b> 전보 </b>을 메신저 를 볼 수 있습니다",
		jaJP: "現時点では私たちのプログラムは、ちょうど上の<a href='https://telegram.org/'>Telegram</a>電報のメッセンジャーで提供されています",
		zhCN: "目前我们的计划是只提供在<a href='https://telegram.org/'>Telegram</a>电报的使者",
	},
	WS_MOTTO: {
		ruRU: "Платежи по долгам целиком и вовремя!",
		enUK: "Know your bottom line & never miss a debt payment!",
		esES: "¡Controle sus pagos y deudas!",
		faIR: "از سود و زیان خود مطلع باشید و هرگز پرداخت بدهی ای را از قلم نیندازید",
		plPL: "Znaj swoją równowagę i nigdy nie przegapisz zapłatę długu!",
		ptPT: "Conheça o seu equilíbrio e nunca perca um pagamento da dívida!",
		deDE: "Wissen, wem man wie viel schuldet!",
		frFR: "Apprenez à connaître votre solde et ne jamais manquer un paiement de la dette!",
		itIT: "Tieni sott'occhio il tuo bilancio e non dimenticarti mai di un debito!",
		koKR: "균형을 알고 및 채무 지불을 놓칠 수 없어!",
		jaJP: "あなたのバランスを知っている＆債務の支払いを見逃すことはありません！",
		zhCN: "了解天平＆不会错过任何一个债务付款！",
	},
	WS_SHORT_DESC: {
		ruRU: "DebtsTracker.io - мобильное приложение и сервис напоминаний для учёта и своевременной выплаты долгов. Отсылает автоматические уведомления вашим должникам по email и SMS.",
		enUK: "DebtsTracker.io is a mobile IOU app & a reminder service that helps to track your debts, credits & assets. Sends automated email & SMS reminders to your debtors.",
		esES: "DebtsTracker.io es un servicio de aplicaciones móviles y recordatorios que ayuda a realizar un seguimiento de sus deudas y créditos. Envía notificaciones por correo electrónico y SMS automatizados a sus deudores.",
		faIR: "DebtsTracker.io یک برنامه موبایل و سرویس یادآور می باشد که به شما کمک می کند تا بدهی ها و اعتبارات خود را ردیابی نمایید. همچنین ایمیل و پیام کوتاه یادآوری به بدهکاران ارسال می کند.",
		plPL: "DebtsTracker.io to aplikacje mobilne i przypomnienia usługa, która pozwala na śledzenie swoich długów i kredytów. Wysyła automatycznych powiadomień e-mail i SMS do swoich dłużników.",
		ptPT: "DebtsTracker.io é um serviço de aplicativos móveis e lembrete de que ajuda a controlar seus débitos e créditos. Envia e-mail e SMS notificações automáticas aos seus devedores.",
		deDE: "DebtsTracker.io ist eine mobile App, die beim Verwalten von persönlichen Schulden hilft - egal ob Sie Geld verleihen oder welches leihen. Sendet automatisierte E-Mail und SMS-Benachrichtigungen an Ihre Schuldner und Gläubiger.",
		frFR: "DebtsTracker.io est une des applications mobiles et rappel service qui permet de suivre vos dettes et crédits. Envoie automatisés email & SMS reminders à vos débiteurs.",
		itIT: "DebtsTracker.io è un servizio di applicazioni mobili che ricordare e aiuta a monitorare i debiti e crediti. Invia notifiche e-mail e SMS automatici ai i vostri debitori.",
		koKR: "DebtsTracker.io 은 채무 및 크레딧을 추적하는 데 도움이 모바일 앱 및 알림 서비스입니다. 당신의 채무자에 자동화 된 이메일 및 SMS 알림을 보냅니다.",
		jaJP: "DebtsTracker.io は、あなたの借金＆クレジットを追跡するのに役立ちますモバイルアプリ＆リマインダーサービスです。あなたの債務者に自動メール＆SMS通知を送信します。",
		zhCN: "DebtsTracker.io 是一个移动应用和提醒服务，帮助跟踪你的债务和信用。发送自动电子邮件和短信通知到您的债务人。",
	},

	WS_HELP_US_TITLE: {
		deDE: "Wie kann man beim DebtsTracker.io Projekt helfen kann",
		enUK: "How you can help to DebtsTracker.io project",
		esES: "Como puedes ayudar a DebtsTracker.io project",
		faIR: "چگونه می توانید به پروژه  DebtsTracker.io کمک کنید.",
		itIT: "Come potete aiutare il progetto DebtsTracker.io",
		ruRU: "Как вы можете помочь проекту DebtsTracker.io",
	},
	WS_ADS_TITLE: {
		deDE: "Werbung @ DebtsTracker.IO",
		enUK: "Ads @ DebtsTracker.IO",
		esES: "Anuncio @ DebtsTracker.IO",
		faIR: "تبلیغات @ DebtsTracker.IO",
		itIT: "Pubblicita' @ DebtsTracker.IO",
		ruRU: "Реклама на DebtsTracker.IO",
	},
	WS_ADS_CONTENT: {
		deDE: `Um Werbung in unserer App zu schalten, schick uns eine Mail an <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		enUK: `To place ads in our app please send us an email to <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		esES: `Para publicar un anuncio en nuestra app escríbenos un e-mail a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		faIR: `برای قراردادن تبلیغات در برنامه ما، درخواست خود را به این آدرس ایمیل فرمایید <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		itIT: `Per inserire della pubblicita nella nostra app inviaci un email a <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
		ruRU: `Чтобы разместить рекламу в нашем приложении пишите нам на <a href="mailto:ads@debtstracker.io">ads@debtstracker.io</a>.`,
	},
	FB_MAKE_RECORD_HEADER: {
		deDE: "Schuldschein",
		enUK: "Record debt",
		esES: "Registrar la deuda",
		faIR: "ثبت بدهی",
		itIT: "Registra il debito",
		ruRU: "Записать долг",
	},
	FB_MAKE_RECORD_HEADLINE: {
		deDE: "Scroll nach links, um weitere Optionen zu sehen.",
		enUK: "Scroll left to see other options.",
		esES: "Desliza a la izquierda para ver otras opciones",
		faIR: "برای دیدن سایر گزینه ها به سمت چپ اسکرول نمایید.",
		itIT: "Scrolla a sinistra per vedere altre opzioni",
		ruRU: "Пролистайте карточки влево чтобы увидеть другие опции.",
	},
	FB_HOW_ARE_THINGS_HEADER: {
		deDE: "Wie geht es dir?",
		enUK: "How are you doing?",
		esES: "¿Cómo va todo?",
		faIR: "حال شما چطوره؟",
		itIT: "Come te la passi?",
		ruRU: "Как идут дела?",
	},
	SNEATBOT_MSG_TXT_START: {
		enUK: `
<b>From bot's creator</b>: Hi %s!

@SneatBot helps to manage your day-to-day family life. Or you can create a space to manage your group/team/community.

I've spend lot's of time to make this bot useful, quick & reliable.I hope you'll like it.

You can learn about new features of the bot in @SneatApp channel where <a href="https://t.me/StarGiveaways_EN">we giveaway 500 🌟 EVERY month</a>.
`,
		ruRU: `
<b>От создателя бота:</b> Привет %s!

@SneatBot помогает организовать вашу семейную жизнь.
Так же можно создать пространство для управления группой/командой/сообществом.

Вы можете узнавать о новых возможностях бота в канале @SneatApp_ru где <a href="https://t.me/StarGiveaways_RU">мы разыгрываем 500 🌟 КАЖДЫЙ месяц</a>.
`,
		koKR: `<b>봇 제작자로부터</b>: 안녕하세요 %s!

@SneatBot은 일상의 가족 생활을 관리하는 데 도움이 됩니다. 또는 그룹/팀/커뮤니티를 관리할 공간을 만들 수 있습니다.

저는 이 봇을 유용하고 빠르고 안정적으로 만들기 위해 많은 시간을 투자했습니다. 마음에 들어 하시기를 바랍니다.

@SneatApp 채널에서 봇의 새로운 기능에 대해 알아볼 수 있습니다. <a href="https://t.me/StarGiveaways_EN">매월 500🌟를 경품으로 드립니다</a>.`,
		esES: `<b>Del creador del bot:</b> ¡Hola, %s!

@SneatBot te ayuda a gestionar tu vida familiar diaria. O puedes crear un espacio para gestionar tu grupo/equipo/comunidad.

He dedicado mucho tiempo a hacer que este bot sea útil, rápido y fiable. Espero que te guste.

Puedes conocer las nuevas funciones del bot en el canal de @SneatApp, donde <a href="https://t.me/StarGiveaways_EN">regalamos 500 🌟 CADA mes</a>.`,
		frFR: `<b>De la part du créateur du bot</b> : Bonjour %s !

@SneatBot vous aide à gérer votre vie de famille au quotidien. Vous pouvez également créer un espace pour gérer votre groupe/équipe/communauté.

J'ai passé beaucoup de temps à rendre ce bot utile, rapide et fiable. J'espère qu'il vous plaira.

Vous pouvez en apprendre davantage sur les nouvelles fonctionnalités du bot sur le canal @SneatApp où <a href="https://t.me/StarGiveaways_FR">nous offrons 500 🌟 CHAQUE mois</a>.`,
		itIT: `<b>Dal creatore del bot</b>: Ciao %s!

@SneatBot ti aiuta a gestire la tua vita familiare quotidiana. Oppure puoi creare uno spazio per gestire il tuo gruppo/team/community.

Ho dedicato molto tempo a rendere questo bot utile, veloce e affidabile. Spero che ti piaccia.

Puoi scoprire le nuove funzionalità del bot nel canale @SneatApp dove <a href="https://t.me/StarGiveaways_EN">regaliamo 500 🌟 OGNI mese</a>.`,
		jaJP: `<b>ボットの作成者より</b>: こんにちは、%s!

@SneatBot は、日々の家族生活の管理に役立ちます。または、グループ/チーム/コミュニティを管理するスペースを作成することもできます。

このボットを便利で、迅速で、信頼できるものにするために、多くの時間を費やしました。気に入っていただければ幸いです。

ボットの新機能については、@SneatApp チャンネルで確認できます。ここでは、<a href="https://t.me/StarGiveaways_EN">毎月 500 🌟 をプレゼント</a>しています。`,
		zhCN: `<b>来自机器人的创建者</b>：嗨 %s！

@SneatBot 帮助您管理日常家庭生活。或者您可以创建一个空间来管理您的群组/团队/社区。

我花了很多时间让这个机器人变得有用、快速和可靠。希望你会喜欢它。

您可以在 @SneatApp 频道了解该机器人的新功能，<a href="https://t.me/StarGiveaways_EN">我们每月赠送 500 🌟</a>。`,
		deDE: `<b>Vom Ersteller des Bots</b>: Hallo %s!

@SneatBot hilft Ihnen, Ihren Familienalltag zu organisieren. Oder Sie können einen Bereich erstellen, in dem Sie Ihre Gruppe/Ihr Team/Ihre Community verwalten können.

Ich habe viel Zeit darauf verwendet, diesen Bot nützlich, schnell und zuverlässig zu machen. Ich hoffe, er gefällt Ihnen.

Sie können sich über neue Funktionen des Bots im @SneatApp-Kanal informieren, wo wir <a href="https://t.me/StarGiveaways_EN">JEDEN Monat 500 🌟 verschenken</a>.`,
		ptPT: `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerir a sua vida familiar quotidiana. Ou pode criar um espaço para gerir o seu grupo/equipa/comunidade.

Gastei muito tempo para tornar este bot útil, rápido e fiável.

Pode conhecer as novas funcionalidades do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">distribuímos 500 🌟 TODOS os meses</a>.`,
		ptBR: `<b>Do criador do bot</b>: Olá %s!

@SneatBot ajuda a gerenciar sua vida familiar cotidiana. Ou você pode criar um espaço para gerenciar seu grupo/equipe/comunidade.

Eu gastei muito tempo para tornar este bot útil, rápido e confiável. Espero que você goste.

Você pode aprender sobre os novos recursos do bot no canal @SneatApp onde <a href="https://t.me/StarGiveaways_EN">nós doamos 500 🌟 TODO mês</a>.`,
		plPL: `<b>Od twórcy bota</b>: Cześć %s!

@SneatBot pomaga zarządzać codziennym życiem rodzinnym. Możesz też utworzyć przestrzeń do zarządzania swoją grupą/zespołem/społecznością.

Poświęciłem dużo czasu, aby ten bot był użyteczny, szybki i niezawodny. Mam nadzieję, że Ci się spodoba.

Możesz dowiedzieć się o nowych funkcjach bota na kanale @SneatApp, gdzie <a href="https://t.me/StarGiveaways_EN">rozdajemy 500 🌟 CO MIESIĄC</a>.`,
		ukUA: `<b>Від творця бота</b>: Привіт, %s!

@SneatBot допомагає керувати повсякденним сімейним життям. Або ви можете створити простір для керування своєю групою/командою/спільнотою.

Я витратив багато часу, щоб зробити цього бота корисним, швидким і надійним. Сподіваюся, він вам сподобається.

Ви можете дізнатися про нові функції бота на каналі @SneatApp, де <a href="https://t.me/StarGiveaways_EN">ми роздаємо 500 🌟 КОЖНОГО місяця</a>.`,
	},
	SpaceCmdText: {
		enUK: "Current space: %s <b>%s</b>",
		ruRU: "Текущее пространство: %s <b>%s</b>",
		esES: "Espacio actual: %s <b>%s</b>",
		faIR: "فضای فعلی: %s <b>%s</b>",
		itIT: "Spazio corrente: %s <b>%s</b>",
		deDE: "Aktueller Raum: %s <b>%s</b>",
		frFR: "Espace actuel: %s <b>%s</b>",
		plPL: "Aktualna przestrzeń: %s <b>%s</b>",
		ptPT: "Espaço atual: %s <b>%s</b>",
		koKR: "현재 공간: %s <b>%s</b>",
		jaJP: "現在のスペース: %s <b>%s</b>",
		zhCN: "当前空间: %s <b>%s</b>",
		ukUA: "Поточний простір: %s <b>%s</b>",
		ptBR: "Espaço atual: %s <b>%s</b>",
		trTR: "Mevcut alan: %s <b>%s</b>",
		idID: "Ruang saat ini: %s <b>%s</b>",
	},
	SpaceCmdBtnContacts: {
		enUK: "Contacts",
		ruRU: "Контакты",
		esES: "Contactos",
		faIR: "مخاطبین",
		itIT: "Contatti",
		deDE: "Kontakte",
		frFR: "Contacts",
		plPL: "Kontakty",
		ptPT: "Contatos",
		koKR: "연락처",
		jaJP: "連絡先",
		zhCN: "联系人",
		ukUA: "Контакти",
		ptBR: "Contatos",
		trTR: "Kişiler",
		idID: "Kontak",
	},
	SpaceCmdBtnMembers: {
		enUK: "Members",
		ruRU: "Участники",
		esES: "Miembros",
		faIR: "اعضا",
		itIT: "Membri",
		deDE: "Mitglieder",
		frFR: "Membres",
		plPL: "Członkowie",
		ptPT: "Membros",
		koKR: "회원",
		jaJP: "メンバー",
		zhCN: "成员",
		ukUA: "Учасники",
		ptBR: "Membros",
		trTR: "Üyeler",
		idID: "Anggota",
	},
	FamilyMembers: {
		enUK: "Family members",
		ruRU: "Члены семьи",
		esES: "Miembros de la familia",
		faIR: "اعضای خانواده",
		itIT: "Membri della famiglia",
		deDE: "Familienmitglieder",
		frFR: "Membres de la famille",
		plPL: "Członkowie rodziny",
		ptPT: "Membros da família",
		koKR: "가족 구성원",
		jaJP: "家族のメンバー",
		zhCN: "家庭成员",
		ukUA: "Члени сім'ї",
		ptBR: "Membros da família",
		trTR: "Aile üyeleri",
		idID: "Anggota keluarga",
	},
	SpaceMembers: {
		enUK: "Space members",
		ruRU: "Участники пространства",
		esES: "Miembros del espacio",
		faIR: "اعضای فضا",
		itIT: "Membri dello spazio",
		deDE: "Raummitglieder",
		frFR: "Membres de l'espace",
		plPL: "Członkowie przestrzeni",
		ptPT: "Membros do espaço",
		koKR: "공간 멤버",
		jaJP: "スペースメンバー",
		zhCN: "空间成员",
		ukUA: "Члени простору",
		ptBR: "Membros do espaço",
		trTR: "Alan üyeleri",
		idID: "Anggota ruang",
	},
	SpaceCmdBtnLists: {
		enUK: "Lists",
		ruRU: "Списки",
		esES: "Listas",
		faIR: "لیست ها",
		itIT: "Elenchi",
		deDE: "Listen",
		frFR: "Listes",
		plPL: "Listy",
		ptPT: "Listas",
		koKR: "목록",
		jaJP: "リスト",
		zhCN: "清单",
		ukUA: "Списки",
		ptBR: "Listas",
		trTR: "Listeler",
		idID: "Daftar",
	},
	SpaceCmdBtnAssets: {
		enUK: "Assets",
		ruRU: "Активы",
		esES: "Activos",
		faIR: "دارایی ها",
		itIT: "Attività",
		deDE: "Vermögenswerte",
		frFR: "Actifs",
		plPL: "Aktywa",
		ptPT: "Ativos",
		koKR: "자산",
		jaJP: "資産",
		zhCN: "资产",
		ukUA: "Активи",
		ptBR: "Ativos",
		trTR: "Varlıklar",
		idID: "Aset",
	},
	SpaceCmdBtnBudget: {
		enUK: "Budget",
		ruRU: "Бюджет",
		esES: "Presupuesto",
		faIR: "بودجه",
		itIT: "Budget",
		deDE: "Budget",
		frFR: "Budget",
		plPL: "Budżet",
		ptPT: "Orçamento",
		koKR: "예산",
		jaJP: "予算",
		zhCN: "预算",
		ukUA: "Бюджет",
		ptBR: "Orçamento",
		trTR: "Bütçe",
		idID: "Anggaran",
	},
	SpaceCmdBtnDebts: {
		enUK: "Debts",
		ruRU: "Долги",
		esES: "Deudas",
		faIR: "بدهی ها",
		itIT: "Debiti",
		deDE: "Schulden",
		frFR: "Dettes",
		plPL: "Długi",
		ptPT: "Dívidas",
		koKR: "부채",
		jaJP: "借金",
		zhCN: "债务",
		ukUA: "Борги",
		ptBR: "Dívidas",
		trTR: "Borçlar",
		idID: "Hutang",
	},
	SpaceCmdBtnCalendar: {
		enUK: "Calendar",
		ruRU: "Календарь",
		esES: "Calendario",
		faIR: "تقویم",
		itIT: "Calendario",
		deDE: "Kalender",
		frFR: "Calendrier",
		plPL: "Kalendarz",
		ptPT: "Calendário",
		koKR: "달력",
		jaJP: "カレンダー",
		zhCN: "日历",
		ukUA: "Календар",
		ptBR: "Calendário",
		trTR: "Takvim",
		idID: "Kalender",
	},
	SpaceCmdBtnTrackers: {
		deDE: "Tracker", // Placeholder
		enUK: "Trackers",
		enUS: "Trackers",      // Placeholder
		esES: "Rastreadores",  // Placeholder
		faIR: "ردیاب\u200cها", // Placeholder
		frFR: "Traqueurs",     // Placeholder
		idID: "Pelacak",       // Placeholder
		itIT: "Tracker",       // Placeholder
		jaJP: "トラッカー",         // Placeholder
		koKR: "추적기",           // Placeholder
		plPL: "Monitorujące",  // Placeholder
		ptBR: "Rastreadores",  // Placeholder
		ruRU: "Трекеры",
		trTR: "İzleyiciler",   // Placeholder
		ukUA: "Трекери",       // Placeholder
		uzUZ: "Kuzatuvchilar", // Placeholder
		zhCN: "追踪器",           // Placeholder
	},
	BtnSpaces: {
		enUK: "Spaces",
		ruRU: "Пространства",
		esES: "Espacios",
		faIR: "فضاها",
		itIT: "Spazi",
		deDE: "Räume",
		frFR: "Espaces",
		plPL: "Przestrzenie",
		ptPT: "Espaços",
		koKR: "공간",
		jaJP: "スペース",
		zhCN: "空间",
		ukUA: "Простори",
		ptBR: "Espaços",
		trTR: "Mekanlar",
		idID: "Ruang",
	},
	SpaceCmdBtnSettings: {
		enUK: "Settings",
		ruRU: "Настройки",
		esES: "Ajustes",
		faIR: "تنظیمات",
		itIT: "Impostazioni",
		deDE: "Einstellungen",
		frFR: "Paramètres",
		plPL: "Ustawienia",
		ptPT: "Configurações",
		koKR: "설정",
		jaJP: "設定",
		zhCN: "设置",
		ukUA: "Налаштування",
		ptBR: "Configurações",
		trTR: "Ayarlar",
		idID: "Pengaturan",
	},
	LIST_CMD_BUY: {
		enUK: "Buy",
		ruRU: "Купить",
		esES: "Comprar",
		faIR: "خرید",
		itIT: "Acquista",
		deDE: "Kaufen",
		frFR: "Acheter",
		plPL: "Kup",
		ptPT: "Comprar",
		koKR: "사다",
		jaJP: "購入",
		zhCN: "购买",
		ukUA: "Купити",
		ptBR: "Comprar",
		trTR: "Satın al",
		idID: "Beli",
	},
	LIST_CMD_TODO: {
		enUK: "ToDo",
		ruRU: "Задачи",
		esES: "Tareas",
		faIR: "وظایف",
		itIT: "Compiti",
		deDE: "Aufgaben",
		frFR: "Tâches",
		plPL: "Zadania",
		ptPT: "Tarefas",
		koKR: "할 일",
		jaJP: "タスク",
		zhCN: "任务",
		ukUA: "Завдання",
		ptBR: "Tarefas",
		trTR: "Görevler",
		idID: "Tugas",
	},
	LIST_CMD_WATCH: {
		enUK: "Watch",
		ruRU: "Смотреть",
		esES: "Ver",
		faIR: "تماشا کنید",
		itIT: "Guarda",
		deDE: "Ansehen",
		frFR: "Regarder",
		plPL: "Oglądaj",
		ptPT: "Assistir",
		koKR: "보기",
		jaJP: "見る",
		zhCN: "观看",
		ukUA: "Дивитися",
		ptBR: "Assistir",
		trTR: "İzle",
		idID: "Menonton",
	},
	LIST_CMD_READ: {
		enUK: "Read",
		ruRU: "Читать",
		esES: "Leer",
		faIR: "خواندن",
		itIT: "Leggi",
		deDE: "Lesen",
		frFR: "Lire",
		plPL: "Czytaj",
		ptPT: "Ler",
		koKR: "읽기",
		jaJP: "読む",
		zhCN: "读",
		ukUA: "Читати",
		ptBR: "Ler",
		trTR: "Oku",
		idID: "Baca",
	},
	ListCmdBtnToRead: {
		enUK: "To read",
		ruRU: "Прочитать",
		esES: "Leer",
		faIR: "برای خواندن",
		itIT: "Da leggere",
		deDE: "Zu lesen",
		frFR: "À lire",
		plPL: "Do przeczytania",
		ptPT: "Para ler",
		koKR: "읽을 것",
		jaJP: "読む",
		zhCN: "阅读",
		ukUA: "Читати",
		ptBR: "Para ler",
		trTR: "Okunacaklar",
		idID: "Untuk dibaca",
	},
	ListCmdBtnToWatch: {
		enUK: "To watch",
		ruRU: "Посмотреть",
		esES: "Ver",
		faIR: "برای تماشا",
		itIT: "Da guardare",
		deDE: "Ansehen",
		frFR: "À regarder",
		plPL: "Do obejrzenia",
		ptPT: "Para assistir",
		koKR: "볼 것",
		jaJP: "見る",
		zhCN: "观看",
		ukUA: "Дивитися",
		ptBR: "Para assistir",
		trTR: "İzlenecekler",
		idID: "Untuk ditonton",
	},
	ListCmdBtnToBuy: {
		enUK: "To buy",
		ruRU: "Купить",
		esES: "Comprar",
		faIR: "برای خرید",
		itIT: "Da comprare",
		deDE: "Zu kaufen",
		frFR: "À acheter",
		plPL: "Do kupienia",
		ptPT: "Para comprar",
		koKR: "구매할 것",
		jaJP: "購入",
		zhCN: "购买",
		ukUA: "Купити",
		ptBR: "Para comprar",
		trTR: "Alınacaklar",
		idID: "Untuk dibeli",
	},
	ListCmdBtnToDo: {
		enUK: "To do",
		ruRU: "Сделать",
		esES: "Hacer",
		faIR: "انجام دادن",
		itIT: "Da fare",
		deDE: "Zu tun",
		frFR: "À faire",
		plPL: "Do zrobienia",
		ptPT: "Para fazer",
		koKR: "할 일",
		jaJP: "やる",
		zhCN: "做",
		ukUA: "Зробити",
		ptBR: "Para fazer",
		trTR: "Yapılacaklar",
		idID: "Untuk dilakukan",
	},
	Groceries: {
		enUK: "Groceries",
		ruRU: "Продукты",
		esES: "Comestibles",
		faIR: "خوار و بار",
		itIT: "Generi alimentari",
		deDE: "Lebensmittel",
		frFR: "Épicerie",
		plPL: "Artykuły spożywcze",
		ptPT: "Comestíveis",
		koKR: "식료품",
		jaJP: "食料品",
		zhCN: "杂货",
		ukUA: "Продукти",
		ptBR: "Comestíveis",
		trTR: "Bakkaliye",
		idID: "Bahan makanan",
	},
	Books: {
		enUK: "Books",
		ruRU: "Книги",
		esES: "Libros",
		faIR: "کتاب ها",
		itIT: "Libri",
		deDE: "Bücher",
		frFR: "Livres",
		plPL: "Książki",
		ptPT: "Livros",
		koKR: "책",
		jaJP: "本",
		zhCN: "书籍",
		ukUA: "Книги",
		ptBR: "Livros",
		trTR: "Kitaplar",
		idID: "Buku",
	},
	Movies: {
		enUK: "Movies",
		ruRU: "Фильмы",
		esES: "Películas",
		faIR: "فیلم ها",
		itIT: "Film",
		deDE: "Filme",
		frFR: "Films",
		plPL: "Filmy",
		ptPT: "Filmes",
		koKR: "영화",
		jaJP: "映画",
		zhCN: "电影",
		ukUA: "Фільми",
		ptBR: "Filmes",
		trTR: "Filmler",
		idID: "Film",
	},
	Tasks: {
		enUK: "Tasks",
		ruRU: "Задачи",
		esES: "Tareas",
		faIR: "وظایف",
		itIT: "Compiti",
		deDE: "Aufgaben",
		frFR: "Tâches",
		plPL: "Zadania",
		ptPT: "Tarefas",
		koKR: "할 일",
		jaJP: "タスク",
		zhCN: "任务",
		ukUA: "Завдання",
		ptBR: "Tarefas",
		trTR: "Görevler",
		idID: "Tugas",
	},
	ListsOfFamily: {
		enUK: "Family lists",
		ruRU: "Семейные списки",
		esES: "Listas familiares",
		faIR: "لیست های خانواده",
		itIT: "Liste familiari",
		deDE: "Familienlisten",
		frFR: "Listes familiales",
		plPL: "Listy rodzinne",
		ptPT: "Listas familiares",
		koKR: "가족 목록",
		jaJP: "家族のリスト",
		zhCN: "家庭清单",
		ukUA: "Сімейні списки",
		ptBR: "Listas familiares",
		trTR: "Aile listeleri",
		idID: "Daftar keluarga",
	},
	ListsOfPrivate: {
		enUK: "Private lists",
		ruRU: "Личные списки",
		esES: "Listas privadas",
		faIR: "لیست های خصوصی",
		itIT: "Liste private",
		deDE: "Private Listen",
		frFR: "Listes privées",
		plPL: "Prywatne listy",
		ptPT: "Listas privadas",
		koKR: "개인 목록",
		jaJP: "プライベートリスト",
		zhCN: "私人清单",
		ukUA: "Особисті списки",
		ptBR: "Listas privadas",
		trTR: "Özel listeler",
		idID: "Daftar pribadi",
	},
	ListsOfSpace: {
		deDE: "Listen @ %s",
		enUK: "Lists @ %s",
		enUS: "Lists @ %s",
		esES: "Listas @ %s",
		faIR: "لیست ها @ %s",
		frFR: "Listes @ %s",
		idID: "Daftar @ %s",
		itIT: "Liste @ %s",
		jaJP: "リスト @ %s",
		koKR: "목록 @ %s",
		plPL: "Listy @ %s",
		ptBR: "Listas @ %s",
		ptPT: "Listas @ %s",
		ruRU: "Списки @ %s",
		trTR: "Listeler @ %s",
		ukUA: "Списки @ %s",
		uzUZ: "Ro'yxatlar @ %s",
		zhCN: "清单 @ %s",
	},
	BtnBackToSpace: {
		enUK: "Back to space",
		ruRU: "К пространству",
		esES: "Volver al espacio",
		frFR: "Retour à l'espace",
		deDE: "Zurück zum Bereich",
		itIT: "Torna allo spazio",
		ptPT: "Voltar ao espaço",
		ptBR: "Voltar ao espaço",
		zhCN: "返回空间",
		jaJP: "スペースに戻る",
		koKR: "공간으로 돌아가기",
		plPL: "Powrót do przestrzeni",
		ukUA: "Назад до простору",
		trTR: "Alana geri dön",
		idID: "Kembali ke ruang",
		faIR: "بازگشت به فضا",
	},
	BtnPrivate: {
		enUK: "Private",
		ruRU: "Личное",
		esES: "Privado",
		faIR: "شخصی",
		itIT: "Privato",
		deDE: "Privat",
		frFR: "Privé",
		plPL: "Prywatne",
		ptPT: "Privado",
		koKR: "개인적인",
		jaJP: "個人",
		zhCN: "私人",
		ukUA: "Особисте",
		trTR: "Özel",
		idID: "Pribadi",
	},
	BtnFamily: {
		enUK: "Family",
		ruRU: "Семья",
		esES: "Familia",
		faIR: "خانواده",
		itIT: "Famiglia",
		deDE: "Familie",
		frFR: "Famille",
		plPL: "Rodzina",
		ptPT: "Família",
		koKR: "가족",
		jaJP: "家族",
		zhCN: "家庭",
		ukUA: "Сім'я",
		ptBR: "Família",
		trTR: "Aile",
		idID: "Keluarga",
	},
	TrackerPushUps: {
		enUK: "Push-ups",
		ruRU: "Отжимания",
		esES: "Flexiones",
		faIR: "پرسه",
		itIT: "Flessioni",
		deDE: "Liegestütze",
		frFR: "Pompes",
		plPL: "Pompki",
		ptPT: "Flexões",
		koKR: "푸시업",
		jaJP: "腕立て伏せ",
		zhCN: "俯卧撑",
		ukUA: "Віджимання",
		ptBR: "Flexões",
		trTR: "Şınav",
		idID: "Push-up",
	},
	TrackerPullUps: {
		enUK: "Pull-ups",
		ruRU: "Подтягивания",
		esES: "Dominadas",
		faIR: "کشیدن",
		itIT: "Trazioni",
		deDE: "Klimmzüge",
		frFR: "Tirages",
		plPL: "Podciągania",
		ptPT: "Pull-ups",
		koKR: "풀업",
		jaJP: "引き上げ",
		zhCN: "引体向上",
		ukUA: "Підтягування",
		ptBR: "Pull-ups",
		trTR: "Çekme",
		idID: "Pull-up",
	},
	TrackerSquats: {
		enUK: "Squats",
		ruRU: "Приседания",
		esES: "Sentadillas",
		faIR: "کرنچ",
		itIT: "Squat",
		deDE: "Kniebeugen",
		frFR: "Squats",
		plPL: "Przysiady",
		ptPT: "Agachamentos",
		koKR: "스쿼트",
		jaJP: "スクワット",
		zhCN: "深蹲",
		ukUA: "Присідання",
		ptBR: "Agachamentos",
		trTR: "Squat",
		idID: "Squat",
	},
	TrackerJumpingJacks: {
		enUK: "Jumping jacks",
		ruRU: "Прыжки на месте",
		esES: "Saltos de tijera",
		faIR: "حرکت پروانه",
		itIT: "Salti con apertura",
		deDE: "Hampelmänner",
		frFR: "Jumping jacks",
		plPL: "Pajacyki",
		ptPT: "Saltos de estrela",
		koKR: "팔 벌려 뛰기",
		jaJP: "ジャンピングジャック",
		zhCN: "开合跳",
		ukUA: "Стрибки з розмахом рук",
		ptBR: "Polichinelos",
		trTR: "Zıplama hareketi",
		idID: "Lompat bintang",
	},
	TrackerFuelCost: {
		enUK: "Fuel Cost",
		ruRU: "Стоимость топлива",
		esES: "Costo del combustible",
		faIR: "هزینه سوخت",
		itIT: "Costo del carburante",
		deDE: "Kraftstoffkosten",
		frFR: "Coût du carburant",
		plPL: "Koszt paliwa",
		ptPT: "Custo do combustível",
		koKR: "연료 비용",
		jaJP: "燃料費",
		zhCN: "燃料成本",
		ukUA: "Вартість палива",
		ptBR: "Custo do combustível",
		trTR: "Yakıt Maliyeti",
		idID: "Biaya bahan bakar",
	},
	TrackerFuelVolume: {
		enUK: "Fuel Volume",
		ruRU: "Объем топлива",
		esES: "Volumen de combustible",
		faIR: "حجم سوخت",
		itIT: "Volume di carburante",
		deDE: "Kraftstoffvolumen",
		frFR: "Volume de carburant",
		plPL: "Objętość paliwa",
		ptPT: "Volume de combustível",
		koKR: "연료 양",
		jaJP: "燃料量",
		zhCN: "燃料容量",
		ukUA: "Об'єм палива",
		ptBR: "Volume de combustível",
		trTR: "Yakıt Hacmi",
		idID: "Volume bahan bakar",
	},
	TrackerMileage: {
		enUK: "Mileage",
		ruRU: "Пробег",
		esES: "Kilometraje",
		faIR: "مسافت پیموده شده",
		itIT: "Chilometraggio",
		deDE: "Kilometerstand",
		frFR: "Kilométrage",
		plPL: "Przebieg",
		ptPT: "Quilometragem",
		koKR: "주행 거리",
		jaJP: "走行距離",
		zhCN: "里程",
		ukUA: "Пробіг",
		ptBR: "Quilometragem",
		trTR: "Kilometre",
		idID: "Jarak tempuh",
	},
	TrackerHeight: {
		enUK: "Height",
		ruRU: "Рост",
		esES: "Altura",
		faIR: "قد",
		itIT: "Altezza",
		deDE: "Höhe",
		frFR: "Taille",
		plPL: "Wzrost",
		ptPT: "Altura",
		koKR: "키",
		jaJP: "身長",
		zhCN: "身高",
		ukUA: "Зріст",
		ptBR: "Altura",
		trTR: "Boy",
		idID: "Tinggi",
	},
	TrackerWeight: {
		enUK: "Weight",
		ruRU: "Вес",
		esES: "Peso",
		faIR: "وزن",
		itIT: "Peso",
		deDE: "Gewicht",
		frFR: "Poids",
		plPL: "Waga",
		ptPT: "Peso",
		koKR: "몸무게",
		jaJP: "体重",
		zhCN: "体重",
		ukUA: "Вага",
		ptBR: "Peso",
		trTR: "Ağırlık",
		idID: "Berat",
	},
	TrackerSpending: {
		enUK: "Spending",
		ruRU: "Расходы",
		esES: "Gastos",
		faIR: "هزینه\u200cها",
		itIT: "Spese",
		deDE: "Ausgaben",
		frFR: "Dépenses",
		plPL: "Wydatki",
		ptPT: "Despesas",
		koKR: "지출",
		jaJP: "支出",
		zhCN: "支出",
		ukUA: "Витрати",
		ptBR: "Despesas",
		trTR: "Harcamalar",
		idID: "Pengeluaran",
	},
	TrackerBloodPressure: {
		enUK: "Blood Pressure",
		ruRU: "Кровяное давление",
		esES: "Presión arterial",
		faIR: "فشار خون",
		itIT: "Pressione sanguigna",
		deDE: "Blutdruck",
		frFR: "Pression artérielle",
		plPL: "Ciśnienie krwi",
		ptPT: "Pressão arterial",
		koKR: "혈압",
		jaJP: "血圧",
		zhCN: "血压",
		ukUA: "Кров'яний тиск",
		ptBR: "Pressão arterial",
		trTR: "Kan basıncı",
		idID: "Tekanan darah",
	},
	TrackerCategoryHealth: {

		enUK: "Health",
		ruRU: "Здоровье",
		esES: "Salud",
		faIR: "سلامتی",
		itIT: "Salute",
		deDE: "Gesundheit",
		frFR: "Santé",
		plPL: "Zdrowie",
		ptPT: "Saúde",
		koKR: "건강",
		jaJP: "健康",
		zhCN: "健康",
		ukUA: "Здоров'я",
		ptBR: "Saúde",
		trTR: "Sağlık",
		idID: "Kesehatan",
	},
	TrackerCategoryFitness: {
		enUK: "Fitness",
		ruRU: "Фитнес",
		esES: "Aptitud física",
		faIR: "تناسب اندام",
		itIT: "Fitness",
		deDE: "Fitness",
		frFR: "Forme physique",
		plPL: "Fitness",
		ptPT: "Aptidão física",
		koKR: "피트니스",
		jaJP: "フィットネス",
		zhCN: "健身",
		ukUA: "Фітнес",
		ptBR: "Aptidão física",
		trTR: "Fitness",
		idID: "Kebugaran fisik",
	},
	TrackerCategoryVehicle: {
		enUK: "Vehicle",
		ruRU: "Транспортное средство",
		esES: "Vehículo",
		faIR: "وسیله نقلیه",
		itIT: "Veicolo",
		deDE: "Fahrzeug",
		frFR: "Véhicule",
		plPL: "Pojazd",
		ptPT: "Veículo",
		koKR: "차량",
		jaJP: "車両",
		zhCN: "车辆",
		ukUA: "Транспортний засіб",
		ptBR: "Veículo",
		trTR: "Araç",
		idID: "Kendaraan",
	},
	CustomTracker: {
		enUK: "Custom tracker",
		ruRU: "Свой трекер",
		esES: "Rastreador personalizado",
		faIR: "ردیاب سفارشی",
		itIT: "Tracker personalizzato",
		deDE: "Benutzerdefinierter Tracker",
		frFR: "Tracker personnalisé",
		plPL: "Niestandardowy tracker",
		ptPT: "Rastreador personalizado",
		koKR: "사용자 정의 추적기",
		jaJP: "カスタムトラッカー",
		zhCN: "自定义追踪器",
		ukUA: "Свій трекер",
		ptBR: "Rastreador personalizado",
		trTR: "Özel izleyici",
		idID: "Pelacak kustom",
	},
	BackToTrackers: {
		enUK: "Back to trackers",
		ruRU: "Назад к трекерам",
		esES: "Volver a los rastreadores",
		faIR: "بازگشت به ردیاب\u200cها",
		itIT: "Torna ai tracker",
		deDE: "Zurück zu Trackern",
		frFR: "Retour aux trackers",
		plPL: "Powrót do trackerów",
		ptPT: "Voltar aos rastreadores",
		koKR: "트래커로 돌아가기",
		jaJP: "トラッカーに戻る",
		zhCN: "返回追踪器",
		ukUA: "Повернутись до трекерів",
		ptBR: "Voltar aos rastreadores",
		trTR: "İzleyicilere geri dön",
		idID: "Kembali ke pelacak",
	},
	AddTracker: {
		enUK: "Add tracker",
		ruRU: "Добавить трекер",
		esES: "Añadir rastreador",
		faIR: "افزودن ردیاب",
		itIT: "Aggiungi tracker",
		deDE: "Tracker hinzufügen",
		frFR: "Ajouter un tracker",
		plPL: "Dodaj tracker",
		ptPT: "Adicionar rastreador",
		koKR: "추적기 추가",
		jaJP: "トラッカーを追加",
		zhCN: "添加追踪器",
		ukUA: "Додати трекер",
		ptBR: "Adicionar rastreador",
		trTR: "İzleyici ekle",
		idID: "Tambahkan pelacak",
	},
	ShareTracker: {
		enUK: "Share tracker",
		ruRU: "Поделиться трекером",
		esES: "Compartir rastreador",
		faIR: "اشتراک\u200cگذاری ردیاب",
		itIT: "Condividi il tracker",
		deDE: "Tracker teilen",
		frFR: "Partager le tracker",
		plPL: "Udostępnij tracker",
		ptPT: "Partilhar rastreador",
		koKR: "추적기 공유",
		jaJP: "トラッカーを共有",
		zhCN: "分享追踪器",
		ukUA: "Поділитись трекером",
		ptBR: "Compartilhar rastreador",
		trTR: "İzleyiciyi paylaş",
		idID: "Bagikan pelacak",
	},
	RecordTrackerValue: {
		enUK: "Record value",
		ruRU: "Записать значение",
		esES: "Registrar valor",
		faIR: "ثبت مقدار",
		itIT: "Registra valore",
		deDE: "Wert aufzeichnen",
		frFR: "Enregistrer la valeur",
		plPL: "Zapisz wartość",
		ptPT: "Registrar valor",
		koKR: "값 기록",
		jaJP: "値を記録",
		zhCN: "记录值",
		ukUA: "Записати значення",
		ptBR: "Registrar valor",
		trTR: "Değeri kaydet",
		idID: "Catat nilai",
	},
	FamilyTrackers: {
		enUK: "Family trackers",
		ruRU: "Семейные трекеры",
		esES: "Rastreadores familiares",
		faIR: "ردیاب\u200cهای خانواده",
		itIT: "Tracker familiari",
		deDE: "Familientracker",
		frFR: "Trackers familiaux",
		plPL: "Trackery rodzinne",
		ptPT: "Rastreadores de família",
		koKR: "가족 추적기",
		jaJP: "家族のトラッカー",
		zhCN: "家庭追踪器",
		ukUA: "Сімейні трекери",
		ptBR: "Rastreadores de família",
		trTR: "Aile izleyicileri",
		idID: "Pelacak keluarga",
	},
	NoActiveTrackers: {
		enUK: "Currently you do not have any active trackers. Add some to start tracking your progress.",
		ruRU: "В настоящее время у вас нет активных трекеров. Добавьте их, чтобы начать отслеживать ваш прогресс.",
		esES: "Actualmente, no tienes rastreadores activos. Agrega algunos para comenzar a rastrear tu progreso.",
		faIR: "در حال حاضر هیچ ردیابی فعال ندارید. چند مورد اضافه کنید تا پیشرفت خود را پیگیری کنید.",
		itIT: "Attualmente non hai tracker attivi. Aggiungine alcuni per iniziare a monitorare i tuoi progressi.",
		deDE: "Derzeit haben Sie keine aktiven Tracker. Fügen Sie einige hinzu, um Ihren Fortschritt zu verfolgen.",
		frFR: "Vous n'avez actuellement aucun tracker actif. Ajoutez-en pour commencer à suivre vos progrès.",
		plPL: "Obecnie nie masz żadnych aktywnych trackerów. Dodaj kilka, aby zacząć śledzić swoje postępy.",
		ptPT: "Atualmente, não tens rastreadores ativos. Adiciona alguns para começar a monitorizar o teu progresso.",
		koKR: "현재 활성화된 추적기가 없습니다. 진행 상황을 추적하려면 몇 가지를 추가하세요.",
		jaJP: "現在、有効なトラッカーはありません。進捗を追跡するには、いくつか追加してください。",
		zhCN: "目前你还没有任何活跃的追踪器。添加一些以开始追踪你的进度。",
		ukUA: "Наразі у вас немає активних трекерів. Додайте кілька, щоб почати відстежувати свій прогрес.",
		ptBR: "Atualmente, você não tem rastreadores ativos. Adicione alguns para começar a rastrear seu progresso.",
		trTR: "Şu anda aktif izleyiciniz yok. İlerlemenizi izlemek için birkaç tane ekleyin.",
		idID: "Saat ini Anda tidak memiliki pelacak aktif. Tambahkan beberapa untuk mulai melacak kemajuan Anda.",
	},
	OneActiveTracker: {
		enUK: "You have one active tracker. Select it to record progress. Or add a new one to track more.",
		ruRU: "У вас есть один активный трекер. Выберите его, чтобы записать прогресс, или добавьте новый для отслеживания большего.",
		esES: "Tienes un rastreador activo. Selecciónalo para registrar el progreso o agrega uno nuevo para rastrear más.",
		faIR: "شما یک ردیاب فعال دارید. آن را انتخاب کنید تا پیشرفت خود را ثبت کنید یا یک ردیاب جدید اضافه کنید تا موارد بیشتری را پیگیری کنید.",
		itIT: "Hai un tracker attivo. Selezionalo per registrare i progressi o aggiungi uno nuovo per monitorare di più.",
		deDE: "Sie haben einen aktiven Tracker. Wählen Sie ihn aus, um Fortschritte aufzuzeichnen, oder fügen Sie einen neuen hinzu, um mehr zu verfolgen.",
		frFR: "Vous avez un tracker actif. Sélectionnez-le pour enregistrer vos progrès ou ajoutez-en un nouveau pour en suivre davantage.",
		plPL: "Masz jeden aktywny tracker. Wybierz go, aby zarejestrować postępy, lub dodaj nowy, aby śledzić więcej.",
		ptPT: "Tens um rastreador ativo. Seleciona-o para registar o progresso, ou adiciona um novo para monitorizar mais.",
		koKR: "활성화된 추적기 하나가 있습니다. 선택하여 진행 상황을 기록하거나 더 많이 추적하려면 새로 추가하세요.",
		jaJP: "アクティブなトラッカーが1つあります。選択して進捗を記録するか、新しいものを追加してさらに追跡してください。",
		zhCN: "你有一个活跃的追踪器。选择它来记录你的进度，或者添加一个新的来追踪更多。",
		ukUA: "У вас є один активний трекер. Виберіть його, щоб записати прогрес, або додайте новий для відстеження більшого.",
		ptBR: "Você tem um rastreador ativo. Selecione-o para registrar o progresso ou adicione um novo para rastrear mais.",
		trTR: "Bir adet aktif izleyiciniz var. İlerlemenizi kaydetmek için bunu seçin veya daha fazla şey izlemek için yeni bir tane ekleyin.",
		idID: "Anda memiliki satu pelacak aktif. Pilih untuk mencatat kemajuan atau tambahkan yang baru untuk melacak lebih banyak.",
	},
	NActiveTrackers: {
		enUK: "You have few active tracker. Select one to record progress.",
		ruRU: "У вас несколько активных трекеров. Выберите один, чтобы записать прогресс.",
		esES: "Tienes varios rastreadores activos. Selecciona uno para registrar el progreso.",
		faIR: "چند ردیاب فعال دارید. یکی را انتخاب کنید تا پیشرفت خود را ثبت کنید.",
		itIT: "Hai alcuni tracker attivi. Selezionane uno per registrare i progressi.",
		deDE: "Sie haben einige aktive Tracker. Wählen Sie einen aus, um Fortschritte aufzuzeichnen.",
		frFR: "Vous avez plusieurs trackers actifs. Sélectionnez-en un pour enregistrer vos progrès.",
		plPL: "Masz kilka aktywnych trackerów. Wybierz jeden, aby zarejestrować postępy.",
		ptPT: "Tens vários rastreadores ativos. Seleciona um para registar o progresso.",
		koKR: "활성화된 추적기 몇 개가 있습니다. 하나를 선택하여 진행 상황을 기록하세요.",
		jaJP: "複数のアクティブなトラッカーがあります。1つ選択して進捗を記録してください。",
		zhCN: "你有几个活跃的追踪器。选择一个来记录你的进度。",
		ukUA: "У вас є кілька активних трекерів. Виберіть один, щоб записати прогрес.",
		ptBR: "Você tem vários rastreadores ativos. Selecione um para registrar o progresso.",
		trTR: "Birkaç aktif izleyiciniz var. İlerlemenizi kaydetmek için birini seçin.",
		idID: "Anda memiliki beberapa pelacak aktif. Pilih salah satu untuk mencatat kemajuan.",
	},
	HintForIntTracker: {
		enUK: "Send an integer number to record it to the tracker",
		ruRU: "Отправьте целое число, чтобы записать его в трекер",
		esES: "Envía un número entero para registrarlo en el rastreador",
		faIR: "یک عدد صحیح ارسال کنید تا آن را در ردیاب ثبت کنید",
		itIT: "Invia un numero intero per registrarlo nel tracker",
		deDE: "Senden Sie eine Ganzzahl, um sie im Tracker zu speichern",
		frFR: "Envoyez un nombre entier pour l'enregistrer dans le tracker",
		plPL: "Wyślij liczbę całkowitą, aby zapisać ją w trackerze",
		ptPT: "Envia um número inteiro para o registar no rastreador",
		koKR: "정수 값을 추적기에 기록하려면 전송하세요",
		jaJP: "整数値を送信してトラッカーに記録してください",
		zhCN: "发送一个整数以将其记录到追踪器中",
		ukUA: "Надішліть ціле число, щоб записати його в трекер",
		ptBR: "Envie um número inteiro para registrá-lo no rastreador",
		trTR: "Bir tamsayı göndererek bunu izleyiciye kaydedin",
		idID: "Kirim bilangan bulat untuk mencatatnya ke pelacak",
	},
	Tracker: {
		enUK: "Tracker",
		ruRU: "Трекер",
		esES: "Rastreador",
		faIR: "ردیاب",
		itIT: "Tracker",
		deDE: "Tracker",
		frFR: "Tracker",
		plPL: "Tracker",
		ptPT: "Rastreador",
		koKR: "추적기",
		jaJP: "トラッカー",
		zhCN: "追踪器",
		ukUA: "Трекер",
		ptBR: "Rastreador",
		trTR: "İzleyici",
		idID: "Pelacak",
	},
	HintToShareTracker: {
		enUK: `\n\nYou can share this tracker either
	🤫 with a friend (<i>a single person would be able to accept</i>)
	🌍 or publicly (<i>anyone with a link can accept, you can cancel it at any time</i>)`,

		ruRU: `\n\nВы можете поделиться этим трекером либо
🤫 с другом (<i>только один человек сможет принять</i>)
🌍 или публично (<i>любой с ссылкой может принять, вы можете отменить это в любое время</i>)`,
		esES: `\n\nPuedes compartir este rastreador
🤫 con un amigo (<i>una sola persona puede aceptarlo</i>)
🌍 o públicamente (<i>cualquiera con el enlace puede aceptarlo, puedes cancelarlo en cualquier momento</i>)`,
		faIR: `\n\nشما می\u200cتوانید این ردیاب را به دو صورت به اشتراک بگذارید
🤫 با یک دوست (<i>فقط یک نفر می\u200cتواند بپذیرد</i>)
🌍 یا به صورت عمومی (<i>هرکسی با لینک می\u200cتواند بپذیرد، شما می\u200cتوانید در هر زمان آن را لغو کنید</i>)`,
		itIT: `\n\nPuoi condividere questo tracker
🤫 con un amico (<i>una sola persona potrà accettarlo</i>)
🌍 oppure pubblicamente (<i>chiunque con il link potrà accettarlo, puoi annullarlo in qualsiasi momento</i>)`,
		deDE: `\n\nSie können diesen Tracker entweder
🤫 mit einem Freund teilen (<i>nur eine Person kann ihn akzeptieren</i>)
🌍 oder öffentlich (<i>jeder mit dem Link kann ihn akzeptieren, Sie können es jederzeit abbrechen</i>)`,
		frFR: `\n\nVous pouvez partager ce tracker
🤫 avec un ami (<i>une seule personne pourra l'accepter</i>)
🌍 ou publiquement (<i>n'importe qui avec un lien peut l'accepter, vous pouvez l'annuler à tout moment</i>)`,
		plPL: `\n\nMożesz udostępnić ten tracker
🤫 znajomemu (<i>tylko jedna osoba będzie mogła przyjąć</i>)
🌍 lub publicznie (<i>każdy z linkiem może przyjąć, możesz to anulować w dowolnym momencie</i>)`,
		ptPT: `\n\nPodes partilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa poderá aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, podes cancelá-lo a qualquer momento</i>)`,
		koKR: `\n\n이 추적기를 공유할 수 있습니다
🤫 친구와 (<i>단 한 사람만 수락 가능</i>)
🌍 또는 공개적으로 (<i>링크를 가진 모든 사람이 수락할 수 있으며, 언제든 취소할 수 있습니다</i>)`,
		jaJP: `\n\nこのトラッカーを以下の方法で共有できます
🤫 友人と (<i>受け取れるのは1人だけ</i>)
🌍 または公開で (<i>リンクを持っている誰でも受け取れます。いつでもキャンセルできます</i>)`,
		zhCN: `\n\n您可以通过以下方式共享此追踪器
🤫 与朋友共享 (<i>只有一个人可以接受</i>)
🌍 或公开共享 (<i>任何拥有链接的人都可以接受，您可以随时取消</i>)`,
		ukUA: `\n\nВи можете поділитися цим трекером
🤫 з другом (<i>лише одна людина зможе прийняти</i>)
🌍 або публічно (<i>кожен, хто має посилання, може прийняти, ви можете скасувати це в будь-який час</i>)`,
		ptBR: `\n\nVocê pode compartilhar este rastreador
🤫 com um amigo (<i>apenas uma pessoa pode aceitar</i>)
🌍 ou publicamente (<i>qualquer pessoa com o link pode aceitar, você pode cancelar a qualquer momento</i>)`,
		trTR: `\n\nBu izleyiciyi şu şekilde paylaşabilirsiniz
🤫 bir arkadaşla (<i>yalnızca bir kişi kabul edebilir</i>)
🌍 veya genel olarak (<i>bağlantıya sahip herkes kabul edebilir, istediğiniz zaman iptal edebilirsiniz</i>)`,
		idID: `\n\nAnda dapat membagikan pelacak ini
🤫 kepada seorang teman (<i>hanya satu orang yang dapat menerima</i>)
🌍 atau secara publik (<i>siapa saja yang memiliki tautan dapat menerimanya, Anda dapat membatalkannya kapan saja</i>)`,
	},
	YourSpaces: {
		deDE: "Ihre Bereiche",
		enUK: "Your spaces",
		enUS: "Your spaces",
		esES: "Tus espacios",
		faIR: "فضاهای شما",
		frFR: "Vos espaces",
		idID: "Ruang Anda",
		itIT: "I tuoi spazi",
		jaJP: "あなたのスペース",
		koKR: "귀하의 공간",
		plPL: "Twoje przestrzenie",
		ptBR: "Seus espaços",
		ptPT: "Os teus espaços",
		ruRU: "Ваши пространства",
		trTR: "Alanlarınız",
		ukUA: "Ваші простори",
		uzUZ: "Ваши пространства",
		zhCN: "您的空间",
	},
	CurrentSpace: {
		enUK: "Current space",
		ruRU: "Текущее пространство",
		esES: "Espacio actual",
		faIR: "فضای کنونی",
		itIT: "Spazio corrente",
		deDE: "Aktueller Bereich",
		frFR: "Espace actuel",
		plPL: "Bieżąca przestrzeń",
		ptPT: "Espaço atual",
		koKR: "현재 공간",
		jaJP: "現在のスペース",
		zhCN: "当前空间",
		ukUA: "Поточний простір",
		ptBR: "Espaço atual",
		trTR: "Mevcut alan",
		idID: "Ruang saat ini",
	},
	ClickToSwitchCurrentSpace: {
		enUK: "Click to switch current space",
		ruRU: "Нажмите, чтобы переключить текущее пространство",
		esES: "Haz clic para cambiar el espacio actual",
		faIR: "برای تغییر فضای کنونی کلیک کنید",
		itIT: "Clicca per cambiare spazio corrente",
		deDE: "Klicken Sie hier, um den aktuellen Bereich zu wechseln",
		frFR: "Cliquez pour changer l'espace actuel",
		plPL: "Kliknij, aby zmienić bieżącą przestrzeń",
		ptPT: "Clique para mudar o espaço atual",
		koKR: "현재 공간을 전환하려면 클릭하세요",
		jaJP: "クリックで現在のスペースを切り替えます",
		zhCN: "点击切换当前空间",
		ukUA: "Натисніть, щоб змінити поточний простір",
		ptBR: "Clique para alterar o espaço atual",
		trTR: "Mevcut alanı değiştirmek için tıklayın",
		idID: "Klik untuk mengganti ruang saat ini",
	},
	Family: {
		enUK: "Family",
		ruRU: "Семья",
		esES: "Familia",
		faIR: "خانواده",
		itIT: "Famiglia",
		deDE: "Familie",
		frFR: "Famille",
		plPL: "Rodzina",
		ptPT: "Família",
		koKR: "가족",
		jaJP: "家族",
		zhCN: "家庭",
		ukUA: "Родина",
		ptBR: "Família",
		trTR: "Aile",
		idID: "Keluarga",
	},
	Private: {
		enUK: "Private",
		ruRU: "Личное",
		esES: "Privado",
		faIR: "خصوصی",
		itIT: "Privato",
		deDE: "Privat",
		frFR: "Privé",
		plPL: "Prywatne",
		ptPT: "Privado",
		koKR: "개인",
		jaJP: "プライベート",
		zhCN: "私人",
		ukUA: "Приватне",
		ptBR: "Privado",
		trTR: "Özel",
		idID: "Pribadi",
	},
	ChooseSpace: {
		enUK: "Choose space to make it active for other commands.",
		ruRU: "Выберите пространство, чтобы сделать его активным для других команд.",
		esES: "Elija un espacio para activarlo para otros comandos.",
		faIR: "یک فضا را انتخاب کنید تا برای دستورات دیگر فعال شود.",
		itIT: "Scegli uno spazio per renderlo attivo per altri comandi.",
		deDE: "Wählen Sie einen Bereich, um ihn für andere Befehle zu aktivieren.",
		frFR: "Choisissez un espace pour le rendre actif pour d'autres commandes.",
		plPL: "Wybierz przestrzeń, aby była aktywna dla innych poleceń.",
		ptPT: "Escolha um espaço para torná-lo ativo para outros comandos.",
		koKR: "다른 명령에 활성화하려면 공간을 선택하세요.",
		jaJP: "他のコマンド用にアクティブにするスペースを選択してください。",
		zhCN: "选择一个空间以使其对其他命令有效。",
		ukUA: "Виберіть простір, щоб зробити його активним для інших команд.",
		ptBR: "Escolha um espaço para torná-lo ativo para outros comandos.",
		trTR: "Diğer komutlar için aktif yapmak üzere bir alan seçin.",
		idID: "Pilih ruang untuk menjadikannya aktif untuk perintah lainnya.",
	},
	FamilyContacts: {
		enUK: "Family contacts",
		ruRU: "Семейные контакты",
		esES: "Contactos familiares",
		faIR: "مخاطبین خانواده",
		itIT: "Contatti familiari",
		deDE: "Familienkontakte",
		frFR: "Contacts familiaux",
		plPL: "Kontakty rodzinne",
		ptPT: "Contactos familiares",
		koKR: "가족 연락처",
		jaJP: "家族の連絡先",
		zhCN: "家庭联系人",
		ukUA: "Сімейні контакти",
		ptBR: "Contatos familiares",
		trTR: "Aile kişileriniz",
		idID: "Kontak keluarga",
	},
	YourContacts: {
		enUK: "Your contacts",
		ruRU: "Ваши контакты",
		esES: "Tus contactos",
		faIR: "مخاطبین شما",
		itIT: "I tuoi contatti",
		deDE: "Ihre Kontakte",
		frFR: "Vos contacts",
		plPL: "Twoje kontakty",
		ptPT: "Os seus contactos",
		koKR: "당신의 연락처",
		jaJP: "あなたの連絡先",
		zhCN: "您的联系人",
		ukUA: "Ваші контакти",
		ptBR: "Seus contatos",
		trTR: "Kişileriniz",
		idID: "Kontak Anda",
	},
	ContactsTitle: {
		deDE: "Kontakte",
		enUK: "Contacts",
		enUS: "Contacts",
		esES: "Contactos",
		faIR: "مخاطبین",
		frFR: "Contacts",
		idID: "Kontak",
		itIT: "Contatti",
		jaJP: "連絡先",
		koKR: "연락처",
		plPL: "Kontakty",
		ptBR: "Contatos",
		ptPT: "Contactos",
		ruRU: "Контакты",
		trTR: "Kişiler",
		ukUA: "Контакти",
		uzUZ: "Kontaktlar",
		zhCN: "联系人",
	},
	OpenInApp: {
		enUK: "Open in app",
		ruRU: "Открыть в приложении",
		faIR: "باز کردن در برنامه",
		esES: "Abrir en la aplicación",
		itIT: "Apri nell'app",
		deDE: "In der App öffnen",
		frFR: "Ouvrir dans l'application",
		plPL: "Otwórz w aplikacji",
		ptPT: "Abrir na aplicação",
		koKR: "앱에서 열기",
		jaJP: "アプリで開く",
		zhCN: "在应用中打开",
		ukUA: "Відкрити у додатку",
		ptBR: "Abrir no aplicativo",
		trTR: "Uygulamada aç",
		idID: "Buka di aplikasi",
	},
	ManageInApp: {
		enUK: "Manage in app",
		ruRU: "Управление в приложении",
		esES: "Gestionar en la aplicación",
		faIR: "مدیریت در برنامه",
		itIT: "Gestisci nell'app",
		deDE: "In der App verwalten",
		frFR: "Gérer dans l'application",
		plPL: "Zarządzaj w aplikacji",
		ptPT: "Gerir na aplicação",
		koKR: "앱에서 관리",
		jaJP: "アプリで管理",
		zhCN: "在应用中管理",
		ukUA: "Керуйте в додатку",
		ptBR: "Gerenciar no aplicativo",
		trTR: "Uygulamada yönet",
		idID: "Kelola di aplikasi",
	},
	AddContact: {
		enUK: "Add contact",
		ruRU: "Добавить контакт",
		esES: "Agregar contacto",
		faIR: "اضافه کردن مخاطب",
		itIT: "Aggiungi contatto",
		deDE: "Kontakt hinzufügen",
		frFR: "Ajouter un contact",
		plPL: "Dodaj kontakt",
		ptPT: "Adicionar contacto",
		koKR: "연락처 추가",
		jaJP: "連絡先を追加",
		zhCN: "添加联系人",
		ukUA: "Додати контакт",
		ptBR: "Adicionar contato",
		trTR: "Kişi ekle",
		idID: "Tambah kontak",
	},
	AddMember: {
		deDE: "Mitglied hinzufügen",
		enUK: "Add member",
		enUS: "Add member",
		esES: "Agregar miembro",
		faIR: "اضافه کردن عضو",
		frFR: "Ajouter un membre",
		idID: "Tambah anggota",
		itIT: "Aggiungi membro",
		jaJP: "メンバーを追加",
		koKR: "구성원 추가",
		plPL: "Dodaj członka",
		ptBR: "Adicionar membro",
		ptPT: "Adicionar membro",
		ruRU: "Добавить участника",
		trTR: "Üye ekle",
		ukUA: "Додати учасника",
		uzUZ: "A'zo qo'shish",
		zhCN: "添加成员",
	},
	MyContact: {
		enUK: "My contact",
		ruRU: "Мой контакт",
		esES: "Mi contacto",
		faIR: "مخاطب من",
		itIT: "Il mio contatto",
		deDE: "Mein Kontakt",
		frFR: "Mon contact",
		plPL: "Mój kontakt",
		ptPT: "O meu contacto",
		koKR: "내 연락처",
		jaJP: "私の連絡先",
		zhCN: "我的联系人",
		ukUA: "Мій контакт",
		ptBR: "Meu contato",
		trTR: "Benim kişim",
		idID: "Kontak saya",
	},
	GenderTitled: {
		enUK: "Gender",
		ruRU: "Пол",
		esES: "Género",
		faIR: "جنسیت",
		itIT: "Genere",
		deDE: "Geschlecht",
		frFR: "Genre",
		plPL: "Płeć",
		ptPT: "Género",
		koKR: "성별",
		jaJP: "性別",
		zhCN: "性别",
		ukUA: "Стать",
		ptBR: "Gênero",
		trTR: "Cinsiyet",
		idID: "Jenis Kelamin",
	},
	MaleTitled: {
		enUK: "Male",
		ruRU: "Мужской",
		esES: "Masculino",
		faIR: "مرد",
		itIT: "Maschio",
		deDE: "Männlich",
		frFR: "Homme",
		plPL: "Mężczyzna",
		ptPT: "Masculino",
		koKR: "남성",
		jaJP: "男性",
		zhCN: "男性",
		ukUA: "Чоловік",
		ptBR: "Masculino",
		trTR: "Erkek",
		idID: "Pria",
	},
	FemaleTitled: {
		enUK: "Female",
		ruRU: "Женский",
		esES: "Femenino",
		faIR: "زن",
		itIT: "Femmina",
		deDE: "Weiblich",
		frFR: "Femme",
		plPL: "Kobieta",
		ptPT: "Feminino",
		koKR: "여성",
		jaJP: "女性",
		zhCN: "女性",
		ukUA: "Жінка",
		ptBR: "Feminino",
		trTR: "Kadın",
		idID: "Wanita",
	},
	UnknownTitled: {
		enUK: "Unknown",
		ruRU: "Неизвестно",
		esES: "Desconocido",
		faIR: "ناشناس",
		itIT: "Sconosciuto",
		deDE: "Unbekannt",
		frFR: "Inconnu",
		plPL: "Nieznany",
		ptPT: "Desconhecido",
		koKR: "알 수 없음",
		jaJP: "不明",
		zhCN: "未知",
		ukUA: "Невідомо",
		ptBR: "Desconhecido",
		trTR: "Bilinmiyor",
		idID: "Tidak Diketahui",
	},
	UndisclosedTitled: {
		enUK: "Undisclosed",
		ruRU: "Не разглашено",
		esES: "No revelado",
		faIR: "فاش نشده",
		itIT: "Non divulgato",
		deDE: "Nicht offengelegt",
		frFR: "Non divulgué",
		plPL: "Nieujawnione",
		ptPT: "Não divulgado",
		koKR: "공개되지 않음",
		jaJP: "未公開",
		zhCN: "未披露",
		ukUA: "Не розголошено",
		ptBR: "Não divulgado",
		trTR: "Açıklanmamış",
		idID: "Tidak Diungkapkan",
	},
	InviteToJoinSpace: {
		enUK: "Invite to join",
		ruRU: "Пригласить присоедениться",
		esES: "Invitar a unirse",
		faIR: "دعوت به پیوستن",
		itIT: "Invita a unirti",
		deDE: "Einladen beizutreten",
		frFR: "Inviter à rejoindre",
		plPL: "Zaproś do dołączenia",
		ptPT: "Convidar para juntar-se",
		koKR: "가입 초대",
		jaJP: "参加に招待",
		zhCN: "邀请加入",
		ukUA: "Запросити приєднатися",
		ptBR: "Convidar para participar",
		trTR: "Katılmaya davet et",
		idID: "Undang untuk bergabung",
	},
	BackToContacts: {
		enUK: "Back to contacts",
		ruRU: "Вернуться к контактам",
		esES: "Volver a los contactos",
		faIR: "بازگشت به مخاطبین",
		itIT: "Torna ai contatti",
		deDE: "Zurück zu Kontakten",
		frFR: "Retour aux contacts",
		plPL: "Powrót do kontaktów",
		ptPT: "Voltar aos contatos",
		koKR: "연락처로 돌아가기",
		jaJP: "連絡先に戻る",
		zhCN: "返回联系人",
		ukUA: "Повернутися до контактів",
		ptBR: "Voltar para os contatos",
		trTR: "Kişilere geri dön",
		idID: "Kembali ke kontak",
	},
	BackToMembers: {
		enUK: "Back to members",
		ruRU: "Вернуться к участникам",
		esES: "Volver a los miembros",
		faIR: "بازگشت به اعضا",
		itIT: "Torna ai membri",
		deDE: "Zurück zu Mitgliedern",
		frFR: "Retour aux membres",
		plPL: "Powrót do członków",
		ptPT: "Voltar aos membros",
		koKR: "멤버로 돌아가기",
		jaJP: "メンバーに戻る",
		zhCN: "返回成员",
		ukUA: "Повернутися до учасників",
		ptBR: "Voltar para os membros",
		trTR: "Üyelere geri dön",
		idID: "Kembali ke anggota",
	},
	BackBtnTitle: {
		enUK: "Back",
		ruRU: "Назад",
		esES: "Atrás",
		faIR: "بازگشت",
		itIT: "Indietro",
		deDE: "Zurück",
		frFR: "Retour",
		plPL: "Wstecz",
		ptPT: "Voltar",
		koKR: "뒤로",
		jaJP: "戻る",
		zhCN: "返回",
		ukUA: "Назад",
		ptBR: "Voltar",
		trTR: "Geri",
		idID: "Kembali",
	},
	SelectTrackerToAdd: {
		enUK: "Select a tracker to add",
		ruRU: "Выберите трекер для добавления",
		esES: "Seleccionar un rastreador para agregar",
		faIR: "یک ردیاب برای افزودن انتخاب کنید",
		itIT: "Seleziona un tracker da aggiungere",
		deDE: "Wählen Sie einen Tracker zum Hinzufügen aus",
		frFR: "Sélectionner un tracker à ajouter",
		plPL: "Wybierz tracker do dodania",
		ptPT: "Selecionar um rastreador para adicionar",
		koKR: "추가할 트래커를 선택하세요",
		jaJP: "追加するトラッカーを選択",
		zhCN: "选择要添加的跟踪器",
		ukUA: "Виберіть трекер для додавання",
		ptBR: "Selecionar um rastreador para adicionar",
		trTR: "Eklemek için bir izleyici seçin",
		idID: "Pilih pelacak untuk ditambahkan",
	},
	SelectCategoryForNewTracker: {
		enUK: "Select a category for the new tracker",
		ruRU: "Выберите категорию для нового трекера",
		esES: "Seleccionar una categoría para el nuevo rastreador",
		faIR: "یک دسته برای ردیاب جدید انتخاب کنید",
		itIT: "Seleziona una categoria per il nuovo tracker",
		deDE: "Wählen Sie eine Kategorie für den neuen Tracker aus",
		frFR: "Sélectionnez une catégorie pour le nouveau traqueur",
		plPL: "Wybierz kategorię dla nowego trackera",
		ptPT: "Selecionar uma categoria para o novo rastreador",
		koKR: "새 트래커의 카테고리를 선택하세요",
		jaJP: "新しいトラッカーのカテゴリーを選択",
		zhCN: "为新跟踪器选择一个类别",
		ukUA: "Виберіть категорію для нового трекера",
		ptBR: "Selecionar uma categoria para o novo rastreador",
		trTR: "Yeni izleyici için bir kategori seçin",
		idID: "Pilih kategori untuk pelacak baru",
	},
	BackToLists: {
		enUK: "Back to lists",
		ruRU: "Вернуться к спискам",
		esES: "Volver a las listas",
		faIR: "بازگشت به لیست\u200cها",
		itIT: "Torna alle liste",
		deDE: "Zurück zu den Listen",
		frFR: "Retour aux listes",
		plPL: "Powrót do list",
		ptPT: "Voltar às listas",
		koKR: "목록으로 돌아가기",
		jaJP: "リストに戻る",
		zhCN: "返回列表",
		ukUA: "Повернутися до списків",
		ptBR: "Voltar para as listas",
		trTR: "Listelere geri dön",
		idID: "Kembali ke daftar-daftar",
	},
	BackToList: {
		enUK: "Back to list",
		ruRU: "Вернуться к списку",
		esES: "Volver a la lista",
		faIR: "بازگشت به لیست",
		itIT: "Torna alla lista",
		deDE: "Zurück zur Liste",
		frFR: "Retour à la liste",
		plPL: "Powrót do listy",
		ptPT: "Voltar à lista",
		koKR: "목록으로 돌아가기",
		jaJP: "リストに戻る",
		zhCN: "返回列表",
		ukUA: "Повернутися до списку",
		ptBR: "Voltar para a lista",
		trTR: "Listeye geri dön",
		idID: "Kembali ke daftar",
	},
	ClearList: {
		enUK: "Clear list",
		ruRU: "Очистить список",
		esES: "Limpiar lista",
		faIR: "پاک کردن لیست",
		itIT: "Svuota lista",
		deDE: "Liste löschen",
		frFR: "Effacer la liste",
		plPL: "Wyczyść listę",
		ptPT: "Limpar lista",
		koKR: "목록 지우기",
		jaJP: "リストをクリア",
		zhCN: "清空列表",
		ukUA: "Очистити список",
		ptBR: "Limpar lista",
		trTR: "Listeyi temizle",
		idID: "Bersihkan daftar",
	},
	MarkAsDone: {
		enUK: "Mark as done",
		ruRU: "Отметить как выполненное",
		esES: "Marcar como hecho",
		faIR: "علامت\u200cگذاری به عنوان انجام\u200cشده",
		itIT: "Segna come completato",
		deDE: "Als erledigt markieren",
		frFR: "Marquer comme fait",
		plPL: "Oznacz jako zrobione",
		ptPT: "Marcar como feito",
		koKR: "완료로 표시",
		jaJP: "完了としてマークする",
		zhCN: "标记为完成",
		ukUA: "Позначити як виконане",
		ptBR: "Marcar como concluído",
		trTR: "Tamamlandı olarak işaretle",
		idID: "Tandai sebagai selesai",
	},
	AddStandard: {
		enUK: "Add standard",
		ruRU: "Добавить стандартный",
		esES: "Agregar estándar",
		faIR: "افزودن استاندارد",
		itIT: "Aggiungere standard",
		deDE: "Standard hinzufügen",
		frFR: "Ajouter une norme",
		plPL: "Dodaj standard",
		ptPT: "Adicionar padrão",
		koKR: "표준 추가",
		jaJP: "標準を追加",
		zhCN: "添加标准",
		ukUA: "Додати стандарт",
		ptBR: "Adicionar padrão",
		trTR: "Standart ekle",
		idID: "Tambahkan standar",
	},
	DeleteItems: {
		enUK: "Delete items",
		ruRU: "Удалить элементы",
		faIR: "حذف موارد",
		itIT: "Elimina elementi",
		deDE: "Elemente löschen",
		frFR: "Supprimer les éléments",
		plPL: "Usuń elementy",
		ptPT: "Eliminar itens",
		koKR: "항목 삭제",
		jaJP: "項目を削除",
		zhCN: "删除项目",
		ukUA: "Видалити елементи",
		ptBR: "Excluir itens",
		trTR: "Öğeleri sil",
		idID: "Hapus item",
	},
	YouCanAddItemBySendingMessage: {
		enUK: "You can add items to this list by sending a message to me.",
		ruRU: "Вы можете добавлять элементы в этот список, отправив мне сообщение.",
		esES: "Puedes agregar elementos a esta lista enviándome un mensaje.",
		faIR: "می\u200cتوانید با ارسال یک پیام به من، موارد را به این لیست اضافه کنید.",
		itIT: "Puoi aggiungere elementi a questa lista inviandomi un messaggio.",
		deDE: "Sie können Elemente zu dieser Liste hinzufügen, indem Sie mir eine Nachricht senden.",
		frFR: "Vous pouvez ajouter des éléments à cette liste en m'envoyant un message.",
		plPL: "Możesz dodać elementy do tej listy, wysyłając mi wiadomość.",
		ptPT: "Pode adicionar items a esta lista enviando-me uma mensagem.",
		koKR: "저에게 메시지를 보내서 이 목록에 항목들을 추가할 수 있습니다.",
		jaJP: "私にメッセージを送ることで、このリストに項目を追加できます。",
		zhCN: "您可以通过向我发送消息将项目添加到此列表。",
		ukUA: "Ви можете додати елементи до цього списку, надіславши мені повідомлення.",
		ptBR: "Você pode adicionar itens a esta lista enviando uma mensagem para mim.",
		trTR: "Bu listeye bir mesaj göndererek öğeler ekleyebilirsiniz.",
		idID: "Anda dapat menambahkan item ke daftar ini dengan mengirimkan pesan kepada saya.",
	},
	NoItemsInTheListYet: {
		enUK: "No items in the list yet.",
		ruRU: "Пока нет элементов в списке.",
		esES: "Todavía no hay elementos en la lista.",
		faIR: "هنوز هیچ آیتمی در لیست وجود ندارد.",
		itIT: "Non ci sono ancora elementi nella lista.",
		deDE: "Noch keine Elemente in der Liste.",
		frFR: "Pas encore d'éléments dans la liste.",
		plPL: "Jeszcze brak elementów na liście.",
		ptPT: "Ainda não existem itens na lista.",
		koKR: "목록에 아직 항목이 없습니다.",
		jaJP: "リストに項目がまだありません。",
		zhCN: "列表中尚无项目。",
		ukUA: "Ще немає елементів у списку.",
		ptBR: "Ainda não há itens na lista.",
		trTR: "Henüz listede öğe yok.",
		idID: "Belum ada item dalam daftar.",
	},
	FamilyList: {
		enUK: "Family list",
		ruRU: "Семейный список",
		esES: "Lista familiar",
		faIR: "لیست خانواده",
		itIT: "Elenco familiare",
		deDE: "Familienliste",
		frFR: "Liste familiale",
		plPL: "Lista rodzinna",
		ptPT: "Lista familiar",
		koKR: "가족 목록",
		jaJP: "家族リスト",
		zhCN: "家庭清单",
		ukUA: "Сімейний список",
		ptBR: "Lista da família",
		trTR: "Aile listesi",
		idID: "Daftar keluarga",
	},
	PrivateList: {
		enUK: "Private list",
		ruRU: "Личный список",
		esES: "Lista privada",
		faIR: "لیست خصوصی",
		itIT: "Lista personale",
		deDE: "Private Liste",
		frFR: "Liste privée",
		plPL: "Lista prywatna",
		ptPT: "Lista privada",
		koKR: "개인 목록",
		jaJP: "プライベートリスト",
		zhCN: "私人清单",
		ukUA: "Особистий список",
		ptBR: "Lista privada",
		trTR: "Özel liste",
		idID: "Daftar pribadi",
	},
	Refresh: {
		enUK: "Refresh",
		ruRU: "Обновить",
		esES: "Refrescar",
		faIR: "بروزرسانی",
		itIT: "Aggiorna",
		deDE: "Aktualisieren",
		frFR: "Rafraîchir",
		plPL: "Odśwież",
		ptPT: "Atualizar",
		koKR: "새로고침",
		jaJP: "リフレッシュ",
		zhCN: "刷新",
		ukUA: "Оновити",
		ptBR: "Atualizar",
		trTR: "Yenile",
		idID: "Segarkan",
	},
	AdviseToUseTelegramForTgUsers: {
		enUK: `If the person you want to add uses telegram we advise to select "Choose Telegram User".
Otherwise you can add them manually in Sneat.app.`,
		ruRU: `Если человек, которого вы хотите добавить, использует Telegram, 
		мы рекомендуем выбрать "Выбрать пользователя Telegram". 
		В противном случае вы можете добавить их вручную в Sneat.app.`,
		esES: `Si la persona que deseas agregar usa Telegram, te recomendamos seleccionar "Elegir usuario de Telegram".
		De lo contrario, puedes agregarlos manualmente en Sneat.app.`,
		faIR: `اگر شخصی که می\u200cخواهید اضافه کنید از تلگرام استفاده می\u200cکند، 
ما توصیه می\u200cکنیم گزینه "انتخاب کاربر تلگرام" را انتخاب کنید. 
در غیر این صورت می\u200cتوانید آنها را به صورت دستی در Sneat.app اضافه کنید.`,
		itIT: `Se la persona che vuoi aggiungere utilizza telegram, ti consigliamo di selezionare "Scegli utente Telegram".
		In caso contrario, puoi aggiungerli manualmente in Sneat.app.`,
		deDE: `Wenn die Person, die Sie hinzufügen möchten, Telegram verwendet, empfehlen wir, "Telegram-Benutzer auswählen" auszuwählen.
		Andernfalls können Sie sie manuell in Sneat.app hinzufügen.`,
		frFR: `Si la personne que vous souhaitez ajouter utilise Telegram, nous vous conseillons de sélectionner "Choisir un utilisateur Telegram".
		Sinon, vous pouvez les ajouter manuellement sur Sneat.app.`,
		plPL: `Jeśli osoba, którą chcesz dodać, korzysta z Telegrama, radzimy wybrać "Wybierz użytkownika Telegrama".
		W przeciwnym razie możesz dodać ich ręcznie w aplikacji Sneat.app.`,
		ptPT: `Se a pessoa que deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente na Sneat.app.`,
		koKR: `추가하려는 사람이 Telegram을 사용한다면 "Telegram 사용자 선택"을 선택하시길 권장드립니다.
		그렇지 않을 경우 Sneat.app에서 수동으로 추가할 수 있습니다.`,
		jaJP: `追加したい相手がTelegramを使用している場合、「Telegram ユーザーを選択」を選択することをお勧めします。
		それ以外の場合は、Sneat.app で手動で追加することができます。`,
		zhCN: `如果您想添加的人使用 Telegram，我们建议选择“选择 Telegram 用户”。
		否则，您可以在 Sneat.app 中手动添加他们。`,
		ukUA: `Якщо людина, яку ви хочете додати, використовує Telegram,
		ми радимо вибрати "Вибрати користувача Telegram".
		Інакше, ви можете додати їх вручну на Sneat.app.`,
		ptBR: `Se a pessoa que você deseja adicionar usar o Telegram, recomendamos selecionar "Escolher usuário do Telegram".
		Caso contrário, você pode adicioná-los manualmente no Sneat.app.`,
		trTR: `Eklemek istediğiniz kişi Telegram kullanıyorsa, "Telegram Kullanıcısını Seç" seçeneğini seçmenizi öneririz.
		Aksi takdirde Sneat.app'te manuel olarak ekleyebilirsiniz.`,
		idID: `Jika orang yang ingin Anda tambahkan menggunakan Telegram, kami sarankan untuk memilih "Pilih Pengguna Telegram".
		Jika tidak, Anda dapat menambahkannya secara manual di Sneat.app.`,
	},
	ChooseTelegramUser: {
		enUK: "Choose Telegram user",
		ruRU: "Выбрать пользователя Telegram",
		esES: "Elegir usuario de Telegram",
		faIR: "انتخاب کاربر تلگرام",
		itIT: "Scegli utente Telegram",
		deDE: "Telegram-Benutzer auswählen",
		frFR: "Choisir un utilisateur Telegram",
		plPL: "Wybierz użytkownika Telegrama",
		ptPT: "Escolher usuário do Telegram",
		koKR: "Telegram 사용자 선택",
		jaJP: "Telegram ユーザーを選択",
		zhCN: "选择 Telegram 用户",
		ukUA: "Вибрати користувача Telegram",
		ptBR: "Escolher usuário do Telegram",
		trTR: "Telegram Kullanıcısını Seç",
		idID: "Pilih Pengguna Telegram",
	},
	AddManuallyInSneatApp: {
		enUK: "Add manually in Sneat.app",
		ruRU: "Добавить вручную в Sneat.app",
		esES: "Agregar manualmente en Sneat.app",
		faIR: "به صورت دستی در Sneat.app اضافه کنید",
		itIT: "Aggiungi manualmente in Sneat.app",
		deDE: "Manuell in Sneat.app hinzufügen",
		frFR: "Ajouter manuellement dans Sneat.app",
		plPL: "Dodaj ręcznie w Sneat.app",
		ptPT: "Adicionar manualmente no Sneat.app",
		koKR: "Sneat.app에서 수동으로 추가",
		jaJP: "Sneat.app で手動で追加",
		zhCN: "在 Sneat.app 中手动添加",
		ukUA: "Додати вручну на Sneat.app",
		ptBR: "Adicionar manualmente no Sneat.app",
		trTR: "Sneat.app'te manuel olarak ekle",
		idID: "Tambahkan secara manual di Sneat.app",
	},
	CancelAddingMember: {
		enUK: "Cancel adding member",
		ruRU: "Отменить добавление участника",
		esES: "Cancelar la adición de miembro",
		faIR: "لغو افزودن عضو",
		itIT: "Annulla l'aggiunta di un membro",
		deDE: "Hinzufügen eines Mitglieds abbrechen",
		frFR: "Annuler l'ajout d'un membre",
		plPL: "Anuluj dodawanie członka",
		ptPT: "Cancelar a adição de membro",
		koKR: "멤버 추가 취소",
		jaJP: "メンバー追加をキャンセル",
		zhCN: "取消添加成员",
		ukUA: "Скасувати додавання учасника",
		ptBR: "Cancelar a adição de membro",
		trTR: "Üye eklemeyi iptal et",
		idID: "Batalkan menambahkan anggota",
	},
	CancelAddingContact: {
		enUK: "Cancel adding contact",
		ruRU: "Отменить добавление контакта",
		esES: "Cancelar la adición de contacto",
		faIR: "لغو افزودن مخاطب",
		itIT: "Annulla l'aggiunta di un contatto",
		deDE: "Hinzufügen eines Kontakts abbrechen",
		frFR: "Annuler l'ajout d'un contact",
		plPL: "Anuluj dodawanie kontaktu",
		ptPT: "Cancelar a adição de contato",
		koKR: "연락처 추가 취소",
		jaJP: "連絡先追加をキャンセル",
		zhCN: "取消添加联系人",
		ukUA: "Скасувати додавання контакту",
		ptBR: "Cancelar a adição de contato",
		trTR: "Kişi eklemeyi iptal et",
		idID: "Batalkan menambahkan kontak",
	},
	FamilyDebts: {
		enUK: "Family debts",
		ruRU: "Семейные долги",
		esES: "Deudas familiares",
		faIR: "بدهی های خانواده",
		itIT: "Debiti familiari",
		deDE: "Familienschulden",
		frFR: "Dettes familiales",
		plPL: "Długi rodzinne",
		ptPT: "Dívidas da família",
		koKR: "가족 부채",
		jaJP: "家族の借金",
		zhCN: "家庭债务",
		ukUA: "Сімейні борги",
		ptBR: "Dívidas familiares",
		trTR: "Aile borçları",
		idID: "Hutang keluarga",
	},
	DebtsRelatedContacts: {
		enUK: "Debts related contacts",
		ruRU: "Контакты, связанные с долгами",
		esES: "Contactos relacionados con deudas",
		faIR: "ارتباطات مربوط به بدهی\u200cها",
		itIT: "Contatti legati ai debiti",
		deDE: "Schuldenbezogene Kontakte",
		frFR: "Contacts liés aux dettes",
		plPL: "Kontakty związane z długami",
		ptPT: "Contactos relacionados com dívidas",
		koKR: "부채 관련 연락처",
		jaJP: "借金関連の連絡先",
		zhCN: "债务相关联系人",
		ukUA: "Контакти, пов’язані з боргами",
		ptBR: "Contatos relacionados a dívidas",
		trTR: "Borçlarla ilgili kişiler",
		idID: "Kontak yang terkait dengan hutang",
	},
	BackToDebtsMenu: {
		enUK: "Back to debts menu",
		ruRU: "Вернуться в меню долгов",
		esES: "Volver al menú de deudas",
		faIR: "بازگشت به منوی بدهی\u200cها",
		itIT: "Torna al menu dei debiti",
		deDE: "Zurück zum Schuldenmenü",
		frFR: "Retour au menu des dettes",
		plPL: "Powrót do menu długów",
		ptPT: "Voltar ao menu das dívidas",
		koKR: "부채 메뉴로 돌아가기",
		jaJP: "借金メニューに戻る",
		zhCN: "返回债务菜单",
		ukUA: "Повернутися до меню боргів",
		ptBR: "Voltar ao menu de dívidas",
		trTR: "Borçlar menüsüne geri dön",
		idID: "Kembali ke menu hutang",
	},
}
