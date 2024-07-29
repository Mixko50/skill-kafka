import { Client } from "pg";

export const client = new Client({
        connectionString: process.env.POSTGRES_URI
    }
)

export async function pingDatabase() {
    console.log('Postgres URI:', process.env.POSTGRES_URI)
    try {
        await client.connect()
        console.log('Connected to Postgres database')
        await clearDatabase()
    } catch (error) {
        console.error('Error connecting to database:', error)
    }
}

export async function clearDatabase() {
    try {
        await client.query("DELETE FROM skill")
    } catch (error) {
        console.error('Error connecting to database:', error)
    }
}

export async function insertSampleData() {
    try {
        const sampleData = {
            key: 'e2e_playwright',
            name: 'E2E Playwright',
            description: 'Playwright is a Node.js library to automate the Chromium, WebKit, and Firefox browsers with a single API.',
            logo: 'https://playwright.dev/img/playwright-logo.svg',
            tags: ['node', 'javascript', 'typescript', 'automation', 'testing']
        }
        const insertQuery = 'INSERT INTO skill (key, name, description, logo, tags) values ($1, $2, $3, $4, $5)'
        await client.query(insertQuery, [sampleData.key, sampleData.name, sampleData.description, sampleData.logo, sampleData.tags])
        console.log('Inserted sample data')
    } catch (error) {
        console.error('Error connecting to database:', error)
    }
}

export async function deleteData() {
    try {
        await client.query("DELETE FROM skill")
    } catch (error) {
        console.error('Error connecting to database:', error)
    }
}

export async function closeDatabase() {
    try {
        await client.end()
        console.log('Closed Postgres database connection')
    } catch (error) {
        console.error('Error closing database:', error)
    }
}