#ifndef QTEVENT_H
#define QTEVENT_H

#include <QObject>
#include <QHash>

class QtEvent : public QObject
{
    Q_OBJECT
public:
    explicit QtEvent(QObject *parent = 0);
    void setEvent(int type, void *func);
    virtual bool eventFilter(QObject *target, QEvent *event);
protected:
    QHash<int,void*> m_hash;
};

#endif // QTEVENT_H
