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
  GetGame200Response,
} from '../models/index';
import {
    GetGame200ResponseFromJSON,
    GetGame200ResponseToJSON,
} from '../models/index';

export interface GetGameRequest {
    gameId: string;
}

/**
 * 
 */
export class GameApi extends runtime.BaseAPI {

    /**
     * Get game and steps given id
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

        const response = await this.request({
            path: `/game/{game_id}`.replace(`{${"game_id"}}`, encodeURIComponent(String(requestParameters['gameId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GetGame200ResponseFromJSON(jsonValue));
    }

    /**
     * Get game and steps given id
     */
    async getGame(requestParameters: GetGameRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GetGame200Response> {
        const response = await this.getGameRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
