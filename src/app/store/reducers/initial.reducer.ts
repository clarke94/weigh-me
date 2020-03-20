import { createReducer, on } from '@ngrx/store';
import { active, inactive } from '../actions/initial.actions';

export const initialState = false;

const _initialReducer = createReducer(initialState,
    on(active, state => state = true),
    on(inactive, state => state = false),
)

export function initialReducer(state, action) {
    return _initialReducer(state, action);
}
