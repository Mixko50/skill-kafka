import { expect, test } from '@playwright/test'
import {
    closeDatabase,
    deleteData,
    insertSampleData,
    pingDatabase, testDataKey,
} from '../database/connection';

test.beforeAll(pingDatabase);
test.beforeAll(insertSampleData);

test.afterAll(deleteData);
test.afterAll(closeDatabase);

test.describe('GET /skills/:key', () => {
    test('should response skill with status success', async ({request,}) => {
        const res = await request.get(`/api/v1/skills/` + testDataKey.insertSetupKey)
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "data": expect.objectContaining({
                    "key": testDataKey.insertSetupKey,
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
                    "key": testDataKey.insertSetupKey,
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

test.describe('POST /skills', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.post(`/api/v1/skills`, {
            data: {
                "key": testDataKey.createSkillKey,
                "name": "E2E Jest",
                "description": "Jest is a delightful JavaScript Testing Framework with a focus on simplicity.",
                "logo": "https://jestjs.io/img/jest.svg",
                "tags": ["node", "javascript", "typescript", "testing"]
            }
        })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "creating skill already in progress"
            })
        )
    })
})

test.describe('PUT /skills/:key', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.put(`/api/v1/skills/` + testDataKey.updateSkillKey, {
            data: {
                "name": "E2E Jest",
                "description": "Jest is a delightful JavaScript Testing Framework with a focus on simplicity.",
                "logo": "https://jestjs.io/img/jest.svg",
                "tags": ["node", "javascript", "typescript", "testing"]
            }
        })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "updating skill already in progress"
            })
        )
    })
})

test.describe('PATCH /skills/:key/actions/name', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.patch(`/api/v1/skills/` + testDataKey.updateNameKey + `/actions/name`, {
            data: {
                "name": testDataKey.updateNameKey
            }
        })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "updating skill name already in progress"
            })
        )
    })
})

test.describe('PATCH /skills/:key/actions/description', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.patch(`/api/v1/skills/` + testDataKey.updateDescriptionKey + `/actions/description`, {
            data: {
                "description": testDataKey.updateDescriptionKey
            }
        })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "updating skill description already in progress"
            })
        )
    })
})

test.describe('PATCH /skills/:key/actions/logo', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.patch(`/api/v1/skills/` + testDataKey.updateLogoKey + `/actions/logo`, {
            data: {
                "logo": testDataKey.updateLogoKey
            }
        })
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "updating skill logo already in progress"
            })
        )
    })
})

test.describe('DELETE /skills/:key', () => {
    test('should response with status success', async ({request}) => {
        const res = await request.delete(`/api/v1/skills/` + testDataKey.deleteSkillKey)
        expect(res.ok()).toBeTruthy()
        expect(await res.json()).toEqual(
            expect.objectContaining({
                "status": "success",
                "message": "deleting skill already in progress"
            })
        )
    })
})