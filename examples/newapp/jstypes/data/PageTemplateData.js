
	import {PageTemplate, PageTemplateFilter} from '../apiModel';

let pageTemplateData = function () {
    return {
        pageTemplateFilter: new PageTemplateFilter(),
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
        currentPageTemplateItem: {
            item: new PageTemplate(),
            hasChange: false,
            isSelected: false,
            canDelete: false,
            showDeleteConfirmation: false,
        },
    }
}();

export default pageTemplateData;

