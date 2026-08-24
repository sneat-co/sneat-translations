# Sneat Translations

[![Go CI](https://github.com/sneat-co/sneat-translations/actions/workflows/ci.yml/badge.svg)](https://github.com/sneat-co/sneat-translations/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/sneat-co/sneat-translations)](https://goreportcard.com/report/github.com/sneat-co/sneat-translations)

Sneat Translations is an active Go repository of shared internationalization resources for Sneat platform surfaces. It preserves the legacy DebtsTracker vocabulary and translation history while serving current applications.

<!-- dev-approach:v1 -->
## Our approach to development

We build with our own tooling:

- **[SpecScore](https://specscore.md)** — specify requirements as `SpecScore.md` artifacts
- **[SpecStudio](https://specscore.studio)** — author & manage specs across their lifecycle
- **[inGitDB](https://ingitdb.com)** — store structured data in Git where applicable
- **[DALgo](https://dalgo.io)** — data access layer for Go
- **[cover100.dev](https://cover100.dev)** — drive toward 100% test coverage
- **[DataTug](https://datatug.io)** — query & explore data
<!-- /dev-approach -->

## Quick jump to your language
<a href="#english">EN</a> | <a href="#russian">RU</a> | <a href="#italian">IT</a> | <a href="#farsi">FA</a>

## <a id="english">English</a>
At present, the legacy app supports English and Russian. Contributions for additional languages are welcome. If an existing translation needs correction, please [open an issue](https://github.com/sneat-co/sneat-translations/issues/new).

For Telegram bots and the legacy website, all translation strings are kept in a single file: [trans/translations.go](https://github.com/sneat-co/sneat-translations/blob/main/trans/translations.go).

Translation files for web, iOS, and Android applications will be published separately.


## <a id="russian">Русский</a>
В настоящее время приложение переведено на английский и русский языки.
Мы будем благодарны за pull request с переводом на любой другой язык.
Если вы считаете, что существующий перевод неточен, [сообщите нам](https://github.com/sneat-co/sneat-translations/issues/new).

Для Telegram-ботов и [веб-сайта](http://debtstracker.io/) все строки перевода находятся в одном файле: [trans/translations.go](https://github.com/sneat-co/sneat-translations/blob/main/trans/translations.go).

Переводы для web-, iOS- и Android-приложений будут опубликованы отдельно.


## <a id="italian">Italiano</a>
Al momento, la nostra applicazione è tradotta in inglese e russo.
Saremo lieti di ricevere una pull request per qualsiasi altra lingua.
Se pensi che una traduzione esistente non sia corretta, [faccelo sapere](https://github.com/sneat-co/sneat-translations/issues/new).

Per i bot Telegram e il [sito web](http://debtstracker.io/), tutte le stringhe di traduzione sono in un unico file: [trans/translations.go](https://github.com/sneat-co/sneat-translations/blob/main/trans/translations.go).

I file di traduzione per le app web, iOS e Android saranno pubblicati separatamente.


## <a id="farsi">Farsi/Persian - فارسی</a>
درحال حاضر برنامه ما به زبانهای انگلیسی و روسی ترجمه شده است.
سپاسگزار خواهیم بود اگر شما بتوانید برای هر یک از زبانهای دیگر یک پول ریکوئست ثبت نمایید.
اگر اعتقاد دارید ترجمه در هریک از زبانهای موجود صحیح نمی باشد لطفاً [به ما اطلاع دهید](https://github.com/DebtsTracker/translations/issues/new). 

برای رباتهای (*تلگرام*) و [وب سایت](http://debtstracker.io/) تمام رشته های ترجمه در یک فایل موجود می باشند: https://github.com/DebtsTracker/translations/blob/master/trans/translations.go

درخصوص برنامه (*web, iOS, Android*)  فایلهای ترجمه به زودی منتشر خواند شد.
