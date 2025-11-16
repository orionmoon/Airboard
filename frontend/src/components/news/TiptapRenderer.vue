<template>
  <div class="tiptap-content" v-html="renderedContent"></div>
</template>

<script setup>
import { computed } from 'vue'
import { generateHTML } from '@tiptap/html'
import StarterKit from '@tiptap/starter-kit'
import Link from '@tiptap/extension-link'
import { Table } from '@tiptap/extension-table'
import { TableRow } from '@tiptap/extension-table-row'
import { TableCell } from '@tiptap/extension-table-cell'
import { TableHeader } from '@tiptap/extension-table-header'
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight'
import TextAlign from '@tiptap/extension-text-align'
import { createLowlight, common } from 'lowlight'

// Create lowlight instance with common languages
const lowlight = createLowlight(common)

const props = defineProps({
  content: {
    type: [Object, String],
    required: true
  }
})

const renderedContent = computed(() => {
  if (!props.content) return ''

  // If content is string, assume it's plain text
  if (typeof props.content === 'string') {
    return `<p>${props.content}</p>`
  }

  // Generate HTML from Tiptap JSON
  return generateHTML(props.content, [
    StarterKit.configure({
      codeBlock: false
    }),
    Link.configure({
      openOnClick: false
    }),
    Table,
    TableRow,
    TableCell,
    TableHeader,
    CodeBlockLowlight.configure({
      lowlight
    }),
    TextAlign.configure({
      types: ['heading', 'paragraph']
    })
  ])
})
</script>

<style>
.tiptap-content {
  @apply text-gray-800 dark:text-gray-200;
}

.tiptap-content h1 {
  @apply text-3xl font-bold mb-4 text-gray-900 dark:text-white;
}

.tiptap-content h2 {
  @apply text-2xl font-bold mb-3 text-gray-900 dark:text-white;
}

.tiptap-content h3 {
  @apply text-xl font-bold mb-2 text-gray-900 dark:text-white;
}

.tiptap-content p {
  @apply mb-4 leading-relaxed;
}

.tiptap-content ul {
  @apply list-disc list-inside mb-4 space-y-2;
}

.tiptap-content ol {
  @apply list-decimal list-inside mb-4 space-y-2;
}

.tiptap-content a {
  @apply text-primary-500 hover:text-primary-600 underline;
}

.tiptap-content blockquote {
  @apply border-l-4 border-gray-300 dark:border-gray-600 pl-4 italic my-4 text-gray-700 dark:text-gray-300;
}

.tiptap-content code {
  @apply bg-gray-100 dark:bg-gray-800 px-2 py-1 rounded text-sm font-mono text-red-600 dark:text-red-400;
}

.tiptap-content pre {
  @apply bg-gray-900 dark:bg-gray-950 text-gray-100 p-4 rounded-lg overflow-x-auto mb-4;
}

.tiptap-content pre code {
  @apply bg-transparent text-gray-100 p-0;
}

.tiptap-content table {
  @apply w-full border-collapse mb-4;
}

.tiptap-content th {
  @apply bg-gray-100 dark:bg-gray-800 border border-gray-300 dark:border-gray-600 px-4 py-2 text-left font-semibold;
}

.tiptap-content td {
  @apply border border-gray-300 dark:border-gray-600 px-4 py-2;
}

.tiptap-content hr {
  @apply border-gray-300 dark:border-gray-600 my-6;
}

.tiptap-content strong {
  @apply font-bold;
}

.tiptap-content em {
  @apply italic;
}

.tiptap-content img {
  @apply max-w-full h-auto rounded-lg my-4;
}
</style>
