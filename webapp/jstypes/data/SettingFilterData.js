
	import {SettingFilter, SettingFilterFilter} from '../apiModel';

let settingFilterData = function () {
    return {
        settingFilterFilter: new SettingFilterFilter(),
        panel: {
            type: '',
            show: false,
            create: 'create',
            edit: 'edit',
            request: 'request'
        },
        panelHeaderCreate: 'Создать',
        panelHeaderEdit: 'Изменить',
        panelSubmitButtonTextCreate: 'Создать',
        panelSubmitButtonTextEdit: 'Изменить',
        currentSettingFilterItem: {
            item: new SettingFilter(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default settingFilterData;

