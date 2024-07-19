<script setup lang="ts">
import { ref } from "vue";
import { Form, Field, ErrorMessage } from "vee-validate";
import * as api from "../api";
import { onMounted } from "vue";
import type { Url } from "@/api";

const urls = ref<Url[]>([]);
const longUrl = defineModel<string>("longUrl", {
  default: "",
});

onMounted(async () => {
  try {
    const response = await api.listUrls();
    urls.value = response;
  } catch (e: any) {
    // TODO put this in an error alert model field
    console.error(e);
  }
});

function save() {
  api
    .createUrl(longUrl.value)
    .then((response) => {
      urls.value.unshift(response);
      longUrl.value = "";
    })
    .catch((error) => {
      // TODO put this in an error alert model field
      console.error(error);
    });
}

function deleteUrl(id: number) {
  api
    .deleteUrl(id)
    .then(() => {
      urls.value = urls.value.filter((url) => url.id !== id);
    })
    .catch((error) => {
      // TODO put this in an error alert model field
      console.error(error);
    });
}

function validateLongUrl(value: any) {
  let givenURL;
  try {
    givenURL = new URL(value);
  } catch (error) {
    return false;
  }
  return givenURL.protocol === "http:" || givenURL.protocol === "https:";
}
</script>

<template>
  <div>
    <div class="bg-white py-16 sm:py-24">
      <div class="relative sm:py-16">
        <div aria-hidden="true" class="hidden sm:block">
          <div class="inset-y-0 left-0 w-1/2 bg-gray-50 rounded-r-3xl" />
          <svg
            class="absolute top-8 left-1/2 -ml-3"
            width="404"
            height="392"
            fill="none"
            viewBox="0 0 404 392"
          >
            <defs>
              <pattern
                id="8228f071-bcee-4ec8-905a-2a059a2cc4fb"
                x="0"
                y="0"
                width="20"
                height="20"
                patternUnits="userSpaceOnUse"
              >
                <rect
                  x="0"
                  y="0"
                  width="4"
                  height="4"
                  class="text-gray-200"
                  fill="currentColor"
                />
              </pattern>
            </defs>
            <rect
              width="404"
              height="392"
              fill="url(#8228f071-bcee-4ec8-905a-2a059a2cc4fb)"
            />
          </svg>
        </div>
        <div
          class="mx-auto max-w-md px-4 sm:max-w-3xl sm:px-6 lg:max-w-7xl lg:px-8"
        >
          <div
            class="relative rounded-2xl px-6 py-10 bg-indigo-600 overflow-hidden shadow-xl sm:px-12 sm:py-20"
          >
            <div
              aria-hidden="true"
              class="absolute inset-0 -mt-72 sm:-mt-32 md:mt-0"
            >
              <svg
                class="absolute inset-0 h-full w-full"
                preserveAspectRatio="xMidYMid slice"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 1463 360"
              >
                <path
                  class="text-indigo-500 text-opacity-40"
                  fill="currentColor"
                  d="M-82.673 72l1761.849 472.086-134.327 501.315-1761.85-472.086z"
                />
                <path
                  class="text-indigo-700 text-opacity-40"
                  fill="currentColor"
                  d="M-217.088 544.086L1544.761 72l134.327 501.316-1761.849 472.086z"
                />
              </svg>
            </div>
            <div class="relative">
              <div class="sm:text-center">
                <h2
                  class="text-3xl font-extrabold text-white tracking-tight sm:text-4xl"
                >
                  Enter a long URL to make it short
                </h2>
                <p class="mt-6 mx-auto max-w-2xl text-lg text-indigo-200">
                  URLs must start with http:// or https://
                </p>
              </div>

              <Form
                @submit="save"
                id="url-shortener-form"
                class="mt-12 sm:mx-auto sm:max-w-lg"
              >
                <div class="sm:flex mt-3">
                  <div class="min-w-0 flex-1">
                    <div class="mt-2">
                      <Field
                        :rules="validateLongUrl"
                        v-model.trim="longUrl"
                        type="text"
                        name="longUrl"
                        id="longUrl"
                        class="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                        placeholder="Enter a long URL"
                        aria-describedby="long-url"
                        required
                      />
                    </div>

                    <ErrorMessage name="longUrl" />
                  </div>

                  <div class="mt-2 ml-2">
                    <button
                      type="submit"
                      class="rounded-md bg-white/10 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-white/20"
                    >
                      Shorten
                    </button>
                  </div>
                </div>
              </Form>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-4">
        <div class="flex flex-col">
          <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
            <div
              class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8"
            >
              <div
                class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg"
              >
                <table class="min-w-full divide-y divide-gray-200">
                  <thead class="bg-gray-50">
                    <tr>
                      <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                      >
                        URL
                      </th>
                      <th
                        scope="col"
                        class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                      >
                        To
                      </th>

                      <th scope="col" class="relative px-6 py-3">
                        <span class="sr-only">Go</span>
                      </th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr
                      v-for="(url, index) in urls"
                      :class="index % 2 === 0 ? 'bg-white' : 'bg-gray-50'"
                    >
                      <td class="px-6 py-4 text-sm font-medium text-gray-900">
                        {{ url.short_url }}
                      </td>
                      <td class="px-6 py-4 text-sm text-gray-500">
                        {{ url.long_url }}
                      </td>
                      <td class="px-6 py-4 text-right text-sm font-medium">
                        <a
                          :href="url.long_url"
                          class="text-indigo-600 hover:text-indigo-900"
                          target="_blank"
                          rel="nofollow noopener noreferrer"
                        >
                          Go
                        </a>
                      </td>
                      <td>
                        <button
                          @click="deleteUrl(url.id)"
                          class="text-red-600 hover:text-red-900"
                        >
                          Delete
                        </button>
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
