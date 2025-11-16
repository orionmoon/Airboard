<template>
  <div class="rich-text-editor">
    <!-- Toolbar -->
    <div v-if="editor" class="editor-toolbar">
      <!-- Text formatting -->
      <div class="toolbar-group">
        <button
          type="button"
          @click="editor.chain().focus().toggleBold().run()"
          :class="{ 'is-active': editor.isActive('bold') }"
          class="toolbar-button"
          title="Bold (Ctrl+B)"
        >
          <Icon icon="mdi:format-bold" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleItalic().run()"
          :class="{ 'is-active': editor.isActive('italic') }"
          class="toolbar-button"
          title="Italic (Ctrl+I)"
        >
          <Icon icon="mdi:format-italic" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleStrike().run()"
          :class="{ 'is-active': editor.isActive('strike') }"
          class="toolbar-button"
          title="Strikethrough"
        >
          <Icon icon="mdi:format-strikethrough" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleCode().run()"
          :class="{ 'is-active': editor.isActive('code') }"
          class="toolbar-button"
          title="Inline Code"
        >
          <Icon icon="mdi:code-tags" class="h-5 w-5" />
        </button>
      </div>

      <!-- Headings -->
      <div class="toolbar-group">
        <button
          type="button"
          @click="editor.chain().focus().toggleHeading({ level: 1 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 1 }) }"
          class="toolbar-button"
          title="Heading 1"
        >
          <Icon icon="mdi:format-header-1" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleHeading({ level: 2 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 2 }) }"
          class="toolbar-button"
          title="Heading 2"
        >
          <Icon icon="mdi:format-header-2" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleHeading({ level: 3 }).run()"
          :class="{ 'is-active': editor.isActive('heading', { level: 3 }) }"
          class="toolbar-button"
          title="Heading 3"
        >
          <Icon icon="mdi:format-header-3" class="h-5 w-5" />
        </button>
      </div>

      <!-- Lists -->
      <div class="toolbar-group">
        <button
          type="button"
          @click="editor.chain().focus().toggleBulletList().run()"
          :class="{ 'is-active': editor.isActive('bulletList') }"
          class="toolbar-button"
          title="Bullet List"
        >
          <Icon icon="mdi:format-list-bulleted" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleOrderedList().run()"
          :class="{ 'is-active': editor.isActive('orderedList') }"
          class="toolbar-button"
          title="Ordered List"
        >
          <Icon icon="mdi:format-list-numbered" class="h-5 w-5" />
        </button>
      </div>

      <!-- Blocks -->
      <div class="toolbar-group">
        <button
          type="button"
          @click="editor.chain().focus().toggleCodeBlock().run()"
          :class="{ 'is-active': editor.isActive('codeBlock') }"
          class="toolbar-button"
          title="Code Block"
        >
          <Icon icon="mdi:code-braces" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().toggleBlockquote().run()"
          :class="{ 'is-active': editor.isActive('blockquote') }"
          class="toolbar-button"
          title="Blockquote"
        >
          <Icon icon="mdi:format-quote-close" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().setHorizontalRule().run()"
          class="toolbar-button"
          title="Horizontal Rule"
        >
          <Icon icon="mdi:minus" class="h-5 w-5" />
        </button>
      </div>

      <!-- Link -->
      <div class="toolbar-group">
        <button
          type="button"
          @click="setLink"
          :class="{ 'is-active': editor.isActive('link') }"
          class="toolbar-button"
          title="Add Link"
        >
          <Icon icon="mdi:link-variant" class="h-5 w-5" />
        </button>
        <button
          v-if="editor.isActive('link')"
          type="button"
          @click="editor.chain().focus().unsetLink().run()"
          class="toolbar-button"
          title="Remove Link"
        >
          <Icon icon="mdi:link-variant-off" class="h-5 w-5" />
        </button>
      </div>

      <!-- Table (for later when we want to add tables) -->
      <!-- <div class="toolbar-group">
        <button
          type="button"
          @click="editor.chain().focus().insertTable({ rows: 3, cols: 3, withHeaderRow: true }).run()"
          class="toolbar-button"
          title="Insert Table"
        >
          <Icon icon="mdi:table" class="h-5 w-5" />
        </button>
      </div> -->

      <!-- Undo/Redo -->
      <div class="toolbar-group ml-auto">
        <button
          type="button"
          @click="editor.chain().focus().undo().run()"
          :disabled="!editor.can().undo()"
          class="toolbar-button"
          title="Undo (Ctrl+Z)"
        >
          <Icon icon="mdi:undo" class="h-5 w-5" />
        </button>
        <button
          type="button"
          @click="editor.chain().focus().redo().run()"
          :disabled="!editor.can().redo()"
          class="toolbar-button"
          title="Redo (Ctrl+Y)"
        >
          <Icon icon="mdi:redo" class="h-5 w-5" />
        </button>
      </div>
    </div>

    <!-- Editor content -->
    <EditorContent :editor="editor" class="editor-content" />
  </div>
</template>

<script setup>
import { ref, watch, onMounted, onBeforeUnmount } from 'vue'
import { useEditor, EditorContent } from '@tiptap/vue-3'
import StarterKit from '@tiptap/starter-kit'
import Link from '@tiptap/extension-link'
import { Table } from '@tiptap/extension-table'
import { TableRow } from '@tiptap/extension-table-row'
import { TableCell } from '@tiptap/extension-table-cell'
import { TableHeader } from '@tiptap/extension-table-header'
import CodeBlockLowlight from '@tiptap/extension-code-block-lowlight'
import TextAlign from '@tiptap/extension-text-align'
import Placeholder from '@tiptap/extension-placeholder'
import { createLowlight } from 'lowlight'
import { Icon } from '@iconify/vue'

// Import common languages for syntax highlighting
import javascript from 'highlight.js/lib/languages/javascript'
import typescript from 'highlight.js/lib/languages/typescript'
import python from 'highlight.js/lib/languages/python'
import java from 'highlight.js/lib/languages/java'
import bash from 'highlight.js/lib/languages/bash'
import json from 'highlight.js/lib/languages/json'
import xml from 'highlight.js/lib/languages/xml'
import css from 'highlight.js/lib/languages/css'

// Create lowlight instance and register languages
const lowlight = createLowlight()
lowlight.register('javascript', javascript)
lowlight.register('typescript', typescript)
lowlight.register('python', python)
lowlight.register('java', java)
lowlight.register('bash', bash)
lowlight.register('json', json)
lowlight.register('xml', xml)
lowlight.register('css', css)

const props = defineProps({
  modelValue: {
    type: [String, Object],
    default: ''
  },
  placeholder: {
    type: String,
    default: 'Start writing...'
  },
  editable: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue'])

const editor = useEditor({
  content: props.modelValue,
  editable: props.editable,
  extensions: [
    StarterKit.configure({
      codeBlock: false, // Disable default code block, we use lowlight
      link: false, // Disable default link, we configure it separately
    }),
    Link.configure({
      openOnClick: false,
      HTMLAttributes: {
        class: 'text-blue-600 dark:text-blue-400 underline hover:text-blue-800 dark:hover:text-blue-300',
      },
    }),
    Table.configure({
      resizable: true,
    }),
    TableRow,
    TableCell,
    TableHeader,
    CodeBlockLowlight.configure({
      lowlight,
    }),
    TextAlign.configure({
      types: ['heading', 'paragraph'],
    }),
    Placeholder.configure({
      placeholder: props.placeholder,
    }),
  ],
  onUpdate: ({ editor }) => {
    emit('update:modelValue', editor.getJSON())
  },
  editorProps: {
    attributes: {
      class: 'prose prose-sm sm:prose lg:prose-lg xl:prose-xl dark:prose-invert max-w-none focus:outline-none min-h-[300px] p-4',
    },
  },
})

// Set link
const setLink = () => {
  const previousUrl = editor.value.getAttributes('link').href
  const url = window.prompt('URL', previousUrl)

  if (url === null) {
    return
  }

  if (url === '') {
    editor.value.chain().focus().extendMarkRange('link').unsetLink().run()
    return
  }

  editor.value.chain().focus().extendMarkRange('link').setLink({ href: url }).run()
}

// Watch for external content changes
watch(() => props.modelValue, (value) => {
  if (editor.value) {
    const isSame = JSON.stringify(editor.value.getJSON()) === JSON.stringify(value)
    if (!isSame) {
      editor.value.commands.setContent(value, false)
    }
  }
})

// Watch for editable changes
watch(() => props.editable, (value) => {
  if (editor.value) {
    editor.value.setEditable(value)
  }
})

// Cleanup
onBeforeUnmount(() => {
  if (editor.value) {
    editor.value.destroy()
  }
})
</script>

<style scoped>
.rich-text-editor {
  @apply border border-gray-300 dark:border-gray-600 rounded-lg overflow-hidden bg-white dark:bg-gray-800;
}

.editor-toolbar {
  @apply flex flex-wrap items-center gap-1 p-2 bg-gray-50 dark:bg-gray-900 border-b border-gray-300 dark:border-gray-600;
}

.toolbar-group {
  @apply flex items-center gap-1 border-r border-gray-300 dark:border-gray-600 pr-2 mr-2;
}

.toolbar-group:last-child {
  @apply border-r-0 mr-0 pr-0;
}

.toolbar-button {
  @apply p-2 rounded hover:bg-gray-200 dark:hover:bg-gray-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed text-gray-700 dark:text-gray-300;
}

.toolbar-button.is-active {
  @apply bg-blue-100 dark:bg-blue-900 text-blue-600 dark:text-blue-400;
}

.editor-content {
  @apply min-h-[300px];
}

/* Tiptap editor styles */
.editor-content :deep(.ProseMirror) {
  @apply focus:outline-none;
}

.editor-content :deep(.ProseMirror p.is-editor-empty:first-child::before) {
  content: attr(data-placeholder);
  @apply text-gray-400 dark:text-gray-500 float-left h-0 pointer-events-none;
}

/* Code block styling */
.editor-content :deep(.ProseMirror pre) {
  @apply bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto;
}

.editor-content :deep(.ProseMirror code) {
  @apply bg-gray-100 dark:bg-gray-800 px-1 py-0.5 rounded text-sm;
}

/* Link styling */
.editor-content :deep(.ProseMirror a) {
  @apply text-blue-600 dark:text-blue-400 underline hover:text-blue-800 dark:hover:text-blue-300;
}

/* Table styling */
.editor-content :deep(.ProseMirror table) {
  @apply border-collapse table-auto w-full;
}

.editor-content :deep(.ProseMirror th),
.editor-content :deep(.ProseMirror td) {
  @apply border border-gray-300 dark:border-gray-600 px-3 py-2;
}

.editor-content :deep(.ProseMirror th) {
  @apply bg-gray-100 dark:bg-gray-800 font-semibold;
}
</style>
