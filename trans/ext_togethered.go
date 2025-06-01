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
		"de-DE": "🗓️️ Kalender",
		"en-UK": "🗓️️ Calendar",
		"en-US": "🗓️️ Calendar",
		"es-ES": "🗓️️ Calendario",
		"fa-IR": "🗓️️ تقویم",
		"fr-FR": "🗓️️ Calendrier",
		"id-ID": "🗓️️ Kalender",
		"it-IT": "🗓️️ Calendario",
		"ja-JP": "🗓️️ カレンダー",
		"ko-KO": "🗓️️ 달력",
		"pl-PL": "🗓️️ Kalendarz",
		"pt-BR": "🗓️️ Calendário",
		"pt-PT": "🗓️️ Calendário",
		"ru-RU": "🗓️️ Календарь",
		"tr-TR": "🗓️️ Takvim",
		"ua-UA": "🗓️️ Календар",
		"uz-UZ": "🗓️️ Taqvim",
		"zh-CN": "🗓️️ 日历",
	}
}

func init() {
	/*
		locale items should be ordered by key in ascending order,
		use it as a reference for all values:

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
	var trans = map[string]map[string]string{
		CalendarButtonText: calendarButtonText(),
		NewEventTitle: {
			"de-DE": "Neue Veranstaltung",
			"en-UK": "New Event",
			"en-US": "New Event",
			"es-ES": "Nuevo Evento",
			"fa-IR": "رویداد جدید",
			"fr-FR": "Nouvel Événement",
			"id-ID": "Acara Baru",
			"it-IT": "Nuovo Evento",
			"ja-JP": "新しいイベント",
			"ko-KO": "새 이벤트",
			"pl-PL": "Nowe Wydarzenie",
			"pt-BR": "Novo Evento",
			"ru-RU": "Новое Событие",
			"tr-TR": "Yeni Etkinlik",
			"ua-UA": "Нова Подія",
			"uz-UZ": "Yangi Tadbir",
			"zh-CN": "新事件",
		},
		NewEventText: {
			"de-DE": "@{BOT_CODE} hilft dabei, Datum, Uhrzeit und Ort zu wählen, die für alle Teilnehmer am besten funktionieren.",
			"en-UK": "@{BOT_CODE} helps to choose date, time & place that works best for all participants.",
			"en-US": "@{BOT_CODE} helps to choose date, time & place that works best for all participants.",
			"es-ES": "@{BOT_CODE} ayuda a elegir fecha, hora y lugar que funcione mejor para todos los participantes.",
			"fa-IR": "@{BOT_CODE} به انتخاب تاریخ، زمان و مکانی که برای همه شرکت\u200Cکنندگان مناسب است کمک می\u200Cکند.",
			"fr-FR": "@{BOT_CODE} aide à choisir la date, l'heure et le lieu qui conviennent le mieux à tous les participants.",
			"id-ID": "@{BOT_CODE} membantu memilih tanggal, waktu & tempat yang paling cocok untuk semua peserta.",
			"it-IT": "@{BOT_CODE} aiuta a scegliere data, ora e luogo che funzionano meglio per tutti i partecipanti.",
			"ja-JP": "@{BOT_CODE} は、すべての参加者に最適な日付、時間、場所を選ぶのに役立ちます。",
			"ko-KO": "@{BOT_CODE}는 모든 참가자에게 가장 적합한 날짜, 시간 및 장소를 선택하는 데 도움을 줍니다.",
			"pl-PL": "@{BOT_CODE} pomaga wybrać datę, godzinę i miejsce, które najlepiej pasują wszystkim uczestników.",
			"pt-BR": "@{BOT_CODE} ajuda a escolher data, hora e local que funcionam melhor para todos os participantes.",
			"ru-RU": "@{BOT_CODE} помогает выбрать дату, время и место, которые лучше всего подходят для всех участников.",
			"tr-TR": "@{BOT_CODE} tüm katılımcılar için en uygun tarih, saat ve yeri seçmeye yardımcı olur.",
			"ua-UA": "@{BOT_CODE} допомагає обрати дату, час і місце, які найкраще підходять для всіх учасників.",
			"uz-UZ": "@{BOT_CODE} barcha ishtirokchilar uchun eng mos sana, vaqt va joyni tanlashda yordam beradi.",
			"zh-CN": "@{BOT_CODE} 帮助选择最适合所有参与者的日期、时间和地点。",
		},
		NewEventHint: {
			"de-DE": "Geben Sie den Titel Ihrer neuen Veranstaltung ein:",
			"en-UK": "Enter title of your new event:",
			"en-US": "Enter title of your new event:",
			"es-ES": "Ingrese el título de su nuevo evento:",
			"fa-IR": "عنوان رویداد جدید خود را وارد کنید:",
			"fr-FR": "Entrez le titre de votre nouvel événement:",
			"id-ID": "Masukkan judul acara baru Anda:",
			"it-IT": "Inserisci il titolo del tuo nuovo evento:",
			"ja-JP": "新しいイベントのタイトルを入力してください:",
			"ko-KO": "새 이벤트의 제목을 입력하세요:",
			"pl-PL": "Wprowadź tytuł swojego nowego wydarzenia:",
			"pt-BR": "Digite o título do seu novo evento:",
			"ru-RU": "Введите название вашего нового события:",
			"tr-TR": "Yeni etkinliğinizin başlığını girin:",
			"ua-UA": "Введіть назву вашої нової події:",
			"uz-UZ": "Yangi tadbiringizning sarlavhasini kiriting:",
			"zh-CN": "输入您的新事件标题:",
		},
		TodayButtonText: {
			"de-DE": "🕒 Heute — {DATE}",
			"en-UK": "🕒 Today — {DATE}",
			"en-US": "🕒 Today — {DATE}",
			"es-ES": "🕒 Hoy — {DATE}",
			"fa-IR": "🕒 امروز — {DATE}",
			"fr-FR": "🕒 Aujourd'hui — {DATE}",
			"id-ID": "🕒 Hari ini — {DATE}",
			"it-IT": "🕒 Oggi — {DATE}",
			"ja-JP": "🕒 今日 — {DATE}",
			"ko-KO": "🕒 오늘 — {DATE}",
			"pl-PL": "🕒 Dzisiaj — {DATE}",
			"pt-BR": "🕒 Hoje — {DATE}",
			"pt-PT": "🕒 Hoje — {DATE}",
			"ru-RU": "🕒 Сегодня — {DATE}",
			"tr-TR": "🕒 Bugün — {DATE}",
			"ua-UA": "🕒 Сьогодні — {DATE}",
			"uz-UZ": "🕒 Bugun — {DATE}",
			"zh-CN": "🕒 今天 — {DATE}",
		},
		TomorrowButtonText: {
			"de-DE": "🌅 Morgen — {DATE}",
			"en-UK": "🌅 Tomorrow —  {DATE}",
			"en-US": "🌅 Tomorrow —  {DATE}",
			"es-ES": "🌅 Mañana — {DATE}",
			"fa-IR": "🌅 فردا — {DATE}",
			"fr-FR": "🌅 Demain — {DATE}",
			"id-ID": "🌅 Besok — {DATE}",
			"it-IT": "🌅 Domani — {DATE}",
			"ja-JP": "🌅 明日 — {DATE}",
			"ko-KO": "🌅 내일 — {DATE}",
			"pl-PL": "🌅 Jutro — {DATE}",
			"pt-BR": "🌅 Amanhã — {DATE}",
			"ru-RU": "🌅 Завтра — {DATE}",
			"tr-TR": "🌅 Yarın — {DATE}",
			"ua-UA": "🌅 Завтра — {DATE}",
			"uz-UZ": "🌅 Ertaga — {DATE}",
			"zh-CN": "🌅 明天 — {DATE}",
		},
		SpotGoingToDoActivities: {
			"de-DE": "Vorhaben: {ACTIVITIES}",
			"en-UK": "Going to do: {ACTIVITIES}",
			"en-US": "Going to do: {ACTIVITIES}",
			"es-ES": "Voy a hacer: {ACTIVITIES}",
			"fa-IR": "قصد انجام: {ACTIVITIES}",
			"fr-FR": "Va faire: {ACTIVITIES}",
			"id-ID": "Akan melakukan: {ACTIVITIES}",
			"it-IT": "Intenzione di fare: {ACTIVITIES}",
			"ja-JP": "する予定: {ACTIVITIES}",
			"ko-KO": "할 일: {ACTIVITIES}",
			"pl-PL": "Zamierzam zrobić: {ACTIVITIES}",
			"pt-BR": "Vai fazer: {ACTIVITIES}",
			"pt-PT": "Vai fazer: {ACTIVITIES}",
			"ru-RU": "Собираюсь делать: {ACTIVITIES}",
			"tr-TR": "Yapacağım: {ACTIVITIES}",
			"ua-UA": "Збираюся робити: {ACTIVITIES}",
			"uz-UZ": "Qilmoqchi: {ACTIVITIES}",
			"zh-CN": "打算做: {ACTIVITIES}",
		},
		ChooseSpotToRSVP: {
			"de-DE": "Wählen Sie einen Platz zum Zusagen",
			"en-UK": "Choose a spot to RSVP",
			"en-US": "Choose a spot to RSVP",
			"es-ES": "Elige un lugar para confirmar asistencia",
			"fa-IR": "مکانی برای تایید حضور انتخاب کنید",
			"fr-FR": "Choisissez un lieu pour confirmer votre présence",
			"id-ID": "Pilih tempat untuk konfirmasi kehadiran",
			"it-IT": "Scegli un posto per confermare la presenza",
			"ja-JP": "出席返事をする場所を選択してください",
			"ko-KO": "참석 응답할 장소를 선택하세요",
			"pl-PL": "Wybierz miejsce, aby potwierdzić obecność",
			"pt-BR": "Escolha um local para confirmar presença",
			"ru-RU": "Выберите место для подтверждения участия",
			"tr-TR": "Katılımı onaylamak için bir yer seçin",
			"ua-UA": "Оберіть місце для підтвердження участі",
			"uz-UZ": "Ishtirok etishni tasdiqlash uchun joy tanlang",
			"zh-CN": "选择一个地点来回复邀请",
		},
		TogdIntentPublished: {
			"de-DE": "Sie haben Ihre Absicht erfolgreich veröffentlicht.",
			"en-UK": "You've successfully published your intention.",
			"en-US": "You've successfully published your intention.",
			"es-ES": "Has publicado tu intención exitosamente.",
			"fa-IR": "شما با موفقیت قصد خود را منتشر کردید.",
			"fr-FR": "Vous avez publié votre intention avec succès.",
			"id-ID": "Anda telah berhasil menerbitkan niat Anda.",
			"it-IT": "Hai pubblicato con successo la tua intenzione.",
			"ja-JP": "意図を正常に公開しました。",
			"ko-KO": "의도를 성공적으로 게시했습니다.",
			"pl-PL": "Pomyślnie opublikowałeś swoją intencję.",
			"pt-BR": "Você publicou sua intenção com sucesso.",
			"ru-RU": "Вы успешно опубликовали свое намерение.",
			"tr-TR": "Niyetinizi başarıyla yayınladınız.",
			"ua-UA": "Ви успішно опублікували свій намір.",
			"uz-UZ": "Siz o'z niyatingizni muvaffaqiyatli e'lon qildingiz.",
			"zh-CN": "您已成功发布您的意图。",
		},
		TogdBackToActivities: {
			"de-DE": "🔙 Zurück zu Aktivitäten",
			"en-UK": "🔙 Back to Activities",
			"en-US": "🔙 Back to Activities",
			"es-ES": "🔙 Volver a Actividades",
			"fa-IR": "🔙 بازگشت به فعالیت\u200cها",
			"fr-FR": "🔙 Retour aux Activités",
			"id-ID": "🔙 Kembali ke Aktivitas",
			"it-IT": "🔙 Torna alle Attività",
			"ja-JP": "🔙 アクティビティに戻る",
			"ko-KO": "🔙 활동으로 돌아가기",
			"pl-PL": "🔙 Powrót do Aktywności",
			"pt-BR": "🔙 Voltar às Atividades",
			"ru-RU": "🔙 Назад к Активностям",
			"tr-TR": "🔙 Aktivitelere Dön",
			"ua-UA": "🔙 Назад до Активностей",
			"uz-UZ": "🔙 Faoliyatlarga qaytish",
			"zh-CN": "🔙 返回活动",
		},
		TogdPlansButtonText: {
			"de-DE": "📝 Pläne",
			"en-UK": "📝 Plans",
			"en-US": "📝 Plans",
			"es-ES": "📝 Planes",
			"fa-IR": "📝 برنامه\u200cها",
			"fr-FR": "📝 Plans",
			"id-ID": "📝 Rencana",
			"it-IT": "📝 Piani",
			"ja-JP": "📝 プラン",
			"ko-KO": "📝 계획",
			"pl-PL": "📝 Plany",
			"pt-BR": "📝 Planos",
			"ru-RU": "📝 Планы",
			"tr-TR": "📝 Planlar",
			"ua-UA": "📝 Плани",
			"uz-UZ": "📝 Rejalar",
			"zh-CN": "📝 计划",
		},
		TogdSpotsButtonText: {
			"de-DE": "📍 Orte",
			"en-UK": "📍 Spots",
			"en-US": "📍 Spots",
			"es-ES": "📍 Lugares",
			"fa-IR": "📍 مکان\u200cها",
			"fr-FR": "📍 Lieux",
			"id-ID": "📍 Tempat",
			"it-IT": "📍 Luoghi",
			"ja-JP": "📍 スポット",
			"ko-KO": "📍 장소",
			"pl-PL": "📍 Miejsca",
			"pt-BR": "📍 Locais",
			"pt-PT": "📍 Locais",
			"ru-RU": "📍 Места",
			"tr-TR": "📍 Yerler",
			"ua-UA": "📍 Місця",
			"uz-UZ": "📍 Joylar",
			"zh-CN": "📍 地点",
		},
		RsvpQuestionOnWhatDate: {
			"de-DE": "An welchem Tag werden Sie teilnehmen?",
			"en-UK": "On what day are you going to attend?",
			"en-US": "On what day are you going to attend?",
			"es-ES": "¿Qué día vas a asistir?",
			"fa-IR": "چه روزی قصد شرکت دارید؟",
			"fr-FR": "Quel jour allez-vous assister ?",
			"id-ID": "Pada hari apa Anda akan hadir?",
			"it-IT": "In che giorno parteciperai?",
			"ja-JP": "何日に参加する予定ですか？",
			"ko-KO": "어느 날에 참석할 예정인가요?",
			"pl-PL": "W którym dniu zamierzasz uczestniczyć?",
			"pt-BR": "Em que dia você vai comparecer?",
			"ru-RU": "В какой день вы собираетесь присутствовать?",
			"tr-TR": "Hangi gün katılacaksınız?",
			"ua-UA": "У який день ви збираєтеся відвідати?",
			"uz-UZ": "Qaysi kuni qatnashmoqchisiz?",
			"zh-CN": "您打算哪一天参加？",
		},
		RsvpQuestionAtWhatTime: {
			"de-DE": "Um wie viel Uhr werden Sie ankommen?",
			"en-UK": "At what time are you going to arrive?",
			"en-US": "At what time are you going to arrive?",
			"es-ES": "¿A qué hora vas a llegar?",
			"fa-IR": "در چه ساعتی خواهید رسید؟",
			"fr-FR": "À quelle heure allez-vous arriver ?",
			"id-ID": "Pada jam berapa Anda akan tiba?",
			"it-IT": "A che ora arriverai?",
			"ja-JP": "何時に到着予定ですか？",
			"ko-KO": "몇 시에 도착할 예정인가요?",
			"pl-PL": "O której godzinie zamierzasz przyjść?",
			"pt-BR": "A que horas você vai chegar?",
			"ru-RU": "Во сколько вы собираетесь прибыть?",
			"tr-TR": "Saat kaçta geleceksiniz?",
			"ua-UA": "О котрій годині ви збираєтеся прибути?",
			"uz-UZ": "Soat nechada kelasiz?",
			"zh-CN": "您几点到达？",
		},
		RsvpTimeIsChangeable: {
			"de-DE": "Sie können die Minuten bei Bedarf später ändern.",
			"en-UK": "You would be able to change minutes if needed later.",
			"en-US": "You would be able to change minutes if needed later.",
			"es-ES": "Podrás cambiar los minutos si es necesario más tarde.",
			"fa-IR": "در صورت نیاز بعداً می\u200cتوانید دقایق را تغییر دهید.",
			"fr-FR": "Vous pourrez modifier les minutes plus tard si nécessaire.",
			"id-ID": "Anda dapat mengubah menit jika diperlukan nanti.",
			"it-IT": "Potrai cambiare i minuti se necessario più tardi.",
			"ja-JP": "必要に応じて後で分数を変更できます。",
			"ko-KO": "필요하면 나중에 분을 변경할 수 있습니다.",
			"pl-PL": "W razie potrzeby będziesz mógł później zmienić minuty.",
			"pt-BR": "Você poderá alterar os minutos se necessário mais tarde.",
			"ru-RU": "При необходимости вы сможете изменить минуты позже.",
			"tr-TR": "Gerekirse daha sonra dakikaları değiştirebileceksiniz.",
			"ua-UA": "За потреби ви зможете змінити хвилини пізніше.",
			"uz-UZ": "Kerak bo'lsa keyinroq daqiqalarni o'zgartirishingiz mumkin.",
			"zh-CN": "如果需要，您稍后可以更改分钟数。",
		},
		RsvpResponse100Percent: {
			"de-DE": "Ich werde da sein 💯%",
			"en-UK": "I'll be there 💯%",
			"en-US": "I'll be there 💯%",
			"es-ES": "Estaré allí 💯%",
			"fa-IR": "💯% آنجا خواهم بود",
			"fr-FR": "Je serai là 💯%",
			"id-ID": "Saya akan ada di sana 💯%",
			"it-IT": "Ci sarò 💯%",
			"ja-JP": "💯% そこにいます",
			"ko-KO": "💯% 거기에 있을게요",
			"pl-PL": "Będę tam 💯%",
			"pt-BR": "Estarei lá 💯%",
			"ru-RU": "Я буду там 💯%",
			"tr-TR": "Orada olacağım 💯%",
			"ua-UA": "Я буду там 💯%",
			"uz-UZ": "Men u yerda bo'laman 💯%",
			"zh-CN": "我会在那里 💯%",
		},
		RsvpResponseNotAttending: {
			"de-DE": "Nicht teilnehmend",
			"en-UK": "Not attending",
			"en-US": "Not attending",
			"es-ES": "No asistiré",
			"fa-IR": "شرکت نمی\u200cکنم",
			"fr-FR": "Ne participe pas",
			"id-ID": "Tidak hadir",
			"it-IT": "Non partecipo",
			"ja-JP": "参加しません",
			"ko-KO": "참석하지 않음",
			"pl-PL": "Nie biorę udziału",
			"pt-BR": "Não vou participar",
			"ru-RU": "Не участвую",
			"tr-TR": "Katılmıyorum",
			"ua-UA": "Не беру участь",
			"uz-UZ": "Qatnashmayman",
			"zh-CN": "不参加",
		},
		RsvpResponseMostLikely: {
			"de-DE": "Höchstwahrscheinlich",
			"en-UK": "Most likely",
			"en-US": "Most likely",
			"es-ES": "Muy probable",
			"fa-IR": "خیلی محتمل",
			"fr-FR": "Très probablement",
			"id-ID": "Kemungkinan besar",
			"it-IT": "Molto probabilmente",
			"ja-JP": "おそらく",
			"ko-KO": "아마도",
			"pl-PL": "Najprawdopodobniej",
			"pt-BR": "Muito provável",
			"ru-RU": "Скорее всего",
			"tr-TR": "Büyük ihtimalle",
			"ua-UA": "Найімовірніше",
			"uz-UZ": "Katta ehtimol bilan",
			"zh-CN": "很可能",
		},
		RsvpResponseMaybe: {
			"de-DE": "Vielleicht",
			"en-UK": "Maybe",
			"en-US": "Maybe",
			"es-ES": "Tal vez",
			"fa-IR": "شاید",
			"fr-FR": "Peut-être",
			"id-ID": "Mungkin",
			"it-IT": "Forse",
			"ja-JP": "たぶん",
			"ko-KO": "아마",
			"pl-PL": "Może",
			"pt-BR": "Talvez",
			"pt-PT": "Talvez",
			"ru-RU": "Возможно",
			"tr-TR": "Belki",
			"ua-UA": "Можливо",
			"uz-UZ": "Balki",
			"zh-CN": "也许",
		},
		RsvpResponseUnlikely: {
			"de-DE": "Unwahrscheinlich",
			"en-UK": "Unlikely",
			"en-US": "Unlikely",
			"es-ES": "Poco probable",
			"fa-IR": "بعید",
			"fr-FR": "Peu probable",
			"id-ID": "Tidak mungkin",
			"it-IT": "Improbabile",
			"ja-JP": "ありそうにない",
			"ko-KO": "가능성 낮음",
			"pl-PL": "Mało prawdopodobne",
			"pt-BR": "Improvável",
			"ru-RU": "Маловероятно",
			"tr-TR": "Pek olası değil",
			"ua-UA": "Малоймовірно",
			"uz-UZ": "Kam ehtimol",
			"zh-CN": "不太可能",
		},
		RsvpHowLikelyQuestion: {
			"de-DE": "Wie wahrscheinlich ist es, dass Sie dort sein werden?",
			"en-UK": "How likely is it you are going to be there?",
			"en-US": "How likely is it you are going to be there?",
			"es-ES": "¿Qué tan probable es que vayas a estar allí?",
			"fa-IR": "چقدر احتمال دارد که آنجا باشید؟",
			"fr-FR": "Quelle est la probabilité que vous y soyez ?",
			"id-ID": "Seberapa besar kemungkinan Anda akan ada di sana?",
			"it-IT": "Quanto è probabile che tu ci sia?",
			"ja-JP": "そこにいる可能性はどのくらいですか？",
			"ko-KO": "거기에 있을 가능성이 얼마나 됩니까?",
			"pl-PL": "Jak prawdopodobne jest, że tam będziesz?",
			"pt-BR": "Qual a probabilidade de você estar lá?",
			"ru-RU": "Насколько вероятно, что вы будете там?",
			"tr-TR": "Orada olma olasılığınız ne kadar?",
			"ua-UA": "Наскільки ймовірно, що ви будете там?",
			"uz-UZ": "U yerda bo'lish ehtimoli qanchalik?",
			"zh-CN": "您去那里的可能性有多大？",
		},
		SpotTitle: {
			"de-DE": "Ort: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"en-UK": "Spot: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"en-US": "Spot: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"es-ES": "Lugar: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"fa-IR": "مکان: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"fr-FR": "Lieu: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"id-ID": "Tempat: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"it-IT": "Posto: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"ja-JP": "スポット: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"ko-KO": "장소: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"pl-PL": "Miejsce: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"pt-BR": "Local: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"ru-RU": "Место: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"tr-TR": "Nokta: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"ua-UA": "Місце: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"uz-UZ": "Joy: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
			"zh-CN": "地点: <b>{SPOT_TITLE}</b> — <i>{SPOT_LOCATION}</i>",
		},
		SpotActivities: {
			"de-DE": "Aktivitäten",
			"en-UK": "Activities",
			"en-US": "Activities",
			"es-ES": "Actividades",
			"fa-IR": "فعالیت\u200cها",
			"fr-FR": "Activités",
			"id-ID": "Aktivitas",
			"it-IT": "Attività",
			"ja-JP": "アクティビティ",
			"ko-KO": "활동",
			"pl-PL": "Aktywności",
			"pt-BR": "Atividades",
			"ru-RU": "Активности",
			"tr-TR": "Etkinlikler",
			"ua-UA": "Активності",
			"uz-UZ": "Faoliyatlar",
			"zh-CN": "活动",
		},
		SpotButtonText: {
			"de-DE": "Platz: {SPOT_TITLE}",
			"en-UK": "Spot: {SPOT_TITLE}",
			"en-US": "Spot: {SPOT_TITLE}",
			"es-ES": "Lugar: {SPOT_TITLE}",
			"fa-IR": "مکان: {SPOT_TITLE}",
			"fr-FR": "Lieu: {SPOT_TITLE}",
			"id-ID": "Tempat: {SPOT_TITLE}",
			"it-IT": "Posto: {SPOT_TITLE}",
			"ja-JP": "スポット: {SPOT_TITLE}",
			"ko-KO": "장소: {SPOT_TITLE}",
			"pl-PL": "Miejsce: {SPOT_TITLE}",
			"pt-BR": "Local: {SPOT_TITLE}",
			"ru-RU": "Место: {SPOT_TITLE}",
			"tr-TR": "Yer: {SPOT_TITLE}",
			"ua-UA": "Місце: {SPOT_TITLE}",
			"uz-UZ": "Joy: {SPOT_TITLE}",
			"zh-CN": "地点: {SPOT_TITLE}",
		},
		ShareSpotButtonText: {
			"de-DE": "📤 Ort teilen",
			"en-UK": "📤 Share spot",
			"en-US": "📤 Share spot",
			"es-ES": "📤 Compartir lugar",
			"fa-IR": "📤 اشتراک\u200cگذاری مکان",
			"fr-FR": "📤 Partager lieu",
			"id-ID": "📤 Bagikan tempat",
			"it-IT": "📤 Condividi posto",
			"ja-JP": "📤 スポットを共有",
			"ko-KO": "📤 장소 공유",
			"pl-PL": "📤 Udostępnij miejsce",
			"pt-BR": "📤 Compartilhar local",
			"ru-RU": "📤 Поделиться местом",
			"tr-TR": "📤 Noktayı paylaş",
			"ua-UA": "📤 Поділитися місцем",
			"uz-UZ": "📤 Joyni baham ko'rish",
			"zh-CN": "📤 分享地点",
		},
		PlanEventButtonText: {
			"de-DE": "🎯 Event planen",
			"en-UK": "🎯 Plan event",
			"en-US": "🎯 Plan event",
			"es-ES": "🎯 Planificar evento",
			"fa-IR": "🎯 برنامه\u200cریزی رویداد",
			"fr-FR": "🎯 Planifier événement",
			"id-ID": "🎯 Rencanakan acara",
			"it-IT": "🎯 Pianifica evento",
			"ja-JP": "🎯 イベントを計画",
			"ko-KO": "🎯 이벤트 계획",
			"pl-PL": "🎯 Zaplanuj wydarzenie",
			"pt-BR": "🎯 Planejar evento",
			"ru-RU": "🎯 Запланировать событие",
			"tr-TR": "🎯 Etkinlik planla",
			"ua-UA": "🎯 Запланувати подію",
			"uz-UZ": "🎯 Tadbir rejalashtirish",
			"zh-CN": "🎯 计划活动",
		},
		RemoveFromSpots: {
			"de-DE": "💔 Aus Favoriten entfernen",
			"en-UK": "💔 Remove from spots",
			"en-US": "💔 Remove from spots",
			"es-ES": "💔 Quitar de favoritos",
			"fa-IR": "💔 حذف از مکان\u200cهای منتخب",
			"fr-FR": "💔 Supprimer des favoris",
			"id-ID": "💔 Hapus dari tempat favorit",
			"it-IT": "💔 Rimuovi dai preferiti",
			"ja-JP": "💔 お気に入りから削除",
			"ko-KO": "💔 즐겨찾기에서 제거",
			"pl-PL": "💔 Usuń z ulubionych",
			"pt-BR": "💔 Remover dos favoritos",
			"pt-PT": "💔 Remover dos favoritos",
			"ru-RU": "💔 Удалить из избранного",
			"tr-TR": "💔 Favorilerden kaldır",
			"ua-UA": "💔 Видалити з обраного",
			"uz-UZ": "💔 Sevimlilardan olib tashlash",
			"zh-CN": "💔 从收藏中移除",
		},
		AddToMySpots: {
			"de-DE": "💛 Zu meinen Favoriten hinzufügen",
			"en-UK": "💛 Add to my spots",
			"en-US": "💛 Add to my spots",
			"es-ES": "💛 Añadir a mis favoritos",
			"fa-IR": "💛 افزودن به مکان\u200cهای من",
			"fr-FR": "💛 Ajouter à mes favoris",
			"id-ID": "💛 Tambah ke tempat favorit saya",
			"it-IT": "💛 Aggiungi ai miei preferiti",
			"ja-JP": "💛 お気に入りに追加",
			"ko-KO": "💛 내 즐겨찾기에 추가",
			"pl-PL": "💛 Dodaj do moich ulubionych",
			"pt-BR": "💛 Adicionar aos meus favoritos",
			"pt-PT": "💛 Adicionar aos meus favoritos",
			"ru-RU": "💛 Добавить в мои места",
			"tr-TR": "💛 Favorilerime ekle",
			"ua-UA": "💛 Додати до моїх місць",
			"uz-UZ": "💛 Mening sevimlilarimga qo'shish",
			"zh-CN": "💛 添加到我的收藏",
		},
		FollowButtonText: {
			"de-DE": "👀 Folgen",
			"en-UK": "👀 Follow",
			"en-US": "👀 Follow",
			"es-ES": "👀 Seguir",
			"fa-IR": "👀 دنبال کردن",
			"fr-FR": "👀 Suivre",
			"id-ID": "👀 Ikuti",
			"it-IT": "👀 Segui",
			"ja-JP": "👀 フォロー",
			"ko-KO": "👀 팔로우",
			"pl-PL": "👀 Obserwuj",
			"pt-BR": "👀 Seguir",
			"ru-RU": "👀 Подписаться",
			"tr-TR": "👀 Takip et",
			"ua-UA": "👀 Слідкувати",
			"uz-UZ": "👀 Kuzatish",
			"zh-CN": "👀 关注",
		},
		UnfollowButtonText: {
			"de-DE": "👀 Nicht mehr folgen",
			"en-UK": "👀 Unfollow",
			"en-US": "👀 Unfollow",
			"es-ES": "👀 Dejar de seguir",
			"fa-IR": "👀 لغو دنبال کردن",
			"fr-FR": "👀 Ne plus suivre",
			"id-ID": "👀 Berhenti mengikuti",
			"it-IT": "👀 Non seguire più",
			"ja-JP": "👀 フォロー解除",
			"ko-KO": "👀 언팔로우",
			"pl-PL": "👀 Przestań obserwować",
			"pt-BR": "👀 Deixar de seguir",
			"ru-RU": "👀 Отписаться",
			"tr-TR": "👀 Takibi bırak",
			"ua-UA": "👀 Перестати слідкувати",
			"uz-UZ": "👀 Kuzatishni to'xtatish",
			"zh-CN": "👀 取消关注",
		},
		EventTitle: {
			"de-DE": "Veranstaltung",
			"en-UK": "Event",
			"en-US": "Event",
			"es-ES": "Evento",
			"fa-IR": "رویداد",
			"fr-FR": "Événement",
			"id-ID": "Acara",
			"it-IT": "Evento",
			"ja-JP": "イベント",
			"ko-KO": "이벤트",
			"pl-PL": "Wydarzenie",
			"pt-BR": "Evento",
			"ru-RU": "Мероприятие",
			"tr-TR": "Etkinlik",
			"ua-UA": "Подія",
			"uz-UZ": "Tadbir",
			"zh-CN": "活动",
		},
		EventCreated: {
			"de-DE": "Veranstaltung erstellt",
			"en-UK": "Event created",
			"en-US": "Event created",
			"es-ES": "Evento creado",
			"fa-IR": "رویداد ایجاد شد",
			"fr-FR": "Événement créé",
			"id-ID": "Acara dibuat",
			"it-IT": "Evento creato",
			"ja-JP": "イベントが作成されました",
			"ko-KO": "이벤트 생성됨",
			"pl-PL": "Wydarzenie utworzone",
			"pt-BR": "Evento criado",
			"ru-RU": "Событие создано",
			"tr-TR": "Etkinlik oluşturuldu",
			"ua-UA": "Подію створено",
			"uz-UZ": "Tadbir yaratildi",
			"zh-CN": "活动已创建",
		},
		CancelEvent: {
			"de-DE": "Veranstaltung absagen",
			"en-UK": "Cancel event",
			"en-US": "Cancel event",
			"es-ES": "Cancelar evento",
			"fa-IR": "لغو رویداد",
			"fr-FR": "Annuler l'événement",
			"id-ID": "Batalkan acara",
			"it-IT": "Annulla evento",
			"ja-JP": "イベントをキャンセル",
			"ko-KO": "이벤트 취소",
			"pl-PL": "Anuluj wydarzenie",
			"pt-BR": "Cancelar evento",
			"ru-RU": "Отменить событие",
			"tr-TR": "Etkinliği iptal et",
			"ua-UA": "Скасувати подію",
			"uz-UZ": "Tadbirni bekor qilish",
			"zh-CN": "取消活动",
		},
		BackToEvents: {
			"de-DE": "Zurück zu Veranstaltungen",
			"en-UK": "Back to events",
			"en-US": "Back to events",
			"es-ES": "Volver a eventos",
			"fa-IR": "بازگشت به رویدادها",
			"fr-FR": "Retour aux événements",
			"id-ID": "Kembali ke acara",
			"it-IT": "Torna agli eventi",
			"ja-JP": "イベントに戻る",
			"ko-KO": "이벤트로 돌아가기",
			"pl-PL": "Powrót do wydarzeń",
			"pt-BR": "Voltar aos eventos",
			"ru-RU": "Назад к событиям",
			"tr-TR": "Etkinliklere dön",
			"ua-UA": "Назад до подій",
			"uz-UZ": "Tadbirlarga qaytish",
			"zh-CN": "返回活动",
		},
		EventOptionsButton: {
			"de-DE": "Optionen",
			"en-UK": "Options",
			"en-US": "Options",
			"es-ES": "Opciones",
			"fa-IR": "گزینه\u200cها",
			"fr-FR": "Options",
			"id-ID": "Opsi",
			"it-IT": "Opzioni",
			"ja-JP": "オプション",
			"ko-KO": "옵션",
			"pl-PL": "Opcje",
			"pt-BR": "Opções",
			"ru-RU": "Варианты",
			"tr-TR": "Seçenekler",
			"ua-UA": "Варіанти",
			"uz-UZ": "Variantlar",
			"zh-CN": "选项",
		},
		NewEventOptionButton: {
			"de-DE": "Option hinzufügen",
			"en-UK": "Add option",
			"en-US": "Add option",
			"es-ES": "Añadir opción",
			"fa-IR": "افزودن گزینه",
			"fr-FR": "Ajouter une option",
			"id-ID": "Tambah opsi",
			"it-IT": "Aggiungi opzione",
			"ja-JP": "オプションを追加",
			"ko-KO": "옵션 추가",
			"pl-PL": "Dodaj opcję",
			"pt-BR": "Adicionar opção",
			"ru-RU": "Добавить вариант",
			"tr-TR": "Seçenek ekle",
			"ua-UA": "Додати варіант",
			"uz-UZ": "Variant qo'shish",
			"zh-CN": "添加选项",
		},
		EventVisibility: {
			"de-DE": "Sichtbarkeit: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"en-UK": "Visibility: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"en-US": "Visibility: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"es-ES": "Visibilidad: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"fa-IR": "نمایان\u200cبودن: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"fr-FR": "Visibilité: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"id-ID": "Visibilitas: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"it-IT": "Visibilità: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"ja-JP": "表示設定: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"ko-KO": "공개 범위: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"pl-PL": "Widoczność: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"pt-BR": "Visibilidade: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"ru-RU": "Видимость: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"tr-TR": "Görünürlük: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"ua-UA": "Видимість: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"uz-UZ": "Ko'rinish: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
			"zh-CN": "可见性: {VISIBILITY_TITLE} {VISIBILITY_EMOJI}",
		},
		EventStatus: {
			"de-DE": "Status: {EVENT_STATUS}",
			"en-UK": "Status: {EVENT_STATUS}",
			"en-US": "Status: {EVENT_STATUS}",
			"es-ES": "Estado: {EVENT_STATUS}",
			"fa-IR": "وضعیت: {EVENT_STATUS}",
			"fr-FR": "Statut: {EVENT_STATUS}",
			"id-ID": "Status: {EVENT_STATUS}",
			"it-IT": "Stato: {EVENT_STATUS}",
			"ja-JP": "ステータス: {EVENT_STATUS}",
			"ko-KO": "상태: {EVENT_STATUS}",
			"pl-PL": "Status: {EVENT_STATUS}",
			"pt-BR": "Status: {EVENT_STATUS}",
			"ru-RU": "Статус: {EVENT_STATUS}",
			"tr-TR": "Durum: {EVENT_STATUS}",
			"ua-UA": "Статус: {EVENT_STATUS}",
			"uz-UZ": "Holat: {EVENT_STATUS}",
			"zh-CN": "状态: {EVENT_STATUS}",
		},
		TogdMyProfile: {
			"de-DE": "Mein Profil",
			"en-UK": "My profile",
			"en-US": "My profile",
			"es-ES": "Mi perfil",
			"fa-IR": "پروفایل من",
			"fr-FR": "Mon profil",
			"id-ID": "Profil saya",
			"it-IT": "Il mio profilo",
			"ja-JP": "マイプロフィール",
			"ko-KO": "내 프로필",
			"pl-PL": "Mój profil",
			"pt-BR": "Meu perfil",
			"ru-RU": "Мой профиль",
			"tr-TR": "Benim profilim",
			"ua-UA": "Мій профіль",
			"uz-UZ": "Mening profilim",
			"zh-CN": "我的资料",
		},
		TogdMySpots: {
			"de-DE": "Meine Orte",
			"en-UK": "My spots",
			"en-US": "My spots",
			"es-ES": "Mis lugares",
			"fa-IR": "مکان\u200cهای من",
			"fr-FR": "Mes endroits",
			"id-ID": "Tempat saya",
			"it-IT": "I miei posti",
			"ja-JP": "マイスポット",
			"ko-KO": "내 장소들",
			"pl-PL": "Moje miejsca",
			"pt-BR": "Meus locais",
			"ru-RU": "Мои места",
			"tr-TR": "Benim yerlerim",
			"ua-UA": "Мої місця",
			"uz-UZ": "Mening joylarim",
			"zh-CN": "我的地点",
		},

		TogdMyEvents: {
			"de-DE": "Meine Events",
			"en-UK": "My events",
			"en-US": "My events",
			"es-ES": "Mis eventos",
			"fa-IR": "رویدادهای من",
			"fr-FR": "Mes événements",
			"id-ID": "Acara saya",
			"it-IT": "I miei eventi",
			"ja-JP": "マイイベント",
			"ko-KO": "내 이벤트들",
			"pl-PL": "Moje wydarzenia",
			"pt-BR": "Meus eventos",
			"ru-RU": "Мои события",
			"tr-TR": "Benim etkinliklerim",
			"ua-UA": "Мої події",
			"uz-UZ": "Mening tadbirlarim",
			"zh-CN": "我的活动",
		},
		TogdMyActivities: {
			"de-DE": "Meine Aktivitäten",
			"en-UK": "My activities",
			"en-US": "My activities",
			"es-ES": "Mis actividades",
			"fa-IR": "فعالیت\u200cهای من",
			"fr-FR": "Mes activités",
			"id-ID": "Aktivitas saya",
			"it-IT": "Le mie attività",
			"ja-JP": "マイアクティビティ",
			"ko-KO": "내 활동들",
			"pl-PL": "Moje aktywności",
			"pt-BR": "Minhas atividades",
			"ru-RU": "Мои активности",
			"tr-TR": "Benim aktivitelerim",
			"ua-UA": "Мої активності",
			"uz-UZ": "Mening faoliyatlarim",
			"zh-CN": "我的活动",
		},
		TogdMyPlans: {
			"de-DE": "Meine Pläne",
			"en-UK": "My plans",
			"en-US": "My plans",
			"es-ES": "Mis planes",
			"fa-IR": "برنامه\u200cهای من",
			"fr-FR": "Mes projets",
			"id-ID": "Rencana saya",
			"it-IT": "I miei piani",
			"ja-JP": "マイプラン",
			"ko-KO": "내 계획들",
			"pl-PL": "Moje plany",
			"pt-BR": "Meus planos",
			"pt-PT": "Meus planos",
			"ru-RU": "Мои планы",
			"tr-TR": "Benim planlarım",
			"ua-UA": "Мої плани",
			"uz-UZ": "Mening rejalarim",
			"zh-CN": "我的计划",
		},
		TogdUserProfile: {
			"de-DE": "Benutzerprofil",
			"en-UK": "User profile",
			"en-US": "User profile",
			"es-ES": "Perfil de usuario",
			"fa-IR": "پروفایل کاربر",
			"fr-FR": "Profil utilisateur",
			"id-ID": "Profil pengguna",
			"it-IT": "Profilo utente",
			"ja-JP": "ユーザープロフィール",
			"ko-KO": "사용자 프로필",
			"pl-PL": "Profil użytkownika",
			"pt-BR": "Perfil do usuário",
			"pt-PT": "Perfil do utilizador",
			"ru-RU": "Профиль пользователя",
			"tr-TR": "Kullanıcı profili",
			"ua-UA": "Профіль користувача",
			"uz-UZ": "Foydalanuvchi profili",
			"zh-CN": "用户资料",
		},

		TogdActivitiesOfUser: {
			"de-DE": "Aktivitäten von {USER_NAME}",
			"en-UK": "Activities of {USER_NAME}",
			"en-US": "Activities of {USER_NAME}",
			"es-ES": "Actividades de {USER_NAME}",
			"fa-IR": "فعالیت\u200cهای {USER_NAME}",
			"fr-FR": "Activités de {USER_NAME}",
			"id-ID": "Aktivitas {USER_NAME}",
			"it-IT": "Attività di {USER_NAME}",
			"ja-JP": "{USER_NAME}のアクティビティ",
			"ko-KO": "{USER_NAME}의 활동",
			"pl-PL": "Aktywności użytkownika {USER_NAME}",
			"pt-BR": "Atividades de {USER_NAME}",
			"ru-RU": "Активности пользователя {USER_NAME}",
			"tr-TR": "{USER_NAME} etkinlikleri",
			"ua-UA": "Активності користувача {USER_NAME}",
			"uz-UZ": "{USER_NAME} faoliyatlari",
			"zh-CN": "{USER_NAME}的活动",
		},

		YouHaveNoFavoriteActivities: {
			"de-DE": "Sie haben keine Lieblingsaktivitäten",
			"en-UK": "You have no favorite activities",
			"en-US": "You have no favorite activities",
			"es-ES": "No tienes actividades favoritas",
			"fa-IR": "شما هیچ فعالیت مورد علاقه\u200cای ندارید",
			"fr-FR": "Vous n'avez aucune activité favorite",
			"id-ID": "Anda tidak memiliki aktivitas favorit",
			"it-IT": "Non hai attività preferite",
			"ja-JP": "お気に入りのアクティビティがありません",
			"ko-KO": "좋아하는 활동이 없습니다",
			"pl-PL": "Nie masz ulubionych aktywności",
			"pt-BR": "Você não tem atividades favoritas",
			"ru-RU": "У вас нет любимых активностей",
			"tr-TR": "Favori etkinliğiniz yok",
			"ua-UA": "У вас немає улюблених активностей",
			"uz-UZ": "Sizda sevimli faoliyatlar yo'q",
			"zh-CN": "您没有喜欢的活动",
		},

		InstructionHowToAddActivityInBot: {
			"de-DE": "Um Aktivitäten hinzuzufügen, senden Sie sie durch Kommas getrennt.",
			"en-UK": "To add activities send them separated by comma.",
			"en-US": "To add activities send them separated by comma.",
			"es-ES": "Para agregar actividades envíalas separadas por comas.",
			"fa-IR": "برای افزودن فعالیت\u200cها آنها را با کاما جدا کرده ارسال کنید.",
			"fr-FR": "Pour ajouter des activités, envoyez-les séparées par des virgules.",
			"id-ID": "Untuk menambahkan aktivitas, kirim mereka dipisahkan dengan koma.",
			"it-IT": "Per aggiungere attività, inviale separate da virgole.",
			"ja-JP": "アクティビティを追加するには、カンマで区切って送信してください。",
			"ko-KO": "활동을 추가하려면 쉼표로 구분해서 보내세요.",
			"pl-PL": "Aby dodać aktywności, wyślij je oddzielone przecinkami.",
			"pt-BR": "Para adicionar atividades, envie-as separadas por vírgulas.",
			"ru-RU": "Чтобы добавить активности, отправьте их через запятую.",
			"tr-TR": "Etkinlik eklemek için virgülle ayırarak gönderin.",
			"ua-UA": "Щоб додати активності, надсилайте їх через кому.",
			"uz-UZ": "Faoliyatlarni qo'shish uchun ularni vergul bilan ajratib yuboring.",
			"zh-CN": "要添加活动，请用逗号分隔发送。",
		},

		TogdUserActivities: {
			"en-UK": "Activities of {USER_NAME}",
			"en-US": "User activities",
			"es-ES": "Actividades del usuario",
			"fa-IR": "فعالیت\u200cهای کاربر",
			"fr-FR": "Activités utilisateur",
			"id-ID": "Aktivitas pengguna",
			"it-IT": "Attività utente",
			"ja-JP": "ユーザーアクティビティ",
			"ko-KO": "사용자 활동",
			"pl-PL": "Aktywności użytkownika",
			"pt-BR": "Atividades do usuário",
			"ru-RU": "Активности пользователя",
			"tr-TR": "Kullanıcı aktiviteleri",
			"ua-UA": "Активності користувача",
			"uz-UZ": "Foydalanuvchi faoliyatlari",
			"zh-CN": "用户活动",
		},

		TogdUserEvents: {
			"de-DE": "Benutzerereignisse",
			"en-UK": "User events",
			"en-US": "User events",
			"es-ES": "Eventos del usuario",
			"fa-IR": "رویدادهای کاربر",
			"fr-FR": "Événements utilisateur",
			"id-ID": "Acara pengguna",
			"it-IT": "Eventi utente",
			"ja-JP": "ユーザーイベント",
			"ko-KO": "사용자 이벤트",
			"pl-PL": "Wydarzenia użytkownika",
			"pt-BR": "Eventos do usuário",
			"ru-RU": "События пользователя",
			"tr-TR": "Kullanıcı etkinlikleri",
			"ua-UA": "Події користувача",
			"uz-UZ": "Foydalanuvchi tadbirlari",
			"zh-CN": "用户活动",
		},

		TogdUserPlans: {
			"de-DE": "Benutzerpläne",
			"en-UK": "User plans",
			"en-US": "User plans",
			"es-ES": "Planes del usuario",
			"fa-IR": "برنامه\u200cهای کاربر",
			"fr-FR": "Plans utilisateur",
			"id-ID": "Rencana pengguna",
			"it-IT": "Piani utente",
			"ja-JP": "ユーザープラン",
			"ko-KO": "사용자 계획",
			"pl-PL": "Plany użytkownika",
			"pt-BR": "Planos do usuário",
			"ru-RU": "Планы пользователя",
			"tr-TR": "Kullanıcı planları",
			"ua-UA": "Плани користувача",
			"uz-UZ": "Foydalanuvchi rejalari",
			"zh-CN": "用户计划",
		},

		TogdUserSpots: {
			"de-DE": "Benutzerorte",
			"en-UK": "User spots",
			"en-US": "User spots",
			"es-ES": "Lugares del usuario",
			"fa-IR": "مکان\u200cهای کاربر",
			"fr-FR": "Lieux utilisateur",
			"id-ID": "Tempat pengguna",
			"it-IT": "Luoghi utente",
			"ja-JP": "ユーザースポット",
			"ko-KO": "사용자 장소",
			"pl-PL": "Miejsca użytkownika",
			"pt-BR": "Locais do usuário",
			"ru-RU": "Места пользователя",
			"tr-TR": "Kullanıcı mekanları",
			"ua-UA": "Місця користувача",
			"uz-UZ": "Foydalanuvchi joylari",
			"zh-CN": "用户地点",
		},
		TogdBotWelcome: {
			"de-DE": `Willkommen bei @ToGetheredBot — Ihrem einfachen und intelligenten Planungsassistenten für Treffen mit Freunden, die Organisation von Gruppenaktivitäten oder einfach um anderen mitzuteilen, wo Sie sein werden. Ob Kitesurfen am Strand, Streetball spielen oder ein lockeres Treffen planen, ToGethered macht die Koordination mühelos.

Der Bot bietet zwei Hauptfunktionen:

	1.	<b>Plan-Sharing</b> – Lassen Sie andere wissen, <b>wo und wann Sie irgendwo sein möchten</b>. Sie können einen Ort und einen Zeitrahmen angeben, sodass andere Ihre Pläne sehen und mitmachen können, wenn sie auch verfügbar sind.

	2.	<b>Event-Koordination</b> – Organisieren Sie einfach Gruppenaktivitäten, indem Sie mehrere Zeit- und Ortsoptionen vorschlagen. Der Bot sammelt Stimmen von Teilnehmern und zeigt, welche Kombinationen am besten funktionieren, damit sich die Gruppe auf einen Plan einigen kann, ohne lange Chat-Threads.

Mit @ToGetheredBot wird Planung sozial, sichtbar und reibungslos — perfekt für spontane Sessions oder organisierte Zusammenkünfte.`,
			"en-UK": `Welcome to @ToGetheredBot — your simple and smart planning assistant for meeting up with friends, organising group activities, or just letting others know where you'll be. Whether it's kitesurfing at the beach, playing street basketball, or planning a casual meetup, ToGethered makes coordination effortless.

The bot offers two main features:

	1.	<b>Plan Sharing</b> – Let others know <b>where and when you're planning to be somewhere</b>. You can specify a location and a time range, so others can see your plans and join in if they're available too.

	2.	<b>Event Coordination</b> – Easily organise group activities by suggesting multiple time and place options. The bot collects votes from participants and shows which combinations work best, helping the group agree on a plan without long chat threads.

With @ToGetheredBot, planning becomes social, visible, and frictionless — perfect for spontaneous sessions or organised gatherings.`,
			"en-US": `Welcome to @ToGetheredBot — your simple and smart planning assistant for meeting up with friends, organizing group activities, or just letting others know where you'll be. Whether it's kitesurfing at the beach, playing street basketball, or planning a casual meetup, ToGethered makes coordination effortless.

The bot offers two main features:

	1.	<b>Plan Sharing</b> – Let others know <b>where and when you're planning to be somewhere</b>. You can specify a location and a time range, so others can see your plans and join in if they're available too.

	2.	<b>Event Coordination</b> – Easily organize group activities by suggesting multiple time and place options. The bot collects votes from participants and shows which combinations work best, helping the group agree on a plan without long chat threads.

With @ToGetheredBot, planning becomes social, visible, and frictionless — perfect for spontaneous sessions or organized gatherings.`,
			"es-ES": `Bienvenido a @ToGetheredBot — tu asistente de planificación simple e inteligente para quedar con amigos, organizar actividades grupales o simplemente avisar a otros dónde estarás. Ya sea kitesurf en la playa, jugar baloncesto callejero o planificar una reunión casual, ToGethered hace que la coordinación sea sin esfuerzo.

El bot ofrece dos características principales:

	1.	<b>Compartir Planes</b> – Haz saber a otros <b>dónde y cuándo planeas estar en algún lugar</b>. Puedes especificar una ubicación y un rango de tiempo, para que otros puedan ver tus planes y unirse si también están disponibles.

	2.	<b>Coordinación de Eventos</b> – Organiza fácilmente actividades grupales sugiriendo múltiples opciones de tiempo y lugar. El bot recopila votos de los participantes y muestra qué combinaciones funcionan mejor, ayudando al grupo a acordar un plan sin largos hilos de chat.

Con @ToGetheredBot, la planificación se vuelve social, visible y sin fricciones — perfecto para sesiones espontáneas o reuniones organizadas.`,
			"fa-IR": `به @ToGetheredBot خوش آمدید — دستیار برنامه\u200cریزی ساده و هوشمند شما برای ملاقات با دوستان، سازماندهی فعالیت\u200cهای گروهی، یا فقط اطلاع دادن به دیگران که کجا خواهید بود. چه کایت\u200cسرفینگ در ساحل، بازی بسکتبال خیابانی، یا برنامه\u200cریزی برای یک ملاقات غیررسمی، ToGethered هماهنگی را بدون زحمت می\u200cکند.

ربات دو ویژگی اصلی ارائه می\u200cدهد:

	1.	<b>اشتراک\u200cگذاری برنامه</b> – به دیگران اطلاع دهید <b>کجا و چه زمانی قصد حضور در جایی را دارید</b>. می\u200cتوانید مکان و بازه زمانی مشخص کنید تا دیگران بتوانند برنامه\u200cهای شما را ببینند و در صورت در دسترس بودن به شما بپیوندند.

	2.	<b>هماهنگی رویداد</b> – به راحتی فعالیت\u200cهای گروهی را با پیشنهاد گزینه\u200cهای متعدد زمان و مکان سازماندهی کنید. ربات رای\u200cهای شرکت\u200cکنندگان را جمع\u200cآوری می\u200cکند و نشان می\u200cدهد کدام ترکیبات بهترین عملکرد را دارند، و به گروه کمک می\u200cکند بدون رشته\u200cهای چت طولانی روی یک برنامه توافق کنند.

با @ToGetheredBot، برنامه\u200cریزی اجتماعی، قابل مشاهده و بدون اصطکاک می\u200cشود — عالی برای جلسات خودجوش یا گردهمایی\u200cهای سازمان\u200cیافته.`,
			"fr-FR": `Bienvenue sur @ToGetheredBot — votre assistant de planification simple et intelligent pour rencontrer des amis, organiser des activités de groupe, ou simplement faire savoir aux autres où vous serez. Que ce soit du kitesurf à la plage, jouer au basket de rue, ou planifier une rencontre décontractée, ToGethered rend la coordination sans effort.

Le bot offre deux fonctionnalités principales :

	1.	<b>Partage de Plans</b> – Faites savoir aux autres <b>où et quand vous prévoyez d'être quelque part</b>. Vous pouvez spécifier un lieu et une plage horaire, afin que les autres puissent voir vos plans et se joindre s'ils sont également disponibles.

	2.	<b>Coordination d'Événements</b> – Organisez facilement des activités de groupe en suggérant plusieurs options de temps et de lieu. Le bot collecte les votes des participants et montre quelles combinaisons fonctionnent le mieux, aidant le groupe à s'accorder sur un plan sans longs fils de discussion.

Avec @ToGetheredBot, la planification devient sociale, visible et sans friction — parfait pour des sessions spontanées ou des rassemblements organisés.`,
			"id-ID": `Selamat datang di @ToGetheredBot — asisten perencanaan sederhana dan cerdas Anda untuk bertemu dengan teman, mengorganisir kegiatan grup, atau sekadar memberi tahu orang lain di mana Anda akan berada. Baik itu kitesurfing di pantai, bermain basket jalanan, atau merencanakan pertemuan santai, ToGethered membuat koordinasi menjadi mudah.

Bot ini menawarkan dua fitur utama:

	1.	<b>Berbagi Rencana</b> – Beri tahu orang lain <b>di mana dan kapan Anda berencana berada di suatu tempat</b>. Anda dapat menentukan lokasi dan rentang waktu, sehingga orang lain dapat melihat rencana Anda dan bergabung jika mereka juga tersedia.

	2.	<b>Koordinasi Acara</b> – Dengan mudah mengorganisir kegiatan grup dengan menyarankan beberapa pilihan waktu dan tempat. Bot mengumpulkan suara dari peserta dan menunjukkan kombinasi mana yang paling berhasil, membantu grup menyepakati rencana tanpa thread chat yang panjang.

Dengan @ToGetheredBot, perencanaan menjadi sosial, terlihat, dan tanpa gesekan — sempurna untuk sesi spontan atau pertemuan yang terorganisir.`,
			"it-IT": `Benvenuto su @ToGetheredBot — il tuo assistente di pianificazione semplice e intelligente per incontrare amici, organizzare attività di gruppo, o semplicemente far sapere agli altri dove sarai. Che si tratti di kitesurf in spiaggia, giocare a basket di strada, o pianificare un incontro informale, ToGethered rende il coordinamento senza sforzo.

Il bot offre due funzionalità principali:

	1.	<b>Condivisione Piani</b> – Fai sapere agli altri <b>dove e quando hai intenzione di essere da qualche parte</b>. Puoi specificare una posizione e un intervallo di tempo, così gli altri possono vedere i tuoi piani e unirsi se sono disponibili anche loro.

	2.	<b>Coordinamento Eventi</b> – Organizza facilmente attività di gruppo suggerendo multiple opzioni di tempo e luogo. Il bot raccoglie i voti dai partecipanti e mostra quali combinazioni funzionano meglio, aiutando il gruppo a concordare su un piano senza lunghe discussioni in chat.

Con @ToGetheredBot, la pianificazione diventa sociale, visibile e senza attriti — perfetto per sessioni spontanee o raduni organizzati.`,
			"ja-JP": `@ToGetheredBotへようこそ — 友達との待ち合わせ、グループ活動の企画、または単に自分がどこにいるかを他の人に知らせるためのシンプルでスマートな計画アシスタントです。ビーチでのカイトサーフィン、ストリートバスケットボール、またはカジュアルな集まりの計画など、ToGetheredは調整を簡単にします。

ボットは2つの主要機能を提供します：

	1.	<b>プラン共有</b> – <b>どこでいつ何かをする予定かを</b>他の人に知らせます。場所と時間範囲を指定できるので、他の人があなたの計画を見て、利用可能であれば参加できます。

	2.	<b>イベント調整</b> – 複数の時間と場所のオプションを提案してグループ活動を簡単に企画します。ボットは参加者からの投票を収集し、どの組み合わせが最適かを表示し、長いチャットスレッドなしでグループが計画に合意するのを助けます。

@ToGetheredBotで、計画はソーシャルで、見える化され、摩擦のないものになります — 自発的なセッションや企画された集まりに最適です。`,
			"ko-KO": `@ToGetheredBot에 오신 것을 환영합니다 — 친구들과의 만남, 그룹 활동 조직, 또는 단순히 다른 사람들에게 당신이 어디에 있을지 알려주는 간단하고 스마트한 계획 도우미입니다. 해변에서의 카이트서핑, 길거리 농구, 또는 캐주얼한 만남 계획 등 무엇이든, ToGethered는 조정을 쉽게 만듭니다.

봇은 두 가지 주요 기능을 제공합니다:

	1.	<b>계획 공유</b> – 다른 사람들에게 <b>언제 어디에 있을 계획인지</b> 알려주세요. 위치와 시간 범위를 지정할 수 있어서 다른 사람들이 당신의 계획을 보고 그들도 가능하다면 참여할 수 있습니다.

	2.	<b>이벤트 조정</b> – 여러 시간과 장소 옵션을 제안하여 그룹 활동을 쉽게 조직하세요. 봇은 참가자들의 투표를 수집하고 어떤 조합이 가장 잘 작동하는지 보여주어, 긴 채팅 스레드 없이 그룹이 계획에 동의할 수 있도록 도와줍니다.

@ToGetheredBot과 함께라면 계획이 사회적이고, 가시적이며, 마찰 없이 이루어집니다 — 자발적인 세션이나 조직된 모임에 완벽합니다.`,
			"pl-PL": `Witaj w @ToGetheredBot — Twoim prostym i inteligentnym asystentem planowania spotkań z przyjaciółmi, organizowania działań grupowych lub po prostu informowania innych, gdzie będziesz. Czy to kitesurfing na plaży, granie w koszykówkę uliczną, czy planowanie nieformalnego spotkania, ToGethered sprawia, że koordynacja staje się bezwysiłkowa.

Bot oferuje dwie główne funkcje:

	1.	<b>Udostępnianie Planów</b> – Poinformuj innych <b>gdzie i kiedy planujesz gdzieś być</b>. Możesz określić lokalizację i zakres czasowy, aby inni mogli zobaczyć Twoje plany i dołączyć, jeśli są również dostępni.

	2.	<b>Koordynacja Wydarzeń</b> – Łatwo organizuj działania grupowe, proponując wiele opcji czasu i miejsca. Bot zbiera głosy od uczestników i pokazuje, które kombinacje działają najlepiej, pomagając grupie uzgodnić plan bez długich wątków czatu.

Z @ToGetheredBot planowanie staje się społeczne, widoczne i bezproblemowe — idealne dla spontanicznych sesji lub zorganizowanych spotkań.`,
			"pt-BR": `Bem-vindo ao @ToGetheredBot — seu assistente de planejamento simples e inteligente para encontrar amigos, organizar atividades em grupo, ou simplesmente avisar outros onde você estará. Seja kitesurf na praia, jogar basquete de rua, ou planejar um encontro casual, ToGethered torna a coordenação sem esforço.

O bot oferece duas funcionalidades principais:

	1.	<b>Compartilhamento de Planos</b> – Deixe outros saberem <b>onde e quando você planeja estar em algum lugar</b>. Você pode especificar uma localização e um intervalo de tempo, para que outros possam ver seus planos e participar se também estiverem disponíveis.

	2.	<b>Coordenação de Eventos</b> – Organize facilmente atividades em grupo sugerindo múltiplas opções de tempo e local. O bot coleta votos dos participantes e mostra quais combinações funcionam melhor, ajudando o grupo a concordar com um plano sem longas discussões no chat.

Com @ToGetheredBot, o planejamento se torna social, visível e sem atritos — perfeito para sessões espontâneas ou encontros organizados.`,
			"ru-RU": `Добро пожаловать в @ToGetheredBot — ваш простой и умный помощник для планирования встреч с друзьями, организации групповых мероприятий или просто для того, чтобы сообщить другим, где вы будете. Будь то кайтсёрфинг на пляже, игра в уличный баскетбол или планирование неформальной встречи, ToGethered делает координацию лёгкой.

Бот предлагает две основные функции:

	1.	<b>Обмен планами</b> – Дайте другим знать, <b>где и когда вы планируете где-то быть</b>. Вы можете указать место и временной диапазон, чтобы другие могли видеть ваши планы и присоединиться, если они тоже свободны.

	2.	<b>Координация событий</b> – Легко организуйте групповые мероприятия, предлагая несколько вариантов времени и места. Бот собирает голоса участников и показывает, какие комбинации работают лучше всего, помогая группе договориться о плане без длинных цепочек сообщений.

С @ToGetheredBot планирование становится социальным, видимым и беспрепятственным — идеально подходит для спонтанных сессий или организованных встреч.`,
			"tr-TR": `@ToGetheredBot'a hoş geldiniz — arkadaşlarınızla buluşmak, grup etkinlikleri düzenlemek veya sadece başkalarına nerede olacağınızı bildirmek için basit ve akıllı planlama asistanınız. İster sahilde rüzgar sörfü, ister sokak basketbolu oynamak, ister gündelik bir buluşma planlamak olsun, ToGethered koordinasyonu zahmetsiz hale getirir.

Bot iki ana özellik sunar:

	1.	<b>Plan Paylaşımı</b> – Başkalarına <b>nerede ve ne zaman bir yerde olmayı planladığınızı</b> bildirin. Bir konum ve zaman aralığı belirtebilirsiniz, böylece diğerleri planlarınızı görebilir ve onlar da müsaitse katılabilir.

	2.	<b>Etkinlik Koordinasyonu</b> – Birden fazla zaman ve yer seçeneği önererek grup etkinliklerini kolayca düzenleyin. Bot katılımcılardan oyları toplar ve hangi kombinasyonların en iyi çalıştığını gösterir, grubun uzun sohbet zincirleri olmadan bir plan üzerinde anlaşmasına yardımcı olur.

@ToGetheredBot ile planlama sosyal, görünür ve sürtünmesiz hale gelir — spontane oturumlar veya organize toplantılar için mükemmel.`,
			"ua-UA": `Ласкаво просимо до @ToGetheredBot — вашого простого та розумного помічника для планування зустрічей з друзями, організації групових заходів або просто повідомлення іншим, де ви будете. Чи то кайтсерфінг на пляжі, гра в вуличний баскетбол, чи планування неформальної зустрічі, ToGethered робить координацію легкою.

Бот пропонує дві основні функції:

	1.	<b>Обмін планами</b> – Дайте іншим знати, <b>де і коли ви плануєте десь бути</b>. Ви можете вказати місце та часовий діапазон, щоб інші могли бачити ваші плани та приєднатися, якщо вони теж вільні.

	2.	<b>Координація подій</b> – Легко організовуйте групові заходи, пропонуючи декілька варіантів часу та місця. Бот збирає голоси учасників і показує, які комбінації працюють найкраще, допомагаючи групі домовитися про план без довгих ланцюжків повідомлень.

З @ToGetheredBot планування стає соціальним, видимим та безперешкодним — ідеально підходить для спонтанних сесій або організованих зустрічей.`,
			"uz-UZ": `@ToGetheredBot'ga xush kelibsiz — do'stlaringiz bilan uchrashish, guruh faoliyatlarini tashkil qilish yoki boshqalarga qayerda bo'lishingizni bildirish uchun oddiy va aqlli rejalashtirish yordamchingiz. Plyajda kite surfing, ko'cha basketboli o'ynash yoki oddiy uchrashuv rejalashtirish bo'ladimi, ToGethered muvofiqlashtirish jarayonini osonlashtiradi.

Bot ikkita asosiy xususiyatni taklif qiladi:

	1.	<b>Reja Almashish</b> – Boshqalarga <b>qayerda va qachon biror joyda bo'lishni rejalashtirayotganingizni</b> bildiring. Joylashuv va vaqt oralig'ini belgilashingiz mumkin, shunda boshqalar rejalaringizni ko'rib, agar ular ham bo'sh bo'lsa qo'shilishlari mumkin.

	2.	<b>Tadbir Muvofiqlashtirish</b> – Bir nechta vaqt va joy variantlarini taklif qilib guruh faoliyatlarini osonlikcha tashkil qiling. Bot ishtirokchilardan ovozlarni to'playdi va qaysi kombinatsiyalar eng yaxshi ishlashini ko'rsatadi, guruhga uzun chat zanjirlarisiz reja bo'yicha kelishishga yordam beradi.

@ToGetheredBot bilan rejalashtirish ijtimoiy, ko'rinadigan va muammosiz bo'ladi — o'z-o'zidan yuzaga keladigan sessiyalar yoki tashkil etilgan uchrashuvlar uchun mukammal.`,
			"zh-CN": `欢迎使用@ToGetheredBot — 您简单智能的规划助手，用于与朋友见面、组织团体活动，或只是让其他人知道您将在哪里。无论是在海滩风筝冲浪、打街头篮球，还是规划休闲聚会，ToGethered让协调变得轻松。

机器人提供两个主要功能：

	1.	<b>计划分享</b> – 让其他人知道<b>您计划在何时何地出现</b>。您可以指定地点和时间范围，这样其他人就能看到您的计划，如果他们也有空就可以加入。

	2.	<b>活动协调</b> – 通过建议多个时间和地点选项轻松组织团体活动。机器人收集参与者的投票并显示哪些组合效果最好，帮助团队在不需要长聊天记录的情况下达成计划共识。

使用@ToGetheredBot，规划变得社交化、可视化且无摩擦 — 非常适合自发的聚会或有组织的集会。`,
		},
		TogdMainMenuText: {
			"de-DE": `
Um Ihre Pläne zu teilen, wählen Sie einen Ort oder eine Aktivität und klicken Sie auf die <b>{RSVP}</b>-Schaltfläche.

Um ein Event zu organisieren, erstellen Sie ein <b>{NEW_EVENT}</b> unter "Meine Events".
`,
			"en-UK": `
To share your plans choose a spot or activity and hit the <b>{RSVP}</b> button.

To organize an event create a <b>{NEW_EVENT}</b> from "My Events".
`,
			"en-US": `
To share your plans choose a spot or activity and hit the <b>{RSVP}</b> button.

To organize an event create a <b>{NEW_EVENT}</b> from "My Events".
`,
			"es-ES": `
Para compartir tus planes elige un lugar o actividad y pulsa el botón <b>{RSVP}</b>.

Para organizar un evento crea un <b>{NEW_EVENT}</b> desde "Mis eventos".
`,
			"fa-IR": `
برای به اشتراک گذاری برنامه\u200cهایتان، مکان یا فعالیتی را انتخاب کنید و دکمه <b>{RSVP}</b> را فشار دهید.

برای سازماندهی یک رویداد، یک <b>{NEW_EVENT}</b> از "رویدادهای من" ایجاد کنید.
`,
			"fr-FR": `
Pour partager vos plans, choisissez un lieu ou une activité et cliquez sur le bouton <b>{RSVP}</b>.

Pour organiser un événement, créez un <b>{NEW_EVENT}</b> depuis "Mes événements".
`,
			"id-ID": `
Untuk membagikan rencana Anda, pilih tempat atau aktivitas dan tekan tombol <b>{RSVP}</b>.

Untuk mengorganisir acara, buat <b>{NEW_EVENT}</b> dari "Acara Saya".
`,
			"it-IT": `
Per condividere i tuoi piani scegli un posto o un'attività e premi il pulsante <b>{RSVP}</b>.

Per organizzare un evento crea un <b>{NEW_EVENT}</b> da "I miei eventi".
`,
			"ja-JP": `
プランを共有するには、スポットまたはアクティビティを選択して<b>{RSVP}</b>ボタンを押してください。

イベントを企画するには、「マイイベント」から<b>{NEW_EVENT}</b>を作成してください。
`,
			"ko-KO": `
계획을 공유하려면 장소나 활동을 선택하고 <b>{RSVP}</b> 버튼을 누르세요.

이벤트를 기획하려면 "내 이벤트"에서 <b>{NEW_EVENT}</b>를 만드세요.
`,
			"pl-PL": `
Aby udostępnić swoje plany, wybierz miejsce lub aktywność i kliknij przycisk <b>{RSVP}</b>.

Aby zorganizować wydarzenie, utwórz <b>{NEW_EVENT}</b> z "Moje wydarzenia".
`,
			"pt-BR": `
Para compartilhar seus planos, escolha um local ou atividade e clique no botão <b>{RSVP}</b>.

Para organizar um evento, crie um <b>{NEW_EVENT}</b> em "Meus eventos".
`,
			"pt-PT": `
Para partilhar os seus planos, escolha um local ou atividade e clique no botão <b>{RSVP}</b>.

Para organizar um evento, crie um <b>{NEW_EVENT}</b> em "Os meus eventos".
`,
			"ru-RU": `
Чтобы поделиться своими планами, выберите место или активность и нажмите кнопку <b>{RSVP}</b>.

Чтобы организовать событие, создайте <b>{NEW_EVENT}</b> в разделе "Мои события".
`,
			"tr-TR": `
Planlarınızı paylaşmak için bir yer veya aktivite seçin ve <b>{RSVP}</b> düğmesine tıklayın.

Bir etkinlik düzenlemek için "Etkinliklerim"den <b>{NEW_EVENT}</b> oluşturun.
`,
			"ua-UA": `
Щоб поділитися своїми планами, оберіть місце або активність і натисніть кнопку <b>{RSVP}</b>.

Щоб організувати подію, створіть <b>{NEW_EVENT}</b> з "Мої події".
`,
			"uz-UZ": `
Rejalaringizni baham ko'rish uchun joy yoki faoliyatni tanlang va <b>{RSVP}</b> tugmasini bosing.

Tadbir tashkil qilish uchun "Mening tadbirlarim"dan <b>{NEW_EVENT}</b> yarating.
`,
			"zh-CN": `
要分享您的计划，请选择一个地点或活动，然后点击<b>{RSVP}</b>按钮。

要组织活动，请从"我的活动"创建<b>{NEW_EVENT}</b>。
`,
		},
		RsvpButtonText: {
			"de-DE": "Anmelden",
			"en-UK": "RSVP",
			"en-US": "RSVP",
			"es-ES": "Confirmar",
			"fa-IR": "ثبت نام",
			"fr-FR": "S'inscrire",
			"id-ID": "Daftar",
			"it-IT": "Iscriviti",
			"ja-JP": "参加登録",
			"ko-KO": "등록하기",
			"pl-PL": "Zapisz się",
			"pt-BR": "Inscrever-se",
			"pt-PT": "Inscrever-se",
			"ru-RU": "Отметиться",
			"tr-TR": "Kayıt ol",
			"ua-UA": "Зареєструватися",
			"uz-UZ": "Ro'yxatdan o'tish",
			"zh-CN": "报名",
		},
		NewEventButtonText: {
			"de-DE": "Neues Event",
			"en-UK": "New Event",
			"en-US": "New Event",
			"es-ES": "Nuevo evento",
			"fa-IR": "رویداد جدید",
			"fr-FR": "Nouvel événement",
			"id-ID": "Acara baru",
			"it-IT": "Nuovo evento",
			"ja-JP": "新しいイベント",
			"ko-KO": "새 이벤트",
			"pl-PL": "Nowe wydarzenie",
			"pt-BR": "Novo evento",
			"ru-RU": "Новое событие",
			"tr-TR": "Yeni etkinlik",
			"ua-UA": "Нова подія",
			"uz-UZ": "Yangi tadbir",
			"zh-CN": "新活动",
		},
	}
	for k, v := range trans {
		TRANS[k] = v
	}
}
