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

import { mapValues } from '../runtime';
import type { GameHistory } from './GameHistory';
import {
    GameHistoryFromJSON,
    GameHistoryFromJSONTyped,
    GameHistoryToJSON,
} from './GameHistory';

/**
 * 
 * @export
 * @interface GetGameHistory200Response
 */
export interface GetGameHistory200Response {
    /**
     * 
     * @type {GameHistory}
     * @memberof GetGameHistory200Response
     */
    gameHistory: GameHistory;
}

/**
 * Check if a given object implements the GetGameHistory200Response interface.
 */
export function instanceOfGetGameHistory200Response(value: object): value is GetGameHistory200Response {
    if (!('gameHistory' in value) || value['gameHistory'] === undefined) return false;
    return true;
}

export function GetGameHistory200ResponseFromJSON(json: any): GetGameHistory200Response {
    return GetGameHistory200ResponseFromJSONTyped(json, false);
}

export function GetGameHistory200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetGameHistory200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'gameHistory': GameHistoryFromJSON(json['game_history']),
    };
}

export function GetGameHistory200ResponseToJSON(value?: GetGameHistory200Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'game_history': GameHistoryToJSON(value['gameHistory']),
    };
}

