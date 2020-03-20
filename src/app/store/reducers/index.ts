import {
    ActionReducer,
    ActionReducerMap,
    createFeatureSelector,
    createSelector,
    MetaReducer
} from '@ngrx/store';
import { environment } from '../../../environments/environment';
import { initialReducer } from './initial.reducer';

export interface State {

}

export const reducers: ActionReducerMap<State> = {
    initial: initialReducer
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];
