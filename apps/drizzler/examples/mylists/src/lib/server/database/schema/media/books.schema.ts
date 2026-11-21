import {MediaType} from "@/lib/utils/enums";
import {relations} from "drizzle-orm/_relations";
import {user} from "@/lib/server/database/schema/auth.schema";
import {integer, snakeCase, text} from "drizzle-orm/sqlite-core";
import {mediaTags} from "@/lib/server/database/schema/media-tags.schema";
import {
    commonGenericCols,
    commonGenericIndexes,
    commonMediaCols,
    commonMediaListCols,
    commonMediaListIndexes,
    commonMediaTagsCols,
    commonMediaTagsIndexes
} from "@/lib/server/database/schema/media/_helper";


export const books = snakeCase.table("books", {
    pages: integer().notNull(),
    language: text(),
    publishers: text(),
    apiId: text().unique().notNull(),
    ...commonMediaCols(MediaType.BOOKS),
});


export const booksList = snakeCase.table("books_list", {
    actualPage: integer(),
    redo: integer().default(0).notNull(),
    total: integer("total").default(0).notNull(),
    ...commonMediaListCols(books.id, MediaType.BOOKS),
}, (table) => commonMediaListIndexes(table, MediaType.BOOKS));


export const booksGenre = snakeCase.table("books_genre", {
    ...commonGenericCols(books.id),
}, (table) => commonGenericIndexes(table, "books_genre"));


export const booksAuthors = snakeCase.table("books_authors", {
    ...commonGenericCols(books.id),
}, (table) => commonGenericIndexes(table, "books_authors"));


export const booksTags = snakeCase.table("books_tags", {
    ...commonMediaTagsCols(books.id),
}, (table) => commonMediaTagsIndexes(table, MediaType.BOOKS));


export const booksRelations = relations(books, ({ many }) => ({
    booksAuthors: many(booksAuthors),
    booksGenres: many(booksGenre),
    booksTags: many(booksTags),
    booksLists: many(booksList),
}));


export const booksListRelations = relations(booksList, ({ one }) => ({
    user: one(user, {
        fields: [booksList.userId],
        references: [user.id]
    }),
    book: one(books, {
        fields: [booksList.mediaId],
        references: [books.id]
    }),
}));


export const booksAuthorsRelations = relations(booksAuthors, ({ one }) => ({
    book: one(books, {
        fields: [booksAuthors.mediaId],
        references: [books.id]
    }),
}));


export const booksGenreRelations = relations(booksGenre, ({ one }) => ({
    book: one(books, {
        fields: [booksGenre.mediaId],
        references: [books.id]
    }),
}));


export const booksTagsRelations = relations(booksTags, ({ one }) => ({
    tag: one(mediaTags, {
        fields: [booksTags.tagId],
        references: [mediaTags.id],
    }),
    user: one(user, {
        fields: [booksTags.userId],
        references: [user.id]
    }),
    book: one(books, {
        fields: [booksTags.mediaId],
        references: [books.id]
    }),
}));
