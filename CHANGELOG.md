## NGSI Go v0.8.4-next

-   Improve: Update e2e test (#189)
-   Update: Update documentatin (#188)
-   Hardening: Add geoproxy command (#187)
-   Hardening: Add tokenproxy command (#186)
-   Hardening: Add singleLine option to broker and server cmd (#185)
-   Hardening: Suuport APIKEY (#184)
-   Improve: Refactor regproxy and update regproxy example (#183)
-   Improve: Add server, health and config sub-cmd to regproxy cmd (#182)
-   Improve: Refactor token manager (#181)
-   Improve: Set default FIWARE Service Path to "/" in cp command (#180)
-   Improve: Set default FIWARE Service Path to "/" in rm command (#179)
-   Hardening: Support Kong (#178)
-   Hardening: Add insecureSkipVerify option (#177)
-   Fix: Fix URL path join (#176)
-   Fix: Fix EOF error when ngsi-go-config.json is empty (#175)
-   Hardening: Add replace option in regproxy (#174)
-   Hardening: Support WSO2 (#173)
-   Documentation: Update documentation (#172)
-   Documentation: Add example for Keyrock (#171)
-   Documentation: Add example for Telefonica security stack (#170)
-   Hardening: Support Keycloak (#169)
-   Hardening: Add revoke token option (#168)
-   Fix: Fix error information (#167)
-   Use refresh token to renew token for Keyrock (#166)
-   Update: Refactoring token manager (#165)
-   Hardening: Add PreviousArgs sub-cmd in settings (#164)
-   Documentation: Add example for regproxy (#163)
-   Hardening: Support basic authentication (#162)
-   Update: Refactoring token manager (#161)
-   Fix: Print stack messages correctly (#160)
-   Hardening: Add WireCloud command (#159)

## NGSI Go v0.8.4 - June 20, 2021

-   Hardening: Add registration proxy server (#156)
-   Fix: Fix not to stay messages in buffer for Stdout at receiver command (#155)
-   Fix: Fix cache parameter (#154)

## NGSI Go v0.8.3 - June 12, 2021

-   Fix: Set Accept Header to application/json in batch request (#151)
-   Hardening: Add skipForwarding option for Orion (#150)
-   Fix: Fix message in cp command (#149)
-   Update: Update test process (#148)
-   Update: Update Orion version to 3.1.0 (#146)
-   Hardening: Adds redirectSignOutUri option in applications command (#145)
-   Update: applications-pep-proxies.md (#142)

## NGSI Go v0.8.2 (FIWARE Release 8.0.0) - June 6, 2021

-   Hardening: feature to copy NGSIv2 entities as NGSI-LD entities (#138)
-   Improve quality of documents (#136)
-   Fix typo in flag message (#135)
-   Fix misuse of unbuffered os.Signal channel in testing tools (#134)

## NGSI Go v0.8.1 - April 23, 2021

-   Hardening: Show obfuscated password in broker/server command (#131)
-   Update: ADD tutorial for keystone IDM (#123)
-   Hardening: Support Thinking Cities authentication API (#126)
-   Update: documentation about APIs mapping (#125)
-   Update: Add help message about --brokerType option (#124)
-   Update: Update Orion version to 3.0.0 (#121)
-   Fix: Fix 'sevrice' to 'service' typo (#120)
-   Update: Orion version to 2.6.1 (#119)
-   Hardening: NGSI-LD batch_create (#118)
-   Fix: documentation typo (#116)

## NGSI Go v0.8.0 (FIWARE_7.9.2) - Mar 15, 2021

-   Update: Replace --attrName with --attr (#113)
-   Hardening: Add status check and supports stdout flush to rm command (#112)
-   Hardening: Add NGSI-LD and NGSI v1 copy mode to cp command (#111)
-   Fix: Fix documentation typos (#110)
-   Hardening: Add ngsi v1 mode to rm command (#109)
-   Fix: status code of batch operation in NGSI-LD (#108)
-   Update: Update Orion version to 2.6.0 (#107)
-   Hardening: Add temporal command (#106)
-   Update: documentation about the specification to map FIWARE Open APIs with NGSI Go commands (#105)
-   Hardening: Add --acceptGeoJson option into list entities and get entity command (#104)
-   Fix: typos in some files and formats some source code files by gofmt (#103)
-   Hardening: Update list entities command (#102)
-   Update: Remove tentative fix in types command (#101) 

## NGSI Go v0.7.0 - Feb 28, 2021

-   Hardening: Update list entities command (#98)
-   Update: Add @context (#97)
-   Fix: unmarshal error in types command for Orion-LD (#96)
-   Hardening: Add a period as the allowed character in a server/broker name (#95)
-   Hardening: add feature to remove multiple types in rm command (#94)
-   Hardening: Add buffering feature (#93)
-   Hardening: Add scorpio command (#92)
-   Hardening: Add Cygnus command (#91)
-   Update: Update Orion-LD version to 0.6.1 (#90)
-   Hardening: Adds Keyrock command (#89)
-   Update: Replace attrs=id with attrs=__NONE (#88)
-   Update: Add sleep in IoT Agent test to avoid e2e test failing (#86)
-   Fix: nil pointer dereference in template subscription (#85)
-   Update: Update golang 1.16.0 (#84)
-   Update: Prepare to support WireCloud command (#82)
-   Update: Prepare to support Cygnus command (#81)
-   Update: Prepare to support Keyrock command (#80)
-   Update: Add description about Perseo in documentation (#79)
-   Hardening: Add perseo command (#78)
-   Fix: typos in messages (#77)
-   Hardening: Adds NGSILD-Tenant feature (#76)
-   Hardening: Add IoT Agent Provision command (#75)
-   Fix: HTML tag in documentation (#74)
-   Update: config, link and documentation about QuantumLeap (#73)
-   Fix: fixes documentation (#72)

## NGSI Go v0.6.0 - Feb 3, 2021

-   Update: Add e2e test cases (#69)
-   Hardening: Add time series command (#68)
-   Update: Update ngsi-test command and Dockerfile (#66)
-   Hardening: Update receiver command (#65)
-   Hardening: Clear previous args when host not registered (#64)
-   Fix: Fix update attribute feature (#63)
-   Hardening: Clear previous args when host is URL (#62)
-   Hardening: Add e2e test cases (#61)
-   Hardening: Update CI process (#60)
-   Update: Refactoring NGSI Go and e2e test (#59)
-   Update: Update Markdown files (#58)
-   Update: Improve e2e test (#57)
-   Update: Build ngsi-go with CGO_ENABLED=0 (#56)
-   Hardening: Updates broker command (#55)
-   Update: Update Dockerfile for test tool (#54)
-   Update: Improve e2e test (#53)
-   Update: Move 3000_management into cases (#52)
-   Hardening: Convert tenant name to lowercase (#51)
-   Hardening: Update safe string feature (#50)
-   Update: Add feature to remove intermediate images for test tool (#49)
-   Update: Update docker-compose.yml for e2e test (#48)
-   Update: Update dockerfile for test tool (#47)
-   Hardening: Update e2e test (#46)
-   Update: Add note about FIWARE Service (#42)
-   Update: Add missing backslashes to example commands (#41)
-   Fix: error that null is displayed when number of registrations is 0 (#39)

## NGSI Go v0.5.0 - 2 Jan, 2021

-   Update: Update copyright (#38)
-   Update: Refactoring unit test cases (#37)
-   Hardening: Add --pretty option to 'list types' and 'get type' command (#36)
-   Hardening: Add list types command for NGSI-LD (#35)
-   Update: Refactoring unit test case (#34)
-   Fix: Fix config error (#33)
-   Hardening: Add --context option (#32)
-   Fix: Fix missing line feed (#31)
-   Hardening: Add feature to remove NGSI-LD entities (#30)
-   Hardening: Enable verbose mode when some options are specified (#29)
-   Update: Remove unused options in list types command (#28)
-   Fix: Fix 400 Bad Request  in rm command. (#27)
-   Update: Update Orion version to 2.5.2 (#26)
-   Hardening: Add pretty option (#25)
-   Hardening: Update safeString feature (#24)

## NGSI Go v0.4.0 - 16 Dec, 2020

-   Hardening: Check broker type in entity command. (#23)
-   Hardening: Update subscriptionV2 (#22)
-   Hardening: Update notification receiver (#21)
-   Update: Update build process (#20)
-   Hardening: Update notification receiver (#19)
-   Hardening: Add notification receiver (#18)
-   Hardening: Add registrationLD command (#17)
-   Hardening: Update subscriptionLD command (#16)
-   Hardening: Add subscriptionLD command (#15)
-   hardening: Add upsert entity command (#14)


## NGSI Go v0.3.0 (FIWARE_7.9.1)- 6 Dec 2020

-   Update: Update documentation (#13)
-   Hardening: Add admin command (#12)
-   Hardening: Output version info with logging level higher than Info (#11)
-   Fix: Fix getting cached token (#10)
-   Hardening: Add count option to subscriptions list command (#9)
-   Hardening: Add E2E test (#8)
-   Fix: Fix typos in funcName (#7)

## NGSI Go v0.2.0 - 23 Nov, 2020

-   Update: Improve code quality (#6)
-   Hardening: Improve registration command (#5)
-   Fix: Fix scope (FIWARE ServicePath) error (#4)
-   Fix: Fix some small bugs (#3)
-   Fix: Update documentation (#2)
-   Fix: Update README.md (#1)

## NGSI Go v0.1.0 - 15 Nov, 2020

-   Initial release
