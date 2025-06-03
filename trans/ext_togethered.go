package trans

const (
	TogdBotWelcome = "TogdBotWelcome"

	TogdMyProfile    = "TogdMyProfile"
	TogdMyActivities = "TogdMyActivities"
	TogdMyEvents     = "TogdMyEvents"
	TogdMyPlans      = "TogdMyPlans"
	TogdMySpots      = "TogdMySpots"

	TogdPlansButtonText = "TogdPlansButtonText"
	TogdSpotsButtonText = "TogdSpotsButtonText"

	TogdUserProfile    = "TogdUserProfile"
	TogdUserActivities = "TogdUserActivities"
	TogdUserEvents     = "TogdUserEvents"
	TogdUserPlans      = "TogdUserPlans"
	TogdUserSpots      = "TogdUserSpots"

	TogdActivitiesOfUser             = "TogdActivitiesOfUser"
	YouHaveNoFavoriteActivities      = "YouHaveNoFavoriteActivities"
	InstructionHowToAddActivityInBot = "InstructionHowToAddActivityInBot"

	TogdMainMenuText = "TogdMainMenuText"

	RsvpButtonText     = "RsvpButtonText"
	NewEventButtonText = "NewEventButtonText"

	NewEventTitle = "NewEventTitle"
	NewEventText  = "NewEventText"
	NewEventHint  = "NewEventHint"

	EventOptionsButton   = "EventOptionsButton"
	NewEventOptionButton = "NewEventOptionButton"
	EventVisibility      = "EventVisibility"
	EventStatus          = "EventStatus"
	CancelEvent          = "CancelEvent"
	BackToEvents         = "BackToEvents"
	EventCreated         = "EventCreated"
	EventTitle           = "EventTitle"
	PlanEventButtonText  = "PlanEventButtonText"

	RemoveFromSpots = "RemoveFromSpots"
	AddToMySpots    = "AddToMySpots"

	FollowButtonText   = "FollowButtonText"
	UnfollowButtonText = "UnfollowButtonText"

	ShareSpotButtonText = "ShareSpotButtonText"

	SpotButtonText          = "SpotButtonText"
	SpotTitle               = "SpotTitle"
	SpotActivities          = "SpotActivities"
	SpotGoingToDoActivities = "SpotGoingToDoActivities"
	TogdBackToActivities    = "TogdBackToActivities"

	RsvpHowLikelyQuestion    = "RsvpHowLikelyQuestion"
	RsvpResponse100Percent   = "RsvpResponse100Percent"
	RsvpResponseNotAttending = "RsvpResponseNotAttending"
	RsvpResponseMostLikely   = "RsvpResponseMostLikely"
	RsvpResponseMaybe        = "RsvpResponseMaybe"
	RsvpResponseUnlikely     = "RsvpResponseUnlikely"

	RsvpQuestionOnWhatDate = "RsvpQuestionOnWhatDate"
	RsvpQuestionAtWhatTime = "RsvpQuestionAtWhatTime"
	RsvpTimeIsChangeable   = "RsvpTimeIsChangeable"

	ChooseSpotToRSVP = "ChooseSpotToRSVP"

	TogdIntentPublished = "TogdIntentPublished"

	TodayButtonText    = "TodayButtonText"
	TomorrowButtonText = "TomorrowButtonText"
)

const CalendarButtonText = "CalendarButtonText"

func calendarButtonText() map[string]string {
	return map[string]string{
		deDE: "🗓️️ Kalender",
		enUK: "🗓️️ Calendar",
		esES: "🗓️️ Calendario",
		faIR: "🗓️️ تقویم",
		frFR: "🗓️️ Calendrier",
		idID: "🗓️️ Kalender",
		itIT: "🗓️️ Calendario",
		jaJP: "🗓️️ カレンダー",
		koKR: "🗓️️ 달력",
		plPL: "🗓️️ Kalendarz",
		ptBR: "🗓️️ Calendário",
		ptPT: "🗓️️ Calendário",
		ruRU: "🗓️️ Календарь",
		trTR: "🗓️️ Takvim",
		ukUA: "🗓️️ Календар",
		uzUZ: "🗓️️ Taqvim",
		zhCN: "🗓️️ 日历",
	}
}

func init() {
	/*
		locale items should be ordered by key in ascending order,
		use it as a reference for all values:

			- deDE
			- enUK:
			- enUS:
			- esES:
			- faIR:
			- "fr-FR":
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
	var trans = map[string]map[string]string{
		CalendarButtonText: calendarButtonText(),
		NewEventTitle: {
			deDE: "Neue Veranstaltung",
			enUK: "New Event",
			esES: "Nuevo Evento",
			faIR: "رویداد جدید",
			frFR: "Nouvel Événement",
			idID: "Acara Baru",
			itIT: "Nuovo Evento",
			jaJP: "新しいイベント",
			koKR: "새 이벤트",
			plPL: "Nowe Wydarzenie",
			ptBR: "Novo Evento",
			ruRU: "Новое Событие",
			trTR: "Yeni Etkinlik",
			ukUA: "Нова Подія",
			uzUZ: "Yangi Tadbir",
			zhCN: "新事件",
		},
		NewEventText: {
			deDE: "@{BOT_CODE} hilft dabei, Datum, Uhrzeit und Ort zu wählen, die für alle Teilnehmer am besten funktionieren.",
			enUK: "@{BOT_CODE} helps to choose date, time & place that works best for all participants.",
			esES: "@{BOT_CODE} ayuda a elegir fecha, hora y lugar que funcione mejor para todos los participantes.",
			faIR: "@{BOT_CODE} به انتخاب تاریخ، زمان و مکانی که برای همه شرکت\u200Cکنندگان مناسب است کمک می\u200Cکند.",
			frFR: "@{BOT_CODE} aide à choisir la date, l'heure et le lieu qui conviennent le mieux à tous les participants.",
			idID: "@{BOT_CODE} membantu memilih tanggal, waktu & tempat yang paling cocok untuk semua peserta.",
			itIT: "@{BOT_CODE} aiuta a scegliere data, ora e luogo che funzionano meglio per tutti i partecipanti.",
			jaJP: "@{BOT_CODE} は、すべての参加者に最適な日付、時間、場所を選ぶのに役立ちます。",
			koKR: "@{BOT_CODE}는 모든 참가자에게 가장 적합한 날짜, 시간 및 장소를 선택하는 데 도움을 줍니다.",
			plPL: "@{BOT_CODE} pomaga wybrać datę, godzinę i miejsce, które najlepiej pasują wszystkim uczestników.",
			ptBR: "@{BOT_CODE} ajuda a escolher data, hora e local que funcionam melhor para todos os participantes.",
			ruRU: "@{BOT_CODE} помогает выбрать дату, время и место, которые лучше всего подходят для всех участников.",
			trTR: "@{BOT_CODE} tüm katılımcılar için en uygun tarih, saat ve yeri seçmeye yardımcı olur.",
			ukUA: "@{BOT_CODE} допомагає обрати дату, час і місце, які найкраще підходять для всіх учасників.",
			uzUZ: "@{BOT_CODE} barcha ishtirokchilar uchun eng mos sana, vaqt va joyni tanlashda yordam beradi.",
			zhCN: "@{BOT_CODE} 帮助选择最适合所有参与者的日期、时间和地点。",
		},
		NewEventHint: {
			deDE: "Geben Sie den Titel Ihrer neuen Veranstaltung ein:",
			enUK: "Enter title of your new event:",
			esES: "Ingrese el título de su nuevo evento:",
			faIR: "عنوان رویداد جدید خود را وارد کنید:",
			frFR: "Entrez le titre de votre nouvel événement:",
			idID: "Masukkan judul acara baru Anda:",
			itIT: "Inserisci il titolo del tuo nuovo evento:",
			jaJP: "新しいイベントのタイトルを入力してください:",
			koKR: "새 이벤트의 제목을 입력하세요:",
			plPL: "Wprowadź tytuł swojego nowego wydarzenia:",
			ptBR: "Digite o título do seu novo evento:",
			ruRU: "Введите название вашего нового события:",
			trTR: "Yeni etkinliğinizin başlığını girin:",
			ukUA: "Введіть назву вашої нової події:",
			uzUZ: "Yangi tadbiringizning sarlavhasini kiriting:",
			zhCN: "输入您的新事件标题:",
		},
		TodayButtonText: {
			deDE: "🕒 Heute — {DATE}",
			enUK: "🕒 Today — {DATE}",
			esES: "🕒 Hoy — {DATE}",
			faIR: "🕒 امروز — {DATE}",
			frFR: "🕒 Aujourd'hui — {DATE}",
			idID: "🕒 Hari ini — {DATE}",
			itIT: "🕒 Oggi — {DATE}",
			jaJP: "🕒 今日 — {DATE}",
			koKR: "🕒 오늘 — {DATE}",
			plPL: "🕒 Dzisiaj — {DATE}",
			ptBR: "🕒 Hoje — {DATE}",
			ptPT: "🕒 Hoje — {DATE}",
			ruRU: "🕒 Сегодня — {DATE}",
			trTR: "🕒 Bugün — {DATE}",
			ukUA: "🕒 Сьогодні — {DATE}",
			uzUZ: "🕒 Bugun — {DATE}",
			zhCN: "🕒 今天 — {DATE}",
		},
		TomorrowButtonText: {
			deDE: "🌅 Morgen — {DATE}",
			enUK: "🌅 Tomorrow —  {DATE}",
			esES: "🌅 Mañana — {DATE}",
			faIR: "🌅 فردا — {DATE}",
			frFR: "🌅 Demain — {DATE}",
			idID: "🌅 Besok — {DATE}",
			itIT: "🌅 Domani — {DATE}",
			jaJP: "🌅 明日 — {DATE}",
			koKR: "🌅 내일 — {DATE}",
			plPL: "🌅 Jutro — {DATE}",
			ptBR: "🌅 Amanhã — {DATE}",
			ruRU: "🌅 Завтра — {DATE}",
			trTR: "🌅 Yarın — {DATE}",
			ukUA: "🌅 Завтра — {DATE}",
			uzUZ: "🌅 Ertaga — {DATE}",
			zhCN: "🌅 明天 — {DATE}",
		},
		SpotGoingToDoActivities: {
			deDE: "Vorhaben: {ACTIVITIES}",
			enUK: "Going to do: {ACTIVITIES}",
			esES: "Voy a hacer: {ACTIVITIES}",
			faIR: "قصد انجام: {ACTIVITIES}",
			frFR: "Va faire: {ACTIVITIES}",
			idID: "Akan melakukan: {ACTIVITIES}",
			itIT: "Intenzione di fare: {ACTIVITIES}",
			jaJP: "する予定: {ACTIVITIES}",
			koKR: "할 일: {ACTIVITIES}",
			plPL: "Zamierzam zrobić: {ACTIVITIES}",
			ptBR: "Vai fazer: {ACTIVITIES}",
			ptPT: "Vai fazer: {ACTIVITIES}",
			ruRU: "Собираюсь делать: {ACTIVITIES}",
			trTR: "Yapacağım: {ACTIVITIES}",
			ukUA: "Збираюся робити: {ACTIVITIES}",
			uzUZ: "Qilmoqchi: {ACTIVITIES}",
			zhCN: "打算做: {ACTIVITIES}",
		},
		ChooseSpotToRSVP: {
			deDE: "Wählen Sie einen Platz zum Zusagen",
			enUK: "Choose a spot to RSVP",
			esES: "Elige un lugar para confirmar asistencia",
			faIR: "مکانی برای تایید حضور انتخاب کنید",
			frFR: "Choisissez un lieu pour confirmer votre présence",
			idID: "Pilih tempat untuk konfirmasi kehadiran",
			itIT: "Scegli un posto per confermare la presenza",
			jaJP: "出席返事をする場所を選択してください",
			koKR: "참석 응답할 장소를 선택하세요",
			plPL: "Wybierz miejsce, aby potwierdzić obecność",
			ptBR: "Escolha um local para confirmar presença",
			ruRU: "Выберите место для подтверждения участия",
			trTR: "Katılımı onaylamak için bir yer seçin",
			ukUA: "Оберіть місце для підтвердження участі",
			uzUZ: "Ishtirok etishni tasdiqlash uchun joy tanlang",
			zhCN: "选择一个地点来回复邀请",
		},
		TogdIntentPublished: {
			deDE: "Sie haben Ihre Absicht erfolgreich veröffentlicht.",
			enUK: "You've successfully published your intention.",
			esES: "Has publicado tu intención exitosamente.",
			faIR: "شما با موفقیت قصد خود را منتشر کردید.",
			frFR: "Vous avez publié votre intention avec succès.",
			idID: "Anda telah berhasil menerbitkan niat Anda.",
			itIT: "Hai pubblicato con successo la tua intenzione.",
			jaJP: "意図を正常に公開しました。",
			koKR: "의도를 성공적으로 게시했습니다.",
			plPL: "Pomyślnie opublikowałeś swoją intencję.",
			ptBR: "Você publicou sua intenção com sucesso.",
			ruRU: "Вы успешно опубликовали свое намерение.",
			trTR: "Niyetinizi başarıyla yayınladınız.",
			ukUA: "Ви успішно опублікували свій намір.",
			uzUZ: "Siz o'z niyatingizni muvaffaqiyatli e'lon qildingiz.",
			zhCN: "您已成功发布您的意图。",
		},
		TogdBackToActivities: {
			deDE: "🔙 Zurück zu Aktivitäten",
			enUK: "🔙 Back to Activities",
			esES: "🔙 Volver a Actividades",
			faIR: "🔙 بازگشت به فعالیت\u200cها",
			frFR: "🔙 Retour aux Activités",
			idID: "🔙 Kembali ke Aktivitas",
			itIT: "🔙 Torna alle Attività",
			jaJP: "🔙 アクティビティに戻る",
			koKR: "🔙 활동으로 돌아가기",
			plPL: "🔙 Powrót do Aktywności",
			ptBR: "🔙 Voltar às Atividades",
			ruRU: "🔙 Назад к Активностям",
			trTR: "🔙 Aktivitelere Dön",
			ukUA: "🔙 Назад до Активностей",
			uzUZ: "🔙 Faoliyatlarga qaytish",
			zhCN: "🔙 返回活动",
		},
		TogdPlansButtonText: {
			deDE: "📝 Pläne",
			enUK: "📝 Plans",
			esES: "📝 Planes",
			faIR: "📝 برنامه\u200cها",
			frFR: "📝 Plans",
			idID: "📝 Rencana",
			itIT: "📝 Piani",
			jaJP: "📝 プラン",
			koKR: "📝 계획",
			plPL: "📝 Plany",
			ptBR: "📝 Planos",
			ruRU: "📝 Планы",
			trTR: "📝 Planlar",
			ukUA: "📝 Плани",
			uzUZ: "📝 Rejalar",
			zhCN: "📝 计划",
		},
		TogdSpotsButtonText: {
			deDE: "📍 Orte",
			enUK: "📍 Spots",
			esES: "📍 Lugares",
			faIR: "📍 مکان\u200cها",
			frFR: "📍 Lieux",
			idID: "📍 Tempat",
			itIT: "📍 Luoghi",
			jaJP: "📍 スポット",
			koKR: "📍 장소",
			plPL: "📍 Miejsca",
			ptBR: "📍 Locais",
			ptPT: "📍 Locais",
			ruRU: "📍 Места",
			trTR: "📍 Yerler",
			ukUA: "📍 Місця",
			uzUZ: "📍 Joylar",
			zhCN: "📍 地点",
		},
		RsvpQuestionOnWhatDate: {
			deDE: "An welchem Tag werden Sie teilnehmen?",
			enUK: "On what day are you going to attend?",
			esES: "¿Qué día vas a asistir?",
			faIR: "چه روزی قصد شرکت دارید؟",
			frFR: "Quel jour allez-vous assister ?",
			idID: "Pada hari apa Anda akan hadir?",
			itIT: "In che giorno parteciperai?",
			jaJP: "何日に参加する予定ですか？",
			koKR: "어느 날에 참석할 예정인가요?",
			plPL: "W którym dniu zamierzasz uczestniczyć?",
			ptBR: "Em que dia você vai comparecer?",
			ruRU: "В какой день вы собираетесь присутствовать?",
			trTR: "Hangi gün katılacaksınız?",
			ukUA: "У який день ви збираєтеся відвідати?",
			uzUZ: "Qaysi kuni qatnashmoqchisiz?",
			zhCN: "您打算哪一天参加？",
		},
		RsvpQuestionAtWhatTime: {
			deDE: "Um wie viel Uhr werden Sie ankommen?",
			enUK: "At what time are you going to arrive?",
			esES: "¿A qué hora vas a llegar?",
			faIR: "در چه ساعتی خواهید رسید؟",
			frFR: "À quelle heure allez-vous arriver ?",
			idID: "Pada jam berapa Anda akan tiba?",
			itIT: "A che ora arriverai?",
			jaJP: "何時に到着予定ですか？",
			koKR: "몇 시에 도착할 예정인가요?",
			plPL: "O której godzinie zamierzasz przyjść?",
			ptBR: "A que horas você vai chegar?",
			ruRU: "Во сколько вы собираетесь прибыть?",
			trTR: "Saat kaçta geleceksiniz?",
			ukUA: "О котрій годині ви збираєтеся прибути?",
			uzUZ: "Soat nechada kelasiz?",
			zhCN: "您几点到达？",
		},
		RsvpTimeIsChangeable: {
			deDE: "Sie können die Minuten bei Bedarf später ändern.",
			enUK: "You would be able to change minutes if needed later.",
			esES: "Podrás cambiar los minutos si es necesario más tarde.",
			faIR: "در صورت نیاز بعداً می\u200cتوانید دقایق را تغییر دهید.",
			frFR: "Vous pourrez modifier les minutes plus tard si nécessaire.",
			idID: "Anda dapat mengubah menit jika diperlukan nanti.",
			itIT: "Potrai cambiare i minuti se necessario più tardi.",
			jaJP: "必要に応じて後で分数を変更できます。",
			koKR: "필요하면 나중에 분을 변경할 수 있습니다.",
			plPL: "W razie potrzeby będziesz mógł później zmienić minuty.",
			ptBR: "Você poderá alterar os minutos se necessário mais tarde.",
			ruRU: "При необходимости вы сможете изменить минуты позже.",
			trTR: "Gerekirse daha sonra dakikaları değiştirebileceksiniz.",
			ukUA: "За потреби ви зможете змінити хвилини пізніше.",
			uzUZ: "Kerak bo'lsa keyinroq daqiqalarni o'zgartirishingiz mumkin.",
			zhCN: "如果需要，您稍后可以更改分钟数。",
		},
		RsvpResponse100Percent: {
			deDE: "Ich werde da sein 💯%",
			enUK: "I'll be there 💯%",
			esES: "Estaré allí 💯%",
			faIR: "💯% آنجا خواهم بود",
			frFR: "Je serai là 💯%",
			idID: "Saya akan ada di sana 💯%",
			itIT: "Ci sarò 💯%",
			jaJP: "💯% そこにいます",
			koKR: "💯% 거기에 있을게요",
			plPL: "Będę tam 💯%",
			ptBR: "Estarei lá 💯%",
			ruRU: "Я буду там 💯%",
			trTR: "Orada olacağım 💯%",
			ukUA: "Я буду там 💯%",
			uzUZ: "Men u yerda bo'laman 💯%",
			zhCN: "我会在那里 💯%",
		},
		RsvpResponseNotAttending: {
			deDE: "Nicht teilnehmend",
			enUK: "Not attending",
			esES: "No asistiré",
			faIR: "شرکت نمی\u200cکنم",
			frFR: "Ne participe pas",
			idID: "Tidak hadir",
			itIT: "Non partecipo",
			jaJP: "参加しません",
			koKR: "참석하지 않음",
			plPL: "Nie biorę udziału",
			ptBR: "Não vou participar",
			ruRU: "Не участвую",
			trTR: "Katılmıyorum",
			ukUA: "Не беру участь",
			uzUZ: "Qatnashmayman",
			zhCN: "不参加",
		},
		RsvpResponseMostLikely: {
			deDE: "Höchstwahrscheinlich",
			enUK: "Most likely",
			esES: "Muy probable",
			faIR: "خیلی محتمل",
			frFR: "Très probablement",
			idID: "Kemungkinan besar",
			itIT: "Molto probabilmente",
			jaJP: "おそらく",
			koKR: "아마도",
			plPL: "Najprawdopodobniej",
			ptBR: "Muito provável",
			ruRU: "Скорее всего",
			trTR: "Büyük ihtimalle",
			ukUA: "Найімовірніше",
			uzUZ: "Katta ehtimol bilan",
			zhCN: "很可能",
		},
		RsvpResponseMaybe: {
			deDE: "Vielleicht",
			enUK: "Maybe",
			esES: "Tal vez",
			faIR: "شاید",
			frFR: "Peut-être",
			idID: "Mungkin",
			itIT: "Forse",
			jaJP: "たぶん",
			koKR: "아마",
			plPL: "Może",
			ptBR: "Talvez",
			ptPT: "Talvez",
			ruRU: "Возможно",
			trTR: "Belki",
			ukUA: "Можливо",
			uzUZ: "Balki",
			zhCN: "也许",
		},
		RsvpResponseUnlikely: {
			deDE: "Unwahrscheinlich",
			enUK: "Unlikely",
			esES: "Poco probable",
			faIR: "بعید",
			frFR: "Peu probable",
			idID: "Tidak mungkin",
			itIT: "Improbabile",
			jaJP: "ありそうにない",
			koKR: "가능성 낮음",
			plPL: "Mało prawdopodobne",
			ptBR: "Improvável",
			ruRU: "Маловероятно",
			trTR: "Pek olası değil",
			ukUA: "Малоймовірно",
			uzUZ: "Kam ehtimol",
			zhCN: "不太可能",
		},
		RsvpHowLikelyQuestion: {
			deDE: "Wie wahrscheinlich ist es, dass Sie dort sein werden?",
			enUK: "How likely is it you are going to be there?",
			esES: "¿Qué tan probable es que vayas a estar allí?",
			faIR: "چقدر احتمال دارد که آنجا باشید؟",
			frFR: "Quelle est la probabilité que vous y soyez ?",
			idID: "Seberapa besar kemungkinan Anda akan ada di sana?",
			itIT: "Quanto è probabile che tu ci sia?",
			jaJP: "そこにいる可能性はどのくらいですか？",
			koKR: "거기에 있을 가능성이 얼마나 됩니까?",
			plPL: "Jak prawdopodobne jest, że tam będziesz?",
			ptBR: "Qual a probabilidade de você estar lá?",
			ruRU: "Насколько вероятно, что вы будете там?",
			trTR: "Orada olma olasılığınız ne kadar?",
			ukUA: "Наскільки ймовірно, що ви будете там?",
			uzUZ: "U yerda bo'lish ehtimoli qanchalik?",
			zhCN: "您去那里的可能性有多大？",
		},
		SpotTitle: {
			deDE: "Ort: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			enUK: "Spot: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			esES: "Lugar: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			faIR: "مکان: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			frFR: "Lieu: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			idID: "Tempat: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			itIT: "Posto: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			jaJP: "スポット: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			koKR: "장소: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			plPL: "Miejsce: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			ptBR: "Local: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			ruRU: "Место: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			trTR: "Nokta: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			ukUA: "Місце: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			uzUZ: "Joy: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			zhCN: "地点: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		},
		SpotActivities: {
			deDE: "Aktivitäten",
			enUK: "Activities",
			esES: "Actividades",
			faIR: "فعالیت\u200cها",
			frFR: "Activités",
			idID: "Aktivitas",
			itIT: "Attività",
			jaJP: "アクティビティ",
			koKR: "활동",
			plPL: "Aktywności",
			ptBR: "Atividades",
			ruRU: "Активности",
			trTR: "Etkinlikler",
			ukUA: "Активності",
			uzUZ: "Faoliyatlar",
			zhCN: "活动",
		},
		SpotButtonText: {
			deDE: "Platz: {SPOT_TITLE}",
			enUK: "Spot: {SPOT_TITLE}",
			esES: "Lugar: {SPOT_TITLE}",
			faIR: "مکان: {SPOT_TITLE}",
			frFR: "Lieu: {SPOT_TITLE}",
			idID: "Tempat: {SPOT_TITLE}",
			itIT: "Posto: {SPOT_TITLE}",
			jaJP: "スポット: {SPOT_TITLE}",
			koKR: "장소: {SPOT_TITLE}",
			plPL: "Miejsce: {SPOT_TITLE}",
			ptBR: "Local: {SPOT_TITLE}",
			ruRU: "Место: {SPOT_TITLE}",
			trTR: "Yer: {SPOT_TITLE}",
			ukUA: "Місце: {SPOT_TITLE}",
			uzUZ: "Joy: {SPOT_TITLE}",
			zhCN: "地点: {SPOT_TITLE}",
		},
		ShareSpotButtonText: {
			deDE: "📤 Ort teilen",
			enUK: "📤 Share spot",
			esES: "📤 Compartir lugar",
			faIR: "📤 اشتراک\u200cگذاری مکان",
			frFR: "📤 Partager lieu",
			idID: "📤 Bagikan tempat",
			itIT: "📤 Condividi posto",
			jaJP: "📤 スポットを共有",
			koKR: "📤 장소 공유",
			plPL: "📤 Udostępnij miejsce",
			ptBR: "📤 Compartilhar local",
			ruRU: "📤 Поделиться местом",
			trTR: "📤 Noktayı paylaş",
			ukUA: "📤 Поділитися місцем",
			uzUZ: "📤 Joyni baham ko'rish",
			zhCN: "📤 分享地点",
		},
		PlanEventButtonText: {
			deDE: "🎯 Event planen",
			enUK: "🎯 Plan event",
			esES: "🎯 Planificar evento",
			faIR: "🎯 برنامه\u200cریزی رویداد",
			frFR: "🎯 Planifier événement",
			idID: "🎯 Rencanakan acara",
			itIT: "🎯 Pianifica evento",
			jaJP: "🎯 イベントを計画",
			koKR: "🎯 이벤트 계획",
			plPL: "🎯 Zaplanuj wydarzenie",
			ptBR: "🎯 Planejar evento",
			ruRU: "🎯 Запланировать событие",
			trTR: "🎯 Etkinlik planla",
			ukUA: "🎯 Запланувати подію",
			uzUZ: "🎯 Tadbir rejalashtirish",
			zhCN: "🎯 计划活动",
		},
		RemoveFromSpots: {
			deDE: "💔 Aus Favoriten entfernen",
			enUK: "💔 Remove from spots",
			esES: "💔 Quitar de favoritos",
			faIR: "💔 حذف از مکان\u200cهای منتخب",
			frFR: "💔 Supprimer des favoris",
			idID: "💔 Hapus dari tempat favorit",
			itIT: "💔 Rimuovi dai preferiti",
			jaJP: "💔 お気に入りから削除",
			koKR: "💔 즐겨찾기에서 제거",
			plPL: "💔 Usuń z ulubionych",
			ptBR: "💔 Remover dos favoritos",
			ptPT: "💔 Remover dos favoritos",
			ruRU: "💔 Удалить из избранного",
			trTR: "💔 Favorilerden kaldır",
			ukUA: "💔 Видалити з обраного",
			uzUZ: "💔 Sevimlilardan olib tashlash",
			zhCN: "💔 从收藏中移除",
		},
		AddToMySpots: {
			deDE: "💛 Zu meinen Favoriten hinzufügen",
			enUK: "💛 Add to my spots",
			esES: "💛 Añadir a mis favoritos",
			faIR: "💛 افزودن به مکان\u200cهای من",
			frFR: "💛 Ajouter à mes favoris",
			idID: "💛 Tambah ke tempat favorit saya",
			itIT: "💛 Aggiungi ai miei preferiti",
			jaJP: "💛 お気に入りに追加",
			koKR: "💛 내 즐겨찾기에 추가",
			plPL: "💛 Dodaj do moich ulubionych",
			ptBR: "💛 Adicionar aos meus favoritos",
			ptPT: "💛 Adicionar aos meus favoritos",
			ruRU: "💛 Добавить в мои места",
			trTR: "💛 Favorilerime ekle",
			ukUA: "💛 Додати до моїх місць",
			uzUZ: "💛 Mening sevimlilarimga qo'shish",
			zhCN: "💛 添加到我的收藏",
		},
		FollowButtonText: {
			deDE: "👀 Folgen",
			enUK: "👀 Follow",
			esES: "👀 Seguir",
			faIR: "👀 دنبال کردن",
			frFR: "👀 Suivre",
			idID: "👀 Ikuti",
			itIT: "👀 Segui",
			jaJP: "👀 フォロー",
			koKR: "👀 팔로우",
			plPL: "👀 Obserwuj",
			ptBR: "👀 Seguir",
			ruRU: "👀 Подписаться",
			trTR: "👀 Takip et",
			ukUA: "👀 Слідкувати",
			uzUZ: "👀 Kuzatish",
			zhCN: "👀 关注",
		},
		UnfollowButtonText: {
			deDE: "👀 Nicht mehr folgen",
			enUK: "👀 Unfollow",
			esES: "👀 Dejar de seguir",
			faIR: "👀 لغو دنبال کردن",
			frFR: "👀 Ne plus suivre",
			idID: "👀 Berhenti mengikuti",
			itIT: "👀 Non seguire più",
			jaJP: "👀 フォロー解除",
			koKR: "👀 언팔로우",
			plPL: "👀 Przestań obserwować",
			ptBR: "👀 Deixar de seguir",
			ruRU: "👀 Отписаться",
			trTR: "👀 Takibi bırak",
			ukUA: "👀 Перестати слідкувати",
			uzUZ: "👀 Kuzatishni to'xtatish",
			zhCN: "👀 取消关注",
		},
		EventTitle: {
			deDE: "Veranstaltung",
			enUK: "Event",
			esES: "Evento",
			faIR: "رویداد",
			frFR: "Événement",
			idID: "Acara",
			itIT: "Evento",
			jaJP: "イベント",
			koKR: "이벤트",
			plPL: "Wydarzenie",
			ptBR: "Evento",
			ruRU: "Мероприятие",
			trTR: "Etkinlik",
			ukUA: "Подія",
			uzUZ: "Tadbir",
			zhCN: "活动",
		},
		EventCreated: {
			deDE: "Veranstaltung erstellt",
			enUK: "Event created",
			esES: "Evento creado",
			faIR: "رویداد ایجاد شد",
			frFR: "Événement créé",
			idID: "Acara dibuat",
			itIT: "Evento creato",
			jaJP: "イベントが作成されました",
			koKR: "이벤트 생성됨",
			plPL: "Wydarzenie utworzone",
			ptBR: "Evento criado",
			ruRU: "Событие создано",
			trTR: "Etkinlik oluşturuldu",
			ukUA: "Подію створено",
			uzUZ: "Tadbir yaratildi",
			zhCN: "活动已创建",
		},
		CancelEvent: {
			deDE: "Veranstaltung absagen",
			enUK: "Cancel event",
			esES: "Cancelar evento",
			faIR: "لغو رویداد",
			frFR: "Annuler l'événement",
			idID: "Batalkan acara",
			itIT: "Annulla evento",
			jaJP: "イベントをキャンセル",
			koKR: "이벤트 취소",
			plPL: "Anuluj wydarzenie",
			ptBR: "Cancelar evento",
			ruRU: "Отменить событие",
			trTR: "Etkinliği iptal et",
			ukUA: "Скасувати подію",
			uzUZ: "Tadbirni bekor qilish",
			zhCN: "取消活动",
		},
		BackToEvents: {
			deDE: "Zurück zu Veranstaltungen",
			enUK: "Back to events",
			esES: "Volver a eventos",
			faIR: "بازگشت به رویدادها",
			frFR: "Retour aux événements",
			idID: "Kembali ke acara",
			itIT: "Torna agli eventi",
			jaJP: "イベントに戻る",
			koKR: "이벤트로 돌아가기",
			plPL: "Powrót do wydarzeń",
			ptBR: "Voltar aos eventos",
			ruRU: "Назад к событиям",
			trTR: "Etkinliklere dön",
			ukUA: "Назад до подій",
			uzUZ: "Tadbirlarga qaytish",
			zhCN: "返回活动",
		},
		EventOptionsButton: {
			deDE: "Optionen",
			enUK: "Options",
			esES: "Opciones",
			faIR: "گزینه\u200cها",
			frFR: "Options",
			idID: "Opsi",
			itIT: "Opzioni",
			jaJP: "オプション",
			koKR: "옵션",
			plPL: "Opcje",
			ptBR: "Opções",
			ruRU: "Варианты",
			trTR: "Seçenekler",
			ukUA: "Варіанти",
			uzUZ: "Variantlar",
			zhCN: "选项",
		},
		NewEventOptionButton: {
			deDE: "Option hinzufügen",
			enUK: "Add option",
			esES: "Añadir opción",
			faIR: "افزودن گزینه",
			frFR: "Ajouter une option",
			idID: "Tambah opsi",
			itIT: "Aggiungi opzione",
			jaJP: "オプションを追加",
			koKR: "옵션 추가",
			plPL: "Dodaj opcję",
			ptBR: "Adicionar opção",
			ruRU: "Добавить вариант",
			trTR: "Seçenek ekle",
			ukUA: "Додати варіант",
			uzUZ: "Variant qo'shish",
			zhCN: "添加选项",
		},
		EventVisibility: {
			deDE: "Sichtbarkeit: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			enUK: "Visibility: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			esES: "Visibilidad: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			faIR: "نمایان\u200cبودن: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			frFR: "Visibilité: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			idID: "Visibilitas: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			itIT: "Visibilità: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			jaJP: "表示設定: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			koKR: "공개 범위: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			plPL: "Widoczność: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			ptBR: "Visibilidade: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			ruRU: "Видимость: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			trTR: "Görünürlük: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			ukUA: "Видимість: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			uzUZ: "Ko'rinish: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			zhCN: "可见性: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		},
		EventStatus: {
			deDE: "Status: {EVENT_STATUS}",
			enUK: "Status: {EVENT_STATUS}",
			esES: "Estado: {EVENT_STATUS}",
			faIR: "وضعیت: {EVENT_STATUS}",
			frFR: "Statut: {EVENT_STATUS}",
			idID: "Status: {EVENT_STATUS}",
			itIT: "Stato: {EVENT_STATUS}",
			jaJP: "ステータス: {EVENT_STATUS}",
			koKR: "상태: {EVENT_STATUS}",
			plPL: "Status: {EVENT_STATUS}",
			ptBR: "Status: {EVENT_STATUS}",
			ruRU: "Статус: {EVENT_STATUS}",
			trTR: "Durum: {EVENT_STATUS}",
			ukUA: "Статус: {EVENT_STATUS}",
			uzUZ: "Holat: {EVENT_STATUS}",
			zhCN: "状态: {EVENT_STATUS}",
		},
		TogdMyProfile: {
			deDE: "Mein Profil",
			enUK: "My profile",
			esES: "Mi perfil",
			faIR: "پروفایل من",
			frFR: "Mon profil",
			idID: "Profil saya",
			itIT: "Il mio profilo",
			jaJP: "マイプロフィール",
			koKR: "내 프로필",
			plPL: "Mój profil",
			ptBR: "Meu perfil",
			ruRU: "Мой профиль",
			trTR: "Benim profilim",
			ukUA: "Мій профіль",
			uzUZ: "Mening profilim",
			zhCN: "我的资料",
		},
		TogdMySpots: {
			deDE: "Meine Orte",
			enUK: "My spots",
			esES: "Mis lugares",
			faIR: "مکان\u200cهای من",
			frFR: "Mes endroits",
			idID: "Tempat saya",
			itIT: "I miei posti",
			jaJP: "マイスポット",
			koKR: "내 장소들",
			plPL: "Moje miejsca",
			ptBR: "Meus locais",
			ruRU: "Мои места",
			trTR: "Benim yerlerim",
			ukUA: "Мої місця",
			uzUZ: "Mening joylarim",
			zhCN: "我的地点",
		},

		TogdMyEvents: {
			deDE: "Meine Events",
			enUK: "My events",
			esES: "Mis eventos",
			faIR: "رویدادهای من",
			frFR: "Mes événements",
			idID: "Acara saya",
			itIT: "I miei eventi",
			jaJP: "マイイベント",
			koKR: "내 이벤트들",
			plPL: "Moje wydarzenia",
			ptBR: "Meus eventos",
			ruRU: "Мои события",
			trTR: "Benim etkinliklerim",
			ukUA: "Мої події",
			uzUZ: "Mening tadbirlarim",
			zhCN: "我的活动",
		},
		TogdMyActivities: {
			deDE: "Meine Aktivitäten",
			enUK: "My activities",
			esES: "Mis actividades",
			faIR: "فعالیت\u200cهای من",
			frFR: "Mes activités",
			idID: "Aktivitas saya",
			itIT: "Le mie attività",
			jaJP: "マイアクティビティ",
			koKR: "내 활동들",
			plPL: "Moje aktywności",
			ptBR: "Minhas atividades",
			ruRU: "Мои активности",
			trTR: "Benim aktivitelerim",
			ukUA: "Мої активності",
			uzUZ: "Mening faoliyatlarim",
			zhCN: "我的活动",
		},
		TogdMyPlans: {
			deDE: "Meine Pläne",
			enUK: "My plans",
			esES: "Mis planes",
			faIR: "برنامه\u200cهای من",
			frFR: "Mes projets",
			idID: "Rencana saya",
			itIT: "I miei piani",
			jaJP: "マイプラン",
			koKR: "내 계획들",
			plPL: "Moje plany",
			ptBR: "Meus planos",
			ptPT: "Meus planos",
			ruRU: "Мои планы",
			trTR: "Benim planlarım",
			ukUA: "Мої плани",
			uzUZ: "Mening rejalarim",
			zhCN: "我的计划",
		},
		TogdUserProfile: {
			deDE: "Benutzerprofil",
			enUK: "User profile",
			esES: "Perfil de usuario",
			faIR: "پروفایل کاربر",
			frFR: "Profil utilisateur",
			idID: "Profil pengguna",
			itIT: "Profilo utente",
			jaJP: "ユーザープロフィール",
			koKR: "사용자 프로필",
			plPL: "Profil użytkownika",
			ptBR: "Perfil do usuário",
			ptPT: "Perfil do utilizador",
			ruRU: "Профиль пользователя",
			trTR: "Kullanıcı profili",
			ukUA: "Профіль користувача",
			uzUZ: "Foydalanuvchi profili",
			zhCN: "用户资料",
		},

		TogdActivitiesOfUser: {
			deDE: "Aktivitäten von {USER_NAME}",
			enUK: "Activities of {USER_NAME}",
			esES: "Actividades de {USER_NAME}",
			faIR: "فعالیت\u200cهای {USER_NAME}",
			frFR: "Activités de {USER_NAME}",
			idID: "Aktivitas {USER_NAME}",
			itIT: "Attività di {USER_NAME}",
			jaJP: "{USER_NAME}のアクティビティ",
			koKR: "{USER_NAME}의 활동",
			plPL: "Aktywności użytkownika {USER_NAME}",
			ptBR: "Atividades de {USER_NAME}",
			ruRU: "Активности пользователя {USER_NAME}",
			trTR: "{USER_NAME} etkinlikleri",
			ukUA: "Активності користувача {USER_NAME}",
			uzUZ: "{USER_NAME} faoliyatlari",
			zhCN: "{USER_NAME}的活动",
		},

		YouHaveNoFavoriteActivities: {
			deDE: "Sie haben keine Lieblingsaktivitäten",
			enUK: "You have no favorite activities",
			esES: "No tienes actividades favoritas",
			faIR: "شما هیچ فعالیت مورد علاقه\u200cای ندارید",
			frFR: "Vous n'avez aucune activité favorite",
			idID: "Anda tidak memiliki aktivitas favorit",
			itIT: "Non hai attività preferite",
			jaJP: "お気に入りのアクティビティがありません",
			koKR: "좋아하는 활동이 없습니다",
			plPL: "Nie masz ulubionych aktywności",
			ptBR: "Você não tem atividades favoritas",
			ruRU: "У вас нет любимых активностей",
			trTR: "Favori etkinliğiniz yok",
			ukUA: "У вас немає улюблених активностей",
			uzUZ: "Sizda sevimli faoliyatlar yo'q",
			zhCN: "您没有喜欢的活动",
		},

		InstructionHowToAddActivityInBot: {
			deDE: "Um Aktivitäten hinzuzufügen, senden Sie sie durch Kommas getrennt.",
			enUK: "To add activities send them separated by comma.",
			esES: "Para agregar actividades envíalas separadas por comas.",
			faIR: "برای افزودن فعالیت\u200cها آنها را با کاما جدا کرده ارسال کنید.",
			frFR: "Pour ajouter des activités, envoyez-les séparées par des virgules.",
			idID: "Untuk menambahkan aktivitas, kirim mereka dipisahkan dengan koma.",
			itIT: "Per aggiungere attività, inviale separate da virgole.",
			jaJP: "アクティビティを追加するには、カンマで区切って送信してください。",
			koKR: "활동을 추가하려면 쉼표로 구분해서 보내세요.",
			plPL: "Aby dodać aktywności, wyślij je oddzielone przecinkami.",
			ptBR: "Para adicionar atividades, envie-as separadas por vírgulas.",
			ruRU: "Чтобы добавить активности, отправьте их через запятую.",
			trTR: "Etkinlik eklemek için virgülle ayırarak gönderin.",
			ukUA: "Щоб додати активності, надсилайте їх через кому.",
			uzUZ: "Faoliyatlarni qo'shish uchun ularni vergul bilan ajratib yuboring.",
			zhCN: "要添加活动，请用逗号分隔发送。",
		},

		TogdUserActivities: {
			enUK: "Activities of {USER_NAME}",
			esES: "Actividades del usuario",
			faIR: "فعالیت\u200cهای کاربر",
			frFR: "Activités utilisateur",
			idID: "Aktivitas pengguna",
			itIT: "Attività utente",
			jaJP: "ユーザーアクティビティ",
			koKR: "사용자 활동",
			plPL: "Aktywności użytkownika",
			ptBR: "Atividades do usuário",
			ruRU: "Активности пользователя",
			trTR: "Kullanıcı aktiviteleri",
			ukUA: "Активності користувача",
			uzUZ: "Foydalanuvchi faoliyatlari",
			zhCN: "用户活动",
		},

		TogdUserEvents: {
			deDE: "Benutzerereignisse",
			enUK: "User events",
			esES: "Eventos del usuario",
			faIR: "رویدادهای کاربر",
			frFR: "Événements utilisateur",
			idID: "Acara pengguna",
			itIT: "Eventi utente",
			jaJP: "ユーザーイベント",
			koKR: "사용자 이벤트",
			plPL: "Wydarzenia użytkownika",
			ptBR: "Eventos do usuário",
			ruRU: "События пользователя",
			trTR: "Kullanıcı etkinlikleri",
			ukUA: "Події користувача",
			uzUZ: "Foydalanuvchi tadbirlari",
			zhCN: "用户活动",
		},

		TogdUserPlans: {
			deDE: "Benutzerpläne",
			enUK: "User plans",
			esES: "Planes del usuario",
			faIR: "برنامه\u200cهای کاربر",
			frFR: "Plans utilisateur",
			idID: "Rencana pengguna",
			itIT: "Piani utente",
			jaJP: "ユーザープラン",
			koKR: "사용자 계획",
			plPL: "Plany użytkownika",
			ptBR: "Planos do usuário",
			ruRU: "Планы пользователя",
			trTR: "Kullanıcı planları",
			ukUA: "Плани користувача",
			uzUZ: "Foydalanuvchi rejalari",
			zhCN: "用户计划",
		},

		TogdUserSpots: {
			deDE: "Benutzerorte",
			enUK: "User spots",
			esES: "Lugares del usuario",
			faIR: "مکان\u200cهای کاربر",
			frFR: "Lieux utilisateur",
			idID: "Tempat pengguna",
			itIT: "Luoghi utente",
			jaJP: "ユーザースポット",
			koKR: "사용자 장소",
			plPL: "Miejsca użytkownika",
			ptBR: "Locais do usuário",
			ruRU: "Места пользователя",
			trTR: "Kullanıcı mekanları",
			ukUA: "Місця користувача",
			uzUZ: "Foydalanuvchi joylari",
			zhCN: "用户地点",
		},
		TogdBotWelcome: {
			deDE: `Willkommen bei @ToGetheredBot — Ihrem einfachen und intelligenten Planungsassistenten für Treffen mit Freunden, die Organisation von Gruppenaktivitäten oder einfach um anderen mitzuteilen, wo Sie sein werden. Ob Kitesurfen am Strand, Streetball spielen oder ein lockeres Treffen planen, ToGethered macht die Koordination mühelos.

Der Bot bietet zwei Hauptfunktionen:

	1.	<b>Plan-Sharing</b> – Lassen Sie andere wissen, <b>wo und wann Sie irgendwo sein möchten</b>. Sie können einen Ort und einen Zeitrahmen angeben, sodass andere Ihre Pläne sehen und mitmachen können, wenn sie auch verfügbar sind.

	2.	<b>Event-Koordination</b> – Organisieren Sie einfach Gruppenaktivitäten, indem Sie mehrere Zeit- und Ortsoptionen vorschlagen. Der Bot sammelt Stimmen von Teilnehmern und zeigt, welche Kombinationen am besten funktionieren, damit sich die Gruppe auf einen Plan einigen kann, ohne lange Chat-Threads.

Mit @ToGetheredBot wird Planung sozial, sichtbar und reibungslos — perfekt für spontane Sessions oder organisierte Zusammenkünfte.`,
			enUK: `Welcome to @ToGetheredBot — your simple and smart planning assistant for meeting up with friends, organising group activities, or just letting others know where you'll be. Whether it's kitesurfing at the beach, playing street basketball, or planning a casual meetup, ToGethered makes coordination effortless.

The bot offers two main features:

	1.	<b>Plan Sharing</b> – Let others know <b>where and when you're planning to be somewhere</b>. You can specify a location and a time range, so others can see your plans and join in if they're available too.

	2.	<b>Event Coordination</b> – Easily organise group activities by suggesting multiple time and place options. The bot collects votes from participants and shows which combinations work best, helping the group agree on a plan without long chat threads.

With @ToGetheredBot, planning becomes social, visible, and frictionless — perfect for spontaneous sessions or organised gatherings.`,

			esES: `Bienvenido a @ToGetheredBot — tu asistente de planificación simple e inteligente para quedar con amigos, organizar actividades grupales o simplemente avisar a otros dónde estarás. Ya sea kitesurf en la playa, jugar baloncesto callejero o planificar una reunión casual, ToGethered hace que la coordinación sea sin esfuerzo.

El bot ofrece dos características principales:

	1.	<b>Compartir Planes</b> – Haz saber a otros <b>dónde y cuándo planeas estar en algún lugar</b>. Puedes especificar una ubicación y un rango de tiempo, para que otros puedan ver tus planes y unirse si también están disponibles.

	2.	<b>Coordinación de Eventos</b> – Organiza fácilmente actividades grupales sugiriendo múltiples opciones de tiempo y lugar. El bot recopila votos de los participantes y muestra qué combinaciones funcionan mejor, ayudando al grupo a acordar un plan sin largos hilos de chat.

Con @ToGetheredBot, la planificación se vuelve social, visible y sin fricciones — perfecto para sesiones espontáneas o reuniones organizadas.`,
			faIR: `به @ToGetheredBot خوش آمدید — دستیار برنامه\u200cریزی ساده و هوشمند شما برای ملاقات با دوستان، سازماندهی فعالیت\u200cهای گروهی، یا فقط اطلاع دادن به دیگران که کجا خواهید بود. چه کایت\u200cسرفینگ در ساحل، بازی بسکتبال خیابانی، یا برنامه\u200cریزی برای یک ملاقات غیررسمی، ToGethered هماهنگی را بدون زحمت می\u200cکند.

ربات دو ویژگی اصلی ارائه می\u200cدهد:

	1.	<b>اشتراک\u200cگذاری برنامه</b> – به دیگران اطلاع دهید <b>کجا و چه زمانی قصد حضور در جایی را دارید</b>. می\u200cتوانید مکان و بازه زمانی مشخص کنید تا دیگران بتوانند برنامه\u200cهای شما را ببینند و در صورت در دسترس بودن به شما بپیوندند.

	2.	<b>هماهنگی رویداد</b> – به راحتی فعالیت\u200cهای گروهی را با پیشنهاد گزینه\u200cهای متعدد زمان و مکان سازماندهی کنید. ربات رای\u200cهای شرکت\u200cکنندگان را جمع\u200cآوری می\u200cکند و نشان می\u200cدهد کدام ترکیبات بهترین عملکرد را دارند، و به گروه کمک می\u200cکند بدون رشته\u200cهای چت طولانی روی یک برنامه توافق کنند.

با @ToGetheredBot، برنامه\u200cریزی اجتماعی، قابل مشاهده و بدون اصطکاک می\u200cشود — عالی برای جلسات خودجوش یا گردهمایی\u200cهای سازمان\u200cیافته.`,
			frFR: `Bienvenue sur @ToGetheredBot — votre assistant de planification simple et intelligent pour rencontrer des amis, organiser des activités de groupe, ou simplement faire savoir aux autres où vous serez. Que ce soit du kitesurf à la plage, jouer au basket de rue, ou planifier une rencontre décontractée, ToGethered rend la coordination sans effort.

Le bot offre deux fonctionnalités principales :

	1.	<b>Partage de Plans</b> – Faites savoir aux autres <b>où et quand vous prévoyez d'être quelque part</b>. Vous pouvez spécifier un lieu et une plage horaire, afin que les autres puissent voir vos plans et se joindre s'ils sont également disponibles.

	2.	<b>Coordination d'Événements</b> – Organisez facilement des activités de groupe en suggérant plusieurs options de temps et de lieu. Le bot collecte les votes des participants et montre quelles combinaisons fonctionnent le mieux, aidant le groupe à s'accorder sur un plan sans longs fils de discussion.

Avec @ToGetheredBot, la planification devient sociale, visible et sans friction — parfait pour des sessions spontanées ou des rassemblements organisés.`,
			idID: `Selamat datang di @ToGetheredBot — asisten perencanaan sederhana dan cerdas Anda untuk bertemu dengan teman, mengorganisir kegiatan grup, atau sekadar memberi tahu orang lain di mana Anda akan berada. Baik itu kitesurfing di pantai, bermain basket jalanan, atau merencanakan pertemuan santai, ToGethered membuat koordinasi menjadi mudah.

Bot ini menawarkan dua fitur utama:

	1.	<b>Berbagi Rencana</b> – Beri tahu orang lain <b>di mana dan kapan Anda berencana berada di suatu tempat</b>. Anda dapat menentukan lokasi dan rentang waktu, sehingga orang lain dapat melihat rencana Anda dan bergabung jika mereka juga tersedia.

	2.	<b>Koordinasi Acara</b> – Dengan mudah mengorganisir kegiatan grup dengan menyarankan beberapa pilihan waktu dan tempat. Bot mengumpulkan suara dari peserta dan menunjukkan kombinasi mana yang paling berhasil, membantu grup menyepakati rencana tanpa thread chat yang panjang.

Dengan @ToGetheredBot, perencanaan menjadi sosial, terlihat, dan tanpa gesekan — sempurna untuk sesi spontan atau pertemuan yang terorganisir.`,
			itIT: `Benvenuto su @ToGetheredBot — il tuo assistente di pianificazione semplice e intelligente per incontrare amici, organizzare attività di gruppo, o semplicemente far sapere agli altri dove sarai. Che si tratti di kitesurf in spiaggia, giocare a basket di strada, o pianificare un incontro informale, ToGethered rende il coordinamento senza sforzo.

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
			plPL: `Witaj w @ToGetheredBot — Twoim prostym i inteligentnym asystentem planowania spotkań z przyjaciółmi, organizowania działań grupowych lub po prostu informowania innych, gdzie będziesz. Czy to kitesurfing na plaży, granie w koszykówkę uliczną, czy planowanie nieformalnego spotkania, ToGethered sprawia, że koordynacja staje się bezwysiłkowa.

Bot oferuje dwie główne funkcje:

	1.	<b>Udostępnianie Planów</b> – Poinformuj innych <b>gdzie i kiedy planujesz gdzieś być</b>. Możesz określić lokalizację i zakres czasowy, aby inni mogli zobaczyć Twoje plany i dołączyć, jeśli są również dostępni.

	2.	<b>Koordynacja Wydarzeń</b> – Łatwo organizuj działania grupowe, proponując wiele opcji czasu i miejsca. Bot zbiera głosy od uczestników i pokazuje, które kombinacje działają najlepiej, pomagając grupie uzgodnić plan bez długich wątków czatu.

Z @ToGetheredBot planowanie staje się społeczne, widoczne i bezproblemowe — idealne dla spontanicznych sesji lub zorganizowanych spotkań.`,
			ptBR: `Bem-vindo ao @ToGetheredBot — seu assistente de planejamento simples e inteligente para encontrar amigos, organizar atividades em grupo, ou simplesmente avisar outros onde você estará. Seja kitesurf na praia, jogar basquete de rua, ou planejar um encontro casual, ToGethered torna a coordenação sem esforço.

O bot oferece duas funcionalidades principais:

	1.	<b>Compartilhamento de Planos</b> – Deixe outros saberem <b>onde e quando você planeja estar em algum lugar</b>. Você pode especificar uma localização e um intervalo de tempo, para que outros possam ver seus planos e participar se também estiverem disponíveis.

	2.	<b>Coordenação de Eventos</b> – Organize facilmente atividades em grupo sugerindo múltiplas opções de tempo e local. O bot coleta votos dos participantes e mostra quais combinações funcionam melhor, ajudando o grupo a concordar com um plano sem longas discussões no chat.

Com @ToGetheredBot, o planejamento se torna social, visível e sem atritos — perfeito para sessões espontâneas ou encontros organizados.`,
			ruRU: `Добро пожаловать в @ToGetheredBot — ваш простой и умный помощник для планирования встреч с друзьями, организации групповых мероприятий или просто для того, чтобы сообщить другим, где вы будете. Будь то кайтсёрфинг на пляже, игра в уличный баскетбол или планирование неформальной встречи, ToGethered делает координацию лёгкой.

Бот предлагает две основные функции:

	1.	<b>Обмен планами</b> – Дайте другим знать, <b>где и когда вы планируете где-то быть</b>. Вы можете указать место и временной диапазон, чтобы другие могли видеть ваши планы и присоединиться, если они тоже свободны.

	2.	<b>Координация событий</b> – Легко организуйте групповые мероприятия, предлагая несколько вариантов времени и места. Бот собирает голоса участников и показывает, какие комбинации работают лучше всего, помогая группе договориться о плане без длинных цепочек сообщений.

С @ToGetheredBot планирование становится социальным, видимым и беспрепятственным — идеально подходит для спонтанных сессий или организованных встреч.`,
			trTR: `@ToGetheredBot'a hoş geldiniz — arkadaşlarınızla buluşmak, grup etkinlikleri düzenlemek veya sadece başkalarına nerede olacağınızı bildirmek için basit ve akıllı planlama asistanınız. İster sahilde rüzgar sörfü, ister sokak basketbolu oynamak, ister gündelik bir buluşma planlamak olsun, ToGethered koordinasyonu zahmetsiz hale getirir.

Bot iki ana özellik sunar:

	1.	<b>Plan Paylaşımı</b> – Başkalarına <b>nerede ve ne zaman bir yerde olmayı planladığınızı</b> bildirin. Bir konum ve zaman aralığı belirtebilirsiniz, böylece diğerleri planlarınızı görebilir ve onlar da müsaitse katılabilir.

	2.	<b>Etkinlik Koordinasyonu</b> – Birden fazla zaman ve yer seçeneği önererek grup etkinliklerini kolayca düzenleyin. Bot katılımcılardan oyları toplar ve hangi kombinasyonların en iyi çalıştığını gösterir, grubun uzun sohbet zincirleri olmadan bir plan üzerinde anlaşmasına yardımcı olur.

@ToGetheredBot ile planlama sosyal, görünür ve sürtünmesiz hale gelir — spontane oturumlar veya organize toplantılar için mükemmel.`,
			ukUA: `Ласкаво просимо до @ToGetheredBot — вашого простого та розумного помічника для планування зустрічей з друзями, організації групових заходів або просто повідомлення іншим, де ви будете. Чи то кайтсерфінг на пляжі, гра в вуличний баскетбол, чи планування неформальної зустрічі, ToGethered робить координацію легкою.

Бот пропонує дві основні функції:

	1.	<b>Обмін планами</b> – Дайте іншим знати, <b>де і коли ви плануєте десь бути</b>. Ви можете вказати місце та часовий діапазон, щоб інші могли бачити ваші плани та приєднатися, якщо вони теж вільні.

	2.	<b>Координація подій</b> – Легко організовуйте групові заходи, пропонуючи декілька варіантів часу та місця. Бот збирає голоси учасників і показує, які комбінації працюють найкраще, допомагаючи групі домовитися про план без довгих ланцюжків повідомлень.

З @ToGetheredBot планування стає соціальним, видимим та безперешкодним — ідеально підходить для спонтанних сесій або організованих зустрічей.`,
			uzUZ: `@ToGetheredBot'ga xush kelibsiz — do'stlaringiz bilan uchrashish, guruh faoliyatlarini tashkil qilish yoki boshqalarga qayerda bo'lishingizni bildirish uchun oddiy va aqlli rejalashtirish yordamchingiz. Plyajda kite surfing, ko'cha basketboli o'ynash yoki oddiy uchrashuv rejalashtirish bo'ladimi, ToGethered muvofiqlashtirish jarayonini osonlashtiradi.

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
			deDE: `
Um Ihre Pläne zu teilen, wählen Sie einen Ort oder eine Aktivität und klicken Sie auf die <b>{RSVP}</b>-Schaltfläche.

Um ein Event zu organisieren, erstellen Sie ein <b>{NEW_EVENT}</b> unter "Meine Events".
`,
			enUK: `
To share your plans choose a spot or activity and hit the <b>{RSVP}</b> button.

To organize an event create a <b>{NEW_EVENT}</b> from "My Events".
`,
			esES: `
			Para compartir tus planes elige un lugar o actividad y pulsa el botón <b>{RSVP}</b>.

			Para organizar un evento crea un <b>{NEW_EVENT}</b> desde "Mis eventos".
			`,
			faIR: `
			برای به اشتراک گذاری برنامه\u200cهایتان، مکان یا فعالیتی را انتخاب کنید و دکمه <b>{RSVP}</b> را فشار دهید.

			برای سازماندهی یک رویداد، یک <b>{NEW_EVENT}</b> از "رویدادهای من" ایجاد کنید.
			`,
			frFR: `
			Pour partager vos plans, choisissez un lieu ou une activité et cliquez sur le bouton <b>{RSVP}</b>.

			Pour organiser un événement, créez un <b>{NEW_EVENT}</b> depuis "Mes événements".
			`,
			idID: `
			Untuk membagikan rencana Anda, pilih tempat atau aktivitas dan tekan tombol <b>{RSVP}</b>.

			Untuk mengorganisir acara, buat <b>{NEW_EVENT}</b> dari "Acara Saya".
			`,
			itIT: `
			Per condividere i tuoi piani scegli un posto o un'attività e premi il pulsante <b>{RSVP}</b>.

			Per organizzare un evento crea un <b>{NEW_EVENT}</b> da "I miei eventi".
			`,
			jaJP: `
			プランを共有するには、スポットまたはアクティビティを選択して<b>{RSVP}</b>ボタンを押してください。

			イベントを企画するには、「マイイベント」から<b>{NEW_EVENT}</b>を作成してください。
			`,
			koKR: `
			계획을 공유하려면 장소나 활동을 선택하고 <b>{RSVP}</b> 버튼을 누르세요.

			이벤트를 기획하려면 "내 이벤트"에서 <b>{NEW_EVENT}</b>를 만드세요.
			`,
			plPL: `
			Aby udostępnić swoje plany, wybierz miejsce lub aktywność i kliknij przycisk <b>{RSVP}</b>.

			Aby zorganizować wydarzenie, utwórz <b>{NEW_EVENT}</b> z "Moje wydarzenia".
			`,
			ptBR: `
			Para compartilhar seus planos, escolha um local ou atividade e clique no botão <b>{RSVP}</b>.

			Para organizar um evento, crie um <b>{NEW_EVENT}</b> em "Meus eventos".
			`,
			ptPT: `
			Para partilhar os seus planos, escolha um local ou atividade e clique no botão <b>{RSVP}</b>.

			Para organizar um evento, crie um <b>{NEW_EVENT}</b> em "Os meus eventos".
			`,
			ruRU: `
			Чтобы поделиться своими планами, выберите место или активность и нажмите кнопку <b>{RSVP}</b>.

			Чтобы организовать событие, создайте <b>{NEW_EVENT}</b> в разделе "Мои события".
			`,
			trTR: `
			Planlarınızı paylaşmak için bir yer veya aktivite seçin ve <b>{RSVP}</b> düğmesine tıklayın.

			Bir etkinlik düzenlemek için "Etkinliklerim"den <b>{NEW_EVENT}</b> oluşturun.
			`,
			ukUA: `
			Щоб поділитися своїми планами, оберіть місце або активність і натисніть кнопку <b>{RSVP}</b>.

			Щоб організувати подію, створіть <b>{NEW_EVENT}</b> з "Мої події".
			`,
			uzUZ: `
			Rejalaringizni baham ko'rish uchun joy yoki faoliyatni tanlang va <b>{RSVP}</b> tugmasini bosing.

			Tadbir tashkil qilish uchun "Mening tadbirlarim"dan <b>{NEW_EVENT}</b> yarating.
			`,
			zhCN: `
			要分享您的计划，请选择一个地点或活动，然后点击<b>{RSVP}</b>按钮。

			要组织活动，请从"我的活动"创建<b>{NEW_EVENT}</b>。
		`,
		},
		RsvpButtonText: {
			deDE: "Anmelden",
			enUK: "RSVP",
			esES: "Confirmar",
			faIR: "ثبت نام",
			frFR: "S'inscrire",
			idID: "Daftar",
			itIT: "Iscriviti",
			jaJP: "参加登録",
			koKR: "등록하기",
			plPL: "Zapisz się",
			ptBR: "Inscrever-se",
			ptPT: "Inscrever-se",
			ruRU: "Отметиться",
			trTR: "Kayıt ol",
			ukUA: "Зареєструватися",
			uzUZ: "Ro'yxatdan o'tish",
			zhCN: "报名",
		},
		NewEventButtonText: {
			deDE: "Neues Event",
			enUK: "New Event",
			esES: "Nuevo evento",
			faIR: "رویداد جدید",
			frFR: "Nouvel événement",
			idID: "Acara baru",
			itIT: "Nuovo evento",
			jaJP: "新しいイベント",
			koKR: "새 이벤트",
			plPL: "Nowe wydarzenie",
			ptBR: "Novo evento",
			ruRU: "Новое событие",
			trTR: "Yeni etkinlik",
			ukUA: "Нова подія",
			uzUZ: "Yangi tadbir",
			zhCN: "新活动",
		},
	}
	for k, v := range trans {
		TRANS[k] = v
	}
}
