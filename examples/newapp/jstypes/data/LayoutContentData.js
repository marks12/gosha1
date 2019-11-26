
	import {LayoutContent, LayoutContentFilter} from '../apiModel';

let layoutContentData = function () {
    return {
        layoutContentFilter: new LayoutContentFilter(),
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
        currentLayoutContentItem: {
            item: new LayoutContent(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default layoutContentData;

