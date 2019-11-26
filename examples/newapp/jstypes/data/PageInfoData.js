
	import {PageInfo, PageInfoFilter} from '../apiModel';

let pageInfoData = function () {
    return {
        pageInfoFilter: new PageInfoFilter(),
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
        currentPageInfoItem: {
            item: new PageInfo(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageInfoData;

