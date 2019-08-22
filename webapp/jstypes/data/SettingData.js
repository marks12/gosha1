
	import {Setting, SettingFilter} from '../apiModel';

let settingData = function () {
    return {
        settingFilter: new SettingFilter(),
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
        currentSettingItem: {
            item: new Setting(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default settingData;

