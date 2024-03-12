package main

import "sync"

func setLanguageStrings(LANGUAGE LANG) {
	if LANGUAGE == EN {
		aboutPageTitle = aboutPageEN
		setupPageTitle = setupPageEN
		calendarPageTitle = calendarPageEN
		exportPageTitle = exportPageEN
		settingsPageTitle = settingsPageEN
		thankYouText = thankYouEN
		notImplementedYet = notImplementedYetEN
	}
	if LANGUAGE == GER {
		aboutPageTitle = aboutPageGER
		setupPageTitle = setupPageGER
		calendarPageTitle = calendarPageGER
		exportPageTitle = exportPageGER
		settingsPageTitle = settingsPageGER
		thankYouText = thankYouGER
		notImplementedYet = notImplementedYetGER
	}
}

// boilerplate ebiten function: init stuff
func init() {
	KBDC = newKBDCursor()

	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		setLanguageStrings(LANGUAGE)
		wg.Done()
	}()
	go func() {
		Logos = logoLoader()
		wg.Done()
	}()
	go func() {
		thanksImage = imageLoader(thanks)
		wg.Done()
	}()
	go func() {
		mpRegular = fontLoader(30, "Regular")
		//mpBlack = fontLoader(30, "Black")
		mpBold = fontLoader(30, "Bold")
		mpExtraBold = fontLoader(70, "ExtraBold")
		// mpExtraLight = fontLoader(30, "ExtraLight")
		// mpLight = fontLoader(30, "Light")
		// mpMedium = fontLoader(30, "Medium")
		// mpSemiBold = fontLoader(30, "SemiBold")
		// mpThin = fontLoader(30, "Thin")
		wg.Done()
	}()
	wg.Wait()
	WriteDB("test/dev.DB", "devbucket", "devkey", []byte("devdata"))
}
