package trans

const (
	TogdBotWelcome = "TogdBotWelcome"

	TogdMyProfile    = "TogdMyProfile"
	TogdMyActivities = "TogdMyActivities"
	TogdMyEvents     = "TogdMyEvents"
	TogdMyPlans      = "TogdMyPlans"
	TogdMySpots      = "TogdMySpots"

	TogdUserProfile    = "TogdUserProfile"
	TogdUserActivities = "TogdUserActivities"
	TogdUserEvents     = "TogdUserEvents"
	TogdUserPlans      = "TogdUserPlans"
	TogdUserSpots      = "TogdUserSpots"

	TogdMainMenuText = "TogdMainMenuText"

	RsvpButtonText     = "RsvpButtonText"
	NewEventButtonText = "NewEventButtonText"
)

func init() {
	/*
		Proper order of locale keys in var trans, use it as a reference for all values:
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
			"ru-RU": "Профиль пользователя",
			"tr-TR": "Kullanıcı profili",
			"ua-UA": "Профіль користувача",
			"uz-UZ": "Foydalanuvchi profili",
			"zh-CN": "用户资料",
		},

		TogdUserActivities: {
			"de-DE": "Benutzeraktivitäten",
			"en-UK": "User activities",
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
