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
import type { Game } from './Game';
import {
    GameFromJSON,
    GameFromJSONTyped,
    GameToJSON,
} from './Game';
import type { GameStep } from './GameStep';
import {
    GameStepFromJSON,
    GameStepFromJSONTyped,
    GameStepToJSON,
} from './GameStep';
import type { GameStatistic } from './GameStatistic';
import {
    GameStatisticFromJSON,
    GameStatisticFromJSONTyped,
    GameStatisticToJSON,
} from './GameStatistic';

/**
 * 
 * @export
 * @interface GetGameStatistics200Response
 */
export interface GetGameStatistics200Response {
    /**
     * 
     * @type {Array<GameStep>}
     * @memberof GetGameStatistics200Response
     */
    steps: Array<GameStep>;
    /**
     * 
     * @type {Game}
     * @memberof GetGameStatistics200Response
     */
    game: Game;
    /**
     * 
     * @type {GameStatistic}
     * @memberof GetGameStatistics200Response
     */
    statistics?: GameStatistic;
}

/**
 * Check if a given object implements the GetGameStatistics200Response interface.
 */
export function instanceOfGetGameStatistics200Response(value: object): value is GetGameStatistics200Response {
    if (!('steps' in value) || value['steps'] === undefined) return false;
    if (!('game' in value) || value['game'] === undefined) return false;
    return true;
}

export function GetGameStatistics200ResponseFromJSON(json: any): GetGameStatistics200Response {
    return GetGameStatistics200ResponseFromJSONTyped(json, false);
}

export function GetGameStatistics200ResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GetGameStatistics200Response {
    if (json == null) {
        return json;
    }
    return {
        
        'steps': ((json['steps'] as Array<any>).map(GameStepFromJSON)),
        'game': GameFromJSON(json['game']),
        'statistics': json['statistics'] == null ? undefined : GameStatisticFromJSON(json['statistics']),
    };
}

export function GetGameStatistics200ResponseToJSON(value?: GetGameStatistics200Response | null): any {
    if (value == null) {
        return value;
    }
    return {
        
        'steps': ((value['steps'] as Array<any>).map(GameStepToJSON)),
        'game': GameToJSON(value['game']),
        'statistics': GameStatisticToJSON(value['statistics']),
    };
}

