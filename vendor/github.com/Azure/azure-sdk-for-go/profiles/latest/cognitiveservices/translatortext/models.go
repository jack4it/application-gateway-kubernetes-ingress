// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package translatortext

import original "github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v3.0/translatortext"

type BaseClient = original.BaseClient
type BreakSentenceResultItem = original.BreakSentenceResultItem
type BreakSentenceTextInput = original.BreakSentenceTextInput
type DetectResultItem = original.DetectResultItem
type DetectTextInput = original.DetectTextInput
type DictionaryExampleResultItem = original.DictionaryExampleResultItem
type DictionaryExampleResultItemExamplesItem = original.DictionaryExampleResultItemExamplesItem
type DictionaryExampleTextInput = original.DictionaryExampleTextInput
type DictionaryLookupResultItem = original.DictionaryLookupResultItem
type DictionaryLookupResultItemTranslationsItem = original.DictionaryLookupResultItemTranslationsItem
type DictionaryLookupResultItemTranslationsItemBackTranslationsItem = original.DictionaryLookupResultItemTranslationsItemBackTranslationsItem
type DictionaryLookupTextInput = original.DictionaryLookupTextInput
type ErrorMessage = original.ErrorMessage
type ErrorMessageError = original.ErrorMessageError
type LanguagesResult = original.LanguagesResult
type LanguagesResultDictionary = original.LanguagesResultDictionary
type LanguagesResultDictionaryLanguageCode = original.LanguagesResultDictionaryLanguageCode
type LanguagesResultDictionaryLanguageCodeTranslationsItem = original.LanguagesResultDictionaryLanguageCodeTranslationsItem
type LanguagesResultTranslation = original.LanguagesResultTranslation
type LanguagesResultTranslationLanguageCode = original.LanguagesResultTranslationLanguageCode
type LanguagesResultTransliteration = original.LanguagesResultTransliteration
type LanguagesResultTransliterationLanguageCode = original.LanguagesResultTransliterationLanguageCode
type LanguagesResultTransliterationLanguageCodeScriptsItem = original.LanguagesResultTransliterationLanguageCodeScriptsItem
type LanguagesResultTransliterationLanguageCodeScriptsItemToScriptsItem = original.LanguagesResultTransliterationLanguageCodeScriptsItemToScriptsItem
type ListBreakSentenceResultItem = original.ListBreakSentenceResultItem
type ListDetectResultItem = original.ListDetectResultItem
type ListDictionaryExampleResultItem = original.ListDictionaryExampleResultItem
type ListDictionaryLookupResultItem = original.ListDictionaryLookupResultItem
type ListTranslateResultAllItem = original.ListTranslateResultAllItem
type ListTransliterateResultItem = original.ListTransliterateResultItem
type TranslateResultAllItem = original.TranslateResultAllItem
type TranslateResultAllItemDetectedLanguage = original.TranslateResultAllItemDetectedLanguage
type TranslateResultAllItemTranslationsItem = original.TranslateResultAllItemTranslationsItem
type TranslateResultAllItemTranslationsItemAlignment = original.TranslateResultAllItemTranslationsItemAlignment
type TranslateResultAllItemTranslationsItemSentLen = original.TranslateResultAllItemTranslationsItemSentLen
type TranslateResultAllItemTranslationsItemSentLenSrcSentLenItem = original.TranslateResultAllItemTranslationsItemSentLenSrcSentLenItem
type TranslateResultAllItemTranslationsItemSentLenTransSentLenItem = original.TranslateResultAllItemTranslationsItemSentLenTransSentLenItem
type TranslateResultAllItemTranslationsItemTransliteration = original.TranslateResultAllItemTranslationsItemTransliteration
type TranslateResultItem = original.TranslateResultItem
type TranslateResultItemTranslationItem = original.TranslateResultItemTranslationItem
type TranslateTextInput = original.TranslateTextInput
type TranslatorClient = original.TranslatorClient
type TransliterateResultItem = original.TransliterateResultItem
type TransliterateTextInput = original.TransliterateTextInput

func New(endpoint string) BaseClient {
	return original.New(endpoint)
}
func NewTranslatorClient(endpoint string) TranslatorClient {
	return original.NewTranslatorClient(endpoint)
}
func NewWithoutDefaults(endpoint string) BaseClient {
	return original.NewWithoutDefaults(endpoint)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}