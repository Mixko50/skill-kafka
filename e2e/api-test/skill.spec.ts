import {expect, test} from '@playwright/test'
import {closeDatabase, deleteData, insertSampleData, pingDatabase} from '../database/connection';

test.beforeAll(pingDatabase);
test.beforeAll(insertSampleData);

test.afterAll(deleteData);
test.afterAll(closeDatabase);

test.describe('GET /skills/e2e_playwright', () => {
    test('should response skill with status success', async ({request,}) => {
        const res = await request.get(`/api/v1/skills/e2e_playwright`)
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": expect.objectContaining({
                    "key": "e2e_playwright",
                    "name": "E2E Playwright",
                    "description": "Playwright is a Node.js library to automate the Chromium, WebKit, and Firefox browsers with a single API.",
                    "logo": "https://playwright.dev/img/playwright-logo.svg",
                    "tags": expect.arrayContaining([
                        "node",
                        "javascript",
                        "typescript",
                        "automation",
                        "testing"
                    ])
                })
            }))
    })
})

test.describe('GET /skills', () => {
    test('should response all skills with status success', async ({request,}) => {
        const res = await request.get(`/api/v1/skills`)
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": expect.arrayContaining([expect.objectContaining({
                    "key": "e2e_playwright",
                    "name": "E2E Playwright",
                    "description": "Playwright is a Node.js library to automate the Chromium, WebKit, and Firefox browsers with a single API.",
                    "logo": "https://playwright.dev/img/playwright-logo.svg",
                    "tags": expect.arrayContaining([
                        "node",
                        "javascript",
                        "typescript",
                        "automation",
                        "testing"
                    ])
                })])
            }))
    })
})