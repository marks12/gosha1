
	import {Authenticator, AuthenticatorFilter} from '../apiModel';

let authenticatorData = function () {
    return {
        authenticatorFilter: new AuthenticatorFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentAuthenticatorItem: {
            item: new Authenticator(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default authenticatorData;

