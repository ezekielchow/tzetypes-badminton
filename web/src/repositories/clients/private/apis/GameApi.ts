/* tslint:disable */
/* eslint-disable */
/**
 * tzetypes-badminton
 * tzetypes-badminton
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  AddGameSteps201Response,
  AddGameStepsRequestSchema,
  CreateOrUpdateGameHistoryRequestSchema,
  EndGameRequest,
  GameStartRequestSchema,
  GetGame200Response,
  GetGameHistory200Response,
  GetRecentStatistics200Response,
  ListActiveGames200Response,
  StartGame201Response,
} from '../models/index';
import {
    AddGameSteps201ResponseFromJSON,
    AddGameSteps201ResponseToJSON,
    AddGameStepsRequestSchemaFromJSON,
    AddGameStepsRequestSchemaToJSON,
    CreateOrUpdateGameHistoryRequestSchemaFromJSON,
    CreateOrUpdateGameHistoryRequestSchemaToJSON,
    EndGameRequestFromJSON,
    EndGameRequestToJSON,
    GameStartRequestSchemaFromJSON,
    GameStartRequestSchemaToJSON,
    GetGame200ResponseFromJSON,
    GetGame200ResponseToJSON,
    GetGameHistory200ResponseFromJSON,
    GetGameHistory200ResponseToJSON,
    GetRecentStatistics200ResponseFromJSON,
    GetRecentStatistics200ResponseToJSON,
    ListActiveGames200ResponseFromJSON,
    ListActiveGames200ResponseToJSON,
    StartGame201ResponseFromJSON,
    StartGame201ResponseToJSON,
} from '../models/index';

export interface AddGameStepsRequest {
    gameId: string;
    addGameStepsRequestSchema: AddGameStepsRequestSchema;
}

export interface CreateOrUpdateGameHistoryRequest {
    gameId: string;
    createOrUpdateGameHistoryRequestSchema: CreateOrUpdateGameHistoryRequestSchema;
}

export interface DeleteGameStepsRequest {
    gameId: string;
    requestBody: Array<string>;
}

export interface EndGameOperationRequest {
    gameId: string;
    endGameRequest?: EndGameRequest;
}

export interface GetGameRequest {
    gameId: string;
}

export interface GetGameHistoryRequest {
    gameId: string;
}

export interface StartGameRequest {
    gameStartRequestSchema: GameStartRequestSchema;
}

/**
 * 
 */
export class GameApi extends runtime.BaseAPI {

    /**
     */
    async addGameStepsRaw(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AddGameSteps201Response>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling addGameSteps().'
            );
        }

        if (requestParameters['addGameStepsRequestSchema'] == null) {
            throw new runtime.RequiredError(
                'addGameStepsRequestSchema',
                'Required parameter "addGameStepsRequestSchema" was null or undefined when calling addGameSteps().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}/steps`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: AddGameStepsRequestSchemaToJSON(requestParameters['addGameStepsRequestSchema']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AddGameSteps201ResponseFromJSON(jsonValue));
    }

    /**
     */
    async addGameSteps(requestParameters: AddGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AddGameSteps201Response> {
        const response = await this.addGameStepsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createOrUpdateGameHistoryRaw(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetGameHistory200Response>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling createOrUpdateGameHistory().'
            );
        }

        if (requestParameters['createOrUpdateGameHistoryRequestSchema'] == null) {
            throw new runtime.RequiredError(
                'createOrUpdateGameHistoryRequestSchema',
                'Required parameter "createOrUpdateGameHistoryRequestSchema" was null or undefined when calling createOrUpdateGameHistory().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}/history`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: CreateOrUpdateGameHistoryRequestSchemaToJSON(requestParameters['createOrUpdateGameHistoryRequestSchema']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetGameHistory200ResponseFromJSON(jsonValue));
    }

    /**
     */
    async createOrUpdateGameHistory(requestParameters: CreateOrUpdateGameHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetGameHistory200Response> {
        const response = await this.createOrUpdateGameHistoryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async deleteGameStepsRaw(requestParameters: DeleteGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling deleteGameSteps().'
            );
        }

        if (requestParameters['requestBody'] == null) {
            throw new runtime.RequiredError(
                'requestBody',
                'Required parameter "requestBody" was null or undefined when calling deleteGameSteps().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}/steps/delete`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: requestParameters['requestBody'],
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteGameSteps(requestParameters: DeleteGameStepsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteGameStepsRaw(requestParameters, initOverrides);
    }

    /**
     * End game by setting \'isEnded\' to \'true\'
     */
    async endGameRaw(requestParameters: EndGameOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling endGame().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}/end`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: EndGameRequestToJSON(requestParameters['endGameRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * End game by setting \'isEnded\' to \'true\'
     */
    async endGame(requestParameters: EndGameOperationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.endGameRaw(requestParameters, initOverrides);
    }

    /**
     * Get game
     */
    async getGameRaw(requestParameters: GetGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetGame200Response>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling getGame().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetGame200ResponseFromJSON(jsonValue));
    }

    /**
     * Get game
     */
    async getGame(requestParameters: GetGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetGame200Response> {
        const response = await this.getGameRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getGameHistoryRaw(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetGameHistory200Response>> {
        if (requestParameters['gameId'] == null) {
            throw new runtime.RequiredError(
                'gameId',
                'Required parameter "gameId" was null or undefined when calling getGameHistory().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/{game_id}/history`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetGameHistory200ResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGameHistory(requestParameters: GetGameHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetGameHistory200Response> {
        const response = await this.getGameHistoryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getRecentStatisticsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GetRecentStatistics200Response>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/recent-statistics`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetRecentStatistics200ResponseFromJSON(jsonValue));
    }

    /**
     */
    async getRecentStatistics(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetRecentStatistics200Response> {
        const response = await this.getRecentStatisticsRaw(initOverrides);
        return await response.value();
    }

    /**
     * List active games
     */
    async listActiveGamesRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ListActiveGames200Response>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game/active`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ListActiveGames200ResponseFromJSON(jsonValue));
    }

    /**
     * List active games
     */
    async listActiveGames(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ListActiveGames200Response> {
        const response = await this.listActiveGamesRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async startGameRaw(requestParameters: StartGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<StartGame201Response>> {
        if (requestParameters['gameStartRequestSchema'] == null) {
            throw new runtime.RequiredError(
                'gameStartRequestSchema',
                'Required parameter "gameStartRequestSchema" was null or undefined when calling startGame().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("BearerAuth", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/game`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: GameStartRequestSchemaToJSON(requestParameters['gameStartRequestSchema']),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => StartGame201ResponseFromJSON(jsonValue));
    }

    /**
     */
    async startGame(requestParameters: StartGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<StartGame201Response> {
        const response = await this.startGameRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
