# MQTT farmer aid

L'idea base del progetto è quella di realizzare una libreria che permetta di interagire con diverse tipologie di sensori e attuatori in go in modo uniforme. Il target principale del progetto dovrebbero essere dei sensori e attuatori simili a quelli che si utilizzano nei campi coltivati o nelle serre.

Saranno inserite delle interfacce che permettano di interagire coi sensori e con gli attuatori in vario modo oltre che di accedere alle informazioni che espongono e scegliere quali informazioni devono esporre.
I diversi sensori e attuatori disporranno della possibilità di comunicare fra di loro grazie all'iscrizione a specifici topic.

Periodicamente o al verificarsi di specifiche condizioni, i sensori leggono informazioni e le pubblicano. I sensori possono pubblicare le informazioni usando diversi formati, fra quelli supportati sarà presente il JSON.

Per gestire il meccanismo publish/subscribe sarà usato il protocollo MQTT. Il motivo della scelta di questo protocollo risiede nel fatto che è adatto ad una situazione in cui si hanno dispositivi poco potenti come quelli in esame.

Il broker consigliato sarà mosquitto (oppure rabbitmq se si dovesse rivelare più adatto) e i dati da lui ricevuti saranno anche salvati all'interno di un database così da permettere agli utenti di poter visualizzare informazioni sull'ambiente di loro interesse.

Sarà integrata una web-app che permetta di eseguire alcune operazioni quali:

- vedere le informazioni dei sensori in modo user-friendly
- attivare gli attuatori
- spengere/accendere specifici sensori e attuatori.

A questo riguardo sarà integrato anche un sistema di autenticazione e autorizzazione per l'accesso ai dati.
