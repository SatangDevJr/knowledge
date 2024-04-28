import { createI18n } from "vue-i18n";
import en from "./locales/en.json";
import th from "./locales/th.json";

export const SUPPORT_LOCALES = [
  {
    langcode: "th",
    langname: "ไทย",
  },
  {
    langcode: "en",
    langname: "English",
  },
];

export const i18n = createI18n({
  legacy: false,
  globalInjection: true,
  locale: "th",
  fallbackLocale: "en",
  messages: {
    en,
    th,
  },

})

export function useI18n() {
  return i18n.global;
}
